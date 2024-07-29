package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// CasesSearchRequest represents the request structure for searching cases.
type CasesSearchRequest struct {
	// Starting index for pagination.
	FromIndex *int64 `json:"fromIndex,omitempty"`
	// Number of items per page for pagination.
	PageSize *int64 `json:"pageSize,omitempty"`
	// Date-time range filter for cases.
	DateTimeRange *DateTimeRange `json:"dateTimeRange,omitempty"`
	// Additional filter criteria for cases.
	Filter *string `json:"filter,omitempty"`
	// Application associated with the cases.
	App *string `json:"app,omitempty"`
	// Flag indicating whether to include sub-cases in the search results.
	ShowSubCases *bool `json:"showSubCases"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewCasesSearchRequest creates a new CasesSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCasesSearchRequest() *CasesSearchRequest {
	this := CasesSearchRequest{}
	return &this
}

// NewCasesSearchRequestWithDefaults creates a new CasesSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCasesSearchRequestWithDefaults() *CasesSearchRequest {
	this := CasesSearchRequest{}
	return &this
}

// GetFromIndex returns the FromIndex field value if set, zero value otherwise.
func (o *CasesSearchRequest) GetFromIndex() int64 {
	if o == nil || o.FromIndex == nil {
		var ret int64
		return ret
	}
	return *o.FromIndex
}

// GetFromIndexOk returns a tuple with the FromIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesSearchRequest) GetFromIndexOk() (*int64, bool) {
	if o == nil || o.FromIndex == nil {
		return nil, false
	}
	return o.FromIndex, true
}

// HasFromIndex returns a boolean if a field has been set.
func (o *CasesSearchRequest) HasFromIndex() bool {
	return o != nil && o.FromIndex != nil
}

// SetFromIndex gets a reference to the given int64 and assigns it to the FromIndex field.
func (o *CasesSearchRequest) SetFromIndex(v int64) {
	o.FromIndex = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CasesSearchRequest) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesSearchRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CasesSearchRequest) HasPageSize() bool {
	return o != nil && o.PageSize != nil
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *CasesSearchRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *CasesSearchRequest) GetDateTimeRange() DateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret DateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesSearchRequest) GetDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *CasesSearchRequest) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DateTimeRange field.
func (o *CasesSearchRequest) SetDateTimeRange(v DateTimeRange) {
	o.DateTimeRange = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CasesSearchRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesSearchRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CasesSearchRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *CasesSearchRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *CasesSearchRequest) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesSearchRequest) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *CasesSearchRequest) HasApp() bool {
	return o != nil && o.App != nil
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *CasesSearchRequest) SetApp(v string) {
	o.App = &v
}

// GetShowSubCases returns the ShowSubCases field value if set, zero value otherwise.
func (o *CasesSearchRequest) GetShowSubCases() bool {
	if o == nil || o.ShowSubCases == nil {
		var ret bool
		return ret
	}
	return *o.ShowSubCases
}

// GetShowSubCasesOk returns a tuple with the ShowSubCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesSearchRequest) GetShowSubCasesOk() (*bool, bool) {
	if o == nil || o.ShowSubCases == nil {
		return nil, false
	}
	return o.ShowSubCases, true
}

// HasShowSubCases returns a boolean if a field has been set.
func (o *CasesSearchRequest) HasShowSubCases() bool {
	return o != nil && o.ShowSubCases != nil
}

// SetShowSubCases gets a reference to the given bool and assigns it to the ShowSubCases field.
func (o *CasesSearchRequest) SetShowSubCases(v bool) {
	o.ShowSubCases = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *CasesSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *CasesSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *CasesSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CasesSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.FromIndex != nil {
		toSerialize["fromIndex"] = o.FromIndex
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.DateTimeRange != nil {
		toSerialize["dateTimeRange"] = o.DateTimeRange
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.ShowSubCases != nil {
		toSerialize["showSubCases"] = o.ShowSubCases
	}
	if o.SmartRestRequestContext != nil {
		toSerialize["smartRestRequestContext"] = o.SmartRestRequestContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *CasesSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromIndex               *int64         `json:"fromIndex,omitempty"`
		PageSize                *int64         `json:"pageSize,omitempty"`
		DateTimeRange           *DateTimeRange `json:"dateTimeRange,omitempty"`
		Filter                  *string        `json:"filter,omitempty"`
		App                     *string        `json:"app,omitempty"`
		ShowSubCases            *bool          `json:"showSubCases"`
		SmartRestRequestContext *string        `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"fromIndex", "pageSize", "dateTimeRange", "filter", "app", "showSubCases", "smartRestRequestContext"})
	} else {
		return err
	}

	o.FromIndex = all.FromIndex
	o.PageSize = all.PageSize
	hasInvalidField := false
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	o.Filter = all.Filter
	o.App = all.App
	o.ShowSubCases = all.ShowSubCases
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
