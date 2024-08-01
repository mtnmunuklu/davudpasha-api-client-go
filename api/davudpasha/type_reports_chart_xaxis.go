package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsChartXAxis defines the configuration for the X-axis of a chart in reports.
type ReportsChartXAxis struct {
	// Label for the x-axis.
	Label *string `json:"Label,omitempty"`
	// Interval for x-axis values.
	Interval *int64 `json:"Interval,omitempty"`
	// Angle for x-axis labels.
	Angle *int64 `json:"Angle,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsChartXAxis creates a new ReportsChartXAxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsChartXAxis() *ReportsChartXAxis {
	this := ReportsChartXAxis{}
	return &this
}

// NewReportsChartXAxisWithDefaults creates a new ReportsChartXAxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsChartXAxisWithDefaults() *ReportsChartXAxis {
	this := ReportsChartXAxis{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ReportsChartXAxis) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartXAxis) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ReportsChartXAxis) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ReportsChartXAxis) SetLabel(v string) {
	o.Label = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ReportsChartXAxis) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartXAxis) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ReportsChartXAxis) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *ReportsChartXAxis) SetInterval(v int64) {
	o.Interval = &v
}

// GetAngle returns the Angle field value if set, zero value otherwise.
func (o *ReportsChartXAxis) GetAngle() int64 {
	if o == nil || o.Angle == nil {
		var ret int64
		return ret
	}
	return *o.Angle
}

// GetAngleOk returns a tuple with the Angle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartXAxis) GetAngleOk() (*int64, bool) {
	if o == nil || o.Angle == nil {
		return nil, false
	}
	return o.Angle, true
}

// HasAngle returns a boolean if a field has been set.
func (o *ReportsChartXAxis) HasAngle() bool {
	return o != nil && o.Angle != nil
}

// SetAngle gets a reference to the given int64 and assigns it to the Angle field.
func (o *ReportsChartXAxis) SetAngle(v int64) {
	o.Angle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsChartXAxis) MarshalJSON() ([]byte, error) {
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
	if o.Angle != nil {
		toSerialize["Angle"] = o.Angle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsChartXAxis) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label    *string `json:"Label,omitempty"`
		Interval *int64  `json:"Interval,omitempty"`
		Angle    *int64  `json:"Angle,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Label", "Interval", "Angle"})
	} else {
		return err
	}

	o.Label = all.Label
	o.Interval = all.Interval
	o.Angle = all.Angle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
