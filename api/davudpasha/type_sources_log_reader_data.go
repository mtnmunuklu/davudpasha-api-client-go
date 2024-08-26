package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SourcesLogReaderData struct {
	FileInfo                         *SourcesFileInfo      `json:"FileInfo,omitempty"`
	MessageFieldType                 common.NullableString `json:"MessageFieldType,omitempty"`
	TransformFrequencyInSeconds      *int64                `json:"TransformFrequencyInSeconds,omitempty"`
	DeleteArchivedLogFileAfterNHours *int64                `json:"DeleteArchivedLogFileAfterNHours,omitempty"`
	DeleteOutputFileAfterNHours      *int64                `json:"DeleteOutputFileAfterNHours,omitempty"`
	MaxParallelProcessing            *int64                `json:"MaxParallelProcessing,omitempty"`
	EventLogNameOrFileName           *int64                `json:"EventLogNameOrFileName,omitempty"`
	Filter                           common.NullableString `json:"Filter,omitempty"`
	DataFormat                       *string               `json:"DataFormat,omitempty"`
	Credential                       *string               `json:"Credential,omitempty"`
	IdColumnType                     *string               `json:"IdColumnType,omitempty"`
	ParseJson                        *bool                 `json:"ParseJson,omitempty"`
	IndexName                        *string               `json:"IndexName,omitempty"`
	IdColumnName                     *string               `json:"IdColumnName,omitempty"`
	IdColumnFormat                   *string               `json:"IdColumnFormat,omitempty"`
	Source                           *string               `json:"Source,omitempty"`
	LgsID                            *string               `json:"lgsID,omitempty"`
	FirstStartDate                   common.NullableString `json:"FirstStartDate,omitempty"`
	Partition                        *int64                `json:"Partition,omitempty"`
	Topic                            *string               `json:"Topic,omitempty"`
	ApiUrl                           *string               `json:"ApiUrl,omitempty"`
	IdFieldType                      *string               `json:"IdFieldType,omitempty"`
	IdFieldName                      *string               `json:"IdFieldName,omitempty"`
	ReturnType                       *string               `json:"ReturnType,omitempty"`
	TenantId                         *string               `json:"TenantId,omitempty"`
	ClientId                         *string               `json:"ClientId,omitempty"`
	ClientSecret                     *string               `json:"ClientSecret,omitempty"`
	AccessToken                      common.NullableString `json:"AccessToken,omitempty"`
	EventDateFieldName               common.NullableString `json:"EventDateFieldName,omitempty"`
	LogCharLength                    *int64                `json:"LogCharLength,omitempty"`
	DhcpFileNames                    *string               `json:"DhcpFileNames,omitempty"`
	DhcpChangeTime                   *string               `json:"DhcpChangeTime,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourcesLogReaderData creates a new SourcesLogReaderData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcesLogReaderData() *SourcesLogReaderData {
	this := SourcesLogReaderData{}
	return &this
}

// NewSourcesLogReaderDataWithDefaults creates a new SourcesLogReaderData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcesLogReaderDataWithDefaults() *SourcesLogReaderData {
	this := SourcesLogReaderData{}
	return &this
}

// GetFileInfo returns the FileInfo field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetFileInfo() SourcesFileInfo {
	if o == nil || o.FileInfo == nil {
		var ret SourcesFileInfo
		return ret
	}
	return *o.FileInfo
}

// GetFileInfoOk returns a tuple with the FileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetFileInfoOk() (*SourcesFileInfo, bool) {
	if o == nil || o.FileInfo == nil {
		return nil, false
	}
	return o.FileInfo, true
}

// HasFileInfo returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasFileInfo() bool {
	return o != nil && o.FileInfo != nil
}

// SetFileInfo gets a reference to the given SourcesFileInfo and assigns it to the FileInfo field.
func (o *SourcesLogReaderData) SetFileInfo(v SourcesFileInfo) {
	o.FileInfo = &v
}

// GetMessageFieldType returns the MessageFieldType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesLogReaderData) GetMessageFieldType() string {
	if o == nil || o.MessageFieldType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MessageFieldType.Get()
}

// GetMessageFieldTypeOk returns a tuple with the MessageFieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesLogReaderData) GetMessageFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageFieldType.Get(), o.MessageFieldType.IsSet()
}

// HasMessageFieldType returns a boolean if a MessageFieldType has been set.
func (o *SourcesLogReaderData) HasMessageFieldType() bool {
	return o != nil && o.MessageFieldType.IsSet()
}

// SetMessageFieldType gets a reference to the given common.NullableString and assigns it to the MessageFieldType field.
func (o *SourcesLogReaderData) SetMessageFieldType(v string) {
	o.MessageFieldType.Set(&v)
}

// SetMessageFieldTypeNil sets the value for MessageFieldType to be an explicit nil.
func (o *SourcesLogReaderData) SetMessageFieldTypeNil() {
	o.MessageFieldType.Set(nil)
}

// UnSetMessageFieldType ensures that no value is present for MessageFieldType, not even an explicit nil.
func (o *SourcesLogReaderData) UnSetMessageFieldType() {
	o.MessageFieldType.UnSet()
}

// GetTransformFrequencyInSeconds returns the TransformFrequencyInSeconds field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetTransformFrequencyInSeconds() int64 {
	if o == nil || o.TransformFrequencyInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.TransformFrequencyInSeconds
}

// GetTransformFrequencyInSecondsOk returns a tuple with the TransformFrequencyInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetTransformFrequencyInSecondsOk() (*int64, bool) {
	if o == nil || o.TransformFrequencyInSeconds == nil {
		return nil, false
	}
	return o.TransformFrequencyInSeconds, true
}

// HasTransformFrequencyInSeconds returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasTransformFrequencyInSeconds() bool {
	return o != nil && o.TransformFrequencyInSeconds != nil
}

// SetTransformFrequencyInSeconds gets a reference to the given int64 and assigns it to the TransformFrequencyInSeconds field.
func (o *SourcesLogReaderData) SetTransformFrequencyInSeconds(v int64) {
	o.TransformFrequencyInSeconds = &v
}

// GetDeleteArchivedLogFileAfterNHours returns the DeleteArchivedLogFileAfterNHours field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetDeleteArchivedLogFileAfterNHours() int64 {
	if o == nil || o.DeleteArchivedLogFileAfterNHours == nil {
		var ret int64
		return ret
	}
	return *o.DeleteArchivedLogFileAfterNHours
}

// GetDeleteArchivedLogFileAfterNHoursOk returns a tuple with the DeleteArchivedLogFileAfterNHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetDeleteArchivedLogFileAfterNHoursOk() (*int64, bool) {
	if o == nil || o.DeleteArchivedLogFileAfterNHours == nil {
		return nil, false
	}
	return o.DeleteArchivedLogFileAfterNHours, true
}

// HasDeleteArchivedLogFileAfterNHours returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasDeleteArchivedLogFileAfterNHours() bool {
	return o != nil && o.DeleteArchivedLogFileAfterNHours != nil
}

// SetDeleteArchivedLogFileAfterNHours gets a reference to the given int64 and assigns it to the DeleteArchivedLogFileAfterNHours field.
func (o *SourcesLogReaderData) SetDeleteArchivedLogFileAfterNHours(v int64) {
	o.DeleteArchivedLogFileAfterNHours = &v
}

// GetDeleteOutputFileAfterNHours returns the DeleteOutputFileAfterNHours field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetDeleteOutputFileAfterNHours() int64 {
	if o == nil || o.DeleteOutputFileAfterNHours == nil {
		var ret int64
		return ret
	}
	return *o.DeleteOutputFileAfterNHours
}

// GetDeleteOutputFileAfterNHoursOk returns a tuple with the DeleteOutputFileAfterNHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetDeleteOutputFileAfterNHoursOk() (*int64, bool) {
	if o == nil || o.DeleteOutputFileAfterNHours == nil {
		return nil, false
	}
	return o.DeleteOutputFileAfterNHours, true
}

// HasDeleteOutputFileAfterNHours returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasDeleteOutputFileAfterNHours() bool {
	return o != nil && o.DeleteOutputFileAfterNHours != nil
}

// SetDeleteOutputFileAfterNHours gets a reference to the given int64 and assigns it to the DeleteOutputFileAfterNHours field.
func (o *SourcesLogReaderData) SetDeleteOutputFileAfterNHours(v int64) {
	o.DeleteOutputFileAfterNHours = &v
}

// GetMaxParallelProcessing returns the MaxParallelProcessing field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetMaxParallelProcessing() int64 {
	if o == nil || o.MaxParallelProcessing == nil {
		var ret int64
		return ret
	}
	return *o.MaxParallelProcessing
}

// GetMaxParallelProcessingOk returns a tuple with the MaxParallelProcessing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetMaxParallelProcessingOk() (*int64, bool) {
	if o == nil || o.MaxParallelProcessing == nil {
		return nil, false
	}
	return o.MaxParallelProcessing, true
}

// HasMaxParallelProcessing returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasMaxParallelProcessing() bool {
	return o != nil && o.MaxParallelProcessing != nil
}

// SetMaxParallelProcessing gets a reference to the given int64 and assigns it to the MaxParallelProcessing field.
func (o *SourcesLogReaderData) SetMaxParallelProcessing(v int64) {
	o.MaxParallelProcessing = &v
}

// GetEventLogNameOrFileName returns the EventLogNameOrFileName field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetEventLogNameOrFileName() int64 {
	if o == nil || o.EventLogNameOrFileName == nil {
		var ret int64
		return ret
	}
	return *o.EventLogNameOrFileName
}

// GetEventLogNameOrFileNameOk returns a tuple with the EventLogNameOrFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetEventLogNameOrFileNameOk() (*int64, bool) {
	if o == nil || o.EventLogNameOrFileName == nil {
		return nil, false
	}
	return o.EventLogNameOrFileName, true
}

// HasEventLogNameOrFileName returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasEventLogNameOrFileName() bool {
	return o != nil && o.EventLogNameOrFileName != nil
}

// SetEventLogNameOrFileName gets a reference to the given int64 and assigns it to the EventLogNameOrFileName field.
func (o *SourcesLogReaderData) SetEventLogNameOrFileName(v int64) {
	o.EventLogNameOrFileName = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesLogReaderData) GetFilter() string {
	if o == nil || o.Filter.Get() == nil {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesLogReaderData) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a Filter has been set.
func (o *SourcesLogReaderData) HasFilter() bool {
	return o != nil && o.Filter.IsSet()
}

// SetFilter gets a reference to the given common.NullableString and assigns it to the Filter field.
func (o *SourcesLogReaderData) SetFilter(v string) {
	o.Filter.Set(&v)
}

// SetFilterNil sets the value for Filter to be an explicit nil.
func (o *SourcesLogReaderData) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnSetFilter ensures that no value is present for Filter, not even an explicit nil.
func (o *SourcesLogReaderData) UnSetFilter() {
	o.Filter.UnSet()
}

// GetDataFormat returns the DataFormat field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetDataFormat() string {
	if o == nil || o.DataFormat == nil {
		var ret string
		return ret
	}
	return *o.DataFormat
}

// GetDataFormatOk returns a tuple with the DataFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetDataFormatOk() (*string, bool) {
	if o == nil || o.DataFormat == nil {
		return nil, false
	}
	return o.DataFormat, true
}

// HasDataFormat returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasDataFormat() bool {
	return o != nil && o.DataFormat != nil
}

// SetDataFormat gets a reference to the given string and assigns it to the DataFormat field.
func (o *SourcesLogReaderData) SetDataFormat(v string) {
	o.DataFormat = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetCredential() string {
	if o == nil || o.Credential == nil {
		var ret string
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetCredentialOk() (*string, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasCredential() bool {
	return o != nil && o.Credential != nil
}

// SetCredential gets a reference to the given string and assigns it to the Credential field.
func (o *SourcesLogReaderData) SetCredential(v string) {
	o.Credential = &v
}

// GetIdColumnType returns the IdColumnType field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetIdColumnType() string {
	if o == nil || o.IdColumnType == nil {
		var ret string
		return ret
	}
	return *o.IdColumnType
}

// GetIdColumnTypeOk returns a tuple with the IdColumnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetIdColumnTypeOk() (*string, bool) {
	if o == nil || o.IdColumnType == nil {
		return nil, false
	}
	return o.IdColumnType, true
}

// HasIdColumnType returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasIdColumnType() bool {
	return o != nil && o.IdColumnType != nil
}

// SetIdColumnType gets a reference to the given string and assigns it to the IdColumnType field.
func (o *SourcesLogReaderData) SetIdColumnType(v string) {
	o.IdColumnType = &v
}

// GetParseJson returns the ParseJson field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetParseJson() bool {
	if o == nil || o.ParseJson == nil {
		var ret bool
		return ret
	}
	return *o.ParseJson
}

// GetParseJsonOk returns a tuple with the ParseJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetParseJsonOk() (*bool, bool) {
	if o == nil || o.ParseJson == nil {
		return nil, false
	}
	return o.ParseJson, true
}

// HasParseJson returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasParseJson() bool {
	return o != nil && o.ParseJson != nil
}

// SetParseJson gets a reference to the given bool and assigns it to the ParseJson field.
func (o *SourcesLogReaderData) SetParseJson(v bool) {
	o.ParseJson = &v
}

// GetIndexName returns the IndexName field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetIndexName() string {
	if o == nil || o.IndexName == nil {
		var ret string
		return ret
	}
	return *o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetIndexNameOk() (*string, bool) {
	if o == nil || o.IndexName == nil {
		return nil, false
	}
	return o.IndexName, true
}

// HasIndexName returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasIndexName() bool {
	return o != nil && o.IndexName != nil
}

// SetIndexName gets a reference to the given string and assigns it to the IndexName field.
func (o *SourcesLogReaderData) SetIndexName(v string) {
	o.IndexName = &v
}

// GetIdColumnName returns the IdColumnName field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetIdColumnName() string {
	if o == nil || o.IdColumnName == nil {
		var ret string
		return ret
	}
	return *o.IdColumnName
}

// GetIdColumnNameOk returns a tuple with the IdColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetIdColumnNameOk() (*string, bool) {
	if o == nil || o.IdColumnName == nil {
		return nil, false
	}
	return o.IdColumnName, true
}

// HasIdColumnName returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasIdColumnName() bool {
	return o != nil && o.IdColumnName != nil
}

// SetIdColumnName gets a reference to the given string and assigns it to the IdColumnName field.
func (o *SourcesLogReaderData) SetIdColumnName(v string) {
	o.IdColumnName = &v
}

// GetIdColumnFormat returns the IdColumnFormat field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetIdColumnFormat() string {
	if o == nil || o.IdColumnFormat == nil {
		var ret string
		return ret
	}
	return *o.IdColumnFormat
}

// GetIdColumnFormatOk returns a tuple with the IdColumnFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetIdColumnFormatOk() (*string, bool) {
	if o == nil || o.IdColumnFormat == nil {
		return nil, false
	}
	return o.IdColumnFormat, true
}

// HasIdColumnFormat returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasIdColumnFormat() bool {
	return o != nil && o.IdColumnFormat != nil
}

// SetIdColumnFormat gets a reference to the given string and assigns it to the IdColumnFormat field.
func (o *SourcesLogReaderData) SetIdColumnFormat(v string) {
	o.IdColumnFormat = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SourcesLogReaderData) SetSource(v string) {
	o.Source = &v
}

// GetLgsID returns the LgsID field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetLgsID() string {
	if o == nil || o.LgsID == nil {
		var ret string
		return ret
	}
	return *o.LgsID
}

// GetLgsIDOk returns a tuple with the LgsID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetLgsIDOk() (*string, bool) {
	if o == nil || o.LgsID == nil {
		return nil, false
	}
	return o.LgsID, true
}

// HasLgsID returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasLgsID() bool {
	return o != nil && o.LgsID != nil
}

// SetLgsID gets a reference to the given string and assigns it to the LgsID field.
func (o *SourcesLogReaderData) SetLgsID(v string) {
	o.LgsID = &v
}

// GetFirstStartDate returns the FirstStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesLogReaderData) GetFirstStartDate() string {
	if o == nil || o.FirstStartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstStartDate.Get()
}

// GetFirstStartDateOk returns a tuple with the FirstStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesLogReaderData) GetFirstStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstStartDate.Get(), o.FirstStartDate.IsSet()
}

// HasFirstStartDate returns a boolean if a FirstStartDate has been set.
func (o *SourcesLogReaderData) HasFirstStartDate() bool {
	return o != nil && o.FirstStartDate.IsSet()
}

// SetFirstStartDate gets a reference to the given common.NullableString and assigns it to the FirstStartDate field.
func (o *SourcesLogReaderData) SetFirstStartDate(v string) {
	o.FirstStartDate.Set(&v)
}

// SetFirstStartDateNil sets the value for FirstStartDate to be an explicit nil.
func (o *SourcesLogReaderData) SetFirstStartDateNil() {
	o.FirstStartDate.Set(nil)
}

// UnSetFirstStartDate ensures that no value is present for FirstStartDate, not even an explicit nil.
func (o *SourcesLogReaderData) UnSetFirstStartDate() {
	o.FirstStartDate.UnSet()
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetPartition() int64 {
	if o == nil || o.Partition == nil {
		var ret int64
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetPartitionOk() (*int64, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasPartition() bool {
	return o != nil && o.Partition != nil
}

// SetPartition gets a reference to the given int64 and assigns it to the Partition field.
func (o *SourcesLogReaderData) SetPartition(v int64) {
	o.Partition = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasTopic() bool {
	return o != nil && o.Topic != nil
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *SourcesLogReaderData) SetTopic(v string) {
	o.Topic = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetApiUrl() string {
	if o == nil || o.ApiUrl == nil {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetApiUrlOk() (*string, bool) {
	if o == nil || o.ApiUrl == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasApiUrl() bool {
	return o != nil && o.ApiUrl != nil
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *SourcesLogReaderData) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetIdFieldType returns the IdFieldType field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetIdFieldType() string {
	if o == nil || o.IdFieldType == nil {
		var ret string
		return ret
	}
	return *o.IdFieldType
}

// GetIdFieldTypeOk returns a tuple with the IdFieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetIdFieldTypeOk() (*string, bool) {
	if o == nil || o.IdFieldType == nil {
		return nil, false
	}
	return o.IdFieldType, true
}

// HasIdFieldType returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasIdFieldType() bool {
	return o != nil && o.IdFieldType != nil
}

// SetIdFieldType gets a reference to the given string and assigns it to the IdFieldType field.
func (o *SourcesLogReaderData) SetIdFieldType(v string) {
	o.IdFieldType = &v
}

// GetIdFieldName returns the IdFieldName field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetIdFieldName() string {
	if o == nil || o.IdFieldName == nil {
		var ret string
		return ret
	}
	return *o.IdFieldName
}

// GetIdFieldNameOk returns a tuple with the IdFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetIdFieldNameOk() (*string, bool) {
	if o == nil || o.IdFieldName == nil {
		return nil, false
	}
	return o.IdFieldName, true
}

// HasIdFieldName returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasIdFieldName() bool {
	return o != nil && o.IdFieldName != nil
}

// SetIdFieldName gets a reference to the given string and assigns it to the IdFieldName field.
func (o *SourcesLogReaderData) SetIdFieldName(v string) {
	o.IdFieldName = &v
}

// GetReturnType returns the ReturnType field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetReturnType() string {
	if o == nil || o.ReturnType == nil {
		var ret string
		return ret
	}
	return *o.ReturnType
}

// GetReturnTypeOk returns a tuple with the ReturnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetReturnTypeOk() (*string, bool) {
	if o == nil || o.ReturnType == nil {
		return nil, false
	}
	return o.ReturnType, true
}

// HasReturnType returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasReturnType() bool {
	return o != nil && o.ReturnType != nil
}

// SetReturnType gets a reference to the given string and assigns it to the ReturnType field.
func (o *SourcesLogReaderData) SetReturnType(v string) {
	o.ReturnType = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasTenantId() bool {
	return o != nil && o.TenantId != nil
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *SourcesLogReaderData) SetTenantId(v string) {
	o.TenantId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SourcesLogReaderData) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasClientSecret() bool {
	return o != nil && o.ClientSecret != nil
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *SourcesLogReaderData) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesLogReaderData) GetAccessToken() string {
	if o == nil || o.AccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesLogReaderData) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a AccessToken has been set.
func (o *SourcesLogReaderData) HasAccessToken() bool {
	return o != nil && o.AccessToken.IsSet()
}

// SetAccessToken gets a reference to the given common.NullableString and assigns it to the AccessToken field.
func (o *SourcesLogReaderData) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}

// SetAccessTokenNil sets the value for AccessToken to be an explicit nil.
func (o *SourcesLogReaderData) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnSetAccessToken ensures that no value is present for AccessToken, not even an explicit nil.
func (o *SourcesLogReaderData) UnSetAccessToken() {
	o.AccessToken.UnSet()
}

// GetEventDateFieldName returns the EventDateFieldName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourcesLogReaderData) GetEventDateFieldName() string {
	if o == nil || o.EventDateFieldName.Get() == nil {
		var ret string
		return ret
	}
	return *o.EventDateFieldName.Get()
}

// GetEventDateFieldNameOk returns a tuple with the EventDateFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourcesLogReaderData) GetEventDateFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventDateFieldName.Get(), o.EventDateFieldName.IsSet()
}

// HasEventDateFieldName returns a boolean if a EventDateFieldName has been set.
func (o *SourcesLogReaderData) HasEventDateFieldName() bool {
	return o != nil && o.EventDateFieldName.IsSet()
}

// SetEventDateFieldName gets a reference to the given common.NullableString and assigns it to the EventDateFieldName field.
func (o *SourcesLogReaderData) SetEventDateFieldName(v string) {
	o.EventDateFieldName.Set(&v)
}

// SetEventDateFieldNameNil sets the value for EventDateFieldName to be an explicit nil.
func (o *SourcesLogReaderData) SetEventDateFieldNameNil() {
	o.EventDateFieldName.Set(nil)
}

// UnSetEventDateFieldName ensures that no value is present for EventDateFieldName, not even an explicit nil.
func (o *SourcesLogReaderData) UnSetEventDateFieldName() {
	o.EventDateFieldName.UnSet()
}

// GetLogCharLength returns the LogCharLength field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetLogCharLength() int64 {
	if o == nil || o.LogCharLength == nil {
		var ret int64
		return ret
	}
	return *o.LogCharLength
}

// GetLogCharLengthOk returns a tuple with the LogCharLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetLogCharLengthOk() (*int64, bool) {
	if o == nil || o.LogCharLength == nil {
		return nil, false
	}
	return o.LogCharLength, true
}

// HasLogCharLength returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasLogCharLength() bool {
	return o != nil && o.LogCharLength != nil
}

// SetLogCharLength gets a reference to the given int64 and assigns it to the LogCharLength field.
func (o *SourcesLogReaderData) SetLogCharLength(v int64) {
	o.LogCharLength = &v
}

// GetDhcpFileNames returns the DhcpFileNames field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetDhcpFileNames() string {
	if o == nil || o.DhcpFileNames == nil {
		var ret string
		return ret
	}
	return *o.DhcpFileNames
}

// GetDhcpFileNamesOk returns a tuple with the DhcpFileNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetDhcpFileNamesOk() (*string, bool) {
	if o == nil || o.DhcpFileNames == nil {
		return nil, false
	}
	return o.DhcpFileNames, true
}

// HasDhcpFileNames returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasDhcpFileNames() bool {
	return o != nil && o.DhcpFileNames != nil
}

// SetDhcpFileNames gets a reference to the given string and assigns it to the DhcpFileNames field.
func (o *SourcesLogReaderData) SetDhcpFileNames(v string) {
	o.DhcpFileNames = &v
}

// GetDhcpChangeTime returns the DhcpChangeTime field value if set, zero value otherwise.
func (o *SourcesLogReaderData) GetDhcpChangeTime() string {
	if o == nil || o.DhcpChangeTime == nil {
		var ret string
		return ret
	}
	return *o.DhcpChangeTime
}

// GetDhcpChangeTimeOk returns a tuple with the DhcpChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesLogReaderData) GetDhcpChangeTimeOk() (*string, bool) {
	if o == nil || o.DhcpChangeTime == nil {
		return nil, false
	}
	return o.DhcpChangeTime, true
}

// HasDhcpChangeTime returns a boolean if a field has been set.
func (o *SourcesLogReaderData) HasDhcpChangeTime() bool {
	return o != nil && o.DhcpChangeTime != nil
}

// SetDhcpChangeTime gets a reference to the given string and assigns it to the DhcpChangeTime field.
func (o *SourcesLogReaderData) SetDhcpChangeTime(v string) {
	o.DhcpChangeTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcesLogReaderData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.FileInfo != nil {
		toSerialize["FileInfo"] = o.FileInfo
	}
	if o.MessageFieldType.IsSet() {
		toSerialize["MessageFieldType"] = o.MessageFieldType.Get()
	}
	if o.TransformFrequencyInSeconds != nil {
		toSerialize["TransformFrequencyInSeconds"] = o.TransformFrequencyInSeconds
	}
	if o.DeleteArchivedLogFileAfterNHours != nil {
		toSerialize["DeleteArchivedLogFileAfterNHours"] = o.DeleteArchivedLogFileAfterNHours
	}
	if o.DeleteOutputFileAfterNHours != nil {
		toSerialize["DeleteOutputFileAfterNHours"] = o.DeleteOutputFileAfterNHours
	}
	if o.MaxParallelProcessing != nil {
		toSerialize["MaxParallelProcessing"] = o.MaxParallelProcessing
	}
	if o.EventLogNameOrFileName != nil {
		toSerialize["EventLogNameOrFileName"] = o.EventLogNameOrFileName
	}
	if o.Filter.IsSet() {
		toSerialize["Filter"] = o.Filter.Get()
	}
	if o.DataFormat != nil {
		toSerialize["DataFormat"] = o.DataFormat
	}
	if o.Credential != nil {
		toSerialize["Credential"] = o.Credential
	}
	if o.IdColumnType != nil {
		toSerialize["IdColumnType"] = o.IdColumnType
	}
	if o.ParseJson != nil {
		toSerialize["ParseJson"] = o.ParseJson
	}
	if o.IndexName != nil {
		toSerialize["IndexName"] = o.IndexName
	}
	if o.IdColumnName != nil {
		toSerialize["IdColumnName"] = o.IdColumnName
	}
	if o.IdColumnFormat != nil {
		toSerialize["IdColumnFormat"] = o.IdColumnFormat
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.LgsID != nil {
		toSerialize["lgsID"] = o.LgsID
	}
	if o.FirstStartDate.IsSet() {
		toSerialize["FirstStartDate"] = o.FirstStartDate.Get()
	}
	if o.Partition != nil {
		toSerialize["Partition"] = o.Partition
	}
	if o.Topic != nil {
		toSerialize["Topic"] = o.Topic
	}
	if o.ApiUrl != nil {
		toSerialize["ApiUrl"] = o.ApiUrl
	}
	if o.IdFieldType != nil {
		toSerialize["IdFieldType"] = o.IdFieldType
	}
	if o.IdFieldName != nil {
		toSerialize["IdFieldName"] = o.IdFieldName
	}
	if o.ReturnType != nil {
		toSerialize["ReturnType"] = o.ReturnType
	}
	if o.TenantId != nil {
		toSerialize["TenantId"] = o.TenantId
	}
	if o.ClientId != nil {
		toSerialize["ClientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["ClientSecret"] = o.ClientSecret
	}
	if o.AccessToken.IsSet() {
		toSerialize["AccessToken"] = o.AccessToken.Get()
	}
	if o.EventDateFieldName.IsSet() {
		toSerialize["EventDateFieldName"] = o.EventDateFieldName.Get()
	}
	if o.LogCharLength != nil {
		toSerialize["LogCharLength"] = o.LogCharLength
	}
	if o.DhcpFileNames != nil {
		toSerialize["DhcpFileNames"] = o.DhcpFileNames
	}
	if o.DhcpChangeTime != nil {
		toSerialize["DhcpChangeTime"] = o.DhcpChangeTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourcesLogReaderData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileInfo                         *SourcesFileInfo      `json:"FileInfo,omitempty"`
		MessageFieldType                 common.NullableString `json:"MessageFieldType,omitempty"`
		TransformFrequencyInSeconds      *int64                `json:"TransformFrequencyInSeconds,omitempty"`
		DeleteArchivedLogFileAfterNHours *int64                `json:"DeleteArchivedLogFileAfterNHours,omitempty"`
		DeleteOutputFileAfterNHours      *int64                `json:"DeleteOutputFileAfterNHours,omitempty"`
		MaxParallelProcessing            *int64                `json:"MaxParallelProcessing,omitempty"`
		EventLogNameOrFileName           *int64                `json:"EventLogNameOrFileName,omitempty"`
		Filter                           common.NullableString `json:"Filter,omitempty"`
		DataFormat                       *string               `json:"DataFormat,omitempty"`
		Credential                       *string               `json:"Credential,omitempty"`
		IdColumnType                     *string               `json:"IdColumnType,omitempty"`
		ParseJson                        *bool                 `json:"ParseJson,omitempty"`
		IndexName                        *string               `json:"IndexName,omitempty"`
		IdColumnName                     *string               `json:"IdColumnName,omitempty"`
		IdColumnFormat                   *string               `json:"IdColumnFormat,omitempty"`
		Source                           *string               `json:"Source,omitempty"`
		LgsID                            *string               `json:"lgsID,omitempty"`
		FirstStartDate                   common.NullableString `json:"FirstStartDate,omitempty"`
		Partition                        *int64                `json:"Partition,omitempty"`
		Topic                            *string               `json:"Topic,omitempty"`
		ApiUrl                           *string               `json:"ApiUrl,omitempty"`
		IdFieldType                      *string               `json:"IdFieldType,omitempty"`
		IdFieldName                      *string               `json:"IdFieldName,omitempty"`
		ReturnType                       *string               `json:"ReturnType,omitempty"`
		TenantId                         *string               `json:"TenantId,omitempty"`
		ClientId                         *string               `json:"ClientId,omitempty"`
		ClientSecret                     *string               `json:"ClientSecret,omitempty"`
		AccessToken                      common.NullableString `json:"AccessToken,omitempty"`
		EventDateFieldName               common.NullableString `json:"EventDateFieldName,omitempty"`
		LogCharLength                    *int64                `json:"LogCharLength,omitempty"`
		DhcpFileNames                    *string               `json:"DhcpFileNames,omitempty"`
		DhcpChangeTime                   *string               `json:"DhcpChangeTime,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"FileInfo", "MessageFieldType", "TransformFrequencyInSeconds", "DeleteArchivedLogFileAfterNHours", "DeleteOutputFileAfterNHours", "MaxParallelProcessing", "EventLogNameOrFileName", "Filter", "DataFormat", "Credential", "IdColumnType", "ParseJson", "IndexName", "IdColumnName", "IdColumnFormat", "Source", "lgsID", "FirstStartDate", "Partition", "Topic", "ApiUrl", "IdFieldType", "IdFieldName", "ReturnType", "TenantId", "ClientId", "ClientSecret", "AccessToken", "EventDateFieldName", "LogCharLength", "DhcpFileNames", "DhcpChangeTime"})
	} else {
		return err
	}

	o.FileInfo = all.FileInfo
	o.MessageFieldType = all.MessageFieldType
	o.TransformFrequencyInSeconds = all.TransformFrequencyInSeconds
	o.DeleteArchivedLogFileAfterNHours = all.DeleteArchivedLogFileAfterNHours
	o.DeleteOutputFileAfterNHours = all.DeleteOutputFileAfterNHours
	o.MaxParallelProcessing = all.MaxParallelProcessing
	o.EventLogNameOrFileName = all.EventLogNameOrFileName
	o.Filter = all.Filter
	o.DataFormat = all.DataFormat
	o.Credential = all.Credential
	o.IdColumnType = all.IdColumnType
	o.ParseJson = all.ParseJson
	o.IndexName = all.IndexName
	o.IdColumnName = all.IdColumnName
	o.IdColumnFormat = all.IdColumnFormat
	o.Source = all.Source
	o.LgsID = all.LgsID
	o.FirstStartDate = all.EventDateFieldName
	o.Partition = all.Partition
	o.Topic = all.Topic
	o.ApiUrl = all.ApiUrl
	o.IdFieldType = all.IdFieldType
	o.IdFieldName = all.IdFieldName
	o.ReturnType = all.ReturnType
	o.TenantId = all.TenantId
	o.ClientId = all.ClientId
	o.ClientSecret = all.ClientSecret
	o.AccessToken = all.AccessToken
	o.EventDateFieldName = all.EventDateFieldName
	o.LogCharLength = all.LogCharLength
	o.DhcpFileNames = all.DhcpFileNames
	o.DhcpChangeTime = all.DhcpChangeTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
