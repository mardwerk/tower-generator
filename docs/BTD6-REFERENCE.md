# BTD6 reference

[btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas) owns the BTD6 captures, source cleanup and BTD6 Profile.
The initial reference is capture 56.3, Steam build 24829026, at revision [`a380413`](https://github.com/KyleDerZweite/btd6-atlas/tree/a380413ed809c98654acec7aa37cba40f80bf5c5/data/56.3-build-24829026).
Atlas game data and derived tables are licensed CC BY-NC 4.0. Cite the capture, build, actual source commit and source files when importing or using values.

The historical generator used this capture for its BTD6-based Profile. Its source mapping and deliberate departures remain available at the [generator history tag](https://github.com/mardwerk/tower-generator/tree/generator-before-validator-reset-2026-09-29/docs/BTD6-REFERENCE.md).

Atlas's BTD6 Profile describes BTD6 data. It is separate from Tower Generator's future Default Profile and the generic Profile Template supplied by [td-profile](https://github.com/mardwerk/td-profile).
No Default Profile or normalized Tower mapping is accepted in current source. Review Dart Monkey with Kyle through [#85](https://github.com/mardwerk/tower-generator/issues/85) before expanding the import.

[Profile integration](PROFILE.md) documents the shared validator setup. Keep generated records local and gitignored; do not publish the BTD6 corpus in this repository.
