package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// QueriesSearchResponse is the response object with all queries matching the request.
type QueriesSearchResponse struct {
	Items        []QueriesItems        `json:"Items,omitempty"`
	FailedItems  []json.RawMessage     `json:"FailedItems,omitempty"`
	SuccessItems []QueriesSuccessItems `json:"SuccessItems,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewQueriesSearchResponse creates a new QueriesSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueriesSearchResponse() *QueriesSearchResponse {
	this := QueriesSearchResponse{}
	return &this
}

// NewQueriesSearchResponseWithDefault creates a new QueriesSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueriesSearchResponseWithDefault() *QueriesSearchResponse {
	this := QueriesSearchResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *QueriesSearchResponse) GetItems() []QueriesItems {
	if o == nil || o.Items == nil {
		var ret []QueriesItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesSearchResponse) GetItemsOk() (*[]QueriesItems, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QueriesSearchResponse) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []QueriesItems and assigns it to the Items field.
func (o *QueriesSearchResponse) SetItems(v []QueriesItems) {
	o.Items = v
}

// GetFailedItems returns the FailedItems field value if set, zero value otherwise.
func (o *QueriesSearchResponse) GetFailedItems() []json.RawMessage {
	if o == nil || o.FailedItems == nil {
		var ret []json.RawMessage
		return ret
	}
	return o.FailedItems
}

// GetFailedItemsOk returns a tuple with the FailedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesSearchResponse) GetFailedItemsOk() (*[]json.RawMessage, bool) {
	if o == nil || o.FailedItems == nil {
		return nil, false
	}
	return &o.FailedItems, true
}

// HasFailedItems returns a boolean if a field has been set.
func (o *QueriesSearchResponse) HasFailedItems() bool {
	return o != nil && o.FailedItems != nil
}

// SetFailedItems gets a reference to the given []json.RawMessage and assigns it to the FailedItems field.
func (o *QueriesSearchResponse) SetFailedItems(v []json.RawMessage) {
	o.FailedItems = v
}

// GetSuccessItems returns the SuccessItems field value if set, zero value otherwise.
func (o *QueriesSearchResponse) GetSuccessItems() []QueriesSuccessItems {
	if o == nil || o.SuccessItems == nil {
		var ret []QueriesSuccessItems
		return ret
	}
	return o.SuccessItems
}

// GetSuccessItemsOk returns a tuple with the SuccessItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesSearchResponse) GetSuccessItemsOk() (*[]QueriesSuccessItems, bool) {
	if o == nil || o.SuccessItems == nil {
		return nil, false
	}
	return &o.SuccessItems, true
}

// HasSuccessItems returns a boolean if a field has been set.
func (o *QueriesSearchResponse) HasSuccessItems() bool {
	return o != nil && o.SuccessItems != nil
}

// SetSuccessItems gets a reference to the given []QueriesSuccessItems and assigns it to the SuccessItems field.
func (o *QueriesSearchResponse) SetSuccessItems(v []QueriesSuccessItems) {
	o.SuccessItems = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueriesSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if o.FailedItems != nil {
		toSerialize["FailedItems"] = o.FailedItems
	}
	if o.FailedItems != nil {
		toSerialize["SuccessItems"] = o.SuccessItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *QueriesSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items        []QueriesItems        `json:"Items,omitempty"`
		FailedItems  []json.RawMessage     `json:"FailedItems,omitempty"`
		SuccessItems []QueriesSuccessItems `json:"SuccessItems,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Items", "FailedItems", "SuccessItems"})
	} else {
		return err
	}

	o.Items = all.Items
	o.FailedItems = all.FailedItems
	o.SuccessItems = all.SuccessItems

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
