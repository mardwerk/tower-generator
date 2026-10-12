# AI taste and autonomous generation of units, mechanics, maps, and games

## 1. Identity, coverage, and research question

**Provisional reference key: UG-R-game-design**

This is a handover label, not an actual ChatGPT conversation ID or a guarantee of uniqueness. No assigned thread ID or original thread title is visible in the current-thread messages.

The available record contains two substantive user requests, two assistant literature reviews, and the current handover request. There are no visible attachments, project repositories, execution logs, generated game files, or prototype results. The earlier assistant responses are research summaries and proposals, not primary experimental evidence. Their original browsing traces are unavailable.

External checks for this handover were conducted on September 24, 2026. Inspection depth varies by source and is recorded in Section 8. No repository was executed. Some publisher pages were inaccessible; author manuscripts or abstracts were used where available.

Every record belongs to the `UG-R-game-design` namespace. Full IDs appear at their definitions. Short references such as E02 or S06 below refer only to records under that prefix.

### Original and expanded questions

The original question concerned whether AI can propose genuinely novel solutions and exercise useful design judgment, especially for units and games. The user explicitly wanted outputs that could be inspired by existing material while remaining novel. The assistant interpreted "taste" as selecting ideas that are interesting, coherent, and appropriate to a design vision. That interpretation was an assistant framing, not an agreed technical definition. Sources: S01 and S02.

The second request expanded the investigation to autonomous creation of 3D maps, units, mechanics, and entire games, especially RTS and tower defense. Questions about executable representations, automated playtesting, human preference, engine integration, and autonomy emerged during that expansion. Sources: S03 and S04.

The thread's contribution is a literature map and a set of evaluation cautions. It does not establish a working Unit-Generator or select its generation approach.

### Current project context, supplied now

S05 establishes Mardwerk as the broader ecosystem, Towerright as its private, closed-source SaaS platform, and Unit-Generator as the component under research. The intended output is a tower defense unit, with mechanics potentially included when requested. Research into maps and entire games does not expand that product scope automatically.

The intended public product is standalone FOSS with a usable, configurable default. Its basic workflow must not require private Towerright components. This does not establish offline operation, free operation, or independence from external model APIs. The specially tuned Towerright offering is intended to remain private, but the meaning of "tuned" is unresolved.

The immediate goal is a functioning prototype. A later paper is an ambition. Superiority over directly prompting an LLM with reference information remains untested. Towerright-first development followed by extraction is a proposed sequence, not an approved architecture or a decision made in the earlier discussion.

## 2. Findings and competing approaches

All approaches below were explored as research possibilities. None was approved or rejected by the owner as Unit-Generator's architecture.

### Search within a supplied mechanical vocabulary

The RTS unit paper searches numerical attributes and trigger-effect abilities. The TD papers generate maps, tower locations, and waves inside human-designed games. Formal mechanic synthesis instead expresses actions through explicit preconditions and effects. These approaches require a representation, permitted operations, and evaluable objectives.

Their advantage is that candidates can be executed and checked. Their limitation is that the representation determines what can be invented. They apply when the relevant design choices can be expressed without repeatedly changing the underlying system. Evidence: E02, E07, E09.

### Code-level mechanic invention

Mechanic Miner changes accessible program state. Pixie generates code around annotated fields and insertion points in Unity projects. Mortar uses language models as code-variation operators and evaluates mechanics through assembled games.

This approach permits interactions beyond a small catalogue of named abilities, but still depends on accessible code, test environments, and evaluation objectives. Pixie also shows that code generation need not involve an LLM. Successful experimental code may require reimplementation before production use. Evidence: E05 and E09.

### Learned recombination and quality-diversity search

Conceptual Expansion learns approximate representations of existing games and recombines them. GAVEL uses a language model to modify executable Ludii rules and maintains different promising candidates through quality-diversity search.

These methods depend on seed examples, a workable representation, and criteria for retaining candidates. The unresolved question is whether their differences correspond to meaningful gameplay differences rather than different syntax or inactive rules. Reconstruction of existing games is weaker evidence for invention than successful evaluation of new ones. Evidence: E03.

Quality-diversity concerns search and retention. It is not an alternative to using an LLM, a simulator, or a human evaluator. Those choices could coexist, but this thread did not establish that the extra machinery is worthwhile for Unit-Generator.

### Different accounts of design judgment

The conversation considered several ways to select designs: hand-written quality measures, agent-performance differences, models of player experience, human preferences, and foundation-model judgments of interestingness.

These are not interchangeable definitions of taste. The human-invention study compares associative proposal models with models that incorporate simulated gameplay. OMNI uses foundation models to prioritize interesting learning tasks. Mortar uses differences in agent performance, with a small human comparison. Each addresses a different evaluation target. Evidence: E04 through E06.

### Learned level construction and structured 3D assembly

PCGRL learns a policy that edits levels. Its 3D extension incorporates movement constraints. The database-driven 3D framework instead combines reusable components, structural constraints, and repair procedures.

These approaches assume computable validity conditions and a restricted environment. Infinigen and Holodeck address scene creation and placement, not competitive RTS balance. Visual coherence, navigability, and strategic quality remain separate questions. Evidence: E08.

Reinforcement learning as a level generator is also distinct from reinforcement learning as a playtester. PCGRL and Mechanic Maker 2.0 use it for different jobs.

### Autonomous environment creation and engine prototyping

OMNI-EPIC chooses and implements subsequent learning environments. DreamGarden provides an editable, semi-autonomous Unreal workflow. AutoUE coordinates asset retrieval, scene construction, gameplay implementation, and testing.

These systems address task selection and implementation breadth. They do not, by themselves, establish reliable design judgment for RTS or TD. Genie represents another output category: action-conditioned generated worlds rather than conventional editable game rules and engine projects. Evidence: E10 and E12.

The thread did not contain an experiment comparing a "purely agentic" architecture with a "conventional pipeline." That binary should not be imported from another conversation.

## 3. Claim and evidence record

The external findings below were reported in S02 or S04 and checked again during this handover. They remain results reported by the paper authors, not replications or Mardwerk results.

### UG-R-game-design-E01. No project implementation or validated advantage is established

Origin: visible thread, S01 through S05.

What is inspectable consists of literature summaries, illustrative prompts, evaluation suggestions, and a proposed experiment. No code execution, generated unit dataset, human test, or comparison against direct prompting appears. Absence of visible results does not establish that no work exists elsewhere.

### UG-R-game-design-E02. Restricted RTS unit generation has been demonstrated

Sources: S06, originally discussed in S02 and S04.

The paper generated ten microRTS units using human-specified triggers and effects and evaluated them with automated matches. Methods, results, and limitations were inspected. Evaluation used one map, and the strongest agent often neglected generated units. Its supposed skill levels were not validated as human skill models. General RTS usefulness, TD transfer, human enjoyment, and historical originality remain unverified. ([arXiv][1])

### UG-R-game-design-E03. Independent generation does not establish historical originality

Sources: S07, S08, S35.

Ludi produced 1,389 candidates and identified 19 viable games. Its authors reported a new Yavalath rule but also a rediscovery of Irensei through Ndengrod. GAVEL produced executable rule combinations, but its advantage in novel occupied search regions was not statistically significant across three runs. Conceptual Expansion evaluated reconstruction of existing games. These are different strengths of evidence, not equivalent demonstrations of novelty. ([Cameron Browne][2])

### UG-R-game-design-E04. Simulation-informed judgment is supported, but not established as universal taste

