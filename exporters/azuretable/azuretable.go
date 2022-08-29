package azuretable

import (
	//"context"
	//"fmt"
	//"os"

	//"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func ConnectionStringSetup(s string) (serviceClient *aztables.ServiceClient) {
	serviceClient, err := aztables.NewServiceClientFromConnectionString(s, nil)
	if err != nil {
		panic(err)
	}
	return
}

func NewEntity(serviceClient *aztables.ServiceClient) (resp aztables.AddEntityResponse) {
	myEntity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: "001234",
			RowKey:       "RedMarker",
		},
		Properties: map[string]interface{}{
			"Stock":        15,
			"Price":        9.99,
			"Comments":     "great product",
			"OnSale":       true,
			"ReducedPrice": 7.99,
			"PurchaseDate": aztables.EDMDateTime(time.Date(2021, time.August, 21, 1, 1, 0, 0, time.UTC)),
			//"BinaryRepresentation": aztables.EDMBinary([]byte{"Bytesliceinfo"}),
		},
	}
	marshalled, err := json.Marshal(myEntity)
	log.Printf("%v", err)

	client := serviceClient.NewClient("myTable")

	resp, err = client.AddEntity(context.Background(), marshalled, nil)
	log.Printf("%v", err)
	return
}

func Query(serviceClient *aztables.ServiceClient) {
    client := serviceClient.NewClient("myTable")


	filter := "PartitionKey eq 'markers' or RowKey eq 'id-001'"
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Select: to.Ptr("RowKey,Value,Product,Available"),
		Top:    to.Ptr(int32((15))),
	}

	pager := client.NewListEntitiesPager(options)

    pager.NextPage(context.Background())
	// for pager.NextPage() {
	// 	resp := pager.PageResponse()
	// 	fmt.Printf("Received: %v entities\n", len(resp.Entities))

	// 	for _, entity := range resp.Entities {
	// 		var myEntity aztables.EDMEntity
	// 		err = json.Unmarshal(entity, &myEntity)
	// 		log.Printf("%v", err)

	// 		fmt.Printf("Received: %v, %v, %v, %v\n", myEntity.Properties["RowKey"], myEntity.Properties["Value"], myEntity.Properties["Product"], myEntity.Properties["Available"])
	// 	}
	// }

	// if pager.Err() != nil {
	// 	// handle error...
	// }
}
