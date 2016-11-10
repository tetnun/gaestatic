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
            -> Edit **bucket** permissions
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
| 1 | example/1 | Most simple example. This use static files on GCS. | 
| 2 | example/2 | This use static files on GCS and GAE. |

#### Example - app.go

```golang
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
  config_auth_bucket: {YOUR_BASIC_AUTH_GCS_BUCKET_NAME}
  config_auth_object_root: {YOUR_BASIC_AUTH_OBJECT_ROOT}
  config_pub_dir: {YOUR_PUB_DIR}
  config_pub_bucket: {YOUR_PUB_BUCKET_NAME}
  config_pub_object_root: {YOUR_PUB_OBJECT_ROOT}
```

| Name | Example Value | Notes |
| --- | --- | --- |
| YOUR_APPLICATION | example-gaestatic | see GAE golang document |
| YOUR_APPLICATION_VERSION | 1 | see GAE golang document |
| YOUR_BASIC_AUTH_DIR | /apps/ | URI Path (Basic Auth) |
| YOUR_BASIC_AUTH_REALM | 'example realm' | Basic Auth Realm |
| YOUR_BASIC_AUTH_USER_NAME | user123 | Basic Auth Username |
| YOUR_BASIC_AUTH_PASSWORD | pass123 | Basic Auth Password |
| YOUR_BASIC_AUTH_GCS_BUCKET_NAME | example-gcs-apps | GCS Bucket (Basic Auth) |
| YOUR_BASIC_AUTH_OBJECT_ROOT | apps/ | GCS Object Prefix (Basic Auth) |
| YOUR_PUB_DIR | / | URI Path (non Basic Auth) |
| YOUR_PUB_BUCKET_NAME | example-gcs-pub | GCS Bucket (non Basic Auth) |
| YOUR_PUB_OBJECT_ROOT | pub/ | GCS Object Prefix (non Basic Auth) |

### URL mapping

| No.| Basic Auth | URL | Local Location |
| --: | :-: | --- | --- |
| 1 | Yes | https://{YOUR_APPLICATION}.appspot.com/apps/ | gs://{YOUR_BASIC_AUTH_GCS_BUCKET_NAME}/{YOUR_BASIC_AUTH_OBJECT_ROOT} |
| 2 | No | https://{YOUR_APPLICATION}.appspot.com/ (\*1.)| gs://{YOUR_PUB_BUCKET_NAME}/{YOUR_PUB_OBJECT_ROOT} |

*1. exclude URL (No.1)

##### ex.

| No. | Basic Auth | URL | Local Location |
| --: | :-: | --- | --- |
| 1 | Yes | https://example-gaestatic.appspot.com/apps/ | gs://example-gcs-apps/apps/ |
| 2 | No | https://example-gaestatic.appspot.com/pub/ | gs://example-gcs-apps/pub/ 



