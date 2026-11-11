# Fieldnotes — a Guten example

A fictional personal journal about walks, small projects, and paying attention. Four Markdown articles, an about page, shared navigation and footer, and one custom post preview. No JavaScript, external fonts, or build tools.

From the repository root, with Guten installed:

```sh
cp -R docs/EXAMPLES my-site
cd my-site
guten -peek
```

Visit **http://localhost:5000**. Stop with Ctrl+C and rerun after edits.

- `index.html` lists posts with `{{ journal }}`.
- `previews/journal.html` controls each entry's markup.
- `posts/*.md` supplies the articles; `templates/post.html` renders them.
- `posts/a-small-home-on-the-web.md` uses fenced code blocks, inline code, a table, and a local illustration, all written in Markdown.
- `includes/` shares navigation and footer across pages.
- `assets/css/style.css` contains all styling; `about.html` is a plain root page.

Edit these files to make it yours. Run `guten -build` and publish the contents of `out/` at your domain root.
