# gaestatic
Google App Engine (static file with BasicAuth)

## Requirements

* Google App Engine (golang)
* Google Cloud Storage
    * **READ** Permission from **GAE** for **bucket** and **object**.
    * Edit bucket permission.

        ex.)
    
            Google Cloud Platform Console
            -> Storage
            -> Select Bucket
            -> Edit bucket permissions
            -> Add item
            ENTITY=User,
            NAME=YOUR GAE Service Account (eg. example@appspot.gserviceaccount.com),
            ACCESS=Read
        
    * Edit **object** permission.

        ex.)
    
            Google Cloud Platform Console
            -> Storage
            -> Select Bucket
            -> Edit object permissions
            -> Add item
            ENTITY=User,
            NAME=YOUR GAE Service Account (eg. example@appspot.gserviceaccount.com),
            ACCESS=Read
        
    
## Example

### Files

```
+- app
  +- app.go
+- app.yaml
```

### Examples

| No. | path | Notes |
| --: | --- | --- |
| 1 | [examples/1](https://github.com/tetnun/gaestatic/tree/master/examples/1) | Simple example. This use static files on GCS. | 
| 2 | [examples/2](https://github.com/tetnun/gaestatic/tree/master/examples/2) | Comprecated example. This use static files on GCS and GAE. |

#### Example - app.go

```go
package app

import (
    tgs "github.com/tetnun/gaestatic/x/gaestatic"
)

// GAE Entry Point
func init() {
    tgs.Init()
}
```

#### Example - app.yaml

```yaml
application: {YOUR_APPLICATION}
version: {YOUR_APPLICATION_VERSION}
runtime: go
api_version: go1
threadsafe: no

handlers:
- url: /.*
  script: _go_app
  secure: always

env_variables:
  config_default_html: index.html
  config_auth_dir: {YOUR_BASIC_AUTH_DIR}
  config_auth_realm: {YOUR_BASIC_AUTH_REALM}
  config_auth_user: {YOUR_BASIC_AUTH_USER_NAME}
  config_auth_pass: {YOUR_BASIC_AUTH_PASSWORD}
  config_pub_dir: {YOUR_PUB_DIR}
  config_storage_type: {STORAGE_TYPE}
  config_auth_file_path: {YOUR_BASIC_AUTH_FILE_PATH}
  config_pub_file_path: {YOUR_PUB_FILE_PATH}
  config_auth_gcs_bucket: {YOUR_BASIC_AUTH_GCS_BUCKET_NAME}
  config_auth_gcs_object_root: {YOUR_BASIC_AUTH_GCS_OBJECT_ROOT}
  config_pub_gcs_bucket: {YOUR_PUB_GCS_BUCKET_NAME}
  config_pub_gcs_object_root: {YOUR_PUB_GCS_OBJECT_ROOT}
```

| Name | Example Value | Notes |
| --- | --- | --- |
| YOUR_APPLICATION | example-gaestatic | see GAE golang document |
| YOUR_APPLICATION_VERSION | 1 | see GAE golang document |
| YOUR_BASIC_AUTH_DIR | /apps/ | URI Path (Basic Auth) |
| YOUR_BASIC_AUTH_REALM | 'example realm' | Basic Auth Realm |
| YOUR_BASIC_AUTH_USER_NAME | user123 | Basic Auth Username |
| YOUR_BASIC_AUTH_PASSWORD | pass123 | Basic Auth Password |
| YOUR_PUB_DIR | / | URI Path (non Basic Auth) |
| STORAGE_TYPE | 'file', 'gcs', 'gd' | 'file' : Local File, 'gcs' : GCS, 'blob' : GCS Blobstore, ''gd' : Google Drive (Reserved)|
| YOUR_BASIC_AUTH_FILE_PATH | apps/ | File Path Prefix (Basic Auth) |
| YOUR_PUB_FILE_PATH | apps/ | File Path Prefix (non Basic Auth) |
| YOUR_BASIC_AUTH_GCS_BUCKET_NAME | example-gcs-apps | GCS Bucket (Basic Auth) |
| YOUR_BASIC_AUTH_GCS_OBJECT_ROOT | apps/ | GCS Object Prefix (Basic Auth) |
| YOUR_PUB_GCS_BUCKET_NAME | example-gcs-pub | GCS Bucket (non Basic Auth) |
| YOUR_PUB_GCS_OBJECT_ROOT | pub/ | GCS Object Prefix (non Basic Auth) |

### URL mapping

| No.| Basic Auth | URL | Local Location |
| --: | :-: | --- | --- |
| 1 | Yes | https://{YOUR_APPLICATION}.appspot.com/apps/ | gs://{YOUR_BASIC_AUTH_GCS_BUCKET_NAME}/{YOUR_BASIC_AUTH_GCS_OBJECT_ROOT} |
| 2 | No | https://{YOUR_APPLICATION}.appspot.com/ (\*1.)| gs://{YOUR_PUB_GCS_BUCKET_NAME}/{YOUR_PUB_GCS_OBJECT_ROOT} |

*1. exclude URL (No.1)

##### ex.

| No. | Basic Auth | URL | Local Location |
| --: | :-: | --- | --- |
| 1 | Yes | https://example-gaestatic.appspot.com/apps/ | gs://example-gcs-apps/apps/ |
| 2 | No | https://example-gaestatic.appspot.com/pub/ | gs://example-gcs-apps/pub/ 



### iOS OTA Download URL


#### example values

| Item        | Notes                                 | Example                                  |
| ----------- | ------------------------------------- | ---------------------------------------- |
| GAE_APP_URL | GAE App URL                           | https://example.appspot.com              |
| PLIST_DIR   | config_plist_dir in app.yaml          | \/\_\_plist\_\_\/                        |
| BUNDLE_ID   | 'bundle-identifier' in plist          | com.example.sample                       |
| VERSION     | 'bundle-version' in plist             | 1.0                                      |
| IPA_PATH    | Path part of 'url' in plist           | apps/ios/sample.ipa on GAE_APP_URL (*1.) |
| IMG_PATH    | path part of 'display-image' in plist | apps/ios/sample.png on GAE_APP_URL (*2.) |
| TITLE       | 'title' in plist                      | SampleTitle                              |

*1. absolute url is https://example.appspot.com/apps/ios/sample.ipa

*2. absolute url is https://example.appspot.com/apps/ios/sample.png


#### PLIST URL format

${GAE_APP_URL}${PLIST_DIR}/${BUNDLE_ID}/${VERSION}/${IPA_PATH}/${IMG_PATH}/x.plist?title=${TITLE}

ex.) plist

    https://example.appspot.com/__plist__/com.example.sample/1.0/apps/ios/sample.ipa/apps/ios/image.png/x.plist?title=SampleTitle

ex.) download url

    itms-services://?action=download-manifest&url=https://example.appspot.com/__plist__/com.example.sample/1.0/apps/ios/sample.ipa/apps/ios/image.png/x.plist%3Ftitle%3DSampleTitle
