package main

import (
	"flag"
	"fmt"
	"guten/internal/gen"
	"guten/internal/init_project"
	"guten/internal/serve"
	"os"
)

// guten is a static site generator. Run with -build to generate a static site,
// or -peek to build and serve on localhost:5000.
func main() {
	build := flag.Bool("build", false, "Generate the static website")
	peek := flag.Bool("peek", false, "Serve the web pages")
	initDir := flag.String("init", "", "Generates a new Guten projects in the given directory")
	flag.Parse()

	if *initDir != "" {
		if err := init_project.Init(initDir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if *build {
		if err := gen.GenerateWebsite(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		os.Exit(0)
	}

	if *peek {
		if err := serve.Serve(); err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		os.Exit(0)
	}
}
