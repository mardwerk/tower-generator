"""Reproduce the UG-T02 fixture audit without modifying the original files.

Run from any directory with Python 3. JSON Schema's additional probe requires
jsonschema; without it the probe is reported as not run. No network is used.
This tests a static reference resolver, not a game or an AI generator.
"""
from __future__ import annotations

import copy
import hashlib
import importlib.util
import json
import platform
import shutil
import sys
from pathlib import Path
from typing import Any


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def main() -> None:
    root = Path(__file__).resolve().parents[1]
    original = root / "original_reference"
    required = ("check_reference.py", "generalized_wizard.unit.json", "unit-draft.schema.json",
                "validation_report.json", "resolved_builds.json")
    missing = [name for name in required if not (original / name).is_file()]
    if missing:
        raise FileNotFoundError(f"Missing original reference files: {missing}")
    out = Path(__file__).resolve().parent / "reproduced"
    out.mkdir(exist_ok=True)
    for name in required:
        shutil.copy2(original / name, out / name)
    module_spec = importlib.util.spec_from_file_location("ug_t02_original_checker", out / "check_reference.py")
    if module_spec is None or module_spec.loader is None:
        raise ImportError("Could not load the original checker")
    checker = importlib.util.module_from_spec(module_spec)
    module_spec.loader.exec_module(checker)
    input_path = out / "generalized_wizard.unit.json"
    report = checker.run(input_path)
    spec = json.loads(input_path.read_text())
    reversed_order: dict[str, Any]
    try:
        checker.resolve(spec, (2, 0, 0), ["A2", "A1"])
        reversed_order = {"outcome": "accepted", "scope": "membership checked, sequential prerequisites not checked"}
    except Exception as exc:
        reversed_order = {"outcome": "rejected", "exception": f"{type(exc).__name__}: {exc}"}
    binding_probe: dict[str, Any]
    try:
        import jsonschema
    except ImportError:
        binding_probe = {"outcome": "not_run", "reason": "jsonschema is not installed"}
    else:
        probe = copy.deepcopy(spec)
        for group in probe["binding_groups"].values():
            group["status"] = "bound"
            group["values"] = None
        schema = json.loads((out / "unit-draft.schema.json").read_text())
        try:
            jsonschema.validate(probe, schema)
            binding_probe = {"outcome": "accepted", "scope": "document shape, not typed binding values"}
        except jsonschema.ValidationError as exc:
            binding_probe = {"outcome": "rejected", "error": exc.message}
    result = {
        "origin": "new verification during consolidation, not the historical run",
        "python": sys.version,
        "platform": platform.platform(),
        "input_sha256": sha256(input_path),
        "rerun_report": report,
        "saved_report_byte_identical": (out / "validation_report.json").read_bytes() == (original / "validation_report.json").read_bytes(),
        "saved_builds_byte_identical": (out / "resolved_builds.json").read_bytes() == (original / "resolved_builds.json").read_bytes(),
        "reversed_prerequisite_probe": reversed_order,
        "null_bound_values_schema_probe": binding_probe,
        "interpretation": "No runtime, sequential-state, combat, balance or model-generation evidence follows. The original resolver normalizes node order before composition."
    }
    dest = out / "audit_reproduction.json"
    dest.write_text(json.dumps(result, indent=2) + "\n")
    print(json.dumps({"audit_output": str(dest), "report_matches": result["saved_report_byte_identical"],
                      "builds_match": result["saved_builds_byte_identical"],
                      "reversed_order": reversed_order["outcome"],
                      "null_binding_schema": binding_probe["outcome"]}, indent=2))


if __name__ == "__main__":
    main()
