package images

import (
	"fmt"
	"net/http"
	"upload-demo/httpd/status"
)

func canAcceptContentLength(r *http.Request) (statusCode int, statusText string) {
	const (
		_          = iota // ignore first value by assigning to blank identifier
		KB float32 = 1 << (10 * iota)
		MB
	)

	var cl float32
	_, _ = fmt.Sscanf(r.Header.Get("Content-Length"), "%f", &cl)

	switch {
	case cl <= 0:
		statusCode = status.LengthRequired
		statusText = "Length Required"
	case cl > 10*MB:
		statusCode = status.RequestEntityTooLarge
		statusText = "Request Entity Too Large"
	}

	return
}
