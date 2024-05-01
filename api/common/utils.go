package common

import (
	"bytes"
	"io"
	"net/http"
)

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
