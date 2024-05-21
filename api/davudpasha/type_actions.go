package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type Actions struct {
	Actions      []Action                    `json:"Actions,omitempty"`
	ActionRefIds common.NullableList[string] `json:"ActionRefIds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}
