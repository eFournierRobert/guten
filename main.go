package main

import (
	"flag"
	"fmt"
	"guten/internal/gen"
	"guten/internal/serve"
)

func main() {
	build := flag.Bool("build", false, "Generate the static website")
	peak := flag.Bool("peak", false, "Serve the webpage")
	flag.Parse()

	if *build {
		if err := gen.GenerateWebsite(); err != nil {
			fmt.Println(err)
		}
	}

	if *peak {
		if err := serve.Serve(); err != nil {
			fmt.Println(err)
		}
	}
}
