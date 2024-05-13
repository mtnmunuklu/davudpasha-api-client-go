package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SuccessItems struct {
	Application *string  `json:"Application,omitempty"`
	Keys        []string `json:"Keys,omitempty"`
	SourceType  *string  `json:"SourceType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewSuccessItems creates a new SuccessItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSuccessItems() *SuccessItems {
	this := SuccessItems{}
	return &this
}

// NewSuccessItemsWithDefaults creates a new SuccessItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSuccessItemsWithDefault() *SuccessItems {
	this := SuccessItems{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *SuccessItems) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessItems) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *SuccessItems) HasApplication() bool {
	return o != nil && o.Application != nil
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *SuccessItems) SetApplication(v string) {
	o.Application = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *SuccessItems) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessItems) GetKeysOk() (*[]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return &o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *SuccessItems) HasKeys() bool {
	return o != nil && o.Keys != nil
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *SuccessItems) SetKeys(v []string) {
	o.Keys = v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *SuccessItems) GetSourceType() string {
	if o == nil || o.SourceType == nil {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessItems) GetSourceTypeOk() (*string, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *SuccessItems) HasSourceType() bool {
	return o != nil && o.SourceType != nil
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *SuccessItems) SetSourceType(v string) {
	o.SourceType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SuccessItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Application != nil {
		toSerialize["Application"] = o.Application
	}
	if o.Keys != nil {
		toSerialize["Keys"] = o.Keys
	}
	if o.SourceType != nil {
		toSerialize["SourceType"] = o.SourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SuccessItems) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Application *string  `json:"Application,omitempty"`
		Keys        []string `json:"Keys,omitempty"`
		SourceType  *string  `json:"SourceType,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Application", "Keys", "SourceType"})
	} else {
		return err
	}

	o.Application = all.Application
	o.Keys = all.Keys
	o.SourceType = all.SourceType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
