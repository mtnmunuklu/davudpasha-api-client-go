package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// GeoLocationsSearchRequest represents a search request structure for geographical locations.
type GeoLocationsSearchRequest struct {
	// Query filter for searching geographical locations.
	Query *QueryFilter `json:"query,omitempty"`
	// Filter for searching sources.
	Filter *string `json:"filter,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGeoLocationsSearchRequest creates a new GeoLocationsSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeoLocationsSearchRequest() *GeoLocationsSearchRequest {
	this := GeoLocationsSearchRequest{}
	return &this
}

// NewGeoLocationsSearchRequestWithDefaults creates a new GeoLocationsSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeoLocationsSearchRequestWithDefaults() *GeoLocationsSearchRequest {
	this := GeoLocationsSearchRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *GeoLocationsSearchRequest) GetQuery() QueryFilter {
	if o == nil || o.Query == nil {
		var ret QueryFilter
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsSearchRequest) GetQueryOk() (*QueryFilter, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *GeoLocationsSearchRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given QueryFilter and assigns it to the Query field.
func (o *GeoLocationsSearchRequest) SetQuery(v QueryFilter) {
	o.Query = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GeoLocationsSearchRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsSearchRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *GeoLocationsSearchRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *GeoLocationsSearchRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *GeoLocationsSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLocationsSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *GeoLocationsSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *GeoLocationsSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeoLocationsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.SmartRestRequestContext != nil {
		toSerialize["smartRestRequestContext"] = o.SmartRestRequestContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *GeoLocationsSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query *QueryFilter `json:"query,omitempty"`
		// Filter for searching sources.
		Filter *string `json:"filter,omitempty"`
		// Context for the Smart REST request.
		SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"query", "filter", "smartRestRequestContext"})
	} else {
		return err
	}

	o.Query = all.Query
	o.Filter = all.Filter
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
