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
	body := davudpasha.GeoLocationsSearchRequest{
		Filter: common.PtrString("test-geolocation"),
		Query: &davudpasha.QueryFilter{
			QuerySQL:        common.PtrString(""),
			SearchAfterKeys: []string{""},
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}
	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewGeoLocationsApi(apiClient)
	resp, r, err := api.SearchGeoLocations(ctx, *davudpasha.NewSearchGeoLocationsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoLocationsApi.SearchGeoLocations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `GeoLocationsApi.SearchGeoLocations`:\n%s\n", responseContent)

}
