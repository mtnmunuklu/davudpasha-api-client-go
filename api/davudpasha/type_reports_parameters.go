package davudpasha

import "encoding/json"

type ReportsParameters struct {
	Parameters []string         `json:"Parameters,omitempty"`
	Datas      *json.RawMessage `json:"Datas,omitempty"`
	IsActive   *bool            `json:"IsActive,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
