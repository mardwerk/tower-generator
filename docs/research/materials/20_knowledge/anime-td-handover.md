# Anime TD handover (UG-anime-td)

Reviewed summary for [issue #1](https://github.com/mardwerk/project/issues/1).
Reviewed on 2026-09-26 against the raw import at commit [`d72f64f`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge), verified byte for byte against the [import manifest](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/MANIFEST.json).
Use it as the reviewed record of the "Unit Forge" Luffy demonstrator and its proposals; do not use it as evidence of ONE PIECE canon, game balance, a working name-to-unit generator, or an approved architecture.

## Question and scope

The source conversation asked whether a standalone system could turn an anime character name into a complete tower defense unit concept, with Monkey D. Luffy as the example.
The delivered artifact (dated 2026-09-23) is a self-contained research website, a 38-entry bibliography, 10 qualitative architecture assessments, a starter skill, a demo game contract, a Luffy unit specification, a strict sample schema and a single-tower combat kernel.
A later [handover](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/handover.md) (dated 2026-09-24) re-inspected it, reran the kernel into `rerun/`, and added a new recommendation (D07).

In scope: all 31 files under [`anime-td-handover/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover) (the review request said 30; the folder and the import manifest both hold 31).
Out of scope: the conversation messages S41, S42 and S47, which are not in the import; the original `unit_forge_research.zip`, which is not in this folder (a copy with the same SHA-256 exists in the cold store and was consulted only as historical provenance); the private cold store, which was not read.

The other anime TD demo revision, `drafts/90-knowledge/anime-td-systems-verification/`, is reviewed separately in [anime-td-systems-verification.md](anime-td-systems-verification.md).
No evidence, test result or conclusion is pooled between the two reviews.

## Summary for current use

### Verified findings

- **Confirmed**: every present file listed in [handover_manifest.json][hm] (13 of 14 entries) and all 16 entries in [artifact_manifest.json][am] match by SHA-256 and size.
  The 14th handover entry, `original/unit_forge_research.zip`, is absent from this folder.
- **Confirmed**: [anime_td_research.html][html] and [index.html][idx] are byte-identical (`0a93e0cc7a04`).
  The 17 extracted files equal the members of the cold-store `unit_forge_research.zip` (SHA-256 `ec6fcdf38518`, as recorded in [verification.json][ver]) byte for byte.
- **Confirmed**: in `original/` versus `rerun/`, five of six shared files are byte-identical (`demo_game_contract.json`, `luffy.unit.json`, `simulator.js`, `tests.js`, `unit.schema.json`).
  Only `kernel_test_report.json` differs, and only in `executed_at` (2026-09-23T10:27:31Z versus 2026-09-24T09:14:17Z).
- **Confirmed**: `node tests.js` (Node v22.22.2) passes 20 of 20 checks in scratch copies of both folders.
  Apart from `executed_at` and `environment`, each new report equals the recorded report, including all 12 illustrative fixtures.
  The "60 configurations" are 5 tiers x 6 fixed scenarios x active on/off inside one test.
  They are not 60 scenarios or a character benchmark.
- **Confirmed**: [luffy.unit.json][unit] validates against [unit.schema.json][schema] with Python `jsonschema` 4.26.0 (Draft 2020-12, meta-schema check included) and with Ajv 8.20.0 in strict 2020 mode.
  [validation_report.json][vrep] is the kernel report plus one `format_validation` block.
- **Confirmed**: `python3 build_site.py`, run on a scratch copy, regenerates `index.html` and the parent `anime_td_research.html` byte for byte.
  It also rewrites `sources.json` and `README.md` in place; on the delivered files this leaves them unchanged.
- **Confirmed** (partial): headless Chromium (Playwright 1.56.1), loading the scratch `index.html` via `file://` at 1440x1050 and 390x844, reported no page errors, no console errors, no external requests and no horizontal overflow.
  It showed 12 chapters.
  The 28 interaction checks in [browser_test_report.json][btr] were not rerun and remain **Unverified**, because the original test driver is absent.
- **Confirmed**: the Luffy demo facts in the handover E02 row match the files: 5 tiers, deploy cost 600, cumulative cost 5,600, range 4.0 rising to 5.5, 20 Hz ticks, ground-only targets, one tower on a straight path.
  The archive has no license file.
- **Contradicted**: neither screenshot shows the delivered build.
  [report-desktop.png][pngd] shows 12 approaches, 32 sources, "26 checks passed", "Rubber Vanguard", 450 credits, 3.2 m range, a 9-chapter navigation and a dark theme.
  The delivered HTML shows 10, 38, 20/20, "Rubber Captain", 600, 4.0, 12 chapters and a light theme.
  [report-mobile.png][pngm] also shows the dark "A character name in.
  A game-ready unit out." layout, and that text is absent from the delivered HTML, template and app script.
  The screenshots come from an earlier build that is not available.
  They are historical and are not evidence, as the owner confirmed on 2026-09-26.

### Canon and research claims

- **Unverified**: the two canon claims in the Luffy provenance are C1, "rubber-bodied", and C2, "captain of the Straw Hat Pirates with a signature straw hat".
  Both are attributed to S35.
  The S35 page exists (**Metadata only**), but its character section could not be read because the host is blocked.
  The official character directory (S40) was never identified.
  No Gear, Haki or Devil Fruit claim appears in the delivered files.
  The desktop screenshot footnote "no Gear" is an unchecked canon allusion from a non-matching build.
- **Confirmed** at the primary page today: S01, S02, S06, S09, S16 and S37 (titles, bylines and dates, plus the attributed notes where read).
  Examples: S01 now says its tooling landscape "has changed since December 2024", and S16 reports a 90.2% internal eval gain and about 15x token use for multi-agent runs.
  Also confirmed: S39 meta-schema identity and S49, plus the README content of S18, S24 and S28 at today's repository HEAD.
  The TRELLIS.2 README states Linux-only testing and at least 24 GB of GPU memory.
  These are vendor or methods sources, not tower defense evidence.
- **Metadata only**: 29 sources, including all four arXiv papers (S11 to S14), the Google blog (S15, including its "180 configurations" claim), both Bandai Namco sites (S35, S36) and most vendor documentation.
  Their attributed content was not read.

### Proposals (not evidence)

- **Proposal**: the architecture comparison.
  [architectures.json][arch] scores 10 approaches on five 1 to 5 scales.
  The handover presents 8 of them as generator alternatives and treats "evaluator-optimizer" and "simulation-led search" as methods.
  The scores, weights and budget defaults are assumptions, not measurements.
- **Proposal**: the historical recommendation of a hybrid unit compiler (D03, which the delivered HTML also states), the game contract prerequisite (D02), the simple starting paths (D04), separate validation and approval states (D05) and conditional expansion (D06).
  No owner approval is recorded.
- **Proposal**: [demo_game_contract.json][contract] (`demo_td_v1`), the sample-specific [unit.schema.json][schema] (not a general ability language), and the starter [SKILL.md][skill], whose seven host tools are not implemented.
- **Proposal**: all Luffy TD mechanics and numbers are self-declared inventions: Stretch Punch line hits, Captain's Resolve (every 4th hit x1.5) and Elastic Barrage (6 pulses at 0.4 ratio).
  The same applies to the asset briefs.
- **Proposal**: the 540-job benchmark (6 systems x 30 characters x 3 runs) was never executed.
  D07, the handover author's 2026-09-24 pilot recommendation, awaits owner approval.
  D07 proposes a small grounded pilot compared against direct prompting.

**Rely on** the kernel, schema and build as a reproducible reference demonstrator for this one Luffy format.
**Do not rely on** the screenshots, the 26-check claim, the canon claims, the architecture scores or the bibliography as proof that any architecture works for this product.

## Review record

### Method and access

Read in full: `handover.md`, `handover_manifest.json`, `verification.json`, `source_register.md` (identical to handover section 7 except for two trailing blank lines), the `README.md`, `artifact_manifest.json`, `browser_test_report.json`, `tests.js` and `build_site.py` in the extracted folder, and both PNGs, viewed as images.
Read in part or by field: `source_register.json` (all 49 entries cross-checked against `sources.json`), `sources.json`, `architectures.json`, `luffy.unit.json`, `demo_game_contract.json`, `SKILL.md`, `simulator.js`, `validation_report.json` and the HTML, searched for specific strings.
`site_template.html` and `site_app.js` were searched only.

All runs used copies under the scratch folder `anime-td-handover/`: the kernel tests, both schema validators, `build_site.py` and a headless Chromium load.
Nothing was written under `drafts/`.

Network on 2026-09-26: fetched `www.anthropic.com`, `platform.claude.com`, `code.claude.com` and `json-schema.org` directly, and three READMEs via `raw.githubusercontent.com`.
All other cited hosts returned no connection (curl status 000).
`github.com` web pages returned 403.
The WebSearch tool supplied metadata for the rest.

`source_register.json` agrees with `sources.json` on every ID and URL, and it keeps each original record verbatim.
It retitles 13 entries; 9 of those say so in their notes, while 4 minor ones (S10, S19, S20, S33) do not.
Its dates are internally consistent: all 38 sources are marked accessed on 2026-09-23 and bibliographically checked on 2026-09-24.

### Sources checked

| Source | Cited in | Check (2026-09-26) | Status | Note |
| --- | --- | --- | --- | --- |
| S01 Building effective agents (Anthropic, 2024-12-19) | sources.json, handover | Page fetched | Confirmed | Byline "Erik S. and Barry Zhang"; tooling-change note present |
| S02 Equipping agents ... with Agent Skills (2025-10-16) | same | Page fetched | Confirmed | Three authors match; progressive disclosure and security sections present |
| S06 Structured outputs (Anthropic docs) | same | Page fetched | Confirmed | Title matches; no design-quality claim made |
| S09 Agent SDK overview (Anthropic docs) | same | Page fetched | Confirmed | Title matches |
| S16 How we built our multi-agent research system (2025-06-13) | same | Page fetched | Confirmed | Six authors, 90.2% eval gain and about 15x tokens stated |
| S37 Scaling Managed Agents (2026-04-08) | same, build_site.py | Page fetched | Confirmed | Three authors; session, harness and sandbox separation described |
| S18 DSPy, S24 ComfyUI, S28 TRELLIS.2 | same | README via raw.githubusercontent.com at HEAD 8de96ea, 79be670, 75fbf01 | Confirmed | Current HEAD, not the unpinned 2026-09-23 view |
| S39 Draft 2020-12 meta-schema, S49 JSON Schema Specification | unit.schema.json, handover | Fetched | Confirmed | `$id` equals the schema's `$schema`; raw fetch worked today, although the handover's browser could not render it |
| S11 RAG (Lewis et al., arXiv 2005.11401) | same | WebSearch | Metadata only | 12 authors and NeurIPS 2020 match; version v4 not checked |
| S12 Quest for Content (arXiv 2311.04710), S13 MAP-Elites (arXiv 1504.04909), S14 MAS failures (arXiv 2503.13657) | same | WebSearch | Metadata only | Titles, authors and IDs match; S14 v3 date not checked |
| S15 Scaling agent systems (Google Research blog) | same | WebSearch | Metadata only | Kim and Liu, 2026-01-28; "180 configurations" seen only in a snippet |
| S35 ONE PIECE ODYSSEY site; S36 PIRATE WARRIORS 4 site (Bandai Namco) | luffy.unit.json, sources.json | WebSearch; host blocked | Metadata only | Page titles match the register; Luffy character section not read |
| S03, S04, S05, S07, S08, S10, S17, S19 to S23, S25, S27, S29 to S34, S38 (vendor and official docs) | same | WebSearch | Metadata only | Titles and URLs appear in results; content unread |
| S26 Meshy Rigging API | same | WebSearch | Metadata only | "Rigging API" title appears at `/en/api/rigging-and-animation`, not at the cited `/en/api/rigging` |
| S40 official character directory | luffy.unit.json | No URL recorded | Unverified | Never identified |
| S41, S42, S47 conversation messages | handover | Not in import | Unverified | D01 to D07 quotes cannot be checked |
| S43 HTML, S44 archive | handover | Hashes, rerun, rebuild | Confirmed | S44 via the cold-store zip |
| S45, S46 screenshots | handover | Viewed; compared with headless render | Contradicted | They show a different build from S43 and S44 |
| S48 handover verification | handover | Reproduced | Confirmed | Every verification.json claim matches today's rerun except the browser, which was not rerun |
| `original/unit_forge_research.zip` link | handover.md line 6, handover_manifest.json | Path check | Stale link | Absent from this folder; a same-hash copy is in the cold store at [`zips/unit_forge_research.zip`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/99-cold-store/zips/unit_forge_research.zip) |
| Canon claims C1, C2; "no Gear" footnote | luffy.unit.json, report-desktop.png | S35 section unreadable | Unverified | Game adaptation choices are classified as Proposal |

### Contradictions

- The screenshots ([S45 and S46 rows][h260]) disagree with the delivered [HTML][html] and [artifact_manifest.json][am] on counts, unit name, cost, range, checks, layout and theme.
  The handover [E07 row][h75] records the desktop mismatch.
  This review adds three points.
  The mobile screenshot also shows the non-delivered layout.
  The desktop PNG is 1440x1000, while the browser report's viewport is 1440x1050.
  And [site_template.html][tmpl] still says "36 primary sources", which [build_site.py][build] patches to 38, so three source counts (32, 36, 38) exist across revisions.
  Revision order remains unestablished.
- The handover [attachment line][h6] and [handover_manifest.json][hm] list `original/unit_forge_research.zip`, but the import does not contain it at that path.
- The extracted [README.md][readme] "Rebuild the HTML" section says the build writes `index.html` and a standalone file.
  [build_site.py][build] also rewrites `sources.json` and `README.md` and hard-codes the S37 and S38 entries.
  This is minor, and idempotent on the delivered files.
- No disagreement was found between [verification.json][ver], the [recorded kernel report][krep], the [rerun report][rrep] and today's reruns.
  Only the Node version (v22.16.0 versus v22.22.2) and timestamps differ.

### Gaps and unresolved claims

- The 28 browser interaction checks need the original driver or a new Playwright script to be confirmed.
  Today's load check covers only errors, requests and overflow.
- C1 and C2 need the S35 Luffy section or an official ONE PIECE source.
  S40 is unidentified.
  Version and spoiler scope (Q04) is open.
- 29 sources are Metadata only.
  The original 2026-09-23 page snapshots cannot be reconstructed, and repositories were never pinned.
- The conversation sources (S41, S42, S47) are missing.
  D01 to D06 attributions rest on the handover author's reading.
  D07 and Q01 to Q08 await owner decisions ([Q table][hq]).
- The following were never supplied: a generalization beyond Luffy, the D07 direct-prompt comparison, the 540-job study, engine integration, balance calibration, rights clearance and a license.
- The build that produced the screenshots and the "26 checks" claim (Q06) is not available.

## Raw evidence

Paths are relative to [`anime-td-handover/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover).
`x/` means [`original/extracted/unit_forge_research/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research).

| Raw file | SHA-256 (first 12 hex) | Role | Current use |
| --- | --- | --- | --- |
| [handover.md][h] | 5aa247a3c314 | Handover narrative, D/E/Q/S records | Replaced by this summary |
| [handover_manifest.json][hm] | 098b365f7c8f | Handover file hashes | Historical only |
| [source_register.md][srm] | 3d926695052b | Copy of handover section 7 | Replaced by this summary |
| [source_register.json][srj] | 29bc6fbd3224 | Machine-readable register, 49 IDs | Keep as raw data |
| [verification.json][ver] | f0f4ea1b8ee3 | 2026-09-24 integrity and rerun record | Historical only |
| [original/anime_td_research.html][html] | 0a93e0cc7a04 | Delivered report | Keep as raw data |
| [original/report-desktop.png][pngd] | 85285c82b86b | Screenshot, non-matching build | Historical only |
| [original/report-mobile.png][pngm] | d5116c644332 | Screenshot, non-matching build | Historical only |
| [x/README.md][readme] | 443dc183f425 | Artifact readme and limits | Keep as raw data |
| [x/SKILL.md][skill] | b471e046629b | Starter skill | Keep as raw data (Proposal) |
| [x/architectures.json][arch] | 3024524c0e03 | 10 scored approaches | Keep as raw data (Proposal) |
| [x/artifact_manifest.json][am] | 8a7ab81d5bcc | Artifact hashes and counts | Keep as raw data |
| [x/browser_test_report.json][btr] | 575feddba282 | 28 recorded browser checks | Historical only |
| [x/build_site.py][build] | 9f68b3257cf1 | HTML builder | Keep as raw data |
| [x/demo_game_contract.json][contract] | 4e2597bacba0 | Demo game rules | Keep as raw data (Proposal) |
| [x/index.html][idx] | 0a93e0cc7a04 | Same bytes as the delivered report | Keep as raw data |
| [x/kernel_test_report.json][krep] | b74f21b2e3b8 | 2026-09-23 kernel result | Keep as raw data |
| [x/luffy.unit.json][unit] | de9cce3a5588 | Luffy reference spec | Keep as raw data (Proposal) |
| [x/simulator.js][sim] | 04d0372b3b26 | Combat kernel | Keep as raw data |
| [x/site_app.js][app] | 2c408982e697 | Site script | Keep as raw data |
| [x/site_template.html][tmpl] | 5a515690644a | Site template | Keep as raw data |
| [x/sources.json][src] | 2744566c1000 | 38-entry bibliography | Replaced by this summary |
| [x/tests.js][tests] | c9b2a1ef798b | Kernel tests | Keep as raw data |
| [x/unit.schema.json][schema] | f325c2d572f0 | Sample schema | Keep as raw data (Proposal) |
| [x/validation_report.json][vrep] | 8340bcc84ffa | Kernel plus schema result | Keep as raw data |
| [rerun/demo_game_contract.json][r1] | 4e2597bacba0 | Rerun copy, same as original | Keep as raw data |
| [rerun/kernel_test_report.json][rrep] | 6b3ce1ff60a0 | 2026-09-24 rerun result | Historical only |
| [rerun/luffy.unit.json][r2] | de9cce3a5588 | Rerun copy, same as original | Keep as raw data |
| [rerun/simulator.js][r3] | 04d0372b3b26 | Rerun copy, same as original | Keep as raw data |
| [rerun/tests.js][r4] | c9b2a1ef798b | Rerun copy, same as original | Keep as raw data |
| [rerun/unit.schema.json][r5] | f325c2d572f0 | Rerun copy, same as original | Keep as raw data |

[h]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/handover.md
[h6]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/handover.md?plain=1#L6
[h75]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/handover.md?plain=1#L75-L76
[hq]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/handover.md?plain=1#L107-L118
[h260]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/handover.md?plain=1#L260-L261
[hm]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/handover_manifest.json
[srm]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/source_register.md
[srj]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/source_register.json
[ver]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/verification.json
[html]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/anime_td_research.html
[pngd]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/report-desktop.png
[pngm]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/report-mobile.png
[readme]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/README.md
[skill]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/SKILL.md
[arch]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/architectures.json
[am]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/artifact_manifest.json
[btr]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/browser_test_report.json
[build]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/build_site.py
[contract]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/demo_game_contract.json
[idx]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/index.html
[krep]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/kernel_test_report.json
[unit]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/luffy.unit.json
[sim]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/simulator.js
[app]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/site_app.js
[tmpl]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/site_template.html
[src]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/sources.json
[tests]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/tests.js
[schema]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/unit.schema.json
[vrep]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/original/extracted/unit_forge_research/validation_report.json
[r1]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/rerun/demo_game_contract.json
[rrep]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/rerun/kernel_test_report.json
[r2]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/rerun/luffy.unit.json
[r3]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/rerun/simulator.js
[r4]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/rerun/tests.js
[r5]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/anime-td-handover/rerun/unit.schema.json
