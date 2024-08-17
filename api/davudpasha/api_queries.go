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

// QueriesApi service type.
type QueriesApi common.Service

// SearchQueriesOptionalParameters holds optional parameters for SearchQueries.
type SearchQueriesOptionalParameters struct {
	Body *QueriesSearchRequest
}

// NewSearchQueriesOptionalParameters creates an empty struct for parameters.
func NewSearchQueriesOptionalParameters() *SearchQueriesOptionalParameters {
	this := SearchQueriesOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SearchQueriesOptionalParameters) WithBody(body QueriesSearchRequest) *SearchQueriesOptionalParameters {
	r.Body = &body
	return r
}

// SaveQueriesOptionalParameters holds optional parameters for SaveQueries.
type SaveQueriesOptionalParameters struct {
	Body *QueriesSaveRequest
}

// NewSaveQueriesOptionalParameters creates an empty struct for parameters.
func NewSaveQueriesOptionalParameters() *SaveQueriesOptionalParameters {
	this := SaveQueriesOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SaveQueriesOptionalParameters) WithBody(body QueriesSaveRequest) *SaveQueriesOptionalParameters {
	r.Body = &body
	return r
}

// SearchQueries searches for Queries.
//
// @Summary Search Queries
// @Description Search for Queries based on a filter.
// @Tags Queries
// @Accept  json
// @Produce  json
// @Param body body QueriesSearchRequest true "Queries Search Request"
// @Success 200 {object} QueriesSearchResponse "Successful operation"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 429 {object} ErrorResponse "Too Many Requests"
// @Router /ICSiemQueryAct/GetList [post]
// @Security ApiKeyAuth
func (a *QueriesApi) SearchQueries(ctx _context.Context, o ...SearchQueriesOptionalParameters) (QueriesSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   QueriesSearchResponse
		optionalParams        SearchQueriesOptionalParameters
		localVarInterfaceCode = "ICSiemQueryAct"
		localVarMethodName    = "GetList"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type SearchQueriesOptionalParameters is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "QueriesApi.SearchQueries")
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

// SaveQueries saves Queries.
//
// @Summary Save Queries
// @Description Save Queries based on the provided data.
// @Tags Queries
// @Accept  json
// @Produce  json
// @Param body body QueriesSaveRequest true "Queries Save Request"
// @Success 200 {object} QueriesSaveResponse "Successful operation"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 429 {object} ErrorResponse "Too Many Requests"
// @Router /ICSiemQueryAct/Save [post]
// @Security ApiKeyAuth
func (a *QueriesApi) SaveQueries(ctx _context.Context, o ...SaveQueriesOptionalParameters) (QueriesSaveResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   QueriesSaveResponse
		optionalParams        SaveQueriesOptionalParameters
		localVarInterfaceCode = "ICSiemQueryAct"
		localVarMethodName    = "Save"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type SaveQueriesOptionalParameters is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "QueriesApi.SaveQueries")
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

// NewEventsApi returns QueriesApi.
func NewQueriesApi(client *common.APIClient) *QueriesApi {
	return &QueriesApi{
		Client: client,
	}
}
