# API-only reset and backend decision

Kyle selected pure Go with REST, SSE and a bounded in-process research queue on 2026-10-03.
One `serve` backend owns research, Wiki lookup, saving and review; clients need no browser session.
The [product rules](PRODUCT.md) own the requirements, and [#106](https://github.com/mardwerk/tower-generator/issues/106) coordinates workflow acceptance.
Product implementation remains pending that workflow review.

## Completed reset

Kyle approved [PR #120](https://github.com/mardwerk/tower-generator/pull/120), applied directly to main at `b085ba8`.
The old Python product, complete web/Lab tree, Node tooling, Usopp experiment and broken setup are removed.
The deleted Default Profile is retired; a new Profile and Tower design follow the research/Wiki workflow.
The Usopp authoring scope in [#85](https://github.com/mardwerk/tower-generator/issues/85) is retired; it no longer receives new generation reviews.

Git history preserves the retired implementation.
Independent research documents remain under [docs/research](research/README.md).
Kyle's subsequent cleanup removed the snapshot wrapper, manifest and Gojo experiment; Git history preserves them.
Local credentials and collected Wiki data were untouched, and known retired local tooling was removed.
System Python is the default for independent helpers; uv is optional, without a required project virtual environment.

## Recovery check and outcome

A two-step fake-provider experiment compared Go checkpoints with the Temporal Go SDK.
Actual processes were killed and restarted; Temporal's Worker and development Service reused the same persisted database.
No paid service was called.

- Both runners reused a recorded identity response and completed synthesis without repeating identity.
- When a paid response was not recorded, Go's demonstration stopped with uncertain spend; Temporal's selected retry policy repeated the call.
- Those outcomes reflect different retry policies. Neither framework guarantees recovery of an unrecorded paid response.
- The experiment did not test power loss, Wiki commit safety, throughput or production hosting.

All four cases passed.
Kyle said restart recovery was not a major requirement and chose pure Go, REST, SSE and an in-process queue.
The experiment code, generated reports and downloaded runtime were discarded; Temporal and checkpoint machinery are not product dependencies.

## Remaining workflow review

Specify queue limits and full-queue admission, cancellation, clarification, progress retention and interruption outcomes before implementation.
Review safe Wiki updates and browser-client access separately from research execution.
Kyle requested proposals for revision refusal, research preconditions and Markdown versus SQLite storage, followed by Palme's review.
After that review, Kyle selected authoritative Markdown with immutable source captures for evidence, notes and human-review records.
Direct file edits require stopping `serve`; viewing files remains unrestricted, and live changes use the API.
Publication details, offline edit validation and durability remain decisions in #116; these selections do not accept the full storage proposal.
Non-loopback binding is being reconsidered in #119; it must not be recommended, and its behavior remains undecided.

| Record | Current scope |
| --- | --- |
| [#113](https://github.com/mardwerk/tower-generator/issues/113) | Same-character coordination; busy-mark details remain proposed |
| [#115](https://github.com/mardwerk/tower-generator/issues/115) | API save and review revision preconditions |
| [#116](https://github.com/mardwerk/tower-generator/issues/116) | Markdown authority and direct-edit boundary selected; publication and durability pending |
| [#117](https://github.com/mardwerk/tower-generator/issues/117) | Completed by retirement and documentation repair |
| [#118](https://github.com/mardwerk/tower-generator/issues/118) | Research revision preconditions; original website-field claim corrected |
| [#119](https://github.com/mardwerk/tower-generator/issues/119) | Host, Origin and separate browser-client behavior |
