package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ScheduleConfig represents the configuration for scheduling tasks.
type ScheduleConfig struct {
	// Full date and time of the schedule.
	ScheduleFullDateTime *string `json:"ScheduleFullDateTime,omitempty"`
	// Type of the schedule.
	ScheduleType *string `json:"ScheduleType,omitempty"`
	// Time ticks.
	TimeTics *int64 `json:"TimeTics,omitempty"`
	// Time for the schedule.
	Time *string `json:"Time,omitempty"`
	// Date string for the schedule.
	DateStr *string `json:"DateStr,omitempty"`
	// Days for the schedule.
	Days []string `json:"Days,omitempty"`
	// Numbers representing days for the schedule.
	DaysNumber []int64 `json:"DaysNumber,omitempty"`
	// Number of the day.
	DayNo *int64 `json:"DayNo,omitempty"`
	// Day of the schedule.
	Day *string `json:"Day,omitempty"`
	// Number representing the day.
	DayNumber *int64 `json:"DayNumber,omitempty"`
	// Type of the week for the schedule.
	WeekType *string `json:"WeekType,omitempty"`
	// Type of time for the schedule.
	TimeType *string `json:"TimeType,omitempty"`
	// Value representing time for the schedule.
	TimeValue *int64 `json:"TimeValue,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewScheduleConfig creates a new ScheduleConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleConfig() *ScheduleConfig {
	this := ScheduleConfig{}
	return &this
}

// NewScheduleConfigWithDefaults creates a new ScheduleConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleConfigWithDefaults() *ScheduleConfig {
	this := ScheduleConfig{}
	return &this
}

// GetScheduleFullDateTime returns the ScheduleFullDateTime field value if set, zero value otherwise.
func (o *ScheduleConfig) GetScheduleFullDateTime() string {
	if o == nil || o.ScheduleFullDateTime == nil {
		var ret string
		return ret
	}
	return *o.ScheduleFullDateTime
}

// GetScheduleFullDateTimeOk returns a tuple with the ScheduleFullDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetScheduleFullDateTimeOk() (*string, bool) {
	if o == nil || o.ScheduleFullDateTime == nil {
		return nil, false
	}
	return o.ScheduleFullDateTime, true
}

// HasScheduleFullDateTime returns a boolean if a field has been set.
func (o *ScheduleConfig) HasScheduleFullDateTime() bool {
	return o != nil && o.ScheduleFullDateTime != nil
}

// SetScheduleFullDateTime gets a reference to the given string and assigns it to the ScheduleFullDateTime field.
func (o *ScheduleConfig) SetScheduleFullDateTime(v string) {
	o.ScheduleFullDateTime = &v
}

// GetScheduleType returns the ScheduleType field value if set, zero value otherwise.
func (o *ScheduleConfig) GetScheduleType() string {
	if o == nil || o.ScheduleType == nil {
		var ret string
		return ret
	}
	return *o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetScheduleTypeOk() (*string, bool) {
	if o == nil || o.ScheduleType == nil {
		return nil, false
	}
	return o.ScheduleType, true
}

// HasScheduleType returns a boolean if a field has been set.
func (o *ScheduleConfig) HasScheduleType() bool {
	return o != nil && o.ScheduleType != nil
}

// SetScheduleType gets a reference to the given string and assigns it to the ScheduleType field.
func (o *ScheduleConfig) SetScheduleType(v string) {
	o.ScheduleType = &v
}

// GetTimeTics returns the TimeTics field value if set, zero value otherwise.
func (o *ScheduleConfig) GetTimeTics() int64 {
	if o == nil || o.TimeTics == nil {
		var ret int64
		return ret
	}
	return *o.TimeTics
}

// GetTimeTicsOk returns a tuple with the TimeTics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetTimeTicsOk() (*int64, bool) {
	if o == nil || o.TimeTics == nil {
		return nil, false
	}
	return o.TimeTics, true
}

// HasTimeTics returns a boolean if a field has been set.
func (o *ScheduleConfig) HasTimeTics() bool {
	return o != nil && o.TimeTics != nil
}

// SetTimeTics gets a reference to the given int64 and assigns it to the TimeTics field.
func (o *ScheduleConfig) SetTimeTics(v int64) {
	o.TimeTics = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ScheduleConfig) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ScheduleConfig) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *ScheduleConfig) SetTime(v string) {
	o.Time = &v
}

// GetDateStr returns the DateStr field value if set, zero value otherwise.
func (o *ScheduleConfig) GetDateStr() string {
	if o == nil || o.DateStr == nil {
		var ret string
		return ret
	}
	return *o.DateStr
}

// GetDateStrOk returns a tuple with the DateStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetDateStrOk() (*string, bool) {
	if o == nil || o.DateStr == nil {
		return nil, false
	}
	return o.DateStr, true
}

// HasDateStr returns a boolean if a field has been set.
func (o *ScheduleConfig) HasDateStr() bool {
	return o != nil && o.DateStr != nil
}

// SetDateStr gets a reference to the given string and assigns it to the DateStr field.
func (o *ScheduleConfig) SetDateStr(v string) {
	o.DateStr = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *ScheduleConfig) GetDays() []string {
	if o == nil || o.Days == nil {
		var ret []string
		return ret
	}
	return o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetDaysOk() (*[]string, bool) {
	if o == nil || o.Days == nil {
		return nil, false
	}
	return &o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *ScheduleConfig) HasDays() bool {
	return o != nil && o.Days != nil
}

// SetDays gets a reference to the given []string and assigns it to the Days field.
func (o *ScheduleConfig) SetDays(v []string) {
	o.Days = v
}

// GetDaysNumber returns the DaysNumber field value if set, zero value otherwise.
func (o *ScheduleConfig) GetDaysNumber() []int64 {
	if o == nil || o.DaysNumber == nil {
		var ret []int64
		return ret
	}
	return o.DaysNumber
}

// GetDaysNumberOk returns a tuple with the DaysNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetDaysNumberOk() (*[]int64, bool) {
	if o == nil || o.DaysNumber == nil {
		return nil, false
	}
	return &o.DaysNumber, true
}

// HasDaysNumber returns a boolean if a field has been set.
func (o *ScheduleConfig) HasDaysNumber() bool {
	return o != nil && o.DaysNumber != nil
}

// SetDaysNumber gets a reference to the given []int64 and assigns it to the DaysNumber field.
func (o *ScheduleConfig) SetDaysNumber(v []int64) {
	o.DaysNumber = v
}

// GetDayNo returns the DayNo field value if set, zero value otherwise.
func (o *ScheduleConfig) GetDayNo() int64 {
	if o == nil || o.DayNo == nil {
		var ret int64
		return ret
	}
	return *o.DayNo
}

// GetDayNoOk returns a tuple with the DayNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetDayNoOk() (*int64, bool) {
	if o == nil || o.DayNo == nil {
		return nil, false
	}
	return o.DayNo, true
}

// HasDayNo returns a boolean if a field has been set.
func (o *ScheduleConfig) HasDayNo() bool {
	return o != nil && o.DayNo != nil
}

// SetDayNo gets a reference to the given int64 and assigns it to the DayNo field.
func (o *ScheduleConfig) SetDayNo(v int64) {
	o.DayNo = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *ScheduleConfig) GetDay() string {
	if o == nil || o.Day == nil {
		var ret string
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetDayOk() (*string, bool) {
	if o == nil || o.Day == nil {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *ScheduleConfig) HasDay() bool {
	return o != nil && o.Day != nil
}

// SetDay gets a reference to the given string and assigns it to the Day field.
func (o *ScheduleConfig) SetDay(v string) {
	o.Day = &v
}

// GetDayNumber returns the DayNumber field value if set, zero value otherwise.
func (o *ScheduleConfig) GetDayNumber() int64 {
	if o == nil || o.DayNumber == nil {
		var ret int64
		return ret
	}
	return *o.DayNumber
}

// GetDayNumberOk returns a tuple with the DayNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetDayNumberOk() (*int64, bool) {
	if o == nil || o.DayNumber == nil {
		return nil, false
	}
	return o.DayNumber, true
}

// HasDayNumber returns a boolean if a field has been set.
func (o *ScheduleConfig) HasDayNumber() bool {
	return o != nil && o.DayNumber != nil
}

// SetDayNumber gets a reference to the given int64 and assigns it to the DayNumber field.
func (o *ScheduleConfig) SetDayNumber(v int64) {
	o.DayNumber = &v
}

// GetWeekType returns the WeekType field value if set, zero value otherwise.
func (o *ScheduleConfig) GetWeekType() string {
	if o == nil || o.WeekType == nil {
		var ret string
		return ret
	}
	return *o.WeekType
}

// GetWeekTypeOk returns a tuple with the WeekType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetWeekTypeOk() (*string, bool) {
	if o == nil || o.WeekType == nil {
		return nil, false
	}
	return o.WeekType, true
}

// HasWeekType returns a boolean if a field has been set.
func (o *ScheduleConfig) HasWeekType() bool {
	return o != nil && o.WeekType != nil
}

// SetWeekType gets a reference to the given string and assigns it to the WeekType field.
func (o *ScheduleConfig) SetWeekType(v string) {
	o.WeekType = &v
}

// GetTimeType returns the TimeType field value if set, zero value otherwise.
func (o *ScheduleConfig) GetTimeType() string {
	if o == nil || o.TimeType == nil {
		var ret string
		return ret
	}
	return *o.TimeType
}

// GetTimeTypeOk returns a tuple with the TimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetTimeTypeOk() (*string, bool) {
	if o == nil || o.TimeType == nil {
		return nil, false
	}
	return o.TimeType, true
}

// HasTimeType returns a boolean if a field has been set.
func (o *ScheduleConfig) HasTimeType() bool {
	return o != nil && o.TimeType != nil
}

// SetTimeType gets a reference to the given string and assigns it to the TimeType field.
func (o *ScheduleConfig) SetTimeType(v string) {
	o.TimeType = &v
}

// GetTimeValue returns the TimeValue field value if set, zero value otherwise.
func (o *ScheduleConfig) GetTimeValue() int64 {
	if o == nil || o.TimeValue == nil {
		var ret int64
		return ret
	}
	return *o.TimeValue
}

// GetTimeValueOk returns a tuple with the TimeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetTimeValueOk() (*int64, bool) {
	if o == nil || o.TimeValue == nil {
		return nil, false
	}
	return o.TimeValue, true
}

// HasTimeValue returns a boolean if a field has been set.
func (o *ScheduleConfig) HasTimeValue() bool {
	return o != nil && o.TimeValue != nil
}

// SetTimeValue gets a reference to the given int64 and assigns it to the TimeValue field.
func (o *ScheduleConfig) SetTimeValue(v int64) {
	o.TimeValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ScheduleFullDateTime != nil {
		toSerialize["ScheduleFullDateTime"] = o.ScheduleFullDateTime
	}
	if o.ScheduleType != nil {
		toSerialize["ScheduleType"] = o.ScheduleType
	}
	if o.TimeTics != nil {
		toSerialize["TimeTics"] = o.TimeTics
	}
	if o.Time != nil {
		toSerialize["Time"] = o.Time
	}
	if o.DateStr != nil {
		toSerialize["DateStr"] = o.DateStr
	}
	if o.Days != nil {
		toSerialize["Days"] = o.Days
	}
	if o.DaysNumber != nil {
		toSerialize["DaysNumber"] = o.DaysNumber
	}
	if o.DayNo != nil {
		toSerialize["DayNo"] = o.DayNo
	}
	if o.Day != nil {
		toSerialize["Day"] = o.Day
	}
	if o.DayNumber != nil {
		toSerialize["DayNumber"] = o.DayNumber
	}
	if o.WeekType != nil {
		toSerialize["WeekType"] = o.WeekType
	}
	if o.TimeType != nil {
		toSerialize["TimeType"] = o.TimeType
	}
	if o.TimeValue != nil {
		toSerialize["TimeValue"] = o.TimeValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ScheduleConfig) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
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
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ScheduleFullDateTime", "ScheduleType", "TimeTics", "Time", "DateStr", "Days", "DaysNumber", "DayNo", "Day", "DayNumber", "WeekType", "TimeType", "TimeValue"})
	} else {
		return err
	}

	o.ScheduleFullDateTime = all.ScheduleFullDateTime
	o.ScheduleType = all.ScheduleType
	o.TimeTics = all.TimeTics
	o.Time = all.Time
	o.DateStr = all.DateStr
	o.Days = all.Days
	o.DaysNumber = all.DaysNumber
	o.DayNo = all.DayNo
	o.Day = all.Day
	o.DayNumber = all.DayNumber
	o.WeekType = all.WeekType
	o.TimeType = all.TimeType
	o.TimeValue = all.TimeValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableScheduleConfig handles when a null is used for ScheduleConfig.
type NullableScheduleConfig struct {
	value *ScheduleConfig
	isSet bool
}

// Get returns the associated value.
func (v NullableScheduleConfig) Get() *ScheduleConfig {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableScheduleConfig) Set(val *ScheduleConfig) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableScheduleConfig) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableScheduleConfig) UnSet() {
	v.value = nil
	v.isSet = false
}

// NewNullableScheduleConfig initializes the struct as if Set has been called.
func NewNullableScheduleConfig(val *ScheduleConfig) *NullableScheduleConfig {
	return &NullableScheduleConfig{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableScheduleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableScheduleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
