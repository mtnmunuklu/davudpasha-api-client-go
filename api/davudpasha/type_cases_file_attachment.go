package davudpasha

type CasesFileAttachment struct {
	FileId        *string `json:"FileId,omitempty"`
	FileName      *string `json:"FileName,omitempty"`
	MD5Hash       *string `json:"MD5Hash,omitempty"`
	SHA256Hash    *string `json:"SHA256Hash,omitempty"`
	FileSize      *string `json:"FileSize"`
	FileExtension *string `json:"FileExtension,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{}
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{}
}
