# Default Profile

This is Tower Generator's initial Default Profile, copied from Atlas's verified BTD6 Profile.
It retains BTD6 field names, exact source model contracts, three-path ordinary progression and the broader Atlas settings.

The Profile identity is `tower-generator-default`, revision `0.1.0`, with checker interface 5.
All settings, source contracts and shared schemas initially match the imported Atlas Profile; only the manifest identity and revision differ.
Generation policy and additional numerical limits remain to be developed.

## Source and licenses

The imported Profile comes from [Atlas commit de82968](https://github.com/KyleDerZweite/btd6-atlas/tree/de829684232157967fd66f0e999a45df3a669c63/profile).
Its model contracts were derived from BTD6 capture 56.3, Steam build 24829026.
The shared schemas and checker originate in [td-profile v1.0.2](https://github.com/mardwerk/td-profile/releases/tag/v1.0.2), checker 5.0.1.

Source code and reusable schemas use the preserved [MIT license](licenses/btd6-atlas.LICENSE).
Capture-derived contracts and tables retain Atlas's [source notice](licenses/btd6-atlas.NOTICE) and CC BY-NC 4.0 terms.
Upstream checker notices and dependency licenses remain in `licenses/`.

## Check and iterate

Run `python3 setup.py` from the repository root to build the latest released native checker, author Usopp and validate the example.
See [Profile integration](../../docs/PROFILE.md) for commands, local validation and the ownership of changes.

The executable lives at `default/profile-validator`, outside this Profile directory.
The installed software identity is recorded in `default/validator-release.json`.
The Profile and generated game-data stay separate inside the default workspace.

Dart is the verified reference, with 64 ordinary states and 15 ordinary upgrades.
Its unchanged Atlas scoring scope also includes Paragon and supporting data, for 99 files; a generation example can exclude Paragon explicitly.
The score measures declared data compliance. It does not establish balance, behavior simulation or character fidelity.
