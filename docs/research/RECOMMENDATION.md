# Research judgment and recommendation

**Status:** research-led judgment (not yet validated by experiment). **Date:** 2026-10-02. This document records what the established literature points to, my assessment of the most promising direction, and the concrete first step. It follows the project's convention of distinguishing source-reported findings from cross-source inferences and open questions.

---

## 1. What the literature now establishes

### 1.1 The core task is unpopulated — but its subproblems are not

No prior study tests automated **character-name → TD/RTS unit with abilities / statistics / upgrades / explanation, under initially undefined game constraints**. The closest computational ancestor is [RTSUNITS22](RESEARCH_LIBRARY.md#rtsunits22), which generates microRTS units by searching six numeric attributes and hand-authored causes/effects under MCTS evaluation — but it uses no reference character, no ability concepts, no explanation, and no cross-game adaptation. The closest conceptual ancestor is [CHARACTERS03](RESEARCH_LIBRARY.md#characters03), which *explicitly* frames "translating an established character into mechanics" and names concrete translation slots (possible/impossible actions, object interaction, attack damage, damage-triggered strength, stealth, goals, weaknesses, skill progression, and progression changing later choices). **Source-reported:** the gap between these two is a genuine gap — the slot vocabulary from CHARACTERS03 is never automated, and RTSUNITS22's slots never get character grounding.

**Cross-source inference:** the project's task is not a novel invention so much as a *controlled combination* of three established components (character grounding, constrained generation, triaged evaluation) applied to an input specification (a character name) that is weaker than what any of those components assumes. That is exactly where novelty is defensible: the novelty lives in the *bridge* (character → constraints → generation) and in proving it works, not in discovering a new genre of problem.

### 1.2 Each subproblem has precedent

| Subproblem | Established precedents | What they give the project |
|---|---|---|
| Character grounding / provenance | [RAG20](RESEARCH_LIBRARY.md#rag20), [CROSS24](RESEARCH_LIBRARY.md#cross24), [INCHARACTER24](RESEARCH_LIBRARY.md#incharacter24), [AUSTEN24](RESEARCH_LIBRARY.md#austen24), [ROLELLM24](RESEARCH_LIBRARY.md#rolellm24), [MIMIC22](RESEARCH_LIBRARY.md#mimic22), [FICTIONRAG26](RESEARCH_LIBRARY.md#fictionrag26), [DREAM26](RESEARCH_LIBRARY.md#dream26) | Retrieval + structured profiling + factual verification pipelines; the hard problem is source-grounding and version selection |
| Character→mechanics translation | [CHARACTERS03](RESEARCH_LIBRARY.md#characters03), [NPC07](RESEARCH_LIBRARY.md#npc07), [THON19](RESEARCH_LIBRARY.md#thon19) | Named translation slots; version-ambiguity semantics; the warning that believability can *conflict* with gameplay goals |
| Constraint-guided generation | [MECHGEN14](RESEARCH_LIBRARY.md#mechgen14), [GGDG24](RESEARCH_LIBRARY.md#ggdg24), [BOSS17](RESEARCH_LIBRARY.md#boss17), [MORTAR26](RESEARCH_LIBRARY.md#mortar26), [CONSI25](RESEARCH_LIBRARY.md#consi25), [SCRIPT25](RESEARCH_LIBRARY.md#script25), [BEHBTS26](RESEARCH_LIBRARY.md#behbtsunreal26) | Composable mechanic representations, grammar-guided LLM output, program-synthesis well-formedness, LLM+quality-diversity mechanic evolution, optimization-based constraint assembly |
| Partial/undefined constraints | [ICA23](RESEARCH_LIBRARY.md#ica23), [GENCON24](RESEARCH_LIBRARY.md#gencon24), [TANAGRA11](RESEARCH_LIBRARY.md#tanagra11), [AUTOBG26](RESEARCH_LIBRARY.md#autobg26) | Active-elicitation frameworks — but all assume a *candidate language* the user can judge, stronger than the project's "no vocabulary yet" assumption |
| Evaluation / playtesting | [PCGBENCH25](RESEARCH_LIBRARY.md#pcgbench25), [MCTS15](RESEARCH_LIBRARY.md#mcts15), [PERSONAS19](RESEARCH_LIBRARY.md#personas19), [RUNTIME26](RESEARCH_LIBRARY.md#runtime26) | Problem-specific metrics, MCTS agent skill tiers, runtime agent triage loops |
| Evaluation validity | [GAVEL24](RESEARCH_LIBRARY.md#gavel24), [MORTAR26](RESEARCH_LIBRARY.md#mortar26), [CSI14](RESEARCH_LIBRARY.md#csi14), [Flake & Fried 2020](RESEARCH_LIBRARY.md#m12), [Lakens 2017](RESEARCH_LIBRARY.md#m10), [M09](RESEARCH_LIBRARY.md#m09) | Proxy metrics do not suffice (GAVEL24); a skill-ordering proxy disagreed with a human score in MORTAR26; CSI measures tool *support* not design quality; measurement validity is the controlling constraint; balance/faithfulness claims are *equivalence* claims |
| Character drift & fidelity | [CHARHALL26](RESEARCH_LIBRARY.md#charhalluc26), [OOC25](RESEARCH_LIBRARY.md#ooc25), [persona_weaver26](RESEARCH_LIBRARY.md#persona26), [SKILLS26](RESEARCH_LIBRARY.md#skills26) | LLM output drifts out of character; atomic-level fidelity testing detects what aggregate scores miss; LLM content is behaviorally homogeneous by default |

### 1.3 Three empirical risks, each documented in the literature

1. **Fidelity ≠ playability.** [NPC07](RESEARCH_LIBRARY.md#npc07) explicitly cautions that adding believable-character patterns would not necessarily improve the experience, because believability can conflict with gameplay goals. [MORTAR26](RESEARCH_LIBRARY.md#mortar26) and [GAVEL24](RESEARCH_LIBRARY.md#gavel24) independently show simulation/proxy measures disagree with human judgment. → A unit can be character-faithful and still a bad unit. **Faithfulness and playability require separate instruments.**
2. **Automated proxies are insufficient and must be human-validated.** [MORTAR26](RESEARCH_LIBRARY.md#mortar26): a five-agent skill ordering favored the higher-ranked game in two of three human-comparison pairs, and disagreed in one. [GAVEL24](RESEARCH_LIBRARY.md#gavel24): the authors explicitly warn their proxy metrics are far from sufficient evidence of interesting games. → **Human expert judgment is the adjudicator; agents are triage only.**
3. **LLMs hallucinate abilities and drift.** [CHARHALL26](RESEARCH_LIBRARY.md#charhalluc26) (CHARM benchmark): most hallucination is a *compliance* failure (the model knows the boundary but answers anyway), not a recognition failure. [OOC25](RESEARCH_LIBRARY.md#ooc25): atomic-level evaluation catches subtle persona drift whole-response scores miss. [persona_weaver26](RESEARCH_LIBRARY.md#persona26): unmitigated LLM character generation produces behaviorally homogeneous, overly-prosocial outputs. → **The pipeline needs explicit constraint enforcement plus atomic-fidelity checks; nothing should be accepted as faithful without it.**

### 1.4 What remains open

- No validated composite for faithfulness + coherence + balance + usefulness (the sub-problems have instruments, but nothing combines them).
- No game-design-specific inter-rater-reliability instrument was found (character adaptation and human-evaluation searches returned mostly medical/agricultural hits); the library must state this as a coverage limit.
- How to elicit constraints from a designer who does not yet know the vocabulary (ICA23/GENCON24 assume a candidate language) is not settled.
- Which claims require human judgment vs. simulation vs. source checks is still open.

---

## 2. The most promising approach (my judgment)

**Recommendation: a tiered, human-in-the-loop, constraint-guided generation pipeline with three explicit layers and an audit-trace explanation layer.**

```
character name → [1] character grounding & dossier → [2] constraint-guided mechanic generation
     → [3] triaged evaluation (agent triage, then human arbiter) → output unit + explanation trace
```

**Layer 1 — Character grounding and dossier (RAG + knowledge graph).** Retrieve character claims from canonical source pages, decompose them into atomic claims with provenance (the FACTSCORE23 approach), and assemble a structured dossier: powers/abilities, relationships, narrative role, weaknesses, progression. **Version selection is a required first decision** (THON19): the same name can have many versions across media; a generated unit must state which version it is grounded in. **Why this layer first:** [ROLELLM24](RESEARCH_LIBRARY.md#rolellm24), [MIMIC22](RESEARCH_LIBRARY.md#mimic22), and [FICTIONRAG26](RESEARCH_LIBRARY.md#fictionrag26) all show that retrieval from source material is the hard part of faithful character work; generation without retrieval hallucinates.

**Layer 2 — Constraint-guided mechanic generation (LLM + formal representation + search).** The LLM's role is to *translate* character dossier content into a formal, composable mechanic specification (action-like effects per [MECHGEN14](RESEARCH_LIBRARY.md#mechgen14); grammar-guided output per [GGDG24](RESEARCH_LIBRARY.md#ggdg24); behavior-tree structure per [BEHBTS26](RESEARCH_LIBRARY.md#behbtsunreal26)). Then a constraint solver / planner checks well-formedness and consistency, and a search/evolution step (MORTAR26's quality-diversity approach, or CONSI25's optimization-based assembly) explores the neighborhood for diversity. **Why this hybrid:** pure free-form LLM generation is unbounded and inconsistent (CHARHALL26, OOC25); pure constraint programming lacks semantic character grounding; the hybrid uses the LLM where it is strong (semantic translation) and formal methods where they are strong (well-formedness, search).

**Layer 3 — Triaged evaluation.** Automated agent-based triage first: MCTS agents at differentiated skill budgets (MCTS15) or runtime autonomous agents (RUNTIME26) filter out broken/imbalanced drafts cheaply. Then **human expert judgment** adjudicates the shortlist for faithfulness, coherence, originality, balance, and usefulness. This ordering follows GAVEL24 and MORTAR26: proxies are triage, humans are judges. Metrics are defined **per target game** (PCGBENCH25): there is no universal quality score.

**Layer 4 — Explanation / audit trace (provenance + atomic facts).** Every generated mechanic is explained by tracing it back to the character dossier claim(s) it expresses; factual claims are verifiable against source (FACTSCORE23-style). This is the "explanation of character fidelity" output the project promises, and it is also the mechanism that makes faithfulness *auditable* rather than asserted.

**Alternative approaches considered and rejected:**

- **Pure LLM free-generation** (LLM → unit, no constraints): fails on consistency, drift, and homogeneity (CHARHALL26, OOC25, persona_weaver26). **Rejected as the core mechanism**; the LLM is a translator, not the generator.
- **Pure learned generation / RL (PCGRL20):** presupposes a fixed editing vocabulary, rewards, and evaluative criteria. The project's "initially undefined" constraints are explicitly stronger. **Rejected as the core mechanism**; useful later for tuning within a stable vocabulary.
- **Pure constraint programming (MECHGEN14 alone):** needs design requirements and an action vocabulary supplied in advance; cannot ground in a character. **Rejected as the core mechanism**; kept as the well-formedness layer.
- **Learning a representation from the 27 games first:** requires reverse-engineering and assumes a shared schema the literature says is unproven (STATE_OF_ART.md, open question). **Deferred**: the pipeline starts with one game/one character.

**This is a defensible, falsifiable contribution, not a claim of solved novelty:** the reproducible framework, the preregistered evaluation protocol, and the first empirical evidence on the faithfulness-vs-playability trade-off (which NPC07 hypothesized qualitatively and which MORTAR26/GAVEL24 made urgent).

---

## 3. The recommended first experiment (judgment on concretization)

**Character:** **Satoru Gojo** (*Jujutsu Kaisen*). **Why:** maximal ability documentation (reliable RAG grounding), a clear "overpowered/special-grade" status that stresses balance testing, and a rich ability vocabulary (Limitless, Hollow Purple, Domain Expansion) that maps naturally onto tower-type / upgrade-type units.

**Game:** **Plants vs. Zombies**. **Why:** the most-studied TD in the corpus (TDSTRATEGY24, TOWERMIND26); a well-known tower/ability vocabulary; and automated-playtesting precedent means agent-based triage is immediately feasible.

**First experiment scope:** generate a single Gojo-inspired tower unit with a base form and two upgrade branches; evaluate (a) faithfulness to the dossier (atomic-fidelity check vs. source-verified claims, per CHARHALL26/OOC25), (b) playability and balance under MCTS agent triage (MCTS15 skill tiers), (c) human judgment of faithfulness, coherence, originality, usefulness on a shortlisted set (CSI14 for usefulness-support, per the measurement-validity standard Flake & Fried 2020, with equivalence framing per Lakans 2017).

**Precondition:** preregister the evaluation questions, metrics, and the human-study protocol before running outcomes (M11, M09). This is the single most important methodological step, and it is cheap to do.

---

## 4. Concrete next steps and what is needed from Kyle

1. **Approve the layered approach** (Section 2) and the first experimental case (Gojo → Plants vs. Zombies, Section 3).
2. **Authorize the preregistration step** for the first study's evaluation protocol before any outcomes are seen.
3. **Optional but valuable:** a one-time access check on the two SSRN/ACM records the character-evaluation review flagged (NPCREAL25; ACS/other) so the human-evaluation section of the library is not left on access-failure.

---

## 5. What this judgment is NOT

- It is not proof that the approach will succeed. NPC07's faithfulness-vs-playability conflict claim is a hypothesis this experiment is designed to test, not a settled fact.
- It is not a claim that the literature search is complete. Two snowball rounds hit an OpenAlex API cap, and no Scopus/Web of Science/Google Scholar sweep was run; newer citing papers may be missing.
- It does not select all downstream decisions (full representation schema, all 27-game corpus reverse-engineering, training strategy). Those remain open pending the first experimental results.
