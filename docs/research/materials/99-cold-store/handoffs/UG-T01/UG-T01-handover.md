# Generic generator research: UG-T01-handover

Selected source excerpts retain their original identifiers, quotations, inspection status, and proposal labels. This split grants no publication or reuse rights. The private companion contains game-specific proposals that remain in Planning outside current game scope. Original source: `planning/99-cold-store/handoffs/UG-T01/UG-T01-handover.md`.

# UG-T01: Character-based unit generation, replaceable rules, and the concept-first prototype

Internal consolidation record for Mardwerk's Unit-Generator. Prepared September 24, 2026.

Classification: internal working record. Contains private design preferences and links to an owner-only attachment. Do not copy this handover wholesale into the public repository or a paper without disclosure review.

## 1. Thread overview and coverage

**Thread ID:** UG-T01. **Descriptive title:** From AI taste research to a configurable, skill-led tower-defense unit prototype.

**Original question.** The owner asked for actual research on AI taste, novel solutions, and unit/game design, then asked how to adapt documented anime abilities into a tower-defense unit with BTD6-like progression. The conversation eventually became an implementation and scope discussion: how much judgment belongs in a skill, how much control belongs in replaceable rules, and how to preserve useful designs in a standalone tool.

**Contribution.** This thread supplies a detailed design skill, a character-scoped prompt, one owner-approved-as-promising generated concept, a private contrasting interpretation, formal progression examples, research leads, successive repository-review reports, a reported implementation update, and a proposed vocabulary. It does not demonstrate repeatable superiority over direct LLM prompting, runtime correctness, balance, or end-to-end replacement of the game's rules.

**Origin convention.** "Owner statement" and "assistant proposal" refer to the visible historical conversation. "Current context" refers only to the product context supplied with this consolidation request. "Inference" identifies our interpretation. "Current recommendation" and "new verification" identify work performed for this handover. The current product context governs the recommendation, but is not retroactively entered as a historical decision.

**Material actually reviewed.** The visible original user and assistant messages were reviewed, together with all five available attachments: the previous implementation handoff, normalized prototype baseline, private owner notes, broad research catalogue, and proposed CONTEXT.md (UG-T01-S84 through S88). The owner's actual CONTEXT.md is present verbatim in a message (UG-T01-S89). The source catalogue's entire 38-entry numbered collection and additional leads were retained. Selected primary bibliographic records and relevant text were checked afresh for section 7.

**Coverage limitations.** Many older tool responses are redacted. Earlier assistant reports of source or repository inspection are therefore reports, not independently recoverable audits. Several citations are malformed or opaque. Original ChatGPT generation logs, model configuration, repository test output, retained failure artifacts, the bytes of commit `739cae3`, and the standalone `PR10_REVIEW` file are not available here. No repository tests, models, simulations, or human studies were executed during this consolidation. Other Project conversation summaries are visible as context, but were not treated as original UG-T01 messages or mined for extra decisions. Attachment dates are file dates, not invented dates for the original messages. References to source text being "inspected" in the catalogue describe that document's earlier investigation; they are not automatically this handover's verification status.

### Central idea

A character reference supplies recognizable actions and relationships, while a game's rules determine how a unit can progress and behave. The generator should adapt the character within those rules without reducing every ability to damage, speed, or range. The owner supplied a skill-only Luffy concept that was nearly useful enough for a public default, but subsequent application runs reportedly produced no accepted designs. The practical problem is therefore to reproduce and preserve useful conceptual design while making the governing rules genuinely replaceable. Code can check explicit constraints; it cannot by itself establish character fidelity, taste, novelty, or balance. A functioning prototype comes before a publication-ready study.

## 2. Ideas, approaches, and changes of direction

### Main approaches explored

**Character-to-mechanics translation.** Research and a source dossier would distinguish canon facts from invented game applications, then map ability relationships into progression. Intended benefit: stronger source fit than attaching names to generic statistics. Costs: reference scope, ambiguity, and interpretation quality remain real problems. Retained as a design direction, not demonstrated as an automatic research capability. Locator: owner, "we would need to 'transform' rather then come up with fully novel characters"; assistant, "character-to-mechanics translation system." Related: UG-T01-D01, D05, E01, S01, S03, S09.

**Executable, search-and-simulation-first generation.** Early assistant proposals combined structured attacks, whole-tree generation, numerical prices, search, and simulated play. They aimed to verify behavior and balance rather than trust descriptions. They required a runtime, reference economy, capable agents, and a sufficiently broad representation. The assistant later withdrew these as prerequisites for the concept prototype. The owner did not approve the early architecture as a whole. Numerical authoring remained a possible separate route. Related: UG-T01-D11, E05, S85 sections 4 and 7.

**Generic combat-role Luffy.** Jet Rhythm, Giant Impact, and Elastic Geometry divided the unit into cadence, impact, and trajectory behavior. The owner rejected this result as "tasteless and basic" and lacking a consistent figure. This rejection concerns that interpretation, not a universal ban on attack-geometry design. Related: UG-T01-D05.


**A long, reusable design skill.** The owner supplied a detailed concept-design skill, a Wano-scoped character brief, and later the generated unit. It governs causal adaptation, progression, pierce, crosspaths, support boundaries, and prose preservation. Intended benefit: usable concepts without a universal executable language. Costs: instruction-following and self-review reliability are unmeasured. The owner explicitly called it a prototype, then judged one output nearly sufficient for the public default. Related: UG-T01-D07, D08, E01, E02.

**Fixed workflow with bounded agent work.** The assistant proposed a controller that fixes rules, required checks, and budgets while the model interprets the character, researches missing facts, and revises defects. Separate path agents and a manager-agent hierarchy were discouraged as initial complexity. This remained advice, not an owner-approved final architecture. Recommendations also changed from multiple mandatory layouts and repairs to a simpler single-draft baseline. Related: UG-T01-D12, Q12.

**Replaceable rule packages plus a design interpretation.** Proposed RulePack/DesignPlan and later Definition/Profile distinctions separate game law from source-to-system mapping. Intended benefit: one generator can support the public default and external game-specific rules. Dependency: replacement must reach prompts, checks, revision, rendering, and retained artifacts, not only a loader. The owner accepted the replaceability goal and later reported optional implementation work; the exact abstraction and end-to-end result remain unverified. Related: UG-T01-D02, D13, D14, Q01.

**Compact recipe generation.** PR #6 was described in earlier assistant reviews as a constrained, code-selected generation experiment. It could improve controllability within a smaller vocabulary, at the risk of narrowing character interpretation. The owner later reported adapting and merging it as a separate experiment, with no default promotion. It was not globally rejected and was not accepted as the primary design process. Related: UG-T01-D15, E08.

**Concept-first preservation, then optional formalization.** After the positive sample, the assistant recommended keeping rich conceptual behavior and separately identifying what a numerical backend cannot express. The owner reported implementing opt-in concept generation. This supports a concept route, not an assertion that the sample is executable. Preserve permitted-but-unimplemented behavior separately from proposals to change the game's rules. Related: UG-T01-D13, E05, Q02, Q13, Q16.

**Research project after a working prototype.** The owner considered a separate research project/paper and then explicitly postponed research to prioritize a working replaceable system. Earlier assistant suggestions for study variants, sample counts, and experimental infrastructure were never accepted as a protocol. Current product context retains a later paper goal. Related: UG-T01-D17, section 8.

### Consequential corrections to preserve

The later skill permits plain early stat upgrades, bounded early utility, a T3 that substantially develops an existing behavior, and T4/T5 that change role. This conflicts with the assistant's earlier insistence on a new T3 mechanic followed only by deepening it. Do not resurrect that earlier template as a universal rule (UG-T01-D10).

The public positive example uses different Gears in different branches. 

The initial phrase "NON-COMMERCIAL PRIVATE USE ONLY" was later clarified by the owner as a rights concern around outputs while the generator should be fully FOSS. No compatible software or content license was selected. A belief that private EU use is unproblematic was challenged, not established (UG-T01-D04).

### Terminology present, and terminology still proposed

The owner's supplied CONTEXT.md defines **Unit** as content describing a game actor or deployable combat element; **Unit draft** connects role, attacks, progression, and abilities, with values where supplied rules justify them; **Mechanic** is a declared behavior; **Mechanic proposal** is an addition suggested for a Game Definition, not approval or implementation evidence (UG-T01-S89).

"Full unit" originally sometimes implied executable behavior, numbers, and a simulator. The skill and near-acceptable example instead demonstrate a complete qualitative concept at a declared scope. No single universal completion contract was settled. "Usable default" has a qualitative anchor in that example, not an approved quantitative threshold (UG-T01-E01; Q02, Q10).

The owner defines **Definition**, **Game Definition**, and **Profile**, but used "replaceable core," "ruleset," and "DSL" broadly. The assistant's later distinction between a stable **Core**, replaceable **Game Definition**, a bounded **Profile**, and a **Definition package** is an unaccepted vocabulary proposal, not the current owner contract. The same applies to proposed **Strategy**, **Deliverable**, **Interpretation**, **Consumer**, and one-model-call **Attempt** terminology (UG-T01-D20; S88). "DSL" refers in context to a domain-specific language, not a selected technology.

"Tuning" covers rules, prompts, examples, parameters, model configuration, or potentially other techniques. Model-weight training, a chosen fine-tuning method, and a guaranteed local/offline route are not established in this thread. "Novelty," "taste," "validity," "accepted," and "balanced" were sometimes used loosely; the later discussion explicitly separates them. Neither an assistant's preference nor a checker result is Towerright project Acceptance.

## 3. Decision record

Approval and evidence are separate below. "Accepted" identifies an explicit owner intention or reported maintainer action, not proof of technical success. Reconsideration conditions are recorded only where discussed; otherwise none is established.

### UG-T01-D01. Character-name input with a complete unit result

**Status: accepted product intention.** Owner locator: "the default flow is 'Anime Name in' -> 'Full TowerDefense Unit out'." The character provides documented material to adapt rather than requiring invention of a new identity. Exact completeness, retrieval policy, and numerical/runtime scope were not initially settled. Evidence: the request and E01, not a reliable name-only application run. Reconsideration condition: not established.

### UG-T01-D02. Reusable generator with replaceable game-specific rules

**Status: accepted intention; implementation completeness unresolved.** Owner: "the entire idea is to have a 'Replaceable' ruleset / dsl" and later "get this working as intented with the fully replacable core." Alternative: character-specific or fixed-game framework logic. The stated purpose is reuse for generalizable concepts and distinct public/private configurations. No fixed public API, arbitrary-domain guarantee, or verified end-to-end replacement follows. Evidence: E08 is an implementation report. Related Q01, Q15, Q16.

### UG-T01-D03. BTD6-derived public prototype structure

**Status: accepted as a prototype direction; conflicting exactness claims unresolved.** Owner first requested the "exact BTD6 concept," later supplied a specific design skill and approved-as-promising its output. The prototype uses three paths, five tiers, restricted crosspaths, and middle-path activation rules. "Exact BTD6" and "BTD6-inspired" were not reconciled through a compatibility specification. The skill's project conventions must not be represented as verified rules of every BTD6 tower. Public Apex scope also remains open. Related E02, E04, Q04, Q06.

### UG-T01-D04. Fully FOSS generator, separate concern about third-party reuse

**Status: accepted FOSS intention; initial non-commercial wording conflicting/clarified in scope.** Owner: "the generator should be fully FOSS BUT ... Default Ruleset ... COULD be against the copyright." This clarifies the concern without choosing an actual license. The earlier "NON-COMMERCIAL PRIVATE USE ONLY" restriction must not silently become the software's license. The owner's private-use legal assumption is not an accepted legal finding. No license compatibility or permission to copy BTD6/anime material was established. Related Q09 and S58 through S70, S93, S95.

### UG-T01-D05. Reject the first generic Luffy organization

**Status: rejected.** Owner: "it doesnt really feel like a nice unit design ... tasteless and basic." Rationale: failure to capture the spirit and consistent figure. Evidence is a direct preference judgment, not a study. Reconsideration was not specified; later acceptance of another geometry-rich design is not retrospective approval of this one.

### UG-T01-D07. Retain the skill-only Luffy output as a positive default example

**Status: accepted, qualified qualitative judgment.** Owner: "as the default generation, this quality would be nearly enough. With a few optimizations that could work arleady." This is not final approval of every mechanic, a balance finding, or a universal template. The producing model and run conditions are unknown. Related E01; S86 sections B and C.

### UG-T01-D08. Treat the supplied skill as a prototype, not final specification

**Status: accepted experiment boundary.** Owner: "THIS ALSO ONYL A PROTOTYPE AND RESEARCH, NOT ANY FINAL VERSION!" Later positive feedback does not approve every rule forever. The skill is inspectable and reusable as a baseline. Automatic reliability and superiority are untested. Related E02, Q04, Q10.

### UG-T01-D09. Keep the private project outside the public repository

**Status: accepted boundary.** Owner: "the 'private' part shouldnt be in the main unit-generator repo." Assistant suggestions to keep outputs, traces, corpora, profiles, and exact private mappings in a separate checkout elaborate that boundary; their individual storage contracts were not approved. Earlier apparent private content in PR history remains an audit issue, not solved by deleting the latest file. Related Q14; S87.

### UG-T01-D10. Do not force every T3 to introduce a new mechanic or every later tier to repeat it

**Status: superseded assistant prescription; prototype replacement supplied by owner.** The owner's skill says T3 "can introduce a behavior, substantially improve an existing one, or establish a different role" and T4/T5 "are not restricted to scaling the tier-3 mechanic." This replaces the earlier assistant's narrow design formula within the proposed skill, not through a finalized repository policy. Related E02.

### UG-T01-D11. Make an executable runtime and full numerical balance prerequisites

**Status: deferred assistant proposal, not an owner-approved architecture.** Early responses required executable rules, simulation, and reference-game calibration. The later handoff says numerical work is an explicit mode or later formalization, not a prerequisite for a successful concept. Owner implementation of opt-in concepts supports that narrower route, not all later design advice. Related D13, E05.

### UG-T01-D12. Use a bounded controller, a reusable skill, and targeted repairs

**Status: proposed.** Assistant: "fully automatic workflow with limited agent autonomy." Later advice favors one complete-concept baseline before compulsory planning, ranking, or multiple agents. Benefits and limits were argued, not measured. Add stages only if they address a diagnosed failure and justify their cost was an explicit proposed reconsideration rule. Owner did not select a final controller or call budget. Related Q12.

### UG-T01-D13. Add concept generation and artifact operations as opt-in features

**Status: accepted/reported implemented by the owner.** Owner update at `739cae3`: "Added opt-in concept generation, configurable rules, save/reload, revision, rendering and attempt evidence capture." This establishes a maintainer report, not independent verification of feature completeness or quality. Related E08, E09, Q01.

### UG-T01-D14. Keep PR #10 as an optional interpretation experiment

**Status: accepted/reported merged.** Owner: "Merged PR #10 as an optional interpretation experiment." Earlier reports narrowed it from automatic layout discovery to explicit caller-selected interpretation. Its value, full external-rule support, and effects on live design remain unverified. Related E07, E08; S76, S81.

