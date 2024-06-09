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
	Roles                            common.NullableList[string] `json:"Roles,omitempty"`
	Data                             *json.RawMessage            `json:"Data,omitempty"`
	ExtData                          common.NullableString       `json:"ExtData,omitempty"`
	Logs                             []CasesLog                  `json:"Logs,omitempty"`
	Comments                         []CasesComment              `json:"Comments,omitempty"`
	CommentsToAppend                 common.NullableList[string] `json:"CommentsToAppend,omitempty"`
	NumberOfRepetitions              *int64                      `json:"NumberOfRepetitions,omitempty"`
	SimilarityHash                   *string                     `json:"SimilarityHash,omitempty"`
	Tags                             common.NullableList[string] `json:"Tags,omitempty"`
	FileAttachments                  []CasesFileAttachment       `json:"FileAttachments,omitempty"`
	ReminderPeriotsAsHour            common.NullableInt64        `json:"ReminderPeriotsAsHour,omitempty"`
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

// NewCasesItem creates a new CasesItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCasesItem() *CasesItem {
	this := CasesItem{}
	return &this
}

// NewCasesItemWithDefaults creates a new CasesItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCasesItemWithDefaults() *CasesItem {
	this := CasesItem{}
	return &this
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *CasesItem) GetCaseId() string {
	if o == nil || o.CaseId == nil {
		var ret string
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetCaseIdOk() (*string, bool) {
	if o == nil || o.CaseId == nil {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *CasesItem) HasCaseId() bool {
	return o != nil && o.CaseId != nil
}

// SetCaseId gets a reference to the given string and assigns it to the CaseId field.
func (o *CasesItem) SetCaseId(v string) {
	o.CaseId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CasesItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CasesItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CasesItem) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a Description has been set.
func (o *CasesItem) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *CasesItem) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *CasesItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnSetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *CasesItem) UnSetDescription() {
	o.Description.UnSet()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CasesItem) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CasesItem) HasOwner() bool {
	return o != nil && o.Owner != nil
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *CasesItem) SetOwner(v string) {
	o.Owner = &v
}

// GetOccuredDate returns the OccuredDate field value if set, zero value otherwise.
func (o *CasesItem) GetOccuredDate() string {
	if o == nil || o.OccuredDate == nil {
		var ret string
		return ret
	}
	return *o.OccuredDate
}

// GetOccuredDateOk returns a tuple with the OccuredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetOccuredDateOk() (*string, bool) {
	if o == nil || o.OccuredDate == nil {
		return nil, false
	}
	return o.OccuredDate, true
}

// HasOccuredDate returns a boolean if a field has been set.
func (o *CasesItem) HasOccuredDate() bool {
	return o != nil && o.OccuredDate != nil
}

// SetOccuredDate gets a reference to the given string and assigns it to the OccuredDate field.
func (o *CasesItem) SetOccuredDate(v string) {
	o.OccuredDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *CasesItem) GetCreatedDate() string {
	if o == nil || o.CreatedDate == nil {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetCreatedDateOk() (*string, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CasesItem) HasCreatedDate() bool {
	return o != nil && o.CreatedDate != nil
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *CasesItem) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *CasesItem) GetModifiedDate() string {
	if o == nil || o.ModifiedDate == nil {
		var ret string
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetModifiedDateOk() (*string, bool) {
	if o == nil || o.ModifiedDate == nil {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *CasesItem) HasModifiedDate() bool {
	return o != nil && o.ModifiedDate != nil
}

// SetModifiedDate gets a reference to the given string and assigns it to the ModifiedDate field.
func (o *CasesItem) SetModifiedDate(v string) {
	o.ModifiedDate = &v
}

// GetReminderDate returns the ReminderDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetReminderDate() string {
	if o == nil || o.ReminderDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReminderDate.Get()
}

// GetReminderDateOk returns a tuple with the ReminderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetReminderDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReminderDate.Get(), o.ReminderDate.IsSet()
}

// HasReminderDate returns a boolean if a ReminderDate has been set.
func (o *CasesItem) HasReminderDate() bool {
	return o != nil && o.ReminderDate.IsSet()
}

// SetReminderDate gets a reference to the given datadog.NullableString and assigns it to the ReminderDate field.
func (o *CasesItem) SetReminderDate(v string) {
	o.ReminderDate.Set(&v)
}

// SetReminderDateNil sets the value for ReminderDate to be an explicit nil.
func (o *CasesItem) SetReminderDateNil() {
	o.ReminderDate.Set(nil)
}

// UnSetReminderDate ensures that no value is present for ReminderDate, not even an explicit nil.
func (o *CasesItem) UnSetReminderDate() {
	o.ReminderDate.UnSet()
}

// GetExpectedFinishTime returns the ExpectedFinishTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetExpectedFinishTime() string {
	if o == nil || o.ExpectedFinishTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExpectedFinishTime.Get()
}

// GetExpectedFinishTimeOk returns a tuple with the ExpectedFinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetExpectedFinishTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedFinishTime.Get(), o.ExpectedFinishTime.IsSet()
}

// HasExpectedFinishTime returns a boolean if a ExpectedFinishTime has been set.
func (o *CasesItem) HasExpectedFinishTime() bool {
	return o != nil && o.ExpectedFinishTime.IsSet()
}

// SetExpectedFinishTime gets a reference to the given datadog.NullableString and assigns it to the ExpectedFinishTime field.
func (o *CasesItem) SetExpectedFinishTime(v string) {
	o.ExpectedFinishTime.Set(&v)
}

// SetExpectedFinishTimeNil sets the value for ExpectedFinishTime to be an explicit nil.
func (o *CasesItem) SetExpectedFinishTimeNil() {
	o.ExpectedFinishTime.Set(nil)
}

