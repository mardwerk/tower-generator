"""Check authored crosspaths and purchase mappings independently of the checker."""

from copy import deepcopy
import hashlib
import json
from pathlib import Path
import runpy
import tempfile
import unittest


PROJECT = Path(__file__).resolve().parents[1]
AUTHOR = runpy.run_path(str(PROJECT / "src/author.py"))
DESIGN = json.loads((PROJECT / "default/usopp.json").read_text())


class AuthoringTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.temporary = tempfile.TemporaryDirectory(prefix="tower-authoring-test-")
        cls.output = Path(cls.temporary.name)
        cls.data = AUTHOR["generate"](DESIGN, PROJECT.parent / "btd6-atlas", cls.output)

    @classmethod
    def tearDownClass(cls):
        cls.temporary.cleanup()

    def tower(self, name):
        return json.loads((self.data / "Towers/Usopp" / (name + ".json")).read_text())

    def test_crosspath_preserves_both_paths(self):
        tower = self.tower("Usopp-230")
        weapon = next(AUTHOR["models"](tower, "WeaponModel"))
        self.assertEqual(tower["cost"], 1930)
        self.assertEqual(tower["tiers"], [2, 3, 0])
        self.assertEqual(weapon["projectile"]["pierce"], 5)
        self.assertEqual(weapon["emission"]["count"], 3)
        self.assertAlmostEqual(weapon["rate"], 0.68)
        self.assertEqual(len(tower["appliedUpgrades"]), 5)

    def test_sniper_camo_and_range_reach(self):
        tower = self.tower("Usopp-025")
        self.assertEqual(tower["range"], 86)
        self.assertTrue(all(not model["isActive"] for model in AUTHOR["models"](tower, "FilterInvisibleModel")))
        travel = next(AUTHOR["models"](tower, "TravelStraitModel"))
        self.assertGreaterEqual(travel["speed"] * travel["lifespan"], tower["range"])
        weapon = next(AUTHOR["models"](tower, "WeaponModel"))
        damage = next(AUTHOR["models"](weapon, "DamageModel"))
        self.assertEqual(damage["damage"], 7)
        self.assertEqual(damage["CappedDamage"], damage["damage"])

    def test_third_purchases_replace_the_attack(self):
        base = next(AUTHOR["models"](self.tower("Usopp"), "WeaponModel"))
        blast = next(AUTHOR["models"](self.tower("Usopp-300"), "WeaponModel"))
        contact = next(AUTHOR["models"](blast, "CreateProjectileOnContactModel"))
        self.assertEqual(blast["projectile"]["pierce"], 1)
        self.assertEqual(contact["projectile"]["radius"], 12)
        self.assertEqual(contact["projectile"]["pierce"], 14)
        self.assertFalse(list(AUTHOR["models"](self.tower("Usopp-200"), "CreateProjectileOnContactModel")))
        spread = next(AUTHOR["models"](self.tower("Usopp-030"), "WeaponModel"))
        self.assertEqual(spread["emission"]["count"], 3)
        sniper = next(AUTHOR["models"](self.tower("Usopp-003"), "WeaponModel"))
        self.assertNotEqual(sniper["projectile"]["display"], base["projectile"]["display"])
        self.assertEqual(next(AUTHOR["models"](sniper, "TravelStraitModel"))["speed"], 400)
        self.assertEqual(self.tower("Usopp-003")["range"], 66)
        self.assertEqual(sniper["projectile"]["pierce"], 6)

    def test_explosion_crosspaths_and_capstone(self):
        for tier, capacity, radius, damage in [(3, 14, 12, 1), (4, 20, 16, 2), (5, 32, 22, 5)]:
            for secondary in [0, 1, 2]:
                tower = self.tower(f"Usopp-{tier}{secondary}0")
                weapon = next(AUTHOR["models"](tower, "WeaponModel"))
                blast = next(AUTHOR["models"](weapon, "CreateProjectileOnContactModel"))["projectile"]
                self.assertEqual(blast["pierce"], capacity)
                self.assertEqual(blast["radius"], radius)
                self.assertEqual(next(AUTHOR["models"](blast, "DamageModel"))["damage"], damage)
                self.assertAlmostEqual(weapon["rate"], [1, .85, .68][secondary])
                ranged = self.tower(f"Usopp-{tier}0{secondary}")
                ranged_blast = next(AUTHOR["models"](ranged, "CreateProjectileOnContactModel"))["projectile"]
                self.assertEqual(ranged_blast["pierce"], capacity + (1 if secondary == 2 else 0))
                self.assertEqual(ranged_blast["radius"], radius)
                self.assertEqual(ranged["range"], [36, 44, 48][secondary])

    def test_complete_legal_states_and_matching_purchases(self):
        states = [self.tower(path.stem) for path in (self.data / "Towers/Usopp").glob("*.json")]
        self.assertEqual(len(states), 64)
        upgrades = {path.stem: json.loads(path.read_text()) for path in (self.data / "Upgrades").glob("*.json")}
        self.assertEqual(len(upgrades), 15)
        by_name = {tower["name"]: tower for tower in states}
        for tower in states:
            tiers = tower["tiers"]
            self.assertFalse(tower["isParagon"])
            self.assertIsNone(tower["paragonUpgrade"])
            self.assertLessEqual(sum(t > 0 for t in tiers), 2)
            self.assertLessEqual(sorted(tiers)[1], 2)
            self.assertEqual(tower["IsBaseTower"], not any(tiers))
            applied = [identifier for identifier, upgrade in upgrades.items() if upgrade["tier"] < tiers[upgrade["path"]]]
            self.assertEqual(set(tower["appliedUpgrades"]), set(applied))
            for edge in tower["upgrades"]:
                target = by_name[edge["tower"]]
                upgrade = upgrades[edge["upgrade"]]
                difference = [b - a for a, b in zip(tiers, target["tiers"])]
                self.assertEqual(sum(difference), 1)
                self.assertEqual(difference[upgrade["path"]], 1)
                self.assertEqual(upgrade["tier"] + 1, target["tiers"][upgrade["path"]])

    def test_revision_changes_data_and_readable_output(self):
        revision = deepcopy(DESIGN)
        revision["base"]["damage"] = 2
        revised = self.output / "revision"
        AUTHOR["generate"](revision, PROJECT.parent / "btd6-atlas", revised)
        tower = json.loads((revised / "game-data/Towers/Usopp/Usopp.json").read_text())
        self.assertEqual(next(AUTHOR["models"](tower, "DamageModel"))["damage"], 2)
        self.assertIn("| 0-0-0 | 250 | 2 |", (revised / "TOWER.md").read_text())

    def test_character_source_is_pinned_in_generated_provenance(self):
        snapshot = self.output / "character-source.json"
        provenance = json.loads((self.output / "source.json").read_text())
        self.assertEqual(provenance["characterSource"]["id"], "usopp")
        self.assertEqual(provenance["characterSource"]["sha256"], hashlib.sha256(snapshot.read_bytes()).hexdigest())
        self.assertEqual(json.loads(snapshot.read_text())["character"]["work"], "One Piece")

    def test_invalid_design_fails_before_writing(self):
        for changes in [{"damge": 1}, {"damage": float("nan")}, {"intervalMultiplier": 0}]:
            design = deepcopy(DESIGN)
            design["paths"][0]["upgrades"][0]["changes"] = changes
            with self.assertRaises(ValueError):
                AUTHOR["validate_design"](design)


if __name__ == "__main__":
    unittest.main()
