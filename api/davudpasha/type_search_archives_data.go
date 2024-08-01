package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SearchArchivesData represents data structure for searching archives.
type SearchArchivesData struct {
	// Query contains the search query parameters.
	Query *SearchArchivesQuery `json:"Query,omitempty"`
	// Status represents the status of the search operation.
	Status *SearchArchivesStatus `json:"Status,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSearchArchivesData creates a new SearchArchivesData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchArchivesData() *SearchArchivesData {
	this := SearchArchivesData{}
	return &this
}

// NewSearchArchivesDataWithDefaults creates a new SearchArchivesData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSearchArchivesDataWithDefaults() *SearchArchivesData {
	this := SearchArchivesData{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SearchArchivesData) GetQuery() SearchArchivesQuery {
	if o == nil || o.Query == nil {
		var ret SearchArchivesQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesData) GetQueryOk() (*SearchArchivesQuery, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SearchArchivesData) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given SearchArchivesQuery and assigns it to the Query field.
func (o *SearchArchivesData) SetQuery(v SearchArchivesQuery) {
	o.Query = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SearchArchivesData) GetStatus() SearchArchivesStatus {
	if o == nil || o.Status == nil {
		var ret SearchArchivesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesData) GetStatusOk() (*SearchArchivesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SearchArchivesData) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SearchArchivesStatus and assigns it to the Status field.
func (o *SearchArchivesData) SetStatus(v SearchArchivesStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SearchArchivesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SearchArchivesData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query  *SearchArchivesQuery  `json:"Query,omitempty"`
		Status *SearchArchivesStatus `json:"Status,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Query", "Status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = all.Query
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
