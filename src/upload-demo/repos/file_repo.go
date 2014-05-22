package repos

import "upload-demo/models"

type FileRepo interface {
	Create() (models.File, error)
	OneById(models.FileId) (models.File, error)
	RemoveId(models.FileId) error
}
