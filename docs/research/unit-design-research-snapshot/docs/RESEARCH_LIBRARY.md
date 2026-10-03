# Research library

This is the authoritative inventory of retained scholarly references. The initial search was on 2026-09-27, with targeted reading through 2026-09-28. Coverage and reading depth remain incomplete. The [search log](SEARCH_LOG.md) records those limits and outstanding work. This is not an exhaustive systematic review. Cite the stable paper IDs below in other documents.

How to read the entries:

- Each paper has one primary group. Tags show connections across groups.
- “Direct” studies game content, design, mechanics, or relevant game evaluation. “Adjacent” studies a different task with an explicit connection.
- Claims follow the stated inspection depth. An abstract supports only what it states; reading through an index is identified.
- “Scope inference” is our assessment of a study's scope, not necessarily a limitation claimed by its authors.
- Publication status is recorded only where verified. An arXiv record alone does not establish peer review. Inclusion does not select a project method.

## Browse by topic

- [PCG and broad maps](#procedural-content-generation-and-broad-maps)
- [Tower Defense and RTS evidence](#tower-defense-rts-units-and-game-specific-evidence)
- [Automated design and mechanics](#automated-design-mechanics-and-language-guided-generation)
- [Rule representations and corpora](#rule-representations-and-reusable-corpora)
- [Character understanding and fidelity](#character-understanding-and-fidelity)
- [Grounded knowledge and constraints](#grounded-knowledge-and-acquiring-constraints)
- [Evaluation and playtesting](#evaluation-factuality-and-automated-playtesting)
- [Decision models and RLCD](#decision-models-and-the-rlcd-acronym)
- [Research methodology](#research-methodology)
- [Non-paper resources](#non-paper-resources)
- [Paper index by title](#paper-index)

## Procedural content generation and broad maps

### sbpcg11

**SBPCG11. Search-Based Procedural Content Generation: A Taxonomy and Survey**

Julian Togelius, Georgios N. Yannakakis, Kenneth O. Stanley, and Cameron Browne. 2011. *IEEE Transactions on Computational Intelligence and AI in Games*. DOI: [10.1109/TCIAIG.2011.2148116](https://doi.org/10.1109/TCIAIG.2011.2148116). Direct survey. Tags: search, representation, evaluation.

- **Problem and contribution:** organizes evolutionary and other metaheuristic generation of digital and nondigital game content. Its taxonomy distinguishes the generated content, its representation, and the quality or fitness evaluation.
- **Method and evaluation:** literature synthesis, not a new generator benchmark.
- **Relevance:** establishes vocabulary for comparing approaches without assuming learning is required.
- **Limitation:** its historical coverage cannot establish the present state of the field; its open problems need checking against later work.
- **Inspection:** bibliographic metadata in Crossref and abstract reproduced by OpenAlex; original full text not inspected.

### pcgml18

**PCGML18. Procedural Content Generation via Machine Learning (PCGML)**

Adam Summerville, Sam Snodgrass, Matthew Guzdial, Christoffer Holmgård, Amy K. Hoover, Aaron Isaksen, Andy Nealen, and Julian Togelius. 2018. *IEEE Transactions on Games*. DOI: [10.1109/TG.2018.2846639](https://doi.org/10.1109/TG.2018.2846639). [Manuscript, arXiv:1702.00539](https://arxiv.org/abs/1702.00539), initially posted 2017. Direct survey. Tags: learning, datasets, representation, cross-game adaptation.

- **Problem and contribution:** defines PCGML as generation using models trained on existing content and compares approaches across data representations and learning methods. It covers functional content, including levels, interactive fiction, and cards.
- **Method and evaluation:** synthesis of existing systems. In the inspected discussion, the authors caution that methods are difficult to compare because there is no agreed test for generator quality, and training and generation costs are inconsistently reported. They identify small datasets, limited data, multilayer learning, and style transfer as open problems. The conclusion notes the then-heavy concentration on two-dimensional levels and the need to investigate rules, items, and characters.
- **Relevance:** a foundation for interpreting learning claims and corpus requirements. These are historical observations, not evidence that the same gaps remain unaddressed in 2026.
- **Inspection:** abstract and selected discussion/conclusion sections of the primary arXiv HTML; publication metadata checked in Crossref.

### pcgqd19

**PCGQD19. Procedural Content Generation through Quality Diversity**

Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius, and Georgios N. Yannakakis. 2019. *IEEE Conference on Games*. DOI: [10.1109/CIG.2019.8848053](https://doi.org/10.1109/CIG.2019.8848053). [arXiv:1907.04053](https://arxiv.org/abs/1907.04053). Direct survey and position paper. Tags: search, diversity, mixed initiative.

- **Problem and contribution:** explains algorithms that search for many high-quality solutions spread across a space defined by behavior metrics. It reviews applications to game content and play and discusses future challenges.
- **Method and evaluation:** conceptual comparison and review of existing applications, rather than evidence that one metric space captures all useful diversity.
- **Relevance:** clarifies the difference between optimizing quality and characterizing varied possibilities.
- **Scope inference:** the chosen behavior metrics determine which differences the archive can express; this does not answer how to discover unfamiliar mechanics.
- **Inspection:** primary arXiv abstract and Crossref metadata.

### pcgrl20

**PCGRL20. PCGRL: Procedural Content Generation via Reinforcement Learning**

Ahmed Khalifa, Philip Bontrager, Sam Earle, and Julian Togelius. 2020. *AAAI Conference on Artificial Intelligence and Interactive Digital Entertainment*, 16(1). DOI: [10.1609/aiide.v16i1.7416](https://doi.org/10.1609/aiide.v16i1.7416). [arXiv:2001.09212](https://arxiv.org/abs/2001.09212). Direct method paper. Tags: reinforcement learning, representation, reward design.

- **Problem and contribution:** treats iterative level editing as a sequential decision problem and learns a generator policy.
- **Method:** PPO with narrow, turtle, and wide editing representations, tested on binary maps, Zelda, and Sokoban. The inspected experiments train three models per representation/problem for 100 million frames. Evaluation uses predefined connectivity, object-count, path-length, and solver criteria. The conclusion reports many playable levels but difficulty generating hard Zelda and Sokoban levels. The discussion explicitly says designing an appropriate reward remains a problem, like defining fitness in search-based PCG.
- **Relevance:** shows what learning a generator can mean when examples are scarce. It presupposes a content domain, editing actions, and evaluative criteria. It supplies no evidence that these can remain unspecified.
- **Inspection:** abstract and selected experiments, discussion, and conclusion sections of primary HTML; Crossref venue and DOI.

### dlpcg20

**DLPCG20. Deep Learning for Procedural Content Generation**

Jialin Liu, Sam Snodgrass, Ahmed Khalifa, Sebastian Risi, Georgios N. Yannakakis, and Julian Togelius. 2020 online publication. *Neural Computing and Applications*. DOI: [10.1007/s00521-020-05383-8](https://doi.org/10.1007/s00521-020-05383-8). [arXiv:2010.04548](https://arxiv.org/abs/2010.04548). Direct survey. Tags: deep learning, hybrid methods, content types.

- **Problem and contribution:** surveys deep learning for content generation, including combinations with traditional generation and interactive settings.
- **Method and evaluation:** literature review across content types, not a single comparative experiment.
- **Relevance:** complements SBPCG11 and PCGML18 and provides terminology for following deep learning citations.
- **Limitation:** abstract-level inspection supports its scope but not a detailed ranking of methods; later LLM work requires separate coverage.
- **Inspection:** primary arXiv abstract, journal reference, and DOI; bibliographic record cross-checked in OpenAlex.

### consi25

**CONSI25. Constraint Is All You Need: Optimization-Based 3D Level Generation with LLMs**

Authors (metadata only via Crossref). 2025. *Proceedings of the 20th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3723498.3723840](https://doi.org/10.1145/3723498.3723840). Direct method paper. Tags: LLM, constraints, 3D level generation, optimization, GLDL.

- **Problem and contribution:** Integrating high-level design intentions and game mechanics into complex 3D environments remains hard for PCG. This paper introduces a framework that transforms narrative-level descriptions into playable 3D levels.
- **Method and evaluation:** LLMs parse natural-language descriptions into a structured Game Level Description Language (GLDL) capturing spatial constraints; level generation is modeled as a Facility Layout Optimization problem so placements and configurations adhere to specified design criteria. Experiments include automated constraint evaluations and agent-based simulations, reporting feasibility and stability of constraints extracted from textual descriptions.
- **Relevance:** Directly relevant to this project's "initially undefined constraints": shows LLM-driven constraint extraction and optimization-based assembly can produce valid 3D level layouts from prose.
- **Scope inference:** evaluated on 3D level generation, not character-grounded unit design; the GLDL vocabulary is supplied, not discovered.
- **Inspection:** primary abstract (OpenAlex), Crossref metadata checked. Full methods/effect sizes not inspected.

### coopen25

**COOPEN25. Procedural Content Generation for Cooperative Games—A Systematic Review**

Authors (metadata only via Crossref). 2025. *IEEE Transactions on Games*. DOI: [10.1109/TG.2025.3530419](https://doi.org/10.1109/TG.2025.3530419). Direct systematic review. Tags: cooperative games, systematic review, mixed-initiative, autonomy.

- **Problem and contribution:** Systematic review of PCG for cooperative games, comparing content generation for cooperative versus other game types.
- **Method and evaluation:** Systematic literature review; results show a lack of research specifically on cooperative-game PCG despite varied use cases, and methods span autonomous generation and co-creative/mixed-initiative design.
- **Relevance:** Maps the state of cooperative content generation; confirms this is under-explored, and that mixed-initiative design is a recognized path — relevant to eliciting constraints from a human alongside automated generation.
- **Scope inference:** does not address character-grounded unit design; generalizes across genres.
- **Inspection:** primary abstract (OpenAlex); full review protocol and included-study extraction not inspected.

### latent23

**LATENT23. Learning latent representations for controllable combinational creativity and game design**

Authors (metadata only via Crossref). 2023. DOI: [10.17760/d20581905](https://doi.org/10.17760/d20581905). Direct method paper. Tags: latent representations, combinational creativity, control, game design.

- **Problem and contribution:** Applies latent variable models to combinational creativity, where existing artifacts are recombined into new ones — a model relevant to recombining a character's abilities into a unit.
- **Method and evaluation:** Not fully inspected; abstract indicates latent modeling applied to creative tasks and game design.
- **Relevance:** Latent-space combination is a candidate mechanism for recombining character traits into mechanical designs.
- **Scope inference:** Not verified as game-specific beyond the title/abstract; needs full-text inspection before method claims.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### diverse25

**DIVERSE25. Diverse Level Generation via Machine Learning of Quality Diversity**

Authors (metadata only via Crossref). 2025. *Proceedings of the 20th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3723498.3723843](https://doi.org/10.1145/3723498.3723843). Direct method paper. Tags: quality diversity, machine learning, level generation, QD.

- **Problem and contribution:** Investigates whether generative ML techniques can replicate the power of evolutionary algorithms for discovering good and diverse game content.
- **Method and evaluation:** Not fully inspected; compares ML-based diversity discovery against quality-diversity baselines.
- **Relevance:** Addresses the diversity-of-output question for a generator: is the unit/ability space covered, or concentrated in a few archetypes?
- **Scope inference:** Level-generation focus; diversity defined by behavior metrics supplied by the authors.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### spellspace25

**SPELLSPACE25. Exploring the Possibility Space of 1 Billion Spells**

Authors (metadata only via Crossref). 2025. *Proceedings of the 20th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3723498.3723783](https://doi.org/10.1145/3723498.3723783). Direct method paper. Tags: spell generation, combinatorial mechanics, possibility space.

- **Problem and contribution:** Explores large combinatorial spaces of game mechanics (spells) to characterize the reachable possibility space and the distribution of interesting outcomes.
- **Method and evaluation:** Not fully inspected; abstract indicates exploration of a large spell-combination space.
- **Relevance:** Spells are a natural analogue to character abilities; studying a combinatorial ability space speaks directly to the "abilities" output type of this project.
- **Scope inference:** Combinatorial spell space, not character-grounded generation; interestingness still subjectively assessed.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### spellforger25

**SPELLFORGER25. SpellForger: Prompting Custom Spell Properties In-Game using BERT supervised-trained model**

Authors (metadata only via Crossref). 2025. DOI: [10.5753/sbgames_estendido.2025.14890](https://doi.org/10.5753/sbgames_estendido.2025.14890). Direct method paper. Tags: spell properties, in-game generation, BERT, supervised prompting.

- **Problem and contribution:** Generates custom spell properties in-game using a BERT model trained via supervised prompting.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** In-game, procedural generation of ability-like properties; adjacent to generating unit abilities conditioned on character traits.
- **Scope inference:** Single-game spell system; not character-grounded.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### gameplay25

**GAMEPLAY25. Gameplay Evolution: A Game Design Method Based on Evolutionary Theory**

Authors (metadata only via Crossref). 2025. *IEEE Conference on Games*. DOI: [10.1109/CoG64752.2025.11114235](https://doi.org/10.1109/CoG64752.2025.11114235). Direct method paper. Tags: evolutionary theory, game design method, evolution, design.

- **Problem and contribution:** Proposes a game design method grounded in evolutionary theory; details not fully inspected.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Evolutionary metaphors for design space search are a recurring pattern in PCG (SBPCG11, quality diversity); this adds a design-theory framing.
- **Scope inference:** Method proposal; needs full text to assess novelty and applicability.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### rpgagent26

**RPGAGENT26. RPGAgent: Driving Coherent Story-to-Play Generation with an LLM-Based Multi-Agent System**

Authors (metadata only via Crossref). 2026. DOI: [10.1145/3772318.3790326](https://doi.org/10.1145/3772318.3790326). Direct method paper. Tags: story-to-play, multi-agent LLM, narrative, coherence, RPG.

- **Problem and contribution:** Uses a multi-agent LLM system to generate coherent "story-to-play" content — bridging narrative and interactive gameplay.
- **Method and evaluation:** Not fully inspected; abstract indicates multi-agent coordination focused on coherence.
- **Relevance:** The "coherent story-to-play" task is conceptually close to deriving coherent unit design from a character's narrative role; coherence is exactly the dimension this project must evidence.
- **Scope inference:** RPG narrative/play, not character-grounded TD/RTS units; coherence measured by the authors.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### enemy26

**ENEMY26. An Exploration of Collision-based Enemy Morphology Generation**

Authors (metadata only via Crossref). 2026. DOI: [10.1145/3815598.3815636](https://doi.org/10.1145/3815598.3815636). Direct method paper. Tags: enemy generation, morphology, collision, procedural.

- **Problem and contribution:** Explores procedural generation of enemy morphology based on collision properties.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Enemy/monster morphology generation is a direct analogue to unit-body/ability generation; collision-based grounding is one way to tie geometry/behavior to function.
- **Scope inference:** Morphology-level, not semantic/character-grounded.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### skills26

**SKILLS26. Generate Diverse Skills with Large Language Models**

Authors (metadata only via Crossref). 2026. *Proceedings of the 20th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3815598.3815673](https://doi.org/10.1145/3815598.3815673). Direct method paper. Tags: skill generation, LLM, diversity, game skills.

- **Problem and contribution:** Generates diverse skills using large language models, addressing the homogeneity problem common to LLM-generated game content.
- **Method and evaluation:** Not fully inspected; title/abstract indicate an LLM-based skill-generation method targeting diversity.
- **Relevance:** Skills are the unit's ability set; diversity-of-skills is a direct output quality for this project, and the homogeneity problem (also seen in PERSONAWEAVER) is a key risk to mitigate.
- **Scope inference:** Skill taxonomy unspecified; not character-grounded.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### gptgames25

**GPTGAMES25. GPT for Games: An Updated Scoping Review (2020–2024)**

Authors (metadata only via Crossref). 2025. *IEEE Transactions on Games*. DOI: [10.1109/TG.2025.3563780](https://doi.org/10.1109/TG.2025.3563780). Direct scoping review. Tags: GPT, LLM, games, scoping review, 2020-2024.

- **Problem and contribution:** Scoping review of GPT-family model use in games for the period 2020–2024, updating the earlier LLMGAMES24 survey.
- **Method and evaluation:** Systematic scoping of the literature; not fully inspected.
- **Relevance:** Provides the current (2025) state of LLM-in-games research; essential background for justifying LLM-based design as a component.
- **Scope inference:** Broad LLM-in-games coverage; no focus on character grounding.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### ubcl26

**UBCL26. UBCL: A Reinforcement Learning Framework for Controllable and Diverse Player Behaviors**

Authors (metadata only via Crossref). 2026. *IEEE Transactions on Games*. DOI: [10.1109/TG.2026.3703355](https://doi.org/10.1109/TG.2026.3703355). Direct method paper. Tags: controllable behaviors, diverse behaviors, RL, player modeling.

- **Problem and contribution:** A reinforcement-learning framework for producing controllable and diverse player behaviors — a player-modeling tool rather than a generator.
- **Method and evaluation:** Not fully inspected; abstract indicates RL training of behavior policies with controllable and diverse outcomes.
- **Relevance:** Player/agent behavior diversity is needed both for the character-grounded unit's playstyle and for automated evaluation agents; this shows how to make agent behavior both controllable and diverse.
- **Scope inference:** Player behavior, not unit generation; control parameters unspecified.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### bsp26

**BSP26. Two-Scale Controllability of a Parameterized BSP Dungeon Generator for Player Flow Optimization**

Authors (metadata only via Crossref). 2026. *IEEE Transactions on Games*. DOI: [10.1109/TG.2026.3702534](https://doi.org/10.1109/TG.2026.3702534). Direct method paper. Tags: BSP, dungeon generation, controllability, player flow, two-scale.

- **Problem and contribution:** Studies two-scale controllability of a parameterized BSP dungeon generator for optimizing player flow.
- **Method and evaluation:** Not fully inspected; title/abstract indicate parameterized generation with controllability analysis at two scales.
- **Relevance:** Controllability is a central open problem for this project (PCGBENCH25 defines it per problem); a two-scale analysis of how much control a generator can actually honor is directly instructive for evaluating any character-to-unit mapping.
- **Scope inference:** Dungeon/level geometry; controllability about player flow, not character fidelity.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### gameplay26

**GAMEPLAY26. From Gameplay Traces to Game Mechanics: Causal Induction with Large Language Models**

Authors (metadata only via arXiv API). 2026. arXiv:2602.00190. Direct method paper. Tags: causal induction, gameplay traces, VGDL, reverse-engineering rules, LLM.

- **Problem and contribution:** Investigates whether LLMs can reverse-engineer Video Game Description Language (VGDL) rules from gameplay traces (causal induction), reducing redundancy by selecting nine representative GVGAI games via semantic embeddings/clustering.
- **Method and evaluation:** Compares direct code generation from observations vs. a two-stage method that first infers a structural causal model (SCM) then translates to VGDL; evaluated across prompting strategies and context regimes. The SCM-based approach more often produces VGDL descriptions closer to ground truth, with preference win rates up to 81% in blind evaluations and fewer logically inconsistent rules.
- **Relevance:** Reverse-engineering rules from observation data is a direct analogue of deriving unit behavior from observed play, and the SCM-two-stage pipeline may generalize to deriving mechanics from a character dossier.
- **Scope inference:** VGDL rules from gameplay traces, not character grounding; 81% win rate is a narrow setting.
- **Inspection:** primary arXiv abstract.

### multiobj25

**MULTIOBJ25. Multi-Objective Instruction-Aware Representation Learning in Procedural Content Generation RL**

Authors (metadata only via arXiv API). 2025. arXiv:2508.09193v3. Direct method paper. Tags: instructed PCG, multi-objective RL, instruction encoding, controllability.

- **Problem and contribution:** Instructed PCG-RL struggles to leverage rich textual instructions under complex multi-objective conditions, limiting controllability.
- **Method and evaluation:** Proposes MIPCGRL: multi-objective representation learning incorporating sentence embeddings as conditions, with multi-label classification and multi-head regression. Reports up to 13.8% improvement in controllability with multi-objective instructions.
- **Relevance:** Natural-language instructions as control signals for PCG — directly relevant to feeding character-derived descriptions into a generator; controllability is the open problem PCGBENCH25 isolates.
- **Scope inference:** Two-dimensional PCG environments; instructions control generation, not character fidelity.
- **Inspection:** primary arXiv abstract.

### runtime26

**RUNTIME26. Runtime Evaluation of Procedural Content Generation in an Endless Runner Game Using Autonomous Agents**

Authors (metadata only via arXiv API). 2026. arXiv:2605.01783v1. Direct method paper. Tags: runtime PCG, autonomous agents, evaluation, endless runner, WFC.

- **Problem and contribution:** PCG introduces an evaluation problem — generated content may become unbalanced, blocked, repetitive, or unsolvable.
- **Method and evaluation:** Momentum, an endless-runner with runtime terrain generation, environment spawning, and autonomous agent-based evaluation in the same loop; tile/object placement uses a Wave Function Collapse-inspired constraint mechanism; two autonomous agents (aerial geometric scanner and ground-traversal nav agent) inspect generated content before it reaches the player, using ray casting, physics sweeps, and crash reporting.
- **Relevance:** Demonstrates generation and validation unified in a single runtime loop with autonomous agent triage — a candidate architecture for continuous unit-balance validation.
- **Scope inference:** Endless runner terrain, not character-grounded units.
- **Inspection:** primary arXiv abstract.

### zero3d25

**ZERO3D25. Zero-shot 3D Map Generation with LLM Agents: A Dual-Agent Architecture for Procedural Content Generation**

Authors (metadata only via arXiv API). 2025. arXiv:2512.10501v2. Direct method paper. Tags: zero-shot PCG, LLM agents, 3D map generation, actor-critic.

- **Problem and contribution:** LLMs fail to bridge the semantic gap between abstract user instructions and strict PCG parameter specifications.
- **Method and evaluation:** A training-free actor-critic agent pair that autonomously reasons over tool parameters and refines configurations to align with human design preferences; validated on 3D map generation, establishing a benchmark for instruction-following in PCG; outperforms single-agent baselines with diverse, structurally valid environments.
- **Relevance:** Zero-shot LLM agents as PCG-tool controllers — a candidate for translating character-derived instructions into concrete design parameters without training.
- **Scope inference:** 3D map parameters, not character grounding; training-free architecture.
- **Inspection:** primary arXiv abstract.

### instruct25

**INSTRUCT25. IPCGRL: Language-Instructed Reinforcement Learning for Procedural Level Generation**

Authors (metadata only via arXiv API). 2025. *IEEE Conference on Games*. DOI: [10.1109/CoG64752.2025.11114105](https://doi.org/10.1109/CoG64752.2025.11114105). Direct method paper. Tags: instructed RL, level generation, natural language, controllability.

- **Problem and contribution:** Natural language as a controllability modality for content-generation models in PCG-RL.
- **Method and evaluation:** Not fully inspected; title/abstract indicate language-instructed RL for procedural level generation.
- **Relevance:** Language-instructed generation is a direct line of work to the project's prose-to-design goal.
- **Scope inference:** Level generation, not units; RL-based.
- **Inspection:** abstract-level via arXiv/OpenAlex.

### cppn26

**CPPN26. CPPN2WFC: Extending Wave Function Collapse to Generate Globally Coherent Content**

Authors (metadata only via arXiv API / EPYC). 2026. *Proceedings of the Genetic and Evolutionary Computation Conference*. DOI: [10.1109/CEC...](https://doi.org/10.1109/CEC...) (venue DOI truncated). Direct method paper. Tags: Wave Function Collapse, CPPN, global coherence, constraint satisfaction.

- **Problem and contribution:** Extends Wave Function Collapse with Composite Pattern Neural Networks (CPPNs) to generate globally coherent content.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** WFC is a constraint-satisfaction PCG method; CPPN extension addresses global coherence — relevant to ensuring generated units behave coherently rather than locally valid but globally broken.
- **Scope inference:** Content coherence, not character grounding.
- **Inspection:** abstract-level via arXiv/OpenAlex.

### llmgames24

**LLMGAMES24. Large Language Models and Games: A Survey and Roadmap**

Roberto Gallotta, Graham Todd, Marvin Zammit, Sam Earle, Antonios Liapis, Julian Togelius, and Georgios N. Yannakakis. 2024 early access. *IEEE Transactions on Games*. DOI: [10.1109/TG.2024.3461510](https://doi.org/10.1109/TG.2024.3461510). [arXiv:2402.18659](https://arxiv.org/abs/2402.18659). Direct survey. Tags: foundation models, generation, agents, research map.

- **Problem and contribution:** surveys the roles LLMs can occupy in and for games, their potential, and limitations.
- **Method and evaluation:** literature synthesis and research agenda.
- **Relevance:** helps distinguish a model that designs content from one that plays, evaluates, or converses in a game.
- **Scope inference:** evidence for one of those roles cannot establish competence in the others. A survey's proposed roadmap is not an adopted project direction.
- **Inspection:** primary arXiv abstract and publication fields. Detailed classifications and recommendations have not been extracted from full text.

## Tower Defense, RTS units, and game-specific evidence

### rtsunits22

**RTSUNITS22. Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search**

Kynan Sorochan and Matthew Guzdial. 2022. arXiv manuscript; publication venue not verified. [arXiv:2212.03387](https://arxiv.org/abs/2212.03387). Direct method paper. Tags: unit generation, balance, search, simulation.

- **Problem and contribution:** generates new units for microRTS and studies automated approximations of usefulness and balance.
- **Method:** search with MCTS gameplay evaluation; the paper presents ten generated units. Its fitness evaluation first gives only one agent the unit and then gives both agents access, measuring construction, survival, and wins. Each fitness call uses two rounds of ten games. A separate evaluation varies MCTS budgets to approximate strong, medium, and weak players and uses 100 games per matchup/round/unit. The authors report evidence of utility and skill ordering in parts of this evaluation, but their usefulness assumption fails in the Strong-versus-Strong and Medium-versus-Medium matchups. The strongest agent builds the generated units least often, averaging 31.7%. The paper reruns one Strong-versus-Strong batch because its average unit-construction count of 7.4 in 100 games falls below the stated threshold of 25. These details qualify the overall usefulness and balance claim.

The paper explicitly limits its conclusions: it uses one standard map, no human playtesting, and restricted computational budgets. It says balance for humans remains unclear. Its greedy hill-climbing search operates over six numerical attributes and a predefined set of four causes and four effects. Names are assigned to the units after generation.

- **Relevance:** unusually direct prior work on the eventual output type. It does not study character identity, unknown mechanics, explanations, or adaptation across the user's games.
- **Inspection:** primary arXiv HTML, including representation, search, evaluator, evaluation, limitations, and conclusion sections, revisited on 2026-09-28. Crossref's title search returned an unrelated MCTS paper, which was rejected rather than used as publication metadata.

### tdci11

**TDCI11. Computational intelligence and tower defence games**

Phillipa Avery, Julian Togelius, Elvis Alistar, and Robert Pieter van Leeuwen. 2011. *IEEE Congress of Evolutionary Computation*. DOI: [10.1109/CEC.2011.5949738](https://doi.org/10.1109/CEC.2011.5949738). Direct foundational paper. Tags: genre analysis, experience-driven PCG.

- **Problem and contribution:** introduces TD as a computational intelligence testbed, classifies TD components, and presents a prototype based on experience-driven PCG.
- **Method and evaluation:** the inspected abstract describes a classification and prototype; evaluation details and outcomes were not available in the inspected material.
- **Relevance:** an early TD-specific conceptual source whose categories can be compared with contemporary games later.
- **Scope inference:** its classification should not be adopted as a universal schema for the varied 27-game corpus.
- **Inspection:** Crossref bibliographic record and abstract reproduced by OpenAlex; original full text not inspected.

### tdlevels19

**TDLEVELS19. Automatic generation of tower defense levels using PCG**

Simon Liu, Li Chaoran, Li Yue, Ma Heng, Hou Xiao, Shen Yiming, Wang Licong, Chen Ze, Guo Xianghao, Lu Hengtong, Du Yu, and Tang Qinting. 2019. *14th International Conference on the Foundations of Digital Games*. Author names are reproduced in Crossref's order and spelling. DOI: [10.1145/3337722.3337723](https://doi.org/10.1145/3337722.3337723). Direct method paper. Tags: levels, simulation, game-specific generation.

- **Problem and contribution:** automates level construction for *Kingdom Rush: Frontiers* by analyzing road maps, tower locations, and monster sequences, generating corresponding components, and assembling them into levels. The abstract describes Monte Carlo search for automated playability testing. Evaluation details and numerical findings have not been inspected.
- **Relevance:** directly connects PCG and automated testing to a TD game, but it generates levels rather than character-grounded units.
- **Scope inference:** its three building blocks are specific to the modeled game, not evidence of shared mechanics across the corpus.
- **Inspection:** Crossref publication record and abstract reproduced by OpenAlex; full text not read.

### tdga21

**TDGA21. Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms**

Vid Kraner, Iztok Fister, and Lucija Brezočnik. 2021. *New Technologies, Development and Application IV*, Lecture Notes in Networks and Systems. DOI: [10.1007/978-3-030-75275-0_54](https://doi.org/10.1007/978-3-030-75275-0_54). Direct bibliographic lead retained for its close domain match. Tags: Tower Defense, genetic algorithms.

- **Problem, contribution, method, evaluation, and findings:** not analyzed beyond verified publication metadata. The title identifies a custom TD game and genetic algorithms; no detailed claims are inferred from it.
- **Relevance:** potentially direct TD generation evidence requiring primary text inspection.
- **Limitation:** the current entry cannot support any claim about effectiveness, generated content types, or a game called “Save the Sheep.” Inspection: Crossref bibliographic record; OpenAlex record contained no abstract.

### tdstrategy24

**TDSTRATEGY24. Reinforcement Learning for High-Level Strategic Control in Tower Defense Games**

Joakim Bergdahl, Alessandro Sestini, and Linus Gisslén. 2024. *IEEE Conference on Games*. DOI: [10.1109/CoG60054.2024.10645621](https://doi.org/10.1109/CoG60054.2024.10645621). [arXiv:2406.07980](https://arxiv.org/abs/2406.07980). Direct playtesting paper. Tags: Plants vs. Zombies, reinforcement learning, scripted agents.

- **Problem and contribution:** combines scripted control with reinforcement learning for automated gameplay testing.
- **Method and evaluation:** tests the hybrid agent in *Plants vs. Zombies* over 40 levels. The abstract reports a 57.12% success rate versus 47.95% for heuristic AI and emphasizes the difficulty of training a general agent for this puzzle-like game.
- **Relevance:** empirical TD evidence that automated evaluation depends on agent competence and task structure.
- **Scope inference:** agent success is not a measure of fidelity to a fictional character or the quality of generated unit designs.
- **Inspection:** primary arXiv abstract and Crossref publication metadata.

### towermind26

**TOWERMIND26. TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents**

Dawei Wang, Chengming Zhou, Di Zhao, Xinyuan Liu, Marci Chi Ma, Gary Ushaw, and Richard Davison. 2026. *Proceedings of the AAAI Conference on Artificial Intelligence*, 40(31), according to the arXiv DOI record. DOI: [10.1609/aaai.v40i31.39818](https://doi.org/10.1609/aaai.v40i31.39818). [arXiv:2601.05899](https://arxiv.org/abs/2601.05899). Direct benchmark paper. Tags: TD, LLM agents, multimodal observations, validity.

- **Problem and contribution:** a lightweight configurable TD environment with text, structured-state, and pixel observations.
- **Method and evaluation:** five benchmark levels, five seeds per model and condition, and five human experts; evaluates score and valid-action rate, with additional PPO and Ape-X DQN experiments. The authors report a large gap between tested LLMs and human experts and identify weak planning validation, inefficient actions, and missed opportunities to accomplish multiple goals with one action. Valid actions and effective play differ in the results.
- **Relevance:** direct evidence for the difficulty of TD reasoning and a possible future benchmark resource to inspect.
- **Scope inference:** this evaluates game-playing agents under fixed rules, not generated designs. The inspected text also labels Gemini differently between its main evaluation and appendix tables; that discrepancy should be resolved before reusing model-specific comparisons.
- **Inspection:** primary abstract and selected evaluation/conclusion sections, not the full paper or environment implementation.

### tdnovel22

**TDNOVEL22. A Novel Procedural Content Generation Algorithm for Tower Defense Games**

Authors (metadata only via Crossref). 2022. *Proceedings of the 17th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3564982.3564993](https://doi.org/10.1145/3564982.3564993). Direct method paper. Tags: tower defense, PCG, roguelite.

- **Problem and contribution:** Presents a PCG algorithm for an isometric tower-defense game with roguelite elements.
- **Method and evaluation:** Not fully inspected; abstract indicates a PCG algorithm for TD level/content generation.
- **Relevance:** One of few TD-specific PCG methods; closest genre match to the corpus.
- **Scope inference:** Level/content generation, not character-grounded unit design; TD-specific.
- **Inspection:** abstract-level via Crossref/OpenAlex. Note: the library already contains TDLEVELS19 (Kingdom Rush levels) and TDGA21 (TD genetic algorithms) — this is a distinct third TD-PCG work.

### tdwavegen22

**TDWAVE22. A NEAT Approach to Wave Generation in Tower Defense Games**

Authors (metadata only via Crossref). 2022. DOI: [10.1109/IMET54801.2022.9929595](https://doi.org/10.1109/IMET54801.2022.9929595). Direct method paper. Tags: tower defense, wave generation, NEAT, enemy waves.

- **Problem and contribution:** Uses NeuroEvolution of Augmenting Topologies (NEAT) for generating enemy waves in tower defense games.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Wave (enemy roster) generation is closer to unit design than level generation; NEAT for wave design is an alternative to the LLM/mechanic-evolution approaches.
- **Scope inference:** Enemy waves only, not towers/abilities; TD-specific.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### monster26

**MONSTER26. Application of machine learning to monster level prediction in tabletop RPG game design**

Jolanta Śliwa and Jakub Adamczyk. 2026, first posted July 10; version 3 inspected through the API. arXiv manuscript; venue not verified. [arXiv:2607.09196](https://arxiv.org/abs/2607.09196). Direct design-support paper. Tags: unit attributes, balance, tabular learning, datasets.

- **Problem and contribution:** predicts ordinal monster levels from numerical attributes in publicly available *Pathfinder Second Edition* data.
- **Method and evaluation:** compares regressors with rounding, ordinal methods, and neural losses under chronological and expanding-window protocols. The abstract reports that tree ensembles outperform linear and neural approaches and discusses feature importance and errors.
- **Relevance:** close adjacent evidence on learning quantitative assessments of designed entities.
- **Scope inference:** predicting a published level is not a direct demonstration of balanced combat, generation quality, or faithfulness to a character; this evidence remains tied to one rules system.
- **Inspection:** primary abstract; dataset size and exact accuracy not inspected.

## Automated design, mechanics, and language-guided generation

### mechgen14

**MECHGEN14. Automatic Game Design via Mechanic Generation**

Alexander Zook and Mark O. Riedl. 2014. *AAAI Conference on Artificial Intelligence*, 28(1). DOI: [10.1609/aaai.v28i1.8788](https://doi.org/10.1609/aaai.v28i1.8788). [arXiv:1908.01420](https://arxiv.org/abs/1908.01420) is a later 2019 manuscript deposit of the same work, not a separate 2019 contribution. Direct method paper. Tags: constraints, planning, representation, cross-domain mechanics.

- **Problem and contribution:** formalizes mechanic generation through a composable representation inspired by planning actions. A constraint solver generates mechanics meeting design requirements and a planner checks playability requirements. Demonstrations model RPG, platformer, and combined RPG/platformer mechanics. The inspected conclusion calls for richer playability requirements over action/state trajectories and reasoning about potential gameplay outcomes.
- **Relevance:** foundational evidence for separating rules about the form of mechanics from evidence about their use in play.
- **Scope inference:** predefined modeling primitives and requirements differ from discovering an initially undefined foundation.
- **Inspection:** primary abstract and conclusion, with AAAI DOI metadata independently checked via OpenAlex. The publisher page request failed; no full-paper reading is claimed.

### boss17

**BOSS17. Program Synthesis as a Generative Method**

Eric Butler, Kristin Siu, and Alexander Zook. 2017. *Proceedings of the 12th International Conference on the Foundations of Digital Games*, FDG '17, 10 pages. DOI: [10.1145/3102071.3102076](https://doi.org/10.1145/3102071.3102076). [Author manuscript](https://www.ericbutler.net/assets/papers/fdg2017_boss.pdf). Direct method and feasibility paper. Tags: boss generation, behavior, program synthesis, validity, supplied constraints.

- **Problem and contribution:** generates well-formed programs through grammar expansion constrained by types and variable scope. A hand-authored domain language represents boss morphology, state, and finite-state behavior; function definitions are supplied rather than generated. Examples constrain weak points to the boss body, distinguish speed and duration, and prevent health increasing when damage occurs.
- **Evaluation:** worked examples in a *Moldorm*-like domain and a complete generated *MegaMan*-like boss illustrate behavior composition, including movement, projectiles, and escalation after damage through shared variable references. These are feasibility demonstrations, not comparative tests of balance, enjoyment, or character fidelity.
- **Source report:** section 6.2.2 distinguishes well-formedness from the literal values that remain critical to boss design; section 7 leaves heuristics for higher-quality artifacts to future work. The Figure 5 caption identifies the displayed Cut Man sprite as arbitrary and without semantic significance.
- **Relevance:** close prior work on generating game entities with behavior while making supplied domain knowledge and validity assumptions explicit.
- **Scope inference:** well-formed code and constraints do not by themselves establish playability, balance, identity fidelity, or discovery of unfamiliar mechanics.
- **Inspection:** primary PDF text pp. 1-2 and 4-9, including methods, examples, conclusion, and the Figure 5 caption; figures were not visually inspected. DOI and publication metadata checked in Crossref and the manuscript on 2026-09-28. The ACM landing page returned HTTP 403; the author's publication page supplied an accessible manuscript.

### concept18

**CONCEPT18. Automated Game Design via Conceptual Expansion**

Matthew Guzdial and Mark Riedl. 2018. *AAAI Conference on Artificial Intelligence and Interactive Digital Entertainment*, 14(1). DOI: [10.1609/aiide.v14i1.13022](https://doi.org/10.1609/aiide.v14i1.13022). [arXiv:1809.02232](https://arxiv.org/abs/1809.02232). Direct method paper. Tags: transfer, analogy, learned representations, incomplete specifications.

- **Problem and contribution:** recombines learned approximate game representations using conceptual expansion. Evaluation reconstructs a third game from knowledge of two games and a partial specification of the target, using first-level gameplay videos from *Super Mario Bros.*, *Kirby's Adventure*, and *Mega Man*. Baselines include conceptual blending, nearest neighbor, and a genetic algorithm. The paper reports closer matches to target game graphs. Its authors explicitly say reconstructing an existing game is an imperfect proxy for novel design and that a human study and broader game diversity are needed.
- **Relevance:** concrete cross-game combination and partial-specification work. Its evaluation measures representation distance, not fictional-character fidelity or broad adaptation across genres.
- **Inspection:** abstract and selected evaluation/discussion/conclusion sections of primary HTML; publication record checked in OpenAlex.

### gavel24

**GAVEL24. GAVEL: Generating Games via Evolution and Language Models**

Graham Todd, Alexander Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers, and Julian Togelius. 2024. *Advances in Neural Information Processing Systems 37*. [arXiv:2407.09388](https://arxiv.org/abs/2407.09388). Crossref proceedings DOI: [10.52202/079017-3515](https://doi.org/10.52202/079017-3515). Direct method paper. Tags: language models, fine-tuning, quality diversity, rules, evaluation.

- **Problem and contribution:** generates games within Ludii by combining a trained language model with evolutionary search.
- **Method:** mutate and recombine rules, retain games in a quality-diversity archive, and evaluate compilability, playability, balance, decisiveness, completion, agency, coverage, and strategic-depth proxies. Evaluation compares three runs with an alternative mutation-selection policy, pure sampling, and GPT-4o sampling; the paper reports better archive quality and coverage for GAVEL under those metrics.

The authors explicitly warn that metric satisfaction is far from sufficient evidence of interesting games and use expert filtering. Other reported limits include unused generated components, shared ancestry of held-out seeds, and dependence on the archive representation. Short MCTS rollouts and move limits also exclude some potentially interesting games.

- **Relevance:** central evidence for combining generation and evaluation while keeping those judgments distinct. It assumes a rich but fixed game language and does not establish open-ended mechanics discovery or character adaptation.
- **Inspection:** primary abstract and selected evaluation, experiments, and limitations sections; Crossref publication record.

### mm23

**MM23. Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules**

Johor Jara Gonzalez, Seth Cooper, and Matthew Guzdial. 2023. *AAAI Conference on Artificial Intelligence and Interactive Digital Entertainment*, 19(1). DOI: [10.1609/aiide.v19i1.27522](https://doi.org/10.1609/aiide.v19i1.27522). [arXiv:2309.09476](https://arxiv.org/abs/2309.09476). Direct method paper. Tags: rule generation, reinforcement learning, human-proxy evaluation.

- **Problem and contribution:** investigates RL as an approximation of human learning during evaluation of generated rules, replacing static play approximations. It recreates Mechanic Maker in Unity and compares generated rule sets with an A* agent baseline. The abstract reports distinct rules that may be more usable by humans. That tentative wording is preserved: it is not proof of human usability.
- **Relevance:** shows that the evaluator can change what a generator produces.
- **Scope inference:** evidence that different agents prefer different rules cannot by itself identify which evaluator matches the intended audience.
- **Inspection:** primary arXiv abstract and Crossref publication metadata.

### ggdg24

**GGDG24. Grammar-based Game Description Generation using Large Language Models**

Tsunehiko Tanaka and Edgar Simo-Serra. 2024 early access. *IEEE Transactions on Games*. DOI: [10.1109/TG.2024.3520214](https://doi.org/10.1109/TG.2024.3520214). [arXiv:2407.17404](https://arxiv.org/abs/2407.17404). Direct method paper. Tags: natural-language specifications, grammars, validity.

- **Problem and contribution:** converts natural-language descriptions into game-description-language output.
- **Method:** progressively construct a minimal grammar from the language specification, then iteratively refine output through grammar-guided generation and a parser that identifies valid subsequences and candidate symbols. The abstract reports better performance than direct LLM output.
- **Relevance:** closely related to translating prose into formal game descriptions.
- **Scope inference:** grammatical correctness does not establish semantic fidelity, complete mechanics, or good play; the grammar is supplied. Detailed evaluation tasks and scores were not inspected.
- **Inspection:** primary abstract and arXiv journal/DOI fields.

### chatpcg24

**CHATPCG24. ChatPCG: Large Language Model-Driven Reward Design for Procedural Content Generation**

In-Chang Baek, Tae-Hwa Park, Jin-Ha Noh, Cheong-Mok Bae, and Kyung-Joong Kim. 2024. *IEEE Conference on Games*. DOI: [10.1109/CoG60054.2024.10645619](https://doi.org/10.1109/CoG60054.2024.10645619). [arXiv:2406.11875](https://arxiv.org/abs/2406.11875). Direct method paper. Tags: reward design, language models, reinforcement learning.

- **Problem and contribution:** uses an LLM to generate reward functions tailored to a specified game's content-generation task and integrates them with deep RL. The abstract reports support for multiplayer game content generation and claims evidence of task understanding. Evaluation details and numerical results have not been inspected.
- **Relevance:** illustrates research on reducing reward-authoring work.
- **Scope inference:** this does not show that an LLM can identify all relevant constraints or validate its own reward against designer intent.
- **Inspection:** primary abstract and Crossref publication metadata.

### pcgrllm26

**PCGRLLM26. PCGRLLM: Large Language Model-Driven Reward Design for Procedural Content Generation Reinforcement Learning**

In-Chang Baek, Sung-Hyun Kim, Sam Earle, Zehua Jiang, Jin-Ha Noh, Julian Togelius, and Kyung-Joong Kim. 2026. *IEEE Transactions on Games*. DOI: [10.1109/TG.2026.3695197](https://doi.org/10.1109/TG.2026.3695197). [arXiv:2502.10906](https://arxiv.org/abs/2502.10906), first posted 2025. Direct method paper. Tags: language models, reward design, instructions, reinforcement learning.

- **Problem and contribution:** extends earlier language-model reward design with feedback and reasoning-oriented prompts. The abstract reports a story-to-reward task in a two-dimensional environment, comparison of two language models and prompting variants, and better performance than its earlier structure, approaching the study's human comparison. Evaluation details and the meaning of that comparison were not inspected.
- **Relevance:** a recent example of translating narrative descriptions into evaluative signals for PCG, distinct from generating complete units.
- **Scope inference:** its specified two-dimensional task and reward interface do not establish an ability to discover every game constraint or interpret fictional characters.
- **Inspection:** primary arXiv abstract, arXiv related DOI, and exact Crossref publication record, checked 2026-09-27.

### mortar26

**MORTAR26. Mortar: Evolving Mechanics for Automatic Game Design**

Muhammad U. Nasir, Yuchen Li, Steven James, and Julian Togelius. First posted 2025-12-31; arXiv identifier assigned in 2026. Manuscript; venue not verified. [arXiv:2601.00105](https://arxiv.org/abs/2601.00105). Direct method paper. Tags: mechanics, quality diversity, language models, skill ordering.

- **Problem and contribution:** evolves mechanics with a language model and quality-diversity search, then composes games through tree search in a fixed top-down framework.
- **Method:** screening checks syntax, runtime behavior, and MCTS play. Game fitness uses Kendall's tau to compare the ordering of five agents, three MCTS budgets plus random and no-action agents, with an expected skill ordering. Constrained Importance Through Search assigns zero to unvisited mechanic subsets when estimating contribution.
- **Evaluation and limits:** the inspected user study has ten participants compare three pairs drawn from six games across five dimensions. The paper's aggregate positive-minus-frustration score favors the higher-tau game in two pairs and the lower-tau game in one; the five individual dimensions do not all agree. This small comparison does not validate the proxy across the full generated space. Section 6 identifies the absence of designer control in the autonomous evolution process. An internal discrepancy remains unresolved: the method lists eight mechanic types while the experiment setup says nine.
- **Relevance:** current work generating mechanics as reusable components and judging them in composed games.
- **Scope inference:** skill ordering does not establish character fidelity or universal design quality.
- **Inspection:** primary arXiv abstract and sections 2.2, 2.4, 3, 4.1, and 6 on 2026-09-28; full effect-size and supplementary analysis not completed.

### tanagra11

**TANAGRA11. Tanagra: Reactive Planning and Constraint Solving for Mixed-Initiative Level Design**

Gillian Smith, Jim Whitehead, and Michael Mateas. 2011. *IEEE Transactions on Computational Intelligence and AI in Games* 3(3):201-215. DOI: [10.1109/TCIAIG.2011.2159716](https://doi.org/10.1109/TCIAIG.2011.2159716). Direct mixed-initiative system paper. Tags: human designer, constraints, level design, authoring.

- **Problem and contribution:** combines a human designer's edits with reactive level generation for two-dimensional platformers. The abstract describes planning and numerical constraint solving that respond to geometry and pacing changes while preserving modeled playability.
- **Method and evaluation:** the paper reports an expressive-range evaluation; its detailed procedure and results were not inspected.
- **Relevance:** primary evidence that mixed-initiative design can keep a human in the loop while enforcing a specified game's constraints.
- **Scope inference:** this level-authoring tool assumes a player model and platformer geometry grammar, and does not establish a method for discovering character concepts or generating TD units.
- **Inspection:** Crossref publication metadata and abstract reproduced by OpenAlex; IEEE full text not read. A 2010 Tanagra demonstration exists, but this retained journal article is the primary system account in this library.

### cad26

**CAD26. Procedural Content Metageneration via Program Search and Continual Abstraction Discovery**

Matthew Siper, Ahmed Khalifa, and Julian Togelius. 2026, posted August 18. arXiv manuscript; venue not verified. [arXiv:2608.17947](https://arxiv.org/abs/2608.17947). Direct method paper. Tags: generators, abstraction discovery, reusable primitives, language models.

- **Problem and contribution:** evolves complete Python generators using LLM mutation and crossover and extracts reusable primitives into a run-specific helper module.
- **Method and evaluation:** a 2-by-2 comparison of Continual Abstraction Discovery and a fixed hand-written domain API, covering Sokoban, Zelda, Dangerous Dave, and Lode Runner; the abstract reports 160 completed runs and increased mean final-best fitness in all eight domain/API comparisons.
- **Relevance:** directly relevant to the question of fixed primitives versus acquired abstractions.
- **Scope inference:** utilities learned within four specified level domains do not establish open-ended discovery of character mechanics or transfer among TD games.
- **Inspection:** primary abstract only; fitness, compute matching, and abstraction correctness need deeper inspection.

### autobg26

**AUTOBG26. AutoBG: A Board Game Design Assistant with Interactive Ideation, Iterative Rulebook Generation, and Individualized Feedback**

Zizhen Li, Chuanhao Li, Yibin Wang, Jianwen Sun, Yukang Feng, Fanrui Zhang, Mingzhu Sun, Yifei Huang, and Kaipeng Zhang. 2026, first posted June 1. arXiv manuscript; venue not verified. [arXiv:2606.01976](https://arxiv.org/abs/2606.01976). Direct mixed-initiative paper. Tags: incomplete specifications, rulebooks, feedback, training.

- **Problem and contribution:** integrates ideation, rulebook realization, criticism, and personalized feedback. The abstract describes 2.2K structured rulebooks, 180K filtered player reviews, evaluation on 207 held-out games, and a 30-person user study. It reports better outputs than tested baselines and favorable participant feedback.
- **Relevance:** current evidence about natural-language game-design artifacts and support from vague initial intent.
- **Scope inference:** synthetic audience feedback and critic-approved revisions require separate validity checks; the abstract's claim that rulebooks approach published quality cannot be accepted as a universal result.
- **Inspection:** primary abstract only; evaluation construction, data leakage, and human-study details remain uninspected.

### playchar25

**PLAYCHAR25. Generative Methods for Creating Adaptive Playable Characters in Service Games**

Authors (metadata only via Crossref). 2025. *Automatic Documentation and Mathematical Linguistics*. DOI: [10.3103/s0...](https://doi.org/10.3103/s0...) (journal DOI truncated in source). Direct method paper. Tags: playable character generation, adaptive characters, service games, generative.

- **Problem and contribution:** As games-as-a-service require constant content updates, automating generation of adaptive playable characters is framed as urgent; the paper reviews approaches to adaptive playable-character generation.
- **Method and evaluation:** Not fully inspected; abstract indicates a review of existing approaches to automating adaptive playable character generation.
- **Relevance:** The single closest computational analogue to this project's core task: automating generation of playable characters, explicitly framed as a generation-and-adaptation problem.
- **Scope inference:** Service-game playable characters in general; "adaptive" likely means dynamically adapting, not adapting-from-a-reference-character. Does not address fidelity to a named source character.
- **Inspection:** abstract-level via Crossref/OpenAlex. Access note: low-citation journal venue; verify full text.

### starcharm25

**STARCHARM25. Democratizing Game Modding with GenAI: A Case Study of StarCharM, a Stardew Valley Character Modding Assistant**

Authors (metadata only via Crossref). 2025. *Proceedings of the ACM on Human-Computer Interaction*. DOI: [10.1145/37...](https://doi.org/10.1145/37...) (DOI truncated in source). Adjacent mixed-initiative paper. Tags: game modding, character modding, GenAI, HCI, user study.

- **Problem and contribution:** StarCharM is a GenAI-assisted character-modding tool for *Stardew Valley*; the paper reports on enabling non-experts to create character mods with AI support.
- **Method and evaluation:** Case study with human users; details not fully inspected.
- **Relevance:** Adjacent evidence that GenAI can assist in adapting/modifying character content for a specific game, with a user-study evaluation path.
- **Scope inference:** Modding support (human-led), not autonomous unit generation; specific to one game.
- **Inspection:** abstract-level via Crossref/OpenAlex; full text not inspected.

### devwhat26

**DEVWHAT26. What game developers actually want from procedural level generation tools**

Authors (metadata only via Crossref). 2026. *Proceedings of the 20th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3815598.3815682](https://doi.org/10.1145/3815598.3815682). Adjacent human-study paper. Tags: developer needs, user study, PCG tools, requirements.

- **Problem and contribution:** Investigates what game developers actually want from PCG tools, based on empirical study.
- **Method and evaluation:** Not fully inspected; abstract indicates a user-facing needs assessment.
- **Relevance:** Human-centered evidence for what constraints/designers care about; directly informs how to structure constraint elicitation and what to present to users.
- **Scope inference:** General PCG tool needs, not character-grounded units.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### spacetime26

**SPACETIME26. Constraint-Based Four-Dimensional Spacetime Blending of Learned Game Mechanics Along Orthogonal Axes**

Authors (metadata only via Crossref). 2026. *Proceedings of the 20th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3815598.3815702](https://doi.org/10.1145/3815598.3815702). Direct method paper. Tags: mechanics blending, spacetime, constraints, learned mechanics.

- **Problem and contribution:** Blends learned game mechanics along orthogonal axes in a four-dimensional spacetime under constraints.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Mechanical blending/composition is exactly what a character→unit adapter must do (recombine a character's abilities); constraint-based blending addresses well-formedness.
- **Scope inference:** Learned mechanics blending, not character-grounded; four axes unspecified.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### hades26

**HADES26. Vanquish Your Past: Shifted Imitation Learning in Hades**

Authors (metadata only via Crossref). 2026. DOI: [10.1145/3815598.3815650](https://doi.org/10.1145/3815598.3815650). Direct method paper. Tags: imitation learning, Hades, gameplay adaptation.

- **Problem and contribution:** Applies shifted imitation learning to reproduce/adapt *Hades* gameplay; details not fully inspected.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Imitation/adaptation of an existing game's behavior is a related transfer problem; shifts in the learning regime (as the title implies) may generalize to adapting character behavior.
- **Scope inference:** Single-game behavior reproduction, not character-to-unit design.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### behbts26

**BEHBTS26. Large Language Models for Behavior Trees Generation in Unreal Engine**

Authors (metadata only via Crossref). 2026. DOI: [10.1145/3815598.3815631](https://doi.org/10.1145/3815598.3815631). Direct method paper. Tags: behavior trees, LLM, Unreal Engine, AI behavior.

- **Problem and contribution:** Generates behavior trees for game AI using large language models within Unreal Engine.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Behavior trees are a standard representational form for unit behavior; LLM-driven behavior-tree generation is a concrete path to turning character-driven behavior specifications into executable unit logic.
- **Scope inference:** Behavior trees for generic AI, not character-grounded; Unreal Engine specific.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### invention25

**INVENTION25. Generation and Evaluation in the Human Invention Process through the Lens of Game Design**

Katherine M. Collins, Graham Todd, Cedegao E. Zhang, Adrian Weller, Julian Togelius, Junyi Chu, Lionel Wong, Thomas L. Griffiths, and Joshua B. Tenenbaum. 2025. CogSci conference nonarchival paper, as labeled in the arXiv record. [arXiv:2508.10914](https://arxiv.org/abs/2508.10914). Adjacent empirical and modeling paper. Tags: human creativity, game invention, evaluation, seed examples.

- **Problem and contribution:** studies early human invention of two-player grid strategy games, asking how people propose and assess new rules. The abstract describes over 450 participant-created games after exposure to a seed set and compares an associative proposal account with a bounded model-based evaluation account. It reports that population-level creations are better described by a model including estimates of game quality. Detailed study design, fit statistics, and alternative explanations were not inspected.
- **Relevance:** provides empirical context for the distinction between generating candidate ideas and judging them.
- **Scope inference:** novice human creation of grid games does not predict what characters or TD designers will judge faithful or useful.
- **Inspection:** primary arXiv abstract and publication-status comment, checked 2026-09-27.

## Rule representations and reusable corpora

### vgdl13

**VGDL13. A video game description language for model-based or interactive learning**

Tom Schaul. 2013. *IEEE Conference on Computational Intelligence in Games*. DOI: [10.1109/CIG.2013.6633610](https://doi.org/10.1109/CIG.2013.6633610). Direct representation paper. Tags: game languages, ontology, simulation, benchmark.

- **Problem and contribution:** proposes PyVGDL, a high-level language and interpreter for two-dimensional games, with building blocks, dynamics, and collision interactions. The abstract describes examples of classic games and agent interfaces supporting several learning settings.
- **Relevance:** primary evidence about an existing executable rule vocabulary and its connection to generation and evaluation.
- **Scope inference:** expressiveness is bounded by its supplied ontology; adoption would be a separate technical decision.
- **Inspection:** Crossref publication metadata and abstract reproduced in OpenAlex. Full language and evaluation details not inspected.

### ludii19

**LUDII19. An Overview of the Ludii General Game System**

Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers, and Cameron Browne. 2019. arXiv manuscript; venue not verified in this pass. [arXiv:1907.00240](https://arxiv.org/abs/1907.00240). Direct system paper. Tags: game languages, ludemes, traditional games, agents.

- **Problem and contribution:** describes Ludii's support for modeling, modifying, and playing traditional strategy games in the Digital Ludeme Project, whose stated ambition covers over 1,000 games.
- **Method and evaluation:** system overview, as described by the abstract; performance and coverage claims have not been independently analyzed.
- **Relevance:** context for GAVEL24 and a concrete example of a reusable cross-game vocabulary.
- **Scope inference:** a system for traditional strategy games cannot be assumed to represent the full TD/RTS corpus or anime character concepts.
- **Inspection:** primary arXiv abstract and record.

### vglc16

**VGLC16. The VGLC: The Video Game Level Corpus**

Adam James Summerville, Sam Snodgrass, Michael Mateas, and Santiago Ontañón. 2016. arXiv manuscript; venue not verified in this pass. [arXiv:1606.07487](https://arxiv.org/abs/1606.07487). Direct dataset paper. Tags: datasets, representation, PCGML.

- **Problem and contribution:** makes game levels available in formats intended to be easy to parse for learning and other game-AI research.
- **Method and evaluation:** dataset contribution; the inspected abstract does not specify a generator comparison.
- **Relevance:** shows how reusable research corpora can support comparisons and learning.
- **Scope inference:** levels are a different unit of evidence from complete gameplay-unit designs, and a machine-readable format does not establish comparable semantics across games.
- **Inspection:** primary arXiv abstract. This paper entry is distinct from any future inspection of the accompanying repository, its data licenses, or its current contents.

### behbtsunreal26

**BEHBTSUN26. Large Language Models for Behavior Trees Generation in Unreal Engine**

Authors (metadata only via Crossref). 2026. DOI: [10.1145/3815598.3815631](https://doi.org/10.1145/3815598.3815631). (Duplicate entry — see BEHBTS26 above; retained here as a rules-representation angle.) Direct method paper. Tags: behavior trees, LLM, Unreal, game AI.

- **Problem and contribution:** LLMs generate behavior trees for game AI in Unreal Engine.
- **Relevance:** Behavior-tree representation for unit behavior.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### hdpcg26

**HDPCG26. High Dimensional Procedural Content Generation**

Kaijie Xu and Clark Verbrugge. 2026. FDG publication DOI found through forward citation search: [10.1145/3815598.3815606](https://doi.org/10.1145/3815598.3815606). [arXiv:2602.18943](https://arxiv.org/abs/2602.18943), first posted February 21. Direct representation and generation paper. Tags: mechanics, temporal dynamics, reachability, extensibility.

- **Problem and contribution:** makes nongeometric gameplay dimensions explicit in a joint state space. It studies additional spatial/layer dimensions and temporal dynamics, with reachability checks, time-expanded graphs, and generation/evaluation pipelines. The abstract reports experiments and playable Unity cases.
- **Relevance:** directly challenges the assumption that game content can be assessed only as static geometry.
- **Scope inference:** the two instantiated directions do not demonstrate representation of every novel mechanic.
- **Inspection:** primary arXiv abstract and OpenAlex DOI match; full publication metadata, methods, and results require inspection. Preprint and publication are one retained work.

### tdprofile26

**TDPROFILE26. td-profile: A Profile Validator and Reusable Schemas for Tower-Defense Game Data**

Kyle DerZweite (mardwerk). 2026. [GitHub: mardwerk/td-profile](https://github.com/mardwerk/td-profile), v1.0.2; MIT license. Non-paper resource (software artifact). Tags: tower-defense, profile schema, validation, constraint specification, reusable game data, progression rules.

- **Problem and contribution:** provides a copyable, self-contained **Profile** for one TD game that combines reusable JSON schemas with per-game settings (collection paths, model bindings, required fields, references, progression limits, numerical units, score weights). A validator checks supplied game-data records against the Profile *without rewriting them*; `score-tower` scores an individual tower family's contract compliance (100/100 = complete) — a score measures declared-data compliance, not gameplay balance or simulation.
- **Method and evaluation:** the Profile documents are `manifest.json` (interface version + document registry), `collections.json` (collections, layouts, scopes), `mechanics.json` (mechanic catalog: model types, roles, canonical field mappings, required fields), `rules.json` (progression operations `purchaseProgression` and `linearProgression`, with bounded state-space exploration limits of 100k states / 64 paths), `references.json` (reference resolution), `numerical-units.json` (typed units: credits, seconds, distance, count + field bindings), `scoring.json` (category weights). The validator (Go, ~6k LOC in `atlasvalidate/`, MIT, extracted from btd6-atlas) implements schema-field, required-field, layout, reference, progression-rule, unit-binding, and model-coverage checks; an optional `model-contracts` document enables strict source contracts (exact raw field shapes). Known types: `tower`, `upgrade`, `purchase`, `attack`, `weapon`, `projectile`, `damage`, `ability`, `enemy`, `map`.
- **Relevance:** a concrete, production-grade instance of the "formal, composable mechanic representation" Layer 2 calls for (MECHGEN14-style action-like effects: `attack`/`weapon`/`projectile`/`damage`/`ability` with `sourceUpgrade` references). It closes the ICA23/GENCON24 gap by supplying the candidate vocabulary that layer assumes but does not provide. Its typed quantities map naturally onto PCGBENCH25-style per-game metrics: placement cost → credits, range → distance, cooldown/interval → seconds, damage amount → count. It is the natural game-data representation layer between Layer 1 grounding (dossier) and Layer 2 generation (the generator emits `Towers/<familyId>/<id>.json` and `Upgrades/<id>.json` conforming to the Profile).
- **Scope inference:** scoped to tower defense; the model catalog is a reasonable TD skeleton but would need extension for RTS-style abilities or mechanics beyond TD lanes. Passing a score establishes conformance, not balance, fun, or character faithfulness.
- **Limitation:** the Profile must still be authored for the target game; the starter template covers a generic game with one tower path through tiers 0–2. It does not claim to validate every source field unless strict source contracts are supplied.
- **Inspection:** full README.md, AGENTS.md, docs/PROFILE.md, docs/SETUP.md, docs/RELEASE.md, NOTICE.md, LICENSE; complete `profile/` directory (manifest, collections, mechanics, rules, references, numerical-units, scoring) plus `profile/README.md`; `examples/minimal-game/` (Bolt tower states + Power upgrades, including an embedded ability with `sourceUpgrade`); scanned validator core (`validate.go`, `rules.go`, `scoring.go`, `selection.go`, `profile.go`, `model_contracts.go` and their tests); schema catalog enumerated (`game-data/`, `profile/`, `scopes/` schemas). Does not include a full test-suite walkthrough, and the BTD6 Profile is excluded from the repo per NOTICE.md.

## Character understanding and fidelity

### characters03

**CHARACTERS03. Characters in Computer Games: Toward Understanding Interpretation and Design**

Petri Lankoski, Satu Heliö, and Inger Ekman. 2003. *Proceedings of DiGRA 2003 Conference: Level Up*. DOI: [10.26503/dl.v2003i1.101](https://doi.org/10.26503/dl.v2003i1.101). [Primary paper](https://dl.digra.org/index.php/dl/article/download/101/101). Direct conceptual and design-analysis paper. Tags: character interpretation, adaptation, mechanics, goals, progression.

- **Problem and contribution:** examines how a playable protagonist's character is expressed through predefined behavior, goals, possible and impossible actions, and characterization.
- **Method:** qualitative analysis of existing games and a proposed design approach drawing on dramatic writing. The sections "Possible and impossible actions" and "The Protagonist and Game Mechanics" explicitly analyze translating Hulk and Bruce Banner's contrasting abilities into attack strength, object interaction, stealth, and temporary stronger attacks after Hulk receives damage. The design discussion also considers weaknesses, motives, and skill development, including how progression can change later choices in *Deus Ex*.
- **Source report:** the authors argue that game mechanics affect how a character is interpreted and provide concrete examples of character aspects translated into mechanics. These are analytical examples and design arguments, not measured fidelity gains, a player study, or an automated-generation result.
- **Relevance:** direct prior discussion of adapting established characters into gameplay, extending this library beyond dialogue and profile benchmarks.
- **Scope inference:** it does not demonstrate reliable automated translation, Tower Defense unit adaptation, or a validated character-fidelity measure. Its proposed character dimensions are one historical design approach, not this project's representation.
- **Inspection:** complete primary article text, including examples, design proposal, conclusion, and references; Crossref metadata checked on 2026-09-28. Crossref's 2025 deposit date is not the 2003 publication date.

### npc07

**NPC07. Gameplay Design Patterns for Believable Non-Player Characters**

Petri Lankoski and Staffan Björk. 2007. *Situated Play, Proceedings of DiGRA 2007 Conference*, pp. 416-423. DOI: [10.26503/dl.v2007i1.262](https://doi.org/10.26503/dl.v2007i1.262). [Primary paper](https://dl.digra.org/index.php/dl/article/download/262/262). Direct qualitative game-analysis paper. Tags: character behavior, believability, social interaction, design goals.

- **Problem and contribution:** derives gameplay design patterns from failures and successes of believable behavior in an analysis of Claudette Perrick in *The Elder Scrolls IV: Oblivion*.
- **Method:** play and inspection of interactions, interpreted through theoretical requirements for perceiving intentional agents. The patterns include initiative, an own agenda, emotional attachment, contextual conversation, and personal development driven by goals. The introduction explicitly excludes a study of how players interpret the game. The conclusion cautions that adding the identified patterns would not necessarily improve the playing experience, because believability can conflict with particular gameplay goals, and describes the analysis as an initial survey with incompletely documented patterns.
- **Relevance:** character behavior and social roles can have mechanical consequences beyond combat or visual appearance.
- **Scope inference:** believability within a game and fidelity to an external reference character are different questions; this qualitative case study validates neither a universal checklist nor a scoring instrument.
- **Inspection:** primary PDF abstract, introduction, background, method, selected character-analysis passages, and conclusions, pp. 416-422; publication metadata checked in Crossref on 2026-09-28.

### thon19

**THON19. Transmedia characters: Theory and analysis**

Jan-Noël Thon. 2019. *Frontiers of Narrative Studies* 5(2):176-199. DOI: [10.1515/fns-2019-0012](https://doi.org/10.1515/fns-2019-0012). [Publication copy in the author's institutional repository](https://research.uca.ac.uk/5516/1/%5B25094890%20-%20Frontiers%20of%20Narrative%20Studies%5D%20Transmedia%20characters_%20Theory%20and%20analysis.pdf). Adjacent theoretical paper. Tags: adaptation, character identity, versions, canon, interpretation.

- **Problem and contribution:** proposes analyzing characters as they appear in particular works before treating versions across media as one character.
- **Method:** theoretical argument illustrated through Sherlock Holmes, Batman, and Lara Croft. In "Correlating characters," the author distinguishes redundancy, expansion, and modification. Modification introduces contradictions that prevent the representations from belonging to a single consistent storyworld. The conclusion stresses that a shared name does not establish shared character identity and that audience knowledge, authorship, and normative discourse affect how versions are related.
- **Relevance:** provides a substantive account of ambiguity in a character-name input and of the relationship between source versions.
- **Scope inference:** a fidelity judgment can depend on which version or relation is intended; this is a project implication of the theory, not an experimentally validated fidelity metric. The paper does not propose automatic unit design or show how to choose gameplay abilities. Its framework remains one theoretical perspective, not an adopted ontology.
- **Inspection:** primary publication abstract and opening discussion, "Correlating characters" pp. 187-191, and conclusion pp. 193-194; Crossref metadata checked on 2026-09-28. Other case details and the full reference list were not exhaustively analyzed.

### cross24

**CROSS24. Evaluating Character Understanding of Large Language Models via Character Profiling from Fictional Works**

Xinfeng Yuan, Siyu Yuan, Yuhan Cui, Tianhe Lin, Xintao Wang, Rui Xu, Jiangjie Chen, and Deqing Yang. 2024. *Conference on Empirical Methods in Natural Language Processing*. DOI: [10.18653/v1/2024.emnlp-main.456](https://doi.org/10.18653/v1/2024.emnlp-main.456). [arXiv:2404.12726](https://arxiv.org/abs/2404.12726). Adjacent benchmark paper. Tags: character extraction, profiles, narrative, fidelity, evaluation.

- **Problem and contribution:** evaluates character understanding through profile generation, using CroSS, a dataset of 126 profiles from novels. It evaluates factual consistency and downstream motivation recognition, comparing summarization methods and models. The inspected paper uses an LLM consistency evaluator, with a 50-sample human comparison reporting Pearson correlation 0.752. It separately evaluates multiple-choice motivation questions and ablates profile dimensions. Event information has the largest reported effect among the ablated dimensions.

Reported failures include character and relationship misidentification, omitted key information, misinterpreted events, and distorted characterization. The authors flag their four chosen dimensions, possible training-data exposure, and evaluator-model bias.

- **Relevance:** character understanding includes motives, relationships, events, and personality, which matters for the corpus's noncombat characters.
- **Scope inference:** good profiles do not establish faithful translation into gameplay, and its dimensions are not this project's schema.
- **Inspection:** primary abstract and selected evaluation, experiment, error-analysis, conclusion, and limitations sections; Crossref publication metadata.

### incharacter24

**INCHARACTER24. InCharacter: Evaluating Personality Fidelity in Role-Playing Agents through Psychological Interviews**

Xintao Wang, Yunze Xiao, Jen-tse Huang, Siyu Yuan, Rui Xu, Haoran Guo, Quan Tu, Yaying Fei, Ziang Leng, Wei Wang, Jiangjie Chen, Cheng Li, and Yanghua Xiao. 2024. *62nd Annual Meeting of the Association for Computational Linguistics*, Volume 1: Long Papers. DOI: [10.18653/v1/2024.acl-long.102](https://doi.org/10.18653/v1/2024.acl-long.102). [arXiv:2310.17976](https://arxiv.org/abs/2310.17976), first posted 2023. Adjacent evaluation paper. Tags: personality, role playing, fidelity.

- **Problem and contribution:** evaluates personality fidelity through interviews tied to psychological scales, complementing tests of character knowledge and speaking style. Experiments cover 32 characters and 14 scales; the abstract reports up to 80.7% accuracy against human-perceived personalities.
- **Relevance:** demonstrates that “character fidelity” has multiple operational definitions.
- **Scope inference:** dialogue-based personality assessments cannot be assumed valid for abilities, behavior rules, or upgrade designs. Scale construction, annotator agreement, and the precise accuracy denominator require full-text inspection before reuse.
- **Inspection:** primary abstract and Crossref publication metadata.

### austen24

**AUSTEN24. Evaluating Computational Representations of Character: An Austen Character Similarity Benchmark**

Funing Yang and Carolyn Jane Anderson. 2024. *4th International Conference on Natural Language Processing for Digital Humanities*. DOI: [10.18653/v1/2024.nlp4dh-1.3](https://doi.org/10.18653/v1/2024.nlp4dh-1.3). [arXiv:2408.16131](https://arxiv.org/abs/2408.16131). Adjacent benchmark paper. Tags: social roles, literary interpretation, representation.

- **Problem and contribution:** AustenAlike evaluates character representations through structural, social, and literary-expert similarity groupings.
- **Method:** compare features from BookNLP and FanfictionNLP with benchmark groupings and GPT-4 rankings. The abstract reports that broad social and narrative similarities are captured but expert-defined similarities challenge all tested systems.
- **Relevance:** shows a specific difficulty in representing literary character similarity beyond broad social categories.
- **Scope inference:** results are restricted to Austen novels and a similarity task; transfer to anime, manga, or gameplay design is unestablished.
- **Inspection:** primary abstract and Crossref publication metadata.

### rolellm24

**ROLELLM24. RoleLLM: Benchmarking, Eliciting, and Enhancing Role-Playing Abilities of Large Language Models**

Noah Wang, Z.Y. Peng, Haoran Que, Jiaheng Liu, Wangchunshu Zhou, Yuhan Wu, Hongcheng Guo, Ruitong Gan, Zehao Ni, Jian Yang, Man Zhang, Zhaoxiang Zhang, Wanli Ouyang, Ke Xu, Wenhao Huang, Jie Fu, and Junran Peng. 2024. *Findings of the Association for Computational Linguistics: ACL 2024*, pp. 14743-14777. DOI: [10.18653/v1/2024.findings-acl.878](https://doi.org/10.18653/v1/2024.findings-acl.878). [Primary ACL record](https://aclanthology.org/2024.findings-acl.878/). Adjacent benchmark and method paper. Tags: fictional characters, character knowledge, style, fine-tuning, evaluation.

- **Problem and contribution:** constructs role profiles for 100 characters, generates role-specific instructions and style examples, and uses them to build RoleBench with 168,093 samples. The abstract describes role-conditioned instruction tuning and reports improved role-playing performance for the resulting models. Method and evaluation details beyond those stated in the abstract were not inspected.
- **Relevance:** separates character knowledge, speech style, and model conditioning as measurable concerns.
- **Scope inference:** dialogue role-playing scores do not establish gameplay-unit fidelity, mechanical coherence, or coverage of non-dialogue narrative roles.
- **Inspection:** primary ACL publication metadata and abstract, checked 2026-09-27.

### omnic26

**OMNIC26. OMNICharacter++: Toward a Comprehensive Benchmark for Realistic Role-Playing Agents**

Authors (metadata only via Crossref). 2026. *IEEE Transactions on Pattern Analysis and Machine Intelligence*. Direct benchmark paper. Tags: role-playing, character benchmark, fidelity, realism.

- **Problem and contribution:** Extends the OMNICharacter benchmark toward a comprehensive evaluation of realistic role-playing agents.
- **Method and evaluation:** Not fully inspected; abstract indicates an expanded benchmark for role-playing agent realism.
- **Relevance:** A recent, venue-strong benchmark for character-role realism; relevant to establishing a fidelity evaluation for character-grounded units.
- **Scope inference:** Dialogue role-playing agents, not gameplay units; realism ≠ balance.
- **Inspection:** abstract-level via Crossref/OpenAlex; TPAMI venue (verify).

### llmnpccmp26

**LLMNPCCMP26. Comparing LLM-Driven and Script-Based Non-Player Characters Under Controlled Information**

Authors (metadata only via Crossref). 2026. *Applied Sciences* 16(14):7254. Direct method paper. Tags: NPC comparison, LLM-driven, script-based, controlled information.

- **Problem and contribution:** Compares LLM-driven versus script-based NPC behavior under controlled information conditions.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** A controlled comparison framework for character-grounded agents; directly useful as a template for how to evaluate a character-grounded unit against a scripted baseline.
- **Scope inference:** NPC dialogue/behavior, not unit combat design.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### fictionrag26

**FICTIONRAG26. FictionRAG: A Stateful Metacognitive Framework for High-Fidelity Long-Narrative Role-Playing**

Authors (metadata only via Crossref). 2026. *Algorithms* 19(5):383. Direct method paper. Tags: RAG, role-playing, fidelity, stateful memory, metacognitive.

- **Problem and contribution:** A stateful, metacognitive retrieval-augmented generation framework for high-fidelity long-narrative role-playing; reduces character drift over long interactions.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Character drift is a risk for long-form unit explanations and for maintaining character consistency across generated upgrades; RAG with stateful memory is a concrete mitigation strategy for the knowledge-acquisition question.
- **Scope inference:** Narrative role-play, not gameplay mechanics.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### dream26

**DREAM26. DREAM: LLM-based Dynamic Role-playing via Event-Aware Memory Graph**

Authors (metadata only via Crossref). 2026. DOI: [10.1145/3770855.3818027](https://doi.org/10.1145/3770855.3818027). Direct method paper. Tags: role-playing, event-aware memory graph, LLM, dynamic role-play.

- **Problem and contribution:** Dynamic role-playing via an event-aware memory graph built around LLM agents.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Event-aware memory is the kind of structured character knowledge (events, relationships) that CROSS24 shows matters most for character understanding — relevant to building the character dossier.
- **Scope inference:** Dialogue role-play, not unit design.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### baijia26

**BAIJIA26. BaiJia: An Open Role-Playing Platform of Chinese Historical Characters**

Authors (metadata only via Crossref). 2026. DOI: [10.1145/3774905.3793125](https://doi.org/10.1145/3774905.3793125). Adjacent system/paper. Tags: role-playing platform, historical characters, open platform.

- **Problem and contribution:** An open platform for role-playing as Chinese historical characters, grounding characters in historical source material.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Demonstrates grounding named/historical characters in source material for behavior — analogous to grounding named fictional characters in fandom sources for this project.
- **Scope inference:** Platform/system, not unit generation; historical characters, not fictional IPs.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### questgram21

**QUESTGRAM21. Questgram [Qg]: Toward a Mixed-Initiative Quest Generation Tool**

Authors (metadata only via Crossref). 2021. DOI: [10.1145/3472538.3472544](https://doi.org/10.1145/3472538.3472544). Direct mixed-initiative system paper. Tags: quest generation, mixed-initiative, interaction.

- **Problem and contribution:** A mixed-initiative tool for quest generation that supports human-AI collaboration in designing quests.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Mixed-initiative design is a candidate interface for eliciting the initially undefined constraints; quest design overlaps with unit/ability design goals.
- **Scope inference:** Quests, not units; mixed-initiative interface, not autonomous generation.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### trpgpcg20

**TRPGPCG20. Tabletop Roleplaying Games as Procedural Content Generators**

Authors (metadata only via Crossref). 2020. *IEEE Conference on Games*. DOI: [10.1145/3402942.3409605](https://doi.org/10.1145/3402942.3409605). Direct method paper. Tags: TTRPG, procedural content generation, character creation.

- **Problem and contribution:** Treats tabletop roleplaying games themselves as procedural content generators, generating content (characters, encounters) through TTRPG rules.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** TTRPG character/encounter generation is a close analogue to generating character-grounded game units from defined traits.
- **Scope inference:** TTRPG content, not video-game TD/RTS units; character-grounded via TTRPG stat blocks.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### rpgcreature23

**RPGCREATURE23. RPG Creature Design: Cross-System Analysis and Conversion**

Authors (metadata only via Crossref). 2023. *Proceedings of the Genetic and Evolutionary Computation Conference Companion*. DOI: [10.1145/3631085.3631332](https://doi.org/10.1145/3631085.3631332). Direct method paper. Tags: creature design, cross-system conversion, TTRPG.

- **Problem and contribution:** Analyzes and converts RPG creature designs across different game systems — cross-system character/creature adaptation.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Direct analogue of the cross-game adaptation problem: mapping a creature/character to a new system's mechanical vocabulary.
- **Scope inference:** TTRPG systems, not TD/RTS units; conversion, not generation from scratch.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### ignite23

**IGNITE23. "It Has to Ignite Their Creativity": Opportunities for Generative Tools for Game Masters**

Authors (metadata only via Crossref). 2023. *Proceedings of the 18th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3582437.3587204](https://doi.org/10.1145/3582437.3587204). Adjacent HCI/paper. Tags: game masters, generative tools, HCI, creativity.

- **Problem and contribution:** Qualitative study of opportunities for generative AI tools for tabletop game masters.
- **Method and evaluation:** Not fully inspected; qualitative study.
- **Relevance:** Game-master tools are a form of character/NPC design support under ambiguous constraints; user needs from this study inform the mixed-initiative interface.
- **Scope inference:** Tabletop game masters, not video games; qualitative, not automated.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### fighting25

**FIGHTING25. A Survey on the Development of Human-Centered Artificial Agents for Fighting Games**

Authors (metadata only via Crossref). 2025. *IEEE Transactions on Games*. DOI: [10.1109/TG.2025.3577552](https://doi.org/10.1109/TG.2025.3577552). Direct survey. Tags: fighting games, human-centered agents, survey, player modeling.

- **Problem and contribution:** Survey of human-centered artificial agents for fighting games — player models and opponent design grounded in human behavior.
- **Method and evaluation:** Survey; not fully inspected.
- **Relevance:** Human-centered agent design for a fast-paced combat genre is adjacent to evaluating TD/RTS units against human-like players; informs the evaluation-agent design.
- **Scope inference:** Fighting-game agents, not unit generation.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### charhalluc26

**CHARHALL26. CHARM: Character Hallucination for Multicultural Role Play Benchmark**

Authors (metadata only via arXiv API). 2026. arXiv:2609.01352v1. Adjacent benchmark paper. Tags: character hallucination, role play, benchmark, boundary awareness, multicultural.

- **Problem and contribution:** Distinguishes whether character hallucination errors arise from failure to recognize a knowledge boundary vs. failure to comply despite recognition.
- **Method and evaluation:** Introduces CHARM, a multicultural benchmark of 40 real and fictional characters from five cultural-linguistic regions, validated by native reviewers; probes Temporal and Cross-Universe boundary types with abstention-enabled multiple-choice questions; two-stage evaluation separating Boundary-Awareness from Boundary-Compliance. Evaluations across six LLMs show hallucination is driven predominantly by compliance failures (parametric overrides).
- **Relevance:** Boundary-awareness vs. compliance failure is directly analogous to a unit claiming an ability the character does not have (hallucinated ability) vs. omitting one it does have; the two-stage protocol maps onto a fidelity test.
- **Scope inference:** Dialogue role-play knowledge boundaries, not gameplay abilities; benchmark, not a unit-generator.
- **Inspection:** primary arXiv abstract.

### ooc25

**OOC25. Spotting Out-of-Character Behavior: Atomic-Level Evaluation of Persona Fidelity in Open-Ended Generation**

Authors (metadata only via arXiv API). 2025. arXiv:2506.19352v1. Adjacent evaluation paper. Tags: out-of-character, persona fidelity, atomic evaluation, granularity.

- **Problem and contribution:** Existing persona-fidelity evaluations assign single scores to whole responses and miss subtle in-consistencies in long-form generation.
- **Method and evaluation:** Proposes an atomic-level evaluation framework quantifying persona fidelity with three key metrics measuring persona alignment and consistency within and across generations; demonstrates the framework detects inconsistencies prior methods overlook; analyzes how task structure and persona desirability influence adaptability.
- **Relevance:** Atomic-level fidelity evaluation is a concrete template for testing whether a generated unit consistently expresses the character across its abilities, stats, and upgrades — exactly the consistency check this project needs.
- **Scope inference:** Persona fidelity in text generation, not gameplay units; metrics not validated for mechanics.
- **Inspection:** primary arXiv abstract.

### mimic22

**MIMIC22. Meet Your Favorite Character: Open-domain Chatbot Mimicking Fictional Characters with only a Few Utterances**

Seungju Han, Beomsu Kim, Jin Yong Yoo, Seokjun Seo, Sangbum Kim, Enkhbayar Erdenee, and Buru Chang. 2022. *Proceedings of the 2022 Conference of the North American Chapter of the Association for Computational Linguistics: Human Language Technologies*, pp. 5114-5132. DOI: [10.18653/v1/2022.naacl-main.377](https://doi.org/10.18653/v1/2022.naacl-main.377). [Primary ACL record](https://aclanthology.org/2022.naacl-main.377/). Adjacent method paper. Tags: fictional characters, few-shot evidence, dialogue style, retrieval.

- **Problem and contribution:** addresses mimicking a fictional character's speech when only a few utterances are available. Pseudo Dialog Prompting places retrieved character utterances into a constructed dialogue context, and the abstract reports both human and automatic evaluation against baselines for stylistic resemblance. Detailed samples, comparison scores, and limitations were not inspected.
- **Relevance:** shows one evaluated form of character-conditioned behavior under limited source material.
- **Scope inference:** resemblance of chatbot responses cannot establish the quality or faithfulness of abilities, statistics, upgrades, or gameplay explanations.
- **Inspection:** primary ACL publication metadata and abstract, checked 2026-09-27.

## Grounded knowledge and acquiring constraints

### rag20

**RAG20. Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks**

Patrick Lewis, Ethan Perez, Aleksandra Piktus, Fabio Petroni, Vladimir Karpukhin, Naman Goyal, Heinrich Küttler, Mike Lewis, Wen-tau Yih, Tim Rocktäschel, Sebastian Riedel, and Douwe Kiela. 2020. arXiv manuscript inspected; NeurIPS publication not independently checked in this pass. [arXiv:2005.11401](https://arxiv.org/abs/2005.11401). Adjacent method paper. Tags: retrieval, knowledge updating, generation, provenance.

- **Problem and contribution:** combines parametric generation with retrieval from a nonparametric Wikipedia index. It compares retrieval fixed across a sequence with retrieval that can vary by token, fine-tunes on knowledge-intensive tasks, and reports gains on open-domain QA and more factual generated text than a parametric-only comparator.
- **Relevance:** directly addresses the relationship between stored knowledge, updateable evidence, and generation, one of the project's unresolved choices.
- **Scope inference:** results on these NLP tasks do not guarantee complete or canon-sensitive fictional-character research, and retrieved passages do not automatically validate a gameplay interpretation.
- **Inspection:** primary arXiv abstract. No recommendation to adopt retrieval follows from retention.

### ica23

**ICA23. Learning to Learn in Interactive Constraint Acquisition**

Dimos Tsouros, Senne Berden, and Tias Guns. 2024. *Proceedings of the AAAI Conference on Artificial Intelligence*, 38(8). DOI: [10.1609/aaai.v38i8.28655](https://doi.org/10.1609/aaai.v38i8.28655). [arXiv:2312.10795](https://arxiv.org/abs/2312.10795), first posted 2023. Adjacent method paper. Tags: constraint discovery, elicitation, active learning.

- **Problem and contribution:** reduces the number of user queries needed to learn a constraint model.
- **Method:** statistical classifiers predict whether expressions from a candidate bias are constraints and guide query generation, scope finding, and constraint finding. The abstract reports reductions in queries of up to 72%.
- **Relevance:** formal research on learning requirements that are not supplied as a complete model.
- **Scope inference:** the method operates over a candidate language or bias and queryable judgments; this is different from discovering entirely new gameplay concepts or ambiguous preferences.
- **Inspection:** primary abstract and Crossref publication metadata; detailed benchmark composition and failure cases not inspected. A related 2026 JAIR paper by the same authors, DOI [10.1613/jair.1.19524](https://doi.org/10.1613/jair.1.19524), appears to extend this work, but its full text was inaccessible; its precise differences remain unchecked and it is not counted as a separate retained study.

### gencon24

**GENCON24. Generalizing Constraint Models in Constraint Acquisition**

Dimos Tsouros, Senne Berden, Steven Prestwich, and Tias Guns. 2024 manuscript; subsequent venue not verified. [arXiv:2412.14950](https://arxiv.org/abs/2412.14950). Adjacent method paper. Tags: reusable constraints, generalization, interpretable models.

- **Problem and contribution:** GenCon learns parameterized models rather than a single instance's set of constraints.
- **Method:** classify candidate constraints under parameterizations, extract decision rules for suitable classifiers, or use generate-and-test for other classifiers. The abstract reports high accuracy and robustness to noisy input instances.
- **Relevance:** distinguishes reuse across instances from repeatedly acquiring isolated rule sets.
- **Scope inference:** generalizing instances of the same problem does not establish cross-game transfer or extension of the underlying mechanic vocabulary.
- **Inspection:** primary arXiv abstract; exact datasets, accuracy figures, and generalization boundaries require fuller reading.

## Evaluation, factuality, and automated playtesting

### csi14

**CSI14. Quantifying the Creativity Support of Digital Tools through the Creativity Support Index**

Erin Cherry and Celine Latulipe. 2014. *ACM Transactions on Computer-Human Interaction* 21(4):1-25. DOI: [10.1145/2617588](https://doi.org/10.1145/2617588). Adjacent instrument-development paper. Tags: designer evaluation, creative tools, measurement, human judgment.

- **Problem and contribution:** introduces the Creativity Support Index, a psychometric survey of how a tool supports a user engaged in creative work. The abstract names six dimensions: Exploration, Expressiveness, Immersion, Enjoyment, Results Worth Effort, and Collaboration. It describes iterative development and validation, deployment scenarios, and use of scores to identify aspects of tool support needing attention. Validation samples, reliability, scoring procedures, and detailed results have not been inspected.
- **Relevance:** a concrete source on evaluating a creative tool's support for its users, complementing measures of generated content.
- **Scope inference:** such support is a different construct from character fidelity, game balance, or originality of a unit, and this abstract does not establish validity for this project's designers or tasks. Retention does not select the instrument.
- **Inspection:** abstract deposited in Crossref and exact publication metadata on 2026-09-28. The primary ACM page returned HTTP 403; full text remains unread.

### pcgbench25

**PCGBENCH25. The Procedural Content Generation Benchmark: An Open-source Testbed for Generative Challenges in Games**

Ahmed Khalifa, Roberto Gallotta, Matthew Barthet, Antonios Liapis, Julian Togelius, and Georgios N. Yannakakis. 2025. *Proceedings of the 20th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3723498.3723794](https://doi.org/10.1145/3723498.3723794). [arXiv:2503.21474](https://arxiv.org/abs/2503.21474). Direct benchmark paper. Tags: PCG, benchmark, quality, diversity, controllability, validity.

- **Problem and contribution:** offers a common interface for 12 game-content problems and variants spanning levels, structures, and simple arcade rules. Each problem defines its own representation, controls, and functions for quality, diversity, and controllability.
- **Method and evaluation:** compares random search, an evolution strategy, and a genetic algorithm under several fitness combinations across ten runs per condition. The paper reports that difficulty and algorithm tradeoffs vary by problem and objective.
- **Relevance:** a concrete comparison framework that makes domain-specific metrics explicit instead of implying one universal score.
- **Scope inference:** benchmark success within its predefined problems does not measure character fidelity or adaptability to an unrepresented ability.
- **Inspection:** primary PDF abstract, benchmark interface, experiment and result sections, and conclusion; Crossref publication record, checked 2026-09-27. The accompanying repository and licenses were not inspected.

### factscore23

**FACTSCORE23. FActScore: Fine-grained Atomic Evaluation of Factual Precision in Long Form Text Generation**

Sewon Min, Kalpesh Krishna, Xinxi Lyu, Mike Lewis, Wen-tau Yih, Pang Wei Koh, Mohit Iyyer, Luke Zettlemoyer, and Hannaneh Hajishirzi. 2023. arXiv manuscript inspected; publication venue not checked in this pass. [arXiv:2305.14251](https://arxiv.org/abs/2305.14251). Adjacent evaluation paper. Tags: evidence, factuality, character knowledge.

- **Problem and contribution:** evaluates mixed factual and unfactual long-form output by decomposing it into atomic claims and measuring the fraction supported by a reliable source.
- **Method and evaluation:** human annotation of generated biographies and an automated retrieval-plus-LM estimator, later applied to 6,500 generations from 13 models. The abstract reports less than 2% estimation error in the studied setting.
- **Relevance:** provides a concrete distinction between plausible prose and supported character claims.
- **Scope inference:** precision alone does not measure omitted facts, interpretive adequacy, or gameplay fidelity; the source's coverage matters.
- **Inspection:** primary abstract. These reported figures must not be generalized to anime/manga dossiers without validation.

### mcts15

**MCTS15. Monte-Carlo Tree Search for Simulation-based Strategy Analysis**

Alexander Zook, Brent Harrison, and Mark O. Riedl. 2015. *Foundations of Digital Games*, as identified on the primary manuscript. [arXiv:1908.01423](https://arxiv.org/abs/1908.01423), deposited August 4, 2019. [Primary manuscript HTML](https://arxiv.org/html/1908.01423v1). DOI not verified. Direct game-analysis paper. Tags: automated playtesting, player skill, balance, agent limitations.

- **Problem and contribution:** uses MCTS agents with different computation budgets to inspect gameplay strategies and potential design issues at four levels: summaries, individual actions, action chains, and action spaces.
- **Method and evaluation:** tests a modified *Scrabble* with perfect information, a full dictionary, and a 150-point target, and *Cardonomicon* with perfect information and identical fixed 20-card decks. Each game uses three rollout budgets and 100 games per agent pairing. The authors find no first-turn advantage in their modified *Scrabble*, a substantial disadvantage for the second player in *Cardonomicon*, and no significant action combinations in the latter. Their explanation in terms of missing card synergy is an interpretation of those observations.

The limitations section states that rollout budget models deliberation rather than perception, motor skill, or memory; the tested domains are fully observable. Extending this to real-time games would require deciding what proxy represents player skill.

- **Relevance:** a primary source behind automated unit-evaluation work that helps qualify what differing MCTS budgets measure.
- **Scope inference:** budget-dependent behavior does not validate the agents as human skill groups or establish a universal balance measure.
- **Inspection:** primary manuscript abstract, methods section 3, metrics section 4, case studies section 5, and limitations section 6 on 2026-09-28. The 2015 paper and 2019 arXiv deposit are one retained work. Crossref's title lookup returned unrelated records; these were rejected. The HTML contains a merge-conflict artifact in the Figure 7 caption, so that caption was not used as evidence.

### personas19

**PERSONAS19. Automated Playtesting With Procedural Personas Through MCTS With Evolved Heuristics**

Christoffer Holmgård, Michael Cerny Green, Antonios Liapis, and Julian Togelius. 2019 issue publication, 2018 early record. *IEEE Transactions on Games*. DOI: [10.1109/TG.2018.2808198](https://doi.org/10.1109/TG.2018.2808198). [arXiv:1802.06881](https://arxiv.org/abs/1802.06881). Direct evaluation method. Tags: player modeling, simulation, playstyles.

- **Problem and contribution:** creates synthetic playtesters with different playstyles.
- **Method:** evolve MCTS node-selection heuristics for procedural personas and demonstrate different behavior across a corpus of levels.
- **Relevance:** shows why a single strong agent need not represent all relevant player behavior.
- **Scope inference:** enacted styles and fast automated feedback do not establish that the personas faithfully reproduce a particular human population. Detailed human validation was not inspected.
- **Inspection:** primary arXiv abstract and Crossref publication metadata.

### npcreal25

**NPCREAL25. An Empirical Evaluation of AI-Powered Non-Player Characters' Perceived Realism and Performance**

Authors (metadata only via Crossref). 2025. SSRN Electronic Journal. DOI: [10.2139/ssrn.5148461](https://doi.org/10.2139/ssrn.5148461). Direct empirical paper. Tags: NPC realism, human evaluation, perceived performance, empirical.

- **Problem and contribution:** Empirical study of how users perceive the realism and performance of AI-powered NPCs.
- **Method and evaluation:** Empirical human evaluation; details not fully inspected.
- **Relevance:** Perceived realism of AI-generated behavior is a human-judgment metric relevant to the "explanation" and "faithfulness" dimensions of a character-grounded unit.
- **Scope inference:** NPC dialogue/behavior realism, not unit design; SSRN preprint (verify publication venue).
- **Inspection:** abstract-level via Crossref/OpenAlex. Access note: SSRN returned 403 in the prior pass; full text not inspected.

### explai26

**EXPLAI26. Transparent and Adaptive Gaming: The Role of Explainable AI and Dynamic Difficulty**

Authors (metadata only via Crossref). 2026. *ACM Computing Surveys*. DOI: [10.1145/3841631](https://doi.org/10.1145/3841631). Direct survey. Tags: explainable AI, dynamic difficulty, survey, adaptive gaming.

- **Problem and contribution:** Survey of explainable AI and dynamic difficulty adjustment in games.
- **Method and evaluation:** Survey; not fully inspected.
- **Relevance:** Explanation of AI-generated game content (the "explanation of character fidelity" output) is explicitly discussed; dynamic difficulty adjustment is relevant to the "progression and upgrades" topic noted as thin in the library.
- **Scope inference:** Broad XAI/DDA coverage, not character-grounded units.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### onboarding25

**ONBOARD25. Support Autonomy: Exploring Player Perspectives on AI-Supported Onboarding in Video Games**

Authors (metadata only via Crossref/OpenAlex). 2025. DOI: [10.1145/3706598.3713576](https://doi.org/10.1145/3706598.3713576). Adjacent HCI paper. Tags: player perspectives, AI-supported onboarding, HCI, player experience.

- **Problem and contribution:** Explores player perspectives on AI-supported onboarding in video games (details not fully inspected).
- **Method and evaluation:** Not fully inspected; title/abstract indicate a player-experience study of AI-supported onboarding.
- **Relevance:** The "explanation" output of a character-to-unit generator must be understandable to players; player perspectives on AI support inform how to present generated content and its rationale.
- **Scope inference:** Onboarding, not unit design; player perspectives, not generation quality.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### dynamicnpc25

**DYNAMICNPC25. Leveraging Large Language Models for Dynamic NPC Interactions in 2D RPGs**

Authors (metadata only via Crossref/OpenAlex). 2025. *IEEE Conference on Games*. DOI: [10.1109/SEEDA-CECN68644.2025.11329608](https://doi.org/10.1109/SEEDA-CECN68644.2025.11329608). Direct method paper. Tags: NPC interactions, LLM, 2D RPG, dynamic.

- **Problem and contribution:** Uses LLMs to generate dynamic NPC interactions in 2D RPGs.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** LLM-driven NPC behavior in a 2D game setting is adjacent to generating dynamic behavior for generated units; provides a concrete integration point for LLM + game engine.
- **Scope inference:** NPC dialogue/interaction, not unit combat or balance.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### genai3d26

**GENAI3D26. A deep generative approach to personalized super mario level design**

Authors (metadata only via arXiv API / Scientific Reports). 2026. DOI: [10.1038/s41598-026-46199-1](https://doi.org/10.1038/s41598-026-46199-1). Direct method paper. Tags: generative AI, personalized level design, Mario, deep learning.

- **Problem and contribution:** Personalized level design for Super Mario Bros. using a deep generative approach.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Personalization is analogous to tailoring a generated unit to a specific character; demonstrates deep generative methods applied to an existing game engine's content.
- **Scope inference:** Single-game level personalization; not character grounding.
- **Inspection:** abstract-level via arXiv/OpenAlex.

### mixedinit25

**MIXEDINIT25. Boosting Mixed-Initiative Co-Creativity in Game Design: A Tutorial**

Authors (metadata only via Crossref/OpenAlex). 2025. *ACM Computing Surveys*. DOI: [10.1145/3759914](https://doi.org/10.1145/3759914). Direct tutorial/survey. Tags: mixed-initiative, co-creativity, game design, tutorial.

- **Problem and contribution:** Tutorial/survey on boosting mixed-initiative co-creativity in game design.
- **Method and evaluation:** Tutorial synthesizing prior work; not fully inspected.
- **Relevance:** Consolidates the mixed-initiative design literature into one accessible source — useful for designing the constraint-elicitation interface with the human designer.
- **Scope inference:** General game design co-creativity, not character grounding.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### script25

**SCRIPT25. ScriptDoctor: Automatic Generation of PuzzleScript Games via Large Language Models and Tree Search**

Authors (metadata only via Crossref/OpenAlex). 2025. *IEEE Conference on Games*. DOI: [10.1109/CoG64752.2025.11114269](https://doi.org/10.1109/CoG64752.2025.11114269). Direct method paper. Tags: PuzzleScript, LLM, tree search, automatic game generation.

- **Problem and contribution:** Generates PuzzleScript games via LLMs combined with tree search.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** LLM + tree search for game generation with a supplied grammar is directly analogous to grammar-guided generation (GGDG24) with an added search for validity — a concrete architecture candidate for ensuring generated mechanics are well-formed.
- **Scope inference:** PuzzleScript, not TD/RTS units; grammar supplied.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### seamless25

**SEAMLESS25. Seamless Tutorial: Contextual State Transition Generation Based on Player Internal Knowledge**

Authors (metadata only via Crossref/OpenAlex). 2025. *IEEE Transactions on Games*. DOI: [10.1109/TG.2025.3596055](https://doi.org/10.1109/TG.2025.3596055). Direct method paper. Tags: tutorial generation, state transitions, player knowledge, context.

- **Problem and contribution:** Generates seamless tutorials based on contextual state transitions and inferred player knowledge.
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Tutorial/instruction generation from internal state is analogous to generating explanations for a unit's design; player-model-based adaptation is a candidate for the "progression and upgrades" topic.
- **Scope inference:** Tutorial content, not unit generation; specific to tutorials.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### xgen26

**XGEN26. AI-directed procedural content generation for personalized XR in industrial-heritage museums**

Authors (metadata only via Crossref/OpenAlex). 2026. *Frontiers in Virtual Reality*. DOI: [10.3389/frvir.2026.1788711](https://doi.org/10.3389/frvir.2026.1788711). Direct method paper. Tags: AI-directed PCG, XR, personalization, industrial heritage.

- **Problem and contribution:** AI-directed PCG for personalized XR experiences in industrial-heritage museums (details not fully inspected).
- **Method and evaluation:** Not fully inspected.
- **Relevance:** Demonstrates AI-directed PCG for personalization outside games; the personalization pipeline may transfer.
- **Scope inference:** XR museum content, not games.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### genaiarch26

**GENAIARCH26. Computational game content generation and socially responsible design: Current practices, ...**

Authors (metadata only via Crossref/OpenAlex). 2026. *Computers & Graphics*. DOI: [10.1016/j.cag.2026.104683](https://doi.org/10.1016/j.cag.2026.104683). Direct review paper. Tags: game content generation, socially responsible design, review, ethics.

- **Problem and contribution:** Reviews current practices in computational game content generation with a focus on socially responsible design.
- **Method and evaluation:** Review; not fully inspected.
- **Relevance:** Socially responsible design considerations are relevant to the explanation/ethics dimension of character-grounded generation (avoiding harmful stereotypes in character-derived units).
- **Scope inference:** General content generation ethics, not character grounding.
- **Inspection:** abstract-level via Crossref/OpenAlex.

### era23

**ERA23. The Right Variety: Improving Expressive Range Analysis with Metric Selection Methods**

Oliver Withington and Laurissa Tokarchuk. 2023. *18th International Conference on the Foundations of Digital Games*. DOI: [10.1145/3582437.3582453](https://doi.org/10.1145/3582437.3582453). [arXiv:2304.02366](https://arxiv.org/abs/2304.02366). Direct evaluation paper. Tags: diversity, metric validity, visualization.

- **Problem and contribution:** examines how choosing two metrics affects expressive-range visualizations. The authors propose three quantifiable selection criteria and investigate their usefulness for ranking metric pairs. They describe the study as an early exploration and argue that metric choice should be justified.
- **Relevance:** supports caution when claiming that a generator covers a broad or novel design space from a projection.
- **Scope inference:** informative metric pairs do not establish semantic diversity, originality, or usefulness to designers.
- **Inspection:** primary abstract and Crossref publication metadata.

## Decision models and the RLCD acronym

### jev26

**JEV26. Just Ask Jev: Reinforcement Learning for Calibrated Decisions as a Zero-Shot Detector of AI Alignment Failures**

Ruoqi Guo, Yi Liu, Gelei Deng, Yuekang Li, Lida Zhao, Yutao Wu, Simin Chen, Ying Zhang, and Leo Yu Zhang. 2026. arXiv preprint, marked “under review as a conference paper at ICLR 2027” in the manuscript; acceptance is not established. [arXiv:2609.29429](https://arxiv.org/abs/2609.29429). Adjacent evaluation paper. Tags: typed decisions, calibration, language-model judges, source identity.

- **Problem and contribution:** evaluates whether TypeSafe AI's Jev, described as a model for typed probabilistic decisions, can detect language-model alignment failures. The paper introduces RLCDAlignBench with ten failure categories drawn from 44 benchmarks and five target models.
- **Method and evaluation:** varies question wording separately from supplied context and compares detection rankings, human-labeled subsets, and cost against other scorers. The authors report median AUROC 0.886 for a generic question over the applicable benchmark subset and lower inference cost in their comparison. These are study-specific findings, not validated game-design performance. The authors explicitly limit coverage to one Jev version, English benchmarks, 2B to 7B target models, and mostly scorer-derived labels.
- **Relevance:** independently confirms the identity of the Jev research lead and shows a possible class of probabilistic evaluation tasks, while leaving its training method and suitability for unit design untested.
- **Inspection:** primary arXiv abstract, introduction, figure description, conclusion, and limitations in the PDF, checked 2026-09-27.

### rlcd23

**RLCD23. RLCD: Reinforcement Learning from Contrastive Distillation for Language Model Alignment**

Kevin Yang, Dan Klein, Asli Celikyilmaz, Nanyun Peng, and Yuandong Tian. 2024. *International Conference on Learning Representations*, poster, verified in the [OpenReview publication record](https://openreview.net/forum?id=v3XXtxWKi6). [arXiv:2307.12950](https://arxiv.org/abs/2307.12950), first posted 2023. Adjacent alignment method paper. Tags: acronym disambiguation, preference learning, language-model alignment.

- **Problem and contribution:** proposes a way to construct preference pairs without direct human preference labels.
- **Method:** prompt a model in contrasting positive and negative ways, train a preference model from the resulting pairs, then use reinforcement learning to improve a base language model. The abstract reports comparisons with RLAIF and context distillation on harmlessness, helpfulness, and story-outline generation at two model scales. Detailed experimental setup and limitations were not inspected.
- **Relevance:** this is a different expansion of “RLCD” from TypeSafe's “Reinforcement Learning for Calibrated Decisions.” Its retention prevents the terminology collision from contaminating the Jev lead; its alignment results do not establish any unit-design method.
- **Inspection:** primary arXiv abstract and official OpenReview venue record, checked 2026-09-27.

## Research methodology

These sources concern research practice. They do not supply a selected unit-generation method. Associated official reporting resources are distinguished from inspected scholarly paper content.

### m01

**M01. Guidelines for conducting systematic mapping studies in software engineering: An update**

Kai Petersen, Sairam Vakkalanka, and Ludwik Kuzniarz. 2015. *Information and Software Technology* 64:1-18. DOI: [10.1016/j.infsof.2015.03.007](https://doi.org/10.1016/j.infsof.2015.03.007). Research methodology. Tags: systematic mapping, coverage, literature review.

- **Problem and contribution:** the verified publication record identifies methodological guidance for systematic mapping in software engineering. Substantive recommendations have not been inspected.
- **Method and evaluation:** not inspected; a bibliographic reference list does not establish the paper's own methods or findings.
- **Findings and limitations:** no substantive finding is attributed here. ScienceDirect returned HTTP 403, and OpenAlex supplied neither an abstract nor an accessible full-text location.
- **Relevance:** retained as a methodological source for assessing a future formal mapping protocol and the present review's coverage limits.
- **Inspection:** bibliographic record through Crossref and OpenAlex on 2026-09-27.

### m02

**M02. Guidelines for Snowballing in Systematic Literature Studies and a Replication in Software Engineering**

Claes Wohlin. 2014. *Proceedings of the 18th International Conference on Evaluation and Assessment in Software Engineering*, EASE '14. DOI: [10.1145/2601248.2601268](https://doi.org/10.1145/2601248.2601268). [Author manuscript](https://www.wohlin.eu/ease14.pdf). Research methodology. Tags: citation search, coverage, replication.

- **Problem and contribution:** inconsistent terminology and database selection can leave relevant work undiscovered. The paper specifies backward and forward snowballing and evaluates it through replication of a systematic review.
- **Method and evaluation:** inspected sections describe selecting a diverse start set, screening references and citing papers, resolving inclusion before using a paper as a seed, and keeping iterations traceable. The replication concerns cross-company and within-company software effort estimation; detailed replication results were not inspected.
- **Findings and limitations:** the abstract reports successful application. The discussion emphasizes imperfect search samples and difficult start-set selection, including communities that do not cite one another.
- **Relevance:** supports citation trails across this project's distinct research communities. Saturating one cluster cannot establish exhaustive coverage.
- **Inspection:** abstract, sections 1-3, and opening of section 4 in the author manuscript; Crossref metadata checked on 2026-09-27.

### m03

**M03. PRISMA Extension for Scoping Reviews (PRISMA-ScR): Checklist and Explanation**

Andrea C. Tricco, Erin Lillie, Wasifa Zarin, Kelly K. O'Brien, Heather Colquhoun, Danielle Levac, David Moher, Micah D. J. Peters, Tanya Horsley, Laura Weeks, Susanne Hempel, Elie A. Akl, Christine Chang, Jessie McGowan, Lesley Stewart, Lisa Hartling, Adrian Aldcroft, Michael G. Wilson, Chantelle Garritty, Simon Lewin, Christina M. Godfrey, Marilyn T. Macdonald, Etienne V. Langlois, Karla Soares-Weiser, Jo Moriarty, Tammy Clifford, Özge Tunçalp, and Sharon E. Straus. 2018. *Annals of Internal Medicine* 169(7):467-473. DOI: [10.7326/M18-0850](https://doi.org/10.7326/M18-0850). Research methodology. Tags: scoping review, reporting, transparency.

- **Problem and contribution:** reporting guidance for scoping reviews, which synthesize evidence and assess a literature's scope.
- **Method and evaluation:** the paper's development process and evaluation were not inspected.
- **Findings and limitations:** the associated [official PRISMA-ScR resource](https://www.prisma-statement.org/scoping) lists 20 essential and two optional reporting items covering objectives, eligibility, sources, searches, selection, charting, results, synthesis, and limitations. These details come from that official resource rather than a complete reading of the paper.
- **Relevance:** a reporting reference for a broad exploratory review; inclusion does not establish that this session is a completed or PRISMA-compliant scoping review.
- **Inspection:** bibliographic record and associated official checklist overview; full explanatory paper not inspected. Checked on 2026-09-27.

### m04

**M04. The PRISMA 2020 statement: An updated guideline for reporting systematic reviews**

Matthew J. Page, Joanne E. McKenzie, Patrick M. Bossuyt, Isabelle Boutron, Tammy C. Hoffmann, Cynthia D. Mulrow, Larissa Shamseer, Jennifer M. Tetzlaff, Elie A. Akl, Sue E. Brennan, Roger Chou, Julie Glanville, Jeremy M. Grimshaw, Asbjørn Hróbjartsson, Manoj M. Lalu, Tianjing Li, Elizabeth W. Loder, Evan Mayo-Wilson, Steve McDonald, Luke A. McGuinness, Lesley A. Stewart, James Thomas, Andrea C. Tricco, Vivian A. Welch, Penny Whiting, and David Moher. 2021. *PLOS Medicine* 18(3):e1003583. DOI: [10.1371/journal.pmed.1003583](https://doi.org/10.1371/journal.pmed.1003583). Research methodology. Tags: systematic review, reporting, transparency.

- **Problem and contribution:** updated reporting guidance for systematic reviews, accompanied by a checklist and expanded recommendations.
- **Method and evaluation:** development methods were not inspected.
- **Findings and limitations:** the [official checklist page](https://www.prisma-statement.org/prisma-2020-checklist) was inspected; the BMJ article route returned HTTP 403. Checklist completion is not itself evidence of substantive review quality.
- **Relevance:** transparent reporting of searches and selection, while retaining this mapping pass's actual completion status.
- **Version note:** the statement was co-published; its BMJ and PLOS versions are one retained work.
- **Inspection:** bibliographic record and official checklist overview, not the complete paper. Checked on 2026-09-27.

### m05

**M05. Generating Research Questions Through Problematization**

Mats Alvesson and Jörgen Sandberg. 2011. *Academy of Management Review*. DOI: [10.5465/amr.2009.0188](https://doi.org/10.5465/amr.2009.0188). Research methodology. Tags: problem formulation, assumptions, research gaps.

- **Problem and contribution:** the abstract contrasts locating gaps in theories with challenging their assumptions and proposes problematization as a methodology for developing research questions.
- **Method and evaluation:** detailed procedure, examples, and supporting analysis were not inspected.
- **Findings and limitations:** conceptual methodological work from management research; the abstract's argument does not establish that assumption challenges guarantee useful questions or publication.
- **Relevance:** helps distinguish an absent application from an unsupported assumption or unresolved theoretical issue. It does not settle this project's technical questions.
- **Inspection:** abstract through OpenAlex and metadata through Crossref; publisher returned HTTP 403. Checked on 2026-09-27.

### m06

**M06. Structure-Mapping: A Theoretical Framework for Analogy**

Dedre Gentner. 1983. *Cognitive Science*. DOI: [10.1207/s15516709cog0702_3](https://doi.org/10.1207/s15516709cog0702_3). Research methodology and adjacent theory. Tags: analogy, reframing, relational structure.

- **Problem and contribution:** a theory of how an analogy's interpretation follows from its parts. The abstract distinguishes relational mappings from shared object attributes and emphasizes higher-order relations.
- **Method and evaluation:** detailed formalization, examples, and evaluation were not inspected.
- **Findings and limitations:** the abstract presents a theoretical account, not evidence of successful fictional-character adaptation.
- **Relevance:** disciplined discussion of future cross-field analogies can state corresponding relations and where they fail. This does not select a representation or adaptation algorithm.
- **Inspection:** abstract and bibliographic record through Crossref, corroborated by OpenAlex, on 2026-09-27.

### m07

**M07. Troubling Trends in Machine Learning Scholarship**

Zachary C. Lipton and Jacob Steinhardt. 2018. Position paper presented at ICML 2018, The Debates. [arXiv:1807.03341v2](https://arxiv.org/abs/1807.03341). [Manuscript HTML](https://arxiv.org/html/1807.03341v2). Research methodology. Tags: claims, ablations, novelty, scientific writing.

- **Problem and contribution:** critical analysis of explanation/speculation confusion, failure to isolate empirical gains, obfuscating mathematics, and misleading terminology.
- **Method and evaluation:** examples and argumentation. The authors explicitly do not claim a balanced empirical census of ML scholarship.
- **Findings and limitations:** inspected sections explain how bundled changes and unequal hyperparameter tuning obscure the cause of an improvement, and why speculative explanations need labels. These are argued methodological concerns, not a universal experimental recipe.
- **Relevance:** helps distinguish a useful implementation, new application, empirical improvement, and scientific explanation, each with different evidence requirements.
- **Inspection:** abstract and selected manuscript sections 1-3.2, including the authors' disclaimers, on 2026-09-27.

### m08

**M08. Deep Reinforcement Learning That Matters**

Peter Henderson, Riashat Islam, Philip Bachman, Joelle Pineau, Doina Precup, and David Meger. 2018. *Proceedings of the AAAI Conference on Artificial Intelligence* 32(1). DOI: [10.1609/aaai.v32i1.11694](https://doi.org/10.1609/aaai.v32i1.11694). [arXiv:1709.06560](https://arxiv.org/abs/1709.06560). Research methodology and adjacent evaluation. Tags: baselines, variance, reproducibility, reinforcement learning.

- **Problem and contribution:** stochasticity and implementation choices complicate reproducible, fair deep-RL comparisons.
- **Method and evaluation:** inspected setup covers TRPO, DDPG, PPO, and ACKTR in continuous-control settings, examining hyperparameters, seeds, environments, and codebases. Detailed results across all settings were not inspected.
- **Findings and limitations:** the abstract reports variability that complicates interpreting apparent improvements. Evidence concerns selected deep-RL settings and does not establish a trial count, test, or agent for unit-design research.
- **Relevance:** future stochastic training, search, or automated playtesting would require attention to resources and uncertainty; inclusion does not recommend RL.
- **Version note:** selected reading used arXiv v3, revised January 2019, alongside AAAI 2018 metadata; version differences were not exhaustively compared.
- **Inspection:** abstract, introduction, and experimental setup; Crossref metadata checked on 2026-09-27.

### m09

**M09. Improving Reproducibility in Machine Learning Research (A Report from the NeurIPS 2019 Reproducibility Program)**

Joelle Pineau, Philippe Vincent-Lamarre, Koustuv Sinha, Vincent Larivière, Alina Beygelzimer, Florence d'Alché-Buc, Emily Fox, and Hugo Larochelle. 2021. *Journal of Machine Learning Research* 22(164):1-20. Stable record: JMLR `v22/20-303`. [Publication](https://www.jmlr.org/papers/v22/20-303.html). [Paper](https://www.jmlr.org/papers/volume22/20-303/20-303.pdf). Research methodology. Tags: reproducibility, reporting, negative results, claim limits.

- **Problem and contribution:** reports a conference reproducibility program incorporating code submission, a reproducibility challenge, and a checklist.
- **Method and evaluation:** conference case study with participation and reviewer-response analysis.
- **Findings and limitations:** encouraging participation indicators do not establish that the interventions caused better research quality; the authors explicitly state that causal evidence is not conclusive. Inspected discussion covers underspecified procedures and metrics, selective reporting, adaptive overfitting, and overclaiming.
- **Relevance:** preserve information needed to assess and reproduce a particular claim without treating a checklist or released code as a validity certificate.
- **Inspection:** publication abstract and selected PDF pages 1-3 and 12-15 on 2026-09-27.

### m10

**M10. Equivalence Tests: A Practical Primer for t Tests, Correlations, and Meta-Analyses**

Daniël Lakens. 2017. *Social Psychological and Personality Science* 8(4):355-362. DOI: [10.1177/1948550617697177](https://doi.org/10.1177/1948550617697177). [Full text in PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC5502906/). Research methodology. Tags: negative results, uncertainty, meaningful effects, equivalence.

- **Problem and contribution:** a nonsignificant difference is often interpreted incorrectly as absence of an effect. The article presents practical equivalence-testing guidance, including two one-sided tests, bounds, power analysis, and software.
- **Method and evaluation:** inspected sections explain bounds, testing logic, and four possible combinations of significance and equivalence outcomes. Detailed worked examples and software were not evaluated.
- **Findings and limitations:** nonsignificance alone cannot establish no meaningful difference; equivalence claims depend on justified bounds and sufficient information.
- **Relevance:** preserves distinctions among negative, equivalent, and inconclusive findings. This does not select the project's statistical approach.
- **Inspection:** abstract, introduction, and primary explanation of the TOST procedure and outcome types on 2026-09-27.

### m11

**M11. The preregistration revolution**

Brian A. Nosek, Charles R. Ebersole, Alexander C. DeHaven, and David T. Mellor. 2018. *Proceedings of the National Academy of Sciences*. DOI: [10.1073/pnas.1708274114](https://doi.org/10.1073/pnas.1708274114). Research methodology. Tags: hypotheses, exploration, confirmation, preregistration.

- **Problem and contribution:** confusing hypotheses generated from observed data with predictions tested against new observations weakens interpretation. The abstract describes preregistration and strategies for distinguishing these cases, including pre-existing data.
- **Method and evaluation:** detailed examples, strategies, and supporting evidence were not inspected.
- **Findings and limitations:** the abstract supports the conceptual distinction; its advocacy is not independent causal proof that preregistration resolves every validity problem.
- **Relevance:** retain this mapping's exploratory status and distinguish later predictions from retrospective explanations.
- **Inspection:** abstract and bibliographic record through Crossref on 2026-09-27.

### m12

**M12. Measurement Schmeasurement: Questionable Measurement Practices and How to Avoid Them**

Jessica Kay Flake and Eiko I. Fried. 2020. *Advances in Methods and Practices in Psychological Science*. DOI: [10.1177/2515245920952393](https://doi.org/10.1177/2515245920952393). Research methodology. Tags: measurement, construct validity, evaluation, transparency.

- **Problem and contribution:** defines questionable measurement practices and offers questions intended to make measurement choices assessable.
- **Method and evaluation:** the abstract describes an argument and practical questions; detailed examples and the question list were not inspected.
- **Findings and limitations:** the abstract connects opaque measurement choices to threats to construct, internal, external, and statistical-conclusion validity. It does not validate metrics for game-unit design.
- **Relevance:** fidelity, creativity, coherence, and usefulness need defensible measurement interpretations; a score's name alone does not establish what it measures.
- **Inspection:** abstract and bibliographic record through Crossref on 2026-09-27.

### m13

**M13. How to read a paper**

S. Keshav. 2007. *ACM SIGCOMM Computer Communication Review*. DOI: [10.1145/1273445.1273458](https://doi.org/10.1145/1273445.1273458). Research methodology. Tags: paper reading, literature practice, inspection depth.

- **Problem and contribution:** the abstract presents a practical three-pass reading method and discusses its use in a literature survey.
- **Method and evaluation:** detailed steps and any supporting evaluation were not inspected.
- **Findings and limitations:** no effectiveness claim or detailed procedure is attributed from the abstract. A guessed institutional mirror returned HTTP 404.
- **Relevance:** guidance for staged reading and distinguishing preliminary discovery from substantive inspection.
- **Inspection:** abstract and bibliographic record through Crossref on 2026-09-27.

## Non-paper resources

The following are discovery or access resources, not peer-reviewed evidence and not a second paper inventory.

- [arXiv](https://arxiv.org/) and its [API](https://export.arxiv.org/api/query): primary manuscript records and abstracts; HTML sections were read where stated. Posting a manuscript does not certify peer review.
- [Crossref REST API](https://api.crossref.org/): authoritative DOI deposit metadata used for titles, authors, dates, and venues. Exact-title results were screened; unrelated matches were rejected.
- [OpenAlex](https://openalex.org/): discovery index used for abstract access and citation trails. Its abstracts and citation links are secondary records, and its repository/venue labels can be wrong.
- [ACL Anthology](https://aclanthology.org/): publication destination identified by verified ACL DOIs for CROSS24 through AUSTEN24. The local annotations state which primary manuscript sections were actually inspected.
- [TypeSafe AI's Jev announcement](https://typesafe.ai/blog/introducing-system-one-models-and-jev) and [model documentation](https://docs.typesafe.ai/models): vendor primary sources for the product name and its expansion of RLCD as “Reinforcement Learning for Calibrated Decisions.” The announcement's speed, calibration, and hallucination claims are vendor claims, not independent evidence for game-design quality or a public account of the complete training method.
- [Official Microsoft Research writing resource](https://www.microsoft.com/en-us/research/academic-program/write-great-research-paper/): expert educational guidance used in [METHODOLOGY.md](METHODOLOGY.md), not a peer-reviewed study.
- [Steam storefront](https://store.steampowered.com/) and the [official Arknights: Endfield site](https://endfield.gryphline.com/en-us/): non-paper title checks for the [user supplied corpus](CORPORA.md). The game rules have not been compared in this phase.
- [td-profile](https://github.com/mardwerk/td-profile): Kyle DerZweite's MIT-licensed tower-defense Profile Validator and reusable schemas (v1.0.2); a copyable Profile plus a validator that checks game-data records against declared mechanics, rules, references and numerical units without rewriting the data. Inspected 2026-10-02; also catalogued as [TDPROFILE26](#tdprofile26) under Rule representations and reusable corpora.

Repositories and datasets named by the retained papers have not thereby been audited for licensing, completeness, reproducibility, or suitability. Game documentation and the user-provided corpora belong in [CORPORA.md](CORPORA.md), not in the scholarly paper inventory.

## Paper index

Use this list when a paper title is easier to remember than its citation ID. The linked ID headings keep citations stable.

**Procedural content generation and broad maps**

- [SBPCG11](#sbpcg11): Search-Based Procedural Content Generation: A Taxonomy and Survey
- [PCGML18](#pcgml18): Procedural Content Generation via Machine Learning (PCGML)
- [PCGQD19](#pcgqd19): Procedural Content Generation through Quality Diversity
- [PCGRL20](#pcgrl20): PCGRL: Procedural Content Generation via Reinforcement Learning
- [DLPCG20](#dlpcg20): Deep Learning for Procedural Content Generation
- [LLMGAMES24](#llmgames24): Large Language Models and Games: A Survey and Roadmap

**Tower Defense, RTS units, and game-specific evidence**

- [RTSUNITS22](#rtsunits22): Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search
- [TDCI11](#tdci11): Computational intelligence and tower defence games
- [TDLEVELS19](#tdlevels19): Automatic generation of tower defense levels using PCG
- [TDGA21](#tdga21): Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms
- [TDSTRATEGY24](#tdstrategy24): Reinforcement Learning for High-Level Strategic Control in Tower Defense Games
- [TOWERMIND26](#towermind26): TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents
- [MONSTER26](#monster26): Application of machine learning to monster level prediction in tabletop RPG game design

**Automated design, mechanics, and language-guided generation**

- [MECHGEN14](#mechgen14): Automatic Game Design via Mechanic Generation
- [BOSS17](#boss17): Program Synthesis as a Generative Method
- [CONCEPT18](#concept18): Automated Game Design via Conceptual Expansion
- [GAVEL24](#gavel24): GAVEL: Generating Games via Evolution and Language Models
- [MM23](#mm23): Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules
- [GGDG24](#ggdg24): Grammar-based Game Description Generation using Large Language Models
- [CHATPCG24](#chatpcg24): ChatPCG: Large Language Model-Driven Reward Design for Procedural Content Generation
- [PCGRLLM26](#pcgrllm26): PCGRLLM: Large Language Model-Driven Reward Design for Procedural Content Generation Reinforcement Learning
- [MORTAR26](#mortar26): Mortar: Evolving Mechanics for Automatic Game Design
- [TANAGRA11](#tanagra11): Tanagra: Reactive Planning and Constraint Solving for Mixed-Initiative Level Design
- [CAD26](#cad26): Procedural Content Metageneration via Program Search and Continual Abstraction Discovery
- [AUTOBG26](#autobg26): AutoBG: A Board Game Design Assistant with Interactive Ideation, Iterative Rulebook Generation, and Individualized Feedback
- [INVENTION25](#invention25): Generation and Evaluation in the Human Invention Process through the Lens of Game Design

**Rule representations and reusable corpora**

- [VGDL13](#vgdl13): A video game description language for model-based or interactive learning
- [LUDII19](#ludii19): An Overview of the Ludii General Game System
- [VGLC16](#vglc16): The VGLC: The Video Game Level Corpus
- [HDPCG26](#hdpcg26): High Dimensional Procedural Content Generation
- [TDPROFILE26](#tdprofile26): td-profile: A Profile Validator and Reusable Schemas for Tower-Defense Game Data

**Character understanding and fidelity**

- [CHARACTERS03](#characters03): Characters in Computer Games: Toward Understanding Interpretation and Design
- [NPC07](#npc07): Gameplay Design Patterns for Believable Non-Player Characters
- [THON19](#thon19): Transmedia characters: Theory and analysis
- [CROSS24](#cross24): Evaluating Character Understanding of Large Language Models via Character Profiling from Fictional Works
- [INCHARACTER24](#incharacter24): InCharacter: Evaluating Personality Fidelity in Role-Playing Agents through Psychological Interviews
- [AUSTEN24](#austen24): Evaluating Computational Representations of Character: An Austen Character Similarity Benchmark
- [ROLELLM24](#rolellm24): RoleLLM: Benchmarking, Eliciting, and Enhancing Role-Playing Abilities of Large Language Models
- [MIMIC22](#mimic22): Meet Your Favorite Character: Open-domain Chatbot Mimicking Fictional Characters with only a Few Utterances

**Grounded knowledge and acquiring constraints**

- [RAG20](#rag20): Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks
- [ICA23](#ica23): Learning to Learn in Interactive Constraint Acquisition
- [GENCON24](#gencon24): Generalizing Constraint Models in Constraint Acquisition

**Evaluation, factuality, and automated playtesting**

- [CSI14](#csi14): Quantifying the Creativity Support of Digital Tools through the Creativity Support Index
- [PCGBENCH25](#pcgbench25): The Procedural Content Generation Benchmark: An Open-source Testbed for Generative Challenges in Games
- [FACTSCORE23](#factscore23): FActScore: Fine-grained Atomic Evaluation of Factual Precision in Long Form Text Generation
- [MCTS15](#mcts15): Monte-Carlo Tree Search for Simulation-based Strategy Analysis
- [PERSONAS19](#personas19): Automated Playtesting With Procedural Personas Through MCTS With Evolved Heuristics
- [ERA23](#era23): The Right Variety: Improving Expressive Range Analysis with Metric Selection Methods

**Decision models and the RLCD acronym**

- [JEV26](#jev26): Just Ask Jev: Reinforcement Learning for Calibrated Decisions as a Zero-Shot Detector of AI Alignment Failures
- [RLCD23](#rlcd23): RLCD: Reinforcement Learning from Contrastive Distillation for Language Model Alignment

**Research methodology**

- [M01](#m01): Guidelines for conducting systematic mapping studies in software engineering: An update
- [M02](#m02): Guidelines for Snowballing in Systematic Literature Studies and a Replication in Software Engineering
- [M03](#m03): PRISMA Extension for Scoping Reviews (PRISMA-ScR): Checklist and Explanation
- [M04](#m04): The PRISMA 2020 statement: An updated guideline for reporting systematic reviews
- [M05](#m05): Generating Research Questions Through Problematization
- [M06](#m06): Structure-Mapping: A Theoretical Framework for Analogy
- [M07](#m07): Troubling Trends in Machine Learning Scholarship
- [M08](#m08): Deep Reinforcement Learning That Matters
- [M09](#m09): Improving Reproducibility in Machine Learning Research (A Report from the NeurIPS 2019 Reproducibility Program)
- [M10](#m10): Equivalence Tests: A Practical Primer for t Tests, Correlations, and Meta-Analyses
- [M11](#m11): The preregistration revolution
- [M12](#m12): Measurement Schmeasurement: Questionable Measurement Practices and How to Avoid Them
- [M13](#m13): How to read a paper
## Papers added on 2026-10-01
Entries below were added during the 2026-10-01 search pass. They are abstract-level finds; DOIs and venues are from Crossref or OpenAlex unless noted, and full-text inspection is pending.


### Procedural content generation and broad maps (added 2026-10-01)
- [consi25](RESEARCH_LIBRARY.md#consi25): CONSI25. Constraint Is All You Need: Optimization-Based 3D Level Generation with LLMs
- [coopen25](RESEARCH_LIBRARY.md#coopen25): COOPEN25. Procedural Content Generation for Cooperative Games—A Systematic Review
- [latent23](RESEARCH_LIBRARY.md#latent23): LATENT23. Learning latent representations for controllable combinational creativity and game design
- [diverse25](RESEARCH_LIBRARY.md#diverse25): DIVERSE25. Diverse Level Generation via Machine Learning of Quality Diversity
- [spellspace25](RESEARCH_LIBRARY.md#spellspace25): SPELLSPACE25. Exploring the Possibility Space of 1 Billion Spells
- [spellforger25](RESEARCH_LIBRARY.md#spellforger25): SPELLFORGER25. SpellForger: Prompting Custom Spell Properties In-Game using BERT supervised-trained model
- [gameplay25](RESEARCH_LIBRARY.md#gameplay25): GAMEPLAY25. Gameplay Evolution: A Game Design Method Based on Evolutionary Theory
- [rpgagent26](RESEARCH_LIBRARY.md#rpgagent26): RPGAGENT26. RPGAgent: Driving Coherent Story-to-Play Generation with an LLM-Based Multi-Agent System
- [enemy26](RESEARCH_LIBRARY.md#enemy26): ENEMY26. An Exploration of Collision-based Enemy Morphology Generation
- [skills26](RESEARCH_LIBRARY.md#skills26): SKILLS26. Generate Diverse Skills with Large Language Models
- [gptgames25](RESEARCH_LIBRARY.md#gptgames25): GPTGAMES25. GPT for Games: An Updated Scoping Review (2020–2024)
- [ubcl26](RESEARCH_LIBRARY.md#ubcl26): UBCL26. UBCL: A Reinforcement Learning Framework for Controllable and Diverse Player Behaviors
- [bsp26](RESEARCH_LIBRARY.md#bsp26): BSP26. Two-Scale Controllability of a Parameterized BSP Dungeon Generator for Player Flow Optimization

### Automated design, mechanics, and language-guided generation (added 2026-10-01)
- [playchar25](RESEARCH_LIBRARY.md#playchar25): PLAYCHAR25. Generative Methods for Creating Adaptive Playable Characters in Service Games
- [starcharm25](RESEARCH_LIBRARY.md#starcharm25): STARCHARM25. Democratizing Game Modding with GenAI: A Case Study of StarCharM, a Stardew Valley Character Modding Assistant
- [devwhat26](RESEARCH_LIBRARY.md#devwhat26): DEVWHAT26. What game developers actually want from procedural level generation tools
- [spacetime26](RESEARCH_LIBRARY.md#spacetime26): SPACETIME26. Constraint-Based Four-Dimensional Spacetime Blending of Learned Game Mechanics Along Orthogonal Axes
- [hades26](RESEARCH_LIBRARY.md#hades26): HADES26. Vanquish Your Past: Shifted Imitation Learning in Hades
- [behbts26](RESEARCH_LIBRARY.md#behbts26): BEHBTS26. Large Language Models for Behavior Trees Generation in Unreal Engine

### Character understanding and fidelity (added 2026-10-01)
- [omnic26](RESEARCH_LIBRARY.md#omnic26): OMNIC26. OMNICharacter++: Toward a Comprehensive Benchmark for Realistic Role-Playing Agents
- [llmnpccmp26](RESEARCH_LIBRARY.md#llmnpccmp26): LLMNPCCMP26. Comparing LLM-Driven and Script-Based Non-Player Characters Under Controlled Information
- [fictionrag26](RESEARCH_LIBRARY.md#fictionrag26): FICTIONRAG26. FictionRAG: A Stateful Metacognitive Framework for High-Fidelity Long-Narrative Role-Playing
- [dream26](RESEARCH_LIBRARY.md#dream26): DREAM26. DREAM: LLM-based Dynamic Role-playing via Event-Aware Memory Graph
- [baijia26](RESEARCH_LIBRARY.md#baijia26): BAIJIA26. BaiJia: An Open Role-Playing Platform of Chinese Historical Characters
- [questgram21](RESEARCH_LIBRARY.md#questgram21): QUESTGRAM21. Questgram [Qg]: Toward a Mixed-Initiative Quest Generation Tool
- [trpgpcg20](RESEARCH_LIBRARY.md#trpgpcg20): TRPGPCG20. Tabletop Roleplaying Games as Procedural Content Generators
- [rpgcreature23](RESEARCH_LIBRARY.md#rpgcreature23): RPGCREATURE23. RPG Creature Design: Cross-System Analysis and Conversion
- [ignite23](RESEARCH_LIBRARY.md#ignite23): IGNITE23. "It Has to Ignite Their Creativity": Opportunities for Generative Tools for Game Masters
- [fighting25](RESEARCH_LIBRARY.md#fighting25): FIGHTING25. A Survey on the Development of Human-Centered Artificial Agents for Fighting Games

### Evaluation, factuality, and automated playtesting (added 2026-10-01)
- [npcreal25](RESEARCH_LIBRARY.md#npcreal25): NPCREAL25. An Empirical Evaluation of AI-Powered Non-Player Characters' Perceived Realism and Performance
- [explai26](RESEARCH_LIBRARY.md#explai26): EXPLAI26. Transparent and Adaptive Gaming: The Role of Explainable AI and Dynamic Difficulty

### Rule representations and reusable corpora (added 2026-10-01)
- [behbtsunreal26](RESEARCH_LIBRARY.md#behbtsunreal26): BEHBTSUN26. Large Language Models for Behavior Trees Generation in Unreal Engine

### Tower Defense, RTS units, and game-specific evidence (added 2026-10-01)
- [tdnovel22](RESEARCH_LIBRARY.md#tdnovel22): TDNOVEL22. A Novel Procedural Content Generation Algorithm for Tower Defense Games
- [tdwavegen22](RESEARCH_LIBRARY.md#tdwavegen22): TDWAVE22. A NEAT Approach to Wave Generation in Tower Defense Games

### Procedural content generation and broad maps (added 2026-10-01)
- [gameplay26](RESEARCH_LIBRARY.md#gameplay26): GAMEPLAY26. From Gameplay Traces to Game Mechanics: Causal Induction with Large Language Models
- [multiobj25](RESEARCH_LIBRARY.md#multiobj25): MULTIOBJ25. Multi-Objective Instruction-Aware Representation Learning in Procedural Content Generation RL
- [runtime26](RESEARCH_LIBRARY.md#runtime26): RUNTIME26. Runtime Evaluation of Procedural Content Generation in an Endless Runner Game Using Autonomous Agents
- [zero3d25](RESEARCH_LIBRARY.md#zero3d25): ZERO3D25. Zero-shot 3D Map Generation with LLM Agents: A Dual-Agent Architecture for Procedural Content Generation
- [instruct25](RESEARCH_LIBRARY.md#instruct25): INSTRUCT25. IPCGRL: Language-Instructed Reinforcement Learning for Procedural Level Generation
- [cppn26](RESEARCH_LIBRARY.md#cppn26): CPPN26. CPPN2WFC: Extending Wave Function Collapse to Generate Globally Coherent Content

### Character understanding and fidelity (added 2026-10-01)
- [charhalluc26](RESEARCH_LIBRARY.md#charhalluc26): CHARHALL26. CHARM: Character Hallucination for Multicultural Role Play Benchmark
- [ooc25](RESEARCH_LIBRARY.md#ooc25): OOC25. Spotting Out-of-Character Behavior: Atomic-Level Evaluation of Persona Fidelity in Open-Ended Generation

### Evaluation, factuality, and automated playtesting (added 2026-10-01)
- [onboarding25](RESEARCH_LIBRARY.md#onboarding25): ONBOARD25. Support Autonomy: Exploring Player Perspectives on AI-Supported Onboarding in Video Games
- [dynamicnpc25](RESEARCH_LIBRARY.md#dynamicnpc25): DYNAMICNPC25. Leveraging Large Language Models for Dynamic NPC Interactions in 2D RPGs
- [genai3d26](RESEARCH_LIBRARY.md#genai3d26): GENAI3D26. A deep generative approach to personalized super mario level design
- [mixedinit25](RESEARCH_LIBRARY.md#mixedinit25): MIXEDINIT25. Boosting Mixed-Initiative Co-Creativity in Game Design: A Tutorial
- [script25](RESEARCH_LIBRARY.md#script25): SCRIPT25. ScriptDoctor: Automatic Generation of PuzzleScript Games via Large Language Models and Tree Search
- [seamless25](RESEARCH_LIBRARY.md#seamless25): SEAMLESS25. Seamless Tutorial: Contextual State Transition Generation Based on Player Internal Knowledge
- [xgen26](RESEARCH_LIBRARY.md#xgen26): XGEN26. AI-directed procedural content generation for personalized XR in industrial-heritage museums
- [genaiarch26](RESEARCH_LIBRARY.md#genaiarch26): GENAIARCH26. Computational game content generation and socially responsible design: Current practices, ...
