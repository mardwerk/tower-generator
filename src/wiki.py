"""Local Markdown character Wiki. The website invokes this same CLI."""

import argparse
from copy import deepcopy
from concurrent.futures import ThreadPoolExecutor
from datetime import datetime, timezone
import hashlib
import json
from pathlib import Path
import re
import sys
import tempfile
import unicodedata

import yaml

from src.research import MAX_DOCUMENT, MAX_SOURCE, PROJECT, fetch_document, text, validate_source, validate_url

WIKI = PROJECT / "wiki"
CATEGORIES = PROJECT / "docs/research-categories.yaml"


def slug(value):
    value = unicodedata.normalize("NFKC", value).casefold()
    value = re.sub(r"[^\w]+|_+", "-", value, flags=re.UNICODE).strip("-")
    if not value or len(value) > 120:
        raise ValueError("Use a name that produces a readable folder up to 120 characters")
    return value


def entry_key(character):
    return f"{slug(character['work'])}/{slug(character['name'])}"


def safe_path(root, relative):
    root = Path(root).absolute()
    parts = relative.split("/")
    if any(not part or part in {".", ".."} or "\\" in part for part in parts):
        raise ValueError("Invalid Wiki path")
    path = root
    for part in parts:
        path = path / part
        if path.is_symlink():
            raise ValueError("Wiki files and folders cannot be symbolic links")
    if not path.resolve().is_relative_to(root.resolve()):
        raise ValueError("Wiki path leaves its directory")
    return path


def entry_path(key, root=WIKI):
    if not isinstance(key, str) or len(key.split("/")) != 2 or any(slug(p) != p for p in key.split("/")):
        raise ValueError("Use a Wiki key such as one-piece/usopp")
    return safe_path(root, key)


def markdown(metadata, body):
    return "---\n" + yaml.safe_dump(metadata, allow_unicode=True, sort_keys=False).rstrip() + "\n---\n\n" + body.strip() + "\n"


def parse_markdown(value):
    value = text(value, "Markdown", MAX_SOURCE).replace("\r\n", "\n")
    if len(value.encode()) > MAX_SOURCE or not value.startswith("---\n") or "\n---\n" not in value[4:]:
        raise ValueError("Markdown needs YAML frontmatter between --- lines")
    header, body = value[4:].split("\n---\n", 1)
    metadata = yaml.safe_load(header)
    if not isinstance(metadata, dict):
        raise ValueError("YAML frontmatter must be an object")
    return metadata, body.strip()


def read_text(path):
    if path.stat().st_size > MAX_SOURCE:
        raise ValueError("Wiki file is too large")
    return path.read_text(encoding="utf-8")


def source_document(document):
    paragraphs = [p.strip() for p in re.split(r"\n+", document["text"]) if p.strip()]
    # A content-based passage id survives unrelated additions and refreshes.
    passages = []
    seen = set()
    for paragraph in paragraphs:
        identifier = "p-" + hashlib.sha256(paragraph.encode()).hexdigest()[:16]
        if identifier not in seen:
            passages.append({"id": identifier, "text": paragraph})
            seen.add(identifier)
    metadata = {key: value for key, value in document.items() if key != "text"}
    body = f"# {document['title']}\n\n" + "\n\n".join(f"<!-- passage:{p['id']} -->\n{p['text']}" for p in passages)
    return markdown(metadata, body)


def parse_source(value, identifier):
    meta, body = parse_markdown(value)
    if meta.get("id") != identifier or not re.fullmatch(r"[a-z0-9][a-z0-9-]{0,100}", identifier):
        raise ValueError("Invalid source id")
    text(meta.get("title"), "Source title", 300)
    if meta.get("language", "en") != "en":
        raise ValueError("Character research uses English source documents")
    if meta.get("access") not in {"retrieved", "supplied"}:
        raise ValueError("Source access must be retrieved or supplied")
    for key in ["url", "resolvedUrl"]:
        if meta.get(key) is not None:
            validate_url(meta[key])
    date = meta.get("retrievedAt")
    if meta["access"] == "retrieved" and (not meta.get("url") or not date):
        raise ValueError("Retrieved sources need a URL and date")
    if date is not None and (not isinstance(date, str) or datetime.fromisoformat(date.replace("Z", "+00:00")).tzinfo is None):
        raise ValueError("Source dates need a timezone-aware text value")
    parts = re.split(r"<!-- passage:([a-z0-9-]+) -->\n", body)
    passages = [{"id": parts[i], "text": parts[i + 1].strip()} for i in range(1, len(parts), 2)]
    if not passages or len({p['id'] for p in passages}) != len(passages):
        raise ValueError("Sources need unique passage markers")
    for passage in passages:
        text(passage["text"], "Source passage", MAX_DOCUMENT)
    if sum(len(p["text"]) for p in passages) > MAX_DOCUMENT:
        raise ValueError("Source text is too long")
    return {**meta, "passages": passages}


