package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SystemLogsSearchResponse struct {
	Items      []SystemLogsItem `json:"Items,omitempty"`
	TotalCount *int64           `json:"TotalCount,omitempty"`
	PageSize   *int64           `json:"PageSize,omitempty"`
	FromIndex  *int64           `json:"FromIndex,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSystemLogsSearchResponse creates a new SystemLogsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSystemLogsSearchResponse() *SystemLogsSearchResponse {
	this := SystemLogsSearchResponse{}
	return &this
}

// NewSystemLogsSearchResponseWithDefaults creates a new SystemLogsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSystemLogsSearchResponseWithDefaults() *SystemLogsSearchResponse {
	this := SystemLogsSearchResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SystemLogsSearchResponse) GetItems() []SystemLogsItem {
	if o == nil || o.Items == nil {
		var ret []SystemLogsItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchResponse) GetItemsOk() (*[]SystemLogsItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SystemLogsSearchResponse) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []SystemLogsItem and assigns it to the Items field.
func (o *SystemLogsSearchResponse) SetItems(v []SystemLogsItem) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *SystemLogsSearchResponse) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchResponse) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *SystemLogsSearchResponse) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *SystemLogsSearchResponse) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SystemLogsSearchResponse) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchResponse) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SystemLogsSearchResponse) HasPageSize() bool {
	return o != nil && o.PageSize != nil
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *SystemLogsSearchResponse) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetFromIndex returns the FromIndex field value if set, zero value otherwise.
func (o *SystemLogsSearchResponse) GetFromIndex() int64 {
	if o == nil || o.FromIndex == nil {
		var ret int64
		return ret
	}
	return *o.FromIndex
}

// GetFromIndexOk returns a tuple with the FromIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsSearchResponse) GetFromIndexOk() (*int64, bool) {
	if o == nil || o.FromIndex == nil {
		return nil, false
	}
	return o.FromIndex, true
}

// HasFromIndex returns a boolean if a field has been set.
func (o *SystemLogsSearchResponse) HasFromIndex() bool {
	return o != nil && o.FromIndex != nil
}

// SetFromIndex gets a reference to the given int64 and assigns it to the FromIndex field.
func (o *SystemLogsSearchResponse) SetFromIndex(v int64) {
	o.FromIndex = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SystemLogsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if o.TotalCount != nil {
		toSerialize["TotalCount"] = o.TotalCount
	}
	if o.PageSize != nil {
		toSerialize["PageSize"] = o.PageSize
	}
	if o.FromIndex != nil {
		toSerialize["FromIndex"] = o.FromIndex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SystemLogsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      []SystemLogsItem `json:"Items,omitempty"`
		TotalCount *int64           `json:"TotalCount,omitempty"`
		PageSize   *int64           `json:"PageSize,omitempty"`
		FromIndex  *int64           `json:"FromIndex,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Items", "TotalCount", "PageSize", "FromIndex"})
	} else {
		return err
	}

	o.Items = all.Items
	o.TotalCount = all.TotalCount
	o.PageSize = all.PageSize
	o.FromIndex = all.FromIndex

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
