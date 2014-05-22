package main

import (
	"net/http"
	"upload-demo/config"
	"upload-demo/httpd"
	"upload-demo/log"
	"upload-demo/mongo"

	"labix.org/v2/mgo"
)

var (
	mongoAddr = "127.0.0.1"
	mongoDB   = "upload-demo"
	httpdAddr = "127.0.0.1:9000"
)

func main() {
	s, err := mgo.Dial(mongoAddr)
	log.PanicIf(err)
	defer s.Close()

	db := s.DB(mongoDB)
	config.Init(
		mongo.NewFileRepo(db),
		mongo.NewImageRepo(db),
	)

	router := httpd.Routes.Router()
	log.Printf("listening at http://%v", httpdAddr)
	log.PanicIf(http.ListenAndServe(httpdAddr, router))
}
