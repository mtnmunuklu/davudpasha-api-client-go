package davudpasha

type ReportsChartVisualization struct {
	ChartType   *string  `json:"ChartType,omitempty"`
	SerieColors []string `json:"SerieColors,omitempty"`
	LineWidth   *int64   `json:"LineWidth,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
