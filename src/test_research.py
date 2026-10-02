"""Check public fetching and the retained historical source reader."""

import unittest
from unittest.mock import MagicMock, patch

from src import research


class SourceTests(unittest.TestCase):
    def test_public_web_boundary(self):
        with patch.object(research.socket, "getaddrinfo", return_value=[(2, 1, 6, "", ("127.0.0.1", 80))]):
            with self.assertRaisesRegex(ValueError, "public web"):
                research.public_url("http://localhost/canon")
        for url in ["file:///etc/passwd", "https://user:secret@example.com", "http://example.com:5173"]:
            with self.assertRaises(ValueError):
                research.validate_url(url)

    def test_collection_requests_english_and_rejects_explicit_japanese_pages(self):
        response = MagicMock()
        response.headers.get_content_type.return_value = "text/html"
        response.headers.get_content_charset.return_value = "utf-8"
        response.read.return_value = '<html lang="ja"><title>Japanese page</title><p>日本語の情報</p></html>'.encode()
        response.url = "https://example.com/character"
        with patch.object(research, "public_url"), patch.object(research, "build_opener") as factory:
            factory.return_value.open.return_value.__enter__.return_value = response
            with self.assertRaisesRegex(ValueError, "English source"):
                research.fetch_document(response.url)
            request = factory.return_value.open.call_args.args[0]
            self.assertEqual(request.get_header("Accept-language"), "en")

    def test_legacy_source_remains_readable_for_authoring_and_migration(self):
        source = research.load_source("usopp")
        self.assertEqual(source["character"]["name"], "Usopp")
        self.assertEqual(len(source["documents"]), 1)
        with self.assertRaises(ValueError):
            research.load_source("../outside")


if __name__ == "__main__":
    unittest.main()