// UnSetExpectedFinishTime ensures that no value is present for ExpectedFinishTime, not even an explicit nil.
func (o *CasesItem) UnSetExpectedFinishTime() {
	o.ExpectedFinishTime.UnSet()
}

// GetRealFinishTime returns the RealFinishTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetRealFinishTime() string {
	if o == nil || o.RealFinishTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.RealFinishTime.Get()
}

// GetRealFinishTimeOk returns a tuple with the RealFinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetRealFinishTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RealFinishTime.Get(), o.RealFinishTime.IsSet()
}

// HasRealFinishTime returns a boolean if a RealFinishTime has been set.
func (o *CasesItem) HasRealFinishTime() bool {
	return o != nil && o.RealFinishTime.IsSet()
}

// SetRealFinishTime gets a reference to the given datadog.NullableString and assigns it to the RealFinishTime field.
func (o *CasesItem) SetRealFinishTime(v string) {
	o.RealFinishTime.Set(&v)
}

// SetRealFinishTimeNil sets the value for RealFinishTime to be an explicit nil.
func (o *CasesItem) SetRealFinishTimeNil() {
	o.RealFinishTime.Set(nil)
}

// UnSetRealFinishTime ensures that no value is present for RealFinishTime, not even an explicit nil.
func (o *CasesItem) UnSetRealFinishTime() {
	o.RealFinishTime.UnSet()
}

// GetLastSentReminderEmailDate returns the LastSentReminderEmailDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetLastSentReminderEmailDate() string {
	if o == nil || o.LastSentReminderEmailDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastSentReminderEmailDate.Get()
}

// GetLastSentReminderEmailDateOk returns a tuple with the LastSentReminderEmailDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetLastSentReminderEmailDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSentReminderEmailDate.Get(), o.LastSentReminderEmailDate.IsSet()
}

// HasLastSentReminderEmailDate returns a boolean if a LastSentReminderEmailDate has been set.
func (o *CasesItem) HasLastSentReminderEmailDate() bool {
	return o != nil && o.LastSentReminderEmailDate.IsSet()
}

// SetLastSentReminderEmailDate gets a reference to the given datadog.NullableString and assigns it to the LastSentReminderEmailDate field.
func (o *CasesItem) SetLastSentReminderEmailDate(v string) {
	o.LastSentReminderEmailDate.Set(&v)
}

// SetLastSentReminderEmailDateNil sets the value for LastSentReminderEmailDate to be an explicit nil.
func (o *CasesItem) SetLastSentReminderEmailDateNil() {
	o.LastSentReminderEmailDate.Set(nil)
}

// UnSetLastSentReminderEmailDate ensures that no value is present for LastSentReminderEmailDate, not even an explicit nil.
func (o *CasesItem) UnSetLastSentReminderEmailDate() {
	o.LastSentReminderEmailDate.UnSet()
}

// GetPriorityValueLastCalculationDate returns the PriorityValueLastCalculationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetPriorityValueLastCalculationDate() string {
	if o == nil || o.PriorityValueLastCalculationDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PriorityValueLastCalculationDate.Get()
}

// GetPriorityValueLastCalculationDateOk returns a tuple with the PriorityValueLastCalculationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetPriorityValueLastCalculationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriorityValueLastCalculationDate.Get(), o.PriorityValueLastCalculationDate.IsSet()
}

// HasPriorityValueLastCalculationDate returns a boolean if a PriorityValueLastCalculationDate has been set.
func (o *CasesItem) HasPriorityValueLastCalculationDate() bool {
	return o != nil && o.PriorityValueLastCalculationDate.IsSet()
}

// SetPriorityValueLastCalculationDate gets a reference to the given datadog.NullableString and assigns it to the PriorityValueLastCalculationDate field.
func (o *CasesItem) SetPriorityValueLastCalculationDate(v string) {
	o.PriorityValueLastCalculationDate.Set(&v)
}

// SetPriorityValueLastCalculationDateNil sets the value for PriorityValueLastCalculationDate to be an explicit nil.
func (o *CasesItem) SetPriorityValueLastCalculationDateNil() {
	o.PriorityValueLastCalculationDate.Set(nil)
}

// UnSetPriorityValueLastCalculationDate ensures that no value is present for PriorityValueLastCalculationDate, not even an explicit nil.
func (o *CasesItem) UnSetPriorityValueLastCalculationDate() {
	o.PriorityValueLastCalculationDate.UnSet()
}

// GetCaseType returns the CaseType field value if set, zero value otherwise.
func (o *CasesItem) GetCaseType() string {
	if o == nil || o.CaseType == nil {
		var ret string
		return ret
	}
	return *o.CaseType
}

// GetCaseTypeOk returns a tuple with the CaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetCaseTypeOk() (*string, bool) {
	if o == nil || o.CaseType == nil {
		return nil, false
	}
	return o.CaseType, true
}

// HasCaseType returns a boolean if a field has been set.
func (o *CasesItem) HasCaseType() bool {
	return o != nil && o.CaseType != nil
}

// SetCaseType gets a reference to the given string and assigns it to the CaseType field.
func (o *CasesItem) SetCaseType(v string) {
	o.CaseType = &v
}

// GetCaseCategory returns the CaseCategory field value if set, zero value otherwise.
func (o *CasesItem) GetCaseCategory() string {
	if o == nil || o.CaseCategory == nil {
		var ret string
		return ret
	}
	return *o.CaseCategory
}

// GetCaseCategoryOk returns a tuple with the CaseCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetCaseCategoryOk() (*string, bool) {
	if o == nil || o.CaseCategory == nil {
		return nil, false
	}
	return o.CaseCategory, true
}

// HasCaseCategory returns a boolean if a field has been set.
func (o *CasesItem) HasCaseCategory() bool {
	return o != nil && o.CaseCategory != nil
}

