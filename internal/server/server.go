package server

import (
	"log"
	"net/http"
)

const outDirectory = "out"

func StartServer() error {
	log.Println("Starting server on 127.0.0.1:5000...")

	fs := http.FileServer(http.Dir(outDirectory))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":5000", nil))

	return nil
}
