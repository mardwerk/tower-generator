# Profile integration

Tower Generator uses [td-profile](https://github.com/mardwerk/td-profile)'s generic validator and reusable schemas.
The initial [Default Profile](../default-profile/) copies Atlas's verified BTD6 Profile, with its own identity and revision; further generation work continues through [#85](https://github.com/mardwerk/tower-generator/issues/85).

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

The upstream Template uses one path with tiers 0 through 2. The Default Profile instead inherits Atlas's three-path, five-upgrade ordinary progression, with one Tower file per state and separate Upgrade definitions.
It preserves raw BTD6 fields, exact source contracts and the broader Atlas settings, including presentation fields and required supporting data.
The Template, Atlas's BTD6 Profile and Tower Generator's Default Profile remain separate.

## Iterate across the three repositories

Keep the checkouts as siblings and edit each subject in its owning repository.

```text
mardwerk/
  btd6-atlas/
  td-profile/
  tower-generator/
    default-profile/
    scripts/check-default-profile.py
```

Run the smoke check from Tower Generator's root:

```sh
python3 scripts/check-default-profile.py
python3 scripts/check-default-profile.py --local-checker
```

The first command uses the checksum-verified pinned executable, copied locally with the Profile or supplied by the sibling Atlas checkout.
It compares the Default Profile's shared schema copies with Atlas's and scores Dart against both Profiles, requiring complete 100/100 results.
The second command builds the sibling td-profile working tree into ignored `.tools/`, then checks both Profiles against that working tree's shared schemas and runs the same Dart checks.
Local mode requires Go. Neither mode pulls repositories, rewrites settings or changes source captures.

Use `--atlas /path/to/checkout` and `--td-profile /path/to/checkout` for alternate worktrees.
The script resolves its own repository paths, so it can also be invoked from another working directory.

| Change | Where to edit | Local verification |
| --- | --- | --- |
| BTD6 source facts, bindings or capture contracts | Atlas | Atlas's checks and the cross-repository Dart command |
| Generic checks or reusable schemas | td-profile | Its tests, vet and build, then `--local-checker` |
| Default settings, authoring or generated Towers | Tower Generator | The Default Profile command and candidate checks |

Shared schema changes require explicit updates to the consuming Profiles' local copies; the smoke check reports drift instead of copying files automatically.
Default settings can diverge from Atlas settings as product requirements develop. Keep the inherited raw bindings while using Dart for this baseline.
Use local checker builds while iterating; publish a td-profile release when a change is ready, then deliberately update the pinned checker metadata and affected schema copies.
Record the tested checker and Profile revisions in the owning PR. A release is unnecessary for each local edit.

Once a generated candidate exists, check it alongside the reference:

```sh
python3 scripts/check-default-profile.py \
  --game-data game-data/probe --tower Towers/Probe/Probe.json
```

Supply its complete data directory, including declared dependencies. Add `--local-checker` to exercise pending shared changes.
The command reports checker and Profile identities and leaves diagnostics in the terminal.

The first generated Tower should retain 64 ordinary states and 15 upgrades, with Paragon excluded explicitly.
The inherited unmodified Dart baseline includes Paragon and supporting data, so it checks 99 files; this does not expand the first generation's scope.

## Deferred checks

Kyle will review the implementation details later. The observed gaps are tracked upstream:

- [Purchase and Upgrade-definition consistency](https://github.com/mardwerk/td-profile/issues/1).
- [Applied-upgrade consistency](https://github.com/mardwerk/td-profile/issues/2).
- [Profile-configured numerical limits](https://github.com/mardwerk/td-profile/issues/3).

The current compliance score does not establish these relationships or gameplay balance.

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

The [first authoring experiment](AUTHORING.md) creates an editable Usopp draft with readable output and positive and negative checker reports.
Run `python3 scripts/author-tower.py examples/usopp.json --check` to regenerate it locally.
The experiment keeps the current Profile unchanged; it does not establish general generation or design acceptance.

Use the accepted Dart reference and record its source evidence using [BTD6-REFERENCE.md](BTD6-REFERENCE.md).
Prepare a compact authoring view while keeping the full source available, then generate one different Tower using demonstrated mechanics and explicit state templates.
Keep the 64 ordinary states, 15 upgrades, purchase links and relevant mechanics explicit; exclude Paragon from that generation example.
Keep supplied data unchanged during validation and keep generated data out of Git.

The next product checks should demonstrate a valid sample and failures for a missing required state or upgrade.
Record the release identity, Profile revision and checks actually performed. Document the accepted product contract through [#92](https://github.com/mardwerk/tower-generator/issues/92).
The first generation experiment can proceed while the additional checks are deferred. Game integration remains later work; defer SkillOpt until validator evaluation cases are stable.
