# gaestatic
Google App Engine (static file with BasicAuth)

## Requirements

* Google App Engine
* Google Cloud Storage

## Example

### Files

```
+- app
  +- app.go
+- app.yaml
```

### Examples

| No. | path | Notes |
| --@ | --- | --- | 
| 1 | example/1 | Most simple example. This use static files on GCS. | 
| 2 | example/2 | This use static files on GCS and GAE. |

#### Example - app.go

```
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

```
application: {YOUR_APPLICATION}
version: {YOUR_APPLICATION_VERSION}
runtime: go
api_version: go1
threadsafe: no

handlers:
- url: /apps/.*
  script: _go_app
  secure: always

env_variables:
  config_default_html: index.html
  config_auth_dir: /apps/
  config_auth_realm: {YOUR_BASIC_AUTH_REALM}
  config_auth_user: {YOUR_BASIC_AUTH_USER_NAME}
  config_auth_pass: {YOUR_BASIC_AUTH_PASSWORD}
  config_auth_bucket: {YOUR_BASIC_AUTH_GCS_BUCKET_NAME}
  config_auth_object_root: {YOUR_BASIC_AUTH_OBJECT_ROOT}
  config_pub_dir: /
  config_pub_bucket: {YOUR_PUB_BUCKET_NAME}
  config_pub_object_root: {YOUR_PUB_OBJECT_ROOT}
```

| Name | Example Value |
| --- | --- |
| YOUR_APPLICATION | example-app |
| YOUR_APPLICATION_VERSION | 1 |
| YOUR_BASIC_AUTH_REALM | 'example realm' |
| YOUR_BASIC_AUTH_USER_NAME | user123 |
| YOUR_BASIC_AUTH_PASSWORD | pass123 |
| YOUR_BASIC_AUTH_OBJECT_ROOT | |
| YOUR_BASIC_AUTH_GCS_BUCKET_NAME | example-gcs-a |
| YOUR_BASIC_AUTH_OBJECT_ROOT | apps/ |
| YOUR_PUB_BUCKET_NAME | example-gcs-a |
| YOUR_PUB_OBJECT_ROOT | pub/ |
