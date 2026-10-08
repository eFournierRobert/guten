package server

// Package server provides a simple HTTP file server for serving the generated site.
// It serves static files from the out/ directory on port 5000.

import (
	"log"
	"net/http"
)

const outDirectory = "out"

// StartServer starts an HTTP server on localhost:5000 serving the out/ directory.
// Uses http.FileServer to serve static files. Blocks until interrupted.
func StartServer() error {
	log.Println("Starting server on 127.0.0.1:5000...")

	fs := http.FileServer(http.Dir(outDirectory))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":5000", nil))

	return nil
}
