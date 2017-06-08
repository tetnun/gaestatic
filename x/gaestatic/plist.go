package gaestatic

import (
	"strings"
	"fmt"
	"strconv"
	"net/http"
	"text/template"
	"bytes"
)

const PLIST_TEMPLATE string = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>items</key>
		<array>
			<dict>
				<key>assets</key>
				<array>
					{{if .IpaUrl}}
					<dict>
						<key>kind</key>
						<string>software-package</string>
						<key>url</key>
						<string>{{.IpaUrl}}</string>
					</dict>
					{{end}}
					{{if .DisplayImageUrl}}
					<dict>
						<key>kind</key>
						<string>display-image</string>
						<key>url</key>
						<string>{{.DisplayImageUrl}}</string>
					</dict>
					{{end}}
					{{if .FullSizeImageUrl}}
					<dict>
						<key>kind</key>
						<string>full-size-image</string>
						<key>url</key>
						<string>{{.FullSizeImageUrl}}</string>
					</dict>
					{{end}}
				</array>
				<key>metadata</key>
				<dict>
					{{if .BundleIdentifer}}
					<key>bundle-identifier</key>
					<string>{{.BundleIdentifer}}</string>
					{{end}}
					{{if .BundleVersion}}
					<key>bundle-version</key>
					<string>{{.BundleVersion}}</string>
					<key>kind</key>
					{{end}}
					{{if .Title}}
					<string>software</string>
					<key>title</key>
					<string>{{.Title}}</string>
					{{end}}
				</dict>
			</dict>
		</array>
	</dict>
</plist>
`

type PlistTemplateParams struct {
	// eg. https://example.com/apps/ios/sample.ipa
	IpaUrl string
	// eg. https://example.com/apps/ios/image.png
	DisplayImageUrl string
	// eg. https://example.com/apps/ios/full-image.png
	FullSizeImageUrl string
	// eg. com.example.sample
	BundleIdentifer string
	// eg. 1.0
	BundleVersion string
	// eg. Sample App
	Title string
}

/**
 * Dynamic Plist Handler
 */
func plistHandler(w http.ResponseWriter, r *http.Request) bool {
	var filePath string

	isDone := true

	config := GetAppConfig()
	if config == nil {
		// Internal Server Errror
		w.WriteHeader(500)
		w.Write([]byte("No Config"))
		return isDone
	}

	filePath = strings.Replace(r.URL.Path, config.PubDir, config.PlistDir, -1)
	tmp := strings.SplitN(filePath, "/", 2)
	if len(tmp) < 2 {
		// Bad Request
		w.WriteHeader(400)
		w.Write([]byte("invalid path #1"))
		return isDone
	}
	bundleIdentifer := tmp[0]
	if strings.Contains(tmp[1], "..") {
		// Bad Request
		w.WriteHeader(400)
		w.Write([]byte("invalid path #2"))
		return isDone
	}

	ipaUrl := r.URL
	ipaUrl.Path = "/" + tmp[1]

	params := PlistTemplateParams{}
	// http://example.com/{filePath}/{bundleId}/{IpaPath}?title={title}&version={bundleVersion}
	params.Title = r.URL.Query().Get("title")
	params.BundleVersion = r.URL.Query().Get("version")
	params.BundleIdentifer = bundleIdentifer
	params.IpaUrl = ipaUrl.RequestURI()

	tmpl, err := template.New("plist").Parse(PLIST_TEMPLATE)

	if err != nil {
		// Not Found
		w.WriteHeader(501)
		w.Write([]byte(fmt.Sprintf("plist template is invalid.")))
		return isDone
	}

	writer := new(bytes.Buffer)
	err	= tmpl.Execute(writer, params)

	var contentLength string
	if err != nil {
		// Forbidden : サイズ取得失敗
		w.WriteHeader(403)
		w.Write([]byte(fmt.Sprintf("plist params is invalid.")))
		return isDone
	} else {
		contentLength = strconv.FormatInt(int64(writer.Len()), 10)
	}
	contentLength = contentLength + "bytes"

	contentType := GetContentType("_.plist")
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}
	w.Write(writer.Bytes())
	isDone = true
	return isDone
}