### UG-T01-D15. Keep PR #6 as a separate compact-generation experiment

**Status: accepted/reported adapted and merged.** Owner: "Adapted and merged PR #6 as a separate compact-generation experiment." This does not approve hardcoded recipes as the public default. Related E08; S77.

### UG-T01-D16. Do not promote the experiments or freeze a study yet

**Status: accepted/reported maintainer action.** Owner: "The default pipeline remains unchanged" and "no study freeze or default promotion was made." Rationale explicitly accompanies unresolved live quality: three failed candidates, one timeout, zero accepted designs. Reconsideration threshold was not specified. Related E09.

### UG-T01-D17. Postpone formal research to get the intended product working

**Status: accepted priority.** Owner: "i want to postone the 'research' part and get this working." Earlier interest in a separate paper is not withdrawn; research is deferred, not rejected. Current product context also states v1 first and a later paper, but is not the source of this historical decision. No study timeline or publication venue was approved.

### UG-T01-D18. Remove obsolete scratch material and consolidate guidance

**Status: accepted/reported completed.** Owner reports removal of 225 obsolete files, keeping a concept-proof note and 11 useful records. The retained four live attempts were separately said to be preserved. Exact deletion contents, surviving paths, and the scope of "all logs" are not independently known. Do not infer that the available conversation attachments disappeared. Related E10, Q14.

### UG-T01-D19. Use the owner-supplied CONTEXT.md as the historical contributor contract

**Status: accepted as the owner-supplied contract for this discussion; repository correspondence unverified.** Locator: "my current 'CONTEXT.md' is:" followed by the definitions and ownership text. It assigns Unit-Generator unit computation, source evidence, and scoped validation; Towerright project history, cross-tool orchestration, wider validation/balancing, and Acceptance; UnitLab explicit Request preparation. Runtime ownership is not expressly stated in that pasted file, although the research attachment and assistant discuss a Consumer. Related S89.

### UG-T01-D20. Adopt the expanded vocabulary and effective-input/versioning contract

**Status: proposed, not accepted.** `CONTEXT.proposed.md` introduces Core, Definition package, Strategy, Deliverable, Attempt, Consumer, and stronger revision/compatibility rules. The owner requested vocabulary but did not approve this exact replacement. The proposed one-model-invocation meaning of Attempt needs reconciliation with the existing implementation. Related Q15; S88.

## 4. Technical artifacts, research claims, and results

### Inspectable artifacts and reported implementation

| Evidence ID | Item and locator | Status and limits |
| --- | --- | --- |
| UG-T01-E01 | Owner's generated Luffy and approval quotation; S86 sections B-C. | Full concept text inspectable; owner says generated with the skill in ChatGPT. One positive qualitative example. Model/version, settings, hidden context, cost, latency, repeatability, and runtime unknown. |
| UG-T01-E02 | Full skill and Wano brief; S86 sections A-B; original user message. | Written prototype, not executable DSL. The attachment normalizes Markdown escaping and reproduces messages; it is not an independent generation log. |
| UG-T01-E04 | Normal progression predicates and visible Python enumeration. | Small mathematical check shown in the thread: 64 legal upgrade configurations, 19 without T3 and 45 with T3. Not combat simulation. No application implementation is proved by this calculation. |
| UG-T01-E05 | Earlier main review, S85 sections 2-4 and 11. | Assistant-reported static inspection, no immutable main SHA. Fixed 3-by-5 numerical route and one-loop restrictions were reported; not a fresh audit of `739cae3`. |
| UG-T01-E06 | Earlier evaluation notes quoted in S85 section 4.3. | Three development batches reportedly had 0 completed drafts out of 3 subjects each; repeated cases and raw logs unavailable. Do not pool them into a general success estimate. |
| UG-T01-E07 | Updated PR #10 review at `c97fdad8cb052bafa37f6dc92de3179f8ba2f3d3`; assistant's review response. | Static findings and reported 403 offline tests, not reproduced. The described live Luffy run did not supply the new interpretation, so it was not evidence for that feature. `PR10_REVIEW` remained inaccessible. |
| UG-T01-E08 | Owner update beginning "Implemented and pushed to main at 739cae3." | Reports clean tree, opt-in concepts, both experimental PR merges, 436 tests and TypeScript/build/format/doc-link checks passing. No raw checkout or check output inspected here. |
| UG-T01-E09 | Same owner update: "three candidates failed checks and one timed out, with zero accepted designs." | Latest explicit live outcome. Attempts said to be retained, but unavailable here. Failure causes, models, prompts, and definition of "accepted" not established. |
| UG-T01-E10 | Same update: removed 225 files; kept concept-proof note and 11 records. | Owner-reported cleanup. Neither surviving file contents nor interpretation of "all logs" independently checked. |
| UG-T01-E11 | Five available Markdown attachments, S84-S88. | All read in this consolidation. Generated handoffs are proposals or summaries, not evidence their instructions were implemented. Hashes retained in the companion input manifest. |

### Mechanics worth preserving from the positive example

The prototype is an ordinary-unit concept, not an Apex specification. It defines three complete five-tier paths and all six directional crosspath packages. It is evidence of descriptive specificity, not evidence that the mechanics run (E01, E02).

The top path preserves one hit-capacity pool across pursuing turns, records the outbound route, assigns a separate return pool, permits one hit per enemy on each pass, and eventually overlaps two fists according to arm availability. The middle path adds an independently timed Whip, uses a barrage with per-fist capacity, and temporarily moves the attack origin through a manual Boundman deployment while retaining placement occupancy and return/recovery conditions. The bottom path uses a target-hit-and-retraction trigger, one waiting fold, one launch per enemy across copies, an independent landing stamp, and one waiting backstop. Redirection preserves the allied projectile's remaining pierce, travel, lifetime, source effects, and hit history. It excludes several attack types and does not grant the projectile Luffy's detection.

These are regression candidates proposed by the assistant, not a requirement that all future units use them. Flattening an independent attack into a hit-triggered follow-up changes behavior. Preserving a rich concept does not require that a numerical backend already support it. S86 section C and S85 section 12 locate the exact descriptions and proposed preservation tests.

### Equations and parameters actually discussed

For the specific normal progression prototype, let `u=(a,b,c)` with integer tiers in `[0,5]`. Sequential purchases increase one path by one tier; `count(t > 0) <= 2` and `count(t >= 3) <= 1`. These assumptions give E04's 64 configurations and the six maximum crosspath endpoints `5-2-0`, `5-0-2`, `2-5-0`, `0-5-2`, `2-0-5`, `0-2-5`. Timers, targets, resources, and forms create many additional runtime states.


An earlier numerical illustration used placement cost 600 and `roundTo5(basePrice × difficultyMultiplier)`, with example multipliers 0.85, 1.00, 1.10, and 1.25. The example A-path costs 150, 300, 1,800, 6,000, 32,000 led to 2,850 for `3-0-0` and 3,400 for `3-2-0` after adding B1=200 and B2=350. These were assistant-authored unbalanced examples, not verified BTD6 prices or adopted private economics. The later skill intentionally excludes such magnitudes unless requested.

Earlier generation budgets varied: twelve T3 proposals; six layouts/two detailed candidates/three repairs; later four layouts/two candidates/two repairs; finally a one-draft baseline. No budget was approved or optimized. Earlier pilot suggestions of 54 and 72 outputs were unapproved illustrative sample sizes, not conducted experiments or current recommendations.

### Schemas, interfaces, code, and tests

`RulePack`, `DesignProfile`, `ReferencePack`, `DesignPlan`, YAML effect records, a project manifest, and `generate_unit(...)` orchestration pseudocode were suggested. JSON Schema, Zod, and CEL were referenced as possible structural/expression tools, not a settled dependency stack. The thread never establishes that a schema implements an arbitrary mechanic.

The older handoff names existing `AuthorRequest`, `UnitCandidate`, `ModelClient`, CLI, UnitLab, `planned-v1`, and the generic candidate route. Those observations have the scope of E05. The proposed `concept`/`mechanics` Deliverable and `generate`/`redesign`/`prose-edit`/`formalize` operation distinctions were not all proved by the owner's later feature summary. Do not manufacture current API fields from example code.

Tests proposed include progression legality, six crosspath coverage, same-ID/version changed-content detection, active-button restrictions, inheritance, save/reload, prose-edit preservation, complete external-rule propagation, and explicit unsupported-formalization errors. Known-defect injection and positive/negative demonstration scenarios were later suggestions. Passing 436 reported software tests is not evidence that these exact tests all exist or that live design quality is solved (E08, E09).

The older PR #10 review reported fixed layout selection and private-file removal. The later `c97fdad...` review reported caller-selected interpretation, full-content immutability, but prompt omission of concept descriptions, coarse source-ID matching, a three-by-five preparation boundary, and no feature-specific live result. These are historical assistant findings with unavailable underlying code here, not verified remaining defects after `739cae3`. The owner reports both PRs merged; stale earlier "open" pages do not invalidate that report (S76-S81).

### Research claims and their evidence limits

| Evidence ID | Claim actually supportable at its stated scope | Sources and limitation |
| --- | --- | --- |
| UG-T01-E12 | Research demonstrates generated unit configurations and executable game-rule search within supplied representations. | S01 and S03. R1 abstract and GAVEL relevant sections checked here. Neither demonstrates anime fidelity or Unit-Generator's performance. GAVEL's novel-cell advantage was not significant over three runs; its other archive measures were stronger. |
| UG-T01-E13 | Independent generation, an unfamiliar configuration, different gameplay, and historical originality are distinct. | S02, S04-S10, S43; S84 evidence ladder. Ludi's counts and rediscovery example are retained as catalogue-reported results, not newly reproduced. No historical originality test for this project exists. |
| UG-T01-E14 | Modeling human game invention benefits from evaluating proposed consequences, within the studied game class. | S12-S13. S12's v1 dataset description specifies 484 submitted games, not 484 valid instances of every analysis. This is cognitive modeling, not an AI-taste validation. |
| UG-T01-E15 | Novel proposal generation and reliable self-selection are separate capabilities. | S16 and S49-S52. Research-idea ratings and task-specific revision findings do not certify game-design judgments. No preference model trained for this project is shown. |
| UG-T01-E16 | AI assistance can improve individual work or decisions while diversity remains a separate concern. | S40-S42. Fiction, standardized creativity tasks, and Go decisions are not unit-design experiments. Broad future capability claims do not follow. |
| UG-T01-E17 | Evaluation policy, comprehension, contribution, and human preference must not be collapsed. | S07-S08, S17-S19, S22, S37. Catalogue Mortar final-version record says 14 participants; newly checked arXiv v1 section 4.1 says 10. Preserve a version discrepancy rather than treating one as a silent correction. |
| UG-T01-E18 | Self-feedback sometimes improves outputs, while intrinsic self-correction can fail; explicit instruction checks measure a different property from taste. | S49-S52. Self-Refine and the intrinsic-correction study test different settings. Neither establishes this skill's reliability. Checklists paper uses RL from checklist feedback, not merely a long prompt. |
| UG-T01-E19 | Historical capability/cost trends do not guarantee this application's future quality or cost. | S82-S83. Stanford's 2025 webpage reports $20 to $0.07 per million tokens at a specified MMLU threshold, November 2022 to October 2024. That is not an end-to-end accepted-unit cost forecast. |
| UG-T01-E20 | Structural correctness and task completion are not equivalent. | S38, S55-S57. Catalogue/indexed GameASG abstract reports 93.2% best mean runtime checks versus 55.3% strict task success, 26/47; different maxima may belong to different stacks. No game-design preference claim follows. |
| UG-T01-E21 | New literature suggests testing instruction position and reviewer bias rather than trusting longer prompts or reviewer confidence. | Newly identified S96-S97, relevant sections checked. Older QA/chat tasks, not demonstrated errors of the supplied skill or current models. |
| UG-T01-E22 | Unit-Generator is better than direct LLM-plus-wiki prompting. | Hypothesis only. E01 is positive owner feedback on skill-only ChatGPT output; E09 is negative reported application delivery. No matched comparative experiment is available. |

## 5. Recommended approach and v1 boundary

### Recommendation actually reached in this thread

There is no fully approved final architecture. The explicit owner direction is to get the intended replaceable system working, postpone research, retain the useful skill-derived quality target, keep private material outside the public repository, and leave the new routes experimental. The last assistant recommendation was to diagnose the four retained attempts, preserve the known-good concept through artifact operations, and prove replacement with two different Definitions. That recommendation was not subsequently approved or reported executed.

### Current recommendation

This subsection is advice awaiting owner approval, based on this thread and the current product context. It is not a recovered historical decision.

**Product principle:** Make the standalone default produce a useful unit under the user's selected game rules, without silently changing the intended behavior or depending on private platform material.

**Smallest useful flow:** character name plus once-configured public defaults -> explicit source scope/reference input -> generation using the working skill and effective game rules -> checks limited to what the tool can establish -> retained readable Unit result, with revision and export. Additional mechanics or changes to the Game Definition remain optional and requested, not mandatory novelty quotas. A unit still needs enough ordinary attack and progression description to be useful; that is different from requiring new game-wide mechanic proposals.

The public default must itself complete this flow. A bare framework, schema library, or private-only quality path would not meet the supplied current context. Model provider, credentials, hardware, network access, and costs should be documented as actual requirements when chosen, not assumed to be free or offline.

**Necessary now, and why:**

| Component or work | Why it is needed at v1 scope |
| --- | --- |
| Inspect retained failed Runs and their effective inputs. | E09 is the immediate unresolved blocker; another architecture cannot be justified without knowing the failure class. |
| Preserve the complete skill and a reference concept as versioned development inputs. | E01 is the concrete quality anchor. Condensing it or silently formalizing it changes the baseline. |
| A bundled usable public Definition and configurable Profile, resolved into the Request. | The name-only interface needs known game rules; replacement must not rely on hidden private context. |
| Minimal complete Unit representation and readable rendering. | Save/reload and revision must retain triggers, targeting, inheritance, and limitations rather than only names and stats. |
| Scoped checks and bounded repair with failure retention. | Prevent explicit rule violations and lost interactions without claiming taste or balance from a check result. |
| One end-to-end replacement test using a meaningfully different public/synthetic Definition. | Loading another file is not proof that generation, revision, checking, and output follow it. |
| Source facts separated from interpretation and inventions. | A familiar character name does not authorize invented canon, an assumed continuity, or unrestricted reuse. |

Use existing implemented paths where they satisfy these needs. Start diagnostic generation from the complete skill, a frozen reference packet, and one selected model configuration. Keep interpretation and compact Strategies optional. Add further model stages only after a recurring failure identifies their purpose.

**Potential blockers:** raw failed-run evidence may be unavailable; concept checks may enforce unintended numerical or older default restrictions; a schema may erase independent or event-driven behavior; rules may change prompts but not rendering/revision; source resolution or output budgets may fail. These are possibilities to investigate, not newly discovered defects (Q01-Q04, Q11-Q13).

