"""Formal tower/unit specification for Plants vs. Zombies.

Mirrors a PvZ tower draft as a composable, constraint-checkable object.
This is the shared representation that Layer 2 outputs and Layer 3
consumes, keeping well-formedness separate from character semantics.
"""

from dataclasses import dataclass, field
from enum import Enum
from typing import Dict, List, Optional


class TargetType(Enum):
    GROUND = "ground"
    AIR = "air"
    ANY = "any"


class ProjectileType(Enum):
    NORMAL = "normal"
    PENETRATING = "penetrating"      # pierces through zombies
    SPLASH = "splash"                 # area damage on impact
    ZERO = "zero"                     # ignores durability / instant effect


class DamageType(Enum):
    NORMAL = "normal"
    CRUSH = "crush"
    MAGIC = "magic"


@dataclass
class Effect:
    """A single mechanical effect that a tower or upgrade applies.

    Each effect is annotated with which character-dossier claim it expresses,
    so the audit trace (Layer 4) can reconstruct provenance.
    """
    name: str
    effect_type: str                  # e.g. "single_target", "splash", "attract",
    damage: float                     # damage per hit
    damage_type: Optional[DamageType] = None
    target_type: TargetType = TargetType.GROUND
    projectile: Optional[ProjectileType] = None
    range: Optional[float] = None     # tiles, overrides tower range
    duration: Optional[float] = None  # seconds, if applicable
    cooldown_reduction: Optional[float] = None
    cost_delta: Optional[float] = None
    provenance: List[str] = field(default_factory=list)  # dossier claim IDs
    explanation: str = ""             # why this maps to the character ability

    def to_dict(self) -> dict:
        return {"name": self.name, "effect_type": self.effect_type, "damage": self.damage,
                "damage_type": self.damage_type.value if self.damage_type else None,
                "target_type": self.target_type.value,
                "projectile": self.projectile.value if self.projectile else None,
                "range": self.range, "duration": self.duration,
                "cooldown_reduction": self.cooldown_reduction, "cost_delta": self.cost_delta,
                "provenance": list(self.provenance), "explanation": self.explanation}


@dataclass
class UpgradeBranch:
    name: str
    cost: float
    effects: List[Effect] = field(default_factory=list)
    unlock_requirement: Optional[str] = None  # e.g. "defeat 50 zombies"
    provenance: List[str] = field(default_factory=list)

    def to_dict(self) -> dict:
        return {"name": self.name, "cost": self.cost,
                "effects": [e.to_dict() for e in self.effects],
                "unlock_requirement": self.unlock_requirement,
                "provenance": list(self.provenance)}


@dataclass
class TowerDraft:
    """A complete generated unit draft: base form plus upgrade branches."""
    name: str
    tier: str                           # e.g. "premium", "special", "standard"
    cost: float
    range: float
    cooldown: float                     # shots per second = 1/cooldown
    damage: float
    damage_type: DamageType
    target_type: TargetType
    projectile: ProjectileType
    effects: List[Effect] = field(default_factory=list)   # base effects
    branches: Dict[str, UpgradeBranch] = field(default_factory=dict)
    description: str = ""               # flavor text tied to character
    provenance_summary: List[str] = field(default_factory=list)  # top-level claim IDs

    @property
    def dps(self) -> float:
        return self.damage / max(self.cooldown, 1e-6)

    @property
    def cost_efficiency(self) -> float:
        return self.dps / max(self.cost, 1e-6)

    def to_dict(self) -> dict:
        d = {
            "name": self.name, "tier": self.tier, "cost": self.cost,
            "range": self.range, "cooldown": self.cooldown,
            "dps": self.dps, "cost_efficiency": self.cost_efficiency,
            "damage": self.damage, "damage_type": self.damage_type.value,
            "target_type": self.target_type.value,
            "projectile": self.projectile.value, "effects": [e.to_dict() for e in self.effects],
            "branches": {k: v.to_dict() for k, v in self.branches.items()},
            "description": self.description,
            "provenance_summary": self.provenance_summary,
        }
        return d

    def branch_to_dict(self) -> dict:
        return {k: {"name": v.name, "cost": v.cost, "effects": [e.to_dict() for e in v.effects],
                    "unlock_requirement": v.unlock_requirement, "provenance": list(v.provenance)}
                for k, v in self.branches.items()}

    def __repr__(self):
        return (f"<TowerDraft {self.name} cost={self.cost} range={self.range} "
                f"dps={self.dps:.1f} efficiency={self.cost_efficiency:.2f} "
                f"branches={list(self.branches)}>")

