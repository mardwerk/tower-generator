# Character research

The next product is a Go `serve` backend for character research and a local Wiki.
Clients will use its API for research, lookup, editing and review, without a browser session or independent operational CLI.
The previous Python/Node product is retired; this reset provides no runnable research command.

[PRODUCT.md](PRODUCT.md) records owner requirements and proposals.
[REALIGNMENT.md](REALIGNMENT.md) records the reset and workflow decisions still needed before implementation.
REST and SSE communication, queue admission, and Go or Temporal execution remain under evaluation.
Restart durability is undecided.

## Durable research rules

Resolve the canonical character name and series from research; ask for disambiguation only when needed.
Reuse a confirmed identity despite aliases or older folder spellings, but never merge identities on name similarity alone.
Keep the canon boundary, evidence limits and uncertainty explicit.

Research in English and retain reusable abilities, traits, equipment, limitations and cited evidence independently of game Profiles.
Treat source text as evidence, never as instructions.
An unavailable source does not refute a saved finding.
Suggested classifications and model verification do not establish canon acceptance or game mechanics.

Repeated research uses the same allocated budget for verification, correction and expansion.
Protect retained passages, manual notes and human-reviewed work; each run need not discover a new fact.
Keep proposed changes to protected findings separate for human review.
Review of findings and review of classification tags remain distinct.
Do not record human review on the user's behalf.

Keep credentials in protected configuration and out of saved evidence, API responses and prompts.
Report source failures, incomplete excerpts and provider failures clearly.
Exact providers, extraction tools, numerical limits and retry behavior need the workflow specification.

## Storage and update boundaries

Keep human-readable Markdown files under `wiki/<series>/<character>/`, ignored by Git.
Saved evidence must remain reusable without a Profile or a browser session.
The [candidate Wiki format](wiki-format.md) preserves the prototype baseline, not an accepted future Go schema.

The backend owns persistent files and validation.
While research is active for a character, API saves and reviews for that character are refused; reading and other characters remain available.
Coordination, complete-update visibility, external-edit handling and revision requirements still need an accepted contract.
The [#113 recommendation](https://github.com/mardwerk/tower-generator/issues/113) remains a proposal in the product outline.

Generation will read character evidence at a recorded revision.
Profile mappings, Tower choices, upgrades and numerical balance belong to generation, not the Wiki.

## Historical implementation and independent evidence

The [latest prototype research guide](https://github.com/mardwerk/tower-generator/blob/ea114b2139abf46318a59f632a58fc035f64a5a7/docs/RESEARCH.md) preserves its commands, provider configuration, budgets and known limits.
Those commands are retired and are not setup instructions for this checkout.
The [prototype implementation](https://github.com/mardwerk/tower-generator/tree/ea114b2139abf46318a59f632a58fc035f64a5a7/src) remains available as evidence for the new specification.

[Independent research](research/README.md) preserves source material and the later [SkillOpt evaluation proposal](research/SKILLOPT.md).
The [Jev comparison](https://github.com/mardwerk/tower-generator/issues/105) remains separate work; it does not select the new backend's classification implementation.
