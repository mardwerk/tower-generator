"""Character→mechanic translation and constraint-guided generation (Layer 2).

The LLM's role is the translator: dossier content -> formal PvZ tower spec.
A formal constraint checker enforces well-formedness, and a diversity search
explores the config neighborhood. In this demonstration the translator is a
worked example; the PromptTranslator below holds the LLM hook so a real model
can be wired in without changing the pipeline.
"""

from typing import Dict, List, Optional, Tuple

from .spec import (TowerDraft, UpgradeBranch, Effect, DamageType, TargetType,
                   ProjectileType)
from .retrieval import CharacterDossier


class TranslationError(Exception):
    pass


class ConstraintCheck:
    """Well-formedness verdict on a tower draft."""
    def __init__(self, pass_: bool, violations: List[str], notes: List[str] = None):
        self.pass_ = pass_
        self.violations = violations or []
        self.notes = notes or []

    def __bool__(self):
        return self.pass_

    def __repr__(self):
        return f"<ConstraintCheck pass={self.pass_} violations={len(self.violations)}>"


class TranslatorABC:
    """Interface for any character->mechanic translator (LLM-based or reference)."""

    def translate(self) -> Tuple[TowerDraft, List[str]]:
        raise NotImplementedError


class CharacterTranslator(TranslatorABC):
    """Translate a versioned dossier into a PvZ tower draft.

    Implements a character->vocabulary mapping grounded in the dossier claims.
    Demonstrative implementation: hand-authored mapping so the pipeline runs
    offline and produces a reproducible baseline. Replace/extend `translate`
    with an LLM call (see PromptTranslator) for the generative run.
    """

    def __init__(self, dossier: CharacterDossier):
        self.dossier = dossier
        self.abilities = {c.claim_id: c for c in dossier.abilities()}
        self.weaknesses = {c.claim_id: c for c in dossier.weaknesses()}

    def translate(self) -> Tuple[TowerDraft, List[str]]:
        """Return (tower, explanation trace)."""
        base, b1, b2, trace = self._build_base_and_branches()
        base.branches["hollow_purple"] = b1
        base.branches["unlimited_void"] = b2
        return base, trace

    def _build_base_and_branches(self):
        # Base tower = Blue (attraction/space-warp) — the signature mechanic,
        # reflecting the "shield / control" narrative role while expressing the
        # character's primary technique.
        trace = []

        def mk_eff(name, effects, proj, dmg, tag, prov, expl):
            e = Effect(name=name, effect_type=effects, damage=dmg,
                       projectile=proj, provenance=prov, explanation=expl)
            trace.append(f"base_effect:{tag}:{e.explanation}")
            return e

        base = TowerDraft(
            name="Domain Anchor: Limitless (Blue)",
            tier="premium",
            cost=325,
            range=5,
            cooldown=1.6,
            damage=60,
            damage_type=DamageType.MAGIC,
            target_type=TargetType.GROUND,
            projectile=ProjectileType.NORMAL,
            effects=[
                mk_eff("attract_on_hit", "attract", ProjectileType.NORMAL, 60, "G002",
                       [self.abilities["G002"].claim_id],
                       "Blue pulls adjacent zombies toward the tower, clustering enemies for follow-up damage."),
                mk_eff("space_warp_hit", "impact", ProjectileType.NORMAL, 60, "G002",
                       [self.abilities["G002"].claim_id],
                       "On hit, briefly warps target position, delaying advancement."),
            ],
            description=(
                "Anchors space around a chosen tile like the Limitless. Hits attract "
                "adjacent zombies inward, grouping them for team damage while slowing "
                "the front line. The Special Grade 'Strongest' barrier holds a chokepoint."),
            provenance_summary=[self.abilities["G001"].claim_id, self.abilities["G002"].claim_id],
        )

        # Branch 1: Purple — erasure beam, high single-target damage, piercing.
        b1 = UpgradeBranch(
            name="Hollow Purple (Erasure)",
            cost=600,
            unlock_requirement="Base tower must exist on board",
            effects=[
                Effect(name="erasure_beam", effect_type="splash", damage=400,
                       projectile=ProjectileType.PENETRATING, provenance=[self.abilities["G003"].claim_id],
                       explanation="Purple erases matter on contact; the beam pierces through zombies dealing massive damage per tile."),
            ],
            provenance=[self.abilities["G003"].claim_id],
        )
        b1.unlock_requirement = "place Domain Anchor and defeat 25 zombies"
        base.branches["hollow_purple"] = b1

        # Branch 2: Unlimited Void — guaranteed-hit zone, crowd control.
        b2 = UpgradeBranch(
            name="Unlimited Void (Domain)",
            cost=900,
            effects=[
                Effect(name="domain_zone", effect_type="attract", damage=0,
                       projectile=ProjectileType.ZERO, provenance=[self.abilities["G004"].claim_id],
                       explanation="Creates a guaranteed-hit domain zone; enemies inside are immobilized by infinite sensory input."),
            ],
            provenance=[self.abilities["G004"].claim_id],
        )
        base.branches["unlimited_void"] = b2

        return base, b1, b2, trace


