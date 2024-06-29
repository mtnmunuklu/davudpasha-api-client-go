package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ActionDefinationsSearchResponse struct {
	Action             *Action               `json:"Action,omitempty"`
	ActionDefinitionId *string               `json:"ActionDefinitionId,omitempty"`
	Application        *string               `json:"Application,omitempty"`
	Author             *string               `json:"Author,omitempty"`
	Description        common.NullableString `json:"Description,omitempty"`
	IsPublic           *bool                 `json:"IsPublic,omitempty"`
	Name               *string               `json:"Name,omitempty"`
	Type               *string               `json:"Type,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewActionDefinationsSearchResponse creates a new ActionDefinationsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionDefinationsSearchResponse() *ActionDefinationsSearchResponse {
	this := ActionDefinationsSearchResponse{}
	return &this
}

// NewActionDefinationsSearchResponseWithDefaults creates a new ActionDefinationsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionDefinationsSearchResponseWithDefaults() *ActionDefinationsSearchResponse {
	this := ActionDefinationsSearchResponse{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ActionDefinationsSearchResponse) GetAction() Action {
	if o == nil || o.Action == nil {
		var ret Action
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchResponse) GetActionOk() (*Action, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ActionDefinationsSearchResponse) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given Action and assigns it to the Action field.
func (o *ActionDefinationsSearchResponse) SetAction(v Action) {
	o.Action = &v
}

// GetActionDefinitionId returns the ActionDefinitionId field value if set, zero value otherwise.
func (o *ActionDefinationsSearchResponse) GetActionDefinitionId() string {
	if o == nil || o.ActionDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.ActionDefinitionId
}

// GetActionDefinitionIdOk returns a tuple with the ActionDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchResponse) GetActionDefinitionIdOk() (*string, bool) {
	if o == nil || o.ActionDefinitionId == nil {
		return nil, false
	}
	return o.ActionDefinitionId, true
}

// HasActionDefinitionId returns a boolean if a field has been set.
func (o *ActionDefinationsSearchResponse) HasActionDefinitionId() bool {
	return o != nil && o.ActionDefinitionId != nil
}

// SetActionDefinitionId gets a reference to the given string and assigns it to the ActionDefinitionId field.
func (o *ActionDefinationsSearchResponse) SetActionDefinitionId(v string) {
	o.ActionDefinitionId = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ActionDefinationsSearchResponse) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchResponse) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ActionDefinationsSearchResponse) HasApplication() bool {
	return o != nil && o.Application != nil
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ActionDefinationsSearchResponse) SetApplication(v string) {
	o.Application = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ActionDefinationsSearchResponse) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchResponse) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ActionDefinationsSearchResponse) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *ActionDefinationsSearchResponse) SetAuthor(v string) {
	o.Author = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionDefinationsSearchResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ActionDefinationsSearchResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a Description has been set.
func (o *ActionDefinationsSearchResponse) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ActionDefinationsSearchResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ActionDefinationsSearchResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnSetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ActionDefinationsSearchResponse) UnSetDescription() {
	o.Description.UnSet()
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *ActionDefinationsSearchResponse) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchResponse) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *ActionDefinationsSearchResponse) HasIsPublic() bool {
	return o != nil && o.IsPublic != nil
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *ActionDefinationsSearchResponse) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActionDefinationsSearchResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActionDefinationsSearchResponse) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActionDefinationsSearchResponse) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActionDefinationsSearchResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinationsSearchResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActionDefinationsSearchResponse) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActionDefinationsSearchResponse) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionDefinationsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.ActionDefinitionId != nil {
		toSerialize["ActionDefinitionId"] = o.ActionDefinitionId
	}
	if o.Application != nil {
		toSerialize["Application"] = o.Application
	}
	if o.Author != nil {
		toSerialize["Author"] = o.Author
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.IsPublic != nil {
		toSerialize["IsPublic"] = o.IsPublic
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ActionDefinationsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action             *Action               `json:"Action,omitempty"`
		ActionDefinitionId *string               `json:"ActionDefinitionId,omitempty"`
		Application        *string               `json:"Application,omitempty"`
		Author             *string               `json:"Author,omitempty"`
		Description        common.NullableString `json:"Description,omitempty"`
		IsPublic           *bool                 `json:"IsPublic,omitempty"`
		Name               *string               `json:"Name,omitempty"`
		Type               *string               `json:"Type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Action", "ActionDefinitionId", "Application", "Author", "Description", "IsPublic", "Name", "Type"})
	} else {
		return err
	}

	o.Action = all.Action
	o.ActionDefinitionId = all.ActionDefinitionId
	o.Application = all.Application
	o.Author = all.Author
	o.Description = all.Description
	o.IsPublic = all.IsPublic
	o.Name = all.Name
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
