package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type ReportsReportQueryData struct {
	ItemType        *string                     `json:"ItemType,omitempty"`
	QueryID         common.NullableString       `json:"QueryID,omitempty"`
	QueryStr        common.NullableString       `json:"QueryStr,omitempty"`
	Code            *string                     `json:"Code,omitempty"`
	MaxRowCount     *int64                      `json:"MaxRowCount,omitempty"`
	DateTimeRange   *DateTimeRange              `json:"DateTimeRange,omitempty"`
	ScriptPath      common.NullableString       `json:"ScriptPath,omitempty"`
	ScriptArguments common.NullableList[string] `json:"ScriptArguments,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
