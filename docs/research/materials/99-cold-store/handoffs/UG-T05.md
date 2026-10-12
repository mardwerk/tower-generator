# Generic generator research: UG-T05

Selected source excerpts retain their original identifiers, quotations, inspection status, and proposal labels. This split grants no publication or reuse rights. The private companion contains game-specific proposals that remain in Planning outside current game scope. Original source: `planning/99-cold-store/handoffs/UG-T05.md`.

## 1. Thread overview and coverage

**Thread ID:** UG-T05
**Title:** Character-coherent tower defense generation through replaceable rules

The original request was for research about AI taste, novel solutions, and novel decisions, especially in unit and game design. The conversation then examined how those ideas could support the workflow "Anime Name in" to "Full TowerDefense Unit out."

The thread's main contribution is a distinction between **obeying an upgrade structure** and **choosing a coherent interpretation of a character**. The first generated Luffy design satisfied many proposed structural requirements but was explicitly rejected by the owner. The subsequent discussion explored how a reusable generator might discover a more appropriate organization without hardcoding Luffy or his abilities.

### Coverage

Reviewed material consists of the visible user and assistant messages, their inline examples, and recoverable references from the visible research-tool traces. No implementation repository, uploaded specification, generated unit file, or test-results attachment was available in the reviewed material.

Several original web-tool responses are redacted. Their visible queries and URLs sometimes identify a source, but do not establish what was originally read. Status summaries and summaries of reasoning are not treated as independent evidence. Other Project conversation excerpts are not used as evidence of decisions in UG-T05.

Source verification performed for this handover is identified separately in section 7. It does not retrospectively prove that the original answers checked those sources adequately.

This is an internal record. Permission to publish its particular configuration, examples, or tuning material is **not established in this thread**.

### Central idea

A generator can follow every upgrade restriction and still produce an unsatisfying character adaptation. This thread suggests that the generator must first decide how the character's abilities belong together, then express that interpretation under the selected game's rules. The framework should discover such mappings without containing character-specific code. A public default should support the basic name-to-unit workflow, while a replaceable private ruleset can introduce additional systems. Whether this approach produces more useful units than a well-written direct LLM prompt remains untested.

The current owner-supplied product context governs the recommendations below. It is not recorded as historical agreement reached in this thread.

## 2. Ideas, approaches, and changes of direction

### Research-led generation and external evaluation

The initial assistant answer proposed separating idea generation, selection, and verification. It drew on procedural unit generation, game-rule evolution, quality-diversity search, preference modeling, and studies of LLM creativity.

The intended benefit was to avoid equating fluent descriptions with useful novelty. Its main dependencies were a representation of game behavior, appropriate evaluation criteria, and potentially simulations or human judgments.

**Disposition:** Explored, not adopted as a specific algorithm or architecture.

**Locator:** Assistant, opening research answer, "generating something unfamiliar," "judging whether it is worth keeping," and "demonstrating that it creates genuinely different gameplay."

### Complete upgrade trees built around distinct combat behaviors

The first adaptation proposed generating T3 specializations first, selecting three, then constructing auxiliary upgrades, later tiers, crosspaths, and prices. Its Luffy branches were Jet Rhythm, Giant Impact, and Elastic Geometry.

This addressed mechanical differentiation and compatibility with the proposed three-path system. The cost was conceptual: the branches organized combat functions more clearly than they organized the character.

**Disposition:** The concrete Luffy design was explicitly rejected. The rejection does not establish that all of its individual mechanics or all behavior-based generation are unsuitable.

**Locator:** Assistant, "Build this as a character-to-mechanics translation system that searches over complete upgrade trees." Owner, "It doesnt capture the spirit / has an actuall consistent figure."

### Character organization before individual upgrades


The assistant then reframed taste as deciding which abilities are parallel specializations, which are shared progression, and which limitations create an appropriate payoff.

This offers a coherent interpretation of the owner's example. It does not demonstrate that the generator can discover an equally appropriate organization for unfamiliar characters.



### Replaceable rules instead of a Luffy-specific framework

The owner clarified that the broader aim was a replaceable ruleset or DSL and the same framework for generalizable concepts. 

The assistant proposed a `RulePack` defining available systems and constraints, and a generated `DesignPlan` mapping source concepts into them. It also proposed separate reference, preference, runtime, and evaluation components.

The benefit would be reuse without character-specific branches in framework code. The cost is defining the boundary between configuration and capabilities that require new implementation.

**Disposition:** Replaceability is an explicit owner requirement. The named pack architecture is an assistant proposal, not an approved implementation.

**Locator:** Owner, "the entire idea is to have a 'Replaceable' ruleset / dsl." Assistant, "a versioned `RulePack`, paired with a generated `DesignPlan`."

### Automatic research and interpretation from a name

The final substantive user request made the normal interface explicit: the user supplies the anime character's name, not a dossier or preselected path mapping.

The assistant proposed automatic subject resolution, sourced research, relationship extraction, competing layouts, detailed generation, validation, revision, and export. Project configuration would supply the game rules once.

This would reduce per-unit authoring work. It depends on source access, ambiguity handling, an output contract, and a workable model configuration.

**Disposition:** Name-only input is explicit owner intent. The particular multistage flow and search budgets remain proposed.

**Locator:** Owner, "the default flow is 'Anime Name in' -> 'Full TowerDefense Unit out'."

### Expansion into an executable reference game

The last assistant answer defined a complete result as an engine-loadable, playable unit with placeholder presentation. It proposed a runtime, reference enemies, waves, maps, benchmark towers, simulation, and numerical tuning.

Those components could support gameplay evaluation. They would also substantially increase v1 scope.

**Disposition:** Proposed, not accepted. The current product context does not make executable game code, assets, or demonstrated balance part of the default output promise.

**Locator:** Assistant, "an engine-loadable, playable unit with placeholder presentation."

### Terms whose meanings remain unsettled

| Term               | Meaning supported by this thread                                                                                                  | Unresolved distinction                                                                                                                     |
| ------------------ | --------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| Unit               | A character-derived tower defense design with progression under an active ruleset.                                                | The assistant later added executability. The owner did not approve that definition of "full."                                              |
| Mechanic           | Used for both a described gameplay behavior and a typed, executable operation.                                                    | These are different deliverables. Optional mechanic generation must not automatically imply runtime implementation.                        |
| Usable default     | The owner wants the public generator to perform the name-to-unit workflow.                                                        | Installation, provider setup, output completeness, and minimum quality were not specified.                                                 |
| Tuning             | Assistant proposals included prompts, preference examples, generation settings, and numerical balancing.                          | No model-weight training method was selected.                                                                                              |
| Taste              | Initially judgment about interesting and appropriate ideas; later, especially coherent interpretation of the reference character. | No agreed metric or evaluator exists.                                                                                                      |
| Exact BTD6 concept | Owner wording for the intended public default.                                                                                    | The assistant alternated between exactness and a neutral BTD6-inspired contract. Compatibility scope and target version remain unresolved. |

## 3. Decision record

“Accepted” below records clear owner intent. It does not imply implementation or experimental support. No formal owner-approved reconsideration conditions were recorded.

### UG-T05-D01. Name-only normal input

**Status:** Accepted.

The normal workflow should accept an anime character name and return a full tower defense unit.

**Authority and locator:** Owner, "'Anime Name in' -> 'Full TowerDefense Unit out'."

