package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// CasesFileAttachment represents the structure of a file attachment in a case.
type CasesFileAttachment struct {
	// Unique identifier for the file.
	FileID *string `json:"FileId,omitempty"`
	// Name of the file.
	FileName *string `json:"FileName,omitempty"`
	// MD5 hash of the file.
	MD5Hash *string `json:"MD5Hash,omitempty"`
	// SHA-256 hash of the file.
	SHA256Hash *string `json:"SHA256Hash,omitempty"`
	// Size of the file.
	FileSize *string `json:"FileSize"`
	// Extension of the file.
	FileExtension *string `json:"FileExtension,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewCasesFileAttachment creates a new CasesFileAttachment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCasesFileAttachment() *CasesFileAttachment {
	this := CasesFileAttachment{}
	return &this
}

// NewCasesFileAttachmentWithDefaults creates a new CasesFileAttachment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCasesFileAttachmentWithDefaults() *CasesFileAttachment {
	this := CasesFileAttachment{}
	return &this
}

// GetFileID returns the FileID field value if set, zero value otherwise.
func (o *CasesFileAttachment) GetFileID() string {
	if o == nil || o.FileID == nil {
		var ret string
		return ret
	}
	return *o.FileID
}

// GetFileIDOk returns a tuple with the FileID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesFileAttachment) GetFileIDOk() (*string, bool) {
	if o == nil || o.FileID == nil {
		return nil, false
	}
	return o.FileID, true
}

// HasFileID returns a boolean if a field has been set.
func (o *CasesFileAttachment) HasFileID() bool {
	return o != nil && o.FileID != nil
}

// SetFileID gets a reference to the given string and assigns it to the FileID field.
func (o *CasesFileAttachment) SetFileID(v string) {
	o.FileID = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *CasesFileAttachment) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesFileAttachment) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *CasesFileAttachment) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *CasesFileAttachment) SetFileName(v string) {
	o.FileName = &v
}

// GetMD5Hash returns the MD5Hash field value if set, zero value otherwise.
func (o *CasesFileAttachment) GetMD5Hash() string {
	if o == nil || o.MD5Hash == nil {
		var ret string
		return ret
	}
	return *o.MD5Hash
}

// GetMD5HashOk returns a tuple with the MD5Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesFileAttachment) GetMD5HashOk() (*string, bool) {
	if o == nil || o.MD5Hash == nil {
		return nil, false
	}
	return o.MD5Hash, true
}

// HasMD5Hash returns a boolean if a field has been set.
func (o *CasesFileAttachment) HasMD5Hash() bool {
	return o != nil && o.MD5Hash != nil
}

// SetMD5Hash gets a reference to the given string and assigns it to the MD5Hash field.
func (o *CasesFileAttachment) SetMD5Hash(v string) {
	o.MD5Hash = &v
}

// GetSHA256Hash returns the SHA256Hash field value if set, zero value otherwise.
func (o *CasesFileAttachment) GetSHA256Hash() string {
	if o == nil || o.SHA256Hash == nil {
		var ret string
		return ret
	}
	return *o.SHA256Hash
}

// GetSHA256HashOk returns a tuple with the SHA256Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesFileAttachment) GetSHA256HashOk() (*string, bool) {
	if o == nil || o.SHA256Hash == nil {
		return nil, false
	}
	return o.SHA256Hash, true
}

// HasSHA256Hash returns a boolean if a field has been set.
func (o *CasesFileAttachment) HasSHA256Hash() bool {
	return o != nil && o.SHA256Hash != nil
}

// SetSHA256Hash gets a reference to the given string and assigns it to the SHA256Hash field.
func (o *CasesFileAttachment) SetSHA256Hash(v string) {
	o.SHA256Hash = &v
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *CasesFileAttachment) GetFileExtension() string {
	if o == nil || o.FileExtension == nil {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesFileAttachment) GetFileExtensionOk() (*string, bool) {
	if o == nil || o.FileExtension == nil {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *CasesFileAttachment) HasFileExtension() bool {
	return o != nil && o.FileExtension != nil
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *CasesFileAttachment) SetFileExtension(v string) {
	o.FileExtension = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CasesFileAttachment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.FileID != nil {
		toSerialize["FileId"] = o.FileID
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.MD5Hash != nil {
		toSerialize["MD5Hash"] = o.MD5Hash
	}
	if o.SHA256Hash != nil {
		toSerialize["SHA256Hash"] = o.SHA256Hash
	}
	if o.FileSize != nil {
		toSerialize["FileSize"] = o.FileSize
	}
	if o.FileExtension != nil {
		toSerialize["FileExtension"] = o.FileExtension
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *CasesFileAttachment) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileID        *string `json:"FileId,omitempty"`
		FileName      *string `json:"FileName,omitempty"`
		MD5Hash       *string `json:"MD5Hash,omitempty"`
		SHA256Hash    *string `json:"SHA256Hash,omitempty"`
		FileSize      *string `json:"FileSize"`
		FileExtension *string `json:"FileExtension,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"FileId", "FileName", "MD5Hash", "SHA256Hash", "FileSize", "FileExtension"})
	} else {
		return err
	}

	o.FileID = all.FileID
	o.FileName = all.FileName
	o.MD5Hash = all.MD5Hash
	o.SHA256Hash = all.SHA256Hash
	o.FileSize = all.FileSize
	o.FileExtension = all.FileExtension

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
