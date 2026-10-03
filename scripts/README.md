# Python helpers

Put future independent maintenance and research helpers in this directory.
No active helpers remain after retiring the previous application and setup script.

Use system Python by default and prefer its standard library.
Run a helper from the repository root with `python3 -B scripts/name.py`.
No Python project, dependency installation or repository virtual environment is required.

[uv](https://docs.astral.sh/uv/guides/scripts/) is optional.
`uv run --no-project scripts/name.py` skips Python project discovery and may reuse a discovered virtual environment.
If a helper needs a third-party package, declare that dependency in the script with PEP 723 metadata.
uv manages the script's isolated environment automatically; isolation still uses a virtual environment.
Do not add dependencies or Python project files before a helper needs them.

The archived Python experiment stays in [the research snapshot](../docs/research/unit-design-research-snapshot/README.md).
Its eight Python files use only the standard library and bundled modules.
Keep them in their original locations and preserve their bytes so the snapshot manifest remains valid.
They are historical research evidence, not active product code or maintenance helpers.
