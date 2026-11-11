# Start from scratch

[Try the finished example](EXAMPLES/README.md) if you'd rather start with a styled site. Otherwise, install Guten as described in the [README](../README.md), then:

```sh
guten -init my-site
cd my-site
guten -new-post posts/hello.md
```

The scaffold's directories are empty. Create `templates/post.html`:

```html
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{ title }}</title>
</head>
<body>
  <a href="/">Home</a>
  <article>
    <h1>{{ title }}</h1>
    <time datetime="{{ date }}">{{ date }}</time>
    {{ content }}
  </article>
</body>
</html>
```

Edit `posts/hello.md`, keeping the opening and closing `---`:

```markdown
---
date: 2026-11-10
title: Hello, world
excerpt: A first note from my new site.
template: post
tags: [blog]
---

This is my **first post**.
```

Replace the scaffold's `index.html` with:

```html
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>My blog</title>
</head>
<body>
  <h1>My blog</h1>
  {{ blog }}
</body>
</html>
```

Run `guten -peek` and visit **http://localhost:5000**. After edits, stop with Ctrl+C and run it again. Use `guten -build` to generate `out/` for publishing.

For an existing directory, `guten -init .` creates the scaffold there and **overwrites `index.html`**.
