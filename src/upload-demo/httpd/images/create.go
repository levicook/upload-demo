package images

import "net/http"

func Create(w http.ResponseWriter, r *http.Request) {

	// Validations
	// - does this client have permission to create an image?
	// - is this an acceptable request? (multipart/form-data)
	// - is this upload too large? (trusting Content-Length)
	// - is this an image format we accept?

	// Behaviors
	// - Stream valid requests direcly into MongoDB
	// - Stop streaming, and cleanup, if Content-Length was a lie.
	// - Detect and record the file's name.
	// - Detect and record the file's content-type.
	// - Send metadata back to the client.

	// --------------------------------------------

	{ // Perform all of the "easy" validations:

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

	{ // Start streaming the upload!

	}

}
