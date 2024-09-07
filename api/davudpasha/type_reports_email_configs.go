package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsEmailConfigs represents the email configuration settings for reports.
type ReportsEmailConfigs struct {
	// The email address to send the report to.
	Email common.NullableString `json:"Email,omitempty"`
	// The subject line for the report email.
	Subject common.NullableString `json:"Subject,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsEmailConfigs creates a new ReportsEmailConfigs object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsEmailConfigs() *ReportsEmailConfigs {
	this := ReportsEmailConfigs{}
	return &this
}

// NewReportsEmailConfigsWithDefaults creates a new ReportsEmailConfigs object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsEmailConfigsWithDefaults() *ReportsEmailConfigs {
	this := ReportsEmailConfigs{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsEmailConfigs) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsEmailConfigs) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ReportsEmailConfigs) HasEmail() bool {
	return o != nil && o.Email.IsSet()
}

// SetEmail gets a reference to the given common.NullableString and assigns it to the Email field.
func (o *ReportsEmailConfigs) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil.
func (o *ReportsEmailConfigs) SetEmailNil() {
	o.Email.Set(nil)
}

// UnSetEmail ensures that no value is present for Email, not even an explicit nil.
func (o *ReportsEmailConfigs) UnsetEmail() {
	o.Email.Unset()
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsEmailConfigs) GetSubject() string {
	if o == nil || o.Subject.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsEmailConfigs) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *ReportsEmailConfigs) HasSubject() bool {
	return o != nil && o.Subject.IsSet()
}

// SetSubject gets a reference to the given common.NullableString and assigns it to the Subject field.
func (o *ReportsEmailConfigs) SetSubject(v string) {
	o.Subject.Set(&v)
}

// SetSubjectNil sets the value for Subject to be an explicit nil.
func (o *ReportsEmailConfigs) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnSetSubject ensures that no value is present for Subject, not even an explicit nil.
func (o *ReportsEmailConfigs) UnsetSubject() {
	o.Subject.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsEmailConfigs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Email.IsSet() {
		toSerialize["Email"] = o.Email.Get()
	}
	if o.Subject.IsSet() {
		toSerialize["Subject"] = o.Subject.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsEmailConfigs) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email   common.NullableString `json:"Email,omitempty"`
		Subject common.NullableString `json:"Subject,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Email", "Subject"})
	} else {
		return err
	}

	o.Email = all.Email
	o.Subject = all.Subject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