Sources: S10, S11; theoretical context S17.

In the human-invention study, a model incorporating simulated gameplay better explained submitted games than the examined proposal-only account. The inspected work concerns over 450 novice-created, grid-based games and population-level model fit. It does not reveal participants' rejected ideas or establish that they literally used the proposed computation. Experience-driven PCG is a framework for conditioning generation on player experience, not a universal evaluator. ([arXiv][3])

### UG-R-game-design-E05. Agent-performance measures and human preference can disagree

Source: S29.

Mortar evaluates mechanics through their contribution to games and skill-based agent ordering. Its human comparison involved 14 participants, six games, and three pairs. Automated ordering and aggregate preference agreed in two comparisons and disagreed in one. The paper explicitly describes a mechanic-generation research prototype, not a system intended to generate complete games. Its metric is not a validated substitute for player preference. ([RAIL Lab][4])

### UG-R-game-design-E06. Novelty ratings, feasibility, and diversity can move differently

Sources: S12 through S14.

The research-idea study found higher expert novelty ratings for AI proposals, alongside self-evaluation and diversity problems. The writing experiment found improved individual evaluations but greater similarity between AI-assisted stories. The cross-model creativity study also found population-level homogeneity. These findings challenge reliance on fluent output or model variety alone. They do not establish the same effects in TD unit generation. ([ICLR Proceedings][5])

### UG-R-game-design-E07. Direct TD evidence concerns content inside existing rules

Sources: S18 through S21.

Liu's system generates roads, tower locations, and monster sequences in an in-house KRF-style simulator. Its tester uses flat Monte Carlo search, not standard MCTS. Kraner's genetic algorithms generate arrangements and waves; humans created the premise, assets, and game implementation. TowerMind evaluates agents rather than generating games. These papers do not demonstrate unrestricted tower or rule invention. ([ResearchGate][6])

### UG-R-game-design-E08. 3D scene quality, navigation, and strategy require different evidence

Sources: S22 through S28.

The RTS map work studies gameplay-related objectives, with visual evaluations rather than extensive competitive playtesting. The 3D level work checks movement or structural constraints. Infinigen and Holodeck target scene generation and embodied-AI environments. None of these results alone validates tactical terrain, build-location quality, or balanced TD encounters. Inspection ranged from relevant sections to abstracts, as recorded below. ([Springer][7])

### UG-R-game-design-E09. Executable mechanics can still be confusing or implementation-dependent

Sources: S30 through S33.

Mechanic Miner's teleportation example was useful to its solver but poorly understood in its pilot study. Mechanic Maker 2.0 changed generated rules by changing the evaluator, without a human study. Pixie produced mechanics across five case studies, but accessible fields and software structure constrained its discoveries; generated code could bypass object-pooling conventions. Execution success does not establish readability, production readiness, or player value. ([Possibility Space][8])

### UG-R-game-design-E10. Broader autonomy has been demonstrated in narrower settings

Sources: S36 through S39.

OMNI-EPIC reports autonomous learning-environment creation. DreamGarden evaluates semi-autonomous prototyping, including a ten-person usability study. AutoUE reports twenty benchmark tasks. Genie generates interactive environments through a learned world model. The inspected evidence does not demonstrate a complete novel 3D RTS or TD game with convincing strategic and originality evaluation by human players. That is a bounded review conclusion, not proof that no such work exists. ([arXiv][9])

### UG-R-game-design-E11. Evidence for novel solutions exists outside game creation

Sources: S15 and S16.

The Go study analyzes approximately 5.8 million decisions from 1950 to 2021 and associates post-AI changes with better and more novel human decisions. It is observational research within an existing game. FunSearch reports externally evaluated mathematical constructions and algorithmic improvements. These results support the possibility of useful novelty, but neither establishes autonomous TD design taste. ([arXiv][10])

### UG-R-game-design-E12. Passing many checks is weaker than completing a whole game correctly

Source: S44.

GameASG-Bench reports a highest mean runtime-check pass rate of 93.2% and a highest strict task success rate of 55.3% across 47 tasks. These are different aggregate measures and may concern different agent stacks. The paper measures requirement compliance, not strategic depth or originality. The manuscript was inspected; benchmark runs were not reproduced. ([arXiv][11])

### UG-R-game-design-E13. Citation, method adaptation, and replication must remain separate

Sources: S06, S18 through S20, S24, S26, S29 through S32, S36, S40 through S42.

The checked record supports an explicit Mechanic Miner environment recreation in Mechanic Maker 2.0, an extension of PCGRL to 3D generation, and OMNI-EPIC's extension of OMNI. Kraner's references include both Avery and Liu. Mortar cites related mechanic-generation work. These relationships do not establish identical implementations or independent replications of every earlier result. ([arXiv][12])

## 4. Decisions and unresolved choices

### UG-R-game-design-D01. Research scope expansion

Status: accepted as a user instruction, not an architecture decision.

Locator: S03, user, "especially for 3D Maps, Units, Mechanics, entire novel games" and "Especially in the RTS genre / towerdefense subgenre."

Recorded rationale: broaden the original research and inspect references and uses. The user did not approve the conclusions or any implementation by requesting this expansion.

### UG-R-game-design-D02. Headless TD experiment

Status: proposed by the assistant; approval not recorded.

Locator: S04, "My proposed first experiment would be a small headless TD simulator with an editable rule interface."

The proposal was to generate one tower or enemy mechanic together with a map-and-wave scenario. The assistant's rationale was to test an interaction before expanding to whole games. Whether this is necessary or proportionate for the current unit-output prototype remains unresolved.

### UG-R-game-design-D03. Separate generation, evaluation, and diversity management

Status: proposed by the assistant; approval not recorded.

Locators: S02, "Keep originality, quality, and fit separate"; S04, "Use a population of evaluators."

The proposal included explicit rules, comparisons with simpler mechanics, varied testers, precedent notes, and retention of different promising designs. The cited papers motivate parts of this proposal, not the complete architecture. No cost-benefit comparison was conducted.

### UG-R-game-design-D04. Public and private product boundaries

Status: owner-supplied current direction, not a historical decision recovered from this thread.

Locator: S05, "The intended public product is a standalone FOSS project with a genuinely usable default."

This direction is accepted as current context. Its implementation, licensing choices, operating costs, and private tuning mechanism remain untested or unspecified. No additional business rationale is recorded here.

### UG-R-game-design-D05. Towerright-first development and later extraction

Status: proposed by the owner in current context.

Locator: S05, "considering building the first working system for the Towerright pilot, then extracting a reusable standalone Unit-Generator."

No approved extraction boundary, dependency structure, or historical rationale is present in this thread.

### UG-R-game-design-D06. Small direct-prompt comparison

Status: new recommendation in this handover, awaiting approval.

Section 6 describes the comparison. Its purpose is to test the currently unverified product advantage, not to select a project-wide architecture in advance.

No explicit rejection, deferral, or supersession of a generation approach was recorded.

## 5. Experiments and technical artifacts

The thread contains the following inspectable material.

