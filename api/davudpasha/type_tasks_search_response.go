package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// TasksSearchResponse represents the structure for task search responses.
type TasksSearchResponse struct {
	// TaskID represents the identifier of the task.
	TaskID *string `json:"TaskID,omitempty"`
	// ModelType specifies the type of the model associated with the task.
	ModelType *string `json:"ModelType,omitempty"`
	// Name specifies the name of the task.
	Name *string `json:"Name,omitempty"`
	// IsPrivate indicates whether the task is private.
	IsPrivate *bool `json:"IsPrivate,omitempty"`
	// ModelID represents the identifier of the model associated with the task.
	ModelID *string `json:"ModelID,omitempty"`
	// OperationType specifies the type of operation the task performs.
	OperationType *string `json:"OperationType,omitempty"`
	// TaskRunID represents the identifier of the task run.
	TaskRunID common.NullableString `json:"TaskRunID,omitempty"`
	// ScheduleConfig specifies the schedule configuration for the task.
	ScheduleConfig *ScheduleConfig `json:"ScheduleConfig,omitempty"`
	// Active indicates whether the task is active.
	Active *bool `json:"Active,omitempty"`
	// NextPlannedTime specifies the next planned execution time for the task.
	NextPlannedTime *string `json:"NextPlannedTime,omitempty"`
	// LastExecutionStartTime specifies the start time of the last execution of the task.
	LastExecutionStartTime common.NullableString `json:"LastExecutionStartTime,omitempty"`
	// LastSuccesufulExecutionTime specifies the time of the last successful execution of the task.
	LastSuccesufulExecutionTime common.NullableString `json:"LastSuccesufulExecutionTime,omitempty"`
	// LastFailedExecutionTime specifies the time of the last failed execution of the task.
	LastFailedExecutionTime common.NullableString `json:"LastFailedExecutionTime,omitempty"`
	// CreatedDate specifies the creation date of the task.
	CreatedDate *string `json:"CreatedDate,omitempty"`
	// JwtLifeTimeSeconds specifies the JWT lifetime in seconds.
	JwtLifeTimeSeconds common.NullableString `json:"JwtLifeTimeSeconds,omitempty"`
	// ModuleFilter specifies the module filter for the task.
	ModuleFilter *string `json:"ModuleFilter,omitempty"`
	// RemoteInterfaceName specifies the remote interface name for the task.
	RemoteInterfaceName *string `json:"RemoteInterfaceName,omitempty"`
	// RemoteMethodName specifies the remote method name for the task.
	RemoteMethodName *string `json:"RemoteMethodName,omitempty"`
	// Data specifies additional data associated with the task.
	Data common.NullableString `json:"Data,omitempty"`
	// MicroserviceAddress specifies the address of the microservice associated with the task.
	MicroserviceAddress common.NullableString `json:"MicroserviceAddress,omitempty"`
	// Result specifies the result of the task execution.
	Result common.NullableString `json:"Result,omitempty"`
	// DisableRetryIfFail indicates whether retry should be disabled if the task fails.
	DisableRetryIfFail *bool `json:"DisableRetryIfFail,omitempty"`
	// ExtraObject specifies additional object data associated with the task.
	ExtraObject common.NullableString `json:"ExtraObject,omitempty"`
	// ExecutorName specifies the name of the executor handling the task.
	ExecutorName common.NullableString `json:"ExecutorName,omitempty"`
	// ExecutorType specifies the type of the executor handling the task.
	ExecutorType *string `json:"ExecutorType,omitempty"`
	// TenantID represents the identifier of the tenant associated with the task.
	TenantID *string `json:"TenantID,omitempty"`
	// ModuleGUID specifies the GUID of the module associated with the task.
	ModuleGUID common.NullableString `json:"ModuleGuid,omitempty"`
	// OwnerType specifies the type of the owner associated with the task.
	OwnerType *string `json:"OwnerType,omitempty"`
	// OwnerName specifies the name of the owner associated with the task.
	OwnerName *string `json:"OwnerName,omitempty"`
	// ApplicationName specifies the name of the application associated with the task.
	ApplicationName *string `json:"ApplicationName,omitempty"`
	// RunForAllTenants indicates whether the task runs for all tenants.
	RunForAllTenants *bool `json:"RunForAllTenants,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTasksSearchResponse creates a new TasksSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTasksSearchResponse() *TasksSearchResponse {
	this := TasksSearchResponse{}
	return &this
}

// NewTasksSearchResponseWithDefaults creates a new TasksSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTasksSearchResponseWithDefaults() *TasksSearchResponse {
	this := TasksSearchResponse{}
	return &this
}

// GetTaskID returns the TaskID field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetTaskID() string {
	if o == nil || o.TaskID == nil {
		var ret string
		return ret
	}
	return *o.TaskID
}

// GetTaskIDOk returns a tuple with the TaskID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetTaskIDOk() (*string, bool) {
	if o == nil || o.TaskID == nil {
		return nil, false
	}
	return o.TaskID, true
}

// HasTaskID returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasTaskID() bool {
	return o != nil && o.TaskID != nil
}

// SetTaskID gets a reference to the given string and assigns it to the TaskID field.
func (o *TasksSearchResponse) SetTaskID(v string) {
	o.TaskID = &v
}

// GetModelType returns the ModelType field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetModelType() string {
	if o == nil || o.ModelType == nil {
		var ret string
		return ret
	}
	return *o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetModelTypeOk() (*string, bool) {
	if o == nil || o.ModelType == nil {
		return nil, false
	}
	return o.ModelType, true
}

// HasModelType returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasModelType() bool {
	return o != nil && o.ModelType != nil
}

// SetModelType gets a reference to the given string and assigns it to the ModelType field.
func (o *TasksSearchResponse) SetModelType(v string) {
	o.ModelType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TasksSearchResponse) SetName(v string) {
	o.Name = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetIsPrivate() bool {
	if o == nil || o.IsPrivate == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetIsPrivateOk() (*bool, bool) {
	if o == nil || o.IsPrivate == nil {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasIsPrivate() bool {
	return o != nil && o.IsPrivate != nil
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *TasksSearchResponse) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetModelID returns the ModelID field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetModelID() string {
	if o == nil || o.ModelID == nil {
		var ret string
		return ret
	}
	return *o.ModelID
}

// GetModelIDOk returns a tuple with the ModelID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetModelIDOk() (*string, bool) {
	if o == nil || o.ModelID == nil {
		return nil, false
	}
	return o.ModelID, true
}

// HasModelID returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasModelID() bool {
	return o != nil && o.ModelID != nil
}

// SetModelID gets a reference to the given string and assigns it to the ModelID field.
func (o *TasksSearchResponse) SetModelID(v string) {
	o.ModelID = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetOperationType() string {
	if o == nil || o.OperationType == nil {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetOperationTypeOk() (*string, bool) {
	if o == nil || o.OperationType == nil {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasOperationType() bool {
	return o != nil && o.OperationType != nil
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *TasksSearchResponse) SetOperationType(v string) {
	o.OperationType = &v
}

// GetScheduleConfig returns the ScheduleConfig field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetScheduleConfig() ScheduleConfig {
	if o == nil || o.ScheduleConfig == nil {
		var ret ScheduleConfig
		return ret
	}
	return *o.ScheduleConfig
}

// GetScheduleConfigOk returns a tuple with the ScheduleConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetScheduleConfigOk() (*ScheduleConfig, bool) {
	if o == nil || o.ScheduleConfig == nil {
		return nil, false
	}
	return o.ScheduleConfig, true
}

// HasScheduleConfig returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasScheduleConfig() bool {
	return o != nil && o.ScheduleConfig != nil
}

// SetScheduleConfig gets a reference to the given ScheduleConfig and assigns it to the ScheduleConfig field.
func (o *TasksSearchResponse) SetScheduleConfig(v ScheduleConfig) {
	o.ScheduleConfig = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasActive() bool {
	return o != nil && o.Active != nil
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *TasksSearchResponse) SetActive(v bool) {
	o.Active = &v
}

// GetNextPlannedTime returns the NextPlannedTime field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetNextPlannedTime() string {
	if o == nil || o.NextPlannedTime == nil {
		var ret string
		return ret
	}
	return *o.NextPlannedTime
}

// GetNextPlannedTimeOk returns a tuple with the NextPlannedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetNextPlannedTimeOk() (*string, bool) {
	if o == nil || o.NextPlannedTime == nil {
		return nil, false
	}
	return o.NextPlannedTime, true
}

// HasNextPlannedTime returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasNextPlannedTime() bool {
	return o != nil && o.NextPlannedTime != nil
}

// SetNextPlannedTime gets a reference to the given string and assigns it to the NextPlannedTime field.
func (o *TasksSearchResponse) SetNextPlannedTime(v string) {
	o.NextPlannedTime = &v
}

// GetLastExecutionStartTime returns the LastExecutionStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetLastExecutionStartTime() string {
	if o == nil || o.LastExecutionStartTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastExecutionStartTime.Get()
}

// GetLastExecutionStartTimeOk returns a tuple with the LastExecutionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetLastExecutionStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutionStartTime.Get(), o.LastExecutionStartTime.IsSet()
}

// HasLastExecutionStartTime returns a boolean if a LastExecutionStartTime has been set.
func (o *TasksSearchResponse) HasLastExecutionStartTime() bool {
	return o != nil && o.LastExecutionStartTime.IsSet()
}

// SetLastExecutionStartTime gets a reference to the given common.NullableString and assigns it to the LastExecutionStartTime field.
func (o *TasksSearchResponse) SetLastExecutionStartTime(v string) {
	o.LastExecutionStartTime.Set(&v)
}

// SetLastExecutionStartTimeNil sets the value for LastExecutionStartTime to be an explicit nil.
func (o *TasksSearchResponse) SetLastExecutionStartTimeNil() {
	o.LastExecutionStartTime.Set(nil)
}

// UnSetLastExecutionStartTime ensures that no value is present for LastExecutionStartTime, not even an explicit nil.
func (o *TasksSearchResponse) UnsetLastExecutionStartTime() {
	o.LastExecutionStartTime.Unset()
}

// GetLastSuccesufulExecutionTime returns the LastSuccesufulExecutionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetLastSuccesufulExecutionTime() string {
	if o == nil || o.LastSuccesufulExecutionTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastSuccesufulExecutionTime.Get()
}

// GetLastSuccesufulExecutionTimeOk returns a tuple with the LastSuccesufulExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetLastSuccesufulExecutionTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSuccesufulExecutionTime.Get(), o.LastSuccesufulExecutionTime.IsSet()
}

// HasLastSuccesufulExecutionTime returns a boolean if a LastSuccesufulExecutionTime has been set.
func (o *TasksSearchResponse) HasLastSuccesufulExecutionTime() bool {
	return o != nil && o.LastSuccesufulExecutionTime.IsSet()
}

// SetLastSuccesufulExecutionTime gets a reference to the given common.NullableString and assigns it to the LastSuccesufulExecutionTime field.
func (o *TasksSearchResponse) SetLastSuccesufulExecutionTime(v string) {
	o.LastSuccesufulExecutionTime.Set(&v)
}

// SetLastSuccesufulExecutionTimeNil sets the value for LastSuccesufulExecutionTime to be an explicit nil.
func (o *TasksSearchResponse) SetLastSuccesufulExecutionTimeNil() {
	o.LastSuccesufulExecutionTime.Set(nil)
}

// UnSetLastSuccesufulExecutionTime ensures that no value is present for LastSuccesufulExecutionTime, not even an explicit nil.
func (o *TasksSearchResponse) UnsetLastSuccesufulExecutionTime() {
	o.LastSuccesufulExecutionTime.Unset()
}

// GetLastFailedExecutionTime returns the LastFailedExecutionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetLastFailedExecutionTime() string {
	if o == nil || o.LastFailedExecutionTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastFailedExecutionTime.Get()
}

// GetLastFailedExecutionTimeOk returns a tuple with the LastFailedExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetLastFailedExecutionTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFailedExecutionTime.Get(), o.LastFailedExecutionTime.IsSet()
}

// HasLastFailedExecutionTime returns a boolean if a LastFailedExecutionTime has been set.
func (o *TasksSearchResponse) HasLastFailedExecutionTime() bool {
	return o != nil && o.LastFailedExecutionTime.IsSet()
}

// SetLastFailedExecutionTime gets a reference to the given common.NullableString and assigns it to the LastFailedExecutionTime field.
func (o *TasksSearchResponse) SetLastFailedExecutionTime(v string) {
	o.LastFailedExecutionTime.Set(&v)
}

// SetLastFailedExecutionTimeNil sets the value for LastFailedExecutionTime to be an explicit nil.
func (o *TasksSearchResponse) SetLastFailedExecutionTimeNil() {
	o.LastFailedExecutionTime.Set(nil)
}

// UnSetLastFailedExecutionTime ensures that no value is present for LastFailedExecutionTime, not even an explicit nil.
func (o *TasksSearchResponse) UnsetLastFailedExecutionTime() {
	o.LastFailedExecutionTime.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetCreatedDate() string {
	if o == nil || o.CreatedDate == nil {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetCreatedDateOk() (*string, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasCreatedDate() bool {
	return o != nil && o.CreatedDate != nil
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *TasksSearchResponse) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetJwtLifeTimeSeconds returns the JwtLifeTimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetJwtLifeTimeSeconds() string {
	if o == nil || o.JwtLifeTimeSeconds.Get() == nil {
		var ret string
		return ret
	}
	return *o.JwtLifeTimeSeconds.Get()
}

// GetJwtLifeTimeSecondsOk returns a tuple with the JwtLifeTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetJwtLifeTimeSecondsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JwtLifeTimeSeconds.Get(), o.JwtLifeTimeSeconds.IsSet()
}

// HasJwtLifeTimeSeconds returns a boolean if a JwtLifeTimeSeconds has been set.
func (o *TasksSearchResponse) HasJwtLifeTimeSeconds() bool {
	return o != nil && o.JwtLifeTimeSeconds.IsSet()
}

// SetJwtLifeTimeSeconds gets a reference to the given common.NullableString and assigns it to the JwtLifeTimeSeconds field.
func (o *TasksSearchResponse) SetJwtLifeTimeSeconds(v string) {
	o.JwtLifeTimeSeconds.Set(&v)
}

// SetJwtLifeTimeSecondsNil sets the value for JwtLifeTimeSeconds to be an explicit nil.
func (o *TasksSearchResponse) SetJwtLifeTimeSecondsNil() {
	o.JwtLifeTimeSeconds.Set(nil)
}

// UnSetJwtLifeTimeSeconds ensures that no value is present for JwtLifeTimeSeconds, not even an explicit nil.
func (o *TasksSearchResponse) UnsetJwtLifeTimeSeconds() {
	o.JwtLifeTimeSeconds.Unset()
}

// GetModuleFilter returns the ModuleFilter field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetModuleFilter() string {
	if o == nil || o.ModuleFilter == nil {
		var ret string
		return ret
	}
	return *o.ModuleFilter
}

// GetModuleFilterOk returns a tuple with the ModuleFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetModuleFilterOk() (*string, bool) {
	if o == nil || o.ModuleFilter == nil {
		return nil, false
	}
	return o.ModuleFilter, true
}

// HasModuleFilter returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasModuleFilter() bool {
	return o != nil && o.ModuleFilter != nil
}

// SetModuleFilter gets a reference to the given string and assigns it to the ModuleFilter field.
func (o *TasksSearchResponse) SetModuleFilter(v string) {
	o.ModuleFilter = &v
}

// GetRemoteInterfaceName returns the RemoteInterfaceName field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetRemoteInterfaceName() string {
	if o == nil || o.RemoteInterfaceName == nil {
		var ret string
		return ret
	}
	return *o.RemoteInterfaceName
}

// GetRemoteInterfaceNameOk returns a tuple with the RemoteInterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetRemoteInterfaceNameOk() (*string, bool) {
	if o == nil || o.RemoteInterfaceName == nil {
		return nil, false
	}
	return o.RemoteInterfaceName, true
}

// HasRemoteInterfaceName returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasRemoteInterfaceName() bool {
	return o != nil && o.RemoteInterfaceName != nil
}

// SetRemoteInterfaceName gets a reference to the given string and assigns it to the RemoteInterfaceName field.
func (o *TasksSearchResponse) SetRemoteInterfaceName(v string) {
	o.RemoteInterfaceName = &v
}

// GetRemoteMethodName returns the RemoteMethodName field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetRemoteMethodName() string {
	if o == nil || o.RemoteMethodName == nil {
		var ret string
		return ret
	}
	return *o.RemoteMethodName
}

// GetRemoteMethodNameOk returns a tuple with the RemoteMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetRemoteMethodNameOk() (*string, bool) {
	if o == nil || o.RemoteMethodName == nil {
		return nil, false
	}
	return o.RemoteMethodName, true
}

// HasRemoteMethodName returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasRemoteMethodName() bool {
	return o != nil && o.RemoteMethodName != nil
}

// SetRemoteMethodName gets a reference to the given string and assigns it to the RemoteMethodName field.
func (o *TasksSearchResponse) SetRemoteMethodName(v string) {
	o.RemoteMethodName = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a Data has been set.
func (o *TasksSearchResponse) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given common.NullableString and assigns it to the Data field.
func (o *TasksSearchResponse) SetData(v string) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *TasksSearchResponse) SetDataNil() {
	o.Data.Set(nil)
}

// UnSetData ensures that no value is present for Data, not even an explicit nil.
func (o *TasksSearchResponse) UnsetData() {
	o.Data.Unset()
}

// GetMicroserviceAddress returns the MicroserviceAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetMicroserviceAddress() string {
	if o == nil || o.MicroserviceAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.MicroserviceAddress.Get()
}

// GetMicroserviceAddressOk returns a tuple with the MicroserviceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetMicroserviceAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MicroserviceAddress.Get(), o.MicroserviceAddress.IsSet()
}

// HasMicroserviceAddress returns a boolean if a MicroserviceAddress has been set.
func (o *TasksSearchResponse) HasMicroserviceAddress() bool {
	return o != nil && o.MicroserviceAddress.IsSet()
}

// SetMicroserviceAddress gets a reference to the given common.NullableString and assigns it to the MicroserviceAddress field.
func (o *TasksSearchResponse) SetMicroserviceAddress(v string) {
	o.MicroserviceAddress.Set(&v)
}

// SetMicroserviceAddressNil sets the value for MicroserviceAddress to be an explicit nil.
func (o *TasksSearchResponse) SetMicroserviceAddressNil() {
	o.MicroserviceAddress.Set(nil)
}

// UnSetMicroserviceAddress ensures that no value is present for MicroserviceAddress, not even an explicit nil.
func (o *TasksSearchResponse) UnsetMicroserviceAddress() {
	o.MicroserviceAddress.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetResult() string {
	if o == nil || o.Result.Get() == nil {
		var ret string
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a Result has been set.
func (o *TasksSearchResponse) HasResult() bool {
	return o != nil && o.Result.IsSet()
}

// SetResult gets a reference to the given common.NullableString and assigns it to the Result field.
func (o *TasksSearchResponse) SetResult(v string) {
	o.Result.Set(&v)
}

// SetResultNil sets the value for Result to be an explicit nil.
func (o *TasksSearchResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnSetResult ensures that no value is present for Result, not even an explicit nil.
func (o *TasksSearchResponse) UnsetResult() {
	o.Result.Unset()
}

// GetDisableRetryIfFail returns the DisableRetryIfFail field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetDisableRetryIfFail() bool {
	if o == nil || o.DisableRetryIfFail == nil {
		var ret bool
		return ret
	}
	return *o.DisableRetryIfFail
}

// GetDisableRetryIfFailOk returns a tuple with the DisableRetryIfFail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetDisableRetryIfFailOk() (*bool, bool) {
	if o == nil || o.DisableRetryIfFail == nil {
		return nil, false
	}
	return o.DisableRetryIfFail, true
}

// HasDisableRetryIfFail returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasDisableRetryIfFail() bool {
	return o != nil && o.DisableRetryIfFail != nil
}

// SetDisableRetryIfFail gets a reference to the given bool and assigns it to the DisableRetryIfFail field.
func (o *TasksSearchResponse) SetDisableRetryIfFail(v bool) {
	o.DisableRetryIfFail = &v
}

// GetExtraObject returns the ExtraObject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetExtraObject() string {
	if o == nil || o.ExtraObject.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtraObject.Get()
}

// GetExtraObjectOk returns a tuple with the ExtraObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetExtraObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraObject.Get(), o.ExtraObject.IsSet()
}

// HasExtraObject returns a boolean if a ExtraObject has been set.
func (o *TasksSearchResponse) HasExtraObject() bool {
	return o != nil && o.ExtraObject.IsSet()
}

// SetExtraObject gets a reference to the given common.NullableString and assigns it to the ExtraObject field.
func (o *TasksSearchResponse) SetExtraObject(v string) {
	o.ExtraObject.Set(&v)
}

// SetExtraObjectNil sets the value for ExtraObject to be an explicit nil.
func (o *TasksSearchResponse) SetExtraObjectNil() {
	o.ExtraObject.Set(nil)
}

// UnSetExtraObject ensures that no value is present for ExtraObject, not even an explicit nil.
func (o *TasksSearchResponse) UnsetExtraObject() {
	o.ExtraObject.Unset()
}

// GetExecutorName returns the ExecutorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetExecutorName() string {
	if o == nil || o.ExecutorName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExecutorName.Get()
}

// GetExecutorNameOk returns a tuple with the ExecutorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetExecutorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutorName.Get(), o.ExecutorName.IsSet()
}

// HasExecutorName returns a boolean if a ExecutorName has been set.
func (o *TasksSearchResponse) HasExecutorName() bool {
	return o != nil && o.ExecutorName.IsSet()
}

// SetExecutorName gets a reference to the given common.NullableString and assigns it to the ExecutorName field.
func (o *TasksSearchResponse) SetExecutorName(v string) {
	o.ExecutorName.Set(&v)
}

// SetExecutorNameNil sets the value for ExecutorName to be an explicit nil.
func (o *TasksSearchResponse) SetExecutorNameNil() {
	o.ExecutorName.Set(nil)
}

// UnSetExecutorName ensures that no value is present for ExecutorName, not even an explicit nil.
func (o *TasksSearchResponse) UnsetExecutorName() {
	o.ExecutorName.Unset()
}

// GetExecutorType returns the ExecutorType field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetExecutorType() string {
	if o == nil || o.ExecutorType == nil {
		var ret string
		return ret
	}
	return *o.ExecutorType
}

// GetExecutorTypeOk returns a tuple with the ExecutorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetExecutorTypeOk() (*string, bool) {
	if o == nil || o.ExecutorType == nil {
		return nil, false
	}
	return o.ExecutorType, true
}

// HasExecutorType returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasExecutorType() bool {
	return o != nil && o.ExecutorType != nil
}

// SetExecutorType gets a reference to the given string and assigns it to the ExecutorType field.
func (o *TasksSearchResponse) SetExecutorType(v string) {
	o.ExecutorType = &v
}

// GetTenantID returns the TenantID field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetTenantID() string {
	if o == nil || o.TenantID == nil {
		var ret string
		return ret
	}
	return *o.TenantID
}

// GetTenantIDOk returns a tuple with the TenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetTenantIDOk() (*string, bool) {
	if o == nil || o.TenantID == nil {
		return nil, false
	}
	return o.TenantID, true
}

// HasTenantID returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasTenantID() bool {
	return o != nil && o.TenantID != nil
}

// SetTenantID gets a reference to the given string and assigns it to the TenantID field.
func (o *TasksSearchResponse) SetTenantID(v string) {
	o.TenantID = &v
}

// GetModuleGUID returns the ModuleGUID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksSearchResponse) GetModuleGUID() string {
	if o == nil || o.ModuleGUID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleGUID.Get()
}

// GetModuleGUIDOk returns a tuple with the ModuleGUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TasksSearchResponse) GetModuleGUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleGUID.Get(), o.ModuleGUID.IsSet()
}

// HasModuleGUID returns a boolean if a ModuleGUID has been set.
func (o *TasksSearchResponse) HasModuleGUID() bool {
	return o != nil && o.ModuleGUID.IsSet()
}

// SetModuleGUID gets a reference to the given common.NullableString and assigns it to the ModuleGUID field.
func (o *TasksSearchResponse) SetModuleGUID(v string) {
	o.ModuleGUID.Set(&v)
}

// SetModuleGUIDNil sets the value for ModuleGUID to be an explicit nil.
func (o *TasksSearchResponse) SetModuleGUIDNil() {
	o.ModuleGUID.Set(nil)
}

// UnSetModuleGUID ensures that no value is present for ModuleGUID, not even an explicit nil.
func (o *TasksSearchResponse) UnsetModuleGUID() {
	o.ModuleGUID.Unset()
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetOwnerType() string {
	if o == nil || o.OwnerType == nil {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetOwnerTypeOk() (*string, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasOwnerType() bool {
	return o != nil && o.OwnerType != nil
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *TasksSearchResponse) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetOwnerName() string {
	if o == nil || o.OwnerName == nil {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetOwnerNameOk() (*string, bool) {
	if o == nil || o.OwnerName == nil {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasOwnerName() bool {
	return o != nil && o.OwnerName != nil
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *TasksSearchResponse) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasApplicationName() bool {
	return o != nil && o.ApplicationName != nil
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *TasksSearchResponse) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetRunForAllTenants returns the RunForAllTenants field value if set, zero value otherwise.
func (o *TasksSearchResponse) GetRunForAllTenants() bool {
	if o == nil || o.RunForAllTenants == nil {
		var ret bool
		return ret
	}
	return *o.RunForAllTenants
}

// GetRunForAllTenantsOk returns a tuple with the RunForAllTenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksSearchResponse) GetRunForAllTenantsOk() (*bool, bool) {
	if o == nil || o.RunForAllTenants == nil {
		return nil, false
	}
	return o.RunForAllTenants, true
}

// HasRunForAllTenants returns a boolean if a field has been set.
func (o *TasksSearchResponse) HasRunForAllTenants() bool {
	return o != nil && o.RunForAllTenants != nil
}

// SetRunForAllTenants gets a reference to the given bool and assigns it to the RunForAllTenants field.
func (o *TasksSearchResponse) SetRunForAllTenants(v bool) {
	o.RunForAllTenants = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TasksSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.TaskID != nil {
		toSerialize["Items"] = o.TaskID
	}
	if o.ModelType != nil {
		toSerialize["ModelType"] = o.ModelType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.IsPrivate != nil {
		toSerialize["IsPrivate"] = o.IsPrivate
	}
	if o.ModelID != nil {
		toSerialize["ModelID"] = o.ModelID
	}
	if o.OperationType != nil {
		toSerialize["OperationType"] = o.OperationType
	}
	if o.TaskRunID.IsSet() {
		toSerialize["TaskRunID"] = o.TaskRunID.Get()
	}
	if o.ScheduleConfig != nil {
		toSerialize["ScheduleConfig"] = o.ScheduleConfig
	}
	if o.Active != nil {
		toSerialize["Active"] = o.Active
	}
	if o.NextPlannedTime != nil {
		toSerialize["NextPlannedTime"] = o.NextPlannedTime
	}
	if o.LastExecutionStartTime.IsSet() {
		toSerialize["LastExecutionStartTime"] = o.LastExecutionStartTime.Get()
	}
	if o.LastSuccesufulExecutionTime.IsSet() {
		toSerialize["LastSuccesufulExecutionTime"] = o.LastSuccesufulExecutionTime.Get()
	}
	if o.LastFailedExecutionTime.IsSet() {
		toSerialize["LastFailedExecutionTime"] = o.LastFailedExecutionTime.Get()
	}
	if o.CreatedDate != nil {
		toSerialize["CreatedDate"] = o.CreatedDate
	}
	if o.JwtLifeTimeSeconds.IsSet() {
		toSerialize["JwtLifeTimeSeconds"] = o.JwtLifeTimeSeconds
	}
	if o.ModuleFilter != nil {
		toSerialize["ModuleFilter"] = o.ModuleFilter
	}
	if o.RemoteInterfaceName != nil {
		toSerialize["RemoteInterfaceName"] = o.RemoteInterfaceName
	}
	if o.RemoteMethodName != nil {
		toSerialize["RemoteMethodName"] = o.RemoteMethodName
	}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data.Get()
	}
	if o.MicroserviceAddress.IsSet() {
		toSerialize["MicroserviceAddress"] = o.MicroserviceAddress.Get()
	}
	if o.Result.IsSet() {
		toSerialize["Result"] = o.Result.Get()
	}
	if o.DisableRetryIfFail != nil {
		toSerialize["DisableRetryIfFail"] = o.DisableRetryIfFail
	}
	if o.ExtraObject.IsSet() {
		toSerialize["NaExtraObjectme"] = o.ExtraObject.Get()
	}
	if o.ExecutorName.IsSet() {
		toSerialize["ExecutorName"] = o.ExecutorName.Get()
	}
	if o.ExecutorType != nil {
		toSerialize["ExecutorType"] = o.ExecutorType
	}
	if o.TenantID != nil {
		toSerialize["TenantID"] = o.TenantID
	}
	if o.ModuleGUID.IsSet() {
		toSerialize["ModuleGuid"] = o.ModuleGUID.Get()
	}
	if o.OwnerType != nil {
		toSerialize["OwnerType"] = o.OwnerType
	}
	if o.OwnerName != nil {
		toSerialize["OwnerName"] = o.OwnerName
	}
	if o.ApplicationName != nil {
		toSerialize["ApplicationName"] = o.ApplicationName
	}
	if o.RunForAllTenants != nil {
		toSerialize["RunForAllTenants"] = o.RunForAllTenants
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *TasksSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskID                      *string               `json:"TaskID,omitempty"`
		ModelType                   *string               `json:"ModelType,omitempty"`
		Name                        *string               `json:"Name,omitempty"`
		IsPrivate                   *bool                 `json:"IsPrivate,omitempty"`
		ModelID                     *string               `json:"ModelID,omitempty"`
		OperationType               *string               `json:"OperationType,omitempty"`
		TaskRunID                   common.NullableString `json:"TaskRunID,omitempty"`
		ScheduleConfig              *ScheduleConfig       `json:"ScheduleConfig,omitempty"`
		Active                      *bool                 `json:"Active,omitempty"`
		NextPlannedTime             *string               `json:"NextPlannedTime,omitempty"`
		LastExecutionStartTime      common.NullableString `json:"LastExecutionStartTime,omitempty"`
		LastSuccesufulExecutionTime common.NullableString `json:"LastSuccesufulExecutionTime,omitempty"`
		LastFailedExecutionTime     common.NullableString `json:"LastFailedExecutionTime,omitempty"`
		CreatedDate                 *string               `json:"CreatedDate,omitempty"`
		JwtLifeTimeSeconds          common.NullableString `json:"JwtLifeTimeSeconds,omitempty"`
		ModuleFilter                *string               `json:"ModuleFilter,omitempty"`
		RemoteInterfaceName         *string               `json:"RemoteInterfaceName,omitempty"`
		RemoteMethodName            *string               `json:"RemoteMethodName,omitempty"`
		Data                        common.NullableString `json:"Data,omitempty"`
		MicroserviceAddress         common.NullableString `json:"MicroserviceAddress,omitempty"`
		Result                      common.NullableString `json:"Result,omitempty"`
		DisableRetryIfFail          *bool                 `json:"DisableRetryIfFail,omitempty"`
		ExtraObject                 common.NullableString `json:"ExtraObject,omitempty"`
		ExecutorName                common.NullableString `json:"ExecutorName,omitempty"`
		ExecutorType                *string               `json:"ExecutorType,omitempty"`
		TenantID                    *string               `json:"TenantID,omitempty"`
		ModuleGUID                  common.NullableString `json:"ModuleGuid,omitempty"`
		OwnerType                   *string               `json:"OwnerType,omitempty"`
		OwnerName                   *string               `json:"OwnerName,omitempty"`
		ApplicationName             *string               `json:"ApplicationName,omitempty"`
		RunForAllTenants            *bool                 `json:"RunForAllTenants,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"TaskID", "ModelType", "Name", "IsPrivate", "ModelID", "OperationType", "OperationType", "ScheduleConfig", "Active", "NextPlannedTime", "LastExecutionStartTime", "LastSuccesufulExecutionTime", "LastFailedExecutionTime", "CreatedDate", "JwtLifeTimeSeconds", "ModuleFilter", "RemoteInterfaceName", "RemoteMethodName", "Data", "MicroserviceAddress", "Result", "DisableRetryIfFail", "ExtraObject", "ExecutorName", "ExecutorType", "TenantID", "ModuleGuid", "OwnerType", "OwnerName", "ApplicationName", "RunForAllTenants"})
	} else {
		return err
	}

	o.TaskID = all.TaskID
	o.ModelType = all.ModelType
	o.Name = all.Name
	o.IsPrivate = all.IsPrivate
	o.ModelID = all.ModelID
	o.OperationType = all.OperationType
	o.TaskRunID = all.TaskRunID
	hasInvalidField := false
	if all.ScheduleConfig != nil && all.ScheduleConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScheduleConfig = all.ScheduleConfig
	o.Active = all.Active
	o.NextPlannedTime = all.NextPlannedTime
	o.LastExecutionStartTime = all.LastExecutionStartTime
	o.LastSuccesufulExecutionTime = all.LastSuccesufulExecutionTime
	o.LastFailedExecutionTime = all.LastFailedExecutionTime
	o.CreatedDate = all.CreatedDate
	o.JwtLifeTimeSeconds = all.JwtLifeTimeSeconds
	o.ModuleFilter = all.ModuleFilter
	o.RemoteInterfaceName = all.RemoteInterfaceName
	o.RemoteMethodName = all.RemoteMethodName
	o.Data = all.Data
	o.MicroserviceAddress = all.MicroserviceAddress
	o.Result = all.Result
	o.DisableRetryIfFail = all.DisableRetryIfFail
	o.ExtraObject = all.ExtraObject
	o.ExecutorName = all.ExecutorName
	o.ExecutorType = all.ExecutorType
	o.TenantID = all.TenantID
	o.ModuleGUID = all.ModuleGUID
	o.OwnerType = all.OwnerType
	o.OwnerName = all.OwnerName
	o.ApplicationName = all.ApplicationName
	o.RunForAllTenants = all.RunForAllTenants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
