package davudpasha

// QueriesSearchRequest is the object sent with a request to retrieve a list of queries from your organization.
type QueriesSearchRequest struct {
	Filter *string `json:"filter,omitempty"`
	// SmartRestRequestContext is the context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
