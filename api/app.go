package api

import (
	"net/http"
	"strings"
)

func App(w http.ResponseWriter, r *http.Request) {
	filePath := "./app/dist" + r.URL.Path

	if serveFileWithContentType(w, r, filePath) {
		return
	}

	// Default to serving index.html for other paths
	filePath = "./app/dist/index.html"
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, filePath)
}

func serveFileWithContentType(w http.ResponseWriter, r *http.Request, filePath string) bool {
	switch {
	case strings.HasSuffix(r.URL.Path, ".js"):
		w.Header().Set("Content-Type", "text/javascript")
	case strings.HasSuffix(r.URL.Path, ".css"):
		w.Header().Set("Content-Type", "text/css")
	case strings.HasSuffix(r.URL.Path, ".html"):
		w.Header().Set("Content-Type", "text/html")
	case strings.HasSuffix(r.URL.Path, ".svg"):
		w.Header().Set("Content-Type", "image/svg+xml")
	default:
		return false
	}

	http.ServeFile(w, r, filePath)
	return true
}
