package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// TagsItem represents a single tag item structure.
type TagsItem struct {
	// ID represents the identifier of the tag.
	ID *string `json:"ID,omitempty"`
	// Name specifies the name of the tag.
	Name *string `json:"Name,omitempty"`
	// Category specifies the category of the tag.
	Category TagCategory `json:"Category,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewTagsItem creates a new TagsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagsItem() *TagsItem {
	this := TagsItem{}
	return &this
}

// NewTagsItemWithDefaults creates a new TagsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagsItemWithDefaults() *TagsItem {
	this := TagsItem{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TagsItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TagsItem) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TagsItem) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagsItem) SetName(v string) {
	o.Name = &v
}

// GetDateTimeType returns the DateTimeType field value if set, zero value otherwise.
func (o *TagsItem) GetDateTimeType() TagCategory {
	if o == nil {
		var ret TagCategory
		return ret
	}
	return o.Category
}

// GetDateTimeTypeOk returns a tuple with the DateTimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsItem) GetDateTimeTypeOk() (*TagCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetDateTimeType gets a reference to the given string and assigns it to the DateTimeType field.
func (o *TagsItem) SetDateTimeType(v TagCategory) {
	o.Category = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	toSerialize["Category"] = o.Category

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *TagsItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID       *string      `json:"ID,omitempty"`
		Name     *string      `json:"Name,omitempty"`
		Category *TagCategory `json:"Category,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ID", "Name", "Category"})
	} else {
		return err
	}
	o.ID = all.ID
	o.Name = all.Name
	o.Category = *all.Category

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
