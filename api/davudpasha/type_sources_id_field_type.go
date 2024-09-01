package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourcesIDFieldType represents the type of date time range for events.
type SourcesIDFieldType string

// Allowed values for SourcesIDFieldType.
const (
	IDFIELDTYPE_DATETIME SourcesIDFieldType = "DateTime"
)

// allowedIDFieldEnumValues contains the allowed values for SourcesIDFieldType.
var allowedIDFieldEnumValues = []SourcesIDFieldType{
	IDFIELDTYPE_DATETIME,
}

// GetAllowedValues returns the list of possible values.
func (v *SourcesIDFieldType) GetAllowedValues() []SourcesIDFieldType {
	return allowedIDFieldEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourcesIDFieldType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcesIDFieldType(value)
	return nil
}

// NewSourcesIDFieldTypeFromValue returns a pointer to a valid SourcesIDFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcesIDFieldTypeFromValue(v string) (*SourcesIDFieldType, error) {
	ev := SourcesIDFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcesIDFieldType: valid values are %v", v, allowedIDFieldEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcesIDFieldType) IsValid() bool {
	for _, existing := range allowedIDFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourcesIDFieldType value.
func (v SourcesIDFieldType) Ptr() *SourcesIDFieldType {
	return &v
}
