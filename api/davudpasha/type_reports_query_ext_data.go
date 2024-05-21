package davudpasha

type ReportsQueryExtData struct {
	ReportType               *string        `json:"ReportType,omitempty"`
	DateTimeRange            *DateTimeRange `json:"DateTimeRange,omitempty"`
	Alerts                   *bool          `json:"Alerts,omitempty"`
	LogPositions             *bool          `json:"LogPositions,omitempty"`
	SystemInfo               *bool          `json:"SystemInfo,omitempty"`
	LogCounts                *bool          `json:"LogCounts,omitempty"`
	Classifications          *bool          `json:"Classifications,omitempty"`
	Correlation              *bool          `json:"Correlation,omitempty"`
	TimeStamps               *bool          `json:"TimeStamps,omitempty"`
	LogCounts_Chart          *bool          `json:"LogCounts_Chart,omitempty"`
	LogCounts_Table          *bool          `json:"LogCounts_Table,omitempty"`
	Classifications_Chart    *bool          `json:"Classifications_Chart,omitempty"`
	Classifications_SvrChart *bool          `json:"Classifications_SvrChart,omitempty"`
	Classifications_Table    *bool          `json:"Classifications_Table,omitempty"`
	Correlation_AlertChart   *bool          `json:"Correlation_AlertChart,omitempty"`
	Correlation_SvrChart     *bool          `json:"Correlation_SvrChart,omitempty"`
	Correlation_Table        *bool          `json:"Correlation_Table,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
