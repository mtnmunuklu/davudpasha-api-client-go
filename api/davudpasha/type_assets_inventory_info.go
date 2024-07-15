package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// AssetsInventoryInfo represents the inventory information of an asset.
type AssetsInventoryInfo struct {
	// Information about the operating system of the asset.
	OperationSystemInfos *AssetsOperationSystemInfo `json:"OperationSystemInfos,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewAssetsInventoryInfo creates a new AssetsInventoryInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssetsInventoryInfo() *AssetsInventoryInfo {
	this := AssetsInventoryInfo{}
	return &this
}

// NewAssetsInventoryInfoWithDefaults creates a new AssetsInventoryInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssetsInventoryInfoWithDefault() *AssetsInventoryInfo {
	this := AssetsInventoryInfo{}
	return &this
}

// GetOperationSystemInfos returns the OperationSystemInfos field value if set, zero value otherwise.
func (o *AssetsInventoryInfo) GetOperationSystemInfos() AssetsOperationSystemInfo {
	if o == nil || o.OperationSystemInfos == nil {
		var ret AssetsOperationSystemInfo
		return ret
	}
	return *o.OperationSystemInfos
}

// GetOperationSystemInfosOk returns a tuple with the OperationSystemInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInventoryInfo) GetOperationSystemInfosOk() (*AssetsOperationSystemInfo, bool) {
	if o == nil || o.OperationSystemInfos == nil {
		return nil, false
	}
	return o.OperationSystemInfos, true
}

// HasOperationSystemInfos returns a boolean if a field has been set.
func (o *AssetsInventoryInfo) HasOperationSystemInfos() bool {
	return o != nil && o.OperationSystemInfos != nil
}

// SetOperationSystemInfos gets a reference to the given AssetsOperationSystemInfo and assigns it to the OperationSystemInfos field.
func (o *AssetsInventoryInfo) SetOperationSystemInfos(v AssetsOperationSystemInfo) {
	o.OperationSystemInfos = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssetsInventoryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.OperationSystemInfos != nil {
		toSerialize["OperationSystemInfos"] = o.OperationSystemInfos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AssetsInventoryInfo) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		OperationSystemInfos *AssetsOperationSystemInfo `json:"OperationSystemInfos,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"OperationSystemInfos"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OperationSystemInfos != nil && all.OperationSystemInfos.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OperationSystemInfos = all.OperationSystemInfos

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
