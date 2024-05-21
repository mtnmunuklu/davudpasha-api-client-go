package davudpasha

type ReportsQuery struct {
	Name               *string                    `json:"Nane,omitempty"`
	Data               *ReportsQueryData          `json:"Data,omitempty"`
	ShowTable          *bool                      `json:"ShowTable,omitempty"`
	TableVisualization *ReportsTableVisualization `json:"TableVisualization,omitempty"`
	ShowChart          *bool                      `json:"ShowChart,omitempty"`
	ChartVisualization *ReportsChartVisualization `json:"ChartVisualization,omitempty"`
	ExtData            *ReportsQueryExtData       `json:"ExtData,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
