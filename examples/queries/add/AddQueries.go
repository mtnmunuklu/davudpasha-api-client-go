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
	body := davudpasha.QueriesSaveRequest{
		QuerySettings: &davudpasha.QueriesItem{
			Name:   common.PtrString("test-query"),
			Author: common.PtrString("admin"),
			DateTimeRange: &davudpasha.DateTimeRange{
				DateTimeType: davudpasha.DATETIMETYPE_LAST10MINUTES,
			},
			Description: *common.NewNullableString(common.PtrString("tes-description")),
			Query:       common.PtrString("sourcetype='windows-security' eql select * from _source_ where _condition_"),
			QueryType:   davudpasha.QUERYTYPE_NORMAL,
			Tags:        *common.NewNullableList(&[]string{"tag1", "tag2"}),
			KillChainPhase: *common.NewNullableList(&[]string{
				"Initial Access",
				"Execution",
			}),
			MitreTags: *common.NewNullableList(&[]davudpasha.MitreTag{
				{
					Tactic:        common.PtrString("Persistence"),
					Technique:     []string{"T1078", "T1087"},
					SubTechniques: []string{"T1078.001"},
				},
			}),
			Columns: *common.NewNullableList(&[]davudpasha.SelectedColumn{
				{
					Value:       common.PtrString("column1"),
					DisplayText: common.PtrString("column1"),
				},
				{
					Value:       common.PtrString("column2"),
					DisplayText: common.PtrString("column2"),
				},
			}),
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}

	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewQueriesApi(apiClient)
	resp, r, err := api.SaveQueries(ctx, *davudpasha.NewSaveQueriesOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.SaveQueries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `QueriesApi.SaveQueries`:\n%s\n", responseContent)

}
