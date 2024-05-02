package davudpasha

import (
	"encoding/json"
	"fmt"
)

type EventsDateTimeType string

const (
	EVENTSDATETIMETYPE_LAST5MINUTES   EventsDateTimeType = "Last5Minutes"
	EVENTSDATETIMETYPE_LAST10MINUTES  EventsDateTimeType = "Last10Minutes"
	EVENTSDATETIMETYPE_LAST15MINUTES  EventsDateTimeType = "Last15Minutes"
	EVENTSDATETIMETYPE_LAST30MINUTES  EventsDateTimeType = "Last30Minutes"
	EVENTSDATETIMETYPE_LAST60MINUTES  EventsDateTimeType = "Last60Minutes"
	EVENTSDATETIMETYPE_LAST120MINUTES EventsDateTimeType = "Last120Minutes"
	EVENTSDATETIMETYPE_THISHOUR       EventsDateTimeType = "ThisHour"
	EVENTSDATETIMETYPE_LAST6HOUR      EventsDateTimeType = "Last6Hours"
	EVENTSDATETIMETYPE_LAST12HOUR     EventsDateTimeType = "Last12Hours"
	EVENTSDATETIMETYPE_LAST48HOUR     EventsDateTimeType = "Last48Hours"
	EVENTSDATETIMETYPE_LAST72HOUR     EventsDateTimeType = "Last72Hours"
	EVENTSDATETIMETYPE_TODAY          EventsDateTimeType = "Today"
	EVENTSDATETIMETYPE_YESTERDAY      EventsDateTimeType = "Yesterday"
	EVENTSDATETIMETYPE_THISWEEK       EventsDateTimeType = "ThisWeek"
	EVENTSDATETIMETYPE_PREVWEEK       EventsDateTimeType = "PrevWeek"
	EVENTSDATETIMETYPE_THISMONTH      EventsDateTimeType = "ThisMonth"
	EVENTSDATETIMETYPE_LASTMONTH      EventsDateTimeType = "LastMonth"
	EVENTSDATETIMETYPE_ALLTIME        EventsDateTimeType = "AllTime"
	EVENTSDATETIMETYPE_DATERANGE      EventsDateTimeType = "DateRange"
)

var allowedEventsDataTimeEnumValues = []EventsDateTimeType{
	EVENTSDATETIMETYPE_LAST5MINUTES,
	EVENTSDATETIMETYPE_LAST10MINUTES,
	EVENTSDATETIMETYPE_LAST15MINUTES,
	EVENTSDATETIMETYPE_LAST30MINUTES,
	EVENTSDATETIMETYPE_LAST60MINUTES,
	EVENTSDATETIMETYPE_LAST120MINUTES,
	EVENTSDATETIMETYPE_THISHOUR,
	EVENTSDATETIMETYPE_LAST6HOUR,
	EVENTSDATETIMETYPE_LAST12HOUR,
	EVENTSDATETIMETYPE_LAST48HOUR,
	EVENTSDATETIMETYPE_LAST72HOUR,
	EVENTSDATETIMETYPE_TODAY,
	EVENTSDATETIMETYPE_YESTERDAY,
	EVENTSDATETIMETYPE_THISWEEK,
	EVENTSDATETIMETYPE_PREVWEEK,
	EVENTSDATETIMETYPE_THISMONTH,
	EVENTSDATETIMETYPE_LASTMONTH,
	EVENTSDATETIMETYPE_ALLTIME,
	EVENTSDATETIMETYPE_DATERANGE,
}

// GetAllowedValues returns the list of possible values.
func (v *EventsDateTimeType) GetAllowedValues() []EventsDateTimeType {
	return allowedEventsDataTimeEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *EventsDateTimeType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsDateTimeType(value)
	return nil
}

// NewEventsDateTimeTypeFromValue returns a pointer to a valid EventsDateTimeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsDateTimeTypeFromValue(v string) (*EventsDateTimeType, error) {
	ev := EventsDateTimeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsDateTimeType: valid values are %v", v, allowedEventsDataTimeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsDateTimeType) IsValid() bool {
	for _, existing := range allowedEventsDataTimeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to EventsDateTimeType value.
func (v EventsDateTimeType) Ptr() *EventsDateTimeType {
	return &v
}
