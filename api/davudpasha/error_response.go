package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ErrorResponse API Error response
type ErrorResponse struct {
	// A list of errors
	Errors []string `json:"errors"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewErrorResponse creates a new ErrorResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewErrorResponse(errors []string) *ErrorResponse {
	this := ErrorResponse{}
	this.Errors = errors
	return &this
}

// NewErrorResponseWithDefaults creates a new ErrorResponse object.
// This constructor will assign default values to properties that have it defined,
// but it doensn't guarantee that properties requiered by API are set.
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value.
func (o *ErrorResponse) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *ErrorResponse) SetErrors(v []string) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["errors"] = o.Errors

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ErrorResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors *[]string `json:"errors"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Errors == nil {
		return fmt.Errorf("requiered field errors missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"errors"})
	} else {
		return err
	}
	o.Errors = *all.Errors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
