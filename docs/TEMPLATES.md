# Templates, lists, and includes

Guten replaces exact notations such as `{{ title }}`. Keep the spaces shown; this is string substitution, not Go's template language. Unknown notations stay unchanged.

## Post templates

A post's `template: post` selects `templates/post.html`. Available variables:

| Notation | Value |
| --- | --- |
| `{{ title }}` | Post title |
| `{{ date }}` | Date as `YYYY-MM-DD` |
| `{{ excerpt }}` | Post summary |
| `{{ content }}` | Markdown body rendered as HTML |

See the [minimal template](QUICKSTART.md) or the [styled example](EXAMPLES/templates/post.html). Metadata is inserted as-is, without HTML escaping; keep markup-sensitive characters escaped where needed.

## Tag lists and custom previews

Put `{{ blog }}` in a root HTML page, post template, or include to list all posts tagged `blog`, newest first. By default, entries show a linked title and excerpt inside `<div class="blog-previews">`.

To customize each entry, create `previews/blog.html`:

```html
<article>
  <time datetime="{{ date }}">{{ date }}</time>
  <h2><a href="{{ link }}">{{ title }}</a></h2>
  <p>{{ excerpt }}</p>
</article>
```

The wrapper stays the same. Custom previews support `title`, `date`, `excerpt`, and **`link`**, a root-relative post URL such as `/posts/hello.html`. `link` is a preview variable, not a post template variable. Keep previews to per-entry markup; avoid embedding other tag lists in them.

## Shared includes

Create `includes/footer.html`, then insert it with:

```html
{{ include:footer.html }}
```

Includes work in root HTML files and post templates, and include references in Markdown bodies are also expanded. Match the full filename, including its extension. Only files directly in `includes/` and `previews/` are loaded; both directories are optional.

Tag lists in includes are expanded when loaded. Post variables in includes stay literal, so keep them in the post template. Nested includes aren't supported; use self-contained snippets.

Tag lists in Markdown bodies are not expanded. Use them in the surrounding HTML template instead.
