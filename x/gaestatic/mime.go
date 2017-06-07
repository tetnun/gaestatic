//
// @Author Nobuhisa TAKAHASHI
//
package gaestatic

import (
    "regexp"
)

/**
 * 拡張子取得
 */
func GetExt(s string) string {
    rex := regexp.MustCompile(`\.[a-zA-Z0-9]+$`)
    m := rex.FindString(s)

    return m
}

/**
 * パスの文字列から拡張子部分に対応するContent-Typeを返す
 * 対応：
 *  .plist, .ipa, .apk
 */
func GetContentType(s string) string {
    ext := GetExt(s)
    var contentType string
    switch ext {
        case ".plist":
            contentType = "application/x-plist"
        case ".ipa":
            contentType = "application/octet-stream"
        case ".apk":
            contentType = "application/vnd.android.package-archive"
        case ".png":
            contentType = "image/png"
        case ".gif":
            contentType = "image/gif"
        case ".jpg":
            fallthrough
        case ".jpeg":
            contentType = "image/jpeg"
    }
    return contentType
}