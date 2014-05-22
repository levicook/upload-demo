package mongo

import (
	"upload-demo/models"
	"upload-demo/repos"

	"labix.org/v2/mgo"
)

func NewFileRepo(db *mgo.Database) repos.FileRepo {
	return fileRepo{db.GridFS("files")}
}

type fileRepo struct{ gridFS *mgo.GridFS }

func (r fileRepo) Create() (models.File, error) {
	gfs, err := r.gridFS.Create("")
	return &file{gfs}, err
}

func (r fileRepo) OneById(id models.FileId) (models.File, error) {
	gfs, err := r.gridFS.OpenId(fileOid(id))
	return &file{gfs}, err
}

func (r fileRepo) RemoveId(id models.FileId) error {
	return r.gridFS.RemoveId(fileOid(id))
}
