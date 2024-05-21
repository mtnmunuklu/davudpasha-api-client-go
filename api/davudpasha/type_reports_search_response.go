package davudpasha

import (
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ReportsSearchResponse struct {
	ModifiedDate            *string                     `json:"ModifiedDate,omitempty"`
	ReportId                *string                     `jsom:"ReportId,omitempty"`
	Username                *string                     `json:"Username,omitempty"`
	Name                    *string                     `json:"Name,omitempty"`
	Description             *string                     `json:"Description,omitempty"`
	IsActive                *bool                       `json:"IsActive,omitempty"`
	Author                  common.NullableString       `json:"Author,omitempty"`
	ReportLink              *string                     `json:"ReportLink,omitempty"`
	SharedUsersAndGroups    common.NullableList[string] `json:"SharedUsersAndGroups,omitempty"`
	CreatedDate             *string                     `json:"CreatedDate,omitempty"`
	ReportQuery             []ReportsQuery              `json:"ReportQuery,omitempty"`
	ReportData              *ReportsData                `json:"ReportData,omitempty"`
	Schedule                *ScheduleConfig             `json:"Schedule,omitempty"`
	LastGenerationTime      common.NullableString       `json:"LastGenerationTime,omitempty"`
	NextGenerationTime      common.NullableString       `json:"NextGenerationTime,omitempty"`
	LatestReportFile        *ReportsLatestReportFile    `json:"LatestReportFile,omitempty"`
	PageSettings            common.NullableString       `json:"PageSettings,omitempty"`
	Parameters              *ReportsParameters          `json:"Parameters,omitempty"`
	Tags                    []string                    `json:"Tags,omitempty"`
	Actions                 *Actions                    `json:"Actions,omitempty"`
	ModuleFilter            *string                     `json:"ModuleFilter,omitempty"`
	RemoteInterfaceName     *string                     `json:"RemoteInterfaceName,omitempty"`
	RemoteMethodName        *string                     `json:"RemoteMethodName,omitempty"`
	ApplicationName         *string                     `json:"ApplicationName,omitempty"`
	SubExecutorModuleFilter *string                     `json:"SubExecutorModuleFilter,omitempty"`
	Version                 *float64                    `json:"Version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