// SetCaseCategory gets a reference to the given string and assigns it to the CaseCategory field.
func (o *CasesItem) SetCaseCategory(v string) {
	o.CaseCategory = &v
}

// GetCaseCategoryId returns the CaseCategoryId field value if set, zero value otherwise.
func (o *CasesItem) GetCaseCategoryId() string {
	if o == nil || o.CaseCategoryId == nil {
		var ret string
		return ret
	}
	return *o.CaseCategoryId
}

// GetCaseCategoryIdOk returns a tuple with the CaseCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetCaseCategoryIdOk() (*string, bool) {
	if o == nil || o.CaseCategoryId == nil {
		return nil, false
	}
	return o.CaseCategoryId, true
}

// HasCaseCategoryId returns a boolean if a field has been set.
func (o *CasesItem) HasCaseCategoryId() bool {
	return o != nil && o.CaseCategoryId != nil
}

// SetCaseCategoryId gets a reference to the given string and assigns it to the CaseCategoryId field.
func (o *CasesItem) SetCaseCategoryId(v string) {
	o.CaseCategoryId = &v
}

// GetParentCaseId returns the ParentCaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetParentCaseId() string {
	if o == nil || o.ParentCaseId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentCaseId.Get()
}

// GetParentCaseIdOk returns a tuple with the ParentCaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetParentCaseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentCaseId.Get(), o.ParentCaseId.IsSet()
}

// HasParentCaseId returns a boolean if a ParentCaseId has been set.
func (o *CasesItem) HasParentCaseId() bool {
	return o != nil && o.ParentCaseId.IsSet()
}

// SetParentCaseId gets a reference to the given datadog.NullableString and assigns it to the ParentCaseId field.
func (o *CasesItem) SetParentCaseId(v string) {
	o.ParentCaseId.Set(&v)
}

// SetParentCaseIdNil sets the value for ParentCaseId to be an explicit nil.
func (o *CasesItem) SetParentCaseIdNil() {
	o.ParentCaseId.Set(nil)
}

// UnSetParentCaseId ensures that no value is present for ParentCaseId, not even an explicit nil.
func (o *CasesItem) UnSetParentCaseId() {
	o.ParentCaseId.UnSet()
}

// GetParentCaseName returns the ParentCaseName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetParentCaseName() string {
	if o == nil || o.ParentCaseName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentCaseName.Get()
}

// GetParentCaseNameOk returns a tuple with the ParentCaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetParentCaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentCaseName.Get(), o.ParentCaseName.IsSet()
}

// HasParentCaseName returns a boolean if a ParentCaseName has been set.
func (o *CasesItem) HasParentCaseName() bool {
	return o != nil && o.ParentCaseName.IsSet()
}

// SetParentCaseName gets a reference to the given datadog.NullableString and assigns it to the ParentCaseName field.
func (o *CasesItem) SetParentCaseName(v string) {
	o.ParentCaseName.Set(&v)
}

// SetParentCaseNameNil sets the value for ParentCaseName to be an explicit nil.
func (o *CasesItem) SetParentCaseNameNil() {
	o.ParentCaseName.Set(nil)
}

// UnSetParentCaseName ensures that no value is present for ParentCaseName, not even an explicit nil.
func (o *CasesItem) UnSetParentCaseName() {
	o.ParentCaseName.UnSet()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CasesItem) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CasesItem) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CasesItem) SetState(v string) {
	o.State = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *CasesItem) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *CasesItem) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *CasesItem) SetSeverity(v string) {
	o.Severity = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *CasesItem) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *CasesItem) HasApp() bool {
	return o != nil && o.App != nil
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *CasesItem) SetApp(v string) {
	o.App = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetAssignedUserId() string {
	if o == nil || o.AssignedUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssignedUserId.Get()
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetAssignedUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedUserId.Get(), o.AssignedUserId.IsSet()
}

// HasAssignedUserId returns a boolean if a AssignedUserId has been set.
func (o *CasesItem) HasAssignedUserId() bool {
	return o != nil && o.AssignedUserId.IsSet()
}

// SetAssignedUserId gets a reference to the given datadog.NullableString and assigns it to the AssignedUserId field.
func (o *CasesItem) SetAssignedUserId(v string) {
	o.AssignedUserId.Set(&v)
}

// SetAssignedUserIdNil sets the value for AssignedUserId to be an explicit nil.
func (o *CasesItem) SetAssignedUserIdNil() {
	o.AssignedUserId.Set(nil)
}

// UnSetAssignedUserId ensures that no value is present for AssignedUserId, not even an explicit nil.
func (o *CasesItem) UnSetAssignedUserId() {
	o.AssignedUserId.UnSet()
}

// GetAssignedUserName returns the AssignedUserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetAssignedUserName() string {
	if o == nil || o.AssignedUserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssignedUserName.Get()
}

// GetAssignedUserNameOk returns a tuple with the AssignedUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetAssignedUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedUserName.Get(), o.AssignedUserName.IsSet()
}

// HasAssignedUserName returns a boolean if a AssignedUserName has been set.
func (o *CasesItem) HasAssignedUserName() bool {
	return o != nil && o.AssignedUserName.IsSet()
}

// SetAssignedUserName gets a reference to the given datadog.NullableString and assigns it to the AssignedUserName field.
func (o *CasesItem) SetAssignedUserName(v string) {
	o.AssignedUserName.Set(&v)
}

// SetAssignedUserNameNil sets the value for AssignedUserName to be an explicit nil.
func (o *CasesItem) SetAssignedUserNameNil() {
	o.AssignedUserName.Set(nil)
}