**Rationale:** The request is for automatic character adaptation, not a tool requiring the user to prepare the interpretation first.

**Dependencies:** The meaning of "full," source policy, and operational setup remain unresolved. No successful end-to-end run was shown.

### UG-T05-D02. Replaceable rules and a reusable framework

**Status:** Accepted.

The framework should support replaceable rulesets rather than contain a fixed BTD6 structure or Luffy-specific implementation.

**Authority and locator:** Owner, "use the same 'Generator Framework' for all possible/generalizable concepts."

**Rationale:** Reuse across concepts and different game rules.

**Dependencies:** "All possible/generalizable concepts" is an ambition, not a demonstrated capability. The supported domain boundary was not specified.

### UG-T05-D03. BTD6-derived public default

**Status:** Accepted at the level of stated intent; exact scope unresolved.

The owner requested a default based on the BTD6 concept, including three paths of five upgrades, at most two purchased paths, and only one path above T2. T1 and T2 should provide auxiliary upgrades, while T3 establishes a substantial specialization.

**Authority and locator:** Owner, "3 paths x 5 upgrades" and later "the exact BTD6 concept as the FOSS default."

**Alternatives:** The assistant proposed an independently authored, neutrally named, BTD6-inspired pack.

**Dependencies:** Exact game-version compatibility, exceptions, difficulty pricing, and Apex acquisition were not established. The tier responsibilities are owner design preferences, not verified universal BTD6 rules.

### UG-T05-D04. Reject the first Luffy layout

**Status:** Rejected.

The Jet Rhythm, Giant Impact, and Elastic Geometry design did not meet the owner's expectations for character identity and taste.

**Authority and locator:** Owner, "kinda tasteless and basic" and "doesnt capture the spirit."

**Rationale:** The complete unit did not feel like a consistent interpretation of Luffy.

**Evidence:** UG-T05-E01. This is direct qualitative owner feedback on one proposal, not a comparative experiment.

### UG-T05-D08. The generator itself should be fully FOSS

**Status:** Accepted.

**Authority and locator:** Owner, "the generator should be fully FOSS."

**Rationale:** The owner distinguished the generator's licensing from potential third-party-rights issues in referenced content or generated outputs.

**Dependencies:** No actual license was selected. The MIT license appeared only as an assistant example. License compatibility of a future distribution was not examined.

### UG-T05-D09. Non-commercial/private-use wording and rights assumptions

**Status:** Conflicting or unresolved, depending on the affected component.

The owner first requested a BTD6-inspired default for "NON-COMMERCIAL PRIVATE USE ONLY," then clarified that the generator should be fully FOSS and expressed concern about distribution of generated material.

The later clarification establishes the generator's intended FOSS status. It does not conclusively settle the license of every ruleset, example, reference collection, or output.

**Authority and locators:** Owner, the quoted restriction; later, "for private use, it is not a problem. Especially considering EU laws."

**Evidence limitation:** The private-use statement is an owner assumption, not a legal finding. The assistant challenged its generality using legal references. See UG-T05-E14.

### UG-T05-D10. Pack architecture and an explicit interpretation plan

**Status:** Proposed.

The assistant proposed `RulePack`, `ReferencePack`, `DesignProfile`, `DesignPlan`, a domain backend, and versioned configuration.

**Locator:** Assistant, "Separate the framework, ruleset, reference, and taste."

**Rationale:** Separate game permissions, source facts, design preferences, interpretation, and execution semantics.

**Dependencies:** No schema, API, package split, or implementation was approved or exercised.

### UG-T05-D11. Automatic research, layout selection, generation, and repair

**Status:** Proposed as implementation of UG-T05-D01.

**Locator:** Assistant, "The automatic generation flow" and the `generate_unit(subject_name, project)` pseudocode.

**Rationale:** Make source collection and conceptual organization automatic rather than mandatory user inputs.

**Dependencies:** Source access, model choice, reliable selection, bounded failure handling, and the output contract. No stage was shown running.

### UG-T05-D12. Executable output, reference runtime, and simulation

**Status:** Proposed, not accepted.

**Locator:** Assistant, "The default needs a small reference game" and "an engine-loadable, playable unit with placeholder presentation."

**Rationale:** Give statistics, pricing, and gameplay claims an executable reference environment.

**Scope conflict:** This goes beyond the current product context's default promise. That context does not retroactively turn the proposal into a rejected historical decision, but it prevents treating it as a v1 requirement.

### UG-T05-D13. Preference examples, comparative review, and diversity-preserving search

**Status:** Proposed.

**Locator:** Assistant, "Make taste replaceable too" and "How automatic selection gets your sense of taste."

**Rationale:** Select coherent layouts, preserve alternatives, and avoid relying on the generator's own creativity score.

**Dependencies:** Useful examples, held-out evaluation, and reliable measures. No preference dataset, trained evaluator, quality-diversity archive, or demonstrated improvement exists in this thread.

## 4. Technical artifacts, research claims, and results

### Artifact inventory

All project-specific artifacts below are inline proposals. They were written in chat but not shown to have been implemented or run.

| Artifact ID | Artifact and locator                                                                                               | State                                                                                             |
| ----------- | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------- |
| UG-T05-A01  | Upgrade-state equations, assistant section "Your upgrade contract."                                                | Written mathematical specification. Arithmetic independently recomputed during this handover.     |
| UG-T05-A02  | Rubber Brawler example, assistant section "Worked example: Luffy as one complete unit concept."                    | Written design and prices; explicitly unbalanced and subsequently rejected as a character design. |
| UG-T05-A04  | `RulePack`, `DesignPlan`, public/private YAML examples, and relationship vocabulary.                               | Illustrative schemas, not a library API or implemented DSL.                                       |
| UG-T05-A05  | CLI, project manifest, generation pseudocode, diagnostics, and export-bundle names.                                | Written implementation outline; no execution evidence.                                            |
| UG-T05-A06  | Proposed validation, simulation, comparative review, and mutation tests.                                           | Test plans only. No Unit-Generator results or datasets supplied.                                  |

### UG-T05-A01. Upgrade states

The assistant formalized normal progression as:

$$
(a,b,c)\in\{0,1,2,3,4,5\}^3
$$

with at most two nonzero coordinates and at most one coordinate greater than two. Purchases advance one path by one tier.

Under these assumptions there are **64 legal purchased-upgrade configurations**, including the base state. The six highest-total configurations are `5-2-0`, `5-0-2`, `2-5-0`, `0-5-2`, `2-0-5`, and `0-2-5`.

**UG-T05-E03:** An enumeration during this handover reproduced the count and endpoints. This verifies the arithmetic under the stated constraints, not BTD6 compatibility, runtime behavior, balance, or a Unit-Generator implementation.

The immediate two-path lock and the T3 specialization lock are distinct. Under the proposed contract, `1-1-0` already excludes investment in the third path.

### UG-T05-A02. Rejected design and numerical examples

The three paths were:

* Jet Rhythm, rapid normal attacks leading to barrages and recovery.
* Giant Impact, charged heavy attacks and stored power.
* Elastic Geometry, damaging return trajectories, selected endpoints, and delayed returns.

Important illustrative parameters included five successful normal cycles to earn a six-punch barrage, no barrage self-charging, one stored barrage or charged attack, and one active elastic fist. The return-attack schema used a **0.8-second endpoint hold**, no return homing, separate outbound/return hit budgets, and explicit modifier inheritance.

The proposed placement price was **600**. Tier prices were:

