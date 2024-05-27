package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsQueryData represents the data for a report query.
type ReportsQueryData struct {
	// Type of the item.
	ItemType *string `json:"ItemType,omitempty"`
	// ID of the query.
	QueryID common.NullableString `json:"QueryID,omitempty"`
	// Query string.
	QueryStr common.NullableString `json:"QueryStr,omitempty"`
	// Code related to the query.
	Code *string `json:"Code,omitempty"`
	// Maximum row count for the query data.
	MaxRowCount *int64 `json:"MaxRowCount,omitempty"`
	// Date-time range for the query.
	DateTimeRange *DateTimeRange `json:"DateTimeRange,omitempty"`
	// Path to the script related to the query.
	ScriptPath common.NullableString `json:"ScriptPath,omitempty"`
	// Arguments for the script.
	ScriptArguments common.NullableList[string] `json:"ScriptArguments,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewReportsQueryData creates a new ReportsQueryData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsQueryData() *ReportsQueryData {
	this := ReportsQueryData{}
	return &this
}

// NewReportsQueryDataWithDefaults creates a new ReportsQueryData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsQueryDataWithDefaults() *ReportsQueryData {
	this := ReportsQueryData{}
	return &this
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *ReportsQueryData) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryData) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *ReportsQueryData) HasItemType() bool {
	return o != nil && o.ItemType != nil
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *ReportsQueryData) SetItemType(v string) {
	o.ItemType = &v
}

// GetQueryID returns the QueryID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsQueryData) GetQueryID() string {
	if o == nil || o.QueryID.Get() == nil {
		var ret string
		return ret
	}
	return *o.QueryID.Get()
}

// GetQueryIDOk returns a tuple with the QueryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryData) GetQueryIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryID.Get(), o.QueryID.IsSet()
}

// HasQueryID returns a boolean if a field has been set.
func (o *ReportsQueryData) HasQueryID() bool {
	return o != nil && o.QueryID.IsSet()
}

// SetQueryID gets a reference to the given datadog.NullableString and assigns it to the QueryID field.
func (o *ReportsQueryData) SetQueryID(v string) {
	o.QueryID.Set(&v)
}

// SetQueryIDNil sets the value for QueryID to be an explicit nil.
func (o *ReportsQueryData) SetQueryIDNil() {
	o.QueryID.Set(nil)
}

// UnsetQueryID ensures that no value is present for QueryID, not even an explicit nil.
func (o *ReportsQueryData) UnSetQueryID() {
	o.QueryID.UnSet()
}

// GetQueryStr returns the QueryStr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsQueryData) GetQueryStr() string {
	if o == nil || o.QueryStr.Get() == nil {
		var ret string
		return ret
	}
	return *o.QueryStr.Get()
}

// GetQueryStrOk returns a tuple with the QueryStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryData) GetQueryStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryStr.Get(), o.QueryStr.IsSet()
}

// HasQueryStr returns a boolean if a field has been set.
func (o *ReportsQueryData) HasQueryStr() bool {
	return o != nil && o.QueryStr.IsSet()
}

// SetQueryStr gets a reference to the given datadog.NullableString and assigns it to the QueryStr field.
func (o *ReportsQueryData) SetQueryStr(v string) {
	o.QueryStr.Set(&v)
}

// SetQueryStrNil sets the value for QueryStr to be an explicit nil.
func (o *ReportsQueryData) SetQueryStrNil() {
	o.QueryStr.Set(nil)
}

// UnsetQueryStr ensures that no value is present for QueryStr, not even an explicit nil.
func (o *ReportsQueryData) UnSetQueryStr() {
	o.QueryStr.UnSet()
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ReportsQueryData) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryData) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ReportsQueryData) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ReportsQueryData) SetCode(v string) {
	o.Code = &v
}

// GetMaxRowCount returns the MaxRowCount field value if set, zero value otherwise.
func (o *ReportsQueryData) GetMaxRowCount() int64 {
	if o == nil || o.MaxRowCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxRowCount
}

// GetMaxRowCountOk returns a tuple with the MaxRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryData) GetMaxRowCountOk() (*int64, bool) {
	if o == nil || o.MaxRowCount == nil {
		return nil, false
	}
	return o.MaxRowCount, true
}

// HasMaxRowCount returns a boolean if a field has been set.
func (o *ReportsQueryData) HasMaxRowCount() bool {
	return o != nil && o.MaxRowCount != nil
}

// SetMaxRowCount gets a reference to the given int64 and assigns it to the MaxRowCount field.
func (o *ReportsQueryData) SetMaxRowCount(v int64) {
	o.MaxRowCount = &v
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *ReportsQueryData) GetDateTimeRange() DateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret DateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQueryData) GetDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *ReportsQueryData) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DateTimeRange field.
func (o *ReportsQueryData) SetDateTimeRange(v DateTimeRange) {
	o.DateTimeRange = &v
}

// GetScriptPath returns the ScriptPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsQueryData) GetScriptPath() string {
	if o == nil || o.ScriptPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.ScriptPath.Get()
}

// GetScriptPathOk returns a tuple with the ScriptPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryData) GetScriptPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScriptPath.Get(), o.ScriptPath.IsSet()
}

// HasScriptPath returns a boolean if a field has been set.
func (o *ReportsQueryData) HasScriptPath() bool {
	return o != nil && o.ScriptPath.IsSet()
}

// SetScriptPath gets a reference to the given datadog.NullableString and assigns it to the ScriptPath field.
func (o *ReportsQueryData) SetScriptPath(v string) {
	o.ScriptPath.Set(&v)
}

// SetScriptPathNil sets the value for ScriptPath to be an explicit nil.
func (o *ReportsQueryData) SetScriptPathNil() {
	o.ScriptPath.Set(nil)
}

// UnsetScriptPath ensures that no value is present for ScriptPath, not even an explicit nil.
func (o *ReportsQueryData) UnSetScriptPath() {
	o.ScriptPath.UnSet()
}

// GetScriptArguments returns a tuple with the ScriptArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryData) GetScriptArguments() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScriptArguments.Get(), o.ScriptArguments.IsSet()
}

// GetScriptArgumentsOk returns a tuple with the ScriptArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsQueryData) GetScriptArgumentsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScriptArguments.Get(), o.ScriptArguments.IsSet()
}

// HasScriptArguments returns a boolean if a ScriptArguments has been set.
func (o *ReportsQueryData) HasScriptArguments() bool {
	return o != nil && o.ScriptArguments.IsSet()
}

// SetScriptArguments gets a reference to the given datadog.Nullable[]string and assigns it to the ScriptArguments field.
func (o *ReportsQueryData) SetScriptArguments(v []string) {
	o.ScriptArguments.Set(&v)
}

// SetScriptArgumentsNil sets the value for ScriptArguments to be an explicit nil.
func (o *ReportsQueryData) SetScriptArgumentsNil() {
	o.ScriptArguments.Set(nil)
}

// UnsetScriptArguments ensures that no value is present for ScriptArguments, not even an explicit nil.
func (o *ReportsQueryData) UnSetScriptArguments() {
	o.ScriptArguments.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsQueryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ItemType != nil {
		toSerialize["ItemType"] = o.ItemType
	}
	if o.QueryID.IsSet() {
		toSerialize["QueryID"] = o.QueryID.Get()
	}
	if o.QueryStr.IsSet() {
		toSerialize["QueryStr"] = o.QueryStr.Get()
	}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.MaxRowCount != nil {
		toSerialize["MaxRowCount"] = o.MaxRowCount
	}
	if o.DateTimeRange != nil {
		toSerialize["DateTimeRange"] = o.DateTimeRange
	}
	if o.ScriptPath.IsSet() {
		toSerialize["ScriptPath"] = o.ScriptPath.Get()
	}
	if o.ScriptArguments.IsSet() {
		toSerialize["ScriptArguments"] = o.ScriptArguments.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsQueryData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ItemType        *string                     `json:"ItemType,omitempty"`
		QueryID         common.NullableString       `json:"QueryID,omitempty"`
		QueryStr        common.NullableString       `json:"QueryStr,omitempty"`
		Code            *string                     `json:"Code,omitempty"`
		MaxRowCount     *int64                      `json:"MaxRowCount,omitempty"`
		DateTimeRange   *DateTimeRange              `json:"DateTimeRange,omitempty"`
		ScriptPath      common.NullableString       `json:"ScriptPath,omitempty"`
		ScriptArguments common.NullableList[string] `json:"ScriptArguments,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ItemType", "QueryID", "QueryStr", "Code", "MaxRowCount", "DateTimeRange", "ScriptPath", "ScriptArguments"})
	} else {
		return err
	}

	o.ItemType = all.ItemType
	o.QueryID = all.QueryID
	o.QueryStr = all.QueryStr
	o.Code = all.Code
	o.MaxRowCount = all.MaxRowCount
	hasInvalidField := false
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	o.ScriptPath = all.ScriptPath
	o.ScriptArguments = all.ScriptArguments

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