def categories():
    return yaml.safe_load(CATEGORIES.read_text(encoding="utf-8"))


def validate_entry(meta, body, documents, key):
    if type(meta.get("formatVersion")) is not int or meta.get("formatVersion") != 1 or meta.get("kind") != "character-wiki":
        raise ValueError("Expected character-wiki with formatVersion 1")
    for name, limit in [("name", 120), ("work", 120), ("scope", 500)]:
        text(meta.get(name), name, limit)
    if entry_key(meta) != key:
        raise ValueError("Character name and work must match the Wiki folder")
    if meta.get("language", "en") != "en":
        raise ValueError("Character research entries use English")
    aliases = meta.get("aliases", [])
    if not isinstance(aliases, list) or len(aliases) > 20:
        raise ValueError("Aliases must be a list up to 20 names")
    for alias in aliases:
        text(alias, "Alias", 120)
    if meta.get("status") not in {"draft", "reviewed"}:
        raise ValueError("Entry status must be draft or reviewed")
    source_ids = meta.get("sources")
    if not isinstance(source_ids, list) or not 1 <= len(source_ids) <= 10 or len(set(source_ids)) != len(source_ids):
        raise ValueError("An entry needs 1 to 10 unique sources")
    if set(source_ids) != set(documents):
        raise ValueError("The entry's source files are missing")
    abilities = meta.get("abilities", [])
    if not isinstance(abilities, list) or len(abilities) > 100:
        raise ValueError("Abilities must be a list up to 100 records")
    allowed = categories()
    ids = set()
    for ability in abilities:
        if not isinstance(ability, dict):
            raise ValueError("Ability records must be objects")
        identifier = text(ability.get("id"), "Ability id", 120)
        if slug(identifier) != identifier or identifier in ids:
            raise ValueError("Ability ids must be unique slugs")
        ids.add(identifier)
        if ability.get("kind") not in {"ability", "trait", "equipment"}:
            raise ValueError("Record kind must be ability, trait or equipment")
        for field, limit in [("name", 150), ("description", 3000), ("scope", 500)]:
            text(ability.get(field), f"Ability {field}", limit)
        text(ability.get("limitations", ""), "Limitations", 3000, required=False)
        evidence = ability.get("evidence")
        if not isinstance(evidence, list) or not evidence:
            raise ValueError(f"{identifier} needs cited evidence")
        for reference in evidence:
            if not isinstance(reference, dict):
                raise ValueError("Evidence references must be objects")
            file = reference.get("source", "")
            if not isinstance(file, str) or not re.fullmatch(r"sources/[a-z0-9][a-z0-9-]{0,100}\.md", file):
                raise ValueError("Evidence must reference sources/<id>.md")
            document = documents.get(file[8:-3])
            if not document or not any(p["id"] == reference.get("passage") for p in document["passages"]):
                raise ValueError(f"{identifier} references a missing source passage")
        classification = ability.get("classification", {})
        if classification.get("delivery", "unknown") not in allowed["delivery"]:
            raise ValueError("Unknown delivery category")
        functions = classification.get("functions", ["unknown"])
        if not isinstance(functions, list) or not functions or any(f not in allowed["functions"] for f in functions):
            raise ValueError("Unknown function category")
        if classification.get("status", "suggested") not in {"suggested", "reviewed"}:
            raise ValueError("Classification status must be suggested or reviewed")
    text(body, "Entry body", 80_000)
    return meta


def fingerprint(meta, body, files):
    metadata = {k: v for k, v in meta.items() if k not in {"status", "reviewedAt", "reviewedHash"}}
    return hashlib.sha256(json.dumps([metadata, body, files], sort_keys=True, ensure_ascii=False).encode()).hexdigest()


