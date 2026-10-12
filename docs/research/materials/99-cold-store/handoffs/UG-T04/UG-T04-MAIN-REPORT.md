# UG-T04 handover: character adaptation, design taste and preference feedback

Internal consolidation record. Prepared September 24, 2026.
Approval and publication status: this handover records the discussion; it does not approve an architecture, dependency, study, license or disclosure.

## 1. Thread overview and coverage

**Thread ID:** UG-T04.
**Descriptive title:** From constrained character adaptation to learning preferred upgrade-tree organization.

**Original question.** The user asked how to transform a character description or wiki entry into a complete tower defense unit under a supplied Game Definition, including whether a specialized model should be pretrained, post-trained or replaced by another approach. Locator: user, "its rather a transformation then a generation" and "BEST CASE i would like to train an AI model".

**Contribution.** The thread assembled adjacent research, proposed a source-grounded generation and checking process, then identified a more specific unresolved problem: a mechanically plausible unit can still have an unsatisfying interpretation and progression. The user supplied a concrete negative example, a preferred alternative structure and a proposal for comparative human review.

**Material reviewed.** The visible original user requests, the two substantive assistant answers, visible research-tool requests, and both complete uploaded files: `GAME-DESIGN-RESEARCH.md` (UG-T04-S01) and `CONTEXT.md` (UG-T04-S02). The source register also retains identifiable sources appearing only in the earlier research requests, without treating a request as proof that the source was read. Bibliographic checks for this handover used public primary records and, where stated, underlying text.

**Coverage limits.** Many earlier web results are redacted. Assistant process summaries are summaries, not original experimental records. Other project-thread excerpts are not treated as the contents of those threads. No complete earlier generation history, implementation repository, review database, actual character wiki packet, or original model/prompt configuration was available. Relative repository links inside the files do not establish access to their targets. The files' own "inspected" labels are historical claims about their preparation, not fresh verification or reproduced experiments. The current product context supplied with the handover request is treated as owner-supplied context, not a historical decision.

**Central idea, in plain language.** A unit can contain plausible attacks yet feel wrong because its upgrade paths do not express a satisfying interpretation of the character. The user preferred a different overall organization for the Luffy example, not necessarily more elaborate individual attacks. The thread proposed comparing alternative interpretations before fully developing them, retaining what reviewers preferred and why, and using those records to guide later generation. This could begin with editable instructions and examples rather than model training. A leaderboard would summarize reviewed candidates; learning to judge unseen designs would be a separate capability. Whether this improves players' experience remains untested, and a concept review cannot establish gameplay balance or actual play feel.

### Reading and origin conventions

"Historical" means directly present in the visible conversation or supplied files. "Inference" identifies an interpretation of that material. "Current recommendation" and Section 7B are new advice or new sources. Accepted status below is limited to explicit user-stated priorities or preferences. No software implementation proposal received clear user approval.

Locators use the speaker and a distinctive phrase, or a supplied filename and section. No message dates or message numbers have been inferred.

## 2. Ideas, approaches, and changes of direction

### Specialized character-to-unit model

The user's initial ideal was a model taking a character description and producing a complete constrained unit. The assistant proposed a pretrained, code-capable generative LLM, initially guided by sources and examples, with later supervised adaptation. The expected benefit was more consistent translation and less correction. The dependencies were a useful target representation, reviewed transformations and an evaluation that could distinguish good adaptations from merely valid outputs.

The assistant discouraged foundation-model pretraining from scratch. It did not reject training small scorers or predictors from scratch. The user did not accept or reject either choice. Status: **unresolved user aspiration; proposed assistant implementation path**. Locators: user, "Pre-Train from scratch?"; first assistant, "Pretraining a language model from scratch is not the right starting point." Related: UG-T04-D07.

### Constrained semantic adaptation with independent checks

The first assistant proposed source evidence → Interpretation → typed mechanics specification → scoped deterministic checks and bounded repair. Runtime evaluation and numerical calibration would follow when a suitable Consumer existed. The aim was to make the adaptation explicit and prevent unsupported rules or vague effects from being presented as complete mechanics.

This could improve inspectability and scoped rule compliance. It also depends on defined operators and support checks and does not solve preference by itself. A complete DSL, compiler and combat runtime were not demonstrated or approved. Under the current owner context, this is an optional formalization route, not a prerequisite for every conceptual unit. Status: **proposed, not accepted**. Locator: first assistant, "Generate a mechanics specification, not just themed prose." Related: UG-T04-D08.

### Search, diversity and simulation as separate operations

The supplied research catalogue proposed retaining meaningfully different candidates, testing a mechanic's contribution and separating validity from behavioral effects and human preference. Quality-diversity search, novelty search, MCTS, ablation and numerical optimization appeared as possible methods or research precedents.

These could avoid selecting only familiar or high-damage designs. They require useful descriptors and, for gameplay claims, an implemented runtime and capable policies. Most were reading leads or later-stage adaptations, not selected algorithms. The catalogue explicitly did not authorize importing dependencies or changing the pipeline. Status: **retained as research context, implementation unresolved**. Locator: `GAME-DESIGN-RESEARCH.md`, "Proposed adaptations for our workflow" and R1 to R19. Related: UG-T04-E12.

### Refocus on interpretation and progression taste

The user said taste remained the main problem and provided the Luffy tree. They preferred paths organized by the three Haki, defining unlocks at Tier 3, and a paragon combining all three with Gear 5. They also said the individual abilities were not bad and that this was only an example.

The assistant inferred that composition, milestones and culmination were the important missing qualities. This is plausible, but the alternative changes several variables simultaneously. No comparison isolates which variable caused the preference. Status: **explicit priority and example-level preference; causal explanation untested**. Locators: user, "the main issue ... is the 'Taste'" and "This is just an example!" Related: UG-T04-D01, UG-T04-D02, UG-T04-E01 to UG-T04-E03.

