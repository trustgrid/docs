
# Trustgrid Docs: Content Authoring Instructions

When working with pages under the `content` directory, follow these repository-specific guidelines:

## 1. Front Matter Rules
- Do **not** add `weight` or `date` fields to the front matter unless the page is in `content/en/release-notes`.
- Always include:
	- `linkTitle`: Use at most 3 words (prefer 2 or less).
	- `title`: Use a longer, descriptive title (ideally less than 6 words).

## 2. Internal Links
- Use the `{{< relref >}}` shortcode for all references to other pages.
- Remove any trailing `/` from the path and from anchor links.
	- Example: `{{< relref "/docs/nodes#example" >}}` (not `/docs/nodes/#example`).

## 3. Images & Screenshots
- Use the `{{< tgimg >}}` shortcode for screenshots or images.
- Default the width to 50%.
- Suggest a filename and caption, but leave the actual image and final caption for the user to add.
	- Example: `{{< tgimg src="suggested-name.png" width="50%" caption="Describe what this screenshot shows." >}}`

## 4. Formatting Conventions
- Use backticks (`` `text` ``) for user inputs, commands, or code.
- Use bold (`**text**`) for UI elements or items that would be clicked on in the Trustgrid portal.

# How to Contribute

We'd love to accept your patches and contributions to this project. There are
just a few small guidelines you need to follow.

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Community Guidelines

This project follows
[Google's Open Source Community Guidelines](https://opensource.google.com/conduct/).
