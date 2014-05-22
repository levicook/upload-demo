upload-demo
===========

### Safely and efficiently upload, download and serve images over HTTP with Go and MongoDB.

This app demonstrates patterns we've found useful and trustworthy. YMMV.


#### Topics:

Architecture
* Efficiency: Stream all the things!

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
GET      | /images/{imageId}/download | download_image |
GET      | /images/{imageId}/metadata | image_metadata |
