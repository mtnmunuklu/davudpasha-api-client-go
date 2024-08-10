package davudpasha

import (
	"encoding/json"
	"fmt"
)

// AlertsTimeFrameType represents the type of date time range for events.
type AlertsTimeFrameType string

// Allowed values for AlertsTimeFrameType.
const (
	TIMEFRAMETYPE_MINUTES AlertsTimeFrameType = "minutes"
	TIMEFRAMETYPE_HOURS   AlertsTimeFrameType = "hours"
	TIMEFRAMETYPE_DAYS    AlertsTimeFrameType = "days"
	TIMEFRAMETYPE_WEEKS   AlertsTimeFrameType = "weeks"
)

// allowedTimeFrameEnumValues contains the allowed values for AlertsTimeFrameType.
var allowedTimeFrameEnumValues = []AlertsTimeFrameType{
	TIMEFRAMETYPE_MINUTES,
	TIMEFRAMETYPE_HOURS,
	TIMEFRAMETYPE_DAYS,
	TIMEFRAMETYPE_WEEKS,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertsTimeFrameType) GetAllowedValues() []AlertsTimeFrameType {
	return allowedTimeFrameEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *AlertsTimeFrameType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertsTimeFrameType(value)
	return nil
}

// NewAlertsTimeFrameTypeFromValue returns a pointer to a valid AlertsTimeFrameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertsTimeFrameTypeFromValue(v string) (*AlertsTimeFrameType, error) {
	ev := AlertsTimeFrameType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertsTimeFrameType: valid values are %v", v, allowedTimeFrameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertsTimeFrameType) IsValid() bool {
	for _, existing := range allowedTimeFrameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to AlertsTimeFrameType value.
func (v AlertsTimeFrameType) Ptr() *AlertsTimeFrameType {
	return &v
}
