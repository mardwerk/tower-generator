# Issue labels and milestones

Use this catalogue when creating issues, changing labels or assigning milestones in Tower Generator.
It contains the contributor rules needed without access to private Planning.
Labels describe recorded work and decisions. They do not approve a specification or select work for the owner.

## State

Every open issue has exactly one open state.
Choose it from the latest recorded owner decision and record the reason for a change.
Do not infer selection, acceptance, pause or resumption from inactivity, milestones, model suggestions or passing tests.

| Label | Color | Meaning |
| --- | --- | --- |
| `status:active` | `0E8A16` | Owner-selected current work or decision |
| `status:triage` | `D4C5F9` | Unresolved scope, ownership or scheduling; no recorded selection |
| `status:deferred` | `FBCA04` | Retained agreed later work or pending prerequisite |
| `status:parked` | `F9A825` | Future option awaiting owner selection; no explicit exclusion |
| `status:on-hold` | `BFDADC` | Explicit owner pause; reason and resumption conditions |
| `status:out-of-scope` | `F9A825` | Explicit exclusion from named product scope; retained open proposal |

An excluded open proposal has only `status:out-of-scope` as its state.
A missing milestone does not establish exclusion.
An owner pause needs a recorded reason and conditions for resumption.

Use GitHub's native closure when the owning completion rules are met.
Remove open states on closure and retain useful descriptive labels and decision history.
Normal completion needs no terminal label. At most one of these terminal reasons may describe a closed issue.

| Label | Color | Meaning |
| --- | --- | --- |
| `status:invalid` | `B60205` | Incorrect or non-actionable report; recorded evidence and reason |
| `status:duplicate` | `CFD3D7` | Same work owned elsewhere; recorded owning issue |
| `status:declined` | `6A737D` | Explicit owner rejection; recorded decision |

Labels alone never justify closure.
When reopening an issue, remove its terminal reason and assign one current open state.

## Topics and origins

Give each open issue one primary topic for its main responsibility.
Use the most specific documented match and split independent tasks into separate issues.
A triage exception may temporarily leave the topic missing or unresolved, with the classification question recorded in the issue.
Closed issues may retain multiple descriptive topics for their historical scope.

Classify clear content from its actual responsibility, including the body and recorded decisions.
When more than one primary topic remains plausible, record the alternatives and seek owner guidance before changing the assignment.
Topic classification does not change state, importance or specification acceptance.

| Label | Color | Meaning |
| --- | --- | --- |
| `topic:repository` | `7057FF` | Repository organization, contributor rules, tracking and governance |
| `topic:operations` | `7057FF` | Builds, test environments, publishing, deployment and operational data |
| `topic:interface` | `7057FF` | User controls, presentation and interaction |
| `topic:cli` | `7057FF` | CLI commands, flags and serve API |
| `topic:engine` | `7057FF` | Generator schemas, mechanics, checks and prompts; Profile importer and static validator |
| `topic:unitlab` | `7057FF` | Local UnitLab client, server, settings and library |
| `topic:security` | `7057FF` | Security boundaries, credentials and access protection |
| `origin:manga-mayhem` | `C5DEF5` | Explicit Manga Mayhem product or source relationship |

`topic:engine` covers both retained generator work and the current Profile Validator work.
Its scope is wider than the validator. Do not add a separate validator topic for the same responsibility.
`topic:unitlab` includes server work as well as the client.
Use `topic:interface` for interaction or presentation as the primary responsibility, and `topic:unitlab` for the local application's broader workflows and services.
Use `topic:security` when the main responsibility is a security boundary, even if the affected code is in the CLI or Lab.

Origins are optional here and require an explicit relationship to another product or source repository.
Use `origin:manga-mayhem` for work that records that relationship, such as support for a requested game mechanic.
Do not add a Tower Generator self-origin or infer an origin from generic game terminology.
Multiple origins require multiple stated product relationships.

## Type, review and importance

Use at most one change type when it helps explain the work.
A documentation change concerning the engine uses `topic:engine` with `type:documentation`.

| Label | Color | Meaning |
| --- | --- | --- |
| `type:bug` | `D73A4A` | Existing behavior failing an accepted requirement |
| `type:feature` | `1D76DB` | New or changed product behavior |
| `type:documentation` | `0075CA` | Documentation or contributor instructions |
| `type:maintenance` | `0366D6` | Code, dependency or repository upkeep; no primary feature or bug |
| `review:decision` | `FBCA04` | Pending owner review; named question, options or acceptance evidence |
| `review:audit` | `F9A825` | Fresh owner audit of existing decision or artifact; no reaffirmation |
| `importance:high` | `D93F0B` | Explicitly recorded high owner importance |

Review and importance are independent of state.
An active owner decision uses `status:active` and `review:decision`.
Deferred, parked and excluded proposals may also need a named pending review.
Use both review labels only when the issue records two distinct pending tasks.
After a review resolves, remove its review label and record any acceptance or state change separately.
A completed owner review must not become pending again merely because labels change.