| Item                                                                                                            | Status and location                                                                                                                                                                                                            |
| --------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Literature reviews and source links                                                                             | Created and visible in S02 and S04. Their claims require the source qualifications in Sections 3 and 8.                                                                                                                        |
| Illustrative design prompts                                                                                     | Created as text, not executed. S02 proposed a unit that makes retreat timing important without a conventional slowing effect. S04 proposed a tower that makes route-change timing matter. Neither is a generated, tested unit. |
| Rule and evaluation checklists                                                                                  | Proposed in S02 and S04. Fields included triggers, targets, costs, effects, duration, limits, and counterplay. These are not a formal schema or implemented validator.                                                         |
| Headless simulator, mechanic-and-scenario generation, evaluator population, archive, and 3D construction stages | Proposed in S04. No code, configuration, dataset, or run output is visible.                                                                                                                                                    |
| External experimental results                                                                                   | Reported as tested by paper authors. Relevant manuscript sections and selected tables were inspected, but raw experiments were not reproduced.                                                                                 |
| Mardwerk technical artifacts                                                                                    | No code, model checkpoint, executable schema, benchmark dataset, calculation notebook, or output file is present.                                                                                                              |

The external quantities that mattered to the discussion are preserved in E02 through E12 and the source register. They are not project targets.

Reproduction would require the relevant paper version, source code and dependency versions, game rules, seed content, generator settings, model and prompt versions, evaluator policies and budgets, random seeds, candidate-selection procedure, and raw outcomes. Human comparisons also require their recruitment and evaluation protocols. Recording these improves traceability but does not guarantee identical future LLM outputs.

## 6. Implications for the pilot and later research

### What the thread actually recommended

S02 and S04 recommended moving beyond descriptive originality toward explicit mechanics, executable tests, comparison with simpler alternatives, and human evaluation. S04 proposed starting with a mechanic and its supporting scenario in a headless TD environment. It also proposed separate gameplay and 3D representations, varied evaluators, and an archive of different promising designs.

These were assistant recommendations. They were not approved requirements.

### Current recommendation, awaiting approval

Test one limited addition to direct generation: explicit constraint checking followed by a bounded revision. Do not first implement the entire S04 architecture.

A credible baseline would directly prompt the same model with the same character-reference snapshot, TD design brief, requested output format, and instruction about whether mechanics are wanted. It should be a competent prompt, not an intentionally weak one.

Compare that baseline with a workflow that generates the same kind of unit, checks the declared constraints, and revises identified failures. Use the same overall resource allowance where practical. The direct-prompt condition can spend its allowance on independent candidates rather than revision. Apply the same selection procedure to both conditions.

Record differences in source information, model versions, prompts, tool access, retrieval, private examples, calls, tokens, latency, monetary cost, and human intervention. Extra tools or information are treatment differences, not evidence that orchestration alone caused improvement.

The first comparison should distinguish:

* Practical usefulness, including designer editing effort and whether the result can be used.
* Technical correctness, including required fields, rule consistency, and executable behavior where execution is available.
* Diversity and character fidelity, assessed separately across outputs and against supplied reference facts.
* Game balance and scientific novelty, neither of which follows from the preceding measures.

A small simulator becomes informative when the output makes behavioral claims that cannot be checked from the design record. It should not become mandatory merely because S04 proposed one.

Full 3D maps, whole-game generation, reinforcement-learning training, a quality-diversity archive, autonomous balance certification, and model-weight fine-tuning should not become v1 requirements from this thread alone.

For standalone use, the same basic workflow should be tested without private Towerright components. Private tuning can remain an independent choice; the literature does not establish whether it should consist of prompts, data, settings, tools, or weights.

A possible later research contribution is a bounded evaluation of whether explicit checking and revision improves reference-conditioned TD unit outputs at a stated cost. Originality of that contribution would still require a focused prior-work review. It would not support claims about autonomous creation of balanced RTS games.

## 7. Open questions and conflicting evidence

| Question ID              | Question, significance, and resolution                                                                                                                                                                                                                                                                    |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **UG-R-game-design-Q01** | What is the minimum useful output: a design record, engine-ready configuration, or executable behavior? This product choice blocks a meaningful prototype acceptance test. Resolve it with a small set of concrete pilot examples.                                                                        |
| **UG-R-game-design-Q02** | Does Unit-Generator improve on competent direct prompting? This does not block construction, but it blocks the intended advantage claim. D06 would provide initial evidence.                                                                                                                              |
| **UG-R-game-design-Q03** | Which evaluations predict useful player decisions rather than evaluator exploitation? E02, E05, and E09 show different failures. This blocks automated balance or quality claims, not a designer-assisted prototype. Compare tester judgments with observed play when executable mechanics are available. |
| **UG-R-game-design-Q04** | How should character fidelity and mechanical originality be judged when they conflict? No cited experiment directly evaluates this task. It need not block a prototype, but source facts, intended adaptations, and invented features should be distinguishable during review.                            |
| **UG-R-game-design-Q05** | What belongs in the public workflow, and what is privately tuned? This is a product-boundary choice. It must be resolved before a standalone release, although a first pilot can test a provisional boundary. A run without private dependencies is the relevant check.                                   |
| **UG-R-game-design-Q06** | Does iteration improve individual outputs while reducing collection-level diversity? E06 makes this plausible but does not establish it for units. It does not block the first prototype. Compare repeated outputs before adding a diversity archive.                                                     |
| **UG-R-game-design-Q07** | When does additional autonomy justify its cost and failure modes? No thread experiment compares direct prompting, revision, search, or engine-agent workflows under controlled resources. Resolve the smallest difference first through D06 rather than treating autonomy as an objective itself.         |

The apparent conflicts do not reduce to "AI is creative" versus "AI is not creative." Higher novelty ratings can coexist with homogenization. Mechanically successful outputs can be disliked. A representation can support useful invention while also excluding important designs. These distinctions should survive consolidation.

## 8. Source register

### A. Sources already present in this thread

This register preserves all 39 identifiable research works cited or substantively named in the two reviews. OMNI, PCGRL, and the earlier StarCraft map paper originally appeared with incomplete bibliographic detail; their records were recovered during checking.

Repeated links and versions are grouped under one source ID. Tracking parameters have been removed from URLs. The malformed citation strings in the earlier responses cannot be recovered as original source markers; independently recovered publications are identified below.

"Relevant sections inspected" does not mean every appendix or reference was reviewed. Unless stated otherwise, associated code, data, model availability, and licensing were not checked. No listed repository was run.

#### Conversation sources

**UG-R-game-design-S01. Original user research request.**
Type: original message. Locator: user, "actual research paper regarding AI taste and getting novel solutions with AI." Establishes the original question. Message date and direct message URL are unavailable.

**UG-R-game-design-S02. First assistant literature review.**
Type: assistant synthesis. Locator: opening, "there is research directly relevant to this, including a paper specifically on generating RTS units"; later section, "How I would apply the research to designing units." Contains the first twelve research works, illustrative prompts, and unapproved recommendations.

**UG-R-game-design-S03. Expanded user research request.**
Type: original message. Locator: user, "check the references, check their usages" and "entire novel games." Establishes the expanded scope, not agreement with S02.

**UG-R-game-design-S04. Second assistant literature review.**
Type: assistant synthesis. Locator: section "A research-grounded architecture for your specific goal" and statement, "I did not run the repositories." Contains the expanded review and D02/D03. Its stated coverage date is not treated as a message timestamp.

**UG-R-game-design-S05. Current handover request and project context.**
Type: original message. Locator: headings "Current project context" and "Evidence rules." Establishes present intentions and evidence-handling instructions, not decisions made earlier.

#### Novelty, taste, and the first game-design sources

