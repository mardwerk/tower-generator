# Tower Generator

The current website focuses on character research and a reusable local Wiki.
Research a character by name, inspect cited abilities and traits, and record manual review before returning to Tower generation.
Research stays independent of a game's Profile.
The [product outline](docs/PRODUCT.md) records the rough rules under review before further implementation.
It records Go as the selected backend direction, with proposed screens, shared CLI/API responsibilities and bounded concurrency.
The current implementation below still uses Python and a Node web bridge; the Go migration awaits its workflow specification.

For website development, use Python 3.10 or newer and Node.js 22.12 or newer:

```sh
python3 -m venv .venv
source .venv/bin/activate
python3 -m pip install -r requirements.txt
pnpm install
pnpm dev
```

Open the URL printed by Vite. Create shows the character input, Research action and small controls underneath.
Research finds English sources and resolves the series automatically. Repeating it verifies and expands the saved Wiki entry with the same budget.
Research inputs provides optional series/scope hints or supplied references; Wiki looks up local entries.
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
The [initial Tower Generator skill](.agents/skills/tower-generator/SKILL.md) supports cited, reviewable designs within the current authoring limits.
[Focused research](docs/research/README.md) preserves the separate research state and describes later SkillOpt evaluation.
Name-only research uses the configured OpenRouter model and web search. Suggested classifications use the broad local categories; specific [Jev benchmarking](https://github.com/mardwerk/tower-generator/issues/105) remains later work.

The Python CLI owns persistence and validation and can run without Node or the website:

```sh
python3 -B -m src.wiki research usopp
```

`pnpm wiki` is a convenience shortcut. Node currently runs the website tooling and local API bridge, which invokes the Python CLI.
Run `pnpm wiki research usopp` to research or improve Usopp, using `OPENROUTER_API_KEY` and `OPENROUTER_MODEL` from the environment or ignored local `.env`.
Use `pnpm wiki --help`, `pnpm wiki list`, `pnpm wiki show one-piece/usopp`, or `pnpm wiki categories` without a browser.
Use `pnpm wiki migrate` to copy earlier local JSON sources into the new layout; original files and existing Wiki entries are preserved.
The tool has no hidden session or database. Lookup rebuilds a small index from local files.

Run `pnpm test`, `pnpm typecheck`, `pnpm format:check` and `pnpm build` for the focused checks.
`pnpm preview` serves the built website with the same local commands.
Wiki storage is recorded in [#102](https://github.com/mardwerk/tower-generator/issues/102), and automatic research in [#103](https://github.com/mardwerk/tower-generator/issues/103).

The earlier [Usopp experiment](docs/USOPP.md) and [Default Profile](default/profile/) remain in the repository, outside the website's current focus.
[Authoring](docs/AUTHORING.md) describes the separate experiment commands; [Profile integration](docs/PROFILE.md) documents its checker setup.
The [reference attribution](docs/BTD6-REFERENCE.md) records Atlas provenance and licenses.
The previous generator remains in the [generator-before-validator-reset-2026-09-29 tag](https://github.com/mardwerk/tower-generator/tree/generator-before-validator-reset-2026-09-29).
See [issue labels](docs/ISSUE_LABELS.md) for tracking conventions.
