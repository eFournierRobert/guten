package init_project

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

const templatesDir = "templates"
const postsDir = "posts"
const indexFile = "index.html"
const assetsDir = "assets"

func Init(directory *string) error {
	if exist, err := doesPathExist(*directory); exist {
		return fmt.Errorf("error: directory %s already exists", *directory)
	} else if err != nil {
		return fmt.Errorf("error during init: %w", err)
	}

	if err := os.MkdirAll(*directory+"/"+templatesDir, 0740); err != nil {
		return fmt.Errorf("error during init: %w", err)
	}
	if err := os.MkdirAll(*directory+"/"+postsDir, 0740); err != nil {
		return fmt.Errorf("error during init: %w", err)
	}
	if err := os.MkdirAll(*directory+"/"+assetsDir, 0740); err != nil {
		return fmt.Errorf("error during init: %w", err)
	}

	fd, err := os.Create(*directory + "/" + indexFile)
	if err != nil {
		return fmt.Errorf("error during init: %w", err)
	}
	fd.Close()

	return nil
}

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
