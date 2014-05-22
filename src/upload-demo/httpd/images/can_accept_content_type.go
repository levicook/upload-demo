package images

import (
	"fmt"
	"net/http"
	"strings"
	"upload-demo/httpd/status"
)

func canAcceptContentType(r *http.Request) (statusCode int, statusText string) {
	var (
		contentType = r.Header.Get("Content-Type")
		acceptable  = strings.HasPrefix(contentType, "multipart/form-data;")
	)

	if !acceptable {
		statusCode = status.UnsupportedMediaType
		statusText = fmt.Sprintf("Unsupported Media Type: %q", contentType)
	}

	return
}
