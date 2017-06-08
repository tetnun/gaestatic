package gaestatic

import (
	"fmt"
	"strconv"
	"io"
	"net/http"
	"strings"
	"os"
)

/**
 * Use Local File Storage
 */
func fileHandler(w http.ResponseWriter, r *http.Request, isAuth bool) bool {
	var filePath string

	fileConfig := config.FileConfig

	isDone := true

	config := GetAppConfig()
	if config == nil {
		// Internal Server Errror
		w.WriteHeader(500)
		w.Write([]byte("No Config"))
		return isDone
	}

	if isAuth == true {
		// Basic Auth.
		if CheckBasicAuth(r) == false {
			// Authentication
			outputUnauth(w)
			return isDone
		}
		filePath = strings.Replace(r.URL.Path, config.AuthDir, fileConfig.AuthPath, -1)
	} else {
		filePath = strings.Replace(r.URL.Path, config.PubDir, fileConfig.PubPath, -1)
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
		// Forbidden : Unknown Size.
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
