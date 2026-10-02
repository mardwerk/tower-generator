"""Build a small, editable projectile Tower from the verified Dart reference."""

from copy import deepcopy
import hashlib
import json
import math
from pathlib import Path
import re
import shutil
import subprocess
import sys
import tempfile


PROJECT = Path(__file__).resolve().parents[1]
CAPTURE = "data/56.3-build-24829026"
SOURCE = "de829684232157967fd66f0e999a45df3a669c63"
STATS = {"cost", "damage", "pierce", "range", "interval", "projectiles", "camo"}
CHANGES = {"damage", "pierce", "range", "intervalMultiplier", "projectiles", "camo", "blastRadius"}
ATTACKS = {"explosion", "spread", "sniper"}


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
    if {p["upgrades"][2].get("attack") for p in design["paths"]} != ATTACKS:
        raise ValueError("The third purchases must establish explosion, spread and sniper attacks")
    for key, value in design["base"].items():
        validate_value(key, value)
    for path in design["paths"]:
        for tier, upgrade in enumerate(path["upgrades"], 1):
            if "attack" in upgrade and (tier != 3 or upgrade["attack"] not in ATTACKS):
                raise ValueError("Attack replacements belong on the third purchase and must be explosion, spread or sniper")
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
    stats = dict(design["base"], attack="pellet", blastRadius=0)
    for path, count in zip(design["paths"], tiers):
        for upgrade in path["upgrades"][:count]:
            stats["cost"] += upgrade["cost"]
            if "attack" in upgrade:
                stats["attack"] = upgrade["attack"]
            for key, value in upgrade["changes"].items():
                if key == "intervalMultiplier":
                    stats["interval"] *= value
                elif key == "camo":
                    stats["camo"] = stats["camo"] or value
                else:
                    stats[key] += value
    return stats


def apply_stats(tower, stats, arc, projectiles):
    tower["cost"], tower["range"] = stats["cost"], stats["range"]
    attack = next(models(tower, "AttackModel"))
    attack["range"] = stats["range"]
    weapon = attack["weapons"][0]
    weapon["projectile"] = deepcopy(projectiles[stats["attack"]])
    weapon["rate"] = weapon["Rate"] = stats["interval"]
    if stats["projectiles"] > 1:
        weapon["emission"] = deepcopy(arc)
        weapon["emission"]["count"] = weapon["emission"]["Count"] = stats["projectiles"]
    projectile = weapon["projectile"]
    hit = projectile
    if stats["attack"] == "explosion":
        # The carrier expires at first contact; purchased pierce belongs to its blast.
        hit = next(models(projectile, "CreateProjectileOnContactModel"))["projectile"]
        hit["radius"] = stats["blastRadius"]
    hit["pierce"] = hit["CappedPierce"] = stats["pierce"]
    damage = next(models(hit, "DamageModel"))
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
    character_source = None
    if design.get("sourceId"):
        from src.research import load_source
        character_source = load_source(design["sourceId"])
        if character_source["character"]["name"].casefold() != design["name"].casefold():
            raise ValueError("The selected character source does not match the Tower name")
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
    projectiles = {"pellet": next(models(root, "WeaponModel"))["projectile"],
                   "spread": next(models(root, "WeaponModel"))["projectile"],
                   "sniper": next(models(next(s for s in states if s["tiers"] == [0, 0, 3]), "WeaponModel"))["projectile"],
                   "explosion": next(models(read("Towers/BombShooter/BombShooter.json"), "WeaponModel"))["projectile"]}
    upgrade_template = read("Upgrades/Sharp Shots.json")
    text = read("textTable.json")
    resources = read("resources.json")
    data = output / "game-data"
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
        apply_stats(tower, stats, arc, projectiles)
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
    source_record = {"repository": "https://github.com/KyleDerZweite/btd6-atlas", "commit": SOURCE,
                     "capture": manifest["gameVersion"], "build": manifest["steamBuildId"], "files": provenance}
    if character_source:
        write_json(output / "character-source.json", character_source)
        source_record["characterSource"] = {
            "id": character_source["id"],
            "sha256": hashlib.sha256((output / "character-source.json").read_bytes()).hexdigest(),
            "file": "character-source.json",
        }
    else:
        (output / "character-source.json").unlink(missing_ok=True)
    write_json(output / "source.json", source_record)
    description = [f"# {design['name']}", "", design["description"], "",
                   "Draft projectile design. 64 ordinary states, 15 upgrades, no Paragon or Monkey Knowledge.",
                   "Costs and stats are authored choices. Inherited BTD6 visuals and sounds are placeholders.",
                   "Gunpowder stars use Bomb Shooter contact explosions. Leaf shuriken use Dart spread; Kabuto shots use Dart Crossbow projectiles.",
                   "Inherited BTD6 immunity rules remain. Pop Greens are adapted as projectiles, without persistent plants or active abilities.", "",
                   "Changes accumulate in top-middle-bottom order. Interval multipliers multiply; camo detection stays enabled once gained.",
                   "Cost is the base purchase plus every applied upgrade. The interval is seconds per volley.",
                   "Damage and pierce are per projectile, or per explosion for Gunpowder Star. Its carrier expires on first contact.",
                   "Crosspath pierce increases the explosion capacity or the leaf/sniper projectile capacity; reload and range apply to the selected attack.",
                   "Actual Usopp art is pending. The third purchases specify gunpowder shots, Black Kabuto with leaf shuriken, and Sogeking with Kabuto.", "",
                   "| Upgrade | Path | Name | Cost | Effect |", "| --- | --- | --- | --- | --- |", *upgrade_rows, "",
                   "| State | Cost | Damage | Pierce | Range | Interval | Projectiles | Camo | Attack | Blast radius |",
                   "| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |"]
    for tiers, stats in sorted(rows):
        description.append(f"| {'-'.join(map(str, tiers))} | {stats['cost']:g} | {stats['damage']:g} | {stats['pierce']:g} | "
                           f"{stats['range']:g} | {stats['interval']:.6g} | {stats['projectiles']} | {stats['camo']} | {stats['attack']} | {stats['blastRadius']:g} |")
    (output / "TOWER.md").write_text("\n".join(description) + "\n")
    (output / "REFERENCE.md").write_text("\n".join([
        "# Dart reference", "", f"Atlas commit {SOURCE}, capture {manifest['gameVersion']}, Steam build {manifest['steamBuildId']}.",
        "Full source files and byte hashes are recorded in source.json. The original source remains in the Atlas checkout.",
        "This view shows the first weapon only. Higher states can include additional attacks and abilities.",
        "Atlas data attribution: KyleDerZweite/btd6-atlas, CC BY-NC 4.0. See the repository LICENSE.", "",
        "| State | Cost | Range | Interval | Pierce | Damage |", "| --- | --- | --- | --- | --- | --- |", *reference_rows,
    ]) + "\n")
    return data


def check(output, tower_id):
    binary = output / ("profile-validator.exe" if sys.platform == "win32" else "profile-validator")
    identity = json.loads((output / "validator-release.json").read_text())
    if hashlib.sha256(binary.read_bytes()).hexdigest() != identity["executableSha256"]:
        raise ValueError("Installed validator checksum mismatch; run setup.py again")
    data = output / "game-data"
    tower = f"Towers/{tower_id}/{tower_id}.json"

    def report(directory):
        result = subprocess.run([str(binary), "score-tower", "--profile", str(output / "profile"),
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
