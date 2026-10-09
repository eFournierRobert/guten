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

## Build succeeds but `{{ include:... }}` appears in the generated HTML

**Cause:** The include notation doesn't match any file in your `includes/` directory. The most common reasons:

- The file name is misspelled (including the `.html` extension - it's part of the notation: `{{ include:footer.html }}`)
- The file was placed in the wrong directory (top-level `includes/`, no subdirectories)
- The `includes/` directory doesn't exist or is empty

**Fix:** Check that the exact file exists in `includes/`. Includes are matched by the full file name. If a notation never matches, guten leaves it in the output as-is rather than failing the build. Note that nested includes (an `{{ include:... }}` inside an include file) are not supported and will also remain literal.

If you see `{{ title }}`, `{{ date }}` or similar rendered as literal text on pages that use includes: post variables are not expanded inside include files - they are only expanded in template files. Move that markup out of the include into the template.

## Build succeeds but no HTML in output

**Cause:** Missing `index.html` in project root.

**Fix:** Add an `index.html` file to your project root. `guten -init` also creates a placeholder one when scaffolding a new project.

## Assets not copied

**Cause:** Missing `assets/` directory.

**Fix:** Create an `assets/` directory (can be empty if you don't need static assets). Note that `guten -init` already creates this directory when scaffolding a new project.

## Server won't start on port 5000

**Cause:** Port 5000 is already in use.

**Fix:** Stop the other process using that port, or modify the port in `internal/server/server.go`.
