package davudpasha

type AssetsInventoryInfo struct {
	OperationSystemInfos *AssetsOperationSystemInfo `json:"OperationSystemInfos,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
