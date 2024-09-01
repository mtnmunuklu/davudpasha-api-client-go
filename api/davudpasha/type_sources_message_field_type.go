package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourcesMessageFieldType represents the type of date time range for events.
type SourcesMessageFieldType string

// Allowed values for SourcesMessageFieldType.
const (
	MESSAGEFIELDTYPE_XML  SourcesMessageFieldType = "xml"
	MESSAGEFIELDTYPE_DATA SourcesMessageFieldType = "data"
)

// allowedMessagFieldEnumValues contains the allowed values for SourcesMessageFieldType.
var allowedMessagFieldEnumValues = []SourcesMessageFieldType{
	MESSAGEFIELDTYPE_XML,
	MESSAGEFIELDTYPE_DATA,
}

// GetAllowedValues returns the list of possible values.
func (v *SourcesMessageFieldType) GetAllowedValues() []SourcesMessageFieldType {
	return allowedMessagFieldEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourcesMessageFieldType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcesMessageFieldType(value)
	return nil
}

// NewSourcesMessageFieldTypeFromValue returns a pointer to a valid SourcesMessageFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcesMessageFieldTypeFromValue(v string) (*SourcesMessageFieldType, error) {
	ev := SourcesMessageFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcesMessageFieldType: valid values are %v", v, allowedMessagFieldEnumValues)
}

// IsValid returns true if the value is valid for the enum, false otherwise.
func (v SourcesMessageFieldType) IsValid() bool {
	for _, existing := range allowedMessagFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourcesMessageFieldType value.
func (v SourcesMessageFieldType) Ptr() *SourcesMessageFieldType {
	return &v
}

// NullableSourcesMessageFieldType is a struct that wraps SourcesMessageFieldType to support null values.
type NullableSourcesMessageFieldType struct {
	value *SourcesMessageFieldType
	isSet bool
}

// Get returns the value.
func (v NullableSourcesMessageFieldType) Get() *SourcesMessageFieldType {
	return v.value
}

// Set modifies the value and marks it as set.
func (v *NullableSourcesMessageFieldType) Set(val *SourcesMessageFieldType) {
	v.value = val
	v.isSet = true
}

// IsSet returns true if the value has been set.
func (v NullableSourcesMessageFieldType) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSourcesMessageFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSourcesMessageFieldType initializes a new NullableSourcesMessageFieldType with a set value.
func NewNullableSourcesMessageFieldType(val *SourcesMessageFieldType) *NullableSourcesMessageFieldType {
	return &NullableSourcesMessageFieldType{value: val, isSet: true}
}

// MarshalJSON serializes the value if set, otherwise it returns null.
func (v NullableSourcesMessageFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the value and marks it as set.
func (v *NullableSourcesMessageFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
