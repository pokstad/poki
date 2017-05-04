# Poki

Poki is a minimal markdown wiki system written in Go that enables teams to write documentation in concert.

## Features

NONE!

## Roadmap

- Exposes a REST API to manage markdown posts
- Default underlying storage system utilizes the filesystem and git to manage revisions of the wiki
- Ensures serialized writes to individual posts while allowing concurrent reads

### FAQ

- Why?
  - Hardly any wiki systems utilize markdown, which is the main markup language of choice among programmers.
- Why filesystem based persistence?
  - So that existing static site generators (e.g. Hugo) can be used to render the markdown into HTML.
