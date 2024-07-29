package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// QueryFilter represents the filter used for querying events.
type QueryFilter struct {
	// QuerySQL is the SQL query for filtering events.
	QuerySQL *string `json:"QuerySQL,omitempty"`
	// DateTimeRange specifies the date time range for filtering events.
	DateTimeRange *DateTimeRange `json:"DateTimeRange,omitempty"`
	// QueryOptions stores additional query options.
	QueryOptions *QueryOptions `json:"QueryOptions,omitempty"`
	// SearchAfterKeys specifies the keys to use for pagination in search results.
	SearchAfterKeys []string `json:"SearchAfterKeys,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewQueryFilter creates a new QueryFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryFilter() *QueryFilter {
	this := QueryFilter{}
	var querySQL string = "query"
	this.QuerySQL = &querySQL
	return &this
}

// NewQueryFilterWithDefaults creates a new QueryFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryFilterWithDefaults() *QueryFilter {
	this := QueryFilter{}
	var querySQL string = "query"
	this.QuerySQL = &querySQL
	return &this
}

// GetQuerySQL returns the QuerySQL field value if set, zero value otherwise.
func (o *QueryFilter) GetQuerySQL() string {
	if o == nil || o.QuerySQL == nil {
		var ret string
		return ret
	}
	return *o.QuerySQL
}

// GetQuerySQLOk returns a tuple with the QuerySQL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilter) GetQuerySQLOk() (*string, bool) {
	if o == nil || o.QuerySQL == nil {
		return nil, false
	}
	return o.QuerySQL, true
}

// HasQuerySQL returns a boolean if a field has been set.
func (o *QueryFilter) HasQuerySQL() bool {
	return o != nil && o.QuerySQL != nil
}

// SetQuerySQL gets a reference to the given string and assigns it to the QuerySQL field.
func (o *QueryFilter) SetQuerySQL(v string) {
	o.QuerySQL = &v
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *QueryFilter) GetDateTimeRange() DateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret DateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilter) GetDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *QueryFilter) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DateTimeRange field.
func (o *QueryFilter) SetDateTimeRange(v DateTimeRange) {
	o.DateTimeRange = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *QueryFilter) GetQueryOptions() QueryOptions {
	if o == nil || o.QueryOptions == nil {
		var ret QueryOptions
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilter) GetQueryOptionsOk() (*QueryOptions, bool) {
	if o == nil || o.QueryOptions == nil {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *QueryFilter) HasQueryOptions() bool {
	return o != nil && o.QueryOptions != nil
}

// SetQueryOptions gets a reference to the given QueryOptions and assigns it to the QueryOptions field.
func (o *QueryFilter) SetQueryOptions(v QueryOptions) {
	o.QueryOptions = &v
}

// GetSearchAfterKeys returns the SearchAfterKeys field value if set, zero value otherwise.
func (o *QueryFilter) GetSearchAfterKeys() []string {
	if o == nil || o.SearchAfterKeys == nil {
		var ret []string
		return ret
	}
	return o.SearchAfterKeys
}

// GetSearchAfterKeysOk returns a tuple with the SearchAfterKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilter) GetSearchAfterKeysOk() (*[]string, bool) {
	if o == nil || o.SearchAfterKeys == nil {
		return nil, false
	}
	return &o.SearchAfterKeys, true
}

// HasSearchAfterKeys returns a boolean if a field has been set.
func (o *QueryFilter) HasSearchAfterKeys() bool {
	return o != nil && o.SearchAfterKeys != nil
}

// SetSearchAfterKeys gets a reference to the given []string and assigns it to the SearchAfterKeys field.
func (o *QueryFilter) SetSearchAfterKeys(v []string) {
	o.SearchAfterKeys = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.QuerySQL != nil {
		toSerialize["QuerySQL"] = o.QuerySQL
	}
	if o.DateTimeRange != nil {
		toSerialize["DateTimeRange"] = o.DateTimeRange
	}
	if o.QueryOptions != nil {
		toSerialize["QueryOptions"] = o.QueryOptions
	}
	if o.SearchAfterKeys != nil {
		toSerialize["SearchAfterKeys"] = o.SearchAfterKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *QueryFilter) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		QuerySQL        *string        `json:"QuerySQL,omitempty"`
		DateTimeRange   *DateTimeRange `json:"DateTimeRange,omitempty"`
		QueryOptions    *QueryOptions  `json:"QueryOptions,omitempty"`
		SearchAfterKeys []string       `json:"SearchAfterKeys,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.QuerySQL == nil {
		return fmt.Errorf("requiered field QuerySQL is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"QuerySQL", "DateTimeRange", "QueryOptions", "SearchAfterKeys"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	if all.QueryOptions != nil && all.QueryOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryOptions = all.QueryOptions
	o.QuerySQL = all.QuerySQL
	o.SearchAfterKeys = all.SearchAfterKeys

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
