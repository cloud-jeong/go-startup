package handlers

import (
	"net/http"
	"io/ioutil"
	"path/filepath"
	"fmt"
)

func home(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprint(w, "/home: Hello! Your request was processed.")

	fmt.Printf("req path: %s\n", req.URL.Path)

	localPath := "techmag" + req.URL.Path

	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}