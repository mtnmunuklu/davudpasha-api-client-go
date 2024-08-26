package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SourceTypesSaveRequest struct {
	LgsDef *SourceTypesItem `json:"lgsDef,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourceTypesSaveRequest creates a new SourceTypesSaveRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesSaveRequest() *SourceTypesSaveRequest {
	this := SourceTypesSaveRequest{}
	return &this
}

// NewSourceTypesSaveRequestWithDefaults creates a new SourceTypesSaveRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourceTypesSaveRequestWithDefaults() *SourceTypesSaveRequest {
	this := SourceTypesSaveRequest{}
	return &this
}

// GetLgsDef returns the LgsDef field value if set, zero value otherwise.
func (o *SourceTypesSaveRequest) GetLgsDef() SourceTypesItem {
	if o == nil || o.LgsDef == nil {
		var ret SourceTypesItem
		return ret
	}
	return *o.LgsDef
}

// GetLgsDefOk returns a tuple with the LgsDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesSaveRequest) GetLgsDefOk() (*SourceTypesItem, bool) {
	if o == nil || o.LgsDef == nil {
		return nil, false
	}
	return o.LgsDef, true
}

// HasLgsDef returns a boolean if a field has been set.
func (o *SourceTypesSaveRequest) HasLgsDef() bool {
	return o != nil && o.LgsDef != nil
}

// SetLgsDef gets a reference to the given SourceTypesItem and assigns it to the LgsDef field.
func (o *SourceTypesSaveRequest) SetLgsDef(v SourceTypesItem) {
	o.LgsDef = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *SourceTypesSaveRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesSaveRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *SourceTypesSaveRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *SourceTypesSaveRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesSaveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.LgsDef != nil {
		toSerialize["lgsDef"] = o.LgsDef
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
func (o *SourceTypesSaveRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		LgsDef                  *SourceTypesItem `json:"lgsDef,omitempty"`
		SmartRestRequestContext *string          `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LgsDef == nil {
		return fmt.Errorf("requiered field lgsDef is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"lgsDef", "smartRestRequestContext"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.LgsDef != nil && all.LgsDef.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LgsDef = all.LgsDef
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
