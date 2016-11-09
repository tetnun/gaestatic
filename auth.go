package gaestatic

import (
    "net/http"
)

/**
 * Basic認証情報が一致しているかのチェック
 */
func CheckBasicAuth(r *http.Request) bool {
    config := GetAppConfig()
    
    rc := false
    //rc := true
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
