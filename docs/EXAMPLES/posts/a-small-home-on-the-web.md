---
date: 2026-11-10
title: A small home on the web
excerpt: An evening spent making a place for words, with room to leave a few things unfinished.
template: post
tags: [journal]
---

I wanted a place for the notes that outgrow the notebook. Somewhere I could put a few paragraphs, give them a title, and find them again a year from now.

So I cleared a corner of the desk and spent an evening making this little site. The tea went cold before I remembered to drink it.

![An overhead illustration of an open notebook, a pencil, and a cup of tea on a green desk.](/assets/img/writing-desk.svg)

*A notebook, a cup of tea, and enough room to begin.*

## A few files to look after

I started with a copy of this journal rather than an empty folder. Having something on the page made it easier to decide what I wanted to change.

Most evenings, these are the only files I need to think about:

| When I want to… | I open… |
| --- | --- |
| Write a new entry | A Markdown file in `posts/` |
| Change the front page | `index.html` |
| Adjust the colours or type | `assets/css/style.css` |

The article layout lives in `templates/post.html`. Once it felt comfortable to read, I could leave it alone and get back to the words.

## The small routine

From the site folder, I make a new draft:

```sh
guten -new-post posts/an-ordinary-tuesday.md
```

Then I open the file, change the title and excerpt, and add `tags: [journal]` to its frontmatter so it appears on the homepage. Under the closing `---`, I start writing. A heading when the subject changes. A list when there are a few things to remember.

When there is enough to read, I take a look in the browser:

```sh
guten -peek
```

The site is at `http://localhost:5000`. After another round of edits, I stop the server with **Ctrl+C** and run the same command again. A small pause to see the page with fresh eyes.

## Leave a little room

There are still things I could add. There always will be. For now, the pages are readable, the links go somewhere, and there is a place for the next entry.

That seems like a good point to close the laptop and make another cup of tea.
