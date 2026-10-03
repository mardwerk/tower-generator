"""Plants vs. Zombies-lite simulator and MCTS triage agents.

A minimal but playable PvZ-style environment: 5 lanes x 9 columns, wave-based
zombie spawning, targeting towers, and outcome measurement. The MCTS agent
evaluates a tower draft by searching placement and reporting win-rate,
survival, and placement robustness. This mirrors TDSTRATEGY24/TOWERMIND26's
MCTS-based PvZ playtesting and RUNTIME26's runtime-agent triage.

This is a triage approximation, not a fidelity claim: per GAVEL24 and
MORTAR26, agent scores are triage inputs only, adjudicated by humans.
"""

import math
import random
from dataclasses import dataclass, field
from typing import Dict, List, Optional, Tuple

from .spec import TowerDraft, Effect, DamageType, TargetType, ProjectileType


LANES = 5
COLS = 9
WAVE_COUNT = 30


@dataclass
class Zombie:
    zid: int
    lane: int
    x: float
    health: float
    max_health: float
    damage: float
    _speed: float
    armor: float = 0.0
    pushed: float = 0.0
    state: str = "normal"
    stunned_until: float = 0.0
    tag: str = ""

    @property
    def speed(self):
        return self._speed

    @speed.setter
    def speed(self, v):
        self._speed = v


@dataclass
class Tower:
    tower_id: int
    lane: int
    col: float
    draft: TowerDraft
    charge: float = 0.0


@dataclass
class Projectile:
    pid: int
    lane: int
    x: float
    damage: float
    projectile: ProjectileType
    range: float
    effect_type: str = "normal"
    damage_type: DamageType = DamageType.NORMAL


@dataclass
class State:
    zombies: List[Zombie] = field(default_factory=list)
    towers: List[Tower] = field(default_factory=list)
    projectiles: List[Projectile] = field(default_factory=list)
    wave: int = 0
    time: float = 0.0
    zombies_home: int = 0
    alive: bool = True
    seed: int = 0

    def reset(self, tower: TowerDraft, seed: int = 0):
        self.zombies = []
        self.towers = []
        self.projectiles = []
        self.wave = 0
        self.time = 0.0
        self.zombies_home = 0
        self.alive = True
        self.seed = seed

    def place_tower(self, tower: TowerDraft, lane: int, col: float) -> Tower:
        t = Tower(len(self.towers), lane, col, tower)
        self.towers.append(t)
        return t


class PvZWaveSystem:
    def __init__(self, difficulty: str = "normal"):
        self.difficulty = difficulty
        self.multipliers = {"easy": 0.8, "normal": 1.0, "hard": 1.35}

    def next_wave(self, state: State) -> List[Zombie]:
        m = self.multipliers[self.difficulty]
        wave = state.wave + 1
        # PvZ-style horde: grows with the wave and is spread across all lanes so that
        # every placed tower gets targets each wave (a fair test of a reusable unit).
        count = max(3, 2 + wave)
        zombies = []
        for i in range(count):
            # wave 6 is a boss wave with a tank; earlier waves mix variants.
            if wave == 6 and i == 0:
                variant = "tank"
            elif wave == 5 and i % 2 == 0:
                variant = "conehead"
            elif wave >= 4 and i % 4 == 0:
                variant = "fast"
            else:
                variant = "basic"
            lane = (wave - 1 + i) % LANES
            hp, dmg, spd, arm, label = self._variant_stats(variant, wave=wave, m=m)
            z = Zombie(len(state.zombies), lane, COLS, hp, hp, dmg, spd, arm, tag=label)
            zombies.append(z)
        state.wave += 1
        return zombies

    def _variant_stats(self, variant: str, wave: int, m: float) -> Tuple[float, float, float, float, str]:
        base = {
            "basic":  (200, 0.8, 1.0, 0.0, "Basic"),
            "conehead": (350, 0.8, 0.95, 100.0, "Conehead"),
            "fast":   (130, 0.8, 1.6, 0.0, "Speedster"),
            "tank":   (1200, 1.0, 0.6, 200.0, "Tank"),
        }
        hp, dmg, spd, arm, label = base[variant]
        return (hp * (1 + wave * 0.03) * m, dmg, spd, arm, label)


class WaveScheduler:
    def __init__(self, spawn_interval: float = 22.0, first_wave_at: float = 3.0):
        self.spawn_interval = spawn_interval
        self.first_wave_at = first_wave_at
        self.next_spawn = first_wave_at

    def update(self, state: State, dt: float) -> None:
        cond = self.next_spawn <= state.time and state.wave < WAVE_COUNT
        if cond:
            ws = PvZWaveSystem("normal")
            before = len(state.zombies)
            for z in ws.next_wave(state):
                z.x = COLS
                state.zombies.append(z)
            self.next_spawn += self.spawn_interval
        elif state.wave >= WAVE_COUNT and not state.zombies:
            state.alive = False  # victory (all waves survived/cleared)


