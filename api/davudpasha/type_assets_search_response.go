package davudpasha

type AssetsSearchResponse struct {
	Assets []AssetsAsset `json:"assets,omitempty"`
	Count  *int64        `json:"count,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
