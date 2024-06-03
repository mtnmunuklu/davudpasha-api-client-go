package davudpasha

type AssetsOperationSystemInfo struct {
	Caption        *string `json:"Caption,omitempty"`
	Manufacturer   *string `json:"Manufacturer,omitempty"`
	OSArchitecture *string `json:"OSArchitecture,omitempty"`
	ServisPack     *string `json:"ServisPack,omitempty"`
	SerialNumber   *string `json:"SerialNumber,omitempty"`
	Version        *string `json:"Version,omitempty"`
	BuildNumber    *string `json:"BuildNumber,omitempty"`
	Description    *string `json:"Description,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
