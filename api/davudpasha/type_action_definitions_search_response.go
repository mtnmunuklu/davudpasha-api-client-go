package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ActionDefinitionsSearchResponse represents the response structure for searching action definitions.
type ActionDefinitionsSearchResponse struct {
	// Contains the details of the action.
	Action *Action `json:"Action,omitempty"`
	// Unique identifier for the action definition.
	ActionDefinitionID *string `json:"ActionDefinitionId,omitempty"`
	// Specifies the application to which the action definition belongs.
	Application *string `json:"Application,omitempty"`
	// The author of the action definition.
	Author *string `json:"Author,omitempty"`
	// Description of the action definition, which can be null.
	Description common.NullableString `json:"Description,omitempty"`
	// Indicates whether the action definition is public.
	IsPublic *bool `json:"IsPublic,omitempty"`
	// Name of the action definition.
	Name *string `json:"Name,omitempty"`
	// Type of the action definition.
	Type *string `json:"Type,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionDefinitionsSearchResponse creates a new ActionDefinitionsSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionDefinitionsSearchResponse() *ActionDefinitionsSearchResponse {
	this := ActionDefinitionsSearchResponse{}
	return &this
}

// NewActionDefinitionsSearchResponseWithDefaults creates a new ActionDefinitionsSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionDefinitionsSearchResponseWithDefaults() *ActionDefinitionsSearchResponse {
	this := ActionDefinitionsSearchResponse{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ActionDefinitionsSearchResponse) GetAction() Action {
	if o == nil || o.Action == nil {
		var ret Action
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinitionsSearchResponse) GetActionOk() (*Action, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ActionDefinitionsSearchResponse) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given Action and assigns it to the Action field.
func (o *ActionDefinitionsSearchResponse) SetAction(v Action) {
	o.Action = &v
}

// GetActionDefinitionID returns the ActionDefinitionID field value if set, zero value otherwise.
func (o *ActionDefinitionsSearchResponse) GetActionDefinitionID() string {
	if o == nil || o.ActionDefinitionID == nil {
		var ret string
		return ret
	}
	return *o.ActionDefinitionID
}

// GetActionDefinitionIDOk returns a tuple with the ActionDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinitionsSearchResponse) GetActionDefinitionIDOk() (*string, bool) {
	if o == nil || o.ActionDefinitionID == nil {
		return nil, false
	}
	return o.ActionDefinitionID, true
}

// HasActionDefinitionID returns a boolean if a field has been set.
func (o *ActionDefinitionsSearchResponse) HasActionDefinitionID() bool {
	return o != nil && o.ActionDefinitionID != nil
}

// SetActionDefinitionID gets a reference to the given string and assigns it to the ActionDefinitionID field.
func (o *ActionDefinitionsSearchResponse) SetActionDefinitionID(v string) {
	o.ActionDefinitionID = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ActionDefinitionsSearchResponse) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinitionsSearchResponse) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ActionDefinitionsSearchResponse) HasApplication() bool {
	return o != nil && o.Application != nil
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ActionDefinitionsSearchResponse) SetApplication(v string) {
	o.Application = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ActionDefinitionsSearchResponse) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinitionsSearchResponse) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ActionDefinitionsSearchResponse) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *ActionDefinitionsSearchResponse) SetAuthor(v string) {
	o.Author = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionDefinitionsSearchResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ActionDefinitionsSearchResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a Description has been set.
func (o *ActionDefinitionsSearchResponse) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *ActionDefinitionsSearchResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ActionDefinitionsSearchResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnSetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ActionDefinitionsSearchResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *ActionDefinitionsSearchResponse) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinitionsSearchResponse) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *ActionDefinitionsSearchResponse) HasIsPublic() bool {
	return o != nil && o.IsPublic != nil
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *ActionDefinitionsSearchResponse) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActionDefinitionsSearchResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinitionsSearchResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActionDefinitionsSearchResponse) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActionDefinitionsSearchResponse) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActionDefinitionsSearchResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDefinitionsSearchResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActionDefinitionsSearchResponse) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActionDefinitionsSearchResponse) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionDefinitionsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.ActionDefinitionID != nil {
		toSerialize["ActionDefinitionId"] = o.ActionDefinitionID
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
func (o *ActionDefinitionsSearchResponse) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action             *Action               `json:"Action,omitempty"`
		ActionDefinitionID *string               `json:"ActionDefinitionId,omitempty"`
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
	o.ActionDefinitionID = all.ActionDefinitionID
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
