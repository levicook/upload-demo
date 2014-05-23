package images

import (
	"net/http"
	"upload-demo/httpd/status"
)

func canCreateImage(r *http.Request) (statusCode int, statusText string) {
	println(r.Header.Get("X-SECRET"))
	var acceptable = r.Header.Get("X-SECRET") == "sekret"

	if !acceptable {
		statusCode = status.Unauthorized
		statusText = "Unauthorized"
	}

	return
}
