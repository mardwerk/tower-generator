# Anime TD systems verification

Reviewed summary for [issue #1](https://github.com/mardwerk/project/issues/1).
Reviewed on 2026-09-26 against the raw import at commit [`d72f64f`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge), verified byte for byte against the [import manifest](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/MANIFEST.json).
Use this summary as a map of the UG-anime-td-systems handover and its source register; do not use it as proof that the reference unit, schema or tests work, because the demo files are not in this import and nothing could be rerun.

## Question and scope

The group records a 2026-09-24 handover of one earlier assistant conversation.
The question was whether a system could turn a character name ("Monkey D. Luffy") into a tower defense unit detailed enough to implement.
The conversation produced a research HTML page (`anime-unit-forge.html`), a nine-file reference package (`unit-forge-reference-package.zip`: `luffy.unit.json`, `game.contract.json`, `unit.schema.json`, `sources.json`, `reference-combat.js`, `reference-tests.js`, `README.md`, `VALIDATION.md` and the HTML), and a 34-entry source list.
The handover adds a decision record, open questions, a re-annotated source register, and a verification supplement.

In scope: the five files of [`anime-td-systems-verification/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification), all read in full and listed under [Raw evidence](#raw-evidence).
The matching cold-store ZIP was checked as historical provenance only.

Out of scope: the other demo revision in [`anime-td-handover/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover), its ZIPs (`UG-anime-td-handover.zip`, `unit_forge_research.zip`) and the private cold store.
That revision is reviewed separately in [anime-td-handover.md](anime-td-handover.md); no evidence was pooled, and no file from it was used to fill gaps here.

## Summary for current use

### Verified or reproduced here

- **Confirmed**: the five raw files match the import manifest, and are byte-identical to the five files in the cold-store `UG-anime-td-systems-verification.zip`, whose SHA-256 (`1ee982c5ebd0...`) matches `drafts/99-cold-store/zips/MANIFEST.sha256`.
  The ZIP contains only these five files.
- **Unverified**: the recorded rerun (37/37 Node checks under v22.16.0, `schema_valid` and `unit_valid` true, seven embedded exports matching, nine file hashes).
  None of the nine package files is in this import, in the matching ZIP, or anywhere in the `unit-generator` product repository history, so no test, schema validation or hash could be recomputed.
- **Confirmed** (internal consistency only): `node-verification.txt` has 37 `PASS` lines and no failures, matching its `37/37` line.
  `source-verification.json` holds 34 historical entries plus S35 and N01 to N03, as the README and handover state.
  The handover's access summary for S10, S11, S16, S18, S19, S26 to S28, S32 and S33 matches the JSON `access` fields.
  The recorded numbers are arithmetically consistent: 165 x 0.45 = 74.25 rounds to 74; 50 hits x 40 = 2,000; a 4,830 refund on 6,900 invested is 70%; 432 frames at about 60.8 MiB fits 192 x 192 RGBA frames (frame size inferred, not stated in this folder).
- **Confirmed**: the Anthropic articles [S04](https://www.anthropic.com/engineering/building-effective-agents) (workflow versus agent distinction; banner says the tooling landscape has changed), [S05](https://www.anthropic.com/engineering/demystifying-evals-for-ai-agents) (code, model and human graders; repeated trials) and [S33](https://www.anthropic.com/engineering/harness-design-long-running-apps) (planner, generator and evaluator roles; stripping components that are no longer load-bearing) match the register's bylines, dates and attributions.
- **Confirmed**: [AniList terms](https://docs.anilist.co/guide/terms-of-use) (S11, read at the docs source repository): free for non-commercial use; commercial use free below USD 150 revenue per month and licensed above; hoarding, mass collection, backup or storage use, and competing list or tracker services are prohibited.
  Both the original "non-commercial" note and the later "commercial licensing" note are partial.
- **Confirmed**: the [Hunyuan3D 2.1 licence](https://github.com/Tencent-Hunyuan/Hunyuan3D-2.1) (S34) does not apply in the European Union, United Kingdom or South Korea and forbids use of its works or output outside that Territory.
  The draft only said a licence audit was pending.
  Any EU-based use falls outside the licence, which matters because the drafts reason under German copyright law.
- **Confirmed**: Ollama Cloud does not support structured outputs (S07); Anthropic structured outputs cover JSON outputs and strict tool use under stated JSON Schema limitations (S02); the Draft 2020-12 meta-schema URI resolves today and its landing page lists the four authors, 16 June 2022, and the `draft-bhutton-*` documents (S35, N02).
- **Unverified**: all One Piece character claims.
  C01 ("unarmed rubber fighter", S27) and C02 ("Gum Gum Pistol" as the early attack, S28) rest on search-index evidence, then and now; one-piece.com and YouTube were not readable. "East Blue-inspired" is not a canon or era audit.

### Proposals (classified, not verified)

- **Proposal**: the six generation approaches and their numeric rankings (templates, single LLM with schema, grounded workflow, compiler with simulation, multi-agent, tuned model).
  The scores are the assistant's planning estimates.
- **Proposal**: decisions D02 to D05 (reference-game contract, grounded workflow evolving to a compiler, authoring-time generation with `NEEDS_NEW_PRIMITIVE` for unsupported mechanics, deferring agents and training) and D07, the smaller text/specification pilot.
  The handover records only D01, the research task itself, as owner-accepted.
  D06 (no rubber-defense passive) applies to the Luffy example only.
- **Proposal**: the whole E02 reference unit and its formulas: `td-reference-1.0.0`, six linear tiers, 20 ticks per second, 450 placement cost, 4-tick windup, 40 base damage, armor where 100 halves damage, every fourth basic becoming a capped 120-degree sweep from tier 2, and a tier-3 active with six pulses, a 60-degree cone and a 600-tick cooldown.
  These are invented game adaptations, not balance findings or canon.
- **Proposal**: the E06 evaluation plan (40 cases, three trials each, 95% factual precision and 4/5 recognizability targets), the eight-stage pipeline, the asset specification (432 frames), the engine mappings, and the cost-calculator rates.
  None was measured or built.

### Use and limits

A consumer may rely on: the list of open questions Q01 to Q08; the facts-versus-adaptation separation as a pattern; the AniList, Hunyuan3D and Ollama constraints above.
A consumer must not rely on: any claim that the schema, runtime or tests pass today; the reference unit as a model of the product (the product engine uses three upgrade paths of five tiers, not six linear tiers); any One Piece canon claim; any provider, cost or engine claim beyond page identity.

## Review record

### Method and access

On 2026-09-26 the reviewer read all five raw files in full, recomputed their SHA-256 values against `drafts/MANIFEST.json`, recomputed the cold-store ZIP hash, extracted it to scratch and compared bytes.
The reviewer searched this repository (excluding the sibling folder and the private cold store) and the public `unit-generator` clone (main at `01000bf`, all history) for the package file names, the reference key and the recorded hashes; none were found.
No rerun was possible.

Direct fetch worked for www.anthropic.com, platform.claude.com and json-schema.org.
GitHub-hosted documentation sources were read with `git clone`, `git ls-remote` and raw.githubusercontent.com: AniList/docs `03281c0`, ollama/ollama `7af3931`, vllm-project/vllm `bf93e1f`, langchain-ai/docs `264fdf8`, Comfy-Org/docs `b0605b5`, godotengine/godot-docs stable `6d86d7c`, optuna/optuna `0d05a96`, OWASP top-10-for-LLM `99f4395`, Roblox/creator-docs `0b817b5`, pgvector `7db2345`, Hunyuan3D-2.1 `82920d6`.
These are the documentation sources, not the rendered pages, which were blocked like the rest.
The proxy blocked (HTTP 403) developers.openai.com, ai.google.dev, docs.meshy.ai, elevenlabs.io, docs.unity3d.com, dev.epicgames.com, gesetze-im-internet.de, copyright.gov, one-piece.com, youtube.com, arxiv.org and github.com web pages.
For those, WebSearch gave titles, URLs and snippets only.

### Sources checked

Draft IDs are those in the handover [register](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/UG-anime-td-systems-handover.md?plain=1#L139-L305).
All 38 external entries were checked in some form.

| Source | Cited in | Check (2026-09-26) | Status | Note |
|---|---|---|---|---|
| Import manifest, 5 files | Import | SHA-256 recomputed | Confirmed | All match. |
| Cold-store ZIP | `99-cold-store/zips` | SHA-256 recomputed, unzipped, `cmp` | Confirmed | Same five files; no demo files. |
| `node-verification.txt` rerun | E03 | Files missing | Unverified | `reference-combat.js` and `reference-tests.js` absent. |
| `verification.json` hashes, schema result | E03, H07 | Files missing | Unverified | Nine files absent everywhere searched. |
| H07 path `ug-handover-inspection/` | Handover L21 | Path search | Stale link | Content now lives in this folder under plain names. |
| [S01](https://developers.openai.com/api/docs/guides/structured-outputs) OpenAI structured outputs | D03/D04 | Search | Metadata only | Title "Structured model outputs" matches. |
| [S02](https://platform.claude.com/docs/en/build-with-claude/structured-outputs) Anthropic structured outputs | D03/D04 | Fetched | Confirmed | JSON outputs, strict tool use, schema limitations. |
| [S03](https://ai.google.dev/gemini-api/docs/structured-output) Gemini structured outputs | D03/D04 | Search | Metadata only | Title matches; page date conflict unresolved. |
| S04, S05, S33 Anthropic Engineering | D05, E06, Q04 | Fetched | Confirmed | Bylines, dates, passages match. |
| [S06](https://docs.langchain.com/oss/python/langgraph/overview) LangGraph overview | D05 | Docs source | Confirmed | Mixes deterministic and LLM steps; human-in-the-loop, persistence. |
| [S07](https://docs.ollama.com/capabilities/structured-outputs) Ollama | Q06 | Docs source | Confirmed | Cloud lacks structured outputs. |
| [S08](https://docs.vllm.ai/en/latest/features/structured_outputs/) vLLM | Q06 | Docs source | Confirmed | Title matches; xgrammar/guidance backends. |
| [S09](https://github.com/pgvector/pgvector) pgvector | Architecture | Raw README | Confirmed | Vector similarity search for Postgres. |
| [S10](https://docs.anilist.co/guide/introduction) AniList introduction | Q05 | Docs source | Confirmed | Free public GraphQL API. |
| [S11](https://docs.anilist.co/guide/terms-of-use) AniList terms | Q05/Q08 | Docs source | Confirmed | See summary; resolves the note dispute. |
| [S12](https://genai.owasp.org/llmrisk/llm01-prompt-injection/) OWASP LLM01:2025 | D04 | Project repo source | Confirmed | Least privilege; segregate external content. |
| [S13](https://docs.comfy.org/) ComfyUI docs | Assets | Docs source | Confirmed | Title "ComfyUI Official Documentation". |
| [S14](https://developers.openai.com/api/docs/guides/tools-image-generation) OpenAI image generation | Assets | Search | Metadata only | Title matches at platform.openai.com path; exact developers path not shown. |
| [S15](https://ai.google.dev/gemini-api/docs/image-generation) Gemini image generation | Assets | Search | Metadata only | URL exists; indexed title is "Gemini API", not the register label. |
| S16, S17, S18, S32 Meshy text-to-3D, rigging, animation, pricing | Assets, E02 | Search | Metadata only | Titles match; rigging snippet says textured humanoid GLB. |
| [S19](https://elevenlabs.io/docs/overview/capabilities/sound-effects) ElevenLabs | Assets | Search | Metadata only | Title "Sound effects". |
| [S20](https://docs.unity3d.com/6000.0/Documentation/Manual/class-ScriptableObject.html) Unity ScriptableObject | Q07 | Search | Metadata only | Only the 6000.1 manual surfaced. |
| [S21](https://docs.godotengine.org/en/stable/tutorials/scripting/resources.html) Godot Resources | Q07 | Docs source (stable) | Confirmed | Resources are data containers; nodes act. |
| [S22](https://dev.epicgames.com/documentation/en-us/unreal-engine/data-assets-in-unreal-engine) Unreal Data Assets | Q07 | Search | Metadata only | Title matches. |
| [S23](https://create.roblox.com/docs/reference/engine/classes/ModuleScript) Roblox ModuleScript | Q07 | Docs source | Confirmed | Returns exactly one value via `require()`. |
| [S24](https://optuna.readthedocs.io/en/stable/) Optuna | Validation | Docs source | Confirmed | Title matches. |
| [S25](https://www.gesetze-im-internet.de/englisch_urhg/englisch_urhg.html) German Copyright Act, English | Q08 | Search | Metadata only | Snippet has section 23 "sufficient distance" wording. |
| [S26](https://www.copyright.gov/ai/) USCO Copyright and AI | Q08 | Search | Metadata only | Overview page exists; Parts 1 to 3 listed. |
| [S27](https://one-piece.com/greg/o20150520_0399/index.html) Greg column, 20 May 2015 | E02 C01 | Search | Metadata only | Title "第35回：May the Fourth be with you!" matches; body not read. |
| [S28](https://www.youtube.com/watch?v=yB9shgf-PJ0) Gum Gum Pistol clip | E02 C02 | Search | Metadata only | Title and 31 Dec 2023 match; channel not shown. |
| S29, S31 OpenAI and Gemini pricing | Cost calculator | Search | Metadata only | Titles match; rates not adopted. |
| [S30](https://platform.claude.com/docs/en/about-claude/pricing) Anthropic pricing | Cost calculator | Fetched | Confirmed | Title "Pricing"; caching changes cost. |
| [S34](https://github.com/Tencent-Hunyuan/Hunyuan3D-2.1) Hunyuan3D 2.1 | Assets | Raw README and LICENSE | Confirmed | 10/21/29 GB VRAM; licence excludes EU, UK, South Korea. |
| [S35](https://json-schema.org/draft/2020-12/schema) 2020-12 meta-schema | `unit.schema.json` | Fetched | Confirmed | `$id` matches; the 09-24 fetch had failed. |
| [N01](https://www.gesetze-im-internet.de/urhg/__23.html) UrhG section 23 | Q08 | Search | Metadata only | Title "Bearbeitungen und Umgestaltungen". |
| [N02](https://json-schema.org/draft/2020-12) Draft 2020-12 | S35 | Fetched | Confirmed | Authors, date, IDs match. |
| [N03](https://arxiv.org/abs/2506.15442) Hunyuan3D 2.1 paper | S34 | Search plus S34 citation block | Metadata only | Title and ID match; DOI and 53-name list unchecked. |

Counts for the 38 external entries: 19 Confirmed, 19 Metadata only, 0 Contradicted.
Internal records: 2 Confirmed, 2 Unverified, 1 Stale link.

### Contradictions

- [`drafts/90-knowledge/README.md`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/README.md?plain=1#L10) calls this folder a "static demo", but the folder and its ZIP contain no demo.
  [`README-verification.txt`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/README-verification.txt) says the HTML and package were "separate attachments".
- S25: the historical entry records "Full page reviewed" on 2026-09-23, while the 2026-09-24 annotation says "Original URL retrieval failed" ([`source-verification.json` L389-L404](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/source-verification.json#L389-L404)).
  The historical claim remains uncorroborated.
- S11: the historical `use` says the terms mention "non-commercial use"; the annotation warns against a "blanket noncommercial-only rule" ([L165-L180](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/source-verification.json#L165-L180)).
  The docs source contains both, so both notes are incomplete rather than wrong.
- S03: the historical entry dates the page 17 September 2026 and the annotation reads a 23 September footer ([L37-L52](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/source-verification.json#L37-L52)).
  The page was blocked today, so this is unresolved.
- Access depth: the historical entries claim "Full page reviewed" for S08, S14, S15, S21 to S24 and S29 to S31, but the 09-24 annotations reached only page identity or metadata.
  For S15 the annotation could not confirm the heading ([L229-L244](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/source-verification.json#L229-L244)), and today's index title is "Gemini API".
- S35: the annotation says direct retrieval failed; today it succeeded. This is an access difference only.
- The handover sets aside other project knowledge ([L25](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/UG-anime-td-systems-handover.md?plain=1#L25)), and D07 proposes building a small name-to-specification pilot ([L116-L124](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/UG-anime-td-systems-handover.md?plain=1#L116-L124)).
  The [unit-generator](https://github.com/mardwerk/unit-generator) repository had a generation CLI from 2026-09-08 (`091bae8`) and now has `research`, `generate` and `review` commands (`01000bf`).
  D07 is not reconciled with that implementation.

### Gaps and unresolved claims

- The nine package files, including the HTML, are missing from this import, from the matching ZIP and from product history.
  Resolving E03 needs files whose SHA-256 values match [`verification.json`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/verification.json), such as `reference-combat.js` `ac76c7d697ff...`, followed by a rerun with Node and a 2020-12 validator.
- Browser checks at 320 to 1440 px, the H05 screenshots, and the original conversation (H01, H02, H06 quotes) are not in the import and cannot be checked.
- Claims that depend on the missing files cannot be checked: tier-5 DPS 440 and 990, the 450, 6,900 and 4,830 economy values, the claim that the package has no LICENSE file, and the embedded-export equality.
- The canon claims C01 and C02 need readable primary pages (one-piece.com, the Crunchyroll clip) or another licensed primary source.
  Spoiler and variant scope (Q05) is still open.
- The rights questions (Q08) are open: reuse of character material, the missing software licence, AniList commercial thresholds and the Hunyuan3D territory exclusion.
  The legal pages S25, S26 and N01 were checked for metadata only.
- Thirteen blocked documentation pages (OpenAI, Google, Meshy, ElevenLabs, Unity, Unreal) were checked for metadata only.
  A future fetch from an unblocked network would be needed to confirm their content.
- Whether the original 34-entry `sources.json` was preserved "unchanged" cannot be checked without that file (recorded as 12,782 bytes, `dbf5ddccd006...`).

## Raw evidence

| Raw file | SHA-256 (first 12 hex) | Role | Current use |
|---|---|---|---|
| [`UG-anime-td-systems-handover.md`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/UG-anime-td-systems-handover.md) | `249f3ef9e000` | Handover: decisions, evidence, questions, register | Replaced by this summary; keep for the Q01 to Q08 wording |
| [`source-verification.json`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/source-verification.json) | `3640a49b1225` | 34 historical source entries plus S35 and N01 to N03 | Keep as raw data |
| [`verification.json`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/verification.json) | `63a72b7048bf` | Recorded hashes, tool versions and schema result for missing files | Keep as raw data; the only hash record for the missing package |
| [`node-verification.txt`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/node-verification.txt) | `4e13e6dd15fe` | Recorded 37/37 test output, 2026-09-24 | Historical only; not reproduced |
| [`README-verification.txt`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-systems-verification/README-verification.txt) | `ebcadcaa868c` | Scope note for the supplement | Replaced by this summary |
