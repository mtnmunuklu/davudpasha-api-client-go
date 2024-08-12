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

// ReportsApi service type.
type ReportsApi common.Service

// SearchReportsOptionalParameters holds optional parameters for SearchQueries.
type SearchReportsOptionalParameters struct {
	Body *ReportsSearchRequest
}

// NewSearchReportsOptionalParameters creates an empty struct for parameters.
func NewSearchReportsOptionalParameters() *SearchReportsOptionalParameters {
	this := SearchReportsOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SearchReportsOptionalParameters) WithBody(body ReportsSearchRequest) *SearchReportsOptionalParameters {
	r.Body = &body
	return r
}

// SaveReportsOptionalParameters holds optional parameters for SaveQueries.
type SaveReportsOptionalParameters struct {
	Body *ReportsSaveRequest
}

// NewSaveReportsOptionalParameters creates an empty struct for parameters.
func NewSaveReportsOptionalParameters() *SaveReportsOptionalParameters {
	this := SaveReportsOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SaveReportsOptionalParameters) WithBody(body ReportsSaveRequest) *SaveReportsOptionalParameters {
	r.Body = &body
	return r
}

// SearchReports searches for Reports.
//
// @Summary Search Reports
// @Description Search for Reports based on a filter.
// @Tags Reports
// @Accept  json
// @Produce  json
// @Param body body ReportsSearchRequest true "Reports Search Request"
// @Success 200 {array} ReportsItem "Successful operation"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 429 {object} ErrorResponse "Too Many Requests"
// @Router /IDPReportScheduleAct/GetListWithSchedule [post]
// @Security ApiKeyAuth
func (a *ReportsApi) SearchReports(ctx _context.Context, o ...SearchReportsOptionalParameters) ([]ReportsItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   []ReportsItem
		optionalParams        SearchReportsOptionalParameters
		localVarInterfaceCode = "IDPReportScheduleAct"
		localVarMethodName    = "GetListWithSchedule"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type SearchReportsOptionalParameters is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "ReportsApi.SearchReports")
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

// SaveReports saves Reports.
//
// @Summary Save Reports
// @Description Save Reports based on the provided data.
// @Tags Reports
// @Accept  json
// @Produce  json
// @Param body body ReportsSaveRequest true "Reports Save Request"
// @Success 200 {object} ReportsSaveResponse "Successful operation"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 429 {object} ErrorResponse "Too Many Requests"
// @Router /IDPReportScheduleAct/Save [post]
// @Security ApiKeyAuth
func (a *ReportsApi) SaveReports(ctx _context.Context, o ...SaveReportsOptionalParameters) (ReportsSaveResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   ReportsSaveResponse
		optionalParams        SaveReportsOptionalParameters
		localVarInterfaceCode = "IDPReportScheduleAct"
		localVarMethodName    = "Save"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type SaveReportsOptionalParameters is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "ReportsApi.SaveReports")
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

// NewReportsApi returns ReportsApi.
func NewReportsApi(client *common.APIClient) *ReportsApi {
	return &ReportsApi{
		Client: client,
	}
}
