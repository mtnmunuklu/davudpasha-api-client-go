package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// EventsQueryFilter represents the search and filter query settings.
type EventsQueryFilter struct {
	// QuerySQL is the SQL query for filtering events.
	QuerySQL *string `json:"QuerySQL,omitempty"`
	// DateTimeRange specifies the date time range for filtering events.
	DateTimeRange *EventsDateTimeRange `json:"DateTimeRange,omitempty"`
	// QueryOptions stores additional query options.
	QueryOptions *EventsQueryOptions `json:"QueryOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewEventsQueryFilter creates a new EventsQueryFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsQueryFilter() *EventsQueryFilter {
	this := EventsQueryFilter{}
	var querySQL string = "query"
	this.QuerySQL = &querySQL
	return &this
}

// NewEventsQueryFilterWithDefaults creates a new EventsQueryFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsQueryFilterWithDefaults() *EventsQueryFilter {
	this := EventsQueryFilter{}
	var querySQL string = "query"
	this.QuerySQL = &querySQL
	return &this
}

// GetQuerySQL returns the QuerySQL field value if set, zero value otherwise.
func (o *EventsQueryFilter) GetQuerySQL() string {
	if o == nil || o.QuerySQL == nil {
		var ret string
		return ret
	}
	return *o.QuerySQL
}

// GetQuerySQLOk returns a tuple with the QuerySQL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsQueryFilter) GetQuerySQLOk() (*string, bool) {
	if o == nil || o.QuerySQL == nil {
		return nil, false
	}
	return o.QuerySQL, true
}

// HasQuerySQL returns a boolean if a field has been set.
func (o *EventsQueryFilter) HasQuerySQL() bool {
	return o != nil && o.QuerySQL != nil
}

// SetQuerySQL gets a reference to the given string and assigns it to the QuerySQL field.
func (o *EventsQueryFilter) SetQuerySQL(v string) {
	o.QuerySQL = &v
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *EventsQueryFilter) GetDateTimeRange() EventsDateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret EventsDateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsQueryFilter) GetDateTimeRangeOk() (*EventsDateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *EventsQueryFilter) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given EventsDateTimeRange and assigns it to the DateTimeRange field.
func (o *EventsQueryFilter) SetDateTimeRange(v EventsDateTimeRange) {
	o.DateTimeRange = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *EventsQueryFilter) GetQueryOptions() EventsQueryOptions {
	if o == nil || o.QueryOptions == nil {
		var ret EventsQueryOptions
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsQueryFilter) GetQueryOptionsOk() (*EventsQueryOptions, bool) {
	if o == nil || o.QueryOptions == nil {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *EventsQueryFilter) HasQueryOptions() bool {
	return o != nil && o.QueryOptions != nil
}

// SetQueryOptions gets a reference to the given EventsQueryOptions and assigns it to the QueryOptions field.
func (o *EventsQueryFilter) SetQueryOptions(v EventsQueryOptions) {
	o.QueryOptions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsQueryFilter) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *EventsQueryFilter) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		QuerySQL        *string                     `json:"QuerySQL,omitempty"`
		DateTimeRange   *EventsDateTimeRange        `json:"DateTimeRange,omitempty"`
		SearchAfterKeys common.NullableList[string] `json:"SearchAfterKeys,omitempty"`
		QueryOptions    *EventsQueryOptions         `json:"QueryOptions,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.QuerySQL == nil {
		return fmt.Errorf("requiered field QuerySQL is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"QuerySQL", "DateTimeRange", "SearchAfterKeys", "Size", "QueryOptions"})
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
