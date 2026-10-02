"""Portable character evidence, independent of Profiles and Tower designs."""

from datetime import datetime, timezone
import hashlib
from html.parser import HTMLParser
import ipaddress
import json
from pathlib import Path
import re
import socket
from urllib.parse import urlsplit
from urllib.request import HTTPRedirectHandler, ProxyHandler, Request, build_opener

PROJECT = Path(__file__).resolve().parents[1]
SOURCES = PROJECT / "sources"
MAX_SOURCE = 600_000
MAX_DOCUMENT = 80_000


def text(value, field, limit, required=True):
    if not isinstance(value, str) or len(value) > limit or (required and not value.strip()):
        raise ValueError(f"{field} must be {'nonempty ' if required else ''}text up to {limit} characters")
    return value.strip()


def source_path(identifier, directory=SOURCES):
    if not isinstance(identifier, str) or not re.fullmatch(r"[a-z0-9][a-z0-9-]{0,100}", identifier):
        raise ValueError("Invalid source identifier")
    return directory / f"{identifier}.json"


def validate_url(url):
    text(url, "Source URL", 2000)
    parsed = urlsplit(url)
    if parsed.scheme not in {"http", "https"} or not parsed.hostname or parsed.username or parsed.password:
        raise ValueError("Source URLs must be public HTTP or HTTPS links without credentials")
    if parsed.port not in {None, 80, 443}:
        raise ValueError("Source URLs must use standard HTTP or HTTPS ports")
    return parsed


def public_url(url):
    parsed = validate_url(url)
    addresses = socket.getaddrinfo(parsed.hostname, parsed.port or (443 if parsed.scheme == "https" else 80))
    if not addresses or any(not ipaddress.ip_address(item[4][0]).is_global for item in addresses):
        raise ValueError("Research only fetches public web pages")


class PublicRedirect(HTTPRedirectHandler):
    def redirect_request(self, request, fp, code, message, headers, newurl):
        public_url(newurl)
        return super().redirect_request(request, fp, code, message, headers, newurl)


class PageTitle(HTMLParser):
    def __init__(self):
        super().__init__(convert_charrefs=True)
        self.title = ""
        self.in_title = False
        self.language = ""

    def handle_starttag(self, tag, attrs):
        if tag == "title": self.in_title = True
        if tag == "html": self.language = (dict(attrs).get("lang") or "").split("-")[0].lower()

    def handle_endtag(self, tag):
        if tag == "title": self.in_title = False

    def handle_data(self, data):
        if self.in_title: self.title += data


def fetch_document(url):
    public_url(url)
    opener = build_opener(ProxyHandler({}), PublicRedirect())
    request = Request(url, headers={"User-Agent": "Mozilla/5.0 TowerGenerator/0.1", "Accept-Language": "en"})
    with opener.open(request, timeout=15) as response:
        content_type = response.headers.get_content_type()
        if content_type not in {"text/html", "text/plain"}:
            raise ValueError("Research accepts HTML pages and plain text")
        raw = response.read(2_000_001)
        if len(raw) > 2_000_000:
            raise ValueError("Source page is larger than 2 MB; supply a smaller excerpt")
        body = raw.decode(response.headers.get_content_charset() or "utf-8", errors="replace")
        final_url = response.url
    parser = PageTitle()
    if content_type == "text/html":
        parser.feed(body)
        if parser.language and parser.language != "en":
            raise ValueError("Use an English source page for character research")
        from trafilatura import extract
        body = extract(body, include_comments=False, include_tables=True)
        if not body:
            raise ValueError("No readable source text found; supply a focused excerpt")
    if len(body) > MAX_DOCUMENT:
        raise ValueError("Source text is too long; supply a focused excerpt")
    return {
        "id": re.sub(r"[^a-z0-9]+", "-", urlsplit(url).hostname.lower()).strip("-")[:70] + "-" + hashlib.sha256(url.encode()).hexdigest()[:12],
        "title": parser.title.strip()[:300] or urlsplit(url).hostname,
        "url": url,
        "resolvedUrl": final_url,
        "retrievedAt": datetime.now(timezone.utc).isoformat(),
        "access": "retrieved",
        "language": "en",
        "text": text(body, "Source text", MAX_DOCUMENT),
    }


def validate_source(source):
    if not isinstance(source, dict) or type(source.get("formatVersion")) is not int or source.get("formatVersion") != 1 or source.get("kind") != "character-source":
        raise ValueError("Expected character-source JSON with formatVersion 1")
    source_path(source.get("id"))
    character = source.get("character", {})
    for key, limit in [("name", 120), ("work", 120), ("scope", 500)]:
        text(character.get(key), f"Character {key}", limit)
    documents = source.get("documents")
    if not isinstance(documents, list) or not 1 <= len(documents) <= 8:
        raise ValueError("A character source needs 1 to 8 evidence documents")
    identifiers = set()
    for document in documents:
        if not isinstance(document, dict):
            raise ValueError("Evidence documents must be objects")
        identifier = text(document.get("id"), "Document id", 100)
        if identifier in identifiers:
            raise ValueError("Document ids must be unique")
        identifiers.add(identifier)
        text(document.get("title"), "Document title", 300)
        text(document.get("text"), "Document text", MAX_DOCUMENT)
        if document.get("access") not in {"supplied", "retrieved"}:
            raise ValueError("Document access must be supplied or retrieved")
        for key in ["url", "resolvedUrl"]:
            if document.get(key) is not None:
                validate_url(document[key])
        date = document.get("retrievedAt")
        if document["access"] == "retrieved" and (not date or not document.get("url")):
            raise ValueError("Retrieved documents need a URL and retrieval date")
        if date is not None:
            parsed = datetime.fromisoformat(date.replace("Z", "+00:00"))
            if parsed.tzinfo is None:
                raise ValueError("Retrieval dates need a timezone")
    source["notes"] = text(source.get("notes", ""), "Research notes", 8000, required=False)
    if len(json.dumps(source, ensure_ascii=False).encode()) > MAX_SOURCE:
        raise ValueError("Character source is too large")
    return source


def load_source(identifier, directory=SOURCES):
    return validate_source(json.loads(source_path(identifier, directory).read_text()))