### Pairwise review website

The user proposed repeatedly generating the same unit with prompt/system variations, presenting two candidates in a 1v1, and recording overall and aspect-level judgments with an Elo-like ranking.

The assistant developed this into comparisons at three levels: progression interpretations, complete units, and local branch/upgrade edits. It proposed retaining raw judgments, reasons, edits and exact revisions. This could expose useful patterns and create future training data. Its costs include reviewer effort, inconsistent judgments, confounded prompts, selection bias and the risk of learning presentation rather than design.

Status: **user proposal, expanded by assistant; no implementation commitment**. Locators: user, "have 2 generations in a 1v1"; second assistant, "make its core asset the preference data, not the Elo leaderboard." Related: UG-T04-D03 to UG-T04-D05.

### Ranking versus learning taste

The assistant distinguished a ranking of reviewed candidates, evaluation of generation strategies across matched requests, and a learned scorer for unseen unit content. It proposed Bradley-Terry estimation with uncertainty instead of relying solely on sequential Elo updates. The user did not reject Elo or accept Bradley-Terry.

The distinction matters because candidate ratings do not encode how to evaluate a new design. A learned ranker also cannot select a good interpretation absent from the generated pool. Status: **proposed statistical approach; unresolved ranking objective and aggregation policy**. Locator: second assistant, "There are three different systems here." Related: UG-T04-D06, UG-T04-E08.

### Feedback-guided tuning before weight training

The later assistant recommendation was reviewed examples and editable design guidance first, then a prompted or trained ranker, then supervised or preference training if beneficial. Laya was discussed as a possible bounded-choice component, not a whole-unit generator. Distillation was a possible way to make a reviewed multi-step process cheaper.

This reduces the initial training dependency but does not establish that configuration will be sufficient. No model, checkpoint, hosting arrangement or cost budget was selected. Status: **proposed sequence, not an accepted roadmap**. Locator: second assistant, "First: examples and explicit design preferences, without training." Related: UG-T04-D07.

### Terms and unresolved shifts in meaning

`CONTEXT.md` defines a **Unit** as content describing an actor or deployable combat element, not necessarily executable code. A **Unit draft** connects role, attacks, progression and abilities, with values when the supplied rules provide a basis. A **Mechanic** is declared behavior; a **Mechanic proposal** is an unapproved addition to a Game Definition. An **Interpretation** maps source material into organization and mechanics without becoming canon or a rule.

A **Profile** tunes permitted variation without overriding the Definition. **Validation**, **Review**, **Acceptance** and Consumer execution are different operations. These are supplied-file terms, not a claim that the implementation already satisfies them. Locator: `CONTEXT.md`, "Shared terms used here", "Unit terms", and "Generation and evaluation terms".

"Full unit" sometimes meant a complete conceptual design and sometimes an executable mechanics specification in the assistant's answer. The output choice remains unresolved. "Taste" includes a stated concept preference and an aspiration about actual player experience; those are not the same measurement. "Tuning" included prompts, examples, ranking and weight updates, with no chosen implementation. **"Usable default" was not operationally defined in the historical discussion.** The current owner context requires it, but the supported provider, running costs and minimum deliverable are not established here.

## 3. Decision record

Approval status and evidential support are separate. "Accepted" below records a stated user priority or preference, not an implementation authorization.

### UG-T04-D01: Treat perceived character fit and design taste as a central objective

**Status:** Accepted as an explicit user priority.
**Authority and locator:** User, "the main issue ... is the 'Taste'" and "feels like it should to the then actual player/user".
**Rationale:** Earlier generations were not subjectively good enough despite individually plausible abilities.
**Evidence/assumptions:** UG-T04-E01. The wider audience and the relation between concept preference and play feel are not established.
**Reconsideration condition:** None stated.

### UG-T04-D02: Preserve the Haki/Tier-3/paragon alternative as a preferred example

**Status:** Accepted as the user's example-level preference only.
**Authority and locator:** User, "each path corresponds to one of his 3 Haki", "at Tier3", and "paragon ... all three + gear5"; qualified by "This is just an example!"
**Rationale:** The user and a few friends did not find the displayed organization good.
**Alternatives:** The supplied mixed attack/transformation/utility tree versus the described Haki-organized alternative.
**Evidence/assumptions:** UG-T04-E01 to UG-T04-E03. No full alternative, controlled test, canonical verification or gameplay result.
**Reconsideration condition:** None stated. Do not generalize this into a required three-branch taxonomy.

### UG-T04-D03: Explore a pairwise human-review website

**Status:** Proposed by the user.
**Authority and locator:** User, "I also had the idea of creating a small website" and "elo like system".
**Rationale:** Compare generations and identify which designs and aspects work.
**Alternatives:** Overall wins and several aspect-level judgments were mentioned; the exact interface was not settled.
**Evidence/assumptions:** Assumes repeated reviews can produce useful preference data; no dataset or app shown.
**Reconsideration condition:** None stated.

### UG-T04-D04: Compare interpretations before full units, with local revision comparisons

**Status:** Proposed by the assistant.
**Locator:** Second assistant, "Compare the organizing idea before comparing complete units".
**Rationale:** Target the structural issue separately from upgrade detail and permit partial repairs.
**Alternatives:** Ranking complete units only; repeatedly regenerating everything.
**Evidence/assumptions:** UG-T04-E03; UG-T04-S07 offers a semantic-representation analogy, not validation of this interface.
**Condition discussed:** A preferred outline must still be reviewed after expansion because the full unit may disappoint.

### UG-T04-D05: Store attributed judgments and revisions rather than scores alone

