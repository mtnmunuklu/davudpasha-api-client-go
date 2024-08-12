package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsTableVisualization defines the visualization options for a table in a report.
type ReportsTableVisualization struct {
	// Type of the chart.
	ChartType *string `json:"ChartType,omitempty"`
	// Maximum row count.
	MaxRowCount *int64 `json:"MaxRowCount,omitempty"`
	// Columns for the visualization.
	Columns []string `json:"Columns,omitempty"`
	// List of values corresponding to specific columns in a dataset.
	ColumnValues []string `json:"ColumnValues,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
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
func (o *ReportsTableVisualization) GetMaxRowCount() int64 {
	if o == nil || o.MaxRowCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxRowCount
}

// GetMaxRowCountOk returns a tuple with the MaxRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsTableVisualization) GetMaxRowCountOk() (*int64, bool) {
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
func (o *ReportsTableVisualization) SetMaxRowCount(v int64) {
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

// GetColumnValues returns the ColumnValues field value if set, zero value otherwise.
func (o *ReportsTableVisualization) GetColumnValues() []string {
	if o == nil || o.ColumnValues == nil {
		var ret []string
		return ret
	}
	return o.ColumnValues
}

// GetColumnValuesOk returns a tuple with the ColumnValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsTableVisualization) GetColumnValuesOk() (*[]string, bool) {
	if o == nil || o.ColumnValues == nil {
		return nil, false
	}
	return &o.ColumnValues, true
}

// HasColumnValues returns a boolean if a field has been set.
func (o *ReportsTableVisualization) HasColumnValues() bool {
	return o != nil && o.ColumnValues != nil
}

// SetColumnValues gets a reference to the given []string and assigns it to the ColumnValues field.
func (o *ReportsTableVisualization) SetColumnValues(v []string) {
	o.ColumnValues = v
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
	if o.ColumnValues != nil {
		toSerialize["ColumnValues"] = o.ColumnValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsTableVisualization) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChartType    *string  `json:"ChartType,omitempty"`
		MaxRowCount  *int64   `json:"MaxRowCount,omitempty"`
		Columns      []string `json:"Columns,omitempty"`
		ColumnValues []string `json:"ColumnValues,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ChartType", "MaxRowCount", "Columns", "ColumnValues"})
	} else {
		return err
	}

	o.ChartType = all.ChartType
	o.MaxRowCount = all.MaxRowCount
	o.Columns = all.Columns
	o.ColumnValues = all.ColumnValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
