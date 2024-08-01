package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SearchArchivesStatus defines the status structure for search archives.
type SearchArchivesStatus struct {
	// SearchArchiveID represents the identifier of the search archive.
	SearchArchiveID *string `json:"SearchArchiveId,omitempty"`
	// Message contains an optional message related to the search archive status.
	Message common.NullableString `json:"Message,omitempty"`
	// ExecutionStartTimeUtc indicates the UTC start time of the search execution.
	ExecutionStartTimeUtc *string `json:"ExecutionStartTimeUtc,omitempty"`
	// ExecutionDurationMs specifies the duration of the search execution in milliseconds.
	ExecutionDurationMs *int64 `json:"ExecutionDurationMs,omitempty"`
	// Status indicates the current status of the search archive.
	Status *string `json:"Status,omitempty"`
	// PauseRequest specifies if there's a request to pause the search.
	PauseRequest *bool `json:"PauseRequest,omitempty"`
	// FoundLines represents the number of lines found during the search.
	FoundLines *int64 `json:"FoundLines,omitempty"`
	// ScannedLines indicates the number of lines scanned during the search.
	ScannedLines *int64 `json:"ScannedLines,omitempty"`
	// ProcessedFileCount specifies the number of files processed during the search.
	ProcessedFileCount *int64 `json:"ProcessedFileCount,omitempty"`
	// ProcessedFileDate indicates the date of the processed files.
	ProcessedFileDate *string `json:"ProcessedFileDate,omitempty"`
	// ActualStartDate indicates the actual start date of the search execution.
	ActualStartDate *string `json:"ActualStartDate,omitempty"`
	// ActualEndDate specifies the actual end date of the search execution.
	ActualEndDate *string `json:"ActualEndDate,omitempty"`
	// LimitReached specifies if any search limit was reached.
	LimitReached *bool `json:"LimitReached,omitempty"`
	// RunningMachineIP represents the IP address of the machine where the search is running.
	RunningMachineIP *string `json:"RunningMachineIp,omitempty"`
	// TenantID represents the identifier of the tenant associated with the search.
	TenantID *string `json:"TenantID,omitempty"`
	// ID represents the unique identifier of the search status.
	ID *string `json:"_id,omitempty"`
	// Timestamp indicates the timestamp associated with the search status.
	Timestamp *string `json:"Timestamp,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSearchArchivesStatus creates a new SearchArchivesStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchArchivesStatus() *SearchArchivesStatus {
	this := SearchArchivesStatus{}
	return &this
}

// NewSearchArchivesStatusWithDefaults creates a new SearchArchivesStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSearchArchivesStatusWithDefaults() *SearchArchivesStatus {
	this := SearchArchivesStatus{}
	return &this
}

// GetSearchArchiveID returns the SearchArchiveID field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetSearchArchiveI() string {
	if o == nil || o.SearchArchiveID == nil {
		var ret string
		return ret
	}
	return *o.SearchArchiveID
}

// GetSearchArchiveIDOk returns a tuple with the SearchArchiveID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetSearchArchiveIDOk() (*string, bool) {
	if o == nil || o.SearchArchiveID == nil {
		return nil, false
	}
	return o.SearchArchiveID, true
}

// HasSearchArchiveID returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasSearchArchiveID() bool {
	return o != nil && o.SearchArchiveID != nil
}

// SetSearchArchiveID gets a reference to the given string and assigns it to the SearchArchiveI field.
func (o *SearchArchivesStatus) SetSearchArchiveID(v string) {
	o.SearchArchiveID = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchArchivesStatus) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchArchivesStatus) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a Message has been set.
func (o *SearchArchivesStatus) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given common.NullableString and assigns it to the Message field.
func (o *SearchArchivesStatus) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *SearchArchivesStatus) SetMessageNil() {
	o.Message.Set(nil)
}

// UnSetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *SearchArchivesStatus) UnSetMessage() {
	o.Message.UnSet()
}

// GetExecutionStartTimeUtc returns the ExecutionStartTimeUtc field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetExecutionStartTimeUtc() string {
	if o == nil || o.ExecutionStartTimeUtc == nil {
		var ret string
		return ret
	}
	return *o.ExecutionStartTimeUtc
}

// GetExecutionStartTimeUtcOk returns a tuple with the ExecutionStartTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetExecutionStartTimeUtcOk() (*string, bool) {
	if o == nil || o.ExecutionStartTimeUtc == nil {
		return nil, false
	}
	return o.ExecutionStartTimeUtc, true
}

// HasExecutionStartTimeUtc returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasExecutionStartTimeUtc() bool {
	return o != nil && o.ExecutionStartTimeUtc != nil
}

// SetExecutionStartTimeUtc gets a reference to the given string and assigns it to the ExecutionStartTimeUtc field.
func (o *SearchArchivesStatus) SetExecutionStartTimeUtc(v string) {
	o.ExecutionStartTimeUtc = &v
}

// GetExecutionDurationMs returns the ExecutionDurationMs field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetExecutionDurationMs() int64 {
	if o == nil || o.ExecutionDurationMs == nil {
		var ret int64
		return ret
	}
	return *o.ExecutionDurationMs
}

// GetExecutionDurationMsOk returns a tuple with the ExecutionDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetExecutionDurationMsOk() (*int64, bool) {
	if o == nil || o.ExecutionDurationMs == nil {
		return nil, false
	}
	return o.ExecutionDurationMs, true
}

// HasExecutionDurationMs returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasExecutionDurationMs() bool {
	return o != nil && o.ExecutionDurationMs != nil
}

// SetExecutionDurationMs gets a reference to the given int64 and assigns it to the ExecutionDurationMs field.
func (o *SearchArchivesStatus) SetExecutionDurationMs(v int64) {
	o.ExecutionDurationMs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SearchArchivesStatus) SetStatus(v string) {
	o.Status = &v
}

// GetPauseRequest returns the PauseRequest field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetPauseRequest() bool {
	if o == nil || o.PauseRequest == nil {
		var ret bool
		return ret
	}
	return *o.PauseRequest
}

// GetPauseRequestOk returns a tuple with the PauseRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetPauseRequestOk() (*bool, bool) {
	if o == nil || o.PauseRequest == nil {
		return nil, false
	}
	return o.PauseRequest, true
}

// HasPauseRequest returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasPauseRequest() bool {
	return o != nil && o.PauseRequest != nil
}

// SetPauseRequest gets a reference to the given bool and assigns it to the PauseRequest field.
func (o *SearchArchivesStatus) SetPauseRequest(v bool) {
	o.PauseRequest = &v
}

// GetFoundLines returns the FoundLines field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetFoundLines() int64 {
	if o == nil || o.FoundLines == nil {
		var ret int64
		return ret
	}
	return *o.FoundLines
}

// GetFoundLinesOk returns a tuple with the FoundLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetFoundLinesOk() (*int64, bool) {
	if o == nil || o.FoundLines == nil {
		return nil, false
	}
	return o.FoundLines, true
}

// HasFoundLines returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasFoundLines() bool {
	return o != nil && o.FoundLines != nil
}

// SetFoundLines gets a reference to the given int64 and assigns it to the FoundLines field.
func (o *SearchArchivesStatus) SetFoundLines(v int64) {
	o.FoundLines = &v
}

// GetProcessedFileCount returns the ProcessedFileCount field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetProcessedFileCount() int64 {
	if o == nil || o.ProcessedFileCount == nil {
		var ret int64
		return ret
	}
	return *o.ProcessedFileCount
}

// GetProcessedFileCountOk returns a tuple with the ProcessedFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetProcessedFileCountOk() (*int64, bool) {
	if o == nil || o.ProcessedFileCount == nil {
		return nil, false
	}
	return o.ProcessedFileCount, true
}

// HasProcessedFileCount returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasProcessedFileCount() bool {
	return o != nil && o.ProcessedFileCount != nil
}

// SetProcessedFileCount gets a reference to the given int64 and assigns it to the ProcessedFileCount field.
func (o *SearchArchivesStatus) SetProcessedFileCount(v int64) {
	o.ProcessedFileCount = &v
}

// GetProcessedFileDate returns the ProcessedFileDate field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetProcessedFileDate() string {
	if o == nil || o.ProcessedFileDate == nil {
		var ret string
		return ret
	}
	return *o.ProcessedFileDate
}

// GetProcessedFileDateOk returns a tuple with the ProcessedFileDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetProcessedFileDateOk() (*string, bool) {
	if o == nil || o.ProcessedFileDate == nil {
		return nil, false
	}
	return o.ProcessedFileDate, true
}

// HasProcessedFileDate returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasProcessedFileDate() bool {
	return o != nil && o.ProcessedFileDate != nil
}

// SetProcessedFileDate gets a reference to the given string and assigns it to the ProcessedFileDate field.
func (o *SearchArchivesStatus) SetProcessedFileDate(v string) {
	o.ProcessedFileDate = &v
}

// GetActualStartDate returns the ActualStartDate field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetActualStartDate() string {
	if o == nil || o.ActualStartDate == nil {
		var ret string
		return ret
	}
	return *o.ActualStartDate
}

// GetActualStartDateOk returns a tuple with the ActualStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetActualStartDateOk() (*string, bool) {
	if o == nil || o.ActualStartDate == nil {
		return nil, false
	}
	return o.ActualStartDate, true
}

// HasActualStartDate returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasActualStartDate() bool {
	return o != nil && o.ActualStartDate != nil
}

// SetActualStartDate gets a reference to the given string and assigns it to the ActualStartDate field.
func (o *SearchArchivesStatus) SetActualStartDate(v string) {
	o.ActualStartDate = &v
}

// GetActualEndDate returns the ActualEndDate field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetActualEndDate() string {
	if o == nil || o.ActualEndDate == nil {
		var ret string
		return ret
	}
	return *o.ActualEndDate
}

// GetActualEndDateOk returns a tuple with the ActualEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetActualEndDateOk() (*string, bool) {
	if o == nil || o.ActualEndDate == nil {
		return nil, false
	}
	return o.ActualEndDate, true
}

// HasActualEndDate returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasActualEndDate() bool {
	return o != nil && o.ActualEndDate != nil
}

// SetActualEndDate gets a reference to the given string and assigns it to the ActualEndDate field.
func (o *SearchArchivesStatus) SetActualEndDate(v string) {
	o.ActualEndDate = &v
}

// GetLimitReached returns the LimitReached field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetLimitReached() bool {
	if o == nil || o.LimitReached == nil {
		var ret bool
		return ret
	}
	return *o.LimitReached
}

// GetLimitReachedOk returns a tuple with the LimitReached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetLimitReachedOk() (*bool, bool) {
	if o == nil || o.LimitReached == nil {
		return nil, false
	}
	return o.LimitReached, true
}

// HasLimitReached returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasLimitReached() bool {
	return o != nil && o.LimitReached != nil
}

// SetLimitReached gets a reference to the given bool and assigns it to the LimitReached field.
func (o *SearchArchivesStatus) SetLimitReached(v bool) {
	o.LimitReached = &v
}

// GetRunningMachineIP returns the RunningMachineIP field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetRunningMachineIP() string {
	if o == nil || o.RunningMachineIP == nil {
		var ret string
		return ret
	}
	return *o.RunningMachineIP
}

// GetRunningMachineIPOk returns a tuple with the RunningMachineIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetRunningMachineIPOk() (*string, bool) {
	if o == nil || o.RunningMachineIP == nil {
		return nil, false
	}
	return o.RunningMachineIP, true
}

// HasRunningMachineIP returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasRunningMachineIP() bool {
	return o != nil && o.RunningMachineIP != nil
}

// SetRunningMachineIP gets a reference to the given string and assigns it to the RunningMachineIP field.
func (o *SearchArchivesStatus) SetRunningMachineIP(v string) {
	o.RunningMachineIP = &v
}

// GetTenantID returns the TenantID field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetTenantID() string {
	if o == nil || o.TenantID == nil {
		var ret string
		return ret
	}
	return *o.TenantID
}

// GetTenantIDOk returns a tuple with the TenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetTenantIDOk() (*string, bool) {
	if o == nil || o.TenantID == nil {
		return nil, false
	}
	return o.TenantID, true
}

// HasTenantID returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasTenantID() bool {
	return o != nil && o.TenantID != nil
}

// SetTenantID gets a reference to the given string and assigns it to the TenantID field.
func (o *SearchArchivesStatus) SetTenantID(v string) {
	o.TenantID = &v
}

// GetId returns the ID field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *SearchArchivesStatus) SetID(v string) {
	o.ID = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SearchArchivesStatus) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchArchivesStatus) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SearchArchivesStatus) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *SearchArchivesStatus) SetTimestamp(v string) {
	o.Timestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SearchArchivesStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.SearchArchiveID != nil {
		toSerialize["SearchArchiveId"] = o.SearchArchiveID
	}
	if o.Message.IsSet() {
		toSerialize["Message"] = o.Message.Get()
	}
	if o.ExecutionStartTimeUtc != nil {
		toSerialize["ExecutionStartTimeUtc"] = o.ExecutionStartTimeUtc
	}
	if o.ExecutionDurationMs != nil {
		toSerialize["ExecutionDurationMs"] = o.ExecutionDurationMs
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.PauseRequest != nil {
		toSerialize["PauseRequest"] = o.PauseRequest
	}
	if o.FoundLines != nil {
		toSerialize["FoundLines"] = o.FoundLines
	}
	if o.ScannedLines != nil {
		toSerialize["ScannedLines"] = o.ScannedLines
	}
	if o.ProcessedFileCount != nil {
		toSerialize["ProcessedFileCount"] = o.ProcessedFileCount
	}
	if o.ProcessedFileDate != nil {
		toSerialize["ProcessedFileDate"] = o.ProcessedFileDate
	}
	if o.ActualStartDate != nil {
		toSerialize["ActualStartDate"] = o.ActualStartDate
	}
	if o.ActualEndDate != nil {
		toSerialize["ActualEndDate"] = o.ActualEndDate
	}
	if o.LimitReached != nil {
		toSerialize["LimitReached"] = o.LimitReached
	}
	if o.RunningMachineIP != nil {
		toSerialize["RunningMachineIp"] = o.RunningMachineIP
	}
	if o.TenantID != nil {
		toSerialize["TenantID"] = o.TenantID
	}
	if o.ID != nil {
		toSerialize["Id"] = o.ID
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SearchArchivesStatus) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		SearchArchiveID       *string               `json:"SearchArchiveId,omitempty"`
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
		RunningMachineIP      *string               `json:"RunningMachineIp,omitempty"`
		TenantID              *string               `json:"TenantID,omitempty"`
		ID                    *string               `json:"_id,omitempty"`
		Timestamp             *string               `json:"Timestamp,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"SearchArchiveId", "Message", "ExecutionStartTimeUtc", "ExecutionDurationMs", "Status", "PauseRequest", "FoundLines", "ScannedLines", "ProcessedFileCount", "ProcessedFileDate", "ActualStartDate", "ActualEndDate", "LimitReached", "RunningMachineIp", "TenantID", "_id", "Timestamp"})
	} else {
		return err
	}

	o.SearchArchiveID = all.SearchArchiveID
	o.Message = all.Message
	o.ExecutionStartTimeUtc = all.ExecutionStartTimeUtc
	o.ExecutionDurationMs = all.ExecutionDurationMs
	o.Status = all.Status
	o.PauseRequest = all.PauseRequest
	o.FoundLines = all.FoundLines
	o.ScannedLines = all.ScannedLines
	o.ProcessedFileCount = all.ProcessedFileCount
	o.ProcessedFileDate = all.ProcessedFileDate
	o.LimitReached = all.LimitReached
	o.RunningMachineIP = all.RunningMachineIP
	o.TenantID = all.TenantID
	o.ID = all.ID
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
