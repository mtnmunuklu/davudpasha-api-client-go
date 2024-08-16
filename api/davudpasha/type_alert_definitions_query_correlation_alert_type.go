package davudpasha

import (
	"encoding/json"
	"fmt"
)

// AlertDefinitionsQueryCorrelationAlertType represents the type of date time range for events.
type AlertDefinitionsQueryCorrelationAlertType string

// Allowed values for AlertDefinitionsQueryCorrelationAlertType.
const (
	QUERYCORRELATIONALERTTYPE_WHENONEORMOREROW AlertDefinitionsQueryCorrelationAlertType = "WhenOneOrMoreRow"
	QUERYCORRELATIONALERTTYPE_WHENNOROW        AlertDefinitionsQueryCorrelationAlertType = "WhenNoRow"
	QUERYCORRELATIONALERTTYPE_ALERTFOREACHROW  AlertDefinitionsQueryCorrelationAlertType = "AlertForEachRow"
)

// allowedQueryCorrelationAlertEnumValues contains the allowed values for AlertDefinitionsQueryCorrelationAlertType.
var allowedQueryCorrelationAlertEnumValues = []AlertDefinitionsQueryCorrelationAlertType{
	QUERYCORRELATIONALERTTYPE_WHENONEORMOREROW,
	QUERYCORRELATIONALERTTYPE_WHENNOROW,
	QUERYCORRELATIONALERTTYPE_ALERTFOREACHROW,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertDefinitionsQueryCorrelationAlertType) GetAllowedValues() []AlertDefinitionsQueryCorrelationAlertType {
	return allowedQueryCorrelationAlertEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *AlertDefinitionsQueryCorrelationAlertType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertDefinitionsQueryCorrelationAlertType(value)
	return nil
}

// NewAlertDefinitionsQueryCorrelationAlertTypeFromValue returns a pointer to a valid AlertDefinitionsQueryCorrelationAlertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertDefinitionsQueryCorrelationAlertTypeFromValue(v string) (*AlertDefinitionsQueryCorrelationAlertType, error) {
	ev := AlertDefinitionsQueryCorrelationAlertType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertDefinitionsQueryCorrelationAlertType: valid values are %v", v, allowedQueryCorrelationAlertEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertDefinitionsQueryCorrelationAlertType) IsValid() bool {
	for _, existing := range allowedQueryCorrelationAlertEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to AlertDefinitionsQueryCorrelationAlertType value.
func (v AlertDefinitionsQueryCorrelationAlertType) Ptr() *AlertDefinitionsQueryCorrelationAlertType {
	return &v
}
