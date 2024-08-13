package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsQueryExtData represents the extended data for a report query.
type ReportsQueryExtData struct {
	// Type of the report. ReportType: PythonScript, Query, SearchArchive, NonRepudiation, Timestamp, Code
	ReportType ReportsReportType `json:"ReportType,omitempty"`
	// Date-time range for the report.
	DateTimeRange *DateTimeRange `json:"DateTimeRange,omitempty"`
	// Indicates if alerts should be included.
	Alerts *bool `json:"Alerts,omitempty"`
	// Indicates if the log status should be included.
	LogStatus *bool `json:"LogStatus,omitempty"`
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
	// Specifies if the data should be grouped by date.
	GroupByDate common.NullableBool `json:"GroupByDate,omitempty"`
	// Specifies the maximum number of rows to include in the result.
	MaxRowCount common.NullableInt64 `json:"MaxRowCount,omitempty"`
	// Indicates if only faulty files should be shown.
	ShowOnlyFaultyFiles *bool `json:"ShowOnlyFaultyFiles,omitempty"`
	// Indicates if disk information should be included.
	DiskInfo *bool `json:"DiskInfo,omitempty"`
	// List of log IDs to be processed.
	LgsIDs []string `json:"LgsIds,omitempty"`
	// Specifies if any of the provided values are included.
	IncludesAny common.NullableString `json:"IncludesAny,omitempty"`
	// List of values to be included.
	Includes common.NullableList[string] `json:"Includes,omitempty"`
	// List of values to be excluded.
	Excludes common.NullableList[string] `json:"Excludes,omitempty"`
	// Specifies the size of the data to be processed.
	Size *int64 `json:"Size,omitempty"`
	// Specifies the number of parallel processes to use.
	ParallelCount *int64 `json:"ParallelCount,omitempty"`
	// Indicates if the search should be performed on raw logs.
	SearchOnRawLogs *bool `json:"SearchOnRawLogs,omitempty"`
	// List of column names to be included in the report.
	ReportColNames common.NullableList[string] `json:"ReportColNames,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
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
func (o *ReportsQueryExtData) GetReportType() ReportsReportType {
	if o == nil {
		var ret ReportsReportType
		return ret
	}
	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetReportTypeOk() (*ReportsReportType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *ReportsQueryExtData) SetReportType(v ReportsReportType) {
	o.ReportType = v
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

// GetLogStatus returns the LogStatus field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetLogStatus() bool {
	if o == nil || o.LogStatus == nil {
		var ret bool
		return ret
	}
	return *o.LogStatus
}

// GetLogStatusOk returns a tuple with the LogStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetLogStatusOk() (*bool, bool) {
	if o == nil || o.LogStatus == nil {
		return nil, false
	}
	return o.LogStatus, true
}

// HasLogStatus returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasLogStatus() bool {
	return o != nil && o.LogStatus != nil
}

// SetLogStatus gets a reference to the given bool and assigns it to the LogStatus field.
func (o *ReportsQueryExtData) SetLogStatus(v bool) {
	o.LogStatus = &v
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

// GetGroupByDate returns the GroupByDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsQueryExtData) GetGroupByDate() bool {
	if o == nil || o.GroupByDate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.GroupByDate.Get()
}

// GetGroupByDateOk returns a tuple with the GroupByDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetGroupByDateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupByDate.Get(), o.GroupByDate.IsSet()
}

// HasGroupByDate returns a boolean if a GroupByDate has been set.
func (o *ReportsQueryExtData) HasGroupByDate() bool {
	return o != nil && o.GroupByDate.IsSet()
}

// SetGroupByDate gets a reference to the given common.NullableString and assigns it to the GroupByDate field.
func (o *ReportsQueryExtData) SetGroupByDate(v bool) {
	o.GroupByDate.Set(&v)
}

// SetGroupByDateNil sets the value for GroupByDate to be an explicit nil.
func (o *ReportsQueryExtData) SetGroupByDateNil() {
	o.GroupByDate.Set(nil)
}

// UnSetGroupByDate ensures that no value is present for GroupByDate, not even an explicit nil.
func (o *ReportsQueryExtData) UnSetGroupByDate() {
	o.GroupByDate.UnSet()
}

// GetMaxRowCount returns the MaxRowCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsQueryExtData) GetMaxRowCount() int64 {
	if o == nil || o.MaxRowCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MaxRowCount.Get()
}

// GetMaxRowCountOk returns a tuple with the MaxRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetMaxRowCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRowCount.Get(), o.MaxRowCount.IsSet()
}

// HasMaxRowCount returns a boolean if a MaxRowCount has been set.
func (o *ReportsQueryExtData) HasMaxRowCount() bool {
	return o != nil && o.MaxRowCount.IsSet()
}

// SetMaxRowCount gets a reference to the given common.NullableString and assigns it to the MaxRowCount field.
func (o *ReportsQueryExtData) SetMaxRowCount(v int64) {
	o.MaxRowCount.Set(&v)
}

// SetMaxRowCountNil sets the value for MaxRowCount to be an explicit nil.
func (o *ReportsQueryExtData) SetMaxRowCountNil() {
	o.MaxRowCount.Set(nil)
}

// UnSetMaxRowCount ensures that no value is present for MaxRowCount, not even an explicit nil.
func (o *ReportsQueryExtData) UnSetMaxRowCount() {
	o.MaxRowCount.UnSet()
}

// GetShowOnlyFaultyFiles returns the ShowOnlyFaultyFiles field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetShowOnlyFaultyFiles() bool {
	if o == nil || o.ShowOnlyFaultyFiles == nil {
		var ret bool
		return ret
	}
	return *o.ShowOnlyFaultyFiles
}

// GetShowOnlyFaultyFilesOk returns a tuple with the ShowOnlyFaultyFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetShowOnlyFaultyFilesOk() (*bool, bool) {
	if o == nil || o.ShowOnlyFaultyFiles == nil {
		return nil, false
	}
	return o.ShowOnlyFaultyFiles, true
}

// HasShowOnlyFaultyFiles returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasShowOnlyFaultyFiles() bool {
	return o != nil && o.ShowOnlyFaultyFiles != nil
}

// SetShowOnlyFaultyFiles gets a reference to the given bool and assigns it to the ShowOnlyFaultyFiles field.
func (o *ReportsQueryExtData) SetShowOnlyFaultyFiles(v bool) {
	o.ShowOnlyFaultyFiles = &v
}

// GetDiskInfo returns the DiskInfo field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetDiskInfo() bool {
	if o == nil || o.DiskInfo == nil {
		var ret bool
		return ret
	}
	return *o.DiskInfo
}

// GetDiskInfoOk returns a tuple with the DiskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetDiskInfoOk() (*bool, bool) {
	if o == nil || o.DiskInfo == nil {
		return nil, false
	}
	return o.DiskInfo, true
}

// HasDiskInfo returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasDiskInfo() bool {
	return o != nil && o.DiskInfo != nil
}

// SetDiskInfo gets a reference to the given bool and assigns it to the DiskInfo field.
func (o *ReportsQueryExtData) SetDiskInfo(v bool) {
	o.DiskInfo = &v
}

// GetLgsIDs returns the LgsIDs field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetLgsIDs() []string {
	if o == nil || o.LgsIDs == nil {
		var ret []string
		return ret
	}
	return o.LgsIDs
}

// GetLgsIDsOk returns a tuple with the LgsIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetLgsIDsOk() (*[]string, bool) {
	if o == nil || o.LgsIDs == nil {
		return nil, false
	}
	return &o.LgsIDs, true
}

// HasLgsIDs returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasLgsIDs() bool {
	return o != nil && o.LgsIDs != nil
}

// SetLgsIDs gets a reference to the given []string and assigns it to the LgsIDs field.
func (o *ReportsQueryExtData) SetLgsIDs(v []string) {
	o.LgsIDs = v
}

// GetIncludesAny returns the IncludesAny field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsQueryExtData) GetIncludesAny() string {
	if o == nil || o.IncludesAny.Get() == nil {
		var ret string
		return ret
	}
	return *o.IncludesAny.Get()
}

// GetIncludesAnyOk returns a tuple with the IncludesAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetIncludesAnyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludesAny.Get(), o.IncludesAny.IsSet()
}

// HasIncludesAny returns a boolean if a IncludesAny has been set.
func (o *ReportsQueryExtData) HasIncludesAny() bool {
	return o != nil && o.IncludesAny.IsSet()
}

// SetIncludesAny gets a reference to the given common.NullableString and assigns it to the IncludesAny field.
func (o *ReportsQueryExtData) SetIncludesAny(v string) {
	o.IncludesAny.Set(&v)
}

// SetIncludesAnyNil sets the value for IncludesAny to be an explicit nil.
func (o *ReportsQueryExtData) SetIncludesAnyNil() {
	o.IncludesAny.Set(nil)
}

// UnSetIncludesAny ensures that no value is present for IncludesAny, not even an explicit nil.
func (o *ReportsQueryExtData) UnSetIncludesAny() {
	o.IncludesAny.UnSet()
}

// GetIncludes returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetIncludes() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Includes.Get(), o.Includes.IsSet()
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetIncludesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Includes.Get(), o.Includes.IsSet()
}

// HasIncludes returns a boolean if a Includes has been set.
func (o *ReportsQueryExtData) HasIncludes() bool {
	return o != nil && o.Includes.IsSet()
}

// SetIncludes gets a reference to the given common.Nullable[]string and assigns it to the Includes field.
func (o *ReportsQueryExtData) SetIncludes(v []string) {
	o.Includes.Set(&v)
}

// SetIncludesNil sets the value for Includes to be an explicit nil.
func (o *ReportsQueryExtData) SetIncludesNil() {
	o.Includes.Set(nil)
}

// UnSetIncludes ensures that no value is present for Includes, not even an explicit nil.
func (o *ReportsQueryExtData) UnSetIncludes() {
	o.Includes.UnSet()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ReportsQueryExtData) SetSize(v int64) {
	o.Size = &v
}

// GetParallelCount returns the ParallelCount field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetParallelCount() int64 {
	if o == nil || o.ParallelCount == nil {
		var ret int64
		return ret
	}
	return *o.ParallelCount
}

// GetParallelCountOk returns a tuple with the ParallelCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetParallelCountOk() (*int64, bool) {
	if o == nil || o.ParallelCount == nil {
		return nil, false
	}
	return o.ParallelCount, true
}

// HasParallelCount returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasParallelCount() bool {
	return o != nil && o.ParallelCount != nil
}

// SetParallelCount gets a reference to the given int64 and assigns it to the ParallelCount field.
func (o *ReportsQueryExtData) SetParallelCount(v int64) {
	o.ParallelCount = &v
}

// GetExcludes returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetExcludes() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Excludes.Get(), o.Excludes.IsSet()
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetExcludesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Excludes.Get(), o.Excludes.IsSet()
}

// HasExcludes returns a boolean if a Excludes has been set.
func (o *ReportsQueryExtData) HasExcludes() bool {
	return o != nil && o.Excludes.IsSet()
}

// SetExcludes gets a reference to the given common.Nullable[]string and assigns it to the Excludes field.
func (o *ReportsQueryExtData) SetExcludes(v []string) {
	o.Excludes.Set(&v)
}

// SetExcludesNil sets the value for Excludes to be an explicit nil.
func (o *ReportsQueryExtData) SetExcludesNil() {
	o.Excludes.Set(nil)
}

// UnSetExcludes ensures that no value is present for Excludes, not even an explicit nil.
func (o *ReportsQueryExtData) UnSetExcludes() {
	o.Excludes.UnSet()
}

// GetSearchOnRawLogs returns the SearchOnRawLogs field value if set, zero value otherwise.
func (o *ReportsQueryExtData) GetSearchOnRawLogs() bool {
	if o == nil || o.SearchOnRawLogs == nil {
		var ret bool
		return ret
	}
	return *o.SearchOnRawLogs
}

// GetSearchOnRawLogsOk returns a tuple with the SearchOnRawLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryExtData) GetSearchOnRawLogsOk() (*bool, bool) {
	if o == nil || o.SearchOnRawLogs == nil {
		return nil, false
	}
	return o.SearchOnRawLogs, true
}

// HasSearchOnRawLogs returns a boolean if a field has been set.
func (o *ReportsQueryExtData) HasSearchOnRawLogs() bool {
	return o != nil && o.SearchOnRawLogs != nil
}

// SetSearchOnRawLogs gets a reference to the given bool and assigns it to the SearchOnRawLogs field.
func (o *ReportsQueryExtData) SetSearchOnRawLogs(v bool) {
	o.SearchOnRawLogs = &v
}

// GetReportColNames returns a tuple with the ReportColNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetReportColNames() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportColNames.Get(), o.ReportColNames.IsSet()
}

// GetReportColNamesOk returns a tuple with the ReportColNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryExtData) GetReportColNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportColNames.Get(), o.ReportColNames.IsSet()
}

// HasReportColNames returns a boolean if a ReportColNames has been set.
func (o *ReportsQueryExtData) HasReportColNames() bool {
	return o != nil && o.ReportColNames.IsSet()
}

// SetReportColNames gets a reference to the given common.Nullable[]string and assigns it to the ReportColNames field.
func (o *ReportsQueryExtData) SetReportColNames(v []string) {
	o.ReportColNames.Set(&v)
}

// SetReportColNamesNil sets the value for ReportColNames to be an explicit nil.
func (o *ReportsQueryExtData) SetReportColNamesNil() {
	o.ReportColNames.Set(nil)
}

// UnSetReportColNames ensures that no value is present for ReportColNames, not even an explicit nil.
func (o *ReportsQueryExtData) UnSetReportColNames() {
	o.ReportColNames.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsQueryExtData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ReportType.IsValid() {
		toSerialize["ReportType"] = o.ReportType
	}
	if o.DateTimeRange != nil {
		toSerialize["DateTimeRange"] = o.DateTimeRange
	}
	if o.Alerts != nil {
		toSerialize["Alerts"] = o.Alerts
	}
	if o.LogStatus != nil {
		toSerialize["LogStatus"] = o.LogStatus
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
	if o.GroupByDate.IsSet() {
		toSerialize["GroupByDate"] = o.GroupByDate.Get()
	}
	if o.MaxRowCount.IsSet() {
		toSerialize["MaxRowCount"] = o.MaxRowCount.Get()
	}
	if o.ShowOnlyFaultyFiles != nil {
		toSerialize["ShowOnlyFaultyFiles"] = o.ShowOnlyFaultyFiles
	}
	if o.DiskInfo != nil {
		toSerialize["DiskInfo"] = o.DiskInfo
	}
	if o.LgsIDs != nil {
		toSerialize["LgsIDs"] = o.LgsIDs
	}
	if o.IncludesAny.IsSet() {
		toSerialize["IncludesAny"] = o.IncludesAny.Get()
	}
	if o.Includes.IsSet() {
		toSerialize["Includes"] = o.Includes.Get()
	}
	if o.Excludes.IsSet() {
		toSerialize["Excludes"] = o.Excludes.Get()
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.ParallelCount != nil {
		toSerialize["ParallelCount"] = o.ParallelCount
	}
	if o.SearchOnRawLogs != nil {
		toSerialize["SearchOnRawLogs"] = o.SearchOnRawLogs
	}
	if o.ReportColNames.IsSet() {
		toSerialize["ReportColNames"] = o.ReportColNames.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsQueryExtData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReportType              *ReportsReportType          `json:"ReportType,omitempty"`
		DateTimeRange           *DateTimeRange              `json:"DateTimeRange,omitempty"`
		Alerts                  *bool                       `json:"Alerts,omitempty"`
		LogStatus               *bool                       `json:"LogStatus,omitempty"`
		LogPositions            *bool                       `json:"LogPositions,omitempty"`
		SystemInfo              *bool                       `json:"SystemInfo,omitempty"`
		LogCounts               *bool                       `json:"LogCounts,omitempty"`
		Classifications         *bool                       `json:"Classifications,omitempty"`
		Correlation             *bool                       `json:"Correlation,omitempty"`
		TimeStamps              *bool                       `json:"TimeStamps,omitempty"`
		LogCountsChart          *bool                       `json:"LogCounts_Chart,omitempty"`
		LogCountsTable          *bool                       `json:"LogCounts_Table,omitempty"`
		ClassificationsChart    *bool                       `json:"Classifications_Chart,omitempty"`
		ClassificationsSvrChart *bool                       `json:"Classifications_SvrChart,omitempty"`
		ClassificationsTable    *bool                       `json:"Classifications_Table,omitempty"`
		CorrelationAlertChart   *bool                       `json:"Correlation_AlertChart,omitempty"`
		CorrelationSvrChart     *bool                       `json:"Correlation_SvrChart,omitempty"`
		CorrelationTable        *bool                       `json:"Correlation_Table,omitempty"`
		GroupByDate             common.NullableBool         `json:"GroupByDate,omitempty"`
		MaxRowCount             common.NullableInt64        `json:"MaxRowCount,omitempty"`
		ShowOnlyFaultyFiles     *bool                       `json:"ShowOnlyFaultyFiles,omitempty"`
		DiskInfo                *bool                       `json:"DiskInfo,omitempty"`
		LgsIDs                  []string                    `json:"LgsIds,omitempty"`
		IncludesAny             common.NullableString       `json:"IncludesAny,omitempty"`
		Includes                common.NullableList[string] `json:"Includes,omitempty"`
		Excludes                common.NullableList[string] `json:"Excludes,omitempty"`
		Size                    *int64                      `json:"Size,omitempty"`
		ParallelCount           *int64                      `json:"ParallelCount,omitempty"`
		SearchOnRawLogs         *bool                       `json:"SearchOnRawLogs,omitempty"`
		ReportColNames          common.NullableList[string] `json:"ReportColNames,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ReportType", "DateTimeRange", "Alerts", "LogStatus", "LogPositions", "SystemInfo", "LogCounts", "Classifications", "Correlation", "TimeStamps", "LogCountsChart", "LogCountsTable", "ClassificationsChart", "ClassificationsSvrChart", "ClassificationsTable", "CorrelationAlertChart", "CorrelationSvrChart", "CorrelationTable", "GroupByDate", "MaxRowCount", "ShowOnlyFaultyFiles", "DiskInfo", "LgsIds", "IncludesAny", "Includes", "Excludes", "Size", "ParallelCount", "SearchOnRawLogs", "ReportColNames"})
	} else {
		return err
	}
	hasInvalidField := false
	if !all.ReportType.IsValid() {
		hasInvalidField = true
	} else {
		o.ReportType = *all.ReportType
	}
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	o.Alerts = all.Alerts
	o.LogStatus = all.LogStatus
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
	o.GroupByDate = all.GroupByDate
	o.MaxRowCount = all.MaxRowCount
	o.ShowOnlyFaultyFiles = all.ShowOnlyFaultyFiles
	o.DiskInfo = all.DiskInfo
	o.LgsIDs = all.LgsIDs
	o.IncludesAny = all.IncludesAny
	o.Includes = all.Includes
	o.Excludes = all.Excludes
	o.Size = all.Size
	o.ParallelCount = all.ParallelCount
	o.SearchOnRawLogs = all.SearchOnRawLogs
	o.ReportColNames = all.ReportColNames

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
