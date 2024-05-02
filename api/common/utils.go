package common

import (
	"bytes"
	"io"
	"net/http"
)

// PtrBool is a helper routine that returns a pointer to given boolean value.
func PtrBool(v bool) *bool { return &v }

// PtrInt64 is a helper routine that returns a pointer to given integer value.
func PtrInt64(v int64) *int64 { return &v }

// PtrString is a helper routine that returns a pointer to given string value.
func PtrString(v string) *string { return &v }

// Copy the original request so it doesn't get lost when retrying
func copyRequest(r *http.Request, rawBody *[]byte) *http.Request {
	newRequest := *r

	if r.Body == nil || r.Body == http.NoBody {
		return &newRequest
	}
	newRequest.Body = io.NopCloser(bytes.NewBuffer(*rawBody))
	return &newRequest
}

// DeleteKeys helper method to delete keys from a map
func DeleteKeys[V any](obj map[string]V, keysToDelete *[]string) {
	for _, s := range *keysToDelete {
		delete(obj, s)
	}
}
