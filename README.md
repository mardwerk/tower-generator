# Tower Generator

Tower Generator will turn reusable character evidence into inspectable Towers for a selected Profile.
The current focus is a complete Go research and Wiki backend; Tower design follows later.

Kyle selected a greenfield, API-only Go `serve` backend.
Research, lookup, editing and review clients will use its API without a browser session or independent operational CLI.
The previous Python application, Node bridge, website and Default Profile/Usopp experiment are retired.
There is no runnable product backend in this reset.

The [product rules](docs/PRODUCT.md) distinguish owner requirements from proposed design.
The [realignment plan](docs/REALIGNMENT.md) records the reset and remaining decisions.
Kyle selected pure Go with REST, SSE and a bounded in-process research queue.
Restart recovery is not a major requirement; detailed interruption and retention behavior remains to be specified.

The backend will own local, human-readable Markdown files under `wiki/<series>/<character>/`, ignored by Git.
Research remains in English and independent of game Profiles.
Repeated research must protect retained evidence, manual notes and human-reviewed records.
See the [research rules](docs/RESEARCH.md) and [candidate Wiki format](docs/wiki-format.md).

The [Tower design skill](.agents/skills/tower-generator/SKILL.md) supports cited concepts without claiming an implemented exporter.
[Research evidence](docs/research/README.md) and independent helpers under [scripts/](scripts/) remain available for review.
The retained research documents are reference evidence, not current product instructions or a runtime.

[td-profile](https://github.com/mardwerk/td-profile) owns the generic Profile Validator and reusable schemas.
[Profile ownership and BTD6 attribution](docs/PROFILE.md) remain separate from the research backend.
A new Profile and its generation integration are deferred.
See [issue labels](docs/ISSUE_LABELS.md) for tracking conventions.

The [latest prototype](https://github.com/mardwerk/tower-generator/tree/ea114b2139abf46318a59f632a58fc035f64a5a7) preserves the retired application and Lab.
The [complete Default Profile experiment](https://github.com/mardwerk/tower-generator/tree/c7a71f25147ee871bf48790e6e5b0cd6515d8cfa) preserves its contracts, notices and setup.
