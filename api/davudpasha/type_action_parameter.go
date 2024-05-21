package davudpasha

type ActionParameter struct {
	Key   *string `json:"Key,omitempty"`
	Value *string `json:"Value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}
