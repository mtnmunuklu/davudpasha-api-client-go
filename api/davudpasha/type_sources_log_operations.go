package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SourcesLogOperations struct {
	ModuleName *string          `json:"ModuleName,omitempty"`
	Priority   *string          `json:"Priority,omitempty"`
	Parameters *json.RawMessage `json:"Parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSourcesLogOperations creates a new SourcesLogOperations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcesLogOperations() *SourcesLogOperations {
	this := SourcesLogOperations{}
	return &this
}

// NewSourcesLogOperationsWithDefaults creates a new SourcesLogOperations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcesLogOperationsWithDefaults() *SourcesLogOperations {
	this := SourcesLogOperations{}
	return &this
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *SourcesLogOperations) GetModuleName() string {
	if o == nil || o.ModuleName == nil {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogOperations) GetModuleNameOk() (*string, bool) {
	if o == nil || o.ModuleName == nil {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *SourcesLogOperations) HasModuleName() bool {
	return o != nil && o.ModuleName != nil
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *SourcesLogOperations) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SourcesLogOperations) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogOperations) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SourcesLogOperations) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *SourcesLogOperations) SetPriority(v string) {
	o.Priority = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *SourcesLogOperations) GetParameters() json.RawMessage {
	if o == nil || o.Parameters == nil {
		var ret json.RawMessage
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogOperations) GetParametersOk() (*json.RawMessage, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *SourcesLogOperations) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given json.RawMessage and assigns it to the Parameters field.
func (o *SourcesLogOperations) SetParameters(v json.RawMessage) {
	o.Parameters = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcesLogOperations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}

	if o.ModuleName != nil {
		toSerialize["ModuleName"] = o.ModuleName
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourcesLogOperations) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModuleName *string          `json:"ModuleName,omitempty"`
		Priority   *string          `json:"Priority,omitempty"`
		Parameters *json.RawMessage `json:"Parameters,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ModuleName", "Priority", "Parameters"})
	} else {
		return err
	}

	o.ModuleName = all.ModuleName
	o.Priority = all.Priority
	o.Parameters = all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
