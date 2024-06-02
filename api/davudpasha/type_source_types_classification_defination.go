package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type SourceTypesClassificationDefination struct {
	ID                *string                       `json:"ID,omitempty"`
	DefID             *string                       `json:"DefID,omitempty"`
	Name              *string                       `json:"Name,omitempty"`
	Severity          *string                       `json:"Severity,omitempty"`
	MitreTags         common.NullableList[MitreTag] `json:"MitreTags,omitempty"`
	KillChainPhase    common.NullableString         `json:"KillChainPhase,omitempty"`
	MitreCreationDate common.NullableString         `json:"MitreCreationDate,omitempty"`
	FromMarket        *bool                         `json:"FromMarket,omitempty"`
	FromModules       *bool                         `json:"FromModules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}
