package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// SourceTypesItem represents the structure for items in source types.
type SourceTypesItem struct {
	// Name specifies the name of the item.
	Name *string `json:"Name,omitempty"`
	// DefCode specifies the definition code of the item.
	DefCode *string `json:"DefCode,omitempty"`
	// Author specifies the author of the item.
	Author common.NullableString `json:"Author,omitempty"`
	// Icon specifies the icon associated with the item.
	Icon common.NullableString `json:"Icon,omitempty"`
	// CatCode specifies the category code of the item.
	CatCode *string `json:"CatCode,omitempty"`
	// SourceReaderType specifies the type of source reader used.
	SourceReaderType *string `json:"SourceReaderType,omitempty"`
	// ReleaseDate specifies the release date of the item.
	ReleaseDate *string `json:"ReleaseDate,omitempty"`
	// Version specifies the version number of the item.
	Version *float64 `json:"Version,omitempty"`
	// Expressions contains expressions associated with the item.
	Expressions []SourceTypesExpression `json:"Expression,omitempty"`
	// ExpressionFields contains expression fields associated with the item.
	ExpressionFields []SourceTypesExpressionField `json:"ExpressionFields,omitempty"`
	// DashboardItems contains items associated with the dashboard.
	DashboardItems []string `json:"DashboardItems,omitempty"`
	// ClassificationRules contains classification rules associated with the item.
	ClassificationRules []SourceTypesClassificationRule `json:"ClassificationRules,omitempty"`
	// FromMarket indicates if the item is sourced from the market.
	FromMarket *bool `json:"FromMarket,omitempty"`
	// FromModules indicates if the item is sourced from modules.
	FromModules *bool `json:"FromModules,omitempty"`
	// CopyFromMarket indicates if the item is copied from the market.
	CopyFromMarket *bool `json:"CopyFromMarket,omitempty"`
	// HasUpdate indicates if the item has updates available.
	HasUpdate *bool `json:"HasUpdate,omitempty"`
	// ModuleID specifies the ID of the module associated with the item.
	ModuleID common.NullableString `json:"ModuleId,omitempty"`
	// ModuleGGUID specifies the GUID of the module associated with the item.
	ModuleGUID common.NullableString `json:"ModuleGuid,omitempty"`
	// ClassificationDefs contains classification definitions associated with the item.
	ClassificationDefs []SourceTypesClassificationDefination `json:"ClassificationDefs,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewSourceTypesItem creates a new SourceTypesItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceTypesItem() *SourceTypesItem {
	this := SourceTypesItem{}
	return &this
}

// NewSourceTypesItemWithDefaults creates a new SourceTypesItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourceTypesItemWithDefaults() *SourceTypesItem {
	this := SourceTypesItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceTypesItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceTypesItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceTypesItem) SetName(v string) {
	o.Name = &v
}

// GetDefCode returns the DefCode field value if set, zero value otherwise.
func (o *SourceTypesItem) GetDefCode() string {
	if o == nil || o.DefCode == nil {
		var ret string
		return ret
	}
	return *o.DefCode
}

// GetDefCodeOk returns a tuple with the DefCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetDefCodeOk() (*string, bool) {
	if o == nil || o.DefCode == nil {
		return nil, false
	}
	return o.DefCode, true
}

// HasDefCode returns a boolean if a field has been set.
func (o *SourceTypesItem) HasDefCode() bool {
	return o != nil && o.DefCode != nil
}

// SetDefCode gets a reference to the given string and assigns it to the DefCode field.
func (o *SourceTypesItem) SetDefCode(v string) {
	o.DefCode = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesItem) GetAuthor() string {
	if o == nil || o.Author.Get() == nil {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesItem) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a Author has been set.
func (o *SourceTypesItem) HasAuthor() bool {
	return o != nil && o.Author.IsSet()
}

// SetAuthor gets a reference to the given common.NullableString and assigns it to the Author field.
func (o *SourceTypesItem) SetAuthor(v string) {
	o.Author.Set(&v)
}

// SetAuthorNil sets the value for Author to be an explicit nil.
func (o *SourceTypesItem) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnSetAuthor ensures that no value is present for Author, not even an explicit nil.
func (o *SourceTypesItem) UnSetAuthor() {
	o.Author.UnSet()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesItem) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesItem) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a Icon has been set.
func (o *SourceTypesItem) HasIcon() bool {
	return o != nil && o.Icon.IsSet()
}

// SetIcon gets a reference to the given common.NullableString and assigns it to the Icon field.
func (o *SourceTypesItem) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil.
func (o *SourceTypesItem) SetIconNil() {
	o.Icon.Set(nil)
}

// UnSetIcon ensures that no value is present for Icon, not even an explicit nil.
func (o *SourceTypesItem) UnSetIcon() {
	o.Icon.UnSet()
}

// GetCatCode returns the CatCode field value if set, zero value otherwise.
func (o *SourceTypesItem) GetCatCode() string {
	if o == nil || o.CatCode == nil {
		var ret string
		return ret
	}
	return *o.CatCode
}

// GetCatCodeOk returns a tuple with the CatCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetCatCodeOk() (*string, bool) {
	if o == nil || o.CatCode == nil {
		return nil, false
	}
	return o.CatCode, true
}

// HasCatCode returns a boolean if a field has been set.
func (o *SourceTypesItem) HasCatCode() bool {
	return o != nil && o.CatCode != nil
}

// SetCatCode gets a reference to the given string and assigns it to the CatCode field.
func (o *SourceTypesItem) SetCatCode(v string) {
	o.CatCode = &v
}

// GetSourceReaderType returns the SourceReaderType field value if set, zero value otherwise.
func (o *SourceTypesItem) GetSourceReaderType() string {
	if o == nil || o.SourceReaderType == nil {
		var ret string
		return ret
	}
	return *o.SourceReaderType
}

// GetSourceReaderTypeOk returns a tuple with the SourceReaderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetSourceReaderTypeOk() (*string, bool) {
	if o == nil || o.SourceReaderType == nil {
		return nil, false
	}
	return o.SourceReaderType, true
}

// HasSourceReaderType returns a boolean if a field has been set.
func (o *SourceTypesItem) HasSourceReaderType() bool {
	return o != nil && o.SourceReaderType != nil
}

// SetSourceReaderType gets a reference to the given string and assigns it to the SourceReaderType field.
func (o *SourceTypesItem) SetSourceReaderType(v string) {
	o.SourceReaderType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *SourceTypesItem) GetReleaseDate() string {
	if o == nil || o.ReleaseDate == nil {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetReleaseDateOk() (*string, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *SourceTypesItem) HasReleaseDate() bool {
	return o != nil && o.ReleaseDate != nil
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *SourceTypesItem) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SourceTypesItem) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SourceTypesItem) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *SourceTypesItem) SetVersion(v float64) {
	o.Version = &v
}

// GetExpressions returns the Expressions field value if set, zero value otherwise.
func (o *SourceTypesItem) GetExpressions() []SourceTypesExpression {
	if o == nil || o.Expressions == nil {
		var ret []SourceTypesExpression
		return ret
	}
	return o.Expressions
}

// GetExpressionsOk returns a tuple with the Expressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetExpressionsOk() (*[]SourceTypesExpression, bool) {
	if o == nil || o.Expressions == nil {
		return nil, false
	}
	return &o.Expressions, true
}

// HasExpressions returns a boolean if a field has been set.
func (o *SourceTypesItem) HasExpressions() bool {
	return o != nil && o.Expressions != nil
}

// SetExpressions gets a reference to the given []SourceTypesExpression and assigns it to the Expressions field.
func (o *SourceTypesItem) SetExpressions(v []SourceTypesExpression) {
	o.Expressions = v
}

// GetExpressionFields returns the ExpressionFields field value if set, zero value otherwise.
func (o *SourceTypesItem) GetExpressionFields() []SourceTypesExpressionField {
	if o == nil || o.ExpressionFields == nil {
		var ret []SourceTypesExpressionField
		return ret
	}
	return o.ExpressionFields
}

// GetExpressionFieldsOk returns a tuple with the ExpressionFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetExpressionFieldsOk() (*[]SourceTypesExpressionField, bool) {
	if o == nil || o.ExpressionFields == nil {
		return nil, false
	}
	return &o.ExpressionFields, true
}

// HasExpressionFields returns a boolean if a field has been set.
func (o *SourceTypesItem) HasExpressionFields() bool {
	return o != nil && o.ExpressionFields != nil
}

// SetExpressionFields gets a reference to the given []SourceTypesExpressionField and assigns it to the ExpressionFields field.
func (o *SourceTypesItem) SetExpressionFields(v []SourceTypesExpressionField) {
	o.ExpressionFields = v
}

// GetDashboardItems returns the DashboardItems field value if set, zero value otherwise.
func (o *SourceTypesItem) GetDashboardItems() []string {
	if o == nil || o.DashboardItems == nil {
		var ret []string
		return ret
	}
	return o.DashboardItems
}

// GetDashboardItemsOk returns a tuple with the DashboardItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetDashboardItemsOk() (*[]string, bool) {
	if o == nil || o.DashboardItems == nil {
		return nil, false
	}
	return &o.DashboardItems, true
}

// HasDashboardItems returns a boolean if a field has been set.
func (o *SourceTypesItem) HasDashboardItems() bool {
	return o != nil && o.DashboardItems != nil
}

// SetDashboardItems gets a reference to the given []string and assigns it to the DashboardItems field.
func (o *SourceTypesItem) SetDashboardItems(v []string) {
	o.DashboardItems = v
}

// GetClassificationRules returns the ClassificationRules field value if set, zero value otherwise.
func (o *SourceTypesItem) GetClassificationRules() []SourceTypesClassificationRule {
	if o == nil || o.ClassificationRules == nil {
		var ret []SourceTypesClassificationRule
		return ret
	}
	return o.ClassificationRules
}

// GetClassificationRulesOk returns a tuple with the ClassificationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetClassificationRulesOk() (*[]SourceTypesClassificationRule, bool) {
	if o == nil || o.ClassificationRules == nil {
		return nil, false
	}
	return &o.ClassificationRules, true
}

// HasClassificationRules returns a boolean if a field has been set.
func (o *SourceTypesItem) HasClassificationRules() bool {
	return o != nil && o.ClassificationRules != nil
}

// SetClassificationRules gets a reference to the given []SourceTypesClassificationRule and assigns it to the ClassificationRules field.
func (o *SourceTypesItem) SetClassificationRules(v []SourceTypesClassificationRule) {
	o.ClassificationRules = v
}

// GetFromMarket returns the FromMarket field value if set, zero value otherwise.
func (o *SourceTypesItem) GetFromMarket() bool {
	if o == nil || o.FromMarket == nil {
		var ret bool
		return ret
	}
	return *o.FromMarket
}

// GetFromMarketOk returns a tuple with the FromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetFromMarketOk() (*bool, bool) {
	if o == nil || o.FromMarket == nil {
		return nil, false
	}
	return o.FromMarket, true
}

// HasFromMarket returns a boolean if a field has been set.
func (o *SourceTypesItem) HasFromMarket() bool {
	return o != nil && o.FromMarket != nil
}

// SetFromMarket gets a reference to the given bool and assigns it to the FromMarket field.
func (o *SourceTypesItem) SetFromMarket(v bool) {
	o.FromMarket = &v
}

// GetFromModules returns the FromModules field value if set, zero value otherwise.
func (o *SourceTypesItem) GetFromModules() bool {
	if o == nil || o.FromModules == nil {
		var ret bool
		return ret
	}
	return *o.FromModules
}

// GetFromModulesOk returns a tuple with the FromModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetFromModulesOk() (*bool, bool) {
	if o == nil || o.FromModules == nil {
		return nil, false
	}
	return o.FromModules, true
}

// HasFromModules returns a boolean if a field has been set.
func (o *SourceTypesItem) HasFromModules() bool {
	return o != nil && o.FromModules != nil
}

// SetFromModules gets a reference to the given bool and assigns it to the FromModules field.
func (o *SourceTypesItem) SetFromModules(v bool) {
	o.FromModules = &v
}

// GetCopyFromMarket returns the CopyFromMarket field value if set, zero value otherwise.
func (o *SourceTypesItem) GetCopyFromMarket() bool {
	if o == nil || o.CopyFromMarket == nil {
		var ret bool
		return ret
	}
	return *o.CopyFromMarket
}

// GetCopyFromMarketOk returns a tuple with the CopyFromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetCopyFromMarketOk() (*bool, bool) {
	if o == nil || o.CopyFromMarket == nil {
		return nil, false
	}
	return o.CopyFromMarket, true
}

// HasCopyFromMarket returns a boolean if a field has been set.
func (o *SourceTypesItem) HasCopyFromMarket() bool {
	return o != nil && o.CopyFromMarket != nil
}

// SetCopyFromMarket gets a reference to the given bool and assigns it to the CopyFromMarket field.
func (o *SourceTypesItem) SetCopyFromMarket(v bool) {
	o.CopyFromMarket = &v
}

// GetHasUpdate returns the HasUpdate field value if set, zero value otherwise.
func (o *SourceTypesItem) GetHasUpdate() bool {
	if o == nil || o.HasUpdate == nil {
		var ret bool
		return ret
	}
	return *o.HasUpdate
}

// GetHasUpdateOk returns a tuple with the HasUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetHasUpdateOk() (*bool, bool) {
	if o == nil || o.HasUpdate == nil {
		return nil, false
	}
	return o.HasUpdate, true
}

// HasHasUpdate returns a boolean if a field has been set.
func (o *SourceTypesItem) HasHasUpdate() bool {
	return o != nil && o.HasUpdate != nil
}

// SetHasUpdate gets a reference to the given bool and assigns it to the HasUpdate field.
func (o *SourceTypesItem) SetHasUpdate(v bool) {
	o.HasUpdate = &v
}

// GetModuleID returns the ModuleID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesItem) GetModuleID() string {
	if o == nil || o.ModuleID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleID.Get()
}

// GetModuleIDOk returns a tuple with the ModuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesItem) GetModuleIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleID.Get(), o.ModuleID.IsSet()
}

// HasModuleID returns a boolean if a ModuleID has been set.
func (o *SourceTypesItem) HasModuleID() bool {
	return o != nil && o.ModuleID.IsSet()
}

// SetModuleID gets a reference to the given common.NullableString and assigns it to the ModuleID field.
func (o *SourceTypesItem) SetModuleID(v string) {
	o.ModuleID.Set(&v)
}

// SetModuleIDNil sets the value for ModuleID to be an explicit nil.
func (o *SourceTypesItem) SetModuleIDNil() {
	o.ModuleID.Set(nil)
}

// UnSetModuleID ensures that no value is present for ModuleID, not even an explicit nil.
func (o *SourceTypesItem) UnSetModuleID() {
	o.ModuleID.UnSet()
}

// GetModuleGUID returns the ModuleGUID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceTypesItem) GetModuleGUID() string {
	if o == nil || o.ModuleGUID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleGUID.Get()
}

// GetModuleGUIDOk returns a tuple with the ModuleGUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceTypesItem) GetModuleGUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleGUID.Get(), o.ModuleGUID.IsSet()
}

// HasModuleGUID returns a boolean if a ModuleGUID has been set.
func (o *SourceTypesItem) HasModuleGUID() bool {
	return o != nil && o.ModuleGUID.IsSet()
}

// SetModuleGUID gets a reference to the given common.NullableString and assigns it to the ModuleGUID field.
func (o *SourceTypesItem) SetModuleGUID(v string) {
	o.ModuleGUID.Set(&v)
}

// SetModuleGUIDNil sets the value for ModuleGUID to be an explicit nil.
func (o *SourceTypesItem) SetModuleGUIDNil() {
	o.ModuleGUID.Set(nil)
}

// UnSetModuleGUID ensures that no value is present for ModuleGUID, not even an explicit nil.
func (o *SourceTypesItem) UnSetModuleGUID() {
	o.ModuleGUID.UnSet()
}

// GetClassificationDefs returns the ClassificationDefs field value if set, zero value otherwise.
func (o *SourceTypesItem) GetClassificationDefs() []SourceTypesClassificationDefination {
	if o == nil || o.ClassificationDefs == nil {
		var ret []SourceTypesClassificationDefination
		return ret
	}
	return o.ClassificationDefs
}

// GetClassificationDefsOk returns a tuple with the ClassificationDefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesItem) GetClassificationDefsOk() (*[]SourceTypesClassificationDefination, bool) {
	if o == nil || o.ClassificationDefs == nil {
		return nil, false
	}
	return &o.ClassificationDefs, true
}

// HasClassificationDefs returns a boolean if a field has been set.
func (o *SourceTypesItem) HasClassificationDefs() bool {
	return o != nil && o.ClassificationDefs != nil
}

// SetClassificationDefs gets a reference to the given []SourceTypesClassificationDefination and assigns it to the ClassificationDefs field.
func (o *SourceTypesItem) SetClassificationDefs(v []SourceTypesClassificationDefination) {
	o.ClassificationDefs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourceTypesItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.DefCode != nil {
		toSerialize["DefCode"] = o.DefCode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Author.IsSet() {
		toSerialize["Author"] = o.Author.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["Icon"] = o.Icon.Get()
	}
	if o.CatCode != nil {
		toSerialize["CatCode"] = o.CatCode
	}
	if o.SourceReaderType != nil {
		toSerialize["SourceReaderType"] = o.SourceReaderType
	}
	if o.ReleaseDate != nil {
		toSerialize["ReleaseDate"] = o.ReleaseDate
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Expressions != nil {
		toSerialize["Expressions"] = o.Expressions
	}
	if o.ExpressionFields != nil {
		toSerialize["ExpressionFields"] = o.ExpressionFields
	}
	if o.DashboardItems != nil {
		toSerialize["DashboardItems"] = o.DashboardItems
	}
	if o.ClassificationRules != nil {
		toSerialize["ClassificationRules"] = o.ClassificationRules
	}
	if o.FromMarket != nil {
		toSerialize["FromMarket"] = o.FromMarket
	}
	if o.FromModules != nil {
		toSerialize["FromModules"] = o.FromModules
	}
	if o.CopyFromMarket != nil {
		toSerialize["CopyFromMarket"] = o.CopyFromMarket
	}
	if o.HasUpdate != nil {
		toSerialize["HasUpdate"] = o.HasUpdate
	}
	if o.ModuleID.IsSet() {
		toSerialize["ModuleId"] = o.ModuleID.Get()
	}
	if o.ModuleGUID.IsSet() {
		toSerialize["ModuleGuid"] = o.ModuleGUID.Get()
	}
	if o.ClassificationDefs != nil {
		toSerialize["ClassificationDefs"] = o.ClassificationDefs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *SourceTypesItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                *string                               `json:"Name,omitempty"`
		DefCode             *string                               `json:"DefCode,omitempty"`
		Author              common.NullableString                 `json:"Author,omitempty"`
		Icon                common.NullableString                 `json:"Icon,omitempty"`
		CatCode             *string                               `json:"CatCode,omitempty"`
		SourceReaderType    *string                               `json:"SourceReaderType,omitempty"`
		ReleaseDate         *string                               `json:"ReleaseDate,omitempty"`
		Version             *float64                              `json:"Version,omitempty"`
		Expressions         []SourceTypesExpression               `json:"Expression,omitempty"`
		ExpressionFields    []SourceTypesExpressionField          `json:"ExpressionFields,omitempty"`
		DashboardItems      []string                              `json:"DashboardItems,omitempty"`
		ClassificationRules []SourceTypesClassificationRule       `json:"ClassificationRules,omitempty"`
		FromMarket          *bool                                 `json:"FromMarket,omitempty"`
		FromModules         *bool                                 `json:"FromModules,omitempty"`
		CopyFromMarket      *bool                                 `json:"CopyFromMarket,omitempty"`
		HasUpdate           *bool                                 `json:"HasUpdate,omitempty"`
		ModuleID            common.NullableString                 `json:"ModuleId,omitempty"`
		ModuleGUID          common.NullableString                 `json:"ModuleGuid,omitempty"`
		ClassificationDefs  []SourceTypesClassificationDefination `json:"ClassificationDefs,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}
	if all.DefCode == nil {
		return fmt.Errorf("requiered field DefCode is missing")
	}
	if all.CatCode == nil {
		return fmt.Errorf("requiered field CatCode is missing")
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Name", "DefCode", "Author", "Icon", "CatCode", "SourceReaderType", "ReleaseDate", "Version", "Expressions", "ExpressionFields", "DashboardItems", "ClassificationRules", "FromMarket", "FromModules", "CopyFromMarket", "HasUpdate", "ModuleId", "ModuleGuid", "ClassificationDefs"})
	} else {
		return err
	}

	o.Name = all.Name
	o.DefCode = all.DefCode
	o.Author = all.Author
	o.Icon = all.Icon
	o.CatCode = all.CatCode
	o.SourceReaderType = all.SourceReaderType
	o.ReleaseDate = all.ReleaseDate
	o.Version = all.Version
	o.Expressions = all.Expressions
	o.ExpressionFields = all.ExpressionFields
	o.DashboardItems = all.DashboardItems
	o.ClassificationRules = all.ClassificationRules
	o.FromMarket = all.FromMarket
	o.FromModules = all.FromModules
	o.CopyFromMarket = all.CopyFromMarket
	o.HasUpdate = all.HasUpdate
	o.ModuleID = all.ModuleID
	o.ModuleGUID = all.ModuleGUID
	o.ClassificationDefs = all.ClassificationDefs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
