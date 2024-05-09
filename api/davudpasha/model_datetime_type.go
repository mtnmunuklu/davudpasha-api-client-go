package davudpasha

import (
	"encoding/json"
	"fmt"
)

// DateTimeType represents the type of date time range for events.
type DateTimeType string

// Allowed values for DateTimeType.
const (
	DATETIMETYPE_LAST5MINUTES   DateTimeType = "Last5Minutes"
	DATETIMETYPE_LAST10MINUTES  DateTimeType = "Last10Minutes"
	DATETIMETYPE_LAST15MINUTES  DateTimeType = "Last15Minutes"
	DATETIMETYPE_LAST30MINUTES  DateTimeType = "Last30Minutes"
	DATETIMETYPE_LAST60MINUTES  DateTimeType = "Last60Minutes"
	DATETIMETYPE_LAST120MINUTES DateTimeType = "Last120Minutes"
	DATETIMETYPE_THISHOUR       DateTimeType = "ThisHour"
	DATETIMETYPE_LAST6HOUR      DateTimeType = "Last6Hours"
	DATETIMETYPE_LAST12HOUR     DateTimeType = "Last12Hours"
	DATETIMETYPE_LAST48HOUR     DateTimeType = "Last48Hours"
	DATETIMETYPE_LAST72HOUR     DateTimeType = "Last72Hours"
	DATETIMETYPE_TODAY          DateTimeType = "Today"
	DATETIMETYPE_YESTERDAY      DateTimeType = "Yesterday"
	DATETIMETYPE_THISWEEK       DateTimeType = "ThisWeek"
	DATETIMETYPE_PREVWEEK       DateTimeType = "PrevWeek"
	DATETIMETYPE_THISMONTH      DateTimeType = "ThisMonth"
	DATETIMETYPE_LASTMONTH      DateTimeType = "LastMonth"
	DATETIMETYPE_ALLTIME        DateTimeType = "AllTime"
	DATETIMETYPE_DATERANGE      DateTimeType = "DateRange"
)

// allowedEventsDataTimeEnumValues contains the allowed values for DateTimeType.
var allowedEventsDataTimeEnumValues = []DateTimeType{
	DATETIMETYPE_LAST5MINUTES,
	DATETIMETYPE_LAST10MINUTES,
	DATETIMETYPE_LAST15MINUTES,
	DATETIMETYPE_LAST30MINUTES,
	DATETIMETYPE_LAST60MINUTES,
	DATETIMETYPE_LAST120MINUTES,
	DATETIMETYPE_THISHOUR,
	DATETIMETYPE_LAST6HOUR,
	DATETIMETYPE_LAST12HOUR,
	DATETIMETYPE_LAST48HOUR,
	DATETIMETYPE_LAST72HOUR,
	DATETIMETYPE_TODAY,
	DATETIMETYPE_YESTERDAY,
	DATETIMETYPE_THISWEEK,
	DATETIMETYPE_PREVWEEK,
	DATETIMETYPE_THISMONTH,
	DATETIMETYPE_LASTMONTH,
	DATETIMETYPE_ALLTIME,
	DATETIMETYPE_DATERANGE,
}

// GetAllowedValues returns the list of possible values.
func (v *DateTimeType) GetAllowedValues() []DateTimeType {
	return allowedEventsDataTimeEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *DateTimeType) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DateTimeType(value)
	return nil
}

// NewDateTimeTypeFromValue returns a pointer to a valid DateTimeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDateTimeTypeFromValue(v string) (*DateTimeType, error) {
	ev := DateTimeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DateTimeType: valid values are %v", v, allowedEventsDataTimeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DateTimeType) IsValid() bool {
	for _, existing := range allowedEventsDataTimeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to DateTimeType value.
func (v DateTimeType) Ptr() *DateTimeType {
	return &v
}
