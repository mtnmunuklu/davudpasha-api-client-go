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
	api := davudpasha.NewQueriesApi(apiClient)

	searcBody := davudpasha.QueriesSearchRequest{
		Filter:                  common.PtrString("test-query"),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	searchResp, r, err := api.SearchQueries(ctx, *davudpasha.NewSearchQueriesOptionalParameters().WithBody(searcBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.SearchQueries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	saveBody := davudpasha.QueriesSaveRequest{
		QuerySettings: &davudpasha.QueriesItem{
			ID:   searchResp.Items[0].ID,
			Name: common.PtrString("edited-test-query"),
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	saveResp, r, err := api.SaveQueries(ctx, *davudpasha.NewSaveQueriesOptionalParameters().WithBody(saveBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.SaveQueries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(saveResp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `QueriesApi.SaveQueries`:\n%s\n", responseContent)

}
