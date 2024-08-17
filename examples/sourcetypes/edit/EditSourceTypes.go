package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/davudpasha"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewSourceTypesApi(apiClient)

	searchSourceTypesBody := davudpasha.SourceTypesSearchRequest{
		Filter:                  common.PtrString("test-source-type"),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	searchSourceTypesResp, r, err := api.SearchSourceTypes(ctx, *davudpasha.NewSearchSourceTypesOptionalParameters().WithBody(searchSourceTypesBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTypesApi.SearchSourceTypes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	saveSourceTypesBody := davudpasha.SourceTypesSaveRequest{
		LgsDef: &davudpasha.SourceTypesItem{
			Name:    searchSourceTypesResp.Items[0].Name,
			CatCode: common.PtrString("updated-catcode"),
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTypesApi.SaveSourceTypes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(saveSourceTypesBody, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `SourceTypesApi.SaveSourceTypes`:\n%s\n", responseContent)

}
