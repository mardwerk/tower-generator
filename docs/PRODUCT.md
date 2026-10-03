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
- Default binding is `localhost:<port>`. Explicit `0.0.0.0:<port>` is allowed; neither binding uses authentication.
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
Offline lookup reads saved evidence without network research.
The [research rules](RESEARCH.md), [candidate Wiki format](wiki-format.md) and [classification definitions](research-categories.yaml) retain useful prototype evidence for the new contract.

## Proposed API workflow

These operation contracts remain proposals; route names and JSON schemas are not selected.
The backend owns provider access, validation and file commits, using one Wiki location chosen at startup.

| Operation | Input | Outcome |
| --- | --- | --- |
| Lookup and read | Query or saved character key | Saved identity, evidence, review state and revision without research |
| Research | Character query or saved identity, optional scope hints and lower budget | Saved evidence and change summary, or an explicit clarification, busy or failure outcome |
| Save edit | Character key, edited content and expected revision | Validated update or refusal without overwriting saved work |
| Record review | Character key, expected revision and explicit target | Human decision tied to that content; findings and classifications remain distinct |
| Read configuration | Selected safe settings | Effective configuration and categories without secret values |
| Generate, later | Character revision and selected Profile | Tower content and provenance under a future generation contract |

Keep progress separate from the final outcome.
Queue limits and admission when the queue is full, cancellation, duplicate-request handling, clarification continuation, retries and operation retention need review.
Report known consumed usage even when identity discovery ends in ambiguity or a busy outcome.

## Proposed coordination and storage

[#113](https://github.com/mardwerk/tower-generator/issues/113) and merged [PR #114](https://github.com/mardwerk/tower-generator/pull/114) propose a per-character busy mark and revision checks.
Their merged proposal text is evidence for discussion, not blanket acceptance of every design choice.

The proposed policy refuses a second change to a busy character, including another research request.
Saves and reviews require the revision they read; missing or stale revisions cause no write.
Research reserves the confirmed identity and reads its current revision before evidence work.
Cooperative same-character research remains a preference; combining evidence and budgets is not a selected design.

Validate each complete update before committing it.
Review whether API readers may wait briefly during commit, and how startup recovers an interrupted multi-file replacement.
These file guarantees are needed independently of research restart recovery and independently of Temporal.

External edits require a stated boundary.
A proposed fingerprint must cover every file the backend can replace, or the backend must leave unrelated files untouched.
Revision checks cannot guarantee preservation of an uncoordinated edit made after the final check and before replacement.
One writer per Wiki directory, enforcement of that rule and supported durability guarantees remain specification choices.

Keep connection capacity, active research capacity, per-run budget and provider limits separate.
A capacity queue is different from waiting for a busy character.
Never silently reduce a research budget, change the model or omit verification because another run or provider limit intervenes.

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
Preserve [td-profile ownership and Profile rules](PROFILE.md) and [Atlas attribution](BTD6-REFERENCE.md) when generation resumes.

## Workflow acceptance review

Review a fresh character, confirmed alias, ambiguous identity, failed source and repeated research before implementation.
Specify inputs, outcomes, progress, saved files and budget behavior together for each case.
Add concurrent same-character requests, independent characters, missing and stale save revisions, external edits and interruption during commit.
Review disconnect and restart behavior separately; successful saved Wiki content must not depend on a browser remaining connected.

[#106](https://github.com/mardwerk/tower-generator/issues/106) remains the coordination record.
[#115](https://github.com/mardwerk/tower-generator/issues/115), [#116](https://github.com/mardwerk/tower-generator/issues/116), [#118](https://github.com/mardwerk/tower-generator/issues/118) and [#119](https://github.com/mardwerk/tower-generator/issues/119) retain relevant contract questions after prototype removal.
Retirement and documentation repair address [#117](https://github.com/mardwerk/tower-generator/issues/117); they do not implement a replacement Profile.