**Status:** Proposed by the assistant.
**Locator:** Second assistant, "Preserve the raw records" and "A relative winner is not automatically a good training demonstration."
**Rationale:** Preserve disagreement, diagnose local failures and retain usable demonstrations and preference pairs.
**Alternatives:** A leaderboard alone, automatic component wins from an overall win, or inferred training pairs from rank order.
**Evidence/assumptions:** UG-T04-E07 to UG-T04-E09. The proposed rubric and overall-before-aspects ordering are untested.
**Condition discussed:** A recombination of preferred branches is a new candidate needing its own review.

### UG-T04-D06: Consider Bradley-Terry with uncertainty, without treating it as a taste model

**Status:** Proposed by the assistant; Elo remains an unresolved user option.
**Locator:** Second assistant, "start with Bradley-Terry estimation".
**Rationale:** Statistical comparison and uncertainty rather than sequence-sensitive rating updates.
**Alternatives:** Sequential Elo; separate per-reviewer and aspect results.
**Evidence/assumptions:** UG-T04-E08. Tie handling, regularization, comparison connectivity and group weights were not specified.
**Reconsideration condition:** No operational threshold stated. Different reviewers' preferences were to be retained rather than automatically treated as errors.

### UG-T04-D07: Prefer staged adaptation over immediate foundation-model training

**Status:** Proposed by the assistant; model training remains a user interest, not an approved method.
**Locators:** First assistant, "Pretraining ... from scratch is not the right starting point"; second assistant, "examples and explicit design preferences, without training".
**Rationale:** First identify the missing design decisions and obtain reviewed data. Test examples/configuration, then ranking and weight adaptation where useful.
**Alternatives:** Scratch language-model pretraining; direct generative LLM; LoRA/QLoRA; DPO; RL; small scorers; distillation.
**Evidence/assumptions:** UG-T04-E06, UG-T04-E09, UG-T04-E13, UG-T04-E14. No task-specific model comparison or sample-sufficiency evidence.
**Conditions discussed:** RL needs reliable executable evaluation; weight training should address a measured bottleneck; assess a trained scorer against a simpler prompted baseline.

### UG-T04-D08: Keep source interpretation, scoped checks and runtime evidence distinct

**Status:** Proposed architecture in the assistant answer; also a responsibility distinction in the supplied contract, not new user approval.
**Locators:** First assistant, "Train the model to make good adaptations"; `CONTEXT.md`, "Validation" and final responsibility paragraph.
**Rationale:** Avoid unsupported mechanics and avoid using a proposal's explanation as proof of correctness or balance.
**Alternatives:** Unconstrained prose and model self-certification.
**Evidence/assumptions:** UG-T04-E11, UG-T04-E19. The Definition, supported operators and Consumer are not supplied as working implementations.
**Conditions discussed:** Runtime claims wait for Consumer evidence. Current owner context makes code and formal mechanics optional, so the historical executable route cannot be an automatic v1 requirement.

### UG-T04-D09: Evaluate with held-out requests and account for the whole generation process

**Status:** Proposed by the assistant and in the research file.
**Locators:** First assistant, "Start with a measured pilot"; second assistant, "The first experiment I would run"; `GAME-DESIGN-RESEARCH.md`, final workflow comparison.
**Rationale:** Test transfer beyond repeated Luffy variants and avoid hiding sampling, repair or reviewer costs.
**Alternatives:** Same-character tuning only; general model benchmarks; per-response rather than per-useful-unit cost.
**Evidence/assumptions:** No evaluation performed. Historical quantities are preserved in Section 4.
**Conditions discussed:** Hold out whole characters/revisions; distinguish concept review from gameplay review. No approved success threshold.

No visible message approves an implementation architecture, ranking algorithm, training run, model checkpoint or commercial artifact split. The shift toward taste does not prove that the user rejected the earlier constraint-validity work.

## 4. Technical artifacts, research claims, and results

### Artifact and implementation inventory

| Item and locator | What is inspectable | Status and limitation |
| --- | --- | --- |
| UG-T04-E01, user Luffy table | Base attack plus 15 upgrades, three paths with five tiers each | Written output supplied by the user; reported as a generation-trial example. No original run metadata, validation or playtest. |
| UG-T04-E02, user preferred alternative | Haki branch organization, defining Tier-3 unlocks and combined paragon/Gear-5 payoff | Written conceptual preference, not a full alternative tree or tested unit. |
| UG-T04-S01, uploaded research catalogue | Full curated bibliography, stated access limits and project adaptations | Inspectable document; not implementation or reproduced experiments. |
| UG-T04-S02, uploaded context | Terminology, responsibility boundaries and a reported implementation limitation | Inspectable contract. Referenced code was not supplied. |
| UG-T04-D08 generation process | Textual workflow and proposed typed fields/checks | Suggested, not an implemented schema, compiler or running service. |
| UG-T04-D03 to UG-T04-D06 review site | Suggested comparisons, labels, provenance and ranking approaches | No UI, database, API, Elo implementation or Bradley-Terry fit shown. |
| UG-T04-D07 training plan | Candidate models, training methods and data strategy | No training code, checkpoint, dataset, reward model or results produced. |
| UG-T04-D09 pilots | Proposed counts and evaluation conditions | No participant recruitment, judgments or completed experiment shown. |

The separate `UG-T04-LUFFY-EXAMPLE.md` retains the supplied table and preferred-structure wording. It is a consolidation copy of visible user content, not an original generation log or a new design.

### Quantities and assumptions that must survive consolidation

