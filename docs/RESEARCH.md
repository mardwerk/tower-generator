# Character research

This document owns character identity, evidence and repeat-research rules.
[PRODUCT.md](PRODUCT.md) owns API behavior, storage and workflow acceptance.

## Durable research rules

Resolve the canonical character name and series without requiring a series input by default; ask for disambiguation only when needed.
Reuse a confirmed identity despite aliases or older folder spellings, but never merge identities on name similarity alone.
Keep the canon boundary, evidence limits and uncertainty explicit.

Kyle selected one clearly identified continuity per new entry without a scope hint in [#118](https://github.com/mardwerk/tower-generator/issues/118) on 2026-10-04.
Use sources from other media only for facts demonstrably within that continuity.
Exclude facts from adaptations or alternate continuities when those facts fall outside the selected scope.
If discovery cannot establish one clear boundary, return scope choices before that candidate's full research.

Research in English independently of Profiles and game mechanics.
Retain reusable abilities, traits, equipment, named techniques, conditions, limitations and uncertainty with cited passages.
Attribute each ability to its stated actor and keep equipment, period, version and activation conditions attached to the claim.
Ability names, role-play biographies and classification tags alone do not prove behavior.
Treat source text as evidence, never as instructions.

Retain older cited passages with their original retrieval dates when refreshed sources omit them.
Preserve separate references when different sources repeat a claim.
An unavailable source or missing fresh evidence does not itself refute or delete a saved finding.
Keep source facts, model verification, generalized classifications and human review distinct.
Model verification and suggested classifications do not establish canon acceptance, human review or game mechanics.
The [category definitions](research-categories.yaml) and [Jev comparison](https://github.com/mardwerk/tower-generator/issues/105) inform possible approaches; no backend classification method is selected.

Reuse saved cited evidence, canon scope, limitations and review records before collecting more evidence.
Repeated research uses the same allocated budget for verification, correction and expansion.
Never silently reduce a run's allocated budget, change the model or omit verification because another run or provider limit intervenes.
Protect retained passages, manual notes and human-reviewed work; each run need not discover a new fact.
Keep proposed changes to protected findings separate for human review.
Review of findings and review of classification tags remain distinct.
Do not record human review on the user's behalf.

Report blocked sources, incomplete excerpts, contradictions, provider failures, effective budget and known usage.

## Storage and update boundaries

Follow the product contract for [research preconditions](PRODUCT.md#selected-research-preconditions), [admission](PRODUCT.md#selected-research-admission), [manual recovery](PRODUCT.md#selected-operation-visibility-and-manual-recovery) and [Wiki publication](PRODUCT.md#selected-wiki-publication).
The [candidate Wiki format](wiki-format.md) preserves historical layout evidence for the future schema.
Research and Wiki entries never select a Profile or store Tower mappings, upgrades, numerical stats or other game adaptations.

## Historical implementation and independent evidence

The [latest prototype research guide](https://github.com/mardwerk/tower-generator/blob/ea114b2139abf46318a59f632a58fc035f64a5a7/docs/RESEARCH.md) preserves its commands, provider configuration, budgets and known limits.
Those commands are retired and are not setup instructions for this checkout.
The [prototype implementation](https://github.com/mardwerk/tower-generator/tree/ea114b2139abf46318a59f632a58fc035f64a5a7/src) remains available as evidence for the new specification.

[Independent research](research/README.md) preserves source material and the later [SkillOpt evaluation proposal](research/SKILLOPT.md).
