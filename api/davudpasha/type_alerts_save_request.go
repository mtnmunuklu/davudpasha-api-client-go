package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AlertsSaveRequest struct {
	Correlation *AlertsCorrelationData `json:"correlation,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertsSaveRequest creates a new AlertsSaveRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertsSaveRequest() *AlertsSaveRequest {
	this := AlertsSaveRequest{}
	return &this
}

// NewAlertsSaveRequestWithDefaults creates a new AlertsSaveRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertsSaveRequestWithDefaults() *AlertsSaveRequest {
	this := AlertsSaveRequest{}
	return &this
}

// GetCorrelation returns the Correlation field value if set, zero value otherwise.
func (o *AlertsSaveRequest) GetCorrelation() AlertsCorrelationData {
	if o == nil || o.Correlation == nil {
		var ret AlertsCorrelationData
		return ret
	}
	return *o.Correlation
}

// GetCorrelationOk returns a tuple with the Correlation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsSaveRequest) GetCorrelationOk() (*AlertsCorrelationData, bool) {
	if o == nil || o.Correlation == nil {
		return nil, false
	}
	return o.Correlation, true
}

// HasCorrelation returns a boolean if a field has been set.
func (o *AlertsSaveRequest) HasCorrelation() bool {
	return o != nil && o.Correlation != nil
}

// SetCorrelation gets a reference to the given AlertsCorrelationData and assigns it to the Correlation field.
func (o *AlertsSaveRequest) SetCorrelation(v AlertsCorrelationData) {
	o.Correlation = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *AlertsSaveRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsSaveRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *AlertsSaveRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *AlertsSaveRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertsSaveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Correlation != nil {
		toSerialize["correlation"] = o.Correlation
	}
	if o.SmartRestRequestContext != nil {
		toSerialize["smartRestRequestContext"] = o.SmartRestRequestContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AlertsSaveRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Correlation             *AlertsCorrelationData `json:"correlation,omitempty"`
		SmartRestRequestContext *string                `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Correlation == nil {
		return fmt.Errorf("requiered field correlation is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"correlation", "smartRestRequestContext"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Correlation != nil && all.Correlation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Correlation = all.Correlation
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
