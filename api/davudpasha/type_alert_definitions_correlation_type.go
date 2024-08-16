package davudpasha

import (
	"encoding/json"
	"fmt"
)

// AlertDefinitionsCorrelationType represents the type of date time range for events.
type AlertDefinitionsCorrelationType string

// Allowed values for AlertDefinitionsCorrelationType.
const (
	CORRELATIONTYPE_INTERFACEIQUERYCORRELATION      AlertDefinitionsCorrelationType = "Interface IQueryCorrelation"
	CORRELATIONTYPE_INTERFACEIQUERYCORRELATIONCHAIN AlertDefinitionsCorrelationType = "Interface IQueryCorrelationChain"
	CORRELATIONTYPE_INTERFACEILOGSOURCECORRELATION  AlertDefinitionsCorrelationType = "Interface ILogSourceCorrelation"
)

// allowedCorrelationEnumValues contains the allowed values for AlertDefinitionsCorrelationType.
var allowedCorrelationEnumValues = []AlertDefinitionsCorrelationType{
	CORRELATIONTYPE_INTERFACEIQUERYCORRELATION,
	CORRELATIONTYPE_INTERFACEIQUERYCORRELATIONCHAIN,
	CORRELATIONTYPE_INTERFACEILOGSOURCECORRELATION,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertDefinitionsCorrelationType) GetAllowedValues() []AlertDefinitionsCorrelationType {
	return allowedCorrelationEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *AlertDefinitionsCorrelationType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertDefinitionsCorrelationType(value)
	return nil
}

// NewAlertDefinitionsCorrelationTypeFromValue returns a pointer to a valid AlertDefinitionsCorrelationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertDefinitionsCorrelationTypeFromValue(v string) (*AlertDefinitionsCorrelationType, error) {
	ev := AlertDefinitionsCorrelationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertDefinitionsCorrelationType: valid values are %v", v, allowedCorrelationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertDefinitionsCorrelationType) IsValid() bool {
	for _, existing := range allowedCorrelationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to AlertDefinitionsCorrelationType value.
func (v AlertDefinitionsCorrelationType) Ptr() *AlertDefinitionsCorrelationType {
	return &v
}
