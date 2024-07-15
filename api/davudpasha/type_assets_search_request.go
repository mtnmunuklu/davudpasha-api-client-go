package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// AssetsSearchRequest represents the request structure for searching assets.
type AssetsSearchRequest struct {
	// Primary filter for searching assets.
	SearchFilter *string `json:"searchFilter,omitempty"`
	// Secondary filter for searching assets.
	SearchFilter2 *string `json:"searchFilter2,omitempty"`
	// Page number for pagination.
	PageNumber *int64 `json:"pageNumber,omitempty"`
	// Number of items per page for pagination.
	PageSize *int64 `json:"pageSize,omitempty"`
	// Context for the Smart REST request.
	SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}

// NewAssetsSearchRequest creates a new AssetsSearchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssetsSearchRequest() *AssetsSearchRequest {
	this := AssetsSearchRequest{}
	return &this
}

// NewAssetsSearchRequestWithDefaults creates a new AssetsSearchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssetsSearchRequestWithDefaults() *AssetsSearchRequest {
	this := AssetsSearchRequest{}
	return &this
}

// GetSearchFilter returns the SearchFilter field value if set, zero value otherwise.
func (o *AssetsSearchRequest) GetSearchFilter() string {
	if o == nil || o.SearchFilter == nil {
		var ret string
		return ret
	}
	return *o.SearchFilter
}

// GetSearchFilterOk returns a tuple with the SearchFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsSearchRequest) GetSearchFilterOk() (*string, bool) {
	if o == nil || o.SearchFilter == nil {
		return nil, false
	}
	return o.SearchFilter, true
}

// HasSearchFilter returns a boolean if a field has been set.
func (o *AssetsSearchRequest) HasSearchFilter() bool {
	return o != nil && o.SearchFilter != nil
}

// SetSearchFilter gets a reference to the given string and assigns it to the SearchFilter field.
func (o *AssetsSearchRequest) SetSearchFilter(v string) {
	o.SearchFilter = &v
}

// GetSearchFilter2 returns the SearchFilter2 field value if set, zero value otherwise.
func (o *AssetsSearchRequest) GetSearchFilter2() string {
	if o == nil || o.SearchFilter2 == nil {
		var ret string
		return ret
	}
	return *o.SearchFilter2
}

// GetSearchFilter2Ok returns a tuple with the SearchFilter2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsSearchRequest) GetSearchFilter2Ok() (*string, bool) {
	if o == nil || o.SearchFilter2 == nil {
		return nil, false
	}
	return o.SearchFilter2, true
}

// HasSearchFilter2 returns a boolean if a field has been set.
func (o *AssetsSearchRequest) HasSearchFilter2() bool {
	return o != nil && o.SearchFilter2 != nil
}

// SetSearchFilter2 gets a reference to the given string and assigns it to the SearchFilter2 field.
func (o *AssetsSearchRequest) SetSearchFilter2(v string) {
	o.SearchFilter2 = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *AssetsSearchRequest) GetPageNumber() int64 {
	if o == nil || o.PageNumber == nil {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsSearchRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil || o.PageNumber == nil {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *AssetsSearchRequest) HasPageNumber() bool {
	return o != nil && o.PageNumber != nil
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *AssetsSearchRequest) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AssetsSearchRequest) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsSearchRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AssetsSearchRequest) HasPageSize() bool {
	return o != nil && o.PageSize != nil
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *AssetsSearchRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetSmartRestRequestContext returns the SmartRestRequestContext field value if set, zero value otherwise.
func (o *AssetsSearchRequest) GetSmartRestRequestContext() string {
	if o == nil || o.SmartRestRequestContext == nil {
		var ret string
		return ret
	}
	return *o.SmartRestRequestContext
}

// GetSmartRestRequestContextOk returns a tuple with the SmartRestRequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsSearchRequest) GetSmartRestRequestContextOk() (*string, bool) {
	if o == nil || o.SmartRestRequestContext == nil {
		return nil, false
	}
	return o.SmartRestRequestContext, true
}

// HasSmartRestRequestContext returns a boolean if a field has been set.
func (o *AssetsSearchRequest) HasSmartRestRequestContext() bool {
	return o != nil && o.SmartRestRequestContext != nil
}

// SetSmartRestRequestContext gets a reference to the given string and assigns it to the SmartRestRequestContext field.
func (o *AssetsSearchRequest) SetSmartRestRequestContext(v string) {
	o.SmartRestRequestContext = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssetsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}

	if o.SearchFilter != nil {
		toSerialize["searchFilter"] = o.SearchFilter
	}
	if o.SearchFilter2 != nil {
		toSerialize["searchFilter2"] = o.SearchFilter2
	}
	if o.SmartRestRequestContext != nil {
		toSerialize["smartRestRequestContext"] = o.SmartRestRequestContext
	}
	if o.PageNumber != nil {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *AssetsSearchRequest) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		SearchFilter  *string `json:"searchFilter,omitempty"`
		SearchFilter2 *string `json:"searchFilter2,omitempty"`
		PageNumber    *int64  `json:"pageNumber,omitempty"`
		PageSize      *int64  `json:"pageSize,omitempty"`
		// Context for the Smart REST request.
		SmartRestRequestContext *string `json:"smartRestRequestContext,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"searchFilter", "searchFilter2", "pageNumber", "pageSize", "smartRestRequestContext"})
	} else {
		return err
	}

	if all.SearchFilter2 == nil {
		return fmt.Errorf("requiered field searchFilter2 is missing")
	}
	if all.SmartRestRequestContext == nil {
		return fmt.Errorf("requiered field smartRestRequestContext is missing")
	}

	o.SearchFilter = all.SearchFilter
	o.SearchFilter2 = all.SearchFilter2
	o.PageNumber = all.PageNumber
	o.PageSize = all.PageSize
	o.SmartRestRequestContext = all.SmartRestRequestContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
