package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsChartVisualization defines the visualization configuration for a chart in reports.
type ReportsChartVisualization struct {
	// Type of chart.
	ChartType *string `json:"ChartType,omitempty"`
	// Colors for series in the chart.
	SerieColors []string `json:"SerieColors,omitempty"`
	// Line width for the chart.
	LineWidth *int64 `json:"LineWidth,omitempty"`
	// Inner radius for the chart (applicable for pie charts).
	InnerRadius *float64 `json:"InnerRadius,omitempty"`
	// Line type for the chart.
	LineType *string `json:"LineType,omitempty"`
	// Indicates whether to show values on the chart.
	ShowValues *bool `json:"ShowValues,omitempty"`
	// Indicates whether to show null values on the chart.
	ShowNullValues *bool `json:"ShowNullValues,omitempty"`
	// X-axis configuration for the chart.
	XAxis *ReportsChartXAxis `json:"XAxis,omitempty"`
	// Y-axis configuration for the chart.
	YAxis *ReportsChartYAxis `json:"YAxis,omitempty"`
	// Legend configuration for the chart.
	Legend *ReportsChartLegend `json:"Legend,omitempty"`
	// ID of the UI DP module.
	UiDPModuleID *string `json:"UiDPModuleId,omitempty"`
	// Color scheme for the chart.
	ColorScheme common.NullableString `json:"ColorScheme,omitempty"`
	// Additional settings for the chart.
	Settings *string `json:"Settings,omitempty"`
	// Maximum row count for the chart data.
	MaxRowCount *int64 `json:"MaxRowCount,omitempty"`
	// Number of slices for pie charts.
	SlicesNumber *int64 `json:"SlicesNumber,omitempty"`
	// Indicates whether to show labels on the chart.
	ShowLabel *bool `json:"ShowLabel,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsChartVisualization creates a new ReportsChartVisualization object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsChartVisualization() *ReportsChartVisualization {
	this := ReportsChartVisualization{}
	return &this
}

// NewReportsChartVisualizationWithDefaults creates a new ReportsChartVisualization object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsChartVisualizationWithDefaults() *ReportsChartVisualization {
	this := ReportsChartVisualization{}
	return &this
}

// GetChartType returns the ChartType field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetChartType() string {
	if o == nil || o.ChartType == nil {
		var ret string
		return ret
	}
	return *o.ChartType
}

// GetChartTypeOk returns a tuple with the ChartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetChartTypeOk() (*string, bool) {
	if o == nil || o.ChartType == nil {
		return nil, false
	}
	return o.ChartType, true
}

// HasChartType returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasChartType() bool {
	return o != nil && o.ChartType != nil
}

// SetChartType gets a reference to the given string and assigns it to the ChartType field.
func (o *ReportsChartVisualization) SetChartType(v string) {
	o.ChartType = &v
}

// GetSerieColors returns the SerieColors field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetSerieColors() []string {
	if o == nil || o.SerieColors == nil {
		var ret []string
		return ret
	}
	return o.SerieColors
}

// GetSerieColorsOk returns a tuple with the SerieColors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetSerieColorsOk() (*[]string, bool) {
	if o == nil || o.SerieColors == nil {
		return nil, false
	}
	return &o.SerieColors, true
}

// HasSerieColors returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasSerieColors() bool {
	return o != nil && o.SerieColors != nil
}

// SetSerieColors gets a reference to the given []string and assigns it to the SerieColors field.
func (o *ReportsChartVisualization) SetSerieColors(v []string) {
	o.SerieColors = v
}

// GetLineWidth returns the LineWidth field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetLineWidth() int64 {
	if o == nil || o.LineWidth == nil {
		var ret int64
		return ret
	}
	return *o.LineWidth
}

// GetLineWidthOk returns a tuple with the LineWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetLineWidthOk() (*int64, bool) {
	if o == nil || o.LineWidth == nil {
		return nil, false
	}
	return o.LineWidth, true
}

// HasLineWidth returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasLineWidth() bool {
	return o != nil && o.LineWidth != nil
}

// SetLineWidth gets a reference to the given int64 and assigns it to the LineWidth field.
func (o *ReportsChartVisualization) SetLineWidth(v int64) {
	o.LineWidth = &v
}

// GetInnerRadius returns the InnerRadius field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetInnerRadius() float64 {
	if o == nil || o.InnerRadius == nil {
		var ret float64
		return ret
	}
	return *o.InnerRadius
}

// GetInnerRadiusOk returns a tuple with the InnerRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetInnerRadiusOk() (*float64, bool) {
	if o == nil || o.InnerRadius == nil {
		return nil, false
	}
	return o.InnerRadius, true
}

// HasInnerRadius returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasInnerRadius() bool {
	return o != nil && o.InnerRadius != nil
}

// SetInnerRadius gets a reference to the given float64 and assigns it to the InnerRadius field.
func (o *ReportsChartVisualization) SetInnerRadius(v float64) {
	o.InnerRadius = &v
}

// GetLineType returns the LineType field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetLineType() string {
	if o == nil || o.LineType == nil {
		var ret string
		return ret
	}
	return *o.LineType
}

// GetLineTypeOk returns a tuple with the LineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetLineTypeOk() (*string, bool) {
	if o == nil || o.LineType == nil {
		return nil, false
	}
	return o.LineType, true
}

// HasLineType returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasLineType() bool {
	return o != nil && o.LineType != nil
}

// SetLineType gets a reference to the given string and assigns it to the LineType field.
func (o *ReportsChartVisualization) SetLineType(v string) {
	o.LineType = &v
}

// GetShowValues returns the ShowValues field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetShowValues() bool {
	if o == nil || o.ShowValues == nil {
		var ret bool
		return ret
	}
	return *o.ShowValues
}

// GetShowValuesOk returns a tuple with the ShowValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetShowValuesOk() (*bool, bool) {
	if o == nil || o.ShowValues == nil {
		return nil, false
	}
	return o.ShowValues, true
}

// HasShowValues returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasShowValues() bool {
	return o != nil && o.ShowValues != nil
}

// SetShowValues gets a reference to the given bool and assigns it to the ShowValues field.
func (o *ReportsChartVisualization) SetShowValues(v bool) {
	o.ShowValues = &v
}

// GetShowNullValues returns the ShowNullValues field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetShowNullValues() bool {
	if o == nil || o.ShowNullValues == nil {
		var ret bool
		return ret
	}
	return *o.ShowNullValues
}

// GetShowNullValuesOk returns a tuple with the ShowNullValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetShowNullValuesOk() (*bool, bool) {
	if o == nil || o.ShowNullValues == nil {
		return nil, false
	}
	return o.ShowNullValues, true
}

// HasShowNullValues returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasShowNullValues() bool {
	return o != nil && o.ShowNullValues != nil
}

// SetShowNullValues gets a reference to the given bool and assigns it to the ShowNullValues field.
func (o *ReportsChartVisualization) SetShowNullValues(v bool) {
	o.ShowNullValues = &v
}

// GetXAxis returns the XAxis field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetXAxis() ReportsChartXAxis {
	if o == nil || o.XAxis == nil {
		var ret ReportsChartXAxis
		return ret
	}
	return *o.XAxis
}

// GetXAxisOk returns a tuple with the XAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetXAxisOk() (*ReportsChartXAxis, bool) {
	if o == nil || o.XAxis == nil {
		return nil, false
	}
	return o.XAxis, true
}

// HasXAxis returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasXAxis() bool {
	return o != nil && o.XAxis != nil
}

// SetXAxis gets a reference to the given ReportsChartXAxis and assigns it to the XAxis field.
func (o *ReportsChartVisualization) SetXAxis(v ReportsChartXAxis) {
	o.XAxis = &v
}

// GetYAxis returns the YAxis field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetYAxis() ReportsChartYAxis {
	if o == nil || o.YAxis == nil {
		var ret ReportsChartYAxis
		return ret
	}
	return *o.YAxis
}

// GetYAxisOk returns a tuple with the YAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetYAxisOk() (*ReportsChartYAxis, bool) {
	if o == nil || o.YAxis == nil {
		return nil, false
	}
	return o.YAxis, true
}

// HasYAxis returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasYAxis() bool {
	return o != nil && o.YAxis != nil
}

// SetYAxis gets a reference to the given ReportsChartYAxis and assigns it to the YAxis field.
func (o *ReportsChartVisualization) SetYAxis(v ReportsChartYAxis) {
	o.YAxis = &v
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetLegend() ReportsChartLegend {
	if o == nil || o.Legend == nil {
		var ret ReportsChartLegend
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetLegendOk() (*ReportsChartLegend, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given ReportsChartLegend and assigns it to the Legend field.
func (o *ReportsChartVisualization) SetLegend(v ReportsChartLegend) {
	o.Legend = &v
}

// GetUiDPModuleID returns the UiDPModuleID field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetUiDPModuleID() string {
	if o == nil || o.UiDPModuleID == nil {
		var ret string
		return ret
	}
	return *o.UiDPModuleID
}

// GetUiDPModuleIDOk returns a tuple with the UiDPModuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetUiDPModuleIDOk() (*string, bool) {
	if o == nil || o.UiDPModuleID == nil {
		return nil, false
	}
	return o.UiDPModuleID, true
}

// HasUiDPModuleID returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasUiDPModuleID() bool {
	return o != nil && o.UiDPModuleID != nil
}

// SetUiDPModuleID gets a reference to the given string and assigns it to the UiDPModuleID field.
func (o *ReportsChartVisualization) SetUiDPModuleID(v string) {
	o.UiDPModuleID = &v
}

// GetColorScheme returns the ColorScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsChartVisualization) GetColorScheme() string {
	if o == nil || o.ColorScheme.Get() == nil {
		var ret string
		return ret
	}
	return *o.ColorScheme.Get()
}

// GetColorSchemeOk returns a tuple with the ColorScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsChartVisualization) GetColorSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ColorScheme.Get(), o.ColorScheme.IsSet()
}

// HasColorScheme returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasColorScheme() bool {
	return o != nil && o.ColorScheme.IsSet()
}

// SetColorScheme gets a reference to the given common.NullableString and assigns it to the ColorScheme field.
func (o *ReportsChartVisualization) SetColorScheme(v string) {
	o.ColorScheme.Set(&v)
}

// SetColorSchemeNil sets the value for ColorScheme to be an explicit nil.
func (o *ReportsChartVisualization) SetColorSchemeNil() {
	o.ColorScheme.Set(nil)
}

// UnSetColorScheme ensures that no value is present for ColorScheme, not even an explicit nil.
func (o *ReportsChartVisualization) UnsetColorScheme() {
	o.ColorScheme.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetSettings() string {
	if o == nil || o.Settings == nil {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetSettingsOk() (*string, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *ReportsChartVisualization) SetSettings(v string) {
	o.Settings = &v
}

// GetMaxRowCount returns the MaxRowCount field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetMaxRowCount() int64 {
	if o == nil || o.MaxRowCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxRowCount
}

// GetMaxRowCountOk returns a tuple with the MaxRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetMaxRowCountOk() (*int64, bool) {
	if o == nil || o.MaxRowCount == nil {
		return nil, false
	}
	return o.MaxRowCount, true
}

// HasMaxRowCount returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasMaxRowCount() bool {
	return o != nil && o.MaxRowCount != nil
}

// SetMaxRowCount gets a reference to the given int64 and assigns it to the MaxRowCount field.
func (o *ReportsChartVisualization) SetMaxRowCount(v int64) {
	o.MaxRowCount = &v
}

// GetSlicesNumber returns the SlicesNumber field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetSlicesNumber() int64 {
	if o == nil || o.SlicesNumber == nil {
		var ret int64
		return ret
	}
	return *o.SlicesNumber
}

// GetSlicesNumberOk returns a tuple with the SlicesNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetSlicesNumberOk() (*int64, bool) {
	if o == nil || o.SlicesNumber == nil {
		return nil, false
	}
	return o.SlicesNumber, true
}

// HasSlicesNumber returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasSlicesNumber() bool {
	return o != nil && o.SlicesNumber != nil
}

// SetSlicesNumber gets a reference to the given int64 and assigns it to the SlicesNumber field.
func (o *ReportsChartVisualization) SetSlicesNumber(v int64) {
	o.SlicesNumber = &v
}

// GetShowLabel returns the ShowLabel field value if set, zero value otherwise.
func (o *ReportsChartVisualization) GetShowLabel() bool {
	if o == nil || o.ShowLabel == nil {
		var ret bool
		return ret
	}
	return *o.ShowLabel
}

// GetShowLabelOk returns a tuple with the ShowLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsChartVisualization) GetShowLabelOk() (*bool, bool) {
	if o == nil || o.ShowLabel == nil {
		return nil, false
	}
	return o.ShowLabel, true
}

// HasShowLabel returns a boolean if a field has been set.
func (o *ReportsChartVisualization) HasShowLabel() bool {
	return o != nil && o.ShowLabel != nil
}

// SetShowLabel gets a reference to the given bool and assigns it to the ShowLabel field.
func (o *ReportsChartVisualization) SetShowLabel(v bool) {
	o.ShowLabel = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsChartVisualization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ChartType != nil {
		toSerialize["ChartType"] = o.ChartType
	}
	if o.SerieColors != nil {
		toSerialize["SerieColors"] = o.SerieColors
	}
	if o.LineWidth != nil {
		toSerialize["LineWidth"] = o.LineWidth
	}
	if o.InnerRadius != nil {
		toSerialize["InnerRadius"] = o.InnerRadius
	}
	if o.LineType != nil {
		toSerialize["Tags"] = o.LineType
	}
	if o.ShowValues != nil {
		toSerialize["ShowValues"] = o.ShowValues
	}
	if o.ShowNullValues != nil {
		toSerialize["ShowNullValues"] = o.ShowNullValues
	}
	if o.XAxis != nil {
		toSerialize["XAxis"] = o.XAxis
	}
	if o.YAxis != nil {
		toSerialize["YAxis"] = o.YAxis
	}
	if o.Legend != nil {
		toSerialize["Legend"] = o.Legend
	}
	if o.UiDPModuleID != nil {
		toSerialize["UiDPModuleId"] = o.UiDPModuleID
	}
	if o.ColorScheme.IsSet() {
		toSerialize["ColorScheme"] = o.ColorScheme.Get()
	}
	if o.Settings != nil {
		toSerialize["Settings"] = o.Settings
	}
	if o.MaxRowCount != nil {
		toSerialize["MaxRowCount"] = o.MaxRowCount
	}
	if o.SlicesNumber != nil {
		toSerialize["SlicesNumber"] = o.SlicesNumber
	}
	if o.ShowLabel != nil {
		toSerialize["ShowLabel"] = o.ShowLabel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsChartVisualization) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChartType      *string               `json:"ChartType,omitempty"`
		SerieColors    []string              `json:"SerieColors,omitempty"`
		LineWidth      *int64                `json:"LineWidth,omitempty"`
		InnerRadius    *float64              `json:"InnerRadius,omitempty"`
		LineType       *string               `json:"LineType,omitempty"`
		ShowValues     *bool                 `json:"ShowValues,omitempty"`
		ShowNullValues *bool                 `json:"ShowNullValues,omitempty"`
		XAxis          *ReportsChartXAxis    `json:"XAxis,omitempty"`
		YAxis          *ReportsChartYAxis    `json:"YAxis,omitempty"`
		Legend         *ReportsChartLegend   `json:"Legend,omitempty"`
		UiDPModuleID   *string               `json:"UiDPModuleId,omitempty"`
		ColorScheme    common.NullableString `json:"ColorScheme,omitempty"`
		Settings       *string               `json:"Settings,omitempty"`
		MaxRowCount    *int64                `json:"MaxRowCount,omitempty"`
		SlicesNumber   *int64                `json:"SlicesNumber,omitempty"`
		ShowLabel      *bool                 `json:"ShowLabel,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ChartType", "SerieColors", "LineWidth", "InnerRadius", "LineType", "ShowValues", "ShowNullValues", "XAxis", "YAxis", "Legend", "UiDPModuleId", "ColorScheme", "Settings", "MaxRowCount", "SlicesNumber", "ShowLabel"})
	} else {
		return err
	}

	o.ChartType = all.ChartType
	o.SerieColors = all.SerieColors
	o.LineWidth = all.LineWidth
	o.InnerRadius = all.InnerRadius
	o.LineType = all.LineType
	o.ShowValues = all.ShowValues
	o.ShowNullValues = all.ShowNullValues
	hasInvalidField := false
	if all.XAxis != nil && all.XAxis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.XAxis = all.XAxis
	if all.YAxis != nil && all.YAxis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.YAxis = all.YAxis
	if all.Legend != nil && all.Legend.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Legend = all.Legend
	o.UiDPModuleID = all.UiDPModuleID
	o.ColorScheme = all.ColorScheme
	o.Settings = all.Settings
	o.MaxRowCount = all.MaxRowCount
	o.SlicesNumber = all.SlicesNumber
	o.ShowLabel = all.ShowLabel

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
