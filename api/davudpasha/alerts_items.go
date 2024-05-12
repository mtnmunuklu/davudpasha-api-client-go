package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AlertsItems struct {
	LgsName         common.NullableString  `json:"LgsName,omitempty"`
	CorrelationData *AlertsCorrelationData `json:"CorrelationData,omitempty"`
	ActionRequired  *bool                  `json:"ActionRequired,omitempty"`
	ActionMessage   common.NullableString  `json:"ActionMessage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewAlertsItems creates a new AlertsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertsItems() *AlertsItems {
	this := AlertsItems{}
	return &this
}

// NewAlertsItemsWithDefault creates a new AlertsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertsItemsWithDefaults() *AlertsItems {
	this := AlertsItems{}
	return &this
}

// GetLgsNameOk returns a tuple with the LgsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsItems) GetLgsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LgsName.Get(), o.LgsName.IsSet()
}

// HasLgsName returns a boolean if a field has been set.
func (o *AlertsItems) HasLgsName() bool {
	return o != nil && o.LgsName.IsSet()
}

// SetLgsName gets a reference to the given datadog.NullableString and assigns it to the LgsName field.
func (o *AlertsItems) SetLgsName(v string) {
	o.LgsName.Set(&v)
}

// SetLgsNameNil sets the value for LgsName to be an explicit nil.
func (o *AlertsItems) SetLgsNameNil() {
	o.LgsName.Set(nil)
}

// UnsetLgsName ensures that no value is present for LgsName, not even an explicit nil.
func (o *AlertsItems) UnSetLgsName() {
	o.LgsName.UnSet()
}

// GetCorrelationData returns the CorrelationData field value if set, zero value otherwise.
func (o *AlertsItems) GetCorrelationData() AlertsCorrelationData {
	if o == nil || o.CorrelationData == nil {
		var ret AlertsCorrelationData
		return ret
	}
	return *o.CorrelationData
}

// GetCorrelationDataOk returns a tuple with the CorrelationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsItems) GetCorrelationDataOk() (*AlertsCorrelationData, bool) {
	if o == nil || o.CorrelationData == nil {
		return nil, false
	}
	return o.CorrelationData, true
}

// HasCorrelationData returns a boolean if a field has been set.
func (o *AlertsItems) HasCorrelationData() bool {
	return o != nil && o.CorrelationData != nil
}

// SetCorrelationData gets a reference to the given CorrelationData and assigns it to the CorrelationData field.
func (o *AlertsItems) SetCorrelationData(v AlertsCorrelationData) {
	o.CorrelationData = &v
}

// GetActionRequired returns the ActionRequired field value if set, zero value otherwise.
func (o *AlertsItems) GetActionRequired() bool {
	if o == nil || o.ActionRequired == nil {
		var ret bool
		return ret
	}
	return *o.ActionRequired
}

// GetActionRequiredOk returns a tuple with the ActionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsItems) GetActionRequiredOk() (*bool, bool) {
	if o == nil || o.ActionRequired == nil {
		return nil, false
	}
	return o.ActionRequired, true
}

// HasActionRequired returns a boolean if a field has been set.
func (o *AlertsItems) HasActionRequired() bool {
	return o != nil && o.ActionRequired != nil
}

// SetActionRequired gets a reference to the given bool and assigns it to the ActionRequired field.
func (o *AlertsItems) SetActionRequired(v bool) {
	o.ActionRequired = &v
}

// GetActionMessageOk returns a tuple with the ActionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsItems) GetActionMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionMessage.Get(), o.ActionMessage.IsSet()
}

// HasActionMessage returns a boolean if a field has been set.
func (o *AlertsItems) HasActionMessage() bool {
	return o != nil && o.ActionMessage.IsSet()
}

// SetActionMessage gets a reference to the given datadog.NullableString and assigns it to the ActionMessage field.
func (o *AlertsItems) SetActionMessage(v string) {
	o.ActionMessage.Set(&v)
}

// SetActionMessageNil sets the value for ActionMessage to be an explicit nil.
func (o *AlertsItems) SetActionMessageNil() {
	o.ActionMessage.Set(nil)
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.LgsName.IsSet() {
		toSerialize["LgsName"] = o.LgsName.Get()
	}
	if o.CorrelationData != nil {
		toSerialize["CorrelationData"] = o.CorrelationData
	}
	if o.ActionRequired != nil {
		toSerialize["ActionRequired"] = o.ActionRequired
	}
	if o.ActionMessage.IsSet() {
		toSerialize["ActionMessage"] = o.ActionMessage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AlertsItems) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		LgsName         common.NullableString  `json:"LgsName,omitempty"`
		CorrelationData *AlertsCorrelationData `json:"CorrelationData,omitempty"`
		ActionRequired  *bool                  `json:"ActionRequired,omitempty"`
		ActionMessage   common.NullableString  `json:"ActionMessage,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"LgsName", "CorrelationData", "ActionRequired", "ActionMessage"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CorrelationData != nil && all.CorrelationData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CorrelationData = all.CorrelationData
	o.LgsName = all.LgsName
	o.ActionRequired = all.ActionRequired
	o.ActionMessage = all.ActionMessage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
