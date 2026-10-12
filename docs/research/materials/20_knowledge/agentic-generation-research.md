# Agentic generation research

Reviewed summary for [issue #1](https://github.com/mardwerk/project/issues/1).
Reviewed on 2026-09-26 against the raw import at commit [`d72f64f`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge), verified byte for byte against the [import manifest](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/MANIFEST.json).
Use it for the bibliographic status of the cited sources and as the list of open Unit Generator proposals; do not cite it as evidence that any paper finding, generation approach, framework split or license applies to Unit Generator.

## Question and scope

The four files are handovers of earlier chat threads, written on 2026-09-24.
They ask whether Unit Generator should generate tower defense units through a prompt-directed agent instead of a research-paper-based pipeline, how to separate a stable framework from replaceable design logic and game rules, how to handle rules that change mid-generation, what research says about AI design taste and unit, mechanic, map or game generation, and whether structured natural-language decomposition helps.
The drafts README calls them "proposals, not roadmap".

In scope: [`UG-agentic-core.md`][core], [`UG-agentic-generation.md`][gen], [`UG-R-game-design.md`][gd] and the German [`UG-R-natural-language-game-design.md`][nl].
Out of scope, cross-referenced only where they conflict: the reading guide [`game-design-research-2026-09-22.md`](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design-research-2026-09-22.md) (see [framework-and-design-references.md](framework-and-design-references.md)) and [`game-design/`](https://github.com/mardwerk/project/tree/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design) (see [game-and-character-design.md](game-and-character-design.md)).

File codes: C = `UG-agentic-core`, G = `UG-agentic-generation`, R = `UG-R-game-design`, N = `UG-R-natural-language-game-design`, followed by the draft's record ID.

## Summary for current use

### Checked findings

- **Confirmed** (all four read in full): no thread produced code, a schema, a generated unit, a dataset or a measured comparison. "Agentic is simpler", "private tuning helps" and "beats direct prompting" are recommendations or hypotheses, as the drafts state (C E06, G E07, R E01, N E02 to E04).
- **Confirmed**: Unit Generator already has an implementation the drafts did not see.
  Its [ARCHITECTURE.md at `01000bf`](https://github.com/mardwerk/unit-generator/blob/01000bf1aa87feb5b3b19265312f03850472d807/docs/ARCHITECTURE.md) describes one fixed route: research without a model call, prepare one request, draft through a planning call and a mechanics call with a repair budget of 0 to 2, deterministic checks, an optional separate model review, render.
  It has "no automatic review-and-redraft loop".
  This answers C Q03 and G Q04 and is not an agent loop.
  Its existence says nothing about output quality: the owner reports a generation quality regression with no isolated cause, recorded in the [Tower Generator workstream](../30_projects/unit-generator/README.md#earlier-generation-work).
- **Confirmed**: Anthropic's "Building effective agents" separates workflows (predefined code paths) from agents (model-directed) and advises the simplest solution, adding agentic complexity only when its latency and cost pay off.
  It says nothing about Unit Generator.
- **Confirmed**: "Demystifying evals for AI agents" recommends code, model and human graders, multiple trials, not grading exact tool-call paths, and calibrating LLM judges against human experts.
- **Confirmed**: Anthropic reports multi-agent research at about 15 times chat tokens with fast-growing coordination complexity, and Managed Agents separating session log, harness and sandbox behind interfaces.
  Both describe Anthropic's systems, not Unit Generator boundaries.
- **Confirmed**: JSON Schema's documentation says complex formats need a structural phase plus semantic checks in general-purpose code.
  OWASP LLM01:2025 (project source text) treats external content as an injection vector and recommends least privilege, human approval of high-risk actions and segregating untrusted content.
- **Metadata only**: all 45 cited papers exist with matching title, authors, year and identifier.
  None looks fabricated and no identifier resolved to another title.
  Three DOIs are unconfirmed (Mortar, Tanaka and Simo-Serra in IEEE ToG, PCGRLLM).
  The findings attributed to the papers were not read, because arXiv and publisher hosts are blocked.
- **Metadata only**: the cautions the drafts draw from those papers are still the useful reading list.
  Unit and TD generation works only inside supplied vocabularies (R E02, E07).
  Agent skill ordering and human preference disagreed in one of three Mortar pairs (R E05).
  Higher novelty ratings coexist with homogenised output (R E06).
  Grammar guidance raised Ludii parsability from 0.23 to 0.68 but not text similarity or playability (N E05, E06).
  LLM self-evaluation made PCGRLLM's search worse (N E07).
  A 93.2% best check pass rate coexisted with 55.3% best task success in GameASG-Bench (R E12); search abstracts match these two figures.
- **Confirmed** (authors' README): the Si, Yang and Hashimoto repository reports a follow-up execution study in which LLM ideas lost more score than expert ideas after execution and rankings flipped.
  R E06 cites only the earlier novelty result.
- **Confirmed**: cited code carries mixed licenses.
  GAVEL has no root license and bundles a Ludii fork under CC BY-NC-ND 4.0; the Wenger and Kenett and Mechanic Maker 2.0 repositories have no license file; the rest are MIT or Apache-2.0.
  Check before any FOSS reuse.
- **Confirmed** (license texts only): Apache-2.0 section 4 allows different terms for modifications, MIT allows sublicensing and sale with notice, AGPL-3.0 section 13 covers network use.
  No thread selected a license.
  No legal conclusion is drawn here; the fair-use FAQ and Open Source Definition pages were unreachable.

### Proposals (not verified, only classified)

- **Proposal** (all four): compare against a competent direct-prompt baseline with the same model, reference material, output format and comparable budget, over repeated runs.
  Measures: acceptance without repair, editing effort, rule adherence, cost, latency, diversity, character fidelity; balance and novelty judged separately.
  Depends on an output contract and rubric that do not exist.
- **Proposal**, competing treatment arms: one bounded agent with research and revision tools (C D01, D07; G D01, D10); generate, check declared constraints, one bounded revision (R section 6); explicit intermediate decomposition plus schema validator (N).
  The product route is closest to R's arm.
  Rationales rest on Metadata-only papers or Anthropic practitioner guidance.
- **Proposal** (C D02; G D03, D07): split a runner (execution, tools, budgets, records, final checks), a replaceable designer (instructions, examples, model settings), a game profile (rules, output structure, checks) and the application.
  Rationale: public and private configurations without a fork.
  Test: swap designers without editing shared code.
  Cautioned by Metz's "The Wrong Abstraction" (Metadata only).
- **Proposal** (C D03, D04; G D04, D09): version the rules, apply authorised changes only between calls, re-check or regenerate afterwards, and never let the designer weaken a requirement to pass.
  Final checks stay outside the agent: code for structural and mechanical constraints, rubric and human review for quality.
  Depends on an undecided authority model.
- **Proposal** (G E06, Q08): treat researched content as untrusted evidence before giving research tools privileges.
  Supported in general by OWASP (Confirmed); no injection test exists.
- **Proposal** (R D02, D03): a headless TD simulator for one mechanic with its map and waves, an evaluator population and a diversity archive.
  R itself limits this to cases where behavioural claims cannot be checked from the design record.
  Depends on Mortar and Mechanic Maker findings (Metadata only).
- **Proposal** (N): never treat an LLM-only judge as ground truth.
  Depends on PCGRLLM (Metadata only).
  The product's optional model review is not a search signal, but no human calibration of it is recorded.
- **Proposal** (all four): keep fine-tuning, multi-agent teams, plugin infrastructure, RL, quality-diversity archives, 3D maps, whole games, engine code, assets and balance claims out of v1 until a measured failure justifies them.
- **Proposal** (owner context in all four): build a Towerright-first pilot, then extract a standalone FOSS generator whose basic workflow runs without private Towerright components.
  Private "tuning" is undefined.
  Owner direction, not a recorded decision.
- **Proposal** (G E04, C E04): start with about 20 representative briefs (plain characters, transformations, summons, contradictory sources, impossible briefs, a mid-run rule change), reviewed blind.
  No thresholds or rubric were set.

## Review record

### Method and access

Read all four raw files in full and matched their SHA-256 against the manifest and the blobs at `d72f64f`.
Deduplicated 60 cited sources: 45 papers, 14 documentation, license or essay pages, one unidentified corpus.
Fetched the five Anthropic articles, the JSON Schema page and the Apache license and grepped them for bylines, dates and attributed statements.
Ran `git ls-remote` and blobless shallow clones of the 11 cited research repositories and `tsunehiko/rlgdg` to read READMEs and license files, plus OWASP's LLM Top 10 source repository and SPDX license texts via `raw.githubusercontent.com`.
Ran 56 WebSearch queries before the shared session search budget ran out.
Read the public `unit-generator` clone (main `01000bf`) for current architecture and history; cold-store handovers `UG-T01` to `UG-T05` were used only as provenance.

Blocked on 2026-09-26: arxiv.org, doi.org, Crossref, OpenAlex, Semantic Scholar, dblp, opensource.org, copyright.gov, genai.owasp.org, sandimetz.com, martinfowler.com, refactoring.com, publisher sites (IEEE, ACM, AAAI, Springer, Nature, OUP, PMLR), ResearchGate and author homepages.
No repository code was run.

Workspace-relative links: none in the four files, so no Stale link entries.
None of the four files ever existed in `unit-generator`.

### Sources checked

Check codes: Fetch = page fetched today; WS = WebSearch result matched title, authors, year, venue and identifier; Repo = public repository cloned and README or license read.

| Source | Cited in | Check (2026-09-26) | Status | Note |
| --- | --- | --- | --- | --- |
| Building effective agents, Anthropic, 2024-12-19 | C S04, G S10 | Fetch | Confirmed | Byline "Erik S. and Barry Zhang"; old URL redirects; notes tooling changed since 2024 |
| Equipping agents with Agent Skills, 2025-10-16 | C S05 | Fetch | Confirmed | Byline matches; now notes an open standard (2025-12-18) |
| Demystifying evals for AI agents, 2026-01-09 | C S07, G S12 | Fetch | Confirmed | Byline matches |
| Scaling Managed Agents, 2026-04-08 | G S11 | Fetch | Confirmed | Byline matches |
| How we built our multi-agent research system, 2025-06-13 | G S14 | Fetch | Confirmed | Byline matches |
| JSON Schema, "What is a schema?" | C S06 | Fetch | Confirmed | Two-phase validation passage present |
| Apache License 2.0 | C S08, G S-N01 | Fetch | Confirmed | Section 4 wording matches |
| Open Source Definition, OSI | C S09, G S17 | WS | Metadata only | Snippet shows criterion 6 and 2024-02-16 page date; v1.9 date not seen |
| MIT License, OSI | G S15 | SPDX text | Confirmed | Text only; OSI page blocked |
| AGPL-3.0, OSI | G S16 | SPDX text | Confirmed | Section 13 present; claimed redirect unchecked |
| OWASP LLM01:2025 | G S13 | OWASP GitHub source | Confirmed | Website blocked |
| U.S. Copyright Office fair-use FAQ | G S18 | WS | Unverified | Host blocked; only a third-party archive copy found |
| Metz, The Wrong Abstraction, 2016 | G S19 | WS | Metadata only | |
| Fowler, Refactoring | G S20 | WS | Metadata only | Redirect to refactoring.com unchecked |
| Owner's research-paper corpus | C S10, G S09 | None possible | Unverified | No titles supplied |

Repositories (all exist, **Confirmed** for the statements named): `gdrtodd/gavel` (no root license; `ludii_fork` CC BY-NC-ND 4.0), `NoviScl/AI-Researcher` (MIT), `arguslab-duke/llms_are_homogeneous` (no license), `google-deepmind/funsearch` (Apache-2.0; omits LLM and sandbox), `tb6147877/TowerMind` (MIT; benchmark and level editor), `Harcurio/MechanicMiner` (no license; Unity recreation with RL evaluation), `maxencefaldor/omni-epic` (Apache-2.0), `jennyzzt/omni` (MIT), `areal-project/GameASG-Bench` (Apache-2.0; 47 tasks in 12 genres), `tsunehiko/ggdg` (Apache-2.0), `bic4907/pcgrl-llm` (Apache-2.0), `tsunehiko/rlgdg` (README says Conference on Games 2025).
The per-paper table closes this record.

### Contradictions

- Treatment arm: C's historical advice to "make the agentic version your new baseline to beat" ([C L179][core-l179]) conflicts with R's constraint-check-and-revise arm ([R L234-L257][gd-rec]) and N's decomposition arm ([N L82-L88][nl-rec]).
  The product had already chosen otherwise: the since-deleted [`research/GENERATION-APPROACHES.md`](https://github.com/mardwerk/unit-generator/blob/7797add67668325e08355bd7db1cc70e57c8bc33/research/GENERATION-APPROACHES.md) (2026-09-21) set "one compact generation call plus bounded, targeted repair" and said to avoid adding agents wholesale; current ARCHITECTURE.md keeps one fixed route.
- Stale premise: C and G treat the existing approach as unavailable, yet `unit-generator` held a staged authoring CLI from 2026-09-19 (`fb94418`) and planned-v1 authoring with UnitLab from 2026-09-21 (`737e783`).
  The drafts report only what each thread saw.
- Access records: R says it inspected manuscripts of Liu et al., Kraner et al., Avery et al., Galactic Arms Race and Pixie ([R L380-L488][gd-s18]); the sibling guide records HTTP 403 or unavailable text for the same sources two days earlier ([R20 to R25, R32](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design-research-2026-09-22.md?plain=1#L244-L344)).
  Neither could be rechecked.
- Venues: R S38 records no verified peer-reviewed venue for AutoUE ([R L511][gd-s38]); search lists Findings of ACL 2026, pages 2341 to 2364.
  N S08 types RLGDG as a preprint ([N L198][nl-s08]); the authors' README says Conference on Games 2025.
- Missing report: N E02 says the natural-language thread never delivered its announced report ([N L50][nl-e02]).
  The German [`game-design/` report](https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/game-design/user-supplied-natural-language-design.txt), received 2026-09-21 on the same topic with the same `cite...` token style, may be that output; its authorship is unestablished.
- Internal: R says tracking parameters were removed ([R L279][gd-l279]) but footnotes 17 and 27 keep `utm_source=chatgpt.com`.
  G's register jumps from S05 to S09 ([G L345-L349][gen-s09]).
  C D04 and G D06 assumed engine-consumable output that the owner context excludes; both drafts flag this (C Q02, G Q01).

### Gaps and unresolved claims

- Content attributed to every paper (R E02 to E13, N E05 to E07) is unchecked and needs arXiv or publisher access.
  Search abstracts matched only for GameASG-Bench, Sorochan and Guzdial, Collins et al., Si et al. and Shin et al.
- Unconfirmed DOIs: Mortar `10.1145/3795095.3805100` (prefix matches other GECCO 2026 DOIs; the paper is also arXiv:2601.00105, not cited), Tanaka and Simo-Serra `10.1109/TG.2024.3520214` (volume 18(1), pages 30 to 43 and IEEE Xplore document 10807354 matched; one search summary attributed the DOI to an unrelated paper without a supporting page), PCGRLLM `10.1109/TG.2026.3695197`.
- Unchecked after the search budget ran out: the 2026 v2 of Collins et al., PCGRL+'s CoG venue, Mechanic Miner's seven-versus-five participants, Kraner's five players, the Galactic Arms Race player count, OSD v1.9 dating, the fair-use FAQ.
- The underlying chat messages, timestamps and URLs are not included; the owner's paper corpus is unidentified; N has 30 unresolvable `citeturn...` tokens.
- Before a pilot decision: an output contract, a rule-change authority model, a review rubric, and a license and reference-rights review of actual dependencies.

### Per-paper bibliographic checks

All rows checked by WS; "+Repo" means the repository was also read.

| Source | Cited in | Check (2026-09-26) | Status | Note |
| --- | --- | --- | --- | --- |
| Sorochan, Guzdial 2022, RTS units, arXiv:2212.03387, EXAG | R S06 | WS | Metadata only | Ten microRTS units in abstract |
| Browne, Maire 2010, Evolutionary Game Design, TCIAIG 2(1) | R S07 | WS | Metadata only | |
| Todd et al. 2024, GAVEL, NeurIPS, arXiv:2407.09388 | R S08 | WS+Repo | Metadata only | License caveat |
| Gravina et al. 2019, PCG through QD, CoG, arXiv:1907.04053 | R S09 | WS | Metadata only | DOI on ACM listing |
| Collins et al. 2025, human invention, arXiv:2508.10914 | R S10 | WS | Metadata only | CogSci non-archival |
| Yannakakis, Togelius 2011, EDPCG, IEEE TAC 2(3) | R S11 | WS | Metadata only | |
| Si, Yang, Hashimoto 2025, ICLR, arXiv:2409.04109 | R S12 | WS+Repo | Metadata only | Follow-up study above |
| Doshi, Hauser 2024, Sci. Adv. 10(28) eadn5290 | R S13 | WS | Metadata only | Preprint 2312.00506 matches |
| Wenger, Kenett 2026, PNAS Nexus 5(3) pgag042 | R S14 | WS+Repo | Metadata only | |
| Shin et al. 2023, PNAS 120(12); arXiv:2303.07462 | R S15 | WS | Metadata only | |
| Romera-Paredes et al. 2024, FunSearch, Nature 625 | R S16 | WS+Repo | Metadata only | |
| Boden 1998, Creativity and AI, AI 103 | R S17 | WS | Metadata only | |
| Liu et al. 2019, TD levels using PCG, FDG | R S18 | WS | Metadata only | Access disputed |
| Kraner et al. 2021, custom TD game, LNNS 233 | R S19 | WS | Metadata only | |
| Avery et al. 2011, CI and TD games, CEC | R S20 | WS | Metadata only | |
| Wang et al. 2026, TowerMind, AAAI 40(31), arXiv:2601.05899 | R S21 | WS+Repo | Metadata only | |
| Togelius et al. 2013, map generation, GPEM 14 | R S22 | WS | Metadata only | |
| Togelius et al. 2010, StarCraft map space, CIG | R S23 | WS | Metadata only | ITW DOI prefix as indexed |
| Jiang et al. 2022, 3D level generators, FDG, arXiv:2206.13623 | R S24 | WS | Metadata only | |
| Earle et al. 2024, PCGRL+, arXiv:2408.12525 | R S25 | WS | Metadata only | Venue unchecked |
| Xu, Verbrugge 2025, 3D levels, AIIDE, arXiv:2508.18533 | R S26 | WS | Metadata only | |
| Raistrick et al. 2023, Infinigen, CVPR | R S27 | WS | Metadata only | |
| Yang et al. 2024, Holodeck, CVPR | R S28 | WS | Metadata only | |
| Nasir et al. 2026, Mortar, GECCO | R S29 | WS | Metadata only | DOI not found |
| Cook et al. 2013, Mechanic Miner, EvoApplications | R S30 | WS | Metadata only | |
| Jara Gonzalez et al. 2023, Mechanic Maker 2.0, AIIDE | R S31 | WS+Repo | Metadata only | |
| Cook 2025, Pixie, AIIDE 21(1) | R S32 | WS | Metadata only | Access disputed |
| Zook, Riedl 2014, mechanic generation, AAAI | R S33 | WS | Metadata only | |
| Hastings et al. 2009, Galactic Arms Race, TCIAIG 1(4) | R S34 | WS | Metadata only | Access disputed |
| Guzdial, Riedl 2018, conceptual expansion, AIIDE | R S35 | WS | Metadata only | |
| Faldor et al. 2025, OMNI-EPIC, ICLR, arXiv:2405.15568 | R S36 | WS+Repo | Metadata only | |
| Earle et al. 2025, DreamGarden, CHI | R S37 | WS | Metadata only | |
| Yin et al. 2026, AutoUE, arXiv:2603.07106 | R S38 | WS | Metadata only | ACL 2026 Findings per search |
| Bruce et al. 2024, Genie, ICML PMLR 235 | R S39 | WS | Metadata only | |
| Zook et al. 2015, MCTS strategy analysis, FDG | R S40 | WS | Metadata only | |
| Zhang et al. 2024, OMNI, ICLR, arXiv:2306.01711 | R S41 | WS+Repo | Metadata only | |
| Khalifa et al. 2020, PCGRL, AIIDE, arXiv:2001.09212 | R S42 | WS | Metadata only | |
| Anderson, Falconer 2025, Entertainment Computing 54 | R S43 | WS | Metadata only | |
| Zhang et al. 2026, GameASG-Bench, arXiv:2609.21293 | R S44 | WS+Repo | Metadata only | Submitted 2026-09-18 |
| Green et al., critical mechanic discovery, arXiv:1909.03094 | R S45 | WS | Metadata only | |
| Tanaka, Simo-Serra, GGDG, arXiv:2407.17404, IEEE ToG 18(1) | N S04 | WS+Repo | Metadata only | DOI unconfirmed |
| Ebner et al. 2013, VGDL, Dagstuhl Follow-Ups 6 | N S05 | WS | Metadata only | |
| Baek et al. 2024, ChatPCG, CoG, arXiv:2406.11875 | N S06 | WS | Metadata only | |
| Baek et al., PCGRLLM, IEEE ToG, arXiv:2502.10906 | N S07 | WS+Repo | Metadata only | v2 2026-05-25; DOI not found |
| Tanaka, Simo-Serra 2025, RLGDG, arXiv:2503.15783 | N S08 | WS+Repo | Contradicted | Draft: preprint; authors: CoG 2025 |

## Raw evidence

| Raw file | SHA-256 (first 12 hex) | Role | Current use |
| --- | --- | --- | --- |
| [`research/UG-agentic-core.md`][core] | `83354331a624` | Handover on one agent, replaceable core and changing rules | Replaced by this summary; keep raw for proposal history |
| [`research/UG-agentic-generation.md`][gen] | `5adc9bd977cb` | Handover on agent-led generation, framework naming and FOSS/private boundaries | Replaced by this summary; keep raw for proposal history |
| [`research/UG-R-game-design.md`][gd] | `0efc141b4a73` | Literature review on AI taste and unit, mechanic, map and game generation | Replaced by this summary; per-paper notes overlap [framework-and-design-references.md](framework-and-design-references.md) |
| [`research/UG-R-natural-language-game-design.md`][nl] | `c9bb50e69179` | German handover of a thin thread plus leads on grammar-guided generation | Replaced by this summary |

[core]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-agentic-core.md
[core-l179]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-agentic-core.md?plain=1#L179
[gen]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-agentic-generation.md
[gen-s09]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-agentic-generation.md?plain=1#L345-L349
[gd]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-game-design.md
[gd-rec]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-game-design.md?plain=1#L234-L257
[gd-l279]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-game-design.md?plain=1#L279
[gd-s18]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-game-design.md?plain=1#L380-L488
[gd-s38]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-game-design.md?plain=1#L511-L515
[nl]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-natural-language-game-design.md
[nl-e02]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-natural-language-game-design.md?plain=1#L50
[nl-rec]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-natural-language-game-design.md?plain=1#L82-L88
[nl-s08]: https://github.com/mardwerk/project/blob/d72f64f5eb45a854e9074940fd8f24b036800d53/drafts/90-knowledge/research/UG-R-natural-language-game-design.md?plain=1#L198-L204
