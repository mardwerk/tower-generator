# Architecture

How Unit Generator is structured. Where this document and the code disagree, the code is right; fix the document.

## What the Tool does

Unit Generator adapts an existing character into a Tower Defense unit under a selected Profile. The normal input is a character name (or explicit source documents) plus a Profile. The output is a checked unit artifact with its evidence and findings.

It researches sources, prepares one explicit request, drafts with a model, checks the draft deterministically, optionally reviews it with a second model call, and renders it.

It does not:

- simulate combat, waves or balance, or run units at runtime (a Consumer game does that);
- keep project history or orchestrate several tools (Towerright does that);
- author original characters (deferred);
- treat passing checks as balance, quality or acceptance.

The numerical Engine supports exactly 3 paths × 5 tiers. Other shapes need an explicit Engine change; a Profile or UI setting cannot enable them.

## Components

```text
                    mardwerk-unit (one Go binary, src/cli)
  CLI commands ──┐
                 ├── Engine: unit + mechanics ── unit.Model ── provider (OpenRouter, Codex)
  serve (HTTP) ──┘        │
        │                 └── render (Markdown, view data, icon prompts)
        ├── research (Wikipedia, Wikidata, Fandom, URLs) → Sources → prepare
        └── library (Sources, artifacts, icons) and Profiles (data/profiles)

  web client (src/web, TypeScript) ── HTTP /api/v1 ──> serve
```

| Component | Package | Owns | Never does |
| --- | --- | --- | --- |
| Engine | `unit`, `mechanics` | Contract types and schemas, preparation and hashing, drafting (plan, mechanics, repair), checks, review, mechanics resolution | File or network access, environment reads, history |
| Schema layer | `schema` | The contract DSL: strict parsing with Zod-compatible issues, JSON Schema for providers, JavaScript-compatible JSON | Business rules |
| `render` | `render` | Markdown, the web view (usage, per-tier stats, purchase sentences, crosspath builds, revision notes), icon subjects and prompts | Model calls, validation decisions |
| `provider` | `provider` | Model and image calls behind `unit.Model`, `.env` reading, key hints | Deciding what to generate |
| `research` | `research` | Character lookup, source text and images as Sources; explicit document inputs of request files | Applying a Profile |
| `library` | `library` | Managed files in the library folder, arranged by work and character; saved Profiles in the Profiles folder | Reading anything else, saving implicitly |
| `evidence` | `evidence` | Exact model inputs and raw outputs, with `--evidence-dir` only | Anything without that flag |
| CLI | `src/cli` (main) | Arguments, explicit input and output files, exit codes | Business rules |
| `serve` | `server` | Local HTTP routes, session token, host and origin checks, embedded web assets | Jobs, runs or resumable state |
| Web client | `src/web` | Screens, stage orchestration for the user, unsaved session state | Legality, build resolution, prompts, provider calls, validation |

Dependencies point inward. `unit` imports `mechanics` and `schema`; `render`, `research`, `library` and `provider` import `unit` for types; the CLI and `server` wire everything together. The Engine's only interface is `unit.Model`, because model execution is its only external boundary.

## Generation route

There is one route; the Profile supplies the rules it follows.

