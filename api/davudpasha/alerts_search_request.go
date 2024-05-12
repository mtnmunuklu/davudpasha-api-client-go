package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AlertsSearchRequest struct {
	Filter *string `json:"filter,omitempty"`
	// SmartRestRequestContext is the context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewAlertsSearchRequest creates a new AlertsSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertsSearchRequest() *AlertsSearchRequest {
	this := AlertsSearchRequest{}
	return &this
}

// NewAlertsSearchRequestWithDefaults creates a new AlertsSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertsSearchRequestWithDefault() *AlertsSearchRequest {
	this := AlertsSearchRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AlertsSearchRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsSearchRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AlertsSearchRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *AlertsSearchRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *AlertsSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *AlertsSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *AlertsSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
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
func (o *AlertsSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter                  *string `json:"filter,omitempty"`
		SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filter == nil {
		return fmt.Errorf("requiered field filter is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filter", "smartRestRequestContext"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
