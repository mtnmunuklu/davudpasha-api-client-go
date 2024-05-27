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
	ActionRefId common.NullableString `json:"ActionRefId,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
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

// SetData gets a reference to the given datadog.NullableString and assigns it to the Data field.
func (o *Action) SetData(v string) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *Action) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil.
func (o *Action) UnsetData() {
	o.Data.UnSet()
}

// GetActionRefId returns the ActionRefId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Action) GetActionRefId() string {
	if o == nil || o.ActionRefId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionRefId.Get()
}

// GetActionRefIdOk returns a tuple with the ActionRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Action) GetActionRefIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionRefId.Get(), o.ActionRefId.IsSet()
}

// HasActionRefId returns a boolean if a ActionRefId has been set.
func (o *Action) HasActionRefId() bool {
	return o != nil && o.ActionRefId.IsSet()
}

// SetActionRefId gets a reference to the given datadog.NullableString and assigns it to the ActionRefId field.
func (o *Action) SetActionRefId(v string) {
	o.ActionRefId.Set(&v)
}

// SetActionRefIdNil sets the value for ActionRefId to be an explicit nil.
func (o *Action) SetActionRefIdNil() {
	o.ActionRefId.Set(nil)
}

// UnsetActionRefId ensures that no value is present for ActionRefId, not even an explicit nil.
func (o *Action) UnsetActionRefId() {
	o.ActionRefId.UnSet()
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
	if o.ActionRefId.IsSet() {
		toSerialize["ActionRefId"] = o.ActionRefId.Get()
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
		ActionRefId      common.NullableString `json:"ActionRefId,omitempty"`
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
	o.ActionRefId = all.ActionRefId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
