package davudpasha

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SystemLogsApi service type.
type SystemLogsApi common.Service

// SearchSystemLogsOptionalParameters holds optional parameters for SearchSystemLogs.
type SearchSystemLogsOptionalParameters struct {
	Body *SystemLogsSearchRequest
}

// NewSearchSystemLogsOptionalParameters creates an empty struct for parameters.
func NewSearchSystemLogsOptionalParameters() *SearchSystemLogsOptionalParameters {
	this := SearchSystemLogsOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SearchSystemLogsOptionalParameters) WithBody(body SystemLogsSearchRequest) *SearchSystemLogsOptionalParameters {
	r.Body = &body
	return r
}

// SearchSystemLogs search SystemLogs.
// Returns SystemLogs that match an SystemLogs search query.
func (a *SystemLogsApi) SearchSystemLogs(ctx _context.Context, o ...SearchSystemLogsOptionalParameters) (SystemLogsSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   SystemLogsSearchResponse
		optionalParams        SearchSystemLogsOptionalParameters
		localVarInterfaceCode = "ISystemLogAct"
		localVarMethodName    = "Get"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type SearchSystemLogsOptionalParameters is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "SystemLogsApi.SearchSystemLogs")
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

// NewSystemLogsApi returns SystemLogsApi.
func NewSystemLogsApi(client *common.APIClient) *SystemLogsApi {
	return &SystemLogsApi{
		Client: client,
	}
}
