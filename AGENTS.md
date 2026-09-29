# Agent instructions

Read [README.md](README.md) before changing this repository. Mardwerk agents with access to private [Planning](https://github.com/mardwerk/planning) must also read its [AGENTS.md](https://github.com/mardwerk/planning/blob/main/AGENTS.md) and master [CONTEXT.md](https://github.com/mardwerk/planning/blob/main/CONTEXT.md). Contributors can use this repository without Planning access.

The current product work is the Profile Validator tracked in [#85](https://github.com/mardwerk/tower-generator/issues/85). It is not implemented. A Profile is a directory selected by path. Do not treat the removed generator's schemas, CLI, or workflows as the validator contract. Review the Dart Monkey sample and owner decisions before defining the importer, schema, record placement, or validator behavior.

Keep BTD6 facts in [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas); cite the pinned source in [BTD6-REFERENCE.md](docs/BTD6-REFERENCE.md). Preserve the reusable Lab design files in `src/web/ui/`, `src/web/app/styles.css`, and `src/web/public/mardwerk.png` for a future Lab.

Before committing, run `git diff --check` and check changed links. For design file changes, also run `pnpm typecheck` and `pnpm format:check`. Run checks appropriate to any new executable code once it exists.
