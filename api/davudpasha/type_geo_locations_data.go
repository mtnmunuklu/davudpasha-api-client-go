package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type GeoLocationsData struct {
	Id          *string               `json:"id,omitempty"`
	StartIp     *string               `json:"startip,omitempty"`
	EndIp       *string               `json:"endip,omitempty"`
	CountryCode *string               `json:"countrycode,omitempty"`
	Country     *string               `json:"country,omitempty"`
	City        *string               `json:"city,omitempty"`
	Region      *string               `json:"region,omitempty"`
	Latitude    *float64              `json:"latitude,omitempty"`
	Longitude   *float64              `json:"longitude,omitempty"`
	IpRange     common.NullableString `json:"iprange,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewGeoLocationsData creates a new GeoLocationsData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeoLocationsData() *GeoLocationsData {
	this := GeoLocationsData{}
	return &this
}

// NewGeoLocationsDataWithDefaults creates a new GeoLocationsData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeoLocationsDataWithDefaults() *GeoLocationsData {
	this := GeoLocationsData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GeoLocationsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GeoLocationsData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GeoLocationsData) SetId(v string) {
	o.Id = &v
}

// GetStartIp returns the StartIp field value if set, zero value otherwise.
func (o *GeoLocationsData) GetStartIp() string {
	if o == nil || o.StartIp == nil {
		var ret string
		return ret
	}
	return *o.StartIp
}

// GetStartIpOk returns a tuple with the StartIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetStartIpOk() (*string, bool) {
	if o == nil || o.StartIp == nil {
		return nil, false
	}
	return o.StartIp, true
}

// HasStartIp returns a boolean if a field has been set.
func (o *GeoLocationsData) HasStartIp() bool {
	return o != nil && o.StartIp != nil
}

// SetStartIp gets a reference to the given string and assigns it to the StartIp field.
func (o *GeoLocationsData) SetStartIp(v string) {
	o.StartIp = &v
}

// GetEndIp returns the EndIp field value if set, zero value otherwise.
func (o *GeoLocationsData) GetEndIp() string {
	if o == nil || o.EndIp == nil {
		var ret string
		return ret
	}
	return *o.EndIp
}

// GetEndIpOk returns a tuple with the EndIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetEndIpOk() (*string, bool) {
	if o == nil || o.EndIp == nil {
		return nil, false
	}
	return o.EndIp, true
}

// HasEndIp returns a boolean if a field has been set.
func (o *GeoLocationsData) HasEndIp() bool {
	return o != nil && o.EndIp != nil
}

// SetEndIp gets a reference to the given string and assigns it to the EndIp field.
func (o *GeoLocationsData) SetEndIp(v string) {
	o.EndIp = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *GeoLocationsData) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *GeoLocationsData) HasCountryCode() bool {
	return o != nil && o.CountryCode != nil
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *GeoLocationsData) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GeoLocationsData) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GeoLocationsData) HasCountry() bool {
	return o != nil && o.Country != nil
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GeoLocationsData) SetCountry(v string) {
	o.Country = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *GeoLocationsData) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *GeoLocationsData) HasCity() bool {
	return o != nil && o.City != nil
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *GeoLocationsData) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GeoLocationsData) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GeoLocationsData) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GeoLocationsData) SetRegion(v string) {
	o.Region = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *GeoLocationsData) GetLatitude() float64 {
	if o == nil || o.Latitude == nil {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetLatitudeOk() (*float64, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *GeoLocationsData) HasLatitude() bool {
	return o != nil && o.Latitude != nil
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *GeoLocationsData) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *GeoLocationsData) GetLongitude() float64 {
	if o == nil || o.Longitude == nil {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetLongitudeOk() (*float64, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *GeoLocationsData) HasLongitude() bool {
	return o != nil && o.Longitude != nil
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *GeoLocationsData) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeoLocationsData) GetIpRange() string {
	if o == nil || o.IpRange.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpRange.Get()
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GeoLocationsData) GetIpRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpRange.Get(), o.IpRange.IsSet()
}

// HasIpRange returns a boolean if a IpRange has been set.
func (o *GeoLocationsData) HasIpRange() bool {
	return o != nil && o.IpRange.IsSet()
}

// SetIpRange gets a reference to the given datadog.NullableString and assigns it to the IpRange field.
func (o *GeoLocationsData) SetIpRange(v string) {
	o.IpRange.Set(&v)
}

// SetIpRangeNil sets the value for IpRange to be an explicit nil.
func (o *GeoLocationsData) SetIpRangeNil() {
	o.IpRange.Set(nil)
}

// UnSetIpRange ensures that no value is present for IpRange, not even an explicit nil.
func (o *GeoLocationsData) UnSetIpRange() {
	o.IpRange.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o GeoLocationsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.StartIp != nil {
		toSerialize["startip"] = o.StartIp
	}
	if o.EndIp != nil {
		toSerialize["endip"] = o.EndIp
	}
	if o.CountryCode != nil {
		toSerialize["countrycode"] = o.CountryCode
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	if o.IpRange.IsSet() {
		toSerialize["iprange"] = o.IpRange.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *GeoLocationsData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string               `json:"id,omitempty"`
		StartIp     *string               `json:"startip,omitempty"`
		EndIp       *string               `json:"endip,omitempty"`
		CountryCode *string               `json:"countrycode,omitempty"`
		Country     *string               `json:"country,omitempty"`
		City        *string               `json:"city,omitempty"`
		Region      *string               `json:"region,omitempty"`
		Latitude    *float64              `json:"latitude,omitempty"`
		Longitude   *float64              `json:"longitude,omitempty"`
		IpRange     common.NullableString `json:"iprange,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "startip", "endip", "countrycode", "country", "city", "region", "latitude", "iprange"})
	} else {
		return err
	}

	o.Id = all.Id
	o.StartIp = all.StartIp
	o.EndIp = all.EndIp
	o.CountryCode = all.CountryCode
	o.Country = all.Country
	o.City = all.City
	o.Region = all.Region
	o.Latitude = all.Latitude
	o.Longitude = all.Longitude
	o.IpRange = all.IpRange

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
