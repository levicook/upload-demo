package mongo

import (
	"upload-demo/log"

	"labix.org/v2/mgo"
)

var testDB *mgo.Database

func setupTestDB() {
	s, err := mgo.Dial("127.0.0.1")
	log.PanicIf(err)

	db := s.DB("upload-demo-test")
	log.PanicIf(db.DropDatabase())

	testDB = db
}

func teardownDB() {
	if testDB != nil && testDB.Session != nil {
		testDB.Session.Close()
	}
}
