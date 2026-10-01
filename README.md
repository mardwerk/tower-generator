# Tower Generator

Tower Generator is preparing its Default Profile and product checks using [td-profile](https://github.com/mardwerk/td-profile)'s validator and reusable schemas.
The Default Profile and character-to-tower generator are not implemented in current source.

A Profile is a directory that defines a game's Tower format and rules. The shared validator checks a supplied Profile and separate game-data offline.
[Profile integration](docs/PROFILE.md) provides the pinned release setup and verification commands. [#96](https://github.com/mardwerk/tower-generator/issues/96) tracks this preparation; [#85](https://github.com/mardwerk/tower-generator/issues/85) tracks the Default Profile and import work.
The Dart Monkey mapping and product rules still need review with Kyle. [#92](https://github.com/mardwerk/tower-generator/issues/92) tracks documentation of the resulting product contract.

The previous generator and UnitLab remain in the [generator-before-validator-reset-2026-09-29 tag](https://github.com/mardwerk/tower-generator/tree/generator-before-validator-reset-2026-09-29), at commit `c157cacdfdaf5e9405a01c8d16a3a6fa1147685b`.
The former Lab's reusable visual design remains in [`src/web/ui/`](src/web/ui/), [`src/web/app/styles.css`](src/web/app/styles.css), and [`src/web/public/mardwerk.png`](src/web/public/mardwerk.png).
These files are design assets for a future Lab. With Node.js 22 or newer, run `pnpm install`, `pnpm typecheck`, and `pnpm format:check` to check the retained components and styles.

[BTD6 reference and attribution](docs/BTD6-REFERENCE.md) records the Atlas source. Atlas's BTD6 Profile is separate from Tower Generator's future Default Profile.

[Issue labels](docs/ISSUE_LABELS.md) explains the label families and local topics used to track work.
