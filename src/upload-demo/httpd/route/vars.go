package route

import "upload-demo/models"

type Vars map[string]string

func (v Vars) FileId(k string) models.FileId {
	return models.FileId(v[k])
}
