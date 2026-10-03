"""MCTS triage agents (Layer 3).

The MCTS agent evaluates a tower draft by searching placement (lane + column)
and estimating win-rate and robustness under PvZ-like waves. Follows the
precedent of TDSTRATEGY24/TOWERMIND26 (MCTS-based PvZ playtesting) and
RUNTIME26 (autonomous agent triage). Agent scores are triage inputs, not
final adjudication — per GAVEL24 and MORTAR26 humans judge the shortlist.
"""

import math
import random
from dataclasses import dataclass, field
from typing import Dict, List, Optional, Tuple

from .simulator import State, Tower, Projectile, run_simulation, LANES, COLS, WAVE_COUNT, PvZWaveSystem
from .spec import TowerDraft


@dataclass
class MCTSNode:
    state: State
    placement: Optional[Tuple[int, float]]
    draft: TowerDraft
    parent: Optional["MCTSNode"] = None
    children: List["MCTSNode"] = field(default_factory=list)
    visits: int = 0
    wins: float = 0.0


def uct_score(node: MCTSNode) -> float:
    if node.visits == 0:
        return float("inf")
    return (node.wins / max(node.visits, 1)) + math.sqrt(2 * math.log(max(node.parent.visits, 1)) / node.visits)


def create_child(parent: MCTSNode, placement: Tuple[int, float]) -> MCTSNode:
    st = State(seed=parent.state.seed + len(parent.children))
    st.reset(parent.draft, st.seed)
    # one tower per lane (same tower repeated), placed at its best column
    lane, col = placement
    for l in range(LANES):
        st.place_tower(parent.draft, l, col)
    return MCTSNode(state=st, placement=placement, draft=parent.draft, parent=parent)


def rollout(node: MCTSNode, max_time: float = 300.0) -> bool:
    """Run the tower alone against waves; return True if all waves are cleared."""
    placement = node.placement or (0, 5.0)
    res = run_simulation(node.state, node.draft, placement, dt=0.05, max_time=max_time, n_towers=LANES)
    return res.outcome == "victory"


def mcts_evaluate(tower: TowerDraft, simulations: int = 300) -> dict:
    """Search placement across lanes and columns; return aggregate triage scores."""
    root = MCTSNode(state=State(seed=1), placement=None, draft=tower)
    placements = [(l, round(c, 1)) for l in range(LANES) for c in range(2, COLS)]

    for _ in range(simulations):
        node = root
        depth = 0
        while node.children:
            node = max(node.children, key=uct_score)
            depth += 1
            if depth > 3:
                break
        if not node.children:
            p = placements[node.state.seed % len(placements)]
            node.children.append(create_child(node, p))
            node = node.children[-1]
        node.visits += 1
        if rollout(node):
            node.wins += 1
        cur = node.parent
        while cur:
            cur.visits += 1
            cur.wins += node.wins / node.visits
            cur = cur.parent

    stats_by_placement = {}
    def gather(node, acc):
        if node.placement:
            acc.setdefault(node.placement, {"visits": 0, "wins": 0.0})
            acc[node.placement]["visits"] += node.visits
            acc[node.placement]["wins"] += node.wins
        for c in node.children:
            gather(c, acc)
    gather(root, stats_by_placement)

    results = sorted(stats_by_placement.items(), key=lambda kv: kv[1]["wins"] / max(kv[1]["visits"], 1), reverse=True)
    best_place, best_stats = results[0] if results else (None, {"visits": 0, "wins": 0.0})
    win_rate = best_stats["wins"] / max(best_stats["visits"], 1)

    def place_all_lanes(state: State, tower: TowerDraft, col: float) -> None:
        for l in range(LANES):
            state.place_tower(tower, l, col)

    def run_all_lanes(tower: TowerDraft, col: float, seed: int, max_time: float = 120.0) -> SimResult:
        st = State(seed=seed)
        st.reset(tower, st.seed)
        place_all_lanes(st, tower, col)
        return run_simulation(st, tower, (0, col), dt=0.05, max_time=max_time, n_towers=LANES)

    n_detail = min(5, max(3, int(simulations * 0.5)))
    best_col = (best_place or (0, 5.0))[1]
    detailed = [run_all_lanes(tower, best_col, 1000 + i) for i in range(n_detail)]

    avg_dps = sum(r.damage_per_second for r in detailed) / max(len(detailed), 1)
    avg_pressure = sum(r.max_pressure for r in detailed) / max(len(detailed), 1)
    avg_surv = sum(r.survival_rate for r in detailed) / max(len(detailed), 1)

    return {
        "win_rate": round(win_rate, 3),
        "visits": sum(v["visits"] for v in stats_by_placement.values()),
        "top_placement": best_place,
        "avg_dps": round(avg_dps, 1),
        "avg_max_pressure": round(avg_pressure, 1),
        "avg_survival_rate": round(avg_surv, 3),
        "detailed_outcomes": [dict(outcome=r.outcome, waves=r.waves_cleared, dps=r.damage_per_second)
                              for r in detailed],
        "sample_notes": detailed[0].notes if detailed else [],
    }


class AgentTriager:
    """Triage API: runs MCTS evaluation and returns a pass/fail decision."""

    def __init__(self, simulations: int = 300, difficulty: str = "normal"):
        self.simulations = simulations
        self.difficulty = difficulty
        # PvZ-style triage thresholds; calibrate empirically (PCGBENCH25 approach)
        self.thresholds = {"win_rate": 0.5, "survival_rate": 0.6}

    def triage(self, tower: TowerDraft) -> dict:
        """Return a triage report: pass/fail plus diagnostics."""
        report = mcts_evaluate(tower, self.simulations)
        report["difficulty"] = self.difficulty
        report["triage"] = {
            "win_rate": report["win_rate"],
            "survival_rate": report["avg_survival_rate"],
            "pressure_per_wave": report["avg_max_pressure"],
            "passes_win_rate": report["win_rate"] >= self.thresholds["win_rate"],
            "passes_survival": report["avg_survival_rate"] >= self.thresholds["survival_rate"],
        }
        return report
