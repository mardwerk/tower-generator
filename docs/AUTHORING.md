# First Tower authoring experiment

The [Usopp design](../examples/usopp.json) is an editable draft with Impact, Volley and Marksman upgrade paths.
Kyle selected this experiment under [#85](https://github.com/mardwerk/tower-generator/issues/85), using the existing Default Profile and checker.
During this iteration, work directly on `main` and track necessary Profile or dependency changes as issues.
Atlas and td-profile remain unchanged. The draft still needs design review.

Run from Tower Generator with Python 3.10 or newer and the sibling Atlas checkout available:

```sh
python3 scripts/author-tower.py examples/usopp.json --check
```

Use `--atlas /path/to/btd6-atlas` for another checkout.
Without `--check`, the command only generates the local files and clears any previous check report.
Generation uses Python's standard library and the pinned local source; it makes no network requests.
Validation uses the checksum-verified td-profile v1.0.2 executable through the existing [Profile integration](PROFILE.md).

The generated `game-data/usopp/` directory is ignored by Git:

| File | Contents |
| --- | --- |
| `TOWER.md` | Description, 15 upgrades and the stats of all 64 states |
| `REFERENCE.md` | Compact Dart reference showing each state's first weapon |
| `data/` | 64 Tower states, 15 Upgrade definitions and two supporting tables |
| `source.json` | Atlas source commit, capture, build, source files and byte hashes |
| `checks.json` | Full checker reports for the candidate and two missing-record cases |

Edit the base stats, upgrade names, costs or changes in `examples/usopp.json`, then run the command again.
Files for that Tower are regenerated. New display names use stable Upgrade identifiers so renamed upgrades leave no old records behind.
The original Atlas source stays available and unchanged.

## Mapping and limits

The experiment verifies consumed source files byte-for-byte against Atlas commit `de829684232157967fd66f0e999a45df3a669c63`.
It also verifies the supporting tables against the capture manifest's hashes.
The source is BTD6 56.3, Steam build 24829026. See [reference attribution](BTD6-REFERENCE.md).

The 64 ordinary Dart states supply the legal state list and purchase transitions.
Each authored state starts with Dart's base model, then receives the designed stats and explicit purchase links.
Dart's `0-3-0` supplies the demonstrated spread emission; its count is authored.
The `Sharp Shots` Upgrade model supplies the record shape, with authored identifiers, costs, path indices and zero-based tiers.

Upgrade changes accumulate in top-middle-bottom order.
Damage, pierce, range and projectile counts add; interval multipliers multiply; camo detection stays enabled once gained.
State cost is the base purchase plus every applied upgrade.
The mapping updates captured stat aliases, attack range and all invisible-target filters together.
Straight-projectile lifespan is extended when needed to reach the authored range.

Paragon and Monkey Knowledge are excluded.
Inherited Dart art, sounds, animation and immunity settings remain placeholders; there are no plant effects, active abilities or game runtime integration.
The authoring input supports only this small projectile design. It is not the general generator interface or an accepted Profile replacement.
An authoring format without captured presentation requirements is deferred in [#100](https://github.com/mardwerk/tower-generator/issues/100).

`--check` requires a complete candidate score of 100 and a `missing_reference` diagnostic when either a required state or Upgrade definition is removed from a disposable copy.
The candidate files remain unchanged by validation.
The [deferred shared checks](PROFILE.md#deferred-checks) still apply. The score measures declared compliance, not design quality, balance or gameplay.

Run the focused authoring checks with:

```sh
python3 scripts/test-author-tower.py
```

These checks cover crosspath stats, camo and projectile reach, legal states, purchase-to-upgrade mapping, applied upgrades, revisions and invalid design inputs.
They require the same Atlas checkout and do not modify it.