@dataclass
class SimResult:
    outcome: str
    waves_cleared: int
    time_to_clear: float
    zombies_home: int
    total_damage_dealt: float
    damage_per_second: float
    max_pressure: float
    placement: Optional[Tuple[int, float]]
    notes: List[str] = field(default_factory=list)

    @property
    def survival_rate(self) -> float:
        return min(1.0, max(0.0, 1.0 - self.zombies_home / 6.0))


def run_simulation(state: State, tower: TowerDraft, placement: Tuple[int, float],
                   dt: float = 0.05, max_time: float = 600.0, n_towers: int = 1) -> SimResult:
    state.reset(tower, state.seed)
    # one tower per lane when evaluating a reusable unit
    for l in range(min(n_towers, LANES)):
        state.place_tower(tower, l, placement[1])
    scheduler = WaveScheduler(spawn_interval=22.0, first_wave_at=3.0)
    time = 0.0
    max_pressure = 0
    total_damage = 0.0
    projectile_id = 0

    while state.alive and time < max_time:
        scheduler.update(state, dt)
        time += dt
        state.time = time
        max_pressure = max(max_pressure, len(state.zombies))

        # --- Update zombies ---
        for z in state.zombies[:]:
            if z.state == "immobilized" and time >= z.stunned_until:
                z.state = "normal"
            if z.state == "immobilized":
                continue
            if z.pushed > 0.001:
                z.x += z.pushed * dt          # pushed right
                z.pushed -= z.speed * dt      # backlog shrinks as it advances
                if z.pushed < 0: z.pushed = 0
                continue
            z.x -= z.speed * dt             # normal leftward advance
            if z.x <= 0:
                z.x = 0
                state.zombies_home += 1
                state.alive = False
                state.zombies.remove(z)

        # --- Towers fire ---
        for tower_obj in state.towers:
            t = tower_obj
            lane = tower_obj.lane
            t.charge += dt
            # find closest target in lane within range (including the tower's own tile)
            target = None
            closest_x = 1e9
            for z in state.zombies:
                if z.lane == lane and t.col - 0.2 <= z.x <= t.col + t.draft.range + 0.2:
                    if z.x < closest_x:
                        closest_x = z.x
                        target = z
            if target is None:
                continue
            if t.charge >= t.draft.cooldown:
                t.charge = 0.0
                eff = next((e for e in t.draft.effects if e.effect_type in ("attract",)), None)
                if eff is None and t.draft.effects:
                    eff = t.draft.effects[0]
                base_dmg = eff.damage if eff else t.draft.damage
                proj = Projectile(
                    pid=projectile_id, lane=lane, x=tower_obj.col,
                    damage=base_dmg, projectile=t.draft.projectile,
                    range=t.draft.range, effect_type=eff.effect_type if eff else "normal",
                    damage_type=t.draft.damage_type,
                )
                projectile_id += 1

                state.projectiles.append(proj)
        # --- Projectiles move and hit ---
        for p in state.projectiles[:]:
            p.x += 8.0 * 0.05
            if p.x > COLS:
                state.projectiles.remove(p)
                continue
            hit = None
            for z in state.zombies:
                if z.lane == p.lane and abs(z.x - p.x) < 0.6:
                    hit = z
                    break
            if hit is None:
                continue
            # projectile consumed on first hit (normal projectiles pierce no zombies;
            # splash handling below applies its own area damage)
            state.projectiles.remove(p)
            dmg = p.damage
            if hit.armor > 0:
                dmg = max(0.0, p.damage - hit.armor)
            hit.health -= dmg
            total_damage += dmg
            if p.effect_type == "attract":
                for other in state.zombies:
                    if other.zid != hit.zid and other.lane == p.lane and abs(other.x - hit.x) <= 1.2:
                        other.x = max(0, hit.x + 1.3)
                        other.pushed = max(other.pushed, 0.8)
                hit.stunned_until = time + 0.8  # brief stun on the primary target
            elif p.effect_type == "impact":
                hit.pushed = max(hit.pushed, 1.2)
                hit._speed = max(0.0, hit.speed - 0.15)
            elif p.effect_type == "splash":
                for other in state.zombies:
                    if other.lane == p.lane and abs(other.x - p.x) <= 1.5 and other.zid != hit.zid:
                        other.health -= dmg * 0.5
                        total_damage += dmg * 0.5
            if hit.health <= 0:
                state.zombies.remove(hit)

    if state.zombies_home > 0:
        outcome = "defeat"
        waves_cleared = max(0, state.wave - 1)
        notes = [f"{state.zombies_home} zombies reached home"]
    else:
        outcome = "victory"
        waves_cleared = state.wave
        notes = ["all waves cleared"]

    return SimResult(
        outcome=outcome, waves_cleared=waves_cleared, time_to_clear=time,
        zombies_home=state.zombies_home, total_damage_dealt=total_damage,
        damage_per_second=total_damage / max(time, 1e-6), max_pressure=max_pressure,
        placement=placement, notes=notes,
    )
