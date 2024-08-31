package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// AssetsAsset represents the structure of an asset.
type AssetsAsset struct {
	// Unique identifier for the asset.
	ID *string `json:"_id,omitempty"`
	// Timestamp of when the asset was created or last updated.
	Timestamp common.NullableString `json:"Timestamp,omitempty"`
	// List of parent group IDs to which the asset belongs.
	ParentGroupIDs common.NullableList[string] `json:"ParentGroupIds,omitempty"`
	// Unique identifier for the tenant.
	TenantID *string `json:"TenantID,omitempty"`
	// Unique identifier for the asset.
	AssetID *string `json:"AssetId,omitempty"`
	// Identifier for the customer associated with the asset.
	CustomerID common.NullableString `json:"CustomerId,omitempty"`
	// Name of the asset.
	Name *string `json:"Name,omitempty"`
	// IP address of the asset.
	IP *string `json:"Ip,omitempty"`
	// Hostname of the asset.
	HostName *string `json:"HostName,omitempty"`
	// MAC address of the asset.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Information about the operating system of the asset.
	OsInfo *string `json:"OsInfo,omitempty"`
	// Date when the asset was entered into the system.
	EnterDate *string `json:"EnterDate,omitempty"`
	// Indicates whether the asset is passive.
	IsPassive *bool `json:"IsPassive,omitempty"`
	// Indicates whether the asset is deleted.
	IsDeleted *bool `json:"IsDeleted,omitempty"`
	// Credential information for the asset, which can be null.
	Credential common.NullableString `json:"Credential,omitempty"`
	// Identifier for the group to which the asset belongs.
	GroupId *string `json:"GroupId,omitempty"`
	// List of tags associated with the asset.
	Tag []string `json:"Tag,omitempty"`
	// Additional tag data, which can be null.
	TagData common.NullableString `json:"TagData,omitempty"`
	// Unique object GUID for the asset.
	ObjectGUID *string `json:"ObjectGUID,omitempty"`
	// Distinguished name of the asset.
	DistinguishedName *string `json:"DistinguishedName,omitempty"`
	// Indicates whether the asset was manually added.
	IsManuel *bool `json:"IsManuel,omitempty"`
	// Inventory information for the asset.
	InventoryInfo *AssetsInventoryInfo `json:"InventoryInfo,omitempty"`
	// Type of the asset.
	AssetType *string `json:"AssetType,omitempty"`
	// Date of the last access to the asset.
	LastAccessDate *string `json:"LastAccessDate,omitempty"`
	// Indicates whether the asset is manageable.
	IsManage *bool `json:"IsManage,omitempty"`
	// Last user who accessed the asset's session.
	LastUserOfSession *string `json:"LastUserOfSession,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssetsAsset creates a new AssetsAsset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssetsAsset() *AssetsAsset {
	this := AssetsAsset{}
	return &this
}

// NewAssetsAssetWithDefaults creates a new AssetsAsset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssetsAssetWithDefaults() *AssetsAsset {
	this := AssetsAsset{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *AssetsAsset) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AssetsAsset) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *AssetsAsset) SetID(v string) {
	o.ID = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsAsset) GetTimestamp() string {
	if o == nil || o.Timestamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AssetsAsset) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a Timestamp has been set.
func (o *AssetsAsset) HasTimestamp() bool {
	return o != nil && o.Timestamp.IsSet()
}

// SetTimestamp gets a reference to the given common.NullableString and assigns it to the Timestamp field.
func (o *AssetsAsset) SetTimestamp(v string) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil.
func (o *AssetsAsset) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnSetTimestamp ensures that no value is present for Timestamp, not even an explicit nil.
func (o *AssetsAsset) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetParentGroupIDs returns a tuple with the ParentGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AssetsAsset) GetParentGroupIDs() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentGroupIDs.Get(), o.ParentGroupIDs.IsSet()
}

// GetParentGroupIDsOk returns a tuple with the ParentGroupIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AssetsAsset) GetParentGroupIDsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentGroupIDs.Get(), o.ParentGroupIDs.IsSet()
}

// HasParentGroupIDs returns a boolean if a ParentGroupIDs has been set.
func (o *AssetsAsset) HasParentGroupIDs() bool {
	return o != nil && o.ParentGroupIDs.IsSet()
}

// SetParentGroupIDs gets a reference to the given common.Nullable[]string and assigns it to the ParentGroupIDs field.
func (o *AssetsAsset) SetParentGroupIDs(v []string) {
	o.ParentGroupIDs.Set(&v)
}

// SetParentGroupIDsNil sets the value for ParentGroupIDs to be an explicit nil.
func (o *AssetsAsset) SetParentGroupIDsNil() {
	o.ParentGroupIDs.Set(nil)
}

// UnSetParentGroupIDs ensures that no value is present for ParentGroupIDs, not even an explicit nil.
func (o *AssetsAsset) UnsetParentGroupIDs() {
	o.ParentGroupIDs.Unset()
}

// GetTenantID returns the TenantID field value if set, zero value otherwise.
func (o *AssetsAsset) GetTenantID() string {
	if o == nil || o.TenantID == nil {
		var ret string
		return ret
	}
	return *o.TenantID
}

// GetTenantIDOk returns a tuple with the TenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetTenantIDOk() (*string, bool) {
	if o == nil || o.TenantID == nil {
		return nil, false
	}
	return o.TenantID, true
}

// HasTenantID returns a boolean if a field has been set.
func (o *AssetsAsset) HasTenantID() bool {
	return o != nil && o.TenantID != nil
}

// SetTenantID gets a reference to the given string and assigns it to the TenantID field.
func (o *AssetsAsset) SetTenantID(v string) {
	o.TenantID = &v
}

// GetAssetID returns the AssetID field value if set, zero value otherwise.
func (o *AssetsAsset) GetAssetID() string {
	if o == nil || o.AssetID == nil {
		var ret string
		return ret
	}
	return *o.AssetID
}

// GetAssetIDOk returns a tuple with the AssetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetAssetIDOk() (*string, bool) {
	if o == nil || o.AssetID == nil {
		return nil, false
	}
	return o.AssetID, true
}

// HasAssetID returns a boolean if a field has been set.
func (o *AssetsAsset) HasAssetID() bool {
	return o != nil && o.AssetID != nil
}

// SetAssetID gets a reference to the given string and assigns it to the AssetID field.
func (o *AssetsAsset) SetAssetID(v string) {
	o.AssetID = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsAsset) GetCustomerID() string {
	if o == nil || o.CustomerID.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomerID.Get()
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AssetsAsset) GetCustomerIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerID.Get(), o.CustomerID.IsSet()
}

// HasCustomerID returns a boolean if a CustomerID has been set.
func (o *AssetsAsset) HasCustomerID() bool {
	return o != nil && o.CustomerID.IsSet()
}

// SetCustomerID gets a reference to the given common.NullableString and assigns it to the CustomerID field.
func (o *AssetsAsset) SetCustomerID(v string) {
	o.CustomerID.Set(&v)
}

// SetCustomerIDNil sets the value for CustomerID to be an explicit nil.
func (o *AssetsAsset) SetCustomerIDNil() {
	o.CustomerID.Set(nil)
}

// UnSetCustomerID ensures that no value is present for CustomerID, not even an explicit nil.
func (o *AssetsAsset) UnsetCustomerID() {
	o.CustomerID.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetsAsset) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetsAsset) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetsAsset) SetName(v string) {
	o.Name = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *AssetsAsset) GetIp() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *AssetsAsset) HasIP() bool {
	return o != nil && o.IP != nil
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *AssetsAsset) SetIP(v string) {
	o.IP = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *AssetsAsset) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *AssetsAsset) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *AssetsAsset) SetHostName(v string) {
	o.HostName = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *AssetsAsset) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *AssetsAsset) HasMacAddress() bool {
	return o != nil && o.MacAddress != nil
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *AssetsAsset) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetOsInfo returns the OsInfo field value if set, zero value otherwise.
func (o *AssetsAsset) GetOsInfo() string {
	if o == nil || o.OsInfo == nil {
		var ret string
		return ret
	}
	return *o.OsInfo
}

// GetOsInfoOk returns a tuple with the OsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetOsInfoOk() (*string, bool) {
	if o == nil || o.OsInfo == nil {
		return nil, false
	}
	return o.OsInfo, true
}

// HasOsInfo returns a boolean if a field has been set.
func (o *AssetsAsset) HasOsInfo() bool {
	return o != nil && o.OsInfo != nil
}

// SetOsInfo gets a reference to the given string and assigns it to the OsInfo field.
func (o *AssetsAsset) SetOsInfo(v string) {
	o.OsInfo = &v
}

// GetEnterDate returns the EnterDate field value if set, zero value otherwise.
func (o *AssetsAsset) GetEnterDate() string {
	if o == nil || o.EnterDate == nil {
		var ret string
		return ret
	}
	return *o.EnterDate
}

// GetEnterDateOk returns a tuple with the EnterDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetEnterDateOk() (*string, bool) {
	if o == nil || o.EnterDate == nil {
		return nil, false
	}
	return o.EnterDate, true
}

// HasEnterDate returns a boolean if a field has been set.
func (o *AssetsAsset) HasEnterDate() bool {
	return o != nil && o.EnterDate != nil
}

// SetEnterDate gets a reference to the given string and assigns it to the EnterDate field.
func (o *AssetsAsset) SetEnterDate(v string) {
	o.EnterDate = &v
}

// GetIsPassive returns the IsPassive field value if set, zero value otherwise.
func (o *AssetsAsset) GetIsPassive() bool {
	if o == nil || o.IsPassive == nil {
		var ret bool
		return ret
	}
	return *o.IsPassive
}

// GetIsPassiveOk returns a tuple with the IsPassive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetIsPassiveOk() (*bool, bool) {
	if o == nil || o.IsPassive == nil {
		return nil, false
	}
	return o.IsPassive, true
}

// HasIsPassive returns a boolean if a field has been set.
func (o *AssetsAsset) HasIsPassive() bool {
	return o != nil && o.IsPassive != nil
}

// SetIsPassive gets a reference to the given bool and assigns it to the IsPassive field.
func (o *AssetsAsset) SetIsPassive(v bool) {
	o.IsPassive = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AssetsAsset) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AssetsAsset) HasIsDeleted() bool {
	return o != nil && o.IsDeleted != nil
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AssetsAsset) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsAsset) GetCredential() string {
	if o == nil || o.Credential.Get() == nil {
		var ret string
		return ret
	}
	return *o.Credential.Get()
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AssetsAsset) GetCredentialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credential.Get(), o.Credential.IsSet()
}

// HasCredential returns a boolean if a Credential has been set.
func (o *AssetsAsset) HasCredential() bool {
	return o != nil && o.Credential.IsSet()
}

// SetCredential gets a reference to the given common.NullableString and assigns it to the Credential field.
func (o *AssetsAsset) SetCredential(v string) {
	o.Credential.Set(&v)
}

// SetCredentialNil sets the value for Credential to be an explicit nil.
func (o *AssetsAsset) SetCredentialNil() {
	o.Credential.Set(nil)
}

// UnSetCredential ensures that no value is present for Credential, not even an explicit nil.
func (o *AssetsAsset) UnsetCredential() {
	o.Credential.Unset()
}

// GetGroupID returns the GroupID field value if set, zero value otherwise.
func (o *AssetsAsset) GetGroupID() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIDOk returns a tuple with the GroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetGroupIDOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupID returns a boolean if a field has been set.
func (o *AssetsAsset) HasGroupID() bool {
	return o != nil && o.GroupId != nil
}

// SetGroupID gets a reference to the given string and assigns it to the GroupID field.
func (o *AssetsAsset) SetGroupID(v string) {
	o.GroupId = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AssetsAsset) GetTag() []string {
	if o == nil || o.Tag == nil {
		var ret []string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetTagOk() (*[]string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return &o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AssetsAsset) HasTag() bool {
	return o != nil && o.Tag != nil
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *AssetsAsset) SetTag(v []string) {
	o.Tag = v
}

// GetTagData returns the TagData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsAsset) GetTagData() string {
	if o == nil || o.TagData.Get() == nil {
		var ret string
		return ret
	}
	return *o.TagData.Get()
}

// GetTagDataOk returns a tuple with the TagData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AssetsAsset) GetTagDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagData.Get(), o.TagData.IsSet()
}

// HasTagData returns a boolean if a TagData has been set.
func (o *AssetsAsset) HasTagData() bool {
	return o != nil && o.TagData.IsSet()
}

// SetTagData gets a reference to the given common.NullableString and assigns it to the TagData field.
func (o *AssetsAsset) SetTagData(v string) {
	o.TagData.Set(&v)
}

// SetTagDataNil sets the value for TagData to be an explicit nil.
func (o *AssetsAsset) SetTagDataNil() {
	o.TagData.Set(nil)
}

// UnSetTagData ensures that no value is present for TagData, not even an explicit nil.
func (o *AssetsAsset) UnsetTagData() {
	o.TagData.Unset()
}

// GetObjectGUID returns the ObjectGUID field value if set, zero value otherwise.
func (o *AssetsAsset) GetObjectGUID() string {
	if o == nil || o.ObjectGUID == nil {
		var ret string
		return ret
	}
	return *o.ObjectGUID
}

// GetObjectGUIDOk returns a tuple with the ObjectGUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetObjectGUIDOk() (*string, bool) {
	if o == nil || o.ObjectGUID == nil {
		return nil, false
	}
	return o.ObjectGUID, true
}

// HasObjectGUID returns a boolean if a field has been set.
func (o *AssetsAsset) HasObjectGUID() bool {
	return o != nil && o.ObjectGUID != nil
}

// SetObjectGUID gets a reference to the given string and assigns it to the ObjectGUID field.
func (o *AssetsAsset) SetObjectGUID(v string) {
	o.ObjectGUID = &v
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *AssetsAsset) GetDistinguishedName() string {
	if o == nil || o.DistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || o.DistinguishedName == nil {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *AssetsAsset) HasDistinguishedName() bool {
	return o != nil && o.DistinguishedName != nil
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *AssetsAsset) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetIsManuel returns the IsManuel field value if set, zero value otherwise.
func (o *AssetsAsset) GetIsManuel() bool {
	if o == nil || o.IsManuel == nil {
		var ret bool
		return ret
	}
	return *o.IsManuel
}

// GetIsManuelOk returns a tuple with the IsManuel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetIsManuelOk() (*bool, bool) {
	if o == nil || o.IsManuel == nil {
		return nil, false
	}
	return o.IsManuel, true
}

// HasIsManuel returns a boolean if a field has been set.
func (o *AssetsAsset) HasIsManuel() bool {
	return o != nil && o.IsManuel != nil
}

// SetIsManuel gets a reference to the given bool and assigns it to the IsManuel field.
func (o *AssetsAsset) SetIsManuel(v bool) {
	o.IsManuel = &v
}

// GetLastUserOfSession returns the LastUserOfSession field value if set, zero value otherwise.
func (o *AssetsAsset) GetLastUserOfSession() string {
	if o == nil || o.LastUserOfSession == nil {
		var ret string
		return ret
	}
	return *o.LastUserOfSession
}

// GetLastUserOfSessionOk returns a tuple with the LastUserOfSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsAsset) GetLastUserOfSessionOk() (*string, bool) {
	if o == nil || o.LastUserOfSession == nil {
		return nil, false
	}
	return o.LastUserOfSession, true
}

// HasLastUserOfSession returns a boolean if a field has been set.
func (o *AssetsAsset) HasLastUserOfSession() bool {
	return o != nil && o.LastUserOfSession != nil
}

// SetLastUserOfSession gets a reference to the given string and assigns it to the LastUserOfSession field.
func (o *AssetsAsset) SetLastUserOfSession(v string) {
	o.LastUserOfSession = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssetsAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["_id"] = o.ID
	}
	if o.Timestamp.IsSet() {
		toSerialize["Timestamp"] = o.Timestamp.Get()
	}
	if o.ParentGroupIDs.IsSet() {
		toSerialize["ParentGroupIds"] = o.ParentGroupIDs.Get()
	}
	if o.TenantID != nil {
		toSerialize["TenantID"] = o.TenantID
	}
	if o.AssetID != nil {
		toSerialize["AssetId"] = o.AssetID
	}
	if o.CustomerID.IsSet() {
		toSerialize["CustomerId"] = o.CustomerID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.IP != nil {
		toSerialize["Ip"] = o.IP
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.OsInfo != nil {
		toSerialize["OsInfo"] = o.OsInfo
	}
	if o.EnterDate != nil {
		toSerialize["EnterDate"] = o.EnterDate
	}
	if o.IsPassive != nil {
		toSerialize["IsPassive"] = o.IsPassive
	}
	if o.IsDeleted != nil {
		toSerialize["IsDeleted"] = o.IsDeleted
	}
	if o.Credential.IsSet() {
		toSerialize["Credential"] = o.Credential.Get()
	}
	if o.GroupId != nil {
		toSerialize["GroupId"] = o.GroupId
	}
	if o.Tag != nil {
		toSerialize["Tag"] = o.Tag
	}
	if o.TagData.IsSet() {
		toSerialize["TagData"] = o.TagData.Get()
	}
	if o.ObjectGUID != nil {
		toSerialize["ObjectGUID"] = o.ObjectGUID
	}
	if o.DistinguishedName != nil {
		toSerialize["DistinguishedName"] = o.DistinguishedName
	}
	if o.IsManuel != nil {
		toSerialize["IsManuel"] = o.IsManuel
	}
	if o.InventoryInfo != nil {
		toSerialize["InventoryInfo"] = o.InventoryInfo
	}
	if o.AssetType != nil {
		toSerialize["AssetType"] = o.AssetType
	}
	if o.LastAccessDate != nil {
		toSerialize["LastAccessDate"] = o.LastAccessDate
	}
	if o.IsManage != nil {
		toSerialize["IsManage"] = o.IsManage
	}
	if o.LastUserOfSession != nil {
		toSerialize["LastUserOfSession"] = o.LastUserOfSession
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AssetsAsset) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID                *string                     `json:"_id,omitempty"`
		Timestamp         common.NullableString       `json:"Timestamp,omitempty"`
		ParentGroupIDs    common.NullableList[string] `json:"ParentGroupIds,omitempty"`
		TenantID          *string                     `json:"TenantID,omitempty"`
		AssetID           *string                     `json:"AssetId,omitempty"`
		CustomerID        common.NullableString       `json:"CustomerId,omitempty"`
		Name              *string                     `json:"Name,omitempty"`
		IP                *string                     `json:"Ip,omitempty"`
		HostName          *string                     `json:"HostName,omitempty"`
		MacAddress        *string                     `json:"MacAddress,omitempty"`
		OsInfo            *string                     `json:"OsInfo,omitempty"`
		EnterDate         *string                     `json:"EnterDate,omitempty"`
		IsPassive         *bool                       `json:"IsPassive,omitempty"`
		IsDeleted         *bool                       `json:"IsDeleted,omitempty"`
		Credential        common.NullableString       `json:"Credential,omitempty"`
		GroupID           *string                     `json:"GroupId,omitempty"`
		Tag               []string                    `json:"Tag,omitempty"`
		TagData           common.NullableString       `json:"TagData,omitempty"`
		ObjectGUID        *string                     `json:"ObjectGUID,omitempty"`
		DistinguishedName *string                     `json:"DistinguishedName,omitempty"`
		IsManuel          *bool                       `json:"IsManuel,omitempty"`
		InventoryInfo     *AssetsInventoryInfo        `json:"InventoryInfo,omitempty"`
		AssetType         *string                     `json:"AssetType,omitempty"`
		LastAccessDate    *string                     `json:"LastAccessDate,omitempty"`
		IsManage          *bool                       `json:"IsManage,omitempty"`
		LastUserOfSession *string                     `json:"LastUserOfSession,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ID", "Timestamp", "ParentGroupIds", "TenantID", "AssetId", "CustomerId", "Name", "Ip", "HostName", "MacAddress", "OsInfo", "EnterDate", "IsPassive", "IsDeleted", "Credential", "GroupId", "Tag", "TagData", "ObjectGUID", "DistinguishedName", "IsManuel", "InventoryInfo", "AssetType", "LastAccessDate", "IsManage", "LastUserOfSession"})
	} else {
		return err
	}

	o.ID = all.ID
	o.Timestamp = all.Timestamp
	o.ParentGroupIDs = all.ParentGroupIDs
	o.TenantID = all.TenantID
	o.AssetID = all.AssetID
	o.CustomerID = all.CustomerID
	o.Name = all.Name
	o.IP = all.IP
	o.HostName = all.HostName
	o.MacAddress = all.MacAddress
	o.OsInfo = all.OsInfo
	o.EnterDate = all.EnterDate
	o.IsPassive = all.IsPassive
	o.IsDeleted = all.IsDeleted
	o.Credential = all.Credential
	o.GroupId = all.GroupID
	o.Tag = all.Tag
	o.TagData = all.TagData
	o.ObjectGUID = all.ObjectGUID
	o.DistinguishedName = all.DistinguishedName
	o.IsManuel = all.IsManuel
	hasInvalidField := false
	if all.InventoryInfo != nil && all.InventoryInfo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InventoryInfo = all.InventoryInfo
	o.AssetType = all.AssetType
	o.LastAccessDate = all.LastAccessDate
	o.IsManage = all.IsManage
	o.LastUserOfSession = all.LastUserOfSession

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
