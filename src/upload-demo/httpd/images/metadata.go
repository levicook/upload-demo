package images

import (
	"net/http"
	"upload-demo/config"
	"upload-demo/httpd/route"
	"upload-demo/httpd/status"
	"upload-demo/httpd/utils"
	"upload-demo/log"

	"github.com/gorilla/mux"
)

func Metadata(w http.ResponseWriter, r *http.Request) {
	metadata(w, r, route.Vars(mux.Vars(r)))
}

func metadata(w http.ResponseWriter, r *http.Request, v route.Vars) {
	image, err := config.ImageRepo.OneById(v.FileId("imageId"))
	log.PanicIf(err)
	utils.WriteJson(w, status.OK, image)
}
