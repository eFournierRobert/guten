package gen

// Package gen provides the website generator for guten.
// It parses Markdown posts with YAML frontmatter, applies HTML templates,
// and generates a static site in the out/ directory.
// Requires posts/, templates/, assets/, and index.html in the project root.
// The includes/ directory is optional and provides {{ include:filename }}
// reusable HTML snippets.

import (
	"errors"
	"fmt"
	"guten/internal/post"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

// Directory constants for the build process.
const outDir = "out"
const assetsDir = "assets"
const postsDir = "posts"
const templateDir = "templates"
const includesDir = "includes"

// Template variable tags available in template files.
const titleTag = "{{ title }}"
const contentTag = "{{ content }}"
const dateTag = "{{ date }}"
const excerptTag = "{{ excerpt }}"

// Include is a reusable HTML snippet loaded from the includes/ directory.
// It is referenced in templates and root HTML files with the
// {{ include:filename }} notation.
type Include struct {
	name        string
	fileContent []byte
}

// Generator handles the static site generation process.
// It collects posts, builds tag indexes, and generates output files.
type Generator struct {
	posts    []post.Post
	tags     []post.Tag
	includes []Include
}

// GenerateWebsite builds a static site from posts and templates.
// It clears out/, copies assets/, processes posts, and generates HTML.
// Requires: index.html in root, assets/ directory, posts/, and templates/.
func GenerateWebsite() error {
	if err := os.RemoveAll(outDir); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(outDir, postsDir), 0740); err != nil && !os.IsExist(err) {
		return err
	}

	generator := Generator{}

	if err := generator.copyAssets(); err != nil {
		return err
	}

	err := generator.getAllPosts(postsDir)
	if err != nil {
		return err
	}

	// Includes must be collected after posts so that tag notations inside
	// include files ({{ tagname }}) can be expanded with the built tag index.
	if err := generator.getAllIncludesTags(); err != nil {
		return err
	}

	err = generator.copyRootHtml()
	if err != nil {
		return err
	}

	err = generator.generatePosts()
	if err != nil {
		return err
	}

	return nil
}

// copyAssets copies the assets/ directory to out/assets/.
// Required for the build to succeed - the directory must exist.
func (g *Generator) copyAssets() error {
	assetsDest := filepath.Join(outDir, assetsDir)

	if err := os.Mkdir(assetsDest, 0740); err != nil && !os.IsExist(err) {
		return err
	}

	if err := os.CopyFS(assetsDest, os.DirFS(assetsDir)); err != nil && !os.IsExist(err) {
		return fmt.Errorf("error while copying assets directory: %w", err)
	}

	return nil
}

// copyRootHtml copies all HTML files from the project root to out/ and
// replaces {{ tagname }} with tag preview sections (lists of posts with that tag).
// The index.html file must exist in the root for a homepage to be generated.
func (g *Generator) copyRootHtml() error {
	entries, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("error while reading directory: %w", err)
	}

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".html") || entry.IsDir() {
			continue
		}

		filename := entry.Name()
		dest := filepath.Join(outDir, filename)

		content, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("error while opening %s: %w", filename, err)
		}

		fOut, err := os.OpenFile(dest, os.O_TRUNC|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640)
		if err != nil {
			return fmt.Errorf("error while generating %s: %w", filename, err)
		}

		content = g.expandAllTags(content)

		for _, include := range g.includes {
			notation := fmt.Sprintf("{{ include:%s }}", include.name)
			content = []byte(strings.ReplaceAll(string(content), notation, string(include.fileContent)))
		}

		if _, err := fOut.Write(content); err != nil {
			return fmt.Errorf("error while generating index.html: %w", err)
		}
		fOut.Close()
	}

	return nil
}

// generatePosts renders each post using its template and writes the result to out/posts/.
// Template variables {{ title }}, {{ date }}, {{ excerpt }}, and {{ content }} are replaced.
// The {{ content }} variable requires the post content to be converted from Markdown to HTML.
func (g *Generator) generatePosts() error {
	destDir := filepath.Join(outDir, postsDir)

	err := os.Mkdir(destDir, 0740)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("error while creating %s: %w", destDir, err)
	}

	for _, p := range g.posts {
		template, err := os.ReadFile(filepath.Join(templateDir, p.Metadata.Template+".html"))
		templateStr := string(template)
		if err != nil {
			return fmt.Errorf("error while opening template %s: %w", p.Metadata.Template, err)
		}

		outFileName := strings.Replace(p.Path, ".md", ".html", -1)
		dest := filepath.Join(outDir, outFileName)

		fOut, err := os.OpenFile(dest, os.O_TRUNC|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640)
		if err != nil {
			return fmt.Errorf("error while generating new post: %w", err)
		}

		if strings.Contains(templateStr, titleTag) {
			templateStr = strings.ReplaceAll(templateStr, titleTag, p.Metadata.Title)
		}

		if strings.Contains(templateStr, dateTag) {
			templateStr = strings.ReplaceAll(templateStr, dateTag, p.Metadata.Date.Format(time.DateOnly))
		}

		if strings.Contains(templateStr, excerptTag) {
			templateStr = strings.ReplaceAll(templateStr, excerptTag, p.Metadata.Excerpt)
		}

		for _, include := range g.includes {
			notation := fmt.Sprintf("{{ include:%s }}", include.name)
			templateStr = strings.ReplaceAll(templateStr, notation, string(include.fileContent))
		}

		templateStr = string(g.expandAllTags([]byte(templateStr)))

		if strings.Contains(templateStr, contentTag) {
			content, err := p.GetHTMLContent()
			if err != nil {
				return fmt.Errorf("error while generating post %s: %w", p.Path, err)
			}

			templateStr = strings.Replace(templateStr, contentTag, string(content), -1)
		}

		if _, err := fOut.WriteString(templateStr); err != nil {
			return fmt.Errorf("error while generating %s: %w", p.Path, err)
		}

		fOut.Close()
	}

	return nil
}

