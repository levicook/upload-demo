package mongo

import (
	"testing"

	"labix.org/v2/mgo/bson"
)

func Test_safeOid(t *testing.T) {
	oid := safeOid("")
	if oid != bson.ObjectIdHex("000000000000000000000000") {
		t.Fatalf("bad oid: %q", oid)
	}
}
