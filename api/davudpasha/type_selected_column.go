package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SelectedColumn represents a selected column in a query.
type SelectedColumn struct {
	// Value of the column.
	Value *string `json:"Value,omitempty"`
	// Text to display for the column.
	DisplayText *string `json:"DisplayText,omitempty"`
	// Lookup value for the column.
	Lookup common.NullableString `json:"Lookup,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectedColumn creates a new SelectedColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectedColumn() *SelectedColumn {
	this := SelectedColumn{}
	return &this
}

// NewSelectedColumnWithDefaults creates a new SelectedColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectedColumnWithDefaults() *SelectedColumn {
	this := SelectedColumn{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SelectedColumn) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedColumn) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SelectedColumn) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SelectedColumn) SetValue(v string) {
	o.Value = &v
}

// GetDisplayText returns the DisplayText field DisplayText if set, zero DisplayText otherwise.
func (o *SelectedColumn) GetDisplayText() string {
	if o == nil || o.DisplayText == nil {
		var ret string
		return ret
	}
	return *o.DisplayText
}

// GetDisplayTextOk returns a tuple with the DisplayText field DisplayText if set, nil otherwise
// and a boolean to check if the DisplayText has been set.
func (o *SelectedColumn) GetDisplayTextOk() (*string, bool) {
	if o == nil || o.DisplayText == nil {
		return nil, false
	}
	return o.DisplayText, true
}

// HasDisplayText returns a boolean if a field has been set.
func (o *SelectedColumn) HasDisplayText() bool {
	return o != nil && o.DisplayText != nil
}

// SetDisplayText gets a reference to the given string and assigns it to the DisplayText field.
func (o *SelectedColumn) SetDisplayText(v string) {
	o.DisplayText = &v
}

// GetLookup returns the Lookup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectedColumn) GetLookup() string {
	if o == nil || o.Lookup.Get() == nil {
		var ret string
		return ret
	}
	return *o.Lookup.Get()
}

// GetLookupOk returns a tuple with the Lookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SelectedColumn) GetLookupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lookup.Get(), o.Lookup.IsSet()
}

// HasLookup returns a boolean if a field has been set.
func (o *SelectedColumn) HasLookup() bool {
	return o != nil && o.Lookup.IsSet()
}

// SetLookup gets a reference to the given common.NullableString and assigns it to the Lookup field.
func (o *SelectedColumn) SetLookup(v string) {
	o.Lookup.Set(&v)
}

// SetLookupNil sets the value for Lookup to be an explicit nil.
func (o *SelectedColumn) SetLookupNil() {
	o.Lookup.Set(nil)
}

// UnSetLookup ensures that no value is present for Lookup, not even an explicit nil.
func (o *SelectedColumn) UnSetLookup() {
	o.Lookup.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o SelectedColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}

	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	if o.DisplayText != nil {
		toSerialize["DisplayText"] = o.DisplayText
	}
	if o.Lookup.IsSet() {
		toSerialize["Lookup"] = o.Lookup.Get()
	}
	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SelectedColumn) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value       *string               `json:"Value,omitempty"`
		DisplayText *string               `json:"DisplayText,omitempty"`
		Lookup      common.NullableString `json:"Lookup,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.Value == nil {
		return fmt.Errorf("required field Value missing")
	}
	if all.DisplayText == nil {
		return fmt.Errorf("required DisplayText Value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Value", "DisplayText", "Lookup"})
	} else {
		return err
	}

	o.Value = all.Value
	o.DisplayText = all.DisplayText
	o.Lookup = all.Lookup

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