**Original example.** The displayed tree has 16 entries: one base attack and 15 upgrades. Its first path moves through Armament Haki to Snakeman and fist-return mechanics; its second develops speed, kicks and temporary Boundman; its third moves from Observation Haki to range and Gear-5 terrain interactions. This inventory does not claim lore accuracy or rule validity. The preferred alternative places one Haki unlock at each of `3-0-0`, `0-3-0`, and `0-0-3`, then combines all three with Gear 5 in a paragon. The user did not provide its individual later upgrades.

**First assistant data budget.** 200 to 500 curated transformation records and 50 to 100 separate evaluation requests. These were explicitly planning figures, not measured sufficiency or an approved training budget.

**Later assistant review pilot.** Eight development characters, four interpretations per character, all six pairs per character and three reviewers: `8 × 6 × 3 = 144` judgments. Then expand two alternatives per character and reserve four additional characters for a held-out comparison. These are unapproved workload figures, not recruited participants, collected data or a power calculation.

**Model options.** Qwen3-14B and Qwen3-8B were proposed reference candidates; the 8 to 14B range was an engineering starting point. Qwen3.5-9B appeared only in earlier research-tool exploration. No model was selected or benchmarked. LoRA, QLoRA, SFT, DPO, RL and distillation were discussed, not implemented.

**External numbers, not project results.** Delta-Engines reports 175 roles and 502 co-designed training examples with CodeGemma-7B/LoRA. The inspected method gives rank 8, alpha 32, learning rate 1.5e-4, batch size 4 and five epochs. These are that paper's settings, not a configuration adopted here. The Pokémon study reports 49 participants and 196 cards with Qwen3-14B and retrieval over 993 cards. Laya's card reports a 421M root model and a 512-token question budget; its typed-decisions scores are 0.362 for the root, 0.461 for a majority baseline and 0.766 for a specialized checkpoint. These are vendor/task-specific results, not measured Unit-Generator quality. See UG-T04-S04 to UG-T04-S06.

**Other important inherited research quantities.** The catalogue reports Ludi's 1,389 evolved games, 19 viable games and 127 paired comparisons from 27 participants (UG-T04-S102); GAVEL's roughly 48-hour search runs on an RTX8000 plus 16 CPU cores and 128 GB memory, with novelty-cell improvements not significant over three runs (UG-T04-S103); and Mortar's 14 participants comparing six games in three pairs, with proxy/preference agreement on two pairs but disagreement on the third (UG-T04-S117). Those details are attributed to the supplied catalogue, not freshly reproduced results. They do not specify suitable latency, sample size or a taste predictor for this project.

The second assistant's "five friends / 1,000 votes" sentence was hypothetical. Actual reviewer count, number of prior generations and measured failure rate are **not established in this thread**.

### Proposed data records, not implemented schemas

The suggested candidate record includes source scope, resolved Game Definition/Profile revisions, brief, model and settings, concrete prompts, immutable output revision, parent revision and selected Interpretation. The proposed review record includes reviewer, compared revisions, comparison level, overall choice, optional aspect choices, whether either candidate is worth developing, explanation or edit, display order and timestamp.

Suggested overall labels were A, B, no preference and cannot judge. Separate development-worthiness labels were A, B, both or neither. Suggested aspects were character fit, path organization, progression/payoff, individual upgrades and clarity. Neither this rubric nor the order of questions was validated.

The research file's candidate-evidence fields are explicitly a documentation proposal, not a required API schema. No executable schema version, database migration, Elo K-factor, tie model, regularization strength, stopping criterion or reward equation was selected.

### Evidence and claim register

**UG-T04-E01, reported taste failure.** The user supplied an actual text example and reported dissatisfaction by themselves and a few friends. Support: visible user message, "All of the abilities ... arent bad" and "subjectively found 'good'". Strength: direct report plus inspectable design text. Limitation: no controlled evaluation or measured frequency.

**UG-T04-E02, preferred organization.** The user explicitly preferred the Haki/Tier-3/paragon structure for the example. Support: same message. Strength: direct preference. Limitation: does not establish a universal principle, a complete implementation or audience consensus.

**UG-T04-E03, organization as the causal problem.** The assistant inferred that composition, milestone placement and culmination explain the preference. Support: UG-T04-E01 and UG-T04-E02. Strength: plausible interpretation. Limitation: several changes are confounded; no isolating test.

**UG-T04-E04, demonstrated project work.** The available project artifacts are text, examples and research/context files. No working generator modification, review website, votes, trained checkpoint or simulation result is shown. This is an observation about coverage, not a claim that no implementation exists elsewhere.

**UG-T04-E05, adjacent transformation precedents.** Pokémon cards, semantic game generation, grammatical cards, description-to-code and partial card completion address parts of the task. Sources: UG-T04-S05 and UG-T04-S07 to UG-T04-S10. Limit: different inputs, representations and evaluations; no one source validates the complete proposed system.

**UG-T04-E06, specialist training precedent.** Delta-Engines demonstrates a task-specific training route with co-designed examples. Source: UG-T04-S06. The reported code-accuracy judgment is by GPT-4 and conditional on execution success. This is not a measured improvement in this group's taste or fixed-Definition TD validity.

**UG-T04-E07, human and fine-grained feedback as supervision.** Experience-driven PCG and fine-grained language feedback support investigating attributed preference data. Sources: UG-T04-S22, UG-T04-S23 and the older UG-T04-S139. They do not establish the proposed rubric, feedback volume or achievable TD preference accuracy.

**UG-T04-E08, ratings are not content generalization.** Chatbot Arena describes paired comparison, Bradley-Terry estimation and uncertainty; Elo Uncovered examines ordering/hyperparameter sensitivity. Sources: UG-T04-S24 and UG-T04-S25. The claim that a candidate-ID score does not itself evaluate a new unit is a distinction in the proposed system, not a new paper result. No ranking algorithm was run here.

