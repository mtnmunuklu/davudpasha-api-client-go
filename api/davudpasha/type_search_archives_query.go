package davudpasha

import (
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SearchArchivesQuery struct {
	Id                  *string                     `json:"Id,omitempty"`
	Name                *string                     `json:"Name,omitempty"`
	CreateDate          *string                     `json:"CreateDate,omitempty"`
	Includes            common.NullableList[string] `json:"Includes,omitempty"`
	Excludes            common.NullableList[string] `json:"Excludes,omitempty"`
	DateTimeRange       *DateTimeRange              `json:"DateTimeRange,omitempty"`
	LgsIds              common.NullableList[string] `json:"LgsIds,omitempty"`
	LgsNames            common.NullableList[string] `json:"LgsNames,omitempty"`
	Size                *int64                      `json:"Size,omitempty"`
	DeleteAfterDay      *int64                      `json:"DeleteAfterDay,omitempty"`
	ParallelCount       *int64                      `json:"ParallelCount,omitempty"`
	SearchOnRawLogs     *bool                       `json:"SearchOnRawLogs,omitempty"`
	SearchActiveLgsOnly *bool                       `json:"SearchActiveLgsOnly,omitempty"`
	PauseRequest        *bool                       `json:"PauseRequest,omitempty"`
	ReportColNames      common.NullableList[string] `json:"ReportColNames,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
