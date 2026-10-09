// Package init_project provides the function to initialize a
// directory into a new guten project.
package init_project

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Directories and files constants for initialization.
const templatesDir = "templates"
const postsDir = "posts"
const indexFile = "index.html"
const assetsDir = "assets"
const IncludesDir = "includes"

// Init function that inits a guten project in the given website directory.
func Init(directory *string) error {
	if *directory == "" {
		return errors.New("error during init: please provide a directory name")
	}

	if *directory != "." {
		if exist, err := doesPathExist(*directory); exist {
			return fmt.Errorf("error during init: %s already exists", *directory)
		} else if err != nil {
			return fmt.Errorf("error during init: %w", err)
		}
	}

	if err := os.MkdirAll(filepath.Join(*directory, templatesDir), 0740); err != nil {
		return fmt.Errorf("error during init: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(*directory, postsDir), 0740); err != nil {
		return fmt.Errorf("error during init: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(*directory, assetsDir), 0740); err != nil {
		return fmt.Errorf("error during init: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(*directory, IncludesDir), 0740); err != nil {
		return fmt.Errorf("error during init: %w", err)
	}

	if err := os.WriteFile(
		filepath.Join(*directory, indexFile),
		[]byte("<!DOCTYPE html><html><body><h1>Welcome to your new Guten website!</h1></body></html>"),
		0640); err != nil {
		return fmt.Errorf("error during init: %w", err)
	}

	return nil
}

// Checks if the given path exists.
func doesPathExist(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
