package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// DateTimeRange represents a range of date and time.
type DateTimeRange struct {
	// DateTimeType specifies the type of date time range.
	DateTimeType DateTimeType `json:"DateTimeType"`
	// StartDate is the start date of the range.
	StartDate *string `json:"StartDate,omitempty"`
	// EndDate is the end date of the range.
	EndDate *string `json:"EndDate,omitempty"`
	// Field specifies the field related to the date time range.
	Field common.NullableString `json:"Field,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDateTimeRange creates a new DateTimeRange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDateTimeRange(dateTimeType DateTimeType) *DateTimeRange {
	this := DateTimeRange{}
	this.DateTimeType = dateTimeType
	return &this
}

// NewDateTimeRangeWithDefaults creates a new DateTimeRange object.
// This constructor will assign default values to properties that have it defined,
// but it doensn't guarantee that properties requiered by API are set.
func NewDateTimeRangeWithDefaults(dateTimeType DateTimeType) *DateTimeRange {
	this := DateTimeRange{}
	this.DateTimeType = dateTimeType
	return &this
}

// GetDateTimeType returns the DateTimeType field value if set, zero value otherwise.
func (o *DateTimeRange) GetDateTimeType() DateTimeType {
	if o == nil {
		var ret DateTimeType
		return ret
	}
	return o.DateTimeType
}

// GetDateTimeTypeOk returns a tuple with the DateTimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetDateTimeTypeOk() (*DateTimeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateTimeType, true
}

// SetDateTimeType gets a reference to the given string and assigns it to the DateTimeType field.
func (o *DateTimeRange) SetDateTimeType(v DateTimeType) {
	o.DateTimeType = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DateTimeRange) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DateTimeRange) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *DateTimeRange) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DateTimeRange) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DateTimeRange) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *DateTimeRange) SetEndDate(v string) {
	o.EndDate = &v
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeRange) GetField() string {
	if o == nil || o.Field.Get() == nil {
		var ret string
		return ret
	}
	return *o.Field.Get()
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DateTimeRange) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Field.Get(), o.Field.IsSet()
}

// HasField returns a boolean if a field has been set.
func (o *DateTimeRange) HasField() bool {
	return o != nil && o.Field.IsSet()
}

// SetField gets a reference to the given common.NullableString and assigns it to the Field field.
func (o *DateTimeRange) SetField(v string) {
	o.Field.Set(&v)
}

// SetFieldNil sets the value for Field to be an explicit nil.
func (o *DateTimeRange) SetFieldNil() {
	o.Field.Set(nil)
}

// UnSetField ensures that no value is present for Field, not even an explicit nil.
func (o *DateTimeRange) UnsetField() {
	o.Field.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DateTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.DateTimeType.IsValid() {
		toSerialize["DateTimeType"] = o.DateTimeType
	}
	if o.StartDate != nil {
		toSerialize["StartDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["EndDate"] = o.EndDate
	}
	if o.Field.IsSet() {
		toSerialize["Field"] = o.Field.Get()
	}
	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *DateTimeRange) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		DateTimeType *DateTimeType         `json:"DateTimeType"`
		StartDate    *string               `json:"StartDate,omitempty"`
		EndDate      *string               `json:"EndDate,omitempty"`
		Field        common.NullableString `json:"Field,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DateTimeType == nil {
		return fmt.Errorf("required field DateTimeType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"DateTimeType", "StartDate", "EndDate", "Field"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DateTimeType.IsValid() {
		hasInvalidField = true
	} else {
		o.DateTimeType = *all.DateTimeType
	}

	if all.DateTimeType == DATETIMETYPE_DATERANGE.Ptr() {
		if all.StartDate != nil || all.EndDate != nil {
			o.StartDate = all.StartDate
			o.EndDate = all.EndDate
		} else {
			return fmt.Errorf("required field StartDate or EndDate missing")
		}
	}

	o.Field = all.Field

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
