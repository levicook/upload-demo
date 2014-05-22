package images

import (
	"net/http"
	"upload-demo/httpd/status"
)

func canCreateImage(r *http.Request) (statusCode int, statusText string) {
	var acceptable = len(r.Header.Get("X-I-HAVE-THE-SECRET")) > 0

	if !acceptable {
		statusCode = status.Unauthorized
		statusText = "Unauthorized"
	}

	return
}
