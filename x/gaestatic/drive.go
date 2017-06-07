package gaestatic

import (
	"fmt"
	"strconv"
	"io"
	"net/http"
	"strings"
	"os"
	"io/ioutil"
	"log"
)


func valueOrFileContents(value string, filename string) string {
	if value != "" {
		return value
	}
	slurp, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading %q: %v", filename, err)
	}
	return strings.TrimSpace(string(slurp))
}

/**
 * Use Google Drive Storage
 * https://github.com/google/google-api-go-client/blob/master/examples/
 */
func driveHandler(w http.ResponseWriter, r *http.Request, isAuth bool) bool {
	var filePath string

	driveConfig := config.DriveConfig

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
		filePath = strings.Replace(r.URL.Path, config.AuthDir, driveConfig.AuthPath, -1)
	} else {
		filePath = strings.Replace(r.URL.Path, config.PubDir, driveConfig.PubPath, -1)
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
