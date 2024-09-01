package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourcesIDColumnType represents the type of date time range for events.
type SourcesIDColumnType string

// Allowed values for SourcesIDColumnType.
const (
	IDCOLUMNTYPE_DATETIME SourcesIDColumnType = "DateTime"
)

// allowedIDColumnEnumValues contains the allowed values for SourcesIDColumnType.
var allowedIDColumnEnumValues = []SourcesIDColumnType{
	IDCOLUMNTYPE_DATETIME,
}

// GetAllowedValues returns the list of possible values.
func (v *SourcesIDColumnType) GetAllowedValues() []SourcesIDColumnType {
	return allowedIDColumnEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourcesIDColumnType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcesIDColumnType(value)
	return nil
}

// NewSourcesIDColumnTypeFromValue returns a pointer to a valid SourcesIDColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcesIDColumnTypeFromValue(v string) (*SourcesIDColumnType, error) {
	ev := SourcesIDColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcesIDColumnType: valid values are %v", v, allowedIDColumnEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcesIDColumnType) IsValid() bool {
	for _, existing := range allowedIDColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourcesIDColumnType value.
func (v SourcesIDColumnType) Ptr() *SourcesIDColumnType {
	return &v
}
