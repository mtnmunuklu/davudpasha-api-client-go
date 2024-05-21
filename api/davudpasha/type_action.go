package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type Action struct {
	ActionType       *string               `json:"ActionType,omitempty"`
	ActionParameters []ActionParameter     `json:"ActionParameters,omitempty"`
	Data             common.NullableString `json:"Data,omitempty"`
	ActionRefId      common.NullableString `json:"ActionRefId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}
