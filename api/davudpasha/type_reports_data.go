package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsData contains the data and settings for generating reports.
type ReportsData struct {
	// Name of the report.
	Name common.NullableString `json:"Name,omitempty"`
	// Filename of the report.
	FileName common.NullableString `json:"FileName,omitempty"`
	// Creation date of the report.
	CreateDate *string `json:"CreateDate,omitempty"`
	// Page settings of the report.
	Page *ReportsPage `json:"Page,omitempty"`
	// Header content of the report.
	Header common.NullableString `json:"Header,omitempty"`
	// Footer content of the report.
	Footer common.NullableString `json:"Footer,omitempty"`
	// Cover page content of the report.
	CoverPage common.NullableString `json:"CoverPage,omitempty"`
	// Sections included in the report.
	Sections []string `json:"Sections,omitempty"`
	// Type of the report. ReportType: csv
	ReportType *string `json:"ReportType,omitempty"`
	// Theme of the report.
	Theme common.NullableString `json:"Theme,omitempty"`
	// Language of the report.
	Language *string `json:"Language,omitempty"`
	// Indicates whether to add a cover page to the report.
	AddCoverPage *bool `json:"AddCoverPage,omitempty"`
	// Password for the report file.
	FilePassword common.NullableString `json:"FilePassword,omitempty"`
	// ID of the report.
	ReportID common.NullableString `json:"ReportId,omitempty"`
	// Theme of the report.
	ReportTheme common.NullableString `json:"ReportTheme,omitempty"`
	// Maximum row count for the report data.
	MaxRowCount *int64 `json:"MaxRowCount,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportsData creates a new ReportsData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsData() *ReportsData {
	this := ReportsData{}
	return &this
}

// NewReportsDataWithDefaults creates a new ReportsData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsDataWithDefaults() *ReportsData {
	this := ReportsData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReportsData) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given common.NullableString and assigns it to the Name field.
func (o *ReportsData) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ReportsData) SetNameNil() {
	o.Name.Set(nil)
}

// UnSetName ensures that no value is present for Name, not even an explicit nil.
func (o *ReportsData) UnsetName() {
	o.Name.Unset()
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetFileName() string {
	if o == nil || o.FileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *ReportsData) HasFileName() bool {
	return o != nil && o.FileName.IsSet()
}

// SetFileName gets a reference to the given common.NullableString and assigns it to the FileName field.
func (o *ReportsData) SetFileName(v string) {
	o.FileName.Set(&v)
}

// SetFileNameNil sets the value for FileName to be an explicit nil.
func (o *ReportsData) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnSetFileName ensures that no value is present for FileName, not even an explicit nil.
func (o *ReportsData) UnsetFileName() {
	o.FileName.Unset()
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *ReportsData) GetCreateDate() string {
	if o == nil || o.CreateDate == nil {
		var ret string
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsData) GetCreateDateOk() (*string, bool) {
	if o == nil || o.CreateDate == nil {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *ReportsData) HasCreateDate() bool {
	return o != nil && o.CreateDate != nil
}

// SetCreateDate gets a reference to the given string and assigns it to the CreateDate field.
func (o *ReportsData) SetCreateDate(v string) {
	o.CreateDate = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ReportsData) GetPage() ReportsPage {
	if o == nil || o.Page == nil {
		var ret ReportsPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsData) GetPageOk() (*ReportsPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ReportsData) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given ReportsPage and assigns it to the Page field.
func (o *ReportsData) SetPage(v ReportsPage) {
	o.Page = &v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetHeader() string {
	if o == nil || o.Header.Get() == nil {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *ReportsData) HasHeader() bool {
	return o != nil && o.Header.IsSet()
}

// SetHeader gets a reference to the given common.NullableString and assigns it to the Header field.
func (o *ReportsData) SetHeader(v string) {
	o.Header.Set(&v)
}

// SetHeaderNil sets the value for Header to be an explicit nil.
func (o *ReportsData) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnSetHeader ensures that no value is present for Header, not even an explicit nil.
func (o *ReportsData) UnsetHeader() {
	o.Header.Unset()
}

// GetFooter returns the Footer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetFooter() string {
	if o == nil || o.Footer.Get() == nil {
		var ret string
		return ret
	}
	return *o.Footer.Get()
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetFooterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Footer.Get(), o.Footer.IsSet()
}

// HasFooter returns a boolean if a field has been set.
func (o *ReportsData) HasFooter() bool {
	return o != nil && o.Footer.IsSet()
}

// SetFooter gets a reference to the given common.NullableString and assigns it to the Footer field.
func (o *ReportsData) SetFooter(v string) {
	o.Footer.Set(&v)
}

// SetFooterNil sets the value for Footer to be an explicit nil.
func (o *ReportsData) SetFooterNil() {
	o.Footer.Set(nil)
}

// UnSetFooter ensures that no value is present for Footer, not even an explicit nil.
func (o *ReportsData) UnsetFooter() {
	o.Footer.Unset()
}

// GetCoverPage returns the CoverPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetCoverPage() string {
	if o == nil || o.CoverPage.Get() == nil {
		var ret string
		return ret
	}
	return *o.CoverPage.Get()
}

// GetCoverPageOk returns a tuple with the CoverPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetCoverPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverPage.Get(), o.CoverPage.IsSet()
}

// HasCoverPage returns a boolean if a field has been set.
func (o *ReportsData) HasCoverPage() bool {
	return o != nil && o.CoverPage.IsSet()
}

// SetCoverPage gets a reference to the given common.NullableString and assigns it to the CoverPage field.
func (o *ReportsData) SetCoverPage(v string) {
	o.CoverPage.Set(&v)
}

// SetCoverPageNil sets the value for CoverPage to be an explicit nil.
func (o *ReportsData) SetCoverPageNil() {
	o.CoverPage.Set(nil)
}

// UnSetCoverPage ensures that no value is present for CoverPage, not even an explicit nil.
func (o *ReportsData) UnsetCoverPage() {
	o.CoverPage.Unset()
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *ReportsData) GetSections() []string {
	if o == nil || o.Sections == nil {
		var ret []string
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsData) GetSectionsOk() (*[]string, bool) {
	if o == nil || o.Sections == nil {
		return nil, false
	}
	return &o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *ReportsData) HasSections() bool {
	return o != nil && o.Sections != nil
}

// SetSections gets a reference to the given []string and assigns it to the Sections field.
func (o *ReportsData) SetSections(v []string) {
	o.Sections = v
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *ReportsData) GetReportType() string {
	if o == nil || o.ReportType == nil {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsData) GetReportTypeOk() (*string, bool) {
	if o == nil || o.ReportType == nil {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *ReportsData) HasReportType() bool {
	return o != nil && o.ReportType != nil
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *ReportsData) SetReportType(v string) {
	o.ReportType = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetTheme() string {
	if o == nil || o.Theme.Get() == nil {
		var ret string
		return ret
	}
	return *o.Theme.Get()
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetThemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Theme.Get(), o.Theme.IsSet()
}

// HasTheme returns a boolean if a field has been set.
func (o *ReportsData) HasTheme() bool {
	return o != nil && o.Theme.IsSet()
}

// SetTheme gets a reference to the given common.NullableString and assigns it to the Theme field.
func (o *ReportsData) SetTheme(v string) {
	o.Theme.Set(&v)
}

// SetThemeNil sets the value for Theme to be an explicit nil.
func (o *ReportsData) SetThemeNil() {
	o.Theme.Set(nil)
}

// UnSetTheme ensures that no value is present for Theme, not even an explicit nil.
func (o *ReportsData) UnsetTheme() {
	o.Theme.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ReportsData) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsData) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ReportsData) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ReportsData) SetLanguage(v string) {
	o.Language = &v
}

// GetAddCoverPage returns the AddCoverPage field value if set, zero value otherwise.
func (o *ReportsData) GetAddCoverPage() bool {
	if o == nil || o.AddCoverPage == nil {
		var ret bool
		return ret
	}
	return *o.AddCoverPage
}

// GetAddCoverPageOk returns a tuple with the AddCoverPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsData) GetAddCoverPageOk() (*bool, bool) {
	if o == nil || o.AddCoverPage == nil {
		return nil, false
	}
	return o.AddCoverPage, true
}

// HasAddCoverPage returns a boolean if a field has been set.
func (o *ReportsData) HasAddCoverPage() bool {
	return o != nil && o.AddCoverPage != nil
}

// SetAddCoverPage gets a reference to the given bool and assigns it to the AddCoverPage field.
func (o *ReportsData) SetAddCoverPage(v bool) {
	o.AddCoverPage = &v
}

// GetFilePassword returns the FilePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetFilePassword() string {
	if o == nil || o.FilePassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilePassword.Get()
}

// GetFilePasswordOk returns a tuple with the FilePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetFilePasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilePassword.Get(), o.FilePassword.IsSet()
}

// HasFilePassword returns a boolean if a field has been set.
func (o *ReportsData) HasFilePassword() bool {
	return o != nil && o.FilePassword.IsSet()
}

// SetFilePassword gets a reference to the given common.NullableString and assigns it to the FilePassword field.
func (o *ReportsData) SetFilePassword(v string) {
	o.FilePassword.Set(&v)
}

// SetFilePasswordNil sets the value for FilePassword to be an explicit nil.
func (o *ReportsData) SetFilePasswordNil() {
	o.FilePassword.Set(nil)
}

// UnSetFilePassword ensures that no value is present for FilePassword, not even an explicit nil.
func (o *ReportsData) UnsetFilePassword() {
	o.FilePassword.Unset()
}

// GetReportID returns the ReportID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetReportID() string {
	if o == nil || o.ReportID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReportID.Get()
}

// GetReportIDOk returns a tuple with the ReportID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetReportIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportID.Get(), o.ReportID.IsSet()
}

// HasReportID returns a boolean if a field has been set.
func (o *ReportsData) HasReportID() bool {
	return o != nil && o.ReportID.IsSet()
}

// SetReportID gets a reference to the given common.NullableString and assigns it to the ReportID field.
func (o *ReportsData) SetReportID(v string) {
	o.ReportID.Set(&v)
}

// SetReportIDNil sets the value for ReportID to be an explicit nil.
func (o *ReportsData) SetReportIDNil() {
	o.ReportID.Set(nil)
}

// UnSetReportID ensures that no value is present for ReportID, not even an explicit nil.
func (o *ReportsData) UnsetReportID() {
	o.ReportID.Unset()
}

// GetReportTheme returns the ReportTheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsData) GetReportTheme() string {
	if o == nil || o.ReportTheme.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReportTheme.Get()
}

// GetReportThemeOk returns a tuple with the ReportTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportsData) GetReportThemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportTheme.Get(), o.ReportTheme.IsSet()
}

// HasReportTheme returns a boolean if a field has been set.
func (o *ReportsData) HasReportTheme() bool {
	return o != nil && o.ReportTheme.IsSet()
}

// SetReportTheme gets a reference to the given common.NullableString and assigns it to the ReportTheme field.
func (o *ReportsData) SetReportTheme(v string) {
	o.ReportTheme.Set(&v)
}

// SetReportThemeNil sets the value for ReportTheme to be an explicit nil.
func (o *ReportsData) SetReportThemeNil() {
	o.ReportTheme.Set(nil)
}

// UnSetReportTheme ensures that no value is present for ReportTheme, not even an explicit nil.
func (o *ReportsData) UnsetReportTheme() {
	o.ReportTheme.Unset()
}

// GetMaxRowCount returns the MaxRowCount field value if set, zero value otherwise.
func (o *ReportsData) GetMaxRowCount() int64 {
	if o == nil || o.MaxRowCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxRowCount
}

// GetMaxRowCountOk returns a tuple with the MaxRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsData) GetMaxRowCountOk() (*int64, bool) {
	if o == nil || o.MaxRowCount == nil {
		return nil, false
	}
	return o.MaxRowCount, true
}

// HasMaxRowCount returns a boolean if a field has been set.
func (o *ReportsData) HasMaxRowCount() bool {
	return o != nil && o.MaxRowCount != nil
}

// SetMaxRowCount gets a reference to the given int64 and assigns it to the MaxRowCount field.
func (o *ReportsData) SetMaxRowCount(v int64) {
	o.MaxRowCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Name.IsSet() {
		toSerialize["Label"] = o.Name.Get()
	}
	if o.FileName.IsSet() {
		toSerialize["FileName"] = o.FileName.Get()
	}
	if o.CreateDate != nil {
		toSerialize["CreateDate"] = o.CreateDate
	}
	if o.Page != nil {
		toSerialize["Page"] = o.Page
	}
	if o.Header.IsSet() {
		toSerialize["Header"] = o.Header.Get()
	}
	if o.Footer.IsSet() {
		toSerialize["Footer"] = o.Footer.Get()
	}
	if o.CoverPage.IsSet() {
		toSerialize["CoverPage"] = o.CoverPage.Get()
	}
	if o.Sections != nil {
		toSerialize["Sections"] = o.Sections
	}
	if o.ReportType != nil {
		toSerialize["ReportType"] = o.ReportType
	}
	if o.Theme.IsSet() {
		toSerialize["Theme"] = o.Theme.Get()
	}
	if o.Language != nil {
		toSerialize["Language"] = o.Language
	}
	if o.AddCoverPage != nil {
		toSerialize["AddCoverPage"] = o.AddCoverPage
	}
	if o.FilePassword.IsSet() {
		toSerialize["FilePassword"] = o.FilePassword.Get()
	}
	if o.ReportID.IsSet() {
		toSerialize["ReportId"] = o.ReportID.Get()
	}
	if o.ReportTheme.IsSet() {
		toSerialize["ReportTheme"] = o.ReportTheme.Get()
	}
	if o.MaxRowCount != nil {
		toSerialize["MaxRowCount"] = o.MaxRowCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         common.NullableString `json:"Name,omitempty"`
		FileName     common.NullableString `json:"FileName,omitempty"`
		CreateDate   *string               `json:"CreateDate,omitempty"`
		Page         *ReportsPage          `json:"Page,omitempty"`
		Header       common.NullableString `json:"Header,omitempty"`
		Footer       common.NullableString `json:"Footer,omitempty"`
		CoverPage    common.NullableString `json:"CoverPage,omitempty"`
		Sections     []string              `json:"Sections,omitempty"`
		ReportType   *string               `json:"ReportType,omitempty"`
		Theme        common.NullableString `json:"Theme,omitempty"`
		Language     *string               `json:"Language,omitempty"`
		AddCoverPage *bool                 `json:"AddCoverPage,omitempty"`
		FilePassword common.NullableString `json:"FilePassword,omitempty"`
		ReportID     common.NullableString `json:"ReportId,omitempty"`
		ReportTheme  common.NullableString `json:"ReportTheme,omitempty"`
		MaxRowCount  *int64                `json:"MaxRowCount,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Name", "FileName", "CreateDate", "Page", "Header", "Footer", "CoverPage", "Sections", "ReportType", "Theme", "Language", "AddCoverPage", "FilePassword", "ReportId", "ReportTheme", "MaxRowCount"})
	} else {
		return err
	}

	o.Name = all.Name
	o.FileName = all.FileName
	o.CreateDate = all.CreateDate
	hasInvalidField := false
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	o.Header = all.Header
	o.Footer = all.Footer
	o.CoverPage = all.CoverPage
	o.Sections = all.Sections
	o.ReportType = all.ReportType
	o.Theme = all.Theme
	o.Language = all.Language
	o.AddCoverPage = all.AddCoverPage
	o.FilePassword = all.FilePassword
	o.ReportID = all.ReportID
	o.ReportTheme = all.ReportTheme
	o.MaxRowCount = all.MaxRowCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
