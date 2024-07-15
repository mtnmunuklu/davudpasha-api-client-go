package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// LookupsSearchResponse represents a search response structure for lookups.
type LookupsSearchResponse struct {
	// Identifier of the lookup.
	ID *string `json:"Id,omitempty"`
	// Name of the lookup.
	Name *string `json:"Name,omitempty"`
	// Description of the lookup.
	Description common.NullableString `json:"Description,omitempty"`
	// Other values associated with the lookup.
	OtherValues common.NullableString `json:"OtherValues,omitempty"`
	// Raw data of the lookup.
	Data *json.RawMessage `json:"Data,omitempty"`
	// Type of the lookup.
	LookupType LookupType `json:"LookupType,omitempty"`
	// Columns of the lookup.
	Columns []string `json:"Columns,omitempty"`
	// Count of results matching the lookup criteria.
	Count *int64 `json:"Count,omitempty"`
	// Last execution time of the lookup.
	LastExecutionTime *string `json:"LastExecutionTime,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewLookupsSearchResponse creates a new LookupsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLookupsSearchResponse() *LookupsSearchResponse {
	this := LookupsSearchResponse{}
	return &this
}

// NewLookupsSearchResponseWithDefaults creates a new LookupsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLookupsSearchResponseWithDefaults() *LookupsSearchResponse {
	this := LookupsSearchResponse{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *LookupsSearchResponse) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupsSearchResponse) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *LookupsSearchResponse) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *LookupsSearchResponse) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LookupsSearchResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupsSearchResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LookupsSearchResponse) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LookupsSearchResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LookupsSearchResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LookupsSearchResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a Description has been set.
func (o *LookupsSearchResponse) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *LookupsSearchResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *LookupsSearchResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnSetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *LookupsSearchResponse) UnSetDescription() {
	o.Description.UnSet()
}

// GetOtherValues returns the OtherValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LookupsSearchResponse) GetOtherValues() string {
	if o == nil || o.OtherValues.Get() == nil {
		var ret string
		return ret
	}
	return *o.OtherValues.Get()
}

// GetOtherValuesOk returns a tuple with the OtherValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LookupsSearchResponse) GetOtherValuesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtherValues.Get(), o.OtherValues.IsSet()
}

// HasOtherValues returns a boolean if a OtherValues has been set.
func (o *LookupsSearchResponse) HasOtherValues() bool {
	return o != nil && o.OtherValues.IsSet()
}

// SetOtherValues gets a reference to the given common.NullableString and assigns it to the OtherValues field.
func (o *LookupsSearchResponse) SetOtherValues(v string) {
	o.OtherValues.Set(&v)
}

// SetOtherValuesNil sets the value for OtherValues to be an explicit nil.
func (o *LookupsSearchResponse) SetOtherValuesNil() {
	o.OtherValues.Set(nil)
}

// UnSetOtherValues ensures that no value is present for OtherValues, not even an explicit nil.
func (o *LookupsSearchResponse) UnSetOtherValues() {
	o.OtherValues.UnSet()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LookupsSearchResponse) GetData() json.RawMessage {
	if o == nil || o.Data == nil {
		var ret json.RawMessage
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupsSearchResponse) GetDataOk() (*json.RawMessage, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LookupsSearchResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given json.RawMessage and assigns it to the Data field.
func (o *LookupsSearchResponse) SetData(v json.RawMessage) {
	o.Data = &v
}

// GetLookupType returns the LookupType field value if set, zero value otherwise.
func (o *LookupsSearchResponse) GetLookupType() LookupType {
	if o == nil {
		var ret LookupType
		return ret
	}
	return o.LookupType
}

// GetLookupTypeOk returns a tuple with the LookupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupsSearchResponse) GetLookupTypeOk() (*LookupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupType, true
}

// SetLookupType gets a reference to the given string and assigns it to the LookupType field.
func (o *LookupsSearchResponse) SetLookupType(v LookupType) {
	o.LookupType = v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *LookupsSearchResponse) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupsSearchResponse) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *LookupsSearchResponse) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *LookupsSearchResponse) SetColumns(v []string) {
	o.Columns = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *LookupsSearchResponse) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupsSearchResponse) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *LookupsSearchResponse) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *LookupsSearchResponse) SetCount(v int64) {
	o.Count = &v
}

// GetLastExecutionTime returns the LastExecutionTime field value if set, zero value otherwise.
func (o *LookupsSearchResponse) GetLastExecutionTime() string {
	if o == nil || o.LastExecutionTime == nil {
		var ret string
		return ret
	}
	return *o.LastExecutionTime
}

// GetLastExecutionTimeOk returns a tuple with the LastExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupsSearchResponse) GetLastExecutionTimeOk() (*string, bool) {
	if o == nil || o.LastExecutionTime == nil {
		return nil, false
	}
	return o.LastExecutionTime, true
}

// HasLastExecutionTime returns a boolean if a field has been set.
func (o *LookupsSearchResponse) HasLastExecutionTime() bool {
	return o != nil && o.LastExecutionTime != nil
}

// SetLastExecutionTime gets a reference to the given string and assigns it to the LastExecutionTime field.
func (o *LookupsSearchResponse) SetLastExecutionTime(v string) {
	o.LastExecutionTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LookupsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["Id"] = o.ID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.OtherValues.IsSet() {
		toSerialize["OtherValues"] = o.OtherValues.Get()
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if o.LookupType.IsValid() {
		toSerialize["LookupType"] = o.LookupType
	}
	if o.Columns != nil {
		toSerialize["Columns"] = o.Columns
	}
	if o.Count != nil {
		toSerialize["Name"] = o.Count
	}
	if o.LastExecutionTime != nil {
		toSerialize["LastExecutionTime"] = o.LastExecutionTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *LookupsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                *string               `json:"Id,omitempty"`
		Name              *string               `json:"Name,omitempty"`
		Description       common.NullableString `json:"Description,omitempty"`
		OtherValues       common.NullableString `json:"OtherValues,omitempty"`
		Data              *json.RawMessage      `json:"Data,omitempty"`
		LookupType        *LookupType           `json:"LookupType,omitempty"`
		Columns           []string              `json:"Columns,omitempty"`
		Count             *int64                `json:"Count,omitempty"`
		LastExecutionTime *string               `json:"LastExecutionTime,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Id", "Name", "Description", "OtherValues", "Data", "LookupType", "Columns", "Count", "LastExecutionTime"})
	} else {
		return err
	}

	o.ID = all.Id
	o.Name = all.Name
	o.Description = all.Description
	o.OtherValues = all.OtherValues
	o.Data = all.Data
	hasInvalidField := false
	if !all.LookupType.IsValid() {
		hasInvalidField = true
	} else {
		o.LookupType = *all.LookupType
	}
	o.Columns = all.Columns
	o.Count = all.Count
	o.LastExecutionTime = all.LastExecutionTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
