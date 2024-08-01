package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsSearchRequest represents a request to search reports.
type ReportsSearchRequest struct {
	// Filter for the search.
	SearchFilter *string `json:"searchFilter,omitempty"`
	// Name of the application.
	ApplicationName *string `json:"applicationName,omitempty"`
	// Start date for the search.
	StartDate common.NullableString `json:"startDate,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Indicates whether to show passive reports.
	ShowPassive *bool `json:"showPassive,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsSearchRequest creates a new ReportsSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsSearchRequest() *ReportsSearchRequest {
	this := ReportsSearchRequest{}
	return &this
}

// NewReportsSearchRequestWithDefaults creates a new ReportsSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsSearchRequestWithDefaults() *ReportsSearchRequest {
	this := ReportsSearchRequest{}
	return &this
}

// GetSearchFilter returns the SearchFilter field value if set, zero value otherwise.
func (o *ReportsSearchRequest) GetSearchFilter() string {
	if o == nil || o.SearchFilter == nil {
		var ret string
		return ret
	}
	return *o.SearchFilter
}

// GetSearchFilterOk returns a tuple with the SearchFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchRequest) GetSearchFilterOk() (*string, bool) {
	if o == nil || o.SearchFilter == nil {
		return nil, false
	}
	return o.SearchFilter, true
}

// HasSearchFilter returns a boolean if a field has been set.
func (o *ReportsSearchRequest) HasSearchFilter() bool {
	return o != nil && o.SearchFilter != nil
}

// SetSearchFilter gets a reference to the given string and assigns it to the SearchFilter field.
func (o *ReportsSearchRequest) SetSearchFilter(v string) {
	o.SearchFilter = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *ReportsSearchRequest) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchRequest) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *ReportsSearchRequest) HasApplicationName() bool {
	return o != nil && o.ApplicationName != nil
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *ReportsSearchRequest) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsSearchRequest) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsSearchRequest) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a StartDate has been set.
func (o *ReportsSearchRequest) HasStartDate() bool {
	return o != nil && o.StartDate.IsSet()
}

// SetStartDate gets a reference to the given common.NullableString and assigns it to the StartDate field.
func (o *ReportsSearchRequest) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

// SetStartDateNil sets the value for StartDate to be an explicit nil.
func (o *ReportsSearchRequest) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnSetStartDate ensures that no value is present for StartDate, not even an explicit nil.
func (o *ReportsSearchRequest) UnSetStartDate() {
	o.StartDate.UnSet()
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *ReportsSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *ReportsSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *ReportsSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// GetShowPassive returns the ShowPassive field value if set, zero value otherwise.
func (o *ReportsSearchRequest) GetShowPassive() bool {
	if o == nil || o.ShowPassive == nil {
		var ret bool
		return ret
	}
	return *o.ShowPassive
}

// GetShowPassiveOk returns a tuple with the ShowPassive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchRequest) GetShowPassiveOk() (*bool, bool) {
	if o == nil || o.ShowPassive == nil {
		return nil, false
	}
	return o.ShowPassive, true
}

// HasShowPassive returns a boolean if a field has been set.
func (o *ReportsSearchRequest) HasShowPassive() bool {
	return o != nil && o.ShowPassive != nil
}

// SetShowPassive gets a reference to the given bool and assigns it to the ShowPassive field.
func (o *ReportsSearchRequest) SetShowPassive(v bool) {
	o.ShowPassive = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}

	toSerialize["startDate"] = o.StartDate.Get()

	if o.SearchFilter != nil {
		toSerialize["searchFilter"] = o.SearchFilter
	}
	if o.ApplicationName != nil {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if o.SmartRestRequestContext != nil {
		toSerialize["smartRestRequestContext"] = o.SmartRestRequestContext
	}
	if o.ShowPassive != nil {
		toSerialize["showPassive"] = o.ShowPassive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		SearchFilter            *string               `json:"searchFilter,omitempty"`
		ApplicationName         *string               `json:"applicationName,omitempty"`
		StartDate               common.NullableString `json:"startDate,omitempty"`
		SmartRestRequestContext *string               `json:"smartRestRequestContext,omitempty"`
		ShowPassive             *bool                 `json:"showPassive,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"searchFilter", "applicationName", "startDate", "smartRestRequestContext", "showPassive"})
	} else {
		return err
	}

	o.SearchFilter = all.SearchFilter
	o.ApplicationName = all.ApplicationName
	o.StartDate = all.StartDate
	o.ShowPassive = all.ShowPassive

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
