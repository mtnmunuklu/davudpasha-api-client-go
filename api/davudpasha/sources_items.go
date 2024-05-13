package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SourcesItems struct {
	Id                  *string                     `json:"Id,omitempty"`
	Enabled             *bool                       `json:"Enabled,omitempty"`
	Name                *string                     `json:"Name,omitempty"`
	Group               common.NullableString       `json:"Group,omitempty"`
	Author              common.NullableString       `json:"Author,omitempty"`
	LogSourceDefCode    *string                     `json:"LogSourceDefCode,omitempty"`
	LogReaderType       *string                     `json:"LogReaderType,omitempty"`
	Tags                []string                    `json:"Tags,omitempty"`
	AlertTimeout        *int64                      `json:"AlertTimeout,omitempty"`
	LogReaderData       *json.RawMessage            `json:"LogReaderData,omitempty"`
	LogOperations       []SourcesLogOperations      `json:"LogOperations,omitempty"`
	DiscardedLogsConfig *string                     `json:"DiscardedLogsConfig,omitempty"`
	Value               *string                     `json:"value,omitempty"`
	Label               *string                     `json:"label,omitempty"`
	IsDeleted           *bool                       `json:"IsDeleted,omitempty"`
	IsAgentSource       *bool                       `json:"IsAgentSource,omitempty"`
	AgentIds            common.NullableList[string] `json:"AgentIds,omitempty"`
	IndexGroupName      common.NullableString       `json:"IndexGroupName,omitempty"`
	DashboardName       *string                     `json:"dashboardName,omitempty"`
	DashboardId         common.NullableString       `json:"dashboardId,omitempty"`
	AssetTags           common.NullableList[string] `json:"AssetTags,omitempty"`
	LogRemoveTime       common.NullableString       `json:"LogRemoveTime,omitempty"`
	LogRemoveFormat     common.NullableString       `json:"LogRemoveFormat,omitempty"`
	AgentId             common.NullableString       `json:"AgentId,omitempty"`
	WriteRawLogs        *bool                       `json:"WriteRawLogs,omitempty"`
	UseSecondaryWriter  *bool                       `json:"UseSecondaryWriter,omitempty"`
	ParallelOptions     *SourcesParallelOptions     `json:"ParallelOptions,omitempty"`
	BlockCount          *int64                      `json:"BlockCount,omitempty"`
	ScheduleConfig      common.Nu

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}
