package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// TagsSearchResponse represents the response structure for tag search.
type TagsSearchResponse struct {
	// Items contains the list of tags retrieved from the search.
	Items []TagsItem `json:"Items,omitempty"`
	// FailedItems contains a list of failed items in raw JSON format.
	FailedItems []json.RawMessage `json:"FailedItems,omitempty"`
	// SuccessItems contains a list of successful items.
	SuccessItems []SuccessItem `json:"SuccessItems,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewTagsSearchResponse creates a new TagsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagsSearchResponse() *TagsSearchResponse {
	this := TagsSearchResponse{}
	return &this
}

// NewTagsSearchResponseWithDefaults creates a new TagsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagsSearchResponseWithDefaults() *TagsSearchResponse {
	this := TagsSearchResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TagsSearchResponse) GetItems() []TagsItem {
	if o == nil || o.Items == nil {
		var ret []TagsItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsSearchResponse) GetItemsOk() (*[]TagsItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TagsSearchResponse) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []TagsItem and assigns it to the Items field.
func (o *TagsSearchResponse) SetItems(v []TagsItem) {
	o.Items = v
}

// GetFailedItems returns the FailedItems field value if set, zero value otherwise.
func (o *TagsSearchResponse) GetFailedItems() []json.RawMessage {
	if o == nil || o.FailedItems == nil {
		var ret []json.RawMessage
		return ret
	}
	return o.FailedItems
}

// GetFailedItemsOk returns a tuple with the FailedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsSearchResponse) GetFailedItemsOk() (*[]json.RawMessage, bool) {
	if o == nil || o.FailedItems == nil {
		return nil, false
	}
	return &o.FailedItems, true
}

// HasFailedItems returns a boolean if a field has been set.
func (o *TagsSearchResponse) HasFailedItems() bool {
	return o != nil && o.FailedItems != nil
}

// SetFailedItems sets field value
func (o *TagsSearchResponse) SetFailedItems(v []json.RawMessage) {
	o.FailedItems = v
}

// GetSuccessItem returns the SuccessItem field value if set, zero value otherwise.
func (o *TagsSearchResponse) GetSuccessItem() []SuccessItem {
	if o == nil || o.SuccessItems == nil {
		var ret []SuccessItem
		return ret
	}
	return o.SuccessItems
}

// GetSuccessItemOk returns a tuple with the SuccessItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsSearchResponse) GetSuccessItemOk() (*[]SuccessItem, bool) {
	if o == nil || o.SuccessItems == nil {
		return nil, false
	}
	return &o.SuccessItems, true
}

// HasSuccessItem returns a boolean if a field has been set.
func (o *TagsSearchResponse) HasSuccessItem() bool {
	return o != nil && o.SuccessItems != nil
}

// SetSuccessItem sets field value
func (o *TagsSearchResponse) SetSuccessItem(v []SuccessItem) {
	o.SuccessItems = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagsSearchResponse) MarshalJSON() ([]byte, error) {
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
	if o.SuccessItems != nil {
		toSerialize["SuccessItems"] = o.SuccessItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *TagsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items        []TagsItem        `json:"Items,omitempty"`
		FailedItems  []json.RawMessage `json:"FailedItems,omitempty"`
		SuccessItems []SuccessItem     `json:"SuccessItems,omitempty"`
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
