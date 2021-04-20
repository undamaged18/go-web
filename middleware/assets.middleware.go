package middleware

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var allowedHosts = []string{"localhost:4200"}

func Assets(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		referrer, err := url.Parse(r.Referer())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		var count int
		for _, host := range allowedHosts {
			if host == referrer.Host {
				count++
			}
		}
		if count <= 0 {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		stat, err := os.Stat(conf.Paths.Root + r.URL.Path)
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if stat.IsDir() {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		f, err := os.Open(fmt.Sprintf("%s/%s", conf.Paths.Root, r.URL.Path))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		etag := generateEtag(f)                    // Generate the etag for the requested file
		etagCheck := r.Header.Get("If-None-Match") // Get the etag value from the browser
		if etag == etagCheck {                     // Check if the etag values match
			// If the etag values match the file has not changed
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("ETag", etag)
			http.Error(w, http.StatusText(http.StatusNotModified), http.StatusNotModified)
			//http.Redirect(w, r, "/dist/"+r.URL.Path, http.StatusNotModified) // return a HTTP Not Modified (304)
			return
		} else { // If value has changed serve the new file to the browser
			// Set the Cache Control HTTP Headers
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("ETag", etag)
			// Split string into two parts [file_name] [file_ext]
			s := strings.SplitAfter(stat.Name(), ".")
			getContentType(w, s[len(s)-1])
		}

		next.ServeHTTP(w, r)
	})
}


// Generate a SHA1 hash value of a file
func generateEtag(f *os.File) string {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", sha1.Sum(data)) // Return a hashed value of the contents of the file
}


// Set the MIME type of files
func getContentType(w http.ResponseWriter, file string) {
	switch file {
	case "css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case "js":
		w.Header().Set("Content-Type", "application/javascript")
	case "json":
		w.Header().Set("Content-Type", "application/json")
	case "jsonld":
		w.Header().Set("Content-Type", "application/ld+json")
	case "pdf":
		w.Header().Set("Content-Type", "application/pdf")
	case "html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	case "jpg":
		w.Header().Set("Content-Type", "image/jpeg")
	case "jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case "png":
		w.Header().Set("Content-Type", "image/png")
	case "ico":
		w.Header().Set("Content-Type", "image/x-ico")
	case "svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case "gif":
		w.Header().Set("Content-Type", "image/gif")
	case "webp":
		w.Header().Set("Content-Type", "image/webp")
	case "bmp":
		w.Header().Set("Content-Type", "image/bmp")
	case "tiff":
		w.Header().Set("Content-Type", "image/tiff")
	case "tif":
		w.Header().Set("Content-Type", "image/tiff")
	case "apng":
		w.Header().Set("Content-Type", "image/apng")
	case "txt":
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	case "xml":
		w.Header().Set("Content-Type", "application/xml")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}
}

