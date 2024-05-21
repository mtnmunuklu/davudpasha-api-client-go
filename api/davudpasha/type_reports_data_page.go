package davudpasha

type ReportsPage struct {
	IsA3           *bool  `json:"IsA3,omitempty"`
	IsLandscape    *bool  `json:"IsLandscape,omitempty"`
	TopMargin      *int64 `json:"TopMargin,omitempty"`
	BottomMargin   *int64 `json:"BottomMargin,omitempty"`
	LeftMargin     *int64 `json:"LeftMargin,omitempty"`
	RightMargin    *int64 `json:"RightMargin,omitempty"`
	HeaderDistance *int64 `json:"HeaderDistance,omitempty"`
	FooterDistance *int64 `json:"FooterDistance,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
