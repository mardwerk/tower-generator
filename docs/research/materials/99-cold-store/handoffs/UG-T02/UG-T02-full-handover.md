# UG-T02 handover: character adaptation, BTD6 progression and the Wizard reference

Internal consolidation record. Prepared September 24, 2026. Not approved for public release. Current product context is recorded separately as UG-T02-S00. Source IDs resolve in UG-T02-source-register.md and its JSON companion. Artifact locations and new checks are documented in UG-T02-technical-appendix.md.

## 1. Thread overview and coverage

**Thread ID:** UG-T02. **Descriptive title:** Character-to-unit adaptation under a BTD6-derived progression contract.

The original question was whether a model could learn to transform an anime character description or wiki entry into a full tower defense unit valid under a supplied Game Definition. The user asked about missing research, pretraining, post-training, model types and whether model training was the right approach. The conversation then narrowed to a BTD6-like upgrade structure and a request to compare reference towers and write a generalized representation. [UG-T02-S01, UG-T02-S03, UG-T02-S05]

The contribution is a clearer separation between character interpretation, game constraints, mechanics representation and evaluation, plus an inspectable Wizard-derived fixture. It is not a trained generator or evidence that Unit-Generator beats a direct LLM prompt.

**Coverage.** Reviewed three original substantive user messages and three full assistant replies visible before this consolidation request, both uploaded Markdown documents, and the six files inside the generated reference ZIP. The separately uploaded README and JSON match the ZIP copies. Speaker plus distinctive quotation is used where no reliable message date or link is available. The README's September 23, 2026 date is an artifact date, not an inferred chat timestamp. [UG-T02-S01 to UG-T02-S09]

The research catalogue is a secondary summary, not the underlying papers. Its stated historical inspection dates and access statuses are retained separately from this handover's primary-source checks. Related-thread snippets supplied in project context were not treated as original UG-T02 messages or evidence of decisions here. Missing product/API documents, the older survey, reference indexes, Foundation, Towerright's own context and private repositories were not inspected. Original browsing logs and hidden reasoning were not available. The public sources revisited now can differ from earlier versions. [UG-T02-S07; UG-T02-S87 to UG-T02-S95]

**Central idea.** Keep the game's rules explicit, then adapt a character's recognizable behavior into a unit that follows those rules. The model would choose the adaptation rather than decide what the game permits. A BTD6-like profile constrains the output to a base unit, three five-tier paths, legal crosspaths and a separate Paragon. A structured representation can expose missing interactions and support repeatable checks. The Wizard example tests whether that representation can describe varied attacks, companions, resources and transformations. It does not yet show that a model can produce faithful, useful units, that the mechanics run correctly, or that reference prices preserve balance. Those remain separate questions.

## 2. Ideas, approaches, and changes of direction

**Train a specialist adaptation model.** The user described model training as the best case, not a selected implementation. The assistant recommended post-training a pretrained code-capable generative model, initially using supervised examples and LoRA, with QLoRA as a memory-saving option. It proposed Qwen3-14B and Qwen3-8B as comparison candidates. The intended benefit was more consistent source-to-mechanic mapping and less correction. The dependencies were a stable target representation, reviewed paired adaptations and held-out evaluation. No model, provider, hardware or training budget was selected, and no training occurred. [UG-T02-S01, "BEST CASE"; UG-T02-S02, "post-training"; UG-T02-D02]

**Use a bounded decision model.** Laya was considered because the user linked it as an example of a fast decision system. The assistant distinguished selection among supplied options from variable-length generation of a unit family. It retained a possible role for ranking interpretations, selecting mechanic families or routing cases, but did not recommend it as the main generator. This was an assistant assessment, not a user rejection of Laya. A repeated grammar-choice system remained possible but would need its own generation process and data. [UG-T02-S11; UG-T02-D02]

**Separate interpretation from formalization.** The assistant proposed recording signature behavior, its legal implementation and adaptation loss before producing detailed mechanics. This would make a faithful adaptation distinguishable from a themed reskin. It also proposed several different interpretations, followed by selection and precise repair. The cost is additional representation and review work. Separate records did not imply separate models or a mandatory multi-agent system. This remained an architectural proposal. [UG-T02-S02, "Learn an explicit Interpretation"; UG-T02-D05]

**Constrain a BTD6-like family rather than arbitrary game design.** The user proposed three paths, five upgrades per path, reference-relative costs, crosspath rules and a Paragon requiring the three tier-five specializations on the field plus cash. The assistant argued that this provides a clearer training target. It also exposed a larger interaction problem: the output is a whole family of legal builds, not fifteen unrelated ability descriptions. This became the working reference brief but not an approved permanent public default. [UG-T02-S03, "my idea would be"; UG-T02-S04; UG-T02-D04]

**Generalize behavior instead of renaming assets.** The user's projectile example led to separate triggering, targeting, delivery, collision, effects and presentation. The assistant extended this to location-bound zones versus target-bound statuses, autonomous attackers versus track-following proxies, and transformations with explicit replacement and persistence rules. This supports more faithful adaptation, but every additional behavior needs declared semantics and, for executable output, Consumer support. The thread did not authorize unlimited new game rules. [UG-T02-S03; UG-T02-S04; UG-T02-S09, The reusable representation]

**Choose a varied fixture before expanding the catalog.** The user asked to skip Dart Monkey and Super Monkey for this comparison and suggested Ninja, Druid and Wizard. The assistant selected Wizard first, Druid next, and described Ninja as smaller and Engineer as useful for ownership/support/trap rules. Beast Handler and Skywarden were additional comparison candidates. Only Wizard received a complete per-upgrade representation. The selection was qualitative, not a measured catalog-wide maximum of variety. [UG-T02-S05; UG-T02-S06; UG-T02-D03; UG-T02-D06]

**Train preferences, distill or search later.** DPO, numerical strength predictors, distillation, reinforcement learning, quality-diversity archives and reusable mechanic abstractions were possible later components. The recorded reasons for postponing them were missing reviewed preferences, runtime feedback, reliable measures or a demonstrated need. Pretraining a general language model from scratch was discouraged by the assistant. None of these alternatives was formally rejected or deferred by the owner. [UG-T02-S02; UG-T02-D09]

The consequential direction change was from asking which model to train to asking what a valid, expressive training target should look like. The thread ended with a representation and checker, not a return to model training. No evidence establishes that this change abandoned the training aspiration.

**Terminology requiring care.** CONTEXT.md defines a Unit as content describing a game actor or deployable combat element. The BTD6 discussion often uses "unit" for an entire family: base, fifteen upgrades, crosspaths and Paragon. A Mechanic is declared behavior; a Mechanic proposal is an unapproved addition to the Game Definition. Neither means an implemented runtime operator. "Normal mode" was the user's phrase; the assistant interpreted it as unmodified Medium prices without confirmation. "Tuning" covered model weights, profile configuration and numerical calibration, which are different activities. No historical operational definition of a "usable default" was established. The present owner context supplies that requirement but does not settle offline operation or pricing. [UG-T02-S08, Unit terms and lines 51-57; UG-T02-S00]

## 3. Decision record

The status below concerns approval. Evidence strength is stated separately. Only two entries reflect clear owner intent or authorization; the other entries preserve proposals.

| ID | Statement and status | Authority and locator | Rationale, evidence and reconsideration |
| --- | --- | --- | --- |
| UG-T02-D01 | **Accepted task objective:** investigate character-to-TD-unit adaptation under a supplied Game Definition. | User, UG-T02-S01: "we provide a 'Game Definition' in which the character has to be valid". | Explicit desired workflow, not evidence of feasibility or an approved executable deliverable. No reconsideration condition recorded. |
| UG-T02-D02 | **Proposed implementation:** post-train a pretrained generative model rather than pretrain a language model from scratch or use Laya as the primary generator. | Assistant, UG-T02-S02 and UG-T02-S04. User's "BEST CASE" establishes interest, not selection. | Paired adaptation data and external precedents were the rationale. No project experiment. Selection was to depend on measured correction effort, quality and cost. |
| UG-T02-D03 | **Accepted scoped task:** compare BTD6 reference units, omit Dart Monkey and Super Monkey from that comparison, and create a generalized representation. | User, UG-T02-S05: "Check them out and then create the generalized unit representation". | User described them as two extremes. This does not exclude those towers from all later work or approve the assistant's chosen tower. |
| UG-T02-D04 | **Proposed reference contract:** three paths by five upgrades, crosspath restrictions, relative prices and a separate Paragon requiring all three tier-five specializations plus cash. | User, UG-T02-S03: "my idea would be". | Bounds the adaptation target. Later artifact implements this brief, but permanent-default selection, exact world rules and reference version remain unapproved. |
| UG-T02-D05 | **Proposed architecture:** explicit Interpretation, structured mechanics, external scoped checks and precise repairs. | Assistant, UG-T02-S02: "Train the model to make good adaptations"; UG-T02-S04. | Separates source fidelity, legality and execution. Supported by related subtask research, not an end-to-end Unit-Generator result. No recorded owner acceptance. |
| UG-T02-D06 | **Proposed fixture selection:** Wizard first and Druid second. | Assistant, UG-T02-S06; UG-T02-S09 README, Decision. | Broad behavior coverage without starting with a cash-production system. Druid/Engineer could be preferable when economy/control/support receive more weight. No exhaustive comparison or user approval. |
| UG-T02-D07 | **Proposed fixture-specific semantics:** unmodified Medium prices, fixed price scale, single-owner roster/purchase assumptions and explicit Paragon lifecycle/cleanup contracts. | Assistant, UG-T02-S04 and UG-T02-S09. | Reduces ambiguity for the fixture. Some contracts are invented generalizations; exact BTD6 parity is not established. README explicitly calls for approval before these become rules. |
| UG-T02-D08 | **Proposed design freedom:** retain progression/economic constraints but allow character-driven behavior graphs, rather than forcing every character into Wizard's graph. | Assistant, UG-T02-S06, "Strict transformation" versus "Character-driven adaptation". | Prevents the first fixture from prescribing every character's mechanics. User did not resolve how much of "exact upgrade paths" is fixed. |
| UG-T02-D09 | **Proposed sequencing:** defer preference training, RL, distillation, automated diversity search and larger runtime work until their prerequisites exist. | Assistant, UG-T02-S02. | Requires reliable preferences, executable feedback or measured bottlenecks. This is advice to defer, not an owner-approved deferral. |
| UG-T02-D10 | **Proposed evidence plan:** reviewed paired transformations, held-out tests and cost-matched comparisons against direct LLM generation. | Assistant, UG-T02-S02; catalogue Proposed adaptations. | Needed to determine whether training or orchestration helps. No dataset, experiment or success threshold approved. |

There is no owner-approved rejection or supersession of the main alternatives in the visible exchange. Repeated discussion of training, Wizard or a formal representation must not be treated as approval. Current context also prevents reading UG-T02-D02 as a commitment that Towerright's "fine-tuned" offering must use trained weights, or reading UG-T02-D05 as a requirement that every user receive executable mechanics.

## 4. Technical artifacts, research claims, and results

### Artifact and implementation inventory

| Item | Location | Status and interpretation |
| --- | --- | --- |
| Generalized Wizard family | UG-T02-S09, generalized_wizard.unit.json | Written and inspectable. Schema 0.1.0, `parametric_unit_family_draft`, ID `layered_emitter_reference`. Contains 15 upgrades, 12 ordinary components, seven modifier definitions and 48 unbound binding groups. |
| Document schema | UG-T02-S09, unit-draft.schema.json | Written and executed as a document-shape check. Not a complete semantic, typed-binding or runtime validator. |
| Static resolver/checker | UG-T02-S09, check_reference.py | Implemented, originally reported tested, and rerun during this handover. Fixture-specific rules and assertions, not a proven game-independent Engine. |
| Resolved builds/report | UG-T02-S09, resolved_builds.json and validation_report.json | Inspectable saved results. New rerun produced byte-identical files. See UG-T02-E02. |
| Selection and semantics explanation | UG-T02-S09, README.md | Written rationale and contract proposals. Its "Decision" heading is not owner acceptance. |
| New audit | UG-T02-S98, audit_findings.json and reproduce_audit.py | New consolidation work. Includes reproduction, hashes and scope probes. Does not change the original fixture. |
| Generative model, training code and adaptation dataset | No artifact established | Suggested only. No trained checkpoint, generation prompt suite, paired dataset or generated anime unit is present. |
| Consumer, simulator, mod and balance evaluation | No artifact established | Suggested or discussed only. None was run. |

The ordinary-build rule is three integer tiers from zero through five. After sorting descending, the second tier must not exceed two and the third must be zero. Under that encoded rule there are **64 legal ordinary builds**, including the unupgraded build. The Paragon is separate, not an ordinary `5-5-5`. The fixture examines **424 prerequisite-preserving purchase interleavings**. These are structural combinations, not 424 combat trials. [UG-T02-S09; UG-T02-E02]

The preserved price snapshot has base cost **250**; incremental path A prices **175, 450, 1,450, 10,000, 32,000**; path B **300, 800, 3,300, 6,000, 50,000**; path C **300, 300, 1,500, 2,800, 26,500**; and Paragon upgrade **750,000**. The cheapest three uncrosspathed tier-five builds plus Paragon total **886,625** under those assumptions. This is arithmetic from the fixture, not a claim about the latest installed BTD6 version, degree or balance. `exact_game_build` is null. [UG-T02-S09, economy; UG-T02-S32; UG-T02-D07]

The twelve components are `primary`, `payload`, `zone`, `stream`, `reveal`, `temporary_companion`, `permanent_companion`, `alternate_form`, `harvest`, `track_standard`, `track_heavy` and `command`. They separate direct attacks, area impact, persistent locations, target statuses, independent emitters, route-bound attacks, state changes, resource accounting and allied modifiers. Exact modifier masks and many numerical/world rules remain unbound. Adding component behavior does not automatically purchase another path's nodes. [UG-T02-S09, component_definitions and composition_contract]

### Evidence record