| Path             |  T1 |  T2 |    T3 |    T4 |     T5 |
| ---------------- | --: | --: | ----: | ----: | -----: |
| Jet Rhythm       | 150 | 300 | 1,800 | 6,000 | 32,000 |
| Giant Impact     | 200 | 350 | 2,200 | 7,500 | 38,000 |
| Elastic Geometry | 175 | 325 | 2,000 | 6,500 | 34,000 |

The assistant proposed difficulty multipliers of **0.85, 1.00, 1.10, and 1.25**, with each purchase rounded to a multiple of five. These were expressly example settings, not verified BTD6 values. Illustrative totals were `3-0-0 = 2,850` and `3-2-0 = 3,400` at base prices.

**UG-T05-E01:** The owner rejected the design's coherence. No combat, economy, or crosspath measurements accompany it.

The full fifteen upgrade descriptions and repeated prompts are not reproduced here. UG-T05-A02's locator identifies them; their prices and consequential behavioral constraints are retained.

### UG-T05-A04 and A05. Architecture, schemas, and flow

The proposed conceptual relationship vocabulary included `belongs_to_family`, `develops_into`, `can_coexist_with`, `replaces`, `requires`, `expresses_identity`, and `limited_by`.

The proposed distinction was:

* `RulePack`: permitted systems and game constraints.
* `ReferencePack`: source-supported facts and labeled interpretations.
* `DesignProfile`: preferences and examples.
* `DesignPlan`: selected mapping of the reference into game systems.
* Domain backend: executable semantics, where execution is required.

The public YAML example declared three five-tier paths and no shared-form progression. 

The final manifest proposed **six layout candidates, two detailed candidates, and three repair attempts**. An earlier answer suggested **twelve T3 concepts**. These are separate illustrative budgets, not measurements or a settled search policy.

Proposed exports were `unit.json`, `design.md`, `sources.json`, `evaluation.json`, `manifest.json`, and a `presentation/` directory. Proposed repository paths and filenames are not evidence that those files exist.

The pseudocode implemented an outline of research, layout generation, selection, detailed generation, compile/tune/test, bounded repair, and export. It was not runnable as provided because its functions were undeclared implementation contracts.

### UG-T05-A06. Proposed checks and failure cases

The thread proposed checking legal purchases, unsupported operations, missing references, dead crosspath modifiers, contradictory inheritance, resource resets, recursive effects, and dependence on purchase order.

Design-level failure cases included:

* Mechanically valid branches that do not capture the character.
* New-looking effects that never affect play.
* A legal crosspath with no useful contribution.
* A form that silently grants unpurchased specialization benefits.
* An Apex that merely stacks incompatible effects.
* An evaluator unable to use the unit's manual or automatic control features.
* A strong final build with an unusable intermediate purchase sequence.
* A generated candidate weakening its own rules or evaluation criteria.
* A source page being treated as instructions rather than reference data.

All are proposed diagnostics or anticipated failures. They are not observed Unit-Generator runtime failures.

### Research claims and evidence

**UG-T05-E04. Structured search can generate RTS units within a defined ability vocabulary.**
UG-T05-S01 reports ten microRTS units generated using search and MCTS-based evaluation. Its evaluation is tied to supplied triggers/effects and agents; it does not establish human enjoyment or transferable balance for anime tower defense units. ([arXiv][1])

**UG-T05-E05. Rule recombination can produce both apparently original games and rediscoveries.**
UG-T05-S02 reports 1,389 evolved candidates and 19 viable games. Its discussion contrasts Yavalath with Ndengrod's rediscovery of an existing design. This supports keeping independent generation separate from historical originality. It does not establish originality of any Unit-Generator output. ([cambolbro.com][2])

**UG-T05-E06. LLM-assisted search over executable game rules is a relevant precedent, with limited novelty evidence.**
UG-T05-S03 combines language-model mutations with game evaluation and diversity-oriented search. Its increase in novel occupied cells was not statistically significant across three runs, and its human assessment was preliminary. The paper also documents irrelevant generated components. ([arXiv][3])

**UG-T05-E07. Quality-diversity offers a way to retain different useful candidates.**
UG-T05-S04 presents the approach for procedural content generation. Whether its behavior descriptors correspond to meaningful unit differences remains a separate design problem. No such archive was built here. ([arXiv][4])

**UG-T05-E08. Evaluation and preference modeling are relevant to design selection.**
UG-T05-S05 examines more than 450 human-created games and supports a model incorporating evaluation rather than association alone. UG-T05-S06 describes experience-driven generation. Neither establishes an autonomous evaluator with this owner's taste. ([arXiv][5])

**UG-T05-E09. Novel proposal generation and reliable self-selection are different capabilities.**
UG-T05-S07 reports higher expert-rated novelty for generated research proposals in its experiment, alongside self-evaluation and diversity limitations. These were research proposals, not completed successful projects or game units. ([ICLR Proceedings][6])

**UG-T05-E10. Better-rated individual outputs can coexist with reduced collection diversity.**
UG-T05-S08 reports this pattern in short-fiction assistance. UG-T05-S09 finds homogeneity across tested LLM creativity responses. Neither directly tests unit rosters or proves that the effect holds for every model and generation strategy. ([UCL Discovery][7])

**UG-T05-E11. AI exposure is associated with novel, improved human decisions in a game.**
UG-T05-S10 analyzes approximately 5.8 million professional Go decisions from 1950 to 2021. Its historical observational design does not establish the causal effect of this proposed generator, and choosing moves is different from designing units. ([arXiv][8])

**UG-T05-E12. AI-assisted program search can produce externally checkable new results.**
UG-T05-S11 reports mathematical constructions and bin-packing heuristics found through program search. Its evaluators check objectives much more directly than "character coherence" or "designer usefulness." ([Nature][9])

**UG-T05-E13. Explicit character organization may improve unit-generation usefulness.**
This is the thread's central project hypothesis. Its support is the owner's diagnosis and relevant external precedents, not a Unit-Generator experiment.

**UG-T05-E14. Licensing and third-party rights are separate questions.**
The Open Source Definition does not permit a non-commercial restriction on the same FOSS component. The cited EU directives distinguish protected expression from underlying program ideas and provide conditional exceptions, not blanket clearance for private adaptations. Their application to an actual release remains unresolved. ([Open Source Initiative][10])

**UG-T05-E15. No operational Unit-Generator result is established.**
The reviewed record contains design prose, illustrative schemas, equations, pseudocode, research summaries, and owner feedback. It contains no inspectable generated bundle, runtime execution, playtest, baseline comparison, or balance result. The handover's arithmetic check is not a substitute.

## 5. Recommended approach and v1 boundary

### Recommendation actually reached in the thread

There was **no settled implementation architecture**.

The last assistant recommendation was an automatic research-and-generation pipeline with explicit layout selection, structured definitions, a compiler, reference runtime, simulation, repair, and export. The owner did not approve that proposal before requesting this handover.


### Current recommendation

The following is new advice awaiting approval.

**Product principle:** Generate a useful, coherent tower defense unit under the chosen rules before attempting to generate or validate an entire playable implementation.

The smallest useful prototype should complete the public name-to-unit workflow without requiring Towerright, a private character dossier, or a private preference pack.

