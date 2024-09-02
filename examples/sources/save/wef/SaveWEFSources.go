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

	searchAssetsBody := davudpasha.AssetsSearchRequest{
		SearchFilter:            common.PtrString("order by IsManage desc, _id desc"),
		SearchFilter2:           common.PtrString("test-asset"),
		PageNumber:              common.PtrInt64(1),
		PageSize:                common.PtrInt64(10),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	assetsApi := davudpasha.NewAssetsApi(apiClient)
	searchAssetsResp, r, err := assetsApi.SearchAssets(ctx, *davudpasha.NewSearchAssetsOptionalParameters().WithBody(searchAssetsBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.SearchAssets`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	saveSourcesBody := davudpasha.SourcesSaveRequest{
		Lgs: &davudpasha.SourcesItem{
			Enabled:          common.PtrBool(true),
			Name:             common.PtrString("test-wef-source"),
			LogSourceDefCode: common.PtrString("test-wef-lgs-defcode"),
			LogReaderType:    davudpasha.LOGREADERTYPE_WEF,
			AlertTimeout:     common.PtrInt64(120),
			LogReaderData: &davudpasha.SourcesLogReaderData{
				MessageFieldType:                 *davudpasha.NewNullableSourcesMessageFieldType(davudpasha.MESSAGEFIELDTYPE_XML.Ptr()),
				TransformFrequencyInSeconds:      common.PtrInt64(60),
				DeleteArchivedLogFileAfterNHours: common.PtrInt64(30),
				DeleteOutputFileAfterNHours:      common.PtrInt64(30),
				MaxParallelProcessing:            common.PtrInt64(1),
				EventLogNameOrFileName:           common.PtrString("c:\\weflogs\\ForwardedEvents.evtx"),
			},
			DiscardedLogsConfig: davudpasha.DISCARDEDLOGSCONFIG_TODETAILS,
			WriteRawLogs:        common.PtrBool(true),
			UseSecondaryWriter:  common.PtrBool(true),
			IsAgentSource:       common.PtrBool(true),
			AgentIDs:            *common.NewNullableList(&[]string{*searchAssetsResp.Assets[0].AssetID}),
		},
	}
	sourcesApi := davudpasha.NewSourcesApi(apiClient)
	saveSourcesResp, r, err := sourcesApi.SaveSources(ctx, *davudpasha.NewSaveSourcesOptionalParameters().WithBody(saveSourcesBody))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.SaveSources`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(saveSourcesResp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `SourcesApi.SaveSources`:\n%s\n", responseContent)

}
