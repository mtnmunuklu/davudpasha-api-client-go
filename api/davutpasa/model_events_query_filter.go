package davutpasa

// EventsQueryFilter The search and filter query settings
type EventsQueryFilter struct {
	QuerySQL        *string              `json:"QuerySQL,omitempty"`
	DateTimeRange   *EventsDateTimeRange `json:"DateTimeRange,omitempty"`
	SearchAfterKeys []string             `json:"SearchAfterKeys,omitempty"`
	Size            *int64               `json:"Size,omitempty"`
	QueryOptions    *EventsQueryOptions  `json:"QueryOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}
