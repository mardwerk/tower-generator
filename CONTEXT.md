# Unit Generator context

Terms and ownership used across this repository. [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) describes how the pieces fit together.

Each term has one meaning here. Use it only with that meaning and do not substitute a synonym. Other Mardwerk documents use some of these words more broadly, so outside this repository name the scope, such as "Unit Generator Result".

## Terms

| Term | Meaning |
| --- | --- |
| Tool | Unit Generator as a whole: the Engine plus its CLI and `serve` API. It keeps no hidden state between calls. |
| Engine | The code that prepares, drafts, checks and reviews units. It does not run combat; a Consumer does. |
| Definition | The rules of one game that the Engine can check: paths, tiers, legal purchases, supported mechanics, currency and scale. |
| Profile | A reusable, editable file that a user selects for generation. It contains a Definition with the values it allows, plus rules text and the task. The bundled default is read-only; saved Profiles are copies in the Profiles folder (`data/profiles` by default). A game's own Profile ships with that game; a copy saved here is a local copy. |
| Definition profile | The `profile` field inside a Definition: currency, cost, stat and change limits, reference scale and the optional design policy. It is not a Profile; write "Definition profile" or `profile` in full. |
| Sources | The saved result of researching a character: its identity, retrieved source documents and source images. It needs no model call and can be prepared under any Profile. |
| Request | The complete explicit input for one unit: character, source documents, Profile content, confirmed decisions and, for a revision, the previous version and feedback. |
| Prepared request | A validated Request with its input hash. Editing the Request invalidates the hash. |
| Draft | A proposed unit produced by a model from a prepared request. |
| Checked artifact | A draft plus deterministic findings. |
| Result | A checked artifact plus a model review. |
| Artifact | Any of the above saved as versioned JSON. |
| Finding | One recorded issue or observation: what it concerns, how it was established (deterministic or model), severity and outcome (`pass`, `fail`, `unresolved`, `not_checked`). |
| Authoring check | A typed mechanics check that judges how a unit is made, not whether its blueprint can be read: change budgets, early attack identity and capabilities, purchases with no effect or only drawbacks, single changes with no effect, and the design policy. Drafting, repair, `check` and review run them. Reading an artifact (`render`, `view`, `build`) checks only its structure, so an authoring check added later cannot hide a saved Unit sheet. The view (`authoringIssues`) and `render --details` list the authoring checks that fail today as current checks, apart from the stored Findings. |
| Candidate | The readable unit, with its blueprint when mechanics are typed. Drafts, checked artifacts and Results each carry one. |
| Unit sheet | The rendered unit: its name, `0-0-0`, each purchase by build code with exact numbers (labeled apart from the plan's intent for it, when the plan types its technique) and its Proposed mechanics, labeled as not yet supported, every crosspath and, for a revision, patch notes. Findings, usage and provenance are rendered apart from it (`render --details`). |
| Patch notes | The changes a revision made: changed mechanics and the builds they affect, apart from renamed purchases. |
| Run | The record of one model stage inside a Draft or Result: model, timing and usage. |
| Active Ability | An ability that the player activates by clicking it. In a blueprint, owned `abilities` are Active Abilities; automatic effects are modeled as attack stats and statuses, not as abilities. The design policy calls an Active Ability a manual ability (`manualAbilityPath`, `maxManualAbilityPaths`). Those field names stay, and they refer only to Active Abilities. |
| Source techniques | The techniques and forms a prepared request's selected source passages name, found by fixed rules (power section headings, leading terms of entries, technique page titles, signature sentences), each with its passage IDs (`sourceTechniques`). `prepare` derives them; they are evidence for the plan, not a Repertoire. A request that has them splits its sources into passages labeled with their section. Under the design policy's `requireCoreConcepts`, the plan lists each in its Repertoire or in `omittedTechniques` by its name; a Repertoire entry named for a form also lists a technique that the form's own section names as a part of it, when the entry cites that passage. |
| Repertoire | The character's source-backed techniques and forms that a design plan selects (`repertoire`). Each planned purchase names the repertoire technique or base attack it adapts. Every path has a technique of its own; two paths share one only as crosspath synergy. Under `requireCoreConcepts` each entry, and each technique the plan omits, carries its `importance`: `core`, `major` or `minor`. |
| Core concept | A Repertoire entry ranked `core`: one of the one to three concepts without which the character would not feel canonical, judged from the sources. Under the design policy's `requireCoreConcepts` it is the base attack or the technique of at least one purchase that adapts its central effect with a typed change, and it is never omitted whole, only in named aspects. One adapted only by a Proposed mechanic is a Design gap: no build grants a proposal, so the Unit embodies the concept only once the Definition supports that mechanic. A purchase that carries its name but not what it does in play does not adapt it; the review judges that. |
| Build code | A purchase or build written top-middle-bottom. `0-0-0` is the base unit; `x-4-x` is the middle path's fourth purchase, where `x` means unspecified; `1-2-0` is a concrete build. |
| Crosspath | A legal build that buys two paths. Under the default Definition there are 12 early crosspaths, with both paths at their first or second purchase, and 36 advanced ones, with one path further. |
| Early benefits | What a path's first and second purchases improve or unlock, taken together: the multiset of their improvement dimensions and unlocks. Order is ignored, a dimension both purchases improve counts twice, and `lowers`, names, prices and amounts are left out. Under the design policy's `distinctEarlyBenefits`, no two paths may have the same early benefits. Under `exclusiveEarlyBenefits`, which subsumes it, no two paths' early benefits may share a dimension or unlock: an unlock of a dimension a purchase can also improve, such as splash, counts as that dimension, and a targeting change is left out. |
| Path identity | What a path's third purchase adds that makes the path its own. It can be simple, such as the only purchase that adds projectiles, or a new form, and a dimension another path's purchase also raises in passing does not take it away. The review judges it from the rules document's purchase roles; code checks no Path identity, and `exclusiveEarlyBenefits` is the typed floor under it. |
| Mechanic proposal | A suggested addition to a Definition. It is not an approved rule and does not mean the Engine supports it. |
| Proposed mechanic | A Mechanic proposal that one purchase needs, kept on that purchase (`proposedMechanics: [{name, effect, sourceIds}]` on its tier) beside the typed changes that approximate it: `effect` says what it does in play and `sourceIds` cite the passages that describe it. No build grants it and build resolution ignores it. Each is reported as an unresolved `proposed-mechanic` Finding, flagged for Definition expansion, and the review fails one that does not fit its purchase and sources or is not a playable mechanic. A plan milestone may name one, and its purchase keeps it. It can be a third purchase's Path identity, but it does not meet `requireTier3BehaviorChange` or `requireTier5BehaviorChange`: a purchase whose only new capability is proposed is a Design gap. |
| Design gap | Something the design policy requires that a Unit has only as a Proposed mechanic: a third or fifth purchase that a behavior rule (`requireTier3BehaviorChange`, `requireTier5BehaviorChange`) covers, that adds no supported behavior or access and whose only new capability is a Proposed mechanic, or, under `requireCoreConcepts`, a Core concept that the purchases whose technique it is adapt only with Proposed mechanics. It is not playable until the Definition supports the mechanic. It does not fail generation: it is reported as an unresolved Finding, `proposed-capability` for a purchase and `core-concept` for a Core concept. A purchase or Core concept with neither a supported adaptation nor a Proposed mechanic fails its rule. |
| Unsupported mechanic | Behavior a unit's sources call for that its Definition cannot express. When a purchase needs it, it is a Proposed mechanic of that purchase; otherwise it is recorded as a Mechanic proposal in the blueprint's `proposals` and reported as an `unsupported-mechanic` Finding. No build grants it. |
| Bonus damage | A typed, additive +N that an attack adds to each hit on enemies with one enemy property, such as +50 against Hardened. Purchases only add to it: two paths' additions sum, and damage multipliers do not scale it. A Definition's vocabulary lists the properties that accept it (`bonusDamageProperties`); the Default Profile lists Hardened, its adaptation of BTD6 Ceramic, and Blimp, its adaptation of BTD6 MOAB-class, as separate properties: a Blimp is not Hardened, and each property's bonus is its own. A damage type that cannot hurt a property blocks its bonus damage too. BTD6's class bonuses, such as Deadly Precision's +50 to Ceramic and MOAB Press's +4 to MOAB-class, are what it adapts. |
| Library | The folder, chosen by the user (`data/runs/library` by default), where saved Sources, artifacts with their Markdown, icons and portraits live, arranged by work and character. The CLI and `serve` write it; nothing else is stored. |
| Consumer | The game or runtime that executes generated units. |
| Towerright | The separate project that owns project history, multi-tool orchestration, wider evaluation and acceptance. |

## Principles

- A check that fails is a finding, not an error; the operation still completed.
- Checks establish that stated rules are followed, not balance, fun or acceptance. Model reviews are judgments, labeled as such.
- Confirmed decisions, proposals and open questions stay distinct in every artifact.
- Source text is evidence for a character; it does not authorize a game mechanic. Missing behavior becomes a finding or a mechanic proposal, never an invented rule.
- Reloading an artifact uses the rules retained inside it, never today's defaults.

## Ownership

| Owner | Owns | Does not own |
| --- | --- | --- |
| Unit Generator | Unit research, preparation, generation, checks, review, rendering, its library and Profiles folders | Project history, orchestration across tools, runtime behavior |
| Web client | Screens, the user's stage-by-stage orchestration, unsaved session state | Rules, validation, model calls |
| Towerright | Project history, orchestration, evaluation, acceptance | Unit rules |
| Consumer | Runtime execution and its evidence | Generation |
