package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourcesReturnType represents the type of date time range for events.
type SourcesReturnType string

// Allowed values for SourcesReturnType.
const (
	RETURNTYPE_JSON SourcesReturnType = "JSON"
)

// allowedReturnEnumValues contains the allowed values for SourcesReturnType.
var allowedReturnEnumValues = []SourcesReturnType{
	RETURNTYPE_JSON,
}

// GetAllowedValues returns the list of possible values.
func (v *SourcesReturnType) GetAllowedValues() []SourcesReturnType {
	return allowedReturnEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourcesReturnType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcesReturnType(value)
	return nil
}

// NewSourcesReturnTypeFromValue returns a pointer to a valid SourcesReturnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcesReturnTypeFromValue(v string) (*SourcesReturnType, error) {
	ev := SourcesReturnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcesReturnType: valid values are %v", v, allowedReturnEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcesReturnType) IsValid() bool {
	for _, existing := range allowedReturnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourcesReturnType value.
func (v SourcesReturnType) Ptr() *SourcesReturnType {
	return &v
}
