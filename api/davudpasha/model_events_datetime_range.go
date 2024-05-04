package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// EventsDateTimeRange represents the date time range for events.
type EventsDateTimeRange struct {
	// DateTimeType specifies the type of date time range
	DateTimeType EventsDateTimeType `json:"DateTimeType"`
	// StartDate is the start date of the range
	StartDate *string `json:"StartDate,omitempty"`
	// EndDate is the end date of the range
	EndDate *string `json:"EndDate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewEventsDateTimeRange creates a new EventsDateTimeRange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsDateTimeRange(dateTimeType EventsDateTimeType) *EventsDateTimeRange {
	this := EventsDateTimeRange{}
	this.DateTimeType = dateTimeType
	return &this
}

// NewDateTimeRangeWithDefaults creates a new DateTimeRange object.
// This constructor will assign default values to properties that have it defined,
// but it doensn't guarantee that properties requiered by API are set.
func NewEventsDateTimeRangeWithDefault(dateTimeType EventsDateTimeType) *EventsDateTimeRange {
	this := EventsDateTimeRange{}
	this.DateTimeType = dateTimeType
	return &this
}

// GetDateTimeType returns the DateTimeType field value if set, zero value otherwise.
func (o *EventsDateTimeRange) GetDateTimeType() EventsDateTimeType {
	if o == nil {
		var ret EventsDateTimeType
		return ret
	}
	return o.DateTimeType
}

// GetDateTimeTypeOk returns a tuple with the DateTimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsDateTimeRange) GetDateTimeTypeOk() (*EventsDateTimeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateTimeType, true
}

// SetDateTimeType gets a reference to the given string and assigns it to the DateTimeType field.
func (o *EventsDateTimeRange) SetDateTimeType(v EventsDateTimeType) {
	o.DateTimeType = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *EventsDateTimeRange) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsDateTimeRange) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *EventsDateTimeRange) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *EventsDateTimeRange) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *EventsDateTimeRange) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsDateTimeRange) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *EventsDateTimeRange) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *EventsDateTimeRange) SetEndDate(v string) {
	o.EndDate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsDateTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["DateTimeType"] = o.DateTimeType
	if o.StartDate != nil {
		toSerialize["StartDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["EndDate"] = o.EndDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *EventsDateTimeRange) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		DateTimeType *EventsDateTimeType `json:"DateTimeType"`
		StartDate    *string             `json:"StartDate,omitempty"`
		EndDate      *string             `json:"EndDate,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DateTimeType == nil {
		return fmt.Errorf("required field DateTimeType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"DateTimeType", "StartDate", "EndDate"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DateTimeType.IsValid() {
		hasInvalidField = true
	} else {
		o.DateTimeType = *all.DateTimeType
	}

	if all.DateTimeType == EVENTSDATETIMETYPE_DATERANGE.Ptr() {
		o.StartDate = all.StartDate
		o.EndDate = all.EndDate
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
