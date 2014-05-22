package main

import (
	"net/http"
	"upload-demo/httpd"
	"upload-demo/log"
)

func init() {

}

func main() {
	var (
		addr   = "127.0.0.1:9000"
		router = httpd.Routes.Router()
	)

	log.Printf("listening at http://%v", addr)
	log.PanicIf(http.ListenAndServe(addr, router))
}