// UnSetAssignedUserName ensures that no value is present for AssignedUserName, not even an explicit nil.
func (o *CasesItem) UnSetAssignedUserName() {
	o.AssignedUserName.UnSet()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetRoles() []string {
	if o == nil || o.Roles.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Roles.Get()
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetRolesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles.Get(), o.Roles.IsSet()
}

// HasRoles returns a boolean if a Roles has been set.
func (o *CasesItem) HasRoles() bool {
	return o != nil && o.Roles.IsSet()
}

// SetRoles gets a reference to the given datadog.Nullable[]string and assigns it to the Roles field.
func (o *CasesItem) SetRoles(v []string) {
	o.Roles.Set(&v)
}

// SetRolesNil sets the value for Roles to be an explicit nil.
func (o *CasesItem) SetRolesNil() {
	o.Roles.Set(nil)
}

// UnSetRoles ensures that no value is present for Roles, not even an explicit nil.
func (o *CasesItem) UnSetRoles() {
	o.Roles.UnSet()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CasesItem) GetData() json.RawMessage {
	if o == nil || o.Data == nil {
		var ret json.RawMessage
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetDataOk() (*json.RawMessage, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CasesItem) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given json.RawMessage and assigns it to the Data field.
func (o *CasesItem) SetData(v json.RawMessage) {
	o.Data = &v
}

// GetExtData returns the ExtData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetExtData() string {
	if o == nil || o.ExtData.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtData.Get()
}

// GetExtDataOk returns a tuple with the ExtData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetExtDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtData.Get(), o.ExtData.IsSet()
}

// HasExtData returns a boolean if a ExtData has been set.
func (o *CasesItem) HasExtData() bool {
	return o != nil && o.ExtData.IsSet()
}

// SetExtData gets a reference to the given datadog.NullableString and assigns it to the ExtData field.
func (o *CasesItem) SetExtData(v string) {
	o.ExtData.Set(&v)
}

// SetExtDataNil sets the value for ExtData to be an explicit nil.
func (o *CasesItem) SetExtDataNil() {
	o.ExtData.Set(nil)
}

// UnSetExtData ensures that no value is present for ExtData, not even an explicit nil.
func (o *CasesItem) UnSetExtData() {
	o.ExtData.UnSet()
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *CasesItem) GetLogs() []CasesLog {
	if o == nil || o.Logs == nil {
		var ret []CasesLog
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetLogsOk() (*[]CasesLog, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return &o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *CasesItem) HasLogs() bool {
	return o != nil && o.Logs != nil
}

// SetLogs gets a reference to the given []CasesLog and assigns it to the Logs field.
func (o *CasesItem) SetLogs(v []CasesLog) {
	o.Logs = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *CasesItem) GetComments() []CasesComment {
	if o == nil || o.Comments == nil {
		var ret []CasesComment
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetCommentsOk() (*[]CasesComment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return &o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *CasesItem) HasComments() bool {
	return o != nil && o.Comments != nil
}

// SetComments gets a reference to the given []CasesComment and assigns it to the Comments field.
func (o *CasesItem) SetComments(v []CasesComment) {
	o.Comments = v
}

// GetCommentsToAppend returns the CommentsToAppend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetCommentsToAppend() []string {
	if o == nil || o.CommentsToAppend.Get() == nil {
		var ret []string
		return ret
	}
	return *o.CommentsToAppend.Get()
}

// GetCommentsToAppendOk returns a tuple with the CommentsToAppend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetCommentsToAppendOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommentsToAppend.Get(), o.CommentsToAppend.IsSet()
}

// HasCommentsToAppend returns a boolean if a CommentsToAppend has been set.
func (o *CasesItem) HasCommentsToAppend() bool {
	return o != nil && o.CommentsToAppend.IsSet()
}

// SetCommentsToAppend gets a reference to the given datadog.Nullable[]string and assigns it to the CommentsToAppend field.
func (o *CasesItem) SetCommentsToAppend(v []string) {
	o.CommentsToAppend.Set(&v)
}

// SetCommentsToAppendNil sets the value for CommentsToAppend to be an explicit nil.
func (o *CasesItem) SetCommentsToAppendNil() {
	o.CommentsToAppend.Set(nil)
}

// UnSetCommentsToAppend ensures that no value is present for CommentsToAppend, not even an explicit nil.
func (o *CasesItem) UnSetCommentsToAppend() {
	o.CommentsToAppend.UnSet()
}

// GetNumberOfRepetitions returns the NumberOfRepetitions field value if set, zero value otherwise.
func (o *CasesItem) GetNumberOfRepetitions() int64 {
	if o == nil || o.NumberOfRepetitions == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfRepetitions
}

// GetNumberOfRepetitionsOk returns a tuple with the NumberOfRepetitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetNumberOfRepetitionsOk() (*int64, bool) {
	if o == nil || o.NumberOfRepetitions == nil {
		return nil, false
	}
	return o.NumberOfRepetitions, true
}

// HasNumberOfRepetitions returns a boolean if a field has been set.
func (o *CasesItem) HasNumberOfRepetitions() bool {
	return o != nil && o.NumberOfRepetitions != nil
}

// SetNumberOfRepetitions gets a reference to the given int64 and assigns it to the NumberOfRepetitions field.
func (o *CasesItem) SetNumberOfRepetitions(v int64) {
	o.NumberOfRepetitions = &v
}

// GetSimilarityHash returns the SimilarityHash field value if set, zero value otherwise.
func (o *CasesItem) GetSimilarityHash() string {
	if o == nil || o.SimilarityHash == nil {
		var ret string
		return ret
	}
	return *o.SimilarityHash
}

// GetSimilarityHashOk returns a tuple with the SimilarityHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetSimilarityHashOk() (*string, bool) {
	if o == nil || o.SimilarityHash == nil {
		return nil, false
	}
	return o.SimilarityHash, true
}

// HasSimilarityHash returns a boolean if a field has been set.
func (o *CasesItem) HasSimilarityHash() bool {
	return o != nil && o.SimilarityHash != nil
}

// SetSimilarityHash gets a reference to the given string and assigns it to the SimilarityHash field.
func (o *CasesItem) SetSimilarityHash(v string) {
	o.SimilarityHash = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a Tags has been set.
func (o *CasesItem) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.Nullable[]string and assigns it to the Tags field.
func (o *CasesItem) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *CasesItem) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnSetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *CasesItem) UnSetTags() {
	o.Tags.UnSet()
}

