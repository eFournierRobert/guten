# Troubleshooting

| Problem | Check |
| --- | --- |
| `guten: command not found` | Run `go install .` from the repository and add Go's binary directory to `PATH`. |
| `No frontmatter in post` | Every file in `posts/` needs YAML between `---` lines, starting on the first line. Keep other files in `assets/`. |
| Error opening a template | `template: post` needs `templates/post.html`. Initialization creates no templates. |
| Error copying assets or reading posts | Both `assets/` and `posts/` must exist, even if empty. Run commands from the site root. |
| No homepage | Add `index.html` at the site root. Other pages can build without it. |
| A notation appears literally | Check exact spacing, tag spelling, and include filename. Post variables don't expand inside includes. See [Templates](TEMPLATES.md). |
| Changes don't appear | `-peek` builds once. Stop it with Ctrl+C and rerun after editing. |
| Port 5000 is busy | Stop the process using it, then rerun `guten -peek`. There is no port flag. |
| Links break on your host | Publish the contents of `out/` at the domain root, not under a subdirectory. |

Keep source files outside `out/`: every build deletes and recreates that directory.
