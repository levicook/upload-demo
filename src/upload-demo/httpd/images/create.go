package images

import (
	"net/http"
	"upload-demo/config"
	"upload-demo/httpd/status"
	"upload-demo/httpd/utils"
	"upload-demo/log"
	"upload-demo/models"
)

// Create performs
//
// Validations
// - does this client have permission to create an image?
// - is this an acceptable request? (multipart/form-data)
// - is this upload too large? (trusts Content-Length)
// - is this an image format we accept?
//
// Behaviors
// - Stream valid requests direcly into MongoDB
// - Stop streaming, and cleanup, if Content-Length was a lie.
// - Detect and record the file's name.
// - Detect and record the file's content-type.
// - Send metadata back to the client.
func Create(w http.ResponseWriter, r *http.Request) {

	{ // Perform all of the "easy" request validations:

		// Does this request have permission?
		if sc, st := canCreateImage(r); sc > 0 {
			http.Error(w, st, sc)
			return
		}

		// Is this an acceptable request?
		if sc, st := canAcceptContentType(r); sc > 0 {
			http.Error(w, st, sc)
			return
		}

		// Is this upload too large?
		if sc, st := canAcceptContentLength(r); sc > 0 {
			http.Error(w, st, sc)
			return
		}
	}

	// Start streaming our upload...
	mr, err := r.MultipartReader()
	log.PanicIf(err)

	file := streamUpload(mr)

	if sc, st := canAcceptFile(file); sc > 0 {
		config.FileRepo.RemoveId(file.Id())
		http.Error(w, st, sc)
		return
	}

	// Create an image record (think of this as metadata for file)
	image := models.NewImage(file)
	log.PanicIf(config.ImageRepo.Create(image))

	// write our response
	utils.WriteJson(w, status.OK, image)
}
