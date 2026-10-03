"""First experiment: Satoru Gojo (Jujutsu Kaisen) -> Plants vs. Zombies tower.

Per docs/RECOMMENDATION.md Section 3. Runs the full layered pipeline:
  Layer 1 — character grounding & dossier (bundled profile; replace retriever to use live sources)
  Layer 2 — constraint-guided generation (worked-example translator + diversify + check_constraints)
  Layer 3 — MCTS agent triage (placement search; thresholds calibrate empirically)
  Layer 4 — audit-trace explanation (provenance back to atomic claims)

Outputs:
  - preregistered protocol: docs/experiments/001_gojo_pvz/preregistration/protocol.md
  - generated tower drafts: docs/experiments/001_gojo_pvz/outputs/
  - triage results: docs/experiments/001_gojo_pvz/triage_results/

Demonstration note: the translator is a worked example so the pipeline runs
offline; wire an LLM into PromptTranslator to make generation fully generative.
"""

import json
import os
import sys
from datetime import datetime


class EnumEncoder(json.JSONEncoder):
    def default(self, obj):
        from enum import Enum
        if isinstance(obj, Enum):
            return obj.value
        return super().default(obj)


ROOT = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, ROOT)
import lib


OUTPUT_DIR = os.path.join(ROOT, "outputs")
TRIAGE_DIR = os.path.join(ROOT, "triage_results")


def run_experiment(seed: int = 42, simulations: int = 10, diversify_count: int = 3):
    os.makedirs(OUTPUT_DIR, exist_ok=True)
    os.makedirs(TRIAGE_DIR, exist_ok=True)

    # ---- Layer 1: grounding ----
    retriever = lib.BundledProfileRetriever()
    retriever._register("Satoru Gojo", [])  # bundled profile used directly below
    dossier = lib.build_gojo_dossier()
    print(f"[L1] Dossier: {dossier.character} (version {dossier.selected_version}), "
          f"{len(dossier)} claims, versions: {[v.version_key for v in dossier.versions]}")

    # ---- Layer 2: generation ----
    translator = lib.CharacterTranslator(dossier)
    base, trace = translator.translate()
    checks = [lib.check_constraints(base, dossier)]
    drafts = [base]
    print(f"[L2] Base draft: {base} efficiency={base.cost_efficiency:.2f}")
    # Branches are expressed in the base tower's branch registry.

    variants = lib.diversify(base, diversify_count, dossier)
    for v in variants:
        cv = lib.check_constraints(v, dossier)
        checks.append(cv)
        drafts.append(v)
        print(f"[L2] Variant: {v.name} cost={v.cost} efficiency={v.cost_efficiency:.2f} "
              f"constraint_check={cv}")

    # ---- Layer 3: triage ----
    triager = lib.AgentTriager(simulations=simulations, difficulty="normal")
    triage_reports = {}
    for i, d in enumerate(drafts):
        if isinstance(d, lib.UpgradeBranch):
            continue
        rep = triager.triage(d)
        triage_reports[d.name] = rep
        print(f"[L3] Triage {d.name}: win_rate={rep['win_rate']} survival={rep['avg_survival_rate']} "
              f"top_placement={rep['top_placement']} passes={rep['triage']['passes_win_rate']}")

    # ---- Layer 4: audit trace ----
    audit = lib.build_audit_trace(dossier, base)
    print(f"[L4] Faithfulness summary: {audit.faithfulness_summary()}")

    # ---- Write outputs ----
    now = datetime.now().strftime("%Y%m%d-%H%M%S")
    base_out = os.path.join(OUTPUT_DIR, f"gojo_tower_base_{now}.json")
    out = base.to_dict()
    out["branches"] = base.branch_to_dict()
    json.dump(out, open(base_out, "w"), indent=2, cls=EnumEncoder)
    trace_out = os.path.join(OUTPUT_DIR, f"gojo_audit_trace_{now}.json")
    json.dump(audit.to_dict(), open(trace_out, "w"), indent=2, cls=EnumEncoder)
    triage_out = os.path.join(TRIAGE_DIR, f"gojo_triage_{now}.json")
    json.dump(triage_reports, open(triage_out, "w"), indent=2, cls=EnumEncoder)
    print(f"\nWrote: {base_out}, {trace_out}, {triage_out}")
    return dossier, base, drafts, triage_reports, audit


if __name__ == "__main__":
    run_experiment()
