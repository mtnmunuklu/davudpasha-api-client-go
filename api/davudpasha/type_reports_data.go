package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type ReportsData struct {
	Name common.NullableString `json:"Name,"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
