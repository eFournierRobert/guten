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

# Create your posts and templates (see docs/QUICKSTART.md)

# Build or serve
guten -build   # Build to out/
guten -peek    # Build and serve on http://localhost:5000
```

`guten -init` creates a basic `index.html` plus the `assets/`, `posts/` and `templates/` directories. The `posts/` and `templates/` directories start empty and the generated `index.html` is a placeholder - you still need to add your own content before building. You can also initialize the current directory with `guten -init .`.

See [`docs/QUICKSTART.md`](./docs/QUICKSTART.md) for a step-by-step guide.

## Commands

| Flag | What it does |
|------|-------------|
| `-init <dir>` | Scaffolds a new project in `dir`: creates `index.html`, `assets/`, `posts/` and `templates/`. The directory must not already exist, except `guten -init .` which initializes the current directory |
| `-build` | Generates the static site in `out/`. The output directory is completely cleared and rebuilt on each run |
| `-peek` | Builds the site (same as `-build`), then serves `out/` over HTTP at `http://localhost:5000` (Ctrl+C to stop) |

## Documentation

- **[Quick Start](./docs/QUICKSTART.md)** - Get your first site built in 5 minutes
- **[Project Structure](./docs/STRUCTURE.md)** - Required directories and files, post frontmatter, output behavior
- **[Templates Reference](./docs/TEMPLATES.md)** - Template variables, tag previews, reserved tag names
- **[Troubleshooting](./docs/TROUBLESHOOTING.md)** - Common errors and how to fix them
- **[Example Site](./docs/EXAMPLES/)** - Working example posts, template and assets
