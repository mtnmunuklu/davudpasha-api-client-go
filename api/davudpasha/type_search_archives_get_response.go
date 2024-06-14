package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SearchArchivesGetResponse struct {
	Data []SearchArchivesData `json:"Data,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSearchArchivesGetResponse creates a new SearchArchivesGetResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchArchivesGetResponse() *SearchArchivesGetResponse {
	this := SearchArchivesGetResponse{}
	return &this
}

// NewSearchArchivesGetResponseWithDefaults creates a new SearchArchivesGetResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSearchArchivesGetResponseWithDefaults() *SearchArchivesGetResponse {
	this := SearchArchivesGetResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SearchArchivesGetResponse) GetData() []SearchArchivesData {
	if o == nil || o.Data == nil {
		var ret []SearchArchivesData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesGetResponse) GetDataOk() (*[]SearchArchivesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SearchArchivesGetResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []SearchArchivesData and assigns it to the Data field.
func (o *SearchArchivesGetResponse) SetData(v []SearchArchivesData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SearchArchivesGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SearchArchivesGetResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []SearchArchivesData `json:"Data,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Data"})
	} else {
		return err
	}

	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
