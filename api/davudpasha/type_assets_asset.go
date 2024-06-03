package davudpasha

import (
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type AssetsAsset struct {
	ID                *string                     `json:"_id,omitempty"`
	Timestamp         common.NullableString       `json:"Timestamp,omitempty"`
	ParentGroupIds    common.NullableList[string] `json:"ParentGroupIds,omitempty"`
	TenantID          *string                     `json:"TenantID,omitempty"`
	AssetId           *string                     `json:"AssetId,omitempty"`
	CustomerId        common.NullableString       `json:"CustomerId,omitempty"`
	Name              *string                     `json:"Name,omitempty"`
	Ip                *string                     `json:"Ip,omitempty"`
	HostName          *string                     `json:"HostName,omitempty"`
	MacAddress        *string                     `json:"MacAddress,omitempty"`
	OsInfo            *string                     `json:"OsInfo,omitempty"`
	EnterDate         *string                     `json:"EnterDate,omitempty"`
	IsPassive         *bool                       `json:"IsPassive,omitempty"`
	IsDeleted         *bool                       `json:"IsDeleted,omitempty"`
	Credential        common.NullableString       `json:"Credential,omitempty"`
	GroupId           *string                     `json:"GroupId,omitempty"`
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
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
