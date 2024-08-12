package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ReportsFormRef struct {
	Current common.NullableString `json:"current,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsFormRef creates a new ReportsFormRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsFormRef() *ReportsFormRef {
	this := ReportsFormRef{}
	return &this
}

// NewReportsFormRefWithDefaults creates a new ReportsFormRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsFormRefWithDefaults() *ReportsFormRef {
	this := ReportsFormRef{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsFormRef) GetCurrent() string {
	if o == nil || o.Current.Get() == nil {
		var ret string
		return ret
	}
	return *o.Current.Get()
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsFormRef) GetCurrentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Current.Get(), o.Current.IsSet()
}

// HasCurrent returns a boolean if a field has been set.
func (o *ReportsFormRef) HasCurrent() bool {
	return o != nil && o.Current.IsSet()
}

// SetCurrent gets a reference to the given common.NullableString and assigns it to the Current field.
func (o *ReportsFormRef) SetCurrent(v string) {
	o.Current.Set(&v)
}

// SetCurrentNil sets the value for Current to be an explicit nil.
func (o *ReportsFormRef) SetCurrentNil() {
	o.Current.Set(nil)
}

// UnSetCurrent ensures that no value is present for Current, not even an explicit nil.
func (o *ReportsFormRef) UnSetCurrent() {
	o.Current.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsFormRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Current.IsSet() {
		toSerialize["current"] = o.Current.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsFormRef) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Current common.NullableString `json:"current,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"current"})
	} else {
		return err
	}

	o.Current = all.Current

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
