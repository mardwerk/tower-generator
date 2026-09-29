# Tower Generator

Tower Generator is preparing a standalone Profile Validator. The validator is not implemented. The previous character-to-tower generator and UnitLab were removed from active source after Kyle preserved their history at the [generator-before-validator-reset-2026-09-29 tag](https://github.com/mardwerk/tower-generator/tree/generator-before-validator-reset-2026-09-29). That tag points to `c157cacdfdaf5e9405a01c8d16a3a6fa1147685b`.

A Profile is a directory that defines a game's Tower format and rules. The planned validator will check Profile consistency and Tower files against a Profile. The accepted scope and open design questions are tracked in [#85](https://github.com/mardwerk/tower-generator/issues/85), [#90](https://github.com/mardwerk/tower-generator/issues/90), and [#92](https://github.com/mardwerk/tower-generator/issues/92). The exact schema, importer mapping, file placement, and validator interface still need review.

The former Lab's reusable visual design remains in [`src/web/ui/`](src/web/ui/), [`src/web/app/styles.css`](src/web/app/styles.css), and [`src/web/public/mardwerk.png`](src/web/public/mardwerk.png). The screens, workflows, server, and built app were removed. These files are design assets for a future Lab, not a runnable application. With Node.js 22 or newer, run `pnpm install`, `pnpm typecheck`, and `pnpm format:check` to check the retained components and styles.

[BTD6 reference and attribution](docs/BTD6-REFERENCE.md) records the pinned atlas source for future Default Profile work. No Default Profile is implemented in current source.
