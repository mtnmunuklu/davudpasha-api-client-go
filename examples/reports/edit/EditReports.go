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
	api := davudpasha.NewReportsApi(apiClient)

	searchReportsBody := davudpasha.ReportsSearchRequest{
		SearchFilter:            common.PtrString("test-report"),
		ApplicationName:         common.PtrString("csiem"),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	searchReportsResp, r, err := api.SearchReports(ctx, *davudpasha.NewSearchReportsOptionalParameters().WithBody(searchReportsBody))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.SearchReports`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	saveReportsBody := davudpasha.ReportsSaveRequest{
		Application: common.PtrString("csiem"),
		Report: &davudpasha.ReportsItem{
			ReportID: searchReportsResp[0].ReportID,
			Name:     common.PtrString("updated-test-report"),
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	saveReportsResp, r, err := api.SaveReports(ctx, *davudpasha.NewSaveReportsOptionalParameters().WithBody(saveReportsBody))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.SaveReports`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(saveReportsResp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `ReportsApi.SaveReports`:\n%s\n", responseContent)

}
