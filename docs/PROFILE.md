# Profile integration

Tower Generator will use [td-profile](https://github.com/mardwerk/td-profile)'s generic validator and reusable schemas.
This setup establishes a working upstream baseline; the Default Profile and Dart Monkey mapping still need review through [#85](https://github.com/mardwerk/tower-generator/issues/85).

## Ownership and contract

Each repository owns a separate part of the work.

| Repository | Responsibility |
| --- | --- |
| [td-profile](https://github.com/mardwerk/td-profile) | Generic checker, reusable schemas and copyable Profile Template |
| [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas) | BTD6 captures, source cleanup, BTD6 settings and source contracts |
| Tower Generator | Default Profile, imported records, authoring and product checks |

Use the [td-profile v1.0.2 contract](https://github.com/mardwerk/td-profile/blob/v1.0.2/docs/PROFILE.md) and its self-contained Profile Template.
The manifest declares the Profile's identity, revision, checker interface and local dependencies. Settings bind source fields and define collections, references, mechanics, numerical units, progression and scoring.
The checker reads an explicit Profile path and separate game-data without rewriting the records. Validation runs offline after setup.
Propose shared checker or schema changes in td-profile, then adopt a released version here.

The bundled Template uses one path with tiers 0 through 2 and one Tower file per state, with separate Upgrade definitions.
These are editable starter settings. Tower Generator's three-path, five-upgrade mapping and record layout still need the Dart Monkey review.
The Template, Atlas's BTD6 Profile and Tower Generator's future Default Profile remain separate.

## Install the pinned baseline

These commands run from this repository's root on Linux amd64 with [GitHub CLI](https://cli.github.com/), `sha256sum` and `tar` installed.
The release bundles the binary, Profile Template, empty data directory, synthetic example, licenses and release identity; no Go build or sibling checkout is required.

```sh
mkdir -p .tools/td-profile
gh release download v1.0.2 --repo mardwerk/td-profile \
  --pattern td-profile-v1.0.2-linux-amd64.tar.gz \
  --pattern td-profile-v1.0.2-linux-amd64.tar.gz.sha256 \
  --dir .tools/td-profile --skip-existing
(
  cd .tools/td-profile &&
  sha256sum --check td-profile-v1.0.2-linux-amd64.tar.gz.sha256 &&
  tar -xzf td-profile-v1.0.2-linux-amd64.tar.gz
)
```

Release [v1.0.2](https://github.com/mardwerk/td-profile/releases/tag/v1.0.2) uses checker interface 5 and Template revision 1.0.0.
The Linux archive's SHA256 is `830f263683bbe895ed0be51a2e1c9afd4b9cc30b5d93d57d3ad33fb6dfdadd20`.
Keep the bundle's `release.json` and notices with the executable. Other native targets are available from the same release.
`.tools/` and local `game-data/` are gitignored.

## Verify the baseline

Run the released checker against its unchanged Template and supplied data before adapting product settings.

```sh
td_profile_bundle=.tools/td-profile/td-profile-v1.0.2-linux-amd64
"$td_profile_bundle/validator" --profile "$td_profile_bundle/profile" \
  --game-data "$td_profile_bundle/game-data" --format text
"$td_profile_bundle/validator" --profile "$td_profile_bundle/profile" \
  --game-data "$td_profile_bundle/examples/minimal-game/game-data" --format text
"$td_profile_bundle/validator" score-tower --profile "$td_profile_bundle/profile" \
  --game-data "$td_profile_bundle/examples/minimal-game/game-data" \
  --tower Towers/Bolt/Bolt-0.json --format text
```

The empty directory should pass with zero files. The synthetic example should pass and Bolt should score 100/100.
A score measures compliance with declared checks; it does not establish gameplay balance or simulate behavior.

For the negative check, copy the bundled Profile into a temporary directory, remove that copy's `mechanics.json`, and validate the synthetic data using the copy's Profile path.
The command must fail and identify the missing dependency. Remove the temporary copy after checking.

Verified on 2026-10-01 on Linux amd64 with checker 5.0.1, from td-profile commit `48fc5daa723d13d89d33bce4a863ba5bf38a1b39`.
The unchanged Template's dependency SHA256 was `981def0d205941af5535501f93b19abfa73254bb1331f4db1f42647975dae5a7`.

| Check | Observed result |
| --- | --- |
| Empty data | Valid, zero files and zero errors |
| Synthetic example | Valid, five files and zero errors |
| Bolt score | 100/100, complete |
| Missing `mechanics.json` in a temporary Profile copy | Exit 2, missing dependency identified |

## Next product iteration

Start with one reviewed Dart Monkey example and record its source evidence using [BTD6-REFERENCE.md](BTD6-REFERENCE.md).
Keep its 64 ordinary states, 15 upgrades, purchase links and relevant mechanics explicit; exclude Paragon from this sample.
Review how the Atlas data maps into the shared contract before writing the Default Profile or expanding the corpus.
Keep any supplied data unchanged during validation and keep generated import data out of Git.

The next product checks should demonstrate a valid sample and failures for a missing required state or upgrade.
Record the release identity, Profile revision and checks actually performed. Document the accepted product contract through [#92](https://github.com/mardwerk/tower-generator/issues/92).
Character generation and game integration remain later work. Defer SkillOpt until validator evaluation cases are stable.
