upload-demo
===========

#### Safely and efficiently upload, download and serve images over HTTP with Go and MongoDB.

This app demonstrates patterns we've found useful and trustworthy. YMMV.

##### Routing:

Method   | URL | Name
-------- | --- | ----
POST     | /images                    | create_image   |
GET      | /images/{imageId}          | show_image     |
GET      | /images/{imageId}/download | download_image |
GET      | /images/{imageId}/metadata | image_metadata |


##### Topics:

Authorization
* Does this client have permission to create an image?

Validation
* Is this an image format we accept?
* Is this upload too large?

Efficiency
* Stream all the things!