// getAllPosts recursively scans the posts/ directory, parses each Markdown file,
// and collects them into the Generator. It also builds a tag index by collecting
// all posts that share the same tags. Posts are sorted by date (newest first).
func (g *Generator) getAllPosts(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var posts []post.Post
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			err := g.getAllPosts(path)
			if err != nil {
				return fmt.Errorf("error while getting posts in posts directory: %w", err)
			}

			if err := os.Mkdir("out/"+path, 0740); err != nil && !os.IsExist(err) {
				return fmt.Errorf("error while creating post directory: %w", err)
			}

			continue
		}

		p, err := post.New(path)
		if err != nil {
			return err
		}

		for _, t := range p.Metadata.Tags {
			if !slices.ContainsFunc(g.tags, func(a post.Tag) bool {
				return a.Name == t
			}) {
				newTag := post.Tag{Name: t, Posts: []*post.Post{&p}}
				g.tags = append(g.tags, newTag)
				continue
			}

			for i, item := range g.tags {
				if item.Name == t {
					g.tags[i].Posts = append(item.Posts, &p)
					break
				}
			}
		}

		posts = append(posts, p)
	}

	for i := range g.tags {
		slices.SortFunc(g.tags[i].Posts, func(a, b *post.Post) int {
			aDate := a.Metadata.Date
			bDate := b.Metadata.Date

			if aDate.Before(bDate) {
				return 1
			} else if bDate.Before(aDate) {
				return -1
			}

			return 0
		})
	}

	g.posts = append(g.posts, posts...)
	return nil
}

// previewBuilder generates HTML for a tag preview section containing all posts with that tag.
// The output is a div with class "{tag}-previews" containing individual preview divs.
// Each preview has links to the post and displays title + excerpt.
func (g *Generator) previewBuilder(tag string) string {
	builder := strings.Builder{}

	for _, t := range g.tags {
		if tag == t.Name {
			builder.WriteString("<div class=\"" + t.Name + "-previews\">")

			for _, p := range t.Posts {
				builder.WriteString("<div class=\"preview\">\n")
				builder.WriteString("<a href=\"" + p.GetPostLink() + "\">")
				builder.WriteString("<p>" + p.Metadata.Title + "</p>")
				builder.WriteString("</a>\n")
				builder.WriteString("<i>" + p.Metadata.Excerpt + "</i>\n")
				builder.WriteString("</div>\n")
			}

			builder.WriteString("</div>\n")
		}
	}

	return builder.String()
}

// getAllIncludesTags loads all files from the includes/ directory and
// pre-expands {{ tagname }} notations in them, so that includes can display
// tag previews when inserted into templates or root HTML files.
// Projects without an includes/ directory are supported for backward
// compatibility and are treated as having no includes.
// Note: {{ include:filename }} inside an include file is not expanded -
// nested includes are not supported.
func (g *Generator) getAllIncludesTags() error {
	entries, err := os.ReadDir(includesDir)
	if err != nil {
		// Backward compatibility for projects that doesn't have includes/
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	var t []Include
	for _, entry := range entries {
		if !entry.IsDir() {
			content, err := os.ReadFile(filepath.Join(includesDir, entry.Name()))
			if err != nil {
				return err
			}

			content = g.expandAllTags(content)

			t = append(t, Include{
				name:        entry.Name(),
				fileContent: content,
			})
		}
	}

	g.includes = t

	return nil
}

// Expands all tags in the given content bytes and returns the updated bytes
// content.
func (g *Generator) expandAllTags(content []byte) []byte {
	for _, tag := range g.tags {
		notation := fmt.Sprintf("{{ %s }}", tag.Name)
		content = []byte(strings.ReplaceAll(string(content), notation, g.previewBuilder(tag.Name)))
	}

	return content
}
