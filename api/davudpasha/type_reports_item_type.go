package davudpasha

import (
	"encoding/json"
	"fmt"
)

// ReportsItemType represents the type of date time range for events.
type ReportsItemType string

// Allowed values for ReportsItemType.
const (
	ITEMTYPE_QUERY        ReportsItemType = "Query"
	ITEMTYPE_PYTHONSCRIPT ReportsItemType = "PythonScript"
	ITEMTYPE_CUSTOM       ReportsItemType = "Custom"
	ITEMTYPE_CODE         ReportsItemType = "Code"
)

// allowedItemEnumValues contains the allowed values for ReportsItemType.
var allowedItemEnumValues = []ReportsItemType{
	ITEMTYPE_QUERY,
	ITEMTYPE_PYTHONSCRIPT,
	ITEMTYPE_CUSTOM,
	ITEMTYPE_CODE,
}

// GetAllowedValues returns the list of possible values.
func (v *ReportsItemType) GetAllowedValues() []ReportsItemType {
	return allowedItemEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *ReportsItemType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportsItemType(value)
	return nil
}

// NewReportsItemTypeFromValue returns a pointer to a valid ReportsItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportsItemTypeFromValue(v string) (*ReportsItemType, error) {
	ev := ReportsItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportsItemType: valid values are %v", v, allowedItemEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportsItemType) IsValid() bool {
	for _, existing := range allowedItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to ReportsItemType value.
func (v ReportsItemType) Ptr() *ReportsItemType {
	return &v
}
