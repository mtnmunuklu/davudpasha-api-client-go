package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// QueriesSaveResponse represents the response structure when saving queries.
type QueriesSaveResponse struct {
	// Indicates whether the save operation was successful.
	Status *bool `json:"status,omitempty"`
	// The query item that was saved.
	Query *QueriesItem `json:"query,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQueriesSaveResponse creates a new QueriesSaveResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueriesSaveResponse() *QueriesSaveResponse {
	this := QueriesSaveResponse{}
	return &this
}

// NewQueriesSaveResponseWithDefaults creates a new QueriesSaveResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueriesSaveResponseWithDefaults() *QueriesSaveResponse {
	this := QueriesSaveResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueriesSaveResponse) GetStatus() bool {
	if o == nil || o.Status == nil {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesSaveResponse) GetStatusOk() (*bool, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueriesSaveResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *QueriesSaveResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *QueriesSaveResponse) GetQuery() QueriesItem {
	if o == nil || o.Query == nil {
		var ret QueriesItem
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesSaveResponse) GetQueryOk() (*QueriesItem, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *QueriesSaveResponse) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given QueriesItem and assigns it to the Query field.
func (o *QueriesSaveResponse) SetQuery(v QueriesItem) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueriesSaveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *QueriesSaveResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *bool        `json:"status,omitempty"`
		Query  *QueriesItem `json:"query,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("requiered field status is missing")
	}
	if all.Query == nil {
		return fmt.Errorf("requiered field query is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "query"})
	} else {
		return err
	}
	o.Status = all.Status

	hasInvalidField := false
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
