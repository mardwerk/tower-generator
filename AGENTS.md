# Agent instructions

Read [README.md](README.md) before changing this repository. Mardwerk agents with access to private [Planning](https://github.com/mardwerk/planning) must also read its [AGENTS.md](https://github.com/mardwerk/planning/blob/main/AGENTS.md) and master [CONTEXT.md](https://github.com/mardwerk/planning/blob/main/CONTEXT.md).
Contributors can use this repository without Planning access.

When managing issues, labels or milestones, Mardwerk agents with Planning access must read its internal [issue labels standard](https://github.com/mardwerk/planning/blob/main/10_docs/ISSUE_LABELS.md), using `planning/10_docs/ISSUE_LABELS.md` in the local Planning checkout when available and the GitHub link otherwise. Also read the [local topic definitions](docs/ISSUE_LABELS.md) and derive work state from recorded owner decisions.

Use [td-profile](https://github.com/mardwerk/td-profile) for the generic Profile Validator and reusable schemas. Follow the pinned contract and setup in [docs/PROFILE.md](docs/PROFILE.md); propose shared checker or schema changes upstream.
Tower Generator owns [default/profile/](default/profile/), authoring and product checks. Dart Monkey's Atlas verification is accepted; review new generation mappings and product rules with Kyle through [#85](https://github.com/mardwerk/tower-generator/issues/85).
A Profile is a directory selected by path. Keep its dependencies local and its governed Tower records separate.
Do not treat the removed generator's schemas, CLI or workflows as the validator contract.

Keep BTD6 facts, source cleanup and the BTD6 Profile in [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas); cite the capture, build, source commit and files as described in [BTD6-REFERENCE.md](docs/BTD6-REFERENCE.md).
Keep downloaded bundles, runtime binaries and generated game-data out of Git. Put progress and skip reasons in the terminal rather than Tower JSON.
Use the retained Lab design system for the web UI. Each page has one visible heading and no subheadings.
Preserve the reusable Lab design files in `src/web/ui/`, `src/web/app/styles.css`, and `src/web/public/mardwerk.png` for a future Lab.

Wrap prose so each physical line contains at most two sentences.
Use headings when useful and put explanatory text between a heading and any subheading.

Before committing, run `git diff --check` and check changed links. For design file changes, also run `pnpm typecheck` and `pnpm format:check`.
Run checks appropriate to executable changes. For td-profile setup changes, verify the documented empty-data and synthetic-example commands and a missing Profile dependency.
For Default Profile or checker integration changes, run `python3 setup.py` and `pnpm test`.
For web changes, also run `pnpm build` and inspect the local view in the browser.
