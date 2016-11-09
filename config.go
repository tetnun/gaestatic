package gaestatic

import (
    "os"
)

/**
 * アプリ設定格納用struct
 */
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

/**
 * アプリ設定格納用初期化処理
 */
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

var config *AppConfig

/**
 * アプリ設定シングルトン処理
 */
func GetAppConfig() *AppConfig {
    if config == nil {
        config = &AppConfig{}
        config.Initialize()
    }
    
    return config
}
