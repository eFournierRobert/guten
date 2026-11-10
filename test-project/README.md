# Guten development fixture

Keep this directory as a clean, committed reference site. Copy it into the ignored
`dev-project/` directory for everyday experiments.

## Start developing

From the repository root, when `dev-project/` does not already exist:

```bash
cp -R test-project dev-project
cd dev-project
go run .. -peek
```

Open http://localhost:5000. This runs the current Go source, so no installed
`guten` binary is needed. Go must satisfy the version in the repository's `go.mod`.
The preview builds once; it does not watch files. After editing content, templates,
assets, or Go code, stop it with Ctrl+C and rerun `go run .. -peek`.

Other commands, from `dev-project/`:

```bash
# Generate without starting a server
go run .. -build

# Create a post using today's date and the post template
go run .. -new-post posts/scratch.md
```

The new post is untagged by default. Link it manually or add `blog` to its tags to
make it appear in the homepage previews. Rebuild after editing.

Both `dev-project/` and generated `out/` directories are ignored by Git. Changes
to `test-project/` are tracked, so make deliberate fixture improvements there.
To reset the playground, stop the server, remove or rename the existing
`dev-project/`, and copy again from the repository root. Copying onto an existing
directory can nest the fixture instead of replacing it.

## Feature checklist

| Feature | Where to inspect | Expected result |
| --- | --- | --- |
| Root HTML generation | `/`, `/behavior.html` | Both root HTML files are generated with shared navigation and footer |
| Frontmatter and template variables | Any post | Title, date, excerpt, and rendered content match the source |
| Multiple templates | `/posts/markdown.html`, `/posts/about.html`, `/posts/project.html` | Post, Page, and Project template labels |
| Untagged content | `/posts/about.html` | Generated and manually linked, absent from previews |
| Nested posts | `/posts/notes/field-notes.html` | Nested output directory, working links and assets |
| Tag previews in root HTML | Homepage blog section | Nested field notes, Markdown showcase, First entry, in that order |
| Overlapping tags | Homepage features section | Sample project, then Markdown showcase; both also appear under other tags |
| Tags in templates | Blog sidebar on any post-template page | Blog previews expand directly in the template |
| Tags in includes | Shared previews on homepage and posts | Features previews expand inside `includes/related.html` |
| Includes in root HTML and templates | Navigation and footer | Shared markup appears on every page |
| Includes in Markdown content | Markdown showcase | The content-note aside appears within the article |
| Markdown | Markdown showcase | Headings, emphasis, strikethrough, lists, blockquote, fenced code, table, rule, links, image, raw HTML |
| Markdown link rendering | Links in Markdown showcase | Rendered links have `target="_blank"` |
| Copied asset directories | Homepage and `/assets/` URLs | Styles apply, landscape loads, JavaScript status says it loaded successfully |
| Unknown notations | Notation checks page | Unknown tag and missing include remain visible without a build error |
| Literal tags in Markdown | Markdown showcase | Blog notation stays literal in the article while expanding in the sidebar |
| Literal post variables in includes | Notation checks page and Markdown showcase | All four variable notations remain literal |
| Output cleanup | Create `out/stale.txt`, then rebuild | The stale file disappears |

All asset and navigation URLs are root-relative so they work on nested pages.
Includes and templates are consumed during generation, not copied as directories
into `out/`. The fixture README is not emitted either.

## Optional and error checks

Run these experiments only in `dev-project/`; restore each change before moving
to the next. The default fixture is intentionally a successful build.

- **Optional includes:** rename `includes/` to `includes-disabled/`, then build.
  The build succeeds and include notations remain literal. Rename it back afterward.
- **Reserved tags:** add `title`, `content`, `date`, or `excerpt` as a tag in a post.
  Building fails with a reserved-tag error. Remove the tag afterward.
- **Missing template:** temporarily change a post's template to `missing`.
  Building fails when opening that template. Restore the original name.
- **Include subdirectories:** put a snippet in `includes/subdir/only.html` and
  reference it with `{{ include:subdir/only.html }}` in a root page. The notation
  remains literal because only files directly in `includes/` are loaded.
- **Nested includes:** includes are intended as flat snippets; nesting them is
  unsupported. Keep include references in root HTML, templates, or Markdown.
  The current implementation uses sequential replacements, so some nesting may
  expand incidentally depending on filename order; do not rely on that behavior.

To check scaffolding independently, run this from the repository root:

```bash
scratch_dir=$(mktemp -d)
go run . -init "$scratch_dir/site"
```

Inspect the new site's `index.html`, `assets/`, `posts/`, `templates/`, and
`includes/`. Scaffolding creates empty content directories; use this fixture for
a complete buildable site. Delete the temporary directory when finished.
