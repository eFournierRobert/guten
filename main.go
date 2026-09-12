package main

import (
	"flag"
	"fmt"
	"guten/internal/gen"
	"guten/internal/serve"
)

func main() {
	generate := flag.Bool("generate", false, "Generate the static website")
	servePages := flag.Bool("serve", false, "Serve the webpage")
	flag.Parse()

	if *generate {
		if err := gen.GenerateWebsite(); err != nil {
			fmt.Println(err)
		}
	}

	if *servePages {
		if err := serve.Serve(); err != nil {
			fmt.Println(err)
		}
	}
}
