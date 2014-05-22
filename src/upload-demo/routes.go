package main

var routes = Routes{
	{
		"show_home",
		"GET", "/", showHome,
	},
	{
		"create_image",
		"POST", "/images", createImage,
	},
	{
		"show_image",
		"GET", "/images/{imageId}", showImage,
	},
	{
		"image_download",
		"GET", "/images/{imageId}/download", imageDownload,
	},
	{
		"show_image_metadata",
		"GET", "/images/{imageId}/metadata", imageMetadata,
	},
}
