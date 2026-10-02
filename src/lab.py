"""Small local Lab operations shared by the web app and command-line checks."""

from contextlib import redirect_stdout
import json
from pathlib import Path
import shutil
import sys
import tempfile

from src.author import PROJECT, check, generate, validate_design, write_json
from src.research import MAX_SOURCE, collect_source, list_sources, load_source, save_source

DEFAULT = PROJECT / "default"


def build_example(design):
    validate_design(design)
    # ponytail: generalize the workspace after a second authored Tower is accepted.
    if design["id"] != "Usopp":
        raise ValueError("This authoring experiment rebuilds Usopp. Other character sources can be saved and exported.")
    if not design.get("sourceId"):
        raise ValueError("Choose a saved character source before building")
    binary = "profile-validator.exe" if sys.platform == "win32" else "profile-validator"
    if not (DEFAULT / binary).exists():
        raise ValueError("Run python3 setup.py once to install the validator")
    # Check a disposable candidate before changing the saved design or generated files.
    with tempfile.TemporaryDirectory(prefix="tower-lab-") as temporary:
        candidate = Path(temporary)
        shutil.copytree(DEFAULT / "profile", candidate / "profile")
        for name in [binary, "validator-release.json"]:
            shutil.copy2(DEFAULT / name, candidate / name)
        with redirect_stdout(sys.stderr):
            generate(design, PROJECT.parent / "btd6-atlas", candidate)
            check(candidate, design["id"])
        shutil.copytree(candidate / "game-data", DEFAULT / "game-data", dirs_exist_ok=True)
        for name in ["TOWER.md", "REFERENCE.md", "source.json", "character-source.json", "checks.json"]:
            shutil.copy2(candidate / name, DEFAULT / name)
        write_json(DEFAULT / "usopp.json", design)
    return {"message": "Usopp saved, built and checked."}


def perform(action, value):
    if action == "list":
        return {"sources": [{"id": source["id"], "character": source["character"],
                             "documents": len(source["documents"])} for source in list_sources()]}
    if action == "load":
        return load_source(value.get("id"))
    if action == "collect":
        return collect_source(value.get("character", {}), value.get("urls", []),
                              value.get("supplied", ""), value.get("notes", ""))
    if action == "import":
        return save_source(value)
    if action == "build":
        return build_example(value)
    raise ValueError("Unknown Lab operation")


if __name__ == "__main__":
    try:
        raw = sys.stdin.read(MAX_SOURCE + 1)
        if len(raw.encode()) > MAX_SOURCE:
            raise ValueError("Request is too large")
        value = json.loads(raw or "{}")
        if not isinstance(value, dict):
            raise ValueError("Expected a JSON object")
        result = perform(sys.argv[1], value)
        print(json.dumps(result, ensure_ascii=False))
    except Exception as error:
        # Keep fetch and validation failures visible, without partial source saves.
        print(json.dumps({"error": str(error)}, ensure_ascii=False))
        sys.exit(1)
