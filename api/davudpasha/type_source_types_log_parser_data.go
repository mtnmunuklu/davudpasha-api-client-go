package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SourceTypesLogParserData struct {
	Delimiter              *string               `json:"Delimiter,omitempty"`
	MustContain            common.NullableString `json:"MustContain,omitempty"`
	MustContainParts       common.NullableInt64  `json:"MustContainParts,omitempty"`
	Query                  *string               `json:"Query,omitempty"`
	IDColumn               *string               `json:"IDColumn,omitempty"`
	IDColumnType           *string               `json:"IDColumnType,omitempty"`
	TableName              *string               `json:"TableName,omitempty"`
	Limit                  common.NullableInt64  `json:"Limit,omitempty"`
	LimitID                common.NullableString `json:"LimitID,omitempty"`
	IDDateTimeCustomFormat common.NullableString `json:"IDDateTimeCustomFormat,omitempty"`
	CodeStr                *string               `json:"CodeStr,omitempty"`
	RegexStr               *string               `json:"RegexStr,omitempty"`
	ExpressionFieldLengths *string               `json:"ExpressionFieldLengths,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSourceTypesLogParserData creates a new SourceTypesLogParserData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesLogParserData() *SourceTypesLogParserData {
	this := SourceTypesLogParserData{}
	return &this
}

// NewSourceTypesLogParserDataWithDefaults creates a new SourceTypesLogParserData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourceTypesLogParserDataWithDefaults() *SourceTypesLogParserData {
	this := SourceTypesLogParserData{}
	return &this
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise.
func (o *SourceTypesLogParserData) GetDelimiter() string {
	if o == nil || o.Delimiter == nil {
		var ret string
		return ret
	}
	return *o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesLogParserData) GetDelimiterOk() (*string, bool) {
	if o == nil || o.Delimiter == nil {
		return nil, false
	}
	return o.Delimiter, true
}

// HasDelimiter returns a boolean if a field has been set.
func (o *SourceTypesLogParserData) HasDelimiter() bool {
	return o != nil && o.Delimiter != nil
}

// SetDelimiter gets a reference to the given string and assigns it to the Delimiter field.
func (o *SourceTypesLogParserData) SetDelimiter(v string) {
	o.Delimiter = &v
}

// GetMustContain returns the MustContain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesLogParserData) GetMustContain() string {
	if o == nil || o.MustContain.Get() == nil {
		var ret string
		return ret
	}
	return *o.MustContain.Get()
}

// GetMustContainOk returns a tuple with the MustContain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesLogParserData) GetMustContainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MustContain.Get(), o.MustContain.IsSet()
}

// HasMustContain returns a boolean if a MustContain has been set.
func (o *SourceTypesLogParserData) HasMustContain() bool {
	return o != nil && o.MustContain.IsSet()
}

// SetMustContain gets a reference to the given datadog.NullableString and assigns it to the MustContain field.
func (o *SourceTypesLogParserData) SetMustContain(v string) {
	o.MustContain.Set(&v)
}

// SetMustContainNil sets the value for MustContain to be an explicit nil.
func (o *SourceTypesLogParserData) SetMustContainNil() {
	o.MustContain.Set(nil)
}

// UnSetMustContain ensures that no value is present for MustContain, not even an explicit nil.
func (o *SourceTypesLogParserData) UnSetMustContain() {
	o.MustContain.UnSet()
}

// GetMustContainParts returns the MustContainParts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesLogParserData) GetMustContainParts() int64 {
	if o == nil || o.MustContainParts.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MustContainParts.Get()
}

// GetMustContainPartsOk returns a tuple with the MustContainParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesLogParserData) GetMustContainPartsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MustContainParts.Get(), o.MustContainParts.IsSet()
}

// HasMustContainParts returns a boolean if a MustContainParts has been set.
func (o *SourceTypesLogParserData) HasMustContainParts() bool {
	return o != nil && o.MustContainParts.IsSet()
}

// SetMustContainParts gets a reference to the given datadog.Nullableint64 and assigns it to the MustContainParts field.
func (o *SourceTypesLogParserData) SetMustContainParts(v int64) {
	o.MustContainParts.Set(&v)
}

// SetMustContainPartsNil sets the value for MustContainParts to be an explicit nil.
func (o *SourceTypesLogParserData) SetMustContainPartsNil() {
	o.MustContainParts.Set(nil)
}

// UnSetMustContainParts ensures that no value is present for MustContainParts, not even an explicit nil.
func (o *SourceTypesLogParserData) UnSetMustContainParts() {
	o.MustContainParts.UnSet()
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SourceTypesLogParserData) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesLogParserData) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SourceTypesLogParserData) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SourceTypesLogParserData) SetQuery(v string) {
	o.Query = &v
}

// GetIDColumn returns the IDColumn field value if set, zero value otherwise.
func (o *SourceTypesLogParserData) GetIDColumn() string {
	if o == nil || o.IDColumn == nil {
		var ret string
		return ret
	}
	return *o.IDColumn
}

// GetIDColumnOk returns a tuple with the IDColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesLogParserData) GetIDColumnOk() (*string, bool) {
	if o == nil || o.IDColumn == nil {
		return nil, false
	}
	return o.IDColumn, true
}

// HasIDColumn returns a boolean if a field has been set.
func (o *SourceTypesLogParserData) HasIDColumn() bool {
	return o != nil && o.IDColumn != nil
}

// SetIDColumn gets a reference to the given string and assigns it to the IDColumn field.
func (o *SourceTypesLogParserData) SetIDColumn(v string) {
	o.IDColumn = &v
}

// GetIDColumnType returns the IDColumnType field value if set, zero value otherwise.
func (o *SourceTypesLogParserData) GetIDColumnType() string {
	if o == nil || o.IDColumnType == nil {
		var ret string
		return ret
	}
	return *o.IDColumnType
}

// GetIDColumnTypeOk returns a tuple with the IDColumnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesLogParserData) GetIDColumnTypeOk() (*string, bool) {
	if o == nil || o.IDColumnType == nil {
		return nil, false
	}
	return o.IDColumnType, true
}

// HasIDColumnType returns a boolean if a field has been set.
func (o *SourceTypesLogParserData) HasIDColumnType() bool {
	return o != nil && o.IDColumnType != nil
}

// SetIDColumnType gets a reference to the given string and assigns it to the IDColumnType field.
func (o *SourceTypesLogParserData) SetIDColumnType(v string) {
	o.IDColumnType = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *SourceTypesLogParserData) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesLogParserData) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *SourceTypesLogParserData) HasTableName() bool {
	return o != nil && o.TableName != nil
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *SourceTypesLogParserData) SetTableName(v string) {
	o.TableName = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesLogParserData) GetLimit() int64 {
	if o == nil || o.Limit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesLogParserData) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a Limit has been set.
func (o *SourceTypesLogParserData) HasLimit() bool {
	return o != nil && o.Limit.IsSet()
}

// SetLimit gets a reference to the given datadog.Nullableint64 and assigns it to the Limit field.
func (o *SourceTypesLogParserData) SetLimit(v int64) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil.
func (o *SourceTypesLogParserData) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnSetLimit ensures that no value is present for Limit, not even an explicit nil.
func (o *SourceTypesLogParserData) UnSetLimit() {
	o.Limit.UnSet()
}

// GetLimitID returns the LimitID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesLogParserData) GetLimitID() string {
	if o == nil || o.LimitID.Get() == nil {
		var ret string
		return ret
	}
	return *o.LimitID.Get()
}

// GetLimitIDOk returns a tuple with the LimitID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesLogParserData) GetLimitIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitID.Get(), o.LimitID.IsSet()
}

// HasLimitID returns a boolean if a LimitID has been set.
func (o *SourceTypesLogParserData) HasLimitID() bool {
	return o != nil && o.LimitID.IsSet()
}

// SetLimitID gets a reference to the given datadog.Nullablestring and assigns it to the LimitID field.
func (o *SourceTypesLogParserData) SetLimitID(v string) {
	o.LimitID.Set(&v)
}

// SetLimitIDNil sets the value for LimitID to be an explicit nil.
func (o *SourceTypesLogParserData) SetLimitIDNil() {
	o.LimitID.Set(nil)
}

// UnSetLimitID ensures that no value is present for LimitID, not even an explicit nil.
func (o *SourceTypesLogParserData) UnSetLimitID() {
	o.LimitID.UnSet()
}

// GetIDDateTimeCustomFormat returns the IDDateTimeCustomFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesLogParserData) GetIDDateTimeCustomFormat() string {
	if o == nil || o.IDDateTimeCustomFormat.Get() == nil {
		var ret string
		return ret
	}
	return *o.IDDateTimeCustomFormat.Get()
}

// GetIDDateTimeCustomFormatOk returns a tuple with the IDDateTimeCustomFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesLogParserData) GetIDDateTimeCustomFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IDDateTimeCustomFormat.Get(), o.IDDateTimeCustomFormat.IsSet()
}

// HasIDDateTimeCustomFormat returns a boolean if a IDDateTimeCustomFormat has been set.
func (o *SourceTypesLogParserData) HasIDDateTimeCustomFormat() bool {
	return o != nil && o.IDDateTimeCustomFormat.IsSet()
}

// SetIDDateTimeCustomFormat gets a reference to the given datadog.Nullablestring and assigns it to the IDDateTimeCustomFormat field.
func (o *SourceTypesLogParserData) SetIDDateTimeCustomFormat(v string) {
	o.IDDateTimeCustomFormat.Set(&v)
}

// SetIDDateTimeCustomFormatNil sets the value for IDDateTimeCustomFormat to be an explicit nil.
func (o *SourceTypesLogParserData) SetIDDateTimeCustomFormatNil() {
	o.IDDateTimeCustomFormat.Set(nil)
}

// UnSetIDDateTimeCustomFormat ensures that no value is present for IDDateTimeCustomFormat, not even an explicit nil.
func (o *SourceTypesLogParserData) UnSetIDDateTimeCustomFormat() {
	o.IDDateTimeCustomFormat.UnSet()
}

// GetCodeStr returns the CodeStr field value if set, zero value otherwise.
func (o *SourceTypesLogParserData) GetCodeStr() string {
	if o == nil || o.CodeStr == nil {
		var ret string
		return ret
	}
	return *o.CodeStr
}

// GetCodeStrOk returns a tuple with the CodeStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesLogParserData) GetCodeStrOk() (*string, bool) {
	if o == nil || o.CodeStr == nil {
		return nil, false
	}
	return o.CodeStr, true
}

// HasCodeStr returns a boolean if a field has been set.
func (o *SourceTypesLogParserData) HasCodeStr() bool {
	return o != nil && o.CodeStr != nil
}

// SetCodeStr gets a reference to the given string and assigns it to the CodeStr field.
func (o *SourceTypesLogParserData) SetCodeStr(v string) {
	o.CodeStr = &v
}

// GetRegexStr returns the RegexStr field value if set, zero value otherwise.
func (o *SourceTypesLogParserData) GetRegexStr() string {
	if o == nil || o.RegexStr == nil {
		var ret string
		return ret
	}
	return *o.RegexStr
}

// GetRegexStrOk returns a tuple with the RegexStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesLogParserData) GetRegexStrOk() (*string, bool) {
	if o == nil || o.RegexStr == nil {
		return nil, false
	}
	return o.RegexStr, true
}

// HasRegexStr returns a boolean if a field has been set.
func (o *SourceTypesLogParserData) HasRegexStr() bool {
	return o != nil && o.RegexStr != nil
}

// SetRegexStr gets a reference to the given string and assigns it to the RegexStr field.
func (o *SourceTypesLogParserData) SetRegexStr(v string) {
	o.RegexStr = &v
}

// GetExpressionFieldLengths returns the ExpressionFieldLengths field value if set, zero value otherwise.
func (o *SourceTypesLogParserData) GetExpressionFieldLengths() string {
	if o == nil || o.ExpressionFieldLengths == nil {
		var ret string
		return ret
	}
	return *o.ExpressionFieldLengths
}

// GetExpressionFieldLengthsOk returns a tuple with the ExpressionFieldLengths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesLogParserData) GetExpressionFieldLengthsOk() (*string, bool) {
	if o == nil || o.ExpressionFieldLengths == nil {
		return nil, false
	}
	return o.ExpressionFieldLengths, true
}

// HasExpressionFieldLengths returns a boolean if a field has been set.
func (o *SourceTypesLogParserData) HasExpressionFieldLengths() bool {
	return o != nil && o.ExpressionFieldLengths != nil
}

// SetExpressionFieldLengths gets a reference to the given string and assigns it to the ExpressionFieldLengths field.
func (o *SourceTypesLogParserData) SetExpressionFieldLengths(v string) {
	o.ExpressionFieldLengths = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesLogParserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Delimiter != nil {
		toSerialize["Delimiter"] = o.Delimiter
	}
	if o.MustContain.IsSet() {
		toSerialize["MustContain"] = o.MustContain.Get()
	}
	if o.MustContainParts.IsSet() {
		toSerialize["MustContainParts"] = o.MustContainParts.Get()
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}
	if o.IDColumn != nil {
		toSerialize["IDColumn"] = o.IDColumn
	}
	if o.IDColumnType != nil {
		toSerialize["IDColumnType"] = o.IDColumnType
	}
	if o.TableName != nil {
		toSerialize["TableName"] = o.TableName
	}
	if o.Limit.IsSet() {
		toSerialize["Limit"] = o.Limit.Get()
	}
	if o.LimitID.IsSet() {
		toSerialize["LimitID"] = o.LimitID.Get()
	}
	if o.IDDateTimeCustomFormat.IsSet() {
		toSerialize["IDDateTimeCustomFormat"] = o.IDDateTimeCustomFormat.Get()
	}
	if o.CodeStr != nil {
		toSerialize["CodeStr"] = o.CodeStr
	}
	if o.RegexStr != nil {
		toSerialize["RegexStr"] = o.RegexStr
	}
	if o.ExpressionFieldLengths != nil {
		toSerialize["ExpressionFieldLengths"] = o.ExpressionFieldLengths
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourceTypesLogParserData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Delimiter              *string               `json:"Delimiter,omitempty"`
		MustContain            common.NullableString `json:"MustContain,omitempty"`
		MustContainParts       common.NullableInt64  `json:"MustContainParts,omitempty"`
		Query                  *string               `json:"Query,omitempty"`
		IDColumn               *string               `json:"IDColumn,omitempty"`
		IDColumnType           *string               `json:"IDColumnType,omitempty"`
		TableName              *string               `json:"TableName,omitempty"`
		Limit                  common.NullableInt64  `json:"Limit,omitempty"`
		LimitID                common.NullableString `json:"LimitID,omitempty"`
		IDDateTimeCustomFormat common.NullableString `json:"IDDateTimeCustomFormat,omitempty"`
		CodeStr                *string               `json:"CodeStr,omitempty"`
		RegexStr               *string               `json:"RegexStr,omitempty"`
		ExpressionFieldLengths *string               `json:"ExpressionFieldLengths,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Delimiter", "MustContain", "MustContainParts", "Query", "IDColumn", "IDColumnType", "TableName", "Limit", "LimitID", "IDDateTimeCustomFormat", "CodeStr", "RegexStr", "ExpressionFieldLengths"})
	} else {
		return err
	}

	o.Delimiter = all.Delimiter
	o.MustContain = all.MustContain
	o.MustContainParts = all.MustContainParts
	o.Query = all.Query
	o.IDColumn = all.IDColumn
	o.IDColumnType = all.IDColumnType
	o.TableName = all.TableName
	o.Limit = all.Limit
	o.LimitID = all.LimitID
	o.IDDateTimeCustomFormat = all.IDDateTimeCustomFormat
	o.CodeStr = all.CodeStr
	o.RegexStr = all.RegexStr
	o.ExpressionFieldLengths = all.ExpressionFieldLengths

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
