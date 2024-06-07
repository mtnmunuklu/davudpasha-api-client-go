package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type CasesMitreData struct {
	CreationDate   *string               `json:"CreationDate,omitempty"`
	KillChainPhase common.NullableString `json:"KillChainPhase,omitempty"`
	MitreTags      []MitreTag            `json:"MitreTags,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
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
