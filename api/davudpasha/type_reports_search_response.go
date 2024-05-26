package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ReportsSearchResponse struct {
	ModifiedDate            *string                     `json:"ModifiedDate,omitempty"`
	ReportId                *string                     `jsom:"ReportId,omitempty"`
	Username                *string                     `json:"Username,omitempty"`
	Name                    *string                     `json:"Name,omitempty"`
	Description             *string                     `json:"Description,omitempty"`
	IsActive                *bool                       `json:"IsActive,omitempty"`
	Author                  common.NullableString       `json:"Author,omitempty"`
	ReportLink              *string                     `json:"ReportLink,omitempty"`
	SharedUsersAndGroups    common.NullableList[string] `json:"SharedUsersAndGroups,omitempty"`
	CreatedDate             *string                     `json:"CreatedDate,omitempty"`
	ReportQuery             []ReportsQuery              `json:"ReportQuery,omitempty"`
	ReportData              *ReportsData                `json:"ReportData,omitempty"`
	Schedule                *ScheduleConfig             `json:"Schedule,omitempty"`
	LastGenerationTime      common.NullableString       `json:"LastGenerationTime,omitempty"`
	NextGenerationTime      common.NullableString       `json:"NextGenerationTime,omitempty"`
	LatestReportFile        *ReportsLatestReportFile    `json:"LatestReportFile,omitempty"`
	PageSettings            common.NullableString       `json:"PageSettings,omitempty"`
	Parameters              *ReportsParameters          `json:"Parameters,omitempty"`
	Tags                    []string                    `json:"Tags,omitempty"`
	Actions                 *Actions                    `json:"Actions,omitempty"`
	ModuleFilter            *string                     `json:"ModuleFilter,omitempty"`
	RemoteInterfaceName     *string                     `json:"RemoteInterfaceName,omitempty"`
	RemoteMethodName        *string                     `json:"RemoteMethodName,omitempty"`
	ApplicationName         *string                     `json:"ApplicationName,omitempty"`
	SubExecutorModuleFilter *string                     `json:"SubExecutorModuleFilter,omitempty"`
	Version                 *float64                    `json:"Version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewReportsSearchResponse creates a new ReportsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsSearchResponse() *ReportsSearchResponse {
	this := ReportsSearchResponse{}
	return &this
}

// NewReportsSearchResponseWithDefaults creates a new ReportsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsSearchResponseWithDefaults() *ReportsSearchResponse {
	this := ReportsSearchResponse{}
	return &this
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetModifiedDate() string {
	if o == nil || o.ModifiedDate == nil {
		var ret string
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetModifiedDateOk() (*string, bool) {
	if o == nil || o.ModifiedDate == nil {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasModifiedDate() bool {
	return o != nil && o.ModifiedDate != nil
}

// SetModifiedDate gets a reference to the given string and assigns it to the ModifiedDate field.
func (o *ReportsSearchResponse) SetModifiedDate(v string) {
	o.ModifiedDate = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetReportId() string {
	if o == nil || o.ReportId == nil {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetReportIdOk() (*string, bool) {
	if o == nil || o.ReportId == nil {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasReportId() bool {
	return o != nil && o.ReportId != nil
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *ReportsSearchResponse) SetReportId(v string) {
	o.ReportId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ReportsSearchResponse) SetUsername(v string) {
	o.Username = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReportsSearchResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReportsSearchResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasIsActive() bool {
	return o != nil && o.IsActive != nil
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ReportsSearchResponse) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsSearchResponse) GetAuthor() string {
	if o == nil || o.Author.Get() == nil {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsSearchResponse) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasAuthor() bool {
	return o != nil && o.Author.IsSet()
}

// SetAuthor gets a reference to the given datadog.NullableString and assigns it to the Author field.
func (o *ReportsSearchResponse) SetAuthor(v string) {
	o.Author.Set(&v)
}

// SetAuthorNil sets the value for Author to be an explicit nil.
func (o *ReportsSearchResponse) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil.
func (o *ReportsSearchResponse) UnSetAuthor() {
	o.Author.UnSet()
}

// GetReportLink returns the ReportLink field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetReportLink() string {
	if o == nil || o.ReportLink == nil {
		var ret string
		return ret
	}
	return *o.ReportLink
}

// GetReportLinkOk returns a tuple with the ReportLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetReportLinkOk() (*string, bool) {
	if o == nil || o.ReportLink == nil {
		return nil, false
	}
	return o.ReportLink, true
}

// HasReportLink returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasReportLink() bool {
	return o != nil && o.ReportLink != nil
}

// SetReportLink gets a reference to the given string and assigns it to the ReportLink field.
func (o *ReportsSearchResponse) SetReportLink(v string) {
	o.ReportLink = &v
}

// GetSharedUsersAndGroups returns a tuple with the SharedUsersAndGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsSearchResponse) GetSharedUsersAndGroups() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedUsersAndGroups.Get(), o.SharedUsersAndGroups.IsSet()
}

// GetSharedUsersAndGroupsOk returns a tuple with the SharedUsersAndGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsSearchResponse) GetSharedUsersAndGroupsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedUsersAndGroups.Get(), o.SharedUsersAndGroups.IsSet()
}

// HasSharedUsersAndGroups returns a boolean if a SharedUsersAndGroups has been set.
func (o *ReportsSearchResponse) HasSharedUsersAndGroups() bool {
	return o != nil && o.SharedUsersAndGroups.IsSet()
}

// SetSharedUsersAndGroups gets a reference to the given datadog.Nullable[]string and assigns it to the SharedUsersAndGroups field.
func (o *ReportsSearchResponse) SetSharedUsersAndGroups(v []string) {
	o.SharedUsersAndGroups.Set(&v)
}

// SetSharedUsersAndGroupsNil sets the value for SharedUsersAndGroups to be an explicit nil.
func (o *ReportsSearchResponse) SetSharedUsersAndGroupsNil() {
	o.SharedUsersAndGroups.Set(nil)
}

// UnsetSharedUsersAndGroups ensures that no value is present for SharedUsersAndGroups, not even an explicit nil.
func (o *ReportsSearchResponse) UnSetSharedUsersAndGroups() {
	o.SharedUsersAndGroups.UnSet()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetCreatedDate() string {
	if o == nil || o.CreatedDate == nil {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetCreatedDateOk() (*string, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasCreatedDate() bool {
	return o != nil && o.CreatedDate != nil
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ReportsSearchResponse) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetReportQuery returns the ReportQuery field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetReportQuery() []ReportsQuery {
	if o == nil || o.ReportQuery == nil {
		var ret []ReportsQuery
		return ret
	}
	return o.ReportQuery
}

// GetReportQueryOk returns a tuple with the ReportQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetReportQueryOk() (*[]ReportsQuery, bool) {
	if o == nil || o.ReportQuery == nil {
		return nil, false
	}
	return &o.ReportQuery, true
}

// HasReportQuery returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasReportQuery() bool {
	return o != nil && o.ReportQuery != nil
}

// SetReportQuery gets a reference to the given []ReportsQuery and assigns it to the ReportQuery field.
func (o *ReportsSearchResponse) SetReportQuery(v []ReportsQuery) {
	o.ReportQuery = v
}

// GetReportData returns the ReportData field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetReportData() ReportsData {
	if o == nil || o.ReportData == nil {
		var ret ReportsData
		return ret
	}
	return *o.ReportData
}

// GetReportDataOk returns a tuple with the ReportData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetReportDataOk() (*ReportsData, bool) {
	if o == nil || o.ReportData == nil {
		return nil, false
	}
	return o.ReportData, true
}

// HasReportData returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasReportData() bool {
	return o != nil && o.ReportData != nil
}

// SetReportData gets a reference to the given ReportsData and assigns it to the ReportData field.
func (o *ReportsSearchResponse) SetReportData(v ReportsData) {
	o.ReportData = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetSchedule() ScheduleConfig {
	if o == nil || o.Schedule == nil {
		var ret ScheduleConfig
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetScheduleOk() (*ScheduleConfig, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given ScheduleConfig and assigns it to the Schedule field.
func (o *ReportsSearchResponse) SetSchedule(v ScheduleConfig) {
	o.Schedule = &v
}

// GetLastGenerationTime returns the LastGenerationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsSearchResponse) GetLastGenerationTime() string {
	if o == nil || o.LastGenerationTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastGenerationTime.Get()
}

// GetLastGenerationTimeOk returns a tuple with the LastGenerationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsSearchResponse) GetLastGenerationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastGenerationTime.Get(), o.LastGenerationTime.IsSet()
}

// HasLastGenerationTime returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasLastGenerationTime() bool {
	return o != nil && o.LastGenerationTime.IsSet()
}

// SetLastGenerationTime gets a reference to the given datadog.NullableString and assigns it to the LastGenerationTime field.
func (o *ReportsSearchResponse) SetLastGenerationTime(v string) {
	o.LastGenerationTime.Set(&v)
}

// SetLastGenerationTimeNil sets the value for LastGenerationTime to be an explicit nil.
func (o *ReportsSearchResponse) SetLastGenerationTimeNil() {
	o.LastGenerationTime.Set(nil)
}

// UnsetLastGenerationTime ensures that no value is present for LastGenerationTime, not even an explicit nil.
func (o *ReportsSearchResponse) UnSetLastGenerationTime() {
	o.LastGenerationTime.UnSet()
}

// GetNextGenerationTime returns the NextGenerationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsSearchResponse) GetNextGenerationTime() string {
	if o == nil || o.NextGenerationTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextGenerationTime.Get()
}

// GetNextGenerationTimeOk returns a tuple with the NextGenerationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsSearchResponse) GetNextGenerationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextGenerationTime.Get(), o.NextGenerationTime.IsSet()
}

// HasNextGenerationTime returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasNextGenerationTime() bool {
	return o != nil && o.NextGenerationTime.IsSet()
}

// SetNextGenerationTime gets a reference to the given datadog.NullableString and assigns it to the NextGenerationTime field.
func (o *ReportsSearchResponse) SetNextGenerationTime(v string) {
	o.NextGenerationTime.Set(&v)
}

// SetNextGenerationTimeNil sets the value for NextGenerationTime to be an explicit nil.
func (o *ReportsSearchResponse) SetNextGenerationTimeNil() {
	o.NextGenerationTime.Set(nil)
}

// UnsetNextGenerationTime ensures that no value is present for NextGenerationTime, not even an explicit nil.
func (o *ReportsSearchResponse) UnSetNextGenerationTime() {
	o.NextGenerationTime.UnSet()
}

// GetLatestReportFile returns the LatestReportFile field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetLatestReportFile() ReportsLatestReportFile {
	if o == nil || o.LatestReportFile == nil {
		var ret ReportsLatestReportFile
		return ret
	}
	return *o.LatestReportFile
}

// GetLatestReportFileOk returns a tuple with the LatestReportFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetLatestReportFileOk() (*ReportsLatestReportFile, bool) {
	if o == nil || o.LatestReportFile == nil {
		return nil, false
	}
	return o.LatestReportFile, true
}

// HasLatestReportFile returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasLatestReportFile() bool {
	return o != nil && o.LatestReportFile != nil
}

// SetLatestReportFile gets a reference to the given ReportsLatestReportFile and assigns it to the LatestReportFile field.
func (o *ReportsSearchResponse) SetLatestReportFile(v ReportsLatestReportFile) {
	o.LatestReportFile = &v
}

// GetPageSettings returns the PageSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsSearchResponse) GetPageSettings() string {
	if o == nil || o.PageSettings.Get() == nil {
		var ret string
		return ret
	}
	return *o.PageSettings.Get()
}

// GetPageSettingsOk returns a tuple with the PageSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsSearchResponse) GetPageSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSettings.Get(), o.PageSettings.IsSet()
}

// HasPageSettings returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasPageSettings() bool {
	return o != nil && o.PageSettings.IsSet()
}

// SetPageSettings gets a reference to the given datadog.NullableString and assigns it to the PageSettings field.
func (o *ReportsSearchResponse) SetPageSettings(v string) {
	o.PageSettings.Set(&v)
}

// SetPageSettingsNil sets the value for PageSettings to be an explicit nil.
func (o *ReportsSearchResponse) SetPageSettingsNil() {
	o.PageSettings.Set(nil)
}

// UnsetPageSettings ensures that no value is present for PageSettings, not even an explicit nil.
func (o *ReportsSearchResponse) UnSetPageSettings() {
	o.PageSettings.UnSet()
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetParameters() ReportsParameters {
	if o == nil || o.Parameters == nil {
		var ret ReportsParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetParametersOk() (*ReportsParameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given ReportsParameters and assigns it to the Parameters field.
func (o *ReportsSearchResponse) SetParameters(v ReportsParameters) {
	o.Parameters = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ReportsSearchResponse) SetTags(v []string) {
	o.Tags = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetActions() Actions {
	if o == nil || o.Actions == nil {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetActionsOk() (*Actions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *ReportsSearchResponse) SetActions(v Actions) {
	o.Actions = &v
}

// GetModuleFilter returns the ModuleFilter field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetModuleFilter() string {
	if o == nil || o.ModuleFilter == nil {
		var ret string
		return ret
	}
	return *o.ModuleFilter
}

// GetModuleFilterOk returns a tuple with the ModuleFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetModuleFilterOk() (*string, bool) {
	if o == nil || o.ModuleFilter == nil {
		return nil, false
	}
	return o.ModuleFilter, true
}

// HasModuleFilter returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasModuleFilter() bool {
	return o != nil && o.ModuleFilter != nil
}

// SetModuleFilter gets a reference to the given string and assigns it to the ModuleFilter field.
func (o *ReportsSearchResponse) SetModuleFilter(v string) {
	o.ModuleFilter = &v
}

// GetRemoteInterfaceName returns the RemoteInterfaceName field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetRemoteInterfaceName() string {
	if o == nil || o.RemoteInterfaceName == nil {
		var ret string
		return ret
	}
	return *o.RemoteInterfaceName
}

// GetRemoteInterfaceNameOk returns a tuple with the RemoteInterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetRemoteInterfaceNameOk() (*string, bool) {
	if o == nil || o.RemoteInterfaceName == nil {
		return nil, false
	}
	return o.RemoteInterfaceName, true
}

// HasRemoteInterfaceName returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasRemoteInterfaceName() bool {
	return o != nil && o.RemoteInterfaceName != nil
}

// SetRemoteInterfaceName gets a reference to the given string and assigns it to the RemoteInterfaceName field.
func (o *ReportsSearchResponse) SetRemoteInterfaceName(v string) {
	o.RemoteInterfaceName = &v
}

// GetRemoteMethodName returns the RemoteMethodName field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetRemoteMethodName() string {
	if o == nil || o.RemoteMethodName == nil {
		var ret string
		return ret
	}
	return *o.RemoteMethodName
}

// GetRemoteMethodNameOk returns a tuple with the RemoteMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetRemoteMethodNameOk() (*string, bool) {
	if o == nil || o.RemoteMethodName == nil {
		return nil, false
	}
	return o.RemoteMethodName, true
}

// HasRemoteMethodName returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasRemoteMethodName() bool {
	return o != nil && o.RemoteMethodName != nil
}

// SetRemoteMethodName gets a reference to the given string and assigns it to the RemoteMethodName field.
func (o *ReportsSearchResponse) SetRemoteMethodName(v string) {
	o.RemoteMethodName = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasApplicationName() bool {
	return o != nil && o.ApplicationName != nil
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *ReportsSearchResponse) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetSubExecutorModuleFilter returns the SubExecutorModuleFilter field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetSubExecutorModuleFilter() string {
	if o == nil || o.SubExecutorModuleFilter == nil {
		var ret string
		return ret
	}
	return *o.SubExecutorModuleFilter
}

// GetSubExecutorModuleFilterOk returns a tuple with the SubExecutorModuleFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetSubExecutorModuleFilterOk() (*string, bool) {
	if o == nil || o.SubExecutorModuleFilter == nil {
		return nil, false
	}
	return o.SubExecutorModuleFilter, true
}

// HasSubExecutorModuleFilter returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasSubExecutorModuleFilter() bool {
	return o != nil && o.SubExecutorModuleFilter != nil
}

// SetSubExecutorModuleFilter gets a reference to the given string and assigns it to the SubExecutorModuleFilter field.
func (o *ReportsSearchResponse) SetSubExecutorModuleFilter(v string) {
	o.SubExecutorModuleFilter = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ReportsSearchResponse) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSearchResponse) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ReportsSearchResponse) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *ReportsSearchResponse) SetVersion(v float64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ModifiedDate != nil {
		toSerialize["ModifiedDate"] = o.ModifiedDate
	}
	if o.ReportId != nil {
		toSerialize["ReportId"] = o.ReportId
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.IsActive != nil {
		toSerialize["IsActive"] = o.IsActive
	}
	if o.Author.IsSet() {
		toSerialize["Author"] = o.Author.Get()
	}
	if o.ReportLink != nil {
		toSerialize["ReportLink"] = o.ReportLink
	}
	if o.SharedUsersAndGroups.IsSet() {
		toSerialize["SharedUsersAndGroups"] = o.SharedUsersAndGroups.Get()
	}
	if o.CreatedDate != nil {
		toSerialize["CreatedDate"] = o.CreatedDate
	}
	if o.ReportQuery != nil {
		toSerialize["ReportQuery"] = o.ReportQuery
	}
	if o.ReportData != nil {
		toSerialize["ReportData"] = o.ReportData
	}
	if o.Schedule != nil {
		toSerialize["Schedule"] = o.Schedule
	}
	if o.LastGenerationTime.IsSet() {
		toSerialize["LastGenerationTime"] = o.LastGenerationTime.Get()
	}
	if o.NextGenerationTime.IsSet() {
		toSerialize["NextGenerationTime"] = o.NextGenerationTime.Get()
	}
	if o.LatestReportFile != nil {
		toSerialize["LatestReportFile"] = o.LatestReportFile
	}
	if o.PageSettings.IsSet() {
		toSerialize["PageSettings"] = o.PageSettings.Get()
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.ModuleFilter != nil {
		toSerialize["ModuleFilter"] = o.ModuleFilter
	}
	if o.RemoteInterfaceName != nil {
		toSerialize["RemoteInterfaceName"] = o.RemoteInterfaceName
	}
	if o.RemoteMethodName != nil {
		toSerialize["RemoteMethodName"] = o.RemoteMethodName
	}
	if o.ApplicationName != nil {
		toSerialize["ApplicationName"] = o.ApplicationName
	}
	if o.SubExecutorModuleFilter != nil {
		toSerialize["SubExecutorModuleFilter"] = o.SubExecutorModuleFilter
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModifiedDate            *string                     `json:"ModifiedDate,omitempty"`
		ReportId                *string                     `jsom:"ReportId,omitempty"`
		Username                *string                     `json:"Username,omitempty"`
		Name                    *string                     `json:"Name,omitempty"`
		Description             *string                     `json:"Description,omitempty"`
		IsActive                *bool                       `json:"IsActive,omitempty"`
		Author                  common.NullableString       `json:"Author,omitempty"`
		ReportLink              *string                     `json:"ReportLink,omitempty"`
		SharedUsersAndGroups    common.NullableList[string] `json:"SharedUsersAndGroups,omitempty"`
		CreatedDate             *string                     `json:"CreatedDate,omitempty"`
		ReportQuery             []ReportsQuery              `json:"ReportQuery,omitempty"`
		ReportData              *ReportsData                `json:"ReportData,omitempty"`
		Schedule                *ScheduleConfig             `json:"Schedule,omitempty"`
		LastGenerationTime      common.NullableString       `json:"LastGenerationTime,omitempty"`
		NextGenerationTime      common.NullableString       `json:"NextGenerationTime,omitempty"`
		LatestReportFile        *ReportsLatestReportFile    `json:"LatestReportFile,omitempty"`
		PageSettings            common.NullableString       `json:"PageSettings,omitempty"`
		Parameters              *ReportsParameters          `json:"Parameters,omitempty"`
		Tags                    []string                    `json:"Tags,omitempty"`
		Actions                 *Actions                    `json:"Actions,omitempty"`
		ModuleFilter            *string                     `json:"ModuleFilter,omitempty"`
		RemoteInterfaceName     *string                     `json:"RemoteInterfaceName,omitempty"`
		RemoteMethodName        *string                     `json:"RemoteMethodName,omitempty"`
		ApplicationName         *string                     `json:"ApplicationName,omitempty"`
		SubExecutorModuleFilter *string                     `json:"SubExecutorModuleFilter,omitempty"`
		Version                 *float64                    `json:"Version,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ModifiedDate", "ReportId", "Username", "Name", "Description", "IsActive", "Author", "ReportLink", "SharedUsersAndGroups", "CreatedDate", "ReportQuery", "ReportData", "Schedule", "LastGenerationTime", "NextGenerationTime", "LatestReportFile", "PageSettings", "Parameters", "Tags", "Actions", "ModuleFilter", "RemoteInterfaceName", "RemoteMethodName", "ApplicationName", "SubExecutorModuleFilter", "Version"})
	} else {
		return err
	}

	o.ModifiedDate = all.ModifiedDate
	o.ReportId = all.ReportId
	o.Username = all.Username
	o.Name = all.Name
	o.Description = all.Description
	o.IsActive = all.IsActive
	o.Author = all.Author
	o.ReportLink = all.ReportLink
	o.SharedUsersAndGroups = all.SharedUsersAndGroups
	o.CreatedDate = all.CreatedDate
	o.ReportQuery = all.ReportQuery
	hasInvalidField := false
	if all.ReportData != nil && all.ReportData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReportData = all.ReportData
	o.Schedule = all.Schedule
	o.LastGenerationTime = all.LastGenerationTime
	o.NextGenerationTime = all.NextGenerationTime
	o.LatestReportFile = all.LatestReportFile
	o.PageSettings = all.PageSettings
	o.Parameters = all.Parameters
	o.Tags = all.Tags
	o.Actions = all.Actions
	o.ModuleFilter = all.ModuleFilter
	o.RemoteInterfaceName = all.RemoteInterfaceName
	o.RemoteMethodName = all.RemoteMethodName
	o.ApplicationName = all.ApplicationName
	o.SubExecutorModuleFilter = all.SubExecutorModuleFilter
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
