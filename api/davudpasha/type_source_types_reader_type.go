package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourceTypesReaderType represents the type of date time range for events.
type SourceTypesReaderType string

// Allowed values for SourceTypesReaderType.
const (
	READERTYPE_FILE     SourceTypesReaderType = "FILE"
	READERTYPE_NOPARSER SourceTypesReaderType = "NOPARSER"
	READERTYPE_DB       SourceTypesReaderType = "DB"
	READERTYPE_EVENTLOG SourceTypesReaderType = "EVENTLOG"
)

// allowedReaderEnumValues contains the allowed values for SourceTypesReaderType.
var allowedReaderEnumValues = []SourceTypesReaderType{
	READERTYPE_FILE,
	READERTYPE_NOPARSER,
	READERTYPE_DB,
	READERTYPE_EVENTLOG,
}

// GetAllowedValues returns the list of possible values.
func (v *SourceTypesReaderType) GetAllowedValues() []SourceTypesReaderType {
	return allowedReaderEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourceTypesReaderType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourceTypesReaderType(value)
	return nil
}

// NewSourceTypesReaderTypeFromValue returns a pointer to a valid SourceTypesReaderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourceTypesReaderTypeFromValue(v string) (*SourceTypesReaderType, error) {
	ev := SourceTypesReaderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourceTypesReaderType: valid values are %v", v, allowedReaderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourceTypesReaderType) IsValid() bool {
	for _, existing := range allowedReaderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourceTypesReaderType value.
func (v SourceTypesReaderType) Ptr() *SourceTypesReaderType {
	return &v
}