**UG-T02-E01. A complete structural fixture exists, but not a complete executable tower.** Directly supported by UG-T02-S09. Forty-eight unresolved groups include combat constants, propagation masks, resource/world policies and degree profiles. This count is groups, not individual scalar values. No Consumer compatibility or acceptance is demonstrated.

**UG-T02-E02. The saved static report reproduces.** Originally reported by UG-T02-S06, independently checked now through UG-T02-S98. The rerun passed **136 checks**, resolved **64 builds**, examined **424 interleavings**, and reproduced both output files byte for byte. The input JSON SHA-256 is `bd0af24ee9ebe12235345d06aec4d1e3a4e88de84aa53432fe288751d54e4946`. This supports reproducible fixture computation only. Python and environment details are in the audit.

**UG-T02-E03. The checker's apparent guarantees are narrower than its counts suggest.** New code inspection and negative probes, UG-T02-S98. `resolve` sorts purchased nodes into a canonical order before composition. Its equality test therefore does not execute live purchase histories or establish runtime order independence. A direct call with `A2, A1` for a `2-0-0` build is accepted because the function checks node membership, not sequential prerequisites. Separately, the schema accepts groups marked `bound` whose values are still null. `executable_handoff_ready` is always false in the resolver. These are scope limitations, not failed combat experiments; the original README already called this a structural checker.

**UG-T02-E04. A close source-to-card precedent uses retrieval, not demonstrated new adaptation training.** UG-T02-S12 reports Qwen3-14B with a 993-card retrieval collection, and 49 participants creating 196 cards. The paper studies representativeness and aesthetic judgments. It does not establish executable rules, TD crosspaths or balance. Relevant primary sections were inspected in this handover.

**UG-T02-E05. Description-to-character-code fine-tuning has an external precedent.** UG-T02-S13 reports CodeGemma-7B with LoRA and 502 human-AI co-designed examples across 175 roles. Engine execution and model-judged correctness are different tests. The latter is evaluated among executable outputs. The paper supports feasibility of a subtask, not complete Unit-Generator performance or dataset sufficiency.

**UG-T02-E06. Semantic mapping, card formalization, completion and bounded selection are different tasks.** Game-O-Matic maps concept relations through authored patterns; Chaos Cards generates via grammar and uses its neural model to predict strength; Latent Predictor Networks maps already-described mechanics to code; Mystical Tutor completes partial card designs; Laya scores supplied alternatives. [UG-T02-S11; UG-T02-S14 to UG-T02-S17] These sources must not be combined into a claim that one model already learns faithful, balanced anime-to-TD families.

**UG-T02-E07. End-to-end superiority and novelty are not established.** No character-generation experiment, training run or direct-LLM comparison is present. The assistant's statement that it found no verified exact end-to-end demonstration was a bounded search report, not proof that no such work exists or that a future paper would be novel. [UG-T02-S02; UG-T02-S07]

**UG-T02-E08. Reference prices provide constraints, not transferred balance.** This is assistant reasoning, not an experimental result. The proposed equation was `generated price = reference price × scale factor`; the artifact uses a scale of one. Changing mechanics, enemies, waves, income or maps changes the context in which prices matter. No balance evaluation tested this transfer. [UG-T02-S04; UG-T02-S09; UG-T02-S15]

**UG-T02-E09. Wizard was a qualitative selection.** UG-T02-S06 and the README explicitly disclaim an exhaustive numerical diversity census. The complete tree was prepared for Wizard only. Druid, Ninja, Engineer, Beast Handler and Skywarden were mechanism-level comparisons. The Druid Paragon reference corrects an older-catalog exclusion risk; it does not establish Druid's comparative rank. [UG-T02-S33 to UG-T02-S37]

**UG-T02-E10. Dataset quantities were planning suggestions.** The assistant proposed **200-500 reviewed transformation records** and **50-100 evaluation requests**, explicitly not sample-sufficiency thresholds. No such dataset was created or approved. The suggested character-level splits, anonymized evidence tests and counterfactual tests are plans, not results. [UG-T02-S02]

**UG-T02-E11. Format checking is not behavioral checking.** PICARD and XGrammar-2 concern constrained generation; the artifact checks document shape and selected structural properties. A supported syntax does not establish Consumer execution, meaningful upgrades or correct interactions. This distinction is also explicit in CONTEXT.md and the supplied research catalogue. [UG-T02-S08; UG-T02-S09; UG-T02-S21; UG-T02-S22]

**UG-T02-E12. Evaluation quality can change the apparent value of a mechanic.** The catalogue reports this through restricted-play evaluation, Mechanic Maker 2.0, Mortar and related sources. These are external findings or catalogue summaries, not project tests. The useful transfer is to record evaluator capability, scenarios and actual changed choices rather than treat one score as balance or human preference. [UG-T02-S07, R8, R17, R19; UG-T02-S46; UG-T02-S55; UG-T02-S57]

**UG-T02-E13. New literature cautions that grammar constraints can change the sampling distribution and output quality.** Grammar-Aligned Decoding studies this issue. The transfer is to measure fidelity and diversity after constraints, not assume syntax enforcement is neutral. No TD result follows. [New source UG-T02-S96]

**UG-T02-E14. New literature cautions against assuming intrinsic self-correction works.** Large Language Models Cannot Self-Correct Reasoning Yet distinguishes model-only correction from external feedback and discusses cost-matched comparisons. Its tested models/tasks do not prove current models incapable of correction, and do not test the proposed validator-guided repair. [New source UG-T02-S97]

## 5. Recommended approach and v1 boundary

### Recommendation actually reached in this thread

The assistant's final recommendation was to use Wizard as the first representation fixture, Druid as a later coverage test, and eventually train a pretrained generator on reviewed character adaptations within explicit constraints. It preferred character-driven behavior graphs over mandatory Wizard reskins. The immediate artifact was the generalized representation and checker. **There was no settled owner-approved architecture, training method or public default.** [UG-T02-D02; UG-T02-D05; UG-T02-D06; UG-T02-D08]

### Current recommendation

This subsection is new advice based on the thread and current product context. It awaits approval and is not a final architecture for Mardwerk.

**Product principle:** The public tool should turn a character description into a usable unit draft under stated game rules, while making adaptation choices and unsupported claims visible.

The smallest useful prototype is one documented standalone entry point that accepts a character name with source text or a supported source-acquisition route, a bundled Game Definition/Profile, and a design brief. It produces a complete readable unit draft and a compact record of sources, adaptation choices, configuration and scoped findings. With the proposed BTD6-like profile, completeness means the base, three five-tier paths, crosspath behavior and Paragon concept required by that profile. A formal mechanics export or proposal to add new rules remains optional.

The following are necessary for that workflow, not generic infrastructure additions:

| Proposed v1 element | Why it is needed | Boundary |
| --- | --- | --- |
| One versioned public Definition/Profile and documented editable settings | Makes "valid in this world" testable and lets public users tune within explicit limits. | Select the public default and reuse basis first. Do not require exact BTD6 parity or all Wizard capabilities. |
| One configured pretrained-model generation path with public task guidance and examples | Actually performs the character transformation, which the existing fixture does not. | Model-weight training is not a prerequisite. Provider, local support and costs need documentation rather than assumptions. |
| Explicit source scope and short adaptation record | Distinguishes supported character facts from design choices and omissions. | Pasted/source-file input can be the first supported path. Do not advertise name-only research until source acquisition actually works. |
| A complete readable draft plus checks for the claims made | Provides useful output rather than an unfilled schema and prevents false declarations of legality. | Check fields, path rules, references and arithmetic as applicable. Unchecked runtime behavior remains unchecked. Optional formal output must expose unresolved bindings. |
| Retained inputs, model/configuration revision, outputs and findings | Makes failures reproducible and provides evidence for comparing methods or later training. | Local files suffice. Towerright history or orchestration must not be required. |

A repair attempt should respond to a concrete finding and retain the previous draft. Multiple agents, a vector database, a large candidate archive or open-ended retries are not automatic v1 requirements.

The first work should therefore turn the fixture into an honest public draft contract, choose the supported model path, run actual character adaptations, and compare them with a fair direct-LLM baseline. The existing schema/checker may be reused as a fixture test, but its narrow guarantees must not become a generic validation claim. Resolving every one of the Wizard's 48 binding groups is unnecessary for a conceptual v1; producing only empty bindings is not a usable result either.

Defer model pretraining, SFT until a measured need exists, DPO/RL/distillation, a Druid fixture, broad catalog coverage, runtime code generation, combat simulation, visual assets, numerical balance certification and publication-scale studies. These are not prerequisites for the basic workflow. Later executable claims will require real Consumer support and execution evidence.

**Public/private boundary.** The current owner context assigns Mardwerk's specially tuned offering to Towerright. This thread contains no identifiable private tuned checkpoint, private prompt pack or private dataset to classify. The exact artifact boundary remains unresolved. The public core still needs sufficient rules, guidance, examples and checks to perform the promised workflow without Towerright. Do not classify those essentials as private merely because a more effective variant might later be commercial. Towerright-specific service integration and project-wide features are outside standalone v1. Its project Acceptance process must not prevent a standalone user from generating, reviewing and using local output. [UG-T02-S00; UG-T02-S08]

Foundation is explicitly private guidance in the uploaded contract and unavailable here. No license or reuse decision for the code, model, BTD6-derived data, character evidence or training examples was accepted in this thread. Preserve that uncertainty; this handover does not select licenses or assume compatibility. [UG-T02-S08, opening paragraph; UG-T02-S07, Code links; UG-T02-Q02; UG-T02-Q08]

## 6. Open questions, conflicts, and risks

**UG-T02-Q01. What makes a v1 output complete?** Product choice. The earlier discussion emphasizes formal, potentially implementable mechanics; current context does not require code, assets or balance. Decide whether v1 promises a complete conceptual draft, numerical mechanics for a named target, or both through distinct deliverables. This blocks the v1 completion test, not necessarily implementation. Resolve with an approved output example and explicit missing-information policy. [UG-T02-D01; UG-T02-D05; UG-T02-E01; UG-T02-S08]

**UG-T02-Q02. Is the BTD6-derived fixture the public default, a research reference, or both?** Product/reuse choice. The user proposed a reference, not a license or public distribution decision. Exact mechanics, prices, source snapshots, examples and copied code have separate provenance. This can block public release of that particular default. Resolve the intended retained elements, pinned reference basis and permitted reuse before packaging it. No inference about private/non-commercial legality is available from this thread. [UG-T02-D04; UG-T02-D07; UG-T02-S32; UG-T02-S38]

**UG-T02-Q03. Does "exact upgrade paths" fix only progression, or each tower's behavior graph?** Product choice. The former permits character-driven mechanics; the latter tends toward reference transformations/reskins. This changes the training task and representation scope, so resolve before curating a dataset or defining the benchmark. Compare explicit example outputs and obtain owner approval. [UG-T02-D04; UG-T02-D08; UG-T02-E06]

**UG-T02-Q04. Which game semantics must be complete now?** Product and implementation boundary. Wizard includes unresolved resource allocation, modifier masks, event ordering, transformations, Paragon degree rules and world tags. These block an executable target but not necessarily a declared conceptual deliverable. Resolve by naming the deliverable and listing supported, proposed and unimplemented behavior separately. [UG-T02-D07; UG-T02-E01; UG-T02-S09]

**UG-T02-Q05. What can the checker honestly guarantee?** Implementation issue. Canonical resolution is not live sequence execution; shape validation does not validate bound values. This blocks a strong validation/export claim, not source-based draft generation. Either narrow the API/report contract or implement and test the missing checks before using those claims. Retain the new negative probes as regression cases where applicable. [UG-T02-E02; UG-T02-E03; UG-T02-S98]

**UG-T02-Q06. Does model-weight training improve the task enough to justify its data cost?** Research and implementation choice. Reviewed adaptations, definitions of good mappings and held-out characters do not yet exist. This blocks a defensible training investment, not a pretrained-model v1. First measure the untuned path's errors and correction effort; train only against a specified bottleneck. [UG-T02-D02; UG-T02-D10; UG-T02-E05; UG-T02-E10]

**UG-T02-Q07. How does the public tool obtain evidence and invoke a model?** Product/deployment choice. Standalone has not settled name-only retrieval, supported providers, local operation, hardware or running cost. This blocks an honest install-and-run promise. Resolve a supported path and verify it without private infrastructure. A text-input v1 must be labeled as such rather than marketed as automatic wiki research. [UG-T02-S00; UG-T02-S01; UG-T02-D02]

**UG-T02-Q08. Which specific tuned artifacts are Towerright-private?** Commercial/disclosure boundary. The current owner direction is clear, but the checkpoint, prompts, rules, examples, evaluator or configuration to which it applies is not identified. This matters before publication, GitHub packaging or a private/public benchmark comparison. Classify actual assets when they exist, and verify that public generation is independently usable. The older contract's project Acceptance responsibility must not be interpreted as a private authorization gate on standalone output. [UG-T02-S00; UG-T02-S08; UG-T02-S92; UG-T02-S93]

**UG-T02-Q09. What does "better than a direct LLM prompt" mean?** Research and product evidence. No comparator was run. Resolve primary measures, equivalent source/rule information and total resource accounting before making improvement claims. This affects v1 evaluation and later research, not whether a prototype can be built. Constraint success, preserved signature behavior and correction effort are the closest measures supported here. [UG-T02-D10; UG-T02-E07; UG-T02-E13; UG-T02-E14]

**UG-T02-Q10. Does the representation generalize beyond Wizard?** Later research/coverage issue. Twelve components in one fixture do not establish catalog coverage or highest variety. A contrasting character or later Druid/Engineer case can expose missing contracts, but broad catalog completeness is not needed for the first useful workflow. [UG-T02-D06; UG-T02-D08; UG-T02-E09]

## 7. Sources and research leads

### A. Sources already present in this thread

