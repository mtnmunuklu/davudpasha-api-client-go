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
	body := davudpasha.AlertsSearchRequest{
		Filter:                  common.PtrString("test-alert"),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewAlertsApi(apiClient)
	resp, r, err := api.SearchAlerts(ctx, *davudpasha.NewSearchAlertsOptionalParameters().WithBody(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.SearchAlerts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `AlertsApi.SearchAlerts`:\n%s\n", responseContent)

}
