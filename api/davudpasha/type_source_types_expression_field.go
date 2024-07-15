package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourceTypesExpressionField represents the structure for expression fields in source types.
type SourceTypesExpressionField struct {
	// Name specifies the name of the expression field.
	Name *string `json:"Name,omitempty"`
	// ValueName indicates the value name associated with the expression field.
	ValueName common.NullableString `json:"ValueName,omitempty"`
	// Type specifies the type of the expression field.
	Type *string `json:"Type,omitempty"`
	// DateFormat specifies the date format for the expression field.
	DateFormat common.NullableString `json:"DateFormat,omitempty"`
	// DateLang specifies the date language for the expression field.
	DateLang common.NullableString `json:"DateLang,omitempty"`
	// DateTrim specifies if date trimming is enabled for the expression field.
	DateTrim *bool `json:"DateTrim,omitempty"`
	// StartIndex specifies the starting index for the expression field.
	StartIndex *int64 `json:"StartIndex,omitempty"`
	// EndIndex specifies the ending index for the expression field.
	EndIndex *int64 `json:"EndIndex,omitempty"`
	// Used indicates if the expression field is used.
	Used *bool `json:"Used,omitempty"`
	// Index specifies if indexing is enabled for the expression field.
	Index *bool `json:"Index,omitempty"`
	// Normalization specifies the normalization settings for the expression field.
	Normalization NullableSourceTypesNormalization `json:"Normalization,omitempty"`
	// CorrectTwoDigitDayFormatPosition specifies the position for correcting two-digit day format.
	CorrectTwoDigitDayFormatPosition *int64 `json:"CorrectTwoDigitDayFormat_Position,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSourceTypesExpressionField creates a new SourceTypesExpressionField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesExpressionField() *SourceTypesExpressionField {
	this := SourceTypesExpressionField{}
	return &this
}

// NewSourceTypesExpressionFieldWithDefaults creates a new SourceTypesExpressionField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NeSourceTypesExpressionFieldWithDefaults() *SourceTypesExpressionField {
	this := SourceTypesExpressionField{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceTypesExpressionField) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpressionField) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceTypesExpressionField) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceTypesExpressionField) SetName(v string) {
	o.Name = &v
}

// GetValueName returns the ValueName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesExpressionField) GetValueName() string {
	if o == nil || o.ValueName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ValueName.Get()
}

// GetValueNameOk returns a tuple with the ValueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesExpressionField) GetValueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueName.Get(), o.ValueName.IsSet()
}

// HasValueName returns a boolean if a ValueName has been set.
func (o *SourceTypesExpressionField) HasValueName() bool {
	return o != nil && o.ValueName.IsSet()
}

// SetValueName gets a reference to the given common.NullableString and assigns it to the ValueName field.
func (o *SourceTypesExpressionField) SetValueName(v string) {
	o.ValueName.Set(&v)
}

// SetValueNameNil sets the value for ValueName to be an explicit nil.
func (o *SourceTypesExpressionField) SetValueNameNil() {
	o.ValueName.Set(nil)
}

// UnSetValueName ensures that no value is present for ValueName, not even an explicit nil.
func (o *SourceTypesExpressionField) UnSetValueName() {
	o.ValueName.UnSet()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceTypesExpressionField) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpressionField) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceTypesExpressionField) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceTypesExpressionField) SetType(v string) {
	o.Type = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesExpressionField) GetDateFormat() string {
	if o == nil || o.DateFormat.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateFormat.Get()
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesExpressionField) GetDateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateFormat.Get(), o.DateFormat.IsSet()
}

// HasDateFormat returns a boolean if a DateFormat has been set.
func (o *SourceTypesExpressionField) HasDateFormat() bool {
	return o != nil && o.DateFormat.IsSet()
}

// SetDateFormat gets a reference to the given common.NullableString and assigns it to the DateFormat field.
func (o *SourceTypesExpressionField) SetDateFormat(v string) {
	o.DateFormat.Set(&v)
}

// SetDateFormatNil sets the value for DateFormat to be an explicit nil.
func (o *SourceTypesExpressionField) SetDateFormatNil() {
	o.DateFormat.Set(nil)
}

// UnSetDateFormat ensures that no value is present for DateFormat, not even an explicit nil.
func (o *SourceTypesExpressionField) UnSetDateFormat() {
	o.DateFormat.UnSet()
}

// GetDateLang returns the DateLang field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesExpressionField) GetDateLang() string {
	if o == nil || o.DateLang.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateLang.Get()
}

// GetDateLangOk returns a tuple with the DateLang field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesExpressionField) GetDateLangOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateLang.Get(), o.DateLang.IsSet()
}

// HasDateLang returns a boolean if a DateLang has been set.
func (o *SourceTypesExpressionField) HasDateLang() bool {
	return o != nil && o.DateLang.IsSet()
}

// SetDateLang gets a reference to the given common.NullableString and assigns it to the DateLang field.
func (o *SourceTypesExpressionField) SetDateLang(v string) {
	o.DateLang.Set(&v)
}

// SetDateLangNil sets the value for DateLang to be an explicit nil.
func (o *SourceTypesExpressionField) SetDateLangNil() {
	o.DateLang.Set(nil)
}

// UnSetDateLang ensures that no value is present for DateLang, not even an explicit nil.
func (o *SourceTypesExpressionField) UnSetDateLang() {
	o.DateLang.UnSet()
}

// GetDateTrim returns the DateTrim field value if set, zero value otherwise.
func (o *SourceTypesExpressionField) GetDateTrim() bool {
	if o == nil || o.DateTrim == nil {
		var ret bool
		return ret
	}
	return *o.DateTrim
}

// GetDateTrimOk returns a tuple with the DateTrim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpressionField) GetDateTrimOk() (*bool, bool) {
	if o == nil || o.DateTrim == nil {
		return nil, false
	}
	return o.DateTrim, true
}

// HasDateTrim returns a boolean if a field has been set.
func (o *SourceTypesExpressionField) HasDateTrim() bool {
	return o != nil && o.DateTrim != nil
}

// SetDateTrim gets a reference to the given bool and assigns it to the DateTrim field.
func (o *SourceTypesExpressionField) SetDateTrim(v bool) {
	o.DateTrim = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *SourceTypesExpressionField) GetStartIndex() int64 {
	if o == nil || o.StartIndex == nil {
		var ret int64
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpressionField) GetStartIndexOk() (*int64, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *SourceTypesExpressionField) HasStartIndex() bool {
	return o != nil && o.StartIndex != nil
}

// SetStartIndex gets a reference to the given int64 and assigns it to the StartIndex field.
func (o *SourceTypesExpressionField) SetStartIndex(v int64) {
	o.StartIndex = &v
}

// GetEndIndex returns the EndIndex field value if set, zero value otherwise.
func (o *SourceTypesExpressionField) GetEndIndex() int64 {
	if o == nil || o.EndIndex == nil {
		var ret int64
		return ret
	}
	return *o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpressionField) GetEndIndexOk() (*int64, bool) {
	if o == nil || o.EndIndex == nil {
		return nil, false
	}
	return o.EndIndex, true
}

// HasEndIndex returns a boolean if a field has been set.
func (o *SourceTypesExpressionField) HasEndIndex() bool {
	return o != nil && o.EndIndex != nil
}

// SetEndIndex gets a reference to the given int64 and assigns it to the EndIndex field.
func (o *SourceTypesExpressionField) SetEndIndex(v int64) {
	o.EndIndex = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *SourceTypesExpressionField) GetUsed() bool {
	if o == nil || o.Used == nil {
		var ret bool
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpressionField) GetUsedOk() (*bool, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *SourceTypesExpressionField) HasUsed() bool {
	return o != nil && o.Used != nil
}

// SetUsed gets a reference to the given bool and assigns it to the Used field.
func (o *SourceTypesExpressionField) SetUsed(v bool) {
	o.Used = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SourceTypesExpressionField) GetIndex() bool {
	if o == nil || o.Index == nil {
		var ret bool
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpressionField) GetIndexOk() (*bool, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SourceTypesExpressionField) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given bool and assigns it to the Index field.
func (o *SourceTypesExpressionField) SetIndex(v bool) {
	o.Index = &v
}

// GetNormalization returns the Normalization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesExpressionField) GetNormalization() SourceTypesNormalization {
	if o == nil || o.Normalization.Get() == nil {
		var ret SourceTypesNormalization
		return ret
	}
	return *o.Normalization.Get()
}

// GetNormalizationOk returns a tuple with the Normalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesExpressionField) GetNormalizationOk() (*SourceTypesNormalization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Normalization.Get(), o.Normalization.IsSet()
}

// HasNormalization returns a boolean if a Normalization has been set.
func (o *SourceTypesExpressionField) HasNormalization() bool {
	return o != nil && o.Normalization.IsSet()
}

// SetNormalization gets a reference to the given common.NullableString and assigns it to the Normalization field.
func (o *SourceTypesExpressionField) SetNormalization(v SourceTypesNormalization) {
	o.Normalization.Set(&v)
}

// SetNormalizationNil sets the value for Normalization to be an explicit nil.
func (o *SourceTypesExpressionField) SetNormalizationNil() {
	o.Normalization.Set(nil)
}

// UnSetNormalization ensures that no value is present for Normalization, not even an explicit nil.
func (o *SourceTypesExpressionField) UnSetNormalization() {
	o.Normalization.UnSet()
}

// GetCorrectTwoDigitDayFormatPosition returns the CorrectTwoDigitDayFormatPosition field value if set, zero value otherwise.
func (o *SourceTypesExpressionField) GetCorrectTwoDigitDayFormatPosition() int64 {
	if o == nil || o.CorrectTwoDigitDayFormatPosition == nil {
		var ret int64
		return ret
	}
	return *o.CorrectTwoDigitDayFormatPosition
}

// GetCorrectTwoDigitDayFormatPositionOk returns a tuple with the CorrectTwoDigitDayFormatPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpressionField) GetCorrectTwoDigitDayFormatPositionOk() (*int64, bool) {
	if o == nil || o.CorrectTwoDigitDayFormatPosition == nil {
		return nil, false
	}
	return o.CorrectTwoDigitDayFormatPosition, true
}

// HasCorrectTwoDigitDayFormatPosition returns a boolean if a field has been set.
func (o *SourceTypesExpressionField) HasCorrectTwoDigitDayFormatPosition() bool {
	return o != nil && o.CorrectTwoDigitDayFormatPosition != nil
}

// SetCorrectTwoDigitDayFormatPosition gets a reference to the given int64 and assigns it to the CorrectTwoDigitDayFormatPosition field.
func (o *SourceTypesExpressionField) SetCorrectTwoDigitDayFormatPosition(v int64) {
	o.CorrectTwoDigitDayFormatPosition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesExpressionField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ValueName.IsSet() {
		toSerialize["ValueName"] = o.ValueName.Get()
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.DateFormat.IsSet() {
		toSerialize["DateFormat"] = o.DateFormat.Get()
	}
	if o.DateLang.IsSet() {
		toSerialize["DateLang"] = o.DateLang
	}
	if o.DateTrim != nil {
		toSerialize["DateTrim"] = o.DateTrim
	}
	if o.StartIndex != nil {
		toSerialize["StartIndex"] = o.StartIndex
	}
	if o.EndIndex != nil {
		toSerialize["EndIndex"] = o.EndIndex
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	if o.Normalization.IsSet() {
		toSerialize["Normalization"] = o.Normalization.Get()
	}
	if o.CorrectTwoDigitDayFormatPosition != nil {
		toSerialize["CorrectTwoDigitDayFormatPosition"] = o.CorrectTwoDigitDayFormatPosition
	}
	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourceTypesExpressionField) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                             *string                          `json:"Name,omitempty"`
		ValueName                        common.NullableString            `json:"ValueName,omitempty"`
		Type                             *string                          `json:"Type,omitempty"`
		DateFormat                       common.NullableString            `json:"DateFormat,omitempty"`
		DateLang                         common.NullableString            `json:"DateLang,omitempty"`
		DateTrim                         *bool                            `json:"DateTrim,omitempty"`
		StartIndex                       *int64                           `json:"StartIndex,omitempty"`
		EndIndex                         *int64                           `json:"EndIndex,omitempty"`
		Used                             *bool                            `json:"Used,omitempty"`
		Index                            *bool                            `json:"Index,omitempty"`
		Normalization                    NullableSourceTypesNormalization `json:"Normalization,omitempty"`
		CorrectTwoDigitDayFormatPosition *int64                           `json:"CorrectTwoDigitDayFormat_Position,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Name", "ValueName", "Type", "DateFormat", "DateLang", "DateTrim", "StartIndex", "EndIndex", "Used", "Index", "Normalization", "CorrectTwoDigitDayFormatPosition"})
	} else {
		return err
	}

	o.Name = all.Name
	o.ValueName = all.ValueName
	o.Type = all.Type
	o.DateFormat = all.DateFormat
	o.DateLang = all.DateLang
	o.DateTrim = all.DateTrim
	o.StartIndex = all.StartIndex
	o.EndIndex = all.EndIndex
	o.Used = all.Used
	o.Index = all.Index
	o.Normalization = all.Normalization
	o.CorrectTwoDigitDayFormatPosition = all.CorrectTwoDigitDayFormatPosition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
