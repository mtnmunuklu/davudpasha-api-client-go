package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type Actions struct {
	Actions      []Action                    `json:"Actions,omitempty"`
	ActionRefIds common.NullableList[string] `json:"ActionRefIds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
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

// GetActionRefIds returns a tuple with the ActionRefIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Actions) GetActionRefIds() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionRefIds.Get(), o.ActionRefIds.IsSet()
}

// GetActionRefIdsOk returns a tuple with the ActionRefIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Actions) GetActionRefIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionRefIds.Get(), o.ActionRefIds.IsSet()
}

// HasActionRefIds returns a boolean if a ActionRefIds has been set.
func (o *Actions) HasActionRefIds() bool {
	return o != nil && o.ActionRefIds.IsSet()
}

// SetActionRefIds gets a reference to the given datadog.Nullable[]string and assigns it to the ActionRefIds field.
func (o *Actions) SetActionRefIds(v []string) {
	o.ActionRefIds.Set(&v)
}

// SetActionRefIdsNil sets the value for ActionRefIds to be an explicit nil.
func (o *Actions) SetActionRefIdsNil() {
	o.ActionRefIds.Set(nil)
}

// UnsetActionRefIds ensures that no value is present for ActionRefIds, not even an explicit nil.
func (o *Actions) UnSetActionRefIds() {
	o.ActionRefIds.UnSet()
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
	if o.ActionRefIds.IsSet() {
		toSerialize["ActionRefIds"] = o.ActionRefIds.Get()
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
		ActionRefIds common.NullableList[string] `json:"ActionRefIds,omitempty"`
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
	o.ActionRefIds = all.ActionRefIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