// GetFileAttachments returns the FileAttachments field value if set, zero value otherwise.
func (o *CasesItem) GetFileAttachments() []CasesFileAttachment {
	if o == nil || o.FileAttachments == nil {
		var ret []CasesFileAttachment
		return ret
	}
	return o.FileAttachments
}

// GetFileAttachmentsOk returns a tuple with the FileAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetFileAttachmentsOk() (*[]CasesFileAttachment, bool) {
	if o == nil || o.FileAttachments == nil {
		return nil, false
	}
	return &o.FileAttachments, true
}

// HasFileAttachments returns a boolean if a field has been set.
func (o *CasesItem) HasFileAttachments() bool {
	return o != nil && o.FileAttachments != nil
}

// SetFileAttachments gets a reference to the given []CasesFileAttachment and assigns it to the FileAttachments field.
func (o *CasesItem) SetFileAttachments(v []CasesFileAttachment) {
	o.FileAttachments = v
}

// GetReminderPeriotsAsHour returns the ReminderPeriotsAsHour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetReminderPeriotsAsHour() int64 {
	if o == nil || o.ReminderPeriotsAsHour.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ReminderPeriotsAsHour.Get()
}

// GetReminderPeriotsAsHourOk returns a tuple with the ReminderPeriotsAsHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetReminderPeriotsAsHourOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReminderPeriotsAsHour.Get(), o.ReminderPeriotsAsHour.IsSet()
}

// HasReminderPeriotsAsHour returns a boolean if a ReminderPeriotsAsHour has been set.
func (o *CasesItem) HasReminderPeriotsAsHour() bool {
	return o != nil && o.ReminderPeriotsAsHour.IsSet()
}

// SetReminderPeriotsAsHour gets a reference to the given datadog.NullableInt64 and assigns it to the ReminderPeriotsAsHour field.
func (o *CasesItem) SetReminderPeriotsAsHour(v int64) {
	o.ReminderPeriotsAsHour.Set(&v)
}

// SetReminderPeriotsAsHourNil sets the value for ReminderPeriotsAsHour to be an explicit nil.
func (o *CasesItem) SetReminderPeriotsAsHourNil() {
	o.ReminderPeriotsAsHour.Set(nil)
}

// UnSetReminderPeriotsAsHour ensures that no value is present for ReminderPeriotsAsHour, not even an explicit nil.
func (o *CasesItem) UnSetReminderPeriotsAsHour() {
	o.ReminderPeriotsAsHour.UnSet()
}

// GetReminderEmailCount returns the ReminderEmailCount field value if set, zero value otherwise.
func (o *CasesItem) GetReminderEmailCount() int64 {
	if o == nil || o.ReminderEmailCount == nil {
		var ret int64
		return ret
	}
	return *o.ReminderEmailCount
}

// GetReminderEmailCountOk returns a tuple with the ReminderEmailCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetReminderEmailCountOk() (*int64, bool) {
	if o == nil || o.ReminderEmailCount == nil {
		return nil, false
	}
	return o.ReminderEmailCount, true
}

// HasReminderEmailCount returns a boolean if a field has been set.
func (o *CasesItem) HasReminderEmailCount() bool {
	return o != nil && o.ReminderEmailCount != nil
}

// SetReminderEmailCount gets a reference to the given int64 and assigns it to the ReminderEmailCount field.
func (o *CasesItem) SetReminderEmailCount(v int64) {
	o.ReminderEmailCount = &v
}

// GetDataVector returns the DataVector field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetDataVector() string {
	if o == nil || o.DataVector.Get() == nil {
		var ret string
		return ret
	}
	return *o.DataVector.Get()
}

// GetDataVectorOk returns a tuple with the DataVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetDataVectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataVector.Get(), o.DataVector.IsSet()
}

// HasDataVector returns a boolean if a DataVector has been set.
func (o *CasesItem) HasDataVector() bool {
	return o != nil && o.DataVector.IsSet()
}

// SetDataVector gets a reference to the given datadog.NullableString and assigns it to the DataVector field.
func (o *CasesItem) SetDataVector(v string) {
	o.DataVector.Set(&v)
}

// SetDataVectorNil sets the value for DataVector to be an explicit nil.
func (o *CasesItem) SetDataVectorNil() {
	o.DataVector.Set(nil)
}

// UnSetDataVector ensures that no value is present for DataVector, not even an explicit nil.
func (o *CasesItem) UnSetDataVector() {
	o.DataVector.UnSet()
}

// GetFilterTags returns the FilterTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetFilterTags() []string {
	if o == nil || o.FilterTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.FilterTags.Get()
}

// GetFilterTagsOk returns a tuple with the FilterTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetFilterTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterTags.Get(), o.FilterTags.IsSet()
}

// HasFilterTags returns a boolean if a FilterTags has been set.
func (o *CasesItem) HasFilterTags() bool {
	return o != nil && o.FilterTags.IsSet()
}

// SetFilterTags gets a reference to the given datadog.Nullable[]string and assigns it to the FilterTags field.
func (o *CasesItem) SetFilterTags(v []string) {
	o.FilterTags.Set(&v)
}

