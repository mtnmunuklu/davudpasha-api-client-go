package davudpasha

type SourceTypesClassificationRule struct {
	QueryStr         *string `json:"QueryStr,omitempty"`
	ClassificationID *string `json:"ClassificationID,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}
