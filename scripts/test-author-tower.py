"""Check authored crosspaths and purchase mappings independently of the checker."""

from copy import deepcopy
import json
from pathlib import Path
import runpy
import tempfile
import unittest


PROJECT = Path(__file__).resolve().parents[1]
AUTHOR = runpy.run_path(str(PROJECT / "scripts/author-tower.py"))
DESIGN = json.loads((PROJECT / "examples/usopp.json").read_text())


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
        self.assertEqual(tower["range"], 76)
        self.assertTrue(all(not model["isActive"] for model in AUTHOR["models"](tower, "FilterInvisibleModel")))
        travel = next(AUTHOR["models"](tower, "TravelStraitModel"))
        self.assertGreaterEqual(travel["speed"] * travel["lifespan"], tower["range"])
        weapon = next(AUTHOR["models"](tower, "WeaponModel"))
        damage = next(AUTHOR["models"](weapon, "DamageModel"))
        self.assertEqual(damage["damage"], 6)
        self.assertEqual(damage["CappedDamage"], damage["damage"])

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
        tower = json.loads((revised / "data/Towers/Usopp/Usopp.json").read_text())
        self.assertEqual(next(AUTHOR["models"](tower, "DamageModel"))["damage"], 2)
        self.assertIn("| 0-0-0 | 250 | 2 |", (revised / "TOWER.md").read_text())

    def test_invalid_design_fails_before_writing(self):
        for changes in [{"damge": 1}, {"damage": float("nan")}, {"intervalMultiplier": 0}]:
            design = deepcopy(DESIGN)
            design["paths"][0]["upgrades"][0]["changes"] = changes
            with self.assertRaises(ValueError):
                AUTHOR["validate_design"](design)


if __name__ == "__main__":
    unittest.main()
