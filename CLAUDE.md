# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the **Trustgrid documentation site** ([docs.trustgrid.io](https://docs.trustgrid.io)), built with [Hugo](https://gohugo.io/) using the [Google Docsy](https://www.docsy.dev/) theme (imported as a Hugo module, not vendored).

## Build & Development Commands

- **Local dev server:** `hugo server` or `make run`
- **Production build:** `hugo --minify` or `make build`
- **Clean build output:** `make clean` (removes `public/`)
- **Install JS dependencies:** `npm i` (required before build in CI; installs autoprefixer/postcss)

Hugo **extended version** is required (v0.75.0+; CI uses v0.133.1). Go is also required for Hugo module resolution.

There are no tests or linters configured for this project.

## Architecture

### Content Structure

All content lives under `content/en/` with these top-level sections:

- `docs/` — Main documentation, organized to mirror the Trustgrid Portal navigation (nodes, clusters, certificates, alarms, etc.)
- `tutorials/` — Deployment guides, tool usage, and troubleshooting walkthroughs
- `getting-started/` — Conceptual overview of Trustgrid
- `release-notes/` — Versioned release notes
    - `release-notes/node` - Specific release notes for the Trustgrid Appliance Node software
    - `release-notes/agent` - Specific release notes for the Trustgrid Agent Node software
    - `release-notes/cloud` - Specific release notes for the Trustgrid Portal and Control Plane software
- `help-center/` — Support resources

### Content Conventions

- Section landing pages use `_index.md`
- Front matter uses YAML with fields: `title`, `linkTitle`, `weight` (for ordering), `tags`, `categories`, `description`
- Cross-references use Hugo ref shortcodes: `{{<ref "docs/nodes">}}`
- The `{{<readfile>}}` shortcode includes shared markdown fragments (e.g., `content/en/docs/nodes/agent-v-app.md`)

### Custom Shortcodes (`layouts/shortcodes/`)

- **`tgimg`** — Centered, click-to-expand image. Params: `src`, `alt`, `caption`, `width`, `height`, `class`
- **`fields`** / **`field`** — Renders a two-column table (Field Name / Description). Wrap `{{<field "Name">}}description{{</field>}}` inside `{{<fields>}}`
- **`codeblock`** — Code block with copy button
- **`webm`** — Embedded video player. Params: `src`, `width`, `height`, `alt`, `caption`
- **`node-release`** — Release info table. Params: `package-version`, `core-version`, `release`
- **`agent-release`** — Agent release info table. Params: `package-version`, `release`

### Configuration

- `config.toml` — Hugo site config (Docsy theme settings, Lunr offline search, Mermaid diagrams enabled, Google Analytics)
- `go.mod` / `go.sum` — Hugo module dependencies (Docsy theme)

### CI/CD

- **On PR:** Hugo build check (`.github/workflows/ci.yml` and `publish.yml`)
- **On push to main:** Build + deploy to GitHub Pages via `peaceiris/actions-gh-pages`

### Images

Images are stored alongside their markdown files in the same directory (co-located), not in a central assets folder. Use the `tgimg` shortcode to embed them.