The full register is supplied as **UG-T02-source-register.md** and **UG-T02-source-register.json**. It contains a separate entry for every identifiable named paper/reference in the available exchange and attachments, with ordinary URLs, original locator, publication information where verified, content-inspection scope, transfer limits, associated artifacts and linked IDs. The original wording of all 38 numbered catalogue records is retained in the JSON. Incomplete references and missing internal documents are not discarded.

The register distinguishes UG-T02-S01 to UG-T02-S09 (messages and supplied artifacts), UG-T02-S10 to UG-T02-S38 (directly discussed papers, models, tools and game sources), UG-T02-S39 to UG-T02-S76 (catalogue R1-R38), UG-T02-S77 to UG-T02-S86 (ten existing extra papers/leads), UG-T02-S87 to UG-T02-S95 (unavailable internal documents and repository pointers), and UG-T02-S99 (unnamed unresolved research trails). UG-T02-S00 is current owner context; UG-T02-S98 is new audit evidence, not historical literature.

The most directly relevant existing external sources are:

| ID | Source and verification | Why retained |
| --- | --- | --- |
| UG-T02-S12 | Pfau and Vrettis, From LLM-Driven Trading Card Generation to Procedural Relatedness: A Pokémon Case Study, arXiv:2604.27972v1, 2026 preprint. Bibliography verified; relevant sections inspected. | Closest description-to-themed-mechanics input/output precedent, with no executable balance result. |
| UG-T02-S13 | Wu et al., Open Role-Playing with Delta-Engines, arXiv:2408.05842v5, 2024/2025 manuscript. Bibliography verified; training/evaluation sections inspected. | Closest specialist fine-tuning precedent, but engine patching rather than full constrained TD families. |
| UG-T02-S14 | Treanor, Blackford, Mateas and Bogost, Game-O-Matic: Generating Videogames that Represent Ideas, PCG workshop 2012, DOI:10.1145/2538528.2538537. Relevant sections inspected. | Maps meanings and relationships to mechanics; representational success not fully evaluated. |
| UG-T02-S15 | Chen and Guy, Chaos Cards, AIIDE 2020, DOI:10.1609/aiide.v16i1.7430. Relevant sections inspected. | Distinguishes grammar-based generation from neural strength prediction. |
| UG-T02-S16 | Ling et al., Latent Predictor Networks for Code Generation, ACL 2016, DOI:10.18653/v1/P16-1057. Relevant sections inspected. | Separates formalization of already-described rules from adaptation design. |
| UG-T02-S17 | Summerville and Mateas, Mystical Tutor, AIIDE 2016, DOI:10.1609/aiide.v12i1.12851. Bibliography verified; abstract only. | Controllable completion of partial specifications. |
| UG-T02-S39 | Sorochan and Guzdial, Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search, arXiv:2212.03387v1, 2022. Relevant sections inspected. | Bounded unit generation and evaluation, not character interpretation. |

Direct URLs for these priority records:

UG-T02-S12: https://arxiv.org/html/2604.27972v1

UG-T02-S13: https://arxiv.org/html/2408.05842v5

UG-T02-S14: https://mtreanor.com/publications/Game-O-Matic-PCG2012.pdf

UG-T02-S15: https://ojs.aaai.org/index.php/AIIDE/article/view/7430

UG-T02-S16: https://aclanthology.org/P16-1057/

UG-T02-S17: https://ojs.aaai.org/index.php/AIIDE/article/view/12851

UG-T02-S39: https://arxiv.org/html/2212.03387v1

Three bibliographic discrepancies found during consolidation are recorded without rewriting the original catalogue. UG-T02-S53's current record uses "Domain Specific Language" in the Ludax title, while the supplied entry uses "Description Language". UG-T02-S66's arXiv v1 title includes "PCGRL+:" despite the catalogue describing that prefix as an informal label. UG-T02-S84 is by **Matthew C. Fontaine**, not "Michael C. Fontaine", and its title uses "through MAP-Elites" rather than the catalogue's "with MAP-Elites". These are record corrections, not changes to project decisions.

Some DOI and publisher accesses still failed, including the early TD level-generation references and several creativity papers. Their supplied metadata remains explicitly unverified, not silently dropped. A recovered title or abstract does not validate a paper's detailed results. The new register is a targeted bibliographic/claim check, not a full reread or reproduction of every paper in the catalogue.

### B. Newly identified literature leads

**UG-T02-S96. Newly found, not part of the original conversation.** Kanghee Park, Jiayu Wang, Taylor Berg-Kirkpatrick, Nadia Polikarpova and Loris D'Antoni. *Grammar-Aligned Decoding*. NeurIPS 2024; arXiv:2405.21047. Bibliographic record verified through arXiv and proceedings. Relevant method and evaluation-framing sections of v1 inspected. It addresses distortion of the model's conditional distribution by ordinary grammar-constrained decoding. This challenges an assumption in UG-T02-D05 and motivates UG-T02-E13, but does not establish game-design quality or require adding its algorithm to v1. Code was not executed.

https://arxiv.org/html/2405.21047v1

https://openreview.net/forum?id=WAidyyRaes

https://github.com/ebmoon/transformers-GAD

**UG-T02-S97. Newly found, not part of the original conversation.** Jie Huang, Xinyun Chen, Swaroop Mishra, Huaixiu Steven Zheng, Adams Wei Yu, Xinying Song and Denny Zhou. *Large Language Models Cannot Self-Correct Reasoning Yet*. arXiv:2310.01798, 2023; ICLR 2024 record located. ArXiv bibliography verified; the OpenReview page was blocked during this handover. Relevant v1 sections on feedback source, comparisons and limitations inspected. It challenges ungrounded critique loops, not correction informed by a reliable external validator. Its results concern tested older models and reasoning tasks, not all current models or TD design. [UG-T02-E14]

https://arxiv.org/html/2310.01798v1

https://openreview.net/forum?id=IkmD3fKBPQ

The remaining literature gaps are paired character-adaptation data, source-fidelity evaluation under constrained mechanics, controllability across rule changes, full-family interaction testing and trustworthy gameplay evaluation. No exact end-to-end absence or publication novelty is established by this extraction.

### Complete source register supplement

### UG-T02 source register

Internal consolidation supplement. Prepared September 24, 2026. Source IDs are stable within this handover, not global IDs for the Mardwerk project.

This register distinguishes original sources from newly found literature and the new local audit. Bibliographic verification does not mean full-text review. "Abstract only" and "metadata only" are deliberately weaker than inspection of methods/results. The catalogue's historical claims of inspection are preserved as attributed records, not silently adopted as current verification. Papers and public code were not reproduced here; the Wizard reference checker was rerun.

Source text for all 38 numbered catalogue entries is retained in the JSON companion. Exact long author lists absent from the supplied catalogue remain "et al." unless independently resolved; final venues are not inferred from an arXiv deposit. No source or dependency has received a reuse or license approval through this register.

#### Current owner context

##### UG-T02-S00. Current owner-supplied handover context

Current user handover request; no verified message timestamp.

Type: Owner context. Origin: Current product context supplied for consolidation.

Original locator: User: "CURRENT PRODUCT CONTEXT".

Bibliographic record: verified as visible supplied context. Content inspected: full text.

Relevance: Controls current advice, not evidence of older agreement.

Limitations: Does not establish implementation, approvals in earlier messages, or model-weight tuning.

Linked record: Section 5; UG-T02-Q01; UG-T02-Q07

Direct URL: Not supplied or not applicable. Use the message or file locator above.

#### A. Sources already present in this thread

##### UG-T02-S01. Original research request

Visible message. Author indicated in locator. Message date/link unavailable.

Type: Original conversation message. Origin: Already present in this thread.

Original locator: User: "BEST CASE i would like to train an AI model".

Bibliographic record: verified as visible message. Content inspected: full text.

Relevance: Character descriptions to full TD units under a supplied Game Definition; research model/training approaches.

Limitations: Assistant proposals are not owner approval. Historical web tool traces are not available.

Linked record: UG-T02-D01 to UG-T02-D10

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S02. Assistant research and training recommendation

Visible message. Author indicated in locator. Message date/link unavailable.

Type: Original conversation message. Origin: Already present in this thread.

Original locator: Assistant: "post-training an existing generative model is a credible approach".

Bibliographic record: verified as visible message. Content inspected: full text.

Relevance: Source-grounded interpretation, structured mechanics, external checking, candidate models and later training/evaluation options.

Limitations: Assistant proposals are not owner approval. Historical web tool traces are not available.

Linked record: UG-T02-D01 to UG-T02-D10

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S03. BTD6-derived reference proposal

Visible message. Author indicated in locator. Message date/link unavailable.

Type: Original conversation message. Origin: Already present in this thread.

Original locator: User: "keeping the exact upgrad paths, relativ cost structures".

Bibliographic record: verified as visible message. Content inspected: full text.

Relevance: Three paths by five upgrades, crosspaths, Paragon and generalized behavior.

Limitations: Assistant proposals are not owner approval. Historical web tool traces are not available.

Linked record: UG-T02-D01 to UG-T02-D10

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S04. Assistant bounded-family training plan

Visible message. Author indicated in locator. Message date/link unavailable.

Type: Original conversation message. Origin: Already present in this thread.

Original locator: Assistant: "The main distinction is preserving BTD6's progression structure".

Bibliographic record: verified as visible message. Content inspected: full text.

Relevance: Whole-tree design; Medium-price interpretation; fixed constraints outside model; learned adaptation rather than reskin.

Limitations: Assistant proposals are not owner approval. Historical web tool traces are not available.

Linked record: UG-T02-D01 to UG-T02-D10

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S05. Catalog comparison and representation request

Visible message. Author indicated in locator. Message date/link unavailable.

Type: Original conversation message. Origin: Already present in this thread.

Original locator: User: "skip like dart monkey and super monkey" and "create the generalized unit representation".

Bibliographic record: verified as visible message. Content inspected: full text.

Relevance: Scoped comparison and artifact creation explicitly requested.

Limitations: Assistant proposals are not owner approval. Historical web tool traces are not available.

Linked record: UG-T02-D01 to UG-T02-D10

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S06. Assistant Wizard recommendation and artifact handoff

Visible message. Author indicated in locator. Message date/link unavailable.

Type: Original conversation message. Origin: Already present in this thread.

Original locator: Assistant: "Wizard Monkey as the first generalized reference, then Druid as the second".

Bibliographic record: verified as visible message. Content inspected: full text.

Relevance: Qualitative selection, parametric representation and reported static results.

Limitations: Assistant proposals are not owner approval. Historical web tool traces are not available.

Linked record: UG-T02-D01 to UG-T02-D10

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S07. Research sources for unit generation and game design

GAME-DESIGN-RESEARCH.md; declares research checked September 22, 2026; 499 supplied citation lines.

Type: Uploaded secondary research catalogue. Origin: Already present in this thread.

Original locator: Uploaded GAME-DESIGN-RESEARCH.md, R1-R38 and Additional references.

Bibliographic record: verified as inspected attachment. Content inspected: full text.

Relevance: Preserves 38 numbered papers, ten additional named references, proposed evaluation adaptations and access gaps.

Limitations: Its historical "inspected" labels are source assertions, not reproductions or a new full-text audit. No referenced implementation was run by that review.

Linked record: UG-T02-E12; UG-T02-S39 to UG-T02-S86

Direct URL: Not supplied or not applicable. Use the message or file locator above.

Associated files/code/data: /mnt/data/GAME-DESIGN-RESEARCH.md

##### UG-T02-S08. Unit Generator context

CONTEXT.md; contributor contract, undated in its contents; 60 supplied citation lines.

Type: Uploaded project contract. Origin: Already present in this thread.

Original locator: Uploaded CONTEXT.md, Shared terms; Unit terms; Generation and evaluation terms.

Bibliographic record: verified as inspected attachment. Content inspected: full text.

Relevance: Defines Unit, Interpretation, Definition, Validation, Consumer, project Acceptance and ownership.

Limitations: Normative descriptions are not evidence that existing repository modules implement them. Foundation and other linked documents were unavailable.

Linked record: UG-T02-D05; UG-T02-Q01; UG-T02-Q08

Direct URL: Not supplied or not applicable. Use the message or file locator above.

Associated files/code/data: /mnt/data/CONTEXT.md

##### UG-T02-S09. Generalized unit representation: Wizard reference

Assistant-produced bundle; README and JSON declare September 23, 2026; schema_version 0.1.0.

Type: Inspectable technical artifact. Origin: Already present in this thread.

Original locator: Assistant Wizard reply and bundle README sections Status, Reusable representation, Validation.

Bibliographic record: verified as inspected and rerun artifact. Content inspected: full text and code; rerun in this handover.

Relevance: Complete parametric family, costs, structural resolver and saved results.

Limitations: 48 unbound groups; no installed-game export, AI generator, combat execution, training or balance result. Artifact section titled Decision is only an assistant recommendation.

Linked record: UG-T02-D06; UG-T02-D07; UG-T02-E01 to UG-T02-E03

Direct URL: Not supplied or not applicable. Use the message or file locator above.

Associated files/code/data: /mnt/data/generalized_unit_reference.zip; README.md; generalized_wizard.unit.json; unit-draft.schema.json; check_reference.py; resolved_builds.json; validation_report.json

##### UG-T02-S10. caveman skill

JuliusBrussee/caveman, skills/caveman/SKILL.md; mutable main branch; user requested LITE.

Type: Prompt/style resource. Origin: Already present in this thread.

Original locator: Opening user request.

Bibliographic record: verified resource. Content inspected: full text.

Relevance: Presentation preference for the original answer.

Limitations: Not a Unit-Generator task prompt, model-training recipe, architecture or evaluation result.

Linked record: UG-T02-S01

Direct URLs:
https://raw.githubusercontent.com/JuliusBrussee/caveman/refs/heads/main/skills/caveman/SKILL.md

##### UG-T02-S11. Laya

Convai Innovations; convaiinnovations/laya model card; mutable main revision; year/version not pinned in thread.

Type: Model card and model repository. Origin: Already present in this thread.

Original locator: Opening user link; assistant model comparison.

Bibliographic record: verified model-card record. Content inspected: relevant sections.

