# Profile integration

Tower Generator owns [default/profile/](../default/profile/), its authoring and product checks.
[td-profile](https://github.com/mardwerk/td-profile) owns the generic Profile Validator, reusable schemas and copyable Template.
[Atlas](https://github.com/KyleDerZweite/btd6-atlas) owns BTD6 source captures and its separate BTD6 Profile.

The Default Profile has identity `tower-generator-default`, revision `0.1.0` and checker interface 5.
It initially copies Atlas's verified raw BTD6 settings and model contracts.
It governs separate data in `default/game-data/`. Every Profile dependency resolves inside `default/profile/`.
The broader Atlas requirements, including presentation fields and supporting tables, still apply.

## Prepare the native validator

Run from this repository with Python 3.10 or newer, Go, Git and GitHub CLI installed:

```sh
python3 setup.py
```

The script resolves the latest published td-profile release through GitHub CLI.
It detects the current OS and architecture, clones the released tag into a temporary directory and builds with `CGO_ENABLED=0` and `go build -trimpath`.
It checks shared-schema compatibility and validates an empty data directory before installing `default/profile-validator`, or `profile-validator.exe` on Windows.
The temporary source checkout is removed afterward. A matching installed build is reused only when its release, native target and executable hash match.

`default/validator-release.json` records the release, source commit, checker identity, native target and executable checksum.
The executable is a local build of the released source, so its bytes may differ from an upstream release asset.
Upstream notices are installed in `default/licenses/`. Imported Profile notices remain in `default/profile/licenses/`.

Setup then generates and validates the [Usopp experiment](AUTHORING.md).
The sibling Atlas checkout is an authoring source; validation after setup uses only the local default workspace and runs offline.
No code or files in Atlas or td-profile are changed.

Check the local example directly on Linux or macOS:

```sh
./default/profile-validator score-tower --profile default/profile \
  --game-data default/game-data --tower Towers/Usopp/Usopp.json --format text
```

On Windows, use `default/profile-validator.exe` with the same arguments.
The command reports Profile and checker identity and the checks actually performed.

## Checks and deferred work

Setup requires a complete candidate score of 100.
It removes a required ordinary state and Upgrade definition in disposable copies, requiring each to fail with a matching `missing_reference` diagnostic.
It leaves the candidate files unchanged during validation.
The first experiment has 64 ordinary states and 15 upgrades. Paragon and Monkey Knowledge are excluded.

The current score measures declared data compliance, not gameplay balance or runtime behavior.
These shared capabilities remain deferred:

- [Purchase and Upgrade-definition consistency](https://github.com/mardwerk/td-profile/issues/1).
- [Applied-upgrade consistency](https://github.com/mardwerk/td-profile/issues/2).
- [Profile-configured numerical limits](https://github.com/mardwerk/td-profile/issues/3).

Authored purchase and applied-upgrade mappings also have focused checks in Tower Generator.
[Presentation requirements](https://github.com/mardwerk/tower-generator/issues/100) remain a Default Profile follow-up.
Propose shared checker or schema updates upstream and record necessary Profile changes as issues while this authoring experiment continues.
Keep game-specific requirements in the owning Profile.

The native setup and candidate were verified on Linux amd64 with td-profile v1.0.2, checker 5.0.1, source commit `48fc5daa723d13d89d33bce4a863ba5bf38a1b39`.
This does not establish native execution on other platforms.
