# Tower Generator product outline

Working outline for review, recorded on 2026-10-02 in [#106](https://github.com/mardwerk/tower-generator/issues/106).
Kyle selected defining rough rules before further implementation.
The owner requirements below come from his requests; the proposed sections remain open for review.
Kyle selected Go for the backend on 2026-10-02.
On 2026-10-03 he replaced the independent operational CLI with an API-only `serve` backend ([#106](https://github.com/mardwerk/tower-generator/issues/106#issuecomment-5969964581), [#113](https://github.com/mardwerk/tower-generator/issues/113)).
Implementation remains pending the workflow specification.
Proposals from the [#107 review](https://github.com/mardwerk/tower-generator/issues/107) leave the owner requirements and the agreed backend direction unchanged.
The [#113 review](https://github.com/mardwerk/tower-generator/issues/113) proposes the same-character update policy in the concurrency rules.

## Purpose and current focus

Tower Generator should turn reusable character evidence into Towers that a person can understand, inspect and revise for a selected Profile.
The immediate focus is reliable research and local storage; generation remains later work.
Specify the product's behavior before further implementation or screen changes.
The existing application is an experiment that can inform the specification.

## Established owner requirements

1. Keep the product clean and simple. Add a feature or dependency only for a defined workflow.
2. Keep Create minimal: character input, primary action and small controls below it. Put optional inputs behind an explicit expansion.
3. Resolve the character's canonical name and series from research. Ask for disambiguation only when necessary; do not require a series field by default.
4. Retain the previous Lab design system and centered navigation: Create plus, Library, Wiki and Generations, with Settings on the right.
   Keep unused destinations disabled in place. Each page has one visible heading and no subheadings.
5. Research in English and retain reusable abilities, traits, equipment, evidence, limitations and uncertainty independently of game Profiles.
6. Reuse a confirmed existing character despite aliases or older normalized folder names. Similar names alone must not silently merge different identities.
7. Repeated research spends the same allocated budget on verification, correction and expansion.
   It must preserve prior evidence and human-reviewed work; an improvement is not guaranteed on every run.
8. Keep local, human-readable character files under `wiki/<series>/<character>/`, ignored by Git.
   The tool may read and update those files without keeping a hidden session or requiring a database.
9. Suggested generalized classifications may use Jev or a comparable decision model.
   Model classification and verification do not establish canon acceptance or game mechanics.
10. Use Go for the backend. A `serve` command exposes one API capable of handling concurrent clients and requests; there is no independent operational CLI.

## Proposed responsibilities

Use one core for character identity, research, lookup, file validation, saving and review.
Every client reaches it through the `serve` API, so all clients have the same behavior and storage rules.
The website handles forms, navigation, readable evidence and user actions; the backend owns durable files, provider access and validation.
Browser memory may hold an unsaved edit, but it must not be the only copy of saved research.
`serve` uses one Wiki location chosen when it starts, so API requests cannot write to other paths.
Both bindings keep today's browser protection, widened to IP addresses: reject a request whose browser origin differs from its host, or whose host name is neither localhost nor an IP address.
This stops a web page from driving the API through DNS rebinding.
Without authentication, anyone who can reach a `0.0.0.0` binding can read, edit, review and spend the research budget, so use it only on a trusted network.

## Agreed backend direction

Kyle recorded this direction in [#106](https://github.com/mardwerk/tower-generator/issues/106#issuecomment-5969964581) and [#113](https://github.com/mardwerk/tower-generator/issues/113) on 2026-10-03.
Use one Go executable whose `serve` command exposes the API. There is no independent operational CLI; process startup and configuration remain.
Research, lookup, edit and review clients call the API without a browser session: the website, scripts and possibly Towerright later.
`serve` binds to `localhost:<port>` by default and to `0.0.0.0:<port>` only when explicitly configured. Neither binding uses authentication.
While a character is being researched through the API, edits are disabled and saves and reviews for that character are refused. Reading and other characters remain available.
The backend keeps the local portable Markdown Wiki and protects manual notes, retained evidence and human-reviewed records.
Complete the Go Wiki and research workflow first; Tower design follows later.
Parallel requests that research one character cooperatively are a preference, not an accepted design.
Go supports concurrent network work and a compiled executable; the current Python CLI and Node API bridge remain the working experiment until migration.
The target API must run without Node or Python. Frontend build tooling can remain separate.
Serving the built website from the executable is a packaging proposal, not yet a specified requirement.

## Proposed concurrency rules

Handle multiple API connections while allowing independent characters and source fetches to progress concurrently.
Keep four limits separate, so more clients never create unlimited provider calls and load never changes what a run may spend:

- Connection capacity: the API connections and requests the server handles at once. Lookup and reading use it and never wait for research.
- Research capacity: the research runs active at once in one process, with bounded concurrent fetches and model calls across them. Work beyond it waits in a visible queue or is refused with a clear outcome.
- Research budget: the per-run limits on search results, page attempts, model calls and output tokens. Configuration sets it; a caller may lower it but not exceed it, and queueing or other runs never change it.
- Provider limits: the provider's rate limits, quotas and account credit. Reaching one delays or fails a run visibly; it never silently lowers the budget, switches the model or skips verification.

Safe updates to the same character are a separate rule, not a capacity limit.
Today the Node bridge in `src/web/server.ts` allows one change at a time across all characters and refuses the others.
The Python CLI checks the revision only when one is supplied. Research compares revisions before replacing the folder and refuses a new folder that appeared meanwhile.
The [#113 review](https://github.com/mardwerk/tower-generator/issues/113) proposes this policy instead:

- `serve` is the only Tower Generator writer for its Wiki directory. It keeps an in-memory set of characters with a change in progress, with no lock files or database.
- Research marks a character once identity resolution, the first networked step, reveals its key. It also marks a matched older folder that it renames.
- Saves, reviews and manual collection mark the character while they run.
- A change to a marked character, including a second research request, is refused with a busy outcome and changes nothing. The caller can repeat it after the first change ends.
- Reading, lookup and other characters stay available. Read results show the mark, so the website can disable editing.
- Before committing, each change checks that the files still have the revision it read. Otherwise it changes nothing and reports a stale revision.
- A missing entry counts as a revision, so two fresh runs cannot both create the same character.
- Saves and reviews must send the revision they read; a request without one is refused.
- Validate and commit each character update as a complete operation. Readers always see the last committed revision, never a partial update.

Direct edits to the Markdown files, such as in a text editor or through Git, need no coordination with `serve`.
`serve` keeps no cached copy and reads the files again for every request and before every commit.
The revision is a hash of the entry and its source files, as today. An external edit therefore makes an older save, review or research commit fail as stale instead of overwriting it.
An edited reviewed entry reads as draft again, as today. An external edit that lands between the final check and the file replacement is not detected, so avoid editing a character's files while it is being researched.
Run one `serve` per Wiki directory; writes by any other process count as external edits.

A failed, refused or ambiguous research run writes no Wiki files, as today.
Its outcome goes only to the caller and the server terminal. A new character has no file to hold it, and an existing entry keeps the `research` metadata of its last successful run.

Long-running work reports queued, running, completed or failed. This is transient process state, not a session: it ends with the process, while outcomes of successful work are recorded in the character files.

Start with concurrency inside one local Go process and portable files; choose numerical limits from measured provider and extraction behavior.
Distributed workers and a database are not required by this direction.

## Proposed screen responsibilities

The table defines each screen's job, not its final layout or complete control list.
Research is the current workflow; the other destinations remain visible according to the established navigation rule.

| Screen | What the person does | Default presentation | Secondary controls or later scope |
| --- | --- | --- | --- |
| Create | Enter a character and start research | One character input, Research action, small controls beneath it | Optional identity/scope hints and supplied evidence; generation inputs later |
| Wiki | Find and open locally researched characters | Search and a concise list of characters | Opening a saved entry performs offline lookup |
| Character entry within Wiki | Understand the evidence and improve or review the entry | Readable summary, capabilities, limitations and review state | Source passages, research again, editing and explicit human review |
| Library | Intended home for saved Towers | Disabled during the research focus | Define its contents and distinction from Generations before enabling |
| Generations | Intended home for inspecting generation work | Disabled during the research focus | Define retained history and revision behavior before enabling |
| Settings | Configure the local tool | Disabled until its required settings are specified | Model/provider references, storage location and later Profile selection |

Keep paths, token limits, raw metadata and provider diagnostics behind details unless they help the person resolve a problem.
Show a concise busy state, completion summary or actionable error where the action occurred.
Changes to the layout must follow an agreed screen contract, not introduce unrelated controls while implementing a backend feature.

## Proposed operation contracts

Define operations independently of HTTP route names.
Use the same validation and outcomes for every client; the table describes the minimum information rather than a fixed JSON schema.

| Operation | Input | Output or effect |
| --- | --- | --- |
| Lookup | Optional character query and selected Wiki location | Saved identities, summaries and review states; no network research |
| Read character | Saved character key | Portable content, retained evidence and current revision |
| Research or improve | Character query, optional series/scope hints and an optional lower research budget | Resolved identity, saved key, revision, changes, unavailable or incomplete sources, effective budget and reported usage |
| Save edit | Character key, edited content and expected revision | Validated files and a new revision, or an error preserving the saved entry |
| Record review | Character key, expected revision and explicit review target | Human review recorded for that target; finding and classification review remain distinct |
| Read categories/configuration | Selected local configuration | Broad classification definitions and safe effective settings; no secret values |
| Generate and validate, later | Character entry at a recorded revision and a selected Profile | Tower content, provenance including that revision, and declared checks; contract remains to be specified |

API responses should distinguish success, identity ambiguity, invalid input, a busy character, stale revisions and provider failure.
Keep progress separate from the final result, so tools can parse results without interpreting terminal messages.
Long operations report the work states defined in the concurrency rules; polling, streaming and cancellation remain open.
Do not require a browser session for any core operation.

## Proposed storage and evidence rules

Kyle selected the local Markdown Wiki in [#102](https://github.com/mardwerk/tower-generator/issues/102); the existing [Wiki format](wiki-format.md) is the implementation baseline.
Its `work` field and first folder level hold the series named in owner requirement 8.
Retain source passages beside each character; avoid database or MDX requirements.
File layout and schema versions are separate from the choice of programming language.

Preserve manual notes and reviewed findings when research refreshes evidence.
Keep suggested replacements separate, retain cited passages and disclose blocked sources or incomplete excerpts.
Validate the complete update before replacing saved content, and reject a conflicting revision instead of overwriting a newer edit.
Research does not delete a finding only because current sources omit it.
Make inputs such as storage location, canon boundary and research budget explicit and reproducible.
Keep credentials in protected local configuration and out of saved research, browser responses and prompts.

Research requests, Wiki entries and classification categories never name a Profile.
Generation reads a character entry; Profile mappings, Tower choices and numbers stay with the generation, not in the entry.

## Choices to leave open

Distribution, executable name, startup options, HTTP routes and schemas remain undecided.
[#89](https://github.com/mardwerk/tower-generator/issues/89) records `mardwerk-tower` as the preferred executable name.
Go is selected; extraction dependencies and numerical limits still need specification.
Admission when research capacity is full and retries after provider throttling remain open.
Whether the tool needs a spending limit across runs, beyond the provider account's own limits, remains open.
The [#113 review](https://github.com/mardwerk/tower-generator/issues/113) defers these until a demonstrated need:

- Cooperative research, where parallel requests for one character split the evidence and return one combined result.
- Waiting or queueing for a busy character instead of refusing the request.
- Reapplying stale research output to a newer revision; stale output is discarded, as today.
- Enforcing one `serve` per Wiki directory, for example with a lock file at startup.
- Watching the files for external edits; the commit-time revision check handles them.
- Authentication, user accounts and per-user spending limits for a `0.0.0.0` binding.
- A stored record of failed research.

Detailed screen controls, progress transport, cancellation, generation storage and Settings contents also need workflow-specific review.
The current search provider, model-call count and numerical limits are implementation choices, not permanent product rules.
Use [research behavior](RESEARCH.md) and the [Wiki format](wiki-format.md) for current commands and file contracts.

## How to expand this into the detailed specification

First review these rough rules and responsibility boundaries.
Then specify the name-to-character-entry workflow end to end before revising its implementation.
For each workflow, record the screen's visible controls, user action, core operation, inputs, outputs, progress, failures and saved files together.
Use a fresh character, an existing alias, an ambiguous name and a failed source as concrete examples.
Mark each choice as an established requirement, an agreed design or an open proposal, and link implementation issues to that contract.
