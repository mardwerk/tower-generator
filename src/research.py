"""Portable character evidence, independent of Profiles and Tower designs."""

from datetime import datetime, timezone
import hashlib
from html.parser import HTMLParser
import ipaddress
import json
from pathlib import Path
import re
import socket
import unicodedata
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


# ponytail: add site-specific extraction only when a needed source loses evidence.
class PageText(HTMLParser):
    """Retain readable page text without scripts or navigation."""
    def __init__(self):
        super().__init__(convert_charrefs=True)
        self.parts = []
        self.skipped = 0
        self.title = ""
        self.in_title = False

    def handle_starttag(self, tag, attrs):
        if tag in {"script", "style", "nav", "header", "footer", "aside", "noscript"}:
            self.skipped += 1
        if tag == "title":
            self.in_title = True
        if tag in {"p", "div", "section", "li", "br", "h1", "h2", "h3", "h4"}:
            self.parts.append("\n")

    def handle_endtag(self, tag):
        if tag in {"script", "style", "nav", "header", "footer", "aside", "noscript"}:
            self.skipped = max(0, self.skipped - 1)
        if tag == "title":
            self.in_title = False
        if tag in {"p", "div", "section", "li", "h1", "h2", "h3", "h4"}:
            self.parts.append("\n")

    def handle_data(self, data):
        if self.in_title:
            self.title += data
        elif not self.skipped:
            self.parts.append(data)

    def result(self):
        lines = [re.sub(r"\s+", " ", line).strip() for line in "".join(self.parts).splitlines()]
        return "\n".join(line for line in lines if line)


def fetch_document(url):
    public_url(url)
    opener = build_opener(ProxyHandler({}), PublicRedirect())
    request = Request(url, headers={"User-Agent": "Mozilla/5.0 TowerGenerator/0.1"})
    with opener.open(request, timeout=15) as response:
        content_type = response.headers.get_content_type()
        if content_type not in {"text/html", "text/plain"}:
            raise ValueError("Research accepts HTML pages and plain text")
        raw = response.read(2_000_001)
        if len(raw) > 2_000_000:
            raise ValueError("Source page is larger than 2 MB; supply a smaller excerpt")
        body = raw.decode(response.headers.get_content_charset() or "utf-8", errors="replace")
        final_url = response.url
    parser = PageText()
    if content_type == "text/html":
        parser.feed(body)
        body = parser.result()
    if len(body) > MAX_DOCUMENT:
        raise ValueError("Source text is too long; supply a focused excerpt")
    return {
        "id": hashlib.sha256(url.encode()).hexdigest()[:12],
        "title": parser.title.strip()[:300] or urlsplit(url).hostname,
        "url": url,
        "resolvedUrl": final_url,
        "retrievedAt": datetime.now(timezone.utc).isoformat(),
        "access": "retrieved",
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


def list_sources(directory=SOURCES):
    return [validate_source(json.loads(path.read_text())) for path in sorted(directory.glob("*.json"))]


def identity(character):
    return tuple(unicodedata.normalize("NFKC", character[key]).strip().casefold() for key in ["name", "work", "scope"])


def save_source(source, directory=SOURCES):
    validate_source(source)
    path = source_path(source["id"], directory)
    if path.exists() and identity(load_source(source["id"], directory)["character"]) != identity(source["character"]):
        raise ValueError("This source identifier already belongs to another character or scope")
    directory.mkdir(parents=True, exist_ok=True)
    temporary = path.with_suffix(".tmp")
    temporary.write_text(json.dumps(source, indent=2, ensure_ascii=False) + "\n")
    temporary.replace(path)
    return source


def collect_source(character, urls, supplied="", notes="", directory=SOURCES):
    for key, limit in [("name", 120), ("work", 120), ("scope", 500)]:
        character[key] = text(character.get(key), f"Character {key}", limit)
    if not isinstance(urls, list) or len(urls) > 5:
        raise ValueError("Collect at most five URLs at a time")
    urls = list(dict.fromkeys(url.strip() for url in urls if isinstance(url, str) and url.strip()))
    supplied = text(supplied, "Supplied evidence", MAX_DOCUMENT, required=False)
    if not urls and not supplied:
        raise ValueError("Provide source URLs or paste evidence to research this character")
    existing = next((s for s in list_sources(directory) if identity(s["character"]) == identity(character)), None)
    identifier = existing["id"] if existing else "character-" + hashlib.sha256(json.dumps(identity(character)).encode()).hexdigest()[:16]
    documents = {document["id"]: document for document in existing["documents"]} if existing else {}
    for url in urls:
        document = fetch_document(url)
        # Replace a refreshed page even when an imported record gave it another id.
        documents = {key: value for key, value in documents.items() if value.get("url") != url}
        documents[document["id"]] = document
    if supplied:
        documents["supplied"] = {"id": "supplied", "title": "Supplied evidence", "url": None,
                                 "retrievedAt": None, "access": "supplied", "text": supplied}
    source = {"formatVersion": 1, "kind": "character-source", "id": identifier,
              "character": character, "documents": list(documents.values()), "notes": notes}
    return save_source(source, directory)
