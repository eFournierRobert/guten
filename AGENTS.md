# guten - Static Site Generator

## Quick Start

```bash
# Init a new project directory
guten -init my-site

# Create a new post file with default metadata (run from your project directory)
guten -new-post posts/hello-world.md

# Build the website (run from your project directory)
guten -build

# Serve the website (builds and starts HTTP server on :5000)
guten -peek
```

## Architecture

- `main.go` - Entry point with `-init`, `-new-post`, `-build` and `-peek` flags
- `internal/init_project/init.go` - Initializes a directory into a new guten project (`-init` flag)
- `internal/gen/gen.go` - Website generator (parses posts, applies templates, expands `{{ include:filename }}` notations)
- `internal/post/post.go` - Post parsing with YAML frontmatter, Markdown rendering, and new post creation (`-new-post` flag)
- `internal/serve/serve.go` - HTTP server (calls gen, then starts server)
- `internal/server/server.go` - Simple HTTP static file server on port 5000

## Project Structure

**Required files (all scaffolded by `guten -init`):**
- `index.html` in project root - copied to `out/index.html`
- `assets/` directory - copied to `out/assets/`
- `posts/` directory with `.md` files - processed and output to `out/posts/`
- `templates/` directory with `.html` files - referenced by post templates
- `includes/` directory (optional) with `.html` files - reusable snippets inlined where referenced; projects without it still build

**Directory structure:**

```
my-site/
├── index.html          # Required - Homepage
├── assets/             # Required - Static files
│   └── css/, js/, img/ # Your assets
├── posts/              # Markdown posts directory
│   └── *.md            # Post files with YAML frontmatter
├── includes/           # Optional - Reusable HTML includes
│   └── *.html          # Referenced with {{ include:filename }}
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
- **Tags also work**: `{{ tagname }}` is replaced with a tag preview (same as in root HTML files)

## Includes (optional)

Files in `includes/` are reusable HTML snippets referenced with `{{ include:filename }}` in templates and root `.html` files. Includes are **not** copied to `out/` - they are inlined where referenced.

Rules:
- Matched by full file name: `{{ include:footer.html }}` → `includes/footer.html`
- Only files directly in `includes/` are used (subdirectories are skipped)
- Missing `includes/` directory is fine (backward compatibility for existing projects)
- Includes are pre-processed when loaded: `{{ tagname }}` notations inside them ARE expanded (tag previews), even when the include ends up in a post template
- Includes are inserted after all post variables and tag notations are expanded, so they are opaque: nothing inside an include is ever expanded further
- Post variables are NOT expanded inside includes: `{{ title }}`, `{{ date }}`, `{{ excerpt }}`, `{{ content }}` land in the output as literal strings
- Nested includes are NOT supported: `{{ include:... }}` inside an include file is not expanded
- Tag notations inside post Markdown content are never expanded - tags work in templates and root HTML files only. Include notations inside post content are expanded by the include pass at the end of post rendering
- Non-matching notations are left as-is in the output (no build error)

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