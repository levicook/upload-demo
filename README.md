upload-demo
===========

### Safely and efficiently upload, download and serve images over HTTP with Go and MongoDB.

This app demonstrates patterns we've found useful and trustworthy. YMMV.


#### Topics:

Architecture
* Stream all the things!

Request Routing

Authorization
* Does this client have permission to create an image?

Validation
* Is this an image format we accept?
* Is this upload too large?


##### Request Routing

		Method   | URL                        | Name           |
		-------- | ---                        | ----           |
		POST     | /images                    | create_image   |
		GET      | /images/{imageId}          | show_image     |
		GET      | /images/{imageId}/download | image_download |
		GET      | /images/{imageId}/metadata | image_metadata |


```bash
$ curl localhost:9000/images -X POST --form "file=@assets/minimal.pdf"
Unauthorized

$ curl localhost:9000/images -X POST --form "file=@assets/minimal.png"
Unauthorized

$ curl localhost:9000/images -X POST --form "file=@assets/minimal.png" -H "X-SECRET: sekret"
{"id":"537ea1a8565c877315000001","contentType":"image/png","fileName":"minimal.png","size":8427,"createdAt":"2014-05-22T19:17:28.404-06:00"}

$ curl localhost:9000/images/537ea1a8565c877315000001/metadata
{"id":"537ea1a8565c877315000001","contentType":"image/png","fileName":"minimal.png","size":8427,"createdAt":"2014-05-22T19:17:28.404-06:00"}
```
