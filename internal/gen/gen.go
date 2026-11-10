// Package gen provides the website generator for guten.
// It parses Markdown posts with YAML frontmatter, applies HTML templates,
// and generates a static site in the out/ directory.
// It reads posts/, templates/, assets/, and root HTML files from the working
// directory. Optional includes/ files provide {{ include:filename }} snippets;
// optional previews/<tag>.html files customize each post in a tag preview list.
package gen

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
const previewsDir = "previews"

// Template variable tags available in template files.
const titleTag = "{{ title }}"
const contentTag = "{{ content }}"
const dateTag = "{{ date }}"
const excerptTag = "{{ excerpt }}"

// ExtraFile holds an include or custom preview loaded from its source directory.
// name is the full base filename, including its extension; fileContent has tag
// notations expanded at load time, before per-post preview variables are filled.
type ExtraFile struct {
	name        string
	fileContent []byte
}

// Generator handles the static site generation process.
// It collects posts, builds tag indexes, loads includes and custom previews,
// and generates output files. Posts within each tag are sorted newest first.
type Generator struct {
	posts    []post.Post
	tags     []post.Tag
	includes []ExtraFile
	previews []ExtraFile
}

// GenerateWebsite builds a static site from posts and templates.
// It clears out/, copies assets/, processes posts, and renders root HTML files.
// Paths are relative to the working directory. The assets/ and posts/
// directories and each post's templates/<name>.html must exist; includes/ and
// previews/ are optional. Source includes, previews, and templates are not copied.
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

	// Custom previews must be available before root HTML and posts are rendered.
	if err := generator.getAllPreviews(); err != nil {
		return err
	}

	// Includes must be collected after posts so that tag notations inside
	// include files ({{ tagname }}) can be expanded with the built tag index.
	if err := generator.getAllIncludes(); err != nil {
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
// Matching custom previews are used when available. Includes are inserted after
// tag expansion; unknown tag and include notations remain unchanged.
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
// Tag previews expand after title, date, and excerpt substitution but before
// Markdown content is inserted, so tag notations in Markdown remain literal.
// Includes are inserted last, including references within rendered Markdown.
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

		templateStr = string(g.expandAllTags([]byte(templateStr)))

		if strings.Contains(templateStr, contentTag) {
			content, err := p.GetHTMLContent()
			if err != nil {
				return fmt.Errorf("error while generating post %s: %w", p.Path, err)
			}

			templateStr = strings.Replace(templateStr, contentTag, string(content), -1)
		}

		for _, include := range g.includes {
			notation := fmt.Sprintf("{{ include:%s }}", include.name)
			templateStr = strings.ReplaceAll(templateStr, notation, string(include.fileContent))
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
// all posts that share the same tags. Each tag's posts are sorted by date
// (newest first); g.posts retains traversal order.
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

// previewBuilder renders a tag's posts in their indexed order (newest first)
// inside a div with class "<tag>-previews". Unknown tags produce an empty string.
// A loaded previews/<tag>.html supplies the markup repeated for each post;
// otherwise each post uses the default preview.
// Custom markup substitutes {{ title }}, {{ date }} (YYYY-MM-DD), {{ excerpt }},
// and {{ link }} (the root-relative HTML post URL).
func (g *Generator) previewBuilder(tag string) string {
	var preview *ExtraFile
	if g.previews != nil {
		for _, p := range g.previews {
			if p.name == tag+".html" {
				preview = &p
				break
			}
		}
	}

	builder := strings.Builder{}

	for _, t := range g.tags {
		if tag == t.Name {
			builder.WriteString("<div class=\"" + t.Name + "-previews\">")

			for _, p := range t.Posts {
				if preview == nil {
					builder.WriteString("<div class=\"preview\">\n")
					builder.WriteString("<a href=\"/" + p.GetPostLink() + "\">")
					builder.WriteString("<p>" + p.Metadata.Title + "</p>")
					builder.WriteString("</a>\n")
					builder.WriteString("<i>" + p.Metadata.Excerpt + "</i>\n")
					builder.WriteString("</div>\n")
				} else {
					fileContent := string(preview.fileContent)

					fullContent := strings.ReplaceAll(fileContent, titleTag, p.Metadata.Title)
					fullContent = strings.ReplaceAll(fullContent, dateTag, p.Metadata.Date.Format(time.DateOnly))
					fullContent = strings.ReplaceAll(fullContent, excerptTag, p.Metadata.Excerpt)
					fullContent = strings.ReplaceAll(fullContent, "{{ link }}", "/"+p.GetPostLink())
					builder.WriteString(fullContent)
					builder.WriteString("\n")
				}
			}

			builder.WriteString("</div>\n")
		}
	}

	return builder.String()
}

// getAllIncludes loads optional includes/ snippets and expands their tag lists
// using the loaded previews.
func (g *Generator) getAllIncludes() error {
	extra, err := g.getAllExtraFile(includesDir)
	if err != nil {
		return err
	}

	g.includes = extra
	return nil
}

// getAllPreviews loads optional previews/ files for lookup by <tag>.html.
// Tag notations are expanded during loading; per-post variables are substituted
// later by previewBuilder. The new previews are assigned only after all files
// are loaded, so they do not reference one another during this loading pass.
func (g *Generator) getAllPreviews() error {
	extra, err := g.getAllExtraFile(previewsDir)
	if err != nil {
		return err
	}

	g.previews = extra
	return nil
}

// getAllExtraFile reads files directly in directory, in filename order, and
// pre-expands tag notations using the current tag index and loaded previews.
// It skips subdirectories, accepts any file extension, and preserves base names
// for exact matching. A missing directory returns nil, nil; other read errors
// are returned to the caller. Post variables and includes are left unprocessed.
func (g *Generator) getAllExtraFile(directory string) ([]ExtraFile, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		// Includes and custom previews are optional for existing projects.
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	var t []ExtraFile
	for _, entry := range entries {
		if !entry.IsDir() {
			content, err := os.ReadFile(filepath.Join(directory, entry.Name()))
			if err != nil {
				return nil, err
			}

			content = g.expandAllTags(content)

			t = append(t, ExtraFile{
				name:        entry.Name(),
				fileContent: content,
			})
		}
	}

	return t, nil
}

// expandAllTags replaces exact {{ tagname }} notations for known tags with
// previewBuilder output, leaving unknown notations unchanged. Tags are processed
// once in index order; this is not a recursive template expansion.
func (g *Generator) expandAllTags(content []byte) []byte {
	for _, tag := range g.tags {
		notation := fmt.Sprintf("{{ %s }}", tag.Name)
		content = []byte(strings.ReplaceAll(string(content), notation, g.previewBuilder(tag.Name)))
	}

	return content
}
