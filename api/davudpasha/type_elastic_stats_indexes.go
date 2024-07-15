package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ElasticStatsIndexes represents statistics for Elasticsearch indexes.
type ElasticStatsIndexes struct {
	// Number of documents in the index.
	DocsCount *int64 `json:"DocsCount,omitempty"`
	// Health status of the index.
	Health *string `json:"Health,omitempty"`
	// Name of the index.
	Index *string `json:"Index,omitempty"`
	// Store size of the index in megabytes.
	StoreSizeMB *float64 `json:"StoreSizeMB,omitempty"`
	// Status of the index.
	Status *string `json:"Status,omitempty"`
	// Type of the index.
	IndexType *string `json:"IndexType,omitempty"`
	// Replica count for the index.
	ReplicaCount *string `json:"ReplicaCount,omitempty"`
	// Primary shard count for the index.
	PrimaryCount *string `json:"PrimaryCount,omitempty"`
	// Size of primary shards for the index.
	PrimaryStoreSize *string `json:"PrimaryStoreSize,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewElasticStatsIndexes creates a new ElasticStatsIndexes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticStatsIndexes() *ElasticStatsIndexes {
	this := ElasticStatsIndexes{}
	return &this
}

// NewElasticStatsIndexesWithDefaults creates a new ElasticStatsIndexes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticStatsIndexesWithDefaults() *ElasticStatsIndexes {
	this := ElasticStatsIndexes{}
	return &this
}

// GetDocsCount returns the DocsCount field value if set, zero value otherwise.
func (o *ElasticStatsIndexes) GetDocsCount() int64 {
	if o == nil || o.DocsCount == nil {
		var ret int64
		return ret
	}
	return *o.DocsCount
}

// GetDocsCountOk returns a tuple with the DocsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsIndexes) GetDocsCountOk() (*int64, bool) {
	if o == nil || o.DocsCount == nil {
		return nil, false
	}
	return o.DocsCount, true
}

// HasDocsCount returns a boolean if a field has been set.
func (o *ElasticStatsIndexes) HasDocsCount() bool {
	return o != nil && o.DocsCount != nil
}

// SetDocsCount gets a reference to the given int64 and assigns it to the DocsCount field.
func (o *ElasticStatsIndexes) SetDocsCount(v int64) {
	o.DocsCount = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ElasticStatsIndexes) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsIndexes) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *ElasticStatsIndexes) HasHealth() bool {
	return o != nil && o.Health != nil
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *ElasticStatsIndexes) SetHealth(v string) {
	o.Health = &v
}

// GetStoreSizeMB returns the StoreSizeMB field value if set, zero value otherwise.
func (o *ElasticStatsIndexes) GetStoreSizeMB() float64 {
	if o == nil || o.StoreSizeMB == nil {
		var ret float64
		return ret
	}
	return *o.StoreSizeMB
}

// GetStoreSizeMBOk returns a tuple with the StoreSizeMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsIndexes) GetStoreSizeMBOk() (*float64, bool) {
	if o == nil || o.StoreSizeMB == nil {
		return nil, false
	}
	return o.StoreSizeMB, true
}

// HasStoreSizeMB returns a boolean if a field has been set.
func (o *ElasticStatsIndexes) HasStoreSizeMB() bool {
	return o != nil && o.StoreSizeMB != nil
}

// SetStoreSizeMB gets a reference to the given float64 and assigns it to the StoreSizeMB field.
func (o *ElasticStatsIndexes) SetStoreSizeMB(v float64) {
	o.StoreSizeMB = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ElasticStatsIndexes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsIndexes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ElasticStatsIndexes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ElasticStatsIndexes) SetStatus(v string) {
	o.Status = &v
}

// GetIndexType returns the IndexType field value if set, zero value otherwise.
func (o *ElasticStatsIndexes) GetIndexType() string {
	if o == nil || o.IndexType == nil {
		var ret string
		return ret
	}
	return *o.IndexType
}

// GetIndexTypeOk returns a tuple with the IndexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsIndexes) GetIndexTypeOk() (*string, bool) {
	if o == nil || o.IndexType == nil {
		return nil, false
	}
	return o.IndexType, true
}

// HasIndexType returns a boolean if a field has been set.
func (o *ElasticStatsIndexes) HasIndexType() bool {
	return o != nil && o.IndexType != nil
}

// SetIndexType gets a reference to the given string and assigns it to the IndexType field.
func (o *ElasticStatsIndexes) SetIndexType(v string) {
	o.IndexType = &v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *ElasticStatsIndexes) GetReplicaCount() string {
	if o == nil || o.ReplicaCount == nil {
		var ret string
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsIndexes) GetReplicaCountOk() (*string, bool) {
	if o == nil || o.ReplicaCount == nil {
		return nil, false
	}
	return o.ReplicaCount, true
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *ElasticStatsIndexes) HasReplicaCount() bool {
	return o != nil && o.ReplicaCount != nil
}

// SetReplicaCount gets a reference to the given string and assigns it to the ReplicaCount field.
func (o *ElasticStatsIndexes) SetReplicaCount(v string) {
	o.ReplicaCount = &v
}

// GetPrimaryCount returns the PrimaryCount field value if set, zero value otherwise.
func (o *ElasticStatsIndexes) GetPrimaryCount() string {
	if o == nil || o.PrimaryCount == nil {
		var ret string
		return ret
	}
	return *o.PrimaryCount
}

// GetPrimaryCountOk returns a tuple with the PrimaryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsIndexes) GetPrimaryCountOk() (*string, bool) {
	if o == nil || o.PrimaryCount == nil {
		return nil, false
	}
	return o.PrimaryCount, true
}

// HasPrimaryCount returns a boolean if a field has been set.
func (o *ElasticStatsIndexes) HasPrimaryCount() bool {
	return o != nil && o.PrimaryCount != nil
}

// SetPrimaryCount gets a reference to the given string and assigns it to the PrimaryCount field.
func (o *ElasticStatsIndexes) SetPrimaryCount(v string) {
	o.PrimaryCount = &v
}

// GetPrimaryStoreSize returns the PrimaryStoreSize field value if set, zero value otherwise.
func (o *ElasticStatsIndexes) GetPrimaryStoreSize() string {
	if o == nil || o.PrimaryStoreSize == nil {
		var ret string
		return ret
	}
	return *o.PrimaryStoreSize
}

// GetPrimaryStoreSizeOk returns a tuple with the PrimaryStoreSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticStatsIndexes) GetPrimaryStoreSizeOk() (*string, bool) {
	if o == nil || o.PrimaryStoreSize == nil {
		return nil, false
	}
	return o.PrimaryStoreSize, true
}

// HasPrimaryStoreSize returns a boolean if a field has been set.
func (o *ElasticStatsIndexes) HasPrimaryStoreSize() bool {
	return o != nil && o.PrimaryStoreSize != nil
}

// SetPrimaryStoreSize gets a reference to the given string and assigns it to the PrimaryStoreSize field.
func (o *ElasticStatsIndexes) SetPrimaryStoreSize(v string) {
	o.PrimaryStoreSize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticStatsIndexes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.DocsCount != nil {
		toSerialize["DocsCount"] = o.DocsCount
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	if o.StoreSizeMB != nil {
		toSerialize["StoreSizeMB"] = o.StoreSizeMB
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.IndexType != nil {
		toSerialize["IndexType"] = o.IndexType
	}
	if o.ReplicaCount != nil {
		toSerialize["ReplicaCount"] = o.ReplicaCount
	}
	if o.PrimaryCount != nil {
		toSerialize["PrimaryCount"] = o.PrimaryCount
	}
	if o.PrimaryStoreSize != nil {
		toSerialize["PrimaryStoreSize"] = o.PrimaryStoreSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ElasticStatsIndexes) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		DocsCount        *int64   `json:"DocsCount,omitempty"`
		Health           *string  `json:"Health,omitempty"`
		Index            *string  `json:"Index,omitempty"`
		StoreSizeMB      *float64 `json:"StoreSizeMB,omitempty"`
		Status           *string  `json:"Status,omitempty"`
		IndexType        *string  `json:"IndexType,omitempty"`
		ReplicaCount     *string  `json:"ReplicaCount,omitempty"`
		PrimaryCount     *string  `json:"PrimaryCount,omitempty"`
		PrimaryStoreSize *string  `json:"PrimaryStoreSize,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"DocsCount", "Health", "Index", "StoreSizeMB", "Status", "IndexType", "ReplicaCount", "PrimaryCount", "PrimaryStoreSize"})
	} else {
		return err
	}

	o.DocsCount = all.DocsCount
	o.Health = all.Health
	o.Index = all.Index
	o.StoreSizeMB = all.StoreSizeMB
	o.Status = all.Status
	o.IndexType = all.IndexType
	o.ReplicaCount = all.ReplicaCount
	o.PrimaryCount = all.PrimaryCount
	o.PrimaryStoreSize = all.PrimaryStoreSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
