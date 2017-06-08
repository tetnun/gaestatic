//
// @Author Nobuhisa TAKAHASHI
//
package gaestatic

import (
    "os"
)

type StorageType string

type DynamicType string

// Storage Type Local File
const STORAGE_TYPE_FILE StorageType = "file"
// Storage Type Google Cloud Storage
const STORAGE_TYPE_GCS StorageType = "gcs"
// Storage Type Google Drive
const STORAGE_TYPE_GD StorageType = "gd"
// Storage Type Google Drive
const STORAGE_TYPE_DEFAULT StorageType = STORAGE_TYPE_FILE

const DYNAMIC_TYPE_PLIST DynamicType = "plist"


// Config for Local File
type FileAppConfig struct {
    // "config_auth_file_path" - File Path for Auth
    AuthPath string
    // "config_pub_file_path" - File Path for Public
    PubPath string
}

// Config for Google Cloud Storage
type GcsAppConfig struct {
    // "config_auth_gcs_bucket" - GCS Bucket for Auth
    AuthBucket string
    // "config_auth_gcs_object_root" - GCS Object Root Path for Auth
    AuthObjectRoot string
    // "config_pub_gcs_bucket" - GCS Bucket for Public
    PubBucket string
    // "config_pub_gcs_object_root" - GCS Object Root Path for Public
    PubObjectRoot string
}

// Config for Google Drive
type DriveAppConfig struct {
    // "config_client_id" - Client ID
    ClientID string
    // "config_secret" - Client Secret
    ClientSecret string
    // "config_auth_drive_path" - Drive Path for Auth
    AuthPath string
    // "config_pub_drive_path" - Drive Path for Public
    PubPath string
}

// struct for Application Setting
type AppConfig struct {
    DefaultHtml string
    AuthDir string
    AuthRealm string
    AuthUser string
    AuthPass string
    PubDir string
    PlistDir string
    StorageType StorageType

    // Config for Local File
    FileConfig FileAppConfig

    // Config for Google Cloud Storage
    GcsConfig GcsAppConfig

    // Config for Google Drive
    DriveConfig DriveAppConfig
}

// initialize
func (self *AppConfig) Initialize() {
    self.DefaultHtml = os.Getenv("config_default_html")
    self.AuthRealm = os.Getenv("config_auth_realm")
    self.AuthUser = os.Getenv("config_auth_user")
    self.AuthPass = os.Getenv("config_auth_pass")
    self.AuthDir = os.Getenv("config_auth_dir")
    self.PubDir = os.Getenv("config_pub_dir")
    self.PlistDir = os.Getenv("config_plist_dir")
    self.StorageType = StorageType(os.Getenv("config_storage_type"))
    switch self.StorageType {
    case STORAGE_TYPE_GCS:
        gcsConfig := GcsAppConfig{}
        gcsConfig.AuthBucket = os.Getenv("config_auth_gcs_bucket")
        gcsConfig.AuthObjectRoot = os.Getenv("config_auth_gcs_object_root")
        gcsConfig.PubBucket = os.Getenv("config_pub_gcs_bucket")
        gcsConfig.PubObjectRoot = os.Getenv("config_pub_gcs_object_root")
        self.GcsConfig = gcsConfig
    case STORAGE_TYPE_GD:
        driveConfig := DriveAppConfig{}
        driveConfig.ClientID = os.Getenv("config_client_id")
        driveConfig.ClientSecret = os.Getenv("config_client_secret")
        driveConfig.AuthPath = os.Getenv("config_auth_drive_path")
        driveConfig.PubPath = os.Getenv("config_pub_drive_path")
        self.DriveConfig = driveConfig
    case STORAGE_TYPE_FILE:
        fallthrough
    default:
        self.StorageType = STORAGE_TYPE_DEFAULT
        fileConfig := FileAppConfig{}
        fileConfig.AuthPath = os.Getenv("config_auth_file_path")
        fileConfig.PubPath = os.Getenv("config_pub_file_path")
        self.FileConfig = fileConfig
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
