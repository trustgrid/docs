# AGENTS.md

Canonical guidance for AI coding agents working in this repository. This is the
single source of truth: `CLAUDE.md` imports it, and `.github/copilot-instructions.md`
is a symlink to it. Update this file only.

## Project Overview

This is the **Trustgrid documentation site** ([docs.trustgrid.io](https://docs.trustgrid.io)), built with [Hugo](https://gohugo.io/) using the [Google Docsy](https://www.docsy.dev/) theme (imported as a Hugo module, not vendored).

## Build & Development Commands

- **Local dev server:** `hugo server -F` or `make run`
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
    - `release-notes/node` — Specific release notes for the Trustgrid Appliance Node software
    - `release-notes/agent` — Specific release notes for the Trustgrid Agent Node software
    - `release-notes/cloud` — Specific release notes for the Trustgrid Portal and Control Plane software
- `help-center/` — Support resources

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

## Content Authoring Rules

When creating or editing pages under the `content` directory, follow these conventions.

### 1. Front Matter

- Front matter uses YAML, delimited by `---`. **Never** remove the leading/trailing `---`.
- Always include:
  - `linkTitle`: at most 3 words (prefer 2 or fewer).
  - `title`: a longer, descriptive title (ideally fewer than 6 words).
- Other fields: `tags`, `categories`, `description`.
- Do **not** add `weight` or `date` fields unless the page is under `content/en/release-notes`.
- Section landing pages use `_index.md`.

### 2. Internal Links

- Use the `{{<relref >}}` shortcode for all references to other pages (no space after `{{<`).
- Always write relref as a Markdown link with link text, e.g. `[Link Text]({{<relref "docs/nodes#example" >}})`.
- Do **not** use a leading `/` in the relref path.
- Remove any trailing `/` from the path, and never put a slash before an anchor.
  - Example: `[Nodes]({{<relref "docs/nodes#example" >}})` — not `/docs/nodes/#example`, `docs/nodes/#example`, or with a leading `/`.
- Do not guess at pages. Search the repository to find the correct page if unsure.
- The `{{<readfile>}}` shortcode includes shared markdown fragments (e.g., `content/en/docs/nodes/agent-v-app.md`).

### 3. Images & Screenshots

- Use the `{{<tgimg >}}` shortcode for screenshots or images (no space after `{{<`).
- Default the width to 50%.
- Images are co-located with their markdown file in the same directory, not in a central assets folder.
- Suggest a filename and caption, but leave the actual image and final caption for the user to add.
  - Example: `{{<tgimg src="suggested-name.png" width="50%" caption="Describe what this screenshot shows." >}}`

### 4. Formatting

- Use backticks (`` `text` ``) for user inputs, commands, or code.
- Use bold (`**text**`) for UI elements or items that would be clicked on in the Trustgrid portal.

## Writing Style

- No em-dashes unless they give the most concise expression. Restructure otherwise.
- No "Not X, but Y" constructions. Just say the thing.
- No cliches: "single pane of glass," "table stakes," "net-net," "at the end of the day," etc.
- No bold labels at the start of every paragraph. Bold is for things that genuinely need to stand out, not for fake headers inside a section.
- Short sentences over long ones. If you can remove words and keep the meaning, do it.
- If you find yourself hedging with "it's worth noting that" or "importantly," just cut it and say the thing directly.
