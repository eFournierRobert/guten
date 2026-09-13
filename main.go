package main

import (
	"flag"
	"fmt"
	"guten/internal/gen"
	"guten/internal/serve"
)

func main() {
	build := flag.Bool("build", false, "Generate the static website")
	peek := flag.Bool("peek", false, "Serve the web pages")
	flag.Parse()

	if *build {
		if err := gen.GenerateWebsite(); err != nil {
			fmt.Println(err)
		}
	}

	if *peek {
		if err := serve.Serve(); err != nil {
			fmt.Println(err)
		}
	}
}
