# Example Site

This directory contains a working example site that demonstrates guten's features, including **includes**, **tag expansions in templates**, and **multiple templates**.

## Structure

```
EXAMPLES/
├── index.html                        # Homepage - uses tag previews & navbar include
├── assets/
│   └── css/
│       └── style.css                 # Site styles
├── includes/
│   ├── navbar.html                   # Shared navigation bar
│   └── footer.html                   # Shared footer, included via {{ include:footer.html }}
├── posts/
│   ├── hello-world.md                 # First example post
│   ├── a-second-post.md               # Second example post
│   └── demo-tag-in-template.md        # Demo showing tags work in templates
├── templates/
│   ├── post.html                     # Blog post template - uses includes + tag expansion
│   └── project.html                  # Project showcase template
└── out/                              # Generated site (if you run guten -build)
```

## Featured Include Patterns

This example showcases the `{{ include:filename }}` feature:

- **`navbar.html`** - A navigation bar used in BOTH root HTML files and post templates. This demonstrates cross-context reuse.
- **`footer.html`** - A footer included at the end of every page.

**Why this matters:** You define the navigation and footer once, and they appear consistently across your entire site, whether in the homepage, blog posts, or other pages - without CSS frameworks or JavaScript!

## Featured Tag-in-Template Pattern

The `post.html` template includes a sidebar with `{{ blog }}` to show all blog posts. This demonstrates that **tag expansions work directly in post templates** (added in recent versions), not just in root HTML files like `index.html`.

See [TEMPLATES.md](../TEMPLATES.md) for full details on `{{ tagname }}` tag expansions.

## Try the Example

You can use these files as a template for your own site:

```bash
# Copy the example files to a new directory
cp -r docs/EXAMPLES/* /path/to/your/site/
```

Or scaffold a new project with `guten -init` and copy just the content:

```bash
# Scaffold the base structure (directories + index.html)
guten -init my-site
cd my-site

# Copy the example content
cp ../docs/EXAMPLES/index.html .
cp ../docs/EXAMPLES/assets/css/style.css assets/css/
cp ../docs/EXAMPLES/includes/*.html includes/
cp ../docs/EXAMPLES/posts/*.md posts/
cp ../docs/EXAMPLES/templates/*.html templates/

# Run guten
guten -build
# or
guten -peek
```

## Notes

- The `index.html` file uses `{{ blog }}` to show posts tagged with `blog`
- Both `index.html`, `templates/post.html`, and `templates/project.html` use `{{ include:navbar.html }}` and `{{ include:footer.html }}` for shared layout
- `templates/post.html` also demonstrates `{{ blog }}` tag expansion in a template sidebar
- Posts are in `posts/` directory with `.md` extension
- Templates are in `templates/` directory with naming convention `{name}.html`
- `posts/` can contain subdirectories - structure is preserved in output