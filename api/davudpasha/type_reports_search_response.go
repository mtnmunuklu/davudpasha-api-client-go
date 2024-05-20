package davudpasha

import (
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ReportsSearchResponse struct {
	ModifiedDate         *string                     `json:"ModifiedDate,omitempty"`
	ReportId             *string                     `jsom:"ReportId,omitempty"`
	Username             *string                     `json:"Username,omitempty"`
	Name                 *string                     `json:"Name,omitempty"`
	Description          *string                     `json:"Description,omitempty"`
	IsActive             *bool                       `json:"IsActive,omitempty"`
	Author               common.NullableString       `json:"Author,omitempty"`
	ReportLink           *string                     `json:"ReportLink,omitempty"`
	SharedUsersAndGroups common.NullableList[string] `json:"SharedUsersAndGroups,omitempty"`
	CreatedDate          *string                     `json:"CreatedDate,omitempty"`
	ReportQuery          []ReportsReportQuery        `json:"ReportQuery,omitempty"`

	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
