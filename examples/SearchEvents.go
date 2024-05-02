package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/davudpasha"
)

func main() {
	body := davudpasha.EventsListRequest{
		Reason: common.PtrString("query"),
		Query: &davudpasha.EventsQueryFilter{
			QuerySQL: common.PtrString("sourcetype=\"paloalto\" eql select * from _source_ where _condition_ limit 1"),
			DateTimeRange: &davudpasha.EventsDateTimeRange{
				DateTimeType: davudpasha.EVENTSDATETIMETYPE_LAST10MINUTES,
				StartDate:    common.PtrString(""),
				EndDate:      common.PtrString(""),
			},
			Size: common.PtrInt64(10),
			QueryOptions: &davudpasha.EventsQueryOptions{
				ShowHighlight: common.PtrBool(true),
			},
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
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
