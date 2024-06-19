package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type GeoLocationsSearchResponse struct {
	GeolocationDatas []GeoLocationsData `json:"GeolocationDatas,omitempty"`
	AfterKey         *string            `json:"AfterKey,omitempty"`
	TotalQuantity    *int64             `json:"TotalQuantity,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewGeoLocationsSearchResponse creates a new GeoLocationsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeoLocationsSearchResponse() *GeoLocationsSearchResponse {
	this := GeoLocationsSearchResponse{}
	return &this
}

// NewGeoLocationsSearchResponseWithDefaults creates a new GeoLocationsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeoLocationsSearchResponseWithDefaults() *GeoLocationsSearchResponse {
	this := GeoLocationsSearchResponse{}
	return &this
}

// GetGeolocationDatas returns the GeolocationDatas field value if set, zero value otherwise.
func (o *GeoLocationsSearchResponse) GetGeolocationDatas() []GeoLocationsData {
	if o == nil || o.GeolocationDatas == nil {
		var ret []GeoLocationsData
		return ret
	}
	return o.GeolocationDatas
}

// GetGeolocationDatasOk returns a tuple with the GeolocationDatas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsSearchResponse) GetGeolocationDatasOk() (*[]GeoLocationsData, bool) {
	if o == nil || o.GeolocationDatas == nil {
		return nil, false
	}
	return &o.GeolocationDatas, true
}

// HasGeolocationDatas returns a boolean if a field has been set.
func (o *GeoLocationsSearchResponse) HasGeolocationDatas() bool {
	return o != nil && o.GeolocationDatas != nil
}

// SetGeolocationDatas gets a reference to the given []GeoLocationsData and assigns it to the GeolocationDatas field.
func (o *GeoLocationsSearchResponse) SetGeolocationDatas(v []GeoLocationsData) {
	o.GeolocationDatas = v
}

// GetAfterKey returns the AfterKey field value if set, zero value otherwise.
func (o *GeoLocationsSearchResponse) GetAfterKey() string {
	if o == nil || o.AfterKey == nil {
		var ret string
		return ret
	}
	return *o.AfterKey
}

// GetAfterKeyOk returns a tuple with the AfterKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsSearchResponse) GetAfterKeyOk() (*string, bool) {
	if o == nil || o.AfterKey == nil {
		return nil, false
	}
	return o.AfterKey, true
}

// HasAfterKey returns a boolean if a field has been set.
func (o *GeoLocationsSearchResponse) HasAfterKey() bool {
	return o != nil && o.AfterKey != nil
}

// SetAfterKey gets a reference to the given string and assigns it to the AfterKey field.
func (o *GeoLocationsSearchResponse) SetAfterKey(v string) {
	o.AfterKey = &v
}

// GetTotalQuantity returns the TotalQuantity field value if set, zero value otherwise.
func (o *GeoLocationsSearchResponse) GetTotalQuantity() int64 {
	if o == nil || o.TotalQuantity == nil {
		var ret int64
		return ret
	}
	return *o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsSearchResponse) GetTotalQuantityOk() (*int64, bool) {
	if o == nil || o.TotalQuantity == nil {
		return nil, false
	}
	return o.TotalQuantity, true
}

// HasTotalQuantity returns a boolean if a field has been set.
func (o *GeoLocationsSearchResponse) HasTotalQuantity() bool {
	return o != nil && o.TotalQuantity != nil
}

// SetTotalQuantity gets a reference to the given int64 and assigns it to the TotalQuantity field.
func (o *GeoLocationsSearchResponse) SetTotalQuantity(v int64) {
	o.TotalQuantity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeoLocationsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.GeolocationDatas != nil {
		toSerialize["GeolocationDatas"] = o.GeolocationDatas
	}
	if o.AfterKey != nil {
		toSerialize["FailedItems"] = o.AfterKey
	}
	if o.TotalQuantity != nil {
		toSerialize["TotalQuantity"] = o.TotalQuantity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *GeoLocationsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		GeolocationDatas []GeoLocationsData `json:"GeolocationDatas,omitempty"`
		AfterKey         *string            `json:"AfterKey,omitempty"`
		TotalQuantity    *int64             `json:"TotalQuantity,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"GeolocationDatas", "AfterKey", "TotalQuantity"})
	} else {
		return err
	}

	o.GeolocationDatas = all.GeolocationDatas
	o.AfterKey = all.AfterKey
	o.TotalQuantity = all.TotalQuantity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
