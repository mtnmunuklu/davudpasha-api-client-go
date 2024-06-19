package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SystemLogsSearchRequest struct {
	AppName                 *string               `json:"appName,omitempty"`
	Severity                common.NullableString `json:"severity,omitempty"`
	FromIndex               *int64                `json:"fromIndex,omitempty"`
	PageSize                *int64                `json:"pageSize,omitempty"`
	DateTimeRange           *DateTimeRange        `json:"dateTimeRange,omitempty"`
	Filter                  *string               `json:"filter,omitempty"`
	ShowServicesData        *bool                 `json:"showServicesData,omitempty"`
	ShowAllTenantsData      *bool                 `json:"showAllTenantsData,omitempty"`
	SmartRestRequestContext *string               `json:"smartRestRequestContext,omitempty"`
	VisualStyled            *bool                 `json:"visualStyled,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSystemLogsSearchRequest creates a new SystemLogsSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSystemLogsSearchRequest() *SystemLogsSearchRequest {
	this := SystemLogsSearchRequest{}
	return &this
}

// NewSystemLogsSearchRequestWithDefaults creates a new SystemLogsSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSystemLogsSearchRequestWithDefaults() *SystemLogsSearchRequest {
	this := SystemLogsSearchRequest{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetAppName() string {
	if o == nil || o.AppName == nil {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetAppNameOk() (*string, bool) {
	if o == nil || o.AppName == nil {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasAppName() bool {
	return o != nil && o.AppName != nil
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *SystemLogsSearchRequest) SetAppName(v string) {
	o.AppName = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsSearchRequest) GetSeverity() string {
	if o == nil || o.Severity.Get() == nil {
		var ret string
		return ret
	}
	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsSearchRequest) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// HasSeverity returns a boolean if a Severity has been set.
func (o *SystemLogsSearchRequest) HasSeverity() bool {
	return o != nil && o.Severity.IsSet()
}

// SetSeverity gets a reference to the given datadog.NullableString and assigns it to the Severity field.
func (o *SystemLogsSearchRequest) SetSeverity(v string) {
	o.Severity.Set(&v)
}

// SetSeverityNil sets the value for Severity to be an explicit nil.
func (o *SystemLogsSearchRequest) SetSeverityNil() {
	o.Severity.Set(nil)
}

// UnSetSeverity ensures that no value is present for Severity, not even an explicit nil.
func (o *SystemLogsSearchRequest) UnSetSeverity() {
	o.Severity.UnSet()
}

// GetFromIndex returns the FromIndex field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetFromIndex() int64 {
	if o == nil || o.FromIndex == nil {
		var ret int64
		return ret
	}
	return *o.FromIndex
}

// GetFromIndexOk returns a tuple with the FromIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetFromIndexOk() (*int64, bool) {
	if o == nil || o.FromIndex == nil {
		return nil, false
	}
	return o.FromIndex, true
}

// HasFromIndex returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasFromIndex() bool {
	return o != nil && o.FromIndex != nil
}

// SetFromIndex gets a reference to the given int64 and assigns it to the FromIndex field.
func (o *SystemLogsSearchRequest) SetFromIndex(v int64) {
	o.FromIndex = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasPageSize() bool {
	return o != nil && o.PageSize != nil
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *SystemLogsSearchRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetDateTimeRange() DateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret DateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DateTimeRange field.
func (o *SystemLogsSearchRequest) SetDateTimeRange(v DateTimeRange) {
	o.DateTimeRange = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *SystemLogsSearchRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetShowServicesData returns the ShowServicesData field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetShowServicesData() bool {
	if o == nil || o.ShowServicesData == nil {
		var ret bool
		return ret
	}
	return *o.ShowServicesData
}

// GetShowServicesDataOk returns a tuple with the ShowServicesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetShowServicesDataOk() (*bool, bool) {
	if o == nil || o.ShowServicesData == nil {
		return nil, false
	}
	return o.ShowServicesData, true
}

// HasShowServicesData returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasShowServicesData() bool {
	return o != nil && o.ShowServicesData != nil
}

// SetShowServicesData gets a reference to the given bool and assigns it to the ShowServicesData field.
func (o *SystemLogsSearchRequest) SetShowServicesData(v bool) {
	o.ShowServicesData = &v
}

// GetShowAllTenantsData returns the ShowAllTenantsData field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetShowAllTenantsData() bool {
	if o == nil || o.ShowAllTenantsData == nil {
		var ret bool
		return ret
	}
	return *o.ShowAllTenantsData
}

// GetShowAllTenantsDataOk returns a tuple with the ShowAllTenantsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetShowAllTenantsDataOk() (*bool, bool) {
	if o == nil || o.ShowAllTenantsData == nil {
		return nil, false
	}
	return o.ShowAllTenantsData, true
}

// HasShowAllTenantsData returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasShowAllTenantsData() bool {
	return o != nil && o.ShowAllTenantsData != nil
}

// SetShowAllTenantsData gets a reference to the given bool and assigns it to the ShowAllTenantsData field.
func (o *SystemLogsSearchRequest) SetShowAllTenantsData(v bool) {
	o.ShowAllTenantsData = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *SystemLogsSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// GetVisualStyled returns the VisualStyled field value if set, zero value otherwise.
func (o *SystemLogsSearchRequest) GetVisualStyled() bool {
	if o == nil || o.VisualStyled == nil {
		var ret bool
		return ret
	}
	return *o.VisualStyled
}

// GetVisualStyledOk returns a tuple with the VisualStyled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchRequest) GetVisualStyledOk() (*bool, bool) {
	if o == nil || o.VisualStyled == nil {
		return nil, false
	}
	return o.VisualStyled, true
}

// HasVisualStyled returns a boolean if a field has been set.
func (o *SystemLogsSearchRequest) HasVisualStyled() bool {
	return o != nil && o.VisualStyled != nil
}

// SetVisualStyled gets a reference to the given bool and assigns it to the VisualStyled field.
func (o *SystemLogsSearchRequest) SetVisualStyled(v bool) {
	o.VisualStyled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SystemLogsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AppName != nil {
		toSerialize["appName"] = o.AppName
	}
	toSerialize["severity"] = o.Severity.Get()
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
	if o.ShowServicesData != nil {
		toSerialize["showServicesData"] = o.ShowServicesData
	}
	if o.ShowAllTenantsData != nil {
		toSerialize["showAllTenantsData"] = o.ShowAllTenantsData
	}
	if o.SmartRestRequestContext != nil {
		toSerialize["smartRestRequestContext"] = o.SmartRestRequestContext
	}
	if o.VisualStyled != nil {
		toSerialize["visualStyled"] = o.VisualStyled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SystemLogsSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppName                 *string               `json:"appName,omitempty"`
		Severity                common.NullableString `json:"severitiy,omitempty"`
		FromIndex               *int64                `json:"fromIndex,omitempty"`
		PageSize                *int64                `json:"pageSize,omitempty"`
		DateTimeRange           *DateTimeRange        `json:"dateTimeRange,omitempty"`
		Filter                  *string               `json:"filter,omitempty"`
		ShowServicesData        *bool                 `json:"showServicesData,omitempty"`
		ShowAllTenantsData      *bool                 `json:"showAllTenantsData,omitempty"`
		SmartRestRequestContext *string               `json:"smartRestRequestContext,omitempty"`
		VisualStyled            *bool                 `json:"visualStyled,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.Filter == nil {
		return fmt.Errorf("requiered field filter is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"appName", "severitiy", "fromIndex", "fromIndex", "dateTimeRange", "filter", "showServicesData", "showAllTenantsData", "smartRestRequestContext", "visualStyled"})
	} else {
		return err
	}

	o.AppName = all.AppName
	o.Severity = all.Severity
	o.FromIndex = all.FromIndex
	o.PageSize = all.PageSize
	hasInvalidField := false
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	o.Filter = all.Filter
	o.ShowServicesData = all.ShowServicesData
	o.ShowAllTenantsData = all.ShowAllTenantsData
	o.SmartRestRequestContext = all.SmartRestRequestContext
	o.VisualStyled = all.VisualStyled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
