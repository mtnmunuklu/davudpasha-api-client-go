package davudpasha

import (
	"encoding/json"
	"fmt"
)

// AlertsRuleType represents the type of date time range for events.
type AlertsRuleType string

// Allowed values for AlertsRuleType.
const (
	RULETYPE_ANY AlertsRuleType = "any"
)

// allowedRuleEnumValues contains the allowed values for AlertsRuleType.
var allowedRuleEnumValues = []AlertsRuleType{
	RULETYPE_ANY,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertsRuleType) GetAllowedValues() []AlertsRuleType {
	return allowedRuleEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *AlertsRuleType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertsRuleType(value)
	return nil
}

// NewAlertsRuleTypeFromValue returns a pointer to a valid AlertsRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertsRuleTypeFromValue(v string) (*AlertsRuleType, error) {
	ev := AlertsRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertsRuleType: valid values are %v", v, allowedRuleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertsRuleType) IsValid() bool {
	for _, existing := range allowedRuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to AlertsRuleType value.
func (v AlertsRuleType) Ptr() *AlertsRuleType {
	return &v
}