// SetFilterTagsNil sets the value for FilterTags to be an explicit nil.
func (o *CasesItem) SetFilterTagsNil() {
	o.FilterTags.Set(nil)
}

// UnSetFilterTags ensures that no value is present for FilterTags, not even an explicit nil.
func (o *CasesItem) UnSetFilterTags() {
	o.FilterTags.UnSet()
}

// GetMappedFields returns the MappedFields field value if set, zero value otherwise.
func (o *CasesItem) GetMappedFields() json.RawMessage {
	if o == nil || o.MappedFields == nil {
		var ret json.RawMessage
		return ret
	}
	return *o.MappedFields
}

// GetMappedFieldsOk returns a tuple with the MappedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetMappedFieldsOk() (*json.RawMessage, bool) {
	if o == nil || o.MappedFields == nil {
		return nil, false
	}
	return o.MappedFields, true
}

// HasMappedFields returns a boolean if a field has been set.
func (o *CasesItem) HasMappedFields() bool {
	return o != nil && o.MappedFields != nil
}

// SetMappedFields gets a reference to the given json.RawMessage and assigns it to the MappedFields field.
func (o *CasesItem) SetMappedFields(v json.RawMessage) {
	o.MappedFields = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetSourceType() string {
	if o == nil || o.SourceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceType.Get()
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceType.Get(), o.SourceType.IsSet()
}

// HasSourceType returns a boolean if a SourceType has been set.
func (o *CasesItem) HasSourceType() bool {
	return o != nil && o.SourceType.IsSet()
}

// SetSourceType gets a reference to the given datadog.NullableString and assigns it to the SourceType field.
func (o *CasesItem) SetSourceType(v string) {
	o.SourceType.Set(&v)
}

// SetSourceTypeNil sets the value for SourceType to be an explicit nil.
func (o *CasesItem) SetSourceTypeNil() {
	o.SourceType.Set(nil)
}

// UnSetSourceType ensures that no value is present for SourceType, not even an explicit nil.
func (o *CasesItem) UnSetSourceType() {
	o.SourceType.UnSet()
}

// GetPriorityValue returns the PriorityValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetPriorityValue() string {
	if o == nil || o.PriorityValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.PriorityValue.Get()
}

// GetPriorityValueOk returns a tuple with the PriorityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetPriorityValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriorityValue.Get(), o.PriorityValue.IsSet()
}

// HasPriorityValue returns a boolean if a PriorityValue has been set.
func (o *CasesItem) HasPriorityValue() bool {
	return o != nil && o.PriorityValue.IsSet()
}

// SetPriorityValue gets a reference to the given datadog.NullableString and assigns it to the PriorityValue field.
func (o *CasesItem) SetPriorityValue(v string) {
	o.PriorityValue.Set(&v)
}

// SetPriorityValueNil sets the value for PriorityValue to be an explicit nil.
func (o *CasesItem) SetPriorityValueNil() {
	o.PriorityValue.Set(nil)
}

// UnSetPriorityValue ensures that no value is present for PriorityValue, not even an explicit nil.
func (o *CasesItem) UnSetPriorityValue() {
	o.PriorityValue.UnSet()
}

// GetRemainingTimeAsHour returns the RemainingTimeAsHour field value if set, zero value otherwise.
func (o *CasesItem) GetRemainingTimeAsHour() int64 {
	if o == nil || o.RemainingTimeAsHour == nil {
		var ret int64
		return ret
	}
	return *o.RemainingTimeAsHour
}

// GetRemainingTimeAsHourOk returns a tuple with the RemainingTimeAsHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetRemainingTimeAsHourOk() (*int64, bool) {
	if o == nil || o.RemainingTimeAsHour == nil {
		return nil, false
	}
	return o.RemainingTimeAsHour, true
}

// HasRemainingTimeAsHour returns a boolean if a field has been set.
func (o *CasesItem) HasRemainingTimeAsHour() bool {
	return o != nil && o.RemainingTimeAsHour != nil
}

// SetRemainingTimeAsHour gets a reference to the given int64 and assigns it to the RemainingTimeAsHour field.
func (o *CasesItem) SetRemainingTimeAsHour(v int64) {
	o.RemainingTimeAsHour = &v
}

// GetWaitingTimeAsHour returns the WaitingTimeAsHour field value if set, zero value otherwise.
func (o *CasesItem) GetWaitingTimeAsHour() int64 {
	if o == nil || o.WaitingTimeAsHour == nil {
		var ret int64
		return ret
	}
	return *o.WaitingTimeAsHour
}

// GetWaitingTimeAsHourOk returns a tuple with the WaitingTimeAsHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetWaitingTimeAsHourOk() (*int64, bool) {
	if o == nil || o.WaitingTimeAsHour == nil {
		return nil, false
	}
	return o.WaitingTimeAsHour, true
}

// HasWaitingTimeAsHour returns a boolean if a field has been set.
func (o *CasesItem) HasWaitingTimeAsHour() bool {
	return o != nil && o.WaitingTimeAsHour != nil
}

// SetWaitingTimeAsHour gets a reference to the given int64 and assigns it to the WaitingTimeAsHour field.
func (o *CasesItem) SetWaitingTimeAsHour(v int64) {
	o.WaitingTimeAsHour = &v
}

// GetMitreData returns the MitreData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetMitreData() CasesMitreData {
	if o == nil || o.MitreData.Get() == nil {
		var ret CasesMitreData
		return ret
	}
	return *o.MitreData.Get()
}

// GetMitreDataOk returns a tuple with the MitreData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetMitreDataOk() (*CasesMitreData, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitreData.Get(), o.MitreData.IsSet()
}

// HasMitreData returns a boolean if a MitreData has been set.
func (o *CasesItem) HasMitreData() bool {
	return o != nil && o.MitreData.IsSet()
}

