package davudpasha

type ReportsTableVisualization struct {
	ChartType   *string  `json:"ChartType,omitempty"`
	MaxRowCount *string  `json:"MaxRowCount,omitempty"`
	Columns     []string `json:"Columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
