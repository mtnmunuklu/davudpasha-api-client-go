package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// Actions represents a collection of actions.
type Actions struct {
	// List of actions.
	Actions []Action `json:"Actions,omitempty"`
	// List of action reference IDs.
	ActionRefIDs common.NullableList[string] `json:"ActionRefIds,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActions creates a new Actions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActions() *Actions {
	this := Actions{}
	return &this
}

// NewActionsWithDefaults creates a new Actions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionsWithDefaults() *Actions {
	this := Actions{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *Actions) GetActions() []Action {
	if o == nil || o.Actions == nil {
		var ret []Action
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actions) GetActionsOk() (*[]Action, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *Actions) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given []Action and assigns it to the Actions field.
func (o *Actions) SetActions(v []Action) {
	o.Actions = v
}

// GetActionRefIDs returns a tuple with the ActionRefIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Actions) GetActionRefIDs() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionRefIDs.Get(), o.ActionRefIDs.IsSet()
}

// GetActionRefIDsOk returns a tuple with the ActionRefIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Actions) GetActionRefIDsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionRefIDs.Get(), o.ActionRefIDs.IsSet()
}

// HasActionRefIDs returns a boolean if a ActionRefIDs has been set.
func (o *Actions) HasActionRefIDs() bool {
	return o != nil && o.ActionRefIDs.IsSet()
}

// SetActionRefIDs gets a reference to the given common.Nullable[]string and assigns it to the ActionRefIDs field.
func (o *Actions) SetActionRefIDs(v []string) {
	o.ActionRefIDs.Set(&v)
}

// SetActionRefIDsNil sets the value for ActionRefIDs to be an explicit nil.
func (o *Actions) SetActionRefIDsNil() {
	o.ActionRefIDs.Set(nil)
}

// UnSetActionRefIDs ensures that no value is present for ActionRefIDs, not even an explicit nil.
func (o *Actions) UnsetActionRefIDs() {
	o.ActionRefIDs.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o Actions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.ActionRefIDs.IsSet() {
		toSerialize["ActionRefIds"] = o.ActionRefIDs.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *Actions) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions      []Action                    `json:"Actions,omitempty"`
		ActionRefIDs common.NullableList[string] `json:"ActionRefIds,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Actions", "ActionRefIds"})
	} else {
		return err
	}

	o.Actions = all.Actions
	o.ActionRefIDs = all.ActionRefIDs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
