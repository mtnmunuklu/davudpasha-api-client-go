package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// AlertDefinitionsCorrelationData represents correlation data for alerts.
type AlertDefinitionsCorrelationData struct {
	// ID of the alert correlation.
	ID *string `json:"Id,omitempty"`
	// Whether the correlation is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// Name of the correlation.
	Name *string `json:"Name,omitempty"`
	// Description of the correlation.
	Description common.NullableString `json:"Description,omitempty"`
	// Tags associated with the correlation.
	Tags common.NullableList[string] `json:"Tags,omitempty"`
	// Whether the correlation is from the market.
	FromMarket *bool `json:"FromMarket,omitempty"`
	// Whether the correlation is from modules.
	FromModules *bool `json:"FromModules,omitempty"`
	// Whether the correlation has an update.
	HasUpdate *bool `json:"HasUpdate,omitempty"`
	// ID of the module.
	ModuleID common.NullableString `json:"ModuleId,omitempty"`
	// GUID of the module.
	ModuleGUID common.NullableString `json:"ModuleGuid,omitempty"`
	// Risk level of the correlation.
	RiskLevel *int64 `json:"RiskLevel,omitempty"`
	// Maximum alert count.
	MaxAlertCount *int64 `json:"MaxAlertCount,omitempty"`
	// Time frame type for limiter.
	LimiterTimeFrameType common.NullableString `json:"LimiterTimeFrameType,omitempty"`
	// Time frame value for limiter.
	LimiterTimeFrameValue common.NullableInt64 `json:"LimiterTimeFrameValue,omitempty"`
	// Columns used for limiting.
	LimiterColumns common.NullableList[string] `json:"LimiterColumns,omitempty"`
	// Actions related to the correlation.
	Actions *Actions `json:"Actions,omitempty"`
	// Maximum send count.
	MaxSendCount *int64 `json:"MaxSendCount,omitempty"`
	// Raw data in JSON format.
	Data *AlertDefinitionsQueryData `json:"Data,omitempty"`
	// Type of correlation.
	CorrelationType AlertDefinitionsCorrelationType `json:"CorrelationType,omitempty"`
	// Message of the correlation.
	Message *string `json:"Message,omitempty"`
	// Columns grouped by the correlation.
	GroupedColumns []string `json:"GroupedColumns,omitempty"`
	// Options for columns that can be grouped.
	GroupedColumnsOptions []SelectedColumn `json:"GroupedColumnsOptions,omitempty"`
	// Version of the correlation.
	Version *float64 `json:"Version,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertDefinitionsCorrelationData creates a new AlertDefinitionsCorrelationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertDefinitionsCorrelationData() *AlertDefinitionsCorrelationData {
	this := AlertDefinitionsCorrelationData{}
	return &this
}

// NewAlertDefinitionsCorrelationDataWithDefaults creates a new AlertDefinitionsCorrelationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertDefinitionsCorrelationDataWithDefaults() *AlertDefinitionsCorrelationData {
	this := AlertDefinitionsCorrelationData{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *AlertDefinitionsCorrelationData) SetID(v string) {
	o.ID = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertDefinitionsCorrelationData) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertDefinitionsCorrelationData) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertDefinitionsCorrelationData) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsCorrelationData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *AlertDefinitionsCorrelationData) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *AlertDefinitionsCorrelationData) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnSetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *AlertDefinitionsCorrelationData) UnsetDescription() {
	o.Description.Unset()
}

// GetTags returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsCorrelationData) GetTags() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsCorrelationData) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a Tags has been set.
func (o *AlertDefinitionsCorrelationData) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given common.Nullable[]string and assigns it to the Tags field.
func (o *AlertDefinitionsCorrelationData) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *AlertDefinitionsCorrelationData) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnSetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *AlertDefinitionsCorrelationData) UnsetTags() {
	o.Tags.Unset()
}

// GetFromMarket returns the FromMarket field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetFromMarket() bool {
	if o == nil || o.FromMarket == nil {
		var ret bool
		return ret
	}
	return *o.FromMarket
}

// GetFromMarketOk returns a tuple with the FromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetFromMarketOk() (*bool, bool) {
	if o == nil || o.FromMarket == nil {
		return nil, false
	}
	return o.FromMarket, true
}

// HasFromMarket returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasFromMarket() bool {
	return o != nil && o.FromMarket != nil
}

// SetFromMarket gets a reference to the given bool and assigns it to the FromMarket field.
func (o *AlertDefinitionsCorrelationData) SetFromMarket(v bool) {
	o.FromMarket = &v
}

// GetFromModules returns the FromModules field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetFromModules() bool {
	if o == nil || o.FromModules == nil {
		var ret bool
		return ret
	}
	return *o.FromModules
}

// GetFromModulesOk returns a tuple with the FromModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetFromModulesOk() (*bool, bool) {
	if o == nil || o.FromModules == nil {
		return nil, false
	}
	return o.FromModules, true
}

// HasFromModules returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasFromModules() bool {
	return o != nil && o.FromModules != nil
}

// SetFromModules gets a reference to the given bool and assigns it to the FromModules field.
func (o *AlertDefinitionsCorrelationData) SetFromModules(v bool) {
	o.FromModules = &v
}

// GetHasUpdate returns the HasUpdate field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetHasUpdate() bool {
	if o == nil || o.HasUpdate == nil {
		var ret bool
		return ret
	}
	return *o.HasUpdate
}

// GetHasUpdateOk returns a tuple with the HasUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetHasUpdateOk() (*bool, bool) {
	if o == nil || o.HasUpdate == nil {
		return nil, false
	}
	return o.HasUpdate, true
}

// HasHasUpdate returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasHasUpdate() bool {
	return o != nil && o.HasUpdate != nil
}

// SetHasUpdate gets a reference to the given bool and assigns it to the HasUpdate field.
func (o *AlertDefinitionsCorrelationData) SetHasUpdate(v bool) {
	o.HasUpdate = &v
}

// GetModuleID returns the ModuleID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertDefinitionsCorrelationData) GetModuleID() string {
	if o == nil || o.ModuleID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleID.Get()
}

// GetModuleIDOk returns a tuple with the ModuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsCorrelationData) GetModuleIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleID.Get(), o.ModuleID.IsSet()
}

// HasModuleID returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasModuleID() bool {
	return o != nil && o.ModuleID.IsSet()
}

// SetModuleID gets a reference to the given common.NullableString and assigns it to the ModuleID field.
func (o *AlertDefinitionsCorrelationData) SetModuleID(v string) {
	o.ModuleID.Set(&v)
}

// SetModuleIDNil sets the value for ModuleID to be an explicit nil.
func (o *AlertDefinitionsCorrelationData) SetModuleIDNil() {
	o.ModuleID.Set(nil)
}

// UnSetModuleID ensures that no value is present for ModuleID, not even an explicit nil.
func (o *AlertDefinitionsCorrelationData) UnsetModuleID() {
	o.ModuleID.Unset()
}

// GetModuleGUID returns the ModuleGUID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertDefinitionsCorrelationData) GetModuleGUID() string {
	if o == nil || o.ModuleGUID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleGUID.Get()
}

// GetModuleGUIDOk returns a tuple with the ModuleGUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsCorrelationData) GetModuleGUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleGUID.Get(), o.ModuleGUID.IsSet()
}

// HasModuleGUID returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasModuleGUID() bool {
	return o != nil && o.ModuleGUID.IsSet()
}

// SetModuleGUID gets a reference to the given common.NullableString and assigns it to the ModuleGUID field.
func (o *AlertDefinitionsCorrelationData) SetModuleGUID(v string) {
	o.ModuleGUID.Set(&v)
}

// SetModuleGUIDNil sets the value for ModuleGUID to be an explicit nil.
func (o *AlertDefinitionsCorrelationData) SetModuleGUIDNil() {
	o.ModuleGUID.Set(nil)
}

// UnSetModuleGUID ensures that no value is present for ModuleGUID, not even an explicit nil.
func (o *AlertDefinitionsCorrelationData) UnsetModuleGUID() {
	o.ModuleGUID.Unset()
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetRiskLevel() int64 {
	if o == nil || o.RiskLevel == nil {
		var ret int64
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetRiskLevelOk() (*int64, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasRiskLevel() bool {
	return o != nil && o.RiskLevel != nil
}

// SetRiskLevel gets a reference to the given int64 and assigns it to the RiskLevel field.
func (o *AlertDefinitionsCorrelationData) SetRiskLevel(v int64) {
	o.RiskLevel = &v
}

// GetMaxAlertCount returns the MaxAlertCount field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetMaxAlertCount() int64 {
	if o == nil || o.MaxAlertCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxAlertCount
}

// GetMaxAlertCountOk returns a tuple with the MaxAlertCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetMaxAlertCountOk() (*int64, bool) {
	if o == nil || o.MaxAlertCount == nil {
		return nil, false
	}
	return o.MaxAlertCount, true
}

// HasMaxAlertCount returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasMaxAlertCount() bool {
	return o != nil && o.MaxAlertCount != nil
}

// SetMaxAlertCount gets a reference to the given int64 and assigns it to the MaxAlertCount field.
func (o *AlertDefinitionsCorrelationData) SetMaxAlertCount(v int64) {
	o.MaxAlertCount = &v
}

// GetLimiterTimeFrameType returns the LimiterTimeFrameType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertDefinitionsCorrelationData) GetLimiterTimeFrameType() string {
	if o == nil || o.LimiterTimeFrameType.Get() == nil {
		var ret string
		return ret
	}
	return *o.LimiterTimeFrameType.Get()
}

// GetLimiterTimeFrameTypeOk returns a tuple with the LimiterTimeFrameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsCorrelationData) GetLimiterTimeFrameTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimiterTimeFrameType.Get(), o.LimiterTimeFrameType.IsSet()
}

// HasLimiterTimeFrameType returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasLimiterTimeFrameType() bool {
	return o != nil && o.LimiterTimeFrameType.IsSet()
}

// SetLimiterTimeFrameType gets a reference to the given common.NullableString and assigns it to the LimiterTimeFrameType field.
func (o *AlertDefinitionsCorrelationData) SetLimiterTimeFrameType(v string) {
	o.LimiterTimeFrameType.Set(&v)
}

// SetLimiterTimeFrameTypeNil sets the value for LimiterTimeFrameType to be an explicit nil.
func (o *AlertDefinitionsCorrelationData) SetLimiterTimeFrameTypeNil() {
	o.LimiterTimeFrameType.Set(nil)
}

// UnSetLimiterTimeFrameType ensures that no value is present for LimiterTimeFrameType, not even an explicit nil.
func (o *AlertDefinitionsCorrelationData) UnsetLimiterTimeFrameType() {
	o.LimiterTimeFrameType.Unset()
}

// GetLimiterColumns returns a tuple with the LimiterColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsCorrelationData) GetLimiterColumns() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimiterColumns.Get(), o.LimiterColumns.IsSet()
}

// GetLimiterColumnsOk returns a tuple with the LimiterColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsCorrelationData) GetLimiterColumnsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimiterColumns.Get(), o.LimiterColumns.IsSet()
}

// HasLimiterColumns returns a boolean if a LimiterColumns has been set.
func (o *AlertDefinitionsCorrelationData) HasLimiterColumns() bool {
	return o != nil && o.LimiterColumns.IsSet()
}

// SetLimiterColumns gets a reference to the given common.Nullable[]string and assigns it to the LimiterColumns field.
func (o *AlertDefinitionsCorrelationData) SetLimiterColumns(v []string) {
	o.LimiterColumns.Set(&v)
}

// SetLimiterColumnsNil sets the value for LimiterColumns to be an explicit nil.
func (o *AlertDefinitionsCorrelationData) SetLimiterColumnsNil() {
	o.LimiterColumns.Set(nil)
}

// UnSetLimiterColumns ensures that no value is present for LimiterColumns, not even an explicit nil.
func (o *AlertDefinitionsCorrelationData) UnsetLimiterColumns() {
	o.LimiterColumns.Unset()
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetActions() Actions {
	if o == nil || o.Actions == nil {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetActionsOk() (*Actions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *AlertDefinitionsCorrelationData) SetActions(v Actions) {
	o.Actions = &v
}

// GetMaxSendCount returns the MaxSendCount field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetMaxSendCount() int64 {
	if o == nil || o.MaxSendCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxSendCount
}

// GetMaxSendCountOk returns a tuple with the MaxSendCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetMaxSendCountOk() (*int64, bool) {
	if o == nil || o.MaxSendCount == nil {
		return nil, false
	}
	return o.MaxSendCount, true
}

// HasMaxSendCount returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasMaxSendCount() bool {
	return o != nil && o.MaxSendCount != nil
}

// SetMaxSendCount gets a reference to the given int64 and assigns it to the MaxSendCount field.
func (o *AlertDefinitionsCorrelationData) SetMaxSendCount(v int64) {
	o.MaxSendCount = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetData() AlertDefinitionsQueryData {
	if o == nil || o.Data == nil {
		var ret AlertDefinitionsQueryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetDataOk() (*AlertDefinitionsQueryData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given AlertDefinitionsQueryData and assigns it to the Data field.
func (o *AlertDefinitionsCorrelationData) SetData(v AlertDefinitionsQueryData) {
	o.Data = &v
}

// GetCorrelationType returns the CorrelationType field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetCorrelationType() AlertDefinitionsCorrelationType {
	if o == nil {
		var ret AlertDefinitionsCorrelationType
		return ret
	}
	return o.CorrelationType
}

// GetCorrelationTypeOk returns a tuple with the CorrelationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetCorrelationTypeOk() (*AlertDefinitionsCorrelationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationType, true
}

// SetCorrelationType gets a reference to the given string and assigns it to the CorrelationType field.
func (o *AlertDefinitionsCorrelationData) SetCorrelationType(v AlertDefinitionsCorrelationType) {
	o.CorrelationType = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AlertDefinitionsCorrelationData) SetMessage(v string) {
	o.Message = &v
}

// GetGroupedColumns returns the GroupedColumns field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetGroupedColumns() []string {
	if o == nil || o.GroupedColumns == nil {
		var ret []string
		return ret
	}
	return o.GroupedColumns
}

// GetGroupedColumnsOk returns a tuple with the GroupedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetGroupedColumnsOk() (*[]string, bool) {
	if o == nil || o.GroupedColumns == nil {
		return nil, false
	}
	return &o.GroupedColumns, true
}

// HasGroupedColumns returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasGroupedColumns() bool {
	return o != nil && o.GroupedColumns != nil
}

// SetGroupedColumns gets a reference to the given []string and assigns it to the GroupedColumns field.
func (o *AlertDefinitionsCorrelationData) SetGroupedColumns(v []string) {
	o.GroupedColumns = v
}

// GetGroupedColumnsOptions returns the GroupedColumnsOptions field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetGroupedColumnsOptions() []SelectedColumn {
	if o == nil || o.GroupedColumnsOptions == nil {
		var ret []SelectedColumn
		return ret
	}
	return o.GroupedColumnsOptions
}

// GetGroupedColumnsOptionsOk returns a tuple with the GroupedColumnsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetGroupedColumnsOptionsOk() (*[]SelectedColumn, bool) {
	if o == nil || o.GroupedColumnsOptions == nil {
		return nil, false
	}
	return &o.GroupedColumnsOptions, true
}

// HasGroupedColumnsOptions returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasGroupedColumnsOptions() bool {
	return o != nil && o.GroupedColumnsOptions != nil
}

// SetGroupedColumnsOptions gets a reference to the given []SelectedColumn and assigns it to the GroupedColumnsOptions field.
func (o *AlertDefinitionsCorrelationData) SetGroupedColumnsOptions(v []SelectedColumn) {
	o.GroupedColumnsOptions = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AlertDefinitionsCorrelationData) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsCorrelationData) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AlertDefinitionsCorrelationData) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *AlertDefinitionsCorrelationData) SetVersion(v float64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertDefinitionsCorrelationData) MarshalJSON() ([]byte, error) {
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
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["Tags"] = o.Tags.Get()
	}
	if o.FromMarket != nil {
		toSerialize["FromMarket"] = o.FromMarket
	}
	if o.FromModules != nil {
		toSerialize["FromModules"] = o.FromModules
	}
	if o.HasUpdate != nil {
		toSerialize["HasUpdate"] = o.HasUpdate
	}
	if o.ModuleID.IsSet() {
		toSerialize["ModuleId"] = o.ModuleID.Get()
	}
	if o.ModuleGUID.IsSet() {
		toSerialize["ModuleGuid"] = o.ModuleGUID.Get()
	}
	if o.RiskLevel != nil {
		toSerialize["RiskLevel"] = o.RiskLevel
	}
	if o.MaxAlertCount != nil {
		toSerialize["MaxAlertCount"] = o.MaxAlertCount
	}
	if o.LimiterTimeFrameType.IsSet() {
		toSerialize["LimiterTimeFrameType"] = o.LimiterTimeFrameType.Get()
	}
	if o.LimiterTimeFrameValue.IsSet() {
		toSerialize["LimiterTimeFrameValue"] = o.LimiterTimeFrameValue.Get()
	}
	if o.LimiterColumns.IsSet() {
		toSerialize["LimiterColumns"] = o.LimiterColumns.Get()
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.MaxSendCount != nil {
		toSerialize["MaxSendCount"] = o.MaxSendCount
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if o.CorrelationType.IsValid() {
		toSerialize["CorrelationType"] = o.CorrelationType
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.GroupedColumns != nil {
		toSerialize["GroupedColumns"] = o.GroupedColumns
	}
	if o.GroupedColumnsOptions != nil {
		toSerialize["GroupedColumnsOptions"] = o.GroupedColumnsOptions
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
func (o *AlertDefinitionsCorrelationData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID                    *string                          `json:"Id,omitempty"`
		Enabled               *bool                            `json:"Enabled,omitempty"`
		Name                  *string                          `json:"Name,omitempty"`
		Description           common.NullableString            `json:"Description,omitempty"`
		Tags                  common.NullableList[string]      `json:"Tags,omitempty"`
		FromMarket            *bool                            `json:"FromMarket,omitempty"`
		FromModules           *bool                            `json:"FromModules,omitempty"`
		HasUpdate             *bool                            `json:"HasUpdate,omitempty"`
		ModuleID              common.NullableString            `json:"ModuleId,omitempty"`
		ModuleGUID            common.NullableString            `json:"ModuleGuid,omitempty"`
		RiskLevel             *int64                           `json:"RiskLevel,omitempty"`
		MaxAlertCount         *int64                           `json:"MaxAlertCount,omitempty"`
		LimiterTimeFrameType  common.NullableString            `json:"LimiterTimeFrameType,omitempty"`
		LimiterTimeFrameValue common.NullableInt64             `json:"LimiterTimeFrameValue,omitempty"`
		LimiterColumns        common.NullableList[string]      `json:"LimiterColumns,omitempty"`
		Actions               *Actions                         `json:"Actions,omitempty"`
		MaxSendCount          *int64                           `json:"MaxSendCount,omitempty"`
		Data                  *AlertDefinitionsQueryData       `json:"Data,omitempty"`
		CorrelationType       *AlertDefinitionsCorrelationType `json:"CorrelationType,omitempty"`
		Message               *string                          `json:"Message,omitempty"`
		GroupedColumns        []string                         `json:"GroupedColumns,omitempty"`
		GroupedColumnsOptions []SelectedColumn                 `json:"GroupedColumnsOptions,omitempty"`
		Version               *float64                         `json:"Version,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}
	if all.Message == nil {
		return fmt.Errorf("requiered field Message is missing")
	}
	if all.CorrelationType == nil {
		return fmt.Errorf("requiered field CorrelationType is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Id", "Enabled", "Name", "Query", "Description", "Tags", "FromMarket", "FromModules", "HasUpdate", "ModuleId", "ModuleGuid", "RiskLevel", "MaxAlertCount", "LimiterTimeFrameType", "LimiterTimeFrameValue", "LimiterColumns", "Actions", "MaxSendCount", "Data", "CorrelationType", "Message", "GroupedColumns", "GroupedColumnsOptions", "Version"})
	} else {
		return err
	}

	o.ID = all.ID
	o.Enabled = all.Enabled
	o.Name = all.Name
	o.Description = all.Description
	o.Tags = all.Tags
	o.FromMarket = all.FromMarket
	o.FromModules = all.FromModules
	o.HasUpdate = all.HasUpdate
	o.ModuleID = all.ModuleID
	o.ModuleGUID = all.ModuleGUID
	o.RiskLevel = all.RiskLevel
	o.MaxAlertCount = all.MaxAlertCount
	o.LimiterTimeFrameType = all.LimiterTimeFrameType
	o.LimiterTimeFrameValue = all.LimiterTimeFrameValue
	o.LimiterColumns = all.LimiterColumns
	hasInvalidField := false
	if all.Actions != nil && all.Actions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Actions = all.Actions
	o.MaxSendCount = all.MaxSendCount
	o.Data = all.Data
	if !all.CorrelationType.IsValid() {
		hasInvalidField = true
	} else {
		o.CorrelationType = *all.CorrelationType
	}
	o.Message = all.Message
	o.GroupedColumns = all.GroupedColumns
	o.GroupedColumnsOptions = all.GroupedColumnsOptions
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