**UG-T04-E09, personalization without immediate weight updates.** SynthesizeMe! supports testing history-derived prompts/examples; AMPLe supports comparison-based learning over a supplied multi-attribute representation. Sources: UG-T04-S26 and UG-T04-S27. Neither establishes automatic discovery of TD taste dimensions or transfer to this group.

**UG-T04-E10, score optimization can exploit an imperfect evaluator.** Reward-model overoptimization is reported for RL and best-of-N selection in a synthetic gold-reward setup. Source: UG-T04-S28. This is a risk analogy, not an observed Unit-Generator failure.

**UG-T04-E11, scoped checks are not preference or balance.** The distinction is explicitly stated in `CONTEXT.md`, "Validation", the conceptual/mechanics Deliverable paragraph and the ownership paragraph. Source: UG-T04-S02. This is a supplied responsibility contract, not evidence of compliant implementation.

**UG-T04-E12, game evaluators and human preference can disagree.** The supplied catalogue records Mortar's mixed agreement and Mechanic Miner's comprehension problem; Mechanic Maker 2.0 illustrates evaluator-dependent discoveries and fitness exploitation. Sources: UG-T04-S117 to UG-T04-S119, inherited through UG-T04-S01. Full underlying experiments were not re-audited in this handover.

**UG-T04-E13, bounded decision models have a different role.** Laya's official card describes scoring supplied choices and reports checkpoint-specific results. Source: UG-T04-S04. The conclusion that it is a possible ranker rather than the default full-unit generator is an assistant recommendation; no comparative test occurred.

**UG-T04-E14, the proposed training sequence is not an optimum.** LoRA, QLoRA and DPO provide adaptation methods; distillation offers another route. Sources: UG-T04-S15, UG-T04-S16, UG-T04-S18 and UG-T04-S19. Their existence does not prove that the proposed model size, sample budget or sequence is sufficient or cost-effective for this task.

**UG-T04-E15, absence of an end-to-end example is a bounded search result.** The first assistant said no verified full wiki-to-faithful/executable/balanced-TD system had been found. That supports a literature gap to investigate, not proof of nonexistence, project novelty or publication eligibility.

**UG-T04-E16, new review-bias evidence.** Newly found UG-T04-S301 reports that human preferences can respond to presentation and miss correctness defects. Transfer is a reason to separate preference from checks, not evidence that these reviewers were biased.

**UG-T04-E17, new challenge to a universal scalar ranking.** Newly found UG-T04-S302 discusses general preferences outside standard scalar Bradley-Terry assumptions. No cycles or stable subgroups were measured here. UG-T04-S24 already discusses misspecification and a nonparametric extension; this is not a blanket rejection of Bradley-Terry inference.

**UG-T04-E18, new direct designer-preference precedent with limitations.** Newly found UG-T04-S303 studies learned preferences in a dungeon editor and reports noisy/sparse-data and usage limitations. Different content and feedback labels limit transfer; more votes are not established as a guaranteed solution.

**UG-T04-E19, reported repository limitations are unverified code claims.** `CONTEXT.md` says `src/core` still contains a bounded numerical implementation and the experimental RulePack API is a layout descriptor, not a complete executable Definition package. Source: UG-T04-S02, final scope note; unavailable code at UG-T04-S208. No code inspection, test or compatibility check confirms its current state.

## 5. Recommended approach and v1 boundary

### Recommendation actually reached in the thread

There was **no user-approved implementation recommendation**. The latest assistant recommendation was to build the review loop before a training pipeline, compare alternative interpretations, retain raw preferences and edits, and use those records to guide generation. This changed the emphasis of the earlier model-centric advice without establishing that constraint checks or later training were rejected. See UG-T04-D03 to UG-T04-D09.

### Current recommendation

The following is new advice based on this thread and the current owner context. It awaits approval and is not the final architecture for Mardwerk.

**Product principle:** Help a user obtain and refine a recognizable, game-compatible character adaptation. Make preference choices editable, and do not present preference, legality or balance as interchangeable.

The smallest useful workflow is: supply character evidence and a Game Definition, generate an explicit interpretation and complete conceptual unit using the public default, inspect or change the interpretation, regenerate the affected design, and save/export the result. Alternative comparison should be available for improvement, but neither prior votes nor a private preference model should be required to obtain the first useful unit.

| Proposed v1 component | Why it is necessary here |
| --- | --- |
| One documented public default Game Definition/configuration and character-evidence input | A novice needs a workable example; the generator needs explicit rules rather than inventing game assumptions. Other Definitions remain an extension question. |
| One supported, documented model connection with stated running requirements | The standalone project must actually execute its default workflow. Provider, hosting, hardware and costs still require a choice; FOSS does not settle them. |
| A short editable Interpretation inside the generation result | Directly addresses UG-T04-E03 and lets a user preserve branch organization or milestones while revising detail. It need not require multiple agents or a separate model. |
| A complete conceptual unit output with explicit omissions and available numerical basis | Delivers the basic product promise without requiring code, assets, optional mechanic invention or a simulator. New mechanics are generated only when requested. |
| Checks limited to requirements actually encoded by the supported Definition | Distinguishes pass, fail and not checked. Avoids silently turning unsupported proposals into valid units without building a general compiler first. |
| Local revision, review and export records | Makes comparison, correction and continued standalone use possible without Towerright-private history or training assets. |
| A small A/B review mode for internal evaluation | Tests whether interpretation-first guidance helps and captures raw evidence. A hosted social site, global leaderboard and automatic preference learner are not necessary for this first test. |

Build and test the basic public workflow now. In parallel, use the review mode to compare meaningful alternatives, retain user edits and identify recurring failures. The review mode can be a simple experimental utility; it need not become the mandatory main user journey.

