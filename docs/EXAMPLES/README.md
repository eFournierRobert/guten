# Example Site

This directory contains a working example site that shows how to structure your guten project.

## Structure

```
EXAMPLES/
├── index.html           # Homepage with tag previews
├── assets/
│   └── css/
│       └── style.css    # Basic styles
├── posts/
│   ├── hello-world.md   # First example post
│   └── a-second-post.md # Second example post
└── templates/
    └── post.html        # Template for blog posts
```

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
mkdir -p assets/css    # init only creates the empty assets/ directory
cp ../docs/EXAMPLES/index.html .
cp ../docs/EXAMPLES/assets/css/style.css assets/css/
cp ../docs/EXAMPLES/posts/*.md posts/
cp ../docs/EXAMPLES/templates/*.html templates/

# Run guten
guten -build
# or
guten -peek
```

## Notes

- The `index.html` file uses `{{ blog }}` to show all posts tagged with `blog`
- Posts are in `posts/` directory with `.md` extension
- Templates are in `templates/` directory with naming convention `{name}.html`