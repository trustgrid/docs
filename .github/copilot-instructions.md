# GitHub Copilot Agent: Custom Instructions for Trustgrid Docs

These instructions are for the GitHub Copilot agent and Copilot Chat to follow when making edits or generating content for this repository, especially for files under the `content` directory.

## 1. Front Matter
- Do **not** add `weight` or `date` fields to the front matter unless the page is in `content/en/release-notes`.
- Always include:
  - `linkTitle`: Use at most 3 words (prefer 2 or less).
  - `title`: Use a longer, descriptive title (ideally less than 6 words).

## 2. Internal Links
- Use the `{{<relref >}}` shortcode for all references to other pages (no space after `{{<`).
- Always use relref as a Markdown link with link text, e.g. `[Link Text]({{<relref "docs/nodes#example" >}})`.
- Do not use a leading `/` in the path for relref shortcodes.
- Remove any trailing `/` from the path and from anchor links. Never use a slash before an anchor.
  - Example: `[Nodes]({{<relref "docs/nodes#example" >}})` (not `/docs/nodes/#example`, `/docs/nodes/#example/`, or with a leading `/`).

## 3. Images & Screenshots
- Use the `{{<tgimg >}}` shortcode for screenshots or images (no space after `{{<`).
- Default the width to 50%.
- Suggest a filename and caption, but leave the actual image and final caption for the user to add.
  - Example: `{{<tgimg src="suggested-name.png" width="50%" caption="Describe what this screenshot shows." >}}`

## 4. Formatting
- Use backticks (`` `text` ``) for user inputs, commands, or code.
- Use bold (`**text**`) for UI elements or items that would be clicked on in the Trustgrid portal.

---

> Place this file at `.github/copilot-instructions.md` so the Copilot agent and Copilot Chat can reference it for all future content and edit requests.
