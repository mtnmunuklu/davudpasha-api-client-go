package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourcesDataFormat represents the type of date time range for events.
type SourcesDataFormat string

// Allowed values for SourcesDataFormat.
const (
	DATAFORMAT_TEXT SourcesDataFormat = "Text"
)

// allowedDataEnumValues contains the allowed values for SourcesDataFormat.
var allowedDataEnumValues = []SourcesDataFormat{
	DATAFORMAT_TEXT,
}

// GetAllowedValues returns the list of possible values.
func (v *SourcesDataFormat) GetAllowedValues() []SourcesDataFormat {
	return allowedDataEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourcesDataFormat) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcesDataFormat(value)
	return nil
}

// NewSourcesDataFormatFromValue returns a pointer to a valid SourcesDataFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcesDataFormatFromValue(v string) (*SourcesDataFormat, error) {
	ev := SourcesDataFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcesDataFormat: valid values are %v", v, allowedDataEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcesDataFormat) IsValid() bool {
	for _, existing := range allowedDataEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourcesDataFormat value.
func (v SourcesDataFormat) Ptr() *SourcesDataFormat {
	return &v
}
