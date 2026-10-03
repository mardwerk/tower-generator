# Pre-experiment protocol: Gojo -> Plants vs. Zombies tower

**Experiment:** 001_gojo_pvz — generate a Satoru Gojo (Jujutsu Kaisen) tower for
Plants vs. Zombies and triage it. **Status:** preregistered before outcomes.
**Date:** 2026-10-02. **Reference:** docs/RECOMMENDATION.md Sections 3–4.

## Rationale for preregistration
Per Flake & Fried 2020 and Lakans 2017, faithfulness and balance claims are
*equivalence* claims; per GAVEL24 and MORTAR26 proxy metrics disagree with human
judgment. Recording questions, metrics, and the human protocol before seeing
outcomes avoids post-hoc rationalization. This is cheap and must precede any
outcome-driven decision.

## 1. Research questions (fixed before running)

1. **Faithfulness:** Does a constraint-guided translation of a character dossier
   into a PvZ tower cite sourced character-ability claims for its effects, and
   can each cited claim be traced back through the audit trace?
2. **Balance:** Does the same character-grounded tower survive PvZ-like waves
   under MCTS agent triage (win-rate, survival rate, placement robustness)?
3. **Faithfulness-vs-playability trade-off (NPC07):** Does a high-fidelity,
   high-power configuration (Purple/Domain effects) fail triage while a
   faithful-but-constrained variant passes — or does faithful generation happen
   to yield viable units?
4. **Controllability (PCGBENCH25):** How much does perturbing the numeric
   parameters of a faithful concept change triage outcomes — i.e., can a human
   control balance without losing character grounding?

## 2. Hypotheses (directional; no effect sizes claimed)

- H1: The audit trace will show ≥80% of base/branch effects citing at least one
  source-verified ability claim (CHARHALL26/OOC25 atomic-fidelity benchmark).
- H2: The faithful high-power variant will have lower triage win-rate than a
  faithful-but-constrained variant (testing NPC07's conflict warning).
- H3: Agent triage win-rate will correlate with human shortlist judgment, but
  disagreement will occur on at least one pair (mirroring MORTAR26), justifying
  human arbitration.
- H4: Diversified variants will differ in triage pass rate while sharing the
  same character provenance (controllability probe).

## 3. Materials

- Character: Satoru Gojo (version selected in dossier; unit explains its version
  per THON19 semantics).
- Game: Plants vs. Zombies; tower/ability vocabulary per the game's public
  design (tower cost, range, cooldown, damage, upgrade branches).
- Generator: constraint-guided pipeline — dossier -> formal tower spec
  (lib/spec.py) -> well-formedness check (lib/generator.py check_constraints)
  -> diversity search -> triage.
- Triager: MCTS placement search over lanes/columns, 40 simulations, then
  30 full simulations at the best placement (TDSTRATEGY24/TOWERMIND26 precedent).

## 4. Metrics (defined per target game, PCGBENCH25)

**Faithfulness (Layer 4):**
- citation_rate = ability_claims_cited_by_effects / ability_claims_present
- inferred_effects = count of effects with no cited claim (hallucination proxy)
- faithfulness_score = citation_rate (max 1.0)

**Balance (Layer 3, agent triage):**
- win_rate = fraction of rollout/seeded simulations where waves are cleared
- survival_rate = 1 - (zombies_home / 6)
- pressure_per_wave = max zombies on board averaged across detailed runs
- cost_efficiency = dps / cost (lib/spec.TowerDraft)

**Trade-off:** faithfulness_score vs. triage win_rate plotted per variant.

**Controllability:** variance of triage win_rate across the numeric
perturbations of a single faithful concept.

## 5. Pass/fail thresholds (fixed before running)

- Faithfulness pass: citation_rate >= 0.80 (H1).
- Balance pass: triage win_rate >= 0.50 and survival_rate >= 0.60 (H2, H4).
- Trade-off finding: a documented monotonic or non-monotonic relationship
  between faithfulness_score and win_rate across variants; if faithful high-power
  config fails balance while faithful-constrained passes, H2 is supported.

## 6. Human protocol (CSI14 usefulness + measurement validity)

A shortlisted set (max 4 variants: base + up to 2 faithful variants that pass
triage + 1 high-power variant) is presented to a human judge familiar with PvZ.
Judgment criteria, anchored on CSI14 "supportiveness of tool support" and the
project's own vocabulary:

- Faithfulness to Gojo (0–5, atomic claims visible beside each variant)
- Coherence / well-formedness as a PvZ tower (0–5)
- Originality (0–5)
- Usefulness / playability (0–5)
- Perceived balance vs. other premium towers (0–5)

Format: one-page spec per variant (in-game stats + effect description + the
audit-trace provenance column). Judge rates independently; then a brief
consensus discussion is recorded. Inter-rater reliability is a known gap in
this library (coverage limit) — future work should add a second rater and a
game-design RRI instrument.

## 7. Data handling and reporting

- All outputs (drafts, triage reports, audit traces) written to
  experiments/001_gojo_pvz/{outputs,triage_results}/ with timestamps.
- Results are reported as the raw triage numbers plus the human scores; no
  cherry-picking: the full set of diversified variants is shown, not just the
  best.
- Agent triage is labeled triage only (not final adjudication). Human judgment
  is the arbitrator.

## 8. What would change this protocol

If an LLM translator is wired in (PromptTranslator), the same questions/metrics
apply; only the provenance extraction step needs review. If the simulator is
replaced by a real PvZ playtest harness, metrics are re-derived on the same
formulas. Any threshold change after seeing outcomes must be flagged in a
post-hoc note in this document.
