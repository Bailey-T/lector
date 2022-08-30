package azuretable

import (
	//"context"
	//"fmt"
	//"os"

	//"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"context"
	"encoding/json"
	"fmt"
	"log"
	//"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"

	"github.com/drtbz/lector/sources"
)

func ConnectionStringSetup(s string) (serviceClient *aztables.ServiceClient) {
	serviceClient, err := aztables.NewServiceClientFromConnectionString(s, nil)
	if err != nil {
		panic(err)
	}
	return
}

func NewGHEntity(client *aztables.Client, v sources.Upstream, r string) (resp aztables.AddEntityResponse) {
	myEntity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: v.Owner(),
			RowKey:       v.Repo(),
		},
		Properties: map[string]interface{}{
			"Source":  v.Source(),
			"Release": r,
		},
	}
	marshalled, err := json.Marshal(myEntity)
	log.Printf("%v", err)

	resp, err = client.AddEntity(context.Background(), marshalled, nil)
	log.Printf("%v", err)
	return
}

func NewAHEntity(client *aztables.Client, v sources.Upstream, chart, app string) (resp aztables.AddEntityResponse) {
	myEntity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: v.Owner(),
			RowKey:       v.Repo(),
		},
		Properties: map[string]interface{}{
			"Source": v.Source(),
			"Chart":  chart,
			"App":    app,
		},
	}
	marshalled, err := json.Marshal(myEntity)
	log.Printf("%v", err)

	resp, err = client.AddEntity(context.Background(), marshalled, nil)
	log.Printf("%v", err)
	return
}

func Query(client *aztables.Client, v sources.Upstream) (err error) {

	filter := fmt.Sprintf("PartitionKey eq '%v' and RowKey eq '%v'", v.Owner(), v.Repo())
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Select: to.Ptr("Source,Release,Chart,App"),
		Top:    to.Ptr(int32((15))),
	}

	pager := client.NewListEntitiesPager(options)
	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			return err
		}
		pager.NextPage(context.Background())
		// for pager.NextPage() {
		// 	resp := pager.PageResponse()
		// 	fmt.Printf("Received: %v entities\n", len(resp.Entities))

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
                return err
            }
			log.Printf("Received: %v, %v, %v, %v \n %v", myEntity.Properties["Source"], myEntity.Properties["Release"], myEntity.Properties["Chart"], myEntity.Properties["App"], myEntity.Properties)
		}
	}

	// if pager.Err() != nil {
	// 	// handle error...
	// }
    return nil
}

func Setup(tableName string, client *aztables.ServiceClient) (err error) {
	myTable := "lector"
	filter := fmt.Sprintf("TableName ge '%v'", myTable)
	pager := client.NewListTablesPager(&aztables.ListTablesOptions{Filter: &filter})

	pageCount := 1
	pageFound := false
	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			return err
		}
		log.Printf("## Azure Tables Setup:")
		log.Printf("Found %d tables in page #%d\n", len(response.Tables), pageCount)
		for _, table := range response.Tables {
			log.Printf("Found TableName: %s", *table.Name)
			log.Printf("Skipping Table Creation")
			pageFound = true
		}
		pageCount += 1
	}
	if !pageFound {
		resp, err := client.CreateTable(context.TODO(), myTable, nil)
		log.Printf("Creating Table: %s", *resp.TableName)
		if err != nil {
			return err
		}
	}
	return nil
}