**Potential blockers.** The minimum deliverable and supported rule scope are not fixed; no actual source packet or default provider/cost arrangement is supplied; the existing repository scope is only reported; disclosure rights for example data and tuning artifacts are unresolved. These are choices or inspections needed for a usable prototype, not reasons to require a full publication study.

**Defer or exclude from v1 unless a concrete failure demands them.** Foundation-model pretraining, mandatory LoRA/DPO/RL, a trained reward model, large evolutionary/QD search, active-pair optimization, a global Elo leaderboard, an all-purpose mechanics compiler, combat simulation, generated executable game code, visual assets, platform integration and paper-ready participant studies.

**Public/private boundary.** The current owner context places the specially tuned offering inside Towerright. The historical contract assigns Towerright project-wide orchestration, evaluation and Acceptance; it separately allows a Lab to retain local work. This does not establish that public users must lose local saving, revision or choice of usable content. A public default must contain enough guidance and supported configuration to complete the basic workflow without private examples, weights or services owned by Towerright.

Exactly which prompt sets, review datasets, preference profiles, weights, evaluation suites or review-website features are private is **not established in this thread**. Treat Foundation as private, and this consolidation as internal. Decide the boundary for each artifact before release. No software, model, data or asset license was selected or found compatible by this thread.

## 6. Open questions, conflicts, and risks

### UG-T04-Q01: Whose taste should the system improve?
**Type:** Product choice and empirical uncertainty. **Impact:** A scoped v1 objective is needed; general audience claims are later.
Kyle, a few friends, fans and eventual players are not an interchangeable population. Resolve by naming the initial target reviewers and separating group preferences from individual differences. Related: UG-T04-D01, UG-T04-E01, UG-T04-S22.

### UG-T04-Q02: Which structural change caused the Luffy preference?
**Type:** Research uncertainty. **Impact:** Does not block a manually guided prototype; affects learned generalization.
Branch identity, synchronized unlocks and the final Gear-5 payoff changed together. Compare controlled alternatives or collect before/after edits before claiming one general rule. Related: UG-T04-D02, UG-T04-D04, UG-T04-E03.

### UG-T04-Q03: What is the minimum complete output and rule scope?
**Type:** Product/technical choice. **Impact:** Blocks a testable v1 boundary.
The historical discussion alternates between conceptual units and executable specifications. Approve a minimum deliverable, supported Definition, numerical basis and completeness checklist. Do not import a simulator requirement. Related: UG-T04-D08, UG-T04-E11, UG-T04-E19.

### UG-T04-Q04: What source evidence governs an adaptation?
**Type:** Input and research-design choice. **Impact:** Blocks reproducible tests; need not block a prototype accepting user text.
No Luffy wiki entry, edition, cutoff or contradictory-source policy was supplied. Preserve exact source packets and distinguish canon claims from adaptation choices. Related: UG-T04-E02, UG-T04-S211.

### UG-T04-Q05: What makes the public default runnable?
**Type:** Feasibility/product choice. **Impact:** Blocks standalone v1 delivery.
Select and test one documented model connection and state access, hardware, cost and network requirements. No claim of offline or free operation follows from FOSS. Related: UG-T04-D07, UG-T04-S04, UG-T04-S13, UG-T04-S14.

### UG-T04-Q06: How should ties, disagreement and reviewer influence be represented?
**Type:** Statistical and product choice. **Impact:** Basic raw review can ship first; any trusted aggregate ranking depends on a decision.
Elo versus Bradley-Terry, equal reviewers versus equal votes, aspect aggregation and cyclic preferences remain unresolved. Inspect paired data and per-reviewer patterns; retain raw outcomes, counts and display order before choosing stronger assumptions. Related: UG-T04-D05, UG-T04-D06, UG-T04-E08, UG-T04-E17.

### UG-T04-Q07: Are we selecting candidates, evaluating strategies or predicting unseen preferences?
**Type:** Capability definition. **Impact:** Affects the review workflow; learned prediction can wait.
These require different inputs and evidence. State the target for each experiment. A candidate leaderboard must not be described as a trained taste model. Related: UG-T04-D06, UG-T04-E08.

### UG-T04-Q08: Which tuning and review artifacts are public or Towerright-private?
**Type:** Commercial/disclosure choice. **Impact:** Blocks release of affected artifacts, not all local prototyping.
Prompts, demonstrations, reviewer histories, weights, preference profiles, evaluation cases and website implementation have no settled ownership/disclosure split. Approve an artifact-level inventory while preserving an independently useful public default. Related: UG-T04-D07, UG-T04-S207.

### UG-T04-Q09: Does Towerright's project ownership leave a complete standalone local workflow?
**Type:** Scope ambiguity. **Impact:** Blocks public workflow design if interpreted restrictively.
`CONTEXT.md` assigns project history and Acceptance to Towerright but local work to a Lab. Clarify the boundary rather than inferring that a public user cannot save revisions or choose a design. Related: UG-T04-E11, UG-T04-S02, UG-T04-Q08.

### UG-T04-Q10: When, and on what evidence, should training begin?
**Type:** Research/cost choice. **Impact:** Later work, not a prerequisite for basic v1.
No data sufficiency or model advantage is established. Compare example-conditioned generation/ranking with a trained variant on held-out characters, including setup and inference costs. Related: UG-T04-D07, UG-T04-E09, UG-T04-E10, UG-T04-E18.

### UG-T04-Q11: Will a preferred concept feel right in play?
**Type:** Empirical uncertainty. **Impact:** Blocks gameplay-feel and balance claims, not conceptual generation.
Collect separate judgments from actual supported runtime previews/playtests. Do not use concept votes or incapable agents to certify play experience. Related: UG-T04-D01, UG-T04-E11, UG-T04-E12.

