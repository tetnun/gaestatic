//
// @Author Nobuhisa TAKAHASHI
//
package gaestatic

import (
    "os"
)

type StorageType string

// Storage Type Local File
const STORAGE_TYPE_FILE StorageType = "file"
// Storage Type Google Cloud Storage
const STORAGE_TYPE_GCS StorageType = "gcs"
// Storage Type Google Drive
const STORAGE_TYPE_GD StorageType = "gd"
// Storage Type Google Drive
const STORAGE_TYPE_DEFAULT StorageType = STORAGE_TYPE_FILE

// struct for Application Setting
type AppConfig struct {
    DefaultHtml string
    AuthDir string
    AuthRealm string
    AuthUser string
    AuthPass string
    PubDir string
    StorageType StorageType

    // Config for Local File
    // - File Path for Auth ("config_auth_file_path")
    AuthFilePath string
    // - File Path for Public ("config_pub_file_path")
    PubFilePath string

    // Config for Google Cloud Storage
    // - GCS Bucket for Auth ("config_auth_gcs_bucket")
    AuthGcsBucket string
    // - GCS Object Root Path for Auth ("config_auth_gcs_object_root")
    AuthGcsObjectRoot string
    // - GCS Bucket for Public ("config_pub_gcs_bucket")
    PubGcsBucket string
    // - GCS Object Root Path for Public ("config_pub_gcs_object_root")
    PubGcsObjectRoot string

}

// initialize
func (self *AppConfig) Initialize() {
    self.DefaultHtml = os.Getenv("config_default_html")
    self.AuthRealm = os.Getenv("config_auth_realm")
    self.AuthUser = os.Getenv("config_auth_user")
    self.AuthPass = os.Getenv("config_auth_pass")
    self.AuthDir = os.Getenv("config_auth_dir")
    self.PubDir = os.Getenv("config_pub_dir")
    self.StorageType = StorageType(os.Getenv("config_storage_type"))
    switch self.StorageType {
    case STORAGE_TYPE_GCS:
        self.AuthGcsBucket = os.Getenv("config_auth_gcs_bucket")
        self.AuthGcsObjectRoot = os.Getenv("config_auth_gcs_object_root")
        self.PubGcsBucket = os.Getenv("config_pub_gcs_bucket")
        self.PubGcsObjectRoot = os.Getenv("config_pub_gcs_object_root")

    case STORAGE_TYPE_GD:

    case STORAGE_TYPE_FILE:
        fallthrough
    default:
        self.StorageType = STORAGE_TYPE_DEFAULT
        self.AuthFilePath = os.Getenv("config_auth_file_path")
        self.PubFilePath = os.Getenv("config_pub_file_path")
    }
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
