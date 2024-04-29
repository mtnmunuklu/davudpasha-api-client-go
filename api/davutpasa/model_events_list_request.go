package davutpasa

// EventListRequest The object sent with request to retrieve a list of event from your organization.
type EventsListRequest struct {
	Reason *string `json:"reason,omitempty"`
	// The search and filter query settings.
	Query                   *EventsQueryFilter `json:"query,omitempty"`
	SmartRestRequestContext *string            `json:"smartRestRequestContext,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObjects      map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}
