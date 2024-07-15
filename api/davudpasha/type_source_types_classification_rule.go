package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourceTypesClassificationRule defines the structure for classification rules for source types.
type SourceTypesClassificationRule struct {
	// QueryStr represents the query string for the classification rule.
	QueryStr *string `json:"QueryStr,omitempty"`
	// ClassificationID specifies the identifier for the classification rule.
	ClassificationID *string `json:"ClassificationID,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSourceTypesClassificationRule creates a new SourceTypesClassificationRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesClassificationRule() *SourceTypesClassificationRule {
	this := SourceTypesClassificationRule{}
	return &this
}

// NewSourceTypesClassificationRuleWithDefaults creates a new SourceTypesClassificationRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourceTypesClassificationRuleWithDefaults() *SourceTypesClassificationRule {
	this := SourceTypesClassificationRule{}
	return &this
}

// GetQueryStr returns the QueryStr field value if set, zero value otherwise.
func (o *SourceTypesClassificationRule) GetQueryStr() string {
	if o == nil || o.QueryStr == nil {
		var ret string
		return ret
	}
	return *o.QueryStr
}

// GetQueryStrOk returns a tuple with the QueryStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesClassificationRule) GetQueryStrOk() (*string, bool) {
	if o == nil || o.QueryStr == nil {
		return nil, false
	}
	return o.QueryStr, true
}

// HasQueryStr returns a boolean if a field has been set.
func (o *SourceTypesClassificationRule) HasQueryStr() bool {
	return o != nil && o.QueryStr != nil
}

// SetQueryStr gets a reference to the given string and assigns it to the QueryStr field.
func (o *SourceTypesClassificationRule) SetQueryStr(v string) {
	o.QueryStr = &v
}

// GetClassificationID returns the ClassificationID field value if set, zero value otherwise.
func (o *SourceTypesClassificationRule) GetClassificationID() string {
	if o == nil || o.ClassificationID == nil {
		var ret string
		return ret
	}
	return *o.ClassificationID
}

// GetClassificationIDOk returns a tuple with the ClassificationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesClassificationRule) GetClassificationIDOk() (*string, bool) {
	if o == nil || o.ClassificationID == nil {
		return nil, false
	}
	return o.ClassificationID, true
}

// HasClassificationID returns a boolean if a field has been set.
func (o *SourceTypesClassificationRule) HasClassificationID() bool {
	return o != nil && o.ClassificationID != nil
}

// SetClassificationID gets a reference to the given string and assigns it to the ClassificationID field.
func (o *SourceTypesClassificationRule) SetClassificationID(v string) {
	o.ClassificationID = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesClassificationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.QueryStr != nil {
		toSerialize["QueryStr"] = o.QueryStr
	}
	if o.ClassificationID != nil {
		toSerialize["ClassificationID"] = o.ClassificationID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourceTypesClassificationRule) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryStr         *string `json:"QueryStr,omitempty"`
		ClassificationID *string `json:"ClassificationID,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"QueryStr", "ClassificationID"})
	} else {
		return err
	}

	o.QueryStr = all.QueryStr
	o.ClassificationID = all.ClassificationID

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
