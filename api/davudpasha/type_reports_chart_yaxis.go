package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type ReportsChartYAxis struct {
	Label    *string              `json:"Label,omitempty"`
	Interval *int64               `json:"Interval,omitempty"`
	MinValue common.NullableInt64 `json:"MinValue,omitempty"`
	MaxValue common.NullableInt64 `json:"MaxValue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
