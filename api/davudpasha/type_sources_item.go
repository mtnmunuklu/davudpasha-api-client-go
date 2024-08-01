package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourcesItem represents an item in the sources.
type SourcesItem struct {
	// ID of the source item.
	ID *string `json:"Id,omitempty"`
	// Indicates if the source item is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// Name of the source item.
	Name *string `json:"Name,omitempty"`
	// Group of the source item.
	Group common.NullableString `json:"Group,omitempty"`
	// Author of the source item.
	Author common.NullableString `json:"Author,omitempty"`
	// Code defining the log source.
	LogSourceDefCode *string `json:"LogSourceDefCode,omitempty"`
	// Type of the log reader.
	LogReaderType *string `json:"LogReaderType,omitempty"`
	// Tags associated with the source item.
	Tags []string `json:"Tags,omitempty"`
	// Timeout for alerts.
	AlertTimeout *int64 `json:"AlertTimeout,omitempty"`
	// Data related to log reader.
	LogReaderData *json.RawMessage `json:"LogReaderData,omitempty"`
	// Log operations associated with the source item.
	LogOperations []SourcesLogOperation `json:"LogOperations,omitempty"`
	// Configuration for discarded logs.
	DiscardedLogsConfig *string `json:"DiscardedLogsConfig,omitempty"`
	// Value of the source item.
	Value *string `json:"value,omitempty"`
	// Label of the source item.
	Label *string `json:"label,omitempty"`
	// Indicates if the source item is deleted.
	IsDeleted *bool `json:"IsDeleted,omitempty"`
	// Indicates if the source item is an agent source.
	IsAgentSource *bool `json:"IsAgentSource,omitempty"`
	// IDs of associated agents.
	AgentIDs common.NullableList[string] `json:"AgentIds,omitempty"`
	// Name of the index group.
	IndexGroupName common.NullableString `json:"IndexGroupName,omitempty"`
	// Name of the associated dashboard.
	DashboardName *string `json:"dashboardName,omitempty"`
	// ID of the associated dashboard.
	DashboardID common.NullableString `json:"dashboardId,omitempty"`
	// Asset tags associated with the source item.
	AssetTags common.NullableList[string] `json:"AssetTags,omitempty"`
	// Time to remove logs.
	LogRemoveTime common.NullableString `json:"LogRemoveTime,omitempty"`
	// Format to remove logs.
	LogRemoveFormat common.NullableString `json:"LogRemoveFormat,omitempty"`
	// ID of the associated agent.
	AgentID common.NullableString `json:"AgentId,omitempty"`
	// Indicates if raw logs should be written.
	WriteRawLogs *bool `json:"WriteRawLogs,omitempty"`
	// Indicates if a secondary writer should be used.
	UseSecondaryWriter *bool `json:"UseSecondaryWriter,omitempty"`
	// Parallel options for the source item.
	ParallelOptions *ParallelOptions `json:"ParallelOptions,omitempty"`
	// Block count for the source item.
	BlockCount *int64 `json:"BlockCount,omitempty"`
	// Schedule configuration for the source item.
	ScheduleConfig NullableScheduleConfig `json:"ScheduleConfig,omitempty"`
	// Indicates if raw logs should be stored.
	StoreRawLogs *bool `json:"StoreRawLogs,omitempty"`
	// Name of the LGS for storing raw logs.
	StoreRawLogsLgs *string `json:"StoreRawLogsLgs,omitempty"`
	// Indicates if the source item is editable.
	IsEditable *bool `json:"IsEditable,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourcesItem creates a new SourcesItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcesItem() *SourcesItem {
	this := SourcesItem{}
	return &this
}

// NewSourcesItemWithDefaults creates a new SourcesItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcesItemWithDefaults() *SourcesItem {
	this := SourcesItem{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *SourcesItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *SourcesItem) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *SourcesItem) SetID(v string) {
	o.ID = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SourcesItem) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SourcesItem) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SourcesItem) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourcesItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourcesItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourcesItem) SetName(v string) {
	o.Name = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a Group has been set.
func (o *SourcesItem) HasGroup() bool {
	return o != nil && o.Group.IsSet()
}

// SetGroup gets a reference to the given common.NullableString and assigns it to the Group field.
func (o *SourcesItem) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil.
func (o *SourcesItem) SetGroupNil() {
	o.Group.Set(nil)
}

// UnSetGroup ensures that no value is present for Group, not even an explicit nil.
func (o *SourcesItem) UnSetGroup() {
	o.Group.UnSet()
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetAuthor() string {
	if o == nil || o.Author.Get() == nil {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a Author has been set.
func (o *SourcesItem) HasAuthor() bool {
	return o != nil && o.Author.IsSet()
}

// SetAuthor gets a reference to the given common.NullableString and assigns it to the Author field.
func (o *SourcesItem) SetAuthor(v string) {
	o.Author.Set(&v)
}

// SetAuthorNil sets the value for Author to be an explicit nil.
func (o *SourcesItem) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnSetAuthor ensures that no value is present for Author, not even an explicit nil.
func (o *SourcesItem) UnSetAuthor() {
	o.Author.UnSet()
}

// GetLogSourceDefCode returns the LogSourceDefCode field value if set, zero value otherwise.
func (o *SourcesItem) GetLogSourceDefCode() string {
	if o == nil || o.LogSourceDefCode == nil {
		var ret string
		return ret
	}
	return *o.LogSourceDefCode
}

// GetLogSourceDefCodeOk returns a tuple with the LogSourceDefCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetLogSourceDefCodeOk() (*string, bool) {
	if o == nil || o.LogSourceDefCode == nil {
		return nil, false
	}
	return o.LogSourceDefCode, true
}

// HasLogSourceDefCode returns a boolean if a field has been set.
func (o *SourcesItem) HasLogSourceDefCode() bool {
	return o != nil && o.LogSourceDefCode != nil
}

// SetLogSourceDefCode gets a reference to the given string and assigns it to the LogSourceDefCode field.
func (o *SourcesItem) SetLogSourceDefCode(v string) {
	o.LogSourceDefCode = &v
}

// GetLogReaderType returns the LogReaderType field value if set, zero value otherwise.
func (o *SourcesItem) GetLogReaderType() string {
	if o == nil || o.LogReaderType == nil {
		var ret string
		return ret
	}
	return *o.LogReaderType
}

// GetLogReaderTypeOk returns a tuple with the LogReaderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetLogReaderTypeOk() (*string, bool) {
	if o == nil || o.LogReaderType == nil {
		return nil, false
	}
	return o.LogReaderType, true
}

// HasLogReaderType returns a boolean if a field has been set.
func (o *SourcesItem) HasLogReaderType() bool {
	return o != nil && o.LogReaderType != nil
}

// SetLogReaderType gets a reference to the given string and assigns it to the LogReaderType field.
func (o *SourcesItem) SetLogReaderType(v string) {
	o.LogReaderType = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SourcesItem) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SourcesItem) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SourcesItem) SetTags(v []string) {
	o.Tags = v
}

// GetAlertTimeout returns the AlertTimeout field value if set, zero value otherwise.
func (o *SourcesItem) GetAlertTimeout() int64 {
	if o == nil || o.AlertTimeout == nil {
		var ret int64
		return ret
	}
	return *o.AlertTimeout
}

// GetAlertTimeoutOk returns a tuple with the AlertTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetAlertTimeoutOk() (*int64, bool) {
	if o == nil || o.AlertTimeout == nil {
		return nil, false
	}
	return o.AlertTimeout, true
}

// HasAlertTimeout returns a boolean if a field has been set.
func (o *SourcesItem) HasAlertTimeout() bool {
	return o != nil && o.AlertTimeout != nil
}

// SetAlertTimeout gets a reference to the given int64 and assigns it to the AlertTimeout field.
func (o *SourcesItem) SetAlertTimeout(v int64) {
	o.AlertTimeout = &v
}

// GetLogReaderData returns the LogReaderData field value if set, zero value otherwise.
func (o *SourcesItem) GetLogReaderData() json.RawMessage {
	if o == nil || o.LogReaderData == nil {
		var ret json.RawMessage
		return ret
	}
	return *o.LogReaderData
}

// GetLogReaderDataOk returns a tuple with the LogReaderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetLogReaderDataOk() (*json.RawMessage, bool) {
	if o == nil || o.LogReaderData == nil {
		return nil, false
	}
	return o.LogReaderData, true
}

// HasLogReaderData returns a boolean if a field has been set.
func (o *SourcesItem) HasLogReaderData() bool {
	return o != nil && o.LogReaderData != nil
}

// SetLogReaderData gets a reference to the given json.RawMessage and assigns it to the LogReaderData field.
func (o *SourcesItem) SetLogReaderData(v json.RawMessage) {
	o.LogReaderData = &v
}

// GetLogOperations returns the LogOperations field value if set, zero value otherwise.
func (o *SourcesItem) GetLogOperations() []SourcesLogOperation {
	if o == nil || o.LogOperations == nil {
		var ret []SourcesLogOperation
		return ret
	}
	return o.LogOperations
}

// GetLogOperationsOk returns a tuple with the LogOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetLogOperationsOk() (*[]SourcesLogOperation, bool) {
	if o == nil || o.LogOperations == nil {
		return nil, false
	}
	return &o.LogOperations, true
}

// HasLogOperations returns a boolean if a field has been set.
func (o *SourcesItem) HasLogOperations() bool {
	return o != nil && o.LogOperations != nil
}

// SetLogOperations gets a reference to the given []SourcesLogOperation and assigns it to the LogOperations field.
func (o *SourcesItem) SetLogOperations(v []SourcesLogOperation) {
	o.LogOperations = v
}

// GetDiscardedLogsConfig returns the DiscardedLogsConfig field value if set, zero value otherwise.
func (o *SourcesItem) GetDiscardedLogsConfig() string {
	if o == nil || o.DiscardedLogsConfig == nil {
		var ret string
		return ret
	}
	return *o.DiscardedLogsConfig
}

// GetDiscardedLogsConfigOk returns a tuple with the DiscardedLogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetDiscardedLogsConfigOk() (*string, bool) {
	if o == nil || o.DiscardedLogsConfig == nil {
		return nil, false
	}
	return o.DiscardedLogsConfig, true
}

// HasDiscardedLogsConfig returns a boolean if a field has been set.
func (o *SourcesItem) HasDiscardedLogsConfig() bool {
	return o != nil && o.DiscardedLogsConfig != nil
}

// SetDiscardedLogsConfig gets a reference to the given string and assigns it to the DiscardedLogsConfig field.
func (o *SourcesItem) SetDiscardedLogsConfig(v string) {
	o.DiscardedLogsConfig = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SourcesItem) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SourcesItem) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SourcesItem) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SourcesItem) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SourcesItem) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SourcesItem) SetLabel(v string) {
	o.Label = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *SourcesItem) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *SourcesItem) HasIsDeleted() bool {
	return o != nil && o.IsDeleted != nil
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *SourcesItem) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetIsAgentSource returns the IsAgentSource field value if set, zero value otherwise.
func (o *SourcesItem) GetIsAgentSource() bool {
	if o == nil || o.IsAgentSource == nil {
		var ret bool
		return ret
	}
	return *o.IsAgentSource
}

// GetIsAgentSourceOk returns a tuple with the IsAgentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetIsAgentSourceOk() (*bool, bool) {
	if o == nil || o.IsAgentSource == nil {
		return nil, false
	}
	return o.IsAgentSource, true
}

// HasIsAgentSource returns a boolean if a field has been set.
func (o *SourcesItem) HasIsAgentSource() bool {
	return o != nil && o.IsAgentSource != nil
}

// SetIsAgentSource gets a reference to the given bool and assigns it to the IsAgentSource field.
func (o *SourcesItem) SetIsAgentSource(v bool) {
	o.IsAgentSource = &v
}

// GetAgentIDs returns the AgentIDs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetAgentIDs() []string {
	if o == nil || o.AgentIDs.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AgentIDs.Get()
}

// GetAgentIDsOk returns a tuple with the AgentIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetAgentIDsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentIDs.Get(), o.AgentIDs.IsSet()
}

// HasAgentIDs returns a boolean if a AgentIds has been set.
func (o *SourcesItem) HasAgentIDs() bool {
	return o != nil && o.AgentIDs.IsSet()
}

// SetAgentIDs gets a reference to the given common.Nullable[]string and assigns it to the AgentIDs field.
func (o *SourcesItem) SetAgentIDs(v []string) {
	o.AgentIDs.Set(&v)
}

// SetAgentIDsNil sets the value for AgentIDs to be an explicit nil.
func (o *SourcesItem) SetAgentIDsNil() {
	o.AgentIDs.Set(nil)
}

// UnSetAgentIds ensures that no value is present for AgentIds, not even an explicit nil.
func (o *SourcesItem) UnSetAgentIds() {
	o.AgentIDs.UnSet()
}

// GetIndexGroupName returns the IndexGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetIndexGroupName() string {
	if o == nil || o.IndexGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.IndexGroupName.Get()
}

// GetIndexGroupNameOk returns a tuple with the IndexGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetIndexGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexGroupName.Get(), o.IndexGroupName.IsSet()
}

// HasIndexGroupName returns a boolean if a IndexGroupName has been set.
func (o *SourcesItem) HasIndexGroupName() bool {
	return o != nil && o.IndexGroupName.IsSet()
}

// SetIndexGroupName gets a reference to the given common.NullableString and assigns it to the IndexGroupName field.
func (o *SourcesItem) SetIndexGroupName(v string) {
	o.IndexGroupName.Set(&v)
}

// SetIndexGroupNameNil sets the value for IndexGroupName to be an explicit nil.
func (o *SourcesItem) SetIndexGroupNameNil() {
	o.IndexGroupName.Set(nil)
}

// UnSetIndexGroupName ensures that no value is present for IndexGroupName, not even an explicit nil.
func (o *SourcesItem) UnSetIndexGroupName() {
	o.IndexGroupName.UnSet()
}

// GetDashboardName returns the DashboardName field value if set, zero value otherwise.
func (o *SourcesItem) GetDashboardName() string {
	if o == nil || o.DashboardName == nil {
		var ret string
		return ret
	}
	return *o.DashboardName
}

// GetDashboardNameOk returns a tuple with the DashboardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetDashboardNameOk() (*string, bool) {
	if o == nil || o.DashboardName == nil {
		return nil, false
	}
	return o.DashboardName, true
}

// HasDashboardName returns a boolean if a field has been set.
func (o *SourcesItem) HasDashboardName() bool {
	return o != nil && o.DashboardName != nil
}

// SetDashboardName gets a reference to the given string and assigns it to the DashboardName field.
func (o *SourcesItem) SetDashboardName(v string) {
	o.DashboardName = &v
}

// GetDashboardID returns the DashboardID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetDashboardID() string {
	if o == nil || o.DashboardID.Get() == nil {
		var ret string
		return ret
	}
	return *o.DashboardID.Get()
}

// GetDashboardIDOk returns a tuple with the DashboardID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetDashboardIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DashboardID.Get(), o.DashboardID.IsSet()
}

// HasDashboardID returns a boolean if a DashboardID has been set.
func (o *SourcesItem) HasDashboardID() bool {
	return o != nil && o.DashboardID.IsSet()
}

// SetDashboardID gets a reference to the given common.NullableString and assigns it to the DashboardID field.
func (o *SourcesItem) SetDashboardID(v string) {
	o.DashboardID.Set(&v)
}

// SetDashboardIDNil sets the value for DashboardID to be an explicit nil.
func (o *SourcesItem) SetDashboardIDNil() {
	o.DashboardID.Set(nil)
}

// UnSetDashboardID ensures that no value is present for DashboardID, not even an explicit nil.
func (o *SourcesItem) UnSetDashboardID() {
	o.DashboardID.UnSet()
}

// GetAssetTags returns the AssetTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetAssetTags() []string {
	if o == nil || o.AssetTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AssetTags.Get()
}

// GetAssetTagsOk returns a tuple with the AssetTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetAssetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTags.Get(), o.AssetTags.IsSet()
}

// HasAssetTags returns a boolean if a AssetTags has been set.
func (o *SourcesItem) HasAssetTags() bool {
	return o != nil && o.AssetTags.IsSet()
}

// SetAssetTags gets a reference to the given common.Nullable[]string and assigns it to the AssetTags field.
func (o *SourcesItem) SetAssetTags(v []string) {
	o.AssetTags.Set(&v)
}

// SetAssetTagsNil sets the value for AssetTags to be an explicit nil.
func (o *SourcesItem) SetAssetTagsNil() {
	o.AssetTags.Set(nil)
}

// UnSetAssetTags ensures that no value is present for AssetTags, not even an explicit nil.
func (o *SourcesItem) UnSetAssetTags() {
	o.AssetTags.UnSet()
}

// GetLogRemoveTime returns the LogRemoveTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetLogRemoveTime() string {
	if o == nil || o.LogRemoveTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogRemoveTime.Get()
}

// GetLogRemoveTimeOk returns a tuple with the LogRemoveTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetLogRemoveTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogRemoveTime.Get(), o.LogRemoveTime.IsSet()
}

// HasLogRemoveTime returns a boolean if a LogRemoveTime has been set.
func (o *SourcesItem) HasLogRemoveTime() bool {
	return o != nil && o.LogRemoveTime.IsSet()
}

// SetLogRemoveTime gets a reference to the given common.NullableString and assigns it to the LogRemoveTime field.
func (o *SourcesItem) SetLogRemoveTime(v string) {
	o.LogRemoveTime.Set(&v)
}

// SetLogRemoveTimeNil sets the value for LogRemoveTime to be an explicit nil.
func (o *SourcesItem) SetLogRemoveTimeNil() {
	o.LogRemoveTime.Set(nil)
}

// UnSetLogRemoveTime ensures that no value is present for LogRemoveTime, not even an explicit nil.
func (o *SourcesItem) UnSetLogRemoveTime() {
	o.LogRemoveTime.UnSet()
}

// GetLogRemoveFormat returns the LogRemoveFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetLogRemoveFormat() string {
	if o == nil || o.LogRemoveFormat.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogRemoveFormat.Get()
}

// GetLogRemoveFormatOk returns a tuple with the LogRemoveFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetLogRemoveFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogRemoveFormat.Get(), o.LogRemoveFormat.IsSet()
}

// HasLogRemoveFormat returns a boolean if a LogRemoveFormat has been set.
func (o *SourcesItem) HasLogRemoveFormat() bool {
	return o != nil && o.LogRemoveFormat.IsSet()
}

// SetLogRemoveFormat gets a reference to the given common.NullableString and assigns it to the LogRemoveFormat field.
func (o *SourcesItem) SetLogRemoveFormat(v string) {
	o.LogRemoveFormat.Set(&v)
}

// SetLogRemoveFormatNil sets the value for LogRemoveFormat to be an explicit nil.
func (o *SourcesItem) SetLogRemoveFormatNil() {
	o.LogRemoveFormat.Set(nil)
}

// UnSetLogRemoveFormat ensures that no value is present for LogRemoveFormat, not even an explicit nil.
func (o *SourcesItem) UnSetLogRemoveFormat() {
	o.LogRemoveFormat.UnSet()
}

// GetAgentID returns the AgentID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetAgentID() string {
	if o == nil || o.AgentID.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgentID.Get()
}

// GetAgentIDOk returns a tuple with the AgentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetAgentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentID.Get(), o.AgentID.IsSet()
}

// HasAgentID returns a boolean if a AgentID has been set.
func (o *SourcesItem) HasAgentID() bool {
	return o != nil && o.AgentID.IsSet()
}

// SetAgentID gets a reference to the given common.NullableString and assigns it to the AgentID field.
func (o *SourcesItem) SetAgentID(v string) {
	o.AgentID.Set(&v)
}

// SetAgentIDNil sets the value for AgentID to be an explicit nil.
func (o *SourcesItem) SetAgentIDNil() {
	o.AgentID.Set(nil)
}

// UnSetAgentID ensures that no value is present for AgentID, not even an explicit nil.
func (o *SourcesItem) UnSetAgentID() {
	o.AgentID.UnSet()
}

// GetWriteRawLogs returns the WriteRawLogs field value if set, zero value otherwise.
func (o *SourcesItem) GetWriteRawLogs() bool {
	if o == nil || o.WriteRawLogs == nil {
		var ret bool
		return ret
	}
	return *o.WriteRawLogs
}

// GetWriteRawLogsOk returns a tuple with the WriteRawLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetWriteRawLogsOk() (*bool, bool) {
	if o == nil || o.WriteRawLogs == nil {
		return nil, false
	}
	return o.WriteRawLogs, true
}

// HasWriteRawLogs returns a boolean if a field has been set.
func (o *SourcesItem) HasWriteRawLogs() bool {
	return o != nil && o.WriteRawLogs != nil
}

// SetWriteRawLogs gets a reference to the given bool and assigns it to the WriteRawLogs field.
func (o *SourcesItem) SetWriteRawLogs(v bool) {
	o.WriteRawLogs = &v
}

// GetUseSecondaryWriter returns the UseSecondaryWriter field value if set, zero value otherwise.
func (o *SourcesItem) GetUseSecondaryWriter() bool {
	if o == nil || o.UseSecondaryWriter == nil {
		var ret bool
		return ret
	}
	return *o.UseSecondaryWriter
}

// GetUseSecondaryWriterOk returns a tuple with the UseSecondaryWriter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetUseSecondaryWriterOk() (*bool, bool) {
	if o == nil || o.UseSecondaryWriter == nil {
		return nil, false
	}
	return o.UseSecondaryWriter, true
}

// HasUseSecondaryWriter returns a boolean if a field has been set.
func (o *SourcesItem) HasUseSecondaryWriter() bool {
	return o != nil && o.UseSecondaryWriter != nil
}

// SetUseSecondaryWriter gets a reference to the given bool and assigns it to the UseSecondaryWriter field.
func (o *SourcesItem) SetUseSecondaryWriter(v bool) {
	o.UseSecondaryWriter = &v
}

// GetParallelOptions returns the ParallelOptions field value if set, zero value otherwise.
func (o *SourcesItem) GetParallelOptions() ParallelOptions {
	if o == nil || o.ParallelOptions == nil {
		var ret ParallelOptions
		return ret
	}
	return *o.ParallelOptions
}

// GetParallelOptionsOk returns a tuple with the ParallelOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetParallelOptionsOk() (*ParallelOptions, bool) {
	if o == nil || o.ParallelOptions == nil {
		return nil, false
	}
	return o.ParallelOptions, true
}

// HasParallelOptions returns a boolean if a field has been set.
func (o *SourcesItem) HasParallelOptions() bool {
	return o != nil && o.ParallelOptions != nil
}

// SetParallelOptions gets a reference to the given ParallelOptions and assigns it to the ParallelOptions field.
func (o *SourcesItem) SetParallelOptions(v ParallelOptions) {
	o.ParallelOptions = &v
}

// GetBlockCount returns the BlockCount field value if set, zero value otherwise.
func (o *SourcesItem) GetBlockCount() int64 {
	if o == nil || o.BlockCount == nil {
		var ret int64
		return ret
	}
	return *o.BlockCount
}

// GetBlockCountOk returns a tuple with the BlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetBlockCountOk() (*int64, bool) {
	if o == nil || o.BlockCount == nil {
		return nil, false
	}
	return o.BlockCount, true
}

// HasBlockCount returns a boolean if a field has been set.
func (o *SourcesItem) HasBlockCount() bool {
	return o != nil && o.BlockCount != nil
}

// SetBlockCount gets a reference to the given int64 and assigns it to the BlockCount field.
func (o *SourcesItem) SetBlockCount(v int64) {
	o.BlockCount = &v
}

// GetScheduleConfig returns the ScheduleConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesItem) GetScheduleConfig() ScheduleConfig {
	if o == nil || o.ScheduleConfig.Get() == nil {
		var ret ScheduleConfig
		return ret
	}
	return *o.ScheduleConfig.Get()
}

// GetScheduleConfigOk returns a tuple with the ScheduleConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesItem) GetScheduleConfigOk() (*ScheduleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleConfig.Get(), o.ScheduleConfig.IsSet()
}

// HasScheduleConfig returns a boolean if a ScheduleConfig has been set.
func (o *SourcesItem) HasScheduleConfig() bool {
	return o != nil && o.ScheduleConfig.IsSet()
}

// SetScheduleConfig gets a reference to the given common.NullableString and assigns it to the ScheduleConfig field.
func (o *SourcesItem) SetScheduleConfig(v ScheduleConfig) {
	o.ScheduleConfig.Set(&v)
}

// SetScheduleConfigNil sets the value for ScheduleConfig to be an explicit nil.
func (o *SourcesItem) SetScheduleConfigNil() {
	o.ScheduleConfig.Set(nil)
}

// UnSetScheduleConfig ensures that no value is present for ScheduleConfig, not even an explicit nil.
func (o *SourcesItem) UnSetScheduleConfig() {
	o.ScheduleConfig.UnSet()
}

// GetStoreRawLogs returns the StoreRawLogs field value if set, zero value otherwise.
func (o *SourcesItem) GetStoreRawLogs() bool {
	if o == nil || o.StoreRawLogs == nil {
		var ret bool
		return ret
	}
	return *o.StoreRawLogs
}

// GetStoreRawLogsOk returns a tuple with the StoreRawLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetStoreRawLogsOk() (*bool, bool) {
	if o == nil || o.StoreRawLogs == nil {
		return nil, false
	}
	return o.StoreRawLogs, true
}

// HasStoreRawLogs returns a boolean if a field has been set.
func (o *SourcesItem) HasStoreRawLogs() bool {
	return o != nil && o.StoreRawLogs != nil
}

// SetStoreRawLogs gets a reference to the given bool and assigns it to the StoreRawLogs field.
func (o *SourcesItem) SetStoreRawLogs(v bool) {
	o.StoreRawLogs = &v
}

// GetStoreRawLogsLgs returns the StoreRawLogsLgs field value if set, zero value otherwise.
func (o *SourcesItem) GetStoreRawLogsLgs() string {
	if o == nil || o.StoreRawLogsLgs == nil {
		var ret string
		return ret
	}
	return *o.StoreRawLogsLgs
}

// GetStoreRawLogsLgsOk returns a tuple with the StoreRawLogsLgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetStoreRawLogsLgsOk() (*string, bool) {
	if o == nil || o.StoreRawLogsLgs == nil {
		return nil, false
	}
	return o.StoreRawLogsLgs, true
}

// HasStoreRawLogsLgs returns a boolean if a field has been set.
func (o *SourcesItem) HasStoreRawLogsLgs() bool {
	return o != nil && o.StoreRawLogsLgs != nil
}

// SetStoreRawLogsLgs gets a reference to the given string and assigns it to the StoreRawLogsLgs field.
func (o *SourcesItem) SetStoreRawLogsLgs(v string) {
	o.StoreRawLogsLgs = &v
}

// GetIsEditable returns the IsEditable field value if set, zero value otherwise.
func (o *SourcesItem) GetIsEditable() bool {
	if o == nil || o.IsEditable == nil {
		var ret bool
		return ret
	}
	return *o.IsEditable
}

// GetIsEditableOk returns a tuple with the IsEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItem) GetIsEditableOk() (*bool, bool) {
	if o == nil || o.IsEditable == nil {
		return nil, false
	}
	return o.IsEditable, true
}

// HasIsEditable returns a boolean if a field has been set.
func (o *SourcesItem) HasIsEditable() bool {
	return o != nil && o.IsEditable != nil
}

// SetIsEditable gets a reference to the given bool and assigns it to the IsEditable field.
func (o *SourcesItem) SetIsEditable(v bool) {
	o.IsEditable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcesItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["Id"] = o.ID
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Group.IsSet() {
		toSerialize["Group"] = o.Group.Get()
	}
	if o.Author.IsSet() {
		toSerialize["Author"] = o.Author.Get()
	}
	if o.LogSourceDefCode != nil {
		toSerialize["LogSourceDefCode"] = o.LogSourceDefCode
	}
	if o.LogReaderType != nil {
		toSerialize["LogReaderType"] = o.LogReaderType
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.AlertTimeout != nil {
		toSerialize["AlertTimeout"] = o.AlertTimeout
	}
	if o.LogReaderData != nil {
		toSerialize["LogReaderData"] = o.LogReaderData
	}
	if o.LogOperations != nil {
		toSerialize["LogOperations"] = o.LogOperations
	}
	if o.DiscardedLogsConfig != nil {
		toSerialize["DiscardedLogsConfig"] = o.DiscardedLogsConfig
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.IsDeleted != nil {
		toSerialize["IsDeleted"] = o.IsDeleted
	}
	if o.IsAgentSource != nil {
		toSerialize["IsAgentSource"] = o.IsAgentSource
	}
	if o.AgentIDs.IsSet() {
		toSerialize["AgentIds"] = o.AgentIDs.Get()
	}
	if o.IndexGroupName.IsSet() {
		toSerialize["IndexGroupName"] = o.IndexGroupName.Get()
	}
	if o.DashboardName != nil {
		toSerialize["DashboardName"] = o.DashboardName
	}
	if o.DashboardID.IsSet() {
		toSerialize["DashboardID"] = o.DashboardID.Get()
	}
	if o.AssetTags.IsSet() {
		toSerialize["AssetTags"] = o.AssetTags.Get()
	}
	if o.LogRemoveTime.IsSet() {
		toSerialize["LogRemoveTime"] = o.LogRemoveTime.Get()
	}
	if o.LogRemoveFormat.IsSet() {
		toSerialize["LogRemoveFormat"] = o.LogRemoveFormat.Get()
	}
	if o.AgentID.IsSet() {
		toSerialize["AgentId"] = o.AgentID.Get()
	}
	if o.WriteRawLogs != nil {
		toSerialize["WriteRawLogs"] = o.WriteRawLogs
	}
	if o.UseSecondaryWriter != nil {
		toSerialize["UseSecondaryWriter"] = o.UseSecondaryWriter
	}
	if o.ParallelOptions != nil {
		toSerialize["ParallelOptions"] = o.ParallelOptions
	}
	if o.BlockCount != nil {
		toSerialize["BlockCount"] = o.BlockCount
	}
	if o.BlockCount != nil {
		toSerialize["BlockCount"] = o.BlockCount
	}
	if o.ScheduleConfig.IsSet() {
		toSerialize["ScheduleConfig"] = o.ScheduleConfig.Get()
	}
	if o.StoreRawLogs != nil {
		toSerialize["StoreRawLogs"] = o.StoreRawLogs
	}
	if o.StoreRawLogsLgs != nil {
		toSerialize["StoreRawLogsLgs"] = o.StoreRawLogsLgs
	}
	if o.IsEditable != nil {
		toSerialize["IsEditable"] = o.IsEditable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourcesItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID                  *string                     `json:"Id,omitempty"`
		Enabled             *bool                       `json:"Enabled,omitempty"`
		Name                *string                     `json:"Name,omitempty"`
		Group               common.NullableString       `json:"Group,omitempty"`
		Author              common.NullableString       `json:"Author,omitempty"`
		LogSourceDefCode    *string                     `json:"LogSourceDefCode,omitempty"`
		LogReaderType       *string                     `json:"LogReaderType,omitempty"`
		Tags                []string                    `json:"Tags,omitempty"`
		AlertTimeout        *int64                      `json:"AlertTimeout,omitempty"`
		LogReaderData       *json.RawMessage            `json:"LogReaderData,omitempty"`
		LogOperations       []SourcesLogOperation       `json:"LogOperations,omitempty"`
		DiscardedLogsConfig *string                     `json:"DiscardedLogsConfig,omitempty"`
		Value               *string                     `json:"value,omitempty"`
		Label               *string                     `json:"label,omitempty"`
		IsDeleted           *bool                       `json:"IsDeleted,omitempty"`
		IsAgentSource       *bool                       `json:"IsAgentSource,omitempty"`
		AgentIDs            common.NullableList[string] `json:"AgentIds,omitempty"`
		IndexGroupName      common.NullableString       `json:"IndexGroupName,omitempty"`
		DashboardName       *string                     `json:"dashboardName,omitempty"`
		DashboardID         common.NullableString       `json:"dashboardId,omitempty"`
		AssetTags           common.NullableList[string] `json:"AssetTags,omitempty"`
		LogRemoveTime       common.NullableString       `json:"LogRemoveTime,omitempty"`
		LogRemoveFormat     common.NullableString       `json:"LogRemoveFormat,omitempty"`
		AgentID             common.NullableString       `json:"AgentId,omitempty"`
		WriteRawLogs        *bool                       `json:"WriteRawLogs,omitempty"`
		UseSecondaryWriter  *bool                       `json:"UseSecondaryWriter,omitempty"`
		ParallelOptions     *ParallelOptions            `json:"ParallelOptions,omitempty"`
		BlockCount          *int64                      `json:"BlockCount,omitempty"`
		ScheduleConfig      NullableScheduleConfig      `json:"ScheduleConfig,omitempty"`
		StoreRawLogs        *bool                       `json:"StoreRawLogs,omitempty"`
		StoreRawLogsLgs     *string                     `json:"StoreRawLogsLgs,omitempty"`
		IsEditable          *bool                       `json:"IsEditable,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.ID == nil {
		return fmt.Errorf("requiered field Id is missing")
	}
	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Id", "Enabled", "Name", "Group", "Author", "LogSourceDefCode", "LogReaderType", "Tags", "AlertTimeout", "LogReaderData", "LogOperations", "DiscardedLogsConfig", "Value", "Label", "IsDeleted", "IsAgentSource", "AgentIds", "IndexGroupName", "DashboardName", "DashboardID", "AssetTags", "LogRemoveTime", "LogRemoveFormat", "AgentId", "WriteRawLogs", "UseSecondaryWriter", "ParallelOptions", "BlockCount", "ScheduleConfig", "StoreRawLogs", "StoreRawLogsLgs", "IsEditable"})
	} else {
		return err
	}

	o.ID = all.ID
	o.Enabled = all.Enabled
	o.Name = all.Name
	o.Group = all.Group
	o.Author = all.Author
	o.LogSourceDefCode = all.LogSourceDefCode
	o.LogReaderType = all.LogReaderType
	o.Tags = all.Tags
	o.AlertTimeout = all.AlertTimeout
	o.LogReaderData = all.LogReaderData
	o.LogOperations = all.LogOperations
	o.DiscardedLogsConfig = all.DiscardedLogsConfig
	o.Value = all.Value
	o.Label = all.Label
	o.IsDeleted = all.IsDeleted
	o.IsAgentSource = all.IsAgentSource
	o.AgentIDs = all.AgentIDs
	o.IndexGroupName = all.IndexGroupName
	o.DashboardName = all.DashboardName
	o.DashboardID = all.DashboardID
	o.AssetTags = all.AssetTags
	o.LogRemoveTime = all.LogRemoveTime
	o.AgentID = all.AgentID
	o.WriteRawLogs = all.WriteRawLogs
	o.UseSecondaryWriter = all.UseSecondaryWriter
	hasInvalidField := false
	if all.ParallelOptions != nil && all.ParallelOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ParallelOptions = all.ParallelOptions
	o.BlockCount = all.BlockCount
	o.ScheduleConfig = all.ScheduleConfig
	o.StoreRawLogs = all.StoreRawLogs
	o.StoreRawLogsLgs = all.StoreRawLogsLgs
	o.IsEditable = all.IsEditable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
