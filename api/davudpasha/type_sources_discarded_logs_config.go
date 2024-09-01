package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourcesDiscardedLogsConfig represents the type of date time range for events.
type SourcesDiscardedLogsConfig string

// Allowed values for SourcesDiscardedLogsConfig.
const (
	DISCARDEDLOGSCONFIG_TODETAILS SourcesDiscardedLogsConfig = "ToDetails"
)

// allowedDiscardedLogsEnumValues contains the allowed values for SourcesDiscardedLogsConfig.
var allowedDiscardedLogsEnumValues = []SourcesDiscardedLogsConfig{
	DISCARDEDLOGSCONFIG_TODETAILS,
}

// GetAllowedValues returns the list of possible values.
func (v *SourcesDiscardedLogsConfig) GetAllowedValues() []SourcesDiscardedLogsConfig {
	return allowedDiscardedLogsEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourcesDiscardedLogsConfig) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcesDiscardedLogsConfig(value)
	return nil
}

// NewSourcesDiscardedLogsConfigFromValue returns a pointer to a valid SourcesDiscardedLogsConfig
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcesDiscardedLogsConfigFromValue(v string) (*SourcesDiscardedLogsConfig, error) {
	ev := SourcesDiscardedLogsConfig(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcesDiscardedLogsConfig: valid values are %v", v, allowedDiscardedLogsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcesDiscardedLogsConfig) IsValid() bool {
	for _, existing := range allowedDiscardedLogsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourcesDiscardedLogsConfig value.
func (v SourcesDiscardedLogsConfig) Ptr() *SourcesDiscardedLogsConfig {
	return &v
}
