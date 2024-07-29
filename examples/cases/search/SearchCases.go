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
	body := davudpasha.CasesSearchRequest{
		FromIndex: common.PtrInt64(0),
		PageSize:  common.PtrInt64(10),
		DateTimeRange: &davudpasha.DateTimeRange{
			DateTimeType: davudpasha.DATETIMETYPE_TODAY,
		},
		Filter:                  common.PtrString("test-case"),
		App:                     common.PtrString("csiem"),
		ShowSubCases:            common.PtrBool(false),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}

	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewCasesApi(apiClient)
	resp, r, err := api.SearchCases(ctx, *davudpasha.NewSearchCasesOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesApi.SearchCases`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `CasesApi.SearchCases`:\n%s\n", responseContent)

}
