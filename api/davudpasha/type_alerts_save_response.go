package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AlertsSaveResponse struct {
	Status         *bool                 `json:"Status,omitempty"`
	ValidationList []string              `json:"ValidationList,omitempty"`
	Data           common.NullableString `json:"Data,omitempty"`
	OptErrorMsg    common.NullableString `json:"optErrorMsg,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertsSaveResponse creates a new AlertsSaveResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertsSaveResponse() *AlertsSaveResponse {
	this := AlertsSaveResponse{}
	return &this
}

// NewAlertsSaveResponseWithDefaults creates a new AlertsSaveResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertsSaveResponseWithDefaults() *AlertsSaveResponse {
	this := AlertsSaveResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertsSaveResponse) GetStatus() bool {
	if o == nil || o.Status == nil {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsSaveResponse) GetStatusOk() (*bool, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertsSaveResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *AlertsSaveResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetValidationList returns the ValidationList field value if set, zero value otherwise.
func (o *AlertsSaveResponse) GetValidationList() []string {
	if o == nil || o.ValidationList == nil {
		var ret []string
		return ret
	}
	return o.ValidationList
}

// GetValidationListOk returns a tuple with the ValidationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsSaveResponse) GetValidationListOk() (*[]string, bool) {
	if o == nil || o.ValidationList == nil {
		return nil, false
	}
	return &o.ValidationList, true
}

// HasValidationList returns a boolean if a field has been set.
func (o *AlertsSaveResponse) HasValidationList() bool {
	return o != nil && o.ValidationList != nil
}

// SetValidationList gets a reference to the given []string and assigns it to the ValidationList field.
func (o *AlertsSaveResponse) SetValidationList(v []string) {
	o.ValidationList = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsSaveResponse) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsSaveResponse) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a Data has been set.
func (o *AlertsSaveResponse) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given common.NullableString and assigns it to the Data field.
func (o *AlertsSaveResponse) SetData(v string) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *AlertsSaveResponse) SetDataNil() {
	o.Data.Set(nil)
}

// UnSetData ensures that no value is present for Data, not even an explicit nil.
func (o *AlertsSaveResponse) UnSetData() {
	o.Data.UnSet()
}

// GetOptErrorMsg returns the OptErrorMsg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsSaveResponse) GetOptErrorMsg() string {
	if o == nil || o.OptErrorMsg.Get() == nil {
		var ret string
		return ret
	}
	return *o.OptErrorMsg.Get()
}

// GetOptErrorMsgOk returns a tuple with the OptErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsSaveResponse) GetOptErrorMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptErrorMsg.Get(), o.OptErrorMsg.IsSet()
}

// HasOptErrorMsg returns a boolean if a OptErrorMsg has been set.
func (o *AlertsSaveResponse) HasOptErrorMsg() bool {
	return o != nil && o.OptErrorMsg.IsSet()
}

// SetOptErrorMsg gets a reference to the given common.NullableString and assigns it to the OptErrorMsg field.
func (o *AlertsSaveResponse) SetOptErrorMsg(v string) {
	o.OptErrorMsg.Set(&v)
}

// SetOptErrorMsgNil sets the value for OptErrorMsg to be an explicit nil.
func (o *AlertsSaveResponse) SetOptErrorMsgNil() {
	o.OptErrorMsg.Set(nil)
}

// UnSetOptErrorMsg ensures that no value is present for OptErrorMsg, not even an explicit nil.
func (o *AlertsSaveResponse) UnSetOptErrorMsg() {
	o.OptErrorMsg.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertsSaveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.ValidationList != nil {
		toSerialize["ValidationList"] = o.ValidationList
	}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data.Get()
	}
	if o.OptErrorMsg.IsSet() {
		toSerialize["optErrorMsg"] = o.OptErrorMsg.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AlertsSaveResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status         *bool                 `json:"Status,omitempty"`
		ValidationList []string              `json:"ValidationList,omitempty"`
		Data           common.NullableString `json:"Data,omitempty"`
		OptErrorMsg    common.NullableString `json:"optErrorMsg,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Status", "ValidationList", "Data", "optErrorMsg"})
	} else {
		return err
	}

	o.Status = all.Status
	o.ValidationList = all.ValidationList
	o.Data = all.Data
	o.OptErrorMsg = all.OptErrorMsg

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
