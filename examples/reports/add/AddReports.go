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

	queriesApi := davudpasha.NewQueriesApi(apiClient)
	searchQueriesBody := davudpasha.QueriesSearchRequest{
		Filter:                  common.PtrString("test-query"),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	searchQueriesResp, r, err := queriesApi.SearchQueries(ctx, *davudpasha.NewSearchQueriesOptionalParameters().WithBody(searchQueriesBody))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.SearchQueries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	saveReportsBody := davudpasha.ReportsSaveRequest{
		Application: common.PtrString("csiem"),
		Report: &davudpasha.ReportsItem{
			IsActive:            common.PtrBool(true),
			ModuleFilter:        common.PtrString("Interface IDPReportAct"),
			RemoteInterfaceName: common.PtrString("IDPReportAct"),
			RemoteMethodName:    common.PtrString("RunScheduleReport"),
			Name:                common.PtrString("test-report"),
			Description:         *common.NewNullableString(common.PtrString("tes-report-description")),
			ReportQuery: []davudpasha.ReportsQuery{
				{
					Name: searchQueriesResp.Items[0].Name,
					Data: &davudpasha.ReportsQueryData{
						ItemType:      davudpasha.ITEMTYPE_QUERY,
						QueryID:       *common.NewNullableString(searchQueriesResp.Items[0].ID),
						DateTimeRange: searchQueriesResp.Items[0].DateTimeRange,
					},
					ExtData: *davudpasha.NewNullableReportsQueryExtData(&davudpasha.ReportsQueryExtData{
						ReportType: davudpasha.REPORTTYPE_QUERY,
					}),
					ShowTable: common.PtrBool(true),
					TableVisualization: &davudpasha.ReportsTableVisualization{
						ChartType:   common.PtrString("Table"),
						MaxRowCount: common.PtrInt64(10000),
					},
				},
			},
			ReportData: &davudpasha.ReportsData{
				ReportType: common.PtrString("csv"),
				Page: &davudpasha.ReportsPage{
					IsA3:        common.PtrBool(false),
					IsLandscape: common.PtrBool(false),
				},
				AddCoverPage: common.PtrBool(false),
				MaxRowCount:  common.PtrInt64(10000),
				Language:     common.PtrString("en"),
			},
			Version:                 common.PtrFloat64(1),
			SubExecutorModuleFilter: common.PtrString("Interface ICsiemReportAct"),
			Username:                common.PtrString("admin"),
		},
		Schedule: &davudpasha.ScheduleConfig{
			ScheduleType: common.PtrString("none"),
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}

	reportsApi := davudpasha.NewReportsApi(apiClient)
	saveReportsResp, r, err := reportsApi.SaveReports(ctx, *davudpasha.NewSaveReportsOptionalParameters().WithBody(saveReportsBody))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.SaveReports`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(saveReportsResp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `ReportsApi.SaveReports`:\n%s\n", responseContent)

}
