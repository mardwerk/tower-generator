# Generator research: game design research 2026 09 22

Source investigation recorded September 22, 2026. "Inspected" means that relevant primary text was accessed; experiments were not reproduced. The source text was edited later, so the inspection date is not a last-edit date. The separate September 26 review checked metadata and selected software examples; it did not reread the paper findings. No new literature investigation is claimed here.

These are dated source observations and proposals. Preserve their recorded review and rejection status. Current local contracts govern acceptance; these notes neither select implementation nor resume paused development.

### R1. Generating RTS units

Kynan Sorochan and Matthew Guzdial. **Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search.** Experimental AI in Games Workshop, 2022. [arXiv:2212.03387v1](https://arxiv.org/html/2212.03387v1). Inspected: Search Space, Evaluator, Search Method, generated units, evaluation and references.

The system adds a fourth military unit to microRTS. It searches numerical attributes and a human-authored vocabulary of four causes and four effects. Causes include taking damage and dealing damage; effects include resource return, additional attacks, healing and attack-speed changes. Greedy search changes attributes or cause/effect choices. Each candidate evaluation uses 20 games: ten with asymmetric access to the new unit and ten with access for both sides. Later evaluation compares agents with different search budgets. Ten generated units are presented.


Additional limitation from the expanded reading: the evaluation uses one map, and the strongest agent often underuses the new units. Its preference for the original roster can confound unit usefulness with evaluator competence. The search budgets are not validated human skill models; generated names and shared visuals are not evidence of visual-content generation.

Citation trail: the paper explicitly builds on R7 for its evaluation approach, cites R2 for self-play quality assessment, and cites R9 for recombination. It also identifies program synthesis and map-generation work, retained below as leads.

### R3. GAVEL

Graham Todd, Alexander G. Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers and Julian Togelius. **GAVEL: Generating Games Via Evolution and Language Models.** NeurIPS 2024. [arXiv:2407.09388v2](https://arxiv.org/html/2407.09388v2). Inspected: representation, mutation, evaluation, experiments, results, limitations and references. [Code link](https://github.com/gdrtodd/gavel), not executed in this investigation.

GAVEL combines a code language model, Ludii game descriptions and quality-diversity search. It mutates portions of executable rules and evaluates candidates through staged checks and simulated games. Evaluation includes balance, decisiveness, completion, agency, board coverage and a strategic-depth proxy. It retains high-fitness representatives in a concept-based archive.

The paper reports significant improvements over its baselines in QD score and archive coverage, but improvements in novel occupied cells were not statistically significant over three runs. Preliminary expert playtesting found promising examples, including Yavalath-derived line conditions combined with Go-style enclosure capture. Each search run took roughly 48 hours using an RTX8000 GPU and parallel evaluation on 16 CPU cores with 128 GB memory. This is not evidence of interactive latency.

Adaptation: make small, inspectable changes to a candidate's mechanical relationships, then apply independent checks. A generated branch should state what changes relative to the baseline and which purchasing decision it adds. Keep the proposed difference separate from the measured result.


Citation trail: GAVEL cites the PCG-QD survey R4, MAP-Elites R5, Ludi R2, and automatic mechanic generation, now inspected as R21. Its board-game descriptors and fitness formula should not become our unit-design rubric unchanged.

### R4. Quality diversity for PCG

Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius and Georgios N. Yannakakis. **Procedural Content Generation through Quality Diversity.** IEEE Conference on Games, 2019. [arXiv:1907.04053](https://arxiv.org/abs/1907.04053), [HTML](https://arxiv.org/html/1907.04053), [DOI:10.1109/CIG.2019.8848053](https://doi.org/10.1109/CIG.2019.8848053). Inspected: framework, surveyed applications, challenges and references.

This survey explains searching for a collection of varied, good artifacts rather than one global winner. It distinguishes quality criteria, diversity descriptions and feasibility constraints, and discusses levels, decks, agents and other content. It is a framework and survey, not a direct experiment in generating our units.

Adaptation: maintain a small reviewed collection organized by strategic contribution. Possible descriptors include placement dependence, timing sensitivity, target-selection dependence, economic commitment and reliance on another unit. These are hypotheses about useful dimensions. A label like “support” or a different damage value is insufficient evidence of a different decision.


Citation trail: R4 cites MAP-Elites R5, novelty search R6, constrained novelty R10 and Sentient Sketchbook R11. It also points to Talakat, deck-space mapping and interactive dungeon design as domain-specific applications.

### R5. MAP-Elites

Jean-Baptiste Mouret and Jeff Clune. **Illuminating search spaces by mapping elites.** 2015 preprint. [arXiv:1504.04909v1](https://arxiv.org/abs/1504.04909v1). Inspected: algorithm, experimental domains, limitations and references via arXiv HTML. The HTML contains an inconsistent generated date; use the arXiv version record for the 2015 date.

MAP-Elites partitions a user-selected feature space and retains a high-performing candidate for each region. Its experiments concern neural networks and robots, not games. The paper explicitly presents itself as preliminary work.

Adaptation: use the archive idea first as a review aid. Retain one or several promising alternatives per meaningful design category, with evidence and unresolved issues. We do not need a grid implementation or automated evolution to benefit from comparing alternatives.

Limits: archive coverage means coverage of the chosen descriptors. It does not establish strategic diversity or historical novelty. A category populated by several different descriptions may still contain the same behavior. A single quality score within each category also remains a design choice.

Second citation level: R5 cites R6, explaining why retaining behaviorally different stepping stones can help avoid an overly narrow objective.

### R6. Novelty search

Joel Lehman and Kenneth O. Stanley. **Abandoning Objectives: Evolution through the Search for Novelty Alone.** Evolutionary Computation 19(2), pages 189 to 223, 2011. [Primary-text PDF](https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf). Inspected: motivation, behavioral novelty, experimental tasks and discussion. Found in both R4 and R5.

Novelty search rewards behavioral difference rather than progress toward a fixed objective. The experiments show benefits in deceptive maze-navigation and biped-walking tasks. The lesson is that some objectives can guide search toward dead ends, not that objectives are always harmful.

Adaptation: avoid discarding every unusual unit because it initially loses on damage per cost. Preserve a candidate if it could create a different useful decision and its weaknesses remain explicit. Compare behavior, not names, lore or embedding distance alone.

Limits: novelty is neither usefulness nor fit. Our output still needs supported mechanics and project constraints. This paper is a reason to inspect the effects of our ranking criteria, not a reason to remove quality checks.

### R9. Conceptual expansion

Matthew Guzdial and Mark Riedl. **Automated Game Design via Conceptual Expansion.** AIIDE 2018, 14(1), pages 31 to 37. [arXiv:1809.02232v1](https://arxiv.org/html/1809.02232v1). Inspected: representation, expansion, evaluation, discussion and references. Reached through R1.

The system learns approximate game representations and recombines them through conceptual expansion. Evaluation uses Super Mario Bros., Kirby's Adventure and Mega Man, attempting to reconstruct a held-out existing game using knowledge from others and a partially specified target. This is evidence about recombining learned structure, not a demonstration of unrestricted original game invention.


Limits: learned approximations can omit mechanics. The paper does not establish source fidelity, appropriate numerical balance or player preference for character adaptations. Its reconstruction evaluation is particularly important when citing it as evidence for originality.

Further references include Mechanic Miner, now inspected as R18, and conceptual blending work, which remains a lead.

### R10. Constrained novelty search

Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. **Constrained Novelty Search: A Study on Game Content Generation.** Evolutionary Computation 23(1), pages 101 to 129, 2015. [Author PDF](https://antoniosliapis.com/papers/constrained_novelty_search.pdf), [DOI:10.1162/EVCO_a_00123](https://doi.org/10.1162/EVCO_a_00123). Inspected: two-population approach, game-content application and reported qualifications. Found through R4; bibliographic metadata also checked through Crossref.

The approach maintains feasible and infeasible populations, with selection suited to each. Experiments investigate generation of diverse playable levels. Performance depends on the search space and operators; the paper does not claim a universally best configuration.


Limits: our draft/proposal distinction is an organizational adaptation, not an implementation of the paper's algorithm. We should not add dual evolutionary populations merely to store unfinished drafts.

### R14. FunSearch

Bernardino Romera-Paredes et al. **Mathematical discoveries from program search with large language models.** Nature 625, pages 468 to 475, 2024; published online in December 2023. [DOI:10.1038/s41586-023-06924-6](https://www.nature.com/articles/s41586-023-06924-6). Inspected: primary article's method and reported applications. [Code link](https://github.com/google-deepmind/funsearch), not executed here.

FunSearch combines LLM program proposals, evolutionary search and external evaluation. It reports new cap-set constructions and online bin-packing heuristics that improve on widely used baselines. The evaluator is central to the achievement.

Adaptation: keep generation, executable validation and design review as distinct operations. Save the actual evaluation result and conditions rather than letting the proposal's explanation serve as its evidence. Benchmark any search or repair loop against the same budget and inputs used by the simpler generator.

Limits: mathematical validity and algorithmic performance are more directly evaluable than fun or character fidelity. External code checks only establish the properties they actually test. FunSearch is not evidence that an LLM reviewer can replace playtesting.

### R16. LLM idea generation and self-evaluation

Chenglei Si, Diyi Yang and Tatsunori Hashimoto. **Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers.** ICLR 2025; preprint 2024. [arXiv:2409.04109](https://arxiv.org/abs/2409.04109), [v1 inspected](https://arxiv.org/html/2409.04109v1). Inspected: study description and discussion of novelty, diversity and self-evaluation. The inspected evidence is the preprint, not an audit of changes in the final proceedings version.

AI proposals received higher novelty ratings in the experiment, while the paper also identifies limitations in self-evaluation and idea diversity. The evaluated objects were research proposals, not completed projects.

Adaptation: use model critique to find questions and possible defects, then require independent evidence for claims of correctness, usefulness and originality. Compare candidates blind to the generator when practical. This is indirect evidence from research ideation, not game-design validation.

### R18. Mechanic Miner

Michael Cook, Simon Colton, Azalea Raad and Jeremy Gow. **Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design.** EvoApplications 2013. [Author PDF](https://www.possibilityspace.org/papers/evo13.pdf), [DOI:10.1007/978-3-642-37192-9_29](https://doi.org/10.1007/978-3-642-37192-9_29). Inspected: representation, search, generated mechanics and pilot study. Promoted from the original R9 citation lead.

The system inspects program fields and evolves simple two-state mechanics through operations on those fields. It also generates levels suited to the resulting mechanic. Examples include gravity inversion, bouncing and coordinate-based teleportation. The pilot study reports difficulty understanding and enjoying the teleportation mechanic despite its computational usefulness. The text gives inconsistent participant counts between the study introduction and the teleportation result, so no precise sample claim is used here.

### R32. Galactic Arms Race

Erin J. Hastings, Ratan K. Guha and Kenneth O. Stanley. **Automatic Content Generation in the Galactic Arms Race Video Game.** IEEE Transactions on Computational Intelligence and AI in Games 1(4), pages 245 to 263, 2009. [DOI:10.1109/TCIAIG.2009.2038365](https://doi.org/10.1109/TCIAIG.2009.2038365), [supplied manuscript entry](https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game). Access gap: ResearchGate returned HTTP 403.

The supplied review reports evolved particle-system weapons in an online game, guided by weapon use, with more than 1,000 registered players. These deployment and sample claims remain unverified here.

Adaptation to investigate: use player behavior as one feedback source, with context. Frequent use may reflect power, availability or progression incentives rather than preference. Retain explicit judgments and alternatives instead of interpreting usage as a universal quality score.

## Additional references and access gaps

The remaining items from the supplied list are retained for traceability. Their detailed empirical claims were not independently revalidated here, so they do not carry the main recommendations.

| Source | Status in this investigation | Potential use |
| --- | --- | --- |
| Georgios N. Yannakakis and Julian Togelius, **Experience-Driven Procedural Content Generation**, IEEE Transactions on Affective Computing 2(3), 2011. [DOI:10.1109/T-AFFC.2011.6](https://doi.org/10.1109/T-AFFC.2011.6) | Access gap: attempted author copies did not yield the paper | Frame evaluation around specified players and experiences; inspect primary text before adopting its models |
| Anil R. Doshi and Oliver P. Hauser, **Generative AI enhances individual creativity but reduces the collective diversity of novel content**, Science Advances 10(28), 2024. [DOI:10.1126/sciadv.adn5290](https://doi.org/10.1126/sciadv.adn5290) | Publisher returned HTTP 403 | Background motivation for checking roster-level similarity; the supplied summary concerns short fiction |
| Emily Wenger and Yoed N. Kenett, **Large language models are homogeneously creative**, PNAS Nexus 5(3), 2026. [DOI:10.1093/pnasnexus/pgag042](https://doi.org/10.1093/pnasnexus/pgag042) | Publisher returned HTTP 403; metadata retained from supplied list | Lead on cross-model similarity; do not assume different providers yield independent designs |
| Minkyu Shin, Jin Kim, Bas van Opheusden and Thomas L. Griffiths, **Superhuman artificial intelligence can improve human decision-making by increasing novelty**, PNAS 120(12), 2023. [DOI:10.1073/pnas.2214840120](https://doi.org/10.1073/pnas.2214840120) | Publisher returned HTTP 403 | Background on AI influencing play within an existing game; not direct evidence for rule invention |
| Margaret A. Boden, **Creativity and artificial intelligence**, Artificial Intelligence 103(1 to 2), 1998. [DOI:10.1016/S0004-3702(98)00055-1](https://doi.org/10.1016/S0004-3702(98)00055-1) | Publisher returned HTTP 403 | Conceptual vocabulary for combination, exploration and transformation; not an empirical LLM result |


The following leads were found in inspected reference lists. Only their citation context or metadata was checked; their original experiments were not evaluated. The parent paper provides a verified place to recover each reference.

| Lead | Verified citation location | Why follow it next |
| --- | --- | --- |
| Eric Butler, Kristin Siu and Alexander Zook, **Program Synthesis as a Generative Method**, FDG 2017 | R1 references | Generating dynamic behavior rather than numerical attributes alone |
| Ahmed Khalifa, Scott Lee, Andy Nealen and Julian Togelius, **Talakat: Bullet Hell Generation through Constrained MAP-Elites**, GECCO 2018 | R4 reference 18 | Combat-pattern diversity with playability constraints |
| Michael C. Fontaine et al., **Mapping Hearthstone Deck Spaces with MAP-Elites with Sliding Boundaries**, GECCO 2019 | R4 reference 13 | Interacting content collections and descriptor design |
| Julian Togelius et al., **Multiobjective Exploration of the StarCraft Map Space**, CIG 2010 | R1 references | Map objectives that can conflict rather than one universal map score |
| Georgios N. Yannakakis, Antonios Liapis and Constantine Alexopoulos, **Mixed-Initiative Co-Creativity**, FDG 2014 | R3 reference 68 | Designer control, useful alternatives and revision dialogue |


Code links are research entry points. No referenced implementation was installed, benchmarked or license-audited in this investigation. Reusing an idea and copying code, data or assets require separate decisions.

## Proposed adaptations for our workflow

These proposals combine the sources above. No paper validates this entire workflow for our project.

First, attach a short evidence record to a serious candidate. It should identify the character/source scope, source facts, adaptation choices, nearest known mechanic, supported operators and unresolved requirements. State the intended purchasing or tactical decision in one sentence. Explain the condition where another build remains preferable. This extends existing authoring evidence rather than replacing it with a new scoring system. Sources: R1, R3 and R9.

Second, keep separate judgments for source fidelity, mechanical support, meaningful difference, usefulness, readability and project fit. A unit can be faithful but unsupported, supported but redundant, or interesting but inappropriate for the game. Do not average away a failed requirement. Retain a few different promising candidates and their rejection reasons before introducing automated search. Sources: R4, R5, R10 and R11.

Third, specify a testable comparison before calling a mechanic novel. Compare the candidate with its distinctive rule removed and with its nearest conventional alternative. Record the scenario, budget, legal build, expected decision change and a negative case where the feature should not help. Before a runtime exists, label this a predicted scenario, not a simulation result. Sources: R3, R7 and R8.


Finally, compare any proposed generation pipeline with the current baseline using fixed source packets and equal total budgets, including repairs. Measure first-pass mechanical validity, preserved source behavior, distinct purchase reasons, human correction effort, latency and cost. Gameplay claims enter only when runtime or playtest evidence exists. R3 and R14 motivate independent evaluation; these particular project measures come from our needs and the existing generation review.

## Evidence to retain and claims we can make

A compact candidate record can use the following fields. This is a documentation proposal, not a required API schema.


```text
Candidate revision:
Game Definition and Profile revision:
Character source scope and evidence:
Closest known design and reference:
Shared ingredients and changed relationship:
Intended player decision and retained weakness:
Required operators and implementation status:
Comparison condition and scenario:
Expected result:
Observed result and evidence location:
Policy, runtime version, seeds and trial count, if applicable:
Human review or playtest context, if applicable:
Remaining uncertainty and review decision:
```



| Lead | Verified citation location | Why follow it next |
| --- | --- | --- |
| Claim | Minimum appropriate evidence |
| --- | --- |
| New description | Text or presentation differs; no gameplay claim follows |
| New configuration relative to our catalogue | Explicit reference set and structured comparison |
| Supported by our Engine | Applicable checks for declared operators and legal builds, with check scope |
| Creates different gameplay in tested conditions | Executed comparison showing changed decisions or behavior under named conditions |
| Preferred by tested players | Human comparison with participant context, method and uncertainty |
| Historically unprecedented | Broad prior-art investigation beyond this document; generation alone is insufficient |

Use qualified language such as “changed placement choices in these scenarios” or “preferred in this playtest.” The present research supports building better proposal and evaluation loops. It does not establish balanced, enjoyable or historically original tower defense units without project-specific evidence.

Technical documentation and character/product source notes are indexed in Generator references. They are implementation references, not additional empirical game-design evidence. Repository-local framework sources preserve the Unit Generator application and limits.

## Source identifier references

These publication identities supply local citation context. Inspection and reproduction limits are recorded above.

- R1. Kynan Sorochan and Matthew Guzdial. **Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search.** Experimental AI in Games Workshop, 2022. [arXiv:2212.03387v1](https://arxiv.org/html/2212.03387v1). Inspected: Search Space, Evaluator, Search Method, generated units, evaluation and references.
- R2. Cameron Browne and Frédéric Maire. **Evolutionary Game Design.** IEEE Transactions on Computational Intelligence and AI in Games 2(1), pages 1 to 16, 2010. [Author PDF](https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf). Inspected: quality measurements, human comparisons, Experiment III and conclusions. The filename is not the publication year.
- R3. Graham Todd, Alexander G. Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers and Julian Togelius. **GAVEL: Generating Games Via Evolution and Language Models.** NeurIPS 2024. [arXiv:2407.09388v2](https://arxiv.org/html/2407.09388v2). Inspected: representation, mutation, evaluation, experiments, results, limitations and references. [Code link](https://github.com/gdrtodd/gavel), not executed in this investigation.
- R4. Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius and Georgios N. Yannakakis. **Procedural Content Generation through Quality Diversity.** IEEE Conference on Games, 2019. [arXiv:1907.04053](https://arxiv.org/abs/1907.04053), [HTML](https://arxiv.org/html/1907.04053), [DOI:10.1109/CIG.2019.8848053](https://doi.org/10.1109/CIG.2019.8848053). Inspected: framework, surveyed applications, challenges and references.
- R5. Jean-Baptiste Mouret and Jeff Clune. **Illuminating search spaces by mapping elites.** 2015 preprint. [arXiv:1504.04909v1](https://arxiv.org/abs/1504.04909v1). Inspected: algorithm, experimental domains, limitations and references via arXiv HTML. The HTML contains an inconsistent generated date; use the arXiv version record for the 2015 date.
- R6. Joel Lehman and Kenneth O. Stanley. **Abandoning Objectives: Evolution through the Search for Novelty Alone.** Evolutionary Computation 19(2), pages 189 to 223, 2011. [Primary-text PDF](https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf). Inspected: motivation, behavioral novelty, experimental tasks and discussion. Found in both R4 and R5.
- R7. Alexander Zook, Brent Harrison and Mark O. Riedl. **Monte-Carlo Tree Search for Simulation-based Strategy Analysis.** The paper identifies FDG 2015 as its conference; the arXiv deposit is 2019. [arXiv:1908.01423v1](https://arxiv.org/html/1908.01423v1). Inspected: methodology, skill-based metrics, case studies, limitations and references. R1 cites the deposit year.
- R8. Alexander Jaffe, Alex Miller, Erik Andersen, Yun-En Liu, Anna Karlin and Zoran Popović. **Evaluating Competitive Game Balance with Restricted Play.** AIIDE 2012. [Author-hosted PDF](https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf). Inspected: framework, example application and conclusions. Reached through R1 → R7 → R8.
- R9. Matthew Guzdial and Mark Riedl. **Automated Game Design via Conceptual Expansion.** AIIDE 2018, 14(1), pages 31 to 37. [arXiv:1809.02232v1](https://arxiv.org/html/1809.02232v1). Inspected: representation, expansion, evaluation, discussion and references. Reached through R1.
- R10. Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. **Constrained Novelty Search: A Study on Game Content Generation.** Evolutionary Computation 23(1), pages 101 to 129, 2015. [Author PDF](https://antoniosliapis.com/papers/constrained_novelty_search.pdf), [DOI:10.1162/EVCO_a_00123](https://doi.org/10.1162/EVCO_a_00123). Inspected: two-population approach, game-content application and reported qualifications. Found through R4; bibliographic metadata also checked through Crossref.
- R11. Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. **Sentient Sketchbook: Computer-Aided Game Level Authoring.** Foundations of Digital Games, 2013, pages 213 to 220. [Author PDF](https://antoniosliapis.com/papers/sentient_sketchbook.pdf). Inspected: tool interaction, checks, alternatives and small-scale designer study. Found through R4.
- R14. Bernardino Romera-Paredes et al. **Mathematical discoveries from program search with large language models.** Nature 625, pages 468 to 475, 2024; published online in December 2023. [DOI:10.1038/s41586-023-06924-6](https://www.nature.com/articles/s41586-023-06924-6). Inspected: primary article's method and reported applications. [Code link](https://github.com/google-deepmind/funsearch), not executed here.
- R16. Chenglei Si, Diyi Yang and Tatsunori Hashimoto. **Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers.** ICLR 2025; preprint 2024. [arXiv:2409.04109](https://arxiv.org/abs/2409.04109), [v1 inspected](https://arxiv.org/html/2409.04109v1). Inspected: study description and discussion of novelty, diversity and self-evaluation. The inspected evidence is the preprint, not an audit of changes in the final proceedings version.
- R18. Michael Cook, Simon Colton, Azalea Raad and Jeremy Gow. **Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design.** EvoApplications 2013. [Author PDF](https://www.possibilityspace.org/papers/evo13.pdf), [DOI:10.1007/978-3-642-37192-9_29](https://doi.org/10.1007/978-3-642-37192-9_29). Inspected: representation, search, generated mechanics and pilot study. Promoted from the original R9 citation lead.
- R21. Alexander Zook and Mark O. Riedl. **Automatic Game Design via Mechanic Generation.** AAAI 2014. [arXiv:1908.01420v1](https://arxiv.org/html/1908.01420v1), [DOI:10.1609/aaai.v28i1.8788](https://doi.org/10.1609/aaai.v28i1.8788). Inspected: representation, constraint-based generation, planning and demonstration domains. The arXiv deposit is 2019; the paper is from 2014. Promoted from the original R3/R7 citation lead.
- R32. Erin J. Hastings, Ratan K. Guha and Kenneth O. Stanley. **Automatic Content Generation in the Galactic Arms Race Video Game.** IEEE Transactions on Computational Intelligence and AI in Games 1(4), pages 245 to 263, 2009. [DOI:10.1109/TCIAIG.2009.2038365](https://doi.org/10.1109/TCIAIG.2009.2038365), [supplied manuscript entry](https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game). Access gap: ResearchGate returned HTTP 403.

## Metadata review recorded September 26, 2026

This later check did not reread the attributed paper findings or reproduce the experiments. Preserve its confirmed, unverified and contradicted distinctions. WS identifies metadata search results, not full-text inspection. FS, GR and GDR identify the supplied source collections.

| Source | Cited input | Check | Status | Qualification |
| --- | --- | --- | --- | --- |
| R1 Sorochan and Guzdial, arXiv 2212.03387, EXAG 2022 | FS, GDR | WS | Metadata only | Abstract confirms ten microRTS units; cause/effect vocabulary and 20-game evaluation not read |
| R3 GAVEL, arXiv 2407.09388, NeurIPS 2024 | FS, GDR | WS, clone the inspected revision | Metadata only | Repo citation matches; no root LICENSE file; "48 hours on RTX8000" not read |
| R4 Gravina et al., arXiv 1907.04053, CoG 2019, DOI 10.1109/CIG.2019.8848053 | FS, GDR | WS | Metadata only | All identifiers match |
| R7 Zook, Harrison and Riedl, arXiv 1908.01423 | GDR | WS | Metadata only | FDG 2015 version is titled "...Simulation-based Play Strategy Analysis" |
| R14 FunSearch, Nature 625:468-475, DOI 10.1038/s41586-023-06924-6 | GDR | WS, clone the inspected revision | Metadata only | README confirms the repo omits the LLM and execution sandbox |
| R16 Si, Yang and Hashimoto, arXiv 2409.04109, ICLR 2025 | FS, GDR | WS | Metadata only | ICLR 2025 proceedings entry exists |
| R18 Mechanic Miner, EvoApplications 2013, DOI 10.1007/978-3-642-37192-9_29 | FS, GDR | WS | Metadata only | Abstract: levels generated for a discovered mechanic, consistent with "supporting challenge" |
| R21 Zook and Riedl, AAAI 28(1) 2014, arXiv 1908.01420 | GDR | WS | Metadata only | |
| R32 Hastings, Guha and Stanley, TCIAIG 1(4):245-263, 2009 | GDR | WS | Metadata only | DOI and the 1,000-player claim unchecked |
| R2, R5, R6, R8, R9, R10, R11 | GDR | WS | Metadata only | Title, authors, year, venue and cited identifier all match (R2 TCIAIG 2(1); R5 1504.04909v1; R6 EC 19(2); R8 AIIDE 2012; R9 1809.02232; R10 EVCO_a_00123; R11 FDG 2013). Attributed findings not read |
| Lead: Fontaine et al. GECCO 2019 | GDR | WS | Contradicted | Title is "...through MAP-Elites with Sliding Boundaries"; the draft says "with MAP-Elites" |

Hearthstone lead title correction: use "through MAP-Elites with Sliding Boundaries". The earlier reading-guide lead used "with MAP-Elites". The correction concerns source identity, not a reproduced experiment.
