package davudpasha

type SourcesScheduleConfig struct {
	ScheduleFullDateTime *string  `json:"ScheduleFullDateTime,omitempty"`
	ScheduleType         *string  `json:"ScheduleType,omitempty"`
	TimeTics             *int64   `json:"TimeTics,omitempty"`
	Time                 *string  `json:"Time,omitempty"`
	DateStr              *string  `json:"DateStr,omitempty"`
	Days                 []string `json:"Days,omitempty"`
	DaysNumber           []int64  `json:"DaysNumber,omitempty"`
	DayNo                *int64   `json:"DayNo,omitempty"`
	Day                  *string  `json:"Day,omitempty"`
	DayNumber            *int64   `json:"DayNumber,omitempty"`
	WeekType             *string  `json:"WeekType,omitempty"`
	TimeType             *string  `json:"TimeType,omitempty"`
	TimeValue            *int64   `json:"TimeValue,omitempty"`
}
