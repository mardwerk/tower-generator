# Tower Generator product outline

Working outline for review, recorded on 2026-10-02 in [#106](https://github.com/mardwerk/tower-generator/issues/106).
Kyle selected defining rough rules before further implementation.
The owner requirements below come from his requests; the proposed sections remain open for review.
Kyle selected Go for the backend, independent CLI and `serve` API on 2026-10-02.
Implementation remains pending the workflow specification.
Proposals from the [#107 review](https://github.com/mardwerk/tower-generator/issues/107) leave the owner requirements and the agreed backend direction unchanged.

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
10. Use Go for the backend and independent CLI. A `serve` command exposes one API capable of handling concurrent clients and requests.

## Proposed responsibilities

Use one shared core for character identity, research, lookup, file validation, saving and review.
The CLI and web API call that core so they have the same behavior and storage rules.
The website handles forms, navigation, readable evidence and user actions; the backend owns durable files, provider access and validation.
Browser memory may hold an unsaved edit, but it must not be the only copy of saved research.
An independent CLI command runs the core in its own process and does not need a running `serve`.
`serve` uses one Wiki location chosen when it starts, so API requests cannot write to other paths.
It accepts only local requests and rejects other browser origins, as the current local API does. Remote or multi-user access needs its own decision, including authentication.

## Agreed backend direction

Use one Go executable with independent CLI commands and `serve` for the web API.
Both interfaces call the same Go core directly; the website uses that API.
Go supports concurrent network work and a compiled executable; the current Python CLI and Node API bridge remain the working experiment until migration.
The target CLI and API must run without Node or Python. Frontend build tooling can remain separate.
Serving the built website from the executable is a packaging proposal, not yet a specified requirement.

## Proposed concurrency rules

Handle multiple API connections while allowing independent characters and source fetches to progress concurrently.
Keep four limits separate, so more clients never create unlimited provider calls and load never changes what a run may spend:

- Connection capacity: the API connections and requests the server handles at once. Lookup and reading use it and never wait for research.
- Research capacity: the research runs active at once in one process, with bounded concurrent fetches and model calls across them. Work beyond it waits in a visible queue or is refused with a clear outcome.
- Research budget: the per-run limits on search results, page attempts, model calls and output tokens. Configuration sets it; a caller may lower it but not exceed it, and queueing or other runs never change it.
- Provider limits: the provider's rate limits, quotas and account credit. Reaching one delays or fails a run visibly; it never silently lowers the budget, switches the model or skips verification.

Safe updates to the same character are a separate rule, not a capacity limit:

- Each research run, save and review checks before committing that the saved revision is still the one it read. Otherwise it changes nothing and reports a stale revision.
- A missing entry counts as a revision, so two fresh runs cannot both create the same character.
- The check must also hold between separate CLI and `serve` processes. Different characters do not share one global write lock.
- The character key is known only after identity resolution, the first networked step. Within one process, coordinate same-character research from that point; the commit check remains the final safeguard.
- Validate and commit each character update as a complete operation. Readers always see the last committed revision, never a partial update.

Long-running work reports queued, running, completed or failed. This is transient process state, not a session: it ends with the process, while outcomes are recorded in the character files.
The API reports work started through that server; a CLI command reports its own progress and leaves its outcome in the files.

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

Define operations independently of HTTP route names and CLI spelling.
Use the same validation and outcomes for both interfaces; the table describes the minimum information rather than a fixed JSON schema.

| Operation | Input | Output or effect |
| --- | --- | --- |
| Lookup | Optional character query and selected Wiki location | Saved identities, summaries and review states; no network research |
| Read character | Saved character key | Portable content, retained evidence and current revision |
| Research or improve | Character query, optional series/scope hints and an optional lower research budget | Resolved identity, saved key, revision, changes, unavailable or incomplete sources, effective budget and reported usage |
| Save edit | Character key, edited content and expected revision | Validated files and a new revision, or an error preserving the saved entry |
| Record review | Character key, expected revision and explicit review target | Human review recorded for that target; finding and classification review remain distinct |
| Read categories/configuration | Selected local configuration | Broad classification definitions and safe effective settings; no secret values |
| Generate and validate, later | Character entry at a recorded revision and a selected Profile | Tower content, provenance including that revision, and declared checks; contract remains to be specified |

Machine-readable CLI output and API responses should distinguish success, identity ambiguity, invalid input, stale revisions and provider failure.
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

Distribution, executable name, exact command names beyond `serve`, HTTP routes and schemas remain undecided.
[#89](https://github.com/mardwerk/tower-generator/issues/89) records `mardwerk-tower` as the preferred executable name.
Go is selected; extraction dependencies, numerical limits and the mechanism that makes the revision check hold across processes still need specification.
Admission when research capacity is full, a second request for a character already being researched and retries after provider throttling remain open.
Whether stale research output is discarded or reapplied to the newer revision is also open.
Whether the tool needs a spending limit across runs, beyond the provider account's own limits, remains open.
Detailed screen controls, progress transport, cancellation, generation storage and Settings contents also need workflow-specific review.
The current search provider, model-call count and numerical limits are implementation choices, not permanent product rules.
Use [research behavior](RESEARCH.md) and the [Wiki format](wiki-format.md) for current commands and file contracts.

## How to expand this into the detailed specification

First review these rough rules and responsibility boundaries.
Then specify the name-to-character-entry workflow end to end before revising its implementation.
For each workflow, record the screen's visible controls, user action, core operation, inputs, outputs, progress, failures and saved files together.
Use a fresh character, an existing alias, an ambiguous name and a failed source as concrete examples.
Mark each choice as an established requirement, an agreed design or an open proposal, and link implementation issues to that contract.
