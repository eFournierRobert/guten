# Project Structure Guide

This document explains the file and directory structure required for guten to work properly.

> **Tip:** `guten -init <dir>` scaffolds the base structure for you (required directories plus a basic `index.html`). You still need to add your templates, and your posts - create each starter with `guten -new-post <path>` or write them by hand.

## Required Structure

Every guten project must have the following structure:

```
my-site/
├── index.html          # ✅ REQUIRED - Root HTML file
├── assets/             # ✅ REQUIRED - Static files directory
│   └── ...             # Any CSS, JS, images, etc.
├── posts/              # ✅ REQUIRED - Post directory
│   └── *.md            # Markdown post files
├── includes/           # Optional - Reusable HTML includes
│   └── *.html          # Referenced with {{ include:filename }}
└── templates/          # ✅ REQUIRED - Template directory
    └── *.html          # HTML template files
```

## Directory Descriptions

### `index.html` (Required)

- **Type:** File
- **Extension:** `.html`
- **Purpose:** The main entry point for your site
- **Behavior:** Copied directly to `out/index.html` during build

In addition to static content, `index.html` supports special tag syntax that generates preview sections for posts - see [TEMPLATES.md](./TEMPLATES.md) for details and examples.

### `assets/` (Required)

- **Type:** Directory
- **Purpose:** Container for all static files
- **Behavior:** Entire directory is copied to `out/assets/`

The assets directory preserves its structure:

```
assets/
├── css/
│   └── style.css
├── js/
│   └── main.js
└── img/
    └── logo.png
```

After build:
```
out/
└── assets/
    ├── css/style.css
    ├── js/main.js
    └── img/logo.png
```

### `posts/` (Required)

- **Type:** Directory
- **Purpose:** Contains Markdown posts with YAML frontmatter
- **Supported extensions:** `.md`, `.markdown`
- **Supports:** Subdirectories (structure is preserved in output)

Each post file must have this structure:

```markdown
---                           # Opening frontmatter delimiter
date: 2024-01-15              # Post date
title: Post Title            # Post title
excerpt: Summary text        # Short excerpt for previews
template: post               # Template to use (without .html)
tags:                        # Array of tags
  - tag1
  - tag2
---                           # Closing frontmatter delimiter

# Markdown content starts here
```

### Subdirectories in posts/

You can organize posts into subdirectories:

```
posts/
├── 2024/
│   ├── 01-january.md
│   └── 02-february.md
├── tutorials/
│   ├── getting-started.md
│   └── advanced-tips.md
└── news.md
```

After build, the output structure is preserved:

```
out/
└── posts/
    ├── 2024/
    │   ├── 01-january.html
    │   └── 02-february.html
    ├── tutorials/
    │   ├── getting-started.html
    │   └── advanced-tips.html
    └── news.html
```

### `templates/` (Required)

- **Type:** Directory
- **Purpose:** Contains HTML template files for rendering posts
- **Behavior:** Templates are read during build to generate HTML pages

Template naming convention:

| Template File | Referenced by | Use Case |
|---------------|---------------|----------|
| `post.html` | `template: post` | Blog posts |
| `page.html` | `template: page` | Static pages |
| `project.html` | `template: project` | Project showcases |

### `includes/` (Optional)

- **Type:** Directory
- **Purpose:** Contains reusable HTML snippets (headers, footers, navigation, etc.)
- **Behavior:** Files are *not* copied to `out/` - they are inlined into pages during build

Reference an include from any root HTML file or template with the `{{ include:filename }}` notation:

```html
{{ include:footer.html }}   <!-- Replaced with the contents of includes/footer.html -->
```

Tag notations (`{{ tagname }}`) inside include files are always expanded into tag previews when the include is loaded - same in root HTML files and post templates. Post variables (`{{ title }}`, `{{ date }}`, `{{ excerpt }}`, `{{ content }}`) and nested includes (`{{ include:... }}`) are **not** supported inside include files - see [TEMPLATES.md](./TEMPLATES.md) for details. Projects without an `includes/` directory build fine; `guten -init` creates the directory but you can ignore it.

### `out/` (Generated)

- **Type:** Directory (created automatically)
- **Purpose:** Contains the generated static site
- **Behavior:** **Completely cleared and regenerated on each build**

Generated structure:

```
out/
├── index.html              # From root index.html
├── assets/                 # Copied from assets/
│   └── ...                 # Everything from assets/
└── posts/                  # Generated post pages
    └── *.html              # One per post file
```

## Complete Example

Here's a complete example project structure:

```
my-blog/
├── index.html                          # Homepage
│                                       # (shows tag previews)
├── assets/                             # Static files
│   ├── css/
│   │   └── style.css
│   ├── js/
│   │   └── main.js
│   └── img/
│       └── favicon.png
├── posts/                              # Blog posts
│   ├── 2024-01-15-hello-world.md
│   ├── 2024-02-20-second-post.md
│   └── tutorials/
│       └── making-a-template.md
├── includes/                           # Reusable HTML includes (optional)
│   ├── navbar.html
│   └── footer.html
├── templates/                          # HTML templates
│   ├── post.html
│   └── page.html
└── out/                                # Generated site (not versioned)
    ├── index.html
    ├── assets/
    │   ├── css/style.css
    │   ├── js/main.js
    │   └── img/favicon.png
    └── posts/
        ├── 2024-01-15-hello-world.html
        ├── 2024-02-20-second-post.html
        └── tutorials/
            └── making-a-template.html
```

## Common Mistakes

1. **Missing index.html**: If you don't have `index.html` in the root, nothing will be copied to `out/` as your homepage.

2. **Missing assets/**: If `assets/` doesn't exist, guten will create the directory in `out/` but it will be empty.

3. **Wrong template name**: The `template` field in a post must match a template file name (without `.html` extension).

4. **Reserved tag names**: Don't use `title`, `content`, `date`, or `excerpt` as custom tags.

5. **Typo in include name**: If `{{ include:foo.html }}` doesn't match a file in `includes/`, the notation is left as-is in the output. Check the file name in your `includes/` directory (exact spelling, including the `.html` extension).