package davudpasha

type ReportsChartXAxis struct {
	Label    *string `json:"Label,omitempty"`
	Interval *int64  `json:"Interval,omitempty"`
	Angle    *int64  `json:"Angle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
