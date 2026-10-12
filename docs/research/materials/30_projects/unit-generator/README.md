# Towerwerk

Status: On hold. Standalone generator; no Manga Mayhem prerequisite.

The [product rules](https://github.com/mardwerk/towerwerk/blob/main/docs/PRODUCT.md#established-direction) own Tower Generator's selected scope and delivery order.
Agent Skill and proper generator remain two modes, with Skill before generator.
Towerwerk is standalone and its development is on hold; Manga Mayhem is not a prerequisite.
The [SkillOpt plan](https://github.com/mardwerk/towerwerk/blob/main/docs/research/SKILLOPT.md) owns optimization, Kyle's review and the remaining evaluation proposal.
Supported requests, output details and quality criteria remain subject to the product review and game evidence.
The [existing skill](https://github.com/mardwerk/towerwerk/blob/main/.agents/skills/tower-generator/SKILL.md) remains a seed, not a completed implementation of the new scope.

The old Python application, Node web bridge, `src/web`, Lab, Default Profile, Usopp experiment and broken setup have been removed in [PR #120](https://github.com/mardwerk/towerwerk/pull/120).
Git history preserves the removed implementation; independent research evidence and helpers remain.
Existing UI preferences inform future client discussion and do not require retaining the old Lab files.
The later Go backend awaits the [workflow acceptance review](https://github.com/mardwerk/towerwerk/blob/main/docs/PRODUCT.md#workflow-acceptance-review) in [#106](https://github.com/mardwerk/towerwerk/issues/106).

Kyle selected a focused transfer from `unit-design-research` on 2026-10-03 and an initial Tower Generator skill for later SkillOpt optimization.
The [owning transfer issue](https://github.com/mardwerk/towerwerk/issues/108) covers preservation, the draft skill and its evaluation plan.
Keep current research guidance beside Tower Generator and collect more evidence only when a concrete design needs it.
Kyle subsequently consolidated the retained documents under [Tower Generator research](https://github.com/mardwerk/towerwerk/tree/main/docs/research).
The snapshot wrapper, manifest and Gojo code/results are retired and remain in Git history.
Planning retains the [transfer reference](../../99-cold-store/unit-design-research-2026-10-03/README.md).
The retained research proposals and historical Gojo experiment are reference material, not accepted product dependencies.

The [2026-10-03 review](2026-10-03-review.md) records the plan check, readiness evidence and local cleanup.
[Planning #104](https://github.com/mardwerk/planning/issues/104) records the completed transfer and workspace coordination; product issues own current status.
The next useful generation contract should connect a saved Wiki revision to an editable Tower design and supported validated records, with explicit behavior for unsupported mechanics.
It remains a proposal for the product review, rather than an accepted interface.

<a id="pending-workflow-choices"></a>

## Later generator review

The product's [API, storage and workflow rules](https://github.com/mardwerk/towerwerk/blob/main/docs/PRODUCT.md) own selected backend decisions and remaining choices.
The [realignment record](https://github.com/mardwerk/towerwerk/blob/main/docs/REALIGNMENT.md) preserves reset history and the discarded recovery experiment's outcome.
These later contracts do not require implementing the backend before the initial Agent Skill.
The product repository records readiness and acceptance; this page does not establish an additional implementation contract.

## Profile ownership

Tower Generator retired the old Default Profile and Usopp experiment.
The [authoring editor outline](https://github.com/mardwerk/towerwerk/blob/main/docs/PRODUCT.md#selected-authoring-editor-boundary) and [Profile rules](https://github.com/mardwerk/towerwerk/blob/main/docs/PROFILE.md) belong to Towerwerk. A potential editor uses explicitly supplied data; integrations are separate proposals and do not establish a game prerequisite.
Product authoring choices and checks remain Tower Generator's responsibility.
[td-profile](https://github.com/mardwerk/td-profile) owns the generic Profile Validator, reusable schemas and copyable Profile Template.
[btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas) owns the captured BTD6 facts and its game-specific Profile.
The retired setup and authoring experiment do not establish dependencies for the new research backend.

A Profile is a directory selected by path.
Manga Mayhem's game-owned Profile remains separate and deferred; [#92](https://github.com/mardwerk/manga-mayhem/issues/92) tracks that work.
The [MangaMayhem Unit Profile](../mangamayhem-unit-profile/README.md) workstream preserves its drafts without establishing game acceptance.

## Earlier work

The September validator milestone and generation pause no longer describe the current stopping point.
The [Luffy quality finding](https://github.com/mardwerk/towerwerk/issues/22#issuecomment-5847053425), [Default Profile scale concerns](https://github.com/mardwerk/towerwerk/issues/21#issuecomment-5847055300) and [quality review](https://github.com/mardwerk/towerwerk/issues/61) remain historical product evidence.
The cause of the reported quality regression was not isolated in that review.
The earlier [unit design research](UNIT-DESIGN-RESEARCH.md) and [bounded optimization proposal](AUTORESEARCH.md) remain historical references, rather than implemented workflows.

Accepted contracts, implementation and reproducible evidence stay in the owning product repository.
Put incoming unsorted notes and files in [Planning's inbox](../../00_inbox/README.md).
