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

The [source snapshot](unit-design-research-snapshot/README.md) preserves the research checkout's documents and exploratory Gojo experiment, including local-only work.
[manifest.json](unit-design-research-snapshot/manifest.json) records the source commit and each file's SHA-256.
Kyle explicitly authorized publishing the full 56-file snapshot from the private source repository on 2026-10-03.
The manifest records this publication decision; the snapshot remains reference evidence rather than accepted product requirements.
Imported files retain their original wording and links; their claims, instruction boundaries and proposals describe the source project.
They do not establish Tower Generator requirements or override its instructions.

| Reference | Use when needed |
| --- | --- |
| [Research library](unit-design-research-snapshot/docs/RESEARCH_LIBRARY.md) | Locate primary sources and their recorded inspection depth. |
| [Search log](unit-design-research-snapshot/docs/SEARCH_LOG.md) | Check access failures, search limits and unfinished discovery. |
| [Corpora](unit-design-research-snapshot/docs/CORPORA.md) | Recover the complete user-supplied game and character lists. |
| [Inventory](unit-design-research-snapshot/docs/INVENTORY.md) | Trace the source project's evidence provenance. |
| [State of the art](unit-design-research-snapshot/docs/STATE_OF_ART.md) | Find research leads and unresolved questions. |
| [Recommendation](unit-design-research-snapshot/docs/RECOMMENDATION.md) | Inspect an unaccepted architecture and experimental proposal. |
| [Gojo protocol](unit-design-research-snapshot/experiments/001_gojo_pvz/preregistration/protocol.md) | Inspect exploratory history, not a validated product benchmark. |

No scholarly claims were newly verified or promoted by this transfer.
The snapshot contains abstract-level claims, external temporary-file references and inconsistencies between its older foundation scope and later proposals.
Review a primary source before using a claim to justify a product decision.
The Gojo code uses bundled claims and a small custom simulation; its outputs do not establish canon fidelity, Plants vs. Zombies runtime behavior or balance.

Knowledge graphs, constraint solvers, evolutionary search, MCTS and human-study machinery remain source proposals.
Add one only when a current requirement and observed failure justify it.
Keep the snapshot as evidence; write new product guidance beside the relevant implementation.
