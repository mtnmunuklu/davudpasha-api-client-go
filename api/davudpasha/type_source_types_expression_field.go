package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type SourcecTypesExpressionField struct {
	Name                              *string                          `json:"Name,omitempty"`
	ValueName                         common.NullableString            `json:"ValueName,omitempty"`
	Type                              *string                          `json:"Type,omitempty"`
	DateFormat                        common.NullableString            `json:"DateFormat,omitempty"`
	DateLang                          common.NullableString            `json:"DateLang,omitempty"`
	DateTrim                          *bool                            `json:"DateTrim,omitempty"`
	StartIndex                        *int64                           `json:"StartIndex,omitempty"`
	EndIndex                          *int64                           `json:"EndIndex,omitempty"`
	Used                              *bool                            `json:"Used,omitempty"`
	Index                             *bool                            `json:"Index,omitempty"`
	Normalization                     NullableSourceTypesNormalization `json:"Normalization,omitempty"`
	CorrectTwoDigitDayFormat_Position *int64                           `json:"CorrectTwoDigitDayFormat_Position,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}
