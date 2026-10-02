# Tower Generator product outline

Working outline for review, recorded on 2026-10-02 in [#106](https://github.com/mardwerk/tower-generator/issues/106).
Kyle selected defining rough rules before further implementation.
The owner requirements below come from his requests; the proposed sections remain open for review.
This document does not select a backend language or authorize a rewrite.

## Purpose and current focus

Tower Generator should turn reusable character evidence into Towers that a person can understand, inspect and revise for a selected Profile.
The immediate focus is reliable research and local storage; generation remains later work.
Specify the product's behavior before choosing implementation details or changing screens.
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
10. The CLI must work independently of the website. Its implementation language is open.

## Proposed responsibilities

Use one shared core for character identity, research, lookup, file validation, saving and review.
The CLI and web API call that core so they have the same behavior and storage rules.
The website handles forms, navigation, readable evidence and user actions; the backend owns durable files, provider access and validation.
Browser memory may hold an unsaved edit, but it must not be the only copy of saved research.

A CLI with a `serve` command could expose the local API and serve the built website.
Go is a candidate because a compiled binary can simplify installation and include the web assets.
Python can reuse the existing extraction code; Node can share TypeScript code and packages with the web application.
Choose the language after defining packaging and behavior, rather than treating the current bridge as a permanent architecture.

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
| Research or improve | Character query, optional series/scope hints and a bounded budget | Resolved identity, saved key, revision, changes, source limits and available usage information |
| Save edit | Character key, edited content and expected revision | Validated files and a new revision, or an error preserving the saved entry |
| Record review | Character key, expected revision and explicit review target | Human review recorded for that target; finding and classification review remain distinct |
| Read categories/configuration | Selected local configuration | Broad classification definitions and safe effective settings; no secret values |
| Generate and validate, later | Character evidence and a selected Profile | Tower content, provenance and declared checks; contract remains to be specified |

Machine-readable CLI output and API responses should distinguish success, identity ambiguity, invalid input, stale edits and provider failure.
Keep progress separate from the final result, so tools can parse results without interpreting terminal messages.
Long operations need an observable running, completed or failed state; polling, streaming and cancellation details remain open.
Do not require a browser session for any core operation.

## Proposed storage and evidence rules

Keep Markdown with simple metadata as the initial format, using the existing [Wiki format](wiki-format.md) as the implementation baseline.
Retain source passages beside each character; avoid database or MDX requirements.
File layout and schema versions are separate from the choice of programming language.

Preserve manual notes and reviewed findings when research refreshes evidence.
Keep suggested replacements separate, retain cited passages and disclose blocked sources or incomplete excerpts.
Validate the complete update before replacing saved content, and reject a conflicting revision instead of overwriting a newer edit.
Make inputs such as storage location, canon boundary and research limits explicit and reproducible.
Keep credentials in protected local configuration and out of saved research, browser responses and prompts.

## Choices to leave open

Backend language and distribution, executable name, exact command names, HTTP routes and schemas remain undecided.
Detailed screen controls, progress transport, cancellation, generation storage and Settings contents also need workflow-specific review.
The current search provider, model-call count and numerical limits are implementation choices, not permanent product rules.
Use [research behavior](RESEARCH.md) and the [Wiki format](wiki-format.md) for current commands and file contracts.

## How to expand this into the detailed specification

First review these rough rules and responsibility boundaries.
Then specify the name-to-character-entry workflow end to end before revising its implementation.
For each workflow, record the screen's visible controls, user action, core operation, inputs, outputs, progress, failures and saved files together.
Use a fresh character, an existing alias, an ambiguous name and a failed source as concrete examples.
Mark each choice as an established requirement, an agreed design or an open proposal, and link implementation issues to that contract.
