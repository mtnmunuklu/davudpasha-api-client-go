package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ActionDefinationsSearchRequest struct {
	Application *string `json:"application,omitempty"`
	// Filter for searching action definations.
	SearchFilter *string `json:"searchFilter,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewActionDefinationsSearchRequest creates a new ActionDefinationsSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionDefinationsSearchRequest() *ActionDefinationsSearchRequest {
	this := ActionDefinationsSearchRequest{}
	return &this
}

// NewActionDefinationsSearchRequestWithDefaults creates a new ActionDefinationsSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionDefinationsSearchRequestWithDefaults() *ActionDefinationsSearchRequest {
	this := ActionDefinationsSearchRequest{}
	return &this
}

// GetSearchFilter returns the SearchFilter field value if set, zero value otherwise.
func (o *ActionDefinationsSearchRequest) GetSearchFilter() string {
	if o == nil || o.SearchFilter == nil {
		var ret string
		return ret
	}
	return *o.SearchFilter
}

// GetSearchFilterOk returns a tuple with the SearchFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchRequest) GetSearchFilterOk() (*string, bool) {
	if o == nil || o.SearchFilter == nil {
		return nil, false
	}
	return o.SearchFilter, true
}

// HasSearchFilter returns a boolean if a field has been set.
func (o *ActionDefinationsSearchRequest) HasSearchFilter() bool {
	return o != nil && o.SearchFilter != nil
}

// SetSearchFilter gets a reference to the given string and assigns it to the SearchFilter field.
func (o *ActionDefinationsSearchRequest) SetSearchFilter(v string) {
	o.SearchFilter = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *ActionDefinationsSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *ActionDefinationsSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *ActionDefinationsSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ActionDefinationsSearchRequest) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchRequest) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ActionDefinationsSearchRequest) HasApplication() bool {
	return o != nil && o.Application != nil
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ActionDefinationsSearchRequest) SetApplication(v string) {
	o.Application = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionDefinationsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.SearchFilter != nil {
		toSerialize["searchFilter"] = o.SearchFilter
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
func (o *ActionDefinationsSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Application             *string `json:"application,omitempty"`
		SearchFilter            *string `json:"searchFilter,omitempty"`
		SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SearchFilter == nil {
		return fmt.Errorf("requiered field searchFilter is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"application", "searchFilter", "smartRestRequestContext"})
	} else {
		return err
	}
	o.SearchFilter = all.SearchFilter
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
