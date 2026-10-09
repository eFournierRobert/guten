# guten

A small static site generator written in Go. Write posts in Markdown, define layouts with HTML templates, and generate a static site.

## Install

```bash
git clone https://github.com/eFournierRobert/guten.git
cd guten
go install .
```

This will compile the binary and put it in `~/go/bin/`.

## Quick Start

```bash
# Create a new project (scaffolds the base structure)
guten -init my-site
cd my-site

# Create a starter post, then edit the template name, title and content
guten -new-post posts/hello-world.md
vim posts/hello-world.md templates/post.html    # see docs/QUICKSTART.md


# Build or serve
guten -build   # Build to out/
guten -peek    # Build and serve on http://localhost:5000
```

`guten -init` creates a basic `index.html` plus the `assets/`, `posts/`, `templates/` and `includes/` directories. The `assets/`, `posts/`, `templates/` and `includes/` directories start empty and the generated `index.html` is a placeholder - you still need to add your own content before building. `includes/` is optional (reusable HTML snippets for `{{ include:filename }}`) and older projects without it build fine. You can also initialize the current directory with `guten -init .`.

See [`docs/QUICKSTART.md`](./docs/QUICKSTART.md) for a step-by-step guide.

## Commands

| Flag | What it does |
|------|-------------|
| `-init <dir>` | Scaffolds a new project in `dir`: creates `index.html`, `assets/`, `posts/`, `templates/` and `includes/`. The directory must not already exist, except `guten -init .` which initializes the current directory |
| `-new-post <path>` | Writes a new starter post: frontmatter with today's date and the `post` template, plus a placeholder quote as content. If the file already exists, it asks before overwriting; you must then replace the placeholder title, excerpt and content |
| `-build` | Generates the static site in `out/`: root `.html` files get tag previews and includes expanded, posts are rendered through their templates. The output directory is completely cleared and rebuilt on each run |
| `-peek` | Builds the site (same as `-build`), then serves `out/` over HTTP at `http://localhost:5000` (Ctrl+C to stop) |

## Documentation

- **[Quick Start](./docs/QUICKSTART.md)** - Get your first site built in 5 minutes
- **[Project Structure](./docs/STRUCTURE.md)** - Required directories and files, post frontmatter, output behavior
- **[Templates Reference](./docs/TEMPLATES.md)** - Template variables, tag previews, reserved tag names
- **[Troubleshooting](./docs/TROUBLESHOOTING.md)** - Common errors and how to fix them
- **[Example Site](./docs/EXAMPLES/)** - Working example site with includes, tag previews and multiple templates
