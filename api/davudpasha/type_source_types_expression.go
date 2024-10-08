package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourceTypesExpression represents the structure for expressions in source types.
type SourceTypesExpression struct {
	// Represents a key value, used as an identifier.
	Key *float64 `json:"Key,omitempty"`
	// Name specifies the name of the expression.
	Name *string `json:"Name,omitempty"`
	// ExcludeThisExpression indicates whether to exclude this expression.
	ExcludeThisExpression *bool `json:"ExcludeThisExpression,omitempty"`
	// LogParserType specifies the type of the log parser used. Example: Module csiem.parser.regex,
	// Module csiem.parser.delimiter, Module csiem.parser.csv, Module csiem.parser.code, Module csiem.parser.static
	// Module csiem.parser.db, Module csiem.parser.none
	LogParserType SourceTypesParserType `json:"LogParserType,omitempty"`
	// LogParserData contains the log parser data for the expression.
	LogParserData *SourceTypesLogParserData `json:"LogParserData,omitempty"`
	// Mapping specifies the mappings associated with the expression.
	Mapping []SourceTypesMapping `json:"Mapping,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourceTypesExpression creates a new SourceTypesExpression object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesExpression() *SourceTypesExpression {
	this := SourceTypesExpression{}
	return &this
}

// NewSourceTypesExpressionWithDefaults creates a new SourceTypesExpression object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourceTypesExpressionWithDefaults() *SourceTypesExpression {
	this := SourceTypesExpression{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SourceTypesExpression) GetKey() float64 {
	if o == nil || o.Key == nil {
		var ret float64
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpression) GetKeyOk() (*float64, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SourceTypesExpression) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given float64 and assigns it to the Key field.
func (o *SourceTypesExpression) SetKey(v float64) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceTypesExpression) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpression) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceTypesExpression) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceTypesExpression) SetName(v string) {
	o.Name = &v
}

// GetExcludeThisExpression returns the ExcludeThisExpression field value if set, zero value otherwise.
func (o *SourceTypesExpression) GetExcludeThisExpression() bool {
	if o == nil || o.ExcludeThisExpression == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeThisExpression
}

// GetExcludeThisExpressionOk returns a tuple with the ExcludeThisExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpression) GetExcludeThisExpressionOk() (*bool, bool) {
	if o == nil || o.ExcludeThisExpression == nil {
		return nil, false
	}
	return o.ExcludeThisExpression, true
}

// HasExcludeThisExpression returns a boolean if a field has been set.
func (o *SourceTypesExpression) HasExcludeThisExpression() bool {
	return o != nil && o.ExcludeThisExpression != nil
}

// SetExcludeThisExpression gets a reference to the given bool and assigns it to the ExcludeThisExpression field.
func (o *SourceTypesExpression) SetExcludeThisExpression(v bool) {
	o.ExcludeThisExpression = &v
}

// GetLogParserType returns the LogParserType field value if set, zero value otherwise.
func (o *SourceTypesExpression) GetLogParserType() SourceTypesParserType {
	if o == nil {
		var ret SourceTypesParserType
		return ret
	}
	return o.LogParserType
}

// GetLogParserTypeOk returns a tuple with the LogParserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpression) GetLogParserTypeOk() (*SourceTypesParserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogParserType, true
}

// SetLogParserType gets a reference to the given string and assigns it to the LogParserType field.
func (o *SourceTypesExpression) SetLogParserType(v SourceTypesParserType) {
	o.LogParserType = v
}

// GetLogParserData returns the LogParserData field value if set, zero value otherwise.
func (o *SourceTypesExpression) GetLogParserData() SourceTypesLogParserData {
	if o == nil || o.LogParserData == nil {
		var ret SourceTypesLogParserData
		return ret
	}
	return *o.LogParserData
}

// GetLogParserDataOk returns a tuple with the LogParserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpression) GetLogParserDataOk() (*SourceTypesLogParserData, bool) {
	if o == nil || o.LogParserData == nil {
		return nil, false
	}
	return o.LogParserData, true
}

// HasLogParserData returns a boolean if a field has been set.
func (o *SourceTypesExpression) HasLogParserData() bool {
	return o != nil && o.LogParserData != nil
}

// SetLogParserData gets a reference to the given SourceTypesLogParserData and assigns it to the LogParserData field.
func (o *SourceTypesExpression) SetLogParserData(v SourceTypesLogParserData) {
	o.LogParserData = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *SourceTypesExpression) GetMapping() []SourceTypesMapping {
	if o == nil || o.Mapping == nil {
		var ret []SourceTypesMapping
		return ret
	}
	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesExpression) GetMappingOk() (*[]SourceTypesMapping, bool) {
	if o == nil || o.Mapping == nil {
		return nil, false
	}
	return &o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *SourceTypesExpression) HasMapping() bool {
	return o != nil && o.Mapping != nil
}

// SetMapping gets a reference to the given []SourceTypesMapping and assigns it to the Mapping field.
func (o *SourceTypesExpression) SetMapping(v []SourceTypesMapping) {
	o.Mapping = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesExpression) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ExcludeThisExpression != nil {
		toSerialize["ExcludeThisExpression"] = o.ExcludeThisExpression
	}
	if o.LogParserType.IsValid() {
		toSerialize["LogParserType"] = o.LogParserType
	}
	if o.LogParserData != nil {
		toSerialize["LogParserData"] = o.LogParserType
	}
	if o.Mapping != nil {
		toSerialize["Mapping"] = o.Mapping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourceTypesExpression) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key                   *float64                  `json:"Key,omitempty"`
		Name                  *string                   `json:"Name,omitempty"`
		ExcludeThisExpression *bool                     `json:"ExcludeThisExpression,omitempty"`
		LogParserType         *SourceTypesParserType    `json:"LogParserType,omitempty"`
		LogParserData         *SourceTypesLogParserData `json:"LogParserData,omitempty"`
		Mapping               []SourceTypesMapping      `json:"Mapping,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}
	if all.ExcludeThisExpression == nil {
		return fmt.Errorf("requiered field ExcludeThisExpression is missing")
	}
	if all.LogParserType.IsValid() {
		return fmt.Errorf("requiered field LogParserType is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Key", "Name", "ExcludeThisExpression", "LogParserType", "LogParserData", "Mapping"})
	} else {
		return err
	}
	o.Key = all.Key
	o.Name = all.Name
	o.ExcludeThisExpression = all.ExcludeThisExpression
	hasInvalidField := false
	if !all.LogParserType.IsValid() {
		hasInvalidField = true
	} else {
		o.LogParserType = *all.LogParserType
	}
	if all.LogParserData != nil && all.LogParserData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogParserData = all.LogParserData
	o.Mapping = all.Mapping

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
