# Tower Generator

The current website focuses on character research and a reusable local Wiki.
Collect evidence, inspect cited abilities and traits, and record manual review before returning to Tower generation.
Research stays independent of a game's Profile.

Use Python 3.10 or newer and Node.js 22.12 or newer:

```sh
python3 -m venv .venv
source .venv/bin/activate
python3 -m pip install -r requirements.txt
pnpm install
pnpm dev
```

Open the URL printed by Vite. Create shows the character input, Research action and small controls underneath.
Research inputs opens series, scope, English references and optional notes; Wiki looks up local entries.
Edit entry saves readable Markdown with YAML metadata, cited records and suggested classification tags.
Each page has one heading and no subheadings, using the retained Lab design system.
The website does not require Go, an installed checker or the sibling Atlas checkout.

```text
wiki/                         Local, Git ignored
  one-piece/
    usopp/
      README.md               Character summary, cited records and review status
      sources/*.md            Retained source metadata and passages
docs/wiki-format.md           Portable file contract
docs/research-categories.yaml Profile-agnostic classification definitions
```

The [research workflow](docs/RESEARCH.md) explains collection, offline reuse, explicit refresh, editing and manual review.
The [Wiki format](docs/wiki-format.md) documents the human-readable files and citations.
Automatic web search, model synthesis and Jev classification are later steps after the first entries have been reviewed.

The CLI owns persistence and validation. The website invokes that same CLI and provides the user interface.
Use `pnpm wiki --help`, `pnpm wiki list`, `pnpm wiki show one-piece/usopp`, or `pnpm wiki categories` without a browser.
Use `pnpm wiki migrate` to copy earlier local JSON sources into the new layout; original files and existing Wiki entries are preserved.
The tool has no hidden session or database. Lookup rebuilds a small index from local files.

Run `pnpm test`, `pnpm typecheck`, `pnpm format:check` and `pnpm build` for the focused checks.
`pnpm preview` serves the built website with the same local commands.
Current implementation work is recorded in [#102](https://github.com/mardwerk/tower-generator/issues/102).

The earlier [Usopp experiment](docs/USOPP.md) and [Default Profile](default/profile/) remain in the repository, outside the website's current focus.
[Authoring](docs/AUTHORING.md) describes the separate experiment commands; [Profile integration](docs/PROFILE.md) documents its checker setup.
The [reference attribution](docs/BTD6-REFERENCE.md) records Atlas provenance and licenses.
The previous generator remains in the [generator-before-validator-reset-2026-09-29 tag](https://github.com/mardwerk/tower-generator/tree/generator-before-validator-reset-2026-09-29).
See [issue labels](docs/ISSUE_LABELS.md) for tracking conventions.
