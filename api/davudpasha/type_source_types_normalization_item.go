package davudpasha

type SourceTypesNormalizationItem struct {
	Key   *string `json:"Key,omitempty"`
	Value *string `json:"Value,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
