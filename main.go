package main

import (
	"flag"
	"fmt"
	"guten/internal/gen"
	"guten/internal/init_project"
	"guten/internal/serve"
	"os"
)

func main() {
	build := flag.Bool("build", false, "Generate the static website")
	peek := flag.Bool("peek", false, "Serve the web pages")
	initDir := flag.String("init", "", "Generate a new Guten project in the given directory")
	flag.Parse()

	if *initDir != "" {
		if err := init_project.Init(initDir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Project %s initialized successfully!\n", *initDir)
		os.Exit(0)
	}

	if *build {
		if err := gen.GenerateWebsite(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println("Website built successfully!")
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
