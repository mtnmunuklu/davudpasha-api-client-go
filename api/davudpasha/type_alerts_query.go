package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AlertsQuery struct {
	QueryID        *string          `json:"QueryID,omitempty"`
	Query          *string          `json:"Query,omitempty"`
	TimeFrameValue *int64           `json:"TimeFrameValue,omitempty"`
	TimeFrameType  *string          `json:"TimeFrameType,omitempty"`
	QueryColumns   []SelectedColumn `json:"QueryColumns,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertsQuery creates a new AlertsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertsQuery() *AlertsQuery {
	this := AlertsQuery{}
	return &this
}

// NewAlertsQueryWithDefaults creates a new AlertsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertsQueryWithDefaults() *AlertsQuery {
	this := AlertsQuery{}
	return &this
}

// GetQueryID returns the QueryID field value if set, zero value otherwise.
func (o *AlertsQuery) GetQueryID() string {
	if o == nil || o.QueryID == nil {
		var ret string
		return ret
	}
	return *o.QueryID
}

// GetQueryIDOk returns a tuple with the QueryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQuery) GetQueryIDOk() (*string, bool) {
	if o == nil || o.QueryID == nil {
		return nil, false
	}
	return o.QueryID, true
}

// HasQueryID returns a boolean if a field has been set.
func (o *AlertsQuery) HasQueryID() bool {
	return o != nil && o.QueryID != nil
}

// SetQueryID gets a reference to the given string and assigns it to the QueryID field.
func (o *AlertsQuery) SetQueryID(v string) {
	o.QueryID = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *AlertsQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *AlertsQuery) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *AlertsQuery) SetQuery(v string) {
	o.Query = &v
}

// GetTimeFrameValue returns the TimeFrameValue field value if set, zero value otherwise.
func (o *AlertsQuery) GetTimeFrameValue() int64 {
	if o == nil || o.TimeFrameValue == nil {
		var ret int64
		return ret
	}
	return *o.TimeFrameValue
}

// GetTimeFrameValueOk returns a tuple with the TimeFrameValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQuery) GetTimeFrameValueOk() (*int64, bool) {
	if o == nil || o.TimeFrameValue == nil {
		return nil, false
	}
	return o.TimeFrameValue, true
}

// HasTimeFrameValue returns a boolean if a field has been set.
func (o *AlertsQuery) HasTimeFrameValue() bool {
	return o != nil && o.TimeFrameValue != nil
}

// SetTimeFrameValue gets a reference to the given int64 and assigns it to the TimeFrameValue field.
func (o *AlertsQuery) SetTimeFrameValue(v int64) {
	o.TimeFrameValue = &v
}

// GetTimeFrameType returns the TimeFrameType field value if set, zero value otherwise.
func (o *AlertsQuery) GetTimeFrameType() string {
	if o == nil || o.TimeFrameType == nil {
		var ret string
		return ret
	}
	return *o.TimeFrameType
}

// GetTimeFrameTypeOk returns a tuple with the TimeFrameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQuery) GetTimeFrameTypeOk() (*string, bool) {
	if o == nil || o.TimeFrameType == nil {
		return nil, false
	}
	return o.TimeFrameType, true
}

// HasTimeFrameType returns a boolean if a field has been set.
func (o *AlertsQuery) HasTimeFrameType() bool {
	return o != nil && o.TimeFrameType != nil
}

// SetTimeFrameType gets a reference to the given string and assigns it to the TimeFrameType field.
func (o *AlertsQuery) SetTimeFrameType(v string) {
	o.TimeFrameType = &v
}

// GetQueryColumns returns the QueryColumns field value if set, zero value otherwise.
func (o *AlertsQuery) GetQueryColumns() []SelectedColumn {
	if o == nil || o.QueryColumns == nil {
		var ret []SelectedColumn
		return ret
	}
	return o.QueryColumns
}

// GetQueryColumnsOk returns a tuple with the QueryColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQuery) GetQueryColumnsOk() (*[]SelectedColumn, bool) {
	if o == nil || o.QueryColumns == nil {
		return nil, false
	}
	return &o.QueryColumns, true
}

// HasQueryColumns returns a boolean if a field has been set.
func (o *AlertsQuery) HasQueryColumns() bool {
	return o != nil && o.QueryColumns != nil
}

// SetQueryColumns gets a reference to the given []SelectedColumn and assigns it to the QueryColumns field.
func (o *AlertsQuery) SetQueryColumns(v []SelectedColumn) {
	o.QueryColumns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.QueryID != nil {
		toSerialize["QueryID"] = o.QueryID
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}
	if o.TimeFrameValue != nil {
		toSerialize["TimeFrameValue"] = o.TimeFrameValue
	}
	if o.TimeFrameType != nil {
		toSerialize["TimeFrameType"] = o.TimeFrameType
	}
	if o.QueryColumns != nil {
		toSerialize["QueryColumns"] = o.QueryColumns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AlertsQuery) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryID        *string          `json:"QueryID,omitempty"`
		Query          *string          `json:"Query,omitempty"`
		TimeFrameValue *int64           `json:"TimeFrameValue,omitempty"`
		TimeFrameType  *string          `json:"TimeFrameType,omitempty"`
		QueryColumns   []SelectedColumn `json:"QueryColumns,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.QueryID == nil {
		return fmt.Errorf("requiered field QueryID is missing")
	}
	if all.Query == nil {
		return fmt.Errorf("requiered field Query is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"QueryID", "Query", "TimeFrameValue", "TimeFrameType", "QueryColumns"})
	} else {
		return err
	}
	o.QueryID = all.QueryID
	o.Query = all.Query
	o.TimeFrameValue = all.TimeFrameValue
	o.TimeFrameType = all.TimeFrameType
	o.QueryColumns = all.QueryColumns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}