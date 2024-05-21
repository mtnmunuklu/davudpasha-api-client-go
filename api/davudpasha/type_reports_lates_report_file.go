package davudpasha

type ReportsLatestReportFile struct {
	Filename           *string `json:"Filename,omitempty"`
	Status             *string `json:"Status,omitempty"`
	ReportGridFSFileId *string `json:"ReportGridFSFileId,omitempty"`
	RunDate            *string `json:"RunDate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
