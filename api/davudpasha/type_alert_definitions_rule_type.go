package davudpasha

import (
	"encoding/json"
	"fmt"
)

// AlertDefinitionsRuleType represents the type of date time range for events.
type AlertDefinitionsRuleType string

// Allowed values for AlertDefinitionsRuleType.
const (
	RULETYPE_ANY AlertDefinitionsRuleType = "any"
)

// allowedRuleEnumValues contains the allowed values for AlertDefinitionsRuleType.
var allowedRuleEnumValues = []AlertDefinitionsRuleType{
	RULETYPE_ANY,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertDefinitionsRuleType) GetAllowedValues() []AlertDefinitionsRuleType {
	return allowedRuleEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *AlertDefinitionsRuleType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertDefinitionsRuleType(value)
	return nil
}

// NewAlertDefinitionsRuleTypeFromValue returns a pointer to a valid AlertDefinitionsRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertDefinitionsRuleTypeFromValue(v string) (*AlertDefinitionsRuleType, error) {
	ev := AlertDefinitionsRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertDefinitionsRuleType: valid values are %v", v, allowedRuleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertDefinitionsRuleType) IsValid() bool {
	for _, existing := range allowedRuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to AlertDefinitionsRuleType value.
func (v AlertDefinitionsRuleType) Ptr() *AlertDefinitionsRuleType {
	return &v
}
