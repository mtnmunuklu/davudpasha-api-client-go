package davudpasha

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// @title Davudpasha API
// @version 1.0.0
// @description Davudpasha API to demonstrate OpenAPI documentation for client-go
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @basePath /api
// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization

// SearchArchivesApi service type.
type SearchArchivesApi common.Service

// GetSearchArchivesOptionalParameters holds optional parameters for GetSearchArchives.
type GetSearchArchivesOptionalParameters struct {
	Body *SearchArchivesGetRequest
}

// NewGetSearchArchivesOptionalParameters creates an empty struct for parameters.
func NewGetSearchArchivesOptionalParameters() *GetSearchArchivesOptionalParameters {
	this := GetSearchArchivesOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *GetSearchArchivesOptionalParameters) WithBody(body SearchArchivesGetRequest) *GetSearchArchivesOptionalParameters {
	r.Body = &body
	return r
}

// GetSearchArchives retrieves search archives.
//
// @Summary Retrieve Search Archives
// @Description Retrieve search archives based on a filter.
// @Tags Search Archives
// @Accept  json
// @Produce  json
// @Param body body SearchArchivesGetRequest true "Search Archives Get Request"
// @Success 200 {object} SearchArchivesGetResponse "Successful operation"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 429 {object} ErrorResponse "Too Many Requests"
// @Router /ICSiemManagerSearchArchiveAct/GetSearchArchivesAndStatus [post]
// @Security ApiKeyAuth
func (a *SearchArchivesApi) GetSearchArchives(ctx _context.Context, o ...GetSearchArchivesOptionalParameters) (SearchArchivesGetResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   SearchArchivesGetResponse
		optionalParams        GetSearchArchivesOptionalParameters
		localVarInterfaceCode = "ICSiemManagerSearchArchiveAct"
		localVarMethodName    = "GetSearchArchivesAndStatus"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetSearchArchivesOptionalParameters	 is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "SearchArchivesApi.GetSearchArchives")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/" + localVarInterfaceCode + "/" + localVarMethodName

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	if optionalParams.Body != nil {
		localVarPostBody = &optionalParams.Body
	}

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "x-api-key"},
	)

	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewSearchArchivesApi returns SearchArchivesApi.
func NewSearchArchivesApi(client *common.APIClient) *SearchArchivesApi {
	return &SearchArchivesApi{
		Client: client,
	}
}
