package davudpasha

import (
	"encoding/json"
	"fmt"
)

// ReportsReportType represents the type of date time range for events.
type ReportsReportType string

// Allowed values for ReportsReportType.
const (
	REPORTTYPE_QUERY           ReportsReportType = "Query"
	REPORTTYPE_PYTHONSCRIPT    ReportsReportType = "PythonScript"
	REPORTTYPE_SEARCHARCHIVE   ReportsReportType = "SearchArchive"
	REPORTTYPE_NONEREPUDIATION ReportsReportType = "NonRepudiation"
	REPORTTYPE_TIMESTAMP       ReportsReportType = "Timestamp"
	REPORTTYPE_CODE            ReportsReportType = "Code"
)

// allowedReportEnumValues contains the allowed values for ReportsReportType.
var allowedReportEnumValues = []ReportsReportType{
	REPORTTYPE_QUERY,
	REPORTTYPE_PYTHONSCRIPT,
	REPORTTYPE_SEARCHARCHIVE,
	REPORTTYPE_NONEREPUDIATION,
	REPORTTYPE_TIMESTAMP,
	REPORTTYPE_CODE,
}

// GetAllowedValues returns the list of possible values.
func (v *ReportsReportType) GetAllowedValues() []ReportsReportType {
	return allowedReportEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *ReportsReportType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportsReportType(value)
	return nil
}

// NewReportsReportTypeFromValue returns a pointer to a valid ReportsReportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportsReportTypeFromValue(v string) (*ReportsReportType, error) {
	ev := ReportsReportType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportsReportType: valid values are %v", v, allowedReportEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportsReportType) IsValid() bool {
	for _, existing := range allowedReportEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to ReportsReportType value.
func (v ReportsReportType) Ptr() *ReportsReportType {
	return &v
}
