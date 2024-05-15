package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ParallelOptions struct {
	IsActive        *bool  `json:"IsActive,omitempty"`
	TaskCount       *int64 `json:"TaskCount,omitempty"`
	FileBlockSizeMb *int64 `json:"FileBlockSizeMb,omitempty"`
	TimeoutSec      *int64 `json:"TimeoutSec,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{} `json:"-"`
	// AdditionalProperties stores any additional properties not explicitly defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewParallelOptions creates a new ParallelOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParallelOptions() *ParallelOptions {
	this := ParallelOptions{}
	return &this
}

// NewParallelOptionsWithDefault creates a new ParallelOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParallelOptionsWithDefaults() *ParallelOptions {
	this := ParallelOptions{}
	return &this
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ParallelOptions) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParallelOptions) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ParallelOptions) HasIsActive() bool {
	return o != nil && o.IsActive != nil
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ParallelOptions) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetTaskCount returns the TaskCount field value if set, zero value otherwise.
func (o *ParallelOptions) GetTaskCount() int64 {
	if o == nil || o.TaskCount == nil {
		var ret int64
		return ret
	}
	return *o.TaskCount
}

// GetTaskCountOk returns a tuple with the TaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParallelOptions) GetTaskCountOk() (*int64, bool) {
	if o == nil || o.TaskCount == nil {
		return nil, false
	}
	return o.TaskCount, true
}

// HasTaskCount returns a boolean if a field has been set.
func (o *ParallelOptions) HasTaskCount() bool {
	return o != nil && o.TaskCount != nil
}

// SetTaskCount gets a reference to the given int64 and assigns it to the TaskCount field.
func (o *ParallelOptions) SetTaskCount(v int64) {
	o.TaskCount = &v
}

// GetFileBlockSizeMb returns the FileBlockSizeMb field value if set, zero value otherwise.
func (o *ParallelOptions) GetFileBlockSizeMb() int64 {
	if o == nil || o.FileBlockSizeMb == nil {
		var ret int64
		return ret
	}
	return *o.FileBlockSizeMb
}

// GetFileBlockSizeMbOk returns a tuple with the FileBlockSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParallelOptions) GetFileBlockSizeMbOk() (*int64, bool) {
	if o == nil || o.FileBlockSizeMb == nil {
		return nil, false
	}
	return o.FileBlockSizeMb, true
}

// HasFileBlockSizeMb returns a boolean if a field has been set.
func (o *ParallelOptions) HasFileBlockSizeMb() bool {
	return o != nil && o.FileBlockSizeMb != nil
}

// SetFileBlockSizeMb gets a reference to the given int64 and assigns it to the FileBlockSizeMb field.
func (o *ParallelOptions) SetFileBlockSizeMb(v int64) {
	o.FileBlockSizeMb = &v
}

// GetTimeoutSec returns the TimeoutSec field value if set, zero value otherwise.
func (o *ParallelOptions) GetTimeoutSec() int64 {
	if o == nil || o.TimeoutSec == nil {
		var ret int64
		return ret
	}
	return *o.TimeoutSec
}

// GetTimeoutSecOk returns a tuple with the TimeoutSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParallelOptions) GetTimeoutSecOk() (*int64, bool) {
	if o == nil || o.TimeoutSec == nil {
		return nil, false
	}
	return o.TimeoutSec, true
}

// HasTimeoutSec returns a boolean if a field has been set.
func (o *ParallelOptions) HasTimeoutSec() bool {
	return o != nil && o.TimeoutSec != nil
}

// SetTimeoutSec gets a reference to the given int64 and assigns it to the TimeoutSec field.
func (o *ParallelOptions) SetTimeoutSec(v int64) {
	o.TimeoutSec = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParallelOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.IsActive != nil {
		toSerialize["IsActive"] = o.IsActive
	}
	if o.TaskCount != nil {
		toSerialize["TaskCount"] = o.TaskCount
	}
	if o.FileBlockSizeMb != nil {
		toSerialize["FileBlockSizeMb"] = o.FileBlockSizeMb
	}
	if o.TimeoutSec != nil {
		toSerialize["TimeoutSec"] = o.TimeoutSec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ParallelOptions) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsActive        *bool  `json:"IsActive,omitempty"`
		TaskCount       *int64 `json:"TaskCount,omitempty"`
		FileBlockSizeMb *int64 `json:"FileBlockSizeMb,omitempty"`
		TimeoutSec      *int64 `json:"TimeoutSec,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"IsActive", "TaskCount", "FileBlockSizeMb", "TimeoutSec"})
	} else {
		return err
	}

	o.IsActive = all.IsActive
	o.TaskCount = all.TaskCount
	o.FileBlockSizeMb = all.FileBlockSizeMb
	o.TimeoutSec = all.TimeoutSec

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
