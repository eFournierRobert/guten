---
date: 2024-01-15
title: Hello, World!
excerpt: Welcome to my blog powered by guten. This is my very first post.
template: post
tags:
  - blog
  - demo
---

# Hello, World!

Welcome to my blog! This is my very first post using **guten**, a static site generator written in Go.

## What is guten?

Guten is a simple static site generator that:

- Parses Markdown posts with YAML frontmatter
- Applies HTML templates to generate static pages
- Supports subdirectories in posts
- Copies assets and root HTML files to output

## How it works

1. Create posts in `posts/` with YAML frontmatter
2. Create templates in `templates/`
3. Run `guten -build` to generate the static site
4. Done! Your static files are in `out/`

## Next steps

- Add more posts to your `posts/` directory
- Customize your `templates/` for different sections
- Add CSS/JS to `assets/` for styling and interactivity