package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ElasticStatsSearchResponse represents the response structure for searching Elasticsearch statistics.
type ElasticStatsSearchResponse struct {
	// Status of the Elasticsearch statistics search operation.
	Status *string `json:"Status,omitempty"`
	// List of Elasticsearch indexes with their respective statistics.
	Indexes []ElasticStatsIndexes `json:"Indexes,omitempty"`
	// Result of the Elasticsearch statistics search.
	Result *ElasticStatsResult `json:"Result,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewElasticStatsSearchResponse creates a new ElasticStatsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticStatsSearchResponse() *ElasticStatsSearchResponse {
	this := ElasticStatsSearchResponse{}
	return &this
}

// NewElasticStatsSearchResponseWithDefaults creates a new ElasticStatsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticStatsSearchResponseWithDefaults() *ElasticStatsSearchResponse {
	this := ElasticStatsSearchResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ElasticStatsSearchResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsSearchResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ElasticStatsSearchResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ElasticStatsSearchResponse) SetStatus(v string) {
	o.Status = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *ElasticStatsSearchResponse) GetIndexes() []ElasticStatsIndexes {
	if o == nil || o.Indexes == nil {
		var ret []ElasticStatsIndexes
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsSearchResponse) GetIndexesOk() (*[]ElasticStatsIndexes, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *ElasticStatsSearchResponse) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []ElasticStatsIndexes and assigns it to the Indexes field.
func (o *ElasticStatsSearchResponse) SetIndexes(v []ElasticStatsIndexes) {
	o.Indexes = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ElasticStatsSearchResponse) GetResult() ElasticStatsResult {
	if o == nil || o.Result == nil {
		var ret ElasticStatsResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsSearchResponse) GetResultOk() (*ElasticStatsResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ElasticStatsSearchResponse) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given ElasticStatsResult and assigns it to the Result field.
func (o *ElasticStatsSearchResponse) SetResult(v ElasticStatsResult) {
	o.Result = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticStatsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Indexes != nil {
		toSerialize["Indexes"] = o.Indexes
	}
	if o.Result != nil {
		toSerialize["Result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ElasticStatsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status  *string               `json:"Status,omitempty"`
		Indexes []ElasticStatsIndexes `json:"Indexes,omitempty"`
		Result  *ElasticStatsResult   `json:"Result,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Status", "Indexes", "Result"})
	} else {
		return err
	}

	o.Status = all.Status
	o.Indexes = all.Indexes
	hasInvalidField := false
	if all.Result != nil && all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = all.Result

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
