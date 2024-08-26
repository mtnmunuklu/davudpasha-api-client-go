package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SourcesSaveRequest struct {
	Lgs           *SourcesItem `json:"lgs,omitempty"`
	DashboardName *string      `json:"dashboardName,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourcesSaveRequest creates a new SourcesSaveRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcesSaveRequest() *SourcesSaveRequest {
	this := SourcesSaveRequest{}
	return &this
}

// NewSourcesSaveRequestWithDefaults creates a new SourcesSaveRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcesSaveRequestWithDefaults() *SourcesSaveRequest {
	this := SourcesSaveRequest{}
	return &this
}

// GetLgs returns the Lgs field value if set, zero value otherwise.
func (o *SourcesSaveRequest) GetLgs() SourcesItem {
	if o == nil || o.Lgs == nil {
		var ret SourcesItem
		return ret
	}
	return *o.Lgs
}

// GetLgsOk returns a tuple with the Lgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesSaveRequest) GetLgsOk() (*SourcesItem, bool) {
	if o == nil || o.Lgs == nil {
		return nil, false
	}
	return o.Lgs, true
}

// HasLgs returns a boolean if a field has been set.
func (o *SourcesSaveRequest) HasLgs() bool {
	return o != nil && o.Lgs != nil
}

// SetLgs gets a reference to the given SourcesItem and assigns it to the Lgs field.
func (o *SourcesSaveRequest) SetLgs(v SourcesItem) {
	o.Lgs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcesSaveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Lgs != nil {
		toSerialize["Lgs"] = o.Lgs
	}
	if o.SmartRestRequestContext != nil {
		toSerialize["SmartRestRequestContext"] = o.SmartRestRequestContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourcesSaveRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Lgs                     *SourcesItem `json:"lgs,omitempty"`
		SmartRestRequestContext *string      `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"lgs", "smartRestRequestContext"})
	} else {
		return err
	}

	o.Lgs = all.Lgs
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
