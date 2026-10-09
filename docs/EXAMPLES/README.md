# Example Site

This directory contains a working example site that demonstrates guten's features, with a focus on **includes**: shared navigation and footer on every page, and tag previews that expand inside include files.

## Structure

```
EXAMPLES/
├── index.html                        # Homepage - direct tag previews + navbar/footer includes
├── about.html                        # Plain root page - navbar/footer includes, no template needed
├── assets/
│   └── css/
│       └── style.css                 # Site styles
├── includes/
│   ├── navbar.html                   # Shared navigation bar
│   ├── footer.html                   # Shared footer
│   └── tag-cloud.html                # Include whose own tag notations expand (article pages)
├── posts/
│   ├── hello-world.md                # First example post
│   ├── a-second-post.md              # Second example post
│   └── demo-tag-in-template.md       # Demo showing tags work in templates
├── templates/
│   ├── post.html                     # Blog post template - uses includes + tag expansion
│   └── project.html                  # Project showcase template
└── out/                              # Generated site (if you run guten -build)
```

## Featured Include Patterns

Every page of this site - homepage, about page and generated posts - shares the same building blocks:

1. **Shared chrome across page types** - `navbar.html` and `footer.html` are included from root HTML files (`index.html`, `about.html`) *and* from post templates (`post.html`, `project.html`). You define them once and they appear everywhere, without CSS frameworks or JavaScript.

2. **Tag previews inside an include** - `tag-cloud.html` contains tag notations for the `demo` and `features` tags. Tag notations inside include files are expanded when the include is loaded, before it is inserted anywhere, so this one snippet shows live, clickable post previews wherever it is used. It is included only from `templates/post.html`: the homepage already lists posts, and a "More from this site" block belongs at the end of an article, where readers are looking for something else to read.

3. **Tag notations directly in a template** - `templates/post.html` puts a `blog` tag notation in its sidebar, showing that tag notations also expand directly in post templates, not just in includes and root HTML files.

### What does not work inside includes

- **Post variables**: `{{ title }}`, `{{ date }}`, `{{ excerpt }}` and `{{ content }}` are never expanded inside includes - they land in the output as literal strings.
- **Nested includes**: an `{{ include:... }}` notation inside an include file is not expanded.
- **Tag notations in post content**: tag notations inside your Markdown are never expanded; they only work in templates and root HTML files. Include notations in post content are expanded by the include pass at the end of post rendering.

See [TEMPLATES.md](../TEMPLATES.md#includes) for the full includes reference.

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
cp ../docs/EXAMPLES/index.html ../docs/EXAMPLES/about.html .
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

- `index.html` and `about.html` are root HTML files: guten expands tag notations and includes in every root `.html` file, not just `index.html`
- `templates/post.html` uses all three includes and demonstrates tag expansion in its sidebar
- `includes/tag-cloud.html` demonstrates tag notations expanding inside an include file
- Posts are in `posts/` directory with `.md` extension
- Templates are in `templates/` directory with naming convention `{name}.html`
- `posts/` can contain subdirectories - structure is preserved in output
