package mongo

import (
	"fmt"
	"upload-demo/models"
	"upload-demo/repos"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func NewImageRepo(db *mgo.Database) repos.ImageRepo {
	c := db.C("images")
	return imageRepo{c}
}

type imageRepo struct {
	c *mgo.Collection
}

func (r imageRepo) Create(o models.Image) (err error) {
	if errs := o.Errors(); errs.Present() {
		return fmt.Errorf("invalid image %q %q", o, errs)
	}

	return r.c.Insert(imageWrapper{Id: fileOid(o.Id), Image: o})
}

func (r imageRepo) OneById(id models.FileId) (o models.Image, err error) {
	var wrapper imageWrapper

	if err = r.c.FindId(fileOid(id)).One(&wrapper); err != nil {
		return
	}

	o = wrapper.Image

	return
}

type imageWrapper struct {
	Id    bson.ObjectId `bson:"_id"`
	Image models.Image  `bson:"img"`
}
