# Templates Reference

This guide covers everything you need to know about creating and using templates with guten.

## What is a Template?

A template is an HTML file that defines the layout and structure of your generated pages. Post templates contain special tags (like `{{ title }}` and `{{ content }}`) that get replaced with the post's metadata and content during the build. Root HTML files like `index.html` use the same curly brace notation for tag previews instead of post variables. Both contexts can reuse shared markup with `{{ include:filename }}` - see the [Includes](#includes) section.

## Template Files Location

All template files must be placed in the `templates/` directory:

```
templates/
├── post.html        # For blog posts
├── page.html        # For static pages
└── project.html     # For project showcases
```

## Template Tags

### Available Template Variables

| Tag | Description | Example Output |
|-----|-------------|----------------|
| `{{ title }}` | Post title | "My Great Adventure" |
| `{{ date }}` | Post publication date (YYYY-MM-DD) | "2024-01-15" |
| `{{ excerpt }}` | Post excerpt/summary | "A brief description..." |
| `{{ content }}` | Rendered HTML from Markdown body | `<p>Hello <strong>world</strong></p>` |

### Using Multiple Variables

You can combine multiple template tags in a single template:

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }} | My Blog</title>
</head>
<body>
    <article>
        <header>
            <h1>{{ title }}</h1>
            <time datetime="{{ date }}">{{ date }}</time>
        </header>
        
        <div class="excerpt">
            {{ excerpt }}
        </div>
        
        <div class="content">
            {{ content }}
        </div>
    </article>
</body>
</html>
```

### Tags in Root HTML Files

For `index.html` and other root-level HTML files, curly brace tags are replaced with **tag previews** (lists of posts with that tag):

```html
<!-- index.html -->
<html>
<body>
    <h1>My Blog</h1>
    
    {{ tech }}      <!-- Shows all posts tagged 'tech' -->
    {{ cooking }}   <!-- Shows all posts tagged 'cooking' -->
</body>
</html>
```

Each tag preview generates:
```html
<div class="tech-previews">
    <div class="preview">
        <a href="posts/tech-article.html">
            <p>Tech Article Title</p>
        </a>
        <i>A brief excerpt...</i>
    </div>
    <!-- More posts... -->
</div>
```

## Includes

Includes are reusable HTML snippets that let you share markup (headers, footers, navigation, etc.) between templates and root HTML files instead of copy-pasting it.

Files in the `includes/` directory are referenced with the `{{ include:filename }}` notation:

```html
<!-- index.html or templates/post.html -->
<nav>
    {{ include:navbar.html }}   <!-- Inserted from includes/navbar.html -->
</nav>

...

{{ include:footer.html }}       <!-- Inserted from includes/footer.html -->
```

During the build, every `{{ include:filename }}` occurrence is replaced with the full contents of `includes/filename`. If your project has no `includes/` directory, builds work as before - existing projects are unaffected.

### What gets expanded inside an include

- **Tag notations always work**: `{{ tagname }}` inside an include file is expanded into tag previews when the include is loaded (before insertion). **Tag notations also now work directly in post templates** - `generatePosts()` expands them after include insertion. So both of the following produce tag previews:
  ```html
  <!-- In templates/post.html -->
  <aside class="blog-sidebar">{{ blog }}</aside>
  
  <!-- Inside includes/sidebar.html -->
  <div>{{ blog }}</div>
  ```
- **Post variables do not**: `{{ title }}`, `{{ date }}`, `{{ excerpt }}` and `{{ content }}` are not expanded inside includes. They are expanded where the template is rendered, after include insertion, so any `{{ title }}` you put in an include file lands in the output as a literal string. Keep includes free of post variables.
- **Nested includes do not**: an `{{ include:... }}` notation inside an include file is *not* expanded. Compose includes from plain HTML only, not from other includes.

### Example

`includes/footer.html`:

```html
<footer>
    <p>&copy; 2024 My Blog. Built with guten.</p>
</footer>
```

`index.html`:

```html
<body>
    <h1>My Blog</h1>
    {{ blog }}
    {{ include:footer.html }}
</body>
```

The generated `out/index.html` contains the footer markup in place of the include notation.

## Template Reference Examples

### Blog Post Template

`templates/post.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ title }}</title>
    <link rel="stylesheet" href="/assets/css/style.css">
</head>
<body>
    <article class="post">
        <header class="post-header">
            <h1 class="post-title">{{ title }}</h1>
            <time class="post-date">{{ date }}</time>
        </header>
        
        <div class="post-content">
            {{ content }}
        </div>
        
        <footer class="post-footer">
            <p class="post-excerpt">{{ excerpt }}</p>
        </footer>
    </article>
    
    <script src="/assets/js/main.js"></script>
</body>
</html>
```

### Static Page Template

`templates/page.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ title }} - My Site</title>
    <link rel="stylesheet" href="/assets/css/style.css">
</head>
<body>
    <nav>
        <a href="/">Home</a>
        <a href="/about.html">About</a>
    </nav>
    
    <main class="page">
        <h1>{{ title }}</h1>
        <div class="page-content">
            {{ content }}
        </div>
    </main>
    
    <footer>
        <p>{{ excerpt }}</p>
    </footer>
</body>
</html>
```

### Project Showcase Template

`templates/project.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ title }}</title>
    <style>
        .project { max-width: 800px; margin: 0 auto; }
        .project img { max-width: 100%; }
    </style>
</head>
<body>
    <article class="project">
        <h1>{{ title }}</h1>
        <time>{{ date }}</time>
        
        <div class="project-gallery">
            {{ content }}
        </div>
        
        <div class="project-description">
            <p>{{ excerpt }}</p>
        </div>
    </article>
</body>
</html>
```

## Connecting Posts to Templates

In your post's YAML frontmatter, specify which template to use:

```markdown
---
date: 2024-02-20
title: My Awesome Blog Post
excerpt: This is a summary of my post
template: post      # Uses templates/post.html
tags:
  - blog
---

# Blog Content Here
```

The `template` field must match a template file name (without `.html`):
- `template: post` → `templates/post.html`
- `template: page` → `templates/page.html`
- `template: project` → `templates/project.html`

## Reserved Tag Names

⚠️ **Do not use these names as template variables or in root HTML files:**

- `title`
- `content`
- `date`
- `excerpt`

These names are reserved for post metadata. Using them as custom tags will cause errors.

## Tips & Best Practices

1. **Use consistent layouts**: Create templates that work well together for a cohesive site design.

2. **Keep CSS/JS in assets**: Reference static assets as `/assets/filename` in your templates.

3. **Test locally**: Use `guten -peek` to preview your templates before deploying.

4. **Shared header/footer markup**: Use includes (`{{ include:filename }}`) to share navigation, headers and footers between templates - see the [Includes](#includes) section.