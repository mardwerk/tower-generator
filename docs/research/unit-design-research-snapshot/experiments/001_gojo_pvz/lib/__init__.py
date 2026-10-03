"""Package init: expose the unit-design pipeline components."""

from .spec import TowerDraft, UpgradeBranch, Effect, DamageType, TargetType, ProjectileType
from .retrieval import (CharacterDossier, AtomicClaim, VersionProfile,
                        RetrieverABC, BundledProfileRetriever, build_gojo_dossier,
                        build_dossier)
from .generator import (ConstraintCheck, TranslatorABC, PromptTranslator,
                        check_constraints, diversify, CharacterTranslator)
from .simulator import (PvZWaveSystem, WaveScheduler, State, Projectile,
                        run_simulation, SimResult)
from .triege import (MCTSNode, uct_score, create_child, rollout, mcts_evaluate,
                     AgentTriager)
from .explanation import TraceElement, AuditTrace, build_audit_trace

__all__ = [
    "TowerDraft", "UpgradeBranch", "Effect", "DamageType", "TargetType", "ProjectileType",
    "CharacterDossier", "AtomicClaim", "VersionProfile",
    "RetrieverABC", "BundledProfileRetriever", "build_gojo_dossier", "build_dossier",
    "check_constraints", "diversify", "TranslatorABC", "PromptTranslator",
    "CharacterTranslator", "PvZWaveSystem", "WaveScheduler", "State", "run_simulation", "SimResult",
    "AgentTriager", "mcts_evaluate", "build_audit_trace", "AuditTrace", "TraceElement",
]
