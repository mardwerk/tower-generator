# Tower Generator product rules

Kyle selected an API-only Go backend in [#106](https://github.com/mardwerk/tower-generator/issues/106#issuecomment-5969964581) and [#113](https://github.com/mardwerk/tower-generator/issues/113).
This outline records the 2026-10-03 realignment and separates owner decisions from proposals awaiting review.
The [reset and architecture review](REALIGNMENT.md) records the removal scope, evidence and open choices.
Kyle subsequently selected pure Go with REST, SSE and a bounded in-process queue on 2026-10-03.
Product implementation waits for an agreed research workflow.

## Established direction

Tower Generator owns reusable character research, its portable local Wiki and later Tower design for a selected Profile.
Complete the Go research and Wiki workflow first; Tower design follows later.

- One pure Go executable exposes REST operations and SSE progress through `serve`, with a bounded in-process research queue. Startup and configuration remain; there is no independent operational CLI.
- Research, Wiki lookup, saving and review all use the same API. Websites, scripts and possibly Towerright later are clients.
- Requests need no browser session. The backend retains Wiki files and tracks research; detailed interruption and retention behavior remains open.
- Default binding is `localhost:<port>` without authentication. Optional non-loopback binding is under review in [#119](https://github.com/mardwerk/tower-generator/issues/119) and must not be recommended.
- Research blocks saves and reviews for that character. Reading and work on other characters remain available.
- Ambiguous names return a clarification-needed outcome. Each frontend decides how to ask the person.
- Remove the old Python application, Node tooling and entire web tree. Git history preserves the former generator and Lab.
- Retire the old Usopp experiment and deleted Default Profile packaging, including its setup script. A new Profile is deferred.
- Retain independent research evidence and helper scripts. Helpers are outside the main product.

Keep the product simple and add behavior or dependencies for a defined workflow.
The selected queue runs inside the Go backend.
Restart recovery is not a major requirement; no automatic job-resumption guarantee is established by this choice.

## Research and evidence rules

Research in English independently of Profiles and game mechanics.
Resolve the canonical character and series without requiring a series input by default.
Reuse confirmed identities despite aliases or older folder names; similarity alone must not merge different characters.
Make canon scope explicit and distinguish versions where the evidence requires it.

Retain abilities, traits, equipment, named techniques, conditions, limitations and uncertainty with cited passages.
Keep source facts, model verification, generalized classifications and human review distinct.
Jev or a comparable classification method remains a possible approach, not established canon or a game mapping.

Repeated research receives the same allocated budget for verification, correction and expansion.
It preserves retained evidence, manual notes and human-reviewed records; every run need not produce an improvement.
Missing fresh evidence does not itself refute or delete a prior finding.
Report blocked sources, incomplete excerpts, contradictions, effective budget and known provider usage.

Keep local human-readable files under `wiki/<series>/<character>/`, ignored by Git, with retained passages beside the character.
Kyle selected authoritative Markdown with immutable source captures in [#116](https://github.com/mardwerk/tower-generator/issues/116).
Character evidence, manual notes and human-review records belong to the Markdown Wiki; SQLite authority is not selected.
Offline lookup reads saved evidence without network research.
The [research rules](RESEARCH.md), [candidate Wiki format](wiki-format.md) and [classification definitions](research-categories.yaml) retain useful prototype evidence for the new contract.

## API workflow outline

The outline combines selected requirements with operations still awaiting review.
Route names and complete JSON schemas are not selected.
The backend owns provider access, validation and file commits, using one Wiki location chosen at startup.

| Operation | Input | Outcome |
| --- | --- | --- |
| Lookup and read | Query or saved character key | Saved identity, evidence, review state and revision without research |
| Research | Character query for discovery, or saved character key and JSON `expectedRevision`; scope and budget inputs need review | Saved evidence and change summary, existing identity to read, or an explicit clarification, busy or failure outcome |
| Save edit | Saved character key, edited content and JSON `expectedRevision` | Validated update or refusal without overwriting saved work |
| Record review | Saved character key, JSON `expectedRevision` and explicit target | Human decision tied to that content; findings and classifications remain distinct |
| Read configuration | Selected safe settings | Effective configuration and categories without secret values |
| Generate, later | Character revision and selected Profile | Tower content and provenance under a future generation contract |

Keep progress separate from the final outcome.
Queue limits and admission when the queue is full, cancellation, duplicate-request handling, clarification continuation, retries and operation retention need review.
Report known consumed usage even when identity discovery ends in ambiguity or a busy outcome.

## Selected save and review policy

Kyle accepted the [#115](https://github.com/mardwerk/tower-generator/issues/115) policy on 2026-10-03.
Saves and reviews address a saved character by its key and require the caller's revision in a JSON `expectedRevision` field.
A missing revision returns `400 revision_required`; a stale revision returns `409 revision_conflict`.
Both refusals leave saved content and human-review state unchanged.

Serialize competing saves and reviews briefly per character, checking the research reservation and revision before committing through the selected Wiki publication boundary.
Two writes based on the same revision cannot both commit; independent characters do not share a lock through file synchronization.
Only explicit review requests record human review.
Ordinary saves cannot supply human approval or make changed content appear approved by an earlier review.

Clients retain refused drafts and read the latest saved state separately for reconciliation.
They must not silently replace `expectedRevision` and retry the write.
Routes, complete request schemas, exact revision coverage and review invalidation details remain specification work.

## Selected research preconditions

Kyle accepted the [#118](https://github.com/mardwerk/tower-generator/issues/118) revision rule on 2026-10-03.
Research on a saved character requires its saved key and the caller's revision in JSON `expectedRevision`, using the same convention as saves and reviews.
A missing revision returns `400 revision_required`; a stale revision returns `409 revision_conflict`.
These precondition refusals happen before provider work and leave saved state unchanged.

Name queries first check saved names and confirmed aliases locally, before any provider call.
A saved match returns the existing key and revision so the client can read the saved character and request research deliberately.
Discovery must not silently become paid research on an existing character.
This supersedes the earlier proposal that existing-character research simply starts from the current server revision.

Queue reservation timing and retry handling remain separate choices.
Scope hints on saved characters, revision advancement after research without new findings, complete schemas and strict request decoding still need specification.
This decision does not establish request deduplication or guarantee that retrying a failed or interrupted operation avoids another paid call.

## Proposed coordination and storage

[#113](https://github.com/mardwerk/tower-generator/issues/113) and merged [PR #114](https://github.com/mardwerk/tower-generator/pull/114) propose a per-character busy mark and revision checks.
Their merged proposal text is evidence for discussion, not blanket acceptance of every design choice.

The proposed policy refuses a second change to a busy character, including another research request.
The selected save, review and existing-character research policies require the revision the caller read.
Research reservation timing remains under review; research must not substitute the latest revision for the caller's revision.
Cooperative same-character research remains a preference; combining evidence and budgets is not a selected design.

Validate each complete update before committing it.
A persistent execution queue is not selected.

Use the API for changes while `serve` runs; stop it before directly editing Wiki files.
Viewing the files remains unrestricted.
Offline edit validation and review-staleness handling still need the storage specification.
A proposed fingerprint must cover every file the backend can replace, or the backend must leave unrelated files untouched.
Revision checks cannot guarantee preservation of an uncoordinated edit made after the final check and before replacement.
Enforcement of one backend writer per Wiki directory remains a specification choice.

Keep connection capacity, active research capacity, per-run budget and provider limits separate.
A capacity queue is different from waiting for a busy character.
Never silently reduce a research budget, change the model or omit verification because another run or provider limit intervenes.

## Selected Wiki publication

Kyle authorized the [#116](https://github.com/mardwerk/tower-generator/issues/116) publication approach after comparing its long-term maintenance and extension options.
One versioned entry document is the commit point for mutable character state, with explicit references to immutable source captures.
Keep publication and integrity checks in one storage component; research, review and HTTP handlers must not write Wiki files independently.

Install and validate new captures before publishing the entry that references them.
Synchronize capture files and directory entries, including newly created parent directories, before entry publication.
Serialize writers for the same character and check the revision before replacement.
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
Exact field layouts, revision coverage, offline validation, retention and writer-ownership enforcement remain specification work.

## Selected transport and execution

REST and SSE describe communication; Temporal can manage execution behind the same API.
Kyle selected pure Go execution with REST, SSE and a bounded in-process research queue after reviewing the recovery tradeoffs.
Temporal remains an evaluated alternative rather than the current backend.
The [comparison](REALIGNMENT.md) separates client disconnect, backend restart, uncertain paid calls and interrupted Wiki commits.

HTTP routes, SSE event retention, queue limits and full-queue admission, numerical limits, extraction dependencies and packaging remain open.
Hostname access and trusted browser origins need review in [#119](https://github.com/mardwerk/tower-generator/issues/119).
Same-origin checks alone would reject a separate frontend on another port; browser clients need an agreed proxy or CORS policy.

## Later clients and generation

Earlier frontend preferences remain context for a future client contract: minimal Create, optional inputs behind an expansion, centered Create, Library, Wiki and Generations navigation, and Settings on the right.
Each page has one visible heading and no subheadings; unused destinations were to remain disabled.
The reset removes their implementation and does not add frontend work to backend acceptance.

Research never names a Profile or stores Tower choices and numerical adaptations.
Generation consumes a recorded character revision and selects a Profile directory by path.
Preserve [Profile ownership and BTD6 attribution](PROFILE.md) when generation resumes.

## Workflow acceptance review

Review a fresh character, confirmed alias, ambiguous identity, failed source and repeated research before implementation.
Specify inputs, outcomes, progress, saved files and budget behavior together for each case.
Add concurrent same-character requests, independent characters, missing and stale save, review and research revisions, external edits and interruption during commit.
Review disconnect and restart behavior separately; successful saved Wiki content must not depend on a browser remaining connected.

[#106](https://github.com/mardwerk/tower-generator/issues/106) remains the coordination record.
[#115](https://github.com/mardwerk/tower-generator/issues/115), [#116](https://github.com/mardwerk/tower-generator/issues/116), [#118](https://github.com/mardwerk/tower-generator/issues/118) and [#119](https://github.com/mardwerk/tower-generator/issues/119) retain relevant contract questions after prototype removal.
Retirement and documentation repair address [#117](https://github.com/mardwerk/tower-generator/issues/117); they do not implement a replacement Profile.
