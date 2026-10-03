# Candidate Wiki format

The retired prototype stored readable Markdown and YAML under `wiki/<series>/<character>/`, with `README.md` and retained passages in `sources/*.md`.
The local Markdown Wiki is an owner requirement; the prototype's detailed schema is a candidate baseline for the Go specification.
Field names, format versions, identity normalization and review fingerprints remain subject to that review.
The [selected publication rules](PRODUCT.md#selected-wiki-publication) define one entry commit point with immutable captures; the exact schema and implementation verification remain open.
No current reader, validator or migration command is provided by this reset.

The [complete prototype format](https://github.com/mardwerk/tower-generator/blob/ea114b2139abf46318a59f632a58fc035f64a5a7/docs/wiki-format.md) preserves examples and implementation details.
The [product rules](PRODUCT.md) and [research rules](RESEARCH.md) establish the requirements that the future format must serve.

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

Retain older cited passages when refreshed sources omit them, with their original retrieval dates.
Preserve separate references when different sources repeat a claim.
Research must preserve prior evidence and cannot delete a finding solely because a current source omits it.

## Review and classification baseline

Suggested classification tags described generalized delivery and functions independently of any Profile.
A suggested tag or model verification verdict was not a canon fact, human review or game rule.
Review of findings did not automatically approve classifications.
The retained [category definitions](research-categories.yaml) remain evidence for the new classification contract.

The prototype recorded human review separately from model verification.
Human-reviewed records survived repeat research, with proposed replacements kept separately for a person to accept.
Generated summaries occupied a marked block, while manual prose outside that block was preserved.

Saved edits reset entry review to draft, and a content fingerprint detected later changes to entry or source files.
Revision hashes let callers detect stale edits when the expected revision was supplied.
Mandatory API revision details remain under review in [PRODUCT.md](PRODUCT.md).
Complete-update publication is selected there, but these historical files do not establish its implementation guarantees.

## Storage boundaries

Markdown and YAML are data, not executable page content.
Keep dependencies and citations within the selected Wiki; the prototype rejected paths outside it and child symbolic links.
Keep credential values out of entry metadata, source files and prompts.
Profile mappings, upgrades, numerical stats and game adaptations belong outside the Wiki.

Prototype research metadata recorded its last successful allocation, usage, model, source limits, warnings and follow-up questions.
The future operation contract must define progress, failure outcomes and restart behavior separately from saved character evidence.
A research workflow engine or API transport choice does not by itself select a new file schema.
