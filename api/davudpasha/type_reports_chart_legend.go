package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ReportsChartLegend struct {
	LegendPosition *string `json:"LegendPosition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewReportsChartLegend creates a new ReportsChartLegend object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsChartLegend() *ReportsChartLegend {
	this := ReportsChartLegend{}
	return &this
}

// NewReportsChartLegendWithDefault creates a new ReportsChartLegend object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsChartLegendWithDefault() *ReportsChartLegend {
	this := ReportsChartLegend{}
	return &this
}

// GetLegendPosition returns the LegendPosition field value if set, zero value otherwise.
func (o *ReportsChartLegend) GetLegendPosition() string {
	if o == nil || o.LegendPosition == nil {
		var ret string
		return ret
	}
	return *o.LegendPosition
}

// GetLegendPositionOk returns a tuple with the LegendPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartLegend) GetLegendPositionOk() (*string, bool) {
	if o == nil || o.LegendPosition == nil {
		return nil, false
	}
	return o.LegendPosition, true
}

// HasLegendPosition returns a boolean if a field has been set.
func (o *ReportsChartLegend) HasLegendPosition() bool {
	return o != nil && o.LegendPosition != nil
}

// SetLegendPosition gets a reference to the given string and assigns it to the LegendPosition field.
func (o *ReportsChartLegend) SetLegendPosition(v string) {
	o.LegendPosition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsChartLegend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.LegendPosition != nil {
		toSerialize["LegendPosition"] = o.LegendPosition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsChartLegend) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		LegendPosition *string `json:"LegendPosition,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"LegendPosition"})
	} else {
		return err
	}
	o.LegendPosition = all.LegendPosition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
