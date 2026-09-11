package main

import (
	"flag"
	"fmt"
	"guten/internal/gen"
)

func main() {
	generate := flag.Bool("generate", false, "Generate the static website")
	serve := flag.Bool("serve", false, "Serve the webpage")
	flag.Parse()

	if *generate {
		if err := gen.GenerateWebsite(); err != nil {
			fmt.Println(err)
		}
	}

	if *serve {
	}

}
