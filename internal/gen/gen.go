package gen

import (
	"bufio"
	"fmt"
	"guten/internal/post"
	"os"
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
	if err := os.Mkdir(outDir, 0740); err != nil && !os.IsExist(err) {
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

	err = generator.copyIndex()
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
	assetsDest := fmt.Sprintf("%s/%s", outDir, assetsDir)
	if err := os.Mkdir(assetsDest, 0740); err != nil && !os.IsExist(err) {
		return err
	}

	if err := os.CopyFS(assetsDest, os.DirFS(assetsDir)); err != nil && !os.IsExist(err) {
		return fmt.Errorf("error while copying assets directory: %w", err)
	}

	return nil
}

func (g *Generator) copyIndex() error {
	dest := fmt.Sprintf("%s/%s", outDir, indexFile)

	fIn, err := os.Open(indexFile)
	if err != nil {
		return fmt.Errorf("error while opening index.html: %w", err)
	}
	defer fIn.Close()

	fOut, err := os.OpenFile(dest, os.O_TRUNC|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640)
	if err != nil {
		return fmt.Errorf("error while generating index.html: %w", err)
	}
	defer fOut.Close()

	scanner := bufio.NewScanner(fIn)
	for scanner.Scan() {
		line := scanner.Text()

		for _, t := range g.tags {
			notation := fmt.Sprintf("{{ %s }}", t.Name)
			if strings.Contains(line, notation) {
				line = strings.Replace(line, notation, g.previewBuilder(t.Name), -1)
			}
		}

		if _, err := fOut.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("error while generating index.html: %w", err)
		}
	}

	return nil
}

func (g *Generator) generatePosts() error {
	destDir := fmt.Sprintf("%s/%s", outDir, postsDir)

	err := os.Mkdir(destDir, 0740)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("error while creating %s: %w", destDir, err)
	}

	for _, p := range g.posts {
		template, err := os.Open(fmt.Sprintf("%s/%s.html", templateDir, p.Metadata.Template))
		if err != nil {
			return fmt.Errorf("error while opening template %s: %w", p.Metadata.Template, err)
		}
		defer template.Close()

		outFileName := strings.Replace(p.Path, ".md", ".html", -1)
		dest := fmt.Sprintf("out/%s", outFileName)

		fOut, err := os.OpenFile(dest, os.O_TRUNC|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640)
		if err != nil {
			return fmt.Errorf("error while generating new post: %w", err)
		}
		defer fOut.Close()

		templateScanner := bufio.NewScanner(template)

		for templateScanner.Scan() {
			line := templateScanner.Text()

			if strings.Contains(line, titleTag) {
				line = strings.Replace(line, titleTag, p.Metadata.Title, -1)
			}

			if strings.Contains(line, dateTag) {
				line = strings.Replace(line, dateTag, p.Metadata.Date.Format(time.DateOnly), -1)
			}

			if strings.Contains(line, excerptTag) {
				line = strings.Replace(line, excerptTag, p.Metadata.Excerpt, -1)
			}

			if _, err := fOut.WriteString(line + "\n"); err != nil {
				return fmt.Errorf("error while generating %s: %w", p.Path, err)
			}
		}
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
		path := fmt.Sprintf("%s/%s", dir, entry.Name())
		if entry.IsDir() {
			err := g.getAllPosts(path)
			if err != nil {
				return fmt.Errorf("error while getting posts in posts directory: %w", err)
			}
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

			for _, item := range g.tags {
				if item.Name == t {
					item.Posts = append(item.Posts, &p)
					break
				}
			}
		}

		posts = append(posts, p)
	}

	slices.SortFunc(posts, func(a, b post.Post) int {
		aDate := a.Metadata.Date
		bDate := b.Metadata.Date

		if aDate.Before(bDate) {
			return -1
		} else if bDate.Before(aDate) {
			return 1
		}

		return 0
	})

	g.posts = append(g.posts, posts...)
	return nil
}

func (g *Generator) previewBuilder(tag string) string {
	builder := strings.Builder{}

	for _, t := range g.tags {
		if tag == t.Name {
			builder.WriteString("<div class=\"bg-red\">\n")

			for _, p := range t.Posts {
				builder.WriteString("<p>" + p.Metadata.Title + "</p>\n")
				builder.WriteString("<i>" + p.Metadata.Excerpt + "</i>\n")
			}

			builder.WriteString("</div>\n")
		}
	}

	return builder.String()
}
