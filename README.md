# guten

A small static site generator written in Go. Write posts in Markdown, define layouts with HTML templates, and generate a static site.

## Install & Run

```bash
git clone https://codeberg.org/efournierrobert/guten.git
cd guten
go install .
```

Create your own directory with `posts/*.md` files (with YAML frontmatter) and `templates/*.html` files (with `{{ variables }}`), then:

```bash
guten -build    # generate site to out/
guten -peek     # build and serve on :5000
```

## Example post

Create a file `posts/hello-world.md`:

```yaml
---
date: 2026-09-13
title: Hello World
excerpt: My first guten post
template: post
tags:
  - blog
---

This is my **first** post.
```

Create `templates/post.html`:

```html
<!DOCTYPE html>
<h1>{{ title }}</h1>
<p class="date">{{ date }}</p>
{{ content }}
```

Run `guten -build` — output goes to `out/`.
