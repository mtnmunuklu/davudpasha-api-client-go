package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsChartYAxis defines the configuration for the Y-axis of a chart in reports.
type ReportsChartYAxis struct {
	// Label for the y-axis.
	Label *string `json:"Label,omitempty"`
	// Interval for y-axis values.
	Interval *int64 `json:"Interval,omitempty"`
	// Minimum value for the y-axis.
	MinValue common.NullableInt64 `json:"MinValue,omitempty"`
	// Maximum value for the y-axis.
	MaxValue common.NullableInt64 `json:"MaxValue,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewReportsChartYAxis creates a new ReportsChartYAxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsChartYAxis() *ReportsChartYAxis {
	this := ReportsChartYAxis{}
	return &this
}

// NewReportsChartYAxisWithDefaults creates a new ReportsChartYAxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsChartYAxisWithDefaults() *ReportsChartYAxis {
	this := ReportsChartYAxis{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ReportsChartYAxis) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartYAxis) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ReportsChartYAxis) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ReportsChartYAxis) SetLabel(v string) {
	o.Label = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ReportsChartYAxis) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartYAxis) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ReportsChartYAxis) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *ReportsChartYAxis) SetInterval(v int64) {
	o.Interval = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsChartYAxis) GetMinValue() int64 {
	if o == nil || o.MinValue.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MinValue.Get()
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsChartYAxis) GetMinValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinValue.Get(), o.MinValue.IsSet()
}

// HasMinValue returns a boolean if a field has been set.
func (o *ReportsChartYAxis) HasMinValue() bool {
	return o != nil && o.MinValue.IsSet()
}

// SetMinValue gets a reference to the given datadog.NullableString and assigns it to the MinValue field.
func (o *ReportsChartYAxis) SetMinValue(v int64) {
	o.MinValue.Set(&v)
}

// SetMinValueNil sets the value for MinValue to be an explicit nil.
func (o *ReportsChartYAxis) SetMinValueNil() {
	o.MinValue.Set(nil)
}

// UnSetMinValue ensures that no value is present for MinValue, not even an explicit nil.
func (o *ReportsChartYAxis) UnSetMinValue() {
	o.MinValue.UnSet()
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsChartYAxis) GetMaxValue() int64 {
	if o == nil || o.MaxValue.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MaxValue.Get()
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsChartYAxis) GetMaxValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxValue.Get(), o.MaxValue.IsSet()
}

// HasMaxValue returns a boolean if a field has been set.
func (o *ReportsChartYAxis) HasMaxValue() bool {
	return o != nil && o.MaxValue.IsSet()
}

// SetMaxValue gets a reference to the given datadog.NullableString and assigns it to the MaxValue field.
func (o *ReportsChartYAxis) SetMaxValue(v int64) {
	o.MaxValue.Set(&v)
}

// SetMaxValueNil sets the value for MaxValue to be an explicit nil.
func (o *ReportsChartYAxis) SetMaxValueNil() {
	o.MaxValue.Set(nil)
}

// UnSetMaxValue ensures that no value is present for MaxValue, not even an explicit nil.
func (o *ReportsChartYAxis) UnSetMaxValue() {
	o.MaxValue.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsChartYAxis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Interval != nil {
		toSerialize["Interval"] = o.Interval
	}
	if o.MinValue.IsSet() {
		toSerialize["MinValue"] = o.MinValue.Get()
	}
	if o.MaxValue.IsSet() {
		toSerialize["MaxValue"] = o.MaxValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsChartYAxis) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label    *string              `json:"Label,omitempty"`
		Interval *int64               `json:"Interval,omitempty"`
		MinValue common.NullableInt64 `json:"MinValue,omitempty"`
		MaxValue common.NullableInt64 `json:"MaxValue,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Label", "Interval", "MinValue", "MaxValue"})
	} else {
		return err
	}

	o.Label = all.Label
	o.Interval = all.Interval
	o.MinValue = all.MinValue
	o.MaxValue = all.MaxValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
