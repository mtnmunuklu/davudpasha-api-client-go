package davudpasha

import (
	"encoding/json"
	"fmt"
)

// AlertDefinitionsTimeFrameType represents the type of date time range for events.
type AlertDefinitionsTimeFrameType string

// Allowed values for AlertDefinitionsTimeFrameType.
const (
	TIMEFRAMETYPE_MINUTES AlertDefinitionsTimeFrameType = "minutes"
	TIMEFRAMETYPE_HOURS   AlertDefinitionsTimeFrameType = "hours"
	TIMEFRAMETYPE_DAYS    AlertDefinitionsTimeFrameType = "days"
	TIMEFRAMETYPE_WEEKS   AlertDefinitionsTimeFrameType = "weeks"
)

// allowedTimeFrameEnumValues contains the allowed values for AlertDefinitionsTimeFrameType.
var allowedTimeFrameEnumValues = []AlertDefinitionsTimeFrameType{
	TIMEFRAMETYPE_MINUTES,
	TIMEFRAMETYPE_HOURS,
	TIMEFRAMETYPE_DAYS,
	TIMEFRAMETYPE_WEEKS,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertDefinitionsTimeFrameType) GetAllowedValues() []AlertDefinitionsTimeFrameType {
	return allowedTimeFrameEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *AlertDefinitionsTimeFrameType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertDefinitionsTimeFrameType(value)
	return nil
}

// NewAlertDefinitionsTimeFrameTypeFromValue returns a pointer to a valid AlertDefinitionsTimeFrameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertDefinitionsTimeFrameTypeFromValue(v string) (*AlertDefinitionsTimeFrameType, error) {
	ev := AlertDefinitionsTimeFrameType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertDefinitionsTimeFrameType: valid values are %v", v, allowedTimeFrameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertDefinitionsTimeFrameType) IsValid() bool {
	for _, existing := range allowedTimeFrameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to AlertDefinitionsTimeFrameType value.
func (v AlertDefinitionsTimeFrameType) Ptr() *AlertDefinitionsTimeFrameType {
	return &v
}
