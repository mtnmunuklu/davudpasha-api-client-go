package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type QueriesItems struct {
	ID                   *string                              `json:"ID,omitempty"`
	Name                 *string                              `json:"Name,omitempty"`
	Description          common.NullableString                `json:"Description,omitempty"`
	Query                *string                              `json:"Query,omitempty"`
	Columns              common.NullableList[SelectedColumns] `json:"Columns,omitempty"`
	Author               *string                              `json:"Author,omitempty"`
	InsertDate           *string                              `json:"InsertDate,omitempty"`
	LastUpdateDate       *string                              `json:"LastUpdateDate,omitempty"`
	QueryType            *string                              `json:"QueryType,omitempty"`
	DateTimeRange        *DateTimeRange                       `json:"DateTimeRange,omitempty"`
	Tags                 common.NullableList[string]          `json:"Tags,omitempty"`
	MitreTags            common.NullableList[string]          `json:"MitreTags,omitempty"`
	KillChainPhase       common.NullableString                `json:"KillChainPhase,omitempty"`
	FromMarket           *bool                                `json:"FromMarket,omitempty"`
	FromModules          *bool                                `json:"FromModules,omitempty"`
	HasUpdate            *bool                                `json:"HasUpdate,omitempty"`
	ModuleId             common.NullableString                `json:"ModuleId,omitempty"`
	ModuleGuid           common.NullableString                `json:"ModuleGuid,omitempty"`
	SharedUsersAndGroups common.NullableList[string]          `json:"SharedUsersAndGroups,omitempty"`
	Version              *float64                             `json:"Version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewQueriesItems creates a new QueriesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueriesItems() *QueriesItems {
	this := QueriesItems{}
	return &this
}

// NewQueriesItemsWithDefaults creates a new QueriesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueriesItemsWithDefaults() *QueriesItems {
	this := QueriesItems{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *QueriesItems) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *QueriesItems) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *QueriesItems) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueriesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueriesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueriesItems) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItems) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItems) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a Description has been set.
func (o *QueriesItems) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *QueriesItems) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *QueriesItems) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *QueriesItems) UnsetDescription() {
	o.Description.UnSet()
}

// GetColumns returns the Columns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItems) GetColumns() []SelectedColumns {
	if o == nil || o.Columns.Get() == nil {
		var ret []SelectedColumns
		return ret
	}
	return *o.Columns.Get()
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItems) GetColumnsOk() (*[]SelectedColumns, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns.Get(), o.Columns.IsSet()
}

// HasColumns returns a boolean if a Columns has been set.
func (o *QueriesItems) HasColumns() bool {
	return o != nil && o.Columns.IsSet()
}

// SetColumns gets a reference to the given datadog.Nullable[SelectedColumns] and assigns it to the Columns field.
func (o *QueriesItems) SetColumns(v []SelectedColumns) {
	o.Columns.Set(&v)
}

// SetColumnsNil sets the value for Columns to be an explicit nil.
func (o *QueriesItems) SetColumnsNil() {
	o.Columns.Set(nil)
}

// UnsetColumns ensures that no value is present for Columns, not even an explicit nil.
func (o *QueriesItems) UnSetColumns() {
	o.Columns.UnSet()
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *QueriesItems) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *QueriesItems) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *QueriesItems) SetAuthor(v string) {
	o.Author = &v
}

// GetInsertDate returns the InsertDate field value if set, zero value otherwise.
func (o *QueriesItems) GetInsertDate() string {
	if o == nil || o.InsertDate == nil {
		var ret string
		return ret
	}
	return *o.InsertDate
}

// GetInsertDateOk returns a tuple with the InsertDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetInsertDateOk() (*string, bool) {
	if o == nil || o.InsertDate == nil {
		return nil, false
	}
	return o.InsertDate, true
}

// HasInsertDate returns a boolean if a field has been set.
func (o *QueriesItems) HasInsertDate() bool {
	return o != nil && o.InsertDate != nil
}

// SetInsertDate gets a reference to the given string and assigns it to the InsertDate field.
func (o *QueriesItems) SetInsertDate(v string) {
	o.InsertDate = &v
}

// GetLastUpdateDate returns the LastUpdateDate field value if set, zero value otherwise.
func (o *QueriesItems) GetLastUpdateDate() string {
	if o == nil || o.LastUpdateDate == nil {
		var ret string
		return ret
	}
	return *o.LastUpdateDate
}

// GetLastUpdateDateOk returns a tuple with the LastUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetLastUpdateDateOk() (*string, bool) {
	if o == nil || o.LastUpdateDate == nil {
		return nil, false
	}
	return o.LastUpdateDate, true
}

// HasLastUpdateDate returns a boolean if a field has been set.
func (o *QueriesItems) HasLastUpdateDate() bool {
	return o != nil && o.LastUpdateDate != nil
}

// SetLastUpdateDate gets a reference to the given string and assigns it to the LastUpdateDate field.
func (o *QueriesItems) SetLastUpdateDate(v string) {
	o.LastUpdateDate = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *QueriesItems) GetQueryType() string {
	if o == nil || o.QueryType == nil {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetQueryTypeOk() (*string, bool) {
	if o == nil || o.QueryType == nil {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *QueriesItems) HasQueryType() bool {
	return o != nil && o.QueryType != nil
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *QueriesItems) SetQueryType(v string) {
	o.QueryType = &v
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *QueriesItems) GetDateTimeRange() DateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret DateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *QueriesItems) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DateTimeRange field.
func (o *QueriesItems) SetDateTimeRange(v DateTimeRange) {
	o.DateTimeRange = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItems) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItems) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a Tags has been set.
func (o *QueriesItems) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.Nullable[]string and assigns it to the Tags field.
func (o *QueriesItems) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *QueriesItems) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *QueriesItems) UnSetTags() {
	o.Tags.UnSet()
}

// GetMitreTags returns the MitreTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItems) GetMitreTags() []string {
	if o == nil || o.MitreTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.MitreTags.Get()
}

// GetMitreTagsOk returns a tuple with the MitreTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItems) GetMitreTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitreTags.Get(), o.MitreTags.IsSet()
}

// HasMitreTags returns a boolean if a MitreTags has been set.
func (o *QueriesItems) HasMitreTags() bool {
	return o != nil && o.MitreTags.IsSet()
}

// SetMitreTags gets a reference to the given datadog.Nullable[]string and assigns it to the MitreTags field.
func (o *QueriesItems) SetMitreTags(v []string) {
	o.MitreTags.Set(&v)
}

// SetMitreTagsNil sets the value for MitreTags to be an explicit nil.
func (o *QueriesItems) SetMitreTagsNil() {
	o.MitreTags.Set(nil)
}

// UnsetMitreTags ensures that no value is present for MitreTags, not even an explicit nil.
func (o *QueriesItems) UnSetMitreTags() {
	o.MitreTags.UnSet()
}

// GetKillChainPhase returns the KillChainPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItems) GetKillChainPhase() string {
	if o == nil || o.KillChainPhase.Get() == nil {
		var ret string
		return ret
	}
	return *o.KillChainPhase.Get()
}

// GetKillChainPhaseOk returns a tuple with the KillChainPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItems) GetKillChainPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KillChainPhase.Get(), o.KillChainPhase.IsSet()
}

// HasKillChainPhase returns a boolean if a KillChainPhase has been set.
func (o *QueriesItems) HasKillChainPhase() bool {
	return o != nil && o.KillChainPhase.IsSet()
}

// SetKillChainPhase gets a reference to the given datadog.NullableString and assigns it to the KillChainPhase field.
func (o *QueriesItems) SetKillChainPhase(v string) {
	o.KillChainPhase.Set(&v)
}

// SetKillChainPhaseNil sets the value for KillChainPhase to be an explicit nil.
func (o *QueriesItems) SetKillChainPhaseNil() {
	o.KillChainPhase.Set(nil)
}

// UnsetKillChainPhase ensures that no value is present for KillChainPhase, not even an explicit nil.
func (o *QueriesItems) UnsetKillChainPhase() {
	o.KillChainPhase.UnSet()
}

// GetFromMarket returns the FromMarket field value if set, zero value otherwise.
func (o *QueriesItems) GetFromMarket() bool {
	if o == nil || o.FromMarket == nil {
		var ret bool
		return ret
	}
	return *o.FromMarket
}

// GetFromMarketOk returns a tuple with the FromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetFromMarketOk() (*bool, bool) {
	if o == nil || o.FromMarket == nil {
		return nil, false
	}
	return o.FromMarket, true
}

// HasFromMarket returns a boolean if a field has been set.
func (o *QueriesItems) HasFromMarket() bool {
	return o != nil && o.FromMarket != nil
}

// SetFromMarket gets a reference to the given bool and assigns it to the FromMarket field.
func (o *QueriesItems) SetFromMarket(v bool) {
	o.FromMarket = &v
}

// GetFromModules returns the FromModules field value if set, zero value otherwise.
func (o *QueriesItems) GetFromModules() bool {
	if o == nil || o.FromModules == nil {
		var ret bool
		return ret
	}
	return *o.FromModules
}

// GetFromModulesOk returns a tuple with the FromModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetFromModulesOk() (*bool, bool) {
	if o == nil || o.FromModules == nil {
		return nil, false
	}
	return o.FromModules, true
}

// HasFromModules returns a boolean if a field has been set.
func (o *QueriesItems) HasFromModules() bool {
	return o != nil && o.FromModules != nil
}

// SetFromModules gets a reference to the given bool and assigns it to the FromModules field.
func (o *QueriesItems) SetFromModules(v bool) {
	o.FromModules = &v
}

// GetHasUpdate returns the HasUpdate field value if set, zero value otherwise.
func (o *QueriesItems) GetHasUpdate() bool {
	if o == nil || o.HasUpdate == nil {
		var ret bool
		return ret
	}
	return *o.HasUpdate
}

// GetHasUpdateOk returns a tuple with the HasUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetHasUpdateOk() (*bool, bool) {
	if o == nil || o.HasUpdate == nil {
		return nil, false
	}
	return o.HasUpdate, true
}

// HasHasUpdate returns a boolean if a field has been set.
func (o *QueriesItems) HasHasUpdate() bool {
	return o != nil && o.HasUpdate != nil
}

// SetHasUpdate gets a reference to the given bool and assigns it to the HasUpdate field.
func (o *QueriesItems) SetHasUpdate(v bool) {
	o.HasUpdate = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItems) GetModuleId() string {
	if o == nil || o.ModuleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleId.Get()
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItems) GetModuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleId.Get(), o.ModuleId.IsSet()
}

// HasModuleId returns a boolean if a ModuleId has been set.
func (o *QueriesItems) HasModuleId() bool {
	return o != nil && o.ModuleId.IsSet()
}

// SetModuleId gets a reference to the given datadog.NullableString and assigns it to the ModuleId field.
func (o *QueriesItems) SetModuleId(v string) {
	o.ModuleId.Set(&v)
}

// SetModuleIdNil sets the value for ModuleId to be an explicit nil.
func (o *QueriesItems) SetModuleIdNil() {
	o.ModuleId.Set(nil)
}

// UnsetModuleId ensures that no value is present for ModuleId, not even an explicit nil.
func (o *QueriesItems) UnsetModuleId() {
	o.ModuleId.UnSet()
}

// GetModuleGuid returns the ModuleGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItems) GetModuleGuid() string {
	if o == nil || o.ModuleGuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleGuid.Get()
}

// GetModuleGuidOk returns a tuple with the ModuleGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItems) GetModuleGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleGuid.Get(), o.ModuleGuid.IsSet()
}

// HasModuleGuid returns a boolean if a ModuleGuid has been set.
func (o *QueriesItems) HasModuleGuid() bool {
	return o != nil && o.ModuleGuid.IsSet()
}

// SetModuleGuid gets a reference to the given datadog.NullableString and assigns it to the ModuleGuid field.
func (o *QueriesItems) SetModuleGuid(v string) {
	o.ModuleGuid.Set(&v)
}

// SetModuleGuidNil sets the value for ModuleGuid to be an explicit nil.
func (o *QueriesItems) SetModuleGuidNil() {
	o.ModuleGuid.Set(nil)
}

// UnsetModuleGuid ensures that no value is present for ModuleGuid, not even an explicit nil.
func (o *QueriesItems) UnsetModuleGuid() {
	o.ModuleGuid.UnSet()
}

// GetSharedUsersAndGroups returns the SharedUsersAndGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItems) GetSharedUsersAndGroups() []string {
	if o == nil || o.SharedUsersAndGroups.Get() == nil {
		var ret []string
		return ret
	}
	return *o.SharedUsersAndGroups.Get()
}

// GetSharedUsersAndGroupsOk returns a tuple with the SharedUsersAndGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItems) GetSharedUsersAndGroupsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedUsersAndGroups.Get(), o.SharedUsersAndGroups.IsSet()
}

// HasSharedUsersAndGroups returns a boolean if a SharedUsersAndGroups has been set.
func (o *QueriesItems) HasSharedUsersAndGroups() bool {
	return o != nil && o.SharedUsersAndGroups.IsSet()
}

// SetSharedUsersAndGroups gets a reference to the given datadog.Nullable[]string and assigns it to the SharedUsersAndGroups field.
func (o *QueriesItems) SetSharedUsersAndGroups(v []string) {
	o.SharedUsersAndGroups.Set(&v)
}

// SetSharedUsersAndGroupsNil sets the value for SharedUsersAndGroups to be an explicit nil.
func (o *QueriesItems) SetSharedUsersAndGroupsNil() {
	o.SharedUsersAndGroups.Set(nil)
}

// UnsetSharedUsersAndGroups ensures that no value is present for SharedUsersAndGroups, not even an explicit nil.
func (o *QueriesItems) UnSetSharedUsersAndGroups() {
	o.SharedUsersAndGroups.UnSet()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *QueriesItems) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItems) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *QueriesItems) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *QueriesItems) SetVersion(v float64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueriesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}
	if o.Columns.IsSet() {
		toSerialize["Columns"] = o.Columns
	}
	if o.Author != nil {
		toSerialize["Author"] = o.Author
	}
	if o.InsertDate != nil {
		toSerialize["InsertDate"] = o.InsertDate
	}
	if o.LastUpdateDate != nil {
		toSerialize["LastUpdateDate"] = o.LastUpdateDate
	}
	if o.QueryType != nil {
		toSerialize["QueryType"] = o.QueryType
	}
	if o.QueryType != nil {
		toSerialize["QueryType"] = o.QueryType
	}
	if o.DateTimeRange != nil {
		toSerialize["DateTimeRange"] = o.DateTimeRange
	}
	if o.Tags.IsSet() {
		toSerialize["Tags"] = o.Tags
	}
	if o.MitreTags.IsSet() {
		toSerialize["MitreTags"] = o.MitreTags
	}
	if o.KillChainPhase.IsSet() {
		toSerialize["KillChainPhase"] = o.KillChainPhase
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
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.ModuleGuid.IsSet() {
		toSerialize["ModuleGuid"] = o.ModuleGuid
	}
	if o.SharedUsersAndGroups.IsSet() {
		toSerialize["SharedUsersAndGroups"] = o.SharedUsersAndGroups
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
func (o *QueriesItems) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID                   *string                              `json:"ID,omitempty"`
		Name                 *string                              `json:"Name,omitempty"`
		Description          common.NullableString                `json:"Description,omitempty"`
		Query                *string                              `json:"Query,omitempty"`
		Columns              common.NullableList[SelectedColumns] `json:"Columns,omitempty"`
		Author               *string                              `json:"Author,omitempty"`
		InsertDate           *string                              `json:"InsertDate,omitempty"`
		LastUpdateDate       *string                              `json:"LastUpdateDate,omitempty"`
		QueryType            *string                              `json:"QueryType,omitempty"`
		DateTimeRange        *DateTimeRange                       `json:"DateTimeRange,omitempty"`
		Tags                 common.NullableList[string]          `json:"Tags,omitempty"`
		MitreTags            common.NullableList[string]          `json:"MitreTags,omitempty"`
		KillChainPhase       common.NullableString                `json:"KillChainPhase,omitempty"`
		FromMarket           *bool                                `json:"FromMarket,omitempty"`
		FromModules          *bool                                `json:"FromModules,omitempty"`
		HasUpdate            *bool                                `json:"HasUpdate,omitempty"`
		ModuleId             common.NullableString                `json:"ModuleId,omitempty"`
		ModuleGuid           common.NullableString                `json:"ModuleGuid,omitempty"`
		SharedUsersAndGroups common.NullableList[string]          `json:"SharedUsersAndGroups,omitempty"`
		Version              *float64                             `json:"Version,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ID == nil {
		return fmt.Errorf("requiered field ID is missing")
	}
	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}
	if all.Query == nil {
		return fmt.Errorf("requiered field Query is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ID", "Name", "Description", "Query", "Columns", "Author", "InsertDate", "LastUpdateDate", "QueryType", "DateTimeRange", "Tags", "MitreTags", "KillChainPhase", "FromMarket", "FromModules", "ModuleId", "ModuleGuid", "SharedUsersAndGroups", "Version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	o.ID = all.ID
	o.Name = all.Name
	o.Description = all.Description
	o.Query = all.Query
	o.Columns = all.Columns
	o.Author = all.Author
	o.InsertDate = all.InsertDate
	o.LastUpdateDate = all.LastUpdateDate
	o.QueryType = all.QueryType
	o.Tags = all.Tags
	o.MitreTags = all.MitreTags
	o.KillChainPhase = all.KillChainPhase
	o.FromMarket = all.FromMarket
	o.FromModules = all.FromModules
	o.HasUpdate = all.HasUpdate
	o.ModuleId = all.ModuleId
	o.ModuleGuid = all.ModuleGuid
	o.SharedUsersAndGroups = all.SharedUsersAndGroups
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