**UG-R-game-design-S06. Kynan Sorochan and Matthew Guzdial, 2022. *Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search*.**
Workshop paper, Experimental AI in Games; arXiv:2212.03387v1.
[https://arxiv.org/html/2212.03387v1](https://arxiv.org/html/2212.03387v1)

Bibliographic record verified through arXiv. Search, evaluation, results, and limitations were inspected. Appeared in S02's unit-design section and S04's RTS section; supports E02. It uses four supplied triggers, four effects, numerical attributes, and microRTS. Generated names and distinct visual assets were not the research contribution. Human enjoyment and general balance remain unverified. ([arXiv][13])

**UG-R-game-design-S07. Cameron Browne and Frédéric Maire, 2010. *Evolutionary Game Design*.**
Journal paper, IEEE Transactions on Computational Intelligence and AI in Games 2(1), 1 to 16. DOI: 10.1109/TCIAIG.2010.2041928.
[https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf](https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf)

Bibliographic record verified against the publisher record; relevant author-manuscript sections inspected. Appeared in S02 for Ludi, Yavalath, and rediscovery; supports E03. The author PDF retains draft journal placeholders, so those placeholders are not publication metadata. Its preference model concerns a restricted game family, not universal taste. ([IEEE Computer Society][14])

**UG-R-game-design-S08. Graham Todd, Alexander Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers, and Julian Togelius, 2024. *GAVEL: Generating Games Via Evolution and Language Models*.**
Conference paper, NeurIPS 2024; arXiv:2407.09388v2.
[https://arxiv.org/html/2407.09388v2](https://arxiv.org/html/2407.09388v2)

Record verified through arXiv and the authors' repository citation. Methods, results, and limitations inspected. Appeared in both reviews; supports E03. Generated examples combine modified line-completion rules with enclosure capture, but inactive components and limited human evaluation qualify the results. Code, data, and checkpoint links are advertised at [https://github.com/gdrtodd/gavel](https://github.com/gdrtodd/gavel) . ([arXiv][15])

**UG-R-game-design-S09. Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius, and Georgios N. Yannakakis, 2019. *Procedural Content Generation through Quality Diversity*.**
Survey/framework paper, IEEE Conference on Games; arXiv:1907.04053. Original cited DOI: 10.1109/CIG.2019.8848053.
[https://arxiv.org/abs/1907.04053](https://arxiv.org/abs/1907.04053)
[https://antoniosliapis.com/articles/pcgqd.php](https://antoniosliapis.com/articles/pcgqd.php)

The arXiv record and author-hosted article were checked; the supplied conference DOI was not independently resolved. Appeared in S02 to motivate retaining diverse good candidates. It supports a search strategy, not a validated TD diversity measure. Descriptor choice remains a design and evaluation problem. ([arXiv][16])

**UG-R-game-design-S10. Katherine M. Collins, Graham Todd, Cedegao E. Zhang, Adrian Weller, Julian Togelius, Junyi Chu, Lionel Wong, Thomas L. Griffiths, and Joshua B. Tenenbaum, 2025. *Generation and Evaluation in the Human Invention Process through the Lens of Game Design*.**
Non-archival CogSci 2025 presentation and preprint; arXiv:2508.10914.
[https://arxiv.org/html/2508.10914v1](https://arxiv.org/html/2508.10914v1)
[https://arxiv.org/abs/2508.10914](https://arxiv.org/abs/2508.10914)

Record verified; relevant v1 sections inspected. Appeared in S02 for simulation-informed invention; supports E04. Current metadata identifies a 2026 v2 update correcting Figure 2 and reference pointers. That update is not independent evidence, and its full content was not compared. The study's limited representation and presence-only data restrict interpretation. ([arXiv][3])

**UG-R-game-design-S11. Georgios N. Yannakakis and Julian Togelius, 2011. *Experience-Driven Procedural Content Generation*.**
Journal paper, IEEE Transactions on Affective Computing 2(3), 147 to 161. DOI: 10.1109/T-AFFC.2011.6.
[https://yannakakis.net/wp-content/uploads/2019/02/EDPCG.pdf](https://yannakakis.net/wp-content/uploads/2019/02/EDPCG.pdf)

Record verified; framework sections inspected in the author manuscript. Appeared in S02 for player-conditioned evaluation; supports E04. It organizes experience-driven generation and related studies. It does not establish an autonomous evaluator with universal design judgment. ([ACM Digital Library][17])

**UG-R-game-design-S12. Chenglei Si, Diyi Yang, and Tatsunori Hashimoto, 2025. *Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers*.**
Conference paper, ICLR 2025; preprint first posted in 2024, arXiv:2409.04109.
[https://arxiv.org/abs/2409.04109](https://arxiv.org/abs/2409.04109)
[https://proceedings.iclr.cc/paper_files/paper/2025/file/ea94957d81b1c1caf87ef5319fa6b467-Paper-Conference.pdf](https://proceedings.iclr.cc/paper_files/paper/2025/file/ea94957d81b1c1caf87ef5319fa6b467-Paper-Conference.pdf)

Record verified; abstract and relevant evaluation material inspected, not all 90 pages. Appeared in S02; supports E06. It evaluates research proposals, not successful completed projects. The paper links implementation and review scores at [https://github.com/NoviScl/AI-Researcher](https://github.com/NoviScl/AI-Researcher) . ([ICLR Proceedings][5])

**UG-R-game-design-S13. Anil R. Doshi and Oliver P. Hauser, 2024. *Generative AI enhances individual creativity but reduces the collective diversity of novel content*.**
Journal paper, Science Advances 10(28), eadn5290. DOI: 10.1126/sciadv.adn5290.
[https://www.science.org/doi/10.1126/sciadv.adn5290](https://www.science.org/doi/10.1126/sciadv.adn5290)

Publisher metadata verified, but the publisher text was inaccessible. The authors' preprint abstract was inspected at [https://arxiv.org/abs/2312.00506](https://arxiv.org/abs/2312.00506) . Its earlier title is *Generative artificial intelligence enhances creativity but reduces the diversity of novel content*. Appeared in S02; supports E06. The experiment concerns stories, not units or mechanics. ([Science][18])

**UG-R-game-design-S14. Emily Wenger and Yoed N. Kenett, 2026. *Large language models are homogeneously creative*.**
Journal paper, PNAS Nexus 5(3), pgag042. DOI: 10.1093/pnasnexus/pgag042.
[https://academic.oup.com/pnasnexus/article/5/3/pgag042/8529001](https://academic.oup.com/pnasnexus/article/5/3/pgag042/8529001)

Record verified; abstract and relevant publisher text inspected. Appeared as a companion source in S02; supports E06. The tested creativity tasks do not establish diversity behavior for TD designs. The paper links code and data at [https://github.com/arguslab-duke/llms_are_homogeneous](https://github.com/arguslab-duke/llms_are_homogeneous) . ([OUP Academic][19])

**UG-R-game-design-S15. Minkyu Shin, Jin Kim, Bas van Opheusden, and Thomas L. Griffiths, 2023. *Superhuman artificial intelligence can improve human decision-making by increasing novelty*.**
Journal paper, PNAS 120(12), e2214840120. DOI: 10.1073/pnas.2214840120.
[https://www.pnas.org/doi/10.1073/pnas.2214840120](https://www.pnas.org/doi/10.1073/pnas.2214840120)
[https://arxiv.org/abs/2303.07462](https://arxiv.org/abs/2303.07462)

Record verified through institutional publication metadata and arXiv. The preprint abstract was inspected; publisher full text was not recovered. Appeared in S02; supports E11. Its historical Go analysis is not a randomized game-design experiment or evidence of new rule invention. ([CityUHK Scholars][20])

**UG-R-game-design-S16. Bernardino Romera-Paredes, Mohammadamin Barekatain, Alexander Novikov, Matej Balog, M. Pawan Kumar, Emilien Dupont, Francisco J. R. Ruiz, Jordan S. Ellenberg, Pengming Wang, Omar Fawzi, Pushmeet Kohli, and Alhussein Fawzi, 2024. *Mathematical discoveries from program search with large language models*.**
Journal paper, Nature 625, 468 to 475; online publication in December 2023. DOI: 10.1038/s41586-023-06924-6.
[https://www.nature.com/articles/s41586-023-06924-6](https://www.nature.com/articles/s41586-023-06924-6)

Record and relevant full-text sections verified. Appeared in S02 as FunSearch; supports E11. External mathematical and algorithmic evaluation is stronger than descriptive novelty, but does not solve preference evaluation. The paper links code and discovered functions at [https://github.com/google-deepmind/funsearch](https://github.com/google-deepmind/funsearch) . ([Nature][21])

**UG-R-game-design-S17. Margaret A. Boden, 1998. *Creativity and artificial intelligence*.**
Theoretical journal paper, Artificial Intelligence 103(1 to 2), 347 to 356. Original cited DOI: 10.1016/S0004-3702(98)00055-1.
[https://www.sciencedirect.com/science/article/pii/S0004370298000551](https://www.sciencedirect.com/science/article/pii/S0004370298000551)

Bibliographic verification is partial: publisher search metadata was recovered, but primary full text was not retrieved. The detailed citation is retained from S02. It appeared for combinational, exploratory, and transformational creativity. That account remains attributed to the original review, not newly content-verified here, and is not empirical evidence about modern LLMs. ([ScienceDirect][22])

#### Direct RTS and tower-defense research

**UG-R-game-design-S18. Simon Liu, Li Chaoran, Li Yue, Ma Heng, Hou Xiao, Shen Yiming, Wang Licong, Chen Ze, Guo Xianghao, Lu Hengtong, Du Yu, and Tang Qinting, 2019. *Automatic generation of tower defense levels using PCG*.**
Conference paper, Foundations of Digital Games. DOI: 10.1145/3337722.3337723. Author names are retained in the order and form printed in the manuscript.
[https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG](https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG)

Record verified; author manuscript methods, results, and references inspected. Appeared in S04; supports E07/E13. The testbed is an in-house KRF-style simulator, not demonstrated integration into the commercial game's code. The paper presents three generated levels and uses flat Monte Carlo testing. ([ResearchGate][6])

**UG-R-game-design-S19. Vid Kraner, Iztok Fister Jr., and Lucija Brezočnik, 2021. *Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms*.**
Conference chapter, New Technologies, Development and Application IV, LNNS 233. DOI: 10.1007/978-3-030-75275-0_54.
[https://link.springer.com/chapter/10.1007/978-3-030-75275-0_54](https://link.springer.com/chapter/10.1007/978-3-030-75275-0_54)
[https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms](https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms)

Record verified; relevant manuscript sections inspected. Appeared in S04; supports E07/E13. Checking also recovered feedback from five players who experienced the unoptimized version first. That small, ordered comparison does not establish general preference. Unity and Blender implementation does not mean the algorithm invented the game's premise or assets. ([Springer][23])

**UG-R-game-design-S20. Phillipa Avery, Julian Togelius, Elvis Alistar, and Robert Pieter van Leeuwen, 2011. *Computational Intelligence and Tower Defence Games*.**
Conference paper, IEEE Congress on Evolutionary Computation, 1084 to 1091. DOI: 10.1109/CEC.2011.5949738.
[https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games](https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games)

Record verified; abstract and relevant author-manuscript material inspected. Appeared in S04 as the earlier TD foundation; supports E07/E13. It classifies components and presents an experience-driven prototype, rather than unrestricted whole-game invention. A separate author PDF URL failed during checking; the original linked manuscript remained accessible. ([ResearchGate][24])

**UG-R-game-design-S21. Dawei Wang, Chengming Zhou, Di Zhao, Xinyuan Liu, Marci Chi Ma, Gary Ushaw, and Richard Davison, 2026. *TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents*.**
Originally cited as AAAI 2026, 40(31), DOI: 10.1609/aaai.v40i31.39818; arXiv:2601.05899.
[https://arxiv.org/html/2601.05899v1](https://arxiv.org/html/2601.05899v1)

ArXiv record verified; relevant v1 sections inspected. The original AAAI citation is retained, without independently resolving its DOI here. Appeared in S04 as evaluation infrastructure; supports E07. It is an agent benchmark, not a generator. The paper links [https://github.com/tb6147877/TowerMind](https://github.com/tb6147877/TowerMind) . The later arXiv version is not treated as another study. ([arXiv][25])

#### Maps and 3D environments

**UG-R-game-design-S22. Julian Togelius, Mike Preuss, Nicola Beume, Simon Wessing, Johan Hagelbäck, Georgios N. Yannakakis, and Corrado Grappiolo, 2013. *Controllable procedural map generation via multiobjective evolution*.**
Journal paper, Genetic Programming and Evolvable Machines 14, 245 to 277. DOI: 10.1007/s10710-012-9174-5.
[https://link.springer.com/article/10.1007/s10710-012-9174-5](https://link.springer.com/article/10.1007/s10710-012-9174-5)

Record verified; publisher abstract inspected, not subscription full text. Appeared in S04 for RTS map objectives; supports E08. It includes heightmap and StarCraft representations. The reported visual evaluation should not be restated as extensive competitive playtesting. ([Springer][7])

**UG-R-game-design-S23. Julian Togelius, Mike Preuss, Nicola Beume, Simon Wessing, Johan Hagelbäck, and Georgios N. Yannakakis, 2010. *Multiobjective exploration of the StarCraft map space*.**
Conference paper, IEEE Computational Intelligence and Games, 265 to 272. DOI: 10.1109/ITW.2010.5593346.
[https://ieeexplore.ieee.org/document/5593346/](https://ieeexplore.ieee.org/document/5593346/)
[https://julian.togelius.com/Togelius2010Multiobjective.pdf](https://julian.togelius.com/Togelius2010Multiobjective.pdf)

Named, but not fully referenced, in S04. Bibliographic details were recovered from publisher and institutional records. Metadata and author-manuscript search text were inspected, not the complete paper. It relates to E08; it is an earlier map-generation work, not independent evidence about unit generation. ([IEEE Xplore][26])

**UG-R-game-design-S24. Zehua Jiang, Sam Earle, Michael Cerny Green, and Julian Togelius, 2022. *Learning Controllable 3D Level Generators*.**
Conference paper, Foundations of Digital Games. DOI: 10.1145/3555858.3563273; arXiv:2206.13623v3.
[https://arxiv.org/html/2206.13623v3](https://arxiv.org/html/2206.13623v3)

Record verified; relevant sections inspected. Appeared in S04; supports E08/E13. It extends learned level editing into constrained voxel environments with 3D movement considerations. The demonstrations do not establish production-scale RTS map quality. ([ACM Digital Library][27])

**UG-R-game-design-S25. Sam Earle, Zehua Jiang, and Julian Togelius, 2024. *PCGRL+: Scaling, Control and Generalization in Reinforcement Learning Level Generators*.**
Conference paper, IEEE Conference on Games; arXiv:2408.12525.
[https://arxiv.org/abs/2408.12525](https://arxiv.org/abs/2408.12525)

Record verified, including the arXiv venue comment; abstract inspected. Appeared in S04 as a methodological continuation. It concerns scaling and control of level generators, not complete-game design or a demonstrated TD unit workflow. Related to E08. ([arXiv][28])

**UG-R-game-design-S26. Kaijie Xu and Clark Verbrugge, 2025. *A Database-Driven Framework for 3D Level Generation with LLMs*.**
Conference paper, AIIDE 21(1), 367 to 376. DOI: 10.1609/aiide.v21i1.36840; arXiv:2508.18533v1.
[https://ojs.aaai.org/index.php/AIIDE/article/view/36840](https://ojs.aaai.org/index.php/AIIDE/article/view/36840)
[https://arxiv.org/html/2508.18533v1](https://arxiv.org/html/2508.18533v1)

Record verified; methods and limitations inspected. Appeared in S04; supports E08/E13. It uses reusable databases, constrained assembly, and repair. The authors acknowledge curation, remaining unplayable outputs, and no human-player evaluation. Its key-and-lock-style progression is not an RTS strategy evaluation. ([AAAI Publications][29])

**UG-R-game-design-S27. Alexander Raistrick, Lahav Lipson, Zeyu Ma, Lingjie Mei, Mingzhe Wang, Yiming Zuo, Karhan Kayan, Hongyu Wen, Beining Han, Yihan Wang, Alejandro Newell, Hei Law, Ankit Goyal, Kaiyu Yang, and Jia Deng, 2023. *Infinite Photorealistic Worlds using Procedural Generation*.**
Conference paper, CVPR; arXiv:2306.09310v2.
[https://arxiv.org/abs/2306.09310](https://arxiv.org/abs/2306.09310)

Record verified; abstract inspected. Appeared in S04 as Infinigen; supports E08. It generates procedural natural scenes for uses including synthetic visual data. Strategic map quality was not the evaluated target. The authors link code and generated data through [https://infinigen.org](https://infinigen.org) . ([arXiv][30])

**UG-R-game-design-S28. Yue Yang, Fan-Yun Sun, Luca Weihs, Eli VanderBilt, Alvaro Herrasti, Winson Han, Jiajun Wu, Nick Haber, Ranjay Krishna, Lingjie Liu, Chris Callison-Burch, Mark Yatskar, Aniruddha Kembhavi, and Christopher Clark, 2024. *Holodeck: Language Guided Generation of 3D Embodied AI Environments*.**
Conference paper, CVPR; arXiv:2312.09067v2, first posted in 2023.
[https://arxiv.org/abs/2312.09067](https://arxiv.org/abs/2312.09067)

Record verified; abstract inspected. Appeared in S04; supports E08. The system uses language-guided scene planning, retrieved assets, and spatial constraints. Scene preferences and embodied navigation do not validate RTS terrain or TD balance. ([arXiv][31])

#### Mechanics, code, and deployed content

**UG-R-game-design-S29. Muhammad U. Nasir, Yuchen Li, Steven James, and Julian Togelius, 2026. *Mortar: Evolving Mechanics for Automatic Game Design*.**
Conference paper, GECCO. DOI: 10.1145/3795095.3805100.
[https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf](https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf)

Record verified through the author publication page and manuscript. Methods, evaluation, limitations, references, and the human-comparison table were inspected. Appeared in S04; supports E05/E13. Correction to the broad earlier framing: assembling games for mechanic evaluation is not the same as targeting complete-game production. ([RAIL Lab][32])

**UG-R-game-design-S30. Michael Cook, Simon Colton, Azalea Raad, and Jeremy Gow, 2013. *Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design*.**
Conference chapter, EvoApplications. Original cited DOI: 10.1007/978-3-642-37192-9_29.
[https://www.possibilityspace.org/papers/evo13.pdf](https://www.possibilityspace.org/papers/evo13.pdf)

Author manuscript verified and relevant sections inspected; the DOI was not independently resolved. Appeared in S04; supports E09/E13. Its reflection-based search produced gravity, bouncing, and teleportation interactions under restricted operations. A newly noticed reporting ambiguity remains: the pilot paragraph introduces seven participants but later refers to five. No corrected sample size is inferred. ([Possibility Space][8])

**UG-R-game-design-S31. Johor Jara Gonzalez, Seth Cooper, and Matthew Guzdial, 2023. *Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules*.**
Conference paper, AIIDE 19(1), 266 to 275. DOI: 10.1609/aiide.v19i1.27522; arXiv:2309.09476v3.
[https://arxiv.org/html/2309.09476v3](https://arxiv.org/html/2309.09476v3)

Record verified; methods, comparisons, and limitations inspected. Appeared in S04; supports E09/E13. It recreates the earlier environment in Unity and changes evaluation, rather than replicating every original condition. No human study validates the proposed usability advantage. The manuscript links [https://github.com/Harcurio/MechanicMiner](https://github.com/Harcurio/MechanicMiner) . ([arXiv][12])

**UG-R-game-design-S32. Michael Cook, 2025. *Pixie: Code-Level Mechanic Generation for Game Designers*.**
Conference paper, AIIDE 21(1), 206 to 215. DOI: 10.1609/aiide.v21i1.36824.
[https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962](https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962)

Record verified; case studies and limitations inspected. Appeared in S04; supports E09/E13. It uses annotations, code generation, compilation, and search rather than an LLM. Its open-source game case studies are implementation demonstrations, not proof of production-ready output. The object-pooling limitation from S04 is supported by the paper. ([AAAI Publications][33])

**UG-R-game-design-S33. Alexander Zook and Mark O. Riedl, 2014. *Automatic Game Design via Mechanic Generation*.**
Conference paper, AAAI. DOI: 10.1609/aaai.v28i1.8788; later arXiv deposit 1908.01420v1.
[https://ojs.aaai.org/index.php/AAAI/article/view/8788](https://ojs.aaai.org/index.php/AAAI/article/view/8788)
[https://arxiv.org/html/1908.01420v1](https://arxiv.org/html/1908.01420v1)

Record verified; relevant manuscript sections inspected. Appeared in S04 as formal synthesis; relates to E09. It uses planning-like representations and constraints in supplied domains. Its formal vocabulary limits the mechanics it can express. The 2019 deposit is not a separate study or the original publication year. ([AAAI Publications][34])

**UG-R-game-design-S34. Erin J. Hastings, Ratan K. Guha, and Kenneth O. Stanley, 2009. *Automatic Content Generation in the Galactic Arms Race Video Game*.**
Journal paper, IEEE Transactions on Computational Intelligence and AI in Games 1(4), 245 to 263. DOI: 10.1109/TCIAIG.2009.2038365.
[https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game](https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game)

Record verified against publisher metadata; relevant author-manuscript material inspected. Appeared in S04 for deployed player-guided generation. The paper reports over 1,000 registered players and evolved particle-system weapons. It does not represent a controlled 1,000-person preference experiment or invention of the entire game. Player usage is an indirect preference signal. ([IEEE Xplore][35])

#### Complete rule sets, environment autonomy, and engine workflows

**UG-R-game-design-S35. Matthew Guzdial and Mark Riedl, 2018. *Automated Game Design via Conceptual Expansion*.**
Conference paper, AIIDE 14(1), 31 to 37. DOI: 10.1609/aiide.v14i1.13022.
[https://ojs.aaai.org/index.php/AIIDE/article/view/13022](https://ojs.aaai.org/index.php/AIIDE/article/view/13022)

Record verified; publisher abstract inspected. Appeared in S04 for learned recombination; supports E03. The abstract explicitly identifies recreation of existing games as its evaluation. A full-text download attempt failed. The reconstruction result should not be upgraded into demonstrated historical originality. ([AAAI Publications][36])

**UG-R-game-design-S36. Maxence Faldor, Jenny Zhang, Antoine Cully, and Jeff Clune, 2025. *OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code*.**
Conference paper, ICLR 2025; preprint first posted in 2024, arXiv:2405.15568.
[https://arxiv.org/abs/2405.15568](https://arxiv.org/abs/2405.15568)
[https://omni-epic.vercel.app/](https://omni-epic.vercel.app/)

Record verified through arXiv, conference search metadata, and the author project page. Abstract and project description inspected; full-text retrieval failed during this handover. Appeared in S04; supports E10/E13. Its task-learning objective is not equivalent to human-game enjoyment. Associated repository: [https://github.com/maxencefaldor/omni-epic](https://github.com/maxencefaldor/omni-epic) . ([OpenReview][37])

**UG-R-game-design-S37. Sam Earle, Samyak Parajuli, and Andrzej Banburski-Fahey, 2025. *DreamGarden: A Designer Assistant for Growing Games from a Single Prompt*.**
Conference paper, CHI. DOI: 10.1145/3706598.3714233; arXiv:2410.01791v2.
[https://arxiv.org/html/2410.01791v2](https://arxiv.org/html/2410.01791v2)

Record verified; workflow, usability study, and limitations inspected. Appeared in S04; supports E10. It retains human editing and redirection. The study's interaction window was shorter than typical completion time for generation, limiting what its usability results establish about finished games. ([arXiv][38])

**UG-R-game-design-S38. Lei Yin, Wentao Cheng, Zhida Qin, Tianyu Huang, Yidong Li, and Gangyi Ding, 2026. *AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems*.**
Preprint, arXiv:2603.07106. No peer-reviewed venue was verified.
[https://arxiv.org/html/2603.07106v1](https://arxiv.org/html/2603.07106v1)

Record verified; relevant v1 methods and evaluation sections inspected. Appeared in S04; supports E10. PlayGen-20 contains twenty tasks, and evaluation uses substantial automated and model-based assessment. This does not establish strategically novel RTS play. Later arXiv metadata was noted, not counted as another result. ([arXiv][39])

**UG-R-game-design-S39. Jake Bruce, Michael Dennis, Ashley Edwards, Jack Parker-Holder, Yuge Shi, Edward Hughes, Matthew Lai, Aditi Mavalankar, Richie Steigerwald, Chris Apps, Yusuf Aytar, Sarah Bechtle, Feryal Behbahani, Stephanie Chan, Nicolas Heess, Lucy Gonzalez, Simon Osindero, Sherjil Ozair, Scott Reed, Jingwei Zhang, Konrad Zolna, Jeff Clune, Nando de Freitas, Satinder Singh, and Tim Rocktäschel, 2024. *Genie: Generative Interactive Environments*.**
Conference paper, ICML, PMLR 235, 4603 to 4623; arXiv:2402.15391v1. Author forms above follow arXiv; the proceedings expand some names.
[https://arxiv.org/abs/2402.15391](https://arxiv.org/abs/2402.15391)
[https://proceedings.mlr.press/v235/bruce24a.html](https://proceedings.mlr.press/v235/bruce24a.html)

Record verified; abstract inspected. Appeared in S04 as a separate world-model category; supports E10. It does not provide conventional inspectable rules and engine code as its primary output. ([arXiv][40])

#### Referenced foundations and evaluation sources

**UG-R-game-design-S40. Alexander Zook, Brent Harrison, and Mark O. Riedl, 2015. *Monte-Carlo Tree Search for Simulation-based Strategy Analysis*.**
Originally cited as an FDG 2015 paper; later arXiv deposit 1908.01423, dated 2019.
Original URL: [https://www.fdg2015.org/papers/fdg2015_paper_08.pdf](https://www.fdg2015.org/papers/fdg2015_paper_08.pdf)
Recovered author manuscript: [https://faculty.cc.gatech.edu/~riedl/pubs/zook-fdg15.pdf](https://faculty.cc.gatech.edu/~riedl/pubs/zook-fdg15.pdf)
[https://arxiv.org/abs/1908.01423](https://arxiv.org/abs/1908.01423)

Title, authors, manuscript, and arXiv record verified. The original FDG link failed; the 2015 conference attribution is retained from S04. Abstract and introductory material inspected. S04 used it to trace the RTS evaluator's origin; relates to E13. The 2019 repository date must not replace the earlier publication attribution. ([arXiv][41])

**UG-R-game-design-S41. Jenny Zhang, Joel Lehman, Kenneth Stanley, and Jeff Clune, 2024. *OMNI: Open-endedness via Models of human Notions of Interestingness*.**
Conference paper, ICLR 2024; preprint first posted in 2023, arXiv:2306.01711v3.
[https://arxiv.org/abs/2306.01711](https://arxiv.org/abs/2306.01711)

Originally named only as OMNI in S04's lineage discussion. Record recovered and verified; abstract inspected, with conference status checked through proceedings metadata. It concerns selection of interesting learning tasks, not verified game-design taste. Supports E13 and the evaluation alternatives in Section 2. Repository: [https://github.com/jennyzzt/omni](https://github.com/jennyzzt/omni) . ([arXiv][42])

**UG-R-game-design-S42. Ahmed Khalifa, Philip Bontrager, Sam Earle, and Julian Togelius, 2020. *PCGRL: Procedural Content Generation via Reinforcement Learning*.**
Conference paper, AIIDE 2020; arXiv:2001.09212v3.
[https://arxiv.org/abs/2001.09212](https://arxiv.org/abs/2001.09212)

Originally named as the foundational method in S04. Record recovered and verified; abstract inspected. Supports E08/E13. Reinforcement learning trains the level-editing policy here. That is different from training agents to playtest a generated mechanic. No TD unit-generation effectiveness is established. ([arXiv][43])

**UG-R-game-design-S43. Stuart Anderson and Ruth Falconer, 2025. *Emergence of player tactics by expert-guided machine learning: an industry tower defence case study*.**
Journal paper, Entertainment Computing 54, 100963. DOI: 10.1016/j.entcom.2025.100963.
[https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/](https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/)

Record verified through the authors' institutional repository; abstract inspected. Appeared in S04 for behavioral evaluation. It analyzes player tactics rather than generating units. It may inform what to observe in playtests, but does not validate a novelty metric or a generation pipeline. ([Abertay University][44])

**UG-R-game-design-S44. Xiuhui Zhang, Yi Chen, Shusheng Xu, Fan Li, Huan Wang, Tongkai Yang, and Binhang Yuan, 2026. *GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development*.**
Preprint, arXiv:2609.21293v1, submitted September 18, 2026.
[https://arxiv.org/html/2609.21293v1](https://arxiv.org/html/2609.21293v1)

Record verified; benchmark description and reported results inspected. Appeared in S04; supports E12. Its 47 browser-game tasks measure specified runtime behavior, not design taste, originality, or TD balance. The paper links [https://github.com/areal-project/GameASG-Bench](https://github.com/areal-project/GameASG-Bench) . Neither repository execution nor independent replication was performed. ([arXiv][45])

### B. Newly identified research leads

**UG-R-game-design-S45. Newly found during handover. Michael Cerny Green, Ahmed Khalifa, Gabriella A. B. Barros, Tiago Machado, and Julian Togelius. *Automatic Critical Mechanic Discovery Using Playtraces in Video Games*.**
Preprint first posted in 2019; inspected record is arXiv:1909.03094v3, revised in 2020.
[https://arxiv.org/abs/1909.03094](https://arxiv.org/abs/1909.03094)

Found through Mortar's reference 8, which uses the earlier title *Automatic Critical Mechanic Discovery in Video Games*. The persistent identifier connects the title variants. Bibliographic record and abstract were inspected; no full-text assessment or replication was performed.

The specific relevance is determining which mechanics matter in recorded play, a gap behind Q03 and the proposed rule-removal comparison. The abstract reports comparison with human-identified mechanics and use in MCTS across four GVG-AI games. Mortar's reference establishes a citation relationship, not replication, code reuse, or application by Unit-Generator. ([RAIL Lab][4])

## 9. Consolidation handoff

The findings that must survive are the limited but direct unit-generation evidence in E02, the distinction between generated and historically original designs in E03, and the evaluator failures in E05 and E09. E08 and E10 prevent scene generation, executable prototypes, and validated strategic game design from being treated as the same achievement.

The decision status must also survive: D02 and D03 remain assistant proposals. D04 is current owner-supplied direction. D05 remains a proposed development sequence. No generation architecture won in this thread.

This thread establishes a relevant literature base, identifiable mechanisms, and testable uncertainties. It does not establish Unit-Generator's practical advantage, character fidelity, balance, diversity, scientific originality, or standalone implementation.

The next informative step is to resolve Q01 and consider D06 against Q02. That comparison can determine whether additional checking and revision earns its cost before the broader research architecture becomes implementation work.

[1]: https://arxiv.org/html/2212.03387v1 "Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search"
[2]: https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf "https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf"
[3]: https://arxiv.org/html/2508.10914v1 "Generation and Evaluation in the Human Invention Process through the Lens of Game Design"
[4]: https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf "https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf"
[5]: https://proceedings.iclr.cc/paper_files/paper/2025/file/ea94957d81b1c1caf87ef5319fa6b467-Paper-Conference.pdf "https://proceedings.iclr.cc/paper_files/paper/2025/file/ea94957d81b1c1caf87ef5319fa6b467-Paper-Conference.pdf"
[6]: https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG "https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG"
[7]: https://link.springer.com/article/10.1007/s10710-012-9174-5 "https://link.springer.com/article/10.1007/s10710-012-9174-5"
[8]: https://www.possibilityspace.org/papers/evo13.pdf "https://www.possibilityspace.org/papers/evo13.pdf"
[9]: https://arxiv.org/abs/2405.15568 "[2405.15568] OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code"
[10]: https://arxiv.org/abs/2303.07462 "https://arxiv.org/abs/2303.07462"
[11]: https://arxiv.org/html/2609.21293v1 "https://arxiv.org/html/2609.21293v1"
[12]: https://arxiv.org/html/2309.09476v3 "https://arxiv.org/html/2309.09476v3"
[13]: https://arxiv.org/abs/2212.03387 "[2212.03387] Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search"
[14]: https://www.computer.org/csdl/journal/ci/2010/01/05404867/13rRUILLkpS "https://www.computer.org/csdl/journal/ci/2010/01/05404867/13rRUILLkpS"
[15]: https://arxiv.org/html/2407.09388v2 "https://arxiv.org/html/2407.09388v2"
[16]: https://arxiv.org/abs/1907.04053 "[1907.04053] Procedural Content Generation through Quality Diversity"
[17]: https://dl.acm.org/doi/10.1109/T-AFFC.2011.6?utm_source=chatgpt.com "Experience-Driven Procedural Content Generation"
[18]: https://www.science.org/doi/10.1126/sciadv.adn5290 "https://www.science.org/doi/10.1126/sciadv.adn5290"
[19]: https://academic.oup.com/pnasnexus/article/5/3/pgag042/8529001 "https://academic.oup.com/pnasnexus/article/5/3/pgag042/8529001"
[20]: https://scholars.cityu.edu.hk/en/publications/superhuman-artificial-intelligence-can-improve-human-decision-mak/ "https://scholars.cityu.edu.hk/en/publications/superhuman-artificial-intelligence-can-improve-human-decision-mak/"
[21]: https://www.nature.com/articles/s41586-023-06924-6 "Mathematical discoveries from program search with large language models | Nature"
[22]: https://www.sciencedirect.com/science/article/pii/S0004370298000551 "https://www.sciencedirect.com/science/article/pii/S0004370298000551"
[23]: https://link.springer.com/chapter/10.1007/978-3-030-75275-0_54 "https://link.springer.com/chapter/10.1007/978-3-030-75275-0_54"
[24]: https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games "https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games"
[25]: https://arxiv.org/abs/2601.05899 "https://arxiv.org/abs/2601.05899"
[26]: https://ieeexplore.ieee.org/document/5593346/ "https://ieeexplore.ieee.org/document/5593346/"
[27]: https://dl.acm.org/doi/10.1145/3555858.3563273?utm_source=chatgpt.com "Learning Controllable 3D Level Generators"
[28]: https://arxiv.org/abs/2408.12525 "https://arxiv.org/abs/2408.12525"
[29]: https://ojs.aaai.org/index.php/AIIDE/article/view/36840 "A Database-Driven Framework for 3D Level Generation with LLMs | Proceedings of the AAAI Conference on Artificial Intelligence and Interactive Digital Entertainment"
[30]: https://arxiv.org/abs/2306.09310 "[2306.09310] Infinite Photorealistic Worlds using Procedural Generation"
[31]: https://arxiv.org/abs/2312.09067 "[2312.09067] Holodeck: Language Guided Generation of 3D Embodied AI Environments"
[32]: https://www.raillab.org/publication/nasir-2026-mortar/ "Mortar: Evolving Mechanics for Automatic Game Design | RAIL Lab"
[33]: https://ojs.aaai.org/index.php/AIIDE/article/view/36824 "https://ojs.aaai.org/index.php/AIIDE/article/view/36824"
[34]: https://ojs.aaai.org/index.php/AAAI/article/view/8788 "https://ojs.aaai.org/index.php/AAAI/article/view/8788"
[35]: https://ieeexplore.ieee.org/document/5352259/ "https://ieeexplore.ieee.org/document/5352259/"
[36]: https://ojs.aaai.org/index.php/AIIDE/article/view/13022 "Automated Game Design via Conceptual Expansion | Proceedings of the AAAI Conference on Artificial Intelligence and Interactive Digital Entertainment"
[37]: https://openreview.net/forum?id=Y1XkzMJpPd "https://openreview.net/forum?id=Y1XkzMJpPd"
[38]: https://arxiv.org/abs/2410.01791 "https://arxiv.org/abs/2410.01791"
[39]: https://arxiv.org/abs/2603.07106 "https://arxiv.org/abs/2603.07106"
[40]: https://arxiv.org/abs/2402.15391 "[2402.15391] Genie: Generative Interactive Environments"
[41]: https://arxiv.org/abs/1908.01423 "https://arxiv.org/abs/1908.01423"
[42]: https://arxiv.org/abs/2306.01711 "https://arxiv.org/abs/2306.01711"
[43]: https://arxiv.org/abs/2001.09212 "https://arxiv.org/abs/2001.09212"
[44]: https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/ "https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/"
[45]: https://arxiv.org/abs/2609.21293 "[2609.21293] GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development"
