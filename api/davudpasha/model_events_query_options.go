package davudpasha

// EventsQueryOptions The global query options tha are used.
type EventsQueryOptions struct {
	ShowHighlight *bool `json:"ShowHighlight,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}