// SetMitreData gets a reference to the given datadog.NullableString and assigns it to the MitreData field.
func (o *CasesItem) SetMitreData(v CasesMitreData) {
	o.MitreData.Set(&v)
}

// SetMitreDataNil sets the value for MitreData to be an explicit nil.
func (o *CasesItem) SetMitreDataNil() {
	o.MitreData.Set(nil)
}

// UnSetMitreData ensures that no value is present for MitreData, not even an explicit nil.
func (o *CasesItem) UnSetMitreData() {
	o.MitreData.UnSet()
}

// GetAssetList returns the AssetList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetAssetList() []string {
	if o == nil || o.AssetList.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AssetList.Get()
}

// GetAssetListOk returns a tuple with the AssetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetAssetListOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetList.Get(), o.AssetList.IsSet()
}

// HasAssetList returns a boolean if a AssetList has been set.
func (o *CasesItem) HasAssetList() bool {
	return o != nil && o.AssetList.IsSet()
}

// SetAssetList gets a reference to the given datadog.Nullable[]string and assigns it to the AssetList field.
func (o *CasesItem) SetAssetList(v []string) {
	o.AssetList.Set(&v)
}

// SetAssetListNil sets the value for AssetList to be an explicit nil.
func (o *CasesItem) SetAssetListNil() {
	o.AssetList.Set(nil)
}

// UnSetAssetList ensures that no value is present for AssetList, not even an explicit nil.
func (o *CasesItem) UnSetAssetList() {
	o.AssetList.UnSet()
}

// GetApprovals returns the Approvals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasesItem) GetApprovals() []string {
	if o == nil || o.Approvals.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Approvals.Get()
}

// GetApprovalsOk returns a tuple with the Approvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CasesItem) GetApprovalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approvals.Get(), o.Approvals.IsSet()
}

// HasApprovals returns a boolean if a Approvals has been set.
func (o *CasesItem) HasApprovals() bool {
	return o != nil && o.Approvals.IsSet()
}

// SetApprovals gets a reference to the given datadog.Nullable[]string and assigns it to the Approvals field.
func (o *CasesItem) SetApprovals(v []string) {
	o.Approvals.Set(&v)
}

// SetApprovalsNil sets the value for Approvals to be an explicit nil.
func (o *CasesItem) SetApprovalsNil() {
	o.Approvals.Set(nil)
}

// UnSetApprovals ensures that no value is present for Approvals, not even an explicit nil.
func (o *CasesItem) UnSetApprovals() {
	o.Approvals.UnSet()
}

// GetTenantID returns the TenantID field value if set, zero value otherwise.
func (o *CasesItem) GetTenantID() string {
	if o == nil || o.TenantID == nil {
		var ret string
		return ret
	}
	return *o.TenantID
}

// GetTenantIDOk returns a tuple with the TenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetTenantIDOk() (*string, bool) {
	if o == nil || o.TenantID == nil {
		return nil, false
	}
	return o.TenantID, true
}

// HasTenantID returns a boolean if a field has been set.
func (o *CasesItem) HasTenantID() bool {
	return o != nil && o.TenantID != nil
}

