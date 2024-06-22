package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type MitreTag struct {
	Tactic        *string  `json:"Tactic,omitempty"`
	Technique     []string `json:"Technique,omitempty"`
	SubTechniques []string `json:"SubTechniques,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewMitreTag creates a new MitreTag object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMitreTag() *MitreTag {
	this := MitreTag{}
	return &this
}

// NewMitreTagWithDefaults creates a new MitreTag object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMitreTagWithDefaults() *MitreTag {
	this := MitreTag{}
	return &this
}

// GetTechnique returns the Technique field value if set, zero value otherwise.
func (o *MitreTag) GetTechnique() []string {
	if o == nil || o.Technique == nil {
		var ret []string
		return ret
	}
	return o.Technique
}

// GetTechniqueOk returns a tuple with the Technique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MitreTag) GetTechniqueOk() (*[]string, bool) {
	if o == nil || o.Technique == nil {
		return nil, false
	}
	return &o.Technique, true
}

// HasTechnique returns a boolean if a field has been set.
func (o *MitreTag) HasTechnique() bool {
	return o != nil && o.Technique != nil
}

// SetTechnique gets a reference to the given []string and assigns it to the Technique field.
func (o *MitreTag) SetTechnique(v []string) {
	o.Technique = v
}

// GetSubTechniques returns the SubTechniques field value if set, zero value otherwise.
func (o *MitreTag) GetSubTechniques() []string {
	if o == nil || o.SubTechniques == nil {
		var ret []string
		return ret
	}
	return o.SubTechniques
}

// GetSubTechniquesOk returns a tuple with the SubTechniques field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MitreTag) GetSubTechniquesOk() (*[]string, bool) {
	if o == nil || o.SubTechniques == nil {
		return nil, false
	}
	return &o.SubTechniques, true
}

// HasSubTechniques returns a boolean if a field has been set.
func (o *MitreTag) HasSubTechniques() bool {
	return o != nil && o.SubTechniques != nil
}

// SetSubTechniques gets a reference to the given []string and assigns it to the SubTechniques field.
func (o *MitreTag) SetSubTechniques(v []string) {
	o.SubTechniques = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MitreTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Tactic != nil {
		toSerialize["Tactic"] = o.Tactic
	}
	if o.Technique != nil {
		toSerialize["Technique"] = o.Technique
	}
	if o.SubTechniques != nil {
		toSerialize["SubTechniques"] = o.SubTechniques
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *MitreTag) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tactic        *string  `json:"Tactic,omitempty"`
		Technique     []string `json:"Technique,omitempty"`
		SubTechniques []string `json:"SubTechniques,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Tactic", "Technique", "SubTechniques"})
	} else {
		return err
	}

	o.Tactic = all.Tactic
	o.Technique = all.Technique
	o.SubTechniques = all.SubTechniques

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
