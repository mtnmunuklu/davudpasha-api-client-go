package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// ReportsPage defines the layout settings for a page in reports.
type ReportsPage struct {
	// Indicates if the page size is A3.
	IsA3 *bool `json:"IsA3,omitempty"`
	// Indicates if the page orientation is landscape.
	IsLandscape *bool `json:"IsLandscape,omitempty"`
	// Top margin of the page.
	TopMargin *int64 `json:"TopMargin,omitempty"`
	// Bottom margin of the page.
	BottomMargin *int64 `json:"BottomMargin,omitempty"`
	// Left margin of the page.
	LeftMargin *int64 `json:"LeftMargin,omitempty"`
	// Right margin of the page.
	RightMargin *int64 `json:"RightMargin,omitempty"`
	// Distance of the header from the top of the page.
	HeaderDistance *int64 `json:"HeaderDistance,omitempty"`
	// Distance of the footer from the bottom of the page.
	FooterDistance *int64 `json:"FooterDistance,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewReportsPage creates a new ReportsPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsPage() *ReportsPage {
	this := ReportsPage{}
	return &this
}

// NewReportsPageWithDefaults creates a new ReportsPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsPageWithDefaults() *ReportsPage {
	this := ReportsPage{}
	return &this
}

// GetIsA3 returns the IsA3 field value if set, zero value otherwise.
func (o *ReportsPage) GetIsA3() bool {
	if o == nil || o.IsA3 == nil {
		var ret bool
		return ret
	}
	return *o.IsA3
}

// GetIsA3Ok returns a tuple with the IsA3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPage) GetIsA3Ok() (*bool, bool) {
	if o == nil || o.IsA3 == nil {
		return nil, false
	}
	return o.IsA3, true
}

// HasIsA3 returns a boolean if a field has been set.
func (o *ReportsPage) HasIsA3() bool {
	return o != nil && o.IsA3 != nil
}

// SetIsA3 gets a reference to the given bool and assigns it to the IsA3 field.
func (o *ReportsPage) SetIsA3(v bool) {
	o.IsA3 = &v
}

// GetIsLandscape returns the IsLandscape field value if set, zero value otherwise.
func (o *ReportsPage) GetIsLandscape() bool {
	if o == nil || o.IsLandscape == nil {
		var ret bool
		return ret
	}
	return *o.IsLandscape
}

// GetIsLandscapeOk returns a tuple with the IsLandscape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPage) GetIsLandscapeOk() (*bool, bool) {
	if o == nil || o.IsLandscape == nil {
		return nil, false
	}
	return o.IsLandscape, true
}

// HasIsLandscape returns a boolean if a field has been set.
func (o *ReportsPage) HasIsLandscape() bool {
	return o != nil && o.IsLandscape != nil
}

// SetIsLandscape gets a reference to the given bool and assigns it to the IsLandscape field.
func (o *ReportsPage) SetIsLandscape(v bool) {
	o.IsLandscape = &v
}

// GetTopMargin returns the TopMargin field value if set, zero value otherwise.
func (o *ReportsPage) GetTopMargin() int64 {
	if o == nil || o.TopMargin == nil {
		var ret int64
		return ret
	}
	return *o.TopMargin
}

// GetTopMarginOk returns a tuple with the TopMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPage) GetTopMarginOk() (*int64, bool) {
	if o == nil || o.TopMargin == nil {
		return nil, false
	}
	return o.TopMargin, true
}

// HasTopMargin returns a boolean if a field has been set.
func (o *ReportsPage) HasTopMargin() bool {
	return o != nil && o.TopMargin != nil
}

// SetTopMargin gets a reference to the given int64 and assigns it to the TopMargin field.
func (o *ReportsPage) SetTopMargin(v int64) {
	o.TopMargin = &v
}

// GetBottomMargin returns the BottomMargin field value if set, zero value otherwise.
func (o *ReportsPage) GetBottomMargin() int64 {
	if o == nil || o.BottomMargin == nil {
		var ret int64
		return ret
	}
	return *o.BottomMargin
}

// GetBottomMarginOk returns a tuple with the BottomMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPage) GetBottomMarginOk() (*int64, bool) {
	if o == nil || o.BottomMargin == nil {
		return nil, false
	}
	return o.BottomMargin, true
}

// HasBottomMargin returns a boolean if a field has been set.
func (o *ReportsPage) HasBottomMargin() bool {
	return o != nil && o.BottomMargin != nil
}

// SetBottomMargin gets a reference to the given int64 and assigns it to the BottomMargin field.
func (o *ReportsPage) SetBottomMargin(v int64) {
	o.BottomMargin = &v
}

// GetLeftMargin returns the LeftMargin field value if set, zero value otherwise.
func (o *ReportsPage) GetLeftMargin() int64 {
	if o == nil || o.LeftMargin == nil {
		var ret int64
		return ret
	}
	return *o.LeftMargin
}

// GetLeftMarginOk returns a tuple with the LeftMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPage) GetLeftMarginOk() (*int64, bool) {
	if o == nil || o.LeftMargin == nil {
		return nil, false
	}
	return o.LeftMargin, true
}

// HasLeftMargin returns a boolean if a field has been set.
func (o *ReportsPage) HasLeftMargin() bool {
	return o != nil && o.LeftMargin != nil
}

// SetLeftMargin gets a reference to the given int64 and assigns it to the LeftMargin field.
func (o *ReportsPage) SetLeftMargin(v int64) {
	o.LeftMargin = &v
}

// GetRightMargin returns the RightMargin field value if set, zero value otherwise.
func (o *ReportsPage) GetRightMargin() int64 {
	if o == nil || o.RightMargin == nil {
		var ret int64
		return ret
	}
	return *o.RightMargin
}

// GetRightMarginOk returns a tuple with the RightMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPage) GetRightMarginOk() (*int64, bool) {
	if o == nil || o.RightMargin == nil {
		return nil, false
	}
	return o.RightMargin, true
}

// HasRightMargin returns a boolean if a field has been set.
func (o *ReportsPage) HasRightMargin() bool {
	return o != nil && o.RightMargin != nil
}

// SetRightMargin gets a reference to the given int64 and assigns it to the RightMargin field.
func (o *ReportsPage) SetRightMargin(v int64) {
	o.RightMargin = &v
}

// GetHeaderDistance returns the HeaderDistance field value if set, zero value otherwise.
func (o *ReportsPage) GetHeaderDistance() int64 {
	if o == nil || o.HeaderDistance == nil {
		var ret int64
		return ret
	}
	return *o.HeaderDistance
}

// GetHeaderDistanceOk returns a tuple with the HeaderDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPage) GetHeaderDistanceOk() (*int64, bool) {
	if o == nil || o.HeaderDistance == nil {
		return nil, false
	}
	return o.HeaderDistance, true
}

// HasHeaderDistance returns a boolean if a field has been set.
func (o *ReportsPage) HasHeaderDistance() bool {
	return o != nil && o.HeaderDistance != nil
}

// SetHeaderDistance gets a reference to the given int64 and assigns it to the HeaderDistance field.
func (o *ReportsPage) SetHeaderDistance(v int64) {
	o.HeaderDistance = &v
}

// GetFooterDistance returns the FooterDistance field value if set, zero value otherwise.
func (o *ReportsPage) GetFooterDistance() int64 {
	if o == nil || o.FooterDistance == nil {
		var ret int64
		return ret
	}
	return *o.FooterDistance
}

// GetFooterDistanceOk returns a tuple with the FooterDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPage) GetFooterDistanceOk() (*int64, bool) {
	if o == nil || o.FooterDistance == nil {
		return nil, false
	}
	return o.FooterDistance, true
}

// HasFooterDistance returns a boolean if a field has been set.
func (o *ReportsPage) HasFooterDistance() bool {
	return o != nil && o.FooterDistance != nil
}

// SetFooterDistance gets a reference to the given int64 and assigns it to the FooterDistance field.
func (o *ReportsPage) SetFooterDistance(v int64) {
	o.FooterDistance = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.IsA3 != nil {
		toSerialize["IsA3"] = o.IsA3
	}
	if o.IsLandscape != nil {
		toSerialize["IsLandscape"] = o.IsLandscape
	}
	if o.TopMargin != nil {
		toSerialize["TopMargin"] = o.TopMargin
	}
	if o.BottomMargin != nil {
		toSerialize["BottomMargin"] = o.BottomMargin
	}
	if o.LeftMargin != nil {
		toSerialize["LeftMargin"] = o.LeftMargin
	}
	if o.RightMargin != nil {
		toSerialize["RightMargin"] = o.RightMargin
	}
	if o.HeaderDistance != nil {
		toSerialize["HeaderDistance"] = o.HeaderDistance
	}
	if o.FooterDistance != nil {
		toSerialize["FooterDistance"] = o.FooterDistance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsPage) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsA3           *bool  `json:"IsA3,omitempty"`
		IsLandscape    *bool  `json:"IsLandscape,omitempty"`
		TopMargin      *int64 `json:"TopMargin,omitempty"`
		BottomMargin   *int64 `json:"BottomMargin,omitempty"`
		LeftMargin     *int64 `json:"LeftMargin,omitempty"`
		RightMargin    *int64 `json:"RightMargin,omitempty"`
		HeaderDistance *int64 `json:"HeaderDistance,omitempty"`
		FooterDistance *int64 `json:"FooterDistance,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"IsA3", "IsLandscape", "TopMargin", "BottomMargin", "LeftMargin", "RightMargin", "HeaderDistance", "FooterDistance"})
	} else {
		return err
	}

	o.IsA3 = all.IsA3
	o.IsLandscape = all.IsLandscape
	o.TopMargin = all.TopMargin
	o.BottomMargin = all.BottomMargin
	o.LeftMargin = all.LeftMargin
	o.RightMargin = all.RightMargin
	o.HeaderDistance = all.HeaderDistance
	o.FooterDistance = all.FooterDistance

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
