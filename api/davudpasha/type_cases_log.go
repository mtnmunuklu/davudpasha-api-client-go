package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// CasesLog represents the structure of a log entry in a case.
type CasesLog struct {
	// Unique identifier for the log detail.
	DetailID *string `json:"DetailID,omitempty"`
	// Timestamp of when the log entry was recorded.
	Time *string `json:"Time,omitempty"`
	// Source of the log entry.
	Source *string `json:"Source,omitempty"`
	// Content of the log message.
	Message *string `json:"Message,omitempty"`
	// Type or category of the log entry.
	Type *string `json:"Type,omitempty"`
	// List of files associated with the log entry, which can be null.
	Files common.NullableList[string] `json:"Files,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCasesLog creates a new CasesLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCasesLog() *CasesLog {
	this := CasesLog{}
	return &this
}

// NewCasesLogWithDefaults creates a new CasesLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCasesLogWithDefaults() *CasesLog {
	this := CasesLog{}
	return &this
}

// GetDetailID returns the DetailID field value if set, zero value otherwise.
func (o *CasesLog) GetDetailID() string {
	if o == nil || o.DetailID == nil {
		var ret string
		return ret
	}
	return *o.DetailID
}

// GetDetailIDOk returns a tuple with the DetailID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesLog) GetDetailIDOk() (*string, bool) {
	if o == nil || o.DetailID == nil {
		return nil, false
	}
	return o.DetailID, true
}

// HasDetailID returns a boolean if a field has been set.
func (o *CasesLog) HasDetailID() bool {
	return o != nil && o.DetailID != nil
}

// SetDetailID gets a reference to the given string and assigns it to the DetailID field.
func (o *CasesLog) SetDetailID(v string) {
	o.DetailID = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CasesLog) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesLog) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CasesLog) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *CasesLog) SetTime(v string) {
	o.Time = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CasesLog) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesLog) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CasesLog) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CasesLog) SetSource(v string) {
	o.Source = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CasesLog) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesLog) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CasesLog) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CasesLog) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CasesLog) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesLog) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CasesLog) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CasesLog) SetType(v string) {
	o.Type = &v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesLog) GetFiles() []string {
	if o == nil || o.Files.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Files.Get()
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesLog) GetFilesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files.Get(), o.Files.IsSet()
}

// HasFiles returns a boolean if a Files has been set.
func (o *CasesLog) HasFiles() bool {
	return o != nil && o.Files.IsSet()
}

// SetFiles gets a reference to the given common.Nullable[]string and assigns it to the Files field.
func (o *CasesLog) SetFiles(v []string) {
	o.Files.Set(&v)
}

// SetFilesNil sets the value for Files to be an explicit nil.
func (o *CasesLog) SetFilesNil() {
	o.Files.Set(nil)
}

// UnSetFiles ensures that no value is present for Files, not even an explicit nil.
func (o *CasesLog) UnsetFiles() {
	o.Files.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CasesLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.DetailID != nil {
		toSerialize["DetailID"] = o.DetailID
	}
	if o.Time != nil {
		toSerialize["Time"] = o.Time
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Source
	}
	if o.Files.IsSet() {
		toSerialize["Files"] = o.Files.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *CasesLog) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		DetailID *string                     `json:"DetailID,omitempty"`
		Time     *string                     `json:"Time,omitempty"`
		Source   *string                     `json:"Source,omitempty"`
		Message  *string                     `json:"Message,omitempty"`
		Type     *string                     `json:"Type,omitempty"`
		Files    common.NullableList[string] `json:"Files,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"DetailID", "Time", "Source", "Message", "Type", "Files"})
	} else {
		return err
	}

	o.DetailID = all.DetailID
	o.Time = all.Time
	o.Source = all.Source
	o.Message = all.Message
	o.Type = all.Type
	o.Files = all.Files

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
