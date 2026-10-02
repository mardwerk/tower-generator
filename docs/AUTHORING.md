# First Tower authoring experiment

The [Usopp design](../default/usopp.json) is an editable draft with Exploding stars, Pop Greens and Sogeking Styles.
Kyle selected this experiment under [#85](https://github.com/mardwerk/tower-generator/issues/85), using the existing Default Profile and checker.
During this iteration, work directly on `main` and track necessary Profile or dependency changes as issues.
Atlas and td-profile remain unchanged. The draft still needs design review. [Canon and adaptations](USOPP.md) records the sources and the three attack identities.

Run from Tower Generator with Python 3.10 or newer and the sibling Atlas checkout available:

```sh
python3 setup.py
```

Use `--atlas /path/to/btd6-atlas` for another checkout.
Setup builds or reuses the latest released native validator, generates the example and runs positive and negative checks.
[Profile integration](PROFILE.md) records the build and validation behavior.
Run `pnpm dev` to inspect the same local files in the web view, or `pnpm preview` after `pnpm build`.
Use the path selectors or upgrade cards to choose a legal state, and inspect all states or the selected Tower JSON.
Edit design and Build and check provide the same focused authoring flow in the Lab.
The Lab validates a temporary candidate before changing saved files. Editing the file directly and rerunning setup remains supported.
[Character research](RESEARCH.md) explains collecting and reusing saved evidence.

Generated outputs under `default/` are ignored by Git; the design and Profile are tracked:

| File | Contents |
| --- | --- |
| `TOWER.md` | Description, 15 upgrades and the stats of all 64 states |
| `REFERENCE.md` | Compact Dart reference showing each state's first weapon |
| `game-data/` | 64 Tower states, 15 Upgrade definitions and two supporting tables |
| `source.json` | Atlas provenance and the character-source snapshot hash |
| `character-source.json` | Exact reusable research used to author this Tower |
| `checks.json` | Full checker reports for the candidate and two missing-record cases |

Edit the base stats, upgrade names, costs or changes in `default/usopp.json`, then run the command again.
Files for that Tower are regenerated. New display names use stable Upgrade identifiers so renamed upgrades leave no old records behind.
The original Atlas source stays available and unchanged.

## Mapping and limits

The experiment verifies consumed source files byte-for-byte against Atlas commit `de829684232157967fd66f0e999a45df3a669c63`.
It also verifies the supporting tables against the capture manifest's hashes.
The source is BTD6 56.3, Steam build 24829026. See [reference attribution](BTD6-REFERENCE.md).

The 64 ordinary Dart states supply the legal state list and purchase transitions.
Each authored state starts with Dart's base model, then receives the selected projectile, designed stats and explicit purchase links.
The third purchases select Bomb Shooter's contact explosion, Dart's `0-3-0` spread emission, or Dart's `0-0-3` Crossbow projectile.
The Bomb Shooter base record is verified against the same pinned source commit. Its blast uses authored radius, damage and hit capacity.
Its carrier expires at first contact; purchased pierce applies to the explosion. Pop Green volley counts are authored adaptations.
The `Sharp Shots` Upgrade model supplies the record shape, with authored identifiers, costs, path indices and zero-based tiers.

Upgrade changes accumulate in top-middle-bottom order.
Damage, pierce, range, blast radius and projectile counts add; interval multipliers multiply; camo detection stays enabled once gained.
The selected attack retains every early purchase. Pierce applies to the blast or each leaf/sniper projectile; range extends targeting and carrier travel, and reload affects firing.
Range upgrades do not enlarge a blast.
State cost is the base purchase plus every applied upgrade.
The mapping updates captured stat aliases, attack range and all invisible-target filters together.
Straight-projectile lifespan is extended when needed to reach the authored range.

Paragon and Monkey Knowledge are excluded.
Inherited BTD6 art, sounds and animation remain placeholders. Each third purchase records its intended Usopp appearance.
Dart projectiles retain Dart immunities; contact explosions retain Bomb Shooter immunities. No persistent plants, active abilities or game runtime integration are included.
The authoring input supports only this small projectile design. It is not the general generator interface or an accepted Profile replacement.
An authoring format without captured presentation requirements is deferred in [#100](https://github.com/mardwerk/tower-generator/issues/100).

Setup requires a complete candidate score of 100 and a `missing_reference` diagnostic when either a required state or Upgrade definition is removed from a disposable copy.
The candidate files remain unchanged by validation.
The [deferred shared checks](PROFILE.md#checks-and-deferred-work) still apply. The score measures declared compliance, not design quality, balance or gameplay.

Run the focused authoring checks with:

```sh
pnpm test
```

These checks cover distinct third-purchase projectiles, explosion crosspaths, camo and projectile reach, legal states, purchase-to-upgrade mapping, applied upgrades, revisions and invalid design inputs.
They require the same Atlas checkout and do not modify it.
