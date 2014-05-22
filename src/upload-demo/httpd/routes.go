package httpd

import (
	"upload-demo/httpd/home"
	"upload-demo/httpd/images"
	"upload-demo/httpd/route"
)

var Routes = route.Routes{
	{
		"show_home",
		"GET", "/", home.Show,
	},
	{
		"images_create",
		"POST", "/images", images.Create,
	},
	{
		"images_show",
		"GET", "/images/{imageId}", images.Show,
	},
	{
		"images_download",
		"GET", "/images/{imageId}/download", images.Download,
	},
	{
		"images_metadata",
		"GET", "/images/{imageId}/metadata", images.Metadata,
	},
}
