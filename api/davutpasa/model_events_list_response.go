package davutpasa

import "encoding/json"

// EventsListResponse The response object with all events matching the request.
type EventsListResponse struct {
	Data            []json.RawMessage `json:"Data,omitempty"`
	TotalSize       *int64            `json:"TotalSize,omitempty"`
	QueryType       *string           `json:"QueryType,omitempty"`
	QueryFlag       *string           `json:"QueryFlag,omitempty"`
	SelectedColumns []json.RawMessage `json:"SelectedColumns,omitempty"`
	SearchTime      *int64            `json:"SearchTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}
