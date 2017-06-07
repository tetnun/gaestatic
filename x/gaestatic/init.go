//
// @Author Nobuhisa TAKAHASHI
//
package gaestatic

import (
    "fmt"
    "net/http"
)

// Basic Auth Path Handler
func authHandler(w http.ResponseWriter, r *http.Request) {
    var isAuth bool = true
    switch config.StorageType {
    case STORAGE_TYPE_GCS:
        gcsHandler(w, r, isAuth)
    default:
        fileHandler(w, r, isAuth)
    }
}

// Public Path Handler
func pubHandler(w http.ResponseWriter, r *http.Request) {
    var isAuth bool = false
    switch config.StorageType {
    case STORAGE_TYPE_GCS:
        gcsHandler(w, r, isAuth)
    default:
        fileHandler(w, r, isAuth)
    }
}

// All Error Handler
func allErrorHandler(w http.ResponseWriter, r *http.Request) {
    // Internal Server Errror
    w.WriteHeader(500)
}

// Basic Auth Handler
func outputUnauth(w http.ResponseWriter) {
    config := GetAppConfig()
    realm := config.AuthRealm
    if realm == "" {
        realm = "gaestatic realm"
    }
    w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
    w.WriteHeader(401)
}


func Init() {
    var authDir string
    config := GetAppConfig()
    
    if config == nil {
        http.HandleFunc("/", allErrorHandler)
        return        
    }
    authDir = config.AuthDir
    if config.AuthDir != "" {
        http.HandleFunc(authDir, authHandler)        
    }
    http.HandleFunc("/", pubHandler)
}