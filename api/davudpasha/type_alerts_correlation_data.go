package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AlertsCorrelationData struct {
	Id                    *string                     `json:"Id,omitempty"`
	Enabled               *bool                       `json:"Enabled,omitempty"`
	Name                  *string                     `json:"Name,omitempty"`
	Description           *string                     `json:"Description,omitempty"`
	Tags                  []string                    `json:"Tags,omitempty"`
	FromMarket            *bool                       `json:"FromMarket,omitempty"`
	FromModules           *bool                       `json:"FromModules,omitempty"`
	HasUpdate             *bool                       `json:"HasUpdate,omitempty"`
	ModuleId              common.NullableString       `json:"ModuleId,omitempty"`
	ModuleGuid            common.NullableString       `json:"ModuleGuid,omitempty"`
	RiskLevel             *int64                      `json:"RiskLevel,omitempty"`
	MaxAlertCount         *int64                      `json:"MaxAlertCount,omitempty"`
	LimiterTimeFrameType  *string                     `json:"LimiterTimeFrameType,omitempty"`
	LimiterTimeFrameValue *int64                      `json:"LimiterTimeFrameValue,omitempty"`
	LimiterColumns        common.NullableList[string] `json:"LimiterColumns,omitempty"`
	Actions               *Actions                    `json:"Actions,omitempty"`
	MaxSendCount          *int64                      `json:"MaxSendCount,omitempty"`
	Data                  *json.RawMessage            `json:"Data,omitempty"`
	CorrelationType       *string                     `json:"CorrelationType,omitempty"`
	Message               *string                     `json:"Message,omitempty"`
	GroupedColumns        []string                    `json:"GroupedColumns,omitempty"`
	Version               *float64                    `json:"Version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewAlertsCorrelationData creates a new AlertsCorrelationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertsCorrelationData() *AlertsCorrelationData {
	this := AlertsCorrelationData{}
	return &this
}

// NewAlertsCorrelationDataWithDefaults creates a new AlertsCorrelationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertsCorrelationDataWithDefaults() *AlertsCorrelationData {
	this := AlertsCorrelationData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertsCorrelationData) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertsCorrelationData) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertsCorrelationData) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertsCorrelationData) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AlertsCorrelationData) SetTags(v []string) {
	o.Tags = v
}

// GetFromMarket returns the FromMarket field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetFromMarket() bool {
	if o == nil || o.FromMarket == nil {
		var ret bool
		return ret
	}
	return *o.FromMarket
}

// GetFromMarketOk returns a tuple with the FromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetFromMarketOk() (*bool, bool) {
	if o == nil || o.FromMarket == nil {
		return nil, false
	}
	return o.FromMarket, true
}

// HasFromMarket returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasFromMarket() bool {
	return o != nil && o.FromMarket != nil
}

// SetFromMarket gets a reference to the given bool and assigns it to the FromMarket field.
func (o *AlertsCorrelationData) SetFromMarket(v bool) {
	o.FromMarket = &v
}

// GetFromModules returns the FromModules field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetFromModules() bool {
	if o == nil || o.FromModules == nil {
		var ret bool
		return ret
	}
	return *o.FromModules
}

// GetFromModulesOk returns a tuple with the FromModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetFromModulesOk() (*bool, bool) {
	if o == nil || o.FromModules == nil {
		return nil, false
	}
	return o.FromModules, true
}

// HasFromModules returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasFromModules() bool {
	return o != nil && o.FromModules != nil
}

// SetFromModules gets a reference to the given bool and assigns it to the FromModules field.
func (o *AlertsCorrelationData) SetFromModules(v bool) {
	o.FromModules = &v
}

// GetHasUpdate returns the HasUpdate field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetHasUpdate() bool {
	if o == nil || o.HasUpdate == nil {
		var ret bool
		return ret
	}
	return *o.HasUpdate
}

// GetHasUpdateOk returns a tuple with the HasUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetHasUpdateOk() (*bool, bool) {
	if o == nil || o.HasUpdate == nil {
		return nil, false
	}
	return o.HasUpdate, true
}

// HasHasUpdate returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasHasUpdate() bool {
	return o != nil && o.HasUpdate != nil
}

// SetHasUpdate gets a reference to the given bool and assigns it to the HasUpdate field.
func (o *AlertsCorrelationData) SetHasUpdate(v bool) {
	o.HasUpdate = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsCorrelationData) GetModuleId() string {
	if o == nil || o.ModuleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleId.Get()
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetModuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleId.Get(), o.ModuleId.IsSet()
}

// HasModuleId returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasModuleId() bool {
	return o != nil && o.ModuleId.IsSet()
}

// SetModuleId gets a reference to the given datadog.NullableString and assigns it to the ModuleId field.
func (o *AlertsCorrelationData) SetModuleId(v string) {
	o.ModuleId.Set(&v)
}

// SetModuleIdNil sets the value for ModuleId to be an explicit nil.
func (o *AlertsCorrelationData) SetModuleIdNil() {
	o.ModuleId.Set(nil)
}

// UnsetModuleId ensures that no value is present for ModuleId, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetModuleId() {
	o.ModuleId.UnSet()
}

// GetModuleGuid returns the ModuleGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsCorrelationData) GetModuleGuid() string {
	if o == nil || o.ModuleGuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleGuid.Get()
}

// GetModuleGuidOk returns a tuple with the ModuleGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetModuleGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleGuid.Get(), o.ModuleGuid.IsSet()
}

// HasModuleGuid returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasModuleGuid() bool {
	return o != nil && o.ModuleGuid.IsSet()
}

// SetModuleGuid gets a reference to the given datadog.NullableString and assigns it to the ModuleGuid field.
func (o *AlertsCorrelationData) SetModuleGuid(v string) {
	o.ModuleGuid.Set(&v)
}

// SetModuleGuidNil sets the value for ModuleGuid to be an explicit nil.
func (o *AlertsCorrelationData) SetModuleGuidNil() {
	o.ModuleGuid.Set(nil)
}

// UnsetModuleGuid ensures that no value is present for ModuleGuid, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetModuleGuid() {
	o.ModuleGuid.UnSet()
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetRiskLevel() int64 {
	if o == nil || o.RiskLevel == nil {
		var ret int64
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetRiskLevelOk() (*int64, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasRiskLevel() bool {
	return o != nil && o.RiskLevel != nil
}

// SetRiskLevel gets a reference to the given int64 and assigns it to the RiskLevel field.
func (o *AlertsCorrelationData) SetRiskLevel(v int64) {
	o.RiskLevel = &v
}

// GetMaxAlertCount returns the MaxAlertCount field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetMaxAlertCount() int64 {
	if o == nil || o.MaxAlertCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxAlertCount
}

// GetMaxAlertCountOk returns a tuple with the MaxAlertCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetMaxAlertCountOk() (*int64, bool) {
	if o == nil || o.MaxAlertCount == nil {
		return nil, false
	}
	return o.MaxAlertCount, true
}

// HasMaxAlertCount returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasMaxAlertCount() bool {
	return o != nil && o.MaxAlertCount != nil
}

// SetMaxAlertCount gets a reference to the given int64 and assigns it to the MaxAlertCount field.
func (o *AlertsCorrelationData) SetMaxAlertCount(v int64) {
	o.MaxAlertCount = &v
}

// GetLimiterTimeFrameType returns the LimiterTimeFrameType field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetLimiterTimeFrameType() string {
	if o == nil || o.LimiterTimeFrameType == nil {
		var ret string
		return ret
	}
	return *o.LimiterTimeFrameType
}

// GetLimiterTimeFrameTypeOk returns a tuple with the LimiterTimeFrameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetLimiterTimeFrameTypeOk() (*string, bool) {
	if o == nil || o.LimiterTimeFrameType == nil {
		return nil, false
	}
	return o.LimiterTimeFrameType, true
}

// HasLimiterTimeFrameType returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasLimiterTimeFrameType() bool {
	return o != nil && o.LimiterTimeFrameType != nil
}

// SetLimiterTimeFrameType gets a reference to the given string and assigns it to the LimiterTimeFrameType field.
func (o *AlertsCorrelationData) SetLimiterTimeFrameType(v string) {
	o.LimiterTimeFrameType = &v
}

// GetLimiterTimeFrameValue returns the LimiterTimeFrameValue field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetLimiterTimeFrameValue() int64 {
	if o == nil || o.LimiterTimeFrameValue == nil {
		var ret int64
		return ret
	}
	return *o.LimiterTimeFrameValue
}

// GetLimiterTimeFrameValueOk returns a tuple with the LimiterTimeFrameValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetLimiterTimeFrameValueOk() (*int64, bool) {
	if o == nil || o.LimiterTimeFrameValue == nil {
		return nil, false
	}
	return o.LimiterTimeFrameValue, true
}

// HasLimiterTimeFrameValue returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasLimiterTimeFrameValue() bool {
	return o != nil && o.LimiterTimeFrameValue != nil
}

// SetLimiterTimeFrameValue gets a reference to the given int64 and assigns it to the LimiterTimeFrameValue field.
func (o *AlertsCorrelationData) SetLimiterTimeFrameValue(v int64) {
	o.LimiterTimeFrameValue = &v
}

// GetLimiterColumns returns a tuple with the LimiterColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetLimiterColumns() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimiterColumns.Get(), o.LimiterColumns.IsSet()
}

// GetLimiterColumnsOk returns a tuple with the LimiterColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetLimiterColumnsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimiterColumns.Get(), o.LimiterColumns.IsSet()
}

// HasLimiterColumns returns a boolean if a LimiterColumns has been set.
func (o *AlertsCorrelationData) HasLimiterColumns() bool {
	return o != nil && o.LimiterColumns.IsSet()
}

// SetLimiterColumns gets a reference to the given datadog.Nullable[]string and assigns it to the LimiterColumns field.
func (o *AlertsCorrelationData) SetLimiterColumns(v []string) {
	o.LimiterColumns.Set(&v)
}

// SetLimiterColumnsNil sets the value for LimiterColumns to be an explicit nil.
func (o *AlertsCorrelationData) SetLimiterColumnsNil() {
	o.LimiterColumns.Set(nil)
}

// UnsetLimiterColumns ensures that no value is present for LimiterColumns, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetLimiterColumns() {
	o.LimiterColumns.UnSet()
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetActions() Actions {
	if o == nil || o.Actions == nil {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetActionsOk() (*Actions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *AlertsCorrelationData) SetActions(v Actions) {
	o.Actions = &v
}

// GetMaxSendCount returns the MaxSendCount field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetMaxSendCount() int64 {
	if o == nil || o.MaxSendCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxSendCount
}

// GetMaxSendCountOk returns a tuple with the MaxSendCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetMaxSendCountOk() (*int64, bool) {
	if o == nil || o.MaxSendCount == nil {
		return nil, false
	}
	return o.MaxSendCount, true
}

// HasMaxSendCount returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasMaxSendCount() bool {
	return o != nil && o.MaxSendCount != nil
}

// SetMaxSendCount gets a reference to the given int64 and assigns it to the MaxSendCount field.
func (o *AlertsCorrelationData) SetMaxSendCount(v int64) {
	o.MaxSendCount = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetData() json.RawMessage {
	if o == nil || o.Data == nil {
		var ret json.RawMessage
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetDataOk() (*json.RawMessage, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given json.RawMessage and assigns it to the Data field.
func (o *AlertsCorrelationData) SetData(v json.RawMessage) {
	o.Data = &v
}

// GetCorrelationType returns the CorrelationType field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetCorrelationType() string {
	if o == nil || o.CorrelationType == nil {
		var ret string
		return ret
	}
	return *o.CorrelationType
}

// GetCorrelationTypeOk returns a tuple with the CorrelationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetCorrelationTypeOk() (*string, bool) {
	if o == nil || o.CorrelationType == nil {
		return nil, false
	}
	return o.CorrelationType, true
}

// HasCorrelationType returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasCorrelationType() bool {
	return o != nil && o.CorrelationType != nil
}

// SetCorrelationType gets a reference to the given string and assigns it to the CorrelationType field.
func (o *AlertsCorrelationData) SetCorrelationType(v string) {
	o.CorrelationType = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AlertsCorrelationData) SetMessage(v string) {
	o.Message = &v
}

// GetGroupedColumns returns the GroupedColumns field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetGroupedColumns() []string {
	if o == nil || o.GroupedColumns == nil {
		var ret []string
		return ret
	}
	return o.GroupedColumns
}

// GetGroupedColumnsOk returns a tuple with the GroupedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetGroupedColumnsOk() (*[]string, bool) {
	if o == nil || o.GroupedColumns == nil {
		return nil, false
	}
	return &o.GroupedColumns, true
}

// HasGroupedColumns returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasGroupedColumns() bool {
	return o != nil && o.GroupedColumns != nil
}

// SetGroupedColumns gets a reference to the given []string and assigns it to the GroupedColumns field.
func (o *AlertsCorrelationData) SetGroupedColumns(v []string) {
	o.GroupedColumns = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *AlertsCorrelationData) SetVersion(v float64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertsCorrelationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
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
	if o.ModuleId.IsSet() {
		toSerialize["ModuleId"] = o.ModuleId.Get()
	}
	if o.ModuleGuid.IsSet() {
		toSerialize["ModuleGuid"] = o.ModuleGuid.Get()
	}
	if o.RiskLevel != nil {
		toSerialize["RiskLevel"] = o.RiskLevel
	}
	if o.MaxAlertCount != nil {
		toSerialize["MaxAlertCount"] = o.MaxAlertCount
	}
	if o.LimiterTimeFrameType != nil {
		toSerialize["LimiterTimeFrameType"] = o.LimiterTimeFrameType
	}
	if o.LimiterTimeFrameValue != nil {
		toSerialize["LimiterTimeFrameValue"] = o.LimiterTimeFrameValue
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
	if o.CorrelationType != nil {
		toSerialize["CorrelationType"] = o.CorrelationType
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.GroupedColumns != nil {
		toSerialize["GroupedColumns"] = o.GroupedColumns
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
func (o *AlertsCorrelationData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                    *string                     `json:"Id,omitempty"`
		Enabled               *bool                       `json:"Enabled,omitempty"`
		Name                  *string                     `json:"Name,omitempty"`
		Description           *string                     `json:"Description,omitempty"`
		Tags                  []string                    `json:"Tags,omitempty"`
		FromMarket            *bool                       `json:"FromMarket,omitempty"`
		FromModules           *bool                       `json:"FromModules,omitempty"`
		HasUpdate             *bool                       `json:"HasUpdate,omitempty"`
		ModuleId              common.NullableString       `json:"ModuleId,omitempty"`
		ModuleGuid            common.NullableString       `json:"ModuleGuid,omitempty"`
		RiskLevel             *int64                      `json:"RiskLevel,omitempty"`
		MaxAlertCount         *int64                      `json:"MaxAlertCount,omitempty"`
		LimiterTimeFrameType  *string                     `json:"LimiterTimeFrameType,omitempty"`
		LimiterTimeFrameValue *int64                      `json:"LimiterTimeFrameValue,omitempty"`
		LimiterColumns        common.NullableList[string] `json:"LimiterColumns,omitempty"`
		Actions               *Actions                    `json:"Actions,omitempty"`
		MaxSendCount          *int64                      `json:"MaxSendCount,omitempty"`
		Data                  *json.RawMessage            `json:"Data,omitempty"`
		CorrelationType       *string                     `json:"CorrelationType,omitempty"`
		Message               *string                     `json:"Message,omitempty"`
		GroupedColumns        []string                    `json:"GroupedColumns,omitempty"`
		Version               *float64                    `json:"Version,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("requiered field Id is missing")
	}
	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}
	if all.Message == nil {
		return fmt.Errorf("requiered field Message is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Id", "Enabled", "Name", "Query", "Description", "Tags", "FromMarket", "FromModules", "HasUpdate", "ModuleId", "ModuleGuid", "RiskLevel", "MaxAlertCount", "LimiterTimeFrameType", "LimiterTimeFrameValue", "LimiterColumns", "Actions", "MaxSendCount", "Data", "CorrelationType", "Message", "GroupedColumns", "Version"})
	} else {
		return err
	}

	o.Id = all.Id
	o.Enabled = all.Enabled
	o.Name = all.Name
	o.Description = all.Description
	o.Tags = all.Tags
	o.FromMarket = all.FromMarket
	o.FromModules = all.FromModules
	o.HasUpdate = all.HasUpdate
	o.ModuleId = all.ModuleId
	o.ModuleGuid = all.ModuleGuid
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
	o.CorrelationType = all.CorrelationType
	o.Message = all.Message
	o.GroupedColumns = all.GroupedColumns
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
