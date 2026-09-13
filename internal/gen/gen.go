package gen

import (
	"fmt"
	"guten/internal/post"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

const outDir = "out"
const assetsDir = "assets"
const postsDir = "posts"
const indexFile = "index.html"
const templateDir = "templates"

const titleTag = "{{ title }}"
const contentTag = "{{ content }}"
const dateTag = "{{ date }}"
const excerptTag = "{{ excerpt }}"

type Generator struct {
	posts []post.Post
	tags  []post.Tag
}

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

		for _, t := range g.tags {
			notation := fmt.Sprintf("{{ %s }}", t.Name)
			content = []byte(strings.ReplaceAll(string(content), notation, g.previewBuilder(t.Name)))
		}

		if _, err := fOut.Write(content); err != nil {
			return fmt.Errorf("error while generating index.html: %w", err)
		}
		fOut.Close()
	}

	return nil
}

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
			templateStr = strings.Replace(templateStr, titleTag, p.Metadata.Title, -1)
		}

		if strings.Contains(templateStr, dateTag) {
			templateStr = strings.Replace(templateStr, dateTag, p.Metadata.Date.Format(time.DateOnly), -1)
		}

		if strings.Contains(templateStr, excerptTag) {
			templateStr = strings.Replace(templateStr, excerptTag, p.Metadata.Excerpt, -1)
		}

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

	for i, _ := range g.tags {
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
