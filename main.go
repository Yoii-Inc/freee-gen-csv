package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "github.com/taskooh/freee-get-sales/freeeclient"
)

func main() {
	id := int32(56)        // int32 | 受け付けID
	companyId := int32(56) // int32 | 事業所ID

	configuration := openapiclient.NewConfiguration()
	api_client := openapiclient.NewAPIClient(configuration)
	resp, r, err := api_client.JournalsApi.DownloadJournal(context.Background(), id).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.DownloadJournal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadJournal`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `JournalsApi.DownloadJournal`: %v\n", resp)
}
