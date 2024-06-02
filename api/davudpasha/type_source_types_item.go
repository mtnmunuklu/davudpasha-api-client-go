package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type SourceTypesItem struct {
	Name                *string                               `json:"Name,omitempty"`
	DefCode             *string                               `json:"DefCode,omitempty"`
	Author              common.NullableString                 `json:"Author,omitempty"`
	Icon                common.NullableString                 `json:"Icon,omitempty"`
	CatCode             *string                               `json:"CatCode,omitempty"`
	SourceReaderType    *string                               `json:"SourceReaderType,omitempty"`
	ReleaseDate         *string                               `json:"ReleaseDate,omitempty"`
	Version             *float64                              `json:"Version,omitempty"`
	Expressions         []SourcecTypesExpression              `json:"Expression,omitempty"`
	ExpressionFields    []SourcecTypesExpressionField         `json:"ExpressionFields,omitempty"`
	DashboardItems      []string                              `json:"DashboardItems,omitempty"`
	ClassificationRules []SourceTypesClassificationRule       `json:"ClassificationRules,omitempty"`
	FromMarket          *bool                                 `json:"FromMarket,omitempty"`
	FromModules         *bool                                 `json:"FromModules,omitempty"`
	CopyFromMarket      *bool                                 `json:"CopyFromMarket,omitempty"`
	HasUpdate           *bool                                 `json:"HasUpdate,omitempty"`
	ModuleId            common.NullableString                 `json:"ModuleId,omitempty"`
	ModuleGuid          common.NullableString                 `json:"ModuleGuid,omitempty"`
	ClassificationDefs  []SourceTypesClassificationDefination `json:"ClassificationDefs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}
