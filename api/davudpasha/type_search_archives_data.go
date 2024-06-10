package davudpasha

type SearchArchivesData struct {
	Query  *SearchArchivesQuery  `json:"Query,omitempty"`
	Status *SearchArchivesStatus `json:"Status,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