**Deferred or excluded:** a universal executable DSL, compulsory combat simulator, production visual assets, certified balance, automatic whole-game generation, large evolutionary searches, a multi-agent hierarchy, paper-ready study infrastructure, and model-weight training. The numerical route can remain available without becoming a prerequisite for concept success. A public Apex can be added only under an explicit chosen contract; this thread does not settle whether every v1 result must include one.

**Towerright-specific boundary:** the owner's contributor contract assigns broader project history, cross-tool orchestration, wider validation/balancing, and project Acceptance to Towerright. Private game Definitions, exact private source-to-system bindings, private reference collections, and project preferences stay outside the public repository. Current context reserves Mardwerk's specially tuned offering for Towerright without selecting how that tuning works. Do not require Towerright Acceptance or private APIs for a standalone user to receive and locally assess a Unit result.

The visibility and reuse status of character-specific fixtures, generic extension implementations, and any private material already present in public PR history remains unresolved. This handover selects no license and assumes no compatibility. Refer to Q09 and Q14 before exporting it or its examples.

## 6. Open questions, conflicts, and risks

Resolution actions below are proposed evidence or decisions, not accepted requirements.

### UG-T01-Q01. What is replaceable end to end?

**Product/technical choice; can block v1.** Does a selected Definition govern generation, checks, revision, rendering, saved artifacts, and reload, or only prompt text/a loader? Which changes are configuration versus new computation? Resolve with a second substantially different public/synthetic Definition and an actual operation trace, preserving failed incompatibilities. Related D02, D13, E08, D20.

### UG-T01-Q02. What exactly constitutes a complete v1 unit?

**Product choice; blocks acceptance of v1.** Earlier executable-first proposals, the concept-only skill, optional mechanics, and ordinary-unit versus Apex output are not one settled contract. Choose the basic public result, required behavioral detail, optional additions, and honest incomplete/failure states. Do not require numbers, assets, code, or balance merely because earlier examples included them. Related D01, D03, D08, D11, E01.

### UG-T01-Q03. Why did the four live attempts not produce an accepted result?

**Diagnostic uncertainty; immediate blocker.** Distinguish provider/timeout/truncation, source/input, structural, genuine rule, reviewer, and design failures. Inspect actual Requests, full model messages, raw outputs, termination causes, and findings. No cause can be assigned from counts alone. Related E08-E09 and D16.

### UG-T01-Q04. Which public default conventions are authoritative?

**Product choice; can block v1.** Exact BTD6 compatibility versus inspiration; mandatory versus merely permitted middle-path ability; early secondary attacks/armor handling; qualitative versus numeric defaults; shared forms versus upgrade-local transformations. Version an explicit prototype contract and retire or scope contradictory older rules. Related D03, D08, D10, E02, E05.

### UG-T01-Q06. Is a public Apex required, and how is it acquired and composed?

**Product choice; blocks only the deliverable that promises it.** The owner wanted a Paragon equivalent, but the near-acceptable ordinary sample contains no Apex. Decide availability, prerequisites, representation, and scope independently. Related D03, D06, D07, Q02.

### UG-T01-Q07. What does an interpretation or source check actually guarantee?

**Technical/evidence issue; affects optional experiment and fidelity claims.** Historical PR #10 checks were reported to accept broad source IDs/base identity rather than demonstrated specialization fidelity, and to omit semantic descriptions from prompt fragments. Verify actual current code and test adversarial examples. A retained interpretation or valid citation is not proof the generated mechanic follows it. Related D14, E07, S76, S81.

### UG-T01-Q08. What operational environment does standalone use require?

**Product/operational choice; can block basic usability.** Providers, local models, credentials, hosting, hardware, transmission permissions, budgets, and interruption behavior are not settled. Document and test an actual supported route without assuming that FOSS means no running costs or offline operation. Private input transmission requires its own policy. Related D04, D09, E09.

### UG-T01-Q09. Which licenses and publication permissions apply?

**Release/disclosure issue; affects public release and later datasets.** Software licensing, BTD6-like rule implementation, third-party character references, copied assets/text, generated outputs, and private-use exceptions are separate questions. No specific license or legal clearance is established. Resolve by inspecting the actual chosen artifacts and applicable license/legal advice, not a blanket non-commercial label. Related D04, S58-S70, S93, S95.

### UG-T01-Q10. What would "better than a direct LLM plus wiki" mean?

**Product criterion and research uncertainty.** More useful behavior, fewer contradictions, stronger controllability, less correction, and lower accepted-result cost are candidates, not approved metrics. Compare against a competent direct prompt with matched information and disclosed resource differences. Related D07, E01, E22, section 8.

### UG-T01-Q11. How does the default obtain and scope character evidence?

**Technical/product choice; can block the name-only flow.** The Wano brief is supplied evidence, not a demonstrated automatic retrieval run. Resolve ambiguous identities, continuity, spoilers, source availability, uncertain facts, and what is sent to a provider. Test retrieval separately from design on frozen packets. Related D01, E02, S72-S74.

### UG-T01-Q12. How much agent autonomy or extra generation is justified?

**Technical choice and later research question.** Direct skill, multiple layouts, pinned interpretation, compact generation, external feedback, and multi-agent search are alternatives. No optimal call budget or architecture has been demonstrated. Keep simple baselines and diagnose recurring failures; compare added stages at disclosed total cost. Model progress is a reason for replaceable Strategies, not a guaranteed solution date. Related D12, D14-D16, E15, E18-E19.

### UG-T01-Q13. Can revision and formalization preserve the accepted design?

**Technical issue; can block a trustworthy v1.** Independent attacks, route history, trigger timing, resource budgets, and crosspath transfer may be lost in representation or editing. Resolve by round-tripping exact examples and classifying intentional redesign separately from prose change or formalization. No runtime proof follows from a text-preservation test. Related E01, E05, D13.

### UG-T01-Q14. What was retained, deleted, or already disclosed?

**Evidence/disclosure issue; affects audit and privacy.** The owner reports both "all logs" removed and current failed attempts retained; categories may differ, but exact paths are absent. `PR10_REVIEW` and commit bytes were unavailable in earlier review. Historical private-file deletion is not proof removal from public history. Obtain the actual current records and review disclosure before public export. Related D09, D18, E07-E11, S80-S81, S87.

### UG-T01-Q15. Which vocabulary is adopted?

**Contributor-contract choice; affects implementation consistency.** The owner contract and the assistant's proposed replacement differ in Core, package authority, Attempt, Consumer, and Acceptance wording. Decide which terms reflect the code, especially Definition versus allowed Profile variation and successful checks versus project Acceptance. Do not silently treat S88 as adopted. Related D19-D20, S89.

### UG-T01-Q16. Is a missing behavior a rule change or an implementation gap?

**Technical/product distinction; affects creative scope.** A permitted existing game behavior can be unsupported by a numerical backend. It does not automatically require a new Mechanic proposal. Conversely, a well-described proposal can violate the selected game rules. Record permission, specification completeness, operation support, and Consumer support separately. Related D02, E05, D20, S89.

## 7. Sources and research leads

### A. Sources already present in this thread

This register preserves sources found in the visible messages, tool-call URLs, and available attachments. UG-T01-S01 through S38 map one-to-one to the uploaded catalogue's R1 through R38. An entry can be present only as a lead; inclusion is not endorsement or adoption. "Verified" below is limited to the named bibliographic fields. "Content inspected" is reported independently and never means experiments were reproduced. The original catalogue's earlier inspection statement is retained as historical evidence, not silently upgraded to fresh verification.

All fresh checks were performed September 24, 2026. Primary HTML, abstracts, and bibliographic pages were used. No research code or dataset was installed, reproduced, or license-audited; no PDF text was analyzed for this consolidation. Direct URLs are retained even when access failed. "et al." is preserved where the available record does not establish a full author list. New metadata corrections and version differences are written beside the original wording rather than overwriting it.

#### UG-T01-S01. Generating RTS units

