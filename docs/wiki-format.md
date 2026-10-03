# Candidate Wiki format

The retired prototype stored readable Markdown and YAML under `wiki/<series>/<character>/`, with `README.md` and retained passages in `sources/*.md`.
Its detailed schema is a candidate baseline for the Go specification.
Field names, format versions, identity normalization and review fingerprints remain subject to that review.
No current reader, validator or migration command is provided by this reset.

The [complete prototype format](https://github.com/mardwerk/tower-generator/blob/ea114b2139abf46318a59f632a58fc035f64a5a7/docs/wiki-format.md) preserves examples and implementation details.
The [Wiki publication contract](PRODUCT.md#selected-wiki-publication), [save and review policy](PRODUCT.md#selected-save-and-review-policy) and [research rules](RESEARCH.md) govern the future format.

## Evidence and identity baseline

The entry frontmatter recorded the canonical name, series, canon scope, aliases, source references and review metadata.
Its Markdown body held a readable summary, evidence limits, contradictions and open questions.
Abilities, traits and equipment shared a cited-record shape with scope and limitations.

Source files retained titles, URLs, retrieval dates, access information and individually addressable passages.
Redirects and intentionally retained excerpts were explicit; a search excerpt did not imply the full page had been read.
Supplied passages needed origin information for human review.
Content-derived passage IDs avoided renumbering unrelated citations when evidence was added.

The prototype kept one declared scope per entry and allowed narrower scopes on individual records.
It rejected conflicting identities or scopes rather than silently merging them.
Its normalization and older-folder matching are historical implementation choices, not a new Go identity contract.

## Review and classification baseline

Suggested classification tags described generalized delivery and functions.
The prototype recorded human review separately from model verification.
Generated summaries occupied a marked block, while manual prose outside that block was preserved.

Saved edits reset entry review to draft, and a content fingerprint detected later changes to entry or source files.
Revision hashes let callers detect stale edits when the expected revision was supplied.

## Storage boundaries

The prototype rejected paths outside the Wiki and child symbolic links.

Prototype research metadata recorded its last successful allocation, usage, model, source limits, warnings and follow-up questions.
These historical details do not establish implementation guarantees for the Go backend.
