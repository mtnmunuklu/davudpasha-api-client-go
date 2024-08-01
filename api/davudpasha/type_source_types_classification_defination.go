package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourceTypesClassificationDefination represents the classification definition structure for source types.
type SourceTypesClassificationDefination struct {
	// ID represents the identifier of the classification definition.
	ID *string `json:"ID,omitempty"`
	// DefID specifies the definition ID associated with the classification.
	DefID *int64 `json:"DefID,omitempty"`
	// Name specifies the name of the classification definition.
	Name *string `json:"Name,omitempty"`
	// Severity indicates the severity level associated with the classification.
	Severity *string `json:"Severity,omitempty"`
	// MitreTags contains a list of MITRE tags associated with the classification.
	MitreTags common.NullableList[MitreTag] `json:"MitreTags,omitempty"`
	// KillChainPhase specifies the kill chain phase associated with the classification.
	KillChainPhase common.NullableString `json:"KillChainPhase,omitempty"`
	// MitreCreationDate indicates the MITRE creation date of the classification.
	MitreCreationDate common.NullableString `json:"MitreCreationDate,omitempty"`
	// FromMarket specifies if the classification is from the market.
	FromMarket *bool `json:"FromMarket,omitempty"`
	// FromModules specifies if the classification is from modules.
	FromModules *bool `json:"FromModules,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourceTypesClassificationDefination creates a new SourceTypesClassificationDefination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesClassificationDefination() *SourceTypesClassificationDefination {
	this := SourceTypesClassificationDefination{}
	return &this
}

// NewSourceTypesClassificationDefinationWithDefaults creates a new SourceTypesClassificationDefination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourceTypesClassificationDefinationWithDefaults() *SourceTypesClassificationDefination {
	this := SourceTypesClassificationDefination{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *SourceTypesClassificationDefination) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesClassificationDefination) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *SourceTypesClassificationDefination) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *SourceTypesClassificationDefination) SetID(v string) {
	o.ID = &v
}

// GetDefID returns the DefID field value if set, zero value otherwise.
func (o *SourceTypesClassificationDefination) GetDefID() int64 {
	if o == nil || o.DefID == nil {
		var ret int64
		return ret
	}
	return *o.DefID
}

// GetDefIDOk returns a tuple with the DefID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesClassificationDefination) GetDefIDOk() (*int64, bool) {
	if o == nil || o.DefID == nil {
		return nil, false
	}
	return o.DefID, true
}

// HasDefID returns a boolean if a field has been set.
func (o *SourceTypesClassificationDefination) HasDefID() bool {
	return o != nil && o.DefID != nil
}

// SetDefID gets a reference to the given string and assigns it to the DefID field.
func (o *SourceTypesClassificationDefination) SetDefID(v int64) {
	o.DefID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceTypesClassificationDefination) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesClassificationDefination) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceTypesClassificationDefination) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceTypesClassificationDefination) SetName(v string) {
	o.Name = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *SourceTypesClassificationDefination) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesClassificationDefination) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *SourceTypesClassificationDefination) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *SourceTypesClassificationDefination) SetSeverity(v string) {
	o.Severity = &v
}

// GetMitreTags returns a tuple with the MitreTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesClassificationDefination) GetMitreTags() (*[]MitreTag, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitreTags.Get(), o.MitreTags.IsSet()
}

// GetMitreTagsOk returns a tuple with the MitreTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesClassificationDefination) GetMitreTagsOk() (*[]MitreTag, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitreTags.Get(), o.MitreTags.IsSet()
}

// HasMitreTags returns a boolean if a MitreTags has been set.
func (o *SourceTypesClassificationDefination) HasMitreTags() bool {
	return o != nil && o.MitreTags.IsSet()
}

// SetMitreTags gets a reference to the given common.Nullable[]MitreTag and assigns it to the MitreTags field.
func (o *SourceTypesClassificationDefination) SetMitreTags(v []MitreTag) {
	o.MitreTags.Set(&v)
}

// SetMitreTagsNil sets the value for MitreTags to be an explicit nil.
func (o *SourceTypesClassificationDefination) SetMitreTagsNil() {
	o.MitreTags.Set(nil)
}

// UnSetMitreTags ensures that no value is present for MitreTags, not even an explicit nil.
func (o *SourceTypesClassificationDefination) UnSetMitreTags() {
	o.MitreTags.UnSet()
}

// GetKillChainPhase returns the KillChainPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesClassificationDefination) GetKillChainPhase() string {
	if o == nil || o.KillChainPhase.Get() == nil {
		var ret string
		return ret
	}
	return *o.KillChainPhase.Get()
}

// GetKillChainPhaseOk returns a tuple with the KillChainPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesClassificationDefination) GetKillChainPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KillChainPhase.Get(), o.KillChainPhase.IsSet()
}

// HasKillChainPhase returns a boolean if a field has been set.
func (o *SourceTypesClassificationDefination) HasKillChainPhase() bool {
	return o != nil && o.KillChainPhase.IsSet()
}

// SetKillChainPhase gets a reference to the given common.NullableString and assigns it to the KillChainPhase field.
func (o *SourceTypesClassificationDefination) SetKillChainPhase(v string) {
	o.KillChainPhase.Set(&v)
}

// SetKillChainPhaseNil sets the value for KillChainPhase to be an explicit nil.
func (o *SourceTypesClassificationDefination) SetKillChainPhaseNil() {
	o.KillChainPhase.Set(nil)
}

// UnSetKillChainPhase ensures that no value is present for KillChainPhase, not even an explicit nil.
func (o *SourceTypesClassificationDefination) UnSetKillChainPhase() {
	o.KillChainPhase.UnSet()
}

// GetMitreCreationDate returns the MitreCreationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesClassificationDefination) GetMitreCreationDate() string {
	if o == nil || o.MitreCreationDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.MitreCreationDate.Get()
}

// GetMitreCreationDateOk returns a tuple with the MitreCreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesClassificationDefination) GetMitreCreationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitreCreationDate.Get(), o.MitreCreationDate.IsSet()
}

// HasMitreCreationDate returns a boolean if a field has been set.
func (o *SourceTypesClassificationDefination) HasMitreCreationDate() bool {
	return o != nil && o.MitreCreationDate.IsSet()
}

// SetMitreCreationDate gets a reference to the given common.NullableString and assigns it to the MitreCreationDate field.
func (o *SourceTypesClassificationDefination) SetMitreCreationDate(v string) {
	o.MitreCreationDate.Set(&v)
}

// SetMitreCreationDateNil sets the value for MitreCreationDate to be an explicit nil.
func (o *SourceTypesClassificationDefination) SetMitreCreationDateNil() {
	o.MitreCreationDate.Set(nil)
}

// UnSetMitreCreationDate ensures that no value is present for MitreCreationDate, not even an explicit nil.
func (o *SourceTypesClassificationDefination) UnSetMitreCreationDate() {
	o.MitreCreationDate.UnSet()
}

// GetFromMarket returns the FromMarket field value if set, zero value otherwise.
func (o *SourceTypesClassificationDefination) GetFromMarket() bool {
	if o == nil || o.FromMarket == nil {
		var ret bool
		return ret
	}
	return *o.FromMarket
}

// GetFromMarketOk returns a tuple with the FromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesClassificationDefination) GetFromMarketOk() (*bool, bool) {
	if o == nil || o.FromMarket == nil {
		return nil, false
	}
	return o.FromMarket, true
}

// HasFromMarket returns a boolean if a field has been set.
func (o *SourceTypesClassificationDefination) HasFromMarket() bool {
	return o != nil && o.FromMarket != nil
}

// SetFromMarket gets a reference to the given bool and assigns it to the FromMarket field.
func (o *SourceTypesClassificationDefination) SetFromMarket(v bool) {
	o.FromMarket = &v
}

// GetFromModules returns the FromModules field value if set, zero value otherwise.
func (o *SourceTypesClassificationDefination) GetFromModules() bool {
	if o == nil || o.FromModules == nil {
		var ret bool
		return ret
	}
	return *o.FromModules
}

// GetFromModulesOk returns a tuple with the FromModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesClassificationDefination) GetFromModulesOk() (*bool, bool) {
	if o == nil || o.FromModules == nil {
		return nil, false
	}
	return o.FromModules, true
}

// HasFromModules returns a boolean if a field has been set.
func (o *SourceTypesClassificationDefination) HasFromModules() bool {
	return o != nil && o.FromModules != nil
}

// SetFromModules gets a reference to the given bool and assigns it to the FromModules field.
func (o *SourceTypesClassificationDefination) SetFromModules(v bool) {
	o.FromModules = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesClassificationDefination) MarshalJSON() ([]byte, error) {
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
	if o.MitreCreationDate.IsSet() {
		toSerialize["MitreCreationDate"] = o.MitreCreationDate.Get()
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
func (o *SourceTypesClassificationDefination) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID                *string                       `json:"ID,omitempty"`
		DefID             *int64                        `json:"DefID,omitempty"`
		Name              *string                       `json:"Name,omitempty"`
		Severity          *string                       `json:"Severity,omitempty"`
		MitreTags         common.NullableList[MitreTag] `json:"MitreTags,omitempty"`
		KillChainPhase    common.NullableString         `json:"KillChainPhase,omitempty"`
		MitreCreationDate common.NullableString         `json:"MitreCreationDate,omitempty"`
		FromMarket        *bool                         `json:"FromMarket,omitempty"`
		FromModules       *bool                         `json:"FromModules,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.ID == nil {
		return fmt.Errorf("requiered field ID is missing")
	}
	if all.DefID == nil {
		return fmt.Errorf("requiered field DefID is missing")
	}
	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ID", "DefID", "Name", "Severity", "MitreTags", "KillChainPhase", "MitreCreationDate", "FromMarket", "FromModules"})
	} else {
		return err
	}

	o.ID = all.ID
	o.DefID = all.DefID
	o.Name = all.Name
	o.Severity = all.Severity
	o.MitreTags = all.MitreTags
	o.KillChainPhase = all.KillChainPhase
	o.MitreCreationDate = all.MitreCreationDate
	o.FromMarket = all.FromMarket
	o.FromModules = all.FromModules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