def read_entry(key, root=WIKI):
    path = entry_path(key, root)
    raw = read_text(safe_path(root, f"{key}/README.md"))
    meta, body = parse_markdown(raw)
    source_ids = meta.get("sources", [])
    if not isinstance(source_ids, list) or len(source_ids) > 10:
        raise ValueError("Invalid source list")
    files = {}
    documents = {}
    for identifier in source_ids:
        if not isinstance(identifier, str) or not re.fullmatch(r"[a-z0-9][a-z0-9-]{0,100}", identifier):
            raise ValueError("Invalid source id")
        files[identifier] = read_text(safe_path(root, f"{key}/sources/{identifier}.md"))
        documents[identifier] = parse_source(files[identifier], identifier)
    validate_entry(meta, body, documents, key)
    revision = hashlib.sha256(json.dumps([raw, files], sort_keys=True).encode()).hexdigest()
    stale = meta["status"] == "reviewed" and meta.get("reviewedHash") != fingerprint(meta, body, files)
    return {"key": key, "path": str(path), "metadata": {**meta, "status": "draft" if stale else meta["status"]},
            "body": body, "markdown": raw, "documents": list(documents.values()), "revision": revision, "reviewStale": stale}


def list_entries(query="", root=WIKI):
    entries, warnings = [], []
    if not Path(root).exists():
        return {"entries": entries, "warnings": warnings}
    for path in sorted(Path(root).glob("*/*/README.md")):
        key = path.parent.relative_to(root).as_posix()
        try:
            entry = read_entry(key, root)
            meta = entry["metadata"]
            if query.casefold() not in " ".join([meta["name"], meta["work"], *meta.get("aliases", [])]).casefold():
                continue
            entries.append({"key": key, "name": meta["name"], "work": meta["work"], "scope": meta["scope"],
                            "status": meta["status"], "aliases": meta.get("aliases", []), "abilities": len(meta.get("abilities", [])), "sources": len(meta["sources"])})
        except (ValueError, OSError, yaml.YAMLError, TypeError, AttributeError) as error:
            warnings.append(f"{key}: {error}")
    return {"entries": entries, "warnings": warnings}


def atomic_write(path, value):
    path.parent.mkdir(parents=True, exist_ok=True)
    with tempfile.NamedTemporaryFile(mode="w", encoding="utf-8", dir=path.parent, prefix=".wiki-", delete=False) as stream:
        temporary = Path(stream.name)
        stream.write(value)
    try:
        temporary.replace(path)
    finally:
        temporary.unlink(missing_ok=True)


def check_revision(key, expected, root):
    path = entry_path(key, root) / "README.md"
    if path.exists():
        current = read_entry(key, root)
        if expected is not None and current["revision"] != expected:
            raise ValueError("This entry changed. Reopen it before saving your revision.")
        return current
    if expected:
        raise ValueError("This entry no longer exists")
    return None


def save_entry(key, value, expected=None, root=WIKI):
    current = check_revision(key, expected, root)
    if not current:
        raise ValueError("Collect evidence before editing an entry")
    meta, body = parse_markdown(value)
    if meta.get("scope") != current["metadata"]["scope"]:
        raise ValueError("Keep the existing canon scope. Create another character entry for a different version.")
    documents = {d["id"]: d for d in current["documents"]}
    validate_entry(meta, body, documents, key)
    meta.update(status="draft", reviewedAt=None, reviewedHash=None)
    atomic_write(safe_path(root, f"{key}/README.md"), markdown(meta, body))
    return read_entry(key, root)


def review_entry(key, expected=None, root=WIKI):
    current = check_revision(key, expected, root)
    if not current or not current["metadata"].get("abilities"):
        raise ValueError("Add cited abilities or traits before reviewing the entry")
    meta = current["metadata"]
    files = {identifier: read_text(safe_path(root, f"{key}/sources/{identifier}.md")) for identifier in meta["sources"]}
    meta.update(status="reviewed", reviewedAt=datetime.now(timezone.utc).isoformat(), reviewedHash=fingerprint(meta, current["body"], files))
    atomic_write(safe_path(root, f"{key}/README.md"), markdown(meta, current["body"]))
    return read_entry(key, root)


