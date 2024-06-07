package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type CasesItem struct {
	CaseId                           *string                     `json:"CaseId,omitempty"`
	Name                             *string                     `json:"Name,omitempty"`
	Description                      common.NullableString       `json:"Description,omitempty"`
	Owner                            *string                     `json:"Owner,omitempty"`
	OccuredDate                      *string                     `json:"OccuredDate,omitempty"`
	CreatedDate                      *string                     `json:"CreatedDate,omitempty"`
	ModifiedDate                     *string                     `json:"ModifiedDate,omitempty"`
	ReminderDate                     common.NullableString       `json:"ReminderDate,omitempty"`
	ExpectedFinishTime               common.NullableString       `json:"ExpectedFinishTime,omitempty"`
	RealFinishTime                   common.NullableString       `json:"RealFinishTime,omitempty"`
	LastSentReminderEmailDate        common.NullableString       `json:"LastSentReminderEmailDate,omitempty"`
	PriorityValueLastCalculationDate common.NullableString       `json:"PriorityValueLastCalculationDate,omitempty"`
	CaseType                         *string                     `json:"CaseType,omitempty"`
	CaseCategory                     *string                     `json:"CaseCategory,omitempty"`
	CaseCategoryId                   *string                     `json:"CaseCategoryId,omitempty"`
	ParentCaseId                     common.NullableString       `json:"ParentCaseId,omitempty"`
	ParentCaseName                   common.NullableString       `json:"ParentCaseName,omitempty"`
	State                            *string                     `json:"State,omitempty"`
	Severity                         *string                     `json:"Severity,omitempty"`
	App                              *string                     `json:"App,omitempty"`
	AssignedUserId                   common.NullableString       `json:"AssignedUserId,omitempty"`
	AssignedUserName                 common.NullableString       `json:"AssignedUserName,omitempty"`
	Roles                            common.NullableString       `json:"Roles,omitempty"`
	Data                             *json.RawMessage            `json:"Data,omitempty"`
	ExtData                          common.NullableString       `json:"ExtData,omitempty"`
	Logs                             []CasesLog                  `json:"Logs,omitempty"`
	Comments                         []CasesComment              `json:"Comments,omitempty"`
	CommentsToAppend                 common.NullableList[string] `json:"CommentsToAppend,omitempty"`
	NumberOfRepetitions              *int64                      `json:"NumberOfRepetitions,omitempty"`
	SimilarityHash                   *string                     `json:"SimilarityHash,omitempty"`
	Tags                             common.NullableList[string] `json:"Tags,omitempty"`
	FileAttachments                  []CasesFileAttachment       `json:"FileAttachments,omitempty"`
	ReminderPeriotsAsHour            common.NullableString       `json:"ReminderPeriotsAsHour,omitempty"`
	ReminderEmailCount               *int64                      `json:"ReminderEmailCount,omitempty"`
	DataVector                       common.NullableString       `json:"DataVector,omitempty"`
	FilterTags                       common.NullableList[string] `json:"FilterTags,omitempty"`
	MappedFields                     *json.RawMessage            `json:"MappedFields,omitempty"`
	SourceType                       common.NullableString       `json:"SourceType,omitempty"`
	PriorityValue                    common.NullableString       `json:"PriorityValue,omitempty"`
	RemainingTimeAsHour              *int64                      `json:"RemainingTimeAsHour,omitempty"`
	WaitingTimeAsHour                *int64                      `json:"WaitingTimeAsHour,omitempty"`
	MitreData                        NullableCasesMitreData      `json:"MitreData,omitempty"`
	AssetList                        common.NullableList[string] `json:"AssetList,omitempty"`
	Approvals                        common.NullableList[string] `json:"Approvals,omitempty"`
	TenantID                         *string                     `json:"TenantID,omitempty"`
	ID                               *string                     `json:"_id,omitempty"`
	Timestamp                        *string                     `json:"Timestamp,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
