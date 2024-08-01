package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// EventsSearchRequest is the object sent with a request to retrieve a list of events from your organization.
type EventsSearchRequest struct {
	// Reason represents the reason for the event list request.
	Reason *string `json:"reason,omitempty"`
	// Query specifies the search and filter query settings.
	Query *QueryFilter `json:"query,omitempty"`
	// SmartRestRequestContext is the context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventsSearchRequest creates a new EventsSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsSearchRequest() *EventsSearchRequest {
	this := EventsSearchRequest{}
	var reason string = "Query"
	this.Reason = &reason
	var smartRestRequestContext string = "-<SmartRestRequestContext>-"
	this.SmartRestRequestContext = &smartRestRequestContext
	return &this
}

// NewEventsSearchRequestWithDefaults creates a new EventsSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsSearchRequestWithDefaults() *EventsSearchRequest {
	this := EventsSearchRequest{}
	var reason string = "Query"
	this.Reason = &reason
	var smartRestRequestContext string = "-<SmartRestRequestContext>-"
	this.SmartRestRequestContext = &smartRestRequestContext
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EventsSearchRequest) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSearchRequest) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EventsSearchRequest) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EventsSearchRequest) SetReason(v string) {
	o.Reason = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *EventsSearchRequest) GetQuery() QueryFilter {
	if o == nil || o.Query == nil {
		var ret QueryFilter
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSearchRequest) GetQueryOk() (*QueryFilter, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *EventsSearchRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given QueryFilter and assigns it to the Query field.
func (o *EventsSearchRequest) SetQuery(v QueryFilter) {
	o.Query = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *EventsSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *EventsSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *EventsSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
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
func (o *EventsSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason                  *string      `json:"reason,omitempty"`
		Query                   *QueryFilter `json:"query,omitempty"`
		SmartRestRequestContext *string      `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Reason == nil {
		return fmt.Errorf("requiered field reason is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"reason", "query", "smartRestRequestContext"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = all.Query
	o.Reason = all.Reason
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
