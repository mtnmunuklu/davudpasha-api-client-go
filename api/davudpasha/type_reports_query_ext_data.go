package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsQueryExtData represents the extended data for a report query.
type ReportsQueryExtData struct {
	// Type of the report.
	ReportType *string `json:"ReportType,omitempty"`
	// Date-time range for the report.
	DateTimeRange *DateTimeRange `json:"DateTimeRange,omitempty"`
	// Indicates if alerts should be included.
	Alerts *bool `json:"Alerts,omitempty"`
	// Indicates if log positions should be included.
	LogPositions *bool `json:"LogPositions,omitempty"`
	// Indicates if system information should be included.
	SystemInfo *bool `json:"SystemInfo,omitempty"`
	// Indicates if log counts should be included.
	LogCounts *bool `json:"LogCounts,omitempty"`
	// Indicates if classifications should be included.
	Classifications *bool `json:"Classifications,omitempty"`
	// Indicates if correlation should be included.
	Correlation *bool `json:"Correlation,omitempty"`
	// Indicates if timestamps should be included.
	TimeStamps *bool `json:"TimeStamps,omitempty"`
	// Indicates if log counts chart should be included.
	LogCountsChart *bool `json:"LogCounts_Chart,omitempty"`
	// Indicates if log counts table should be included.
	LogCountsTable *bool `json:"LogCounts_Table,omitempty"`
	// Indicates if classifications chart should be included.
	ClassificationsChart *bool `json:"Classifications_Chart,omitempty"`
	// Indicates if classifications server chart should be included.
	ClassificationsSvrChart *bool `json:"Classifications_SvrChart,omitempty"`
	// Indicates if classifications table should be included.
	ClassificationsTable *bool `json:"Classifications_Table,omitempty"`
	// Indicates if correlation alert chart should be included.
	CorrelationAlertChart *bool `json:"Correlation_AlertChart,omitempty"`
	// Indicates if correlation server chart should be included.
	CorrelationSvrChart *bool `json:"Correlation_SvrChart,omitempty"`
	// Indicates if correlation table should be included.
	CorrelationTable *bool `json:"Correlation_Table,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewReportsQueryExtData creates a new ReportsQueryExtData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsQueryExtData() *ReportsQueryExtData {
	this := ReportsQueryExtData{}
	return &this
}

// NewReportsQueryExtDataWithDefaults creates a new ReportsQueryExtData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsQueryExtDataWithDefaults() *ReportsQueryExtData {
	this := ReportsQueryExtData{}
	return &this
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetReportType() string {
	if o == nil || o.ReportType == nil {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetReportTypeOk() (*string, bool) {
	if o == nil || o.ReportType == nil {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasReportType() bool {
	return o != nil && o.ReportType != nil
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *ReportsQueryExtData) SetReportType(v string) {
	o.ReportType = &v
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetDateTimeRange() DateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret DateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DateTimeRange field.
func (o *ReportsQueryExtData) SetDateTimeRange(v DateTimeRange) {
	o.DateTimeRange = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetAlerts() bool {
	if o == nil || o.Alerts == nil {
		var ret bool
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetAlertsOk() (*bool, bool) {
	if o == nil || o.Alerts == nil {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasAlerts() bool {
	return o != nil && o.Alerts != nil
}

// SetAlerts gets a reference to the given bool and assigns it to the Alerts field.
func (o *ReportsQueryExtData) SetAlerts(v bool) {
	o.Alerts = &v
}

// GetLogPositions returns the LogPositions field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetLogPositions() bool {
	if o == nil || o.LogPositions == nil {
		var ret bool
		return ret
	}
	return *o.LogPositions
}

// GetLogPositionsOk returns a tuple with the LogPositions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetLogPositionsOk() (*bool, bool) {
	if o == nil || o.LogPositions == nil {
		return nil, false
	}
	return o.LogPositions, true
}

// HasLogPositions returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasLogPositions() bool {
	return o != nil && o.LogPositions != nil
}

// SetLogPositions gets a reference to the given bool and assigns it to the LogPositions field.
func (o *ReportsQueryExtData) SetLogPositions(v bool) {
	o.LogPositions = &v
}

// GetSystemInfo returns the SystemInfo field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetSystemInfo() bool {
	if o == nil || o.SystemInfo == nil {
		var ret bool
		return ret
	}
	return *o.SystemInfo
}

// GetSystemInfoOk returns a tuple with the SystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetSystemInfoOk() (*bool, bool) {
	if o == nil || o.SystemInfo == nil {
		return nil, false
	}
	return o.SystemInfo, true
}

// HasSystemInfo returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasSystemInfo() bool {
	return o != nil && o.SystemInfo != nil
}

// SetSystemInfo gets a reference to the given bool and assigns it to the SystemInfo field.
func (o *ReportsQueryExtData) SetSystemInfo(v bool) {
	o.SystemInfo = &v
}

// GetLogCounts returns the LogCounts field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetLogCounts() bool {
	if o == nil || o.LogCounts == nil {
		var ret bool
		return ret
	}
	return *o.LogCounts
}

// GetLogCountsOk returns a tuple with the LogCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetLogCountsOk() (*bool, bool) {
	if o == nil || o.LogCounts == nil {
		return nil, false
	}
	return o.LogCounts, true
}

// HasLogCounts returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasLogCounts() bool {
	return o != nil && o.LogCounts != nil
}

// SetLogCounts gets a reference to the given bool and assigns it to the LogCounts field.
func (o *ReportsQueryExtData) SetLogCounts(v bool) {
	o.LogCounts = &v
}

// GetClassifications returns the Classifications field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetClassifications() bool {
	if o == nil || o.Classifications == nil {
		var ret bool
		return ret
	}
	return *o.Classifications
}

// GetClassificationsOk returns a tuple with the Classifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetClassificationsOk() (*bool, bool) {
	if o == nil || o.Classifications == nil {
		return nil, false
	}
	return o.Classifications, true
}

// HasClassifications returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasClassifications() bool {
	return o != nil && o.Classifications != nil
}

// SetClassifications gets a reference to the given bool and assigns it to the Classifications field.
func (o *ReportsQueryExtData) SetClassifications(v bool) {
	o.Classifications = &v
}

// GetCorrelation returns the Correlation field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetCorrelation() bool {
	if o == nil || o.Correlation == nil {
		var ret bool
		return ret
	}
	return *o.Correlation
}

// GetCorrelationOk returns a tuple with the Correlation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetCorrelationOk() (*bool, bool) {
	if o == nil || o.Correlation == nil {
		return nil, false
	}
	return o.Correlation, true
}

// HasCorrelation returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasCorrelation() bool {
	return o != nil && o.Correlation != nil
}

// SetCorrelation gets a reference to the given bool and assigns it to the Correlation field.
func (o *ReportsQueryExtData) SetCorrelation(v bool) {
	o.Correlation = &v
}

// GetTimeStamps returns the TimeStamps field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetTimeStamps() bool {
	if o == nil || o.TimeStamps == nil {
		var ret bool
		return ret
	}
	return *o.TimeStamps
}

// GetTimeStampsOk returns a tuple with the TimeStamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetTimeStampsOk() (*bool, bool) {
	if o == nil || o.TimeStamps == nil {
		return nil, false
	}
	return o.TimeStamps, true
}

// HasTimeStamps returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasTimeStamps() bool {
	return o != nil && o.TimeStamps != nil
}

// SetTimeStamps gets a reference to the given bool and assigns it to the TimeStamps field.
func (o *ReportsQueryExtData) SetTimeStamps(v bool) {
	o.TimeStamps = &v
}

// GetLogCountsChart returns the LogCountsChart field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetLogCountsChart() bool {
	if o == nil || o.LogCountsChart == nil {
		var ret bool
		return ret
	}
	return *o.LogCountsChart
}

// GetLogCountsChartOk returns a tuple with the LogCountsChart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetLogCountsChartOk() (*bool, bool) {
	if o == nil || o.LogCountsChart == nil {
		return nil, false
	}
	return o.LogCountsChart, true
}

// HasLogCountsChart returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasLogCountsChart() bool {
	return o != nil && o.LogCountsChart != nil
}

// SetLogCountsChart gets a reference to the given bool and assigns it to the LogCountsChart field.
func (o *ReportsQueryExtData) SetLogCountsChart(v bool) {
	o.LogCountsChart = &v
}

// GetLogCountsTable returns the LogCountsTable field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetLogCountsTable() bool {
	if o == nil || o.LogCountsTable == nil {
		var ret bool
		return ret
	}
	return *o.LogCountsTable
}

// GetLogCountsTableOk returns a tuple with the LogCountsTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetLogCountsTableOk() (*bool, bool) {
	if o == nil || o.LogCountsTable == nil {
		return nil, false
	}
	return o.LogCountsTable, true
}

// HasLogCountsTable returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasLogCountsTable() bool {
	return o != nil && o.LogCountsTable != nil
}

// SetLogCountsTable gets a reference to the given bool and assigns it to the LogCountsTable field.
func (o *ReportsQueryExtData) SetLogCountsTable(v bool) {
	o.LogCountsTable = &v
}

// GetClassificationsChart returns the ClassificationsChart field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetClassificationsChart() bool {
	if o == nil || o.ClassificationsChart == nil {
		var ret bool
		return ret
	}
	return *o.ClassificationsChart
}

// GetClassificationsChartOk returns a tuple with the ClassificationsChart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetClassificationsChartOk() (*bool, bool) {
	if o == nil || o.ClassificationsChart == nil {
		return nil, false
	}
	return o.ClassificationsChart, true
}

// HasClassificationsChart returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasClassificationsChart() bool {
	return o != nil && o.ClassificationsChart != nil
}

// SetClassificationsChart gets a reference to the given bool and assigns it to the ClassificationsChart field.
func (o *ReportsQueryExtData) SetClassificationsChart(v bool) {
	o.ClassificationsChart = &v
}

// GetClassificationsSvrChart returns the ClassificationsSvrChart field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetClassificationsSvrChart() bool {
	if o == nil || o.ClassificationsSvrChart == nil {
		var ret bool
		return ret
	}
	return *o.ClassificationsSvrChart
}

// GetClassificationsSvrChartOk returns a tuple with the ClassificationsSvrChart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetClassificationsSvrChartOk() (*bool, bool) {
	if o == nil || o.ClassificationsSvrChart == nil {
		return nil, false
	}
	return o.ClassificationsSvrChart, true
}

// HasClassificationsSvrChart returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasClassificationsSvrChart() bool {
	return o != nil && o.ClassificationsSvrChart != nil
}

// SetClassificationsSvrChart gets a reference to the given bool and assigns it to the ClassificationsSvrChart field.
func (o *ReportsQueryExtData) SetClassificationsSvrChart(v bool) {
	o.ClassificationsSvrChart = &v
}

// GetClassificationsTable returns the ClassificationsTable field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetClassificationsTable() bool {
	if o == nil || o.ClassificationsTable == nil {
		var ret bool
		return ret
	}
	return *o.ClassificationsTable
}

// GetClassificationsTableOk returns a tuple with the ClassificationsTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetClassificationsTableOk() (*bool, bool) {
	if o == nil || o.ClassificationsTable == nil {
		return nil, false
	}
	return o.ClassificationsTable, true
}

// HasClassificationsTable returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasClassificationsTable() bool {
	return o != nil && o.ClassificationsTable != nil
}

// SetClassificationsTable gets a reference to the given bool and assigns it to the ClassificationsTable field.
func (o *ReportsQueryExtData) SetClassificationsTable(v bool) {
	o.ClassificationsTable = &v
}

// GetCorrelationAlertChart returns the CorrelationAlertChart field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetCorrelationAlertChart() bool {
	if o == nil || o.CorrelationAlertChart == nil {
		var ret bool
		return ret
	}
	return *o.CorrelationAlertChart
}

// GetCorrelationAlertChartOk returns a tuple with the CorrelationAlertChart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetCorrelationAlertChartOk() (*bool, bool) {
	if o == nil || o.CorrelationAlertChart == nil {
		return nil, false
	}
	return o.CorrelationAlertChart, true
}

// HasCorrelationAlertChart returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasCorrelationAlertChart() bool {
	return o != nil && o.CorrelationAlertChart != nil
}

// SetCorrelationAlertChart gets a reference to the given bool and assigns it to the CorrelationAlertChart field.
func (o *ReportsQueryExtData) SetCorrelationAlertChart(v bool) {
	o.CorrelationAlertChart = &v
}

// GetCorrelationSvrChart returns the CorrelationSvrChart field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetCorrelationSvrChart() bool {
	if o == nil || o.CorrelationSvrChart == nil {
		var ret bool
		return ret
	}
	return *o.CorrelationSvrChart
}

// GetCorrelationSvrChartOk returns a tuple with the CorrelationSvrChart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetCorrelationSvrChartOk() (*bool, bool) {
	if o == nil || o.CorrelationSvrChart == nil {
		return nil, false
	}
	return o.CorrelationSvrChart, true
}

// HasCorrelationSvrChart returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasCorrelationSvrChart() bool {
	return o != nil && o.CorrelationSvrChart != nil
}

// SetCorrelationSvrChart gets a reference to the given bool and assigns it to the CorrelationSvrChart field.
func (o *ReportsQueryExtData) SetCorrelationSvrChart(v bool) {
	o.CorrelationSvrChart = &v
}

// GetCorrelationTable returns the CorrelationTable field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetCorrelationTable() bool {
	if o == nil || o.CorrelationTable == nil {
		var ret bool
		return ret
	}
	return *o.CorrelationTable
}

// GetCorrelationTableOk returns a tuple with the CorrelationTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetCorrelationTableOk() (*bool, bool) {
	if o == nil || o.CorrelationTable == nil {
		return nil, false
	}
	return o.CorrelationTable, true
}

// HasCorrelationTable returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasCorrelationTable() bool {
	return o != nil && o.CorrelationTable != nil
}

// SetCorrelationTable gets a reference to the given bool and assigns it to the CorrelationTable field.
func (o *ReportsQueryExtData) SetCorrelationTable(v bool) {
	o.CorrelationTable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsQueryExtData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ReportType != nil {
		toSerialize["ReportType"] = o.ReportType
	}
	if o.DateTimeRange != nil {
		toSerialize["DateTimeRange"] = o.DateTimeRange
	}
	if o.Alerts != nil {
		toSerialize["Alerts"] = o.Alerts
	}
	if o.LogPositions != nil {
		toSerialize["LogPositions"] = o.LogPositions
	}
	if o.SystemInfo != nil {
		toSerialize["SystemInfo"] = o.SystemInfo
	}
	if o.LogCounts != nil {
		toSerialize["LogCounts"] = o.LogCounts
	}
	if o.Classifications != nil {
		toSerialize["Classifications"] = o.Classifications
	}
	if o.Correlation != nil {
		toSerialize["Correlation"] = o.Correlation
	}
	if o.TimeStamps != nil {
		toSerialize["TimeStamps"] = o.TimeStamps
	}
	if o.LogCountsChart != nil {
		toSerialize["LogCountsChart"] = o.LogCountsChart
	}
	if o.LogCountsTable != nil {
		toSerialize["LogCountsTable"] = o.LogCountsTable
	}
	if o.ClassificationsChart != nil {
		toSerialize["ClassificationsChart"] = o.ClassificationsChart
	}
	if o.ClassificationsSvrChart != nil {
		toSerialize["ClassificationsSvrChart"] = o.ClassificationsSvrChart
	}
	if o.ClassificationsTable != nil {
		toSerialize["ClassificationsTable"] = o.ClassificationsTable
	}
	if o.CorrelationAlertChart != nil {
		toSerialize["CorrelationAlertChart"] = o.CorrelationAlertChart
	}
	if o.CorrelationSvrChart != nil {
		toSerialize["CorrelationSvrChart"] = o.CorrelationSvrChart
	}
	if o.CorrelationTable != nil {
		toSerialize["CorrelationTable"] = o.CorrelationTable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsQueryExtData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReportType              *string        `json:"ReportType,omitempty"`
		DateTimeRange           *DateTimeRange `json:"DateTimeRange,omitempty"`
		Alerts                  *bool          `json:"Alerts,omitempty"`
		LogPositions            *bool          `json:"LogPositions,omitempty"`
		SystemInfo              *bool          `json:"SystemInfo,omitempty"`
		LogCounts               *bool          `json:"LogCounts,omitempty"`
		Classifications         *bool          `json:"Classifications,omitempty"`
		Correlation             *bool          `json:"Correlation,omitempty"`
		TimeStamps              *bool          `json:"TimeStamps,omitempty"`
		LogCountsChart          *bool          `json:"LogCounts_Chart,omitempty"`
		LogCountsTable          *bool          `json:"LogCounts_Table,omitempty"`
		ClassificationsChart    *bool          `json:"Classifications_Chart,omitempty"`
		ClassificationsSvrChart *bool          `json:"Classifications_SvrChart,omitempty"`
		ClassificationsTable    *bool          `json:"Classifications_Table,omitempty"`
		CorrelationAlertChart   *bool          `json:"Correlation_AlertChart,omitempty"`
		CorrelationSvrChart     *bool          `json:"Correlation_SvrChart,omitempty"`
		CorrelationTable        *bool          `json:"Correlation_Table,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ReportType", "DateTimeRange", "Alerts", "LogPositions", "SystemInfo", "LogCounts", "Classifications", "Correlation", "TimeStamps", "LogCountsChart", "LogCountsTable", "ClassificationsChart", "ClassificationsSvrChart", "ClassificationsTable", "CorrelationAlertChart", "CorrelationSvrChart", "CorrelationTable"})
	} else {
		return err
	}

	o.ReportType = all.ReportType
	hasInvalidField := false
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	o.Alerts = all.Alerts
	o.LogPositions = all.LogPositions
	o.SystemInfo = all.SystemInfo
	o.LogCounts = all.LogCounts
	o.Classifications = all.Classifications
	o.Correlation = all.Correlation
	o.TimeStamps = all.TimeStamps
	o.LogCountsChart = all.LogCountsChart
	o.LogCountsTable = all.LogCountsTable
	o.ClassificationsChart = all.ClassificationsChart
	o.ClassificationsSvrChart = all.ClassificationsSvrChart
	o.ClassificationsTable = all.ClassificationsTable
	o.CorrelationAlertChart = all.CorrelationAlertChart
	o.CorrelationSvrChart = all.CorrelationSvrChart
	o.CorrelationTable = all.CorrelationTable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableReportsQueryExtData handles when a null is used for ReportsQueryExtData.
type NullableReportsQueryExtData struct {
	value *ReportsQueryExtData
	isSet bool
}

// Get returns the associated value.
func (v NullableReportsQueryExtData) Get() *ReportsQueryExtData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableReportsQueryExtData) Set(val *ReportsQueryExtData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableReportsQueryExtData) IsSet() bool {
	return v.isSet
}

// UnSet sets the value to nil and resets the set flag/
func (v *NullableReportsQueryExtData) UnSet() {
	v.value = nil
	v.isSet = false
}

// NewNullableReportsQueryExtData initializes the struct as if Set has been called.
func NewNullableReportsQueryExtData(val *ReportsQueryExtData) *NullableReportsQueryExtData {
	return &NullableReportsQueryExtData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableReportsQueryExtData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableReportsQueryExtData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
