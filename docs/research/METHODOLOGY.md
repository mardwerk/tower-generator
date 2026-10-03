# Research methodology

This document records methods available for later research and paper development. It does not choose a technical solution, contribution, experimental character, baseline, metric, or study design. The current work is an exploratory mapping and scoping pass with explicit coverage limits, recorded in the [search log](SEARCH_LOG.md).

All scholarly references below use stable identifiers in the [master library](RESEARCH_LIBRARY.md). Their inspection depths matter: some guidance was read in selected primary sections, while other entries were inspected only through abstracts or bibliographic records. Recommendations here are project-level applications of that guidance, not findings that a particular research process guarantees success.

## Formulating an open problem

The user-provided ambition describes an eventual experience. It does not yet identify a scientific question or establish that a new method is needed. Keep separate the desired output, the phenomena to understand, the assumptions under investigation, and the claim that evidence could eventually support.

Useful research questions identify what is uncertain, the scope in which an answer applies, and what kind of evidence could distinguish plausible answers. A question can concern description, explanation, comparison, representation, measurement, or usefulness. These are possible question types, not a selection among this project's contributions.

[M05](RESEARCH_LIBRARY.md#m05) describes problematization as identifying and challenging assumptions in existing literature. Its abstract supports considering assumptions as a source of questions; its full procedure was not inspected. For this project, statements such as "all relevant games share this mechanic" or "character fidelity is captured by this score" would need evidence. They remain questions until supported.

An unstudied application, a new combination of known components, a useful artifact, and a generalizable scientific finding can have different value. None automatically establishes the others. [M07](RESEARCH_LIBRARY.md#m07) warns that apparent technical gains may have explanations other than the newly proposed component. A defensible contribution eventually needs a clear claim and evidence appropriate to it.

## Mapping the literature and substantiating a gap

[M01](RESEARCH_LIBRARY.md#m01) is a bibliographically verified mapping-method source awaiting substantive reading. [M02](RESEARCH_LIBRARY.md#m02) was inspected for its citation-search procedure and limitations. [M03](RESEARCH_LIBRARY.md#m03) and [M04](RESEARCH_LIBRARY.md#m04) provide reporting guidance through the official PRISMA resources. These sources support a transparent process; citing them does not make this review complete or PRISMA-compliant.

Record the search date, index or source, query wording, eligibility decisions, citation trail, access outcome, and inspection depth. Keep a paper under one primary topic in the master library, with tags for its other connections. Deduplicate manuscript and publication versions before drawing coverage conclusions.

Wohlin [M02](RESEARCH_LIBRARY.md#m02) emphasizes a diverse starting set, including communities that may not cite one another, and traceable rounds of backward and forward citation searching. This matters because game AI, computational creativity, natural-language understanding, requirements research, and human evaluation can use different terminology for related problems. Exhausting one citation cluster does not exhaust the problem.

For later gap claims, distinguish the following:

| Statement | Evidence needed |
| --- | --- |
| A source reports a limitation | The source passage, its context, and the evaluation setting to which it applies |
| Several approaches leave a dimension untested | A documented comparison of their actual evaluations |
| A topic was not found in this review | A bounded statement of searches, screening, access, and date |
| A proposed question is scientifically valuable | An argument about the knowledge it could establish, beyond its absence from the current library |

Absence from a search is not proof that no relevant work exists. A research gap remains provisional when important indexes, citation trails, or full texts are missing. Use surveys to learn vocabulary and locate primary papers; use original papers for important claims about methods and results.

[CHARACTERS03](RESEARCH_LIBRARY.md#characters03) is a direct conceptual precedent for expressing an existing character through possible actions and game mechanics. It does not test an automated method or a fidelity measure. A later gap claim must distinguish these kinds of work, rather than treating the lack of a tested procedure as a lack of prior character-to-mechanics analysis.

Inspection depth should remain visible. [M13](RESEARCH_LIBRARY.md#m13) is a verified abstract-level source on staged paper reading; its detailed three-pass procedure was not read here. An abstract can support preliminary relevance screening but usually cannot support detailed claims about evaluation validity or limitations.

## Comparing formulations without selecting a solution

Disciplined reasoning beyond familiar approaches can use three practices later:

1. Identify the assumptions a formulation needs, where they came from, and what would challenge them. M05 motivates this practice without proving that assumption challenges always produce useful research.
2. Use cross-field analogies by stating corresponding relations and where the correspondence fails. Gentner's structure-mapping theory, [M06](RESEARCH_LIBRARY.md#m06), distinguishes relational structure from superficial shared attributes. Its abstract was inspected; no adaptation algorithm follows from its inclusion here.
3. Compare alternative problem formulations using the same questions: what is given, what remains unknown, what output is claimed, whose judgment matters, and what evidence would count against the claim. This is a project-level reasoning aid informed by M05-M07, not a validated method for discovering novelty.

If later work compares learning before and during use, define the training objectives, data, and model changes before making claims about either. Training remains an open question here.

## Hypotheses, evaluation, and uncertainty

[M11](RESEARCH_LIBRARY.md#m11) distinguishes generating hypotheses from existing observations and testing predictions with new observations. Exploratory work can reveal useful patterns. Its conclusions should retain that exploratory status. Later confirmatory claims need a record of predictions and analysis choices made before the relevant outcomes were observed, with subsequent deviations identified.

A hypothesis should eventually be specific enough that some plausible observations would weaken it. At this stage, preserve uncertainties and competing explanations rather than inventing hypotheses to fit an architecture that has not been selected.

Baseline comparisons ask whether a claimed improvement survives comparison with an appropriate alternative under comparable conditions. Ablations ask which parts account for an observed effect. [M07](RESEARCH_LIBRARY.md#m07) explains how bundled changes and unequal tuning can obscure the source of gains. These principles do not choose the project's baselines or ablations.

[M08](RESEARCH_LIBRARY.md#m08) studies variability from hyperparameters, seeds, environments, and implementations in selected deep-RL settings. If later work involves stochastic models, search, training, or automated players, the relevant variability and resource differences must be reported. The paper does not establish a universal trial count, statistical test, or suitable agent for this project, and its inclusion does not favor reinforcement learning.

Measurement requires its own justification. [M12](RESEARCH_LIBRARY.md#m12) links opaque measurement decisions to threats to validity. A score called "fidelity", "creativity", "coherence", or "usefulness" needs evidence that its observations support that interpretation. Agreement, repeatability, and apparent precision are not by themselves evidence that the intended quality was measured.

The primary evaluation in [MORTAR26](RESEARCH_LIBRARY.md#mortar26) ranks games by how closely five automated agents follow an expected skill order. In its small user study, the authors' aggregate human score favored the higher-ranked game in two pairs and the lower-ranked game in a third. This is a concrete reason to state what a proxy measures and test how it relates to the human judgment claimed, if such a claim is eventually made. The study does not identify a suitable metric for character-grounded unit designs.

When comparing future evaluation approaches, keep these questions distinct:

| Validity question | What it asks |
| --- | --- |
| Construct validity | Do observations support the intended interpretation of the measured quality? |
| Internal validity | Could an alternative explanation account for the observed difference? |
| Statistical-conclusion validity | Does the analysis justify the reported uncertainty and comparison? |
| External validity | To which characters, games, users, and conditions could the claim reasonably extend? |

These are methodological questions, not a chosen evaluation framework. The varied [initial corpora](CORPORA.md) make scope especially consequential: inclusion in a list does not establish representative sampling or cross-game generalization.

[M10](RESEARCH_LIBRARY.md#m10) explains why a nonsignificant result alone cannot establish no meaningful difference. Its equivalence-testing primer is one available tool, requiring justified bounds and sufficient information. The eventual analysis may require another approach. Preserve the distinctions between a supported effect, evidence of practical equivalence, and an inconclusive result.

## Reproducibility and claim limits

[M09](RESEARCH_LIBRARY.md#m09) reports on the NeurIPS 2019 reproducibility program. The authors identify underspecified procedures and metrics, selective reporting, adaptive overfitting, and overclaiming as problems. Their conference case study reports encouraging participation but explicitly lacks conclusive evidence that its interventions caused improvements in research quality. A checklist is therefore useful documentation, not a quality certificate.

Later research records should contain the information needed to assess and reproduce the particular claim. Depending on the work, this may include source revisions, data inclusion decisions, annotation guidance, model and software versions, prompts, tuning budgets, randomization, analysis choices, and failed attempts. These are conditional requirements, not software infrastructure to build now.

Keep negative results, inconclusive observations, and practical failures visible. Distinguish failure to implement an idea, failure under a tested condition, and evidence against a broader claim. Report access gaps and changes in interpretation rather than replacing them with retrospective certainty.

Reproducibility and scientific validity are related but separate. A procedure can produce the same answer consistently while relying on an unsuitable measurement or supporting only a narrow claim.

## Developing a paper from evidence

The eventual paper should explain the problem, the relationship to prior work, a defensible contribution, supporting evidence, and limits. Its title should follow the contribution that the work actually establishes. No target venue or publication claim is selected in this phase.

Simon Peyton Jones's [How to write a great research paper](https://www.microsoft.com/en-us/research/academic-program/write-great-research-paper/) is non-paper educational guidance. The official slides were inspected. They recommend writing to clarify reasoning, stating contributions concretely, connecting claims to evidence, using comprehensible examples, crediting prior work, and seeking criticism. These are useful writing practices; the slides' preferred section order and page allocations are not universal requirements.

For later writing, a claim-to-evidence check can ask whether every central claim has support, whether the support has the same scope as the claim, and whether an alternative explanation remains. This is a practical application of M07 and M09 and the writing guide. It is not a request to draft unsupported results or contribution sections now.

Relevant communities include game AI and procedural content generation, computational creativity, NLP and character understanding, human-computer interaction, requirements and knowledge representation, and machine learning methodology. Examples of publication outlets encountered in these areas include IEEE CoG, AIIDE, FDG, IEEE Transactions on Games, ICCC, ACL-family venues, CHI, and general AI/ML conferences and journals. This is a descriptive orientation, not a target-venue choice; current calls and contribution expectations would require later verification.

## Limits of this methodology review

This session inspected selected primary material for M02, M07-M10 and the non-paper writing guide. M01 remains bibliographic only. M03-M04 were inspected through official reporting resources rather than their complete explanatory papers. M05-M06 and M11-M13 have abstract-level support. The master library gives exact depths and access limitations.

No complete forward-citation sweep of methodology literature was performed. Measurement guidance specific to creative game-design evaluation, recent methodological updates, and unresolved full-text access remain part of the [search coverage record](SEARCH_LOG.md). These limits constrain claims of comprehensiveness; they do not settle any of the project's open technical decisions.
