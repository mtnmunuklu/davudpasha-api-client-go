package davutpasa

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davutpasa-api-client-go/api/common"
)

// EventListRequest The object sent with request to retrieve a list of event from your organization.
type EventsListRequest struct {
	Reason *string `json:"reason,omitempty"`
	// The search and filter query settings.
	Query                   *EventsQueryFilter `json:"query,omitempty"`
	SmartRestRequestContext *string            `json:"smartRestRequestContext,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewEventsListRequest creates a new EventsListRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsListRequest() *EventsListRequest {
	this := EventsListRequest{}
	var reason string = "Query"
	this.Reason = &reason
	var smartRestRequestContext string = "-<SmartRestRequestContext>-"
	this.SmartRestRequestContext = &smartRestRequestContext
	return &this
}

// NewEventsListRequestWithDefaults creates a new EventsListRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsListRequestWithDefaults() *EventsListRequest {
	this := EventsListRequest{}
	var reason string = "Query"
	this.Reason = &reason
	var smartRestRequestContext string = "-<SmartRestRequestContext>-"
	this.SmartRestRequestContext = &smartRestRequestContext
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EventsListRequest) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsListRequest) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EventsListRequest) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EventsListRequest) SetReason(v string) {
	o.Reason = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *EventsListRequest) GetQuery() EventsQueryFilter {
	if o == nil || o.Query == nil {
		var ret EventsQueryFilter
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsListRequest) GetQueryOk() (*EventsQueryFilter, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *EventsListRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given EventsQueryFilter and assigns it to the Query field.
func (o *EventsListRequest) SetQuery(v EventsQueryFilter) {
	o.Query = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *EventsListRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsListRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *EventsListRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *EventsListRequest) SetSmartRestRequestContext(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsListRequest) MarshalJSON() ([]byte, error) {
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
func (o *EventsListRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason                  *string            `json:"reason,omitempty"`
		Query                   *EventsQueryFilter `json:"query,omitempty"`
		SmartRestRequestContext *string            `json:"smartRestRequestContext,omitempty"`
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
