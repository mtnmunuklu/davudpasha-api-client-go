package davudpasha

import (
	"encoding/json"
	"fmt"
)

// AlertsQueryCorrelationAlertType represents the type of date time range for events.
type AlertsQueryCorrelationAlertType string

// Allowed values for AlertsQueryCorrelationAlertType.
const (
	QUERYCORRELATIONALERTTYPE_WHENONEORMOREROW AlertsQueryCorrelationAlertType = "WhenOneOrMoreRow"
	QUERYCORRELATIONALERTTYPE_WHENNOROW        AlertsQueryCorrelationAlertType = "WhenNoRow"
	QUERYCORRELATIONALERTTYPE_ALERTFOREACHROW  AlertsQueryCorrelationAlertType = "AlertForEachRow"
)

// allowedQueryCorrelationAlertEnumValues contains the allowed values for AlertsQueryCorrelationAlertType.
var allowedQueryCorrelationAlertEnumValues = []AlertsQueryCorrelationAlertType{
	QUERYCORRELATIONALERTTYPE_WHENONEORMOREROW,
	QUERYCORRELATIONALERTTYPE_WHENNOROW,
	QUERYCORRELATIONALERTTYPE_ALERTFOREACHROW,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertsQueryCorrelationAlertType) GetAllowedValues() []AlertsQueryCorrelationAlertType {
	return allowedQueryCorrelationAlertEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *AlertsQueryCorrelationAlertType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertsQueryCorrelationAlertType(value)
	return nil
}

// NewAlertsQueryCorrelationAlertTypeFromValue returns a pointer to a valid AlertsQueryCorrelationAlertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertsQueryCorrelationAlertTypeFromValue(v string) (*AlertsQueryCorrelationAlertType, error) {
	ev := AlertsQueryCorrelationAlertType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertsQueryCorrelationAlertType: valid values are %v", v, allowedQueryCorrelationAlertEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertsQueryCorrelationAlertType) IsValid() bool {
	for _, existing := range allowedQueryCorrelationAlertEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to AlertsQueryCorrelationAlertType value.
func (v AlertsQueryCorrelationAlertType) Ptr() *AlertsQueryCorrelationAlertType {
	return &v
}
