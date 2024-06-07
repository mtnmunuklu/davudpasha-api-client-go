package davudpasha

type CasesSearchRequest struct {
	FromIndex     *int64         `json:"fromIndex,omitempty"`
	PageSize      *int64         `json:"pageSize,omitempty"`
	DateTimeRange *DateTimeRange `json:"dateTimeRange,omitempty"`
	Filter        *string        `json:"filter,omitempty"`
	App           *string        `json:"app,omitempty"`
	ShowSubCases  *bool          `json:"showSubCases"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
