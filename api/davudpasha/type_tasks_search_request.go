package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type TasksSearchRequest struct {
	Filter          *string `json:"filter,omitempty"`
	ExtraFilter     *string `json:"extraFilter,omitempty"`
	ShowServiceData *bool   `json:"showServiceData,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
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

// GetShowServiceData returns the ShowServiceData field value if set, zero value otherwise.
func (o *TasksSearchRequest) GetShowServiceData() bool {
	if o == nil || o.ShowServiceData == nil {
		var ret bool
		return ret
	}
	return *o.ShowServiceData
}

// GetShowServiceDataOk returns a tuple with the ShowServiceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchRequest) GetShowServiceDataOk() (*bool, bool) {
	if o == nil || o.ShowServiceData == nil {
		return nil, false
	}
	return o.ShowServiceData, true
}

// HasShowServiceData returns a boolean if a field has been set.
func (o *TasksSearchRequest) HasShowServiceData() bool {
	return o != nil && o.ShowServiceData != nil
}

// SetShowServiceData gets a reference to the given bool and assigns it to the ShowServiceData field.
func (o *TasksSearchRequest) SetShowServiceData(v bool) {
	o.ShowServiceData = &v
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
	if o.ShowServiceData != nil {
		toSerialize["showServiceData"] = o.ShowServiceData
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
		ShowServiceData         *bool   `json:"showServiceData,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"filter", "extraFilter", "showServiceData", "smartRestRequestContext"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.ExtraFilter = all.ExtraFilter
	o.ShowServiceData = all.ShowServiceData
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
