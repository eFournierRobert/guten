package serve

import (
	"fmt"
	"guten/internal/gen"
	"guten/internal/server"
	"log"
)

func Serve() error {
	if err := gen.GenerateWebsite(); err != nil {
		return fmt.Errorf("error while serving website: %w", err)
	}

	if err := server.StartServer(); err != nil {
		log.Fatalf("error while serving http server: %s\n", fmt.Errorf("%w", err))
	}
	return nil
}
