package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// AlertDefinitionsItem represents an item in alerts.
type AlertDefinitionsItem struct {
	// The name of the Log Source associated with the alert.
	LgsName common.NullableString `json:"LgsName,omitempty"`
	// Correlation data related to the alert definitions.
	CorrelationData *AlertDefinitionsCorrelationData `json:"CorrelationData,omitempty"`
	// Indicates whether action is required for this alert.
	ActionRequired *bool `json:"ActionRequired,omitempty"`
	// Message describing the action required for this alert.
	ActionMessage common.NullableString `json:"ActionMessage,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertDefinitionsItem creates a new AlertDefinitionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertDefinitionsItem() *AlertDefinitionsItem {
	this := AlertDefinitionsItem{}
	return &this
}

// NewAlertDefinitionsItemWithDefaults creates a new AlertDefinitionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertDefinitionsItemWithDefaults() *AlertDefinitionsItem {
	this := AlertDefinitionsItem{}
	return &this
}

// GetLgsName returns the LgsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertDefinitionsItem) GetLgsName() string {
	if o == nil || o.LgsName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LgsName.Get()
}

// GetLgsNameOk returns a tuple with the LgsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsItem) GetLgsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LgsName.Get(), o.LgsName.IsSet()
}

// HasLgsName returns a boolean if a field has been set.
func (o *AlertDefinitionsItem) HasLgsName() bool {
	return o != nil && o.LgsName.IsSet()
}

// SetLgsName gets a reference to the given common.NullableString and assigns it to the LgsName field.
func (o *AlertDefinitionsItem) SetLgsName(v string) {
	o.LgsName.Set(&v)
}

// SetLgsNameNil sets the value for LgsName to be an explicit nil.
func (o *AlertDefinitionsItem) SetLgsNameNil() {
	o.LgsName.Set(nil)
}

// UnSetLgsName ensures that no value is present for LgsName, not even an explicit nil.
func (o *AlertDefinitionsItem) UnsetLgsName() {
	o.LgsName.Unset()
}

// GetCorrelationData returns the CorrelationData field value if set, zero value otherwise.
func (o *AlertDefinitionsItem) GetCorrelationData() AlertDefinitionsCorrelationData {
	if o == nil || o.CorrelationData == nil {
		var ret AlertDefinitionsCorrelationData
		return ret
	}
	return *o.CorrelationData
}

// GetCorrelationDataOk returns a tuple with the CorrelationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsItem) GetCorrelationDataOk() (*AlertDefinitionsCorrelationData, bool) {
	if o == nil || o.CorrelationData == nil {
		return nil, false
	}
	return o.CorrelationData, true
}

// HasCorrelationData returns a boolean if a field has been set.
func (o *AlertDefinitionsItem) HasCorrelationData() bool {
	return o != nil && o.CorrelationData != nil
}

// SetCorrelationData gets a reference to the given CorrelationData and assigns it to the CorrelationData field.
func (o *AlertDefinitionsItem) SetCorrelationData(v AlertDefinitionsCorrelationData) {
	o.CorrelationData = &v
}

// GetActionRequired returns the ActionRequired field value if set, zero value otherwise.
func (o *AlertDefinitionsItem) GetActionRequired() bool {
	if o == nil || o.ActionRequired == nil {
		var ret bool
		return ret
	}
	return *o.ActionRequired
}

// GetActionRequiredOk returns a tuple with the ActionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsItem) GetActionRequiredOk() (*bool, bool) {
	if o == nil || o.ActionRequired == nil {
		return nil, false
	}
	return o.ActionRequired, true
}

// HasActionRequired returns a boolean if a field has been set.
func (o *AlertDefinitionsItem) HasActionRequired() bool {
	return o != nil && o.ActionRequired != nil
}

// SetActionRequired gets a reference to the given bool and assigns it to the ActionRequired field.
func (o *AlertDefinitionsItem) SetActionRequired(v bool) {
	o.ActionRequired = &v
}

// GetActionMessage returns the ActionMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertDefinitionsItem) GetActionMessage() string {
	if o == nil || o.ActionMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionMessage.Get()
}

// GetActionMessageOk returns a tuple with the ActionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertDefinitionsItem) GetActionMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionMessage.Get(), o.ActionMessage.IsSet()
}

// HasActionMessage returns a boolean if a field has been set.
func (o *AlertDefinitionsItem) HasActionMessage() bool {
	return o != nil && o.ActionMessage.IsSet()
}

// SetActionMessage gets a reference to the given common.NullableString and assigns it to the ActionMessage field.
func (o *AlertDefinitionsItem) SetActionMessage(v string) {
	o.ActionMessage.Set(&v)
}

// SetActionMessageNil sets the value for ActionMessage to be an explicit nil.
func (o *AlertDefinitionsItem) SetActionMessageNil() {
	o.ActionMessage.Set(nil)
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertDefinitionsItem) MarshalJSON() ([]byte, error) {
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
func (o *AlertDefinitionsItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		LgsName         common.NullableString            `json:"LgsName,omitempty"`
		CorrelationData *AlertDefinitionsCorrelationData `json:"CorrelationData,omitempty"`
		ActionRequired  *bool                            `json:"ActionRequired,omitempty"`
		ActionMessage   common.NullableString            `json:"ActionMessage,omitempty"`
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
