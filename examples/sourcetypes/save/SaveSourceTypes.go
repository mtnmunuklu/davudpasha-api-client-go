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

	body := davudpasha.SourceTypesSaveRequest{
		LgsDef: &davudpasha.SourceTypesItem{
			Name:             common.PtrString("test-source-type"),
			DefCode:          common.PtrString("test-defcode"),
			Version:          common.PtrFloat64(0.1),
			CatCode:          common.PtrString("tes-catcode"),
			SourceReaderType: davudpasha.READERTYPE_FILE,
			Expressions: []davudpasha.SourceTypesExpression{
				{
					LogParserType: davudpasha.PARSERTYPE_DELIMITER,
					LogParserData: &davudpasha.SourceTypesLogParserData{
						Delimiter: common.PtrString("\\s"),
					},
					Mapping: []davudpasha.SourceTypesMapping{
						{
							Field: common.PtrString("column1"),
						},
						{
							Field: common.PtrString("column2"),
						},
					},
				},
			},
			ExpressionFields: []davudpasha.SourceTypesExpressionField{
				{
					Name:  common.PtrString("column1"),
					Used:  common.PtrBool(true),
					Index: common.PtrBool(true),
					Type:  common.PtrString("string"),
				},
				{
					Name:  common.PtrString("column2"),
					Used:  common.PtrBool(true),
					Index: common.PtrBool(true),
					Type:  common.PtrString("string"),
				},
			},
		},
		SmartRestRequestContext: common.PtrString("-<SmartRestRequestContext>-"),
	}

	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewSourceTypesApi(apiClient)
	resp, r, err := api.SaveSourceTypes(ctx, *davudpasha.NewSaveSourceTypesOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTypesApi.SaveSourceTypes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `SourceTypesApi.SaveSourceTypes`:\n%s\n", responseContent)

}
