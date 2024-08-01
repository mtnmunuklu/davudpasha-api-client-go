package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// TasksSearchRequest represents the structure for task search requests.
type TasksSearchRequest struct {
	// Filter specifies the main filter criteria for task search.
	Filter *string `json:"filter,omitempty"`
	// ExtraFilter specifies additional filter criteria for task search.
	ExtraFilter *string `json:"extraFilter,omitempty"`
	// ShowServicesData indicates whether to include service data in the search results.
	ShowServicesData *bool `json:"showServicesData,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTasksSearchRequest creates a new TasksSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTasksSearchRequest() *TasksSearchRequest {
	this := TasksSearchRequest{}
	return &this
}

// NewTasksSearchRequestWithDefaults creates a new TasksSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTasksSearchRequestWithDefaults() *TasksSearchRequest {
	this := TasksSearchRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TasksSearchRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TasksSearchRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *TasksSearchRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetExtraFilter returns the ExtraFilter field value if set, zero value otherwise.
func (o *TasksSearchRequest) GetExtraFilter() string {
	if o == nil || o.ExtraFilter == nil {
		var ret string
		return ret
	}
	return *o.ExtraFilter
}

// GetExtraFilterOk returns a tuple with the ExtraFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchRequest) GetExtraFilterOk() (*string, bool) {
	if o == nil || o.ExtraFilter == nil {
		return nil, false
	}
	return o.ExtraFilter, true
}

// HasExtraFilter returns a boolean if a field has been set.
func (o *TasksSearchRequest) HasExtraFilter() bool {
	return o != nil && o.ExtraFilter != nil
}

// SetExtraFilter gets a reference to the given string and assigns it to the ExtraFilter field.
func (o *TasksSearchRequest) SetExtraFilter(v string) {
	o.ExtraFilter = &v
}

// GetShowServicesData returns the ShowServicesData field value if set, zero value otherwise.
func (o *TasksSearchRequest) GetShowServicesData() bool {
	if o == nil || o.ShowServicesData == nil {
		var ret bool
		return ret
	}
	return *o.ShowServicesData
}

// GetShowServicesDataOk returns a tuple with the ShowServicesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchRequest) GetShowServicesDataOk() (*bool, bool) {
	if o == nil || o.ShowServicesData == nil {
		return nil, false
	}
	return o.ShowServicesData, true
}

// HasShowServicesData returns a boolean if a field has been set.
func (o *TasksSearchRequest) HasShowServicesData() bool {
	return o != nil && o.ShowServicesData != nil
}

// SetShowServicesData gets a reference to the given bool and assigns it to the ShowServicesData field.
func (o *TasksSearchRequest) SetShowServicesData(v bool) {
	o.ShowServicesData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TasksSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.ExtraFilter != nil {
		toSerialize["extraFilter"] = o.ExtraFilter
	}
	if o.ShowServicesData != nil {
		toSerialize["showServicesData"] = o.ShowServicesData
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
func (o *TasksSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter                  *string `json:"filter,omitempty"`
		ExtraFilter             *string `json:"extraFilter,omitempty"`
		ShowServicesData        *bool   `json:"showServicesData,omitempty"`
		SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filter == nil {
		return fmt.Errorf("requiered field Filter is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filter", "extraFilter", "showServicesData", "smartRestRequestContext"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.ExtraFilter = all.ExtraFilter
	o.ShowServicesData = all.ShowServicesData
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