1. `research` finds the character and returns Sources, or a list of choices when the name is ambiguous. No model call. [Sources and passages](#sources-and-passages) describes what it reads.
2. `prepare` combines Sources (or a request file) with a Profile into one explicit request and records its hash. It also derives the request's `sourceTechniques`, which marks its passages as sectioned ([Sources and passages](#sources-and-passages)).
3. `draft` makes a planning call and a mechanics call. Each has a bounded repair budget (0 to 2, default 1). The planning call reads the request's `sourceTechniques`; under `requireCoreConcepts` it ranks its repertoire, and code checks the ranking, the major floor of each strong source technique, the listing of every source technique and a purchase that adapts each core entry. A plan attempt that fails only on items with a known local fix, a source technique listed nowhere or an effect adapted as a promise no purchase of its technique makes, is followed by one targeted correction ([plan_correction.go](../src/cli/internal/unit/plan_correction.go)): it names each failed item and the corrections it allows (list the technique with a rank; promise it on a purchase of that technique or remove it from that `adaptedAs`), it does not use the retry budget and it is not sent when the budget is 0. Its output passes the same plan validation, and the draft records it as an attempt with purpose `plan-correction` and its issues. Code binds the plan, resolves the mechanics and validates every legal build. Invalid output is never published.
4. `check` re-runs the deterministic checks on a draft. Each proposed mechanic of a purchase becomes an unresolved Finding (rule `proposed-mechanic`) that names the purchase and says the Definition must be expanded before any build can grant it. A third or fifth purchase that a behavior rule covers and whose only new capability is proposed adds a [Design gap](../CONTEXT.md), an unresolved Finding (rule `proposed-capability`), instead of failing ([mechanics](MECHANICS.md#default-authoring-policy)). Under `requireCoreConcepts`, a core entry that only proposed mechanics adapt is a Design gap too (rule `core-concept`): the Unit does not embody it until the Definition supports the mechanic.
5. `review` asks a separate model call for a semantic review. The review reads code-resolved facts: `legalBuilds` lists every build the Definition allows with its price, resolved attack and Active Ability, and `purchaseEvidence` the resolved purchase comparisons, with time-averaged Active rates and each side purchase's absolute gain and price beside those of its path's fifth purchase ([mechanics](MECHANICS.md#default-authoring-policy)). `referencePrices` states each path whose five prices equal the Definition's reference sequence (`profile.referenceScale.incrementalUpgradeCosts`) exactly. The review judges from this evidence whether each fourth and fifth purchase's price fits its gain and whether the capstone keeps a reason to buy, counting a side purchase against the capstone only when it adds as much absolute gain or more; code sets no threshold and reports no price finding. A finding about a mechanical mismatch cites the resolved values it relies on (`facts`: build code, field, value), and code checks each one against `legalBuilds`. A review that cites a build the Definition does not allow, such as `3-3-0`, a code that is neither a legal build nor one path's purchase (`1-x-x` to `x-x-5`), such as `1-x-5`, or a wrong value is corrected once in a second call and then rejected; the checked draft is kept. The correction sees the previous review and returns the whole review again. What it may do with a flagged finding depends on the fault: a finding whose only fault is notation, a malformed code such as `5-x-2` or another path's purchase, with no fact disproven, must come back under its ID with only the codes code flagged replaced, each by one code, and every other code, word and field exact; a finding that names an impossible build such as `3-3-0` may be fixed to a legal build, under any ID, or withdrawn; a finding with a wrong or unverifiable fact may be corrected under its ID or withdrawn. When only notation was wrong, the correction adds no finding and its summary may change only at the codes code flagged in it. Every finding whose citations held must come back under its ID unchanged in every field, and code compares each one in full with the previous review; a correction that drops one or changes any field, even only its message, is rejected. Before a Result is published, code rejects any numeric build the Definition does not allow in its summary and findings, including finding IDs, rules and facts. A rejected review keeps the checked draft: `review` reads it from its input, and `author` and `edit` write it beside `-o` as `<output>.checked.json` ([CLI](CLI.md#output-and-exit-codes)). The published findings are then exactly the ones the correction returned, so a Result's summary always describes its findings. Each purchased tier in the review's `unit.paths` carries what the plan says about that purchase: its technique, the sources that technique cites, its typed promises, the promises no effect of that technique adapts (`promisesWithoutEffect`, for a purchase whose technique is a repertoire entry other than the base attack), and its planned text, shown (`adaptation`) or private (`plannedChange`), and its purchase's proposed mechanics (`proposedMechanics`), which the review fails when one does not fit the purchase's technique and the passages it cites, or is lore rather than a playable mechanic. The review also applies the naming rule ([mechanics](MECHANICS.md#default-authoring-policy)): it fails a purchase name that borrows a source concept for a different effect, with the replacement name in the finding's action, and a revision (`edit`) renames the purchase from that finding; code judges no name. It judges each path's progression the same way: a third purchase that defines or distinguishes the path, a fourth that develops it and a fifth that is its pinnacle. It compares each third purchase with its own path's first and second purchases and with the other paths' purchases, and fails one that only adds numbers the path already had, or whose only distinction is a proposed mechanic that is not playable or conflicts with the Definition's rules. `promisesWithoutEffect` is a review cue, not a check: the review fails an adaptation or name that credits the technique with such a promise unless a passage the technique cites describes it. The review's `designPlan` keeps only the plan's decisions across purchases, without the milestones or a compact plan's code-written crosspath contributions, so each planned text appears once, on its tier, and the review judges every tier on its own build code. It also judges each repertoire entry's `effects` against the cited passages; each entry carries `adaptedBy`, the purchases whose technique it is, so the review knows which adaptations code checked and fails one claimed by an entry no purchase names. It judges every whole-technique omission in `omittedTechniques`, major and minor alike, against the Definition's vocabulary and the proposed mechanics a purchase could carry, and fails one whose reason does not hold, such as a speed increase called inexpressible when attack-rate expresses it, or a circular reason such as "this plan does not implement it" (`OmissionRule`); an omission named for a source technique carries `sourceTechnique`, its name, salience and passage IDs. Each purchased tier also lists its typed changes by scope (`changeScope`): `base`, what applies to the automatic attack in every build that owns it, and `boost`, what applies only to the Active Ability, so the review fails a text that limits a base change to the Active (`ActiveScopeRule`, OPUS-NET-61-7). When the plan ranks its repertoire, the review also reads each entry's `importance` and the request's `sourceTechniques`, and fails a core entry whose purchases carry its name but not what it does in play, never judges one embodied by a proposal alone, and challenges every ranking against the passages its source technique cites, failing one they contradict, an implausible downgrade included ([mechanics](MECHANICS.md#default-authoring-policy)). A finding whose text names the same tier on another path instead of its subject's purchase, such as `5-x-x` for `x-5-x`, is corrected like a wrong citation. Prose cues about a later technique or a side purchase's gain are advisory only. They cannot prove what a finding asserts, including negated claims or comparisons per unit of currency. The Result preserves these findings and adds an `unresolved` Finding (rule `review-claim-unread`) with retained timing or gain evidence for human review. These cues cannot trigger correction, withdrawal or rejection; wrong factual values must be established through structured citations ([review_claims.go](../src/cli/internal/unit/review_claims.go)). The review of a revision gets the earlier unit and the requested change, to check that the change was made, but not the earlier review's findings; it judges every finding on the current plan and unit.
6. `render` produces Markdown or view data. Every change a purchase or crosspath makes reads as both values and its delta, such as "Raises damage from 1 to 2 (+1)" or "damage 1 → 2 (+1)": the delta is what the change does, and the values the result when only that change is bought. Multipliers read as percentages: an interval multiplied by 0.85 "attacks 18% faster" (the attack rate changes by 1/0.85 - 1), and an Active Ability's damage multiplier of 2 gives "+100% damage", so a later purchase that changes it reads in percentage points. The unit sheet, the details report and UnitLab's kit stats share these helpers ([change.go](../src/cli/internal/render/change.go)); for a typed unit the details report words each purchase and Active Ability as the sheet does, while the compiled tier benefits and ability descriptions the review reads stay unchanged in the artifact. The unit sheet starts with the character name and `0-0-0`, names each purchase by build code (`3-x-x`, `x-4-x`, `x-x-5`) with the exact numbers of its resolved mechanics, labeled apart from the plan's intent: "Plan: adapts {technique}" ("(base attack)" for the base attack), the plan's source-backed description of the adaptation for a named technique, and an advisory "Review:" sentence when the purchase's name points to another repertoire entry through a word only that entry's name has, then "Resolved:" and the numbers (a Result whose plan predates typed techniques shows the numbers alone), then each of its [Proposed mechanics](../CONTEXT.md) as "Proposed (not yet supported): {name}: {effect}", lists every legal two-path build (under the default, 12 early and 36 advanced) with what each path adds to the other, and holds nothing else: no checks, costs, unsupported or reserved lists, default targeting or absent detection. A revision adds patch notes that keep changed mechanics apart from renamed purchases. Every other sentence and number comes from resolving the blueprint. The planned description is a claim written before the numbers: the review reads it in `unit.paths` and fails one the resolved mechanics or the cited evidence do not support; code drops a build code the text opens with or writes as "In build CODE,", which the sheet already shows; the plan's purchase reasons, weaknesses and capstone notes are private design checks and are not printed. `--details` (and the web app's panels) carry the diagnostics: provenance, review status, findings including unsupported mechanics, usage and evidence.

Every stage is a separate operation that receives the previous artifact explicitly. Nothing depends on an earlier call's presence in memory or on disk. A revision (`edit`, or Edit in the web app) prepares the previous unit, its findings and the feedback into a new request; there is no automatic review-and-redraft loop.

Legacy fields: `authoringMode`, `deliverable` and `operation` in older artifacts are read and kept but no longer select a route. Legacy prose drafts (without a blueprint) can be checked and rendered, but not drafted or reviewed again.

## Sources and passages

`research` makes no model call. It reads, within fixed bounds:

- **Identity.** A Wikipedia article about the character, or the character's entry in a list article. In a list, a heading that equals the name wins over a close heading, and a heading over a prose entry; a parenthetical that lists examples, such as "(e.g., Meliodas, Merlin and Escanor)", names no aliases ([character.go](../src/cli/internal/research/character.go)).
- **Character wiki page.** The Fandom page that the character's Wikidata item links (P6262). Without that link, the page with the character's exact name on the wiki that most of the character's works (P1441) link to; a tie between wikis names none. A page the user supplies (`--fandom URL`, or `fandom` in `/research` and `/character`) is read instead, without the identity lookup. The page and its ability subpage keep the introduction and the ability sections, capped at 24,000 characters: power, form and technique sections first (Devil Fruit, Haki, Gears, forms, techniques, weapons, magic), then other ability sections, then the introduction, overview and miscellaneous sections (such as Luck and Artistic Skill); within a rank, each passage goes to the section with the least text so far. A list entry keeps its own text, and its nested entries are read on their own ([visuals.go](../src/cli/internal/research/visuals.go)).
- **Technique pages.** At most four pages linked from the character's power and ability sections: first the pages those sections name as their own, a bold term an entry starts with or a "Further information" or "Main article" link, the most specific sections first; then links whose titles name an attack family (blade, fire, punch, beam, ice). Each keeps its introduction, description and technique entries up to 4,000 characters, and, where the page splits a section by user, only the character's subsection ([techniques.go](../src/cli/internal/research/techniques.go)).

The Engine splits each source document into passages, the quotable spans the model cites ([evidence.go](../src/cli/internal/unit/evidence.go), [source_sections.go](../src/cli/internal/unit/source_sections.go)). A request that `prepare` made now carries `sourceTechniques`, and its passages are sectioned: each carries its heading path as `section` ("Devil Fruit > Gear 4"; none for an introduction), and headings are never passages. A saved request without `sourceTechniques` keeps the passages it was prepared with, so its hash and its draft's citations stay valid.

The model receives at most 200 passages and 32,000 characters. When a sectioned request has more, the selection keeps, in order: the first passage of each document; for each technique page, its header, first passage and best passages by score, the pages sharing 8,000 characters; each document's introduction in page order; the best passages of each power section and then each ability section, up to 4 passages and 1,500 characters each and 16,000 in all; and then the remaining passages by score, sections taking turns among passages of equal score, power sections first. A section's rank comes from a short keyword list on its headings (`SectionRank`).

`sourceTechniques` lists, in passage order, the techniques and forms the selected passages name: a power section heading that is not only a category, the leading term of an entry in a power or ability section ("Cruel Sun「…」: …"), a technique page's title, or the capitalized name after a signature cue ("his signature attack, the Gum-Gum Pistol"). Each entry has its `name`, up to 12 `passageIds` and `signature`, true when a sentence with a cue such as "signature", "most iconic" or "trademark" names it. Two names are one entry only when one passage proves it ([source_aliases.go](../src/cli/internal/unit/source_aliases.go)): the first passage of a technique page or of a section defines its subject by the other name, as the Rhitta page opens "The Divine Axe Rhitta「…」 is a Sacred Treasure", or one sentence says "X, also called Y", "X (also known as Y)" or "X, otherwise known as Y". Shared words prove nothing, so Gear 2 and Gear 2 Buso, and the kinds of Haki, stay apart. The merged entry keeps the name the passages name first, lists the others in `aliases` and keeps the passage IDs of each name. Each entry then carries its `salience` ([technique_salience.go](../src/cli/internal/unit/technique_salience.go)): `strong` when research followed a technique or form page titled for it, when it is a signature technique, or when its own sections and entries hold at least 10 percent (`StrongSectionShare`) of the text of the power and ability sections of the character articles, the documents that are not technique pages, read in full; otherwise `normal`. A technique's own sections are the text under a heading that names it, subsections included, and the list entries that start with its name. In the v30 Sources of #61, Luffy's Armament and Supreme King Haki hold 14 and 11 percent, his Observation Haki 7 and each named attack under 2; Escanor's The One and Daytime hold 14 and 12 percent. Strong salience is evidence that a technique is at least major, never that it is core ([Salience](../CONTEXT.md)). The plan and the review read it: under `requireCoreConcepts` the plan ranks its techniques as [core concepts](../CONTEXT.md) and lists each source technique ([mechanics](MECHANICS.md#default-authoring-policy)). `prepare` derives it again from the documents, and verifying a prepared request rejects an edited one. A request prepared under Default v30 or earlier has entries without `salience` and without merged aliases; verifying it derives that earlier form, so its hash and its draft's citations stay valid.

## Profiles

- **Profile file.** One JSON document holding the Definition, the rules text and the task. `prepare` copies the Profile's content into the request, so the hash covers it and reloading an artifact never looks a Profile up by ID.
- **Bundled default.** `default`, BTD6-inspired: three paths of five tiers, BTD6 crosspath rules, the character design rules and scale references for several roles, pinned to btd6-atlas capture 56.3 ([BTD6 reference](BTD6-REFERENCE.md)). It is built into the binary and read-only; editing starts from a copy.
- **Saved Profiles.** `<id>.json` files in the Profiles folder (`data/profiles`, `--profiles DIR`), separate from generated runs. Saving runs the same validation as `prepare`, and a saved Profile's rules document must have the ID `profile:<id>`.
- **Web app.** The Profiles tab lists the default and saved Profiles, shows paths × tiers, prices and limits, and edits copies. The Generate form has a compact Profile dropdown with the default preselected.
- **CLI.** `--profile ID`; without it, the bundled default.

## Library

Saved work lives only in a library folder, `data/runs/library` by default. It is set with `--library DIR`, or switched from the web app's Settings; the web app records that choice in `data/runs/lab-settings.json`. Switching is refused while a model stage runs. Records are arranged by source and character: `<work>/<character>/<character>.<stage>.<id>.json`, where `<work>` and `<character>` are readable slugs, `<stage>` is `sources`, `prepared`, `draft`, `checked` or `result` and `<id>` the first 12 hex digits of the artifact's SHA-256 (all 64 if two collide). Each unit also gets a `.md` render beside it, and the character's icons, image receipts and portrait choice live in its `assets/` folder. A `character.json` marker records the exact name and work of each character folder; a character whose slug another identity owns gets `<character>-<hash>`. Slugs keep only letters and digits, so no name can leave the library, and every folder below the library root must be a real directory. The API and listing still identify records by the full SHA-256, which also deduplicates saves. Records from earlier versions (`unitlab-<id>.json` at the root, assets under `assets/unit-<hash>`) stay readable; `library migrate` moves them on request. The server never saves on its own: the web client saves Sources after research and Results when a run completes, and the CLI saves only with `library save`.

## CLI

The binary is `mardwerk-unit` (`go build -o mardwerk-unit ./src/cli`). Commands read explicit inputs and write one artifact to stdout, or to a new `-o` file that is never overwritten (an exclusive temporary file is hard-linked into place). Diagnostics go to stderr. [CLI.md](CLI.md) lists every command and option.

Exit codes: `0` the operation completed (findings may still fail), `1` failure. Configuration comes from flags, then the environment, then `.env` in the working directory. The provider key is read once at startup and never written to artifacts.

## HTTP API

`serve` binds `127.0.0.1` and serves the embedded web client plus the API under `/api/v1`. It injects a fresh session token into the page. The page's Content-Security-Policy allows scripts only from the server, and styles from the server plus `<style>` elements carrying a nonce that is new on every page load; the component library needs those for scroll locking and select menus. Every call must present that token and come from the server's own host and origin; POST calls must also be JSON and stay under the 32 MB body limit. Requests are synchronous and are cancelled when the client disconnects. There are no job, run or history endpoints.

| Method and path | Body | Response |
| --- | --- | --- |
| `GET /health` | – | `{status, version, key: {configured, source, hint}, provider}`, where source is `env`, `env-file`, `settings` or `none` |
| `GET /provider`, `POST /provider` | –, `{provider, apiKey?, model?, reasoning?, imageModel?}` | provider state with selected reasoning |
| `POST /models` | `{provider}` | model catalog and .env defaults for the selected provider |
| `GET /example`, `GET /definition` | – | the example request, the bundled Definition |
| `POST /research` | `{name, choice?, fandom?, previous?}` | Sources or `{kind: "choices", choices}`. With `previous` Sources of the same character, the new lookup extends them: a document retrieved again replaces its copy with the same ID, others stay, and identical documents are kept once |
| `POST /character` | `{name, choice?, fandom?, profileId?\|profile?}` | research and prepare in one call: prepared request or choices |
| `POST /prepare` | `{request, profileId?\|profile?}` or `{sources, profileId?\|profile?}` | prepared request; documents are resolved, `text` or `url`, never `file` |
| `POST /draft` | `{prepared, maxRepairAttempts?}` | draft |
| `POST /check` | `{draft}` | checked artifact |
| `POST /review` | `{checked}` | Result |
| `POST /inspect` | `{artifact, editable?}` | `{kind, artifact}` |
| `POST /render` | `{artifact, details?}` | `{markdown}` |
| `POST /view` | `{artifact}` | `{view, base?, stats?, purchases?, crosspaths?, revision?, authoringIssues?}`: usage summary, design evaluation, the `0-0-0` description, per-tier stat changes with each change's `delta` worded as on the unit sheet, purchase sentences by build code, every legal two-path build, for a revision its mechanics and wording changes and, when the unit's structure is valid but today's [authoring checks](../CONTEXT.md) fail, those issues as `[{path, message}]`. `authoringIssues` is computed when the artifact is read; it is not a stored Finding and is omitted when empty |
| `GET /profiles`, `POST /profiles/save`, `POST /profiles/delete` | –, `{profile}`, `{id}` | `{directory, profiles: [{profile, builtIn, progression}]}` |
| `POST /profiles/apply` | `{request, profileId?\|profile?}` | the edited request under that Profile |
| `GET /library`, `POST /library/configure` | –, `{directory}` | `{directory, entries}`; each entry has its record `path` relative to the folder, and saved Sources their researched `query` |
| `POST /library/migrate` | `{}` | `{records, assets, kept, state}`: records and asset files moved into work and character folders, and files left in place |
| `POST /library/save`, `/load`, `/delete` | `{artifact}`, `{id}`, `{ids}` | entry, `{artifact}`, listing |
| `POST /library/icons` | `{artifact}` | `{directory, icons, portrait?}`; each icon carries its image and Codex prompts |
| `POST /library/portrait/get`, `/library/portrait` | `{artifact}`, `{artifact, referenceId}` | `{portrait?}` |
| `POST /library/icon/generate` | `{artifact, iconKey, model, confirmed: true, destination}` | `{icons, model, usage?}` |

Errors are `{"error": {"code", "message", "details"?, "usage"?}}`. The status is 400 for invalid input, 401 and 403 for the session, host or origin, 404 for an unknown operation, 409 when busy or when the image model or destination changed, 413 for a body that is too large, 415 for a body that is not JSON, and 502 for a model failure. Raw provider payloads, credentials and source text never appear in errors.

## Contracts and versioning

- **Artifacts.** Sources carry `schemaVersion: "1"`. A request, prepared request, draft, checked artifact, Result and Profile carry `schemaVersion: "1"`, or `"2"` exactly when their mechanics Definition is a version 2 Definition with a [Profile-defined vocabulary](MECHANICS.md#profile-defined-vocabulary-version-2); readers choose the schema from that field, so version 1 artifacts stay readable. The contracts are strict schemas in `src/cli/internal/unit` and `src/cli/internal/research`: readers reject unknown keys and unknown versions. A field an older reader would reject bumps the version. Saved artifacts are never rewritten in place.
- **Input hash.** New prepared requests use `jcs-sha256:` (RFC 8785 canonical JSON). The older `sha256:` form (JavaScript `JSON.stringify` of the contract-ordered request) still verifies. Optional request fields added later, such as `sourceTechniques`, are absent from older requests, so their hashes are unchanged.
- **Versions.** Three, no more: the URL prefix for endpoints, `schemaVersion` for artifacts, and the Definition's `revision` inside each prepared request for rules.
- **Findings.** A check that finds problems still completes; failures are findings, not errors.

## Code layout

```text
go.mod                         module github.com/mardwerk/unit-generator
src/cli/                       main: flags, commands, file output, wiring
src/cli/internal/schema/       contract DSL (strict parsing, JSON Schema, JavaScript-compatible JSON)
src/cli/internal/mechanics/    3×5 mechanics: resolve, legal builds, validate, design policy
src/cli/internal/unit/         Engine: contracts, hash, prepare, plan, draft, repair, check, review, prompts
src/cli/internal/render/       Markdown, view data, kit stats, icon prompts
src/cli/internal/provider/     OpenRouter (chat, images), Codex, .env
src/cli/internal/research/     character lookup, Sources, explicit documents and request files
src/cli/internal/library/      library folder and Profiles folder
src/cli/internal/evidence/     --evidence-dir records
src/cli/internal/server/       serve: routes and security checks
src/cli/internal/fixture/      test helper: a scripted reference unit run through every stage
src/web/                       web client: app/, features/, ui/, api/ (React), public/, dist/ (embedded), build.mjs
```

Go dependencies: `golang.org/x/text` (NFKC and NFKD), `github.com/clipperhouse/uax29` (sentence segmentation of evidence), `github.com/PuerkitoBio/goquery` (HTML), `golang.org/x/image` (WebP decoding, thumbnails) and `github.com/BurntSushi/toml` (Codex configuration).

## Web client

The TypeScript client keeps the React screens, the Profiles and Generate tabs, stage orchestration for the user, the current session's unsaved revisions, session import and export, and display-only helpers (portrait ordering, usage formatting, kit comparison). It contains no Engine code, model prompts, provider calls or validation: the server supplies stats, icon prompts, Profile progressions and Profile application.

| Folder | Holds |
| --- | --- |
| `src/web/api/` | The HTTP client, the contract types (`contract.ts`) and read-only helpers over artifacts and usage |
| `src/web/app/` | The entry point, the shell (top bar, activity) and the Tailwind theme (`styles.css`) |
| `src/web/features/` | One folder per area: `authoring` (session state and stage runs), `generate`, `unit`, `library`, `profiles`, `settings` |
| `src/web/ui/` | Generic controls in the [shadcn/ui](https://ui.shadcn.com) style: button, input, select, dropdown menu, dialog, tabs, disclosure, tooltip, badge, alert |

Controls are built on [Radix](https://www.radix-ui.com) primitives and styled with [Tailwind CSS](https://tailwindcss.com) utilities; the theme maps the dark palette onto shadcn/ui's color tokens. `ui/` imports nothing unit-specific, so it can move into a shared package once another generator needs it. `pnpm build` compiles the Tailwind CSS and bundles the client into `src/web/dist`, which is committed so the Go build needs no Node.js.
