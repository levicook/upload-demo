package images

import (
	"fmt"
	"net/http"
	"strings"
	"upload-demo/httpd/status"
)

func canAcceptContentType(r *http.Request) (statusCode int, statusText string) {
	ct := r.Header.Get("Content-Type")

	if !strings.HasPrefix(ct, "multipart/form-data;") {
		statusCode = status.UnsupportedMediaType
		statusText = fmt.Sprintf("Unsupported Media Type: %q", ct)
	}

	return
}
