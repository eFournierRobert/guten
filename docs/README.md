# guten - Static Site Generator

A small static site generator written in Go. Write posts in Markdown, define layouts with HTML templates, and generate a static site.

## Table of Contents

- [Quick Start](#quick-start)
- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Build & Serve](#build--serve)
- [Posts](#posts)
- [Templates](#templates)
- [Template Variables](#template-variables)
- [Output Structure](#output-structure)
- [Reserved Tags](#reserved-tags)
- [Troubleshooting](#troubleshooting)

## Quick Start

```bash
# Clone and install
git clone https://codeberg.org/efournierrobert/guten.git
cd guten
go install .

# Create a project directory with posts and templates
mkdir my-site && cd my-site
mkdir posts templates assets

# Create your content (see examples/)
# ...

# Build the site
guten -build   # generates to out/

# Or serve locally
guten -peek    # builds and starts HTTP server on :5000
```

## Prerequisites

**Before running guten, you must have:**

1. **An `index.html` file in your project root** - This file will be copied to `out/index.html`
2. **An `assets/` directory** - This will be copied to `out/assets/`
3. **A `posts/` directory with `.md` files** - Markdown posts with YAML frontmatter
4. **A `templates/` directory with `.html` files** - Template files for rendering posts

**Minimum required structure:**

```
my-site/
├── index.html          # Your main HTML page (required)
├── assets/             # Static assets directory (required)
│   └── ...             # CSS, JS, images, etc.
├── posts/              # Markdown posts (required)
│   └── *.md
└── templates/          # HTML templates (required)
    └── *.html
```

## Project Structure

| Directory | Description | Required |
|-----------|-------------|----------|
| `posts/` | Markdown files with YAML frontmatter metadata | ✅ Yes |
| `templates/` | HTML template files for rendering posts | ✅ Yes |
| `assets/` | Static files (CSS, JS, images) copied to output | ✅ Yes |
| `index.html` | Root HTML file copied to output root | ✅ Yes |
| `out/` | Generated static site (created by `-build`) | ❌ No |

### Optional directories:

- **`posts/subdir/`** - Subdirectories in posts are supported and preserved in output
- **`assets/css/`, `assets/js/`, `assets/img/**` - Organize static assets as needed

## Build & Serve

### Build Command

```bash
guten -build
```

Generates a static site in the `out/` directory:
- Copies `index.html` to `out/index.html`
- Copies entire `assets/` to `out/assets/`
- Processes all posts in `posts/` and generates static HTML

### Peek Command

```bash
guten -peek
```

1. Builds the website (same as `-build`)
2. Starts an HTTP server on `http://localhost:5000`
3. Serves the generated `out/` directory

Press `Ctrl+C` to stop the server.

## Posts

### Post File Structure

Each post file must contain YAML frontmatter followed by Markdown content:

```markdown
---
date: 2024-01-15
title: My First Post
excerpt: A brief summary of my post
template: post        # References templates/post.html
tags:
  - blog
  - tutorial
---

This is the **Markdown content** of my post.

It supports:
- Headers
- Lists
- Code blocks
- Links
- etc.
```

### Frontmatter Fields

| Field | Required | Description |
|-------|----------|-------------|
| `date` | ✅ Yes | Post date (YYYY-MM-DD format) |
| `title` | ✅ Yes | Post title |
| `excerpt` | ✅ Yes | Short summary (shown in tag previews) |
| `template` | ✅ Yes | Template name (must exist in `templates/`) |
| `tags` | ✅ Yes | Array of tags for categorization |

### Subdirectory Support

Posts can be organized in subdirectories:

```
posts/
├── 2024/
│   └── tech-article.md
├── blog/
│   ├── first-post.md
│   └── second-post.md
└── news.md
```

Subdirectory structure is preserved in the output.

## Templates

### Template Files

Templates are HTML files with special tags that get replaced with post content:

- `templates/post.html` - For blog posts (referenced via `template: post` in post frontmatter)
- `templates/page.html` - For static pages
- `templates/project.html` - For project showcases

### Creating Templates

Create a template file in the `templates/` directory:

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }}</title>
</head>
<body>
    <h1>{{ title }}</h1>
    <p class="date">{{ date }}</p>
    <p class="excerpt">{{ excerpt }}</p>
    <div class="content">
        {{ content }}
    </div>
</body>
</html>
```

## Template Variables

The following variables can be used inside template files:

| Variable | Description | Source |
|----------|-------------|--------|
| `{{ title }}` | Post title | Post frontmatter |
| `{{ date }}` | Post date (formatted as YYYY-MM-DD) | Post frontmatter |
| `{{ excerpt }}` | Post excerpt/summary | Post frontmatter |
| `{{ content }}` | Rendered HTML content from Markdown | Post body |

### Special Tags (for root HTML files)

In `index.html` and other root HTML files, tags are replaced with dynamic content:

- `{{ tagname }}` - Creates a preview section for all posts with that tag
- Multiple tags can be used: `{{ blog }}`, `{{ news }}`, etc.

Example `index.html`:

```html
<!DOCTYPE html>
<html>
<body>
    <h1>My Blog</h1>
    
    {{ blog }}  <!-- Shows list of all posts tagged with 'blog' -->
    {{ news }}  <!-- Shows list of all posts tagged with 'news' -->
</body>
</html>
```

## Output Structure

After running `guten -build`, the `out/` directory contains:

```
out/
├── index.html          # Copied from root index.html
├── assets/             # Copied from assets/
│   └── ...             # All static files
└── posts/              # Generated post pages
    ├── 2024/           # Preserved subdirectories
    │   └── tech-article.html
    ├── blog/
    │   ├── first-post.html
    │   └── second-post.html
    └── news.html
```

**Important:** The output directory is **completely cleared** before each build to prevent stale files.

## Reserved Tags

⚠️ **Warning:** The following tag names are **reserved** and will cause errors:

- `title`
- `content`
- `date`
- `excerpt`

If you use any of these as post tags, guten will return an error:

```
Error: title is a reserved tag
```

Choose alternative names for your tags (e.g., `blog`, `news`, `tutorial`, `project`).

## Troubleshooting

### Error: "No frontmatter in post"

**Cause:** Your post file doesn't start with `---` on the first line.

**Fix:** Ensure your post starts with YAML frontmatter:

```yaml
---
date: 2024-01-15
title: My Post
excerpt: Summary
template: post
tags:
  - blog
---
```

### Error: "template not found"

**Cause:** The `template` field in your post doesn't match any template file.

**Fix:** Create the template file or use an existing one:

- Post specifies `template: post`
- Need `templates/post.html` file

### Build succeeds but no HTML in output

**Cause:** Missing `index.html` in project root.

**Fix:** Add an `index.html` file to your project root.

### Assets not copied

**Cause:** Missing `assets/` directory.

**Fix:** Create an `assets/` directory (can be empty if you don't need static assets).

### Server won't start on port 5000

**Cause:** Port 5000 is already in use.

**Fix:** Stop the other process using that port, or modify the port in `internal/server/server.go`.