Relevance: 421M English decision-head model scores supplied options; considered for ranking/routing rather than open unit generation.

Limitations: Author-reported benchmark: root 0.362 accuracy, majority 0.461, domain-trained checkpoint 0.766. Different checkpoint results must not be conflated. Root input default 512 tokens per question; other variants differ. No Unit-Generator experiment.

Linked record: UG-T02-D02; UG-T02-E06

Direct URLs:
https://huggingface.co/convaiinnovations/laya
https://huggingface.co/convaiinnovations/laya/raw/main/README.md

Associated files/code/data: Root, typed-decisions and multilingual variants documented in model card; no weights run.

##### UG-T02-S12. From LLM-Driven Trading Card Generation to Procedural Relatedness: A Pokémon Case Study

Johannes Pfau and Panagiotis Vrettis. 2026 preprint, arXiv:2604.27972v1, submitted April 30, 2026.

Type: Research preprint. Origin: Already present in this thread.

Original locator: Assistant first research reply, missing paper A.

Bibliographic record: verified. Content inspected: relevant sections: generation pipeline and user study.

Relevance: Close precedent for converting a creature description into themed card properties; Qwen3-14B, retrieval over 993 cards; 49 participants, 196 cards.

Limitations: Representativeness/aesthetic evaluation, not executable correctness or TD balance; not a demonstrated newly trained adaptation model.

Linked record: UG-T02-E04; UG-T02-D02

Direct URLs:
https://arxiv.org/abs/2604.27972
https://arxiv.org/html/2604.27972v1

Associated files/code/data: https://github.com/JohannesPfau/generativePokemonTCG ; card extraction and generation artifacts linked by paper, not executed.

##### UG-T02-S13. Open Role-Playing with Delta-Engines

Hongqiu Wu, Zekai Xu, Tianyang Xu, Shize Wei, Yan Wang, Jiale Hong, Weiqi Wu and Hai Zhao. arXiv:2408.05842, first 2024; inspected v5, March 7, 2025. Publication venue not established in inspected record.

Type: Research preprint/revised manuscript. Origin: Already present in this thread.

Original locator: Assistant first research reply, missing paper B.

Bibliographic record: verified. Content inspected: relevant sections: model training, engine changes and evaluation, especially 6.1 and 6.2.

Relevance: Fine-tunes CodeGemma-7B with LoRA using 502 co-designed examples across 175 roles; description-to-engine-change precedent.

Limitations: Execution testing and model-judged correctness are distinct; correctness is judged among executable outputs. Engine patches are not constrained BTD6 families or demonstrated balance.

Linked record: UG-T02-E05; UG-T02-D02

Direct URLs:
https://arxiv.org/abs/2408.05842
https://arxiv.org/html/2408.05842v5

Associated files/code/data: CodeGemma-7B/codegemma-7b-it named training base; retrieval of engine methods. No training or code run here.

##### UG-T02-S14. Game-O-Matic: Generating Videogames that Represent Ideas

Mike Treanor, Bryan Blackford, Michael Mateas and Ian Bogost. Third Workshop on Procedural Content Generation in Games, 2012. DOI:10.1145/2538528.2538537.

Type: Published workshop paper. Origin: Already present in this thread.

Original locator: Assistant first research reply, missing paper C.

Bibliographic record: verified. Content inspected: relevant sections: micro-rhetorics, recipes, future work; first page visually checked.

Relevance: Explicit concept relationships mapped to human-authored mechanics patterns.

Limitations: Formal evaluation of representational success is future work in inspected text. Not learned wiki-to-tower adaptation.

Linked record: UG-T02-E06; UG-T02-D05

Direct URLs:
https://mtreanor.com/publications/Game-O-Matic-PCG2012.pdf
https://dl.acm.org/doi/10.1145/2538528.2538537
https://profiles.wustl.edu/en/publications/game-o-matic-generating-videogames-that-represent-ideas/

##### UG-T02-S15. Chaos Cards: Creating Novel Digital Card Games through Grammatical Content Generation and Meta-Based Card Evaluation

Tiannan Chen and Stephen J. Guy. AIIDE 16(1), 196-202, 2020. DOI:10.1609/aiide.v16i1.7430.

Type: Published proceedings paper. Origin: Already present in this thread.

Original locator: Assistant first research reply, missing paper D.

Bibliographic record: verified. Content inspected: relevant sections: grammar generation and meta-based evaluation; first page visually checked.

Relevance: Grammar generates executable cards; a recursive neural model predicts strength.

Limitations: The neural model is not the generator. Meta-dependent strength prediction does not establish character fidelity or preference.

Linked record: UG-T02-E06; UG-T02-E08

Direct URLs:
https://ojs.aaai.org/index.php/AIIDE/article/view/7430
https://cdn.aaai.org/ojs/7430/7430-52-10748-1-2-20200923.pdf

Associated files/code/data: Grammar and supplementary-material links in paper; not executed or license-audited.

##### UG-T02-S16. Latent Predictor Networks for Code Generation

Wang Ling, Phil Blunsom, Edward Grefenstette, Karl Moritz Hermann, Tomáš Kočiský, Fumin Wang and Andrew Senior. ACL 2016, 599-609. DOI:10.18653/v1/P16-1057; arXiv:1603.06744v2.

Type: Published conference paper. Origin: Already present in this thread.

Original locator: Assistant first research reply, missing paper E.

Bibliographic record: verified. Content inspected: relevant sections: mixed-input code generation and card datasets.

Relevance: Natural-language plus structured card descriptions mapped into implementation code.

Limitations: Card descriptions already state mechanics; this does not decide how a fictional character should be adapted.

Linked record: UG-T02-E06

Direct URLs:
https://aclanthology.org/P16-1057/
https://arxiv.org/html/1603.06744v2

Associated files/code/data: Magic: The Gathering and Hearthstone code-generation datasets described by paper; historical dataset pointer https://deepmind.com/publications.html, current dataset availability not verified.

##### UG-T02-S17. Mystical Tutor: A Magic: The Gathering Design Assistant via Denoising Sequence-to-Sequence Learning

Adam Summerville and Michael Mateas. AIIDE 12(1), 86-92, 2016. DOI:10.1609/aiide.v12i1.12851.

Type: Published conference paper. Origin: Already present in this thread.

Original locator: Assistant first research reply, missing paper F.

Bibliographic record: verified. Content inspected: abstract only.

Relevance: Partial card specifications completed by a denoising sequence-to-sequence model.

Limitations: Original reply explicitly lacked full-text access. This handover inspected publisher metadata and abstract, not experiments.

Linked record: UG-T02-E06

Direct URLs:
https://ojs.aaai.org/index.php/AIIDE/article/view/12851

##### UG-T02-S18. LoRA: Low-Rank Adaptation of Large Language Models

Edward J. Hu, Yelong Shen, Phillip Wallis, Zeyuan Allen-Zhu, Yuanzhi Li, Shean Wang, Lu Wang and Weizhu Chen. arXiv:2106.09685, 2021 preprint; ICLR 2022 record linked.

Type: Training-method paper. Origin: Already present in this thread.

Original locator: Assistant first and second training recommendations.

Bibliographic record: verified arXiv bibliographic record; venue not independently re-audited. Content inspected: abstract only.

Relevance: Parameter-efficient supervised adaptation of an existing model.

Limitations: Does not provide task-specific data or establish character-adaptation quality.

Linked record: UG-T02-D02

Direct URLs:
https://arxiv.org/abs/2106.09685
https://openreview.net/forum?id=nZeVKeeFYf9

Associated files/code/data: https://github.com/microsoft/LoRA ; not executed.

##### UG-T02-S19. QLoRA: Efficient Finetuning of Quantized LLMs

Tim Dettmers, Artidoro Pagnoni, Ari Holtzman and Luke Zettlemoyer. arXiv:2305.14314, 2023.

Type: Training-method paper/preprint record. Origin: Already present in this thread.

Original locator: Assistant first reply, model training.

Bibliographic record: verified arXiv record; final proceedings status not independently rechecked. Content inspected: abstract only.

Relevance: Adapters trained through a frozen quantized base to reduce training memory.

Limitations: Memory-saving technique, not evidence that a particular unit dataset is sufficient.

Linked record: UG-T02-D02

Direct URLs:
https://arxiv.org/abs/2305.14314

##### UG-T02-S20. Direct Preference Optimization: Your Language Model is Secretly a Reward Model

Rafael Rafailov, Archit Sharma, Eric Mitchell, Stefano Ermon, Christopher D. Manning and Chelsea Finn. arXiv:2305.18290, first 2023; current record v3, 2024.

Type: Preference-training paper/preprint record. Origin: Already present in this thread.

Original locator: Assistant first reply, preference training later.

Bibliographic record: verified arXiv record; final proceedings version not audited. Content inspected: abstract only.

Relevance: Learn from preferred/rejected outputs for the same request.

Limitations: No TD preference dataset, trained model or taste result exists in thread.

Linked record: UG-T02-D09

Direct URLs:
https://arxiv.org/abs/2305.18290

##### UG-T02-S21. PICARD: Parsing Incrementally for Constrained Auto-Regressive Decoding from Language Models

Torsten Scholak, Nathan Schucher and Dzmitry Bahdanau. EMNLP 2021. DOI:10.18653/v1/2021.emnlp-main.779.

Type: Published conference paper. Origin: Already present in this thread.

Original locator: Assistant first and second replies, constrained generation.

Bibliographic record: verified. Content inspected: abstract only.

Relevance: Incremental parsing rejects inadmissible continuations in a text-to-SQL task.

Limitations: Parsing constraints do not prove crosspath behavior, runtime correctness or balance.

Linked record: UG-T02-E11; UG-T02-D05

Direct URLs:
https://aclanthology.org/2021.emnlp-main.779/

##### UG-T02-S22. XGrammar-2: Dynamic and Efficient Structured Generation Engine for Agentic LLMs

Linzhang Li, Yixin Dong, Guanjie Wang, Ziyi Xu, Alexander Jiang and Tianqi Chen. arXiv:2601.04426, 2026. Original first-version title uses "XGrammar 2" without hyphen.

Type: Research manuscript and associated software. Origin: Already present in this thread.

Original locator: Assistant first reply named "XGrammar-2" without a persistent identifier.

Bibliographic record: verified arXiv record; final venue not reverified. Content inspected: relevant introductory sections.

Relevance: Dynamic structured-generation constraints.

Limitations: Efficiency/format validity are not behavioral validity. Exact bibliographic record recovered during handover, not a newly proposed literature lead.

Linked record: UG-T02-E11

Direct URLs:
https://arxiv.org/abs/2601.04426
https://arxiv.org/html/2601.04426v1

Associated files/code/data: https://github.com/mlc-ai/xgrammar ; not executed.

##### UG-T02-S23. Distilling System 2 into System 1

Ping Yu, Jing Xu, Jason Weston and Ilia Kulikov. arXiv:2407.06023, 2024, v3 record; relevant v1 text inspected.

Type: Research preprint. Origin: Already present in this thread.

Original locator: Assistant first reply, fast generation later.

Bibliographic record: verified. Content inspected: relevant sections: distillation method and scope.

Relevance: Train on filtered outputs of costlier inference rather than deploy its full process.

Limitations: Transfer varies by task; no fast Unit-Generator model demonstrated.

Linked record: UG-T02-D09

Direct URLs:
https://arxiv.org/abs/2407.06023
https://arxiv.org/html/2407.06023v1

##### UG-T02-S24. LILO: Learning Interpretable Libraries by Compressing and Documenting Code

Gabriel Grand, Lionel Wong, Maddy Bowers, Theo X. Olausson, Muxin Liu, Joshua B. Tenenbaum and Jacob Andreas. arXiv:2310.19791, first 2023.

Type: Research manuscript. Origin: Already present in this thread.

Original locator: Assistant first reply, reusable abstractions later.

Bibliographic record: verified arXiv record; final venue not audited. Content inspected: abstract only in handover.

Relevance: Potential later reference for reusable mechanic abstractions.

Limitations: An inferred macro is not automatic permission to extend a Consumer or Game Definition.

Linked record: UG-T02-D09

Direct URLs:
https://arxiv.org/abs/2310.19791
https://arxiv.org/html/2310.19791v4

##### UG-T02-S25. Qwen3-14B

Qwen team; Qwen/Qwen3-14B; exact deployed revision not selected.

Type: Official model card. Origin: Already present in this thread.

Original locator: Assistant first reply, candidate model comparison.

Bibliographic record: verified model-card record. Content inspected: relevant model-card sections.

Relevance: Proposed reference models for a controlled experiment.

Limitations: Not a selected dependency or a demonstrated suitable hardware/cost configuration. No weights run.

Linked record: UG-T02-D02; UG-T02-Q07

Direct URLs:
https://huggingface.co/Qwen/Qwen3-14B

##### UG-T02-S26. Qwen3-8B

Qwen team; Qwen/Qwen3-8B; exact deployed revision not selected.

Type: Official model card. Origin: Already present in this thread.

Original locator: Assistant first reply, candidate model comparison.

Bibliographic record: verified model-card record. Content inspected: relevant model-card sections.

Relevance: Proposed reference models for a controlled experiment.

Limitations: Not a selected dependency or a demonstrated suitable hardware/cost configuration. No weights run.

Linked record: UG-T02-D02; UG-T02-Q07

Direct URLs:
https://huggingface.co/Qwen/Qwen3-8B

##### UG-T02-S27. TRL SFTTrainer

Hugging Face TRL documentation; current unpinned documentation.

Type: Official software documentation. Origin: Already present in this thread.

Original locator: Assistant first reply, suggested training implementation.

Bibliographic record: verified documentation. Content inspected: relevant sections.

Relevance: SFT implementation with adapter support.

Limitations: No project installation or training run.

Linked record: UG-T02-D02

Direct URLs:
https://huggingface.co/docs/trl/sft_trainer

##### UG-T02-S28. PEFT

Hugging Face PEFT documentation; library named in original TRL suggestion.

Type: Official software documentation. Origin: Already present in this thread.

