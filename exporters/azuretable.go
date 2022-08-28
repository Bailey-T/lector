package azuretable

import (
    "context"
    "fmt"
    "os"

    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func AzureTableSetup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	accountName, ok := os.LookupEnv("TABLES_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("TABLES_STORAGE_ACCOUNT_NAME could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.table.core.windows.net", accountName)

	service, err := aztables.NewServiceClient(serviceURL, cred, nil)
	if err != nil {
		panic(err)
	}

	// Create a table
	_, err = service.CreateTable(context.TODO(), "fromServiceClient", nil)
	if err != nil {
		panic(err)
	}
}