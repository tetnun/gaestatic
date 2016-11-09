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
            break
        case ".ipa":
            contentType = "application/octet-stream"
            break
        case ".apk":
            contentType = "application/vnd.android.package-archive"
            break
    }
    return contentType
}