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


##### Validation:

Authorization: Does this client have permission to create an image?

Content Type: Is this an image format we'll accept?
MIME Type Detection and validation

Content Length: Is this upload too large?
        Optimist: checking the Content-Length header
        Pessimist: Content-Length header was a lie! Limiting what we're willing to read.
        
