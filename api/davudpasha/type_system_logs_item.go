package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SystemLogsItem represents an item in the system logs.
type SystemLogsItem struct {
	// TS specifies the timestamp of the log item.
	TS *string `json:"ts,omitempty"`
	// Severity specifies the severity level of the log item.
	Severity *string `json:"svr,omitempty"`
	// EventID specifies the event ID associated with the log item.
	EventID *string `json:"eventid,omitempty"`
	// Message contains the main message of the log item.
	Message *string `json:"msg,omitempty"`
	// Title specifies the title associated with the log item.
	Title common.NullableString `json:"title,omitempty"`
	// App specifies the application name related to the log item.
	App *string `json:"app,omitempty"`
	// Version specifies the version information related to the log item.
	Version common.NullableString `json:"ver,omitempty"`
	// Data specifies additional data associated with the log item.
	Data common.NullableString `json:"data,omitempty"`
	// EntityType specifies the entity type associated with the log item.
	EntityType *string `json:"entitytype,omitempty"`
	// EntityID specifies the entity ID associated with the log item.
	EntityID *string `json:"entityid,omitempty"`
	// EntityName specifies the entity name associated with the log item.
	EntityName common.NullableString `json:"entityname,omitempty"`
	// OptID1 specifies optional ID 1 associated with the log item.
	OptID1 *string `json:"optid1,omitempty"`
	// OptID2 specifies optional ID 2 associated with the log item.
	OptID2 *string `json:"optid2,omitempty"`
	// Link specifies a link associated with the log item.
	Link common.NullableString `json:"link,omitempty"`
	// Producer specifies the producer of the log item.
	Producer *string `json:"producer,omitempty"`
	// Reason specifies the reason associated with the log item.
	Reason common.NullableString `json:"reason,omitempty"`
	// DeviceName specifies the device name related to the log item.
	DeviceName common.NullableString `json:"dvcname,omitempty"`
	// DeviceIP specifies the device IP related to the log item.
	DeviceIP *string `json:"dvcip,omitempty"`
	// Tags specifies tags associated with the log item.
	Tags []string `json:"tags,omitempty"`
	// StackTrace specifies the stack trace associated with the log item.
	StackTrace common.NullableString `json:"stacktrace,omitempty"`
	// StackTraceJsonStyled specifies the JSON-styled stack trace associated with the log item.
	StackTraceJsonStyled common.NullableString `json:"stacktrace_JsonStyled,omitempty"`
	// TenantID specifies the ID of the tenant associated with the log item.
	TenantID common.NullableString `json:"TenantID,omitempty"`
	// ID specifies the ID of the log item.
	ID *string `json:"_id,omitempty"`
	// Timestamp specifies the timestamp associated with the log item.
	Timestamp common.NullableString `json:"Timestamp,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSystemLogsItem creates a new SystemLogsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSystemLogsItem() *SystemLogsItem {
	this := SystemLogsItem{}
	return &this
}

// NewSystemLogsItemWithDefaults creates a new SystemLogsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSystemLogsItemWithDefaults() *SystemLogsItem {
	this := SystemLogsItem{}
	return &this
}

// GetTS returns the TS field value if set, zero value otherwise.
func (o *SystemLogsItem) GetTS() string {
	if o == nil || o.TS == nil {
		var ret string
		return ret
	}
	return *o.TS
}

// GetTSOk returns a tuple with the TS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetTSOk() (*string, bool) {
	if o == nil || o.TS == nil {
		return nil, false
	}
	return o.TS, true
}

// HasTS returns a boolean if a field has been set.
func (o *SystemLogsItem) HasTS() bool {
	return o != nil && o.TS != nil
}

// SetTS gets a reference to the given string and assigns it to the TS field.
func (o *SystemLogsItem) SetTS(v string) {
	o.TS = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *SystemLogsItem) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *SystemLogsItem) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity geSeverity a reference to the given string and assigns it to the Severity field.
func (o *SystemLogsItem) SetSeverity(v string) {
	o.Severity = &v
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *SystemLogsItem) GetEventID() string {
	if o == nil || o.EventID == nil {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetEventIDOk() (*string, bool) {
	if o == nil || o.EventID == nil {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *SystemLogsItem) HasEventID() bool {
	return o != nil && o.EventID != nil
}

// SetEventID geEventID a reference to the given string and assigns it to the EventID field.
func (o *SystemLogsItem) SetEventID(v string) {
	o.EventID = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SystemLogsItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SystemLogsItem) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage geMessage a reference to the given string and assigns it to the Message field.
func (o *SystemLogsItem) SetMessage(v string) {
	o.Message = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a Title has been set.
func (o *SystemLogsItem) HasTitle() bool {
	return o != nil && o.Title.IsSet()
}

// SetTitle gets a reference to the given common.NullableString and assigns it to the Title field.
func (o *SystemLogsItem) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil.
func (o *SystemLogsItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnSetTitle ensures that no value is present for Title, not even an explicit nil.
func (o *SystemLogsItem) UnSetTitle() {
	o.Title.UnSet()
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *SystemLogsItem) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *SystemLogsItem) HasApp() bool {
	return o != nil && o.App != nil
}

// SetApp geApp a reference to the given string and assigns it to the App field.
func (o *SystemLogsItem) SetApp(v string) {
	o.App = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a Version has been set.
func (o *SystemLogsItem) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given common.NullableString and assigns it to the Version field.
func (o *SystemLogsItem) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *SystemLogsItem) SetVersionNil() {
	o.Version.Set(nil)
}

// UnSetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *SystemLogsItem) UnSetVersion() {
	o.Version.UnSet()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a Data has been set.
func (o *SystemLogsItem) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given common.NullableString and assigns it to the Data field.
func (o *SystemLogsItem) SetData(v string) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *SystemLogsItem) SetDataNil() {
	o.Data.Set(nil)
}

// UnSetData ensures that no value is present for Data, not even an explicit nil.
func (o *SystemLogsItem) UnSetData() {
	o.Data.UnSet()
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *SystemLogsItem) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *SystemLogsItem) HasEntityType() bool {
	return o != nil && o.EntityType != nil
}

// SetEntityType geEntityType a reference to the given string and assigns it to the EntityType field.
func (o *SystemLogsItem) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityID returns the EntityID field value if set, zero value otherwise.
func (o *SystemLogsItem) GetEntityID() string {
	if o == nil || o.EntityID == nil {
		var ret string
		return ret
	}
	return *o.EntityID
}

// GetEntityIDOk returns a tuple with the EntityID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetEntityIDOk() (*string, bool) {
	if o == nil || o.EntityID == nil {
		return nil, false
	}
	return o.EntityID, true
}

// HasEntityID returns a boolean if a field has been set.
func (o *SystemLogsItem) HasEntityID() bool {
	return o != nil && o.EntityID != nil
}

// SetEntityID geEntityID a reference to the given string and assigns it to the EntityID field.
func (o *SystemLogsItem) SetEntityID(v string) {
	o.EntityID = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetEntityName() string {
	if o == nil || o.EntityName.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityName.Get()
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetEntityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityName.Get(), o.EntityName.IsSet()
}

// HasEntityName returns a boolean if a EntityName has been set.
func (o *SystemLogsItem) HasEntityName() bool {
	return o != nil && o.EntityName.IsSet()
}

// SetEntityName gets a reference to the given common.NullableString and assigns it to the EntityName field.
func (o *SystemLogsItem) SetEntityName(v string) {
	o.EntityName.Set(&v)
}

// SetEntityNameNil sets the value for EntityName to be an explicit nil.
func (o *SystemLogsItem) SetEntityNameNil() {
	o.EntityName.Set(nil)
}

// UnSetEntityName ensures that no value is present for EntityName, not even an explicit nil.
func (o *SystemLogsItem) UnSetEntityName() {
	o.EntityName.UnSet()
}

// GetOptID1 returns the OptID1 field value if set, zero value otherwise.
func (o *SystemLogsItem) GetOptID1() string {
	if o == nil || o.OptID1 == nil {
		var ret string
		return ret
	}
	return *o.OptID1
}

// GetOptID1Ok returns a tuple with the OptID1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetOptID1Ok() (*string, bool) {
	if o == nil || o.OptID1 == nil {
		return nil, false
	}
	return o.OptID1, true
}

// HasOptID1 returns a boolean if a field has been set.
func (o *SystemLogsItem) HasOptID1() bool {
	return o != nil && o.OptID1 != nil
}

// SetOptID1 geOptID1 a reference to the given string and assigns it to the OptID1 field.
func (o *SystemLogsItem) SetOptID1(v string) {
	o.OptID1 = &v
}

// GetOptID2 returns the OptID2 field value if set, zero value otherwise.
func (o *SystemLogsItem) GetOptID2() string {
	if o == nil || o.OptID2 == nil {
		var ret string
		return ret
	}
	return *o.OptID2
}

// GetOptID2Ok returns a tuple with the OptID2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetOptID2Ok() (*string, bool) {
	if o == nil || o.OptID2 == nil {
		return nil, false
	}
	return o.OptID2, true
}

// HasOptID2 returns a boolean if a field has been set.
func (o *SystemLogsItem) HasOptID2() bool {
	return o != nil && o.OptID2 != nil
}

// SetOptID2 geOptID2 a reference to the given string and assigns it to the OptID2 field.
func (o *SystemLogsItem) SetOptID2(v string) {
	o.OptID2 = &v
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a Link has been set.
func (o *SystemLogsItem) HasLink() bool {
	return o != nil && o.Link.IsSet()
}

// SetLink gets a reference to the given common.NullableString and assigns it to the Link field.
func (o *SystemLogsItem) SetLink(v string) {
	o.Link.Set(&v)
}

// SetLinkNil sets the value for Link to be an explicit nil.
func (o *SystemLogsItem) SetLinkNil() {
	o.Link.Set(nil)
}

// UnSetLink ensures that no value is present for Link, not even an explicit nil.
func (o *SystemLogsItem) UnSetLink() {
	o.Link.UnSet()
}

// GetProducer returns the Producer field value if set, zero value otherwise.
func (o *SystemLogsItem) GetProducer() string {
	if o == nil || o.Producer == nil {
		var ret string
		return ret
	}
	return *o.Producer
}

// GetProducerOk returns a tuple with the Producer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetProducerOk() (*string, bool) {
	if o == nil || o.Producer == nil {
		return nil, false
	}
	return o.Producer, true
}

// HasProducer returns a boolean if a field has been set.
func (o *SystemLogsItem) HasProducer() bool {
	return o != nil && o.Producer != nil
}

// SetProducer geProducer a reference to the given string and assigns it to the Producer field.
func (o *SystemLogsItem) SetProducer(v string) {
	o.Producer = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a Reason has been set.
func (o *SystemLogsItem) HasReason() bool {
	return o != nil && o.Reason.IsSet()
}

// SetReason gets a reference to the given common.NullableString and assigns it to the Reason field.
func (o *SystemLogsItem) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil.
func (o *SystemLogsItem) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnSetReason ensures that no value is present for Reason, not even an explicit nil.
func (o *SystemLogsItem) UnSetReason() {
	o.Reason.UnSet()
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetDeviceName() string {
	if o == nil || o.DeviceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceName.Get()
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceName.Get(), o.DeviceName.IsSet()
}

// HasDeviceName returns a boolean if a DeviceName has been set.
func (o *SystemLogsItem) HasDeviceName() bool {
	return o != nil && o.DeviceName.IsSet()
}

// SetDeviceName gets a reference to the given common.NullableString and assigns it to the DeviceName field.
func (o *SystemLogsItem) SetDeviceName(v string) {
	o.DeviceName.Set(&v)
}

// SetDeviceNameNil sets the value for DeviceName to be an explicit nil.
func (o *SystemLogsItem) SetDeviceNameNil() {
	o.DeviceName.Set(nil)
}

// UnSetDeviceName ensures that no value is present for DeviceName, not even an explicit nil.
func (o *SystemLogsItem) UnSetDeviceName() {
	o.DeviceName.UnSet()
}

// GetDeviceIP returns the DeviceIP field value if set, zero value otherwise.
func (o *SystemLogsItem) GetDeviceIP() string {
	if o == nil || o.DeviceIP == nil {
		var ret string
		return ret
	}
	return *o.DeviceIP
}

// GetDeviceIPOk returns a tuple with the DeviceIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetDeviceIPOk() (*string, bool) {
	if o == nil || o.DeviceIP == nil {
		return nil, false
	}
	return o.DeviceIP, true
}

// HasDeviceIP returns a boolean if a field has been set.
func (o *SystemLogsItem) HasDeviceIP() bool {
	return o != nil && o.DeviceIP != nil
}

// SetDeviceIP geDeviceIP a reference to the given string and assigns it to the DeviceIP field.
func (o *SystemLogsItem) SetDeviceIP(v string) {
	o.DeviceIP = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SystemLogsItem) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SystemLogsItem) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SystemLogsItem) SetTags(v []string) {
	o.Tags = v
}

// GetStackTrace returns the StackTrace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetStackTrace() string {
	if o == nil || o.StackTrace.Get() == nil {
		var ret string
		return ret
	}
	return *o.StackTrace.Get()
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetStackTraceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StackTrace.Get(), o.StackTrace.IsSet()
}

// HasStackTrace returns a boolean if a StackTrace has been set.
func (o *SystemLogsItem) HasStackTrace() bool {
	return o != nil && o.StackTrace.IsSet()
}

// SetStackTrace gets a reference to the given common.NullableString and assigns it to the StackTrace field.
func (o *SystemLogsItem) SetStackTrace(v string) {
	o.StackTrace.Set(&v)
}

// SetStackTraceNil sets the value for StackTrace to be an explicit nil.
func (o *SystemLogsItem) SetStackTraceNil() {
	o.StackTrace.Set(nil)
}

// UnSetStackTrace ensures that no value is present for StackTrace, not even an explicit nil.
func (o *SystemLogsItem) UnSetStackTrace() {
	o.StackTrace.UnSet()
}

// GetStackTraceJsonStyled returns the StackTraceJsonStyled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetStackTraceJsonStyled() string {
	if o == nil || o.StackTraceJsonStyled.Get() == nil {
		var ret string
		return ret
	}
	return *o.StackTraceJsonStyled.Get()
}

// GetStackTraceJsonStyledOk returns a tuple with the StackTraceJsonStyled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetStackTraceJsonStyledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StackTraceJsonStyled.Get(), o.StackTraceJsonStyled.IsSet()
}

// HasStackTraceJsonStyled returns a boolean if a StackTraceJsonStyled has been set.
func (o *SystemLogsItem) HasStackTraceJsonStyled() bool {
	return o != nil && o.StackTraceJsonStyled.IsSet()
}

// SetStackTraceJsonStyled gets a reference to the given common.NullableString and assigns it to the StackTraceJsonStyled field.
func (o *SystemLogsItem) SetStackTraceJsonStyled(v string) {
	o.StackTraceJsonStyled.Set(&v)
}

// SetStackTraceJsonStyledNil sets the value for StackTraceJsonStyled to be an explicit nil.
func (o *SystemLogsItem) SetStackTraceJsonStyledNil() {
	o.StackTraceJsonStyled.Set(nil)
}

// UnSetStackTraceJsonStyled ensures that no value is present for StackTraceJsonStyled, not even an explicit nil.
func (o *SystemLogsItem) UnSetStackTraceJsonStyled() {
	o.StackTraceJsonStyled.UnSet()
}

// GetTenantID returns the TenantID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetTenantID() string {
	if o == nil || o.TenantID.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantID.Get()
}

// GetTenantIDOk returns a tuple with the TenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetTenantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantID.Get(), o.TenantID.IsSet()
}

// HasTenantID returns a boolean if a TenantID has been set.
func (o *SystemLogsItem) HasTenantID() bool {
	return o != nil && o.TenantID.IsSet()
}

// SetTenantID gets a reference to the given common.NullableString and assigns it to the TenantID field.
func (o *SystemLogsItem) SetTenantID(v string) {
	o.TenantID.Set(&v)
}

// SetTenantIDNil sets the value for TenantID to be an explicit nil.
func (o *SystemLogsItem) SetTenantIDNil() {
	o.TenantID.Set(nil)
}

// UnSetTenantID ensures that no value is present for TenantID, not even an explicit nil.
func (o *SystemLogsItem) UnSetTenantID() {
	o.TenantID.UnSet()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *SystemLogsItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLogsItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *SystemLogsItem) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID geID a reference to the given string and assigns it to the ID field.
func (o *SystemLogsItem) SetID(v string) {
	o.ID = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemLogsItem) GetTimestamp() string {
	if o == nil || o.Timestamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SystemLogsItem) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a Timestamp has been set.
func (o *SystemLogsItem) HasTimestamp() bool {
	return o != nil && o.Timestamp.IsSet()
}

// SetTimestamp gets a reference to the given common.NullableString and assigns it to the Timestamp field.
func (o *SystemLogsItem) SetTimestamp(v string) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil.
func (o *SystemLogsItem) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnSetTimestamp ensures that no value is present for Timestamp, not even an explicit nil.
func (o *SystemLogsItem) UnSetTimestamp() {
	o.Timestamp.UnSet()
}

// MarshalJSON serializes the struct using spec logic.
func (o SystemLogsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.TS != nil {
		toSerialize["ts"] = o.TS
	}
	if o.Severity != nil {
		toSerialize["svr"] = o.Severity
	}
	if o.EventID != nil {
		toSerialize["eventid"] = o.EventID
	}
	if o.Message != nil {
		toSerialize["msg"] = o.Message
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Version.IsSet() {
		toSerialize["ver"] = o.Version.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.EntityType != nil {
		toSerialize["entitytype"] = o.EntityType
	}
	if o.EntityID != nil {
		toSerialize["entityid"] = o.EntityID
	}
	if o.EntityName.IsSet() {
		toSerialize["entityname"] = o.EntityName.Get()
	}
	if o.OptID1 != nil {
		toSerialize["optid1"] = o.OptID1
	}
	if o.OptID2 != nil {
		toSerialize["optid2"] = o.OptID2
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if o.Producer != nil {
		toSerialize["producer"] = o.Producer
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.DeviceName.IsSet() {
		toSerialize["dvcname"] = o.DeviceName.Get()
	}
	if o.DeviceIP != nil {
		toSerialize["dvcip"] = o.DeviceIP
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.StackTrace.IsSet() {
		toSerialize["stacktrace"] = o.StackTrace.Get()
	}
	if o.StackTraceJsonStyled.IsSet() {
		toSerialize["stacktrace_JsonStyled"] = o.StackTraceJsonStyled.Get()
	}
	if o.TenantID.IsSet() {
		toSerialize["TenantID"] = o.TenantID.Get()
	}
	if o.ID != nil {
		toSerialize["_id"] = o.ID
	}
	if o.Timestamp.IsSet() {
		toSerialize["SuccessItems"] = o.Timestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SystemLogsItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		TS                   *string               `json:"ts,omitempty"`
		Severity             *string               `json:"svr,omitempty"`
		EventID              *string               `json:"eventid,omitempty"`
		Message              *string               `json:"msg,omitempty"`
		Title                common.NullableString `json:"title,omitempty"`
		App                  *string               `json:"app,omitempty"`
		Version              common.NullableString `json:"ver,omitempty"`
		Data                 common.NullableString `json:"data,omitempty"`
		EntityType           *string               `json:"entitytype,omitempty"`
		EntityID             *string               `json:"entityid,omitempty"`
		EntityName           common.NullableString `json:"entityname,omitempty"`
		OptID1               *string               `json:"optid1,omitempty"`
		OptID2               *string               `json:"optid2,omitempty"`
		Link                 common.NullableString `json:"link,omitempty"`
		Producer             *string               `json:"producer,omitempty"`
		Reason               common.NullableString `json:"reason,omitempty"`
		DeviceName           common.NullableString `json:"dvcname,omitempty"`
		DeviceIP             *string               `json:"dvcip,omitempty"`
		Tags                 []string              `json:"tags,omitempty"`
		StackTrace           common.NullableString `json:"stacktrace,omitempty"`
		StackTraceJsonStyled common.NullableString `json:"stacktrace_JsonStyled,omitempty"`
		TenantID             common.NullableString `json:"TenantID,omitempty"`
		ID                   *string               `json:"_id,omitempty"`
		Timestamp            common.NullableString `json:"Timestamp,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ts", "svr", "eventid", "msg", "title", "app", "ver", "data", "entitytype", "entityid", "entityname", "optid1", "optid2", "link", "producer", "reason", "dvcname", "dvcip", "tags", "stacktrace", "stacktrace_JsonStyled", "TenantID", "_id", "Timestamp"})
	} else {
		return err
	}

	o.TS = all.TS
	o.Severity = all.Severity
	o.EventID = all.EventID
	o.Message = all.Message
	o.Title = all.Title
	o.App = all.App
	o.Version = all.Version
	o.Data = all.Data
	o.EntityType = all.EntityType
	o.EntityID = all.EntityID
	o.EntityName = all.EntityName
	o.OptID1 = all.OptID1
	o.OptID2 = all.OptID2
	o.Link = all.Link
	o.Producer = all.Producer
	o.Reason = all.Reason
	o.DeviceName = all.DeviceName
	o.DeviceIP = all.DeviceIP
	o.Tags = all.Tags
	o.StackTrace = all.StackTrace
	o.StackTraceJsonStyled = all.StackTraceJsonStyled
	o.TenantID = all.TenantID
	o.ID = all.ID

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
