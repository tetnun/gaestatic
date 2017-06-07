package gaestatic

import (
	"fmt"
	"strconv"
	"io"
	"net/http"
	"strings"
	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
)

/**
 * Use Google Cloud Storage
 * https://godoc.org/cloud.google.com/go/storage
 */
func gcsHandler(w http.ResponseWriter, r *http.Request, isAuth bool) bool {
	var bucketName string
	var objectName string

	gcsConfig := config.GcsConfig

	isDone := true

	config := GetAppConfig()
	if config == nil {
		// Internal Server Errror
		w.WriteHeader(500)
		w.Write([]byte("No Config"))
		return isDone
	}

	if isAuth == true {
		// Basic認証
		if CheckBasicAuth(r) == false {
			// 認証処理
			outputUnauth(w)
			return isDone
		}
		bucketName = gcsConfig.AuthBucket
		objectName = strings.Replace(r.URL.Path, config.AuthDir, gcsConfig.AuthObjectRoot, -1)
	} else {
		bucketName = gcsConfig.PubBucket
		objectName = strings.Replace(r.URL.Path, config.PubDir, gcsConfig.PubObjectRoot, -1)
	}

	// ローカルは動作しないので未実装扱い
	if appengine.IsDevAppServer() {
		// Not Implemented
		w.WriteHeader(501)
		return isDone
	}

	ctx := appengine.NewContext(r)
	client, err := storage.NewClient(ctx)
	if err != nil {
		// Internal Server Errror
		w.WriteHeader(500)
		return isDone
	}
	defer client.Close()

	// GCS Bucket
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(objectName)

	if _, err := obj.ACL().List(ctx); err != nil {
		// Forbidden : GCS ACL許可されていない
		w.WriteHeader(403)
		w.Write([]byte(fmt.Sprintf("ACL: ObjectName=%s", objectName)))
		return isDone
	}

	var contentLength string
	if attrs, err := obj.Attrs(ctx); err != nil {
		// Forbidden : GCS サイズ取得失敗
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("ACL: ObjectName=%s", objectName)))
		return isDone
	} else {
		contentLength = strconv.FormatInt(attrs.Size, 10)
	}
	contentLength = contentLength + "bytes"

	// Read
	reader, err2 := obj.NewReader(ctx)
	if err2 != nil {
		// Not Found : GCS 読み込みエラー
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("Object Not Found. ObjectName=%s", objectName)))
		return isDone
	}
	defer reader.Close()
	contentType := GetContentType(objectName)
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}
	io.Copy(w, reader)
	isDone = true
	return isDone
}
