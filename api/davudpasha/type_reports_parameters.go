package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ReportsParameters struct {
	Parameters []string         `json:"Parameters,omitempty"`
	Datas      *json.RawMessage `json:"Datas,omitempty"`
	IsActive   *bool            `json:"IsActive,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewReportsParameters creates a new ReportsParameters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsParameters() *ReportsParameters {
	this := ReportsParameters{}
	return &this
}

// NewReportsParametersWithDefaults creates a new ReportsParameters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsParametersWithDefaults() *ReportsParameters {
	this := ReportsParameters{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ReportsParameters) GetParameters() []string {
	if o == nil || o.Parameters == nil {
		var ret []string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsParameters) GetParametersOk() (*[]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ReportsParameters) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []string and assigns it to the Parameters field.
func (o *ReportsParameters) SetParameters(v []string) {
	o.Parameters = v
}

// GetDatas returns the Datas field value if set, zero value otherwise.
func (o *ReportsParameters) GetDatas() json.RawMessage {
	if o == nil || o.Datas == nil {
		var ret json.RawMessage
		return ret
	}
	return *o.Datas
}

// GetDatasOk returns a tuple with the Datas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsParameters) GetDatasOk() (*json.RawMessage, bool) {
	if o == nil || o.Datas == nil {
		return nil, false
	}
	return o.Datas, true
}

// HasDatas returns a boolean if a field has been set.
func (o *ReportsParameters) HasDatas() bool {
	return o != nil && o.Datas != nil
}

// SetDatas gets a reference to the given json.RawMessage and assigns it to the Datas field.
func (o *ReportsParameters) SetDatas(v json.RawMessage) {
	o.Datas = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ReportsParameters) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsParameters) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ReportsParameters) HasIsActive() bool {
	return o != nil && o.IsActive != nil
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ReportsParameters) SetIsActive(v bool) {
	o.IsActive = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}
	if o.Datas != nil {
		toSerialize["Datas"] = o.Datas
	}
	if o.IsActive != nil {
		toSerialize["IsActive"] = o.IsActive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsParameters) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Parameters []string         `json:"Parameters,omitempty"`
		Datas      *json.RawMessage `json:"Datas,omitempty"`
		IsActive   *bool            `json:"IsActive,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Parameters", "Datas", "IsActive"})
	} else {
		return err
	}

	o.Parameters = all.Parameters
	o.Datas = all.Datas
	o.IsActive = all.IsActive

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
