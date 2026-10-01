"""Build a small, editable projectile Tower from the verified Dart reference."""

import argparse
from copy import deepcopy
import hashlib
import json
import math
from pathlib import Path
import re
import runpy
import shutil
import subprocess
import sys
import tempfile


PROJECT = Path(__file__).resolve().parents[1]
CAPTURE = "data/56.3-build-24829026"
SOURCE = "de829684232157967fd66f0e999a45df3a669c63"
STATS = {"cost", "damage", "pierce", "range", "interval", "projectiles", "camo"}
CHANGES = {"damage", "pierce", "range", "intervalMultiplier", "projectiles", "camo"}


def models(value, kind):
    if isinstance(value, dict):
        if value.get("$type", "").split(",")[0].split(".")[-1] == kind:
            yield value
        for child in value.values():
            yield from models(child, kind)
    elif isinstance(value, list):
        for child in value:
            yield from models(child, kind)


def validate_design(design):
    if not re.fullmatch(r"[A-Za-z][A-Za-z0-9]*", design["id"]):
        raise ValueError("Tower id must contain only ASCII letters and digits")
    if set(design["base"]) != STATS:
        raise ValueError(f"Base stats must be {sorted(STATS)}")
    if len(design["paths"]) != 3 or any(len(p["upgrades"]) != 5 for p in design["paths"]):
        raise ValueError("This experiment requires three paths with five upgrades each")
    for key, value in design["base"].items():
        validate_value(key, value)
    for path in design["paths"]:
        for upgrade in path["upgrades"]:
            if not upgrade["name"].strip() or not upgrade["description"].strip():
                raise ValueError("Every upgrade needs a name and description")
            validate_value("cost", upgrade["cost"])
            if not upgrade["changes"] or not set(upgrade["changes"]) <= CHANGES:
                raise ValueError(f"Upgrade changes must use {sorted(CHANGES)}")
            for key, value in upgrade["changes"].items():
                validate_value(key, value)


def validate_value(key, value):
    if key == "camo":
        if not isinstance(value, bool):
            raise ValueError("camo must be true or false")
    elif (isinstance(value, bool) or not isinstance(value, (int, float))
          or not math.isfinite(value) or value < 0):
        raise ValueError(f"{key} must be a finite, nonnegative number")
    elif key in {"interval", "intervalMultiplier", "projectiles"} and value <= 0:
        raise ValueError(f"{key} must be positive")
    elif key == "projectiles" and not isinstance(value, int):
        raise ValueError("projectiles must be an integer")


def state_name(tower_id, tiers):
    return tower_id if not any(tiers) else tower_id + "-" + "".join(map(str, tiers))


def upgrade_id(tower_id, path, tier):
    code = [0, 0, 0]
    code[path] = tier
    return tower_id + "Upgrade" + "".join(map(str, code))


def stats_for(design, tiers):
    stats = dict(design["base"])
    for path, count in zip(design["paths"], tiers):
        for upgrade in path["upgrades"][:count]:
            stats["cost"] += upgrade["cost"]
            for key, value in upgrade["changes"].items():
                if key == "intervalMultiplier":
                    stats["interval"] *= value
                elif key == "camo":
                    stats["camo"] = stats["camo"] or value
                else:
                    stats[key] += value
    return stats


def apply_stats(tower, stats, arc):
    tower["cost"], tower["range"] = stats["cost"], stats["range"]
    attack = next(models(tower, "AttackModel"))
    attack["range"] = stats["range"]
    weapon = attack["weapons"][0]
    weapon["rate"] = weapon["Rate"] = stats["interval"]
    if stats["projectiles"] > 1:
        weapon["emission"] = deepcopy(arc)
        weapon["emission"]["count"] = weapon["emission"]["Count"] = stats["projectiles"]
    projectile = weapon["projectile"]
    projectile["pierce"] = projectile["CappedPierce"] = stats["pierce"]
    damage = next(models(projectile, "DamageModel"))
    damage["damage"] = damage["CappedDamage"] = stats["damage"]
    # Keep the demonstrated straight projectile alive beyond the selected range.
    travel = next(models(projectile, "TravelStraitModel"))
    travel["lifespan"] = travel["Lifespan"] = max(travel["lifespan"], stats["range"] / travel["speed"])
    for invisible in models(tower, "FilterInvisibleModel"):
        invisible["isActive"] = not stats["camo"]


def write_json(path, value):
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(value, indent=2, allow_nan=False) + "\n")


