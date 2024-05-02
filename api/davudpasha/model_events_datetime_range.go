package davudpasha

type EventsDateTimeRange struct {
	DateTimeType *string `json:"DateTimeType,omitempty"`
	StartDate    *string `json:"StartDate,omitempty"`
	EndDate      *string `json:"EndDate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}
