# Project Structure Guide

This document explains the file and directory structure required for guten to work properly.

## Required Structure

Every guten project must have the following structure:

```
my-site/
├── index.html          # ✅ REQUIRED - Root HTML file
├── assets/             # ✅ REQUIRED - Static files directory
│   └── ...             # Any CSS, JS, images, etc.
├── posts/              # ✅ REQUIRED - Post directory
│   └── *.md            # Markdown post files
└── templates/          # ✅ REQUIRED - Template directory
    └── *.html          # HTML template files
```

## Directory Descriptions

### `index.html` (Required)

- **Type:** File
- **Extension:** `.html`
- **Purpose:** The main entry point for your site
- **Behavior:** Copied directly to `out/index.html` during build

You can use special tag syntax in `index.html` to generate tag previews:

```html
<!DOCTYPE html>
<html>
<body>
    <h1>My Blog</h1>
    {{ blog }}      <!-- Posts tagged 'blog' -->
    {{ news }}      <!-- Posts tagged 'news' -->
</body>
</html>
```

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

## Optional Structure

The following are optional but commonly used:

### Custom asset organization

```
assets/
├── css/
│   ├── style.css
│   └── syntax.css
├── js/
│   ├── main.js
│   └── highlight.js
└── fonts/
    └── ...
```

### Posts with metadata

```
posts/
├── drafts/                 # Work in progress
│   └── draft-post.md
├── published/              # Published content
│   └── published-post.md
└── archive/                # Older content
    └── old-post.md
```

### Multiple template types

```
templates/
├── post.html               # Blog posts
├── page.html               # Static pages
├── project.html            # Projects
├── author.html             # Author pages
└── tag.html                # Tag pages
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