# Research for Tower Generator

Research should answer a concrete question needed to design a Tower.
Start with the existing [character Wiki](../RESEARCH.md), the selected Profile and the current [authoring contract](../AUTHORING.md).
The [initial skill](../../.agents/skills/tower-generator/SKILL.md) connects these inputs to a reviewable design.
Its [SkillOpt plan](SKILLOPT.md) describes later measured improvement.

Kyle selected this focused transfer on 2026-10-03 in [#108](https://github.com/mardwerk/tower-generator/issues/108).
The existing product outline and workflow review remain in [#106](https://github.com/mardwerk/tower-generator/issues/106) and [#107](https://github.com/mardwerk/tower-generator/issues/107).
This transfer adds no backend, generic export interface or research prerequisite.

Use the Wiki's cited findings, canon scope, limitations and human review records before collecting more evidence.
Fetch additional evidence only when its absence blocks a specific design choice.
Keep character facts separate from game adaptations, declared Profile compliance and gameplay judgments.
The current Usopp authoring input supports a narrow projectile experiment, rather than arbitrary character export.

## Preserved research state

The private Planning [source snapshot](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/README.md) preserves the research checkout's documents and exploratory Gojo experiment, including local-only work.
[manifest.json](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/manifest.json) records the source commit and each file's SHA-256.
The raw material remains private because the source repository is private and Tower Generator is public.
The initial skill and public contracts do not require archive access.
Imported files retain their original wording and links; their claims, instruction boundaries and proposals describe the source project.
They do not establish Tower Generator requirements or override its instructions.

| Reference | Use when needed |
| --- | --- |
| [Research library](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/docs/RESEARCH_LIBRARY.md) | Locate primary sources and their recorded inspection depth. |
| [Search log](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/docs/SEARCH_LOG.md) | Check access failures, search limits and unfinished discovery. |
| [Corpora](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/docs/CORPORA.md) | Recover the complete user-supplied game and character lists. |
| [Inventory](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/docs/INVENTORY.md) | Trace the source project's evidence provenance. |
| [State of the art](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/docs/STATE_OF_ART.md) | Find research leads and unresolved questions. |
| [Recommendation](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/docs/RECOMMENDATION.md) | Inspect an unaccepted architecture and experimental proposal. |
| [Gojo protocol](https://github.com/mardwerk/planning/blob/main/99-cold-store/unit-design-research-2026-10-03/experiments/001_gojo_pvz/preregistration/protocol.md) | Inspect exploratory history, not a validated product benchmark. |

No scholarly claims were newly verified or promoted by this transfer.
The snapshot contains abstract-level claims, external temporary-file references and inconsistencies between its older foundation scope and later proposals.
Review a primary source before using a claim to justify a product decision.
The Gojo code uses bundled claims and a small custom simulation; its outputs do not establish canon fidelity, Plants vs. Zombies runtime behavior or balance.

Knowledge graphs, constraint solvers, evolutionary search, MCTS and human-study machinery remain source proposals.
Add one only when a current requirement and observed failure justify it.
Keep the snapshot as evidence; write new product guidance beside the relevant implementation.
