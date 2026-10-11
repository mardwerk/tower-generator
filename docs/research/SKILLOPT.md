<a id="skillopt-evaluation-proposal"></a>

# SkillOpt optimization plan

Kyle selected SkillOpt with his review for the initial Agent Skill on 2026-10-11 under the [product delivery order](../PRODUCT.md#established-direction).
The [Tower design skill](../../.agents/skills/tower-generator/SKILL.md) is the existing seed.
Its supported requests, quality criteria, evaluation contract and review process remain to be defined through `grill-with-docs`.
The evaluation design below remains a proposal; selecting SkillOpt does not select its fixtures, scoring, models or spending limits.
Skill optimization is separate from the later API backend and does not select its architecture.
There is no installed adapter, scored Tower Generator benchmark or optimization run from this transfer.

The upstream review used [microsoft/SkillOpt commit fa4ca184573e42ec11472959dd57422381418096](https://github.com/microsoft/SkillOpt/tree/fa4ca184573e42ec11472959dd57422381418096), dated 2026-09-30.
The observations describe that revision, including changes after release v0.2.0.

<a id="evaluation-approach"></a>

## Proposed evaluation approach

SkillOpt changes Markdown instructions rather than model weights.
A target model performs scored tasks; an optimizer proposes bounded edits from their trajectories.
Its default selection gate retains a candidate when its selection score improves.
It still needs tasks, a runner and an evaluator; an explicit initial skill is easier to review.

Use frozen evidence and Profile inputs, with network research disabled, to make failures reproducible.
Keep development, selection and final test examples separate.
Repeated optimization uses the selection set, so it is not an untouched final test.

| Task family | Required behavior |
| --- | --- |
| Supported beam evidence and unsupported teleportation | Preserve the beam and identify the teleportation gap without inventing mechanics. |
| Ability names without operational evidence | Explain the evidence gap without inventing attacks. |
| Sources describing several characters | Attribute capabilities to the correct actor. |
| Equipment or story-period limitations | Preserve those conditions in the design. |
| Requested export without an implemented contract | Produce supported design work and identify the exact interface gap. |
| Historical Usopp crosspath case | Use a frozen historical fixture and state its authoring limits. |

These families describe potential fixtures, not benchmark results.
Review expected outcomes and use different examples in each split.
Record fabricated claims, wrong actors, lost conditions, false refusals and structural failures separately.
Separate provider failures from design errors.
Use deterministic checks only where an implemented contract supports them, and an explicit human rubric for fidelity and design quality.
A validator score does not establish those judgments.

## Integration still needed

The [custom benchmark guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/new-benchmark.md) requires split loading, target rollouts, scoring, persisted conversations, an `EnvAdapter` and configuration.
The adapter must inject the candidate skill into the target prompt.
Product requests do not automatically load a repository skill.

The [installation guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/installation.md) describes upstream setup.
The [first experiment guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/first-experiment.md) describes retained configuration, versions, trajectories and candidate skills.
Any pilot belongs in a separate SkillOpt checkout with reviewed fixtures and fixed model settings.
No runnable Tower Generator training command is specified here.

Review a candidate against the final test set and product contract before adoption.
Keep training artifacts outside product runtime files.
[SkillOpt-Sleep](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/sleep/README.md) is a later option for recurring sessions; textual replay does not prove executable behavior.
For a skill-only pilot, disable memory evolution and establish a provider spending limit before paid calls.
