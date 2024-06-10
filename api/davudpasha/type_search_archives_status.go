package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type SearchArchivesStatus struct {
	SearchArchiveId       *string               `json:"SearchArchiveId,omitempty"`
	Message               common.NullableString `json:"Message,omitempty"`
	ExecutionStartTimeUtc *string               `json:"ExecutionStartTimeUtc,omitempty"`
	ExecutionDurationMs   *int64                `json:"ExecutionDurationMs,omitempty"`
	Status                *string               `json:"Status,omitempty"`
	PauseRequest          *bool                 `json:"PauseRequest,omitempty"`
	FoundLines            *int64                `json:"FoundLines,omitempty"`
	ScannedLines          *int64                `json:"ScannedLines,omitempty"`
	ProcessedFileCount    *int64                `json:"ProcessedFileCount,omitempty"`
	ProcessedFileDate     *string               `json:"ProcessedFileDate,omitempty"`
	ActualStartDate       *string               `json:"ActualStartDate,omitempty"`
	ActualEndDate         *string               `json:"ActualEndDate,omitempty"`
	LimitReached          *bool                 `json:"LimitReached,omitempty"`
	RunningMachineIp      *string               `json:"RunningMachineIp,omitempty"`
	TenantID              *string               `json:"TenantID,omitempty"`
	Id                    *string               `json:"_id,omitempty"`
	Timestamp             *string               `json:"Timestamp,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
