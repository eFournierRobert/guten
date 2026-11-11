# Files and posts

Run Guten from your site's root directory:

```text
my-site/
├── index.html          # Homepage; other root .html pages work too
├── assets/             # Required, even if empty; copied as-is
├── posts/              # Required, even if empty; Markdown posts
│   └── hello.md
├── templates/
│   └── post.html       # Required by posts with template: post
├── includes/           # Optional shared snippets
│   └── footer.html
└── previews/           # Optional markup for tag list entries
    └── blog.html
```

`guten -init` creates all five directories, empty, plus a placeholder `index.html`.

## Post format

Every file under `posts/` is parsed as a post, so keep only Markdown posts there. Start each file with YAML frontmatter:

```markdown
---
date: 2026-11-10
title: A walk by the river
excerpt: Taking the long way home.
template: post
tags: [blog, outdoors]
---

Write your content in **Markdown**.
```

`template: post` selects `templates/post.html`. Use `YYYY-MM-DD` dates. Tags group posts into lists; `tags: []` is fine. The tag names `title`, `content`, `date`, and `excerpt` are reserved and cause build errors.

## Output

Each build clears `out/`, copies `assets/`, renders posts, and processes all root `.html` files. `posts/hello.md` becomes `out/posts/hello.html`; post subdirectories are preserved. Source templates, includes, and previews aren't copied.

A root `index.html` supplies the homepage; without it, posts can still build. Only HTML files directly in the root are processed as standalone pages. Keep other downloadable files in `assets/`.

Use root-relative asset paths such as `/assets/css/style.css` so they work from posts too. Deploy the contents of `out/` at the domain root.
