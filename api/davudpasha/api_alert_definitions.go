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
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host     localhost:5000
// @basePath /api
// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization

// AlertDefinitionsApi service type.
type AlertDefinitionsApi common.Service

// SearchAlertDefinitionsOptionalParamters holds optional parameters for SearchAlertDefinitions.
type SearchAlertDefinitionsOptionalParameters struct {
	Body *AlertDefinitionsSearchRequest
}

// NewSearchAlertDefinitionsOptionalParameters creates an empty struct for parameters.
func NewSearchAlertDefinitionsOptionalParameters() *SearchAlertDefinitionsOptionalParameters {
	this := SearchAlertDefinitionsOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SearchAlertDefinitionsOptionalParameters) WithBody(body AlertDefinitionsSearchRequest) *SearchAlertDefinitionsOptionalParameters {
	r.Body = &body
	return r
}

// SaveAlertDefinitionsOptionalParamters holds optional parameters for SaveAlertDefinitions.
type SaveAlertDefinitionsOptionalParameters struct {
	Body *AlertDefinitionsSaveRequest
}

// NewSaveAlertDefinitionsOptionalParameters creates an empty struct for parameters.
func NewSaveAlertDefinitionsOptionalParameters() *SaveAlertDefinitionsOptionalParameters {
	this := SaveAlertDefinitionsOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SaveAlertDefinitionsOptionalParameters) WithBody(body AlertDefinitionsSaveRequest) *SaveAlertDefinitionsOptionalParameters {
	r.Body = &body
	return r
}

// SearchAlertDefinitions searches alerts.
//
// @Summary Search Alert Definitions
// @Description Search for alerts based on a query.
// @Tags Alert Definitions
// @Accept  json
// @Produce  json
// @Param body body AlertDefinitionsSearchRequest true "Alert Definitions Search Request"
// @Success 200 {object} AlertDefinitionsSearchResponse "Successful operation"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 429 {object} ErrorResponse "Too Many Requests"
// @Router /ICSiemManagerCorrelationAct/GetCorrelationList [post]
// @Security ApiKeyAuth
func (a *AlertDefinitionsApi) SearchAlertDefinitions(ctx _context.Context, o ...SearchAlertDefinitionsOptionalParameters) (AlertDefinitionsSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   AlertDefinitionsSearchResponse
		optionalParams        SearchAlertDefinitionsOptionalParameters
		localVarInterfaceCode = "ICSiemManagerCorrelationAct"
		localVarMethodName    = "GetCorrelationList"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type SearchAlertDefinitionsOptionalParameters is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "AlertDefinitionsApi.SearchAlertDefinitions")
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

// SaveAlertDefinitions saves AlertDefinitions.
//
// @Summary Save Alert Definitions
// @Description Save Alert Definitions based on the provided data.
// @Tags Alert Definitions
// @Accept  json
// @Produce  json
// @Param body body AlertDefinitionsSaveRequest true "Alert Definitions Save Request"
// @Success 200 {object} AlertDefinitionsSaveResponse "Successful operation"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 429 {object} ErrorResponse "Too Many Requests"
// @Router /ICSiemQueryAct/Save [post]
// @Security ApiKeyAuth
func (a *AlertDefinitionsApi) SaveAlertDefinitions(ctx _context.Context, o ...SaveAlertDefinitionsOptionalParameters) (AlertDefinitionsSaveResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   AlertDefinitionsSaveResponse
		optionalParams        SaveAlertDefinitionsOptionalParameters
		localVarInterfaceCode = "ICSiemManagerCorrelationAct"
		localVarMethodName    = "AddOrUpdateCorrelation"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type SaveAlertDefinitionsOptionalParameters is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "AlertDefinitionsApi.SaveAlertDefinitions")
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

// NewAlertDefinitionsApi returns AlertDefinitionsApi.
func NewAlertDefinitionsApi(client *common.APIClient) *AlertDefinitionsApi {
	return &AlertDefinitionsApi{
		Client: client,
	}
}
