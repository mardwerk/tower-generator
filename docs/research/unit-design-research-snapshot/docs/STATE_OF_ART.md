# State of the art map

Sources checked through 2026-09-28. This is a synthesis of the [master research library](RESEARCH_LIBRARY.md), not a claim that the search is exhaustive. The [coverage log](SEARCH_LOG.md) records unfinished searches and reading. No scientific contribution, representation, training process, or generator architecture has been selected.

In this document, **source report** means a claim made by an inspected paper within its study setting. **Cross-source inference** means a comparison we draw from two or more such settings. **Open question** means evidence here does not settle it. A work's source reported limitation is distinguished from our inference about what that work did not test.

## The nearby research does not study one uniform task

| Research thread | Typical input and output in inspected work | What is evaluated | Boundary for this project |
| --- | --- | --- | --- |
| Tower Defense and RTS content | TD levels, tower locations and waves [TDLEVELS19](RESEARCH_LIBRARY.md#tdlevels19), or fixed-format microRTS units [RTSUNITS22](RESEARCH_LIBRARY.md#rtsunits22) | Playability, usage, win rates, or difficulty | Neither task establishes character-grounded designs across different games. |
| General PCG and learned generation | Content encoded in an existing domain [PCGML18](RESEARCH_LIBRARY.md#pcgml18), edits and rewards [PCGRL20](RESEARCH_LIBRARY.md#pcgrl20) | Domain-specific validity and quality functions | Learning a generator does not by itself discover its vocabulary or quality criteria. |
| Rule and mechanic design | Action-like effects [MECHGEN14](RESEARCH_LIBRARY.md#mechgen14), Ludii rules [GAVEL24](RESEARCH_LIBRARY.md#gavel24), or composed code mechanics [MORTAR26](RESEARCH_LIBRARY.md#mortar26) | Playability, search scores, skill separation, sometimes human judgment | Executable rules and simulated play answer different questions from character fidelity. |
| Character interpretation and design | Fictional-character profiles [CROSS24](RESEARCH_LIBRARY.md#cross24), personality interviews [INCHARACTER24](RESEARCH_LIBRARY.md#incharacter24), and qualitative analysis of character actions in games [CHARACTERS03](RESEARCH_LIBRARY.md#characters03) | Factual and behavioral measures in the first two studies; conceptual game analysis in the third | Existing work relates character traits to mechanics, but the listed studies do not validate a general method for adapting a reference character across games. |
| Incomplete requirements | Candidate constraints and user judgments [ICA23](RESEARCH_LIBRARY.md#ica23), generalized models [GENCON24](RESEARCH_LIBRARY.md#gencon24) | Number of queries and constraint recovery or model accuracy | The candidate language and judgments in these studies are more defined than this project's initial constraints. |
| Automated evaluation | MCTS playstyles [PERSONAS19](RESEARCH_LIBRARY.md#personas19), expressive-range metrics [ERA23](RESEARCH_LIBRARY.md#era23), TD agents [TOWERMIND26](RESEARCH_LIBRARY.md#towermind26) | Agent behavior, metric coverage, or task success | None is a validated stand-in for all aspects of a useful, faithful unit design. |

The table compares what studies actually take as given and what they measure. It does not rank methods or identify a preferred research direction.

## Direct game and unit evidence

**Source report.** [TDCI11](RESEARCH_LIBRARY.md#tdci11) treats Tower Defense as a computational-intelligence testbed and classifies components. [TDLEVELS19](RESEARCH_LIBRARY.md#tdlevels19) models road maps, tower locations, and monster sequences for *Kingdom Rush: Frontiers*, then uses Monte Carlo search for automated playability testing. [TDGA21](RESEARCH_LIBRARY.md#tdga21) is a close bibliographic TD lead, but its content and findings have not yet been inspected.

**Scope inference:** these papers cannot establish that the 27 starting games share a mechanics schema. The corpus is intentionally varied, and its game rules were not reverse engineered in this session.

**Source report.** [RTSUNITS22](RESEARCH_LIBRARY.md#rtsunits22) is unusually close to the eventual output type. It searches six numeric attributes and hand-authored cause/effect options for new microRTS units. MCTS play checks whether an agent uses a unit, how long it survives, and whether access changes wins.

The authors present ten units and report evidence of usefulness and balance in their agent setting. That evidence varies with the player: in equal-strength matchups where only one player has the new unit, the Strong and Medium agents show no advantage on average, and the Strong agent builds the unit in only 31.7% of games. The authors suggest that its original-game tuning may limit its use of new abilities. Their explicit limitations include one standard map and no human playtesting; they say human balance is unclear.

**Scope inference:** the paper does not investigate a reference character, explanations, unknown ability concepts, or adaptation among games. This is our observation of the paper's scope, not an author claimed research gap.

[TDSTRATEGY24](RESEARCH_LIBRARY.md#tdstrategy24) and [TOWERMIND26](RESEARCH_LIBRARY.md#towermind26) study agents that play TD games. The former combines heuristics and reinforcement learning in *Plants vs. Zombies*; the latter reports a gap between tested LLM agents and humans on five benchmark levels and distinguishes valid actions from effective play. [MCTS15](RESEARCH_LIBRARY.md#mcts15) varies search budgets to model one aspect of player skill in two fixed games. Its authors explicitly distinguish that approximation from perception, memory, motor skill, and real-time play.

**Cross-source inference:** automated playtesting depends on the competence and behavior of the chosen agent. A playable design and a well played design are not the same observation. These studies do not validate a particular playtester for the user's corpus.

## Generation, mechanics, and the meaning of a reusable foundation

[SBPCG11](RESEARCH_LIBRARY.md#sbpcg11), [PCGML18](RESEARCH_LIBRARY.md#pcgml18), [DLPCG20](RESEARCH_LIBRARY.md#dlpcg20), and [LLMGAMES24](RESEARCH_LIBRARY.md#llmgames24) map different search, learned, and language-model roles in game content generation. [PCGML18](RESEARCH_LIBRARY.md#pcgml18) observed, at its 2018 publication date, a concentration on two-dimensional levels and difficulty comparing generators because quality tests and reported costs differ. That historical finding should not be recast as a current absence of work on units or rules.

**Source report.** [PCGRL20](RESEARCH_LIBRARY.md#pcgrl20) frames level editing as sequential decisions, but its experiments specify editing actions, rewards, and playability or solver criteria. The authors report difficulty with harder *Zelda* and *Sokoban* levels and identify reward design as an open problem in their setting. [MECHGEN14](RESEARCH_LIBRARY.md#mechgen14) generates mechanics under design requirements with an action-like representation and checks playability with planning. Its demonstrations cover RPG, platformer, and a combined domain. [GGDG24](RESEARCH_LIBRARY.md#ggdg24) guides language models with a supplied grammar to improve game-description syntax. Syntax validity does not establish semantic fidelity or game quality.

**Source report.** [BOSS17](RESEARCH_LIBRARY.md#boss17) uses a hand-authored, typed grammar to generate boss form and finite-state behavior, including interactions through shared variables. Its worked examples demonstrate feasible program construction. The authors distinguish well-formed programs from the numeric values needed for good boss design and leave quality heuristics for future work. A Cut Man sprite shown with a generated boss is explicitly arbitrary, with no character meaning in the generated behavior.

**Cross-source inference:** prior work reaches beyond fixed numeric unit attributes, but its supplied vocabulary and validity checks do not establish character-grounded or balanced designs.

[CHATPCG24](RESEARCH_LIBRARY.md#chatpcg24) and its later extension [PCGRLLM26](RESEARCH_LIBRARY.md#pcgrllm26) study language-model assistance in authoring rewards for PCG reinforcement learning. They address a defined reward-generation task, not the selection of every relevant rule or a complete character-derived design.

- [CONCEPT18](RESEARCH_LIBRARY.md#concept18) combines learned game representations using a partial target specification. Its reconstruction tests involve three platform games. The authors explicitly call reconstruction an imperfect stand-in for novel game design and seek broader games and human judgment.
- [GAVEL24](RESEARCH_LIBRARY.md#gavel24) generates Ludii games through language models and evolution; its authors explicitly warn that their proxy metrics do not suffice to show interesting games.
- [MM23](RESEARCH_LIBRARY.md#mm23) reports that evaluating rules with reinforcement learning produces different output from an A* evaluator.
- [MORTAR26](RESEARCH_LIBRARY.md#mortar26) evolves code mechanics in a supplied top-down game framework and judges their contribution in composed games by how consistently five agents rank by expected skill. In its ten-person study of three game pairs, the aggregate human score favored the higher-ranked game twice and the lower-ranked game once. The authors also identify absent designer control as a limitation.

These results support a specific skill-ordering proxy, not a general measure of game or character design quality.

[TANAGRA11](RESEARCH_LIBRARY.md#tanagra11) is an earlier mixed-initiative example: a human edits a platform level while the system regenerates under modeled playability constraints. [AUTOBG26](RESEARCH_LIBRARY.md#autobg26) is a recent abstract-level lead on interactive board-game ideation, rulebooks, and feedback. Both address designer participation in defined settings; neither is evidence that the human or system can automatically resolve this project's undefined mechanics.

[INVENTION25](RESEARCH_LIBRARY.md#invention25) is a preliminary empirical study of how novice participants create grid strategy games from seed examples. Its abstract reports that accounting for quality estimates improves a population-level model of those creations. It informs the distinction between proposing and assessing designs, without specifying how to evaluate this project's units.

**Cross-source inference.** These systems all make some constraints concrete before assessing content: an editing operation, action/effect vocabulary, grammar, Ludii language, evaluator, or complete-game simulator. They demonstrate useful forms of automated design within those settings. Their combined evidence does not select a shared representation for this project, show how to treat an unseen ability, or establish that training is required. [CAD26](RESEARCH_LIBRARY.md#cad26) is a recent abstract-level lead that extracts reusable primitives during program search; [HDPCG26](RESEARCH_LIBRARY.md#hdpcg26) studies layered and temporal dimensions. Both warrant closer reading before any claim that such topics are unstudied or solved.

[VGDL13](RESEARCH_LIBRARY.md#vgdl13) and [LUDII19](RESEARCH_LIBRARY.md#ludii19) show existing machine-readable game descriptions with interpreters or agents. [VGLC16](RESEARCH_LIBRARY.md#vglc16) is a level corpus for game-AI research. They establish that reusable formal systems and datasets exist, with particular coverage boundaries.

**Open question:** what, if anything, can be shared across the supplied games while preserving exceptions and meanings. This document does not answer it.

## Character knowledge, adaptation, and gameplay expression

**Source report.** [CROSS24](RESEARCH_LIBRARY.md#cross24) studies character profiles from fictional works, including attributes, relationships, events, and personality. It reports character and relationship confusion, omitted information, and narrative misinterpretation. [INCHARACTER24](RESEARCH_LIBRARY.md#incharacter24) evaluates personality fidelity of role-playing agents through interviews. [AUSTEN24](RESEARCH_LIBRARY.md#austen24) finds that tested representations capture some broad social similarities while struggling with expert literary similarities in Austen novels. [ROLELLM24](RESEARCH_LIBRARY.md#rolellm24) and [MIMIC22](RESEARCH_LIBRARY.md#mimic22) study character-conditioned dialogue and style. Their measurements differ from each other and from a gameplay unit design.

**Source report.** [CHARACTERS03](RESEARCH_LIBRARY.md#characters03) directly discusses character expression through game rules. Its Hulk and Bruce Banner analysis links contrasting traits to possible actions, object use, attack damage, damage-triggered strength, and detection. It also discusses goals, weaknesses, and development of skills. This is a qualitative design analysis, not a controlled test of fidelity or an automated adaptation method. [THON19](RESEARCH_LIBRARY.md#thon19) distinguishes a character's versions across media and relations among those versions.

**Cross-source inference:** an adaptation may need to state which version and traits it draws from, as well as how the target game's actions express them. These sources provide conceptual distinctions, not a selected representation for this project.

**Source report.** [NPC07](RESEARCH_LIBRARY.md#npc07) analyzes believable behavior in one *Oblivion* character and explicitly cautions that believability can conflict with gameplay goals. Its patterns come from qualitative game analysis, without a study of player interpretations.

**Cross-source inference:** a design can express a character trait while still producing an unwanted play experience. Character interpretation and gameplay quality therefore require distinct evidence.

**Cross-source inference.** A future character case may involve powers, relationships, motives, social role, intelligence, or narrative influence. The user supplied corpus explicitly preserves that variety. The inspected studies make it implausible to equate character fidelity with a single list of combat powers, but they do not validate any particular decomposition of identity for this project. No anime or manga dossier was created here.

[FACTSCORE23](RESEARCH_LIBRARY.md#factscore23) concerns whether atomic factual claims in long text are supported by a source. [RAG20](RESEARCH_LIBRARY.md#rag20) studies retrieval with generation in knowledge-intensive NLP tasks. Those papers clarify a separation between the provenance of a character claim and the quality of a gameplay interpretation. They do not decide whether knowledge should be researched in advance, retrieved, learned in weights, stored in a catalog, or acquired another way.

## Undefined constraints and transfer

[ICA23](RESEARCH_LIBRARY.md#ica23) asks judgments about candidate constraints to learn a model with fewer user queries. [GENCON24](RESEARCH_LIBRARY.md#gencon24) studies generalizing constraint models across instances. These are credible adjacent research on incomplete specifications.

**Cross-source inference:** their candidate expression languages and queryable judgments are stronger assumptions than the undefined mechanics and evolving character concepts in this project. That observation does not rule out their later relevance.

[PCGML18](RESEARCH_LIBRARY.md#pcgml18), [PCGRL20](RESEARCH_LIBRARY.md#pcgrl20), [RAG20](RESEARCH_LIBRARY.md#rag20), and [CONCEPT18](RESEARCH_LIBRARY.md#concept18) demonstrate distinct ways knowledge or experience can affect generation or transfer. Their evaluations concern different outputs and cannot identify one necessary workflow for a name-to-unit experience. Any later comparison of prior learning and learning during use would first need to specify what is learned, from which evidence, and when.

## Evaluation and claims

The literature evaluates playability, balance, agent win rates, skill ordering, validity of a rule language, content diversity, factual support, personality, and human judgments. These are different constructs:

- [GAVEL24](RESEARCH_LIBRARY.md#gavel24) warns against treating proxy fitness as game interestingness. [MORTAR26](RESEARCH_LIBRARY.md#mortar26) records a disagreement between a skill-ordering proxy and an aggregate human score for one game pair.
- [PCGBENCH25](RESEARCH_LIBRARY.md#pcgbench25) defines quality, diversity, and controllability separately for each of its 12 problems. [ERA23](RESEARCH_LIBRARY.md#era23) shows that metric selection changes expressive-range analysis.
- [PERSONAS19](RESEARCH_LIBRARY.md#personas19) offers distinct artificial playstyles without establishing that a synthetic persona represents a human audience. [CSI14](RESEARCH_LIBRARY.md#csi14) concerns a user's creative support from a digital tool, a different construct from one design's quality; only its abstract was inspected here.

The [research methodology](METHODOLOGY.md) and library entries [M08](RESEARCH_LIBRARY.md#m08), [M09](RESEARCH_LIBRARY.md#m09), and [M12](RESEARCH_LIBRARY.md#m12) explain variance, reproducibility, and measurement validity without prescribing this project's evaluation.

**Open questions.** What observations could substantiate character faithfulness, gameplay coherence, balance, usefulness to designers, and cross-game applicability? Which claims require human judgment, simulations, source checks, or another type of evidence? No metric, baseline, ablation, or experiment was chosen in this phase.

## Jev and the two RLCD acronyms

- [TypeSafe AI's September 2026 announcement](https://typesafe.ai/blog/introducing-system-one-models-and-jev) identifies **Jev** as its non-generative “System One” model for typed probabilistic decisions. The vendor expands RLCD as **Reinforcement Learning for Calibrated Decisions** and describes training and performance; those claims are not peer-reviewed method evidence.
- [JEV26](RESEARCH_LIBRARY.md#jev26) is a September 2026 arXiv preprint, marked as under review, about detecting language-model alignment failures. It does not evaluate unit design or establish the model's training details.
- [RLCD23](RESEARCH_LIBRARY.md#rlcd23) expands RLCD as **Reinforcement Learning from Contrastive Distillation**, a separate language-model alignment method.

The acronym collision is resolved. Jev's possible role in this research remains open, and its product claims do not select a method.

## What remains unresolved

The inspected studies establish nearby tasks and some limits within them. They do not establish novelty or publishability for this project. In particular, the review has not determined the shared representation, treatment of unfamiliar concepts, need for training, knowledge acquisition, generation or adaptation architecture, evaluation design, scientific contribution, first experimental character, or first game. These are open questions for a later decision stage.

## New evidence from the 2026-10-01 search pass

This pass added citation-snowballing, a recent arXiv sweep, and a focused character-adaptation/human-evaluation review to the foundation established through 2026-09-28. Its substance is reported in the [search log](SEARCH_LOG.md) and summarized here.

- **Character→mechanics grounding is now better characterized.** The character-adaptation review confirms [CHARACTERS03](RESEARCH_LIBRARY.md#characters03) as the most direct conceptual precedent (explicit "translation" of an established character into actions, object use, attack damage, goals, weaknesses, and skill progression), and [NPC07](RESEARCH_LIBRARY.md#npc07) as the direct warning that character believability can conflict with gameplay goals. It adds evaluated fidelity protocols as the closest validated measures: [InCharacter24](RESEARCH_LIBRARY.md#incharacter24) (psychological-scale interviews, up to 80.7% alignment reported) and the [CHARM](RESEARCH_LIBRARY.md#charhalluc26) benchmark (which separates boundary-awareness from boundary-compliance failures in character hallucination), plus atomic-level persona-fidelity evaluation [OOC25](RESEARCH_LIBRARY.md#ooc25).

- **Fresh 2025–2026 generation and evaluation leads.** [consi25](RESEARCH_LIBRARY.md#consi25) (optimization-based 3D level generation with LLM-extracted constraints), [rpgagent26](RESEARCH_LIBRARY.md#rpgagent26) (multi-agent story-to-play generation with a coherence focus), [skills26](RESEARCH_LIBRARY.md#skills26) (diversity-of-skills generation), [omnic26](RESEARCH_LIBRARY.md#omnic26) (comprehensive role-playing agent benchmark, IEEE TPAMI 2026), [behavior-tree generation for Unreal](RESEARCH_LIBRARY.md#behbtsunreal26), and [spell-space exploration](RESEARCH_LIBRARY.md#spellspace25) extend the generation thread toward behavior and ability content. [runtime26](RESEARCH_LIBRARY.md#runtime26) and [multiobj25](RESEARCH_LIBRARY.md#multiobj25) extend PCG evaluation and instruction-aware control.

- **Evaluation-validity constraints are sharper.** [MORTAR26](RESEARCH_LIBRARY.md#mortar26) and [GAVEL24](RESEARCH_LIBRARY.md#gavel24) again show automated/proxy metrics disagree with human judgment; [Flake & Fried 2020](RESEARCH_LIBRARY.md#m12) sets the measurement-validity standard; [Lakens 2017](RESEARCH_LIBRARY.md#m10) implies balance and faithfulness claims are equivalence claims; [CSI14](RESEARCH_LIBRARY.md#csi14) measures tool support, not design quality. Human judgment remains the adjudicator, agents the triage.

- **Coverage limit stated explicitly.** No game-design-specific inter-rater-reliability validation instrument was found; character-adaptation and human-evaluation searches returned mostly medical/agricultural hits. This is recorded as a coverage limit, not as evidence that no such work exists.

These additions extend coverage of recent LLM-driven generation, character/agent fidelity benchmarks, and human-centered evaluation. They do not close the field (two snowball rounds hit an OpenAlex API cap; no index sweep was run), and claims in the added entries are abstract-level unless noted.

The search log names concrete unfinished searches in semantic adaptation, upgrades and progression, human evaluation, current proceedings, and citation trails. A claim that no prior work addresses character-to-unit design would exceed this review's present coverage.
