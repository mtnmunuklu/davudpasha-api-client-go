package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourcesFileInfo represents the file-related information for processing log files.
type SourcesFileInfo struct {
	// ID of the credential used to access the file.
	CredentialId common.NullableString `json:"CredentialId,omitempty"`
	// Host address where the file is located.
	Host common.NullableString `json:"Host,omitempty"`
	// Port number used to connect to the host.
	Port common.NullableString `json:"Port,omitempty"`
	// Path to the directory where the file is located.
	Path *string `json:"Path,omitempty"`
	// Name of the file to be processed.
	FileName *string `json:"FileName,omitempty"`
	// Format pattern for the file name, if it follows a specific convention.
	FileNameFormat common.NullableString `json:"FileNameFormat,omitempty"`
	// Code page used for encoding the file content.
	CodePage common.NullableString `json:"CodePage,omitempty"`
	// Indicates if the file name is static (i.e., does not change).
	FileNameStatic common.NullableBool `json:"FileNameStatic,omitempty"`
	// Indicates if old files should be deleted after processing.
	DeleteOldFile common.NullableBool `json:"DeleteOldFile,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourcesFileInfo creates a new SourcesFileInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcesFileInfo() *SourcesFileInfo {
	this := SourcesFileInfo{}
	return &this
}

// NewSourcesFileInfoWithDefaults creates a new SourcesFileInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcesFileInfoWithDefaults() *SourcesFileInfo {
	this := SourcesFileInfo{}
	return &this
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFileInfo) GetCredentialId() string {
	if o == nil || o.CredentialId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CredentialId.Get()
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFileInfo) GetCredentialIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialId.Get(), o.CredentialId.IsSet()
}

// HasCredentialId returns a boolean if a CredentialId has been set.
func (o *SourcesFileInfo) HasCredentialId() bool {
	return o != nil && o.CredentialId.IsSet()
}

// SetCredentialId gets a reference to the given common.NullableString and assigns it to the CredentialId field.
func (o *SourcesFileInfo) SetCredentialId(v string) {
	o.CredentialId.Set(&v)
}

// SetCredentialIdNil sets the value for CredentialId to be an explicit nil.
func (o *SourcesFileInfo) SetCredentialIdNil() {
	o.CredentialId.Set(nil)
}

// UnSetCredentialId ensures that no value is present for CredentialId, not even an explicit nil.
func (o *SourcesFileInfo) UnsetCredentialId() {
	o.CredentialId.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFileInfo) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFileInfo) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a Host has been set.
func (o *SourcesFileInfo) HasHost() bool {
	return o != nil && o.Host.IsSet()
}

// SetHost gets a reference to the given common.NullableString and assigns it to the Host field.
func (o *SourcesFileInfo) SetHost(v string) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil.
func (o *SourcesFileInfo) SetHostNil() {
	o.Host.Set(nil)
}

// UnSetHost ensures that no value is present for Host, not even an explicit nil.
func (o *SourcesFileInfo) UnsetHost() {
	o.Host.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFileInfo) GetPort() string {
	if o == nil || o.Port.Get() == nil {
		var ret string
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFileInfo) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a Port has been set.
func (o *SourcesFileInfo) HasPort() bool {
	return o != nil && o.Port.IsSet()
}

// SetPort gets a reference to the given common.NullableString and assigns it to the Port field.
func (o *SourcesFileInfo) SetPort(v string) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil.
func (o *SourcesFileInfo) SetPortNil() {
	o.Port.Set(nil)
}

// UnSetPort ensures that no value is present for Port, not even an explicit nil.
func (o *SourcesFileInfo) UnsetPort() {
	o.Port.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SourcesFileInfo) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesFileInfo) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SourcesFileInfo) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SourcesFileInfo) SetPath(v string) {
	o.Path = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *SourcesFileInfo) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesFileInfo) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *SourcesFileInfo) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *SourcesFileInfo) SetFileName(v string) {
	o.FileName = &v
}

// GetFileNameFormat returns the FileNameFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFileInfo) GetFileNameFormat() string {
	if o == nil || o.FileNameFormat.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileNameFormat.Get()
}

// GetFileNameFormatOk returns a tuple with the FileNameFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFileInfo) GetFileNameFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileNameFormat.Get(), o.FileNameFormat.IsSet()
}

// HasFileNameFormat returns a boolean if a FileNameFormat has been set.
func (o *SourcesFileInfo) HasFileNameFormat() bool {
	return o != nil && o.FileNameFormat.IsSet()
}

// SetFileNameFormat gets a reference to the given common.NullableString and assigns it to the FileNameFormat field.
func (o *SourcesFileInfo) SetFileNameFormat(v string) {
	o.FileNameFormat.Set(&v)
}

// SetFileNameFormatNil sets the value for FileNameFormat to be an explicit nil.
func (o *SourcesFileInfo) SetFileNameFormatNil() {
	o.FileNameFormat.Set(nil)
}

// UnSetFileNameFormat ensures that no value is present for FileNameFormat, not even an explicit nil.
func (o *SourcesFileInfo) UnsetFileNameFormat() {
	o.FileNameFormat.Unset()
}

// GetCodePage returns the CodePage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFileInfo) GetCodePage() string {
	if o == nil || o.CodePage.Get() == nil {
		var ret string
		return ret
	}
	return *o.CodePage.Get()
}

// GetCodePageOk returns a tuple with the CodePage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFileInfo) GetCodePageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodePage.Get(), o.CodePage.IsSet()
}

// HasCodePage returns a boolean if a CodePage has been set.
func (o *SourcesFileInfo) HasCodePage() bool {
	return o != nil && o.CodePage.IsSet()
}

// SetCodePage gets a reference to the given common.NullableString and assigns it to the CodePage field.
func (o *SourcesFileInfo) SetCodePage(v string) {
	o.CodePage.Set(&v)
}

// SetCodePageNil sets the value for CodePage to be an explicit nil.
func (o *SourcesFileInfo) SetCodePageNil() {
	o.CodePage.Set(nil)
}

// UnSetCodePage ensures that no value is present for CodePage, not even an explicit nil.
func (o *SourcesFileInfo) UnsetCodePage() {
	o.CodePage.Unset()
}

// GetFileNameStatic returns the FileNameStatic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFileInfo) GetFileNameStatic() bool {
	if o == nil || o.FileNameStatic.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FileNameStatic.Get()
}

// GetFileNameStaticOk returns a tuple with the FileNameStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFileInfo) GetFileNameStaticOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileNameStatic.Get(), o.FileNameStatic.IsSet()
}

// HasFileNameStatic returns a boolean if a FileNameStatic has been set.
func (o *SourcesFileInfo) HasFileNameStatic() bool {
	return o != nil && o.FileNameStatic.IsSet()
}

// SetFileNameStatic gets a reference to the given common.NullableString and assigns it to the FileNameStatic field.
func (o *SourcesFileInfo) SetFileNameStatic(v bool) {
	o.FileNameStatic.Set(&v)
}

// SetFileNameStaticNil sets the value for FileNameStatic to be an explicit nil.
func (o *SourcesFileInfo) SetFileNameStaticNil() {
	o.FileNameStatic.Set(nil)
}

// UnSetFileNameStatic ensures that no value is present for FileNameStatic, not even an explicit nil.
func (o *SourcesFileInfo) UnsetFileNameStatic() {
	o.FileNameStatic.Unset()
}

// GetDeleteOldFile returns the DeleteOldFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesFileInfo) GetDeleteOldFile() bool {
	if o == nil || o.DeleteOldFile.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOldFile.Get()
}

// GetDeleteOldFileOk returns a tuple with the DeleteOldFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesFileInfo) GetDeleteOldFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeleteOldFile.Get(), o.DeleteOldFile.IsSet()
}

// HasDeleteOldFile returns a boolean if a DeleteOldFile has been set.
func (o *SourcesFileInfo) HasDeleteOldFile() bool {
	return o != nil && o.DeleteOldFile.IsSet()
}

// SetDeleteOldFile gets a reference to the given common.NullableString and assigns it to the DeleteOldFile field.
func (o *SourcesFileInfo) SetDeleteOldFile(v bool) {
	o.DeleteOldFile.Set(&v)
}

// SetDeleteOldFileNil sets the value for DeleteOldFile to be an explicit nil.
func (o *SourcesFileInfo) SetDeleteOldFileNil() {
	o.DeleteOldFile.Set(nil)
}

// UnSetDeleteOldFile ensures that no value is present for DeleteOldFile, not even an explicit nil.
func (o *SourcesFileInfo) UnsetDeleteOldFile() {
	o.DeleteOldFile.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcesFileInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CredentialId.IsSet() {
		toSerialize["CredentialId"] = o.CredentialId.Get()
	}
	if o.Host.IsSet() {
		toSerialize["Host"] = o.Host.Get()
	}
	if o.Port.IsSet() {
		toSerialize["Port"] = o.Port.Get()
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.Port.IsSet() {
		toSerialize["Port"] = o.Port.Get()
	}
	if o.FileNameFormat.IsSet() {
		toSerialize["FileNameFormat"] = o.FileNameFormat.Get()
	}
	if o.CodePage.IsSet() {
		toSerialize["CodePage"] = o.CodePage.Get()
	}
	if o.FileNameStatic.IsSet() {
		toSerialize["FileNameStatic"] = o.FileNameStatic.Get()
	}
	if o.DeleteOldFile.IsSet() {
		toSerialize["DeleteOldFile"] = o.DeleteOldFile.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourcesFileInfo) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		CredentialId   common.NullableString `json:"CredentialId,omitempty"`
		Host           common.NullableString `json:"Host,omitempty"`
		Port           common.NullableString `json:"Port,omitempty"`
		Path           *string               `json:"Path,omitempty"`
		FileName       *string               `json:"FileName,omitempty"`
		FileNameFormat common.NullableString `json:"FileNameFormat,omitempty"`
		CodePage       common.NullableString `json:"CodePage,omitempty"`
		FileNameStatic common.NullableBool   `json:"FileNameStatic,omitempty"`
		DeleteOldFile  common.NullableBool   `json:"DeleteOldFile,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Path == nil {
		return fmt.Errorf("requiered field Path is missing")
	}
	if all.FileName == nil {
		return fmt.Errorf("requiered field FileName is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"CredentialId", "Host", "Port", "Path", "FileName", "FileNameFormat", "CodePage", "FileNameStatic", "DeleteOldFile"})
	} else {
		return err
	}

	o.CredentialId = all.CredentialId
	o.Host = all.Host
	o.Port = all.Port
	o.Path = all.Path
	o.FileName = all.FileName
	o.FileNameFormat = all.FileNameFormat
	o.CodePage = all.CodePage
	o.FileNameStatic = all.FileNameStatic
	o.DeleteOldFile = all.DeleteOldFile

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
