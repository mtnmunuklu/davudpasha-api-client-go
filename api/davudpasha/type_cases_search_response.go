package davudpasha

type CasesSearchResponse struct {
	Items      []CasesItem `json:"Items,omitempty"`
	TotalCount *int64      `json:"TotalCount,omitempty"`
	PageSize   *int64      `json:"PageSize,omitempty"`
	FromIndex  *int64      `json:"FromIndex,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
