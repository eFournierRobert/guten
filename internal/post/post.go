package post

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"go.yaml.in/yaml/v4"
)

var reservedTags = []string{"title", "content", "date", "excerpt"}

type Tag struct {
	Name  string
	Posts []*Post
}

type Metadata struct {
	Date     time.Time
	Tags     []string
	Title    string
	Excerpt  string
	Template string
}

type Post struct {
	Path     string
	Metadata Metadata
}

func New(path string) (Post, error) {
	p := Post{
		path,
		Metadata{},
	}

	if err := p.readMetadata(); err != nil {
		return Post{}, err
	}
	return p, nil
}

func (p *Post) readMetadata() error {
	f, err := os.Open(p.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	if scanner.Text() != "---" {
		return errors.New("No frontmatter in post: " + p.Path)
	}

	line := ""
	builder := strings.Builder{}
	for line != "---" {
		scanner.Scan()
		line = scanner.Text()
		builder.WriteString(line + "\n")
	}

	var metadata Metadata
	if err := yaml.Unmarshal([]byte(builder.String()), &metadata); err != nil {
		return err
	}

	for _, t := range metadata.Tags {
		if slices.Contains(reservedTags, t) {
			return fmt.Errorf("%s is a reserved tag", t)
		}
	}

	p.Metadata = metadata
	return nil
}
