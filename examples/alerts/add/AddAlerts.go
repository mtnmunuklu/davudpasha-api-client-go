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

	searchQueriesBody := davudpasha.QueriesSearchRequest{
		Filter:                  common.PtrString("test-query"),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	queriesApi := davudpasha.NewQueriesApi(apiClient)
	searchQueriesResp, r, err := queriesApi.SearchQueries(ctx, *davudpasha.NewSearchQueriesOptionalParameters().WithBody(searchQueriesBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.SearchQueries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	saveAlertDefinitionsBody := davudpasha.AlertDefinitionsSaveRequest{
		Correlation: &davudpasha.AlertDefinitionsCorrelationData{
			Name:          common.PtrString("test-alert"),
			Description:   *common.NewNullableString(common.PtrString("tes-alert")),
			Tags:          *common.NewNullableList(&[]string{"tag1", "tag2"}),
			RiskLevel:     common.PtrInt64(3),
			MaxAlertCount: common.PtrInt64(4),
			MaxSendCount:  common.PtrInt64(10),
			Data: &davudpasha.AlertDefinitionsQueryData{
				TimeFrameValue:            common.PtrInt64(5),
				TimeFrameType:             davudpasha.TIMEFRAMETYPE_MINUTES,
				RuleType:                  davudpasha.RULETYPE_ANY,
				QueryCorrelationAlertType: davudpasha.QUERYCORRELATIONALERTTYPE_WHENONEORMOREROW,
				QueryID:                   searchQueriesResp.Items[0].ID,
				Query:                     searchQueriesResp.Items[0].Query,
			},
			CorrelationType: davudpasha.CORRELATIONTYPE_INTERFACEIQUERYCORRELATION,
			Message:         common.PtrString("test-alert"),
			Version:         common.PtrFloat64(1.0),
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	alertsApi := davudpasha.NewAlertDefinitionsApi(apiClient)
	alertsResp, r, err := alertsApi.SaveAlertDefinitions(ctx, *davudpasha.NewSaveAlertDefinitionsOptionalParameters().WithBody(saveAlertDefinitionsBody))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionsApi.SaveAlertDefinitions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(alertsResp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionsApi.SaveAlertDefinitions`:\n%s\n", responseContent)

}
