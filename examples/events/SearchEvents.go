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
	body := davudpasha.EventsSearchRequest{
		Reason: common.PtrString("query"),
		Query: &davudpasha.EventsQueryFilter{
			QuerySQL: common.PtrString("sourcetype=\"paloalto\" eql select * from _source_ where _condition_ limit 1"),
			DateTimeRange: &davudpasha.EventsDateTimeRange{
				DateTimeType: davudpasha.EVENTSDATETIMETYPE_LAST10MINUTES,
			},
			QueryOptions: &davudpasha.EventsQueryOptions{
				ShowHighlight: common.PtrBool(true),
			},
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewEventsApi(apiClient)
	resp, r, err := api.SearchEvents(ctx, *davudpasha.NewSearchEventsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.SearchEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.SearchEvents`:\n%s\n", responseContent)
}
