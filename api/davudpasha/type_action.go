package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// Action represents an action.
type Action struct {
	// Type of the action.
	ActionType *string `json:"ActionType,omitempty"`
	// Parameters of the action.
	ActionParameters []ActionParameter `json:"ActionParameters,omitempty"`
	// Data associated with the action.
	Data common.NullableString `json:"Data,omitempty"`
	// Reference ID for the action.
	ActionRefID common.NullableString `json:"ActionRefId,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAction creates a new Action object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAction() *Action {
	this := Action{}
	return &this
}

// NewActionWithDefaults creates a new Action object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionWithDefaults() *Action {
	this := Action{}
	return &this
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *Action) GetActionType() string {
	if o == nil || o.ActionType == nil {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetActionTypeOk() (*string, bool) {
	if o == nil || o.ActionType == nil {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *Action) HasActionType() bool {
	return o != nil && o.ActionType != nil
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *Action) SetActionType(v string) {
	o.ActionType = &v
}

// GetActionParameters returns the ActionParameters field value if set, zero value otherwise.
func (o *Action) GetActionParameters() []ActionParameter {
	if o == nil || o.ActionParameters == nil {
		var ret []ActionParameter
		return ret
	}
	return o.ActionParameters
}

// GetActionParametersOk returns a tuple with the ActionParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetActionParametersOk() (*[]ActionParameter, bool) {
	if o == nil || o.ActionParameters == nil {
		return nil, false
	}
	return &o.ActionParameters, true
}

// HasActionParameters returns a boolean if a field has been set.
func (o *Action) HasActionParameters() bool {
	return o != nil && o.ActionParameters != nil
}

// SetActionParameters gets a reference to the given []ActionParameter and assigns it to the ActionParameters field.
func (o *Action) SetActionParameters(v []ActionParameter) {
	o.ActionParameters = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Action) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Action) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a Data has been set.
func (o *Action) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given common.NullableString and assigns it to the Data field.
func (o *Action) SetData(v string) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *Action) SetDataNil() {
	o.Data.Set(nil)
}

// UnSetData ensures that no value is present for Data, not even an explicit nil.
func (o *Action) UnsetData() {
	o.Data.Unset()
}

// GetActionRefID returns the ActionRefID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Action) GetActionRefID() string {
	if o == nil || o.ActionRefID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionRefID.Get()
}

// GetActionRefIDOk returns a tuple with the ActionRefID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Action) GetActionRefIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionRefID.Get(), o.ActionRefID.IsSet()
}

// HasActionRefID returns a boolean if a ActionRefID has been set.
func (o *Action) HasActionRefID() bool {
	return o != nil && o.ActionRefID.IsSet()
}

// SetActionRefID gets a reference to the given common.NullableString and assigns it to the ActionRefID field.
func (o *Action) SetActionRefID(v string) {
	o.ActionRefID.Set(&v)
}

// SetActionRefIDNil sets the value for ActionRefID to be an explicit nil.
func (o *Action) SetActionRefIDNil() {
	o.ActionRefID.Set(nil)
}

// UnSetActionRefID ensures that no value is present for ActionRefID, not even an explicit nil.
func (o *Action) UnsetActionRefID() {
	o.ActionRefID.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ActionType != nil {
		toSerialize["ActionType"] = o.ActionType
	}
	if o.ActionParameters != nil {
		toSerialize["ActionParameters"] = o.ActionParameters
	}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data.Get()
	}
	if o.ActionRefID.IsSet() {
		toSerialize["ActionRefId"] = o.ActionRefID.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *Action) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionType       *string               `json:"ActionType,omitempty"`
		ActionParameters []ActionParameter     `json:"ActionParameters,omitempty"`
		Data             common.NullableString `json:"Data,omitempty"`
		ActionRefID      common.NullableString `json:"ActionRefId,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ActionType", "ActionParameters", "Data", "ActionRefId"})
	} else {
		return err
	}

	o.ActionType = all.ActionType
	o.ActionParameters = all.ActionParameters
	o.Data = all.Data
	o.ActionRefID = all.ActionRefID

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
