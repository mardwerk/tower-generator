# Game and character design

Reviewed summary for [issue #1](https://github.com/mardwerk/project/issues/1).
Reviewed on 2026-09-26 against the raw import at commit [`d72f64f`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge), verified byte for byte against the [import manifest](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/MANIFEST.json).
Use this summary as unit authoring guidance and as a checked list of the BTD6 examples it relies on; do not cite the user-supplied German report, or the studies and talks it names, as evidence for design, psychology or balance claims.

## Question and scope

The group answers: which authoring principles should shape a generated Tower Defense unit's paths, purchases, feedback and review, and which BTD6 examples illustrate them?
In scope are the four files in [`game-design/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design): the synthesis [README.md](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/README.md), the condensed [CHARACTER-DESIGN-GUIDANCE.md](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/CHARACTER-DESIGN-GUIDANCE.md), [provenance.json](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/provenance.json) and the 59 KB German report [user-supplied-natural-language-design.txt](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/user-supplied-natural-language-design.txt) ("Game Design in natürlicher Sprache").
Out of scope: a full review of the BTD6 snapshot (see [btd6-snapshot.md](btd6-snapshot.md)), the reading guides in `90-knowledge/` (see [framework-and-design-references.md](framework-and-design-references.md)), and playtesting.

## Summary for current use

### Verified findings

- **Confirmed**: the retained report matches its provenance record and the manifest: 58,986 bytes, SHA-256 `4828f196...0b255`.
  All four files are also byte-identical to `research/game-design/` in the public unit-generator repository, added in `b54822c` (2026-09-22) and deleted in `a750989` (2026-09-25).
  Authorship and original publication date remain unknown.
- **Unverified**: the report's 107 inline citation markers use 45 distinct tokens of the form `turnNsearchM` / `turnNviewM`, wrapped in private-use Unicode characters (U+E200 to U+E202).
  The file contains no URLs.
  The tokens cannot be resolved to sources.
- **Metadata only**: 29 named works in the report exist with matching title, speaker or authors per search results, including the MDA paper (Hunicke, LeBlanc, Zubek 2004), Ryan, Rigby and Przybylski (Motivation and Emotion 30(4), 2006) and Ian Schreiber's one-page design docs talk (GDC 2025).
  None of the attributed content was read.
  One further claim (Godot scene composition) is **Confirmed** against the upstream documentation source.
- **Confirmed**: every report section the synthesis says it "adopts" exists (Executive Summary, Die natürliche Sprachform als Designvertrag, Workflow für natürliche Sprache, Mechanics, Systems Design, Mechanics und Systems als Motivationsmaschinen, UI/UX, Audio, Ästhetik und Art Direction, Balancing, Prototyping und Testen).
  The paraphrases are fair condensations.
  Sentences about Engine operators, Profiles and review stages are the synthesis's own additions, classified below as proposals.
- **Confirmed**: all BTD6 upgrade names and qualitative behaviors in the synthesis's mapping table agree with the BTD6 game-data export for v56.0 (text table and tower models) and the snapshot's upgrade catalog.
  This includes the crosspath examples: Sniper `4-2-0` shrapnel carries the MOAB immobilize and can hit Camo; `4-0-2` fires about twice as often as `4-0-0`; Guided Magic and Arcane Blast leave the Phoenix's damage and pierce unchanged; the Phoenix has map-wide attack range but still filters out Camo, even with Monkey Sense.
  Nothing was contradicted.
  The separate Tier 5 ownership rule remains **Unverified** here.
- **Confirmed** (historical): at unit-generator `b54822c` the planning prompt applied the character guidance, and the evaluation code compared all six directional crosspath pairings at main tiers 3 to 5.
  The prompt text now lives in `src/cli/internal/unit/plan_prompt.go` on main (`01000bf`).

### Proposals (authoring guidance; classified, not verified)

- **Proposal**, design contract: describe each purchase as player promise, situation, decision, rule (precondition, state change, cost or risk, limit), feedback and test hypothesis.
  The report's own chain also names required information, edge cases and implementation data.
- **Proposal**, path standard: each path states when to buy it, what changes, what alternative it displaces and what weakness remains.
  Different role labels or stat packages do not make different decisions.
  Tiers 1 and 2 are readable foundations and useful crosspaths; tier 3 commits through a new behavior, a substantial improvement or a narrower job; tiers 4 and 5 develop that payoff without a universal damage multiplier or other paths' capstones.
- **Proposal**, mechanics honesty: record automatic, triggered, configured and manually activated effects separately.
  A named source power grants no Engine operator and no exemption from target, range, obstruction or resource rules; unsupported powers are reserved or omitted explicitly.
- **Proposal**, feedback: state what players must distinguish (ordinary shot versus active window, immunity versus miss, purchased specialization).
  These are consumer requirements; a generated image or costume proves nothing about combat readability and grants no unpurchased effect.
- **Proposal**, evidence separation: deterministic checks establish the encoded contract, semantic review assesses coherence, playtests investigate understanding, choices and balance.
  Damage proxies and prices do not establish enjoyment or balance.
  The five review hypotheses in the [README](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/README.md?plain=1#L38-L50) are untested questions.
- **Proposal**, Profile rule: three five-tier paths, at most two purchased, only one above tier 2, the third closed.
  Legal `3-0-0`, then `3-1-0`, then `3-2-0`; illegal `3-3-0` and `3-2-1`.
  This is identical to the BTD6 ordinary local tuple rule in the snapshot (sorted a <= 5, b <= 2, c = 0; the helper returns true for 300, 310, 320 and false for 330, 321), matches the [MangaMayhem Unit Profile](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/projects/mangamayhem-unit-profile/README.md) page (unchanged since `d72f64f`) and the unit-generator default Definition (`MaxPurchasedPaths: 2, MaxAdvancedPaths: 1, CrosspathTier: 2` at `01000bf`).
  The added tier 1 and 2 "improve the basic attack" restriction is a deliberate Profile adaptation, not a BTD6 rule.
- **Proposal**, profile targets: this guidance is shared by two separate targets, Unit Generator's public BTD6-inspired Default Profile and the MangaMayhem Definition and Profile.
  Ink, transformations and status effects on placed Units belong only to the MangaMayhem Profile and must not become public defaults by implication ([unit-generator#21](https://github.com/mardwerk/unit-generator/issues/21), [#22](https://github.com/mardwerk/unit-generator/issues/22); see [Profile ownership](../30_projects/unit-generator/README.md#profile-ownership)).
- **Proposal**, character guidance in brief: build branches from sourced motions, weapons, transformations and interactions, keeping source fact separate from adaptation (hardening a fist grants no allied damage; perception grants no control).
  Treat pierce as hit capacity with explicit shared or separate budgets; bounces do not refill it; control needs an application cap.
  Consider all six directional crosspaths at tiers 3 to 5 and what each inherits (cadence does not shorten cooldowns; range does not enlarge explosions).
  Check many cheap copies before pricing; bound support and control by eligibility, capacity, locality or downtime; no broad all-source amplification in tiers 1 and 2.
  Prose edits preserve mechanics; no scores or balance certificates.

### Do not rely on

- The report's psychological, empirical and studio-process claims as evidence.
  They are at best **Metadata only**; the report is a secondary, unattributed synthesis.
- The report as a BTD6 source (the synthesis says so itself), or the table's BTD6 facts beyond the v56.0 export checked here.
  Take BTD6 facts from [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas).
- The workspace-relative links in the drafts; all are **Stale link** (see Gaps).

## Review record

### Method and access

Read in full on 2026-09-26: all four raw files, including all 859 lines of the German report.
Recomputed SHA-256 and sizes and compared them with `provenance.json` and `MANIFEST.json`.
Ran 33 WebSearch queries for the named works.
Direct fetches of all 12 primary hosts tried (including aaai.org, selfdeterminationtheory.org, learn.microsoft.com, docs.godotengine.org, gdcvault.com, megacrit.com, youtube.com) were blocked by the proxy.
Fetched from raw.githubusercontent.com: the Godot docs source, Btd6ModHelper `textTable.json` (SHA-256 `c5c221d6e4d4...`, repository HEAD `f818c39`, commit message "56.0", 2026-08-07) and the tower models `SniperMonkey-400/420/402`, `WizardMonkey-040/140/240/042` and `BoomerangMonkey-004`.
Compared them with the snapshot's `upgrade_catalog.csv`.
Ran the snapshot's `btd6_costs.py` from a scratch copy.
Traced linked paths in the unit-generator clone at `01000bf`.

### Sources checked

| Source | Cited in | Check (2026-09-26) | Status | Note |
| --- | --- | --- | --- | --- |
| Retained report bytes | provenance.json, manifest | SHA-256 and size recomputed | Confirmed | Also identical to unit-generator `b54822c` |
| 45 citation tokens | report, passim | Searched file for URLs or a source key | Unverified | No resolvable source; private-use delimiters |
| Hunicke, LeBlanc, Zubek, MDA (2004) | report L5, L800 | WebSearch | Metadata only | AAAI Challenges in Game AI workshop; abstract snippet consistent |
| Ryan, Rigby, Przybylski (2006) | report L15, L288 | WebSearch | Metadata only | DOI 10.1007/s11031-006-9051-8; abstract snippet consistent |
| Schreiber, Four One-Page Design Docs | report L11, L80 | WebSearch | Metadata only | GDC 2025; vision, pillars, loops, resource flow |
| Ingold, 80 DAYS Post-mortem | report L88 | WebSearch | Metadata only | GDC Vault 1021666 |
| Jayanth, Leading Players Astray | bibliography | WebSearch | Metadata only | GDC 2015 |
| Smith and Worch, What Happened Here? | report L90 | WebSearch | Metadata only | GDC 2010 |
| Ingold, Heaven's Vault | bibliography | WebSearch | Metadata only | GDC 2018 |
| Level Design Workshop: Designing Celeste | report L117, L816 | WebSearch | Metadata only | GDC 2017; results credit Maddy Thorson (report: "Matt") |
| Lee, An Approach to Holistic Level Design | report L178 | WebSearch | Metadata only | GDC 2017, Dishonored 2 |
| Into the Breach Design Postmortem | report L146 | WebSearch | Metadata only | GDC 2019, Matthew Davis; four years |
| Cleveland, The Design of Subnautica | report L17, L415 | WebSearch | Metadata only | GDC 2019; pillars, emotion in prototypes |
| Horizon Zero Dawn: A Game Design Postmortem | report L17, L417 | WebSearch | Metadata only | GDC 2018, Eric Boltjes |
| Dead Space UI talk (Ignacio) | report L205 | WebSearch | Metadata only | Actual title begins "Crafting Destruction:"; GDC 2013 |
| Xbox Accessibility Guidelines | report L207, L266 | WebSearch | Metadata only | XAG 103 multi-channel cues; German edition not checked |
| The Art of Monument Valley 2 | report L232 | WebSearch | Metadata only | GDC Vault 1025321 |
| HZD GPU-based procedural placement | report L234 | WebSearch | Metadata only | GDC 2017, Jaap van Muijden |
| The Synthesizer Plugins of COCOON | report L258 | WebSearch | Metadata only | GDC 2025, Jakob Schmid; five FMOD plugins |
| Alan Wake 2: Developing an Audio Vision | report L260 | WebSearch | Metadata only | GDC 2025, Richard Lapington |
| Giovannetti, Slay the Spire metrics talk | report L381 | WebSearch | Metadata only | Mega Crit; GDC video 2019 |
| Mega Crit Slay the Spire 2 Early Access FAQ | report L381 | WebSearch | Metadata only | Snippet: test experimental features, niche problems |
| Mega Crit code retrospective | report L600 | WebSearch | Metadata only | Neowsletter Sept 2024 snippet matches |
| Supergiant, Hades Early Access | report L17, L316 | WebSearch | Metadata only | "Designed for Early Access from the ground up" |
| Playtesting God of War | report L455 | WebSearch | Metadata only | Ed Dearien, GDC Vault 1026280 |
| Riot champion balance framework | report L383 | WebSearch | Metadata only | Average, skilled, elite, pro; pro uses presence; German version not checked |
| Godot design philosophy | report L504 | godot-docs source via raw.githubusercontent.com | Confirmed | Composition, scene inheritance, edits propagate; English source, not the German page |
| Embracing Push Forward Combat in DOOM | report L119 | WebSearch | Metadata only | GDC 2018, Loudy and Campbell |
| Mini Metro: When Less is More | report L203 | WebSearch | Metadata only | GDC 2017, Jamie Churchman |
| Hardspace: Shipbreaker diegetic interface | report L637 | WebSearch | Metadata only | GDC 2021, Vidhi Shah |
| West of Loathing Design Postmortem | report L522 | WebSearch | Metadata only | GDC 2018, Zack Johnson |
| The AI of Dishonored 2 | report L39, L148 | WebSearch | Metadata only | GDC 2017, Couvidou and Sadoulet; not in the report's bibliography |
| Deus Ex: Human Revolution loop talk | report L41, L292 | WebSearch | Unverified | Talk not named; candidate "Reimagining a Classic"; diagram unseen |
| Branching narrative talk (redirect versus define) | report L94 | WebSearch | Unverified | Not named; candidate "All Choice No Consequence" (GDC 2016) |
| Need-frustration research | report L55 | WebSearch | Unverified | Not named; several candidate studies exist |
| Hades patch notes; Live Ops GDC sessions; Destiny 2 and Ragnarök tool talks | report L316, L602, L524 | Not searched | Unverified | Unnamed or generic |
| BTD6 upgrade names in mapping table | README L25-L36 | textTable v56.0, upgrade_catalog.csv | Confirmed | All named upgrades present in both |
| Dart, Sniper tier 1 and 2 separation; Fireball, Sentry Gun | README L28 | Upgrade descriptions | Confirmed | |
| Tier 3 to 5 examples (Spike-o-pult to Permanent Brew) | README L29-L31 | Upgrade descriptions | Confirmed | Ultra-Juggernaut: gigantic ball splits into six Juggernaut balls |
| Tack and Farm path contrasts | README L27 | Upgrade descriptions | Confirmed | Blade Maelstrom and IMF Loan are abilities |
| Sniper `4-2-0` and `4-0-2` | README L32 | Tower models 400/420/402 | Confirmed | Shrapnel has SlowMaimMoabModel, Camo filter off; rate 0.7791 vs 1.59 |
| Phoenix: no bolt modifiers, map-wide, no Camo | README L32-L33 | Tower models 040/140/240/042 | Confirmed | Damage 5, pierce 8 unchanged; range 9999999; FilterInvisible active |
| Boomerang anti-MOAB control is scoped | README L33 | BoomerangMonkey-004 model, text | Confirmed | Press attack filters to MOAB tag, excludes BAD and Camo |
| Tier 5 scene-level ownership rule | README L36 | Not in text table | Unverified | See [btd6-snapshot.md](btd6-snapshot.md) |
| Profile rule | README L32, L36 | Snapshot helper, project page, unit-generator defaults | Proposal | Consistent everywhere checked |
| `../btd6/PATTERNS.md`, `../btd6/BTD6-UNIT-EXAMPLES.md` | README L23 | unit-generator history | Stale link | `research/btd6/`, deleted `f6b5db9` |
| `../../docs/PRODUCT.md` | README L15 | unit-generator history | Stale link | `docs/PRODUCT.md`, deleted `f6b5db9` |
| `../../docs/MECHANICS.md` | README L50 | unit-generator main | Stale link | Not in this import; exists at `docs/MECHANICS.md` on unit-generator main |
| `../../src/core/blueprint/plan.ts` | guidance L23 | unit-generator history | Stale link | Deleted `a750989`; now `plan_prompt.go` |

### Contradictions

- Inside the [report](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/user-supplied-natural-language-design.txt#L816): token `turn20search5` backs Dishonored 2 AI claims (L39, L148, L354, L634), yet the bibliography assigns it to the Celeste workshop (L816, also L117).
  The Dishonored 2 AI talk that search finds is absent from the bibliography.
  At least one attribution is wrong.
- [CHARACTER-DESIGN-GUIDANCE.md](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/CHARACTER-DESIGN-GUIDANCE.md?plain=1#L25) says the Default Profile confines an optional Active Ability "to the middle path from T4".
  The current unit-generator [MECHANICS.md](https://github.com/mardwerk/unit-generator/blob/01000bf1aa87feb5b3b19265312f03850472d807/docs/MECHANICS.md) (same text at `b54822c`) makes `path2` only the default: `manualAbilityPath` can select another path, and `null` forbids boosts.
  The guidance describes the default correctly but not the configurable rule.
  [unit-generator#22](https://github.com/mardwerk/unit-generator/issues/22), opened 2026-09-26, restates middle-path activation from `x-4-x` as the intended default of the shared design rules.
  The historical PATTERNS.md also called middle-path tier 4 abilities a template, not a rule.
- Tension within the guidance: its rules ask to name borrowed purchases and explain inheritance for all six crosspaths ([L15](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/CHARACTER-DESIGN-GUIDANCE.md?plain=1#L15)), while its application record has the model consider them privately and leaves the comparison to code ([L29](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/CHARACTER-DESIGN-GUIDANCE.md?plain=1#L29)).
  The skill's possible early utility is also broader than the preset's tier 1 and 2 policy (acknowledged at L27).
- Currency: the guidance calls cross-copy stacking an unresolved runtime contract.
  The current unit-generator default Definition (`01000bf`) sets slow, burn and stun stacking rules.
  Not reviewed in depth; the guidance may be partly superseded.

### Gaps and unresolved claims

- No named talk, paper or guideline was read in primary form (GDC Vault, YouTube and publisher sites blocked).
  Upgrading any row to Confirmed needs the talk video or slides, the paper PDF, or an accessible mirror.
- The citation tokens are unrecoverable without the original chat export and its source list.
  Six referenced sources are unnamed (need frustration, branching narrative, Deus Ex talk title, Live Ops, AAA tooling, Hades patch notes).
- German-language editions of the Xbox, Riot and Godot material were not accessed.
- BTD6 checks use the Btd6ModHelper model export for v56.0, not in-game observation.
  The Tier 5 ownership rule and Master Double Cross exception were not checked.
- The five review hypotheses have no playtest, scenario or success threshold.
  They stay open until a consumer supplies map, enemies, economy and players.
- Stale link: `../btd6/PATTERNS.md` and `../btd6/BTD6-UNIT-EXAMPLES.md` resolved to unit-generator [`research/btd6/`](https://github.com/mardwerk/unit-generator/tree/b54822c6d7ff7f77d542ecc4a47f5cabd4b881e7/research/btd6) (added `737e783`, deleted `f6b5db9` on 2026-09-26; condensed successor `research/BTD6-REFERENCE.md`, not compared).
  `../../docs/PRODUCT.md` existed in unit-generator until `f6b5db9`.
  `../../docs/MECHANICS.md` exists on unit-generator main.
  `../../src/core/blueprint/plan.ts` was deleted in `a750989`; its prompt now lives in `src/cli/internal/unit/plan_prompt.go`.
  None exists relative to the import.
  The historical PATTERNS.md does state the "no sprite audit" and "proxies are not simulated output" limits the synthesis attributes to it.

## Raw evidence

| Raw file | SHA-256 (first 12 hex) | Role | Current use |
| --- | --- | --- | --- |
| [README.md](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/README.md) | `05a91b44c469` | Synthesis: principles, BTD6 mapping, hypotheses | Replaced by this summary |
| [CHARACTER-DESIGN-GUIDANCE.md](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/CHARACTER-DESIGN-GUIDANCE.md) | `f245f64580fc` | Condensed user-supplied authoring skill and application record | Replaced by this summary for rules; application section historical only |
| [provenance.json](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/provenance.json) | `4f034c45ac95` | Provenance of the report | Keep as raw record |
| [user-supplied-natural-language-design.txt](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/user-supplied-natural-language-design.txt) | `4828f196465e` | German report with unresolved citation tokens | Keep as raw source; consult for templates only, not evidence |
