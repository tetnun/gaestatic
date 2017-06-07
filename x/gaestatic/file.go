package gaestatic

import (
	"fmt"
	"strconv"
	"io"
	"net/http"
	"strings"
	"google.golang.org/appengine"
	"os"
)

/**
 * Use Local File Strage
 */
func fileHandler(w http.ResponseWriter, r *http.Request, isAuth bool) bool {
	var filePath string

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
		filePath = strings.Replace(r.URL.Path, config.AuthDir, config.AuthFilePath, -1)
	} else {
		filePath = strings.Replace(r.URL.Path, config.PubDir, config.PubFilePath, -1)
	}

	// ローカルは動作しないので未実装扱い
	if appengine.IsDevAppServer() {
		// Not Implemented
		w.WriteHeader(501)
		return isDone
	}

	file, err := os.Open(filePath)
	if err != nil {
		// Not Found
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("File Not Found: FilePath=%s", filePath)))
		return isDone
	}
	defer file.Close()

	var contentLength string
	if stat, err := file.Stat(); err != nil {
		// Forbidden : サイズ取得失敗
		w.WriteHeader(403)
		w.Write([]byte(fmt.Sprintf("File Not Found: FilePath=%s", filePath)))
		return isDone
	} else {
		contentLength = strconv.FormatInt(stat.Size(), 10)
	}
	contentLength = contentLength + "bytes"

	contentType := GetContentType(filePath)
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}
	io.Copy(w, file)
	isDone = true
	return isDone
}
