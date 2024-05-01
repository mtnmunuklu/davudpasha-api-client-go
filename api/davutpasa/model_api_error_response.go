package davutpasa

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davutpasa-api-client-go/api/common"
)

// APIErrorResponse API Error response
type APIErrorResponse struct {
	// A list of errors
	Errors []string `json:"errors"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObjects      map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAPIErrorResponse creates a new APIErrorResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPIErrorResponse(errors []string) *APIErrorResponse {
	this := APIErrorResponse{}
	this.Errors = errors
	return &this
}

// NewAPIErrorResponseWithDefaults creates a new APIErrorResponse object.
// This constructor will assign default values to properties that have it defined,
// but it doensn't guarantee that properties requiered by API are set.
func NewAPIErrorResponseWithDefaults() *APIErrorResponse {
	this := APIErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value.
func (o *APIErrorResponse) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *APIErrorResponse) GetErrorsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *APIErrorResponse) SetErrors(v []string) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o APIErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObjects != nil {
		return json.Marshal(o.UnparsedObjects)
	}
	toSerialize["errors"] = o.Errors

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *APIErrorResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors *[]string `json:"errors"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObjects)
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
