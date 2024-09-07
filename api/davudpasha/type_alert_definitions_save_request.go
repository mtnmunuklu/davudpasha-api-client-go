package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// AlertDefinitionsSaveRequest represents the request structure for saving alert definitions.
type AlertDefinitionsSaveRequest struct {
	// Correlation data for the alert definitions.
	Correlation *AlertDefinitionsCorrelationData `json:"correlation,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertDefinitionsSaveRequest creates a new AlertDefinitionsSaveRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertDefinitionsSaveRequest() *AlertDefinitionsSaveRequest {
	this := AlertDefinitionsSaveRequest{}
	return &this
}

// NewAlertDefinitionsSaveRequestWithDefaults creates a new AlertDefinitionsSaveRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertDefinitionsSaveRequestWithDefaults() *AlertDefinitionsSaveRequest {
	this := AlertDefinitionsSaveRequest{}
	return &this
}

// GetCorrelation returns the Correlation field value if set, zero value otherwise.
func (o *AlertDefinitionsSaveRequest) GetCorrelation() AlertDefinitionsCorrelationData {
	if o == nil || o.Correlation == nil {
		var ret AlertDefinitionsCorrelationData
		return ret
	}
	return *o.Correlation
}

// GetCorrelationOk returns a tuple with the Correlation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsSaveRequest) GetCorrelationOk() (*AlertDefinitionsCorrelationData, bool) {
	if o == nil || o.Correlation == nil {
		return nil, false
	}
	return o.Correlation, true
}

// HasCorrelation returns a boolean if a field has been set.
func (o *AlertDefinitionsSaveRequest) HasCorrelation() bool {
	return o != nil && o.Correlation != nil
}

// SetCorrelation gets a reference to the given AlertDefinitionsCorrelationData and assigns it to the Correlation field.
func (o *AlertDefinitionsSaveRequest) SetCorrelation(v AlertDefinitionsCorrelationData) {
	o.Correlation = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *AlertDefinitionsSaveRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionsSaveRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *AlertDefinitionsSaveRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *AlertDefinitionsSaveRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertDefinitionsSaveRequest) MarshalJSON() ([]byte, error) {
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
func (o *AlertDefinitionsSaveRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Correlation             *AlertDefinitionsCorrelationData `json:"correlation,omitempty"`
		SmartRestRequestContext *string                          `json:"smartRestRequestContext,omitempty"`
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