def generate(design, atlas, output):
    validate_design(design)
    capture = atlas / CAPTURE
    manifest = json.loads((capture / "manifest.json").read_text())
    provenance = {}

    def read(relative):
        raw = (capture / "game-data" / relative).read_bytes()
        expected = subprocess.run(
            ["git", "show", f"{SOURCE}:{CAPTURE}/game-data/{relative}"],
            cwd=atlas, capture_output=True, check=True,
        ).stdout
        if raw != expected:
            raise ValueError(f"Source differs from pinned Atlas commit: {relative}")
        digest = hashlib.sha256(raw).hexdigest()
        if relative in manifest.get("topLevelSha256", {}) and digest != manifest["topLevelSha256"][relative]:
            raise ValueError(f"Capture manifest checksum mismatch: {relative}")
        provenance[relative] = digest
        return json.loads(raw)

    states = [read(str(path.relative_to(capture / "game-data")))
              for path in sorted((capture / "game-data/Towers/DartMonkey").glob("*.json"))
              if path.name != "DartMonkey-Paragon.json"]
    states = [state for state in states if not state["isParagon"] and not state["isSubTower"]]
    if len(states) != 64 or len({tuple(s["tiers"]) for s in states}) != 64:
        raise ValueError("Expected 64 distinct ordinary Dart states")
    root = next(s for s in states if s["IsBaseTower"])
    arc = next(models(next(s for s in states if s["tiers"] == [0, 3, 0]), "ArcEmissionModel"))
    upgrade_template = read("Upgrades/Sharp Shots.json")
    text = read("textTable.json")
    resources = read("resources.json")
    data = output / "data"
    (output / "checks.json").unlink(missing_ok=True)
    tower_id = design["id"]
    text[tower_id] = design["name"]
    text[tower_id + " Description"] = design["description"]
    rows = []
    reference_rows = []
    for source in states:
        tiers = source["tiers"]
        tower = deepcopy(root)
        tower["name"] = state_name(tower_id, tiers)
        tower["baseId"] = tower["BaseId"] = tower_id
        tower["tiers"], tower["tier"] = tiers, max(tiers)
        tower["IsBaseTower"] = not any(tiers)
        tower["BeastHandlerLeashMutationId"] = "BeastHandlerContribution" + tower_id
        tower["beastHandlerLeashMutationId"] = None
        tower["mods"] = []
        tower["paragonUpgrade"] = None
        tower["upgrades"] = []
        for edge in source["upgrades"]:
            target = next((s for s in states if s["name"] == edge["tower"]), None)
            if target is None:
                continue
            path = next(i for i in range(3) if target["tiers"][i] != tiers[i])
            purchase = dict(edge)
            purchase["tower"] = state_name(tower_id, target["tiers"])
            purchase["upgrade"] = upgrade_id(tower_id, path, target["tiers"][path])
            tower["upgrades"].append(purchase)
        tower["appliedUpgrades"] = [upgrade_id(tower_id, path, tier)
                                     for path, count in enumerate(tiers) for tier in range(1, count + 1)]
        stats = stats_for(design, tiers)
        apply_stats(tower, stats, arc)
        write_json(data / "Towers" / tower_id / (tower["name"] + ".json"), tower)
        rows.append((tiers, stats))
        weapon = next(models(source, "WeaponModel"))
        projectile = weapon["projectile"]
        reference_rows.append(f"| {'-'.join(map(str, tiers))} | {source['cost']:g} | {source['range']:g} | "
                              f"{weapon['rate']:g} | {projectile['pierce']:g} | "
                              f"{next(models(projectile, 'DamageModel'))['damage']:g} |")
    upgrade_rows = []
    for path, definition in enumerate(design["paths"]):
        for tier, upgrade in enumerate(definition["upgrades"], 1):
            identifier = upgrade_id(tower_id, path, tier)
            record = deepcopy(upgrade_template)
            record.update(name=identifier, LocsKey=identifier, localizedNameOverride="", cost=upgrade["cost"], path=path, tier=tier - 1)
            write_json(data / "Upgrades" / (identifier + ".json"), record)
            text[identifier], text[identifier + " Description"] = upgrade["name"], upgrade["description"]
            code = ["x", "x", "x"]
            code[path] = str(tier)
            upgrade_rows.append(f"| {'-'.join(code)} | {definition['name']} | {upgrade['name']} | {upgrade['cost']} | {upgrade['description']} |")
    write_json(data / "textTable.json", text)
    write_json(data / "resources.json", resources)
    write_json(output / "source.json", {"repository": "https://github.com/KyleDerZweite/btd6-atlas", "commit": SOURCE,
               "capture": manifest["gameVersion"], "build": manifest["steamBuildId"], "files": provenance})
    description = [f"# {design['name']}", "", design["description"], "",
                   "Draft projectile design. 64 ordinary states, 15 upgrades, no Paragon or Monkey Knowledge.",
                   "Costs and stats are authored choices. Inherited Dart visuals and sounds are placeholders.",
                   "Damage uses Dart's existing immunity rules. No active abilities, plant effects, or game integration are included.", "",
                   "Changes accumulate in top-middle-bottom order. Interval multipliers multiply; camo detection stays enabled once gained.",
                   "Cost is the base purchase plus every applied upgrade. The interval is seconds per volley; damage and pierce are per projectile.", "",
                   "| Upgrade | Path | Name | Cost | Effect |", "| --- | --- | --- | --- | --- |", *upgrade_rows, "",
                   "| State | Cost | Damage | Pierce | Range | Interval | Projectiles | Camo |",
                   "| --- | --- | --- | --- | --- | --- | --- | --- |"]
    for tiers, stats in sorted(rows):
        description.append(f"| {'-'.join(map(str, tiers))} | {stats['cost']:g} | {stats['damage']:g} | {stats['pierce']:g} | "
                           f"{stats['range']:g} | {stats['interval']:.6g} | {stats['projectiles']} | {stats['camo']} |")
    (output / "TOWER.md").write_text("\n".join(description) + "\n")
    (output / "REFERENCE.md").write_text("\n".join([
        "# Dart reference", "", f"Atlas commit {SOURCE}, capture {manifest['gameVersion']}, Steam build {manifest['steamBuildId']}.",
        "Full source files and byte hashes are recorded in source.json. The original source remains in the Atlas checkout.",
        "This view shows the first weapon only. Higher states can include additional attacks and abilities.",
        "Atlas data attribution: KyleDerZweite/btd6-atlas, CC BY-NC 4.0. See the repository LICENSE.", "",
        "| State | Cost | Range | Interval | Pierce | Damage |", "| --- | --- | --- | --- | --- | --- |", *reference_rows,
    ]) + "\n")
    return data


