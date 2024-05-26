package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ReportsTableVisualization struct {
	ChartType   *string  `json:"ChartType,omitempty"`
	MaxRowCount *string  `json:"MaxRowCount,omitempty"`
	Columns     []string `json:"Columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewReportsTableVisualization creates a new ReportsTableVisualization object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsTableVisualization() *ReportsTableVisualization {
	this := ReportsTableVisualization{}
	return &this
}

// NewReportsTableVisualizationWithDefaults creates a new ReportsTableVisualization object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsTableVisualizationWithDefaults() *ReportsTableVisualization {
	this := ReportsTableVisualization{}
	return &this
}

// GetChartType returns the ChartType field value if set, zero value otherwise.
func (o *ReportsTableVisualization) GetChartType() string {
	if o == nil || o.ChartType == nil {
		var ret string
		return ret
	}
	return *o.ChartType
}

// GetChartTypeOk returns a tuple with the ChartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsTableVisualization) GetChartTypeOk() (*string, bool) {
	if o == nil || o.ChartType == nil {
		return nil, false
	}
	return o.ChartType, true
}

// HasChartType returns a boolean if a field has been set.
func (o *ReportsTableVisualization) HasChartType() bool {
	return o != nil && o.ChartType != nil
}

// SetChartType gets a reference to the given string and assigns it to the ChartType field.
func (o *ReportsTableVisualization) SetChartType(v string) {
	o.ChartType = &v
}

// GetMaxRowCount returns the MaxRowCount field value if set, zero value otherwise.
func (o *ReportsTableVisualization) GetMaxRowCount() string {
	if o == nil || o.MaxRowCount == nil {
		var ret string
		return ret
	}
	return *o.MaxRowCount
}

// GetMaxRowCountOk returns a tuple with the MaxRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsTableVisualization) GetMaxRowCountOk() (*string, bool) {
	if o == nil || o.MaxRowCount == nil {
		return nil, false
	}
	return o.MaxRowCount, true
}

// HasMaxRowCount returns a boolean if a field has been set.
func (o *ReportsTableVisualization) HasMaxRowCount() bool {
	return o != nil && o.MaxRowCount != nil
}

// SetMaxRowCount gets a reference to the given string and assigns it to the MaxRowCount field.
func (o *ReportsTableVisualization) SetMaxRowCount(v string) {
	o.MaxRowCount = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ReportsTableVisualization) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsTableVisualization) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ReportsTableVisualization) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *ReportsTableVisualization) SetColumns(v []string) {
	o.Columns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsTableVisualization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ChartType != nil {
		toSerialize["ChartType"] = o.ChartType
	}
	if o.MaxRowCount != nil {
		toSerialize["MaxRowCount"] = o.MaxRowCount
	}
	if o.Columns != nil {
		toSerialize["Columns"] = o.Columns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsTableVisualization) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChartType   *string  `json:"ChartType,omitempty"`
		MaxRowCount *string  `json:"MaxRowCount,omitempty"`
		Columns     []string `json:"Columns,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ChartType", "MaxRowCount", "Columns"})
	} else {
		return err
	}

	o.ChartType = all.ChartType
	o.MaxRowCount = all.MaxRowCount
	o.Columns = all.Columns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
