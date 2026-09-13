# guten - Static Site Generator

## Quick Start

```bash
# Build the website (run from test-site/ directory)
cd test-site && go run .. -build

# Serve the website (builds and starts HTTP server on :5000)
cd test-site && go run .. -peak
```

## Architecture

- `main.go` - Entry point with `-build` and `-peak` flags
- `internal/gen/gen.go` - Website generator (parses posts, applies templates)
- `internal/post/post.go` - Post parsing with YAML frontmatter and Markdown rendering
- `internal/serve/serve.go` - HTTP server (calls gen, then starts server)
- `internal/server/server.go` - Simple HTTP static file server on port 5000

## Content Structure

- Posts live in `test-site/posts/` (supports subdirectories)
- Each post needs YAML frontmatter with: `date`, `title`, `excerpt`, `template`, `tags`
- Templates live in `test-site/templates/` (`post.html`, `page.html`, `project.html`)
- Templates support: `{{ title }}`, `{{ content }}`, `{{ date }}`, `{{ excerpt }}`

## Reserved Tags

Tags named `title`, `content`, `date`, or `excerpt` are reserved and will cause errors.