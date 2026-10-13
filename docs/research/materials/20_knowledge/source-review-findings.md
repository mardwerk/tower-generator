# Generator research: source review findings

These observations come from the September 26, 2026 source review. Metadata-only checks did not establish paper findings; reproduced software examples establish only the stated version and scenario. No present-day check or acceptance is implied.

These are dated source observations and proposals. Preserve their recorded review and rejection status. Current local contracts govern acceptance; these notes neither select implementation nor resume paused development.

## Draft corpus review, 2026-09-26

| Source group | Summary | Use it for | Do not use it for |
| --- | --- | --- | --- |
| Agentic generation research | [agentic-generation-research.md](agentic-generation-research.md) | Open Unit Generator proposals and what each depends on; bibliographic status of 45 papers; licenses of cited code | Evidence that any generation approach works for Unit Generator |
| Game and character design | [game-and-character-design.md](game-and-character-design.md) | Unit authoring guidance; BTD6 examples checked against v56.0 game data | Psychology, balance or studio-process claims from the German report |
| BTD6 snapshot | [btd6-snapshot.md](btd6-snapshot.md) | A checked record of the historical snapshot: its prices and local upgrade rule agree with btd6-atlas, apart from one naming error | A BTD6 source (use btd6-atlas), combat stats, crosspath effects, world rules, the proposed kernel as a requirement |

### Findings across groups

- **Product state has moved on.** The drafts describe unit-generator as it stood on 2026-09-22 to 2026-09-24.
  Since then the product was ported to Go (the inspected revision, 2026-09-26), Zod was removed, and most linked docs were deleted or rewritten.
  Its current architecture runs one fixed generation route with bounded repair.
  The drafts' premise that no implementation exists is therefore stale, and their workspace-relative links are recorded as stale.
  Where they disagree, the product repository wins.
  A running route, passing checks or replay parity is not evidence of good unit design: the owner reports a quality regression whose cause is not isolated, recorded in the [Tower Generator workstream](../30_projects/unit-generator/README.md#earlier-generation-work).


- **BTD6 facts.** The BTD6 source is [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas), which unit-generator also names as its authority; the capture manifest has not been accepted yet.
  The frozen snapshot in `drafts/90-knowledge/btd6-snapshot/` agrees with the atlas and two other exports on every price.
  The snapshot has one naming error: it swaps the names of Heli Pilot bottom tiers 1 and 2. btd6-atlas is correct: at the inspected revision, 001 is Faster Darts (250) and 002 is Faster Firing (350).
  The BTD6 examples in the game design synthesis match the v56.0 game data.


- **Citations.** No fabricated identifier was found.
  The two reading-list groups cite 48 and 45 papers, and many appear in both.
  All were checked bibliographically: two titles are wrong and one venue is out of date.
  No paper's attributed findings could be read, because arXiv and publisher hosts were blocked.
  Treat every paper finding in the drafts as unchecked.


- **Licensing flags for reuse.** Cyber Quincy is AGPL-3.0.
  The GAVEL repository has no root license and bundles a Ludii fork under CC BY-NC-ND 4.0.
  The Hunyuan3D 2.1 license excludes the European Union, the United Kingdom and South Korea.


  Cyber Quincy is no longer a BTD6 reference, because btd6-atlas is.
  GAVEL is only a research reference, and Hunyuan3D 2.1 is not used.
  Revisit a flag if its code or data is reused or its tool is adopted.
