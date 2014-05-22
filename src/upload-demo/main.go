package main

import (
	"net/http"
	"upload-demo/log"
)

func main() {

	httpAddr := "127.0.0.1:9000"
	log.Printf("listening at http://%v", httpAddr)
	log.PanicIf(http.ListenAndServe(httpAddr, routes.Router()))
}
