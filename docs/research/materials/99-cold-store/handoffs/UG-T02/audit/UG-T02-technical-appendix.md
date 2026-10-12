# UG-T02 technical appendix

Internal artifact inventory and new local audit. Prepared September 24, 2026. Original files remain unchanged under `original_reference/`. The new audit is UG-T02-S98, not historical experimental evidence. Decision status is recorded in the main report, not inferred from executable files.

## 1. Material actually inspected

The uploaded `generalized_unit_reference.zip` contains six files. Its README and JSON are byte-identical to the separately supplied copies. The original artifact declares a research/draft date of September 23, 2026; this does not establish a message timestamp.

| File | Role | What its existence establishes |
| --- | --- | --- |
| README.md | Selection rationale and proposed contracts | An assistant-authored specification and explicit limitations, not owner acceptance. |
| generalized_wizard.unit.json | Parametric Wizard family | A written behavior organization, progression graph, price snapshot and unresolved binding declarations. |
| unit-draft.schema.json | JSON Schema 2020-12 document schema | Shape constraints on selected fields; not a full mechanics language or typed-binding validator. |
| check_reference.py | Python static reference checker | Implemented fixture computation; standard library core and optional jsonschema check. |
| resolved_builds.json | Saved ordinary configurations | 64 resolved structural plans, costs and unresolved requirements. |
| validation_report.json | Saved check results | Scope-limited results for the hashed input. |

Other inputs inspected are `GAME-DESIGN-RESEARCH.md` and `CONTEXT.md`, retained under `source_inputs/`. Neither is a copy of the unavailable product/API documents to which it links. No private Foundation contents, Towerright runtime, Unit-Generator repository, trained weights or anime source packet was supplied or inspected.

## 2. Representation preserved

The JSON has `schema_version: 0.1.0`, `artifact_kind: parametric_unit_family_draft`, and ID `layered_emitter_reference`. Important sections are `reference`, `status`, `definition`, `economy`, `presentation`, `interpretation`, `base`, `component_definitions`, `modifier_definitions`, `upgrades`, `composition_contract`, `lifecycle_proposal`, `paragon`, `binding_groups`, `capability_requirements`, `known_coverage_gaps`, `export_gate`, `provenance` and `scope_limits`.

This is a proposed standalone document, not demonstrated compatibility with the repository's RulePack API. The supplied CONTEXT.md warns that the experimental RulePack is a layout descriptor, not a complete executable Game Definition package.

### Ordinary progression

Three integer tiers each range from zero to five. Sorted descending, the second must be at most two and the third must be zero. The 64 builds comprise one unupgraded state, fifteen single-path states and 48 two-path states. The checker rejects ordinary `3-3-0`, `2-2-2` and `5-5-5` under this rule. A Paragon is a separate purchase/state transition.

The fixed fixture also assumes a single owner and one tier-five instance per family/path. Cooperative ownership, unlock progression, Monkey Knowledge, discounts and mode restrictions are excluded. These extra scope choices are assistant proposals, not separately confirmed owner requirements.

### Upgrade identity and price snapshot

All costs are incremental unmodified Medium-reference values inside the original artifact. They are not asserted to be the latest game values. `exact_game_build` is null. No game-data export or immutable source revision was supplied.

| Node | Reference identity | Generalized intent | Incremental price |
| --- | --- | --- | ---: |
| A1 | Guided Magic | Guided delivery | 175 |
| A2 | Arcane Blast | Stronger impact | 450 |
| A3 | Arcane Mastery | Sustained specialization | 1,450 |
| A4 | Arcane Spike | Heavy-target specialization | 10,000 |
| A5 | Archmage | Integrated specialist | 32,000 |
| B1 | Fireball | Impact burst | 300 |
| B2 | Wall of Fire | Persistent zone | 800 |
| B3 | Dragon's Breath | Sustained discharge | 3,300 |
| B4 | Summon Phoenix | Temporary companion | 6,000 |
| B5 | Wizard Lord Phoenix | Persistent companion and form | 50,000 |
| C1 | Intense Magic | Extended hit capacity | 300 |
| C2 | Monkey Sense | Expanded perception | 300 |
| C3 | Shimmer | Shared reveal | 1,500 |
| C4 | Necromancer: Unpopped Army | Event-funded attackers | 2,800 |
| C5 | Prince of Darkness | Expanded attackers and command | 26,500 |

Base purchase is 250. Paragon upgrade is 750,000. The assistant also proposed `generated price = reference price × scale factor`; the written fixture sets the factor to one. A common scale preserves ratios before rounding, not gameplay balance.

| Build | Total fixture price |
| --- | ---: |
| 5-2-0 | 45,425 |
| 5-0-2 | 44,925 |
| 2-5-0 | 61,275 |
| 0-5-2 | 61,250 |
| 2-0-5 | 32,275 |
| 0-2-5 | 32,750 |
| 5-0-0 | 44,325 |
| 0-5-0 | 60,650 |
| 0-0-5 | 31,650 |

The last three builds plus the Paragon price total 886,625. This excludes additional sacrifices and says nothing about degree, performance or price parity with a current installation.

### Behavior distinctions

There are twelve ordinary component definitions: `primary`, `payload`, `zone`, `stream`, `reveal`, `temporary_companion`, `permanent_companion`, `alternate_form`, `harvest`, `track_standard`, `track_heavy` and `command`.

