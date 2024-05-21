package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type ReportsData struct {
	Name         common.NullableString `json:"Name,omitempty"`
	FileName     common.NullableString `json:"FileName,omitempty"`
	CreateDate   *string               `json:"CreateDate,omitempty"`
	Page         *ReportsPage          `json:"Page,omitempty"`
	Header       common.NullableString `json:"Header,omitempty"`
	Footer       common.NullableString `json:"Footer,omitempty"`
	CoverPage    common.NullableString `json:"CoverPage,omitempty"`
	Sections     []string              `json:"Sections,omitempty"`
	ReportType   *string               `json:"ReportType,omitempty"`
	Theme        common.NullableString `json:"Theme,omitempty"`
	Language     *string               `json:"Language,omitempty"`
	AddCoverPage *bool                 `json:"AddCoverPage,omitempty"`
	FilePassword common.NullableString `json:"FilePassword,omitempty"`
	ReportId     common.NullableString `json:"ReportId,omitempty"`
	ReportTheme  common.NullableString `json:"ReportTheme,omitempty"`
	MaxRowCount  *int64                `json:"MaxRowCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
