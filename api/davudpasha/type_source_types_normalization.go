package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourceTypesNormalization represents the normalization settings for source types.
type SourceTypesNormalization struct {
	// OtherValue specifies another value related to normalization.
	OtherValue *string `json:"OtherValue,omitempty"`
	// Items contains a list of normalization items.
	Items []SourceTypesNormalizationItem `json:"Items,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourceTypesNormalization creates a new SourceTypesNormalization object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesNormalization() *SourceTypesNormalization {
	this := SourceTypesNormalization{}
	return &this
}

// NewSourceTypesNormalizationWithDefaults creates a new SourceTypesNormalization object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourceTypesNormalizationWithDefaults() *SourceTypesNormalization {
	this := SourceTypesNormalization{}
	return &this
}

// GetOtherValue returns the OtherValue field value if set, zero value otherwise.
func (o *SourceTypesNormalization) GetOtherValue() string {
	if o == nil || o.OtherValue == nil {
		var ret string
		return ret
	}
	return *o.OtherValue
}

// GetOtherValueOk returns a tuple with the OtherValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesNormalization) GetOtherValueOk() (*string, bool) {
	if o == nil || o.OtherValue == nil {
		return nil, false
	}
	return o.OtherValue, true
}

// HasOtherValue returns a boolean if a field has been set.
func (o *SourceTypesNormalization) HasOtherValue() bool {
	return o != nil && o.OtherValue != nil
}

// SetOtherValue gets a reference to the given string and assigns it to the OtherValue field.
func (o *SourceTypesNormalization) SetOtherValue(v string) {
	o.OtherValue = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SourceTypesNormalization) GetItems() []SourceTypesNormalizationItem {
	if o == nil || o.Items == nil {
		var ret []SourceTypesNormalizationItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesNormalization) GetItemsOk() (*[]SourceTypesNormalizationItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SourceTypesNormalization) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []SourceTypesNormalizationItem and assigns it to the Items field.
func (o *SourceTypesNormalization) SetItems(v []SourceTypesNormalizationItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesNormalization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.OtherValue != nil {
		toSerialize["OtherValue"] = o.OtherValue
	}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourceTypesNormalization) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		OtherValue *string                        `json:"OtherValue,omitempty"`
		Items      []SourceTypesNormalizationItem `json:"Items,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.Items == nil {
		return fmt.Errorf("requiered field Items is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"OtherValue", "Items"})
	} else {
		return err
	}

	o.OtherValue = all.OtherValue
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableSourceTypesNormalization handles when a null is used for SourceTypesNormalization.
type NullableSourceTypesNormalization struct {
	value *SourceTypesNormalization
	isSet bool
}

// Get returns the associated value.
func (v NullableSourceTypesNormalization) Get() *SourceTypesNormalization {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSourceTypesNormalization) Set(val *SourceTypesNormalization) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSourceTypesNormalization) IsSet() bool {
	return v.isSet
}

// UnSet sets the value to nil and resets the set flag/
func (v *NullableSourceTypesNormalization) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSourceTypesNormalization initializes the struct as if Set has been called.
func NewNullableSourceTypesNormalization(val *SourceTypesNormalization) *NullableSourceTypesNormalization {
	return &NullableSourceTypesNormalization{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSourceTypesNormalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSourceTypesNormalization) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
