# Tower Generator product rules

This document owns the API and storage rules Kyle selected on 2026-10-03 through [#106](https://github.com/mardwerk/tower-generator/issues/106).
The [realignment record](REALIGNMENT.md) contains reset history and the recovery evaluation.

## Established direction

Tower Generator owns reusable character research, its portable local Wiki and later Tower design for a selected Profile.
Complete the Go research and Wiki workflow first; Tower design follows later.

One backend owns research, Wiki lookup, saving and review through the same API.
Websites, scripts and possibly Towerright later are clients; requests need no browser session.
Add behavior or dependencies for a defined workflow.

## Research and evidence rules

Follow [RESEARCH.md](RESEARCH.md) for identity, evidence, budgets and protection of retained work.
The [candidate Wiki format](wiki-format.md) is historical schema evidence, not an accepted Go format.

## API workflow outline

The backend owns provider access, validation and file commits, using one Wiki location chosen at startup.
Keep credentials in protected configuration and out of evidence, API responses, prompts, generated artifacts and replies.
Treat Markdown and YAML as data, not executable page content; keep dependencies and citations within the selected Wiki.

| Operation | Contract |
| --- | --- |
| Lookup and read | Return saved identity, evidence, review state and revision by query or key, without network research. |
| Research | Apply [preconditions](#selected-research-preconditions), [admission](#selected-research-admission), [paid provider retry rules](#selected-paid-provider-retry-policy) and [publication](#selected-research-publication-revisions). |
| Inspect operations | Use [read-only status, SSE and the recent list](#selected-operation-visibility-and-manual-recovery). |
| Save edit or record review | Apply the [save and review policy](#selected-save-and-review-policy). |
| Read configuration | Return effective safe settings and categories without secret values. |
| Generate, later | Follow the future [generation contract](#later-clients-and-generation). |

Keep progress separate from the final outcome.
Ambiguous names return a clarification-needed outcome; each frontend decides how to ask the person.
Report known consumed usage even when identity discovery ends in ambiguity or a busy outcome.

## Revision preconditions

Saves, reviews and research on a saved character require its saved key and the caller's revision in JSON `expectedRevision`.
A missing revision returns `400 revision_required`; a stale revision returns `409 revision_conflict`.
Both refusals leave saved content and human-review state unchanged; research checks them before provider work.
Never substitute the latest server revision for the caller's revision.

## Selected save and review policy

Kyle accepted the [#115](https://github.com/mardwerk/tower-generator/issues/115) policy on 2026-10-03.
Apply the [revision preconditions](#revision-preconditions) before saving or reviewing.

Check the [research reservation](#selected-research-admission) and commit through the [Wiki publication boundary](#selected-wiki-publication).
Two writes based on the same revision cannot both commit.
Only explicit review requests record human review.
Ordinary saves cannot supply human approval or make changed content appear approved by an earlier review.

Clients retain refused drafts and read the latest saved state separately for reconciliation.
They must not silently replace `expectedRevision` and retry the write.

## Selected research preconditions

Kyle accepted the [#118](https://github.com/mardwerk/tower-generator/issues/118) revision rule on 2026-10-03.
Apply the [revision preconditions](#revision-preconditions) to existing-character research.

Name queries first check saved names and confirmed aliases locally, before any provider call.
A saved match returns the existing key and revision so the client can read the saved character and request research deliberately.
Discovery must not silently become paid research on an existing character.

Kyle selected keeping the saved canon scope for research on an existing character in #118.
Use the entry's saved scope; an omitted or matching scope hint is allowed.
Refuse a conflicting scope hint before provider work and leave saved content unchanged.
Changing the canon boundary requires a separate explicit workflow; repeat research does not change it.

## Selected research admission

Kyle selected reservation at admission in [#118](https://github.com/mardwerk/tower-generator/issues/118) on 2026-10-03.
Reserve a known character when its research request is accepted, including time waiting in the queue.
Keep the reservation until the operation finishes or stops, including its publication work.
Refuse another research, save or review for that character while reserved; reads and work on other characters remain available.

Accept research only when execution capacity or a queue slot is available.
If both are full, refuse without provider work or a character reservation.
Coordinate revision checks, capacity checks and reservation so a refused request leaves no reservation or occupied queue slot.

Kyle selected configurable research limits in #118 on 2026-10-03.
The planned `serve` command defaults to two executing operations and ten queued operations.
Set these limits at startup with `--worker-size <Y>` and `--queue-size <X>`, for example `serve --worker-size 2 --queue-size 10`.
Queue size counts waiting operations separately from executing workers.
These are initial defaults, not measured capacity; revisit them using operation durations, queue waits, capacity refusals and provider limits.
A larger queue reserves more known characters while they wait, which can delay saves and reviews without increasing execution concurrency.
These counts do not select the lifecycle or capacity of clarification waits.

Keep connection capacity, active research capacity, per-run budget and provider limits separate.
A capacity queue is different from waiting for a busy character.
Cooperative same-character research remains a preference; combining evidence and budgets is not a selected design.

## Selected research publication revisions

Kyle selected a new character revision for every successful research publication in [#118](https://github.com/mardwerk/tower-generator/issues/118) on 2026-10-03.
Advance the revision even when a run finds nothing new; the entry still records research provenance and known provider usage.
Preserve applicable human reviews and retained evidence; a new revision does not itself record human approval.
A new research request using the prior revision is then refused as stale before provider work.

Failure or interruption before entry replacement can leave the old revision valid, so revision checks alone do not prevent another paid attempt.
After replacement, follow the [selected Wiki publication failure rules](#selected-wiki-publication), including reporting the visible new revision when durability is unconfirmed.

## Selected operation visibility and manual recovery

Kyle selected manual recovery initially in [#118](https://github.com/mardwerk/tower-generator/issues/118) after the [Astra xhigh review](https://github.com/mardwerk/tower-generator/issues/118#issuecomment-5972581956).
Accepted research returns a server-generated operation ID.
Expose read-only operation status and SSE access, plus a simple bounded recent-operation list from the same in-process records.
Operation reads report the outcome, known usage and any published character revision needed for inspection.

Kyle selected configurable in-memory history and SSE reconnects to current status in #118 on 2026-10-03.
Keep all queued and running operation records plus the latest 100 finished outcomes, ordered by completion.
Make the limit for finished outcomes configurable at startup; use no time expiry.
Later completions evict the oldest finished records when the limit is exceeded; server restart loses the in-process records.
On SSE reconnect, send the operation's current complete status, including its final outcome if retained.
Do not replay missed progress updates.

Clients may reconnect SSE and repeat status reads, but must not automatically resubmit research after an uncertain response or failure.
A person inspects the available outcome and deliberately chooses any new research attempt, which passes normal admission checks and may spend another budget.
An unavailable record or absence from the recent list does not prove that research never ran.
Lost admission responses, record eviction and server restart can leave the outcome uncertain; the list supports inspection without guaranteeing exact recovery.

Defer caller retry IDs until a named client needs automatic submission replay.

## Selected paid provider retry policy

Kyle selected no automatic application or SDK retries for paid calls in [#118](https://github.com/mardwerk/tower-generator/issues/118) on 2026-10-03.
Do not automatically repeat a paid call after a timeout, a dropped or invalid response, `429` or `5xx` errors.
If a paid step's outcome is uncertain, stop further paid work in that operation.
Report known usage separately from unknown usage; unknown usage does not mean zero.

A transient failure may require a deliberate new attempt under the [manual recovery policy](#selected-operation-visibility-and-manual-recovery).
Revisit narrowly defined retries when a selected provider documents that a rejected request performed no paid work.
Verify SDK retry settings and transport replay before claiming adapter compliance; exact adapter behavior remains implementation-verification work.

## Selected Wiki publication

Kyle authorized the [#116](https://github.com/mardwerk/tower-generator/issues/116) publication approach after comparing its long-term maintenance and extension options.
Authoritative Markdown owns character evidence, manual notes and human-review records, with immutable source captures.
Keep readable local files and retained passages under `wiki/<series>/<character>/`, ignored by Git; SQLite authority is not selected.
One versioned entry document is the commit point for mutable character state, with explicit references to immutable source captures.
Keep publication and integrity checks in one storage component; research, review and HTTP handlers must not write Wiki files independently.
Validate each complete update before committing it.

Use the API for changes while `serve` runs; stop it before directly editing Wiki files.
Viewing the files remains unrestricted.
A proposed fingerprint must cover every file the backend can replace, or the backend must leave unrelated files untouched.
Revision checks cannot protect an uncoordinated edit made after the final check and before replacement.

Install and validate new captures before publishing the entry that references them.
Synchronize capture files and directory entries, including newly created parent directories, before entry publication.
Serialize writers for the same character and check the revision before replacement.
Independent characters must not share a lock through file synchronization.
Prepare and synchronize the new entry file on the same filesystem, atomically replace the current entry, then synchronize its containing directory before acknowledging durable completion.

Readers read the entry once and use that complete version's immutable references without read locks.
Keep those captures available while readers can still use them; do not rewrite or delete them while serving.
Do not replace or move a character directory during publication.

A failure before entry replacement leaves the previous entry current, possibly with unused captures.
After replacement, a synchronization failure leaves the new revision visible with durability unconfirmed.
Report that revision and the storage failure; do not claim no write occurred, silently roll back or repeat paid research.

Initial verification covers process interruption and publication on tested local Linux filesystems.
File and directory synchronization is required, but the reported process-kill tests do not establish power-loss behavior or other platform guarantees.
The product implementation must verify its own publication and error handling before claiming those guarantees.

A revision token detects stale changes; it does not promise later retrieval of an overwritten entry.
Select complete entry-snapshot retention before promising historical lookup, rollback or generation from an earlier version.
Snapshots can be retained around the same publication boundary; a derived search index can be rebuilt from authoritative Markdown when needed.

## Selected transport and execution

Kyle selected one pure Go executable with REST, SSE and a bounded in-process research queue.
`serve` provides startup and configuration; there is no independent operational CLI or persistent execution queue.
Restart recovery is not a major requirement, and automatic job resumption is not promised.

## Selected local API access

Kyle accepted the [#119](https://github.com/mardwerk/tower-generator/issues/119) scope on 2026-10-03.
Initially bind only to loopback addresses, with a configurable numeric port and no authentication.
Keep Host validation and browser-origin protection even on loopback.
Require JSON for writes; GET and SSE are read-only.

A local website uses a same-origin proxy to reach the API.
Direct cross-origin browser access and non-loopback binding are deferred until a named client needs them.
Review authentication and trusted origins before expanding that access.

## Later clients and generation

Earlier frontend preferences remain context for a future client contract: minimal Create, optional inputs behind an expansion, centered Create, Library, Wiki and Generations navigation, and Settings on the right.
Each page has one visible heading and no subheadings; unused destinations were to remain disabled.
The reset removes their implementation and does not add frontend work to backend acceptance.

Generation consumes a recorded character revision under the [Profile ownership and BTD6 attribution rules](PROFILE.md).

## Workflow acceptance review

Review a fresh character, confirmed alias, ambiguous identity, failed source and repeated research before implementation.
Specify inputs, outcomes, progress, saved files and budget behavior together for each case.
Add concurrent same-character requests, independent characters, missing and stale save, review and research revisions, external edits and interruption during commit.
Test omitted, matching and conflicting saved-scope hints, malformed Host authorities, IPv6, hostname access, script/native callers and SSE reconnects.
These cases must not imply support for unselected API access.
Review disconnect and restart behavior separately; successful saved Wiki content must not depend on a browser remaining connected.

Product implementation waits for complete workflow acceptance through [#106](https://github.com/mardwerk/tower-generator/issues/106).
The following choices remain open; earlier proposals and merged proposal text do not select them.

| Review | Remaining choices |
| --- | --- |
| [#115](https://github.com/mardwerk/tower-generator/issues/115) | Revision coverage, review invalidation and complete save/review schemas. |
| [#116](https://github.com/mardwerk/tower-generator/issues/116) | Entry format, offline validation, review staleness, historical retention and single-writer enforcement. |
| [#118](https://github.com/mardwerk/tower-generator/issues/118) | Queue and history configuration validation, scheduling and refusal responses, cancellation, clarification continuation, unknown-identity coordination, discovery scope, scope-hint representation and matching, and SSE framing and observer limits. |
| [#119](https://github.com/mardwerk/tower-generator/issues/119) | Exact Host authorities, address families and numeric-port behavior. |
| [#106](https://github.com/mardwerk/tower-generator/issues/106) | Routes, strict request decoding, operation records, provider/model selection, budgets, extraction dependencies, packaging and implementation verification. |
