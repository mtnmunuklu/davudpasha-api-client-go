package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourceTypesParserType represents the type of date time range for events.
type SourceTypesParserType string

// Allowed values for SourceTypesParserType.
const (
	PARSERTYPE_REGEX     SourceTypesParserType = "Module csiem.parser.regex"
	PARSERTYPE_DELIMITER SourceTypesParserType = "Module csiem.parser.delimiter"
	PARSERTYPE_CSV       SourceTypesParserType = "Module csiem.parser.csv"
	PARSERTYPE_CODE      SourceTypesParserType = "Module csiem.parser.code"
	PARSERTYPE_STATIC    SourceTypesParserType = "Module csiem.parser.static"
	PARSERTYPE_DB        SourceTypesParserType = "Module csiem.parser.db"
	PARSERTYPE_NONE      SourceTypesParserType = "Module csiem.parser.none"
)

// allowedParserEnumValues contains the allowed values for SourceTypesParserType.
var allowedParserEnumValues = []SourceTypesParserType{
	PARSERTYPE_REGEX,
	PARSERTYPE_DELIMITER,
	PARSERTYPE_CSV,
	PARSERTYPE_CODE,
	PARSERTYPE_STATIC,
	PARSERTYPE_DB,
	PARSERTYPE_NONE,
}

// GetAllowedValues returns the list of possible values.
func (v *SourceTypesParserType) GetAllowedValues() []SourceTypesParserType {
	return allowedParserEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourceTypesParserType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourceTypesParserType(value)
	return nil
}

// NewSourceTypesParserTypeFromValue returns a pointer to a valid SourceTypesParserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourceTypesParserTypeFromValue(v string) (*SourceTypesParserType, error) {
	ev := SourceTypesParserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourceTypesParserType: valid values are %v", v, allowedParserEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourceTypesParserType) IsValid() bool {
	for _, existing := range allowedParserEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourceTypesParserType value.
func (v SourceTypesParserType) Ptr() *SourceTypesParserType {
	return &v
}
