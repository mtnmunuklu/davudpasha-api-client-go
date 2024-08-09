package davudpasha

import (
	"encoding/json"
	"fmt"
)

// AlertsCorrelationType represents the type of date time range for events.
type AlertsCorrelationType string

// Allowed values for AlertsCorrelationType.
const (
	CORRELATIONTYPE_INTERFACEIQUERYCORRELATION      AlertsCorrelationType = "Interface IQueryCorrelation"
	CORRELATIONTYPE_INTERFACEIQUERYCORRELATIONCHAIN AlertsCorrelationType = "Interface IQueryCorrelationChain"
	CORRELATIONTYPE_INTERFACEILOGSOURCECORRELATION  AlertsCorrelationType = "Interface ILogSourceCorrelation"
)

// allowedCorrelationEnumValues contains the allowed values for AlertsCorrelationType.
var allowedCorrelationEnumValues = []AlertsCorrelationType{
	CORRELATIONTYPE_INTERFACEIQUERYCORRELATION,
	CORRELATIONTYPE_INTERFACEIQUERYCORRELATIONCHAIN,
	CORRELATIONTYPE_INTERFACEILOGSOURCECORRELATION,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertsCorrelationType) GetAllowedValues() []AlertsCorrelationType {
	return allowedCorrelationEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *AlertsCorrelationType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertsCorrelationType(value)
	return nil
}

// NewAlertsCorrelationTypeFromValue returns a pointer to a valid AlertsCorrelationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertsCorrelationTypeFromValue(v string) (*AlertsCorrelationType, error) {
	ev := AlertsCorrelationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertsCorrelationType: valid values are %v", v, allowedCorrelationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertsCorrelationType) IsValid() bool {
	for _, existing := range allowedCorrelationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to AlertsCorrelationType value.
func (v AlertsCorrelationType) Ptr() *AlertsCorrelationType {
	return &v
}
