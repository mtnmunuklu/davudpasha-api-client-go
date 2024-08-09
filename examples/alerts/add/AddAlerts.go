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

	queriesBody := davudpasha.QueriesSearchRequest{
		Filter:                  common.PtrString("test-query"),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	queriesApi := davudpasha.NewQueriesApi(apiClient)
	queriesResp, r, err := queriesApi.SearchQueries(ctx, *davudpasha.NewSearchQueriesOptionalParameters().WithBody(queriesBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.SearchQueries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	alertsBody := davudpasha.AlertsSaveRequest{
		Correlation: &davudpasha.AlertsCorrelationData{
			Name:          common.PtrString("test-alert"),
			Description:   *common.NewNullableString(common.PtrString("tes-alert")),
			Tags:          *common.NewNullableList(&[]string{"tag1", "tag2"}),
			RiskLevel:     common.PtrInt64(3),
			MaxAlertCount: common.PtrInt64(4),
			Data: &davudpasha.AlertsQueryData{
				TimeFrameValue:            common.PtrInt64(5),
				TimeFrameType:             davudpasha.TIMEFRAMETYPE_MINUTES,
				RuleType:                  davudpasha.RULETYPE_ANY,
				QueryCorrelationAlertType: davudpasha.QUERYCORRELATIONALERTTYPE_WHENONEORMOREROW,
				QueryID:                   queriesResp.Items[0].ID,
				Query:                     queriesResp.Items[0].Query,
			},
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	alertsApi := davudpasha.NewAlertsApi(apiClient)
	alertsResp, r, err := alertsApi.SaveAlerts(ctx, *davudpasha.NewSaveAlertsOptionalParameters().WithBody(alertsBody))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.SaveAlerts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(alertsResp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `AlertsApi.SaveAlerts`:\n%s\n", responseContent)

}