# Troubleshooting

Common errors and how to fix them.

## Error: "No frontmatter in post"

**Cause:** Your post file doesn't start with `---` on the first line.

**Fix:** Ensure your post starts with YAML frontmatter:

```yaml
---
date: 2024-01-15
title: My Post
excerpt: Summary
template: post
tags:
  - blog
---
```

## Error: "template not found"

**Cause:** The `template` field in your post doesn't match any template file.

**Fix:** Create the template file or use an existing one:

- Post specifies `template: post`
- Need `templates/post.html` file

## Build succeeds but no HTML in output

**Cause:** Missing `index.html` in project root.

**Fix:** Add an `index.html` file to your project root. `guten -init` also creates a placeholder one when scaffolding a new project.

## Assets not copied

**Cause:** Missing `assets/` directory.

**Fix:** Create an `assets/` directory (can be empty if you don't need static assets). Note that `guten -init` already creates this directory when scaffolding a new project.

## Server won't start on port 5000

**Cause:** Port 5000 is already in use.

**Fix:** Stop the other process using that port, or modify the port in `internal/server/server.go`.
