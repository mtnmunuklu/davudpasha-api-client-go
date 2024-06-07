package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type CasesComment struct {
	DetailID *string                     `json:"DetailID,omitempty"`
	Time     *string                     `json:"Time,omitempty"`
	Source   *string                     `json:"Source,omitempty"`
	Message  *string                     `json:"Message,omitempty"`
	Type     *string                     `json:"Type,omitempty"`
	Files    common.NullableList[string] `json:"Files,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
