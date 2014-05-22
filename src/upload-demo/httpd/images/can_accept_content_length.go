package images

import (
	"fmt"
	"net/http"
	"upload-demo/httpd/status"
)

func canAcceptContentLength(r *http.Request) (statusCode int, statusText string) {
	var cl float32

	fmt.Sscanf(r.Header.Get("Content-Length"), "%f", &cl)

	switch {
	case cl <= 0:
		statusCode = status.LengthRequired
		statusText = "Length Required"
	case cl > MaxSize:
		statusCode = status.RequestEntityTooLarge
		statusText = "Request Entity Too Large"
	}

	return
}