Original locator: Assistant first reply named "PEFT adapters"; standalone URL recovered in handover.

Bibliographic record: verified documentation. Content inspected: relevant introductory sections.

Relevance: Adapter training option.

Limitations: Named tool is not an approved dependency; no execution.

Linked record: UG-T02-D02

Direct URLs:
https://huggingface.co/docs/peft/index

##### UG-T02-S29. BTD6 Internal Structure

gurrenm3/BTD-Mod-Helper community documentation; mutable site.

Type: Implementation reference. Origin: Already present in this thread.

Original locator: Assistant second reply, interpreting normal prices as Medium.

Bibliographic record: verified documentation record. Content inspected: relevant sections.

Relevance: Explains unmodified Medium-price values in default game models.

Limitations: Community tooling documentation, not owner confirmation of intended price basis. No installed version exported.

Linked record: UG-T02-D07; UG-T02-E08

Direct URLs:
https://gurrenm3.github.io/BTD-Mod-Helper/wiki/BTD6-Internal-Structure

##### UG-T02-S30. ModTower.cs

gurrenm3/BTD-Mod-Helper source, mutable master branch.

Type: Community source code. Origin: Already present in this thread.

Original locator: Assistant second reply, legal tier triples.

Bibliographic record: verified source record. Content inspected: relevant code sections.

Relevance: Ordinary upgrade rule comparison for sorted tier values.

Limitations: Not a license audit, project dependency approval, or full official game specification.

Linked record: UG-T02-D04; UG-T02-E02

Direct URLs:
https://raw.githubusercontent.com/gurrenm3/BTD-Mod-Helper/master/BloonsTD6%20Mod%20Helper/Api/Towers/ModTower.cs

##### UG-T02-S31. Making a Custom Tower

gurrenm3/BTD-Mod-Helper community documentation; mutable site.

Type: Tutorial. Origin: Already present in this thread.

Original locator: Assistant second reply, separate appearance and mechanics.

Bibliographic record: verified documentation record. Content inspected: relevant sections.

Relevance: Card Monkey example changes display and pierce separately.

Limitations: Does not establish generator implementation or asset reuse permission.

Linked record: UG-T02-D05

Direct URLs:
https://gurrenm3.github.io/BTD-Mod-Helper/wiki/Making-a-Custom-Tower

##### UG-T02-S32. Cyber Quincy costs.json

hemisemidemipresent/cyberquincy, jsons/costs.json; mutable master branch, no commit pinned.

Type: Community-maintained price dataset. Origin: Already present in this thread.

Original locator: Assistant Wizard reply and bundle price provenance.

Bibliographic record: verified accessible dataset; installed-game parity unverified. Content inspected: relevant Wizard and Paragon entries.

Relevance: Source of populated reference prices.

Limitations: Keep bundle snapshot distinct from current web values; exact game build null. Costs do not prove balance. Reuse terms not audited.

Linked record: UG-T02-D07; UG-T02-E08

Direct URLs:
https://raw.githubusercontent.com/hemisemidemipresent/cyberquincy/master/jsons/costs.json

##### UG-T02-S33. Wizard Monkey (BTD6)

Bloons Wiki/Fandom, community-maintained page; no fixed revision supplied.

Type: Community wiki. Origin: Already present in this thread.

Original locator: Assistant Wizard selection and upgrade mapping.

Bibliographic record: verified accessible page, not official game parity. Content inspected: relevant sections.

Relevance: Mechanism descriptions and cross-checking reference mapping.

Limitations: Mutable page; historical summaries and change logs can refer to different game versions. Not a complete installed-data export.

Linked record: UG-T02-D06; UG-T02-E09

Direct URLs:
https://bloons.fandom.com/wiki/Wizard_Monkey_(BTD6)

##### UG-T02-S34. Bloons TD 6 update notes, version 37.0

Ninja Kiwi update announcement, June 7, 2023; inspected SteamDB reproduction.

Type: Developer announcement mirrored by third party. Origin: Already present in this thread.

Original locator: Assistant Wizard comparison and README provenance.

Bibliographic record: verified announcement record on mirror. Content inspected: relevant sections.

Relevance: Original use: Magus Perfectus resource cycle, transformations and other tower references.

Limitations: A patch description is not complete tower behavior or a frozen all-stat specification.

Linked record: UG-T02-D06; UG-T02-D07; UG-T02-E09

Direct URLs:
https://steamdb.info/patchnotes/11401187/

##### UG-T02-S35. Bloons TD 6 update notes, version 45.0

Ninja Kiwi update announcement, October 9, 2024; inspected SteamDB reproduction.

Type: Developer announcement mirrored by third party. Origin: Already present in this thread.

Original locator: Assistant Wizard comparison and README provenance.

Bibliographic record: verified announcement record on mirror. Content inspected: relevant sections.

Relevance: Original use: Patch descriptions used in the qualitative tower comparison.

Limitations: A patch description is not complete tower behavior or a frozen all-stat specification.

Linked record: UG-T02-D06; UG-T02-D07; UG-T02-E09

Direct URLs:
https://steamdb.info/patchnotes/15962359/

##### UG-T02-S36. Bloons TD 6 update notes, version 54.0

Ninja Kiwi update announcement, April 9, 2026; inspected SteamDB reproduction.

Type: Developer announcement mirrored by third party. Origin: Already present in this thread.

Original locator: Assistant Wizard comparison and README provenance.

Bibliographic record: verified announcement record on mirror. Content inspected: relevant sections.

Relevance: Original use: Druid Paragon Root of all Nature and economy interactions.

Limitations: A patch description is not complete tower behavior or a frozen all-stat specification.

Linked record: UG-T02-D06; UG-T02-D07; UG-T02-E09

Direct URLs:
https://steamdb.info/patchnotes/22719852/

##### UG-T02-S37. Bloons TD 6 Steam announcements

Ninja Kiwi/Steam; mutable all-news page; several updates share URL.

Type: Developer announcements. Origin: Already present in this thread.

Original locator: Assistant Wizard reply and README, current-catalog candidates including Skywarden.

Bibliographic record: verified accessible announcement feed. Content inspected: relevant entries.

Relevance: Reference for current-catalog mechanism examples in historical recommendation.

Limitations: Mutable aggregate is an inadequate long-term locator for exact costs or behavior. No exhaustive catalog census established.

Linked record: UG-T02-E09

Direct URLs:
https://steamcommunity.com/app/960090/allnews/

##### UG-T02-S38. Copyright in Derivative Works and Compilations

U.S. Copyright Office, Circular 14; inspected PDF carries 2020 revision mark.

Type: Government explanatory circular. Origin: Already present in this thread.

Original locator: Assistant first reply, rights boundary.

Bibliographic record: verified. Content inspected: relevant introductory material; first page visually checked.

Relevance: Historical reminder that a generated adaptation is not rights clearance.

Limitations: U.S. explanatory document; not a legal determination of BTD6/anime reuse, EU law, dataset training or a license choice.

Linked record: UG-T02-Q02

Direct URLs:
https://copyright.gov/circs/circ14.pdf

##### UG-T02-S39. Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search

