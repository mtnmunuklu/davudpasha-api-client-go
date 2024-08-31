package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SourcesFilterConditions struct {
	IsOr     *bool                 `json:"IsOr,omitempty"`
	Operator *string               `json:"Operator,omitempty"`
	Variable common.NullableString `json:"Variable,omitempty"`
	Valuable common.NullableString `json:"Valuable,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourcesFilterConditions creates a new SourcesFilterConditions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcesFilterConditions() *SourcesFilterConditions {
	this := SourcesFilterConditions{}
	return &this
}

// NewSourcesFilterConditionsWithDefaults creates a new SourcesFilterConditions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcesFilterConditionsWithDefaults() *SourcesFilterConditions {
	this := SourcesFilterConditions{}
	return &this
}

// GetIsOr returns the IsOr field value if set, zero value otherwise.
func (o *SourcesFilterConditions) GetIsOr() bool {
	if o == nil || o.IsOr == nil {
		var ret bool
		return ret
	}
	return *o.IsOr
}

// GetIsOrOk returns a tuple with the IsOr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesFilterConditions) GetIsOrOk() (*bool, bool) {
	if o == nil || o.IsOr == nil {
		return nil, false
	}
	return o.IsOr, true
}

// HasIsOr returns a boolean if a field has been set.
func (o *SourcesFilterConditions) HasIsOr() bool {
	return o != nil && o.IsOr != nil
}

// SetIsOr gets a reference to the given bool and assigns it to the IsOr field.
func (o *SourcesFilterConditions) SetIsOr(v bool) {
	o.IsOr = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *SourcesFilterConditions) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesFilterConditions) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *SourcesFilterConditions) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *SourcesFilterConditions) SetOperator(v string) {
	o.Operator = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFilterConditions) GetVariable() string {
	if o == nil || o.Variable.Get() == nil {
		var ret string
		return ret
	}
	return *o.Variable.Get()
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFilterConditions) GetVariableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variable.Get(), o.Variable.IsSet()
}

// HasVariable returns a boolean if a Variable has been set.
func (o *SourcesFilterConditions) HasVariable() bool {
	return o != nil && o.Variable.IsSet()
}

// SetVariable gets a reference to the given common.NullableString and assigns it to the Variable field.
func (o *SourcesFilterConditions) SetVariable(v string) {
	o.Variable.Set(&v)
}

// SetVariableNil sets the value for Variable to be an explicit nil.
func (o *SourcesFilterConditions) SetVariableNil() {
	o.Variable.Set(nil)
}

// UnSetVariable ensures that no value is present for Variable, not even an explicit nil.
func (o *SourcesFilterConditions) UnsetVariable() {
	o.Variable.Unset()
}

// GetValuable returns the Valuable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFilterConditions) GetValuable() string {
	if o == nil || o.Valuable.Get() == nil {
		var ret string
		return ret
	}
	return *o.Valuable.Get()
}

// GetValuableOk returns a tuple with the Valuable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFilterConditions) GetValuableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Valuable.Get(), o.Valuable.IsSet()
}

// HasValuable returns a boolean if a Valuable has been set.
func (o *SourcesFilterConditions) HasValuable() bool {
	return o != nil && o.Valuable.IsSet()
}

// SetValuable gets a reference to the given common.NullableString and assigns it to the Valuable field.
func (o *SourcesFilterConditions) SetValuable(v string) {
	o.Valuable.Set(&v)
}

// SetValuableNil sets the value for Valuable to be an explicit nil.
func (o *SourcesFilterConditions) SetValuableNil() {
	o.Valuable.Set(nil)
}

// UnSetValuable ensures that no value is present for Valuable, not even an explicit nil.
func (o *SourcesFilterConditions) UnsetValuable() {
	o.Valuable.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcesFilterConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.IsOr != nil {
		toSerialize["IsOr"] = o.IsOr
	}
	if o.Operator != nil {
		toSerialize["Operator"] = o.Operator
	}
	if o.Variable.IsSet() {
		toSerialize["Variable"] = o.Variable.Get()
	}
	if o.Valuable.IsSet() {
		toSerialize["Valuable"] = o.Valuable.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourcesFilterConditions) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsOr     *bool                 `json:"IsOr,omitempty"`
		Operator *string               `json:"Operator,omitempty"`
		Variable common.NullableString `json:"Variable,omitempty"`
		Valuable common.NullableString `json:"Valuable,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsOr == nil {
		return fmt.Errorf("requiered field IsOr is missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("requiered field Operator is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"IsOr", "Operator", "Variable", "Valuable"})
	} else {
		return err
	}

	o.IsOr = all.IsOr
	o.Operator = all.Operator
	o.Variable = all.Variable
	o.Valuable = all.Valuable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
