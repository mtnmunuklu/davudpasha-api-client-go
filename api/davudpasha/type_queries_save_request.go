package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type QueriesSaveRequest struct {
	QuerySettings *QueriesItem `json:"querySettings,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQueriesSaveRequest creates a new QueriesSaveRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueriesSaveRequest() *QueriesSaveRequest {
	this := QueriesSaveRequest{}
	return &this
}

// NewQueriesSaveRequestWithDefaults creates a new QueriesSaveRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueriesSaveRequestWithDefaults() *QueriesSaveRequest {
	this := QueriesSaveRequest{}
	return &this
}

// GetQuerySettings returns the QuerySettings field value if set, zero value otherwise.
func (o *QueriesSaveRequest) GetQuerySettings() QueriesItem {
	if o == nil || o.QuerySettings == nil {
		var ret QueriesItem
		return ret
	}
	return *o.QuerySettings
}

// GetQuerySettingsOk returns a tuple with the QuerySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesSaveRequest) GetQuerySettingsOk() (*QueriesItem, bool) {
	if o == nil || o.QuerySettings == nil {
		return nil, false
	}
	return o.QuerySettings, true
}

// HasQuerySettings returns a boolean if a field has been set.
func (o *QueriesSaveRequest) HasQuerySettings() bool {
	return o != nil && o.QuerySettings != nil
}

// SetQuerySettings gets a reference to the given QueriesItem and assigns it to the QuerySettings field.
func (o *QueriesSaveRequest) SetQuerySettings(v QueriesItem) {
	o.QuerySettings = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *QueriesSaveRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesSaveRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *QueriesSaveRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *QueriesSaveRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueriesSaveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.QuerySettings != nil {
		toSerialize["querySettings"] = o.QuerySettings
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
func (o *QueriesSaveRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		QuerySettings           *QueriesItem `json:"querySettings,omitempty"`
		SmartRestRequestContext *string      `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.QuerySettings == nil {
		return fmt.Errorf("requiered field querySettings is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"querySettings", "smartRestRequestContext"})
	} else {
		return err
	}
	hasInvalidField := false
	if all.QuerySettings != nil && all.QuerySettings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
