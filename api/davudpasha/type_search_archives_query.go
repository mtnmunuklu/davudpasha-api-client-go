package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SearchArchivesQuery defines the structure for search queries in the archives.
type SearchArchivesQuery struct {
	// ID represents the identifier of the search query.
	ID *string `json:"Id,omitempty"`
	// Name specifies the name of the search query.
	Name *string `json:"Name,omitempty"`
	// CreateDate indicates the creation date of the search query.
	CreateDate *string `json:"CreateDate,omitempty"`
	// Includes contains a list of strings to include in the search.
	Includes common.NullableList[string] `json:"Includes,omitempty"`
	// Excludes contains a list of strings to exclude from the search.
	Excludes common.NullableList[string] `json:"Excludes,omitempty"`
	// DateTimeRange specifies the date and time range for the search.
	DateTimeRange *DateTimeRange `json:"DateTimeRange,omitempty"`
	// LgsIDs contains a list of IDs associated with the logs for the search.
	LgsIDs common.NullableList[string] `json:"LgsIds,omitempty"`
	// LgsNames contains a list of names associated with the logs for the search.
	LgsNames common.NullableList[string] `json:"LgsNames,omitempty"`
	// Size specifies the size of the search query.
	Size *int64 `json:"Size,omitempty"`
	// DeleteAfterDay specifies the number of days after which the search results will be deleted.
	DeleteAfterDay *int64 `json:"DeleteAfterDay,omitempty"`
	// ParallelCount specifies the number of parallel searches to be conducted.
	ParallelCount *int64 `json:"ParallelCount,omitempty"`
	// SearchOnRawLogs indicates whether to search on raw log data.
	SearchOnRawLogs *bool `json:"SearchOnRawLogs,omitempty"`
	// SearchActiveLgsOnly indicates whether to search only on active logs.
	SearchActiveLgsOnly *bool `json:"SearchActiveLgsOnly,omitempty"`
	// PauseRequest indicates whether the search request is paused.
	PauseRequest *bool `json:"PauseRequest,omitempty"`
	// ReportColNames contains a list of column names for the search report.
	ReportColNames common.NullableList[string] `json:"ReportColNames,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSearchArchivesQuery creates a new SearchArchivesQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchArchivesQuery() *SearchArchivesQuery {
	this := SearchArchivesQuery{}
	return &this
}

// NewSearchArchivesQueryWithDefaults creates a new SearchArchivesQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSearchArchivesQueryWithDefaults() *SearchArchivesQuery {
	this := SearchArchivesQuery{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *SearchArchivesQuery) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchArchivesQuery) SetName(v string) {
	o.Name = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetCreateDate() string {
	if o == nil || o.CreateDate == nil {
		var ret string
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetCreateDateOk() (*string, bool) {
	if o == nil || o.CreateDate == nil {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasCreateDate() bool {
	return o != nil && o.CreateDate != nil
}

// SetCreateDate gets a reference to the given string and assigns it to the CreateDate field.
func (o *SearchArchivesQuery) SetCreateDate(v string) {
	o.CreateDate = &v
}

// GetIncludes returns the Includes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchArchivesQuery) GetIncludes() []string {
	if o == nil || o.Includes.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Includes.Get()
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchArchivesQuery) GetIncludesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Includes.Get(), o.Includes.IsSet()
}

// HasIncludes returns a boolean if a Includes has been set.
func (o *SearchArchivesQuery) HasIncludes() bool {
	return o != nil && o.Includes.IsSet()
}

// SetIncludes gets a reference to the given common.Nullable[]string and assigns it to the Includes field.
func (o *SearchArchivesQuery) SetIncludes(v []string) {
	o.Includes.Set(&v)
}

// SetIncludesNil sets the value for Includes to be an explicit nil.
func (o *SearchArchivesQuery) SetIncludesNil() {
	o.Includes.Set(nil)
}

// UnSetIncludes ensures that no value is present for Includes, not even an explicit nil.
func (o *SearchArchivesQuery) UnSetIncludes() {
	o.Includes.UnSet()
}

// GetExcludes returns the Excludes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchArchivesQuery) GetExcludes() []string {
	if o == nil || o.Excludes.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Excludes.Get()
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchArchivesQuery) GetExcludesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Excludes.Get(), o.Excludes.IsSet()
}

// HasExcludes returns a boolean if a Excludes has been set.
func (o *SearchArchivesQuery) HasExcludes() bool {
	return o != nil && o.Excludes.IsSet()
}

// SetExcludes gets a reference to the given common.Nullable[]string and assigns it to the Excludes field.
func (o *SearchArchivesQuery) SetExcludes(v []string) {
	o.Excludes.Set(&v)
}

// SetExcludesNil sets the value for Excludes to be an explicit nil.
func (o *SearchArchivesQuery) SetExcludesNil() {
	o.Excludes.Set(nil)
}

// UnSetExcludes ensures that no value is present for Excludes, not even an explicit nil.
func (o *SearchArchivesQuery) UnSetExcludes() {
	o.Excludes.UnSet()
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetDateTimeRange() DateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret DateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DateTimeRange field.
func (o *SearchArchivesQuery) SetDateTimeRange(v DateTimeRange) {
	o.DateTimeRange = &v
}

// GetLgsIds returns the LgsIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchArchivesQuery) GetLgsIds() []string {
	if o == nil || o.LgsIDs.Get() == nil {
		var ret []string
		return ret
	}
	return *o.LgsIDs.Get()
}

// GetLgsIDsOk returns a tuple with the LgsIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchArchivesQuery) GetLgsIDsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LgsIDs.Get(), o.LgsIDs.IsSet()
}

// HasLgsIDs returns a boolean if a LgsIds has been set.
func (o *SearchArchivesQuery) HasLgsIDs() bool {
	return o != nil && o.LgsIDs.IsSet()
}

// SetLgsIDs gets a reference to the given common.Nullable[]string and assigns it to the LgsIDs field.
func (o *SearchArchivesQuery) SetLgsIDs(v []string) {
	o.LgsIDs.Set(&v)
}

// SetLgsIDsNil sets the value for LgsIDs to be an explicit nil.
func (o *SearchArchivesQuery) SetLgsIDsNil() {
	o.LgsIDs.Set(nil)
}

// UnSetLgsIDs ensures that no value is present for LgsIDs, not even an explicit nil.
func (o *SearchArchivesQuery) UnSetLgsIDs() {
	o.LgsIDs.UnSet()
}

// GetLgsNames returns the LgsNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchArchivesQuery) GetLgsNames() []string {
	if o == nil || o.LgsNames.Get() == nil {
		var ret []string
		return ret
	}
	return *o.LgsNames.Get()
}

// GetLgsNamesOk returns a tuple with the LgsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchArchivesQuery) GetLgsNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LgsNames.Get(), o.LgsNames.IsSet()
}

// HasLgsNames returns a boolean if a LgsNames has been set.
func (o *SearchArchivesQuery) HasLgsNames() bool {
	return o != nil && o.LgsNames.IsSet()
}

// SetLgsNames gets a reference to the given common.Nullable[]string and assigns it to the LgsNames field.
func (o *SearchArchivesQuery) SetLgsNames(v []string) {
	o.LgsNames.Set(&v)
}

// SetLgsNamesNil sets the value for LgsNames to be an explicit nil.
func (o *SearchArchivesQuery) SetLgsNamesNil() {
	o.LgsNames.Set(nil)
}

// UnSetLgsNames ensures that no value is present for LgsNames, not even an explicit nil.
func (o *SearchArchivesQuery) UnSetLgsNames() {
	o.LgsNames.UnSet()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SearchArchivesQuery) SetSize(v int64) {
	o.Size = &v
}

// GetDeleteAfterDay returns the DeleteAfterDay field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetDeleteAfterDay() int64 {
	if o == nil || o.DeleteAfterDay == nil {
		var ret int64
		return ret
	}
	return *o.DeleteAfterDay
}

// GetDeleteAfterDayOk returns a tuple with the DeleteAfterDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetDeleteAfterDayOk() (*int64, bool) {
	if o == nil || o.DeleteAfterDay == nil {
		return nil, false
	}
	return o.DeleteAfterDay, true
}

// HasDeleteAfterDay returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasDeleteAfterDay() bool {
	return o != nil && o.DeleteAfterDay != nil
}

// SetDeleteAfterDay gets a reference to the given int64 and assigns it to the DeleteAfterDay field.
func (o *SearchArchivesQuery) SetDeleteAfterDay(v int64) {
	o.DeleteAfterDay = &v
}

// GetParallelCount returns the ParallelCount field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetParallelCount() int64 {
	if o == nil || o.ParallelCount == nil {
		var ret int64
		return ret
	}
	return *o.ParallelCount
}

// GetParallelCountOk returns a tuple with the ParallelCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetParallelCountOk() (*int64, bool) {
	if o == nil || o.ParallelCount == nil {
		return nil, false
	}
	return o.ParallelCount, true
}

// HasParallelCount returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasParallelCount() bool {
	return o != nil && o.ParallelCount != nil
}

// SetParallelCount gets a reference to the given int64 and assigns it to the ParallelCount field.
func (o *SearchArchivesQuery) SetParallelCount(v int64) {
	o.ParallelCount = &v
}

// GetSearchOnRawLogs returns the SearchOnRawLogs field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetSearchOnRawLogs() bool {
	if o == nil || o.SearchOnRawLogs == nil {
		var ret bool
		return ret
	}
	return *o.SearchOnRawLogs
}

// GetSearchOnRawLogsOk returns a tuple with the SearchOnRawLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetSearchOnRawLogsOk() (*bool, bool) {
	if o == nil || o.SearchOnRawLogs == nil {
		return nil, false
	}
	return o.SearchOnRawLogs, true
}

// HasSearchOnRawLogs returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasSearchOnRawLogs() bool {
	return o != nil && o.SearchOnRawLogs != nil
}

// SetSearchOnRawLogs gets a reference to the given bool and assigns it to the SearchOnRawLogs field.
func (o *SearchArchivesQuery) SetSearchOnRawLogs(v bool) {
	o.SearchOnRawLogs = &v
}

// GetSearchActiveLgsOnly returns the SearchActiveLgsOnly field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetSearchActiveLgsOnly() bool {
	if o == nil || o.SearchActiveLgsOnly == nil {
		var ret bool
		return ret
	}
	return *o.SearchActiveLgsOnly
}

// GetSearchActiveLgsOnlyOk returns a tuple with the SearchActiveLgsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetSearchActiveLgsOnlyOk() (*bool, bool) {
	if o == nil || o.SearchActiveLgsOnly == nil {
		return nil, false
	}
	return o.SearchActiveLgsOnly, true
}

// HasSearchActiveLgsOnly returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasSearchActiveLgsOnly() bool {
	return o != nil && o.SearchActiveLgsOnly != nil
}

// SetSearchActiveLgsOnly gets a reference to the given bool and assigns it to the SearchActiveLgsOnly field.
func (o *SearchArchivesQuery) SetSearchActiveLgsOnly(v bool) {
	o.SearchActiveLgsOnly = &v
}

// GetPauseRequest returns the PauseRequest field value if set, zero value otherwise.
func (o *SearchArchivesQuery) GetPauseRequest() bool {
	if o == nil || o.PauseRequest == nil {
		var ret bool
		return ret
	}
	return *o.PauseRequest
}

// GetPauseRequestOk returns a tuple with the PauseRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesQuery) GetPauseRequestOk() (*bool, bool) {
	if o == nil || o.PauseRequest == nil {
		return nil, false
	}
	return o.PauseRequest, true
}

// HasPauseRequest returns a boolean if a field has been set.
func (o *SearchArchivesQuery) HasPauseRequest() bool {
	return o != nil && o.PauseRequest != nil
}

// SetPauseRequest gets a reference to the given bool and assigns it to the PauseRequest field.
func (o *SearchArchivesQuery) SetPauseRequest(v bool) {
	o.PauseRequest = &v
}

// GetReportColNames returns the ReportColNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchArchivesQuery) GetReportColNames() []string {
	if o == nil || o.ReportColNames.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ReportColNames.Get()
}

// GetReportColNamesOk returns a tuple with the ReportColNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchArchivesQuery) GetReportColNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportColNames.Get(), o.ReportColNames.IsSet()
}

// HasReportColNames returns a boolean if a ReportColNames has been set.
func (o *SearchArchivesQuery) HasReportColNames() bool {
	return o != nil && o.ReportColNames.IsSet()
}

// SetReportColNames gets a reference to the given common.Nullable[]string and assigns it to the ReportColNames field.
func (o *SearchArchivesQuery) SetReportColNames(v []string) {
	o.ReportColNames.Set(&v)
}

// SetReportColNamesNil sets the value for ReportColNames to be an explicit nil.
func (o *SearchArchivesQuery) SetReportColNamesNil() {
	o.ReportColNames.Set(nil)
}

// UnSetReportColNames ensures that no value is present for ReportColNames, not even an explicit nil.
func (o *SearchArchivesQuery) UnSetReportColNames() {
	o.ReportColNames.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o SearchArchivesQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["Id"] = o.ID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.CreateDate != nil {
		toSerialize["CreateDate"] = o.CreateDate
	}
	if o.Includes.IsSet() {
		toSerialize["Includes"] = o.Includes.Get()
	}
	if o.Excludes.IsSet() {
		toSerialize["Excludes"] = o.Excludes.Get()
	}
	if o.DateTimeRange != nil {
		toSerialize["DateTimeRange"] = o.DateTimeRange
	}
	if o.LgsIDs.IsSet() {
		toSerialize["LgsIds"] = o.LgsIDs.Get()
	}
	if o.LgsNames.IsSet() {
		toSerialize["LgsNames"] = o.LgsNames.Get()
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.DeleteAfterDay != nil {
		toSerialize["DeleteAfterDay"] = o.DeleteAfterDay
	}
	if o.ParallelCount != nil {
		toSerialize["ParallelCount"] = o.ParallelCount
	}
	if o.SearchOnRawLogs != nil {
		toSerialize["SearchOnRawLogs"] = o.SearchOnRawLogs
	}
	if o.SearchActiveLgsOnly != nil {
		toSerialize["SearchActiveLgsOnly"] = o.SearchActiveLgsOnly
	}
	if o.PauseRequest != nil {
		toSerialize["PauseRequest"] = o.PauseRequest
	}
	if o.ReportColNames.IsSet() {
		toSerialize["SuccessItems"] = o.ReportColNames.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SearchArchivesQuery) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID                  *string                     `json:"Id,omitempty"`
		Name                *string                     `json:"Name,omitempty"`
		CreateDate          *string                     `json:"CreateDate,omitempty"`
		Includes            common.NullableList[string] `json:"Includes,omitempty"`
		Excludes            common.NullableList[string] `json:"Excludes,omitempty"`
		DateTimeRange       *DateTimeRange              `json:"DateTimeRange,omitempty"`
		LgsIDs              common.NullableList[string] `json:"LgsIds,omitempty"`
		LgsNames            common.NullableList[string] `json:"LgsNames,omitempty"`
		Size                *int64                      `json:"Size,omitempty"`
		DeleteAfterDay      *int64                      `json:"DeleteAfterDay,omitempty"`
		ParallelCount       *int64                      `json:"ParallelCount,omitempty"`
		SearchOnRawLogs     *bool                       `json:"SearchOnRawLogs,omitempty"`
		SearchActiveLgsOnly *bool                       `json:"SearchActiveLgsOnly,omitempty"`
		PauseRequest        *bool                       `json:"PauseRequest,omitempty"`
		ReportColNames      common.NullableList[string] `json:"ReportColNames,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Id", "Name", "CreateDate", "Includes", "Excludes", "DateTimeRange", "LgsIds", "LgsNames", "Size", "DeleteAfterDay", "ParallelCount", "SearchOnRawLogs", "SearchActiveLgsOnly", "PauseRequest", "ReportColNames"})
	} else {
		return err
	}

	o.ID = all.ID
	o.Name = all.Name
	o.CreateDate = all.CreateDate
	o.Includes = all.Includes
	o.Excludes = all.Excludes
	hasInvalidField := false
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	o.LgsIDs = all.LgsIDs
	o.LgsNames = all.LgsNames
	o.Size = all.Size
	o.DeleteAfterDay = all.DeleteAfterDay
	o.ParallelCount = all.ParallelCount
	o.SearchOnRawLogs = all.SearchOnRawLogs
	o.SearchActiveLgsOnly = all.SearchActiveLgsOnly
	o.PauseRequest = all.PauseRequest
	o.ReportColNames = all.ReportColNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