### UG-T04-Q12: What constitutes a fair improvement over direct LLM prompting?
**Type:** Evaluation design. **Impact:** Needed for the claimed advantage and later paper; a small comparison can accompany v1.
Control source and rule information, disclose prompt knowledge, record all attempts/repairs and separate tuning from evaluation characters. Blind presentation where practical. Related: UG-T04-D09, UG-T04-E16.

### UG-T04-Q13: What can be reused or disclosed?
**Type:** Rights, consent and publication boundary. **Impact:** Blocks release/training use of affected content; does not establish any legal violation.
Character material, examples, friends' reviews, third-party datasets and model/code terms require separate review. No license audit or permission record is available. Resolve with provenance, intended-use decisions and appropriate rights review, not assumptions from public URLs. Related: UG-T04-S21, UG-T04-S01, UG-T04-Q08.

### UG-T04-Q14: How are optional new mechanics separated from supported behavior?
**Type:** Product/technical choice. **Impact:** Blocks only the affected optional formalization or mechanic-proposal feature.
Define when the result is a conceptual adaptation, a checked supported specification or an unsupported proposal. A new operator name cannot count as implementation support. Related: UG-T04-D08, UG-T04-E11, UG-T04-S115.

## 7. Sources and research leads

The complete register is supplied in `UG-T04-SOURCE-REGISTER.md` and inside `UG-T04-HANDOVER.md`. It contains 105 records, with direct URLs, provenance, publication identifiers where recoverable, separate bibliographic/content-inspection labels, relevance, limitations and artifact links. No source in that register is an adoption decision.

### A. Sources already present in this thread

UG-T04-S01 and UG-T04-S02 are the two supplied files. UG-T04-S03 to UG-T04-S28 cover the directly linked guidance/model materials and the main transformation, training and preference literature cited in the answers. UG-T04-S29 to UG-T04-S43 retain additional identifiable references from visible research-tool requests, including unresolved leads.

UG-T04-S101 to UG-T04-S138 preserve all 38 numbered research-catalogue entries. UG-T04-S139 to UG-T04-S148 preserve its ten additional named papers. UG-T04-S201 to UG-T04-S211 identify unavailable project references and the missing character source packet rather than inventing their contents.

Priority reading for this thread's central idea is UG-T04-S22 (experience-driven PCG, 2015 extended abstract), UG-T04-S23 (fine-grained feedback), UG-T04-S24 and UG-T04-S25 (ranking and uncertainty), UG-T04-S26 and UG-T04-S27 (preference personalization), and UG-T04-S28 (reward overoptimization). The strongest adjacent generation precedents remain UG-T04-S05 to UG-T04-S10.

Important verification outcomes: the 2015 experience-driven PCG extended abstract is distinct from the 2011 article at UG-T04-S139. The primary Ludax v2 title includes "Domain Specific Language" (UG-T04-S115); PCGRL+ is part of the manuscript title (UG-T04-S128); and the Hearthstone deck paper's first author is Matthew C. Fontaine, not the attachment's "Michael", with "through MAP-Elites" in the title (UG-T04-S146). Original wording remains in the register. The Game-O-Matic DOI is not the Program Synthesis DOI.

Incomplete references to RPGAgent, Progression Balancing, 1001 Nights and Damage optimization in video games remain unresolved. Several publisher/full-text access gaps remain. A verified title is not a verified finding, and none of the experiments or repositories was reproduced.

### B. Newly identified literature leads

**UG-T04-S301, newly found, not part of the original conversation.** Tom Hosking, Phil Blunsom and Max Bartolo, *Human Feedback is not Gold Standard*, ICLR 2024, arXiv:2309.16349v2. Bibliography verified; relevant Sections 2 to 4 inspected. It challenges treating human preference as an independent correctness check. Its language-task results do not establish bias in this review group.
https://arxiv.org/html/2309.16349v2
https://openreview.net/forum?id=7W3GLNImfS

**UG-T04-S302, newly found, not part of the original conversation.** Yixin Liu, Argyris Oikonomou, Weiqiang Zheng, Yang Cai and Arman Cohan, *COMAL: A Convergent Meta-Algorithm for Aligning LLMs with General Preferences*, 2024 preprint, arXiv:2410.23223v2, October 13, 2025. Bibliography verified; relevant preference formulation inspected, not a proof audit. It questions universal scalar rankings under general preferences. It does not justify adding its training algorithm to v1.
https://arxiv.org/html/2410.23223v2

**UG-T04-S303, newly found, not part of the original conversation.** Alberto Alvarez and Jose Font, *Learning the Designer's Preferences to Drive Evolution*, EvoApplications 2020, arXiv:2003.03268v1, DOI:10.1007/978-3-030-43722-0_28. Bibliography verified; relevant Sections 3 to 5 inspected. It studies preference learning in a dungeon editor and reports data/usage limitations. Its 15-student study and indirect labels do not determine the feedback needed for TD character adaptation.
https://arxiv.org/html/2003.03268v1
https://doi.org/10.1007/978-3-030-43722-0_28

Remaining gaps include direct character-to-TD preference evidence, transfer from outline judgments to complete units and play feel, reviewer disagreement, data sufficiency and full-text retrieval for unresolved leads. This extraction is not an exhaustive literature review.

## 8. Preparation for later research

### Existing plans in the thread

The assistant proposed comparisons across direct generation, structured generation with checks and a trained variant. The research file proposed fixed source packets, equal total budgets including repairs and separate measures for validity, fidelity, useful differences, correction effort and cost. The later taste answer proposed interpretation-level comparisons, held-out characters and the 144-judgment pilot retained in Section 4.

None of these plans was approved or run. The current owner's later-paper goal does not convert them into required v1 study protocols.

### Current suggestions, awaiting approval

