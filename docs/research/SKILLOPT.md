<a id="skillopt-evaluation-proposal"></a>

# SkillOpt optimization plan

Kyle selected SkillOpt with his review for the initial Agent Skill on 2026-10-11 under the [product delivery order](../PRODUCT.md#established-direction).
The [Tower design skill](../../.agents/skills/tower-generator/SKILL.md) is the existing seed.
Its detailed request contract, quality criteria, evaluator and adoption evidence remain under review through `grill-with-docs`.
Evaluate the selected character-folder output and any supported question sessions, rather than imposing the later generator's export contract.
Comparisons before rubric development are selected; the proposed evaluation details do not select fixtures, scoring, models or spending limits.
Skill optimization is separate from the later API backend and does not select its architecture.
There is no installed adapter, scored Tower Generator benchmark or optimization run from this transfer.

The upstream review used [microsoft/SkillOpt commit fa4ca184573e42ec11472959dd57422381418096](https://github.com/microsoft/SkillOpt/tree/fa4ca184573e42ec11472959dd57422381418096), dated 2026-09-30.
The observations describe that revision, including changes after release v0.2.0.

<a id="evaluation-approach"></a>

## Proposed evaluation approach

SkillOpt changes Markdown instructions rather than model weights.
A target model performs scored tasks; an optimizer proposes bounded edits from their trajectories.
Its default selection gate retains a candidate when its selection score improves; that gate does not replace Kyle's review for adopting a skill version.
It still needs tasks, a runner and an evaluator; an explicit initial skill is easier to review.

For repeatable design checks, use frozen evidence and any supplied game constraints, with network research disabled.
These checks cover design from supplied material; separate research and conversation cases are needed to assess the complete skill.
Keep development, selection and final test examples separate.
Repeated optimization uses the selection set, so it is not an untouched final test.

| Task family | Required behavior |
| --- | --- |
| Supported beam evidence and unsupported teleportation | Preserve the supported beam and identify the teleportation gap without presenting invented behavior as canon. |
| Ability names without operational evidence | Explain the evidence gap and distinguish any proposed gameplay behavior from character facts. |
| Sources describing several characters | Attribute capabilities to the correct actor. |
| Equipment or story-period limitations | Preserve those conditions in the design. |
| Requested export without an implemented contract | Produce supported design work and identify the exact interface gap. |
| Character-defining behavior conflicts with target-game patterns | Explain the adaptation and ask when the user's intent is unclear. |

These families describe potential fixtures, not benchmark results.
Review expected outcomes and use different examples in each split.
Record fabricated claims, wrong actors, lost conditions, false refusals and structural failures separately.
Separate provider failures from design errors.
Use deterministic checks only where an implemented contract supports them, and an explicit human rubric for fidelity and design quality.
A validator score does not establish those judgments.

<a id="proposed-first-quality-comparison"></a>

## Selected quality direction

Kyle selected comparisons first, followed by a rubric, on 2026-10-11.
The following pilot details are proposals to refine through those comparisons.
Compare the existing and a revised skill on the same tasks, hiding their version names during Kyle's review.
Record concrete errors separately from his preference and its reasons.
For known characters, assess Kyle's selected goals of canon fidelity and closeness to the requested game's patterns and mechanics separately.
Concrete assessment examples remain to be derived from comparisons; for original ideas, there is no external character canon to score.
Other possible criteria include useful gameplay decisions, clear explanations and adaptation to the request.
Assess image quality separately so a stronger illustration does not conceal a weaker design.
Derive short assessment examples from these comparisons before selecting a scoring rubric.

Keep some different characters and request types unseen until a skill version is proposed for adoption.
Compare any automated evaluator with Kyle's judgments before relying on its scores for optimization.
Text and image assessments do not establish gameplay balance or enjoyment; those require implementation and playtesting.
Concrete pilot design, criteria, tasks, scoring scale, evaluator and paid runs remain open.

## Integration still needed

The [custom benchmark guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/new-benchmark.md) requires split loading, target rollouts, scoring, persisted conversations, an `EnvAdapter` and configuration.
The adapter must inject the candidate skill into the target prompt.
Product requests do not automatically load a repository skill.

The [installation guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/installation.md) describes upstream setup.
The [first experiment guide](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/guide/first-experiment.md) describes retained configuration, versions, trajectories and candidate skills.
Any SkillOpt optimization pilot belongs in a separate SkillOpt checkout with reviewed fixtures and fixed model settings.
No runnable Tower Generator training command is specified here.

Review a candidate against the final test set and product contract before adoption.
Keep training artifacts outside product runtime files.
[SkillOpt-Sleep](https://github.com/microsoft/SkillOpt/blob/fa4ca184573e42ec11472959dd57422381418096/docs/sleep/README.md) is a later option for recurring sessions; textual replay does not prove executable behavior.
For a skill-only pilot, disable memory evolution and establish a provider spending limit before paid calls.
