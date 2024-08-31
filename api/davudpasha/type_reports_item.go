package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsItem represents a response from a report search operation.
type ReportsItem struct {
	// Date when the report was last modified.
	ModifiedDate *string `json:"ModifiedDate,omitempty"`
	// ID of the report.
	ReportID *string `json:"ReportId,omitempty"`
	// ID of the report.
	ID *string `json:"ID,omitempty"`
	// Username associated with the report.
	Username *string `json:"Username,omitempty"`
	// Name of the report.
	Name *string `json:"Name,omitempty"`
	// Description of the report.
	Description common.NullableString `json:"Description,omitempty"`
	// Indicates if the report is active.
	IsActive *bool `json:"IsActive,omitempty"`
	// Author of the report.
	Author common.NullableString `json:"Author,omitempty"`
	// Link to access the report.
	ReportLink *string `json:"ReportLink,omitempty"`
	// Users and groups with whom the report is shared.
	SharedUsersAndGroups common.NullableList[string] `json:"SharedUsersAndGroups,omitempty"`
	// Date when the report was created.
	CreatedDate *string `json:"CreatedDate,omitempty"`
	// Queries associated with the report.
	ReportQuery []ReportsQuery `json:"ReportQuery,omitempty"`
	// Data related to the report.
	ReportData *ReportsData `json:"ReportData,omitempty"`
	// Schedule configuration for generating the report.
	Schedule NullableScheduleConfig `json:"Schedule,omitempty"`
	// Email configuration for the report.
	EmailConfigs *ReportsEmailConfigs `json:"EmailConfigs,omitempty"`
	// Date and time when the report was last generated.
	LastGenerationTime common.NullableString `json:"LastGenerationTime,omitempty"`
	// Date and time when the report is scheduled to be generated next.
	NextGenerationTime common.NullableString `json:"NextGenerationTime,omitempty"`
	// Latest file associated with the report.
	LatestReportFile NullableReportsLatestReportFile `json:"LatestReportFile,omitempty"`
	// Page settings for the report.
	PageSettings common.NullableString `json:"PageSettings,omitempty"`
	// Parameters for the report.
	Parameters *ReportsParameters `json:"Parameters,omitempty"`
	// Tags associated with the report.
	Tags common.NullableList[string] `json:"Tags,omitempty"`
	// Actions associated with the report.
	Actions *Actions `json:"Actions,omitempty"`
	// Filter for the module.
	ModuleFilter *string `json:"ModuleFilter,omitempty"`
	// Name of the remote interface.
	RemoteInterfaceName *string `json:"RemoteInterfaceName,omitempty"`
	// Name of the remote method.
	RemoteMethodName *string `json:"RemoteMethodName,omitempty"`
	// Name of the application associated with the report.
	ApplicationName *string `json:"ApplicationName,omitempty"`
	// Filter for the sub-executor module.
	SubExecutorModuleFilter *string `json:"SubExecutorModuleFilter,omitempty"`
	// Version of the report.
	Version        *float64 `json:"Version,omitempty"`
	CheckBoxWeekly []string `json:"chechboxWeekly,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsItem creates a new ReportsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsItem() *ReportsItem {
	this := ReportsItem{}
	return &this
}

// NewReportsItemWithDefaults creates a new ReportsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsItemWithDefaults() *ReportsItem {
	this := ReportsItem{}
	return &this
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *ReportsItem) GetModifiedDate() string {
	if o == nil || o.ModifiedDate == nil {
		var ret string
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetModifiedDateOk() (*string, bool) {
	if o == nil || o.ModifiedDate == nil {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ReportsItem) HasModifiedDate() bool {
	return o != nil && o.ModifiedDate != nil
}

// SetModifiedDate gets a reference to the given string and assigns it to the ModifiedDate field.
func (o *ReportsItem) SetModifiedDate(v string) {
	o.ModifiedDate = &v
}

// GetReportID returns the ReportID field value if set, zero value otherwise.
func (o *ReportsItem) GetReportID() string {
	if o == nil || o.ReportID == nil {
		var ret string
		return ret
	}
	return *o.ReportID
}

// GetReportIDOk returns a tuple with the ReportID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetReportIDOk() (*string, bool) {
	if o == nil || o.ReportID == nil {
		return nil, false
	}
	return o.ReportID, true
}

// HasReportID returns a boolean if a field has been set.
func (o *ReportsItem) HasReportID() bool {
	return o != nil && o.ReportID != nil
}

// SetReportID gets a reference to the given string and assigns it to the ReportID field.
func (o *ReportsItem) SetReportID(v string) {
	o.ReportID = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ReportsItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ReportsItem) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ReportsItem) SetID(v string) {
	o.ID = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ReportsItem) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ReportsItem) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ReportsItem) SetUsername(v string) {
	o.Username = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReportsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReportsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReportsItem) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ReportsItem) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *ReportsItem) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ReportsItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnSetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ReportsItem) UnsetDescription() {
	o.Description.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ReportsItem) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ReportsItem) HasIsActive() bool {
	return o != nil && o.IsActive != nil
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ReportsItem) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsItem) GetAuthor() string {
	if o == nil || o.Author.Get() == nil {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *ReportsItem) HasAuthor() bool {
	return o != nil && o.Author.IsSet()
}

// SetAuthor gets a reference to the given common.NullableString and assigns it to the Author field.
func (o *ReportsItem) SetAuthor(v string) {
	o.Author.Set(&v)
}

// SetAuthorNil sets the value for Author to be an explicit nil.
func (o *ReportsItem) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnSetAuthor ensures that no value is present for Author, not even an explicit nil.
func (o *ReportsItem) UnsetAuthor() {
	o.Author.Unset()
}

// GetReportLink returns the ReportLink field value if set, zero value otherwise.
func (o *ReportsItem) GetReportLink() string {
	if o == nil || o.ReportLink == nil {
		var ret string
		return ret
	}
	return *o.ReportLink
}

// GetReportLinkOk returns a tuple with the ReportLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetReportLinkOk() (*string, bool) {
	if o == nil || o.ReportLink == nil {
		return nil, false
	}
	return o.ReportLink, true
}

// HasReportLink returns a boolean if a field has been set.
func (o *ReportsItem) HasReportLink() bool {
	return o != nil && o.ReportLink != nil
}

// SetReportLink gets a reference to the given string and assigns it to the ReportLink field.
func (o *ReportsItem) SetReportLink(v string) {
	o.ReportLink = &v
}

// GetSharedUsersAndGroups returns a tuple with the SharedUsersAndGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetSharedUsersAndGroups() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedUsersAndGroups.Get(), o.SharedUsersAndGroups.IsSet()
}

// GetSharedUsersAndGroupsOk returns a tuple with the SharedUsersAndGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetSharedUsersAndGroupsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedUsersAndGroups.Get(), o.SharedUsersAndGroups.IsSet()
}

// HasSharedUsersAndGroups returns a boolean if a SharedUsersAndGroups has been set.
func (o *ReportsItem) HasSharedUsersAndGroups() bool {
	return o != nil && o.SharedUsersAndGroups.IsSet()
}

// SetSharedUsersAndGroups gets a reference to the given common.Nullable[]string and assigns it to the SharedUsersAndGroups field.
func (o *ReportsItem) SetSharedUsersAndGroups(v []string) {
	o.SharedUsersAndGroups.Set(&v)
}

// SetSharedUsersAndGroupsNil sets the value for SharedUsersAndGroups to be an explicit nil.
func (o *ReportsItem) SetSharedUsersAndGroupsNil() {
	o.SharedUsersAndGroups.Set(nil)
}

// UnSetSharedUsersAndGroups ensures that no value is present for SharedUsersAndGroups, not even an explicit nil.
func (o *ReportsItem) UnsetSharedUsersAndGroups() {
	o.SharedUsersAndGroups.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ReportsItem) GetCreatedDate() string {
	if o == nil || o.CreatedDate == nil {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetCreatedDateOk() (*string, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ReportsItem) HasCreatedDate() bool {
	return o != nil && o.CreatedDate != nil
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ReportsItem) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetReportQuery returns the ReportQuery field value if set, zero value otherwise.
func (o *ReportsItem) GetReportQuery() []ReportsQuery {
	if o == nil || o.ReportQuery == nil {
		var ret []ReportsQuery
		return ret
	}
	return o.ReportQuery
}

// GetReportQueryOk returns a tuple with the ReportQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetReportQueryOk() (*[]ReportsQuery, bool) {
	if o == nil || o.ReportQuery == nil {
		return nil, false
	}
	return &o.ReportQuery, true
}

// HasReportQuery returns a boolean if a field has been set.
func (o *ReportsItem) HasReportQuery() bool {
	return o != nil && o.ReportQuery != nil
}

// SetReportQuery gets a reference to the given []ReportsQuery and assigns it to the ReportQuery field.
func (o *ReportsItem) SetReportQuery(v []ReportsQuery) {
	o.ReportQuery = v
}

// GetReportData returns the ReportData field value if set, zero value otherwise.
func (o *ReportsItem) GetReportData() ReportsData {
	if o == nil || o.ReportData == nil {
		var ret ReportsData
		return ret
	}
	return *o.ReportData
}

// GetReportDataOk returns a tuple with the ReportData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetReportDataOk() (*ReportsData, bool) {
	if o == nil || o.ReportData == nil {
		return nil, false
	}
	return o.ReportData, true
}

// HasReportData returns a boolean if a field has been set.
func (o *ReportsItem) HasReportData() bool {
	return o != nil && o.ReportData != nil
}

// SetReportData gets a reference to the given ReportsData and assigns it to the ReportData field.
func (o *ReportsItem) SetReportData(v ReportsData) {
	o.ReportData = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsItem) GetSchedule() ScheduleConfig {
	if o == nil || o.Schedule.Get() == nil {
		var ret ScheduleConfig
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetScheduleOk() (*ScheduleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a Schedule has been set.
func (o *ReportsItem) HasSchedule() bool {
	return o != nil && o.Schedule.IsSet()
}

// SetSchedule gets a reference to the given common.NullableScheduleConfig and assigns it to the Schedule field.
func (o *ReportsItem) SetSchedule(v ScheduleConfig) {
	o.Schedule.Set(&v)
}

// SetScheduleNil sets the value for Schedule to be an explicit nil.
func (o *ReportsItem) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnSetSchedule ensures that no value is present for Schedule, not even an explicit nil.
func (o *ReportsItem) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetEmailConfigs returns the EmailConfigs field value if set, zero value otherwise.
func (o *ReportsItem) GetEmailConfigs() ReportsEmailConfigs {
	if o == nil || o.EmailConfigs == nil {
		var ret ReportsEmailConfigs
		return ret
	}
	return *o.EmailConfigs
}

// GetEmailConfigsOk returns a tuple with the EmailConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetEmailConfigsOk() (*ReportsEmailConfigs, bool) {
	if o == nil || o.EmailConfigs == nil {
		return nil, false
	}
	return o.EmailConfigs, true
}

// HasEmailConfigs returns a boolean if a field has been set.
func (o *ReportsItem) HasEmailConfigs() bool {
	return o != nil && o.EmailConfigs != nil
}

// SetEmailConfigs gets a reference to the given ReportsEmailConfigs and assigns it to the EmailConfigs field.
func (o *ReportsItem) SetEmailConfigs(v ReportsEmailConfigs) {
	o.EmailConfigs = &v
}

// GetLastGenerationTime returns the LastGenerationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsItem) GetLastGenerationTime() string {
	if o == nil || o.LastGenerationTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastGenerationTime.Get()
}

// GetLastGenerationTimeOk returns a tuple with the LastGenerationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetLastGenerationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastGenerationTime.Get(), o.LastGenerationTime.IsSet()
}

// HasLastGenerationTime returns a boolean if a field has been set.
func (o *ReportsItem) HasLastGenerationTime() bool {
	return o != nil && o.LastGenerationTime.IsSet()
}

// SetLastGenerationTime gets a reference to the given common.NullableString and assigns it to the LastGenerationTime field.
func (o *ReportsItem) SetLastGenerationTime(v string) {
	o.LastGenerationTime.Set(&v)
}

// SetLastGenerationTimeNil sets the value for LastGenerationTime to be an explicit nil.
func (o *ReportsItem) SetLastGenerationTimeNil() {
	o.LastGenerationTime.Set(nil)
}

// UnSetLastGenerationTime ensures that no value is present for LastGenerationTime, not even an explicit nil.
func (o *ReportsItem) UnsetLastGenerationTime() {
	o.LastGenerationTime.Unset()
}

// GetNextGenerationTime returns the NextGenerationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsItem) GetNextGenerationTime() string {
	if o == nil || o.NextGenerationTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextGenerationTime.Get()
}

// GetNextGenerationTimeOk returns a tuple with the NextGenerationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetNextGenerationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextGenerationTime.Get(), o.NextGenerationTime.IsSet()
}

// HasNextGenerationTime returns a boolean if a field has been set.
func (o *ReportsItem) HasNextGenerationTime() bool {
	return o != nil && o.NextGenerationTime.IsSet()
}

// SetNextGenerationTime gets a reference to the given common.NullableString and assigns it to the NextGenerationTime field.
func (o *ReportsItem) SetNextGenerationTime(v string) {
	o.NextGenerationTime.Set(&v)
}

// SetNextGenerationTimeNil sets the value for NextGenerationTime to be an explicit nil.
func (o *ReportsItem) SetNextGenerationTimeNil() {
	o.NextGenerationTime.Set(nil)
}

// UnSetNextGenerationTime ensures that no value is present for NextGenerationTime, not even an explicit nil.
func (o *ReportsItem) UnsetNextGenerationTime() {
	o.NextGenerationTime.Unset()
}

// GetLatestReportFile returns the LatestReportFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsItem) GetLatestReportFile() ReportsLatestReportFile {
	if o == nil || o.LatestReportFile.Get() == nil {
		var ret ReportsLatestReportFile
		return ret
	}
	return *o.LatestReportFile.Get()
}

// GetLatestReportFileOk returns a tuple with the LatestReportFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetLatestReportFileOk() (*ReportsLatestReportFile, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestReportFile.Get(), o.LatestReportFile.IsSet()
}

// HasLatestReportFile returns a boolean if a LatestReportFile has been set.
func (o *ReportsItem) HasLatestReportFile() bool {
	return o != nil && o.LatestReportFile.IsSet()
}

// SetLatestReportFile gets a reference to the given common.NullableString and assigns it to the LatestReportFile field.
func (o *ReportsItem) SetLatestReportFile(v ReportsLatestReportFile) {
	o.LatestReportFile.Set(&v)
}

// SetLatestReportFileNil sets the value for LatestReportFile to be an explicit nil.
func (o *ReportsItem) SetLatestReportFileNil() {
	o.LatestReportFile.Set(nil)
}

// UnSetLatestReportFile ensures that no value is present for LatestReportFile, not even an explicit nil.
func (o *ReportsItem) UnsetLatestReportFile() {
	o.LatestReportFile.Unset()
}

// GetPageSettings returns the PageSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsItem) GetPageSettings() string {
	if o == nil || o.PageSettings.Get() == nil {
		var ret string
		return ret
	}
	return *o.PageSettings.Get()
}

// GetPageSettingsOk returns a tuple with the PageSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetPageSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSettings.Get(), o.PageSettings.IsSet()
}

// HasPageSettings returns a boolean if a field has been set.
func (o *ReportsItem) HasPageSettings() bool {
	return o != nil && o.PageSettings.IsSet()
}

// SetPageSettings gets a reference to the given common.NullableString and assigns it to the PageSettings field.
func (o *ReportsItem) SetPageSettings(v string) {
	o.PageSettings.Set(&v)
}

// SetPageSettingsNil sets the value for PageSettings to be an explicit nil.
func (o *ReportsItem) SetPageSettingsNil() {
	o.PageSettings.Set(nil)
}

// UnSetPageSettings ensures that no value is present for PageSettings, not even an explicit nil.
func (o *ReportsItem) UnsetPageSettings() {
	o.PageSettings.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ReportsItem) GetParameters() ReportsParameters {
	if o == nil || o.Parameters == nil {
		var ret ReportsParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetParametersOk() (*ReportsParameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ReportsItem) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given ReportsParameters and assigns it to the Parameters field.
func (o *ReportsItem) SetParameters(v ReportsParameters) {
	o.Parameters = &v
}

// GetTags returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetTags() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsItem) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a Tags has been set.
func (o *ReportsItem) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given common.Nullable[]string and assigns it to the Tags field.
func (o *ReportsItem) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *ReportsItem) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnSetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *ReportsItem) UnsetTags() {
	o.Tags.Unset()
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ReportsItem) GetActions() Actions {
	if o == nil || o.Actions == nil {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetActionsOk() (*Actions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ReportsItem) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *ReportsItem) SetActions(v Actions) {
	o.Actions = &v
}

// GetModuleFilter returns the ModuleFilter field value if set, zero value otherwise.
func (o *ReportsItem) GetModuleFilter() string {
	if o == nil || o.ModuleFilter == nil {
		var ret string
		return ret
	}
	return *o.ModuleFilter
}

// GetModuleFilterOk returns a tuple with the ModuleFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetModuleFilterOk() (*string, bool) {
	if o == nil || o.ModuleFilter == nil {
		return nil, false
	}
	return o.ModuleFilter, true
}

// HasModuleFilter returns a boolean if a field has been set.
func (o *ReportsItem) HasModuleFilter() bool {
	return o != nil && o.ModuleFilter != nil
}

// SetModuleFilter gets a reference to the given string and assigns it to the ModuleFilter field.
func (o *ReportsItem) SetModuleFilter(v string) {
	o.ModuleFilter = &v
}

// GetRemoteInterfaceName returns the RemoteInterfaceName field value if set, zero value otherwise.
func (o *ReportsItem) GetRemoteInterfaceName() string {
	if o == nil || o.RemoteInterfaceName == nil {
		var ret string
		return ret
	}
	return *o.RemoteInterfaceName
}

// GetRemoteInterfaceNameOk returns a tuple with the RemoteInterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetRemoteInterfaceNameOk() (*string, bool) {
	if o == nil || o.RemoteInterfaceName == nil {
		return nil, false
	}
	return o.RemoteInterfaceName, true
}

// HasRemoteInterfaceName returns a boolean if a field has been set.
func (o *ReportsItem) HasRemoteInterfaceName() bool {
	return o != nil && o.RemoteInterfaceName != nil
}

// SetRemoteInterfaceName gets a reference to the given string and assigns it to the RemoteInterfaceName field.
func (o *ReportsItem) SetRemoteInterfaceName(v string) {
	o.RemoteInterfaceName = &v
}

// GetRemoteMethodName returns the RemoteMethodName field value if set, zero value otherwise.
func (o *ReportsItem) GetRemoteMethodName() string {
	if o == nil || o.RemoteMethodName == nil {
		var ret string
		return ret
	}
	return *o.RemoteMethodName
}

// GetRemoteMethodNameOk returns a tuple with the RemoteMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetRemoteMethodNameOk() (*string, bool) {
	if o == nil || o.RemoteMethodName == nil {
		return nil, false
	}
	return o.RemoteMethodName, true
}

// HasRemoteMethodName returns a boolean if a field has been set.
func (o *ReportsItem) HasRemoteMethodName() bool {
	return o != nil && o.RemoteMethodName != nil
}

// SetRemoteMethodName gets a reference to the given string and assigns it to the RemoteMethodName field.
func (o *ReportsItem) SetRemoteMethodName(v string) {
	o.RemoteMethodName = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *ReportsItem) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *ReportsItem) HasApplicationName() bool {
	return o != nil && o.ApplicationName != nil
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *ReportsItem) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetSubExecutorModuleFilter returns the SubExecutorModuleFilter field value if set, zero value otherwise.
func (o *ReportsItem) GetSubExecutorModuleFilter() string {
	if o == nil || o.SubExecutorModuleFilter == nil {
		var ret string
		return ret
	}
	return *o.SubExecutorModuleFilter
}

// GetSubExecutorModuleFilterOk returns a tuple with the SubExecutorModuleFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetSubExecutorModuleFilterOk() (*string, bool) {
	if o == nil || o.SubExecutorModuleFilter == nil {
		return nil, false
	}
	return o.SubExecutorModuleFilter, true
}

// HasSubExecutorModuleFilter returns a boolean if a field has been set.
func (o *ReportsItem) HasSubExecutorModuleFilter() bool {
	return o != nil && o.SubExecutorModuleFilter != nil
}

// SetSubExecutorModuleFilter gets a reference to the given string and assigns it to the SubExecutorModuleFilter field.
func (o *ReportsItem) SetSubExecutorModuleFilter(v string) {
	o.SubExecutorModuleFilter = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ReportsItem) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsItem) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ReportsItem) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *ReportsItem) SetVersion(v float64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ModifiedDate != nil {
		toSerialize["ModifiedDate"] = o.ModifiedDate
	}
	if o.ReportID != nil {
		toSerialize["ReportId"] = o.ReportID
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
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
	if o.Schedule.IsSet() {
		toSerialize["Schedule"] = o.Schedule
	}
	if o.LastGenerationTime.IsSet() {
		toSerialize["LastGenerationTime"] = o.LastGenerationTime.Get()
	}
	if o.NextGenerationTime.IsSet() {
		toSerialize["NextGenerationTime"] = o.NextGenerationTime.Get()
	}
	if o.LatestReportFile.IsSet() {
		toSerialize["LatestReportFile"] = o.LatestReportFile.Get()
	}
	if o.PageSettings.IsSet() {
		toSerialize["PageSettings"] = o.PageSettings.Get()
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}
	if o.Tags.IsSet() {
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
func (o *ReportsItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModifiedDate            *string                         `json:"ModifiedDate,omitempty"`
		ReportID                *string                         `jsom:"ReportId,omitempty"`
		Username                *string                         `json:"Username,omitempty"`
		Name                    *string                         `json:"Name,omitempty"`
		Description             common.NullableString           `json:"Description,omitempty"`
		IsActive                *bool                           `json:"IsActive,omitempty"`
		Author                  common.NullableString           `json:"Author,omitempty"`
		ReportLink              *string                         `json:"ReportLink,omitempty"`
		SharedUsersAndGroups    common.NullableList[string]     `json:"SharedUsersAndGroups,omitempty"`
		CreatedDate             *string                         `json:"CreatedDate,omitempty"`
		ReportQuery             []ReportsQuery                  `json:"ReportQuery,omitempty"`
		ReportData              *ReportsData                    `json:"ReportData,omitempty"`
		Schedule                NullableScheduleConfig          `json:"Schedule,omitempty"`
		EmailConfigs            *ReportsEmailConfigs            `json:"EmailConfigs,omitempty"`
		LastGenerationTime      common.NullableString           `json:"LastGenerationTime,omitempty"`
		NextGenerationTime      common.NullableString           `json:"NextGenerationTime,omitempty"`
		LatestReportFile        NullableReportsLatestReportFile `json:"LatestReportFile,omitempty"`
		PageSettings            common.NullableString           `json:"PageSettings,omitempty"`
		Parameters              *ReportsParameters              `json:"Parameters,omitempty"`
		Tags                    common.NullableList[string]     `json:"Tags,omitempty"`
		Actions                 *Actions                        `json:"Actions,omitempty"`
		ModuleFilter            *string                         `json:"ModuleFilter,omitempty"`
		RemoteInterfaceName     *string                         `json:"RemoteInterfaceName,omitempty"`
		RemoteMethodName        *string                         `json:"RemoteMethodName,omitempty"`
		ApplicationName         *string                         `json:"ApplicationName,omitempty"`
		SubExecutorModuleFilter *string                         `json:"SubExecutorModuleFilter,omitempty"`
		Version                 *float64                        `json:"Version,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ModifiedDate", "ReportId", "Username", "Name", "Description", "IsActive", "Author", "ReportLink", "SharedUsersAndGroups", "CreatedDate", "ReportQuery", "ReportData", "Schedule", "LastGenerationTime", "NextGenerationTime", "LatestReportFile", "PageSettings", "Parameters", "Tags", "Actions", "ModuleFilter", "RemoteInterfaceName", "RemoteMethodName", "ApplicationName", "SubExecutorModuleFilter", "Version"})
	} else {
		return err
	}

	o.ModifiedDate = all.ModifiedDate
	o.ReportID = all.ReportID
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
	o.EmailConfigs = all.EmailConfigs
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
