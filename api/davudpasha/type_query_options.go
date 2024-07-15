package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// QueryOptions represents the query options that are used.
type QueryOptions struct {
	// ShowHighlight indicates whether to show highlight.
	ShowHighlight *bool `json:"ShowHighlight,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// GetShowHighlight returns the ShowHighlight field value if set, zero value otherwise.
func (o *QueryOptions) GetShowHighlight() bool {
	if o == nil || o.ShowHighlight == nil {
		var ret bool
		return ret
	}
	return *o.ShowHighlight
}

// GetShowHighlightOk returns a tuple with the ShowHighlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOptions) GetShowHighlightOk() (*bool, bool) {
	if o == nil || o.ShowHighlight == nil {
		return nil, false
	}
	return o.ShowHighlight, true
}

// HasShowHighlight returns a boolean if a field has been set.
func (o *QueryOptions) HasShowHighlight() bool {
	return o != nil && o.ShowHighlight != nil
}

// SetShowHighlight gets a reference to the given string and assigns it to the ShowHighlight field.
func (o *QueryOptions) SetShowHighlight(v bool) {
	o.ShowHighlight = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ShowHighlight != nil {
		toSerialize["ShowHighlight"] = o.ShowHighlight
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *QueryOptions) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ShowHighlight *bool `json:"ShowHighlight,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ShowHighlight"})
	} else {
		return err
	}
	o.ShowHighlight = all.ShowHighlight

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
