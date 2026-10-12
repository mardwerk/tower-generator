# Unit design research

Moved from unit-generator `research/README.md` at `01000bf` on 2026-09-26.
These are working paper notes for the [Unit Generator workstream](README.md), not product contracts.
The product repository documents [BTD6 source attribution](https://github.com/mardwerk/tower-generator/blob/main/docs/PROFILE.md#btd6-source-attribution) alongside Profile ownership.
[unit-generator#26](https://github.com/mardwerk/unit-generator/pull/26) originally moved the BTD6 reference from `research/` to `docs/`.

Current guidance for Unit Generator is on the [workstream page](README.md) and in the [product issues](https://github.com/mardwerk/unit-generator/issues).
Links below to commit `9244cd5` point to historical documents that were removed from the product's main branch.

These working notes cover papers that frame representation and evaluation, plus the pointer to BTD6 facts below.
The natural-language [game-design synthesis](../../99-cold-store/UG-R-natural-language-game-design.md) is in the cold store; [game and character design](../../20_knowledge/game-and-character-design.md) is the reviewed summary.
The historical [next-work note](https://github.com/mardwerk/unit-generator/blob/9244cd5/docs/NEXT-EXPERIMENT.md) recorded the roadmap as of that commit; it is not the current roadmap.

The [bounded optimization proposal](AUTORESEARCH.md) describes a future experiment, not an implemented autonomous training loop.

## Papers most relevant to unit generation

Use these papers to frame representation and evaluation, not to choose a pipeline by reputation.
Findings and access status come from the September 22 investigation; no experiments were reproduced.
The optional workspace catalogue owns the full reading and cross-project connections.
Primary links below remain usable from a standalone checkout.

| Question | Reading | Applicable idea and limit |
| --- | --- | --- |
| How should a generated unit be represented and tested? | [RTS unit generation, R1](https://arxiv.org/html/2212.03387v1) | Explicit attributes and cause/effect abilities. Its one-map microRTS evaluator and bounded vocabulary do not validate character fidelity or our mechanics. |
| Does a mechanic contribute across combinations? | [Mortar, R17](https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf) | Evaluate mechanics within games and retain varied candidates. Automated skill ordering disagreed with human preference in one of three tested pairs; contribution is not preference. |
| What scenario makes a mechanic useful and understandable? | [Mechanic Miner, R18](https://www.possibilityspace.org/papers/evo13.pdf) | Generate a supporting challenge, then test understanding. Solver usefulness did not establish enjoyment; program-field mutations form a bounded search space. |
| Can the evaluator use a new mechanic? | [Mechanic Maker 2.0, R19](https://arxiv.org/html/2309.09476v3) | Evaluator choice changes discovered rules. RL can exploit objectives and is not a validated human proxy. |
| How do we preserve meaningful alternatives? | [GAVEL, R3](https://arxiv.org/html/2407.09388v2) and [quality diversity, R4](https://arxiv.org/abs/1907.04053) | Separate variation, feasibility and evaluation. Archive coverage is not demonstrated strategic diversity; GAVEL's roughly 48-hour runs do not establish interactive latency. |
| What distinguishes many passing checks from a complete result? | [GameASG-Bench, R38](https://arxiv.org/html/2609.21293v1) | Declare acceptance scenarios independently and report whole-task success. The recent preprint measures software compliance, not novelty or fun. |

These are a closer fit than the old card/monster examples.
Runtime claims still require a compatible Consumer; concept work can retain predicted scenarios and attributed review.
The historical [REFINEMENT.md](https://github.com/mardwerk/unit-generator/blob/9244cd5/docs/REFINEMENT.md) described the study protocol at that commit. The [earlier generation work](README.md#earlier-generation-work) remains deferred until after the validator.

## BTD6 facts

BTD6 domain facts live in [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas).
Tower Generator's [Profile documentation](https://github.com/mardwerk/tower-generator/blob/main/docs/PROFILE.md#btd6-source-attribution) records citation requirements and historical source provenance.
Do not add BTD6 facts to this repository.

Generated runs and live provider responses remain in the ignored `data/runs/`; they are not source references or public fixtures.
Manga Mayhem's historical [enemy baseline](https://github.com/mardwerk/manga-mayhem/blob/2017a5b853b5229dd51e5fa2090b53c968db8821/docs/BTD6-ENEMY-BASELINE.md) is preserved at source revision `2017a5b`. Later game-owned [Inkling definitions](https://github.com/mardwerk/manga-mayhem/issues/86) and [wave definitions](https://github.com/mardwerk/manga-mayhem/issues/85) remain parked for a later gameplay milestone.

[Current mechanics documentation](https://github.com/mardwerk/unit-generator/blob/main/docs/MECHANICS.md) owns implementation guarantees.
