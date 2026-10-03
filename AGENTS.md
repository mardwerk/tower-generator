# Agent instructions

Read [README.md](README.md) before changing this repository.
Mardwerk agents with private Planning access must also read its [AGENTS.md](https://github.com/mardwerk/planning/blob/main/AGENTS.md) and master [CONTEXT.md](https://github.com/mardwerk/planning/blob/main/CONTEXT.md).
Public contributors do not need Planning access.

When managing issues, labels or milestones, read [local topic definitions](docs/ISSUE_LABELS.md).
Agents with Planning access must also read its [issue labels standard](https://github.com/mardwerk/planning/blob/main/10_docs/ISSUE_LABELS.md), using the local checkout when available.
Derive work state from explicit, recorded owner decisions.
Relabelling alone does not select work or accept a specification.

Follow [product rules](docs/PRODUCT.md) and the [realignment plan](docs/REALIGNMENT.md).
Go `serve` is the selected API-only direction; there is no independent operational CLI.
The old Python product, Node tooling, full web/Lab and Default Profile/Usopp experiment are retired and preserved in Git history.
Do not present historical commands, schemas or proposals as the new backend contract.
Use the selected pure Go backend with REST, SSE and a bounded in-process research queue.
Specify queue limits, interruption, retention and the research workflow before product implementation.

Use [td-profile](https://github.com/mardwerk/td-profile) for the generic Profile Validator and reusable schemas.
Propose shared checker or schema changes upstream.
Follow [Profile ownership](docs/PROFILE.md); a new Profile and its integration are deferred.
A Profile is a directory selected by path, with local dependencies and separate governed Tower records.
Review current product rules and the research/Wiki workflow with Kyle through [#106](https://github.com/mardwerk/tower-generator/issues/106).
When generation resumes, review new mappings and generation rules with Kyle in the issue that owns the proposed change before implementation.

Keep BTD6 facts, source cleanup and the BTD6 Profile in [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas).
Cite capture, build, source commit and files as described in [Profile attribution](docs/PROFILE.md#btd6-source-attribution).
Keep downloaded bundles, runtime binaries and generated game-data out of Git.
Put progress and skip reasons in the terminal rather than Tower JSON.

Retain independent research helpers under `scripts/`.
Keep references to the retained [research documents](docs/research/README.md) current.
Their source proposals and historical commands do not override this repository's product rules.

Wrap prose so each physical line contains at most two sentences.
Use headings when useful and put explanatory text between a heading and any subheading.
Before committing, run `git diff --check` and check changed links.
Run checks appropriate to executable changes; document what each check establishes.
