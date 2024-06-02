package davudpasha

import "encoding/json"

type SourceTypesNormalization struct {
	OtherValue *string                        `json:"OtherValue,omitempty"`
	Items      []SourceTypesNormalizationItem `json:"Items,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
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
func (v *NullableSourceTypesNormalization) UnSet() {
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