// SetTenantID gets a reference to the given string and assigns it to the TenantID field.
func (o *CasesItem) SetTenantID(v string) {
	o.TenantID = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *CasesItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *CasesItem) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *CasesItem) SetID(v string) {
	o.ID = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CasesItem) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesItem) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CasesItem) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *CasesItem) SetTimestamp(v string) {
	o.Timestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CasesItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CaseId != nil {
		toSerialize["CaseId"] = o.CaseId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}
	if o.OccuredDate != nil {
		toSerialize["OccuredDate"] = o.OccuredDate
	}
	if o.CreatedDate != nil {
		toSerialize["CreatedDate"] = o.CreatedDate
	}
	if o.ModifiedDate != nil {
		toSerialize["ModifiedDate"] = o.ModifiedDate
	}
	if o.ReminderDate.IsSet() {
		toSerialize["ReminderDate"] = o.ReminderDate.Get()
	}
	if o.ExpectedFinishTime.IsSet() {
		toSerialize["ExpectedFinishTime"] = o.ExpectedFinishTime.Get()
	}
	if o.RealFinishTime.IsSet() {
		toSerialize["RealFinishTime"] = o.RealFinishTime.Get()
	}
	if o.LastSentReminderEmailDate.IsSet() {
		toSerialize["LastSentReminderEmailDate"] = o.LastSentReminderEmailDate.Get()
	}
	if o.PriorityValueLastCalculationDate.IsSet() {
		toSerialize["PriorityValueLastCalculationDate"] = o.PriorityValueLastCalculationDate.Get()
	}
	if o.CaseType != nil {
		toSerialize["CaseType"] = o.CaseType
	}
	if o.CaseCategory != nil {
		toSerialize["CaseCategory"] = o.CaseCategory
	}
	if o.CaseCategoryId != nil {
		toSerialize["CaseCategoryId"] = o.CaseCategoryId
	}
	if o.ParentCaseId.IsSet() {
		toSerialize["ParentCaseId"] = o.ParentCaseId.Get()
	}
	if o.ParentCaseName.IsSet() {
		toSerialize["ParentCaseName"] = o.ParentCaseName.Get()
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.App != nil {
		toSerialize["App"] = o.App
	}
	if o.AssignedUserId.IsSet() {
		toSerialize["AssignedUserId"] = o.AssignedUserId.Get()
	}
	if o.AssignedUserName.IsSet() {
		toSerialize["AssignedUserName"] = o.AssignedUserName.Get()
	}
	if o.Roles.IsSet() {
		toSerialize["Roles"] = o.Roles.Get()
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if o.ExtData.IsSet() {
		toSerialize["ExtData"] = o.ExtData.Get()
	}
	if o.Logs != nil {
		toSerialize["Logs"] = o.Logs
	}
	if o.Comments != nil {
		toSerialize["Comments"] = o.Comments
	}
	if o.CommentsToAppend.IsSet() {
		toSerialize["CommentsToAppend"] = o.CommentsToAppend.Get()
	}
	if o.NumberOfRepetitions != nil {
		toSerialize["NumberOfRepetitions"] = o.NumberOfRepetitions
	}
	if o.SimilarityHash != nil {
		toSerialize["SimilarityHash"] = o.SimilarityHash
	}
	if o.Tags.IsSet() {
		toSerialize["Tags"] = o.Tags.Get()
	}
	if o.FileAttachments != nil {
		toSerialize["FileAttachments"] = o.FileAttachments
	}
	if o.ReminderPeriotsAsHour.IsSet() {
		toSerialize["ReminderPeriotsAsHour"] = o.ReminderPeriotsAsHour.Get()
	}
	if o.ReminderEmailCount != nil {
		toSerialize["ReminderEmailCount"] = o.ReminderEmailCount
	}
	if o.DataVector.IsSet() {
		toSerialize["DataVector"] = o.DataVector.Get()
	}
	if o.FilterTags.IsSet() {
		toSerialize["FilterTags"] = o.FilterTags.Get()
	}
	if o.MappedFields != nil {
		toSerialize["MappedFields"] = o.MappedFields
	}
	if o.SourceType.IsSet() {
		toSerialize["SourceType"] = o.SourceType.Get()
	}
	if o.PriorityValue.IsSet() {
		toSerialize["PriorityValue"] = o.PriorityValue.Get()
	}
	if o.RemainingTimeAsHour != nil {
		toSerialize["RemainingTimeAsHour"] = o.RemainingTimeAsHour
	}
	if o.WaitingTimeAsHour != nil {
		toSerialize["WaitingTimeAsHour"] = o.WaitingTimeAsHour
	}
	if o.MitreData.IsSet() {
		toSerialize["MitreData"] = o.MitreData.Get()
	}
	if o.AssetList.IsSet() {
		toSerialize["AssetList"] = o.AssetList.Get()
	}
	if o.Approvals.IsSet() {
		toSerialize["Approvals"] = o.Approvals.Get()
	}
	if o.TenantID != nil {
		toSerialize["TenantID"] = o.TenantID
	}
	if o.ID != nil {
		toSerialize["_id"] = o.ID
	}
	if o.Timestamp != nil {
		toSerialize["TenantID"] = o.TenantID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *CasesItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
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
		Roles                            common.NullableList[string] `json:"Roles,omitempty"`
		Data                             *json.RawMessage            `json:"Data,omitempty"`
		ExtData                          common.NullableString       `json:"ExtData,omitempty"`
		Logs                             []CasesLog                  `json:"Logs,omitempty"`
		Comments                         []CasesComment              `json:"Comments,omitempty"`
		CommentsToAppend                 common.NullableList[string] `json:"CommentsToAppend,omitempty"`
		NumberOfRepetitions              *int64                      `json:"NumberOfRepetitions,omitempty"`
		SimilarityHash                   *string                     `json:"SimilarityHash,omitempty"`
		Tags                             common.NullableList[string] `json:"Tags,omitempty"`
		FileAttachments                  []CasesFileAttachment       `json:"FileAttachments,omitempty"`
		ReminderPeriotsAsHour            common.NullableInt64        `json:"ReminderPeriotsAsHour,omitempty"`
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
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Items", "FailedItems", "SuccessItems"})
	} else {
		return err
	}

	o.CaseId = all.CaseId
	o.Name = all.Name
	o.Description = all.Description
	o.Owner = all.Owner
	o.OccuredDate = all.OccuredDate
	o.CreatedDate = all.CreatedDate
	o.ModifiedDate = all.ModifiedDate
	o.ReminderDate = all.ReminderDate
	o.ExpectedFinishTime = all.ExpectedFinishTime
	o.RealFinishTime = all.RealFinishTime
	o.LastSentReminderEmailDate = all.LastSentReminderEmailDate
	o.PriorityValueLastCalculationDate = all.PriorityValueLastCalculationDate
	o.CaseType = all.CaseType
	o.CaseCategory = all.CaseCategory
	o.CaseCategoryId = all.CaseCategoryId
	o.ParentCaseId = all.ParentCaseId
	o.ParentCaseName = all.ParentCaseName
	o.State = all.State
	o.Severity = all.Severity
	o.App = all.App
	o.AssignedUserId = all.AssignedUserId
	o.AssignedUserName = all.AssignedUserName
	o.Roles = all.Roles
	o.Data = all.Data
	o.ExtData = all.ExtData
	o.Logs = all.Logs
	o.Comments = all.Comments
	o.CommentsToAppend = all.CommentsToAppend
	o.NumberOfRepetitions = all.NumberOfRepetitions
	o.SimilarityHash = all.SimilarityHash
	o.Tags = all.Tags
	o.FileAttachments = all.FileAttachments
	o.ReminderPeriotsAsHour = all.ReminderPeriotsAsHour
	o.ReminderEmailCount = all.ReminderEmailCount
	o.DataVector = all.DataVector
	o.FilterTags = all.FilterTags
	o.MappedFields = all.MappedFields
	o.SourceType = all.SourceType
	o.PriorityValue = all.PriorityValue
	o.RemainingTimeAsHour = all.RemainingTimeAsHour
	o.WaitingTimeAsHour = all.WaitingTimeAsHour
	o.MitreData = all.MitreData
	o.AssetList = all.AssetList
	o.Approvals = all.Approvals
	o.TenantID = all.TenantID
	o.ID = all.ID
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
