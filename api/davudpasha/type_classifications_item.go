package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ClassificationsItem represents an item in classifications.
type ClassificationsItem struct {
	// Unique identifier for the classification item.
	ID *string `json:"ID,omitempty"`
	// Definition ID associated with the classification item.
	DefID *int64 `json:"DefID,omitempty"`
	// Name of the classification item.
	Name *string `json:"Name,omitempty"`
	// Severity level associated with the classification item.
	Severity *string `json:"Severity,omitempty"`
	// List of MITRE tags associated with the classification item, which can be null.
	MitreTags common.NullableList[MitreTag] `json:"MitreTags,omitempty"`
	// Kill chain phase associated with the classification item, which can be null.
	KillChainPhase common.NullableList[string] `json:"KillChainPhase,omitempty"`
	// Flag indicating if the classification item is from the market.
	FromMarket *bool `json:"FromMarket,omitempty"`
	// Flag indicating if the classification item is from modules.
	FromModules *bool `json:"FromModules,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClassificationsItem creates a new ClassificationsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClassificationsItem() *ClassificationsItem {
	this := ClassificationsItem{}
	return &this
}

// NewClassificationsItemWithDefaults creates a new ClassificationsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClassificationsItemWithDefaults() *ClassificationsItem {
	this := ClassificationsItem{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ClassificationsItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassificationsItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ClassificationsItem) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ClassificationsItem) SetID(v string) {
	o.ID = &v
}

// GetDefID returns the DefID field value if set, zero value otherwise.
func (o *ClassificationsItem) GetDefID() int64 {
	if o == nil || o.DefID == nil {
		var ret int64
		return ret
	}
	return *o.DefID
}

// GetDefIDOk returns a tuple with the DefID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassificationsItem) GetDefIDOk() (*int64, bool) {
	if o == nil || o.DefID == nil {
		return nil, false
	}
	return o.DefID, true
}

// HasDefID returns a boolean if a field has been set.
func (o *ClassificationsItem) HasDefID() bool {
	return o != nil && o.DefID != nil
}

// SetDefID gets a reference to the given int64 and assigns it to the DefID field.
func (o *ClassificationsItem) SetDefID(v int64) {
	o.DefID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClassificationsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassificationsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClassificationsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClassificationsItem) SetName(v string) {
	o.Name = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ClassificationsItem) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassificationsItem) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ClassificationsItem) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ClassificationsItem) SetSeverity(v string) {
	o.Severity = &v
}

// GetMitreTags returns the MitreTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClassificationsItem) GetMitreTags() []MitreTag {
	if o == nil || o.MitreTags.Get() == nil {
		var ret []MitreTag
		return ret
	}
	return *o.MitreTags.Get()
}

// GetMitreTagsOk returns a tuple with the MitreTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClassificationsItem) GetMitreTagsOk() (*[]MitreTag, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitreTags.Get(), o.MitreTags.IsSet()
}

// HasMitreTags returns a boolean if a MitreTags has been set.
func (o *ClassificationsItem) HasMitreTags() bool {
	return o != nil && o.MitreTags.IsSet()
}

// SetMitreTags gets a reference to the given common.Nullable[]MitreTag and assigns it to the MitreTags field.
func (o *ClassificationsItem) SetMitreTags(v []MitreTag) {
	o.MitreTags.Set(&v)
}

// SetMitreTagsNil sets the value for MitreTags to be an explicit nil.
func (o *ClassificationsItem) SetMitreTagsNil() {
	o.MitreTags.Set(nil)
}

// UnSetMitreTags ensures that no value is present for MitreTags, not even an explicit nil.
func (o *ClassificationsItem) UnsetMitreTags() {
	o.MitreTags.Unset()
}

// GetKillChainPhase returns the KillChainPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClassificationsItem) GetKillChainPhase() []string {
	if o == nil || o.KillChainPhase.Get() == nil {
		var ret []string
		return ret
	}
	return *o.KillChainPhase.Get()
}

// GetKillChainPhaseOk returns a tuple with the KillChainPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClassificationsItem) GetKillChainPhaseOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KillChainPhase.Get(), o.KillChainPhase.IsSet()
}

// HasKillChainPhase returns a boolean if a KillChainPhase has been set.
func (o *ClassificationsItem) HasKillChainPhase() bool {
	return o != nil && o.KillChainPhase.IsSet()
}

// SetKillChainPhase gets a reference to the given common.Nullable[]string and assigns it to the KillChainPhase field.
func (o *ClassificationsItem) SetKillChainPhase(v []string) {
	o.KillChainPhase.Set(&v)
}

// SetKillChainPhaseNil sets the value for KillChainPhase to be an explicit nil.
func (o *ClassificationsItem) SetKillChainPhaseNil() {
	o.KillChainPhase.Set(nil)
}

// UnSetKillChainPhase ensures that no value is present for KillChainPhase, not even an explicit nil.
func (o *ClassificationsItem) UnsetKillChainPhase() {
	o.KillChainPhase.Unset()
}

// GetFromMarket returns the FromMarket field value if set, zero value otherwise.
func (o *ClassificationsItem) GetFromMarket() bool {
	if o == nil || o.FromMarket == nil {
		var ret bool
		return ret
	}
	return *o.FromMarket
}

// GetFromMarketOk returns a tuple with the FromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassificationsItem) GetFromMarketOk() (*bool, bool) {
	if o == nil || o.FromMarket == nil {
		return nil, false
	}
	return o.FromMarket, true
}

// HasFromMarket returns a boolean if a field has been set.
func (o *ClassificationsItem) HasFromMarket() bool {
	return o != nil && o.FromMarket != nil
}

// SetFromMarket gets a reference to the given bool and assigns it to the FromMarket field.
func (o *ClassificationsItem) SetFromMarket(v bool) {
	o.FromMarket = &v
}

// GetFromModules returns the FromModules field value if set, zero value otherwise.
func (o *ClassificationsItem) GetFromModules() bool {
	if o == nil || o.FromModules == nil {
		var ret bool
		return ret
	}
	return *o.FromModules
}

// GetFromModulesOk returns a tuple with the FromModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassificationsItem) GetFromModulesOk() (*bool, bool) {
	if o == nil || o.FromModules == nil {
		return nil, false
	}
	return o.FromModules, true
}

// HasFromModules returns a boolean if a field has been set.
func (o *ClassificationsItem) HasFromModules() bool {
	return o != nil && o.FromModules != nil
}

// SetFromModules gets a reference to the given bool and assigns it to the FromModules field.
func (o *ClassificationsItem) SetFromModules(v bool) {
	o.FromModules = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClassificationsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.DefID != nil {
		toSerialize["DefID"] = o.DefID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.MitreTags.IsSet() {
		toSerialize["MitreTags"] = o.MitreTags.Get()
	}
	if o.KillChainPhase.IsSet() {
		toSerialize["KillChainPhase"] = o.KillChainPhase.Get()
	}
	if o.FromMarket != nil {
		toSerialize["FromMarket"] = o.FromMarket
	}
	if o.FromModules != nil {
		toSerialize["FromModules"] = o.FromModules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ClassificationsItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID             *string                       `json:"ID,omitempty"`
		DefID          *int64                        `json:"DefID,omitempty"`
		Name           *string                       `json:"Name,omitempty"`
		Severity       *string                       `json:"Severity,omitempty"`
		MitreTags      common.NullableList[MitreTag] `json:"MitreTags,omitempty"`
		KillChainPhase common.NullableList[string]   `json:"KillChainPhase,omitempty"`
		FromMarket     *bool                         `json:"FromMarket,omitempty"`
		FromModules    *bool                         `json:"FromModules,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ID", "DefID", "Name", "Severity", "MitreTags", "KillChainPhase", "FromMarket", "FromModules"})
	} else {
		return err
	}

	o.ID = all.ID
	o.DefID = all.DefID
	o.Name = all.Name
	o.Severity = all.Severity
	o.MitreTags = all.MitreTags
	o.KillChainPhase = all.KillChainPhase
	o.FromMarket = all.FromMarket
	o.FromModules = all.FromModules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
