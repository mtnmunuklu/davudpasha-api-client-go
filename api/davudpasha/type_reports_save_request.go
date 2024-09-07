package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsSaveRequest represents the request structure for saving reports.
type ReportsSaveRequest struct {
	// The application to which the report belongs.
	Application *string `json:"application,omitempty"`
	// The report item to be saved.
	Report *ReportsItem `json:"report,omitempty"`
	// The schedule configuration for the report.
	Schedule *ScheduleConfig `json:"schedule,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsSaveRequest creates a new ReportsSaveRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsSaveRequest() *ReportsSaveRequest {
	this := ReportsSaveRequest{}
	return &this
}

// NewReportsSaveRequestWithDefaults creates a new ReportsSaveRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsSaveRequestWithDefaults() *ReportsSaveRequest {
	this := ReportsSaveRequest{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ReportsSaveRequest) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSaveRequest) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ReportsSaveRequest) HasApplication() bool {
	return o != nil && o.Application != nil
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ReportsSaveRequest) SetApplication(v string) {
	o.Application = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *ReportsSaveRequest) GetReport() ReportsItem {
	if o == nil || o.Report == nil {
		var ret ReportsItem
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSaveRequest) GetReportOk() (*ReportsItem, bool) {
	if o == nil || o.Report == nil {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *ReportsSaveRequest) HasReport() bool {
	return o != nil && o.Report != nil
}

// SetReport gets a reference to the given ReportsItem and assigns it to the Report field.
func (o *ReportsSaveRequest) SetReport(v ReportsItem) {
	o.Report = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ReportsSaveRequest) GetSchedule() ScheduleConfig {
	if o == nil || o.Schedule == nil {
		var ret ScheduleConfig
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSaveRequest) GetScheduleOk() (*ScheduleConfig, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ReportsSaveRequest) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given ScheduleConfig and assigns it to the Schedule field.
func (o *ReportsSaveRequest) SetSchedule(v ScheduleConfig) {
	o.Schedule = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *ReportsSaveRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsSaveRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *ReportsSaveRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *ReportsSaveRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

func (o ReportsSaveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Report != nil {
		toSerialize["report"] = o.Report
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.SmartRestRequestContext != nil {
		toSerialize["smartRestRequestContext"] = o.SmartRestRequestContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsSaveRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Application *string         `json:"application,omitempty"`
		Report      *ReportsItem    `json:"report,omitempty"`
		Schedule    *ScheduleConfig `json:"schedule,omitempty"`
		// Context for the Smart REST request.
		SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Report == nil {
		return fmt.Errorf("requiered field report is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"application", "report", "schedule", "smartRestRequestContext"})
	} else {
		return err
	}

	o.Application = all.Application
	hasInvalidField := false
	if all.Report != nil && all.Report.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Report = all.Report
	if all.Schedule != nil && all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedule = all.Schedule
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
