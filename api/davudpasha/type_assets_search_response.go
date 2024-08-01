package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// AssetsSearchResponse represents the response structure for a search query for assets.
type AssetsSearchResponse struct {
	// List of assets matching the search criteria.
	Assets []AssetsAsset `json:"assets,omitempty"`
	// Total count of assets matching the search criteria.
	Count *int64 `json:"count,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssetsSearchResponse creates a new AssetsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssetsSearchResponse() *AssetsSearchResponse {
	this := AssetsSearchResponse{}
	return &this
}

// NewAssetsSearchResponseWithDefaults creates a new AssetsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssetsSearchResponseWithDefaults() *AssetsSearchResponse {
	this := AssetsSearchResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *AssetsSearchResponse) GetAssets() []AssetsAsset {
	if o == nil || o.Assets == nil {
		var ret []AssetsAsset
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsSearchResponse) GetAssetsOk() (*[]AssetsAsset, bool) {
	if o == nil || o.Assets == nil {
		return nil, false
	}
	return &o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *AssetsSearchResponse) HasAssets() bool {
	return o != nil && o.Assets != nil
}

// SetAssets gets a reference to the given []AssetsAsset and assigns it to the Assets field.
func (o *AssetsSearchResponse) SetAssets(v []AssetsAsset) {
	o.Assets = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AssetsSearchResponse) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsSearchResponse) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AssetsSearchResponse) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *AssetsSearchResponse) SetCount(v int64) {
	o.Count = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssetsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Assets != nil {
		toSerialize["Assets"] = o.Assets
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AssetsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assets []AssetsAsset `json:"assets,omitempty"`
		Count  *int64        `json:"count,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Assets", "Count"})
	} else {
		return err
	}

	o.Assets = all.Assets
	o.Count = all.Count

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
