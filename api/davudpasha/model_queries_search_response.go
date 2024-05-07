package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// QueriesSearchResponse is the response object with all queries matching the request.
type QueriesSearchResponse struct {
	Items        common.NullableList[json.RawMessage] `json:"Items,omitempty"`
	FailedItems  common.NullableList[json.RawMessage] `json:"FailedItems,omitempty"`
	SuccessItems common.NullableList[json.RawMessage] `json:"SuccessItems,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
