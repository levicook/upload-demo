package repos

import "upload-demo/models"

type ImageRepo interface {
	Create(models.Image) error
	OneById(models.FileId) (models.Image, error)
}
