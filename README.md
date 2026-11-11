# Guten

A small static site generator written in Go. Write Markdown posts, style them with HTML and CSS, and build a site you can host anywhere.

## Install

With Go 1.27 or newer:

```sh
git clone https://github.com/eFournierRobert/guten.git
cd guten
go install .
```

Add Go's binary directory (`go env GOBIN`, or `$(go env GOPATH)/bin` when unset) to your `PATH`.

## Try it

From the repository, copy the [example blog](docs/EXAMPLES/README.md):

```sh
cp -R docs/EXAMPLES my-site
cd my-site
guten -peek
```

Open **http://localhost:5000**. Stop with Ctrl+C; rerun `guten -peek` after edits to rebuild. There is no live reload.

## Make it yours

- Edit `index.html` and other root HTML pages.
- Write posts in `posts/` with YAML frontmatter; `template: post` selects `templates/post.html`.
- Put CSS, images, and other static files in `assets/`.
- Use `{{ blog }}` to list posts tagged `blog`, newest first. Customize each entry in `previews/blog.html`.
- Share HTML with `{{ include:footer.html }}` from `includes/footer.html`.

| Command | Action |
| --- | --- |
| `guten -init my-site` | Create an empty scaffold with a placeholder homepage. Add content and templates yourself. |
| `guten -new-post posts/hello.md` | Create a starter post using today's date and the `post` template. Edit its metadata and body. |
| `guten -build` | Rebuild `out/` from the current project. **Deletes the previous `out/` first.** |
| `guten -peek` | Build once and serve `out/` on port 5000. |

Publish the contents of `out/` to a static host. Serve from the domain root; generated post links start with `/posts/`.

[Start from scratch](docs/QUICKSTART.md) · [Files and posts](docs/STRUCTURE.md) · [Templates](docs/TEMPLATES.md) · [Troubleshooting](docs/TROUBLESHOOTING.md)
