package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AlertsQueryData struct {
	TimeFrameValue            *int64                          `json:"TimeFrameValue,omitempty"`
	TimeFrameType             AlertsTimeFrameType             `json:"TimeFrameType,omitempty"`
	RuleType                  AlertsRuleType                  `json:"RuleType,omitempty"`
	QueryCorrelationAlertType AlertsQueryCorrelationAlertType `json:"QueryCorrelationAlertType,omitempty"`
	QueryID                   *string                         `json:"QueryID,omitempty"`
	LgsID                     *string                         `json:"LgsID,omitempty"`
	Query                     *string                         `json:"Query,omitempty"`
	LastReadTime              *string                         `json:"LastReadTime,omitempty"`
	Queries                   []AlertsQuery                   `json:"Queries,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertsQueryData creates a new AlertsQueryData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertsQueryData() *AlertsQueryData {
	this := AlertsQueryData{}
	return &this
}

// NewAlertsQueryDataWithDefaults creates a new AlertsQueryData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertsQueryDataWithDefaults() *AlertsQueryData {
	this := AlertsQueryData{}
	return &this
}

// GetTimeFrameValue returns the TimeFrameValue field value if set, zero value otherwise.
func (o *AlertsQueryData) GetTimeFrameValue() int64 {
	if o == nil || o.TimeFrameValue == nil {
		var ret int64
		return ret
	}
	return *o.TimeFrameValue
}

// GetTimeFrameValueOk returns a tuple with the TimeFrameValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetTimeFrameValueOk() (*int64, bool) {
	if o == nil || o.TimeFrameValue == nil {
		return nil, false
	}
	return o.TimeFrameValue, true
}

// HasTimeFrameValue returns a boolean if a field has been set.
func (o *AlertsQueryData) HasTimeFrameValue() bool {
	return o != nil && o.TimeFrameValue != nil
}

// SetTimeFrameValue gets a reference to the given int64 and assigns it to the TimeFrameValue field.
func (o *AlertsQueryData) SetTimeFrameValue(v int64) {
	o.TimeFrameValue = &v
}

// GetTimeFrameType returns the TimeFrameType field value if set, zero value otherwise.
func (o *AlertsQueryData) GetTimeFrameType() AlertsTimeFrameType {
	if o == nil {
		var ret AlertsTimeFrameType
		return ret
	}
	return o.TimeFrameType
}

// GetTimeFrameTypeOk returns a tuple with the TimeFrameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetTimeFrameTypeOk() (*AlertsTimeFrameType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeFrameType, true
}

// SetTimeFrameType gets a reference to the given string and assigns it to the TimeFrameType field.
func (o *AlertsQueryData) SetTimeFrameType(v AlertsTimeFrameType) {
	o.TimeFrameType = v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *AlertsQueryData) GetRuleType() AlertsRuleType {
	if o == nil {
		var ret AlertsRuleType
		return ret
	}
	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetRuleTypeOk() (*AlertsRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *AlertsQueryData) SetRuleType(v AlertsRuleType) {
	o.RuleType = v
}

// GetQueryCorrelationAlertType returns the QueryCorrelationAlertType field value if set, zero value otherwise.
func (o *AlertsQueryData) GetQueryCorrelationAlertType() AlertsQueryCorrelationAlertType {
	if o == nil {
		var ret AlertsQueryCorrelationAlertType
		return ret
	}
	return o.QueryCorrelationAlertType
}

// GetQueryCorrelationAlertTypeOk returns a tuple with the QueryCorrelationAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetQueryCorrelationAlertTypeOk() (*AlertsQueryCorrelationAlertType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryCorrelationAlertType, true
}

// SetQueryCorrelationAlertType gets a reference to the given string and assigns it to the QueryCorrelationAlertType field.
func (o *AlertsQueryData) SetQueryCorrelationAlertType(v AlertsQueryCorrelationAlertType) {
	o.QueryCorrelationAlertType = v
}

// GetQueryID returns the QueryID field value if set, zero value otherwise.
func (o *AlertsQueryData) GetQueryID() string {
	if o == nil || o.QueryID == nil {
		var ret string
		return ret
	}
	return *o.QueryID
}

// GetQueryIDOk returns a tuple with the QueryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetQueryIDOk() (*string, bool) {
	if o == nil || o.QueryID == nil {
		return nil, false
	}
	return o.QueryID, true
}

// HasQueryID returns a boolean if a field has been set.
func (o *AlertsQueryData) HasQueryID() bool {
	return o != nil && o.QueryID != nil
}

// SetQueryID gets a reference to the given string and assigns it to the QueryID field.
func (o *AlertsQueryData) SetQueryID(v string) {
	o.QueryID = &v
}

// GetLgsID returns the LgsID field value if set, zero value otherwise.
func (o *AlertsQueryData) GetLgsID() string {
	if o == nil || o.LgsID == nil {
		var ret string
		return ret
	}
	return *o.LgsID
}

// GetLgsIDOk returns a tuple with the LgsID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetLgsIDOk() (*string, bool) {
	if o == nil || o.LgsID == nil {
		return nil, false
	}
	return o.LgsID, true
}

// HasLgsID returns a boolean if a field has been set.
func (o *AlertsQueryData) HasLgsID() bool {
	return o != nil && o.LgsID != nil
}

// SetLgsID gets a reference to the given string and assigns it to the LgsID field.
func (o *AlertsQueryData) SetLgsID(v string) {
	o.LgsID = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *AlertsQueryData) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *AlertsQueryData) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *AlertsQueryData) SetQuery(v string) {
	o.Query = &v
}

// GetLastReadTime returns the LastReadTime field value if set, zero value otherwise.
func (o *AlertsQueryData) GetLastReadTime() string {
	if o == nil || o.LastReadTime == nil {
		var ret string
		return ret
	}
	return *o.LastReadTime
}

// GetLastReadTimeOk returns a tuple with the LastReadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetLastReadTimeOk() (*string, bool) {
	if o == nil || o.LastReadTime == nil {
		return nil, false
	}
	return o.LastReadTime, true
}

// HasLastReadTime returns a boolean if a field has been set.
func (o *AlertsQueryData) HasLastReadTime() bool {
	return o != nil && o.LastReadTime != nil
}

// SetLastReadTime gets a reference to the given string and assigns it to the LastReadTime field.
func (o *AlertsQueryData) SetLastReadTime(v string) {
	o.LastReadTime = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *AlertsQueryData) GetQueries() []AlertsQuery {
	if o == nil || o.Queries == nil {
		var ret []AlertsQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsQueryData) GetQueriesOk() (*[]AlertsQuery, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *AlertsQueryData) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []AlertsQuery and assigns it to the Queries field.
func (o *AlertsQueryData) SetQueries(v []AlertsQuery) {
	o.Queries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertsQueryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.TimeFrameValue != nil {
		toSerialize["TimeFrameValue"] = o.TimeFrameValue
	}
	if o.TimeFrameType.IsValid() {
		toSerialize["TimeFrameType"] = o.TimeFrameType
	}
	if o.RuleType.IsValid() {
		toSerialize["RuleType"] = o.RuleType
	}
	if o.QueryCorrelationAlertType.IsValid() {
		toSerialize["QueryCorrelationAlertType"] = o.QueryCorrelationAlertType
	}
	if o.QueryID != nil {
		toSerialize["QueryID"] = o.QueryID
	}
	if o.LgsID != nil {
		toSerialize["LgsID"] = o.LgsID
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}
	if o.TimeFrameValue != nil {
		toSerialize["TimeFrameValue"] = o.TimeFrameValue
	}
	if o.LastReadTime != nil {
		toSerialize["LastReadTime"] = o.LastReadTime
	}
	if o.Queries != nil {
		toSerialize["Queries"] = o.Queries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AlertsQueryData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimeFrameValue            *int64                           `json:"TimeFrameValue,omitempty"`
		TimeFrameType             *AlertsTimeFrameType             `json:"TimeFrameType,omitempty"`
		RuleType                  *AlertsRuleType                  `json:"RuleType,omitempty"`
		QueryCorrelationAlertType *AlertsQueryCorrelationAlertType `json:"QueryCorrelationAlertType,omitempty"`
		QueryID                   *string                          `json:"QueryID,omitempty"`
		LgsID                     *string                          `json:"LgsID,omitempty"`
		Query                     *string                          `json:"Query,omitempty"`
		LastReadTime              *string                          `json:"LastReadTime,omitempty"`
		Queries                   []AlertsQuery                    `json:"Queries,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TimeFrameValue == nil {
		return fmt.Errorf("requiered field TimeFrameValue is missing")
	}
	if all.TimeFrameType == nil {
		return fmt.Errorf("requiered field TimeFrameType is missing")
	}
	if all.QueryCorrelationAlertType == nil {
		return fmt.Errorf("requiered field QueryCorrelationAlertType is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"TimeFrameValue", "TimeFrameType", "RuleType", "QueryCorrelationAlertType", "QueryID", "LgsID", "Query", "LastReadTime", "Queries"})
	} else {
		return err
	}
	o.TimeFrameValue = all.TimeFrameValue
	hasInvalidField := false
	if !all.TimeFrameType.IsValid() {
		hasInvalidField = true
	} else {
		o.TimeFrameType = *all.TimeFrameType
	}

	if !all.RuleType.IsValid() {
		hasInvalidField = true
	} else {
		o.RuleType = *all.RuleType
	}

	if !all.QueryCorrelationAlertType.IsValid() {
		hasInvalidField = true
	} else {
		o.QueryCorrelationAlertType = *all.QueryCorrelationAlertType
	}
	o.QueryID = all.QueryID
	o.LgsID = all.LgsID
	o.Query = all.Query
	o.LastReadTime = all.LastReadTime
	o.Queries = all.Queries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
