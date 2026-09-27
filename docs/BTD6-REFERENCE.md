# BTD6 reference

## BTD6 knowledge authority

BTD6 domain facts live in [btd6-atlas](https://github.com/KyleDerZweite/btd6-atlas), validated against a full 56.3 game capture. Treat it as authoritative for tower roster, tiers, costs, crosspath legality and map geometry. Do not add new BTD6 facts to this repository; add them there.

The atlas keeps each capture under `data/<patch>-build-<id>/` with a manifest, and its derived analyses under `patterns/` (towers, maps and progression). Its exported game data and derived tables are licensed CC BY-NC 4.0, so cite them by patch, build and file instead of copying them here. The pinned capture's manifest has an empty `acceptedBy`. The owner accepted capture 56.3 on 2026-09-26 ([#27](https://github.com/mardwerk/unit-generator/pull/27)) as the Default Profile's numerical BTD6 reference only; that does not accept the atlas's incomplete spatial data.

## Pinned capture

Every BTD6 value in this repository comes from capture **56.3, Steam build 24829026**, at atlas revision [`a380413`](https://github.com/KyleDerZweite/btd6-atlas/tree/a380413ed809c98654acec7aa37cba40f80bf5c5/data/56.3-build-24829026). File names below are relative to that capture's `game-data/` folder: tower files are `Towers/<Tower>/<Tower>-<top><middle><bottom>.json` (`<Tower>.json` for 0-0-0) and upgrade prices are the `cost` of `Upgrades/<Upgrade name>.json`. Prices are Medium incremental purchases.

| Consumer | BTD6 values | Files |
| --- | --- | --- |
| `src/cli/internal/mechanics/defaults.go` | The Definition's reference scale: Dart Monkey 0-0-0 price, damage, interval, range and pierce, and its top-path purchase prices | `Towers/DartMonkey/DartMonkey.json`; `Upgrades/Sharp Shots.json`, `Razor Sharp Shots.json`, `Spike-o-pult.json`, `Juggernaut.json`, `Ultra-Juggernaut.json` |
| `src/cli/internal/unit/defaults_text.go` | The default rules document (`default-td-profile-v34`): scale references for group capacity, precision, attack speed with an allied or a self Active Ability, path identity and a pinnacle, single target with blimp control, control, knockback, burst and support, and the roster price bands; the purchase roles cite Dart Monkey's 0-3-0 Triple Shot as the only Dart Monkey purchase with more than one projectile per attack (`ArcEmissionModel` of 3; every top and bottom file uses a `SingleEmissionModel`), and Mermonkey's top path for the progression from a defining third purchase to a pinnacle (the tentacle attacks from 3-0-0, the `SlowInk` slow at 4-0-0 and the `CreateNearbyWaterModel` of radius 32 that only 5-x-x builds have) | `Towers/DartMonkey/DartMonkey-{100..500,010..050,001..005}.json`, `Towers/Mermonkey/Mermonkey{,-100..-500}.json`, `Towers/BoomerangMonkey/BoomerangMonkey{,-010..-050}.json`, `Towers/SniperMonkey/SniperMonkey{,-100..-500}.json`, `Towers/IceMonkey/IceMonkey{,-100..-500}.json`, `Towers/TackShooter/TackShooter{,-010..-050}.json`, `Towers/MonkeyVillage/MonkeyVillage{,-100,-200,-020}.json`, `Towers/SuperMonkey/SuperMonkey-001.json`, the matching `Upgrades/` files, and `patterns/towers.md` and `patterns/progression.md` at the same revision |
| `src/cli/internal/unit/defaults.go` | The Default Definition's Knockback status: its unit, bounds and immunities follow the `KnockbackModel` values above | `Towers/DartMonkey/DartMonkey-400.json`, `Towers/SuperMonkey/SuperMonkey-001.json`, `Upgrades/Knockback.json` |
| `src/cli/internal/unit/defaults.go` | The Hardened enemy property, adapted from the Ceramic tag, and the Blimp property's bonus damage, adapted from the MOAB-class `Moabs` tag, as two separate properties: `Moab.json` is tagged `Moab` and `Moabs`, not `Ceramic`, and no bloon in the capture has both `Ceramic` and `Moabs`. Bonus damage is an additive +N per hit that follows the additive `DamageModifierForTagModel` bonuses, so it only adds: Deadly Precision's +50 to Ceramic, Juggernaut's +3 to Ceramic and +2 to Fortified, Bionic Boomerang's +1 to Moabs and MOAB Press's +4 to Moabs | `Bloons/Ceramic/Ceramic.json`, `Bloons/Moab/Moab.json`, `Towers/SniperMonkey/SniperMonkey-300.json`, `Towers/DartMonkey/DartMonkey-400.json`, `Towers/BoomerangMonkey/BoomerangMonkey-030.json`, `Towers/BoomerangMonkey/BoomerangMonkey-004.json` |
| `data/reference/dart-monkey.request.json` | A Dart Monkey brief: every purchase, price and supported number of the three paths | `Towers/DartMonkey/` and the matching `Upgrades/` files |
| `src/cli/internal/fixture/testdata/` | The scripted test unit, written from that brief | as above |

The model prompts carry no BTD6 values of their own; they point to the rules document's references.

Names are the display names of the capture's `textTable.json`, which can differ from upgrade IDs and file names: Ice Monkey's second top purchase is the upgrade `Metal Freeze` displayed as Cold Snap, and Bionic Boomerang's file is `Upgrades/Bionc Boomerang.json`. When replacing other hand-compiled values, take names from the text table as well; the retired package below also swapped some names.

### How the references were read

Values were read from the tower models with a throwaway script, not copied from prose: the `AttackModel` range, each weapon's `rate`, emission count, projectile `pierce` and `DamageModel.damage`, the `immuneBloonProperties` bit mask (Lead 1, Black 2, White 4, Purple 8, Frozen 16), `DamageModifierForTagModel` bonuses, the `FilterInvisibleModel` Camo filter, child projectiles, and `AbilityModel` cooldowns with their behaviors' lifespans. Derived comparisons in the rules document assume a target always in range and ignore enemy-class bonuses unless stated:

- **Damage per second** is damage × projectiles ÷ interval. For Sharp Shooter and Crossbow Master a critical shot counts as its critical damage instead of the ordinary hit.
- **Damage capacity per attack** is damage × pierce, plus child projectiles for Ultra-Juggernaut (2 × 6 balls × 2 damage × 50 pierce).
- **Active uptime** is duration ÷ cooldown. Perma Charge's 15-second window is its `DamageUpModel` lifespan of 900 frames.

The atlas itself says per-tier DPS modeling is future work and that price ratios do not establish output multipliers. These comparisons are analytical readings for scale, not balance targets.

### Cross-checks

Medium prices and the local upgrade rule also agree with the reviewed Mardwerk knowledge summary of a September 24 BTD6 price snapshot, which checked every price against this capture ([`knowledge/btd6-snapshot.md`](https://github.com/mardwerk/project/blob/f8b2a47b649f507655ae6a0d4c10fa8850379623/knowledge/btd6-snapshot.md), private and optional). That summary does not validate combat stats or crosspath effects, so none are taken from it. The old default's Perma Charge text (8 extra damage for 15 s) disagrees with this capture's 10.

## Deliberate departures from BTD6

The default Profile follows BTD6 closely: the same 3 × 5 purchase structure, crosspath rule, Medium incremental prices, damage types that cannot hurt listed enemy properties, Camo detection and one placed fifth purchase per path. It departs on purpose where the typed Definition or the design rules differ:

| Area | BTD6 | Default Profile |
| --- | --- | --- |
| Match | Cash, lives and a bloon layer tree | Gold and a shared Health pool of 150; the basic enemy has one layer and no layer tree or leak simulation exists |
| Active Abilities | On the middle fourth purchase of 25 of 26 towers, and on one top and one bottom fourth purchase in the roster | Only the middle path, first at `x-4-x`; one Active Ability, a boost of the purchased attack, which `x-5-x` may modify |
| Attacks | Several attacks, sub-towers and separate ability attacks per tower | One automatic attack with bounded follow-ups and distinct-target volleys |
| Allied effects | Buffs, shared detection and transformations of nearby towers | Unsupported; only the Unit's own attack changes |
| Damage modifiers | Bonuses against Ceramic, Fortified, MOAB-class and other tags, critical-hit counters, rebounds, shot arcs | A bonus against Ceramic becomes typed [bonus damage](MECHANICS.md#bonus-damage) against Hardened, and a bonus against MOAB-class (`Moabs`) becomes bonus damage against Blimp, each an additive +N per hit. They stay separate properties: a Blimp is not Hardened. BTD6 bosses carry `Moabs` too; the Default keeps Boss a separate property, so a Boss takes a Blimp bonus only when the Consumer also marks it Blimp. Fortified, Lead and other tag bonuses, critical-hit counters, rebounds and shot arcs stay unsupported: listed as unsupported mechanics, never folded into ordinary damage |
| Knockback | A `KnockbackModel` with a lifespan and separate multipliers for light, heavy and MOAB-class bloons (Juggernaut 0.15 s at 4, 2 and 0; Super Monkey's Knockback 0.5 s at 1.25, 0.6 and 0.3) | One Knockback status: its magnitude is the light-enemy multiplier, read as a speed multiple, and its duration the lifespan; Blimp and Boss enemies ignore it. The atlas gives the values, not the movement formula |
| Purchase size | An upgrade changes as many properties as it needs; Crossbow Master changes damage, interval, pierce, range and damage type | Each purchase has a change budget set by the Definition profile (`maxChangesPerTier`): the Default Profile allows five from the third purchase and three through the second, enough for Crossbow Master |
| Vocabulary | Many damage types and bloon properties | Four damage types (Sharp, Normal, Explosive, Energy), eight enemy properties, Hardened among them, with Hardened and Blimp accepting bonus damage, and the statuses Slow, Burn, Stun and Knockback |

The atlas leaves one bonus damage question open. No bloon in capture 56.3 has both the `Ceramic` and the `Moabs` tag, so the data cannot show whether BTD6 adds a Ceramic bonus and a MOAB-class bonus to the same hit. The [runtime rule](MECHANICS.md#bonus-damage) that an enemy with both Hardened and Blimp takes both bonuses is this repository's choice (SOL-42-02 on #42), not an atlas fact, and the Default does not make a Blimp Hardened.

A generated unit records its own departures from its source as proposed-mechanic findings on the purchases that need them, or unsupported-mechanic findings for behavior no purchase needs.

## Retired material

The September 20 research was removed on September 26, 2026: the compiled 26-tower package (`btd6_towers.json`, SHA-256 `a2a5e2bb4591…`), role categories, the design baseline, six detailed tower examples, the pattern analysis, the historical fixed-recipe notes, the Dart and Boomerang page snapshots and their provenance file. They remain in history at [research/btd6 in f19af56](https://github.com/mardwerk/unit-generator/tree/f19af56/research/btd6). No active consumer cites them since default Profile version 12; saved artifacts and Profiles made earlier keep the rules text they were prepared with.