Kynan Sorochan and Matthew Guzdial. Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search. Experimental AI in Games Workshop, 2022. arXiv:2212.03387v1 (https://arxiv.org/html/2212.03387v1). Inspected: Search Space, Evaluator, Search Method, generated units, evaluation and references.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R1, starting at file section heading / rendered line 34. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: primary arXiv title, authors, date and workshop comment. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Unit abilities and evaluation inside microRTS; supplied operators and evaluator competence limit transfer. Related: UG-T01-E12; UG-T01-D11.

Associated artifacts: microRTS and ten generated units described by the source; no environment or code executed here.

Direct URLs:
https://arxiv.org/html/2212.03387v1
https://arxiv.org/abs/2212.03387


#### UG-T01-S02. Evolutionary Game Design

Cameron Browne and Frédéric Maire. Evolutionary Game Design. IEEE Transactions on Computational Intelligence and AI in Games 2(1), pages 1 to 16, 2010. Author PDF (https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf). Inspected: quality measurements, human comparisons, Experiment III and conclusions. The filename is not the publication year.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R2, starting at file section heading / rendered line 48. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: institutional bibliographic record. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Human preference and independent rediscovery. The detailed 1,389/19/127/27 counts remain catalogue-attributed here. DOI recovered: 10.1109/TCIAIG.2010.2041928. Related: UG-T01-E13.

Associated artifacts: Ludi; generated Yavalath and Ndengrod, with Irensei as the reported rediscovery precedent. No rule corpus or player data retrieved.

Direct URLs:
https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf
https://cris.maastrichtuniversity.nl/en/publications/c53c0467-e476-4be8-8373-4f7f0aa6ac75


#### UG-T01-S03. GAVEL

Graham Todd, Alexander G. Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers and Julian Togelius. GAVEL: Generating Games Via Evolution and Language Models. NeurIPS 2024. arXiv:2407.09388v2 (https://arxiv.org/html/2407.09388v2). Inspected: representation, mutation, evaluation, experiments, results, limitations and references. Code link (https://github.com/gdrtodd/gavel), not executed in this investigation.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R3, starting at file section heading / rendered line 58. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv record; NeurIPS attribution retained. Content inspected: relevant sections: representation, evaluation, results and limitations, v2.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Executable mutation and candidate archives. Section 6.1 distinguishes significant archive gains from nonsignificant novel-cell gains over three runs. Does not establish automatic character interpretation. Related: UG-T01-E12; UG-T01-D12.

Associated artifacts: Author-linked GAVEL repository contains implementation/artifact entry points; Ludii is the game-description environment. No dependency or dataset reuse approved.

Direct URLs:
https://arxiv.org/html/2407.09388v2
https://github.com/gdrtodd/gavel


#### UG-T01-S04. Quality diversity for PCG

Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius and Georgios N. Yannakakis. Procedural Content Generation through Quality Diversity. IEEE Conference on Games, 2019. arXiv:1907.04053 (https://arxiv.org/abs/1907.04053), HTML (https://arxiv.org/html/1907.04053), DOI:10.1109/CIG.2019.8848053 (https://doi.org/10.1109/CIG.2019.8848053). Inspected: framework, surveyed applications, challenges and references.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R4, starting at file section heading / rendered line 72. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata and DOI. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Survey/framework for preserving useful alternatives; descriptors and quality functions remain domain choices, not validated TD taste measures. Related: UG-T01-E13.

Direct URLs:
https://arxiv.org/abs/1907.04053
https://arxiv.org/html/1907.04053
https://doi.org/10.1109/CIG.2019.8848053


#### UG-T01-S05. MAP-Elites

Jean-Baptiste Mouret and Jeff Clune. Illuminating search spaces by mapping elites. 2015 preprint. arXiv:1504.04909v1 (https://arxiv.org/abs/1504.04909v1). Inspected: algorithm, experimental domains, limitations and references via arXiv HTML. The HTML contains an inconsistent generated date; use the arXiv version record for the 2015 date.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R5, starting at file section heading / rendered line 84. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv v1 record. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: A feature-space archive, not proof that archive coverage equals gameplay novelty. Preliminary preprint, not a TD experiment. Related: UG-T01-E13.

Direct URLs:
https://arxiv.org/abs/1504.04909v1


#### UG-T01-S06. Novelty search

Joel Lehman and Kenneth O. Stanley. Abandoning Objectives: Evolution through the Search for Novelty Alone. Evolutionary Computation 19(2), pages 189 to 223, 2011. Primary-text PDF (https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf). Inspected: motivation, behavioral novelty, experimental tasks and discussion. Found in both R4 and R5.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R6, starting at file section heading / rendered line 96. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: institutional record; full-text evidence not rechecked. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Behavioral novelty in the studied tasks; no warrant to remove game constraints. DOI record recovered as 10.1162/EVCO_a_00025. Related: UG-T01-E13.

Direct URLs:
https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf
https://researcher.itu.dk/da/publications/abandoning-objectives-evolution-through-the-search-for-novelty-al/


#### UG-T01-S07. Simulation-based strategy analysis

Alexander Zook, Brent Harrison and Mark O. Riedl. Monte-Carlo Tree Search for Simulation-based Strategy Analysis. The paper identifies FDG 2015 as its conference; the arXiv deposit is 2019. arXiv:1908.01423v1 (https://arxiv.org/html/1908.01423v1). Inspected: methodology, skill-based metrics, case studies, limitations and references. R1 cites the deposit year.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R7, starting at file section heading / rendered line 106. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; FDG 2015 attribution retained from catalogue. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Action traces and policy differences. A 2019 deposit is not the conference date, and larger MCTS budgets are not validated human-skill models. Related: UG-T01-E17.

Associated artifacts: Scrabble and Cardonomicon case-study environments; no traces or agent implementations executed here.

Direct URLs:
https://arxiv.org/html/1908.01423v1
https://arxiv.org/abs/1908.01423v1


#### UG-T01-S08. Restricted-play evaluation

Alexander Jaffe, Alex Miller, Erik Andersen, Yun-En Liu, Anna Karlin and Zoran Popović. Evaluating Competitive Game Balance with Restricted Play. AIIDE 2012. Author-hosted PDF (https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf). Inspected: framework, example application and conclusions. Reached through R1 → R7 → R8.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R8, starting at file section heading / rendered line 118. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: publisher metadata. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Controlled restriction of options; useful later for contribution tests, not fun certification. DOI recovered: 10.1609/aiide.v8i1.12513; AIIDE 8(1), 26-31. Related: UG-T01-E17.

Direct URLs:
https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf
https://ojs.aaai.org/index.php/AIIDE/article/view/12513


#### UG-T01-S09. Conceptual expansion

Matthew Guzdial and Mark Riedl. Automated Game Design via Conceptual Expansion. AIIDE 2018, 14(1), pages 31 to 37. arXiv:1809.02232v1 (https://arxiv.org/html/1809.02232v1). Inspected: representation, expansion, evaluation, discussion and references. Reached through R1.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R9, starting at file section heading / rendered line 130. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Structured recombination. Held-out existing-game reconstruction is not unrestricted original-game invention. Related: UG-T01-E13.

Associated artifacts: Super Mario Bros., Kirby's Adventure and Mega Man used as the paper's source/reconstruction examples; not a public Unit-Generator dataset.

Direct URLs:
https://arxiv.org/html/1809.02232v1
https://arxiv.org/abs/1809.02232


#### UG-T01-S10. Constrained novelty search

Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. Constrained Novelty Search: A Study on Game Content Generation. Evolutionary Computation 23(1), pages 101 to 129, 2015. Author PDF (https://antoniosliapis.com/papers/constrained_novelty_search.pdf), DOI:10.1162/EVCO_a_00123 (https://doi.org/10.1162/EVCO_a_00123). Inspected: two-population approach, game-content application and reported qualifications. Found through R4; bibliographic metadata also checked through Crossref.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R10, starting at file section heading / rendered line 142. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: unverified in this consolidation; catalogue reports prior Crossref/text checking. Content inspected: not inspected.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Feasible/infeasible search motivates retaining explicit proposals. Merely storing unfinished drafts is not implementing the algorithm. Related: UG-T01-Q16; UG-T01-E13.

Direct URLs:
https://antoniosliapis.com/papers/constrained_novelty_search.pdf
https://doi.org/10.1162/EVCO_a_00123


#### UG-T01-S11. Sentient Sketchbook

Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. Sentient Sketchbook: Computer-Aided Game Level Authoring. Foundations of Digital Games, 2013, pages 213 to 220. Author PDF (https://antoniosliapis.com/papers/sentient_sketchbook.pdf). Inspected: tool interaction, checks, alternatives and small-scale designer study. Found through R4.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R11, starting at file section heading / rendered line 152. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified in part: institutional title/authors/venue; spelling difference retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Institutional title says Computer-Assisted, whereas the supplied catalogue says Computer-Aided. Preserve the alias until final citation export. Map-authoring interaction, not proof of UnitLab usability. Related: UG-T01-Q13.

Direct URLs:
https://antoniosliapis.com/papers/sentient_sketchbook.pdf
https://researcher.itu.dk/en/publications/sentient-sketchbook-computer-assisted-game-level-authoring/


#### UG-T01-S12. Human game invention

Katherine M. Collins et al. Generation and Evaluation in the Human Invention Process through the Lens of Game Design. CogSci 2025. arXiv:2508.10914v1 (https://arxiv.org/html/2508.10914v1). Inspected: dataset, proposal/evaluation model, results, limitations and references.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R12, starting at file section heading / rendered line 162. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv record; publication-status qualification below. Content inspected: relevant sections: v1 dataset, modeling scope and limitations.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: V1 describes 484 submissions; that is not the count in every analysis. Current arXiv comment calls the CogSci paper non-archival; v2 is July 2026. Full authors in current record: Katherine M. Collins, Graham Todd, Cedegao E. Zhang, Adrian Weller, Julian Togelius, Junyi Chu, Lionel Wong, Thomas L. Griffiths, Joshua B. Tenenbaum. Related: UG-T01-E14.

Associated artifacts: Dataset of 484 submitted human games and a Ludax formalization subset; no project-specific participant data collected.

Direct URLs:
https://arxiv.org/html/2508.10914v1


#### UG-T01-S13. Reasoning about unfamiliar games

Cedegao E. Zhang, Katherine M. Collins, Lionel Wong, Adrian Weller and Joshua B. Tenenbaum. People use fast, goal-directed simulation to reason about novel games. CogSci 2024. arXiv:2407.14095v1 (https://arxiv.org/html/2407.14095v1). Inspected: model, human judgment task and discussion. Found through R12.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R13, starting at file section heading / rendered line 174. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv record; v1/v2 author difference retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Current v2 author list includes Mauricio Barba, absent from the catalogue v1 list. Do not silently merge versions. Unfamiliar connect-n judgments are not real-time TD playtests. Related: UG-T01-E14; UG-T01-E17.

Direct URLs:
https://arxiv.org/html/2407.14095v1
https://arxiv.org/abs/2407.14095


#### UG-T01-S14. FunSearch

Bernardino Romera-Paredes et al. Mathematical discoveries from program search with large language models. Nature 625, pages 468 to 475, 2024; published online in December 2023. DOI:10.1038/s41586-023-06924-6 (https://www.nature.com/articles/s41586-023-06924-6). Inspected: primary article's method and reported applications. Code link (https://github.com/google-deepmind/funsearch), not executed here.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R14, starting at file section heading / rendered line 184. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: unverified in this consolidation; primary page retrieval failed. Content inspected: not inspected.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Catalogue reports inspected methods and applications. Mathematical/algorithmic evaluation is not a validated fun or fidelity judge. Associated code was not executed. Related: UG-T01-D11; UG-T01-E12.

Direct URLs:
https://www.nature.com/articles/s41586-023-06924-6
https://github.com/google-deepmind/funsearch


#### UG-T01-S15. Ludax

Graham Todd et al. Ludax: A GPU-Accelerated Description Language for Board Games. arXiv:2506.22609 (https://arxiv.org/abs/2506.22609), first submitted in 2025; v2 inspected (https://arxiv.org/html/2506.22609v2), March 25, 2026. Inspected: representation, compilation, supported game class and references. R12 cites the earlier 2025 version; this record deliberately distinguishes it from the newer text inspected.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R15, starting at file section heading / rendered line 194. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv v2 metadata, with title correction recorded. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Verified v2 title is Ludax: A GPU-Accelerated Domain Specific Language for Board Games, rather than the catalogue wording Description Language. Todd, Alexander G. Padula, Dennis J. N. J. Soemers, Togelius. Turn-based perfect-information scope, not TD runtime compatibility. Related: UG-T01-Q16.

Associated artifacts: Ludax/JAX environment implementation described by the source; not a real-time TD runtime dependency.

Direct URLs:
https://arxiv.org/abs/2506.22609
https://arxiv.org/html/2506.22609v2
https://arxiv.org/abs/2506.22609v2


#### UG-T01-S16. LLM idea generation and self-evaluation

Chenglei Si, Diyi Yang and Tatsunori Hashimoto. Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers. ICLR 2025; preprint 2024. arXiv:2409.04109 (https://arxiv.org/abs/2409.04109), v1 inspected (https://arxiv.org/html/2409.04109v1). Inspected: study description and discussion of novelty, diversity and self-evaluation. The inspected evidence is the preprint, not an audit of changes in the final proceedings version.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R16, starting at file section heading / rendered line 204. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv record; ICLR 2025 attribution retained. Content inspected: relevant sections: v1 abstract, study description and self-evaluation discussion.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Experts judged proposals, not completed discoveries; useful warning about selection and diversity, not evidence of this generator's taste. Related: UG-T01-E15.

Associated artifacts: Authors link code/review records at https://github.com/NoviScl/AI-Researcher ; not downloaded, run, or license-audited.

Direct URLs:
https://arxiv.org/abs/2409.04109
https://arxiv.org/html/2409.04109v1


#### UG-T01-S17. Mortar: mechanics evaluated inside games

Muhammad U. Nasir, Yuchen Li, Steven James and Julian Togelius. Mortar: Evolving Mechanics for Automatic Game Design. GECCO 2026. Primary PDF (https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf), DOI:10.1145/3795095.3805100 (https://doi.org/10.1145/3795095.3805100). Inspected: method, mechanic contribution measure, human comparisons and related work.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R17, starting at file section heading / rendered line 218. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: author-lab GECCO 2026 metadata and arXiv record. Content inspected: relevant sections of arXiv v1, including section 4.1; published PDF not inspected here.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Version discrepancy: catalogue cites the GECCO PDF with 14 participants; arXiv 2601.00105v1 section 4.1 says 10. Do not replace the published-version claim using the earlier preprint. Both concern small games, not reliable TD preference prediction. Related: UG-T01-E17.

Direct URLs:
https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf
https://doi.org/10.1145/3795095.3805100
https://www.raillab.org/publication/nasir-2026-mortar/
https://arxiv.org/html/2601.00105v1


#### UG-T01-S18. Mechanic Miner

Michael Cook, Simon Colton, Azalea Raad and Jeremy Gow. Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design. EvoApplications 2013. Author PDF (https://www.possibilityspace.org/papers/evo13.pdf), DOI:10.1007/978-3-642-37192-9_29 (https://doi.org/10.1007/978-3-642-37192-9_29). Inspected: representation, search, generated mechanics and pilot study. Promoted from the original R9 citation lead.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R18, starting at file section heading / rendered line 228. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: unverified in this consolidation; DOI retrieval failed. Content inspected: not inspected.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Catalogue reports a primary-text inspection and inconsistent participant counts. Mechanic-plus-level generation is a lead for demonstrations, not authority to require a v1 simulator. Related: UG-T01-E17.

Direct URLs:
https://www.possibilityspace.org/papers/evo13.pdf
https://doi.org/10.1007/978-3-642-37192-9_29


#### UG-T01-S19. Mechanic Maker 2.0

Johor Jara Gonzalez, Seth Cooper and Matthew Guzdial. Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules. AIIDE 19(1), pages 266 to 275, 2023. arXiv:2309.09476v3 (https://arxiv.org/html/2309.09476v3), DOI:10.1609/aiide.v19i1.27522 (https://doi.org/10.1609/aiide.v19i1.27522). Inspected: environment recreation, evaluator comparison and limitations.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R19, starting at file section heading / rendered line 236. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; DOI/venue retained from catalogue. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: RL versus A* can change generated mechanics; catalogue notes no human-subject confirmation of usability. Do not equate learning agent with player model. Related: UG-T01-E17.

Associated artifacts: Unity recreation and A*/RL evaluators described by the paper; no Consumer selected for this project.

Direct URLs:
https://arxiv.org/html/2309.09476v3
https://doi.org/10.1609/aiide.v19i1.27522
https://arxiv.org/abs/2309.09476


#### UG-T01-S20. Pixie

Michael Cook. Pixie: Code-Level Mechanic Generation for Game Designers. AIIDE 21(1), pages 206 to 215, 2025. Publisher PDF (https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962), DOI:10.1609/aiide.v21i1.36824 (https://doi.org/10.1609/aiide.v21i1.36824). Access gap: the publisher download failed and an attempted author copy was unavailable. Mortar's related-work discussion independently confirms Pixie as code-level mechanic evolution using a genetic-programming approach.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R20, starting at file section heading / rendered line 244. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: indexed publisher bibliographic record. Content inspected: abstract only; full publisher PDF still unavailable.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: The former full-text access gap is not closed by verifying the abstract. Detailed Unity annotation/object-pooling claims remain attributed to the supplied review. Related: UG-T01-Q16.

Direct URLs:
https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962
https://doi.org/10.1609/aiide.v21i1.36824
https://ojs.aaai.org/index.php/AIIDE/article/view/36824


#### UG-T01-S21. Formal mechanic generation

Alexander Zook and Mark O. Riedl. Automatic Game Design via Mechanic Generation. AAAI 2014. arXiv:1908.01420v1 (https://arxiv.org/html/1908.01420v1), DOI:10.1609/aaai.v28i1.8788 (https://doi.org/10.1609/aaai.v28i1.8788). Inspected: representation, constraint-based generation, planning and demonstration domains. The arXiv deposit is 2019; the paper is from 2014. Promoted from the original R3/R7 citation lead.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R21, starting at file section heading / rendered line 252. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv title/authors/deposit; AAAI 2014 venue retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Formal conditions and planner-based playability. Deposit year 2019 differs from publication 2014. Reachability under a model is not enjoyment. Related: UG-T01-E12; UG-T01-Q16.

Direct URLs:
https://arxiv.org/html/1908.01420v1
https://doi.org/10.1609/aaai.v28i1.8788
https://arxiv.org/abs/1908.01420v1


#### UG-T01-S22. TowerMind

Dawei Wang et al. TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents. AAAI 2026. arXiv:2601.05899v1 (https://arxiv.org/html/2601.05899v1), DOI:10.1609/aaai.v40i31.39818 (https://doi.org/10.1609/aaai.v40i31.39818). Inspected: environment, observation formats, evaluation scope and reported agent failures. Code (https://github.com/tb6147877/TowerMind), not executed.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R22, starting at file section heading / rendered line 260. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv record; supplied AAAI DOI not separately audited. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: TD agent benchmark and observation/action design, not a generator of novel unit rules. The linked environment was not installed or tested. Related: UG-T01-E17.

Associated artifacts: Author-linked TowerMind repository; no benchmark level or agent run performed.

Direct URLs:
https://arxiv.org/html/2601.05899v1
https://doi.org/10.1609/aaai.v40i31.39818
https://github.com/tb6147877/TowerMind
https://arxiv.org/abs/2601.05899v1


#### UG-T01-S23. Tower defense level generation

Simon Liu et al. Automatic generation of tower defense levels using PCG. FDG 2019. DOI:10.1145/3337722.3337723 (https://doi.org/10.1145/3337722.3337723), supplied manuscript entry (https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG). Access gap: ResearchGate returned HTTP 403; the primary text was not independently inspected.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R23, starting at file section heading / rendered line 268. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: unverified in this consolidation; DOI/source access unresolved. Content inspected: not inspected.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Keep author listing as Simon Liu et al.; full author list not recovered. Flat Monte Carlo and implementation details are supplied-review claims. Related: UG-T01-D11.

Associated artifacts: Kingdom Rush: Frontiers-style simulator is a supplied-review description pending primary-text confirmation.

Direct URLs:
https://doi.org/10.1145/3337722.3337723
https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG


#### UG-T01-S24. Save the Sheep

Vid Kraner, Iztok Fister Jr. and Lucija Brezočnik. Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms. NT 2021, Lecture Notes in Networks and Systems 233. DOI:10.1007/978-3-030-75275-0_54 (https://doi.org/10.1007/978-3-030-75275-0_54), supplied manuscript entry (https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms). Access gap: ResearchGate returned HTTP 403.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R24, starting at file section heading / rendered line 276. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: unverified in this consolidation; primary access unresolved. Content inspected: not inspected.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Supplied description distinguishes human-authored game/assets from generated arrangements; no implementation claim independently checked. Related: UG-T01-D11.

Associated artifacts: Save the Sheep, Unity and human-made Blender assets are source-review descriptions, not verified adopted tools.

Direct URLs:
https://doi.org/10.1007/978-3-030-75275-0_54
https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms


#### UG-T01-S25. Early tower defense research

Phillipa Avery, Julian Togelius, Elvis Alistar and Robert Pieter van Leeuwen. Computational Intelligence and Tower Defence Games. IEEE Congress on Evolutionary Computation, 2011. DOI:10.1109/CEC.2011.5949738 (https://doi.org/10.1109/CEC.2011.5949738), supplied manuscript entry (https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games). Access gap: ResearchGate returned HTTP 403.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R25, starting at file section heading / rendered line 284. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: unverified in this consolidation; primary access unresolved. Content inspected: not inspected.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Retain catalogue author spelling and DOI, not unverified reuse claims. Early TD decomposition lead, not a required dependency. Related: UG-T01-D11.

Direct URLs:
https://doi.org/10.1109/CEC.2011.5949738
https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games


#### UG-T01-S26. Multiobjective RTS maps

Julian Togelius et al. Controllable procedural map generation via multiobjective evolution. Genetic Programming and Evolvable Machines 14, pages 245 to 277, 2013. Publisher article and DOI (https://link.springer.com/article/10.1007/s10710-012-9174-5). Abstract and reference list inspected; the full article requires access.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R26, starting at file section heading / rendered line 290. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: publisher metadata. Content inspected: abstract and references only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Full authors: Julian Togelius, Mike Preuss, Nicola Beume, Simon Wessing, Johan Hagelbäck, Georgios N. Yannakakis, Corrado Grappiolo. Conflicting map objectives are relevant later; visual review is not extensive playtesting. Related: UG-T01-D11.

Direct URLs:
https://link.springer.com/article/10.1007/s10710-012-9174-5


#### UG-T01-S27. Controllable 3D level generators

Zehua Jiang, Sam Earle, Michael Cerny Green and Julian Togelius. Learning Controllable 3D Level Generators. FDG 2022. arXiv:2206.13623v3 (https://arxiv.org/html/2206.13623v3), DOI:10.1145/3555858.3563273 (https://doi.org/10.1145/3555858.3563273). Inspected: PCGRL formulation, 3D movement requirements and control experiments.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R27, starting at file section heading / rendered line 298. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; FDG citation retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Movement-aware 3D generation. Human-defined rewards/representations and voxel environments do not establish production RTS navigation. Related: UG-T01-D11.

Associated artifacts: Minecraft-style 3D level tasks; not generated production RTS geometry.

Direct URLs:
https://arxiv.org/html/2206.13623v3
https://doi.org/10.1145/3555858.3563273
https://arxiv.org/abs/2206.13623


#### UG-T01-S28. Scaling PCGRL

Sam Earle, Zehua Jiang and Julian Togelius. Scaling, Control and Generalization in Reinforcement Learning Level Generators. 2024. arXiv:2408.12525v1 (https://arxiv.org/html/2408.12525v1). Inspected: method, scaling and generalization scope. The supplied list calls this "PCGRL+"; the citation here uses the manuscript title.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R28, starting at file section heading / rendered line 306. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv record, with version-title qualification. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Current record title begins PCGRL+: Scaling, Control and Generalization in Reinforcement Learning Level Generators; the catalogue deliberately used a version without the prefix. No RL infrastructure chosen. Related: UG-T01-D11.

Associated artifacts: PCGRL/JAX experiments, map-size variation and frozen tiles are catalogue-reported contexts.

Direct URLs:
https://arxiv.org/html/2408.12525v1
https://arxiv.org/abs/2408.12525


#### UG-T01-S29. Database-driven 3D levels

Kaijie Xu and Clark Verbrugge. A Database-Driven Framework for 3D Level Generation with LLMs. AIIDE 21(1), pages 367 to 376, 2025. arXiv:2508.18533v1 (https://arxiv.org/html/2508.18533v1), DOI:10.1609/aiide.v21i1.36840 (https://doi.org/10.1609/aiide.v21i1.36840). Inspected: assembly pipeline, repair and limitations.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R29, starting at file section heading / rendered line 314. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; AIIDE DOI retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Assembly and navigation repair lead. Catalogue limitations include curation, unplayable outputs and no player study; no engine migration approved. Related: UG-T01-Q13.

Direct URLs:
https://arxiv.org/html/2508.18533v1
https://doi.org/10.1609/aiide.v21i1.36840
https://arxiv.org/abs/2508.18533


#### UG-T01-S30. Infinigen

Alexander Raistrick et al. Infinite Photorealistic Worlds using Procedural Generation. CVPR 2023. arXiv:2306.09310 (https://arxiv.org/abs/2306.09310). Primary HTML inspected for system purpose and generation scope.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R30, starting at file section heading / rendered line 322. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; CVPR attribution retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Synthetic visual environments, not measured tactical map quality. Production assets are outside the agreed immediate direction. Related: UG-T01-D11.

Associated artifacts: Infinigen procedural scene system; no code/assets imported.

Direct URLs:
https://arxiv.org/abs/2306.09310


#### UG-T01-S31. Holodeck

Yue Yang et al. Holodeck: Language Guided Generation of 3D Embodied AI Environments. CVPR 2024. arXiv:2312.09067 (https://arxiv.org/abs/2312.09067). Primary HTML inspected for planning, retrieval, spatial constraints and evaluation scope.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R31, starting at file section heading / rendered line 330. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; CVPR attribution retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Spatial constraints and retrieved assets for embodied environments. Source was not evidence that attractive scenes imply useful TD geometry. Related: UG-T01-D11.

Associated artifacts: Holodeck scene/asset pipeline; no retrieved asset rights assessed.

Direct URLs:
https://arxiv.org/abs/2312.09067


#### UG-T01-S32. Galactic Arms Race

Erin J. Hastings, Ratan K. Guha and Kenneth O. Stanley. Automatic Content Generation in the Galactic Arms Race Video Game. IEEE Transactions on Computational Intelligence and AI in Games 1(4), pages 245 to 263, 2009. DOI:10.1109/TCIAIG.2009.2038365 (https://doi.org/10.1109/TCIAIG.2009.2038365), supplied manuscript entry (https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game). Access gap: ResearchGate returned HTTP 403.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R32, starting at file section heading / rendered line 338. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: unverified in this consolidation; DOI access unresolved. Content inspected: not inspected.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Particle-weapon evolution and more-than-1,000-player claims remain supplied-review claims. Usage can reflect power or availability rather than preference. Related: UG-T01-Q10.

Direct URLs:
https://doi.org/10.1109/TCIAIG.2009.2038365
https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game


#### UG-T01-S33. OMNI-EPIC

Maxence Faldor, Jenny Zhang, Antoine Cully and Jeff Clune. OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code. ICLR 2025. arXiv:2405.15568v3 (https://arxiv.org/html/2405.15568v3). Inspected: archive, task/environment generation, interestingness filtering and evaluation modes.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R33, starting at file section heading / rendered line 346. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; ICLR attribution retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Open-ended environment/task generation. Catalogue qualifies that a long 200-iteration run assumed solvability rather than training agents throughout. No such archive built here. Related: UG-T01-D12.

Associated artifacts: PyBullet task environments and successful/failed task archive; no task reproduction.

Direct URLs:
https://arxiv.org/html/2405.15568v3
https://arxiv.org/abs/2405.15568


#### UG-T01-S34. DreamGarden

Sam Earle, Samyak Parajuli and Andrzej Banburski-Fahey. DreamGarden: A Designer Assistant for Growing Games from a Single Prompt. CHI 2025. arXiv:2410.01791v2 (https://arxiv.org/html/2410.01791v2), DOI:10.1145/3706598.3714233 (https://doi.org/10.1145/3706598.3714233). Inspected: hierarchical planning, Unreal implementation modules and usability study.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R34, starting at file section heading / rendered line 356. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; CHI DOI retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Editable plans and implementation feedback. Ten-participant usability result remains catalogue-attributed, not complete-game quality proof. Related: UG-T01-Q13.

Associated artifacts: Unreal implementation modules and editable plans; no code or assets imported.

Direct URLs:
https://arxiv.org/html/2410.01791v2
https://doi.org/10.1145/3706598.3714233
https://arxiv.org/abs/2410.01791


#### UG-T01-S35. AutoUE

Lei Yin et al. AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems. 2026 preprint. arXiv:2603.07106v1 (https://arxiv.org/html/2603.07106v1). Inspected: engine workflow, documentation retrieval, runtime tests and benchmark. Code (https://github.com/Pluto156/AutoUE), not executed.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R35, starting at file section heading / rendered line 364. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: 2026 preprint. PlayGen-20 and Unreal automation are not an established autonomous novel TD design system. Code link retained but unexecuted. Related: UG-T01-D12.

Associated artifacts: Author-linked AutoUE code and PlayGen-20 benchmark; no execution or licensing audit.

Direct URLs:
https://arxiv.org/html/2603.07106v1
https://github.com/Pluto156/AutoUE
https://arxiv.org/abs/2603.07106v1


#### UG-T01-S36. Genie

Jake Bruce et al. Genie: Generative Interactive Environments. ICML 2024. arXiv:2402.15391 (https://arxiv.org/abs/2402.15391). Primary HTML inspected for model and output representation.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R36, starting at file section heading / rendered line 372. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: arXiv metadata; ICML attribution retained. Content inspected: abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Interactive world-model output differs from explicit game state, editable engine code, and testable resource accounting. Related: UG-T01-Q02.

Associated artifacts: Action-controllable world model, not conventional generated engine code.

Direct URLs:
https://arxiv.org/abs/2402.15391


#### UG-T01-S37. Player tactics in a commercial TD game

Stuart Anderson and Ruth Falconer. Emergence of player tactics by expert-guided machine learning: an industry tower defence case study. Entertainment Computing 54, article 100963, 2025. Institutional record (https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/), DOI:10.1016/j.entcom.2025.100963 (https://doi.org/10.1016/j.entcom.2025.100963). Institutional abstract and metadata inspected; the linked full-text download returned HTTP 403.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R37, starting at file section heading / rendered line 380. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: institutional metadata. Content inspected: institutional abstract only.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Observed tactic mining in an existing commercial TD game; not generated unit validation. Full methods remain inaccessible. Related: UG-T01-E17.

Associated artifacts: Commercial mobile TD network-message/tactic data; raw dataset not supplied.

Direct URLs:
https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/
https://doi.org/10.1016/j.entcom.2025.100963


#### UG-T01-S38. GameASG-Bench

Xiuhui Zhang et al. GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development. Preprint, September 18, 2026. arXiv:2609.21293v1 (https://arxiv.org/html/2609.21293v1). Inspected: evaluation contract, task suite and aggregate results.

Type/status: Research source; publication/preprint status is stated in the retained bibliographic record. Any unverified venue qualification is noted below. Origin: Already present in this thread through the supplied catalogue; see original R record. Locator/use: UG-T01-S84, R38, starting at file section heading / rendered line 388. Also used in the assistant's literature/adaptation discussion where applicable.

Bibliographic record: verified: indexed primary arXiv metadata and abstract; direct HTML access failed. Content inspected: abstract only in this consolidation.

Any "Inspected" or "Access gap" wording in the retained record is the source catalogue's report, not a new verification claim.

Relevance, linked records, and limits: Authors recovered: Xiuhui Zhang, Yi Chen, Shusheng Xu, Fan Li, Huan Wang, Tongkai Yang, Binhang Yuan. September 2026 preprint; 47 tasks and distinct aggregate/task-success maxima do not measure novelty or player preference. Related: UG-T01-E20.

Associated artifacts: Forty-seven browser-game tasks and static/runtime checks reported by the source; no tasks executed here.

Direct URLs:
https://arxiv.org/html/2609.21293v1
https://arxiv.org/abs/2609.21293


#### UG-T01-S39. Experience-Driven Procedural Content Generation

Georgios N. Yannakakis and Julian Togelius (2011). Experience-Driven Procedural Content Generation. IEEE Transactions on Affective Computing 2(3), 147-161. DOI: 10.1109/T-AFFC.2011.6.

Type/status: Published research article Origin: Already present in this thread Locator/use: First literature answer, paper 6; S84 additional references.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Specific desired player experience can guide evaluation. The initial answer described it as read, whereas the supplied catalogue records an access gap. No model of TD taste was demonstrated. Related: UG-T01-E13; UG-T01-Q10.

Direct URLs:
https://doi.org/10.1109/T-AFFC.2011.6
https://yannakakis.net/wp-content/uploads/2019/02/EDPCG.pdf
https://julian.togelius.com/Yannakakis2011Experience.pdf


#### UG-T01-S40. Individual creativity and collective diversity

Anil R. Doshi and Oliver P. Hauser (2024). Generative AI enhances individual creativity but reduces the collective diversity of novel content. Science Advances 10(28), eadn5290. DOI: 10.1126/sciadv.adn5290.

Type/status: Published research article Origin: Already present in this thread Locator/use: First literature answer, paper 8; S84 additional references.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Story-writing result used to motivate collection-level variety. Publisher access was blocked in this extraction; no TD transfer is verified. Related: UG-T01-E16.

Direct URLs:
https://doi.org/10.1126/sciadv.adn5290
https://pmc.ncbi.nlm.nih.gov/articles/PMC11244532/


#### UG-T01-S41. Homogeneously creative LLMs

Emily Wenger and Yoed N. Kenett (2026). Large language models are homogeneously creative. PNAS Nexus 5(3), pgag042. DOI: 10.1093/pnasnexus/pgag042.

Type/status: Published research article Origin: Already present in this thread Locator/use: First literature answer companion to paper 8; S84 additional references.

Bibliographic record: verified: publisher record. Content inspected: relevant sections in publisher HTML.

Relevance, linked records, and limits: The earlier catalogue's 403 is retained historically, but current primary HTML was accessible. Standardized verbal tasks, tested people/models, not a general guarantee about future models or game rosters. Related: UG-T01-E16.

Direct URLs:
https://academic.oup.com/pnasnexus/article/5/3/pgag042/8529001
https://doi.org/10.1093/pnasnexus/pgag042


#### UG-T01-S42. AI and novelty in human Go decisions

Minkyu Shin, Jin Kim, Bas van Opheusden and Thomas L. Griffiths (2023). Superhuman artificial intelligence can improve human decision-making by increasing novelty. PNAS 120(12), e2214840120. DOI: 10.1073/pnas.2214840120.

Type/status: Published article; arXiv version also available Origin: Already present in this thread Locator/use: First literature answer, paper 9; S84 additional references; early tool call also used arXiv 2303.07462.

Bibliographic record: verified: arXiv title/authors/journal record; publisher retrieval failed. Content inspected: abstract only.

Relevance, linked records, and limits: Historical observational Go analysis, approximately 5.8 million decisions in the earlier account. Decisions within an existing game, not creation of new unit rules or randomized proof of causation. Related: UG-T01-E16.

Direct URLs:
https://doi.org/10.1073/pnas.2214840120
https://arxiv.org/abs/2303.07462


#### UG-T01-S43. Boden creativity taxonomy

Margaret A. Boden (1998). Creativity and artificial intelligence. Artificial Intelligence 103(1-2), 347-356. DOI: 10.1016/S0004-3702(98)00055-1.

Type/status: Published conceptual article Origin: Already present in this thread Locator/use: First literature answer, inspired-but-novel section; S84 additional references.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Combination, exploration and transformation were used as vocabulary. Not an empirical test of LLMs, and primary text was not retrieved here. Related: UG-T01-E13.

Direct URLs:
https://doi.org/10.1016/S0004-3702(98)00055-1
https://www.sciencedirect.com/science/article/pii/S0004370298000551


#### UG-T01-S44. Program Synthesis as a Generative Method

Eric Butler, Kristin Siu and Alexander Zook (2017). Program Synthesis as a Generative Method. FDG 2017. DOI lead: 10.1145/3102071.3102076.

Type/status: Conference-paper lead Origin: Already present in this thread Locator/use: S84 additional citation leads, found in R1 references.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: DOI recovered through indexed secondary metadata, not a primary-text review. Potential generation of dynamic behavior; no adopted synthesis dependency. Related: UG-T01-Q16.

Direct URLs:
https://doi.org/10.1145/3102071.3102076


#### UG-T01-S45. Talakat

Ahmed Khalifa, Scott Lee, Andy Nealen and Julian Togelius (2018). Talakat: Bullet Hell Generation through Constrained MAP-Elites. GECCO 2018. DOI: 10.1145/3205455.3205470.

Type/status: Conference paper Origin: Already present in this thread Locator/use: S84 additional citation leads, R4 reference 18.

Bibliographic record: verified in part: indexed ACM DOI record. Content inspected: abstract only.

Relevance, linked records, and limits: Combat-pattern diversity under constraints. The short publisher heading Talakat and supplied expanded title are retained; not TD path or character-fidelity evidence. Related: UG-T01-E13.

Direct URLs:
https://doi.org/10.1145/3205455.3205470


#### UG-T01-S46. Hearthstone deck-space mapping

Michael C. Fontaine et al. (2019). Mapping Hearthstone Deck Spaces with MAP-Elites with Sliding Boundaries. GECCO 2019. Full author list and DOI not recovered in this extraction.

Type/status: Conference-paper lead Origin: Already present in this thread Locator/use: S84 additional citation leads, R4 reference 13.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: URL is the inspected parent survey, not a falsely identified paper landing page. Useful lead for interacting content collections; original experiments not rechecked. Related: UG-T01-E13.

Direct URLs:
https://arxiv.org/abs/1907.04053


#### UG-T01-S47. StarCraft map-space exploration

Julian Togelius et al. (2010). Multiobjective Exploration of the StarCraft Map Space. IEEE CIG 2010. Full metadata not recovered here.

Type/status: Conference-paper lead Origin: Already present in this thread Locator/use: S84 additional citation leads, R1 references.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Parent reference URL only. Conflicting map objectives, not an approved map-generation v1 task. Related: UG-T01-D11.

Direct URLs:
https://arxiv.org/abs/2212.03387


#### UG-T01-S48. Mixed-Initiative Co-Creativity

Georgios N. Yannakakis, Antonios Liapis and Constantine Alexopoulos (2014). Mixed-Initiative Co-Creativity. FDG 2014. DOI not established here.

Type/status: Conference-paper lead Origin: Already present in this thread Locator/use: S84 additional citation leads, GAVEL reference 68.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Parent reference URL only. Designer control and revision; original text not independently reviewed here. Related: UG-T01-Q13.

Direct URLs:
https://arxiv.org/html/2407.09388v2


#### UG-T01-S49. Self-Refine

Aman Madaan, Niket Tandon, Prakhar Gupta, Skyler Hallinan, Luyu Gao, Sarah Wiegreffe, Uri Alon, Nouha Dziri, Shrimai Prabhumoye, Yiming Yang, Shashank Gupta, Bodhisattwa Prasad Majumder, Katherine Hermann, Sean Welleck, Amir Yazdanbakhsh and Peter Clark (2023). Self-Refine: Iterative Refinement with Self-Feedback. arXiv:2303.17651v2.

Type/status: Research paper; preprint record inspected, final venue not reverified here Origin: Already present in this thread Locator/use: Assistant review of skill, section about a separate prose-edit pass.

Bibliographic record: verified: arXiv metadata. Content inspected: abstract only.

Relevance, linked records, and limits: Reports gains from iterative self-feedback on its tasks. Not evidence that another pass preserves this unit's mechanics. Related: UG-T01-E18; UG-T01-D12.

Associated artifacts: Code/data/demo linked by authors at selfrefine.info; not installed or audited.

Direct URLs:
https://arxiv.org/abs/2303.17651v2
https://selfrefine.info/


#### UG-T01-S50. Limits of intrinsic self-correction

Jie Huang, Xinyun Chen, Swaroop Mishra, Huaixiu Steven Zheng, Adams Wei Yu, Xinying Song and Denny Zhou (2024; preprint 2023). Large Language Models Cannot Self-Correct Reasoning Yet. ICLR 2024; arXiv:2310.01798v2.

Type/status: Conference paper with preprint Origin: Already present in this thread Locator/use: Assistant review of skill, contrasted with Self-Refine.

Bibliographic record: verified: primary arXiv metadata. Content inspected: relevant introductory and task-scope sections.

Relevance, linked records, and limits: Examines intrinsic correction without external feedback; failures are conditional on studied tasks/models. It does not contradict every external-feedback or editing workflow. Related: UG-T01-E18.

Direct URLs:
https://arxiv.org/abs/2310.01798v2
https://arxiv.org/html/2310.01798


#### UG-T01-S51. FollowBench

Yuxin Jiang, Yufei Wang, Xingshan Zeng, Wanjun Zhong, Liangyou Li, Fei Mi, Lifeng Shang, Xin Jiang, Qun Liu and Wei Wang (2024). FollowBench: A Multi-level Fine-grained Constraints Following Benchmark for Large Language Models. ACL 2024, 4667-4688. DOI: 10.18653/v1/2024.acl-long.257.

Type/status: Published conference paper and benchmark Origin: Already present in this thread Locator/use: Assistant skill evaluation discussion.

Bibliographic record: verified: ACL record. Content inspected: abstract only.

Relevance, linked records, and limits: Constraint following distinguished from overall quality; LLM-based checking itself needs scrutiny. Benchmark/data not executed here. Related: UG-T01-E18.

Direct URLs:
https://aclanthology.org/2024.acl-long.257/
https://github.com/YJiangcm/FollowBench


#### UG-T01-S52. Checklist feedback

Vijay Viswanathan, Yanchao Sun, Shuang Ma, Xiang Kong, Meng Cao, Graham Neubig and Tongshuang Wu (2025). Checklists Are Better Than Reward Models For Aligning Language Models. NeurIPS 2025; arXiv:2507.18624v2.

Type/status: Conference paper with preprint Origin: Already present in this thread Locator/use: Assistant skill-testing discussion.

Bibliographic record: verified: arXiv metadata and venue comment. Content inspected: abstract only.

Relevance, linked records, and limits: RLCF combines instruction-specific checklist judgments/verifiers into RL rewards. It does not show that pasting this checklist into one prompt has the same effect. Related: UG-T01-E18.

Direct URLs:
https://arxiv.org/abs/2507.18624v2


#### UG-T01-S53. Agent Skills documentation

Agent Skills specification and overview. Maintainer documentation, living pages; no paper or fixed model claim.

Type/status: Technical specification Origin: Already present in this thread Locator/use: Skill versus agent and packaging discussion; S85 section 8.

Bibliographic record: verified: official documentation. Content inspected: relevant format and packaging sections.

Relevance, linked records, and limits: SKILL.md, references and scripts describe packaging. The host must load the material; format conformance is not generation quality. Related: UG-T01-D12; UG-T01-E20.

Direct URLs:
https://agentskills.io/specification
https://agentskills.io/what-are-skills


#### UG-T01-S54. Building effective agents

Anthropic (2024). Building effective agents. Engineering guidance, not a peer-reviewed comparative result for this project.

Type/status: Vendor engineering article Origin: Already present in this thread Locator/use: Assistant bounded-workflow recommendation.

Bibliographic record: verified: official article. Content inspected: relevant workflow-versus-agent and simplicity guidance.

Relevance, linked records, and limits: Supports a practitioner distinction between programmed workflows and model-controlled tool use. No proof of the optimal Unit-Generator architecture. Related: UG-T01-D12.

Direct URLs:
https://www.anthropic.com/engineering/building-effective-agents


#### UG-T01-S55. JSON Schema

JSON Schema, Understanding JSON Schema: Reference. Living documentation.

Type/status: Technical documentation Origin: Already present in this thread Locator/use: Schema/ruleset discussions and S84 framework references.

Bibliographic record: verified: official documentation. Content inspected: reference overview and relevant structural descriptions.

Relevance, linked records, and limits: Structure and constraints, not gameplay semantics or taste. No schema version or library was chosen by this thread. Related: UG-T01-E20; UG-T01-Q16.

Direct URLs:
https://json-schema.org/understanding-json-schema/reference


#### UG-T01-S56. Zod JSON Schema guide

Zod, JSON Schema guide. Living documentation.

Type/status: Technical documentation Origin: Already present in this thread Locator/use: S84 framework references; earlier repository contracts use Zod according to the historical audit.

Bibliographic record: verified: official documentation. Content inspected: relevant conversion/coverage sections.

Relevance, linked records, and limits: Structural conversion support does not establish preservation of a generated mechanic or executable compatibility. Related: UG-T01-E20.

Direct URLs:
https://zod.dev/json-schema


#### UG-T01-S57. CEL overview

Common Expression Language, CEL overview. Living official documentation.

Type/status: Technical documentation Origin: Already present in this thread Locator/use: Ruleset expressions suggested in framework discussion; S84 framework references.

Bibliographic record: verified: official documentation. Content inspected: relevant parsing/checking/evaluation overview.

Relevance, linked records, and limits: A possible bounded expression evaluator against host bindings. It remains an option, not an adopted project dependency or implementation of new combat behavior. Related: UG-T01-E20; UG-T01-Q16.

Direct URLs:
https://cel.dev/
https://cel.dev/overview/cel-overview


#### UG-T01-S58. Open Source Definition

Open Source Initiative, The Open Source Definition.

Type/status: License criteria / organizational primary source Origin: Already present in this thread Locator/use: FOSS versus non-commercial restriction discussion.

Bibliographic record: verified: official document. Content inspected: relevant redistribution and fields-of-endeavor clauses.

Relevance, linked records, and limits: Supports separating FOSS licensing from third-party rights. Does not select a license or clear BTD6/anime content. Related: UG-T01-D04; UG-T01-Q09.

Direct URLs:
https://opensource.org/osd


#### UG-T01-S59. MIT license text

Open Source Initiative, The MIT License.

Type/status: License text Origin: Already present in this thread Locator/use: Assistant gave MIT as an example, not an owner choice.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: No project selection or compatibility audit follows from this reference. Related: UG-T01-D04.

Direct URLs:
https://opensource.org/license/mit


#### UG-T01-S60. Apache License 2.0 text

Open Source Initiative, Apache License 2.0.

Type/status: License text Origin: Already present in this thread Locator/use: Earlier licensing research tool lookup.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Retained lookup, not a selected license or a verified dependency audit. Related: UG-T01-D04.

Direct URLs:
https://opensource.org/license/apache-2-0


#### UG-T01-S61. Free software definition

Free Software Foundation, What is free software?

Type/status: Organizational definition Origin: Already present in this thread Locator/use: Earlier FOSS licensing discussion/lookup.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: No conclusion about third-party content rights follows from software freedoms. Related: UG-T01-D04.

Direct URLs:
https://www.gnu.org/philosophy/free-sw.html


#### UG-T01-S62. GPL tools and output FAQ

Free Software Foundation, GNU GPL FAQ, Can I use GPL-covered editors/tools to develop nonfree programs? Original anchor CanIUseGPLToolsForNF.

Type/status: License FAQ Origin: Already present in this thread Locator/use: Earlier output-versus-tool licensing lookup.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: FAQ concerns the stated licensing situation, not a general output license or legal clearance for generated characters. Related: UG-T01-D04.

Direct URLs:
https://www.gnu.org/licenses/gpl-faq.html#CanIUseGPLToolsForNF


#### UG-T01-S63. EU software directive

European Parliament and Council (2009). Directive 2009/24/EC on the legal protection of computer programs. CELEX:32009L0024.

Type/status: EU legislation Origin: Already present in this thread Locator/use: EU licensing/private-use discussion.

Bibliographic record: verified: official EUR-Lex text. Content inspected: Article 1(2) and relevant recitals.

Relevance, linked records, and limits: Distinguishes expression from underlying ideas/principles under this directive. This does not decide the legality of the project's specific content or national exceptions. Related: UG-T01-D04; UG-T01-Q09.

Direct URLs:
https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX:32009L0024


#### UG-T01-S64. EU information-society copyright directive

European Parliament and Council (2001). Directive 2001/29/EC. CELEX:32001L0029.

Type/status: EU legislation Origin: Already present in this thread Locator/use: Private-copying and distribution discussion.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Current retrieval did not expose substantive text. Historical Article 5 discussion remains attributed; no blanket private-use conclusion is adopted. Related: UG-T01-D04; UG-T01-Q09.

Direct URLs:
https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX:32001L0029
https://eur-lex.europa.eu/legal-content/EN/TXT/HTML/?uri=CELEX:02001L0029-20190606


#### UG-T01-S65. SAS Institute v World Programming

Court of Justice of the European Union (2012), SAS Institute Inc. v World Programming Ltd, C-406/10. CELEX:62010CJ0406.

Type/status: Judgment / historical lookup Origin: Already present in this thread Locator/use: Software-functionality copyright research.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: No present full judgment review or project-specific legal holding claimed. Related: UG-T01-Q09.

Direct URLs:
https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX:62010CJ0406
https://juris.curia.europa.eu/juris/document/document.jsf?docid=122362&doclang=EN


#### UG-T01-S66. Nintendo v PC Box

Court of Justice of the European Union (2014), Nintendo and Others v PC Box and 9Net, C-355/12. CELEX:62012CJ0355.

Type/status: Judgment / historical lookup Origin: Already present in this thread Locator/use: Game-content and software-rights distinction.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Retain source identity; no current analysis of its applicability to a generator. Related: UG-T01-Q09.

Direct URLs:
https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX:62012CJ0355
https://juris.curia.europa.eu/juris/document/document.jsf?docid=146686&doclang=EN


#### UG-T01-S67. ACI Adam

Court of Justice of the European Union (2014), ACI Adam and Others, C-435/12. CELEX:62012CJ0435.

Type/status: Judgment / historical lookup Origin: Already present in this thread Locator/use: Private-copying and lawful-source discussion.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: No private-use clearance is inferred and underlying judgment was not reinspected here. Related: UG-T01-Q09.

Direct URLs:
https://eur-lex.europa.eu/LexUriServ/LexUriServ.do?uri=CELEX:62012CJ0435:EN:HTML


#### UG-T01-S68. German Copyright Act section 23

Germany, Urheberrechtsgesetz, section 23. Living statutory text.

Type/status: National legislation Origin: Already present in this thread Locator/use: Earlier adaptation/copyright lookup.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: No choice of governing national law or legal assessment made in this handover. Related: UG-T01-Q09.

Direct URLs:
https://www.gesetze-im-internet.de/urhg/__23.html


#### UG-T01-S69. US Copyright Office guidance

U.S. Copyright Office, What is Copyright?, What Does Copyright Protect?, and Fair Use Index.

Type/status: Government guidance pages Origin: Already present in this thread Locator/use: Earlier assistant used US expression/ideas and non-commercial-use discussion before the owner emphasized EU law.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Different jurisdiction. Preserve the historical references, but do not substitute them for an EU assessment. Related: UG-T01-Q09.

Direct URLs:
https://www.copyright.gov/what-is-copyright/
https://www.copyright.gov/help/faq/faq-protect.html
https://www.copyright.gov/fair-use/


#### UG-T01-S70. GitHub repository licensing documentation

GitHub Docs, Licensing a repository.

Type/status: Technical/legal-administration documentation Origin: Already present in this thread Locator/use: Earlier assistant could not identify an actual project license in the inspected snapshot.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Public visibility is not itself a license decision; no current checkout license audit performed. Related: UG-T01-D04; UG-T01-Q09.

Direct URLs:
https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/licensing-a-repository


#### UG-T01-S71. Bloons TD 6 product and reference pages

Ninja Kiwi / Steam, Bloons TD 6 product listing; support search and update lookups.

Type/status: Game/product reference Origin: Already present in this thread Locator/use: Owner's progression inspiration; assistant used the product page for three paths and Paragons; skill also names Dart, Super Monkey, Engineer, Desperado and Ice patterns.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: No exact versioned mechanics/conformance specification or complete official crosspath/cost review was established. Named tower examples are owner-supplied patterns, not newly verified claims. Related: UG-T01-D03; UG-T01-Q04.

Direct URLs:
https://store.steampowered.com/app/960090/Bloons_TD_6/
https://support.ninjakiwi.com/hc/en-us/search?query=paragon


#### UG-T01-S72. Official Luffy profile

ONE PIECE official character profile: Monkey D. Luffy.

Type/status: Character reference Origin: Already present in this thread Locator/use: Character-grounding research and S84 framework references.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Last URL was an attempted alternate path; successful access varied historically. Not a complete through-Wano source dossier or verification of invented TD applications. Related: UG-T01-E02; UG-T01-Q11.

Direct URLs:
https://one-piece.com/character/luffy/index.html
https://one-piece.com/character/luffy/index.html?lang=en
https://one-piece.com/character/monkey_d_luffy/index.html


#### UG-T01-S73. Katakuri future-sight feature

Crunchyroll (2020). FEATURE: How Luffy vs. Katakuri Puts an Interesting Spin on Anime Power-Ups. Author not recovered here.

Type/status: Licensed-media editorial/source lead Origin: Already present in this thread Locator/use: Assistant Observation-Haki interpretation; S84 labels supplied feature an unverified follow-up.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Original wording and paths retained. No source fact is promoted from a failed lookup. Related: UG-T01-D06; UG-T01-Q11.

Direct URLs:
https://www.crunchyroll.com/news/features/2020/9/18/feature-how-luffy-vs-katakuri-puts-an-interesting-spin-on-anime-power-ups
https://www.crunchyroll.com/pl/news/features/2020/9/18/feature-how-luffy-vs-katakuri-puts-an-interesting-spin-on-anime-power-ups


#### UG-T01-S74. Other character-source lookups

Official episode pages and alternate source pages attempted during earlier Luffy research; not a verified combined canon dossier.

Type/status: Lookup inventory / mixed source types Origin: Already present in this thread Locator/use: Visible old tool calls; many returned bodies are redacted.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Wikipedia entries are secondary leads, not primary canon. Page identity/access and any factual use remain unverified here. Decorative image search did not establish research evidence. Related: UG-T01-Q11.

Direct URLs:
https://one-piece.com/anime/516/index.html
https://one-piece.com/anime/726/index.html
https://one-piece.com/anime/869/index.html
https://one-piece.com/anime/1028/index.html
https://www.toei-anim.co.jp/tv/onep/episode/summary/516/
https://www.toei-anim.co.jp/tv/onep/episode/summary/869/
https://www.toei-anim.co.jp/tv/onep/episode/summary/1071/
https://www.toei-animation.com/catalog/one-piece/
https://www.netflix.com/tudum/articles/one-piece-cast
https://en.bandainamcoent.eu/one-piece/news/one-piece-odyssey-meet-the-straw-hat-crew
https://oppw4-20-en.bn-ent.net/character/luffy/
https://en.onepiece-base.com/characters/monkey-d-luffy
https://en.wikipedia.org/wiki/Monkey_D._Luffy
https://en.wikipedia.org/wiki/Luffy%27s_Peak_-_Attained%21_Fifth_Gear


#### UG-T01-S75. Unit-Generator repository and historical main-file audit

Mardwerk, unit-generator repository. Historical inspected starter revision 2026-09-21-design-v9; no immutable main SHA obtained for the first handoff.

Type/status: Repository / source and documentation Origin: Already present in this thread Locator/use: Owner repository link; S85 sections 2-4 and source list.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: available prior assistant audit and handoff only; current source not inspected.

Relevance, linked records, and limits: Dynamic main URLs are locators, not evidence of current code. No build, tests or live model calls performed for this consolidation. Historical numerical coverage and early-upgrade restrictions require checking against the requested revision. Related: UG-T01-E05; UG-T01-E06.

Direct URLs:
https://github.com/mardwerk/unit-generator
https://raw.githubusercontent.com/mardwerk/unit-generator/main/docs/API.md
https://raw.githubusercontent.com/mardwerk/unit-generator/main/src/core/schemas.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/main/src/core/draft.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/main/src/core/default-profile.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/main/src/core/check.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/main/src/core/mechanics/schemas.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/main/docs/MECHANICS.md
https://raw.githubusercontent.com/mardwerk/unit-generator/main/docs/PIPELINE-EVALUATION.md
https://raw.githubusercontent.com/mardwerk/unit-generator/main/docs/TESTING.md
https://raw.githubusercontent.com/mardwerk/unit-generator/main/CONTEXT.md
https://raw.githubusercontent.com/mardwerk/unit-generator/main/research/game-design/CHARACTER-DESIGN-GUIDANCE.md
https://raw.githubusercontent.com/mardwerk/unit-generator/main/AGENTS.md
https://raw.githubusercontent.com/mardwerk/unit-generator/main/package.json


#### UG-T01-S76. PR #10 and interpretation experiment

Mardwerk Unit-Generator PR #10. Historical observations at c7216f3 and later c97fdad8cb052bafa37f6dc92de3179f8ba2f3d3; owner later reports merge into 739cae3.

Type/status: Pull request / reported code review Origin: Already present in this thread Locator/use: Earlier and revised assistant reviews; owner merge update.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: historical reports and available handoff; commit source not reinspected.

Relevance, linked records, and limits: Do not carry historical defects forward as verified current defects. Private-file introduction/removal is a disclosure lead, not an audited history rewrite. Related: UG-T01-D14; UG-T01-E07; UG-T01-Q07.

Direct URLs:
https://github.com/mardwerk/unit-generator/pull/10
https://github.com/mardwerk/unit-generator/pull/10/commits
https://github.com/mardwerk/unit-generator/tree/c97fdad8cb052bafa37f6dc92de3179f8ba2f3d3
https://raw.githubusercontent.com/mardwerk/unit-generator/c97fdad/src/core/design.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/c97fdad/src/core/rulepack.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/c97fdad/src/core/prepare.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/c97fdad/src/core/blueprint/plan.ts
https://raw.githubusercontent.com/mardwerk/unit-generator/c97fdad/tests/rulepack-layout.test.ts
https://github.com/mardwerk/unit-generator/blob/c97fdad8cb052bafa37f6dc92de3179f8ba2f3d3/docs/RULEPACK-DESIGN-PLAN.md


#### UG-T01-S77. PR #6 compact-generation experiment

Mardwerk Unit-Generator PR #6. Earlier visible head described as 3e5d73b; owner reports adapted merge in 739cae3.

Type/status: Pull request / historical review Origin: Already present in this thread Locator/use: Assistant recipe-route discussion and owner implementation update.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: historical reports only.

Relevance, linked records, and limits: Keep opening recipe description separate from later evidence/simulator changes and adapted merge. No current comparative generation result inspected. Related: UG-T01-D15; UG-T01-E08.

Direct URLs:
https://github.com/mardwerk/unit-generator/pull/6


#### UG-T01-S78. Issue #4

Mardwerk Unit-Generator issue #4, larger progression support as described by the historical handoff. Exact current title not recovered.

Type/status: Issue / scope reference Origin: Already present in this thread Locator/use: S85 section 11 and initial repository response.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: historical handoff only.

Relevance, linked records, and limits: Distinguishes presentation/layout support from runtime or numerical support; not proof either is complete. Related: UG-T01-Q01.

Direct URLs:
https://github.com/mardwerk/unit-generator/issues/4


#### UG-T01-S79. Issue #9

Mardwerk Unit-Generator issue #9, unrelated cleanup as described by the historical handoff. Exact current title not recovered.

Type/status: Issue reference Origin: Already present in this thread Locator/use: S85 section 11.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: historical handoff only.

Relevance, linked records, and limits: Retained for completeness; no scope requirement inferred. Related: UG-T01-D18.

Direct URLs:
https://github.com/mardwerk/unit-generator/issues/9


#### UG-T01-S80. Owner-reported main update 739cae3

Owner implementation report: main at 739cae3; clean working tree; 436 tests and other checks passed; four live attempts yielded no accepted designs.

Type/status: Primary owner report, not a raw code/test artifact Origin: Already present in this thread Locator/use: Owner message beginning "Implemented and pushed to main at 739cae3."

Bibliographic record: verified as a visible owner report; code/test claims not independently verified. Content inspected: full owner message.

Relevance, linked records, and limits: Commit URL is a locator derived from the owner's hash, not a claim that the commit was fetched. Earlier inaccessible/stale pages do not disprove the report. Related: UG-T01-D13 through UG-T01-D18; UG-T01-E08 through UG-T01-E10.

Direct URLs:
https://github.com/mardwerk/unit-generator/commit/739cae3


#### UG-T01-S81. PR10_REVIEW on main

Standalone document named PR10_REVIEW by the owner. Exact repository path and file bytes not available.

Type/status: Missing requested review artifact Origin: Already present in this thread Locator/use: Owner asked to inspect updated PR #10 and "a PR10_REVIEW on the main branch."

Bibliographic record: unverified: artifact not retrieved. Content inspected: not inspected.

Relevance, linked records, and limits: Do not invent a .md path or equate a response-to-review document with this original review. Its issues/resolutions cannot be certified from this conversation. Related: UG-T01-E07; UG-T01-Q14.

Direct URL: none available. Use the exact message or local-file locator above.


#### UG-T01-S82. AI Index 2025 price trend

Stanford HAI (2025), AI Index Report, Research and Development webpage, item 6.

Type/status: Institutional report summary Origin: Already present in this thread Locator/use: Assistant response to owner's model-progress concern.

Bibliographic record: verified: official page. Content inspected: relevant webpage price passage; full PDF not inspected.

Relevance, linked records, and limits: Benchmark-specific $20 to $0.07 per-million-token comparison. Do not inherit the webpage's approximate elapsed-month wording without checking dates; no future unit-design cost guarantee follows. Related: UG-T01-E19.

Direct URLs:
https://hai.stanford.edu/ai-index/2025-ai-index-report/research-and-development


#### UG-T01-S83. AI Index 2026

Stanford HAI (2026), AI Index Report.

Type/status: Institutional report overview Origin: Already present in this thread Locator/use: Assistant response to model-progress concern.

Bibliographic record: verified: official overview page. Content inspected: overview only, not full report.

Relevance, linked records, and limits: General capability trends are not evidence that this exact unit-design task will be solved within months. Related: UG-T01-E19.

Direct URLs:
https://hai.stanford.edu/ai-index/2026-ai-index-report


#### UG-T01-S84. Attached research catalogue

Research sources for unit generation and game design. Supplied Markdown document, internally dated September 22, 2026. Authorship not independently established by upload.

Type/status: User-uploaded curated secondary research record Origin: Already present in this thread Locator/use: Attachment uploaded when owner asked what useful material was still missed.

Bibliographic record: verified: exact available file and its internal date, not every embedded claim. Content inspected: full supplied text.

Relevance, linked records, and limits: Local file: markdown(1).md eingefügt. Contains R1-R38, additional references, citation trails and labeled Adaptations. Explicitly not systematic or an implementation commitment. Project-relative links are not proof of access to those projects. Related: UG-T01-E11 through UG-T01-E20.

Direct URL: none available. Use the exact message or local-file locator above.


#### UG-T01-S85. Earlier implementation handoff

Unit Generator: implementation and research handoff. Assistant-generated Markdown, internally dated September 22, 2026.

Type/status: Proposed internal handoff / historical secondary audit Origin: Already present in this thread Locator/use: Assistant response to positive sample and repository request.

Bibliographic record: verified: available attachment; proposals not adopted by title. Content inspected: full supplied file.

Relevance, linked records, and limits: Local file: unit-generator-handoff.md. Preserve reported inspection limits and mutable URLs. Owner later reports deleting repository handoffs; this conversation copy remains available. Related: UG-T01-E05; UG-T01-E06; UG-T01-E11.

Direct URL: none available. Use the exact message or local-file locator above.


#### UG-T01-S86. Frozen skill and positive example

Unit Generator: frozen prompt and positive example. Assistant-produced normalized copy of owner-provided text, internally dated September 22, 2026.

Type/status: Prompt, reference brief and generated-design fixture Origin: Already present in this thread Locator/use: Original user skill and later positive sample; attachment sections A-D.

Bibliographic record: verified: available text and correspondence to visible messages; generation provenance not independently verified. Content inspected: full supplied file.

Relevance, linked records, and limits: Local file: unit-generator-prototype-baseline.md. Escaping normalized. Does not establish model/version, runtime, balance or public redistribution rights. Related: UG-T01-D07; UG-T01-D08; UG-T01-E01; UG-T01-E02.

Direct URL: none available. Use the exact message or local-file locator above.


#### UG-T01-S88. Proposed vocabulary replacement

Unit Generator context, CONTEXT.proposed.md. Assistant proposal produced in the last architecture response.

Type/status: Proposed contributor contract Origin: Already present in this thread Locator/use: Assistant response after the owner supplied current CONTEXT.md and 739cae3 report.

Bibliographic record: verified: available proposal, not adopted contract. Content inspected: full supplied file.

Relevance, linked records, and limits: Local file: CONTEXT.proposed.md. Additional terms and compatibility rules require owner approval. Do not claim it is the current repository file. Related: UG-T01-D20; UG-T01-Q15.

Direct URL: none available. Use the exact message or local-file locator above.


#### UG-T01-S89. Owner-supplied CONTEXT.md

Unit Generator context, pasted verbatim by the owner in the message reporting 739cae3.

Type/status: Primary supplied contributor contract Origin: Already present in this thread Locator/use: Owner: "my current CONTEXT.md is:" followed by two term tables and ownership paragraph.

Bibliographic record: verified: visible owner-supplied text; live file identity not verified. Content inspected: full pasted text.

Relevance, linked records, and limits: The pasted definitions, not S88, establish the historical vocabulary supplied by the owner. Foundation/Towerright private files are not available by implication. Related: UG-T01-D19; UG-T01-Q15.

Direct URL: none available. Use the exact message or local-file locator above.


#### UG-T01-S90. Visible conversation messages and earlier summaries

Original visible user messages, assistant responses, visible calculations, and redacted tool-call traces in UG-T01. Dates and message URLs not provided for most original turns.

Type/status: Primary conversation record plus attributed assistant proposals/reports Origin: Already present in this thread Locator/use: Use speaker plus distinctive quotations in sections 1-6. Examples: "tasteless and basic"; "NON-COMMERCIAL PRIVATE USE ONLY"; "ONLY A PROTOTYPE"; "this quality would be nearly enough"; "Implemented and pushed to main at 739cae3"; "postone the research part".

Bibliographic record: verified: statements visibly present; underlying unavailable tool results remain unverified. Content inspected: available conversation text.

Relevance, linked records, and limits: Other-thread summaries are not original UG-T01 evidence. Older assistant source-check claims do not substitute for primary text, and no timestamps are invented. Related: UG-T01-D01 through UG-T01-D20; UG-T01-E01 through UG-T01-E11.

Direct URL: none available. Use the exact message or local-file locator above.


#### UG-T01-S91. Cooking Up Creativity

Moran Mizrahi, Chen Shani, Gabriel Stanovsky, Dan Jurafsky and Dafna Shahaf (2025). Cooking Up Creativity: Enhancing LLM Creativity through Structured Recombination. arXiv:2504.20643. Current v2 comment: accepted at TACL, pre-MIT Press publication version.

Type/status: Research paper / preprint; original v1 URL was used in an early tool call Origin: Already present in this thread Locator/use: Early literature search opened arXiv:2504.20643v1; no clear claim tied to it survived the redacted result.

Bibliographic record: verified: current arXiv metadata; original v1 content not recovered. Content inspected: abstract only.

Relevance, linked records, and limits: DishCOVER structured recombination in the culinary domain. This is an already-present lookup recovered during consolidation, not a newly found thread source or proof of TD creativity. Related: UG-T01-E13; UG-T01-Q12.

Direct URLs:
https://arxiv.org/html/2504.20643v1
https://arxiv.org/abs/2504.20643v2


#### UG-T01-S92. PICARD

Torsten Scholak, Nathan Schucher and Dzmitry Bahdanau (2021). PICARD: Parsing Incrementally for Constrained Auto-Regressive Decoding from Language Models. EMNLP 2021, 9895-9901. DOI: 10.18653/v1/2021.emnlp-main.779.

Type/status: Published conference paper Origin: Already present in this thread Locator/use: An earlier framework-research tool call opened this exact ACL record; tool output is redacted in the historical transcript.

Bibliographic record: verified: ACL bibliographic record. Content inspected: abstract only.

Relevance, linked records, and limits: Constrained decoding for SQL, not semantic correctness or design quality. No adoption or implementation shown. Related: UG-T01-E20; UG-T01-Q16.

Associated artifacts: Author-linked code, Spider/CoSQL tasks and fine-tuned T5 context; no code run or license audit.

Direct URLs:
https://aclanthology.org/2021.emnlp-main.779/
https://github.com/ElementAI/picard


#### UG-T01-S93. Ninja Kiwi terms and fan-content searches

Ninja Kiwi terms/terms-of-service and support fan-content searches. No permission document or successful rights clearance recovered.

Type/status: Unresolved policy/source lookups Origin: Already present in this thread Locator/use: Initial licensing discussion.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Do not infer a fan-use license, private-use safe harbor, or right to redistribute exact content from an attempted lookup. Related: UG-T01-Q09.

Direct URLs:
https://ninjakiwi.com/terms
https://ninjakiwi.com/terms-of-service
https://support.ninjakiwi.com/hc/en-us/search?query=fan%20content


#### UG-T01-S94. Other incomplete research and project-relative leads

Conceptual blending, play-trace similarity, quantitative balance, game metrics without players, program synthesis, and unnamed source references appear as leads without recoverable full citations. Project-relative documents also appear in S84.

Type/status: Incomplete lead inventory Origin: Already present in this thread Locator/use: S84 citation trails and project links; older opaque/redacted web citations.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Preserve as unresolved rather than inventing authors/identifiers. Project-relative paths: unit-generator/docs/PRODUCT.md; towerright/CONTEXT.md; unit-generator/research/GENERATION-APPROACHES.md; unit-generator/research/FRAMEWORK-SOURCES.md; unit-generator/docs/GENERATOR-RESHAPE.md; unit-generator/docs/REFINEMENT.md; unit-generator/docs/PR6-REVIEW.md. Their content was not supplied by those links. Private Foundation and Towerright are not assumed accessible. Related: UG-T01-Q10; UG-T01-Q14.

Direct URL: none available. Use the exact message or local-file locator above.


#### UG-T01-S95. WIPO legal-text mirrors

WIPO Lex records and an attempted legal-text PDF used in the earlier EU discussion. Record-to-statute mapping is not independently revalidated here.

Type/status: Legal source mirrors / unresolved lookup provenance Origin: Already present in this thread Locator/use: Earlier legal research calls.

Bibliographic record: unverified in this consolidation; retained from the conversation. Content inspected: not inspected.

Relevance, linked records, and limits: Retain original identifiers without substituting them for verified directive text. No PDF was inspected during this consolidation. Related: UG-T01-Q09.

Direct URLs:
https://www.wipo.int/wipolex/en/legislation/details/1453
https://www.wipo.int/wipolex/en/legislation/details/8612
https://wipolex-res.wipo.int/edocs/lexdocs/laws/en/eu/eu309en_1.pdf

### B. Newly identified literature leads

These two papers fill specific gaps in prompt and reviewer evaluation. They are new to this consolidation, not historical support for decisions in UG-T01. They neither justify a new mandatory agent stage nor establish an error in the current implementation.

#### UG-T01-S96. Lost in the Middle

Nelson F. Liu, Kevin Lin, John Hewitt, Ashwin Paranjape, Michele Bevilacqua, Fabio Petroni and Percy Liang (2024; preprint 2023). Lost in the Middle: How Language Models Use Long Contexts. Transactions of the Association for Computational Linguistics 12, 157-173. DOI: 10.1162/tacl_a_00638; arXiv:2307.03172.

Type/status: Research paper / new literature lead Origin: Newly found, not part of the original conversation. Locator/use: Identified during this consolidation, September 24, 2026.

Bibliographic record: verified: primary bibliographic records. Content inspected: relevant introduction, task and results sections in publisher text.

Relevance, linked records, and limits: Fills a gap in assumptions about the long Skill: QA/key-value experiments found sensitivity to information position. This motivates testing prompt assembly and placement rather than assuming longer context preserves every rule. Older tasks/models, not proof that the supplied skill fails or that compression improves TD designs. Related: UG-T01-E21; UG-T01-Q10; UG-T01-Q12.

Direct URLs:
https://aclanthology.org/2024.tacl-1.9/
https://direct.mit.edu/tacl/article/doi/10.1162/tacl_a_00638/119630/Lost-in-the-Middle-How-Language-Models-Use-Long
https://arxiv.org/abs/2307.03172


#### UG-T01-S97. Judging LLM-as-a-Judge

Lianmin Zheng, Wei-Lin Chiang, Ying Sheng, Siyuan Zhuang, Zhanghao Wu, Yonghao Zhuang, Zi Lin, Zhuohan Li, Dacheng Li, Eric P. Xing, Hao Zhang, Joseph E. Gonzalez and Ion Stoica (2023). Judging LLM-as-a-Judge with MT-Bench and Chatbot Arena. NeurIPS 2023 Datasets and Benchmarks; arXiv:2306.05685v4.

Type/status: Research paper / new literature lead Origin: Newly found, not part of the original conversation. Locator/use: Identified during this consolidation, September 24, 2026.

Bibliographic record: verified: primary bibliographic records. Content inspected: relevant sections 3.3-3.4 and abstract.

Relevance, linked records, and limits: Fills the reviewer-validation gap. The paper reports useful human agreement alongside position and verbosity vulnerabilities. Section 3.3 explicitly does not determine self-enhancement bias conclusively; do not turn a discussed possibility into a proven universal bias. Conversation-quality tests are not designer judgments about game units. Related: UG-T01-E21; UG-T01-Q10; UG-T01-Q12.

Associated artifacts: MT-Bench, expert votes and Chatbot Arena preference conversations linked by the paper; no dataset downloaded or used here.

Direct URLs:
https://arxiv.org/abs/2306.05685v4
https://arxiv.org/html/2306.05685v4

### Literature gaps and omissions

A later broader review should examine character-to-mechanics adaptation, semantic preservation through structured generation and editing, natural-language mechanic specification, designer preference measurement, current-model constraint following, and useful evaluation policies for real-time TD. The catalogue's paper-access gaps and version differences should be resolved before publication citations are exported.

The register is not an exhaustive literature review. It does not import every reference in every cited paper. Unnamed citation trails, source pages with unrecoverable opaque markers, search-result bodies that were redacted, and private project links remain unresolved under UG-T01-S94. Decorative image lookups, tool-installation suggestions, generic search strings, and website boilerplate are not research evidence and are omitted. Named substantive lookup URLs were retained even when they never supported a final claim. Newly identified bibliographic details do not convert an earlier lookup into a historically adopted method.


## 8. Preparation for later research

### Existing research plans in the thread

The owner considered building a working proof of concept, then conducting a separate research project that might produce a paper. The later explicit direction was to postpone research and finish the intended product (UG-T01-D17). No formal protocol, study freeze, publication venue, participant recruitment, or approved sample size is established.

The assistant suggested studying how representations, rules, interpretation choices, and repair affect source-based unit design. A possible title was "Preserving design intent in LLM-generated game units under replaceable rulesets." This is an assistant-proposed framing, not an established original contribution or an owner-approved paper title.

Other earlier proposals compared skill-only output, structured output, and bounded repair; separately tested prose editing; retained fixed source packets; separated generation quality from selector quality; and kept product, study, and private-project repositories distinct. The historical 54-output and 72-output pilot suggestions were illustrative and unapproved. They must not be used as existing study designs.

### Current research recommendation, explicitly new advice

Keep research deferred as a project phase. Retain enough operational evidence during v1 work to support later analysis without turning everyday debugging into a publication prerequisite.

**Research question suggested by this thread:** Under a stated game contract and character reference, which combinations of skill guidance, structure, and external checks improve designer usefulness and preserve behavior compared with competent direct LLM prompting?

A potential contribution is a reproducible adaptation and revision task with explicit rule variation and a record of what behavior was preserved, lost, or deliberately changed. That contribution's novelty is not established. It needs a later focused related-work review, including UG-T01-S09, S16, S49-S52, S91-S92 and the newly found S96-S97.

**Existing evidence versus hypotheses.** There is inspectable text for one positively judged skill-only concept (E01), inspectable skill instructions (E02), a small progression enumeration (E04), historical audit reports (E05-E07), and an owner-reported software update with unsuccessful live delivery (E08-E09). External papers establish results only in their own settings. Superior public-default quality, useful automatic interpretation, preserved semantics through every operation, reliable cross-ruleset generation, balanced play, and lower accepted-result cost remain hypotheses (E22).

### A fair baseline and smallest useful next evaluation

The next immediate evaluation should be diagnostic, not a paper study: inspect the four retained failed attempts; run the known positive concept through save/reload/render/edit operations under its intended rules; and compare a fresh direct-skill generation with the application route using the same fixed reference input and game constraints. This is proposed work. No new sample size or success threshold is approved here.

For a later comparison, retain a **competent direct LLM plus wiki/reference baseline**. Give it the same actual user goal, selected game rules, source scope, output requirements, and access to relevant facts. Do not handicap it with a one-line prompt while giving the framework a curated source packet and hundreds of extra instructions. The working long skill itself is an important strong baseline.

Two comparisons can answer different questions. An information-matched comparison isolates guidance, representation, checks, and repair. A practical name-only comparison includes retrieval quality and its costs. Report both input differences rather than describing one as the other. 

Record model/provider and exact model version, API or chat product, system and user messages, sampling and output limits, available tools, browsing/source snapshots, continuity, strategy, total calls, repair attempts, tokens, operational timeouts, and known cost. The producing configuration of the existing ChatGPT sample is unknown; it cannot be reconstructed from the prose alone. Compare at a disclosed resource budget and include failed attempts, not just selected outputs. No model-provider or data-transmission authorization is implied by this recommendation.

### Proposed measures, not approved success criteria

| Measure | What to observe | What it does not establish |
| --- | --- | --- |
| Constraint satisfaction | Correct progression, activation placement, explicit hit policies, and applicable required sections. | That the unit is appealing or balanced. |
| Source fidelity | Supported facts and coherent adaptation relationships; no invented canon. | Historical novelty or legal clearance. |
| Controllability | Compliance with changed Definitions, supplied preferences, and targeted revisions. | That a rule change is implemented by a Consumer. |
| Mechanic preservation | Triggers, independence, targeting, hit budgets, resources, and inheritance retained through editing/formalization. | Semantic equivalence beyond the tested representation. |
| Designer usefulness | Attributed preference, reasons, and correction effort for an intended use. | A universal taste score or representative player preference. |
| Diversity | Different useful behavior or purchase reasons across retained outputs, with a declared comparison basis. | Historical originality from names or embedding distances. |
| Operational usefulness | Completion/failure rates, time and model cost per usable or owner-accepted result, with the acceptance rule stated. | Lower prices for all future models or a positive result when the denominator is zero. |

Gameplay effectiveness or balance enters only after a compatible Consumer, defined rules/economy/maps/waves, competent policies, and actual tests exist. Stronger automated play and human comprehension/preference must remain separate outcomes. No numerical balance threshold is proposed here.

Independent review or blinded pairwise comparisons can help later. Reviewer familiarity with the character and TD games should be recorded. Repeated outputs share references and conditions; fifteen upgrades are not automatically fifteen independent observations. Do not treat an LLM review certificate as external evidence of its own correctness. New sources S96-S97 motivate checking instruction position and reviewer order/verbosity effects; they do not prescribe a new v1 stage.

### Evidence to retain from future work

Retain code revision; effective Game Definition/Profile and skill content or secured fingerprints; source snapshots and scope; exact Requests and model messages; model configuration; raw outputs; every repair; check implementation/version and findings; criteria for acceptance or rejection; owner/reviewer reasons; operational failures and termination state; and cost/usage when available. If gameplay is tested later, retain the Consumer version, maps, waves, policies, seeds and actual traces. Expected behavior must remain separate from observed behavior.

Separate development/tuning examples from held-out evaluation examples. The repeated Luffy case is a development reference, not sufficient held-out evidence. Keep private evaluations separate from public-default results.

A public-default study should be reproducible using public or permission-cleared reference packets and the public configuration. A Towerright-tuned study may rely on private rules, examples, model configuration, or assets; report that dependence and its reproducibility limit without assuming those assets must be released. Do not attribute a private-only result to the public default or require a private platform to reproduce a public claim.

## 9. Consolidation handoff


**Owner approval or conflict resolution remains needed:** the exact public Deliverable and default conventions, public Apex scope, the Definition/Profile/Core vocabulary, which checks are authoritative, the final generation Strategy, source/provider requirements, and artifact licensing/disclosure (Q01-Q04, Q06-Q09, Q14-Q16). The assistant's proposed architecture and CONTEXT.md are not approved by being included in this handover.

**Next concrete action:** inspect the retained four failed attempts at the reported revision, using their effective inputs, outputs, termination evidence, and findings; classify each failure before changing the architecture (Q03). In parallel, verify that the known positive concept survives the existing artifact operations and that a second Definition changes the real end-to-end flow (Q01, Q13). These are product-validation steps proposed here, not a study freeze or a promise that the tests have run.

This handover preserves historical evidence and proposed directions. It does not approve a final architecture, settle third-party rights, verify `739cae3`, or demonstrate the project's advantage over direct LLM prompting.
