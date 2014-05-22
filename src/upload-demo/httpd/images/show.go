package images

import (
	"io"
	"net/http"
	"upload-demo/config"
	"upload-demo/httpd/route"
	"upload-demo/log"

	"github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
	show(w, r, route.Vars(mux.Vars(r)))
}

func show(w http.ResponseWriter, r *http.Request, v route.Vars) {
	file, err := config.FileRepo.OneById(v.FileId("imageId"))
	log.PanicIf(err)
	defer file.Close()

	// write our response
	w.Header().Set("Content-Type", file.ContentType())
	io.Copy(w, file)
}