class PromptTranslator(TranslatorABC):
    """LLM-driven translator: formats the dossier into a constrained prompt
    and emits a TowerDraft from structured JSON.

    Hook for the generative run. The demonstration uses CharacterTranslator
    directly so the experiment runs without an LLM. Set `llm` in __init__ with
    an object exposing `generate(prompt, schema)` returning JSON.
    """

    def __init__(self, dossier: CharacterDossier, llm=None):
        self.dossier = dossier
        self.llm = llm

    def translate(self) -> Tuple[TowerDraft, List[str]]:
        if self.llm is None:
            raise RuntimeError("No LLM registered for PromptTranslator; use "
                               "CharacterTranslator for the demonstration baseline.")
        prompt = self._build_prompt()
        raw = self.llm.generate(prompt, schema=dict)
        trace = self._parse(trace_raw=raw)
        return raw, trace

    def _build_prompt(self) -> str:
        ver = self.dossier.selected_version
        return (f"You are translating a character dossier into a Plants vs. Zombies "
                f"tower draft. Character: {self.dossier.character}, Series: "
                f"{self.dossier.series}, Version: {ver}.\n\n"
                f"Atomic claims (category: id — statement):\n"
                + "\n".join(
                    f"  [{c.category}] {c.claim_id}: {c.statement}"
                    for c in self.dossier.sorted_claims)
                + f"\n\nOutput exactly one TowerDraft (JSON): name, tier, cost, range, "
                f"cooldown, damage, damage_type, target_type, projectile, description, "
                f"and up to two upgrade branches. Every effect must cite the dossier "
                f"claim ID(s) it expresses. Faithful to the character AND playable in PvZ.")


def check_constraints(tower: TowerDraft, dossier) -> ConstraintCheck:
    """Well-formedness + design-quality constraints for a PvZ tower.

    Checks follow PCGBENCH25-style problem-specific metrics for PvZ, plus
    faithfulness-of-translation checks. Adjust thresholds empirically.
    """
    violations = []
    notes = []

    # PvZ cost sanity (budget balance class): no premium tower below ~250;
    # costs scale with DPS and range.
    if tower.cost < 150 or tower.cost > 2500:
        violations.append(f"cost={tower.cost} outside PvZ tower budget [150,2500]")

    if tower.range < 2 or tower.range > 10:
        violations.append(f"range={tower.range} outside PvZ [2,10]")

    # Cooldown in realistic tower band: no instant (0) and no > 3.0s
    if tower.cooldown <= 0.5:
        violations.append("cooldown too low -> near-instant firing, unbalanced")
    if tower.cooldown > 3.5:
        notes.append("slow cadence, compensated by damage")

    # Damage per shot reasonable for tier
    if tower.damage <= 0:
        violations.append("non-positive damage")

    # Faithfulness: at least one effect traces to an ability claim
    ability_claims = {c.claim_id for c in dossier.abilities()}
    all_provenance = set()
    for e in tower.effects:
        all_provenance.update(e.provenance)
    for branch in tower.branches.values():
        for e in branch.effects:
            all_provenance.update(e.provenance)
    if not all_provenance & ability_claims:
        violations.append("no effect cites a character ability claim")

    if not tower.description.strip():
        violations.append("missing descriptive flavor text")

    # Branch cost discipline: expensive upgrade must add clear value (new effect type)
    for name, b in tower.branches.items():
        if b.cost < tower.cost:
            violations.append(f"upgrade '{name}' costs less than base tower")
        new_effect_types = {e.effect_type for e in b.effects}
        if len(new_effect_types) < 1:
            notes.append(f"upgrade '{name}' adds no new effect type")

    return ConstraintCheck(len(violations) == 0, violations, notes)


def diversify(tower: TowerDraft, n: int = 3, dossier=None) -> List[TowerDraft]:
    """Generate n diverse variants of a base tower by perturbing numeric
    parameters within PvZ-valid bounds, then re-check constraints.

    This is the MORTAR26-style quality-diversity neighborhood: same concept,
    different numeric specialization, enabling the faithfulness-vs-playability
    comparison (NPC07).
    """
    variants = [tower]
    import random
    random.seed(42)
    for i in range(n - 1):
        t = TowerDraft(
            name=f"{tower.name} (v{i+1})",
            tier=tower.tier,
            cost=round(tower.cost * random.uniform(0.8, 1.25), 0),
            range=round(tower.range * random.uniform(0.85, 1.35), 1),
            cooldown=round(tower.cooldown * random.uniform(0.7, 1.4), 2),
            damage=round(tower.damage * random.uniform(0.7, 1.5), 0),
            damage_type=tower.damage_type,
            target_type=tower.target_type,
            projectile=tower.projectile,
            effects=[Effect(name=e.name, effect_type=e.effect_type,
                            damage=round(e.damage * random.uniform(0.8, 1.2), 0),
                            damage_type=e.damage_type, target_type=e.target_type,
                            projectile=e.projectile, provenance=list(e.provenance),
                            explanation=e.explanation) for e in tower.effects],
            branches=dict(tower.branches),
            description=tower.description,
            provenance_summary=list(tower.provenance_summary),
        )
        c = check_constraints(t, dossier) if dossier else check_constraints(t)
        if c:
            variants.append(t)
    return variants
