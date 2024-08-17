package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsQuery represents a report query.
type ReportsQuery struct {
	// Represents a key value, used as an identifier.
	Key *float64 `json:"Key,omitempty"`
	// Name of the query.
	Name *string `json:"Nane,omitempty"`
	// Data related to the query.
	Data *ReportsQueryData `json:"Data,omitempty"`
	// Indicates if the table should be shown.
	ShowTable *bool `json:"ShowTable,omitempty"`
	// Visualization settings for the table.
	TableVisualization *ReportsTableVisualization `json:"TableVisualization,omitempty"`
	// Reference to the form used for generating the report.
	FormRef *FormRef `json:"FormRef"`
	// Indicates if the chart should be shown.
	ShowChart *bool `json:"ShowChart,omitempty"`
	// Visualization settings for the chart.
	ChartVisualization *ReportsChartVisualization `json:"ChartVisualization,omitempty"`
	// Extended data for the query.
	ExtData NullableReportsQueryExtData `json:"ExtData,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsQuery creates a new ReportsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsQuery() *ReportsQuery {
	this := ReportsQuery{}
	return &this
}

// NewReportsQueryWithDefaults creates a new ReportsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsQueryWithDefaults() *ReportsQuery {
	this := ReportsQuery{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ReportsQuery) GetKey() float64 {
	if o == nil || o.Key == nil {
		var ret float64
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetKeyOk() (*float64, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ReportsQuery) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given float64 and assigns it to the Key field.
func (o *ReportsQuery) SetKey(v float64) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReportsQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReportsQuery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReportsQuery) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReportsQuery) GetData() ReportsQueryData {
	if o == nil || o.Data == nil {
		var ret ReportsQueryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetDataOk() (*ReportsQueryData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReportsQuery) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given ReportsQueryData and assigns it to the Data field.
func (o *ReportsQuery) SetData(v ReportsQueryData) {
	o.Data = &v
}

// GetShowTable returns the ShowTable field value if set, zero value otherwise.
func (o *ReportsQuery) GetShowTable() bool {
	if o == nil || o.ShowTable == nil {
		var ret bool
		return ret
	}
	return *o.ShowTable
}

// GetShowTableOk returns a tuple with the ShowTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetShowTableOk() (*bool, bool) {
	if o == nil || o.ShowTable == nil {
		return nil, false
	}
	return o.ShowTable, true
}

// HasShowTable returns a boolean if a field has been set.
func (o *ReportsQuery) HasShowTable() bool {
	return o != nil && o.ShowTable != nil
}

// SetShowTable gets a reference to the given bool and assigns it to the ShowTable field.
func (o *ReportsQuery) SetShowTable(v bool) {
	o.ShowTable = &v
}

// GetTableVisualization returns the TableVisualization field value if set, zero value otherwise.
func (o *ReportsQuery) GetTableVisualization() ReportsTableVisualization {
	if o == nil || o.TableVisualization == nil {
		var ret ReportsTableVisualization
		return ret
	}
	return *o.TableVisualization
}

// GetTableVisualizationOk returns a tuple with the TableVisualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetTableVisualizationOk() (*ReportsTableVisualization, bool) {
	if o == nil || o.TableVisualization == nil {
		return nil, false
	}
	return o.TableVisualization, true
}

// HasTableVisualization returns a boolean if a field has been set.
func (o *ReportsQuery) HasTableVisualization() bool {
	return o != nil && o.TableVisualization != nil
}

// SetTableVisualization gets a reference to the given ReportsTableVisualization and assigns it to the TableVisualization field.
func (o *ReportsQuery) SetTableVisualization(v ReportsTableVisualization) {
	o.TableVisualization = &v
}

// GetFormRef returns the FormRef field value if set, zero value otherwise.
func (o *ReportsQuery) GetFormRef() FormRef {
	if o == nil || o.FormRef == nil {
		var ret FormRef
		return ret
	}
	return *o.FormRef
}

// GetFormRefOk returns a tuple with the FormRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetFormRefOk() (*FormRef, bool) {
	if o == nil || o.FormRef == nil {
		return nil, false
	}
	return o.FormRef, true
}

// HasFormRef returns a boolean if a field has been set.
func (o *ReportsQuery) HasFormRef() bool {
	return o != nil && o.FormRef != nil
}

// SetFormRef gets a reference to the given FormRef and assigns it to the FormRef field.
func (o *ReportsQuery) SetFormRef(v FormRef) {
	o.FormRef = &v
}

// GetShowChart returns the ShowChart field value if set, zero value otherwise.
func (o *ReportsQuery) GetShowChart() bool {
	if o == nil || o.ShowChart == nil {
		var ret bool
		return ret
	}
	return *o.ShowChart
}

// GetShowChartOk returns a tuple with the ShowChart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetShowChartOk() (*bool, bool) {
	if o == nil || o.ShowChart == nil {
		return nil, false
	}
	return o.ShowChart, true
}

// HasShowChart returns a boolean if a field has been set.
func (o *ReportsQuery) HasShowChart() bool {
	return o != nil && o.ShowChart != nil
}

// SetShowChart gets a reference to the given bool and assigns it to the ShowChart field.
func (o *ReportsQuery) SetShowChart(v bool) {
	o.ShowChart = &v
}

// GetChartVisualization returns the ChartVisualization field value if set, zero value otherwise.
func (o *ReportsQuery) GetChartVisualization() ReportsChartVisualization {
	if o == nil || o.ChartVisualization == nil {
		var ret ReportsChartVisualization
		return ret
	}
	return *o.ChartVisualization
}

// GetChartVisualizationOk returns a tuple with the ChartVisualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetChartVisualizationOk() (*ReportsChartVisualization, bool) {
	if o == nil || o.ChartVisualization == nil {
		return nil, false
	}
	return o.ChartVisualization, true
}

// HasChartVisualization returns a boolean if a field has been set.
func (o *ReportsQuery) HasChartVisualization() bool {
	return o != nil && o.ChartVisualization != nil
}

// SetChartVisualization gets a reference to the given ReportsChartVisualization and assigns it to the ChartVisualization field.
func (o *ReportsQuery) SetChartVisualization(v ReportsChartVisualization) {
	o.ChartVisualization = &v
}

// GetExtData returns the ExtData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsQuery) GetExtData() ReportsQueryExtData {
	if o == nil || o.ExtData.Get() == nil {
		var ret ReportsQueryExtData
		return ret
	}
	return *o.ExtData.Get()
}

// GetExtDataOk returns a tuple with the ExtData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQuery) GetExtDataOk() (*ReportsQueryExtData, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtData.Get(), o.ExtData.IsSet()
}

// HasExtData returns a boolean if a ExtData has been set.
func (o *ReportsQuery) HasExtData() bool {
	return o != nil && o.ExtData.IsSet()
}

// SetExtData gets a reference to the given common.NullableString and assigns it to the ExtData field.
func (o *ReportsQuery) SetExtData(v ReportsQueryExtData) {
	o.ExtData.Set(&v)
}

// SetExtDataNil sets the value for ExtData to be an explicit nil.
func (o *ReportsQuery) SetExtDataNil() {
	o.ExtData.Set(nil)
}

// UnSetExtData ensures that no value is present for ExtData, not even an explicit nil.
func (o *ReportsQuery) UnSetExtData() {
	o.ExtData.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if o.ShowTable != nil {
		toSerialize["ShowTable"] = o.ShowTable
	}
	if o.TableVisualization != nil {
		toSerialize["TableVisualization"] = o.TableVisualization
	}
	if o.FormRef != nil {
		toSerialize["FormRef"] = o.FormRef
	}
	if o.ShowChart != nil {
		toSerialize["ShowChart"] = o.ShowChart
	}
	if o.ChartVisualization != nil {
		toSerialize["ChartVisualization"] = o.ChartVisualization
	}
	if o.ExtData.IsSet() {
		toSerialize["ExtData"] = o.ExtData.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsQuery) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key                *float64                    `json:"Key,omitempty"`
		Name               *string                     `json:"Nane,omitempty"`
		Data               *ReportsQueryData           `json:"Data,omitempty"`
		ShowTable          *bool                       `json:"ShowTable,omitempty"`
		TableVisualization *ReportsTableVisualization  `json:"TableVisualization,omitempty"`
		FormRef            *FormRef                    `json:"FormRef"`
		ShowChart          *bool                       `json:"ShowChart,omitempty"`
		ChartVisualization *ReportsChartVisualization  `json:"ChartVisualization,omitempty"`
		ExtData            NullableReportsQueryExtData `json:"ExtData,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Key", "Name", "Data", "ShowTable", "TableVisualization", "FormRef", "ShowChart", "ChartVisualization", "ExtData"})
	} else {
		return err
	}

	o.Key = all.Key
	o.Name = all.Name
	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.ShowTable = all.ShowTable
	if all.TableVisualization != nil && all.TableVisualization.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TableVisualization = all.TableVisualization
	o.FormRef = all.FormRef
	o.ShowChart = all.ShowChart
	if all.ChartVisualization != nil && all.ChartVisualization.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChartVisualization = all.ChartVisualization
	o.ExtData = all.ExtData

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
