---
date: 2024-02-10
title: Markdown showcase
excerpt: Formatting, local assets, and includes inside Markdown content.
template: post
tags:
  - blog
  - features
---

## Text and headings

A paragraph with **bold**, *emphasis*, ~~strikethrough~~, and `inline code`.

### A smaller heading

> A blockquote should have its own visual treatment.

## Lists

- One unordered item
- Another item
    - A nested item

1. Edit a source file.
2. Rebuild the site.
3. Refresh the browser.

## Fenced code

```go
func greeting() string {
    return "Hello from Guten"
}
```

## Table

| Input | Output |
| --- | --- |
| Markdown | HTML |
| Templates | Pages |
| Assets | Copied files |

---

## Links and images

[Local project page](/posts/project.html) and [external Go documentation](https://go.dev/doc/).
Rendered Markdown links open in a new tab.

![Geometric hills beneath a golden sun](/assets/img/landscape.svg)

## Raw HTML

<p class="html-example">This paragraph is raw HTML inside Markdown.</p>

## Notation behavior

The following tag remains literal here, while the blog tag in the surrounding
template expands: `{{ blog }}`.

The next include is written directly in the Markdown source:

{{ include:content-note.html }}

This include shows that post variables stay literal inside includes, even when
the include is inserted into a rendered post:

{{ include:literal-variables.html }}
