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
