package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AssetsOperationSystemInfo struct {
	Caption        *string `json:"Caption,omitempty"`
	Manufacturer   *string `json:"Manufacturer,omitempty"`
	OSArchitecture *string `json:"OSArchitecture,omitempty"`
	ServisPack     *string `json:"ServisPack,omitempty"`
	SerialNumber   *string `json:"SerialNumber,omitempty"`
	Version        *string `json:"Version,omitempty"`
	BuildNumber    *string `json:"BuildNumber,omitempty"`
	Description    *string `json:"Description,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewAssetsOperationSystemInfo creates a new AssetsOperationSystemInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssetsOperationSystemInfo() *AssetsOperationSystemInfo {
	this := AssetsOperationSystemInfo{}
	return &this
}

// NewAssetsOperationSystemInfoWithDefaults creates a new AssetsOperationSystemInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssetsOperationSystemInfoWithDefaults() *AssetsOperationSystemInfo {
	this := AssetsOperationSystemInfo{}
	return &this
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *AssetsOperationSystemInfo) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsOperationSystemInfo) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *AssetsOperationSystemInfo) HasCaption() bool {
	return o != nil && o.Caption != nil
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *AssetsOperationSystemInfo) SetCaption(v string) {
	o.Caption = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *AssetsOperationSystemInfo) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsOperationSystemInfo) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *AssetsOperationSystemInfo) HasManufacturer() bool {
	return o != nil && o.Manufacturer != nil
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *AssetsOperationSystemInfo) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetOSArchitecture returns the OSArchitecture field value if set, zero value otherwise.
func (o *AssetsOperationSystemInfo) GetOSArchitecture() string {
	if o == nil || o.OSArchitecture == nil {
		var ret string
		return ret
	}
	return *o.OSArchitecture
}

// GetOSArchitectureOk returns a tuple with the OSArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsOperationSystemInfo) GetOSArchitectureOk() (*string, bool) {
	if o == nil || o.OSArchitecture == nil {
		return nil, false
	}
	return o.OSArchitecture, true
}

// HasOSArchitecture returns a boolean if a field has been set.
func (o *AssetsOperationSystemInfo) HasOSArchitecture() bool {
	return o != nil && o.OSArchitecture != nil
}

// SetOSArchitecture gets a reference to the given string and assigns it to the OSArchitecture field.
func (o *AssetsOperationSystemInfo) SetOSArchitecture(v string) {
	o.OSArchitecture = &v
}

// GetServisPack returns the ServisPack field value if set, zero value otherwise.
func (o *AssetsOperationSystemInfo) GetServisPack() string {
	if o == nil || o.ServisPack == nil {
		var ret string
		return ret
	}
	return *o.ServisPack
}

// GetServisPackOk returns a tuple with the ServisPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsOperationSystemInfo) GetServisPackOk() (*string, bool) {
	if o == nil || o.ServisPack == nil {
		return nil, false
	}
	return o.ServisPack, true
}

// HasServisPack returns a boolean if a field has been set.
func (o *AssetsOperationSystemInfo) HasServisPack() bool {
	return o != nil && o.ServisPack != nil
}

// SetServisPack gets a reference to the given string and assigns it to the ServisPack field.
func (o *AssetsOperationSystemInfo) SetServisPack(v string) {
	o.ServisPack = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *AssetsOperationSystemInfo) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsOperationSystemInfo) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *AssetsOperationSystemInfo) HasSerialNumber() bool {
	return o != nil && o.SerialNumber != nil
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *AssetsOperationSystemInfo) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AssetsOperationSystemInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsOperationSystemInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AssetsOperationSystemInfo) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AssetsOperationSystemInfo) SetVersion(v string) {
	o.Version = &v
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *AssetsOperationSystemInfo) GetBuildNumber() string {
	if o == nil || o.BuildNumber == nil {
		var ret string
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsOperationSystemInfo) GetBuildNumberOk() (*string, bool) {
	if o == nil || o.BuildNumber == nil {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *AssetsOperationSystemInfo) HasBuildNumber() bool {
	return o != nil && o.BuildNumber != nil
}

// SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.
func (o *AssetsOperationSystemInfo) SetBuildNumber(v string) {
	o.BuildNumber = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AssetsOperationSystemInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsOperationSystemInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AssetsOperationSystemInfo) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AssetsOperationSystemInfo) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssetsOperationSystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Caption != nil {
		toSerialize["Caption"] = o.Caption
	}
	if o.Manufacturer != nil {
		toSerialize["Manufacturer"] = o.Manufacturer
	}
	if o.OSArchitecture != nil {
		toSerialize["OSArchitecture"] = o.OSArchitecture
	}
	if o.ServisPack != nil {
		toSerialize["ServisPack"] = o.ServisPack
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.BuildNumber != nil {
		toSerialize["BuildNumber"] = o.BuildNumber
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AssetsOperationSystemInfo) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Caption        *string `json:"Caption,omitempty"`
		Manufacturer   *string `json:"Manufacturer,omitempty"`
		OSArchitecture *string `json:"OSArchitecture,omitempty"`
		ServisPack     *string `json:"ServisPack,omitempty"`
		SerialNumber   *string `json:"SerialNumber,omitempty"`
		Version        *string `json:"Version,omitempty"`
		BuildNumber    *string `json:"BuildNumber,omitempty"`
		Description    *string `json:"Description,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Caption", "Manufacturer", "OSArchitecture", "ServisPack", "SerialNumber", "Version", "BuildNumber", "Description"})
	} else {
		return err
	}

	o.Caption = all.Caption
	o.Manufacturer = all.Manufacturer
	o.OSArchitecture = all.OSArchitecture
	o.ServisPack = all.ServisPack
	o.SerialNumber = all.SerialNumber
	o.Version = all.Version
	o.BuildNumber = all.BuildNumber
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
