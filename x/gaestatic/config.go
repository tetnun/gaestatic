//
// @Author Nobuhisa TAKAHASHI
//
package gaestatic

import (
    "os"
)

// struct for Application Setting
type AppConfig struct {
    DefaultHtml string
    AuthDir string
    AuthRealm string
    AuthUser string
    AuthPass string
    AuthBucket string
    AuthObjectRoot string
    PubDir string
    PubBucket string
    PubObjectRoot string
}

// initialize
func (self *AppConfig) Initialize() {
    self.DefaultHtml = os.Getenv("config_default_html")
    self.AuthRealm = os.Getenv("config_auth_realm")
    self.AuthUser = os.Getenv("config_auth_user")
    self.AuthPass = os.Getenv("config_auth_pass")
    self.AuthDir = os.Getenv("config_auth_dir")
    self.AuthBucket = os.Getenv("config_auth_bucket")
    self.AuthObjectRoot = os.Getenv("config_auth_object_root")
    self.PubDir = os.Getenv("config_pub_dir")
    self.PubBucket = os.Getenv("config_pub_bucket")
    self.PubObjectRoot = os.Getenv("config_pub_object_root")
}

// singleton
var config *AppConfig

// create and get singleton
func GetAppConfig() *AppConfig {
    if config == nil {
        config = &AppConfig{}
        config.Initialize()
    }
    
    return config
}
