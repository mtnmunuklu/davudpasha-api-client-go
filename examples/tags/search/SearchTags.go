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
	body := davudpasha.TagsSearchRequest{
		Category:                davudpasha.TAGCATEGORY_ALL,
		SearchFilter:            common.PtrString("test-tag"),
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewTagsApi(apiClient)
	resp, r, err := api.SearchTags(ctx, *davudpasha.NewSearchTagsOptionalParameters().WithBody(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.SearchTags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `TagsApi.SearchTags`:\n%s\n", responseContent)

}