The valuable distinctions are independent attack clocks; target access versus damage immunity; acquisition range versus projectile travel; carrier collision versus area impact; location-owned zones versus target-owned statuses; independently targeting companions versus route-following attack proxies; and the persistence of attacks/resources during transformation.

Seven modifier definitions express A1 guidance, A2 impact, C1 hit capacity, C2 perception, C3 collection range, C4 pool scaling and C5 pool scaling. Candidate targets are recorded, but exact propagation masks require bindings. A5 adds stream/reveal behavior without purchasing B3/C3. B5 retires the temporary companion when it introduces the persistent companion and form. C5 supersedes the earlier pool-scaling modifier. These are written structural choices; they are not executed combat interactions.

### Paragon contract

The generalized Paragon has Charge, Spend, Release and Awakened states or actions. Its contract distinguishes gathering resource, consuming resource, releasing a recorded pool once, temporarily replacing emissions, companion recovery and restoring permitted state. Eligibility requires three distinct same-family tier-five specializations within the chosen owner scope and enough cash. The written transaction also describes sacrifice evidence and cleanup.

Some of these mechanics are inspired by the cited Magus reference; exact event guards, cleanup, persistence, ownership and state migration are proposed generalized contracts. The Python eligibility function tests selected conditions. It does not execute the complete purchase transaction, all temporal transitions or degree calculation.

### Unresolved bindings

There are 48 binding groups. Each declares a purpose, required field names/types, values and required evidence. A count of groups is not a count of scalar constants. Missing material includes numeric attacks, target masks, modifier masks, world/event policy, resource accounting and degree-dependent profiles.

A binding declaration is not proof that a Consumer implements it. A group marked `bound` is not automatically complete: the present schema checks document shape rather than validating every declared field/value relation.

## 3. Reproduction and negative probes

Run `python audit/reproduce_audit.py` from the extracted handover package. It copies the original reference to `audit/reproduced/`, runs the original checker, compares outputs and executes two new probes. It does not contact a model provider or game runtime. Python's jsonschema package is optional for the original structural checker and needed for the additional null-binding schema probe. If absent, that probe reports not run.

The audit in this handover used Python 3.13.5. Platform and complete result details are in `audit/audit_findings.json` and `audit/reproduced/audit_reproduction.json`.

Original JSON SHA-256:

`bd0af24ee9ebe12235345d06aec4d1e3a4e88de84aa53432fe288751d54e4946`

### Reproduced results: UG-T02-E02

The original checker reports 136 successful structural checks, 64 resolved ordinary builds and 424 prerequisite-preserving interleavings. The rerun's `resolved_builds.json` and `validation_report.json` are byte-identical to the original saved files in this environment. The report still declares 48 unbound groups, executable handoff false, Consumer execution not evaluated, balance not evaluated and Acceptance not requested.

This does not establish code coverage, semantic completeness, game compatibility, human fidelity or balance. Many checks are fixture-specific assertions rather than independent general properties.

### New probe 1: canonical composition is not live sequence validation

`resolve(spec, (2, 0, 0), ['A2', 'A1'])` returns a structural plan rather than rejecting the reversed prerequisite order. The function first checks that the supplied nodes match the expected set exactly once, then sorts them by tier and path. It does not apply them as the supplied sequence of live purchases.

The `purchase_interleavings` generator itself emits only prerequisite-preserving orders. Equality over its 424 generated orders therefore checks a canonical static plan, not preservation of cooldowns, resource pools, prior emissions or purchase-time events. This limitation does not invalidate its advertised static use; it limits how the test count should be described.

### New probe 2: bound status does not establish valid values

A copy of the fixture with every binding group marked `bound`, while all values remain null, passes the document-shape schema. The schema explicitly does not claim full binding validation. The resolver's `executable_handoff_ready` field is hardcoded false, so it is also not a functioning export-readiness computation that becomes true after all bindings are supplied.

These probes are new handover evidence, not failures previously observed by a user, and not new requirements automatically accepted into the project. They should become regression tests only if the corresponding stronger guarantees are adopted.

## 4. Training and evaluation proposals not implemented

The assistant suggested a pretrained code-capable generator, LoRA/QLoRA supervised adaptation, reviewed paired source-to-unit transformations, optional completion/formalization/repair tasks, and later preference training or distillation. It mentioned Qwen3-14B and Qwen3-8B as candidate comparison points. There is no selected model revision, provider, quantization, context limit, training dataset, training script or checkpoint in this package.

The proposed 200-500 training records and 50-100 evaluation requests were planning figures, not measured sufficiency or approved study sizes. Exact crosspath outputs are related states of a family, not independent adaptation examples. Copies or paraphrases of a character should not inflate evidence of generalization. These are recorded evaluation cautions, not executed dataset splits.

The intended comparison included direct LLM generation, a structured generation/checking path and the same path after training. No scores, human participants, source-fidelity measurements, cost measurements, gameplay trials or successful anime-to-unit transformations exist in this thread. The public versus private performance split therefore has no experimental result to report.

## 5. Disclosure and reuse

This is an internal handover. The inclusion of files and references does not authorize GitHub or paper publication. No license was selected for a generator, model, BTD6-derived data, prompts, examples or adapted characters. The supplied CONTEXT.md explicitly names Foundation as private guidance, but that material is not present. No Towerright-private tuned checkpoint, prompt pack or dataset was available to classify. Any public release needs a separate artifact-level boundary decision.
