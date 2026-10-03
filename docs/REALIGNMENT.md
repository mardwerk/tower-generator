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

Specify queue limits and refusal responses, cancellation, clarification, progress retention and interruption outcomes before implementation.
Review safe Wiki updates and browser-client access separately from research execution.
Kyle requested proposals for revision refusal, research preconditions and Markdown versus SQLite storage, followed by Palme's review.
After that review, Kyle selected authoritative Markdown with immutable source captures for evidence, notes and human-review records.
Direct file edits require stopping `serve`; viewing files remains unrestricted, and live changes use the API.
Kyle authorized the single-entry publication approach after a comparison focused on long-term maintenance and extension.
The [selected publication rules](PRODUCT.md#selected-wiki-publication) install immutable captures first, serialize writers and atomically replace the entry while readers use a complete old or new version.
File and directory synchronization is required; failure after replacement reports visible publication with durability unconfirmed.
Initial verification targets tested local Linux filesystems; power loss and other platforms require their own evidence.
Historical entry retention can be added around this boundary, but previous entry retrieval is not currently promised.
Exact format, offline validation, retention and implementation verification remain work in #116.
Kyle accepted the [save and review policy](PRODUCT.md#selected-save-and-review-policy) in #115.
Saves and reviews use saved character keys and JSON `expectedRevision`; missing revisions return 400 and stale revisions return 409 without changing saved work.
Competing writes serialize per character, only explicit review requests record human review, and clients retain refused drafts.
Complete schemas and revision coverage remain to be specified and verified.
Kyle accepted the [research preconditions](PRODUCT.md#selected-research-preconditions) in #118.
Saved-character research requires its saved key and JSON `expectedRevision`; missing revisions return 400 and stale revisions return 409 before provider work, leaving saved state unchanged.
Name queries check saved names and confirmed aliases locally and return an existing key and revision instead of silently commissioning research.
Kyle selected [reservation at admission](PRODUCT.md#selected-research-admission) for known characters, including time waiting in the queue.
While reserved, another research, save or review is refused until the operation finishes or stops; reads and other characters remain available.
Full execution and queue capacity causes refusal without provider work or a character reservation.
Kyle selected [a new revision for every successful research publication](PRODUCT.md#selected-research-publication-revisions), including runs with no new findings.
The entry records research provenance and known usage while preserving applicable human reviews.
An unpublished failure can leave the old revision valid, so revision checks alone do not prevent another paid attempt.
Kyle selected [manual recovery initially](PRODUCT.md#selected-operation-visibility-and-manual-recovery), with server-generated operation IDs, read-only status/SSE and a simple bounded recent-operation list.
Clients may reconnect reads but must not automatically resubmit research after an uncertain response; a person deliberately chooses any new paid attempt.
Caller retry IDs are deferred until a named client needs automatic submission replay.
Exact retention limits, provider-call retries, unknown-identity coordination and numerical queue limits remain owner choices.
Kyle accepted [loopback-only API access](PRODUCT.md#selected-local-api-access) in #119, with a configurable numeric port, Host validation and browser-origin protection.
Writes require JSON, GET and SSE are read-only, and a local website uses a same-origin proxy.
Non-loopback binding and direct cross-origin access are deferred until a named client needs them, followed by authentication and trusted-origin review.
Exact Host authorities, address-family behavior and implementation verification remain specification work.

| Record | Current scope |
| --- | --- |
| [#113](https://github.com/mardwerk/tower-generator/issues/113) | Known-character admission reservation selected; remaining coordination details need specification |
| [#115](https://github.com/mardwerk/tower-generator/issues/115) | Save and review policy selected; complete schemas and implementation verification pending |
| [#116](https://github.com/mardwerk/tower-generator/issues/116) | Markdown authority, direct-edit boundary and publication selected; format and verification pending |
| [#117](https://github.com/mardwerk/tower-generator/issues/117) | Completed by retirement and documentation repair |
| [#118](https://github.com/mardwerk/tower-generator/issues/118) | Research preconditions, admission, publication revisions and manual recovery selected; remaining research contract and verification pending |
| [#119](https://github.com/mardwerk/tower-generator/issues/119) | Loopback-only access selected; exact Host/address behavior and implementation verification pending |
