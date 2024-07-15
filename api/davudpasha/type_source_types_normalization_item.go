package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourceTypesNormalizationItem represents an item in the normalization settings for source types.
type SourceTypesNormalizationItem struct {
	// Key specifies the normalization key.
	Key *string `json:"Key,omitempty"`
	// Value specifies the normalization value.
	Value *string `json:"Value,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSourceTypesNormalizationItem creates a new SourceTypesNormalizationItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesNormalizationItem() *SourceTypesNormalizationItem {
	this := SourceTypesNormalizationItem{}
	return &this
}

// NewSourceTypesNormalizationItemWithDefaults creates a new SourceTypesNormalizationItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourceTypesNormalizationItemWithDefaults() *SourceTypesNormalizationItem {
	this := SourceTypesNormalizationItem{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SourceTypesNormalizationItem) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesNormalizationItem) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SourceTypesNormalizationItem) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SourceTypesNormalizationItem) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SourceTypesNormalizationItem) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesNormalizationItem) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SourceTypesNormalizationItem) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SourceTypesNormalizationItem) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesNormalizationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourceTypesNormalizationItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key   *string `json:"Key,omitempty"`
		Value *string `json:"Value,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.Key == nil {
		return fmt.Errorf("requiered field Key is missing")
	}
	if all.Value == nil {
		return fmt.Errorf("requiered field Value is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Key", "Value"})
	} else {
		return err
	}

	o.Key = all.Key
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
