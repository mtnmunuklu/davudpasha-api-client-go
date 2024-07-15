package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// GeoLocationsData represents geographical location data.
type GeoLocationsData struct {
	// Unique identifier for the geographical location data.
	ID *string `json:"id,omitempty"`
	// Start IP address of the geographical range.
	StartIP *string `json:"startip,omitempty"`
	// End IP address of the geographical range.
	EndIP *string `json:"endip,omitempty"`
	// ISO 3166-1 alpha-2 country code.
	CountryCode *string `json:"countrycode,omitempty"`
	// Country name.
	Country *string `json:"country,omitempty"`
	// City name.
	City *string `json:"city,omitempty"`
	// Region or state name.
	Region *string `json:"region,omitempty"`
	// Latitude coordinate of the geographical location.
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude coordinate of the geographical location.
	Longitude *float64 `json:"longitude,omitempty"`
	// IP range associated with the geographical location.
	IPRange common.NullableString `json:"iprange,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
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

// GetID returns the ID field value if set, zero value otherwise.
func (o *GeoLocationsData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *GeoLocationsData) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *GeoLocationsData) SetID(v string) {
	o.ID = &v
}

// GetStartIP returns the StartIP field value if set, zero value otherwise.
func (o *GeoLocationsData) GetStartIP() string {
	if o == nil || o.StartIP == nil {
		var ret string
		return ret
	}
	return *o.StartIP
}

// GetStartIPOk returns a tuple with the StartIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetStartIPOk() (*string, bool) {
	if o == nil || o.StartIP == nil {
		return nil, false
	}
	return o.StartIP, true
}

// HasStartIP returns a boolean if a field has been set.
func (o *GeoLocationsData) HasStartIP() bool {
	return o != nil && o.StartIP != nil
}

// SetStartIP gets a reference to the given string and assigns it to the StartIP field.
func (o *GeoLocationsData) SetStartIP(v string) {
	o.StartIP = &v
}

// GetEndIP returns the EndIP field value if set, zero value otherwise.
func (o *GeoLocationsData) GetEndIP() string {
	if o == nil || o.EndIP == nil {
		var ret string
		return ret
	}
	return *o.EndIP
}

// GetEndIPOk returns a tuple with the EndIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsData) GetEndIPOk() (*string, bool) {
	if o == nil || o.EndIP == nil {
		return nil, false
	}
	return o.EndIP, true
}

// HasEndIP returns a boolean if a field has been set.
func (o *GeoLocationsData) HasEndIP() bool {
	return o != nil && o.EndIP != nil
}

// SetEndIp gets a reference to the given string and assigns it to the EndIp field.
func (o *GeoLocationsData) SetEndIp(v string) {
	o.EndIP = &v
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
	if o == nil || o.IPRange.Get() == nil {
		var ret string
		return ret
	}
	return *o.IPRange.Get()
}

// GetIPRangeOk returns a tuple with the IPRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GeoLocationsData) GetIPRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IPRange.Get(), o.IPRange.IsSet()
}

// HasIPRange returns a boolean if a IPRange has been set.
func (o *GeoLocationsData) HasIPRange() bool {
	return o != nil && o.IPRange.IsSet()
}

// SetIPRange gets a reference to the given common.NullableString and assigns it to the IPRange field.
func (o *GeoLocationsData) SetIPRange(v string) {
	o.IPRange.Set(&v)
}

// SetIPRangeNil sets the value for IPRange to be an explicit nil.
func (o *GeoLocationsData) SetIPRangeNil() {
	o.IPRange.Set(nil)
}

// UnSetIPRange ensures that no value is present for IPRange, not even an explicit nil.
func (o *GeoLocationsData) UnSetIPRange() {
	o.IPRange.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o GeoLocationsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.StartIP != nil {
		toSerialize["startip"] = o.StartIP
	}
	if o.EndIP != nil {
		toSerialize["endip"] = o.EndIP
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
	if o.IPRange.IsSet() {
		toSerialize["iprange"] = o.IPRange.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *GeoLocationsData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID          *string               `json:"id,omitempty"`
		StartIP     *string               `json:"startip,omitempty"`
		EndIP       *string               `json:"endip,omitempty"`
		CountryCode *string               `json:"countrycode,omitempty"`
		Country     *string               `json:"country,omitempty"`
		City        *string               `json:"city,omitempty"`
		Region      *string               `json:"region,omitempty"`
		Latitude    *float64              `json:"latitude,omitempty"`
		Longitude   *float64              `json:"longitude,omitempty"`
		IPRange     common.NullableString `json:"iprange,omitempty"`
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

	o.ID = all.ID
	o.StartIP = all.StartIP
	o.EndIP = all.EndIP
	o.CountryCode = all.CountryCode
	o.Country = all.Country
	o.City = all.City
	o.Region = all.Region
	o.Latitude = all.Latitude
	o.Longitude = all.Longitude
	o.IPRange = all.IPRange

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
