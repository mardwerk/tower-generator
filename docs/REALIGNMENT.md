# API-only reset and architecture review

Review prepared on 2026-10-03 from ten parallel GPT-6.1-sol reviews, followed by focused recovery comparisons.
The [product rules](PRODUCT.md) own the public requirements; [#106](https://github.com/mardwerk/tower-generator/issues/106) coordinates acceptance.
This document records decisions, evidence and proposals without selecting an execution engine.

## Owner decisions and reset boundary

Kyle selected one Go `serve` backend for research, lookup, saving and review, with all operational clients using its API.
No browser session is required; the backend retains Wiki files and tracks running research.
Ambiguous identities require clarification, and each frontend decides how to obtain the answer.

Kyle confirmed removing the entire old product implementation, including all `src/`, the Lab files, Node packages and the Python research/Wiki CLI.
He selected retirement of the old Usopp experiment and deleted Default Profile rather than restoration.
Retain independent research evidence and helpers; a new Profile and Tower generation remain later work.
Prepare review PRs, then push the agreed reset directly to Tower Generator main only after Kyle reviews the concrete result.
Planning changes remain in a separate PR.

The reset removes `setup.py`, `requirements.txt`, `package.json`, `pnpm-lock.yaml`, Prettier configuration, `default/usopp.json`, `sources/usopp.json`, and the Usopp authoring guides.
The locally moved `scripts/setup.py` contains the same retired setup code and is also retired when applying the reset to the main checkout.
Do not run retired setup or Node checks after removal.

Keep local credentials and collected `wiki/` files untouched.
Remove only known retired local tooling and generated artifacts during the agreed main-checkout cleanup; do not treat every ignored file as disposable.
Standalone active helpers belong under `scripts/`; system Python is sufficient for standard-library scripts, with uv optional and no required project virtual environment.
Manifest-covered research Python remains in its original archive location to preserve provenance.

The [latest prototype revision](https://github.com/mardwerk/tower-generator/tree/ea114b2139abf46318a59f632a58fc035f64a5a7) preserves the Python application and Lab.
The [pre-deletion Profile revision](https://github.com/mardwerk/tower-generator/tree/c7a71f25147ee871bf48790e6e5b0cd6515d8cfa/default/profile) preserves all 50 Profile files.
The imported [research snapshot](research/unit-design-research-snapshot/README.md) remains byte-for-byte evidence, not a product specification.

## Separate communication from execution

REST accepts commands and reads saved content and operation status.
SSE can deliver progress updates; polling can expose the same status to scripts.
Neither mechanism selects how work persists, retries or survives a restart.
Temporal can orchestrate execution behind the same REST and SSE API.

A proposed operation ID lets clients find work independently of a browser session.
Accepted research should use a server-owned context; [Go request contexts](https://pkg.go.dev/net/http#Request.Context) end when the client disconnects or the handler returns.
A proposed caller request ID distinguishes retransmission after a lost response from deliberate repeated research with a fresh budget.
The lifetime of that deduplication guarantee must match the chosen operation retention policy.

For SSE, current-state snapshots on connection and reconnect are the smallest proposal.
Replaying every missed event needs explicit event retention and cursor expiry.
[Last-Event-ID](https://html.spec.whatwg.org/multipage/server-sent-events.html) supports reconnect protocols but does not itself store events or recover jobs.

## Execution comparison

Kyle requested further evaluation before choosing restart guarantees.
These are alternatives for execution behind the same API, not accepted architecture.

| Choice | Stored operation state | Restart behavior | Runtime ownership |
| --- | --- | --- | --- |
| Process-local Go | In-memory inputs, progress and outcome | Unfinished research ends; the client must submit another request | One Go process and the Wiki directory |
| Narrow durable Go | Versioned operation records and immutable completed step responses | Resume recorded steps; pause interrupted calls with uncertain outcomes | One Go process, recovery code and a local operational store |
| Temporal | Workflow history plus retained large responses where needed | Replay recorded results and continue pending steps or waits | API/Worker code plus Temporal Cloud or an operated Temporal Service |

An intermediate durable-record option can preserve accepted work and status while marking running jobs interrupted on restart.
That is less than checkpoint resumption and should not be advertised as recovery of paid steps.
A durable queue can preserve waiting work across restart; an in-memory queue cannot.
Neither reconstructs a completed provider response that was never saved.

| Scenario | Process-local Go | Narrow durable Go | Temporal |
| --- | --- | --- | --- |
| Browser or script disconnects | Continues with a job-owned context | Same | Same |
| Restart between recorded paid steps | Starts over on a new request | Reuses recorded responses | Replays recorded Activity results |
| Provider succeeds before the coordinator records success | Charge and result may be unknown | Attempt record exposes uncertainty | Activity retry can repeat the charge |
| Crash during Wiki replacement | Needs independent file recovery | Needs independent file recovery | Commit Activity needs independent file reconciliation |
| Provider throttles for minutes | In-process bounded wait | Saved retry time and budget | Durable timer or bounded retry |
| Clarification arrives tomorrow | Requires the process to remain alive or a new request | Can retain a pending question | Can retain a waiting Workflow |

Job recovery, billing safety and Wiki recovery are separate guarantees.
Neither durable option guarantees a single paid charge if the provider succeeds before its response is recorded.
An enforced provider idempotency key or retrievable provider operation can close that gap; otherwise the contract must account for uncertain spend.
Temporal [Activities must tolerate repetition](https://docs.temporal.io/activity-definition#idempotency), and its [default retry policy](https://docs.temporal.io/encyclopedia/retry-policies) needs explicit limits for paid calls.

A narrow Go runner could use fixed identity/search, evidence-fetch, synthesis/verification, validation/staging and commit stages.
It needs persisted attempts, original budgets, completed responses, pending answers, retry times and a versioned execution record.
Old records must be migrated, supported by their old executor or stopped explicitly after a code change.
Temporal handles execution history but still requires [replay-compatible code and Worker versioning](https://docs.temporal.io/production-deployment/worker-deployments/worker-versioning).

Prefer process-local execution only if repeating interrupted paid work is acceptable.
Consider narrow durable Go for one fixed workflow and one local writer if its recovery rules remain small.
Consider Temporal when durable human waits, independent retries, deployments and multiple Workers make custom recovery and versioning substantial.
These are review recommendations, not a selection.

[Temporal Cloud pricing](https://temporal.io/pricing), checked on 2026-10-03, lists no base monthly fee, usage-based Actions and history storage, and Developer support at 10% of usage.
Worker hosting and research-provider charges remain separate.
Self-hosting adds persistence, backups and upgrades; the [embedded SQLite server](https://docs.temporal.io/self-hosted-guide/embedded-server) is for development and testing.
Measure research cost, duration, response sizes, clarification delays, restart frequency and expected concurrency before relying on cost estimates.

## Proposed comparative experiment

A small experiment can test recovery before a product architecture is selected.
Use a fake paid provider with its own invocation ledger and the same small research sequence in narrow Go and Temporal implementations.
Do not call a paid model or build the product API.

Inject interruption after admission, after provider success but before recording, after response persistence, after moving the old Wiki directory, and after installing its replacement.
Replay the same request ID and inspect provider call counts, operation state and Wiki hashes.
Add an overnight clarification wait and a code-version change while work is pending.
Compare recovery correctness, custom application code, migration obligations and deployment setup.
This experiment remains proposed; the reset does not introduce either runner.

## Coordination and browser review

The simplest proposed same-character policy is a busy mark plus mandatory save/review revisions.
Admission at overall capacity may refuse immediately or use a bounded visible queue; this is separate from queueing a busy character.
Cooperative research requires evidence merging and budget rules beyond the busy mark.

API reads can remain coherent with a short storage lock during final commit, while network work stages outside it.
That proposal permits brief read waiting, not waiting for the entire research run.
Recoverable staging and backup metadata can repair a killed process before the API serves requests again.
Portable [os.Rename](https://pkg.go.dev/os#Rename) does not atomically replace a populated directory; power-loss durability requires a separate platform and synchronization contract.

The prototype fingerprint omitted arbitrary preserved files such as `notes.md`, although research replaced the whole directory.
Fingerprint every file that can be replaced or leave unrelated files untouched.
Neither a busy mark nor Temporal coordinates an arbitrary external editor in the final check-to-write interval.

The proposed localhost-or-literal-IP Host check rejects an attacker hostname even if it rebinds to the LAN address.
Exact configured hostnames such as `machine.local` need a separate choice; accepting arbitrary names because they resolve locally would weaken the rule.
A separate browser frontend on another port has a different [origin](https://url.spec.whatwg.org/#origin).
Choose an explicit trusted-origin [CORS policy](https://fetch.spec.whatwg.org/#cors-protocol) or a same-origin client proxy; nonbrowser callers may omit Origin.
Keep Kyle's selected unauthenticated binding behavior unchanged.

## Linked issue disposition

The reset preserves unresolved requirements instead of treating prototype deletion as implementation.

| Record | Reset treatment and remaining review |
| --- | --- |
| [#106](https://github.com/mardwerk/tower-generator/issues/106) | API-only scope supersedes its older CLI wording; keep workflow acceptance open |
| [#113](https://github.com/mardwerk/tower-generator/issues/113) and [#114](https://github.com/mardwerk/tower-generator/pull/114) | Preserve the coordination recommendation as a proposal until explicitly accepted |
| [#115](https://github.com/mardwerk/tower-generator/issues/115) | Retiring the CLI removes that implementation; mandatory API save/review revisions remain proposed acceptance cases |
| [#116](https://github.com/mardwerk/tower-generator/issues/116) | Keep coherent reads and interrupted commit recovery in Go storage review |
| [#117](https://github.com/mardwerk/tower-generator/issues/117) | Owner selected retirement; remove broken setup and current Profile links, defer replacement |
| [#118](https://github.com/mardwerk/tower-generator/issues/118) | Correct the website-field claim and decide whether existing-entry research starts from current state or an opened revision |
| [#119](https://github.com/mardwerk/tower-generator/issues/119) | Keep Host, Origin, hostname and separate-browser-client cases for Go acceptance |

The code linked in #118 sent a revision on manual collection, not on name-only research.
An arbitrary extra revision field would still be ignored; retiring that code does not select the new API precondition.
The [pinned frontend](https://github.com/mardwerk/tower-generator/blob/ea114b2139abf46318a59f632a58fc035f64a5a7/src/web/app/app.tsx#L169-L182) preserves the evidence.

## Next owner review

Review this reset PR independently of the execution-engine choice.
Before implementing Go, decide restart guarantees, clarification continuation, admission at capacity, cancellation, progress retention and browser-client access.
Then specify one name-to-Wiki workflow and its concrete acceptance cases.
No unresolved architecture choice becomes accepted merely by merging the reset documentation.
