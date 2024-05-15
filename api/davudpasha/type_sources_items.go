package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type SourcesItems struct {
	Id                  *string                       `json:"Id,omitempty"`
	Enabled             *bool                         `json:"Enabled,omitempty"`
	Name                *string                       `json:"Name,omitempty"`
	Group               common.NullableString         `json:"Group,omitempty"`
	Author              common.NullableString         `json:"Author,omitempty"`
	LogSourceDefCode    *string                       `json:"LogSourceDefCode,omitempty"`
	LogReaderType       *string                       `json:"LogReaderType,omitempty"`
	Tags                []string                      `json:"Tags,omitempty"`
	AlertTimeout        *int64                        `json:"AlertTimeout,omitempty"`
	LogReaderData       *json.RawMessage              `json:"LogReaderData,omitempty"`
	LogOperations       []SourcesLogOperations        `json:"LogOperations,omitempty"`
	DiscardedLogsConfig *string                       `json:"DiscardedLogsConfig,omitempty"`
	Value               *string                       `json:"value,omitempty"`
	Label               *string                       `json:"label,omitempty"`
	IsDeleted           *bool                         `json:"IsDeleted,omitempty"`
	IsAgentSource       *bool                         `json:"IsAgentSource,omitempty"`
	AgentIds            common.NullableList[string]   `json:"AgentIds,omitempty"`
	IndexGroupName      common.NullableString         `json:"IndexGroupName,omitempty"`
	DashboardName       *string                       `json:"dashboardName,omitempty"`
	DashboardId         common.NullableString         `json:"dashboardId,omitempty"`
	AssetTags           common.NullableList[string]   `json:"AssetTags,omitempty"`
	LogRemoveTime       common.NullableString         `json:"LogRemoveTime,omitempty"`
	LogRemoveFormat     common.NullableString         `json:"LogRemoveFormat,omitempty"`
	AgentId             common.NullableString         `json:"AgentId,omitempty"`
	WriteRawLogs        *bool                         `json:"WriteRawLogs,omitempty"`
	UseSecondaryWriter  *bool                         `json:"UseSecondaryWriter,omitempty"`
	ParallelOptions     *SourcesParallelOptions       `json:"ParallelOptions,omitempty"`
	BlockCount          *int64                        `json:"BlockCount,omitempty"`
	ScheduleConfig      NullableSourcesScheduleConfig `json:"ScheduleConfig,omitempty"`
	StoreRawLogs        *bool                         `json:"StoreRawLogs,omitempty"`
	StoreRawLogsLgs     *string                       `json:"StoreRawLogsLgs,omitempty"`
	IsEditable          *bool                         `json:"IsEditable,omitempty"`

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSourcesItems creates a new SourcesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcesItems() *SourcesItems {
	this := SourcesItems{}
	return &this
}

// NewSourcesItemsWithDefaults creates a new SourcesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcesItemsWithDefault() *SourcesItems {
	this := SourcesItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourcesItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourcesItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourcesItems) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SourcesItems) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItems) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SourcesItems) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SourcesItems) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourcesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourcesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourcesItems) SetName(v string) {
	o.Name = &v
}
