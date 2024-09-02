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
	body := davudpasha.SourcesSaveRequest{
		Lgs: &davudpasha.SourcesItem{
			Enabled:          common.PtrBool(true),
			Name:             common.PtrString("test-file-source"),
			LogSourceDefCode: common.PtrString("test-file-lgs-defcode"),
			LogReaderType:    davudpasha.LOGREADERTYPE_MSGRAPH,
			AlertTimeout:     common.PtrInt64(120),
			LogReaderData: &davudpasha.SourcesLogReaderData{
				ApiUrl:       common.PtrString("/auditLogs/directoryAudits"),
				IdFieldType:  davudpasha.IDFIELDTYPE_DATETIME.Ptr(),
				IdFieldName:  common.PtrString("activityDateTime"),
				ReturnType:   davudpasha.RETURNTYPE_JSON.Ptr(),
				TenantId:     common.PtrString("04111867-3f4b-48d5"),
				ClientId:     common.PtrString("2d5bc44f-61a1-4c89"),
				ClientSecret: common.PtrString("EFI8Q~QSWI"),
				// Use CredentialApi for search credential id.
				Credential: common.PtrString("6YRSGRMA"),
			},
			DiscardedLogsConfig: davudpasha.DISCARDEDLOGSCONFIG_TODETAILS,
			WriteRawLogs:        common.PtrBool(true),
			UseSecondaryWriter:  common.PtrBool(true),
		},
	}

	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewSourcesApi(apiClient)
	resp, r, err := api.SaveSources(ctx, *davudpasha.NewSaveSourcesOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.SaveSources`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `SourcesApi.SaveSources`:\n%s\n", responseContent)

}