def collect(character, urls, supplied="", summary="", refresh=False, expected=None, root=WIKI):
    for field, limit in [("name", 120), ("work", 120), ("scope", 500)]:
        character[field] = text(character.get(field), field, limit)
    key = entry_key(character)
    current = check_revision(key, expected, root)
    if current and any(unicodedata.normalize("NFKC", current["metadata"][field]).casefold() != unicodedata.normalize("NFKC", character[field]).casefold() for field in ["name", "work", "scope"]):
        raise ValueError("This folder already has another identity or canon scope; use a distinct character/version name")
    if not isinstance(urls, list) or len(urls) > 10 or any(not isinstance(u, str) for u in urls):
        raise ValueError("Collect up to 10 URLs")
    urls = list(dict.fromkeys(u.strip() for u in urls if u.strip()))
    supplied = text(supplied, "Supplied evidence", MAX_DOCUMENT, required=False)
    summary = text(summary, "Summary", 80_000, required=False)
    if not urls and not supplied:
        raise ValueError("Provide source URLs or supplied evidence")
    documents = {d["id"]: d for d in current["documents"]} if current else {}
    files = {i: read_text(safe_path(root, f"{key}/sources/{i}.md")) for i in documents}
    original_files = dict(files)
    skipped, updated = 0, 0
    pending = [url for url in urls if refresh or not any(d.get("url") == url for d in documents.values())]
    for url in pending:
        validate_url(url)
    # Reuse saved pages; fetch independent new pages with bounded concurrency.
    with ThreadPoolExecutor(max_workers=4) as pool:
        fetched = dict(zip(pending, pool.map(fetch_document, pending)))
    for url in urls:
        validate_url(url)
        existing = next((d for d in documents.values() if d.get("url") == url), None)
        if existing and not refresh:
            skipped += 1
            continue
        document = fetched[url]
        if existing:
            document["id"] = existing["id"]
        identifier = document["id"]
        value = source_document(document)
        if existing and current:
            fresh = parse_source(value, identifier)
            cited = {r["passage"] for a in current["metadata"].get("abilities", [])
                     for r in a["evidence"] if r["source"] == f"sources/{identifier}.md"}
            missing = [p for p in existing["passages"] if p["id"] in cited
                       and not any(n["id"] == p["id"] for n in fresh["passages"])]
            if missing:
                header, source_body = parse_markdown(value)
                old_dates = {p["id"]: p.get("retrievedAt") for p in existing.get("retainedPassages", [])}
                header["retainedPassages"] = [{"id": p["id"], "retrievedAt": old_dates.get(p["id"], existing.get("retrievedAt"))} for p in missing]
                source_body += "\n\n" + "\n\n".join(f"<!-- passage:{p['id']} -->\n{p['text']}" for p in missing)
                value = markdown(header, source_body)
        documents[identifier] = parse_source(value, identifier)
        files[identifier] = value
        updated += 1
    if supplied:
        value = source_document({"id": "supplied", "title": "Supplied evidence", "url": None,
                                 "retrievedAt": None, "access": "supplied", "language": "en", "text": supplied})
        documents["supplied"] = parse_source(value, "supplied")
        if files.get("supplied") != value:
            files["supplied"] = value
            updated += 1
    meta = deepcopy(current["metadata"]) if current else {
        "formatVersion": 1, "kind": "character-wiki", **character, "aliases": [],
        "language": "en", "status": "draft", "reviewedAt": None, "abilities": []}
    meta["sources"] = list(documents)
    body = summary or (current["body"] if current else f"# {character['name']}\n\nEvidence collected. Add cited findings and record their limits before review.")
    if updated or (summary and (not current or body != current["body"])):
        meta.update(status="draft", reviewedAt=None, reviewedHash=None)
    validate_entry(meta, body, documents, key)
    # Fetch and validate the full batch before writing. Never erase manual findings.
    check_revision(key, current["revision"] if current else expected, root)
    for identifier, value in files.items():
        if original_files.get(identifier) != value:
            atomic_write(safe_path(root, f"{key}/sources/{identifier}.md"), value)
    atomic_write(safe_path(root, f"{key}/README.md"), markdown(meta, body))
    return {"entry": read_entry(key, root), "fetched": updated, "reused": skipped}


