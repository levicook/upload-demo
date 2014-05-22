package models

func NewImage(f File) Image {
	return Image{
		Id:          f.Id(),
		ContentType: f.ContentType(),
		FileName:    f.Name(),
		Size:        f.Size(),
		CreatedAt:   f.UploadDate(),
	}
}
