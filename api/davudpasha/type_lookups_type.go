package davudpasha

import (
	"encoding/json"
	"fmt"
)

// LookupType represents the type of date time range for events.
type LookupType string

// Allowed values for LookupType.
const (
	LookupType_IPSUBNET        LookupType = "IP_Subnet"
	LookupType_FILEJSON        LookupType = "File_JSON"
	LookupType_FILEJSV         LookupType = "File_CSV"
	LookupType_DATABASE        LookupType = "Database"
	LookupType_STATIC          LookupType = "Static"
	LookupType_BLACKLIST       LookupType = "Blacklist"
	LookupType_STATICURL       LookupType = "Static_URL"
	LookupType_ACTIVEDIRECTORY LookupType = "ActiveDirectory"
	LookupType_ACTIVELIST      LookupType = "Active_List"
)

// allowedLookupEnumValues contains the allowed values for LookupType.
var allowedLookupEnumValues = []LookupType{
	LookupType_IPSUBNET,
	LookupType_FILEJSON,
	LookupType_FILEJSV,
	LookupType_DATABASE,
	LookupType_BLACKLIST,
	LookupType_ACTIVEDIRECTORY,
	LookupType_ACTIVELIST,
}

// GetAllowedValues returns the list of possible values.
func (v *LookupType) GetAllowedValues() []LookupType {
	return allowedLookupEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *LookupType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LookupType(value)
	return nil
}

// NewLookupTypeFromValue returns a pointer to a valid LookupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLookupTypeFromValue(v string) (*LookupType, error) {
	ev := LookupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LookupType: valid values are %v", v, allowedLookupEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LookupType) IsValid() bool {
	for _, existing := range allowedLookupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to LookupType value.
func (v LookupType) Ptr() *LookupType {
	return &v
}
