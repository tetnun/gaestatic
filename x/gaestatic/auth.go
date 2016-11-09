//
// @Author Nobuhisa TAKAHASHI
//
package gaestatic

import (
    "net/http"
)

// Check Basic Auth
func CheckBasicAuth(r *http.Request) bool {
    config := GetAppConfig()    
    rc := false
    username, password, ok := (*r).BasicAuth()
    if ok == false {
        return rc
    }
    if username != config.AuthUser {
        return rc
    }
    if password != config.AuthPass {
        return rc
    }
    rc = true
    return rc
}
