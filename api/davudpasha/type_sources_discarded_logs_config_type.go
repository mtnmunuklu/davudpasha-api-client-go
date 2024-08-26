package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourcesDiscardedLogsConfigType represents the type of date time range for events.
type SourcesDiscardedLogsConfigType string

// Allowed values for SourcesDiscardedLogsConfigType.
const (
	DISCARDEDLOGSCONFIGTYPE_TODETAILS SourcesDiscardedLogsConfigType = "ToDetails"
)

// allowedDiscardedLogsConfigEnumValues contains the allowed values for SourcesDiscardedLogsConfigType.
var allowedDiscardedLogsConfigEnumValues = []SourcesDiscardedLogsConfigType{
	DISCARDEDLOGSCONFIGTYPE_TODETAILS,
}

// GetAllowedValues returns the list of possible values.
func (v *SourcesDiscardedLogsConfigType) GetAllowedValues() []SourcesDiscardedLogsConfigType {
	return allowedDiscardedLogsConfigEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourcesDiscardedLogsConfigType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcesDiscardedLogsConfigType(value)
	return nil
}

// NewSourcesDiscardedLogsConfigTypeFromValue returns a pointer to a valid SourcesDiscardedLogsConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcesDiscardedLogsConfigTypeFromValue(v string) (*SourcesDiscardedLogsConfigType, error) {
	ev := SourcesDiscardedLogsConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcesDiscardedLogsConfigType: valid values are %v", v, allowedDiscardedLogsConfigEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcesDiscardedLogsConfigType) IsValid() bool {
	for _, existing := range allowedDiscardedLogsConfigEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourcesDiscardedLogsConfigType value.
func (v SourcesDiscardedLogsConfigType) Ptr() *SourcesDiscardedLogsConfigType {
	return &v
}
