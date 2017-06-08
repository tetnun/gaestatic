package gaestatic

import (
	"strings"
	"fmt"
	"strconv"
	"net/http"
	"text/template"
	"bytes"
	"net/url"
	"regexp"
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
					<string>software</string>
					<key>title</key>
					{{if .Title}}
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
 * http://example.com/{filePath}/{bundleId}/{bundleVersion}/{IpaPath}/{imagePath}/x.plist
 *   ?title={title}&full-image={fullSizeImagePath}
 * IpaPath : extension must be '.ipa'
 */
func plistHandler(w http.ResponseWriter, r *http.Request) bool {
	const DYNAMIC_PLIST_POSTFIX = "/x.plist"
	isDone := true

	config := GetAppConfig()
	if config == nil {
		// Internal Server Errror
		w.WriteHeader(500)
		w.Write([]byte("No Config"))
		return isDone
	}

	query := r.URL.Query()

	filePath := strings.Replace(r.URL.Path, config.PlistDir, "", 1)
	tmp := strings.SplitN(filePath, "/", 3)
	if len(tmp) < 3 {
		// Bad Request
		w.WriteHeader(400)
		w.Write([]byte("invalid path #1"))
		return isDone
	}
	// Bundle Identifer (Required)
	bundleIdentifer := tmp[0]
	// Bundle Version (Required)
	bundleVersion := tmp[1]
	// Parse
	filePath = tmp[2]

	//
	rex := regexp.MustCompile(`^(.+\.ipa)/(.+)/x.plist$`)
	if !rex.MatchString(filePath) {
		// Bad Request
		w.WriteHeader(400)
		w.Write([]byte("invalid path #2"))
		return isDone
	}
	tmp = rex.FindStringSubmatch(filePath)
	ipaPath := tmp[1]
	imagePath := tmp[2]


	if len(bundleIdentifer) == 0 {
		// Bad Request
		w.WriteHeader(400)
		w.Write([]byte("bundle-identifier is required"))
		return isDone
	}

	if len(bundleVersion) == 0 {
		// Bad Request
		w.WriteHeader(400)
		w.Write([]byte("bundle-version is required"))
		return isDone
	}

	// Display Image Url (Required)
	if len(imagePath) == 0 {
		// Bad Request
		w.WriteHeader(400)
		w.Write([]byte("display-image is required"))
		return isDone
	}

	// Full Image Url (Optional)
	fullImageUrl := query.Get("full-image")
	// Title (Optional)
	title := query.Get("title")

	ipaUrl, _ := url.Parse(r.URL.String())
	if !r.URL.IsAbs() {
		ipaUrl.Scheme = "http"
		ipaUrl.Host = r.Host
	}
	ipaUrl.RawQuery = ""
	ipaUrl.Path = "/" + ipaPath

	imageUrl, _ := url.Parse(r.URL.String())
	if !r.URL.IsAbs() {
		imageUrl.Scheme = "http"
		imageUrl.Host = r.Host
	}
	imageUrl.RawQuery = ""
	imageUrl.Path = "/" + imagePath

	params := PlistTemplateParams{}
	params.Title = title
	params.BundleVersion = bundleVersion
	params.BundleIdentifer = bundleIdentifer
	params.IpaUrl = ipaUrl.String()
	params.DisplayImageUrl = imageUrl.String()
	params.FullSizeImageUrl = fullImageUrl

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