Apply `importance:high` only from an explicit owner decision.
No importance label means no special importance is recorded.
Keep initial uncertainty in triage and named pending decisions in Review; there is no `question` label.

## Milestones and views

A milestone describes one repository outcome with a short name, included scope, observable completion criteria and a coordination issue.
Use dates only with owner agreement.
Assign at most one milestone per issue and only when the work directly contributes to that outcome.
A milestone assignment does not activate work or accept a contract.

Keep one current outcome for this product and create later milestones only when their scope is concrete.
Follow Tower Generator's recorded scope decisions and keep milestone scope in this repository.
Close a milestone when its criteria and required work are complete, not because its closed percentage is high.
Review deferred or cancelled items individually before moving them to a later milestone or removing their assignment.

Start with saved or bookmarked GitHub searches.
Use `is:issue is:open label:"status:active"` for current work and substitute another state for triage or later work.
Use `is:issue is:open label:"review:decision"` for pending owner decisions and `review:audit` for audits.
Narrow a search with topic, type, origin or milestone filters when useful.
Labels remain authoritative for state and importance. Any mirrored Project fields need a synchronization owner and rule.
Moving a card or changing a view does not make a scope or acceptance decision.

## Migration and extensions

Install this shared base and the documented local topics and origin before migrating assignments.
Use exact names and meanings from the tables, with six-digit hexadecimal colors and no `#` in API values.
Keep GitHub descriptions consistent with these meanings.
The same name and meaning must use the same color across repositories.

Preserve recorded scope, owner decisions, reviews, importance, acceptance and historical information.
A label migration does not close, reject, accept, pause or resume work.
Inspect closed issues too, remove open states and preserve exclusion history in the recorded decision.
Apply terminal reasons only when the evidence supports them.

| Legacy label | Mapping or adaptation |
| --- | --- |
| `bug`, `enhancement`, `documentation` | `type:bug`, `type:feature`, `type:documentation`; select at most one supported change type |
| `cli`, `engine`, `unitlab`, `security` | Corresponding `topic:*` with the same responsibility |
| `ux` | `topic:interface` |
| `deferred`, `triage`, `on hold` | Corresponding `status:*` from recorded decisions; later work is distinct from an owner pause |
| `out-of-scope` | On an open excluded proposal, one `status:out-of-scope` replaces other open states |
| `needs review`, `audit: needs review` | `review:decision`, `review:audit`; preserve the named pending tasks |
| `invalid`, `duplicate`, `wontfix` | On closed issues, the evidenced `status:invalid`, `status:duplicate` or `status:declined` reason |
| `question` | Remove the label; use triage for unresolved initial scope or Review for a named pending decision |
| `idea` | Retain the rough-idea information in the body and classify state independently |
| `help wanted`, `good first issue` | Retain useful contribution instructions in the body and retire the labels |
| `accessibility` | Choose the actual responsibility; add a dedicated topic only for a recurring filter need |

If legacy labels cannot map without interpreting or changing an owner decision, record the conflict and propose an adaptation to the owner.
Keep the affected decision intact until resolved; continue independent migrations.
Update filters and references, then retire legacy aliases after their assignments migrate.

Prefer an existing label before adding a recurring local topic or origin filter.
Owner review is required for additions or changed meanings. Update this catalogue and GitHub descriptions together.
A local addition needs its exact name, meaning, color, examples and valid combinations.
Use `7057FF` for local topics and `C5DEF5` for origins.
New categories, states or importance meanings require a shared-standard decision, not a competing local lifecycle.

## Examples

These examples describe combinations supported by recorded scope decisions.
They do not accept a pending specification or implementation.

| Work | Labels and interpretation |
| --- | --- |
| [Profile Validator #85](https://github.com/mardwerk/tower-generator/issues/85) | `status:active`, `topic:engine`; preserve any separately recorded pending review |
| [Validator requirements #92](https://github.com/mardwerk/tower-generator/issues/92) | `status:active`, `topic:engine`, `type:documentation`; documentation can concern the engine |
| [Future Lab #86](https://github.com/mardwerk/tower-generator/issues/86) | `status:deferred`, `topic:unitlab`; both client and server remain in this responsibility |
| Retained Manga Mayhem mechanic request | `status:deferred`, `topic:engine`, `origin:manga-mayhem` when those later-work and product relationships are recorded |

[Larger progression testing #4](https://github.com/mardwerk/tower-generator/issues/4) validated an ambiguous topic during the 2026-09-30 migration.
It retains both engine validation and future Lab presentation checks.
Kyle selected `topic:engine` as its primary topic and retained all Lab requirements in the body.
Renaming its two legacy topics alone could not settle that choice.

For [the future executable name #89](https://github.com/mardwerk/tower-generator/issues/89), Kyle selected `status:deferred` as the single state.
The unresolved final name and implementation details remain in the body.
This preserves the later-work decision without using a second state for the remaining questions.

Before adopting automatic classification, validate an ambiguous example against the issue body and the owner's decisions.
Record the alternatives, the chosen responsibility or unresolved question, and the reason in the coordination issue.
Do not adopt automation that silently settles ambiguous topics or changes owner decisions.
