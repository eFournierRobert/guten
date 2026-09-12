package server

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

const outDirectory = "out"
const indexFile = "index.html"

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method + r.RequestURI)

	if r.Method != http.MethodGet {
		log.Println("invalid request: " + r.RequestURI)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	requestedPath := r.RequestURI
	if requestedPath == "/" {
		requestedPath = indexFile
	} else {
		requestedPath = requestedPath[1:]
	}

	f, err := os.OpenFile(requestedPath, os.O_RDONLY, 0440)
	if err != nil {
		log.Printf("error opening file: %s\n", fmt.Errorf("%w", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	var contentType = ""
	switch path.Ext(requestedPath) {
	case "html":
		contentType = "text/html"
	case "css":
		contentType = "text/css"
	case "ico":
		contentType = "image/vnd.microsoft.icon"
	case "js":
		contentType = "text/javascript"
	case "md":
		contentType = "text/markdown"
	case "png":
		contentType = "image/png"
	case "svg":
		contentType = "image/svg+xml"
	case "txt":
		contentType = "text/plain"
	case "jpeg", "jpg":
		contentType = "image/jpeg"
	}
	w.Header().Set("Content-Type", contentType)

	scanner := bufio.NewScanner(f)
	w.WriteHeader(http.StatusOK)
	for scanner.Scan() {
		w.Write(scanner.Bytes())
	}
	return
}

func StartServer() error {
	if err := os.Chdir(outDirectory); err != nil {
		return fmt.Errorf("error while starting http server: %w", err)
	}

	log.Println("Starting server on 127.0.0.1:5000...")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))

	return nil
}