| Proposed v1 component                                                                         | Why it is needed                                                                                                                                     |
| --------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| An editable public rules configuration and an explicit unit-output contract.                  | Replaceability and completeness cannot be tested without knowing which rules apply and which fields the result must contain.                         |
| Subject resolution and a small sourced reference record.                                      | The promised name-only interface requires the system to obtain its own character information and distinguish evidence from invention.                |
| A character-organization step followed by unit generation.                                    | This directly addresses UG-T05-D04 instead of merely producing more polished versions of the rejected approach.                                      |
| Deterministic checks for the configured structure and internal consistency.                   | These can catch path-count, tier, reference, and contradictory-output errors without requiring a game engine.                                        |
| A readable unit result, configuration/source provenance, and explicit unresolved assumptions. | The output must be useful to a designer and inspectable when it fails.                                                                               |
| A working documented model route and a small comparison exercise.                             | The default must actually run, and its claimed advantage requires initial evidence. This does not settle offline operation or require a local model. |

The default unit result should include the base concept and the progression required by the selected default pack. The owner still needs to approve exactly which statistics, costs, and interaction details constitute completeness. When numerical values are included, their assumed game context and provisional status should be visible.

Detailed or newly proposed mechanics should remain an optional output request. Basic descriptions needed to understand the unit do not imply an executable mechanics language.

### Build or test now

Implement one complete public flow and demonstrate that changing a public configuration changes the generated structure without character-specific framework code. Test source fidelity, rule adherence, coherent organization, and whether the result reduces designer revision work.

Use the rejected Luffy example as a development case, not as the sole demonstration of generalization. A generator that memorizes the preferred answer has not demonstrated the proposed capability.

### Defer or exclude from v1

Do not make a combat runtime, simulator, production art, animation generation, proven balance, model-weight training, learned preference model, quality-diversity search, or a multi-agent organization mandatory.

These proposals remain useful research or engineering options. Their necessity depends on the approved output contract and observed failures, rather than their presence in the literature.

Towerright integration, accounts, billing, and unrelated Mardwerk products are outside this prototype boundary.

### Public and Towerright-specific material

The public default must contain enough rules, instructions, and documented configuration to produce its promised output. Its quality must not depend on access to private examples or hidden Towerright prompts.

Under the current product context, Mardwerk's specially tuned offering belongs in Towerright. 

The public framework's ability to accept extensions should not be confused with permission to publish a particular private extension. No license is selected or compatibility assumed in this handover.

## 6. Open questions, conflicts, and risks

### UG-T05-Q01. What constitutes a "full" unit?

The owner requested a full unit; the assistant later required an executable one. The current product context does not require execution.

**Impact:** Blocks a stable v1 acceptance contract.

**Resolution:** Approve a concrete example output and distinguish required unit fields, optional mechanic detail, implementation assets, and unproven balance.

**Related:** UG-T05-D01, UG-T05-D12, UG-T05-E15.

### UG-T05-Q02. How exact is the public BTD6-derived default?

The owner said "exact BTD6 concept"; the assistant proposed a neutral inspired implementation. No target version, exception set, or complete Apex policy was established.

**Impact:** Affects default scope, testing, maintenance, and public documentation.

**Resolution:** Define the rules contract explicitly and identify which claims are compatibility claims versus original design conventions.

**Related:** UG-T05-D03, UG-T05-S16.

### UG-T05-Q04. How will automatic character research work?

Continuity, spoiler cutoff, source hierarchy, ambiguity handling, retrieval access, and source-failure behavior remain proposals rather than policy.

**Impact:** Can block the name-only workflow or create confident but incorrect adaptations.

**Resolution:** Select an initial source policy and test subject resolution and factual claims against available source material.

**Related:** UG-T05-D01, UG-T05-D11, UG-T05-S17, UG-T05-S18.

### UG-T05-Q05. What model and operating environment make the default usable?

No provider, local-model requirement, hardware target, cost envelope, or hosting mode was agreed.

**Impact:** Blocks a credible installation and running-cost statement, but does not justify assuming offline operation.

**Resolution:** Implement and document one supported route, including credentials, dependencies, expected costs, and failure behavior. Keep the configuration replaceable.

**Related:** UG-T05-D08, UG-T05-D11.

### UG-T05-Q06. Can the system select coherent layouts without copying the Luffy example?

The owner can identify a preferred organization. The proposed model-based selector has not demonstrated that capability.

**Impact:** Central research uncertainty and potential reason the product may not outperform a direct prompt.

**Resolution:** Compare held-out character results, including sources whose abilities do not naturally fit three parallel disciplines.

**Related:** UG-T05-D04, UG-T05-D05, UG-T05-D13, UG-T05-E09, UG-T05-E13.

### UG-T05-Q07. How much mechanical meaning belongs in the DSL?

A descriptive configuration, a typed unit specification, and an executable behavior language have different complexity and guarantees.

**Impact:** Can substantially change v1 feasibility and the meaning of validation.

**Resolution:** Decide the v1 representation first. Reject claims of executability for unsupported behaviors, and distinguish structural validation from gameplay validation.

**Related:** UG-T05-D10, UG-T05-D12, UG-T05-A04.

### UG-T05-Q09. Which restrictions apply to software, packs, references, and outputs?

The thread distinguishes FOSS software from third-party content concerns, but does not settle component licenses or establish private-use clearance.

**Impact:** Affects distribution and reuse. It need not block a prototype using material for which appropriate rights are available.

**Resolution:** Separate the artifacts and review the intended release under the applicable jurisdiction. Do not use a "private/non-commercial" label as a substitute for that analysis.

**Related:** UG-T05-D08, UG-T05-D09, UG-T05-E14, UG-T05-S19 through UG-T05-S24.

### UG-T05-Q10. What would count as "better"?

Novel wording, familiar ability names, legal upgrade trees, useful unit designs, and balanced gameplay are different outcomes.

**Impact:** Determines whether the central product advantage can be evaluated.

**Resolution:** Select measures relevant to the approved output contract and compare against a strong direct-LLM-plus-wiki baseline. Balance requires an explicitly defined game and separate evidence.

**Related:** UG-T05-E01, UG-T05-E10, UG-T05-E13, UG-T05-E15.

## 7. Sources and research leads

### A. Sources already present in this thread

Verification labels below describe work performed for this handover. “Bibliographic record verified” does not mean every finding or supplementary artifact was checked. Unless stated otherwise, associated repositories or datasets were identified but not executed or audited.

#### UG-T05-S01. Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search

Kynan Sorochan and Matthew Guzdial, 2022. Experimental AI in Games Workshop, as identified in the arXiv record. arXiv:2212.03387v1.

