package post

// Package post handles parsing of Markdown posts with YAML frontmatter.
// Each post file must start with YAML frontmatter (between --- delimiters)
// containing date, title, excerpt, template, and tags fields.
// Markdown content is converted to HTML using the gomarkdown library.

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"go.yaml.in/yaml/v4"
)

// reservedTags contains tag names that cannot be used in posts.
// These names are used for template variables and would cause conflicts.
var reservedTags = []string{"title", "content", "date", "excerpt"}

// Tag represents a collection of posts sharing the same tag.
type Tag struct {
	Name  string  // Tag name (used in {{ tagname }} template syntax)
	Posts []*Post // Posts belonging to this tag
}

// Metadata contains YAML frontmatter data for a post.
type Metadata struct {
	Date     time.Time // Post publication date
	Tags     []string  // Array of tag names for categorization
	Title    string    // Post title
	Excerpt  string    // Short summary (shown in tag previews)
	Template string    // Template name (must match templates/{name}.html)
}

// Post represents a Markdown file with YAML frontmatter.
// The Path field contains the relative path to the post file.
type Post struct {
	Path     string   // File path to the post
	Metadata Metadata // Parsed YAML frontmatter
}

var defaultQuotes = []string{
	"Every page begins with a blank one.",
	"TODO: write something interesting.",
	"A thought worth keeping.",
	"Begin anywhere.",
	"One thing at a time.",
	"The best time to write was yesterday.",
	"Make something. Write it down.",
	"Less, but better.",
	"Nothing fancy. Just words.",
	"There is no perfect first draft.",
	"Small things add up.",
	"Ideas are cheap. Ship them.",

	/// Ken Thompson
	"\"One of my most productive days was throwing away 1,000 lines of code.\"\n-Ken Thompson",
}

// New creates a Post from a file path by parsing its YAML frontmatter.
// Returns an error if the file has no frontmatter or has reserved tag names.
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

// NewEmpty creates a new post file at path with default metadata
// (today's date, "post" template) and a random placeholder quote as content.
// If path already exists as a file, it asks before overwriting; directories are never overwritten.
func NewEmpty(path string) (Post, error) {
	f, err := os.Stat(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Post{}, err
	}
	if err == nil {
		if f.IsDir() {
			return Post{}, errors.New("cannot overwrite an existing directory")
		}

		var userInput string
		fmt.Printf("Error: %s already exists\n", path)
		fmt.Print("Do you want to overwrite it? [y/N] ")
		fmt.Scanln(&userInput)

		if strings.ToLower(userInput) != "y" {
			return Post{}, errors.New("refused to overwrite existing file")
		}
	}

	metadata := Metadata{
		Date:     time.Now(),
		Tags:     []string{},
		Title:    "New post",
		Excerpt:  "One great excerpt",
		Template: "post",
	}

	p := Post{
		Path:     path,
		Metadata: metadata,
	}

	content := p.buildPostMdContent()

	if err := os.WriteFile(path, content, 0640); err != nil {
		return Post{}, err
	}
	return p, nil
}

// buildPostMdContent writes the post's YAML frontmatter followed by a random
// placeholder quote, ready to be saved as the post file.
func (p *Post) buildPostMdContent() []byte {
	var sb strings.Builder
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("date: %s\n", p.Metadata.Date.Format(time.DateOnly)))
	sb.WriteString(fmt.Sprintf("title: %s\n", p.Metadata.Title))
	sb.WriteString(fmt.Sprintf("excerpt: %s\n", p.Metadata.Excerpt))
	sb.WriteString(fmt.Sprintf("template: %s\n", p.Metadata.Template))
	sb.WriteString("tags: \n")
	sb.WriteString("---\n")

	sb.WriteString("\n" + defaultQuotes[rand.Intn(len(defaultQuotes))])

	return []byte(sb.String())
}

// GetHTMLContent reads the Markdown content from the post file and converts it to HTML.
// It skips the YAML frontmatter and parses the remaining Markdown content.
// Uses gomarkdown with CommonExtensions and adds target="_blank" to links.
func (p *Post) GetHTMLContent() ([]byte, error) {
	extensions := parser.CommonExtensions | parser.NoEmptyLineBeforeBlock
	mdParser := parser.NewWithExtensions(extensions)

	f, err := os.Open(p.Path)
	if err != nil {
		return nil, fmt.Errorf("error while opening %s: %w", p.Path, err)
	}

	scanner := bufio.NewScanner(f)
	p.skipFrontMatter(scanner)
	var contentBuilder strings.Builder

	for scanner.Scan() {
		contentBuilder.WriteString(scanner.Text() + "\n")
	}

	doc := mdParser.Parse([]byte(contentBuilder.String()))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer), nil
}

// GetPostLink returns the HTML output path for this post by replacing .md with .html.
func (p *Post) GetPostLink() string {
	return strings.ReplaceAll(p.Path, ".md", ".html")
}

// skipFrontMatter advances the scanner past the YAML frontmatter block.
// It assumes the file starts with "---" for the opening delimiter.
func (p *Post) skipFrontMatter(s *bufio.Scanner) {
	s.Scan()
	if s.Text() == "---" {
		for s.Scan() {
			if s.Text() == "---" {
				break
			}
		}
	}
}

// readMetadata reads and parses the YAML frontmatter from the post file.
// Returns an error if the file has no frontmatter or uses reserved tag names.
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
