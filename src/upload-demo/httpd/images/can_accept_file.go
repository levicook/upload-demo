package images

import (
	"fmt"
	"upload-demo/httpd/status"
	"upload-demo/mime"
	"upload-demo/models"
)

func canAcceptFile(file models.File) (statusCode int, statusText string) {
	var (
		contentType = file.ContentType()
		acceptable  = mime.IsImage(contentType)
	)

	if !acceptable {
		statusCode = status.UnsupportedMediaType
		statusText = fmt.Sprintf("Unsupported Media Type: %q", contentType)
	}

	return
}
