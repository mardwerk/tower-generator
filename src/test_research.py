"""Check reusable evidence, offline reuse, refresh failures and source boundaries."""

from copy import deepcopy
import json
from pathlib import Path
import tempfile
import unittest
from unittest.mock import patch

from src import research


class ResearchTests(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory(prefix="tower-research-test-")
        self.directory = Path(self.temporary.name)
        self.character = {"name": "Usopp", "work": "One Piece", "scope": "Manga through Dressrosa"}

    def tearDown(self):
        self.temporary.cleanup()

    def supplied(self):
        return research.collect_source(deepcopy(self.character), [], "Source passage about Kabuto.", directory=self.directory)

    def test_save_export_import_and_offline_reuse(self):
        source = self.supplied()
        exported = json.loads(json.dumps(source))
        with patch.object(research, "fetch_document", side_effect=AssertionError("Unexpected fetch")), \
             patch.object(research.socket, "getaddrinfo", side_effect=AssertionError("Unexpected DNS")):
            self.assertEqual(research.load_source(source["id"], self.directory), source)
            self.assertEqual(research.list_sources(self.directory), [source])
            self.assertEqual(research.save_source(exported, self.directory), source)

    def test_same_character_scope_reuses_record(self):
        source = self.supplied()
        again = research.collect_source(deepcopy(self.character), [], "Updated evidence.", directory=self.directory)
        self.assertEqual(again["id"], source["id"])
        self.assertEqual(len(research.list_sources(self.directory)), 1)
        other_scope = {**self.character, "scope": "Manga through Wano"}
        different = research.collect_source(other_scope, [], "Other period.", directory=self.directory)
        self.assertNotEqual(different["id"], source["id"])

    def test_failed_refresh_preserves_existing_evidence(self):
        source = self.supplied()
        before = research.source_path(source["id"], self.directory).read_bytes()
        with patch.object(research, "fetch_document", side_effect=ValueError("Page inaccessible")):
            with self.assertRaisesRegex(ValueError, "Page inaccessible"):
                research.collect_source(deepcopy(self.character), ["https://example.com/canon"], directory=self.directory)
        self.assertEqual(research.source_path(source["id"], self.directory).read_bytes(), before)

    def test_refresh_replaces_requested_page_and_retains_other_documents(self):
        source = self.supplied()
        document = {"id": "page", "title": "Canon page", "url": "https://example.com/canon", "access": "retrieved",
                    "retrievedAt": "2026-10-02T10:00:00+00:00", "text": "Old evidence"}
        source["documents"].append(document)
        research.save_source(source, self.directory)
        updated = {**document, "text": "New evidence"}
        with patch.object(research, "fetch_document", return_value=updated) as fetch:
            refreshed = research.collect_source(deepcopy(self.character), [document["url"]], directory=self.directory)
        fetch.assert_called_once_with(document["url"])
        self.assertEqual(len(refreshed["documents"]), 2)
        self.assertEqual(refreshed["documents"][0]["text"], "Source passage about Kabuto.")
        self.assertEqual(refreshed["documents"][1]["text"], "New evidence")

    def test_invalid_imports_cannot_overwrite_other_identity_or_escape_directory(self):
        source = self.supplied()
        other = deepcopy(source)
        other["character"]["name"] = "Luffy"
        with self.assertRaises(ValueError):
            research.save_source(other, self.directory)
        for identifier in ["../outside", "/tmp/outside", "usopp.json"]:
            with self.assertRaises(ValueError):
                research.save_source({**source, "id": identifier}, self.directory)
        with self.assertRaises(ValueError):
            research.validate_source({**source, "formatVersion": 2})
        with self.assertRaises(ValueError):
            research.validate_source({**source, "documents": [source["documents"][0]] * 2})

    def test_public_web_boundary_and_page_extraction(self):
        with patch.object(research.socket, "getaddrinfo", return_value=[(2, 1, 6, "", ("127.0.0.1", 80))]):
            with self.assertRaisesRegex(ValueError, "public web"):
                research.public_url("http://localhost/canon")
        for url in ["file:///etc/passwd", "https://user:secret@example.com", "http://example.com:5173"]:
            with self.assertRaises(ValueError):
                research.validate_url(url)
        parser = research.PageText()
        parser.feed('<title>Character</title><nav>Menu</nav><main><p>Kabuto &amp; Pop Greens</p>'
                    '<script>bad()</script><p>パチンコ</p></main>')
        self.assertEqual(parser.title, "Character")
        self.assertEqual(parser.result(), "Kabuto & Pop Greens\nパチンコ")


if __name__ == "__main__":
    unittest.main()
