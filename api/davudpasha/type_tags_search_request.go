package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type TagsSearchRequest struct {
	Category     TagCategory `json:"category,omitempty"`
	SearchFilter *string     `json:"searchFilter,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewTagsSearchRequest creates a new TagsSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagsSearchRequest() *TagsSearchRequest {
	this := TagsSearchRequest{}
	return &this
}

// NewTagsSearchRequestWithDefaults creates a new TagsSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagsSearchRequestWithDefaults() *TagsSearchRequest {
	this := TagsSearchRequest{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *TagsSearchRequest) GetCategory() TagCategory {
	if o == nil {
		var ret TagCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsSearchRequest) GetCategoryOk() (*TagCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *TagsSearchRequest) SetCategory(v TagCategory) {
	o.Category = v
}

// GetSearchFilter returns the SearchFilter field value if set, zero value otherwise.
func (o *TagsSearchRequest) GetSearchFilter() string {
	if o == nil || o.SearchFilter == nil {
		var ret string
		return ret
	}
	return *o.SearchFilter
}

// GetSearchFilterOk returns a tuple with the SearchFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsSearchRequest) GetSearchFilterOk() (*string, bool) {
	if o == nil || o.SearchFilter == nil {
		return nil, false
	}
	return o.SearchFilter, true
}

// HasSearchFilter returns a boolean if a field has been set.
func (o *TagsSearchRequest) HasSearchFilter() bool {
	return o != nil && o.SearchFilter != nil
}

// SetSearchFilter gets a reference to the given string and assigns it to the SearchFilter field.
func (o *TagsSearchRequest) SetSearchFilter(v string) {
	o.SearchFilter = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *TagsSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *TagsSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *TagsSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	if o.SearchFilter != nil {
		toSerialize["searchFilter"] = o.SearchFilter
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
func (o *TagsSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category     *TagCategory `json:"category,omitempty"`
		SearchFilter *string      `json:"searchFilter,omitempty"`
		// Context for the Smart REST request.
		SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("requiered field category is missing")
	}
	if all.SearchFilter == nil {
		return fmt.Errorf("requiered field searchFilter is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"category", "searchFilter", "smartRestRequestContext"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.SearchFilter = all.SearchFilter
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