def check(atlas, output, tower_id):
    smoke = runpy.run_path(str(PROJECT / "scripts/check-default-profile.py"))
    binary, schemas = smoke["checker"](argparse.Namespace(local_checker=False, atlas=atlas))
    smoke["check_schemas"](schemas, [atlas / "profile", PROJECT / "default-profile"])
    data = output / "data"
    tower = f"Towers/{tower_id}/{tower_id}.json"

    def report(directory):
        result = subprocess.run([str(binary), "score-tower", "--profile", str(PROJECT / "default-profile"),
                                 "--game-data", str(directory), "--tower", tower], capture_output=True, text=True)
        if not result.stdout:
            raise ValueError(result.stderr.strip() or "Checker produced no report")
        return result.returncode, json.loads(result.stdout)

    code, valid = report(data)
    write_json(output / "checks.json", {"candidate": valid})
    if code or not valid["score"]["complete"] or valid["score"]["points"] != 100:
        for error in valid["errors"]:
            print(error, file=sys.stderr)
        raise ValueError("Candidate did not pass complete validation")
    results = {"candidate": valid}
    for label, missing in [("missingState", f"Towers/{tower_id}/{tower_id}-100.json"),
                            ("missingUpgrade", f"Upgrades/{upgrade_id(tower_id, 0, 1)}.json")]:
        with tempfile.TemporaryDirectory(prefix="tower-authoring-") as temporary:
            copy = Path(temporary) / "data"
            shutil.copytree(data, copy)
            (copy / missing).unlink()
            code, invalid = report(copy)
        if code == 0 or not any(error["code"] == "missing_reference" and Path(missing).stem in error["message"]
                                for error in invalid["errors"]):
            raise ValueError(f"Checker did not diagnose {label}")
        results[label] = invalid
    write_json(output / "checks.json", results)
    print(f"Candidate: {valid['score']['points']}/100, {valid['filesChecked']} files. Missing state and upgrade rejected.")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("design", type=Path)
    parser.add_argument("--atlas", type=Path, default=PROJECT.parent / "btd6-atlas")
    parser.add_argument("--check", action="store_true", help="validate the candidate and missing-state/upgrade cases")
    args = parser.parse_args()
    design = json.loads(args.design.read_text())
    validate_design(design)
    output = PROJECT / "game-data" / design["id"].lower()
    generate(design, args.atlas.resolve(), output)
    print(f"Authored {design['name']}: {output / 'TOWER.md'}")
    if args.check:
        check(args.atlas.resolve(), output, design["id"])


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, KeyError, StopIteration, subprocess.CalledProcessError) as error:
        sys.exit(f"author-tower: {error}")
