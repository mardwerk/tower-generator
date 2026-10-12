# BTD6 snapshot

Reviewed summary for [issue #1](https://github.com/mardwerk/project/issues/1).
Reviewed on 2026-09-26 against the raw import at commit [`d72f64f`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge), verified byte for byte against the [import manifest](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/MANIFEST.json).
The BTD6 source for Mardwerk is [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas): take BTD6 facts from its pinned captures, not from this snapshot.
Use this summary as the review record of a historical snapshot.
It shows where the snapshot's Medium prices and local upgrade rule agree with the atlas and two other exports, and where they do not.
Do not use it for combat stats, crosspath effects or world rules, or as evidence that the proposed kernel works.

## Question and scope

The group is a 2026-09-24 research snapshot with two parts.
The first is a Medium-price catalog of BTD6 towers, upgrades, heroes and Paragons, with a pricing and legality helper.
The second is a proposed "compositional kernel" (six record kinds, ten effect instructions) for describing tower behavior.
It answers two questions.
What does each ordinary build cost at Medium?
Which upgrade tuples are locally legal?
It also proposes a data model that could express BTD6-like units.

In scope: all nine files of [`btd6-snapshot/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot), listed under [Raw evidence](#raw-evidence).
The cold-store `99-cold-store/zips/BTD6_model_and_catalog.zip` was opened as provenance only.
It holds byte-identical copies of the same nine files, dated 2026-09-24 11:49 to 11:54, workbook included.

Out of scope: the other draft groups, the private cold store, and a full review of the [unit-generator](https://github.com/mardwerk/unit-generator) product repository.
For unit-generator, only its BTD6 pointer was checked (see below).

## Summary for current use

### What was confirmed

These results describe the snapshot.
For current BTD6 facts, cite [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas) by capture and file.

- **Confirmed**: the catalog has Medium prices for 26 standard towers (base price plus three paths of five incremental tiers, 390 upgrades).
  All 416 prices equal cyberquincy [`costs.json` at `d0a813e`](https://github.com/hemisemidemipresent/cyberquincy/blob/d0a813e4fd900552e960481621c64a35846f108e/jsons/costs.json).
  They also equal the fees in btd6-game-data `Upgrades/*.json` and the tower model `cost` fields at [`f818c39` ("56.0")](https://github.com/Btd6ModHelper/btd6-game-data/tree/f818c39e6d7cc4e0f8fb02a2e3d1a929bcfe6a24), and the btd6-atlas 56.3 capture at [`a380413`](https://github.com/KyleDerZweite/btd6-atlas/tree/a380413ed809c98654acec7aa37cba40f80bf5c5/data/56.3-build-24829026).
  Prices are incremental.
  A tier-5 fee excludes tiers 1 to 4 and the base.
- **Confirmed**: the 18 hero deployment prices equal cyberquincy [`heroes.json`](https://github.com/hemisemidemipresent/cyberquincy/blob/d0a813e4fd900552e960481621c64a35846f108e/jsons/heroes.json) and the hero model `cost` fields in both game-data exports.
  The 13 Paragon fees equal btd6-game-data `<Tower> Paragon.json`, cyberquincy `paragon_costs.json` and the atlas.
  The catalog's Paragon baseline is three uncrosspathed tier-5 parents plus the fee (for example, Dart 95,415 + 150,000 = 245,415).
  This is arithmetic only and does not fix Paragon degree.
- **Confirmed**: "Medium" is the right label for these numbers. cyberquincy applies a difficulty multiplier of 1 for Medium to `costs.json` and `heroes.json` values (`helpers/bloons-general.js`).
  The draft excludes Monkey Knowledge, support discounts, relics, social perks and special-mode adjustments.
  Easy, Hard and Impoppable prices are not in the catalog.
- **Confirmed**: the local tuple rule is: sort the three tiers descending into a, b, c; then a <= 5, b <= 2 and c = 0.
  This gives 64 tuples per tower, 000 included, and 1,664 tower and tuple rows.
  Recomputation agrees.
  For each of the 26 towers, the set of exported tower model files in btd6-game-data (56.0) and btd6-atlas (56.3) is exactly these 64 tuples.
  The rule does not check placement, ownership, unlocks, world tier-5 limits, uniqueness or donor requirements.
  A Paragon is a separate transaction, not a legal 555.
- **Confirmed**: [`btd6_costs.py`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/btd6_costs.py) reproduces every stated result.
  Dart 203 is 1,405.
  Beast Handler with three uncrosspathed tier-4 donors is 107,920 (500), 111,280 (050) and 79,680 (005).
  Monkey Village 005 is marked incomplete at a known 15,200 unless the caller supplies a Monkeyopolis quote (25,200 with the example 10,000).
  Its `--self-test` output matches [`validation_result.json`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/validation_result.json).
- **Confirmed**: the CSVs and the 11-sheet workbook agree with `btd6_catalog.json` field for field.
  That includes the error below, so do not treat them as independent sources.
- **Contradicted**: 388 of 390 upgrade display names match btd6-game-data `textTable.json`.
  Heli Pilot bottom path tiers 1 and 2 are swapped, together with their semantic sketches.
  The game data has 001 = Faster Darts (250) and 002 = Faster Firing (350), where the catalog has 001 = Faster Firing and 002 = Faster Darts.
  The prices by tier are correct.
  The error is in this frozen snapshot only: btd6-atlas at `a380413` has 001 = Faster Darts (250) and 002 = Faster Firing (350), checked in its `Upgrades` files on 2026-09-26.
  Take the names from btd6-atlas.

### What the snapshot does not establish

- The draft's own limits still apply (**Confirmed** as its stated scope).
  The catalog does not transcribe every damage, rate, pierce, timing, targeting, crosspath or hidden-stat field.
  It is not a simulator.
  Each upgrade has one qualitative `semantic_sketch` and pattern tags.
  A per-model URL is an audit reference, not a claim that each field was read.
  Hero records omit the 20 level deltas and Geraldo's shop.
  Paragon records omit degree stats and ability state machines.
  Powers Pro, Legends, map or quest units, modes and skins are not inventoried.
- **Unverified**: the special-case pricing semantics.
  Monkeyopolis is treated as a dynamic price; the static export only shows fee 5000 and a `MonkeyopolisModel` with farm parameters.
  Beast Handler tier 5 is said to need three extra same-path tier-4 handlers; the 500 model has `BeastHandlerPetModel.requiredCount: 3`, which fits the claim but does not prove the placement or path rules.
  The Sun Temple and True Sun God sacrifice exclusions and the claim that hero deployment price differs from unlock currency and paid XP were not checked.
- **Unverified**: the claims that "selected v56 changes [were] cross-checked with official notes" and that "a layout preview was reviewed".
  Neither has a record, and the Reddit and Steam sources are blocked.
- **Proposal**: the 390 semantic sketches and `dominant_patch_patterns` tags.
  Only the Heli pair and Sharp Shots were checked against models.
  Combat stats, crosspath effects, timing, stacking and world rules are absent from the data.

### Proposals (classified, not verified)

- **Proposal**: [`kernel_spec.json`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/kernel_spec.json) describes itself as a "design specification, not an implemented game engine" and makes no minimality claim.
  It has six persistent record kinds.
  Template holds typed stats, tags, components and rules.
  Entity is a live instance with owner, state, position and timers.
  Rule is event plus guard plus selector plus ordered actions.
  Modifier is a source-keyed patch with stacking, expiry and inheritance.
  Relation is a typed link with a contribution policy.
  Transaction is a guarded, priced, atomic purchase.
  It has ten effect instructions: `write_state`, `spawn`, `despawn`, `damage`, `apply_modifier`, `remove_modifier`, `set_motion`, `transform`, `update_relation` and `schedule_event`.
  It also declares 17 event phases, a nested expression language with bounded control, and seven intrinsic runtime services (geometry, clocks, motion, damage resolution, stat resolution, transactions, seeded randomness).
- **Proposal**: [`example_behaviors.json`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/example_behaviors.json) gives symbolic Dart, Bomb, homing and upgrade snippets, not stat-complete definitions.
  Its one numeric fact is **Confirmed**: Sharp Shots raises the Dart projectile's pierce from 2 to 3 in the game models.
- **Proposal**: the catalog's 12 upgrade patterns, 14 semantic capabilities, 11 relation families, 21 derived-entity families and the README's "semantic obligations".
  The counts are **Confirmed**, but the content is authored design, and no source row supports it.
- **Proposal**: of the 20 validation cases, T01 to T04 are marked `executed_by_catalog_helper` and do run in the helper.
  T05 to T20 are marked `specified_not_run` and are test specifications for a runtime that does not exist.
  The draft says no gameplay equivalence benchmark was run.

### Relation to unit-generator and btd6-atlas

unit-generator (main at `01000bf`) neither uses nor cites this snapshot; its history never mentions `btd6_catalog`, `btd6-snapshot` or the workbook.
It carried a separate hand-compiled 2026-09-20 package (`btd6_towers.json`, SHA-256 `a2a5e2bb4591...`: 26 towers, 390 upgrades, 13 Paragons, calculated Easy, Hard and Impoppable prices, 71 ability records), retired on 2026-09-26.
[`research/BTD6-REFERENCE.md`](https://github.com/mardwerk/unit-generator/blob/01000bf1aa87feb5b3b19265312f03850472d807/research/BTD6-REFERENCE.md) now names btd6-atlas (56.3 capture) as the authority for roster, tiers, costs and crosspath legality.
The retired package has the same Medium prices as this catalog, except that it leaves Monkeyopolis unpriced.
It calls Druid 500 "Superstorm" (the internal ID) and has the same Heli Pilot swap.
For product BTD6 facts, use btd6-atlas; its capture manifest still has an empty `acceptedBy`.
This snapshot agrees with the atlas on every price checked.

## Review record

### Method and access

- Read in full: README, `btd6_costs.py`, `kernel_spec.json`, `example_behaviors.json` and `validation_result.json`.
  Parsed every record of `btd6_catalog.json`, both CSVs and all 11 workbook sheets with scripts.
  All nine files match the manifest SHA-256 and the cold-store zip.
- Ran `--self-test`, the five README commands and the Beast 050 and 005 donor quotes on scratch copies (Python 3.11.15).
  The self-test asserts legality examples, the 64-tuple count, Dart 203, the Village 005 quotes, the three Beast totals, 26 towers with 390 unique upgrade IDs, 3x5 path dimensions and path price equal to upgrade price per tier.
  It does not check the CSVs, workbook, heroes, Paragons, other counts or any upstream source.
  Its returned counts are literals, and `python -O` skips the assertions.
- Recomputed from the data: all counts, the tuple rule, all 1,664 CSV rows against the helper (cost, completeness, requirements, warnings), `upgrade_catalog.csv` against the JSON (every field), pure-path subtotals and Paragon totals.
  Workbook (openpyxl 3.1.5 in a scratch venv): 416 formulas, all with cached values and no error values; a scripted evaluation of every formula equals the cached values and the JSON.
  Monkeyopolis's subtotal is the literal "DYNAMIC".
- Upstream, 2026-09-26: fetched `costs.json`, `heroes.json`, `textTable.json` and 27 model files from `raw.githubusercontent.com`.
  Blobless clones were used to read upstream history and to check all 403 cited model files and 790 upgrade files.
  Today's HEADs: cyberquincy master `d0a813e4fd900552e960481621c64a35846f108e` (2026-09-18), btd6-game-data main `f818c39e6d7cc4e0f8fb02a2e3d1a929bcfe6a24` (2026-08-07, "56.0") and btd6-atlas main `a380413ed809c98654acec7aa37cba40f80bf5c5` (2026-09-25).
- Blocked: `reddit.com` and `steamcommunity.com` (proxy CONNECT 403).
  The WebSearch tool was unavailable because the session's search budget was already spent.
  No search-based metadata check of the v56 date or the v55 notes was possible.

### Sources checked

| Source | Cited in | Check (2026-09-26) | Status | Note |
| --- | --- | --- | --- | --- |
| cyberquincy `jsons/costs.json` | README; every tower and upgrade row | Fetched; compared 26 base and 390 upgrade prices | Confirmed | 416/416 equal. Blob `00c4271d8a07` has been unchanged since [`1c8e46f`](https://github.com/hemisemidemipresent/cyberquincy/commit/1c8e46f28d0856181dc1841f834e0c2e7ff9324b) (2026-08-30) and is the same at HEAD, so the 2026-09-24 retrieval saw this content. |
| cyberquincy `services/wiki/costs_scraper.js` | Not cited | Read at `d0a813e` | Confirmed | Since [`c20a697`](https://github.com/hemisemidemipresent/cyberquincy/commit/c20a69716da9283a720c26e42cdb0c33e3a84883), `costs.json` has been filled from the bloonswiki.com cargo table, a fan wiki. The draft calls it the "main price snapshot" without saying this. |
| cyberquincy `helpers/bloons-general.js` | Not cited | Read at `d0a813e` | Confirmed | Medium multiplier is 1, so file values are Medium. |
| cyberquincy `jsons/heroes.json` | README; 18 hero rows | Fetched; compared | Confirmed | 18/18. Blob `59f45f22c311` unchanged since `47a26de` (2026-07-19). |
| cyberquincy `jsons/paragon_costs.json` | Not cited | Compared | Confirmed | 13/13 Paragon fees equal. |
| btd6-game-data `textTable.json` | README; `name_source` of 390 rows | Fetched (identical to `f818c39`); matched by path, tier and fee | Contradicted | 388 match. Heli Pilot 001 and 002 names are swapped (see Contradictions). |
| btd6-game-data `Upgrades/*.json` | README; Paragon `cost_source` | Sparse checkout at `f818c39` | Confirmed | 390/390 upgrade fees and 13/13 Paragon fees equal. |
| btd6-game-data per-model URLs | `mechanic_source` (390); Paragon `mechanic_reference` (13) | All 403 present at `f818c39`. 27 fetched from raw (one per tower plus Heli 001), byte-identical. | Confirmed | `tiers`, `baseId` and the base `cost` field match. The per-model `cost` is the base tower price, and the upgrade fee is in `Upgrades/`. |
| btd6-game-data `Towers/<id>/` file sets | Not cited | Tree listing | Confirmed | Exactly the 64 legal tuples for each of 26 towers. The atlas 56.3 capture gives the same result. |
| btd6-game-data `Skins/DanDMonke/DanDMonkeHeMan.json`, `UpdatePopupDescription55` | README He-Man sentence | Read | Confirmed | SkinData has `baseTowerName: DanDMonke` and full name "Prince Adam". The in-game v55 text says "New Hero, Dan D'Monke and He-Man skin". |
| "Mechanically identical" He-Man skin | README ([L29-L30](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/README.md?plain=1#L29-L30)) | Skin file read | Unverified | The skin carries asset swaps and `mutatorId: morph_mutator`, whose effect was not traced. Steam notes are blocked. |
| v56 update and date | README header; catalog `version_scope` | `UpdatePopupDescription56` in the game text; btd6-game-data commit "56.0" | Metadata only | A v56 exists (Skywarden, Portable Lake Pro, map Ascent). The export commit is dated 2026-08-07, but the release date is not established. |
| Reddit v56 notes URL | README | curl | Unverified | Blocked (403). Post ID and content unchecked. |
| Steam news URL | README | curl | Unverified | Blocked (403). |
| btd6-atlas 56.3 capture | Not cited (named in the [90-knowledge README](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/README.md)) | Blobless clone at `a380413` | Confirmed | All 26 base, 390 fee, 13 Paragon and 18 hero prices equal. Its game data has the same Heli order as btd6-game-data. |
| Game models for Monkeyopolis and Beast Handler tier 5 | README special cases | Read `MonkeyVillage-005` and `BeastHandler-500` | Unverified | The models have the matching behaviors (`MonkeyopolisModel`; `requiredCount: 3`), but a static export cannot confirm purchase-time pricing or donor rules. |
| unit-generator `research/BTD6-REFERENCE.md` | Not cited | Read at `01000bf`; history search | Confirmed | Does not use this snapshot. Differences are listed above. |

### Contradictions

- Heli Pilot bottom tiers 1 and 2.
  [`btd6_catalog.json` L5621-L5667](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/btd6_catalog.json#L5621-L5667) has 001 "Faster Firing" (sketch "shorten firing interval") and 002 "Faster Darts".
  [`upgrade_catalog.csv` L162-L163](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/upgrade_catalog.csv?plain=1#L162-L163) and the workbook Upgrades sheet say the same.
  The catalog's own cited model, [`HeliPilot-001.json`](https://raw.githubusercontent.com/Btd6ModHelper/btd6-game-data/main/Towers/HeliPilot/HeliPilot-001.json), has `appliedUpgrades: ["Faster Darts"]`.
  `Upgrades/Faster Darts.json` is path 2, tier index 0, cost 250, and btd6-atlas 56.3 agrees.
  Prices by tier are unaffected.
- Druid 500 name.
  The catalog has "Monarch of Storms"; unit-generator's retired 2026-09-20 package has "Superstorm".
  The game text maps internal ID `Superstorm` to display name "Monarch of Storms" (in btd6-game-data since its 50.0 export), which supports this catalog.
- Price provenance.
  The README presents cyberquincy as the "main price snapshot" beside game-model exports.
  The cyberquincy scraper shows those prices come from bloonswiki.com.
  The values still agree with two independent game-data exports, so this is a provenance gap, not a price error.

### Gaps and unresolved claims

- The draft pins no upstream commit and no game build.
  The review infers the pins: cyberquincy `d0a813e` (the `costs.json` content dates from 2026-08-30) and btd6-game-data `f818c39` (56.0).
  All prices also match the atlas 56.3 capture of 2026-09-25, whose manifest is not yet accepted.
  A later patch could change any of this.
- The official v56 and v55 notes (Reddit, Steam) could not be read, and no web search was possible.
  The "selected v56 changes" check and the layout preview have no record.
- These were not checked: combat stats, crosspath deltas, timing, stacking, hero level progressions, Geraldo items, Paragon degree rules, the actual Monkeyopolis quote, Beast donor placement, and Sun Temple sacrifice semantics.
  Checking them would need an in-game build-pinned test or a reviewed model-field extraction.
- 388 of 390 semantic sketches and pattern tags were not reviewed against models.
- The kernel has no implementation, schema or conformance test.
  The example snippets use only declared events (`attack_ready`, `after_hit`) and declared instructions.
  However, their field names (`select`, `actions`, `cost`, `patch`) do not follow the spec's record fields (`selector`, `ordered_actions`, `cost_expression`, `atomic_effects`).
  The upgrade patch operator `add` is not in the instruction set, and the bomb projectile has no collision block.

## Raw evidence

All paths below are under [`drafts/90-knowledge/btd6-snapshot/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot) at `d72f64f`.

| Raw file | SHA-256 (first 12 hex) | Role | Current use |
| --- | --- | --- | --- |
| [README.md](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/README.md?plain=1) | `63a9f3cd6f4f` | Scope, pricing rules, kernel overview, sources | Replaced by this summary |
| [btd6_catalog.json](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/btd6_catalog.json) | `10998ce47e23` | Master data: towers, upgrades, heroes, Paragons, authored patterns and cases | Keep as raw data. Historical only; take BTD6 facts from btd6-atlas. Its Heli Pilot 001/002 names are swapped. |
| [upgrade_catalog.csv](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/upgrade_catalog.csv) | `0d9505b19131` | Flat export of the 390 upgrades | Keep as raw data (same Heli error) |
| [locally_legal_builds.csv](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/locally_legal_builds.csv) | `a7268be8abbe` | 1,664 helper quotes | Keep as raw data (regenerable from the helper) |
| [BTD6_minimal_representation.xlsx](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/BTD6_minimal_representation.xlsx) | `1ab17bd56da2` | 11-sheet workbook view of the catalog | Keep as raw data (same Heli error) |
| [btd6_costs.py](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/btd6_costs.py) | `32350cc06c40` | Pricing and local legality helper | Keep as raw code. Reproduced; scope limited to own-tower prices. |
| [validation_result.json](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/validation_result.json) | `67d13653614d` | Recorded self-test output | Historical only (reproduced today) |
| [kernel_spec.json](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/kernel_spec.json) | `1ab732fb329d` | Proposed kernel | Keep as proposal input, not a requirement |
| [example_behaviors.json](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/btd6-snapshot/example_behaviors.json) | `fd37833efaf1` | Symbolic behavior examples | Keep as proposal input, not a requirement |
