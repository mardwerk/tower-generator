"""Verify reusable Markdown storage and independent CLI operations."""

from copy import deepcopy
import json
from pathlib import Path
import subprocess
import tempfile
import unittest
from unittest.mock import patch

from src import wiki


class WikiTests(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory(prefix="character-wiki-test-")
        self.root = Path(self.temporary.name) / "wiki"
        self.character = {"name": "Usopp", "work": "One Piece", "scope": "Manga through Dressrosa"}
        self.entry = wiki.collect(deepcopy(self.character), [], "Kabuto launches physical ammunition.\nObservation Haki senses presences.", root=self.root)["entry"]

    def tearDown(self):
        self.temporary.cleanup()

    def add_ability(self, entry=None):
        entry = entry or self.entry
        meta, body = wiki.parse_markdown(entry["markdown"])
        meta["abilities"] = [{"id": "kabuto", "kind": "equipment", "name": "Kabuto", "scope": meta["scope"],
                              "description": "A slingshot that launches physical ammunition.",
                              "evidence": [{"source": "sources/supplied.md", "passage": entry["documents"][0]["passages"][0]["id"]}],
                              "classification": {"delivery": "projectile", "functions": ["damage"], "status": "suggested"}}]
        return wiki.save_entry(entry["key"], wiki.markdown(meta, body), entry["revision"], self.root)

    def test_markdown_round_trip_and_offline_lookup(self):
        entry = self.add_ability()
        self.assertTrue((self.root / "one-piece/usopp/README.md").exists())
        self.assertTrue((self.root / "one-piece/usopp/sources/supplied.md").exists())
        with patch.object(wiki, "fetch_document", side_effect=AssertionError("Unexpected fetch")):
            self.assertEqual(wiki.read_entry(entry["key"], self.root)["metadata"], entry["metadata"])
            result = wiki.list_entries("One Piece", self.root)
        self.assertEqual(result["entries"][0]["key"], "one-piece/usopp")

    def test_existing_pages_are_reused_and_batch_failure_preserves_files(self):
        document = {"id": "official", "title": "Character", "url": "https://example.com/usopp", "access": "retrieved",
                    "retrievedAt": "2026-10-02T10:00:00+00:00", "text": "Usopp uses a slingshot."}
        with patch.object(wiki, "fetch_document", return_value=document):
            entry = wiki.collect(deepcopy(self.character), [document["url"]], root=self.root)["entry"]
        with patch.object(wiki, "fetch_document", side_effect=AssertionError("Unexpected fetch")):
            reused = wiki.collect(deepcopy(self.character), [document["url"], document["url"]], root=self.root)
        self.assertEqual(reused["reused"], 1)
        before = {p.relative_to(self.root): p.read_bytes() for p in self.root.rglob("*.md")}
        with patch.object(wiki, "fetch_document", side_effect=[{**document, "text": "Changed evidence."}, ValueError("Page inaccessible")]):
            with self.assertRaisesRegex(ValueError, "Page inaccessible"):
                wiki.collect(deepcopy(self.character), [document["url"], "https://example.com/other"], refresh=True, root=self.root)
        self.assertEqual(before, {p.relative_to(self.root): p.read_bytes() for p in self.root.rglob("*.md")})
        self.assertEqual(wiki.read_entry(entry["key"], self.root)["metadata"]["abilities"], [])

    def test_passage_ids_survive_unrelated_additions(self):
        before = wiki.parse_source(wiki.source_document({"id": "sample", "title": "Example", "access": "supplied", "text": "First fact.\nSecond fact."}), "sample")
        after = wiki.parse_source(wiki.source_document({"id": "sample", "title": "Example", "access": "supplied", "text": "New opening.\nFirst fact.\nSecond fact."}), "sample")
        self.assertEqual(before["passages"], after["passages"][1:])

    def test_refresh_retains_cited_prior_passages_and_resets_review(self):
        document = {"id": "official", "title": "Character", "url": "https://example.com/usopp", "access": "retrieved",
                    "retrievedAt": "2026-10-02T10:00:00+00:00", "text": "Kabuto launches ammunition."}
        with patch.object(wiki, "fetch_document", return_value=document):
            entry = wiki.collect(deepcopy(self.character), [document["url"]], root=self.root)["entry"]
        meta, body = wiki.parse_markdown(entry["markdown"])
        passage = next(d for d in entry["documents"] if d["id"] == "official")["passages"][0]
        meta["abilities"] = [{"id": "kabuto", "kind": "equipment", "name": "Kabuto", "scope": meta["scope"],
                              "description": "Launches ammunition.", "evidence": [{"source": "sources/official.md", "passage": passage["id"]}]}]
        entry = wiki.save_entry(entry["key"], wiki.markdown(meta, body), entry["revision"], self.root)
        wiki.review_entry(entry["key"], entry["revision"], self.root)
        with patch.object(wiki, "fetch_document", return_value={**document, "retrievedAt": "2026-10-03T10:00:00+00:00", "text": "The page removed the technique details."}):
            refreshed = wiki.collect(deepcopy(self.character), [document["url"]], refresh=True, root=self.root)["entry"]
        source = next(d for d in refreshed["documents"] if d["id"] == "official")
        self.assertIn(passage, source["passages"])
        self.assertEqual(source["retainedPassages"][0]["retrievedAt"], document["retrievedAt"])
        self.assertEqual(refreshed["metadata"]["status"], "draft")

    def test_invalid_references_categories_and_scope_do_not_overwrite(self):
        entry = self.add_ability()
        before = (self.root / entry["key"] / "README.md").read_bytes()
        for mutation in ["reference", "category", "scope"]:
            meta, body = wiki.parse_markdown(entry["markdown"])
            if mutation == "reference": meta["abilities"][0]["evidence"][0]["passage"] = "missing"
            if mutation == "category": meta["abilities"][0]["classification"]["delivery"] = "telepathy-damage"
            if mutation == "scope": meta["scope"] = "Another period"
            with self.assertRaises(ValueError):
                wiki.save_entry(entry["key"], wiki.markdown(meta, body), entry["revision"], self.root)
        self.assertEqual(before, (self.root / entry["key"] / "README.md").read_bytes())
        with self.assertRaisesRegex(ValueError, "canon scope"):
            wiki.collect({**self.character, "scope": "Anime only"}, [], "Other evidence", root=self.root)

    def test_stale_edits_and_direct_file_changes_require_new_review(self):
        entry = self.add_ability()
        reviewed = wiki.review_entry(entry["key"], entry["revision"], self.root)
        self.assertEqual(reviewed["metadata"]["status"], "reviewed")
        self.assertEqual(reviewed["metadata"]["abilities"][0]["classification"]["status"], "suggested")
        with self.assertRaisesRegex(ValueError, "changed"):
            wiki.save_entry(entry["key"], entry["markdown"], entry["revision"], self.root)
        path = self.root / entry["key"] / "README.md"
        path.write_text(path.read_text() + "\nAn additional unresolved question.\n")
        changed = wiki.read_entry(entry["key"], self.root)
        self.assertTrue(changed["reviewStale"])
        self.assertEqual(changed["metadata"]["status"], "draft")
        self.assertEqual(wiki.list_entries(root=self.root)["entries"][0]["status"], "draft")

    def test_path_boundaries_and_invalid_entry_do_not_hide_valid_entries(self):
        for key in ["../outside", "/tmp/x", "one-piece/../../outside", "one-piece/Usopp"]:
            with self.assertRaises(ValueError): wiki.read_entry(key, self.root)
        outside = Path(self.temporary.name) / "outside"
        outside.mkdir()
        (self.root / "one-piece/linked").symlink_to(outside, target_is_directory=True)
        with self.assertRaisesRegex(ValueError, "symbolic"):
            wiki.entry_path("one-piece/linked", self.root)
        bad = self.root / "one-piece/bad/README.md"
        bad.parent.mkdir()
        bad.write_text("No frontmatter")
        result = wiki.list_entries(root=self.root)
        self.assertEqual(len(result["entries"]), 1)
        self.assertEqual(len(result["warnings"]), 1)

    def test_cli_migration_and_lookup_are_repeatable_and_preserve_originals(self):
        def command(*args, target=None):
            result = subprocess.run(["python3", "-B", "-m", "src.wiki", "--wiki", str(target or self.root), "--json", *args], cwd=wiki.PROJECT, capture_output=True, text=True)
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            return json.loads(result.stdout)
        empty = Path(self.temporary.name) / "imported"
        empty.mkdir()
        original = (wiki.PROJECT / "sources/usopp.json").read_bytes()
        (empty / "usopp.json").write_bytes(original)
        imported_root = Path(self.temporary.name) / "migrated-wiki"
        result = command("migrate", "--sources", str(empty), target=imported_root)
        self.assertEqual(result["imported"], ["one-piece/usopp"])
        self.assertEqual(len(command("show", "one-piece/usopp", target=imported_root)["documents"]), 1)
        result = command("migrate", "--sources", str(empty), target=imported_root)
        self.assertEqual(result["skipped"], ["one-piece/usopp"])
        self.assertEqual((empty / "usopp.json").read_bytes(), original)
        self.assertEqual(command("show", "one-piece/usopp")["key"], self.entry["key"])
        self.assertEqual(command("list", "--query", "Usopp")["entries"][0]["name"], "Usopp")


if __name__ == "__main__":
    unittest.main()
