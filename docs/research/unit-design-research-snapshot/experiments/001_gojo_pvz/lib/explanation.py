"""Audit-trace and explanation layer (Layer 4).

Reconstructs every generated mechanic back to the character-dossier claims it
expresses, so character fidelity is auditable rather than asserted (per
CHARHALL26 / OOC25 atomic-fidelity practice). Produces a structured trace that
separates verified provenance from generator inference.
"""

from dataclasses import dataclass, field
from typing import Dict, List, Optional

from .retrieval import CharacterDossier, AtomicClaim
from .spec import TowerDraft, UpgradeBranch, Effect


@dataclass
class TraceElement:
    kind: str
    source_id: str
    statement: str
    confidence: str
    category: str = ""
    provenance: List[str] = field(default_factory=list)

    def __dict__(self):
        d = dict(self.__dataclass_fields__)
        out = {}
        for k in d:
            v = getattr(self, k)
            out[k] = list(v) if isinstance(v, list) else v
        return out


@dataclass
class AuditTrace:
    character: str
    version: str
    claims: List[AtomicClaim]
    elements: List[TraceElement] = field(default_factory=list)

    def to_dict(self) -> dict:
        return {
            "character": self.character, "version": self.version,
            "claims": [c.__dict__ for c in self.claims],
            "elements": [e.__dict__() for e in self.elements],
        }

    def faithfulness_summary(self) -> dict:
        """Aggregate: how many claimed abilities vs. how many effects cite them."""
        ability_ids = {c.claim_id for c in self.claims if c.category == "abilities"}
        cited = set()
        inferred = 0
        for e in self.elements:
            if e.kind in ("effect", "branch"):
                if any(p in ability_ids for p in e.provenance):
                    cited.add(e.source_id)
                else:
                    inferred += 1
        return {
            "ability_claims_present": len(ability_ids),
            "ability_claims_cited_by_effects": len(cited),
            "effects_with_inferred_provenance": inferred,
            "citation_rate": round(len(cited) / max(len(ability_ids), 1), 3),
        }


def build_audit_trace(dossier: CharacterDossier, tower: TowerDraft) -> AuditTrace:
    trace = AuditTrace(
        character=dossier.character,
        version=dossier.selected_version,
        claims=dossier.sorted_claims,
    )

    def add(kind, eid, statement, conf, category="", provenance=None):
        trace.elements.append(TraceElement(kind, eid, statement, conf, category,
                                           provenance or []))

    # Add the dossier's own claims
    for c in dossier.sorted_claims:
        add("claim", c.claim_id, c.statement, c.confidence, c.category)

    # Add base effects with provenance
    for i, e in enumerate(tower.effects):
        add("effect", f"base_effect_{i}", e.explanation,
            "plausibly-inferred" if e.provenance else "generator-inferred",
            "abilities" if e.provenance else "",
            provenance=list(e.provenance))

    # Add branches
    for bname, branch in tower.branches.items():
        add("branch", f"branch:{bname}", f"upgrade branch '{bname}', cost {branch.cost}",
            "generator-inferred", "", provenance=list(branch.provenance))
        for e in branch.effects:
            add("effect", f"branch:{bname}:{e.name}", e.explanation,
                "plausibly-inferred" if e.provenance else "generator-inferred",
                "abilities" if e.provenance else "",
                provenance=list(e.provenance))

    return trace


if __name__ == "__main__":
    from .retrieval import build_gojo_dossier
    from .generator import CharacterTranslator
    d = build_gojo_dossier()
    base, _ = CharacterTranslator(d).translate()
    tr = build_audit_trace(d, base)
    print(tr.character, "->", base.name)
    for el in tr.elements[:6]:
        print(f"  [{el.kind}] {el.source_id}: {el.statement[:70]} ({el.confidence})")
    print("faithfulness:", tr.faithfulness_summary())