URL: [https://arxiv.org/abs/2212.03387](https://arxiv.org/abs/2212.03387)
Inspected text: [https://arxiv.org/html/2212.03387v1](https://arxiv.org/html/2212.03387v1)

**Bibliographic record verified. Content inspected: relevant sections**, including search representation, evaluation, results, and limitations.

The opening research answer's first paper. Directly relevant to UG-T05-E04 and the proposed structured-unit evaluation in UG-T05-D12. The microRTS testbed and ten generated units are described; this is not a dataset of anime adaptations or human-rated tower defense designs. ([arXiv][11])

#### UG-T05-S02. Evolutionary Game Design

Cameron Browne and Frédéric Maire, 2010. Cited in the thread as *IEEE Transactions on Computational Intelligence and AI in Games*, 2(1), 1-16.

Author manuscript: [https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf](https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf)

**Bibliographic record unverified in full:** title and authors confirmed from the manuscript; final issue metadata was not independently resolved. **Content inspected: relevant sections**, including the candidate counts and the Yavalath/Ndengrod discussion.

The opening research answer used Ludi to connect evaluation, invention, and rediscovery. Supports UG-T05-E05. Its preference model and restricted game family do not establish general design taste. No Ludi implementation was run. ([cambolbro.com][2])

#### UG-T05-S03. GAVEL: Generating Games Via Evolution and Language Models

Graham Todd, Alexander Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers, and Julian Togelius, 2024. NeurIPS 2024. arXiv:2407.09388v2.

URL: [https://arxiv.org/abs/2407.09388](https://arxiv.org/abs/2407.09388)
Text: [https://arxiv.org/html/2407.09388v2](https://arxiv.org/html/2407.09388v2)
Code: [https://github.com/gdrtodd/gavel](https://github.com/gdrtodd/gavel)

**Bibliographic record verified. Content inspected: relevant sections**, including evaluation and limitations.

Used repeatedly to motivate executable rule mutation and search, especially UG-T05-E06 and UG-T05-D10. Ludii games, generated examples, and a model checkpoint are associated artifacts; none was run here.

An original tool trace attempted `https://arxiv.org/html/2407.09388v3`. That version was not verified; v2 is the inspected version. ([arXiv][12])

#### UG-T05-S04. Procedural Content Generation through Quality Diversity

Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius, and Georgios N. Yannakakis, 2019. *IEEE Conference on Games*. DOI: 10.1109/CIG.2019.8848053. arXiv:1907.04053v1.

URL: [https://arxiv.org/abs/1907.04053](https://arxiv.org/abs/1907.04053)
Original alternative: [https://antoniosliapis.com/articles/pcgqd.php](https://antoniosliapis.com/articles/pcgqd.php)

**Bibliographic record verified. Content inspected: abstract only** during this verification.

The opening research answer and subsequent architecture proposals used it to motivate retaining different useful designs, UG-T05-E07 and UG-T05-D13. It is a framework/survey contribution, not evidence that the proposed behavior measures capture character coherence. ([arXiv][4])

#### UG-T05-S05. Generation and Evaluation in the Human Invention Process through the Lens of Game Design

Katherine M. Collins, Graham Todd, Cedegao E. Zhang, Adrian Weller, Julian Togelius, Junyi Chu, Lionel Wong, Thomas L. Griffiths, and Joshua B. Tenenbaum, 2025. Non-archival CogSci conference paper and arXiv preprint, arXiv:2508.10914.

Original version inspected: [https://arxiv.org/html/2508.10914v1](https://arxiv.org/html/2508.10914v1)
Record: [https://arxiv.org/abs/2508.10914](https://arxiv.org/abs/2508.10914)

**Bibliographic record verified. Content inspected: relevant sections** of v1.

Used to connect invention with evaluation, UG-T05-E08. The study concerns human-created games and preliminary cognitive modeling, not an expert AI designer. A later revision exists, but its changes are not silently substituted for the version used in the conversation. ([arXiv][13])

#### UG-T05-S06. Experience-Driven Procedural Content Generation

Georgios N. Yannakakis and Julian Togelius, 2011. *IEEE Transactions on Affective Computing*, 2(3), 147-161. DOI: 10.1109/T-AFFC.2011.6.

Author manuscript: [https://yannakakis.net/wp-content/uploads/2019/02/EDPCG.pdf](https://yannakakis.net/wp-content/uploads/2019/02/EDPCG.pdf)

**Bibliographic record verified at title, author, year, venue, and DOI level; final pagination was not separately rechecked. Content inspected: relevant introductory sections.**

The opening research answer used it to make evaluation conditional on player experience and preferences, UG-T05-E08. It is a framework rather than evidence of universal aesthetic judgment.

An earlier trace also attempted `https://julian.togelius.com/Yannakakis2011Experience.pdf`; that alternative was not verified. 

#### UG-T05-S07. Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers

Chenglei Si, Diyi Yang, and Tatsunori Hashimoto. ICLR 2025; preprint first posted in 2024. arXiv:2409.04109.

Record: [https://arxiv.org/abs/2409.04109](https://arxiv.org/abs/2409.04109)
Published paper: [https://proceedings.iclr.cc/paper_files/paper/2025/file/ea94957d81b1c1caf87ef5319fa6b467-Paper-Conference.pdf](https://proceedings.iclr.cc/paper_files/paper/2025/file/ea94957d81b1c1caf87ef5319fa6b467-Paper-Conference.pdf)
Code and materials: [https://github.com/NoviScl/AI-Researcher](https://github.com/NoviScl/AI-Researcher)

**Bibliographic record verified. Content inspected: relevant sections**, not the entire long appendix.

Used in the opening research answer and later warnings about self-selection, UG-T05-E09 and UG-T05-D13. Expert judgments of proposed NLP research do not directly measure unit usefulness, implementation correctness, or game balance. ([ICLR Proceedings][6])

#### UG-T05-S08. Generative AI enhances individual creativity but reduces the collective diversity of novel content

Anil R. Doshi and Oliver P. Hauser, 2024. *Science Advances*, 10(28), eadn5290. DOI: 10.1126/sciadv.adn5290.

Original URL: [https://www.science.org/doi/10.1126/sciadv.adn5290](https://www.science.org/doi/10.1126/sciadv.adn5290)
Original alternative: [https://pmc.ncbi.nlm.nih.gov/articles/PMC11244532/](https://pmc.ncbi.nlm.nih.gov/articles/PMC11244532/)
Verified institutional copy: [https://discovery.ucl.ac.uk/10195027/1/Generative%20AI%20enhances%20individual%20creativity%20but%20reduces%20the%20collective%20diversity%20of%20novel%20content.pdf](https://discovery.ucl.ac.uk/10195027/1/Generative%20AI%20enhances%20individual%20creativity%20but%20reduces%20the%20collective%20diversity%20of%20novel%20content.pdf)
Data/code identifier: [https://doi.org/10.5061/dryad.qfttdz0pm](https://doi.org/10.5061/dryad.qfttdz0pm)

**Bibliographic record verified. Content inspected: relevant sections**, including results, story similarity, discussion, and data availability.

The original answer used this to warn about collection-level homogeneity, UG-T05-E10. Its short-fiction task and participant population limit transfer to expert game design. ([UCL Discovery][7])

#### UG-T05-S09. Large language models are homogeneously creative

Emily Wenger and Yoed N. Kenett, 2026. *PNAS Nexus*, 5(3), pgag042. DOI: 10.1093/pnasnexus/pgag042.

URL: [https://academic.oup.com/pnasnexus/article/5/3/pgag042/8529001](https://academic.oup.com/pnasnexus/article/5/3/pgag042/8529001)
Associated repository: [https://github.com/arguslab-duke/llms_are_homogeneous](https://github.com/arguslab-duke/llms_are_homogeneous)

**Bibliographic record verified. Content inspected: relevant sections**, including the abstract, discussion, and materials statement.

Already present as a companion source in the opening research answer, not newly introduced here. Relevant to UG-T05-E10. Its tested models and standardized tasks do not justify a universal claim about all creative systems or future models. ([OUP Academic][14])

#### UG-T05-S10. Superhuman artificial intelligence can improve human decision-making by increasing novelty

Minkyu Shin, Jin Kim, Bas van Opheusden, and Thomas L. Griffiths, 2023. *PNAS*, 120(12), e2214840120. DOI: 10.1073/pnas.2214840120. arXiv:2303.07462.

URL: [https://www.pnas.org/doi/10.1073/pnas.2214840120](https://www.pnas.org/doi/10.1073/pnas.2214840120)
Accessible manuscript: [https://arxiv.org/pdf/2303.07462](https://arxiv.org/pdf/2303.07462)
Associated materials: [https://osf.io/xpf3q/](https://osf.io/xpf3q/)

**Bibliographic record verified. Content inspected: relevant sections.**

The opening answer used this for novel decisions inspired by AI, UG-T05-E11. It is an observational Go analysis, not unit generation or a randomized test of the proposed workflow. ([arXiv][15])

#### UG-T05-S11. Mathematical discoveries from program search with large language models

Bernardino Romera-Paredes, Mohammadamin Barekatain, Alexander Novikov, Matej Balog, M. Pawan Kumar, Emilien Dupont, Francisco J. R. Ruiz, Jordan S. Ellenberg, Pengming Wang, Omar Fawzi, Pushmeet Kohli, and Alhussein Fawzi. *Nature*, 625, 468-475, 2024; published online in 2023. DOI: 10.1038/s41586-023-06924-6.

URL: [https://www.nature.com/articles/s41586-023-06924-6](https://www.nature.com/articles/s41586-023-06924-6)
Code: [https://github.com/google-deepmind/funsearch](https://github.com/google-deepmind/funsearch)

**Bibliographic record verified. Content inspected: relevant sections**, including the method summary and code/data statements.

Used as evidence for externally verifiable novelty, UG-T05-E12. Its objective evaluators do not solve subjective unit-design selection. No associated code was run. ([Nature][9])

#### UG-T05-S12. Creativity and artificial intelligence

Margaret A. Boden, 1998. *Artificial Intelligence*, 103(1-2), 347-356. DOI: 10.1016/S0004-3702(98)00055-1.

URL: [https://www.sciencedirect.com/science/article/pii/S0004370298000551](https://www.sciencedirect.com/science/article/pii/S0004370298000551)

**Bibliographic record verified from publisher metadata. Content inspected: metadata only** in this handover; full text was inaccessible.

The original answer used its distinction between combinatorial, exploratory, and transformational creativity. Retained as a conceptual reference for UG-T05-E13, not an empirical finding about modern LLMs. The original interpretation was not fully re-audited here. ([ScienceDirect][16])

#### UG-T05-S13. Cooking Up Creativity

The original trace opened a paper whose v1 title is *Cooking Up Creativity: A Cognitively-Inspired Approach for Enhancing LLM Creativity through Structured Representations*. The current record uses *Cooking Up Creativity: Enhancing LLM Creativity through Structured Recombination*.

Moran Mizrahi, Chen Shani, Gabriel Stanovsky, Dan Jurafsky, and Dafna Shahaf, 2025. arXiv:2504.20643.

Original URL: [https://arxiv.org/html/2504.20643v1](https://arxiv.org/html/2504.20643v1)
Record: [https://arxiv.org/abs/2504.20643](https://arxiv.org/abs/2504.20643)

**Bibliographic record verified. Content inspected: abstract and representation-related sections.**

This was present in a visible tool trace, not developed into an explicit recommendation. Recipe recombination is a related representation lead, not evidence for Unit-Generator. Publication venue and reusable artifacts were not established here. ([arXiv][17])

#### UG-T05-S14. Abandoning Objectives: Evolution Through the Search for Novelty Alone

Joel Lehman and Kenneth O. Stanley, 2011. *Evolutionary Computation*, 19(2), 189-223. DOI: 10.1162/EVCO_a_00025.

Institutional record: [https://stars.library.ucf.edu/facultybib2010/1530/](https://stars.library.ucf.edu/facultybib2010/1530/)

**Bibliographic record verified. Content inspected: abstract only.**

The title was recoverable from original search queries containing "Abandoning Objectives" and "Novelty." It was not an adopted algorithm in the final discussion. Relevant to the novelty-search background of UG-T05-D13, but novelty without a task-quality objective does not itself establish designer usefulness. ([Stars][18])

#### UG-T05-S15. "Mechanic Maker"

**Unresolved research lead from an original search query.**

Original wording included `"Mechanic Maker" game paper`. No identifiable title, author list, year, venue, persistent identifier, or direct paper URL was recovered.

**Bibliographic record unverified. Content inspected: not inspected.**

It must not be treated as supporting evidence or silently replaced with a different mechanic-generation paper.

#### UG-T05-S16. Bloons TD 6

Ninja Kiwi, official developer-provided Steam product description. Commercial game reference, not a research paper.

URL: [https://store.steampowered.com/app/960090/Bloons_TD_6/](https://store.steampowered.com/app/960090/Bloons_TD_6/)

**Bibliographic record verified as a product reference. Content inspected: relevant product-description sections.**

Used to support the broad three-upgrade-path and Paragon inspiration, UG-T05-D03. The description does not verify every crosspath rule, cost multiplier, exception, or version-specific behavior proposed in this thread. ([Steam Store][19])

#### UG-T05-S17. Official Luffy character profile

ONE PIECE.com character page, cited by the earlier assistant as an official reference.

Original URL: [https://one-piece.com/character/luffy/index.html](https://one-piece.com/character/luffy/index.html)

**Bibliographic record unverified in this handover. Content inspected: not inspected; retrieval failed.**

Originally used for character anchors. No verified source snapshot or complete character dossier exists in this thread. 

#### UG-T05-S18. FEATURE: How Luffy vs Katakuri Puts An Interesting Spin On Anime Power-Ups

Daniel Dockery, September 18, 2020. Crunchyroll editorial feature.

Original URL: [https://www.crunchyroll.com/pl/news/features/2020/9/18/feature-how-luffy-vs-katakuri-puts-an-interesting-spin-on-anime-power-ups](https://www.crunchyroll.com/pl/news/features/2020/9/18/feature-how-luffy-vs-katakuri-puts-an-interesting-spin-on-anime-power-ups)

**Bibliographic record verified from publisher metadata. Content inspected: metadata only; article body was not recovered.**

The assistant used it for the future-sight interpretation of Observation. It is a secondary editorial reference, not the original episode or manga passage, and it does not establish the proposed targeting or activation mechanics. ([Crunchyroll][20])

#### UG-T05-S19. The Open Source Definition

Open Source Initiative, definition, version 1.9.

URL: [https://opensource.org/osd](https://opensource.org/osd)

**Bibliographic record verified. Content inspected: full text.**

Used in the licensing discussion, UG-T05-D08, UG-T05-D09, and UG-T05-E14. It defines open-source licensing criteria; it does not grant rights to third-party characters or game content. ([Open Source Initiative][10])

#### UG-T05-S20. The MIT License

Open Source Initiative license text; SPDX identifier MIT.

URL: [https://opensource.org/license/mit](https://opensource.org/license/mit)

**Bibliographic record verified. Content inspected: full text.**

The assistant offered it as an example. No selection occurred, and its permissions do not establish rights in material the project does not own. Relevant to UG-T05-D08. ([Open Source Initiative][21])

#### UG-T05-S21. What is Copyright?

U.S. Copyright Office explanatory reference.

URL: [https://www.copyright.gov/what-is-copyright/](https://www.copyright.gov/what-is-copyright/)

**Bibliographic record verified. Content inspected: relevant sections.**

Used by an earlier assistant to distinguish protected expression from ideas and systems. This is a U.S. reference and does not settle the owner's EU question. Related to UG-T05-E14. ([U.S. Copyright Office][22])

#### UG-T05-S22. U.S. Copyright Office Fair Use Index

U.S. Copyright Office legal-information resource.

URL: [https://www.copyright.gov/fair-use/](https://www.copyright.gov/fair-use/)

**Bibliographic record verified. Content inspected: relevant introductory sections.**

Used for the warning that non-commercial use does not automatically establish fair use. It is not an EU private-copying authority or clearance for a particular output. Related to UG-T05-Q09. ([U.S. Copyright Office][23])

#### UG-T05-S23. Directive 2009/24/EC on the legal protection of computer programs

European Parliament and Council, April 23, 2009. EU legislation, WIPO Lex record 8612.

URL: [https://www.wipo.int/wipolex/en/legislation/details/8612](https://www.wipo.int/wipolex/en/legislation/details/8612)

**Bibliographic record verified. Content inspected: relevant articles**, including Article 1(2) and program-use exceptions.

The final assistant answer used it to distinguish program expression from underlying ideas and principles. Relevant to UG-T05-E14. It does not determine whether a particular game adaptation copies protected expression or address every applicable right. ([WIPO][24])

#### UG-T05-S24. Directive 2001/29/EC on copyright and related rights in the information society

European Parliament and Council, May 22, 2001. EU legislation, WIPO Lex record 1453.

URL: [https://www.wipo.int/wipolex/en/legislation/details/1453](https://www.wipo.int/wipolex/en/legislation/details/1453)

**Bibliographic record verified. Content inspected: relevant articles**, including Article 5(2)(b) and Article 5(5).

Used to qualify the private-use assumption. It allows conditional national exceptions rather than creating blanket authorization for private adaptations. No national-law analysis or actual release review occurred. Related to UG-T05-D09 and UG-T05-E14. ([WIPO][25])

#### UG-T05-S25. JSON Schema reference

JSON Schema project, *Understanding JSON Schema*, technical documentation. No document-version pin was selected.

URL: [https://json-schema.org/understanding-json-schema/reference](https://json-schema.org/understanding-json-schema/reference)

**Bibliographic record verified. Content inspected: reference index and overview material.**

Proposed for structural validation in UG-T05-D10. It does not provide game semantics, a taste evaluator, or balance validation. ([JSON Schema][26])

#### UG-T05-S26. Common Expression Language overview

CEL project, official technical documentation.

URL: [https://cel.dev/overview/cel-overview](https://cel.dev/overview/cel-overview)

**Bibliographic record verified. Content inspected: relevant overview sections.**

Proposed as an expression-language option for typed, host-defined conditions. It was not selected or integrated. The host would still need to define the game-specific semantics. Related to UG-T05-D10. ([Common Expression Language][27])

#### UG-T05-S27. Unrecoverable citation and search material

Some original responses contain opaque citation references, redacted search results, and a malformed citation placeholder around the Ludi discussion.

**Bibliographic record unverified. Content inspected: not inspected.**

Identifiable papers and URLs have been registered above. The remaining material cannot support additional claims. The decorative image preview, repeated status messages, and duplicate links are not treated as research artifacts.

### B. Newly identified literature leads

These sources were found during handover preparation. They were not part of the original discussion.

#### UG-T05-S28. Automated Game Design via Conceptual Expansion

**Newly found, not part of the original conversation.**

Matthew Guzdial and Mark Riedl, 2018. *AIIDE*, 14(1), 31-37. DOI: 10.1609/aiide.v14i1.13022. arXiv:1809.02232.

Publisher: [https://ojs.aaai.org/index.php/AIIDE/article/view/13022](https://ojs.aaai.org/index.php/AIIDE/article/view/13022)
Text: [https://arxiv.org/html/1809.02232v1](https://arxiv.org/html/1809.02232v1)

**Bibliographic record verified. Content inspected: relevant approach, evaluation, and discussion sections.**

**UG-T05-E16:** The paper combines learned game representations through conceptual expansion. Its evaluation centers on reconstructing existing games, which limits claims about demonstrated creative quality.

It fills a representation/recombination gap around UG-T05-E13 and shows why a possible future contribution must be compared with earlier automated adaptation work. Code availability was not audited. ([AAAI Publications][28])

#### UG-T05-S29. Large Language Models Cannot Self-Correct Reasoning Yet

**Newly found, not part of the original conversation.**

Jie Huang, Xinyun Chen, Swaroop Mishra, Huaixiu Steven Zheng, Adams Wei Yu, Xinying Song, and Denny Zhou. ICLR 2024; preprint first posted in 2023. arXiv:2310.01798v2.

Record: [https://arxiv.org/abs/2310.01798](https://arxiv.org/abs/2310.01798)
Text: [https://arxiv.org/html/2310.01798v2](https://arxiv.org/html/2310.01798v2)

**Bibliographic record verified. Content inspected: relevant sections**, including intrinsic self-correction and baseline comparisons.

**UG-T05-E17:** In its tested reasoning settings, self-correction without external feedback could fail or degrade results. The paper also motivates stronger direct-prompt baselines.

This challenges unvalidated review-and-repair loops in UG-T05-D11 and UG-T05-D13. It does not prove that all creative critique, external validation, or newer models fail. Associated code was not audited. ([arXiv][29])

### Remaining literature gaps

A broader review should address character-faithful adaptation, evaluation of source relationships rather than mere fact inclusion, preference learning for game designers, expressive but controllable generation languages, and the relation between roster diversity and individual-unit usefulness.

None of the inspected sources establishes the exact anime-name-to-tower-defense-unit workflow proposed here. This extraction is not an exhaustive literature review, and the references inside every cited paper have not been exhaustively registered.

## 8. Preparation for later research

### Existing research plans versus new suggestions

Historically, the owner requested research about taste, novelty, and unit/game design. The assistant proposed simulations, human review, diversity-preserving search, and future comparative evaluation.

A formal Unit-Generator research protocol, sample size, success threshold, dataset, or publication claim is **not established in this thread**. The functioning-v1-first and later-paper goals come from the current product context.

### Research question suggested by the thread

**Proposed question:** Does explicit organization of source-character relationships under a replaceable ruleset produce more useful and coherent tower defense units than a strong direct LLM prompt supplied with the same rules and reference information?

A possible contribution is a representation and workflow that separates source evidence, character organization, game constraints, and unit output. Its novelty must be assessed against prior work such as UG-T05-S03 and UG-T05-S28. It is not established merely by naming these components.

### What has evidence and what remains hypothetical

There is direct evidence that the owner disliked one structurally developed proposal and preferred a different organization, UG-T05-E01 and UG-T05-E02.

External papers establish relevant possibilities and limitations in other settings, UG-T05-E04 through UG-T05-E12. They do not establish that Unit-Generator is more coherent, more novel, faster to use, or better balanced.

The central benefit remains UG-T05-E13. Automatic taste selection, generalization across characters, and reliable adaptation across rulesets are also hypotheses.

### Smallest useful next evaluation

The next evaluation should compare actual v1 outputs with a **direct-LLM-plus-wiki baseline**, not with an intentionally vague prompt.

For the baseline, provide a clearly identified character reference, the active rules, the same requested output fields, and competent instructions to preserve character identity. Record the exact prompt. Do not withhold constraints or important source information merely to create a weaker comparison.

Two questions should remain distinct:

**Matched-information comparison.** Supply both methods with the same reference snapshot and rules. This tests whether the proposed organization step helps beyond information access.

**End-to-end workflow comparison.** Begin from the name and measure the complete research-and-generation experience. This tests automation value, but source retrieval differences then become part of the treatment.

The smallest initial exercise need not be publication-ready. No sample size or pass threshold is approved here.

### Measures relevant to this thread

| Proposed measure        | What it can establish                                                                          | What it cannot establish alone                     |
| ----------------------- | ---------------------------------------------------------------------------------------------- | -------------------------------------------------- |
| Constraint satisfaction | Whether the unit follows the configured structure and explicit restrictions.                   | Taste, usefulness, or balance.                     |
| Source fidelity         | Whether character claims are supported and invented adaptations are labeled.                   | Whether the chosen interpretation is the best one. |
| Character coherence     | Whether the base, paths, and optional systems form a recognizable consistent design.           | Historical originality or legal permission.        |
| Designer usefulness     | Preference, required corrections, and work needed to make the output usable.                   | Universal player enjoyment.                        |
| Controllability         | Whether changing a rule or preference produces the intended change without unrelated breakage. | Generality beyond the tested changes.              |
| Diversity               | Whether repeated outputs differ in meaningful organization or behavior.                        | Quality of those differences.                      |

These measures are proposals. Definitions, rating instructions, and success thresholds await approval.

Game balance should enter the evaluation only when the game rules, opponents, economy, evaluator, and relevant play conditions are defined. Plausible numbers and convincing prose are not balance evidence.

### Fair comparison and retained records

Record the source snapshot or an appropriately retained representation, continuity cutoff, ruleset, prompts, model/provider/version, generation parameters, token usage, candidate counts, revision attempts, tools, elapsed time, and monetary or compute cost.

A workflow that spends substantially more resources should not be described as an architecture-only improvement without reporting that difference. Model changes and extra reference information must also remain visible.

Retain failed outputs as well as selected ones. Save code revisions, resolved configuration versions, intermediate layouts, final units, validation results, reviewer criteria, and reasons for rejection.


### Public versus private research results

Results obtainable using the public default should be reported separately from results requiring Towerright's private profiles, rules, reference assets, or tuning.

A private improvement can be valuable without being publicly reproducible. Its dependency should be stated, and it should not be used to imply that the public default has the same measured capability.

Retaining evidence does not imply publishing every source snapshot or private asset. Publication permissions, third-party rights, and Towerright disclosure decisions remain separate questions.

## 9. Consolidation handoff


Owner approval is still needed for the output boundary in **UG-T05-Q01**, default compatibility scope in **UG-T05-Q02**, public/private artifact boundary in **UG-T05-Q03**, and licensing/reuse questions in **UG-T05-Q09**. 

The next concrete action is to approve one public-default unit-output example, then implement the smallest name-to-unit workflow in section 5 and compare its output with the matched-information baseline in section 8. That tests **UG-T05-E13** without first requiring a game engine, Towerright integration, or a publication-ready study.

[1]: https://arxiv.org/html/2212.03387v1 "Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search"
[2]: https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf "https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf"
[3]: https://arxiv.org/html/2407.09388v2 "https://arxiv.org/html/2407.09388v2"
[4]: https://arxiv.org/abs/1907.04053 "[1907.04053] Procedural Content Generation through Quality Diversity"
[5]: https://arxiv.org/html/2508.10914v1 "https://arxiv.org/html/2508.10914v1"
[6]: https://proceedings.iclr.cc/paper_files/paper/2025/file/ea94957d81b1c1caf87ef5319fa6b467-Paper-Conference.pdf "https://proceedings.iclr.cc/paper_files/paper/2025/file/ea94957d81b1c1caf87ef5319fa6b467-Paper-Conference.pdf"
[7]: https://discovery.ucl.ac.uk/10195027/1/Generative%20AI%20enhances%20individual%20creativity%20but%20reduces%20the%20collective%20diversity%20of%20novel%20content.pdf "https://discovery.ucl.ac.uk/10195027/1/Generative%20AI%20enhances%20individual%20creativity%20but%20reduces%20the%20collective%20diversity%20of%20novel%20content.pdf"
[8]: https://arxiv.org/pdf/2303.07462 "https://arxiv.org/pdf/2303.07462"
[9]: https://www.nature.com/articles/s41586-023-06924-6 "Mathematical discoveries from program search with large language models | Nature"
[10]: https://opensource.org/osd "https://opensource.org/osd"
[11]: https://arxiv.org/abs/2212.03387 "https://arxiv.org/abs/2212.03387"
[12]: https://arxiv.org/abs/2407.09388 "https://arxiv.org/abs/2407.09388"
[13]: https://arxiv.org/abs/2508.10914 "[2508.10914] Generation and Evaluation in the Human Invention Process through the Lens of Game Design"
[14]: https://academic.oup.com/pnasnexus/article/5/3/pgag042/8529001 "oup.silverchair-cdn.com"
[15]: https://arxiv.org/abs/2303.07462 "[2303.07462] Superhuman Artificial Intelligence Can Improve Human Decision Making by Increasing Novelty"
[16]: https://www.sciencedirect.com/science/article/pii/S0004370298000551 "https://www.sciencedirect.com/science/article/pii/S0004370298000551"
[17]: https://arxiv.org/abs/2504.20643 "https://arxiv.org/abs/2504.20643"
[18]: https://stars.library.ucf.edu/facultybib2010/1530/ "https://stars.library.ucf.edu/facultybib2010/1530/"
[19]: https://store.steampowered.com/app/960090/Bloons_TD_6/ "https://store.steampowered.com/app/960090/Bloons_TD_6/"
[20]: https://www.crunchyroll.com/th/news/features/2020/9/18/feature-how-luffy-vs-katakuri-puts-an-interesting-spin-on-anime-power-ups?srsltid=AU7gw4W0Ds3S3fE2ountP7Qn_4l2JY-t5Ij6dw-IqZPfu2AUJe71CdRS "https://www.crunchyroll.com/th/news/features/2020/9/18/feature-how-luffy-vs-katakuri-puts-an-interesting-spin-on-anime-power-ups?srsltid=AU7gw4W0Ds3S3fE2ountP7Qn_4l2JY-t5Ij6dw-IqZPfu2AUJe71CdRS"
[21]: https://opensource.org/license/mit "https://opensource.org/license/mit"
[22]: https://www.copyright.gov/what-is-copyright/ "https://www.copyright.gov/what-is-copyright/"
[23]: https://www.copyright.gov/fair-use/ "https://www.copyright.gov/fair-use/"
[24]: https://www.wipo.int/wipolex/en/legislation/details/8612 "https://www.wipo.int/wipolex/en/legislation/details/8612"
[25]: https://www.wipo.int/wipolex/en/legislation/details/1453 "https://www.wipo.int/wipolex/en/legislation/details/1453"
[26]: https://json-schema.org/understanding-json-schema/reference "https://json-schema.org/understanding-json-schema/reference"
[27]: https://cel.dev/overview/cel-overview "https://cel.dev/overview/cel-overview"
[28]: https://ojs.aaai.org/index.php/AIIDE/article/view/13022 "https://ojs.aaai.org/index.php/AIIDE/article/view/13022"
[29]: https://arxiv.org/abs/2310.01798 "https://arxiv.org/abs/2310.01798"
