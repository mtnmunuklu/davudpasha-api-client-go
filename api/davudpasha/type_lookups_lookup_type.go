package davudpasha

import (
	"encoding/json"
	"fmt"
)

// LookupsLookupType represents the type of date time range for events.
type LookupsLookupType string

// Allowed values for LookupsLookupType.
const (
	LOOKUPTYPE_IPSUBNET        LookupsLookupType = "IP_Subnet"
	LOOKUPTYPE_FILEJSON        LookupsLookupType = "File_JSON"
	LOOKUPTYPE_FILEJSV         LookupsLookupType = "File_CSV"
	LOOKUPTYPE_DATABASE        LookupsLookupType = "Database"
	LOOKUPTYPE_STATIC          LookupsLookupType = "Static"
	LOOKUPTYPE_BLACKLIST       LookupsLookupType = "Blacklist"
	LOOKUPTYPE_STATICURL       LookupsLookupType = "Static_URL"
	LOOKUPTYPE_ACTIVEDIRECTORY LookupsLookupType = "ActiveDirectory"
	LOOKUPTYPE_ACTIVELIST      LookupsLookupType = "Active_List"
)

// allowedLookupEnumValues contains the allowed values for LookupsLookupType.
var allowedLookupEnumValues = []LookupsLookupType{
	LOOKUPTYPE_IPSUBNET,
	LOOKUPTYPE_FILEJSON,
	LOOKUPTYPE_FILEJSV,
	LOOKUPTYPE_DATABASE,
	LOOKUPTYPE_BLACKLIST,
	LOOKUPTYPE_ACTIVEDIRECTORY,
	LOOKUPTYPE_ACTIVELIST,
}

// GetAllowedValues returns the list of possible values.
func (v *LookupsLookupType) GetAllowedValues() []LookupsLookupType {
	return allowedLookupEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *LookupsLookupType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LookupsLookupType(value)
	return nil
}

// NewLookupsLookupTypeFromValue returns a pointer to a valid LookupsLookupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLookupsLookupTypeFromValue(v string) (*LookupsLookupType, error) {
	ev := LookupsLookupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LookupsLookupType: valid values are %v", v, allowedLookupEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LookupsLookupType) IsValid() bool {
	for _, existing := range allowedLookupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to LookupsLookupType value.
func (v LookupsLookupType) Ptr() *LookupsLookupType {
	return &v
}
