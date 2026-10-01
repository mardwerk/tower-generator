# Tower Generator

Tower Generator's initial [Default Profile](default-profile/) copies Atlas's verified BTD6 Profile and uses [td-profile](https://github.com/mardwerk/td-profile)'s validator and reusable schemas.
The character-to-tower generator and additional authoring checks remain to be implemented.

A Profile is a directory that defines a game's Tower format and rules. The shared validator checks a supplied Profile and separate game-data offline.
[Profile integration](docs/PROFILE.md) provides the pinned release setup and cross-repository iteration commands. [#98](https://github.com/mardwerk/tower-generator/issues/98) tracks this preparation; [#85](https://github.com/mardwerk/tower-generator/issues/85) tracks further Profile and generation work.
Dart Monkey is the verified Atlas reference. Additional product rules and the first generated Tower remain to be developed with Kyle.
[#92](https://github.com/mardwerk/tower-generator/issues/92) tracks documentation of the resulting product contract.

Run `python3 scripts/check-default-profile.py` with the sibling Atlas checkout present to check Dart against both Profiles.
Use `--local-checker` to rebuild the sibling td-profile checker while developing shared checks.

The previous generator and UnitLab remain in the [generator-before-validator-reset-2026-09-29 tag](https://github.com/mardwerk/tower-generator/tree/generator-before-validator-reset-2026-09-29), at commit `c157cacdfdaf5e9405a01c8d16a3a6fa1147685b`.
The former Lab's reusable visual design remains in [`src/web/ui/`](src/web/ui/), [`src/web/app/styles.css`](src/web/app/styles.css), and [`src/web/public/mardwerk.png`](src/web/public/mardwerk.png).
These files are design assets for a future Lab. With Node.js 22 or newer, run `pnpm install`, `pnpm typecheck`, and `pnpm format:check` to check the retained components and styles.

[BTD6 reference and attribution](docs/BTD6-REFERENCE.md) records the Atlas source. Atlas's BTD6 Profile and Tower Generator's Default Profile have independent identities and revisions.

[Issue labels](docs/ISSUE_LABELS.md) explains the label families and local topics used to track work.
