package davudpasha

import (
	"encoding/json"
	"fmt"
)

// QueriesQueryType represents the type of date time range for events.
type QueriesQueryType string

// Allowed values for QueriesQueryType.
const (
	QUERYTYPE_NORMAL QueriesQueryType = "Normal"
)

// allowedQueryEnumValues contains the allowed values for QueriesQueryType.
var allowedQueryEnumValues = []QueriesQueryType{
	QUERYTYPE_NORMAL,
}

// GetAllowedValues returns the list of possible values.
func (v *QueriesQueryType) GetAllowedValues() []QueriesQueryType {
	return allowedQueryEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *QueriesQueryType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueriesQueryType(value)
	return nil
}

// NewQueriesQueryTypeFromValue returns a pointer to a valid QueriesQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueriesQueryTypeFromValue(v string) (*QueriesQueryType, error) {
	ev := QueriesQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueriesQueryType: valid values are %v", v, allowedQueryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueriesQueryType) IsValid() bool {
	for _, existing := range allowedQueryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to QueriesQueryType value.
func (v QueriesQueryType) Ptr() *QueriesQueryType {
	return &v
}
