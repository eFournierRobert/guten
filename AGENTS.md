# guten - Static Site Generator

## Quick Start

```bash
# Init a new project directory
guten -init my-site

# Build the website (run from your project directory)
guten -build

# Serve the website (builds and starts HTTP server on :5000)
guten -peek
```

## Architecture

- `main.go` - Entry point with `-init`, `-build` and `-peek` flags
- `internal/init_project/init.go` - Initializes a directory into a new guten project (`-init` flag)
- `internal/gen/gen.go` - Website generator (parses posts, applies templates)
- `internal/post/post.go` - Post parsing with YAML frontmatter and Markdown rendering
- `internal/serve/serve.go` - HTTP server (calls gen, then starts server)
- `internal/server/server.go` - Simple HTTP static file server on port 5000

## Project Structure

**Required files (all scaffolded by `guten -init`):**
- `index.html` in project root - copied to `out/index.html`
- `assets/` directory - copied to `out/assets/`
- `posts/` directory with `.md` files - processed and output to `out/posts/`
- `templates/` directory with `.html` files - referenced by post templates

**Directory structure:**

```
my-site/
├── index.html          # Required - Homepage
├── assets/             # Required - Static files
│   └── css/, js/, img/ # Your assets
├── posts/              # Markdown posts directory
│   └── *.md            # Post files with YAML frontmatter
└── templates/          # HTML templates directory
    └── *.html          # Template files
```

## Post Format

Each post file must have YAML frontmatter:

```yaml
---
date: 2024-01-15
title: My Post
excerpt: Short summary
template: post
tags:
  - tag1
  - tag2
---

# Markdown content starts here
```

## Templates

Templates live in `templates/` directory:
- `templates/post.html` - For posts with `template: post`
- `templates/page.html` - For pages
- `templates/project.html` - For project showcases

Templates support:
- `{{ title }}` - Replaced with post title
- `{{ content }}` - Replaced with rendered HTML content
- `{{ date }}` - Replaced with post date
- `{{ excerpt }}` - Replaced with post excerpt

## Reserved Tags

Tags named `title`, `content`, `date`, or `excerpt` are **reserved** and will cause errors.

Use alternative tag names like `blog`, `news`, `tutorial`, etc.

## Output Structure

After running `guten -build`:

```
out/
├── index.html      # From root index.html
├── assets/         # From assets/
└── posts/          # Generated HTML from posts/
    └── *.html
```

**Note:** The `out/` directory is completely cleared before each build.