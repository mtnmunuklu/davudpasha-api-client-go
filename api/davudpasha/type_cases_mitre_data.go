package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// CasesMitreData represents MITRE-related data associated with a case.
type CasesMitreData struct {
	// Date when the MITRE data was created.
	CreationDate *string `json:"CreationDate,omitempty"`
	// Kill chain phase associated with the MITRE data, which can be null.
	KillChainPhase common.NullableList[string] `json:"KillChainPhase,omitempty"`
	// List of MITRE tags associated with the case.
	MitreTags []MitreTag `json:"MitreTags,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewCasesMitreData creates a new CasesMitreData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCasesMitreData() *CasesMitreData {
	this := CasesMitreData{}
	return &this
}

// NewCasesMitreDataWithDefaults creates a new CasesMitreData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCasesMitreDataWithDefaults() *CasesMitreData {
	this := CasesMitreData{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *CasesMitreData) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesMitreData) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *CasesMitreData) HasCreationDate() bool {
	return o != nil && o.CreationDate != nil
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *CasesMitreData) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetKillChainPhase returns the KillChainPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesMitreData) GetKillChainPhase() []string {
	if o == nil || o.KillChainPhase.Get() == nil {
		var ret []string
		return ret
	}
	return *o.KillChainPhase.Get()
}

// GetKillChainPhaseOk returns a tuple with the KillChainPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesMitreData) GetKillChainPhaseOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KillChainPhase.Get(), o.KillChainPhase.IsSet()
}

// HasKillChainPhase returns a boolean if a KillChainPhase has been set.
func (o *CasesMitreData) HasKillChainPhase() bool {
	return o != nil && o.KillChainPhase.IsSet()
}

// SetKillChainPhase gets a reference to the given common.Nullable[]string and assigns it to the KillChainPhase field.
func (o *CasesMitreData) SetKillChainPhase(v []string) {
	o.KillChainPhase.Set(&v)
}

// SetKillChainPhaseNil sets the value for KillChainPhase to be an explicit nil.
func (o *CasesMitreData) SetKillChainPhaseNil() {
	o.KillChainPhase.Set(nil)
}

// UnSetKillChainPhase ensures that no value is present for KillChainPhase, not even an explicit nil.
func (o *CasesMitreData) UnSetKillChainPhase() {
	o.KillChainPhase.UnSet()
}

// GetMitreTags returns the MitreTags field value if set, zero value otherwise.
func (o *CasesMitreData) GetMitreTags() []MitreTag {
	if o == nil || o.MitreTags == nil {
		var ret []MitreTag
		return ret
	}
	return o.MitreTags
}

// GetMitreTagsOk returns a tuple with the MitreTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesMitreData) GetMitreTagsOk() (*[]MitreTag, bool) {
	if o == nil || o.MitreTags == nil {
		return nil, false
	}
	return &o.MitreTags, true
}

// HasMitreTags returns a boolean if a field has been set.
func (o *CasesMitreData) HasMitreTags() bool {
	return o != nil && o.MitreTags != nil
}

// SetMitreTags gets a reference to the given []MitreTag and assigns it to the MitreTags field.
func (o *CasesMitreData) SetMitreTags(v []MitreTag) {
	o.MitreTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CasesMitreData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CreationDate != nil {
		toSerialize["CreationDate"] = o.CreationDate
	}
	if o.KillChainPhase.IsSet() {
		toSerialize["KillChainPhase"] = o.KillChainPhase.Get()
	}
	if o.MitreTags != nil {
		toSerialize["MitreTags"] = o.MitreTags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *CasesMitreData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreationDate   *string                     `json:"CreationDate,omitempty"`
		KillChainPhase common.NullableList[string] `json:"KillChainPhase,omitempty"`
		MitreTags      []MitreTag                  `json:"MitreTags,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"CreationDate", "KillChainPhase", "MitreTags"})
	} else {
		return err
	}

	o.CreationDate = all.CreationDate
	o.KillChainPhase = all.KillChainPhase
	o.MitreTags = all.MitreTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableCasesMitreData handles when a null is used for CasesMitreData.
type NullableCasesMitreData struct {
	value *CasesMitreData
	isSet bool
}

// Get returns the associated value.
func (v NullableCasesMitreData) Get() *CasesMitreData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCasesMitreData) Set(val *CasesMitreData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCasesMitreData) IsSet() bool {
	return v.isSet
}

// UnSet sets the value to nil and resets the set flag/
func (v *NullableCasesMitreData) UnSet() {
	v.value = nil
	v.isSet = false
}

// NewNullableCasesMitreData initializes the struct as if Set has been called.
func NewNullableCasesMitreData(val *CasesMitreData) *NullableCasesMitreData {
	return &NullableCasesMitreData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCasesMitreData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCasesMitreData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
