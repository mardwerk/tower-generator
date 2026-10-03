# Improve the initial skill with SkillOpt

Start with the [initial Tower Generator skill](../../.agents/skills/tower-generator/SKILL.md) and a few frozen tasks.
Improve its instructions when observed failures justify a change.
The current skill can produce reviewable designs; general generation/export still needs the contract review in [#106](https://github.com/mardwerk/tower-generator/issues/106).

The upstream review used [microsoft/SkillOpt commit fa4ca184573e42ec11472959dd57422381418096](https://github.com/microsoft/SkillOpt/tree/fa4ca184573e42ec11472959dd57422381418096), dated 2026-09-30.
These findings describe that source revision, which includes changes after release v0.2.0.
No package installation or optimization run was performed for this transfer.

## What it does

SkillOpt changes Markdown instructions rather than model weights.
A target model performs scored tasks; an optimizer reflects on the trajectories and proposes bounded additions, deletions or replacements.
The default validation gate retains a candidate when its selection score improves.
It can start with an empty file, but still needs tasks, a runner and an evaluator.
For this product, an explicit seed is easier to review.

The [first experiment guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/first-experiment.md) describes output including `best_skill.md`, resolved configuration, versions, step records and trajectories.
Use the [source installation guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/installation.md) for Python 3.10 or newer and model backend setup.
Model settings are explicit; support for a compatible backend does not establish that `gpt-6.1-sol` has been tested upstream.

## First evaluation scope

Use frozen character evidence and Profile inputs, with network research disabled, to make failures reproducible.
Keep development, selection and final test examples separate.
The selection set is used repeatedly by optimization and is not an untouched final test.

| Task family | Required behavior |
| --- | --- |
| Useful beam evidence plus unsupported teleportation | Keep the supported beam, identify the teleportation gap and invent no mechanic. |
| Ability names without operational evidence | Explain the evidence gap without inventing attacks. |
| A source describing several characters | Attribute capabilities to the correct actor. |
| A technique limited by equipment or story period | Preserve the condition in the design. |
| Persistent plants requested in the Usopp experiment | Report the authoring limit while preserving supported changes. |
| Usopp explosion crosspath revision | Keep contact behavior, blast pierce, firing interval and range consistent. |

These are task families for fixture creation, not scored benchmark results.
Use different concrete examples in each split and review their expected outcomes before training.
Record fabricated claims, wrong actors, lost conditions, false refusals and structural failures separately.
Report provider failures separately from design errors.
Run deterministic authoring/Profile checks where supported; use an explicit human rubric for fidelity and design quality.
Do not replace those judgments with a validator score.

## Integration still needed

The [custom benchmark guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/new-benchmark.md) requires split loading, target rollouts, scoring, persisted conversations, a registered `EnvAdapter` and configuration.
Each result contains `id`, `hard` and `soft`; learning also needs `predictions/<id>/conversation.json`.
The adapter must inject the candidate skill into the target's prompt.
Existing Tower Generator commands do not automatically load this skill.

After that adapter and configuration exist in a separate SkillOpt checkout, the proposed invocation is:

```sh
python scripts/train.py \
  --config configs/tower_generator/default.yaml \
  --out_root outputs/tower_generator_first_run
```

Those paths do not exist yet; this is not a runnable Tower Generator command.
Start with one epoch, a small edit budget, fixed model settings and reviewed selection criteria.
Review the resulting skill against the final test set and the product contract before adoption.
Keep training artifacts outside product runtime files.

[SkillOpt-Sleep](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/sleep/README.md) can later learn from recurring sessions and target a repository skill.
Its default replay is textual, so it does not prove that Tower authoring works in a fresh checkout.
Disable memory evolution for a skill-only pilot.
Input-count limits do not cap spend, and real-backend dry runs can make provider calls; establish an independent provider limit before a paid pilot.