Kynan Sorochan and Matthew Guzdial. Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search. Experimental AI in Games Workshop, 2022. arXiv:2212.03387v1 (https://arxiv.org/html/2212.03387v1). Inspected: Search Space, Evaluator, Search Method, generated units, evaluation and references.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R1; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: relevant sections: search space, evaluator and method.

Relevance: represent a candidate through explicit triggers, targets, effects, costs, scope and limits.

Limitations:  this is invention within a supplied vocabulary. Its utility measure includes unit lifetime, which may undervalue disposable or burst units. Symmetric win rates cannot by themselves rule out a mandatory purchase or a collapsed strategy space. Those are our criticisms of transferring the evaluator, not additional results reported by the authors. Tower defense also lacks the same adversarial purchase structure. Incoming damage, death and movement must not be imported into a Definition that does not support them.  the evaluation uses one map, and the strongest agent often underuses the new units. Its preference for the original roster can confound unit usefulness with evaluator competence. The search budgets are not validated human skill models; generated names and shared visuals are not evidence of visual-content generation.

Linked record: UG-T02-E06

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2212.03387v1
https://arxiv.org/abs/2212.03387v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S40. Evolutionary Game Design

Cameron Browne and Frédéric Maire. Evolutionary Game Design. IEEE Transactions on Computational Intelligence and AI in Games 2(1), pages 1 to 16, 2010. Author PDF (https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf). Inspected: quality measurements, human comparisons, Experiment III and conclusions. The filename is not the publication year.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R2; supplied original entry retained in JSON.

Bibliographic record: verified title/authors in primary manuscript; publication metadata with placeholders remains supplied unless shown. Content inspected: abstract and first page.

Relevance: collect pairwise designer comparisons with reasons tied to exact revisions.

Limitations:  preference prediction is specific to the sample and game family. Generating a viable game, predicting a preference and proving historical originality are different achievements. Its learned coefficients are not tower defense quality weights. Do not treat 19 viable games as 19 independently established historical inventions.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S41. GAVEL: Generating Games Via Evolution and Language Models

Graham Todd, Alexander G. Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers and Julian Togelius. GAVEL: Generating Games Via Evolution and Language Models. NeurIPS 2024. arXiv:2407.09388v2 (https://arxiv.org/html/2407.09388v2). Inspected: representation, mutation, evaluation, experiments, results, limitations and references. Code link (https://github.com/gdrtodd/gavel), not executed in this investigation.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R3; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: make small, inspectable changes to a candidate's mechanical relationships, then apply independent checks.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-E11; UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2407.09388v2
https://github.com/gdrtodd/gavel
https://arxiv.org/abs/2407.09388v2

Associated files/code/data: https://github.com/gdrtodd/gavel

##### UG-T02-S42. Procedural Content Generation through Quality Diversity

Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius and Georgios N. Yannakakis. Procedural Content Generation through Quality Diversity. IEEE Conference on Games, 2019. arXiv:1907.04053 (https://arxiv.org/abs/1907.04053), HTML (https://arxiv.org/html/1907.04053), DOI:10.1109/CIG.2019.8848053 (https://doi.org/10.1109/CIG.2019.8848053). Inspected: framework, surveyed applications, challenges and references.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R4; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: maintain a small reviewed collection organized by strategic contribution.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/abs/1907.04053
https://arxiv.org/html/1907.04053
https://doi.org/10.1109/CIG.2019.8848053

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S43. Illuminating search spaces by mapping elites

Jean-Baptiste Mouret and Jeff Clune. Illuminating search spaces by mapping elites. 2015 preprint. arXiv:1504.04909v1 (https://arxiv.org/abs/1504.04909v1). Inspected: algorithm, experimental domains, limitations and references via arXiv HTML. The HTML contains an inconsistent generated date; use the arXiv version record for the 2015 date.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R5; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: use the archive idea first as a review aid.

Limitations:  archive coverage means coverage of the chosen descriptors. It does not establish strategic diversity or historical novelty. A category populated by several different descriptions may still contain the same behavior. A single quality score within each category also remains a design choice.

Linked record: UG-T02-D09

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/abs/1504.04909v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S44. Abandoning Objectives: Evolution through the Search for Novelty Alone

Joel Lehman and Kenneth O. Stanley. Abandoning Objectives: Evolution through the Search for Novelty Alone. Evolutionary Computation 19(2), pages 189 to 223, 2011. Primary-text PDF (https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf). Inspected: motivation, behavioral novelty, experimental tasks and discussion. Found in both R4 and R5.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R6; supplied original entry retained in JSON.

Bibliographic record: verified title/authors in primary manuscript; publication metadata with placeholders remains supplied unless shown. Content inspected: abstract and first page.

Relevance: avoid discarding every unusual unit because it initially loses on damage per cost.

Limitations:  novelty is neither usefulness nor fit. Our output still needs supported mechanics and project constraints. This paper is a reason to inspect the effects of our ranking criteria, not a reason to remove quality checks.

Linked record: UG-T02-D09

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S45. Monte-Carlo Tree Search for Simulation-based Strategy Analysis

Alexander Zook, Brent Harrison and Mark O. Riedl. Monte-Carlo Tree Search for Simulation-based Strategy Analysis. The paper identifies FDG 2015 as its conference; the arXiv deposit is 2019. arXiv:1908.01423v1 (https://arxiv.org/html/1908.01423v1). Inspected: methodology, skill-based metrics, case studies, limitations and references. R1 cites the deposit year.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R7; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: in future runtime evaluations, record purchase sequences, target choices, ability timing and map-position choices.

Limitations:  more search is an approximation of skill, not a validated model of every human skill level. A bot can miss a mechanic's value or exploit a rule in ways players would not. Simulation diagnoses behavior under a policy; it does not establish enjoyment.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/1908.01423v1
https://arxiv.org/abs/1908.01423v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S46. Evaluating Competitive Game Balance with Restricted Play

Alexander Jaffe, Alex Miller, Erik Andersen, Yun-En Liu, Anna Karlin and Zoran Popović. Evaluating Competitive Game Balance with Restricted Play. AIIDE 2012. Author-hosted PDF (https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf). Inspected: framework, example application and conclusions. Reached through R1 → R7 → R8.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R8; supplied original entry retained in JSON.

Bibliographic record: verified title/authors in primary manuscript; publication metadata with placeholders remains supplied unless shown. Content inspected: abstract and first page.

Relevance: compare a legal build with the same build when its distinctive capability is unavailable.

Limitations:  no observed difference can mean irrelevance, weak agents, unsuitable scenarios or insufficient trials. A large difference can indicate excessive power rather than good design. Pair outcome measurements with changed decisions and retained weaknesses.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S47. Automated Game Design via Conceptual Expansion

Matthew Guzdial and Mark Riedl. Automated Game Design via Conceptual Expansion. AIIDE 2018, 14(1), pages 31 to 37. arXiv:1809.02232v1 (https://arxiv.org/html/1809.02232v1). Inspected: representation, expansion, evaluation, discussion and references. Reached through R1.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R9; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: record which relationships change when translating an inspiration.

Limitations:  learned approximations can omit mechanics. The paper does not establish source fidelity, appropriate numerical balance or player preference for character adaptations. Its reconstruction evaluation is particularly important when citing it as evidence for originality.

Linked record: UG-T02-E06

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/1809.02232v1
https://arxiv.org/abs/1809.02232v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S48. Constrained Novelty Search: A Study on Game Content Generation

Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. Constrained Novelty Search: A Study on Game Content Generation. Evolutionary Computation 23(1), pages 101 to 129, 2015. Author PDF (https://antoniosliapis.com/papers/constrained_novelty_search.pdf), DOI:10.1162/EVCO_a_00123 (https://doi.org/10.1162/EVCO_a_00123). Inspected: two-population approach, game-content application and reported qualifications. Found through R4; bibliographic metadata also checked through Crossref.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R10; supplied original entry retained in JSON.

Bibliographic record: verified title/authors in primary manuscript; publication metadata with placeholders remains supplied unless shown. Content inspected: abstract and first page.

Relevance: keep a supported candidate collection and a separate set of mechanic proposals with actionable missing requirements.

Limitations:  our draft/proposal distinction is an organizational adaptation, not an implementation of the paper's algorithm. We should not add dual evolutionary populations merely to store unfinished drafts.

Linked record: UG-T02-D05

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://antoniosliapis.com/papers/constrained_novelty_search.pdf
https://doi.org/10.1162/EVCO_a_00123

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S49. Sentient Sketchbook: Computer-Aided Game Level Authoring

Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. Sentient Sketchbook: Computer-Aided Game Level Authoring. Foundations of Digital Games, 2013, pages 213 to 220. Author PDF (https://antoniosliapis.com/papers/sentient_sketchbook.pdf). Inspected: tool interaction, checks, alternatives and small-scale designer study. Found through R4.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R11; supplied original entry retained in JSON.

Bibliographic record: verified title/authors in primary manuscript; publication metadata with placeholders remains supplied unless shown. Content inspected: abstract and first page.

Relevance: show the difference between a candidate and its parent, including changed paths, reachable areas, placement opportunities and constraint failures.

Limitations:  the study is small and map-specific. Suggestion quality cannot be inferred from validity alone. A system that supplies legal alternatives but repeatedly ignores confirmed intent remains a poor design partner.

Linked record: UG-T02-D05

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://antoniosliapis.com/papers/sentient_sketchbook.pdf

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S50. Generation and Evaluation in the Human Invention Process through the Lens of Game Design

Katherine M. Collins et al. Generation and Evaluation in the Human Invention Process through the Lens of Game Design. CogSci 2025. arXiv:2508.10914v1 (https://arxiv.org/html/2508.10914v1). Inspected: dataset, proposal/evaluation model, results, limitations and references.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R12; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: ask for a concrete scenario and predicted player response alongside each candidate.

Limitations:  this is preliminary cognitive modeling, with access to submitted games rather than the full set of ideas people considered and rejected. It neither proves that humans literally execute the model nor demonstrates expert AI design taste. Translation into a formal language introduces another source of error.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2508.10914v1
https://arxiv.org/abs/2508.10914v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S51. People use fast, goal-directed simulation to reason about novel games

Cedegao E. Zhang, Katherine M. Collins, Lionel Wong, Adrian Weller and Joshua B. Tenenbaum. People use fast, goal-directed simulation to reason about novel games. CogSci 2024. arXiv:2407.14095v1 (https://arxiv.org/html/2407.14095v1). Inspected: model, human judgment task and discussion. Found through R12.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R13; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: separate "can a novice understand the intended decision?" from "does expert play retain strategic depth?" A short scenario walkthrough can support the first question; it cannot establish the second.

Limitations:  simple board-game judgments are not evidence that short simulations can predict the appeal of complex real-time tower defense. The transfer requires validation with our players and content.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2407.14095v1
https://arxiv.org/abs/2407.14095v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S52. Mathematical discoveries from program search with large language models

Bernardino Romera-Paredes et al. Mathematical discoveries from program search with large language models. Nature 625, pages 468 to 475, 2024; published online in December 2023. DOI:10.1038/s41586-023-06924-6 (https://www.nature.com/articles/s41586-023-06924-6). Inspected: primary article's method and reported applications. Code link (https://github.com/google-deepmind/funsearch), not executed here.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R14; supplied original entry retained in JSON.

Bibliographic record: verified publisher record. Content inspected: relevant article sections.

Relevance: keep generation, executable validation and design review as distinct operations.

Limitations:  mathematical validity and algorithmic performance are more directly evaluable than fun or character fidelity. External code checks only establish the properties they actually test. FunSearch is not evidence that an LLM reviewer can replace playtesting.

Linked record: UG-T02-E11

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://www.nature.com/articles/s41586-023-06924-6
https://github.com/google-deepmind/funsearch

Associated files/code/data: https://github.com/google-deepmind/funsearch

##### UG-T02-S53. Ludax: A GPU-Accelerated Description Language for Board Games

Graham Todd et al. Ludax: A GPU-Accelerated Description Language for Board Games. arXiv:2506.22609 (https://arxiv.org/abs/2506.22609), first submitted in 2025; v2 inspected (https://arxiv.org/html/2506.22609v2), March 25, 2026. Inspected: representation, compilation, supported game class and references. R12 cites the earlier 2025 version; this record deliberately distinguishes it from the newer text inspected.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R15; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: take the representation lesson: executable meaning comes from defined operators and their implementation.

Limitations:  this is not a ready-made real-time tower defense runtime. Its performance claims should not be transferred to our event model, maps or agents. Consider runtime semantics before acceleration or a new language dependency.

Linked record: UG-T02-E11

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label. Current arXiv v2 record says "Ludax: A GPU-Accelerated Domain Specific Language for Board Games"; supplied title says "Description Language". Authors in the verified record are Graham Todd, Alexander G. Padula, Dennis J.N.J. Soemers and Julian Togelius. Retain the supplied v2 date/version separately.

Direct URLs:
https://arxiv.org/abs/2506.22609
https://arxiv.org/html/2506.22609v2

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S54. Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers

Chenglei Si, Diyi Yang and Tatsunori Hashimoto. Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers. ICLR 2025; preprint 2024. arXiv:2409.04109 (https://arxiv.org/abs/2409.04109), v1 inspected (https://arxiv.org/html/2409.04109v1). Inspected: study description and discussion of novelty, diversity and self-evaluation. The inspected evidence is the preprint, not an audit of changes in the final proceedings version.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R16; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: use model critique to find questions and possible defects, then require independent evidence for claims of correctness, usefulness and originality.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/abs/2409.04109
https://arxiv.org/html/2409.04109v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S55. Mortar: Evolving Mechanics for Automatic Game Design

Muhammad U. Nasir, Yuchen Li, Steven James and Julian Togelius. Mortar: Evolving Mechanics for Automatic Game Design. GECCO 2026. Primary PDF (https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf), DOI:10.1145/3795095.3805100 (https://doi.org/10.1145/3795095.3805100). Inspected: method, mechanic contribution measure, human comparisons and related work.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R17; supplied original entry retained in JSON.

Bibliographic record: verified author/institutional publication record. Content inspected: metadata only in this handover.

Relevance: assess a proposed tower in several map, wave and roster contexts.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf
https://doi.org/10.1145/3795095.3805100
https://www.raillab.org/publication/nasir-2026-mortar/

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S56. Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design

Michael Cook, Simon Colton, Azalea Raad and Jeremy Gow. Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design. EvoApplications 2013. Author PDF (https://www.possibilityspace.org/papers/evo13.pdf), DOI:10.1007/978-3-642-37192-9_29 (https://doi.org/10.1007/978-3-642-37192-9_29). Inspected: representation, search, generated mechanics and pilot study. Promoted from the original R9 citation lead.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R18; supplied original entry retained in JSON.

Bibliographic record: verified title/authors in primary manuscript; publication metadata with placeholders remains supplied unless shown. Content inspected: abstract and first page.

Relevance: propose a mechanic together with a scenario that demonstrates what it enables.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://www.possibilityspace.org/papers/evo13.pdf
https://doi.org/10.1007/978-3-642-37192-9_29

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S57. Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules

Johor Jara Gonzalez, Seth Cooper and Matthew Guzdial. Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules. AIIDE 19(1), pages 266 to 275, 2023. arXiv:2309.09476v3 (https://arxiv.org/html/2309.09476v3), DOI:10.1609/aiide.v19i1.27522 (https://doi.org/10.1609/aiide.v19i1.27522). Inspected: environment recreation, evaluator comparison and limitations.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R19; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: distinguish whether an evaluator can learn a new rule from whether it performs well after learning.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2309.09476v3
https://doi.org/10.1609/aiide.v19i1.27522
https://arxiv.org/abs/2309.09476v3

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S58. Pixie: Code-Level Mechanic Generation for Game Designers

Michael Cook. Pixie: Code-Level Mechanic Generation for Game Designers. AIIDE 21(1), pages 206 to 215, 2025. Publisher PDF (https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962), DOI:10.1609/aiide.v21i1.36824 (https://doi.org/10.1609/aiide.v21i1.36824). Access gap: the publisher download failed and an attempted author copy was unavailable. Mortar's related-work discussion independently confirms Pixie as code-level mechanic evolution using a genetic-programming approach.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R20; supplied original entry retained in JSON.

Bibliographic record: verified publisher record. Content inspected: abstract only; original primary-PDF access gap retained.

Relevance: expose a small, explicit mechanic interface with test scenarios rather than arbitrary runtime internals.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-Q05

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962
https://doi.org/10.1609/aiide.v21i1.36824
https://ojs.aaai.org/index.php/AIIDE/article/view/36824

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S59. Automatic Game Design via Mechanic Generation

Alexander Zook and Mark O. Riedl. Automatic Game Design via Mechanic Generation. AAAI 2014. arXiv:1908.01420v1 (https://arxiv.org/html/1908.01420v1), DOI:10.1609/aaai.v28i1.8788 (https://doi.org/10.1609/aaai.v28i1.8788). Inspected: representation, constraint-based generation, planning and demonstration domains. The arXiv deposit is 2019; the paper is from 2014. Promoted from the original R3/R7 citation lead.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R21; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: express preconditions, resource changes and effects so we can ask whether a proposed tactical outcome is reachable.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-E11

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/1908.01420v1
https://doi.org/10.1609/aaai.v28i1.8788
https://arxiv.org/abs/1908.01420v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S60. TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents

Dawei Wang et al. TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents. AAAI 2026. arXiv:2601.05899v1 (https://arxiv.org/html/2601.05899v1), DOI:10.1609/aaai.v40i31.39818 (https://doi.org/10.1609/aaai.v40i31.39818). Inspected: environment, observation formats, evaluation scope and reported agent failures. Code (https://github.com/tb6147877/TowerMind), not executed.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R22; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: inspect its action/observation contract and invalid-action reporting before designing an evaluation adapter.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-E12

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2601.05899v1
https://doi.org/10.1609/aaai.v40i31.39818
https://github.com/tb6147877/TowerMind
https://arxiv.org/abs/2601.05899v1

Associated files/code/data: https://github.com/tb6147877/TowerMind

##### UG-T02-S61. Automatic generation of tower defense levels using PCG

Simon Liu et al. Automatic generation of tower defense levels using PCG. FDG 2019. DOI:10.1145/3337722.3337723 (https://doi.org/10.1145/3337722.3337723), supplied manuscript entry (https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG). Access gap: ResearchGate returned HTTP 403; the primary text was not independently inspected.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R23; supplied original entry retained in JSON.

Bibliographic record: unverified in this handover; supplied DOI retained after access failure. Content inspected: not inspected.

Relevance: evaluate paths, tower opportunities, waves and starting money jointly.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://doi.org/10.1145/3337722.3337723
https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S62. Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms

Vid Kraner, Iztok Fister Jr. and Lucija Brezočnik. Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms. NT 2021, Lecture Notes in Networks and Systems 233. DOI:10.1007/978-3-030-75275-0_54 (https://doi.org/10.1007/978-3-030-75275-0_54), supplied manuscript entry (https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms). Access gap: ResearchGate returned HTTP 403.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R24; supplied original entry retained in JSON.

Bibliographic record: unverified in this handover; supplied DOI retained after access failure. Content inspected: not inspected.

Relevance: separate generated arrangements from the human-authored rules and assets that enable them.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://doi.org/10.1007/978-3-030-75275-0_54
https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S63. Computational Intelligence and Tower Defence Games

Phillipa Avery, Julian Togelius, Elvis Alistar and Robert Pieter van Leeuwen. Computational Intelligence and Tower Defence Games. IEEE Congress on Evolutionary Computation, 2011. DOI:10.1109/CEC.2011.5949738 (https://doi.org/10.1109/CEC.2011.5949738), supplied manuscript entry (https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games). Access gap: ResearchGate returned HTTP 403.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R25; supplied original entry retained in JSON.

Bibliographic record: unverified in this handover; supplied DOI retained after access failure. Content inspected: not inspected.

Relevance: Background or research lead retained from the supplied catalogue; not an approved project component..

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://doi.org/10.1109/CEC.2011.5949738
https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S64. Controllable procedural map generation via multiobjective evolution

Julian Togelius et al. Controllable procedural map generation via multiobjective evolution. Genetic Programming and Evolvable Machines 14, pages 245 to 277, 2013. Publisher article and DOI (https://link.springer.com/article/10.1007/s10710-012-9174-5). Abstract and reference list inspected; the full article requires access.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R26; supplied original entry retained in JSON.

Bibliographic record: verified publisher record. Content inspected: abstract and reference list; full article inaccessible.

Relevance: retain map tradeoffs instead of assuming defensibility, open movement and access to resources improve together.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label. Publisher author list: Julian Togelius, Mike Preuss, Nicola Beume, Simon Wessing, Johan Hagelbäck, Georgios N. Yannakakis and Corrado Grappiolo.

Direct URLs:
https://link.springer.com/article/10.1007/s10710-012-9174-5

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S65. Learning Controllable 3D Level Generators

Zehua Jiang, Sam Earle, Michael Cerny Green and Julian Togelius. Learning Controllable 3D Level Generators. FDG 2022. arXiv:2206.13623v3 (https://arxiv.org/html/2206.13623v3), DOI:10.1145/3555858.3563273 (https://doi.org/10.1145/3555858.3563273). Inspected: PCGRL formulation, 3D movement requirements and control experiments.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R27; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: evaluate geometry under the actual movement model.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2206.13623v3
https://doi.org/10.1145/3555858.3563273
https://arxiv.org/abs/2206.13623v3

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S66. Scaling, Control and Generalization in Reinforcement Learning Level Generators

Sam Earle, Zehua Jiang and Julian Togelius. Scaling, Control and Generalization in Reinforcement Learning Level Generators. 2024. arXiv:2408.12525v1 (https://arxiv.org/html/2408.12525v1). Inspected: method, scaling and generalization scope. The supplied list calls this "PCGRL+"; the citation here uses the manuscript title.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R28; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: preserve fixed designer constraints and test across map sizes if a learned map generator becomes useful.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label. Bibliographic correction: arXiv v1 title begins "PCGRL+: Scaling, Control and Generalization in Reinforcement Learning Level Generators". The supplied catalogue drops "PCGRL+:" and describes it as an informal label. That assertion is contradicted by the v1 title. The arXiv record also states publication at IEEE Conference on Games, 2024; final proceedings text was not inspected.

Direct URLs:
https://arxiv.org/html/2408.12525v1
https://arxiv.org/abs/2408.12525v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S67. A Database-Driven Framework for 3D Level Generation with LLMs

Kaijie Xu and Clark Verbrugge. A Database-Driven Framework for 3D Level Generation with LLMs. AIIDE 21(1), pages 367 to 376, 2025. arXiv:2508.18533v1 (https://arxiv.org/html/2508.18533v1), DOI:10.1609/aiide.v21i1.36840 (https://doi.org/10.1609/aiide.v21i1.36840). Inspected: assembly pipeline, repair and limitations.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R29; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: specify logical connectivity and progression before realizing them as geometry.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2508.18533v1
https://doi.org/10.1609/aiide.v21i1.36840
https://arxiv.org/abs/2508.18533v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S68. Infinite Photorealistic Worlds using Procedural Generation

Alexander Raistrick et al. Infinite Photorealistic Worlds using Procedural Generation. CVPR 2023. arXiv:2306.09310 (https://arxiv.org/abs/2306.09310). Primary HTML inspected for system purpose and generation scope.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R30; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: treat visual generation as a consumer of gameplay constraints.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/abs/2306.09310

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S69. Holodeck: Language Guided Generation of 3D Embodied AI Environments

Yue Yang et al. Holodeck: Language Guided Generation of 3D Embodied AI Environments. CVPR 2024. arXiv:2312.09067 (https://arxiv.org/abs/2312.09067). Primary HTML inspected for planning, retrieval, spatial constraints and evaluation scope.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R31; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: use explicit spatial relationships when placing decorative and functional assets.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/abs/2312.09067

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S70. Automatic Content Generation in the Galactic Arms Race Video Game

Erin J. Hastings, Ratan K. Guha and Kenneth O. Stanley. Automatic Content Generation in the Galactic Arms Race Video Game. IEEE Transactions on Computational Intelligence and AI in Games 1(4), pages 245 to 263, 2009. DOI:10.1109/TCIAIG.2009.2038365 (https://doi.org/10.1109/TCIAIG.2009.2038365), supplied manuscript entry (https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game). Access gap: ResearchGate returned HTTP 403.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R32; supplied original entry retained in JSON.

Bibliographic record: unverified in this handover; supplied DOI retained after access failure. Content inspected: not inspected.

Relevance: use player behavior as one feedback source, with context.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://doi.org/10.1109/TCIAIG.2009.2038365
https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S71. OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code

Maxence Faldor, Jenny Zhang, Antoine Cully and Jeff Clune. OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code. ICLR 2025. arXiv:2405.15568v3 (https://arxiv.org/html/2405.15568v3). Inspected: archive, task/environment generation, interestingness filtering and evaluation modes.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R33; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: retain both failed and successful candidate histories and use them to propose a new testable decision.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2405.15568v3
https://arxiv.org/abs/2405.15568v3

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S72. DreamGarden: A Designer Assistant for Growing Games from a Single Prompt

Sam Earle, Samyak Parajuli and Andrzej Banburski-Fahey. DreamGarden: A Designer Assistant for Growing Games from a Single Prompt. CHI 2025. arXiv:2410.01791v2 (https://arxiv.org/html/2410.01791v2), DOI:10.1145/3706598.3714233 (https://doi.org/10.1145/3706598.3714233). Inspected: hierarchical planning, Unreal implementation modules and usability study.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R34; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: keep plans, confirmed choices and intermediate artifacts inspectable.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2410.01791v2
https://doi.org/10.1145/3706598.3714233
https://arxiv.org/abs/2410.01791v2

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S73. AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems

Lei Yin et al. AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems. 2026 preprint. arXiv:2603.07106v1 (https://arxiv.org/html/2603.07106v1). Inspected: engine workflow, documentation retrieval, runtime tests and benchmark. Code (https://github.com/Pluto156/AutoUE), not executed.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R35; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: pair engine documentation with runtime evidence when implementing a Handoff.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2603.07106v1
https://github.com/Pluto156/AutoUE
https://arxiv.org/abs/2603.07106v1

Associated files/code/data: https://github.com/Pluto156/AutoUE

##### UG-T02-S74. Genie: Generative Interactive Environments

Jake Bruce et al. Genie: Generative Interactive Environments. ICML 2024. arXiv:2402.15391 (https://arxiv.org/abs/2402.15391). Primary HTML inspected for model and output representation.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R36; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: keep it as a separate reference for interactive visual exploration.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/abs/2402.15391

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S75. Emergence of player tactics by expert-guided machine learning: an industry tower defence case study

Stuart Anderson and Ruth Falconer. Emergence of player tactics by expert-guided machine learning: an industry tower defence case study. Entertainment Computing 54, article 100963, 2025. Institutional record (https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/), DOI:10.1016/j.entcom.2025.100963 (https://doi.org/10.1016/j.entcom.2025.100963). Institutional abstract and metadata inspected; the linked full-text download returned HTTP 403.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R37; supplied original entry retained in JSON.

Bibliographic record: verified institutional record. Content inspected: abstract only; full-text gap retained.

Relevance: record contextual action sequences that let us identify tactics.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-D09; later research only

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/
https://doi.org/10.1016/j.entcom.2025.100963

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S76. GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development

Xiuhui Zhang et al. GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development. Preprint, September 18, 2026. arXiv:2609.21293v1 (https://arxiv.org/html/2609.21293v1). Inspected: evaluation contract, task suite and aggregate results.

Type: Research paper or literature lead; see preserved bibliography. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, R38; supplied original entry retained in JSON.

Bibliographic record: verified arXiv bibliographic record; supplied venue/year may require separate proceedings check. Content inspected: abstract only.

Relevance: specify acceptance scenarios independently of generated implementations and report both check-level results and complete-task success.

Limitations: This is not a Unit-Generator implementation or evidence of complete character-to-TD adaptation. See the original catalogue entry for attributed limits and access status.

Linked record: UG-T02-E11

Verification note: The original catalogue reports its own inspection status in the preserved bibliographic field. That status is not transferred to the current inspection label.

Direct URLs:
https://arxiv.org/html/2609.21293v1
https://arxiv.org/abs/2609.21293v1

Associated files/code/data: No implementation or dataset was executed, benchmarked or license-audited in this handover.

##### UG-T02-S77. Experience-Driven Procedural Content Generation

Georgios N. Yannakakis and Julian Togelius. IEEE Transactions on Affective Computing 2(3), 2011. DOI:10.1109/T-AFFC.2011.6.

Type: Existing research/access-gap lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, Additional references and access gaps.

Bibliographic record: unverified independently; DOI access failed in handover. Content inspected: not inspected.

Relevance: Player-specific evaluation framing.

Limitations: Original review also reports access gaps. Retained as leads, not supporting verified empirical claims.

Linked record: UG-T02-E12; UG-T02-Q09

Direct URLs:
https://doi.org/10.1109/T-AFFC.2011.6

##### UG-T02-S78. Generative AI enhances individual creativity but reduces the collective diversity of novel content

Anil R. Doshi and Oliver P. Hauser. Science Advances 10(28), 2024. DOI:10.1126/sciadv.adn5290.

Type: Existing research/access-gap lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, Additional references and access gaps.

Bibliographic record: unverified independently; DOI access failed in handover. Content inspected: not inspected.

Relevance: Check roster-level similarity; supplied claims concern short fiction, not game mechanics.

Limitations: Original review also reports access gaps. Retained as leads, not supporting verified empirical claims.

Linked record: UG-T02-E12; UG-T02-Q09

Direct URLs:
https://doi.org/10.1126/sciadv.adn5290

##### UG-T02-S79. Large language models are homogeneously creative

Emily Wenger and Yoed N. Kenett. PNAS Nexus 5(3), 2026. DOI:10.1093/pnasnexus/pgag042.

Type: Existing research/access-gap lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, Additional references and access gaps.

Bibliographic record: unverified independently; DOI access failed in handover. Content inspected: not inspected.

Relevance: Cross-model diversity may require evaluation rather than assuming provider independence.

Limitations: Original review also reports access gaps. Retained as leads, not supporting verified empirical claims.

Linked record: UG-T02-E12; UG-T02-Q09

Direct URLs:
https://doi.org/10.1093/pnasnexus/pgag042

##### UG-T02-S80. Superhuman artificial intelligence can improve human decision-making by increasing novelty

Minkyu Shin, Jin Kim, Bas van Opheusden and Thomas L. Griffiths. PNAS 120(12), 2023. DOI:10.1073/pnas.2214840120.

Type: Existing research/access-gap lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, Additional references and access gaps.

Bibliographic record: unverified independently; DOI access failed in handover. Content inspected: not inspected.

Relevance: AI influence on play within existing rules is not direct rule invention.

Limitations: Original review also reports access gaps. Retained as leads, not supporting verified empirical claims.

Linked record: UG-T02-E12; UG-T02-Q09

Direct URLs:
https://doi.org/10.1073/pnas.2214840120

##### UG-T02-S81. Creativity and artificial intelligence

Margaret A. Boden. Artificial Intelligence 103(1-2), 1998. DOI:10.1016/S0004-3702(98)00055-1.

Type: Existing research/access-gap lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, Additional references and access gaps.

Bibliographic record: unverified independently; DOI access failed in handover. Content inspected: not inspected.

Relevance: Conceptual creativity vocabulary; not empirical LLM performance.

Limitations: Original review also reports access gaps. Retained as leads, not supporting verified empirical claims.

Linked record: UG-T02-E12; UG-T02-Q09

Direct URLs:
https://doi.org/10.1016/S0004-3702(98)00055-1

##### UG-T02-S82. Program Synthesis as a Generative Method

Eric Butler, Kristin Siu and Alexander Zook. FDG 2017. DOI:10.1145/3102071.3102076, recovered from publisher search record.

Type: Existing citation-trail lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, additional leads; cited by R1.

Bibliographic record: unverified complete record; partial publisher search match. Content inspected: metadata only; search record, full publisher page failed.

Relevance: Generating behavior rather than numerical attributes.

Limitations: No underlying experiment inspected; do not infer solver capabilities from title.

Linked record: UG-T02-E06

Direct URLs:
https://dl.acm.org/doi/10.1145/3102071.3102076

##### UG-T02-S83. Talakat: Bullet Hell Generation through Constrained Map-Elites

Ahmed Khalifa, Scott Lee, Andy Nealen and Julian Togelius. GECCO 2018. arXiv:1806.04718; DOI:10.1145/3205455.3205470.

Type: Existing citation-trail lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, additional leads; R4 reference 18.

Bibliographic record: verified arXiv bibliographic record. Content inspected: abstract only.

Relevance: Executable attack-pattern diversity with constraints.

Limitations: Bullet-hell level generation, not source-grounded TD units.

Linked record: UG-T02-D09

Direct URLs:
https://arxiv.org/abs/1806.04718
https://dl.acm.org/doi/10.1145/3205455.3205470

##### UG-T02-S84. Mapping Hearthstone Deck Spaces through MAP-Elites with Sliding Boundaries

Matthew C. Fontaine, Scott Lee, L. B. Soros, Fernando De Mesentier Silva, Julian Togelius and Amy K. Hoover. GECCO 2019. arXiv:1904.10656; DOI:10.1145/3321707.3321794.

Type: Existing citation-trail lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, additional leads; R4 reference 13.

Bibliographic record: verified arXiv bibliographic record. Content inspected: abstract only.

Relevance: Interactions among content and choice of diversity descriptors.

Limitations: Deck search, not character-to-mechanic adaptation.

Linked record: UG-T02-D09

Verification note: Original catalogue says "Michael C. Fontaine" and "with MAP-Elites". Verified record says "Matthew C. Fontaine" and "through MAP-Elites". Original wording is preserved here explicitly rather than silently corrected.

Direct URLs:
https://arxiv.org/abs/1904.10656
https://dl.acm.org/doi/10.1145/3321707.3321794

##### UG-T02-S85. Multiobjective Exploration of the StarCraft Map Space

Julian Togelius, Mike Preuss, Nicola Beume, Simon Wessing, Johan Hagelbäck and Georgios N. Yannakakis. IEEE CIG 2010, 265-272. DOI:10.1109/ITW.2010.5593346, institutional record.

Type: Existing citation-trail lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, additional leads; R1 references.

Bibliographic record: verified institutional bibliographic record. Content inspected: metadata only in handover.

Relevance: Conflicting map objectives.

Limitations: No primary experiment reviewed here and not a unit generator.

Linked record: UG-T02-D09

Direct URLs:
https://julian.togelius.com/Togelius2010Multiobjective.pdf
https://www.wi.uni-muenster.de/research/publications/87283

##### UG-T02-S86. Mixed-Initiative Co-Creativity

Georgios N. Yannakakis, Antonios Liapis and Constantine Alexopoulos. FDG 2014.

Type: Existing citation-trail lead. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md, additional leads; R3 reference 68.

Bibliographic record: verified institutional bibliographic record. Content inspected: metadata only in handover.

Relevance: Designer intent, interaction and revision control.

Limitations: No current full-text or usability-result audit.

Linked record: UG-T02-D05

Direct URLs:
https://www.um.edu.mt/library/oar/handle/123456789/29459
https://www.fdg2014.org/papers/fdg2014_paper_37.pdf

##### UG-T02-S87. Unit-Generator product requirements

Repository-local reference; date and authors unavailable.

Type: Unavailable internal document. Origin: Already present in this thread.

Original locator: CONTEXT.md and/or GAME-DESIGN-RESEARCH.md link or mention: unit-generator/docs/PRODUCT.md; CONTEXT.md relative link docs/PRODUCT.md.

Bibliographic record: unverified; missing material. Content inspected: not inspected.

Relevance: Referenced product contract, not available.

Limitations: Cannot substantiate implementation, license decisions, private assets or other thread decisions.

Linked record: UG-T02-Q08

Verification note: Do not replace this private source with web search or infer contents from its title.

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S88. API compatibility

Repository-local reference; date and authors unavailable.

Type: Unavailable internal document. Origin: Already present in this thread.

Original locator: CONTEXT.md and/or GAME-DESIGN-RESEARCH.md link or mention: docs/API.md.

Bibliographic record: unverified; missing material. Content inspected: not inspected.

Relevance: Referenced implementation/API detail, not available.

Limitations: Cannot substantiate implementation, license decisions, private assets or other thread decisions.

Linked record: UG-T02-Q08

Verification note: Do not replace this private source with web search or infer contents from its title.

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S89. Archived September 21 generator survey

Repository-local reference; date and authors unavailable.

Type: Unavailable internal document. Origin: Already present in this thread.

Original locator: CONTEXT.md and/or GAME-DESIGN-RESEARCH.md link or mention: unit-generator/research/archive/GENERATOR-SURVEY-2026-09-21.md.

Bibliographic record: unverified; missing material. Content inspected: not inspected.

Relevance: Older card/monster examples and license notes mentioned, not inspected.

Limitations: Cannot substantiate implementation, license decisions, private assets or other thread decisions.

Linked record: UG-T02-Q08

Verification note: Do not replace this private source with web search or infer contents from its title.

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S90. Generator references

Repository-local reference; date and authors unavailable.

Type: Unavailable internal document. Origin: Already present in this thread.

Original locator: CONTEXT.md and/or GAME-DESIGN-RESEARCH.md link or mention: GENERATOR-REFERENCES.md; fragment #framework-and-character-sources.

Bibliographic record: unverified; missing material. Content inspected: not inspected.

Relevance: Shared technical/data/source index, not available.

Limitations: Cannot substantiate implementation, license decisions, private assets or other thread decisions.

Linked record: UG-T02-Q08

Verification note: Do not replace this private source with web search or infer contents from its title.

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S91. Framework sources

Repository-local reference; date and authors unavailable.

Type: Unavailable internal document. Origin: Already present in this thread.

Original locator: CONTEXT.md and/or GAME-DESIGN-RESEARCH.md link or mention: unit-generator/research/FRAMEWORK-SOURCES.md.

Bibliographic record: unverified; missing material. Content inspected: not inspected.

Relevance: Repository-local applications and limitations, not available.

Limitations: Cannot substantiate implementation, license decisions, private assets or other thread decisions.

Linked record: UG-T02-Q08

Verification note: Do not replace this private source with web search or infer contents from its title.

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S92. Towerright context

Repository-local reference; date and authors unavailable.

Type: Unavailable internal document. Origin: Already present in this thread.

Original locator: CONTEXT.md and/or GAME-DESIGN-RESEARCH.md link or mention: towerright/CONTEXT.md.

Bibliographic record: unverified; missing material. Content inspected: not inspected.

Relevance: Private-platform context referenced by catalogue, not available.

Limitations: Cannot substantiate implementation, license decisions, private assets or other thread decisions.

Linked record: UG-T02-Q08

Verification note: Do not replace this private source with web search or infer contents from its title.

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S93. Foundation

Repository-local reference; date and authors unavailable.

Type: Unavailable internal document. Origin: Already present in this thread.

Original locator: CONTEXT.md and/or GAME-DESIGN-RESEARCH.md link or mention: No path or URL supplied.

Bibliographic record: unverified; missing material. Content inspected: not inspected.

Relevance: CONTEXT.md identifies Foundation as private internal guidance unnecessary for contributors.

Limitations: Cannot substantiate implementation, license decisions, private assets or other thread decisions.

Linked record: UG-T02-Q08

Verification note: Do not replace this private source with web search or infer contents from its title.

Direct URL: Not supplied or not applicable. Use the message or file locator above.

##### UG-T02-S94. Mardwerk Map Generator repository

Mardwerk repository named in CONTEXT.md.

Type: Related-repository pointer. Origin: Already present in this thread.

Original locator: CONTEXT.md, Related repositories.

Bibliographic record: unverified repository contents. Content inspected: not inspected.

Relevance: Ecosystem pointer only.

Limitations: No dependency or integration requirement follows.

Linked record: UG-T02-Q08

Direct URLs:
https://github.com/mardwerk/map-generator

##### UG-T02-S95. Mardwerk organization metadata

Mardwerk .github repository named in CONTEXT.md.

Type: Related-repository pointer. Origin: Already present in this thread.

Original locator: CONTEXT.md, Related repositories.

Bibliographic record: unverified repository contents. Content inspected: not inspected.

Relevance: Organization pointer only.

Limitations: No license or architecture claim inferred.

Linked record: UG-T02-Q08

Direct URLs:
https://github.com/mardwerk/.github

##### UG-T02-S99. Unresolved generic research trails

Original phrases: "conceptual blending work"; "play-trace similarity, quantitative balance, game metrics without players and MCTS methods"; "interactive dungeon design".

Type: Incomplete literature references. Origin: Already present in this thread.

Original locator: GAME-DESIGN-RESEARCH.md R9, R7 and R4 respectively.

Bibliographic record: unverified; no uniquely identifiable paper supplied. Content inspected: not inspected.

Relevance: Retains incomplete references without inventing titles or authors.

Limitations: Cannot be cited as empirical support. Recover exact references from parent papers in later review.

Linked record: UG-T02-S42; UG-T02-S45; UG-T02-S47

Direct URL: Not supplied or not applicable. Use the message or file locator above.

#### B. Newly identified literature leads

##### UG-T02-S96. Grammar-Aligned Decoding

Kanghee Park, Jiayu Wang, Taylor Berg-Kirkpatrick, Nadia Polikarpova and Loris D'Antoni. NeurIPS 2024; arXiv:2405.21047; relevant v1 text inspected.

Type: Published research paper. Origin: Newly found, not part of the original conversation.

Original locator: New handover literature search.

Bibliographic record: verified arXiv and proceedings record. Content inspected: relevant sections: constrained distributions, algorithm and evaluation framing.

Relevance: Challenges treating constrained decoding as a neutral quality-preserving wrapper.

Limitations: Distributional and task-quality results are not human design taste or TD balance results. Does not require implementing ASAp for v1.

Linked record: UG-T02-E13; UG-T02-Q09

Direct URLs:
https://arxiv.org/abs/2405.21047
https://arxiv.org/html/2405.21047v1
https://openreview.net/forum?id=WAidyyRaes

Associated files/code/data: https://github.com/ebmoon/transformers-GAD ; not executed.

##### UG-T02-S97. Large Language Models Cannot Self-Correct Reasoning Yet

Jie Huang, Xinyun Chen, Swaroop Mishra, Huaixiu Steven Zheng, Adams Wei Yu, Xinying Song and Denny Zhou. arXiv:2310.01798, 2023; ICLR 2024 record located.

Type: Research paper. Origin: Newly found, not part of the original conversation.

Original locator: New handover literature search.

Bibliographic record: verified arXiv record; OpenReview full page blocked in this handover. Content inspected: relevant sections: intrinsic versus external feedback, comparison guidance and limitations.

Relevance: Tests the assumption that adding model critique necessarily corrects errors; motivates external findings and cost-matched baselines.

Limitations: Older models and reasoning tasks, specifically intrinsic correction without external feedback. Not a universal impossibility claim and not a test of the proposed validator-guided repair.

Linked record: UG-T02-E14; UG-T02-D05

Direct URLs:
https://arxiv.org/abs/2310.01798
https://arxiv.org/html/2310.01798v1
https://openreview.net/forum?id=IkmD3fKBPQ

#### New local verification source

##### UG-T02-S98. UG-T02 handover artifact audit

New local inspection and rerun for this consolidation, September 24, 2026.

Type: New verification artifact. Origin: New verification during consolidation.

Original locator: audit/audit_findings.json; audit/reproduce_audit.py.

Bibliographic record: verified inspectable results. Content inspected: full code inspection and local execution.

Relevance: Reproduces saved reference results and documents two scope probes.

Limitations: Tests the fixture checker, not a game or generator. Does not retroactively turn assistant design proposals into accepted decisions.

Linked record: UG-T02-E01; UG-T02-E02; UG-T02-E03

Direct URL: Not supplied or not applicable. Use the message or file locator above.

Associated files/code/data: audit_findings.json; rerun outputs; reproduction script; hash manifest

#### Literature gaps for later review

The closest verified precedents cover semantic representation, description-to-code generation, card design, structured decoding and evaluation separately. They do not establish the whole character-to-constrained-unit-family workflow. A later review should examine adaptation datasets, source-faithfulness judgments, constraint-preserving preference training, crosspath interaction coverage, evaluator competence, and public/private reproducibility. The incomplete references and primary-access gaps above remain open. This extraction is not an exhaustive literature review or a novelty determination.


## 8. Preparation for later research

### Existing plans in the thread

The assistant proposed reviewed transformation examples, complete requests including rule revisions and source scope, and a comparison among a template/retrieval baseline, direct pretrained generation, a structured generation/checking system, and that system after fine-tuning. It proposed separate measures for completeness/legality, source fidelity, mechanical distinction, gameplay usefulness, human correction effort and cost. It also proposed character-level held-out splits, anonymized/counterfactual inputs, and later mechanic-removal tests with capable policies. None was implemented or approved as a v1 gate. [UG-T02-D10; UG-T02-E10; UG-T02-S07, Proposed adaptations]

The thread suggests the research question: **Does explicit character interpretation plus a constrained unit-family representation improve faithful, useful adaptation compared with direct LLM generation, and does learned specialization add further benefit?** This is a reconstruction of the thread's research direction, not a registered study.

Potential contributions include a source-to-mechanic adaptation dataset, a full-family representation with scoped checks, or evidence about the contribution of training versus constraints and examples. None is asserted to be novel. Existing local evidence establishes a reproducible structural fixture only. External papers support separate techniques; generator superiority, source fidelity, strategic value and balance remain hypotheses.

### New suggestions for the next evaluation

The smallest useful next comparison is an actual source-to-draft run through the public prototype against a direct-LLM-plus-wiki baseline. Use cases outside the examples used to construct the prompt or fixture. This handover does not specify an unapproved sample size or success threshold.

Do not handicap the baseline. For a controlled comparison, give both systems the same character evidence, scope, Game Definition, design brief, reference prices and requested deliverable. Use the same model/version where possible and let the direct baseline receive a well-written prompt. Account for all calls, retries, repair tokens, source acquisition and human intervention. If the practical baseline is literally wiki-only, report that separately: it tests the value of extra game information as well as the method. A more expensive, different-model or privately tuned system needs its differences stated rather than attributing the entire improvement to the representation.

Proposed, unapproved measures are complete-family constraint success, preservation of source-defined signature behavior, unsupported or invented capabilities, completeness of crosspath explanations, reviewer correction time and edits, blind designer usefulness judgments, and total run cost/latency. Diversity should concern different supported behaviors or purchasing reasons, not cosmetic names or embedding distance alone. No gameplay-balance measure should be reported before a named runtime and scenarios actually execute the unit.

Later comparisons can isolate structured checking, public examples and model-weight tuning separately. A valid but generic reskin should not count as faithful adaptation merely because it passes a schema. A readable novel description should not count as a new implemented mechanic.

Retain code and Definition revisions, model/provider identifiers, adapters or checkpoint hashes where available, prompts and inference settings, original evidence packets, drafts and repairs, parsed results, check scopes, failures, reviewer criteria, recorded edits, seeds and resource use. Keep tuning characters and closely related variants separate from evaluation cases. Preserve unsuccessful results as well as selected examples.

Report public-default results separately from Towerright-private results. Private tuning can remain private, but the evaluation must identify which result depends on it and what an external researcher can reproduce. A public checkpoint or prompt change should not be silently combined with private examples, review or platform services and labeled the public default. [UG-T02-S00; UG-T02-Q08; UG-T02-Q09]

## 9. Consolidation handoff

Do not lose the distinction between semantic adaptation and reskinning, between a whole upgrade family and isolated upgrades, or between declared mechanics, static validation and runtime evidence. Preserve the Wizard fixture and its reproducible results together with the limits in UG-T02-E01 to UG-T02-E03.

Owner approval is still needed for the deliverable boundary, the BTD6-derived public default, the amount of fixed behavior, the model/deployment path and specific private tuned assets. These are UG-T02-Q01 to UG-T02-Q08, not settled decisions. Wizard selection and model training remain UG-T02-D06 and UG-T02-D02 proposals.

The next concrete action supported by this record is to approve one public draft contract and run a real character adaptation through it, using Wizard as a representation test rather than a mandatory behavior template. Retain the result, scoped findings and matched direct-LLM output under UG-T02-Q09 before committing to model-weight training or expanding the reference catalog.