def migrate(directory=PROJECT / "sources", root=WIKI):
    imported, skipped = [], []
    for path in sorted(Path(directory).glob("*.json")):
        source = validate_source(json.loads(read_text(path)))
        character = source["character"]
        key = entry_key(character)
        if (entry_path(key, root) / "README.md").exists():
            skipped.append(key)
            continue
        files = {d["id"]: source_document(d) for d in source["documents"]}
        documents = {i: parse_source(value, i) for i, value in files.items()}
        meta = {"formatVersion": 1, "kind": "character-wiki", **character, "aliases": [],
                "language": "en", "status": "draft", "reviewedAt": None, "sources": list(files), "abilities": []}
        body = f"# {character['name']}\n\n" + (source["notes"] or "Imported evidence. Add cited findings before review.")
        validate_entry(meta, body, documents, key)
        for identifier, value in files.items():
            atomic_write(safe_path(root, f"{key}/sources/{identifier}.md"), value)
        atomic_write(safe_path(root, f"{key}/README.md"), markdown(meta, body))
        imported.append(key)
    return {"imported": imported, "skipped": skipped}


def main():
    parser = argparse.ArgumentParser(prog="python3 -m src.wiki", description="Collect, inspect and review a local Markdown character Wiki")
    parser.add_argument("--wiki", type=Path, default=WIKI, help="Local Wiki directory")
    parser.add_argument("--json", action="store_true", help="Return JSON for tools and the website")
    commands = parser.add_subparsers(dest="action", required=True)
    listing = commands.add_parser("list", help="Offline name, work and alias lookup")
    listing.add_argument("--query", default="")
    show = commands.add_parser("show", help="Read an entry without fetching")
    show.add_argument("key")
    collection = commands.add_parser("collect", help="Collect supplied URLs or passages; no automatic synthesis")
    collection.add_argument("--name")
    collection.add_argument("--work")
    collection.add_argument("--scope")
    collection.add_argument("--url", action="append", default=[])
    collection.add_argument("--evidence", default="")
    collection.add_argument("--summary", default="")
    collection.add_argument("--refresh", action="store_true", help="Explicitly fetch saved URLs again")
    collection.add_argument("--expected-revision")
    collection.add_argument("--input", type=Path, help="JSON request file, or - for stdin")
    save = commands.add_parser("save", help="Validate and save edited Markdown as a draft")
    save.add_argument("key")
    save.add_argument("--file", type=Path)
    save.add_argument("--expected-revision")
    save.add_argument("--input", type=Path)
    review = commands.add_parser("review", help="Record explicit human review of cited findings")
    review.add_argument("key")
    review.add_argument("--expected-revision")
    migration = commands.add_parser("migrate", help="Copy legacy JSON into Markdown without deleting originals")
    migration.add_argument("--sources", type=Path, default=PROJECT / "sources")
    commands.add_parser("categories", help="Read profile-agnostic classification definitions")
    args = parser.parse_args()
    try:
        value = {}
        if getattr(args, "input", None):
            raw = sys.stdin.read(MAX_SOURCE + 1) if str(args.input) == "-" else read_text(args.input)
            if len(raw.encode()) > MAX_SOURCE:
                raise ValueError("Request is too large")
            value = json.loads(raw)
            if not isinstance(value, dict):
                raise ValueError("Expected a JSON request object")
        if args.action == "list": result = list_entries(args.query, args.wiki)
        elif args.action == "show": result = read_entry(args.key, args.wiki)
        elif args.action == "categories": result = categories()
        elif args.action == "migrate": result = migrate(args.sources, args.wiki)
        elif args.action == "review": result = review_entry(args.key, args.expected_revision, args.wiki)
        elif args.action == "save":
            raw = value.get("markdown") if value else read_text(args.file) if args.file else None
            result = save_entry(args.key, raw, value.get("revision", args.expected_revision), args.wiki)
        else:
            character = value.get("character", {"name": args.name, "work": args.work, "scope": args.scope})
            result = collect(character, value.get("urls", args.url), value.get("supplied", args.evidence),
                             value.get("summary", args.summary), value.get("refresh", args.refresh),
                             value.get("revision", args.expected_revision), args.wiki)
        if args.json: print(json.dumps(result, ensure_ascii=False))
        elif args.action in {"show", "save", "review"}: print(result["markdown"], end="")
        elif args.action == "collect": print(result["entry"]["path"] + "/README.md")
        else: print(yaml.safe_dump(result, allow_unicode=True, sort_keys=False), end="")
    except Exception as error:
        print(json.dumps({"error": str(error)}) if args.json else f"Error: {error}", file=sys.stdout if args.json else sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
