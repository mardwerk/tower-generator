# Tower Generator

Create and inspect the editable [Usopp draft](default/usopp.json) with the [Default Profile](default/profile/).
The first experiment has 64 ordinary states, 15 upgrades and no Paragon.
[Usopp canon and adaptations](docs/USOPP.md) explains the exploding stars, Pop Greens and Sogeking Styles.
General character generation and game integration remain later work.

Use Python 3.10 or newer, Go, Git, GitHub CLI and Node.js 22.12 or newer:

```sh
python3 setup.py
pnpm install
pnpm dev
```

Open the local URL printed by Vite.
The web view shows upgrade paths, legal crosspaths, stats, Tower JSON and validation results from the generated files.
Edit `default/usopp.json`, run `python3 setup.py` again, then click Reload in the web view.

`setup.py` builds the latest published [td-profile](https://github.com/mardwerk/td-profile) release for the current OS and architecture.
It reuses a verified matching build, generates Usopp using the sibling [Atlas](https://github.com/KyleDerZweite/btd6-atlas) checkout, and checks the candidate and missing-record cases.
Use `--atlas /path/to/btd6-atlas` for another checkout.

The default workspace keeps related files together:

```text
default/
  usopp.json              Editable design
  profile/               Self-contained Profile
  game-data/             Generated Tower and Upgrade records
  profile-validator      Native checker, .exe on Windows
  validator-release.json Installed release and build identity
  checks.json            Candidate and negative-case reports
  TOWER.md               Readable Tower
  REFERENCE.md           Compact Dart reference
  source.json            Atlas provenance
  licenses/              Installed checker notices
```

Generated data, reports and executables stay local and ignored by Git.
[Authoring](docs/AUTHORING.md) explains the mapping and current limits; [Profile integration](docs/PROFILE.md) explains the checker.
[Reference attribution](docs/BTD6-REFERENCE.md) records the Atlas source and licenses.

Run `pnpm test`, `pnpm typecheck`, `pnpm format:check` and `pnpm build` for the focused checks.
`pnpm preview` serves the built web view against the same local default files.

The previous generator remains in the [generator-before-validator-reset-2026-09-29 tag](https://github.com/mardwerk/tower-generator/tree/generator-before-validator-reset-2026-09-29).
The web view reuses its retained Lab theme and components.
Product work is tracked in [#85](https://github.com/mardwerk/tower-generator/issues/85), documentation in [#92](https://github.com/mardwerk/tower-generator/issues/92), and deferred presentation changes in [#100](https://github.com/mardwerk/tower-generator/issues/100).
See [issue labels](docs/ISSUE_LABELS.md) for tracking conventions.
