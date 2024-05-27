package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsLatestReportFile contains information about the latest report file.
type ReportsLatestReportFile struct {
	// Filename of the latest report file.
	Filename *string `json:"Filename,omitempty"`
	// Status of the latest report file.
	Status *string `json:"Status,omitempty"`
	// ID of the latest report file in GridFS.
	ReportGridFSFileId *string `json:"ReportGridFSFileId,omitempty"`
	// Run date of the latest report file.
	RunDate *string `json:"RunDate,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewReportsLatestReportFile creates a new ReportsLatestReportFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsLatestReportFile() *ReportsLatestReportFile {
	this := ReportsLatestReportFile{}
	return &this
}

// NewReportsLatestReportFileWithDefaults creates a new ReportsLatestReportFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsLatestReportFileWithDefaults() *ReportsLatestReportFile {
	this := ReportsLatestReportFile{}
	return &this
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ReportsLatestReportFile) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsLatestReportFile) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ReportsLatestReportFile) HasFilename() bool {
	return o != nil && o.Filename != nil
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ReportsLatestReportFile) SetFilename(v string) {
	o.Filename = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReportsLatestReportFile) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsLatestReportFile) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportsLatestReportFile) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReportsLatestReportFile) SetStatus(v string) {
	o.Status = &v
}

// GetReportGridFSFileId returns the ReportGridFSFileId field value if set, zero value otherwise.
func (o *ReportsLatestReportFile) GetReportGridFSFileId() string {
	if o == nil || o.ReportGridFSFileId == nil {
		var ret string
		return ret
	}
	return *o.ReportGridFSFileId
}

// GetReportGridFSFileIdOk returns a tuple with the ReportGridFSFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsLatestReportFile) GetReportGridFSFileIdOk() (*string, bool) {
	if o == nil || o.ReportGridFSFileId == nil {
		return nil, false
	}
	return o.ReportGridFSFileId, true
}

// HasReportGridFSFileId returns a boolean if a field has been set.
func (o *ReportsLatestReportFile) HasReportGridFSFileId() bool {
	return o != nil && o.ReportGridFSFileId != nil
}

// SetReportGridFSFileId gets a reference to the given string and assigns it to the ReportGridFSFileId field.
func (o *ReportsLatestReportFile) SetReportGridFSFileId(v string) {
	o.ReportGridFSFileId = &v
}

// GetRunDate returns the RunDate field value if set, zero value otherwise.
func (o *ReportsLatestReportFile) GetRunDate() string {
	if o == nil || o.RunDate == nil {
		var ret string
		return ret
	}
	return *o.RunDate
}

// GetRunDateOk returns a tuple with the RunDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsLatestReportFile) GetRunDateOk() (*string, bool) {
	if o == nil || o.RunDate == nil {
		return nil, false
	}
	return o.RunDate, true
}

// HasRunDate returns a boolean if a field has been set.
func (o *ReportsLatestReportFile) HasRunDate() bool {
	return o != nil && o.RunDate != nil
}

// SetRunDate gets a reference to the given string and assigns it to the RunDate field.
func (o *ReportsLatestReportFile) SetRunDate(v string) {
	o.RunDate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsLatestReportFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.ReportGridFSFileId != nil {
		toSerialize["ReportGridFSFileId"] = o.ReportGridFSFileId
	}
	if o.RunDate != nil {
		toSerialize["RunDate"] = o.RunDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsLatestReportFile) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filename           *string `json:"Filename,omitempty"`
		Status             *string `json:"Status,omitempty"`
		ReportGridFSFileId *string `json:"ReportGridFSFileId,omitempty"`
		RunDate            *string `json:"RunDate,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Filename", "Status", "ReportGridFSFileId", "RunDate"})
	} else {
		return err
	}

	o.Filename = all.Filename
	o.Status = all.Status
	o.ReportGridFSFileId = all.ReportGridFSFileId
	o.RunDate = all.RunDate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
