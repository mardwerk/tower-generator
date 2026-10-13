# Generator research: framework and design references

These observations come from the September 26, 2026 source review. Metadata-only checks did not establish paper findings; reproduced software examples establish only the stated version and scenario. No present-day check or acceptance is implied.

These are dated source observations and proposals. Preserve their recorded review and rejection status. Current local contracts govern acceptance; these notes neither select implementation nor resume paused development.

## Summary for current use

### Verified findings


- **Contradicted**: FS says "the project already uses Zod 4".
  True on 2026-09-22 (unit-generator locked zod 4.6.5 at the inspected revision), but the Go port the inspected revision on 2026-09-26 removed Zod.
  Current main the inspected revision has a Go schema layer with "Zod-compatible issues".
- **Confirmed**: CEL separates parsing, optional type checking against declared variables and functions, and evaluation in a binding environment ([cel-spec](https://github.com/google/cel-spec) `doc/langdef.md` at the inspected revision).
  The [cel-go README](https://raw.githubusercontent.com/google/cel-go/master/README.md) recommends parsing and checking ahead of time, and says the resulting `Program` is "stateless, thread-safe, and cachable".
  That supports "compile once, evaluate repeatedly" as a recommendation.
  The cited cel.dev page was not read.


- **Confirmed**: The [Cyber Quincy](https://github.com/hemisemidemipresent/cyberquincy) repository and its `helpers/` folder exist (master the inspected revision, 2026-09-18).
  It is a BTD6 Discord bot licensed AGPL-3.0, not a data package.
  The drafts do not mention the license, and whether its data is accurate is **Unverified**.


- **Confirmed** (historical): At unit-generator the inspected revision (2026-09-22), `docs/PRODUCT.md` kept combat simulation and game-wide balance out of scope, `scripts/evaluate-btd6-roles.mjs` defaulted to `research/btd6/raw/btd6_towers.json`, and `docs/PR6-REVIEW.md` recorded PR 6 as merged without the simulator gate, as GDR and GR say.
  All three files were deleted by 2026-09-26.

### Proposals and guidance (not evidence)

- **Proposal**: Every "Adaptation" paragraph in GDR is the authors' synthesis, not a paper result.
  The same holds for the workflow, the candidate evidence record and the claim-versus-evidence ladder (GDR L451-L496).
  The ladder is the most reusable piece: it separates a new description, catalogue novelty, Engine support, tested gameplay difference, player preference and historical novelty.

### Use rules

- You may rely on the the recorded CEL facts, and on the bibliographic identity of the papers apart from the flagged rows.
- Do not cite any paper finding from GDR as checked: study sizes, runtimes, participant counts and win rates were not read here.
  Do not treat the Steam, One Piece or Crunchyroll pages, or the BTD6 data leads, as inspected or accurate.
  Take BTD6 facts from [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas) instead of these leads.
  Do not follow the workspace-relative links.
  Obsolete workspace locators are not usable current product references.

### Sources checked

| Source | Cited in | Check (2026-09-26) | Status | Note |
| --- | --- | --- | --- | --- |
| Ninja Kiwi BTD6 maps API | GR | Blocked | Unverified | Contents and terms unknown |
| Cyber Quincy repo and `helpers/` | GR | Clone the inspected revision | Confirmed | Exists; Discord bot; AGPL-3.0; data accuracy unchecked |
| Bloons Wiki BTD6 | GR | Blocked | Unverified | |
| `~/CodingProjects/btd6-atlas/` | KR | `git ls-remote`; unit-generator `research/BTD6-REFERENCE.md` | Contradicted | KR: "future", "when it publishes". A public repo exists (HEAD the inspected revision) and unit-generator main calls it authoritative. See [btd6-snapshot.md](btd6-snapshot.md) |
| Product claims (PRODUCT.md, evaluate script, PR 6) | GDR L14, GR L17, L28 | unit-generator the inspected revision | Confirmed | Historical only; files now deleted |


| Source | Cited in | Check (2026-09-26) | Status | Note |
| --- | --- | --- | --- | --- |
| R1 Sorochan and Guzdial, arXiv 2212.03387, EXAG 2022 | FS, GDR | WS | Metadata only | Abstract confirms ten microRTS units; cause/effect vocabulary and 20-game evaluation not read |
| R3 GAVEL, arXiv 2407.09388, NeurIPS 2024 | FS, GDR | WS, clone the inspected revision | Metadata only | Repo citation matches; no root LICENSE file; "48 hours on RTX8000" not read |
| R4 Gravina et al., arXiv 1907.04053, CoG 2019, DOI 10.1109/CIG.2019.8848053 | FS, GDR | WS | Metadata only | All identifiers match |
| R7 Zook, Harrison and Riedl, arXiv 1908.01423 | GDR | WS | Metadata only | FDG 2015 version is titled "...Simulation-based Play Strategy Analysis" |
| R12 Collins et al., arXiv 2508.10914, CogSci 2025 | GDR | WS | Metadata only | Abstract says "over 450" games; draft's 484 not checked |
| R14 FunSearch, Nature 625:468-475, DOI 10.1038/s41586-023-06924-6 | GDR | WS, clone the inspected revision | Metadata only | README confirms the repo omits the LLM and execution sandbox |
| R15 Ludax, arXiv 2506.22609 | GDR | WS, clone the inspected revision | Contradicted | Title is ["...Domain Specific Language for Board Games"](https://arxiv.org/abs/2506.22609); the draft says "Description Language". v2 date (2026-03-25) unchecked |
| R16 Si, Yang and Hashimoto, arXiv 2409.04109, ICLR 2025 | FS, GDR | WS | Metadata only | ICLR 2025 proceedings entry exists |
| R17 Mortar, Nasir et al., GECCO 2026 | FS, GDR | WS | Metadata only | Authors and venue match; DOI 10.1145/3795095.3805100 not found (GECCO '26 prefix); uncited arXiv 2601.00105 exists |
| R18 Mechanic Miner, EvoApplications 2013, DOI 10.1007/978-3-642-37192-9_29 | FS, GDR | WS | Metadata only | Abstract: levels generated for a discovered mechanic, consistent with "supporting challenge" |
| R19 Mechanic Maker 2.0, arXiv 2309.09476, AIIDE 2023 | GDR | WS | Metadata only | OJS id 27522 matches DOI; pages unchecked |
| R20 Pixie, Cook, AIIDE 21(1):206-215 | GDR | WS | Metadata only | Abstract: annotated Unity projects, evolved compiled snippets; full text not read |
| R21 Zook and Riedl, AAAI 28(1) 2014, arXiv 1908.01420 | GDR | WS | Metadata only | |
| R22 TowerMind, arXiv 2601.05899, DOI 10.1609/aaai.v40i31.39818 | GDR | WS, clone the inspected revision | Metadata only | Repo names Dawei Wang as first author; MIT license |
| R23 Liu et al., FDG 2019, DOI 10.1145/3337722.3337723 | GDR | WS | Metadata only | Abstract confirms Kingdom Rush: Frontiers components and Monte Carlo testing |
| R24 Kraner et al., NT 2021, DOI 10.1007/978-3-030-75275-0_54 | GDR | WS | Metadata only | Abstract confirms Save the Sheep, Unity, and GA-generated map, placement and waves |
| R25 Avery et al., CEC 2011, DOI 10.1109/CEC.2011.5949738 | GDR | WS | Metadata only | Abstract mentions a TD component classification and an EDPCG prototype |
| R27 Jiang et al., arXiv 2206.13623, FDG 2022 | GDR | WS | Metadata only | DOI 10.1145/3555858.3563273 unconfirmed |
| R28 Earle, Jiang and Togelius, arXiv 2408.12525 | GDR | WS | Metadata only | arXiv title starts "PCGRL+:" (IEEE CoG 2024 per snippet); draft calls "PCGRL+" a supplied label, gives no venue |
| R29 Xu and Verbrugge, arXiv 2508.18533, AIIDE 21(1):367-376 | GDR | WS | Metadata only | DOI matches the ACM listing |
| R32 Hastings, Guha and Stanley, TCIAIG 1(4):245-263, 2009 | GDR | WS | Metadata only | DOI and the 1,000-player claim unchecked |
| R35 AutoUE, Yin et al., arXiv 2603.07106 | GDR | WS, clone the inspected revision | Metadata only | Also in ACL 2026 Findings; draft says "2026 preprint". Repo: 20 demo games, Apache-2.0 |
| R37 Anderson and Falconer, Entertainment Computing 54:100963 | GDR | WS | Metadata only | ScienceDirect lists it as open access, so the full-text gap may be closable |
| R38 GameASG-Bench, arXiv 2609.21293 | FS, GDR | WS | Metadata only | Authors and 2026-09-18 date match; abstract numbers match; not reproduced |
| R2, R5, R6, R8, R9, R10, R11, R13, R26, R30, R31, R33, R34, R36 | GDR | WS | Metadata only | Title, authors, year, venue and cited identifier all match (R2 TCIAIG 2(1); R5 1504.04909v1; R6 EC 19(2); R8 AIIDE 2012; R9 1809.02232; R10 EVCO_a_00123; R11 FDG 2013; R13 2407.14095; R26 GPEM 14; R30 2306.09310; R31 2312.09067; R33 2405.15568; R34 10.1145/3706598.3714233; R36 2402.15391). Attributed findings not read |
| Further references: Yannakakis and Togelius 2011; Doshi and Hauser 2024; Wenger and Kenett 2026; Shin et al. 2023; Boden 1998 | GDR | WS | Metadata only | Each cited DOI, volume and issue match |
| Leads: Butler, Siu and Zook FDG 2017; Talakat GECCO 2018; Togelius et al. CIG 2010; Yannakakis, Liapis and Alexopoulos FDG 2014 | GDR | WS | Metadata only | |
| Lead: Fontaine et al. GECCO 2019 | GDR | WS | Contradicted | Title is "...through MAP-Elites with Sliding Boundaries"; the draft says "with MAP-Elites" |

### Contradictions

- Dates: FS and GDR say they were checked on September 22, and GR and GDR carry `2026-09-22` in their names.
  All three link to `archive/GENERATOR-SURVEY-2026-09-21.md`, a path that first existed on 2026-09-25.
  FS equals unit-generator's the inspected revision version (2026-09-25) except for one rewritten sentence, so the texts were last edited on or after 2026-09-25.
  The check date is not the text date.
- Reading-guide ownership: GDR L30 says its papers supersede the September 21 survey "as the research reading guide".
  FS L3 and the cold-store survey give that role to unit-generator's research index.
  Current unit-generator `research/README.md` says the framework notes and survey now live in this private repository.
- FS L8 says "The project already uses Zod 4", but unit-generator main no longer depends on Zod.
  KR treats btd6-atlas as an unpublished local project, but unit-generator main cites a public btd6-atlas repository.


- GR L28 revises an earlier cold-store version (`scratch-legacy-2026-09-22/moved-framework-section.md`).
  That version said "PR 6 adds a narrow simulator"; GR says the merged experiment excludes the simulator.
  The unit-generator PR 6 review at the inspected revision supports GR.

### Gaps and unresolved claims

- The Steam, One Piece, Crunchyroll, Ninja Kiwi and Bloons Wiki pages are unchecked: the hosts are blocked and the search budget ran out.
  Checking them needs a browser session or an allowed host.
- No paper finding was read.
  This covers counts (1,389 games, 19 viable, 14 participants, 484 games), runtimes (GAVEL 48 hours) and limitations.
  Checking them needs arXiv or publisher access.
  The overlapping `research/` review in [agentic-generation-research.md](agentic-generation-research.md) may have read some of these papers.
- Still unconfirmed: the Mortar DOI, the R27 and R32 DOIs, R19 pages, the Ludax v2 date and GAVEL's license position (the repo has no root LICENSE file).
- Pixie, R23, R24, R25 and R32 primary texts remain access gaps, as the drafts say.
  R37 may be readable through ScienceDirect.
- The reuse terms of Cyber Quincy (AGPL-3.0) and the Ninja Kiwi API need a licensing decision before any data import.


- An owner should decide whether GDR or the `research/` handover (`UG-R-game-design.md`, checked 2026-09-24) is the canonical game-design reading guide.

Source identifiers: FS means Framework source notes; GR means Generator references; GDR means the game-design research collection; KR means the source collection index. These names identify cited inputs, not current product contracts.
