package davudpasha

import (
	"encoding/json"
	"fmt"
)

// TagCategory represents the category of tag.
type TagCategory string

// Allowed values for type TagCategory string.
const (
	TAGCATEGORY_ALL    TagCategory = "All"
	TAGCATEGORY_QUERY  TagCategory = "Query"
	TAGCATEGORY_ALERT  TagCategory = "Alert"
	TAGCATEGORY_REPORT TagCategory = "Report"
)

var allowedTagCategoryEnumValues = []TagCategory{
	TAGCATEGORY_ALL,
	TAGCATEGORY_QUERY,
	TAGCATEGORY_ALERT,
	TAGCATEGORY_REPORT,
}

// GetAllowedValues returns the list of possible values.
func (v *TagCategory) GetAllowedValues() []TagCategory {
	return allowedTagCategoryEnumValues
}

// UnMarshalJSON deserializes the given payload.
func (v *TagCategory) UnMarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagCategory(value)
	return nil
}

// NewTagCategoryFromValue returns a pointer to a valid TagCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagCategoryFromValue(v string) (*TagCategory, error) {
	ev := TagCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagCategory: valid values are %v", v, allowedTagCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagCategory) IsValid() bool {
	for _, existing := range allowedTagCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns references to TagCategory value.
func (v TagCategory) Ptr() *TagCategory {
	return &v
}
