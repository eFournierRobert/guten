package serve

// Package serve provides the HTTP server functionality for guten.
// It combines website generation with HTTP serving for local preview.
// The server runs on port 5000 and serves files from the out/ directory.

import (
	"fmt"
	"guten/internal/gen"
	"guten/internal/server"
	"log"
)

// Serve generates the website and starts an HTTP server on localhost:5000.
// It builds the site first, then serves the generated out/ directory.
func Serve() error {
	if err := gen.GenerateWebsite(); err != nil {
		return fmt.Errorf("error while serving website: %w", err)
	}

	if err := server.StartServer(); err != nil {
		log.Fatalf("error while serving http server: %s\n", fmt.Errorf("%w", err))
	}
	return nil
}