**Research question.** Does explicit comparison and retention of character interpretations improve the usefulness and perceived character fit of generated TD upgrade trees, compared with a competent direct-LLM-plus-wiki workflow under the same game constraints?

**Potential contribution.** A representation and feedback process that links source evidence, progression organization, local design edits and attributed preferences could be worth studying. The thread does not establish that this combination is novel. It may also reveal when interpretation-level feedback fails to predict preference for complete units.

**Evidence versus hypothesis.** The user preference and dissatisfaction report are real thread evidence, UG-T04-E01 and UG-T04-E02. The proposed causal explanation, improved generalization, useful learned taste, required data volume, lower correction cost and better player experience remain hypotheses. External precedents supply methods and warnings, not Unit-Generator results.

**Smallest useful next evaluation.** First compare genuinely different interpretations for the same source and rule packet, including the user's preferred structure and credible alternatives. Retain whole-design preferences, development-worthiness, reasons and edits. Expand promising alternatives and check whether outline preference survives in full units. Then compare the revised process against the unchanged competent baseline on characters not used to write the guidance. No new sample size or success threshold is approved by this handover.

### Baseline and resource accounting

The direct-LLM baseline should receive the same character source material, Game Definition, brief, requested deliverable and relevant known user preferences. It should use a competent instruction to produce a full unit, not an intentionally vague or rule-starved prompt. If one condition is explicitly told "one Haki per path" while the other is not, label the result as an instruction/personalization difference rather than independent discovery of better structure.

Record the exact model/provider/version, source and Definition revisions, prompts, decoding settings, retrieval examples and output constraints. For a controlled method comparison, hold the model and available information constant and account for the complete attempt/repair/selection budget. For an actual workflow comparison, report differences in sampling, human selection, setup/training, latency and cost rather than pretending the resources are identical.

Do not cherry-pick a best-of-many pipeline output against the first baseline response without charging the difference. Do not force both systems through private guidance while calling the result a test of the public default.

### Proposed outcomes, not approved success criteria

The thread supports measuring scoped constraint satisfaction and completeness; source fidelity and visible adaptation losses; preferred path organization and progression payoff; controllability through revisions; distinct useful interpretations; designer correction time or edit burden; and total cost/latency per useful result.

Record ties, uncertainty and whether neither candidate is worth developing. A relative winner need not be an acceptable unit. Separate human preference from independent checks. A polished explanation is not evidence that a mechanic works. Reserve gameplay effect and balance measures for explicitly defined, executed Consumer scenarios; preserve the policy and scenario limits of those results.

A later study should report reviewer population and disagreement rather than claiming universal taste. It should test previously unseen characters, avoid leakage across paraphrases/revisions and include cases where the preferred Luffy organizing pattern is inappropriate. Whether supplied-evidence counterfactuals are followed can test grounding separately from memorized character associations.

### Records to retain

Retain code/strategy revisions; provider/model/checkpoint and configuration versions; prompts and retrieved examples; source packets and hashes; Game Definition/Profile revisions; all attempts, failed checks and repairs; outputs and parent revisions; candidate-selection decisions; rubric versions; attributed reviews, display order, ties and reasons; costs and human effort; and the tuning/evaluation split.

If runtime testing is later added, retain Consumer version, scenario/map/wave/roster, policy, seeds, trial counts and raw traces. Distinguish predicted outcomes from observed ones. Data retention does not itself authorize public disclosure.

**Public and private reproducibility.** Report results from the public default separately from results using Towerright-private prompts, data, profiles, weights or evaluation assets. Publishable descriptions can identify unavailable dependencies and methodological limits without assuming those assets must be released. A private advantage is not evidence of the public default's performance; a runnable public baseline does not reproduce private tuning.

## 9. Consolidation handoff

**Do not lose:** the actual reported failure and its limited evidence, UG-T04-E01; the example-specific preferred structure, UG-T04-D02; the distinction between interpretation quality and individual upgrade quality, UG-T04-E03; and the separation of raw preferences, aggregate ratings and prediction for unseen designs, UG-T04-D05 and UG-T04-D06.

**Owner decisions still needed:** approve or decline the review-loop proposal, UG-T04-D03 to UG-T04-D05; choose the minimum deliverable and runnable default, UG-T04-Q03 and UG-T04-Q05; define the initial taste audience, UG-T04-Q01; and resolve local/public versus Towerright-private artifacts, UG-T04-Q08 and UG-T04-Q09. No selected algorithm or training commitment should be carried into another thread as settled.

**Next concrete action:** assemble one frozen source-and-Definition request, retain the supplied negative example and preferred-structure record, and use a minimal local comparison/revision workflow to evaluate credible alternative interpretations. In parallel, make the public default complete that request without private tuning. Preserve the evidence needed for UG-T04-Q02 and UG-T04-Q12 before committing to model-weight training.

---

### File and source-register integrity

The supplied files were read at their actual mounted paths. Their hashes identify the reviewed attachments, not repository commits.

- `/mnt/data/GAME-DESIGN-RESEARCH.md`: SHA-256 `8b1611fa74f65009d3c71644220f74880a46d95457302888ce3bd02e4e7c27cc`.
- `/mnt/data/CONTEXT.md`: SHA-256 `fee3f4ebbcaf0f7c44013db8102ca98510676e8723a399e95ee323f7a946ee50`.

The source register retains all 38 numbered catalogue entries, its ten additional named papers, the identifiable substantive references in the visible answers and research requests, and unavailable internal/source references. It does not expand every cited paper's entire bibliography. Generic search phrases with no identifiable source, inaccessible other-thread contents, redacted tool results, unavailable repository files and unprovided wiki material are not reconstructed. This is a traceable consolidation with a bounded literature check, not an exhaustive literature review.
