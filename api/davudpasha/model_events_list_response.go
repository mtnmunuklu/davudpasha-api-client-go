package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// EventsListResponse is the response object with all events matching the request.
type EventsListResponse struct {
	// Data contains the list of events.
	Data common.NullableList[json.RawMessage] `json:"Data,omitempty"`
	// TotalSize is the total number of events.
	TotalSize *int64 `json:"TotalSize,omitempty"`
	// QueryType is the type of query used.
	QueryType *string `json:"QueryType,omitempty"`
	// SelectedColumns contains the selected columns.
	SelectedColumns common.NullableList[json.RawMessage] `json:"SelectedColumns,omitempty"`
	// SearchTime is the time taken for the search.
	SearchTime *int64 `json:"SearchTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewEventsListResponse creates a new EventsListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsListResponse() *EventsListResponse {
	this := EventsListResponse{}
	return &this
}

// NewEventsListResponseWithDefaults creates a new EventsListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsListResponseWithDefault() *EventsListResponse {
	this := EventsListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventsListResponse) GetData() []json.RawMessage {
	if o == nil || o.Data.Get() == nil {
		var ret []json.RawMessage
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsListResponse) GetDataOk() (*[]json.RawMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *EventsListResponse) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given []json.RawMessage and assigns it to the Data field.
func (o *EventsListResponse) SetData(v []json.RawMessage) {
	o.Data.Set(&v)
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *EventsListResponse) GetTotalSize() int64 {
	if o == nil || o.TotalSize == nil {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsListResponse) GetTotalSizeOk() (*int64, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *EventsListResponse) HasTotalSize() bool {
	return o != nil && o.TotalSize != nil
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *EventsListResponse) SetTotalSize(v int64) {
	o.TotalSize = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *EventsListResponse) GetQueryType() string {
	if o == nil || o.QueryType == nil {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsListResponse) GetQueryTypeOk() (*string, bool) {
	if o == nil || o.QueryType == nil {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *EventsListResponse) HasQueryType() bool {
	return o != nil && o.QueryType != nil
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *EventsListResponse) SetQueryType(v string) {
	o.QueryType = &v
}

// GetSelectedColumns returns the SelectedColumns field value if set, zero value otherwise.
func (o *EventsListResponse) GetSelectedColumns() []json.RawMessage {
	if o == nil || o.SelectedColumns.Get() == nil {
		var ret []json.RawMessage
		return ret
	}
	return *o.SelectedColumns.Get()
}

// GetSelectedColumnsOk returns a tuple with the SelectedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsListResponse) GetSelectedColumnsOk() (*[]json.RawMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedColumns.Get(), o.SelectedColumns.IsSet()
}

// HasSelectedColumns returns a boolean if a field has been set.
func (o *EventsListResponse) HasSelectedColumns() bool {
	return o != nil && o.SelectedColumns.IsSet()
}

// SetSelectedColumns gets a reference to the given []json.RawMessage and assigns it to the SelectedColumns field.
func (o *EventsListResponse) SetSelectedColumns(v []json.RawMessage) {
	o.SelectedColumns.Set(&v)
}

// GetSearchTime returns the SearchTime field value if set, zero value otherwise.
func (o *EventsListResponse) GetSearchTime() int64 {
	if o == nil || o.SearchTime == nil {
		var ret int64
		return ret
	}
	return *o.SearchTime
}

// GetSearchTimeOk returns a tuple with the SearchTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsListResponse) GetSearchTimeOk() (*int64, bool) {
	if o == nil || o.SearchTime == nil {
		return nil, false
	}
	return o.SearchTime, true
}

// HasSearchTime returns a boolean if a field has been set.
func (o *EventsListResponse) HasSearchTime() bool {
	return o != nil && o.SearchTime != nil
}

// SetSearchTime gets a reference to the given int64 and assigns it to the SearchTime field.
func (o *EventsListResponse) SetSearchTime(v int64) {
	o.SearchTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data
	}
	if o.TotalSize != nil {
		toSerialize["TotalSize"] = o.TotalSize
	}
	if o.QueryType != nil {
		toSerialize["QueryType"] = o.QueryType
	}
	if o.SelectedColumns.IsSet() {
		toSerialize["SelectedColumns"] = o.SelectedColumns
	}
	if o.SearchTime != nil {
		toSerialize["SearchTime"] = o.SearchTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *EventsListResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data            common.NullableList[json.RawMessage] `json:"Data,omitempty"`
		TotalSize       *int64                               `json:"TotalSize,omitempty"`
		QueryType       *string                              `json:"QueryType,omitempty"`
		QueryFlag       *string                              `json:"QueryFlag,omitempty"`
		SelectedColumns common.NullableList[json.RawMessage] `json:"SelectedColumns,omitempty"`
		SearchTime      *int64                               `json:"SearchTime,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Data", "TotalSize", "QueryType", "QueryFlag", "SelectedColumns", "SearchTime"})
	} else {
		return err
	}

	o.Data = all.Data
	o.TotalSize = all.TotalSize
	o.QueryType = all.QueryType
	o.SelectedColumns = all.SelectedColumns
	o.SearchTime = all.SearchTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
