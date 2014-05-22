package config

import "upload-demo/repos"

var (
	FileRepo  repos.FileRepo
	ImageRepo repos.ImageRepo
)

func Init(
	fileRepo repos.FileRepo,
	imageRepo repos.ImageRepo,
) {
	FileRepo = fileRepo
	ImageRepo = imageRepo
}
