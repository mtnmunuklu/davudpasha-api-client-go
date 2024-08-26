package davudpasha

import (
	"encoding/json"
	"fmt"
)

// SourcesLogReaderType represents the type of date time range for events.
type SourcesLogReaderType string

// Allowed values for SourcesLogReaderType.
const (
	LOGREADERTYPE_HTTPCOLLECTOR SourcesLogReaderType = "Module csiem.reader.http"
	LOGREADERTYPE_EVENTLOG      SourcesLogReaderType = "Module csiem.reader.eventlog"
	LOGREADERTYPE_SSH           SourcesLogReaderType = "Module csiem.reader.ssh"
	LOGREADERTYPE_WEF           SourcesLogReaderType = "Module csiem.reader.eventlogfile"
	LOGREADERTYPE_DATABASE      SourcesLogReaderType = "Module csiem.reader.db"
	LOGREADERTYPE_INDEX         SourcesLogReaderType = "Module csiem.reader.index"
	LOGREADERTYPE_OFFICE365     SourcesLogReaderType = "Module csiem.reader.office365web"
	LOGREADERTYPE_KAFKA         SourcesLogReaderType = "Module csiem.reader.kafka"
	LOGREADERTYPE_FTP           SourcesLogReaderType = "Module csiem.reader.ftp"
	LOGREADERTYPE_FILE          SourcesLogReaderType = "Module csiem.reader.file"
	LOGREADERTYPE_MSGRAPH       SourcesLogReaderType = "Module csiem.reader.msgraph"
	LOGREADERTYPE_MSB           SourcesLogReaderType = "Module csiem.reader.msb"
	LOGREADERTYPE_FIXEDLENGTH   SourcesLogReaderType = "Module csiem.reader.static"
	LOGREADERTYPE_DHCP          SourcesLogReaderType = "Module csiem.reader.dhcp"
)

// allowedLogReaderEnumValues contains the allowed values for SourcesLogReaderType.
var allowedLogReaderEnumValues = []SourcesLogReaderType{
	LOGREADERTYPE_HTTPCOLLECTOR,
	LOGREADERTYPE_EVENTLOG,
	LOGREADERTYPE_SSH,
	LOGREADERTYPE_WEF,
	LOGREADERTYPE_DATABASE,
	LOGREADERTYPE_INDEX,
	LOGREADERTYPE_OFFICE365,
	LOGREADERTYPE_KAFKA,
	LOGREADERTYPE_FTP,
	LOGREADERTYPE_FILE,
	LOGREADERTYPE_MSGRAPH,
	LOGREADERTYPE_MSB,
	LOGREADERTYPE_DHCP,
}

// GetAllowedValues returns the list of possible values.
func (v *SourcesLogReaderType) GetAllowedValues() []SourcesLogReaderType {
	return allowedLogReaderEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *SourcesLogReaderType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcesLogReaderType(value)
	return nil
}

// NewSourcesLogReaderTypeFromValue returns a pointer to a valid SourcesLogReaderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcesLogReaderTypeFromValue(v string) (*SourcesLogReaderType, error) {
	ev := SourcesLogReaderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcesLogReaderType: valid values are %v", v, allowedLogReaderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcesLogReaderType) IsValid() bool {
	for _, existing := range allowedLogReaderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to SourcesLogReaderType value.
func (v SourcesLogReaderType) Ptr() *SourcesLogReaderType {
	return &v
}
