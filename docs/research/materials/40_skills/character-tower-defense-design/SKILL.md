---
name: character-tower-defense-design
description: Design character-based tower-defense units with distinct upgrade paths, complete crosspath behavior, and mechanic-focused descriptions. Edit existing concepts without losing their effects.
---

# Character-based tower-defense unit design

Design what the unit does. Keep design reasoning, balance review, and usage advice out of the unit description.

## Profile boundary

This skill provides shared character-design guidance and the proposed output shape for Unit Generator's public, BTD6-inspired Default Profile.
The rules under "Default Profile conventions" are Default Profile assumptions, not universal game rules.

MangaMayhem has a separate Definition and Profile.
Apply its selected rules for Ink, transformations, placed-unit status effects, and any different ability or output contract.
Do not carry those rules into the public default, or impose a conflicting default rule on MangaMayhem.
When a request does not identify a Profile, use the public Default Profile and name that choice outside the generated unit description.

## Respect the task

A prose edit is not a redesign.
Preserve supplied mechanics and organization unless the user requests changes.
Do not delete attacks, triggers, secondary effects, restrictions, interactions, or duration relationships to shorten the text.

During a redesign, change mechanics deliberately.
When a change summary is requested, distinguish changed mechanics from rewritten prose.
Do not silently resolve contradictions in supplied material.

## Default Profile conventions

Use these defaults internally. Do not print them as a rules section.

- A unit has top, middle, and bottom paths, each with five upgrades.
  A build can purchase two paths, but only one can advance beyond its second upgrade.
- Only the middle path grants player-activated abilities, beginning at `x-4-x`.
  `x-5-x` develops or replaces that ability.
  Top and bottom paths use automatic attacks, passive effects, and conditional triggers.
  A targeting control is not an activation button.
- Purchased effects persist unless an upgrade explicitly changes or replaces them.
  Unmentioned properties remain unchanged.
  Unpurchased features are absent by default.
- Enemy-targeted attacks use the player's selected targeting and normal movement prediction.
  Do not repeat these defaults in attack descriptions.
- Omit prices, damage values, attack intervals, percentages, cooldown lengths, and balance ratings unless requested.
  Use qualitative values such as higher damage, more pierce, faster attacks, a wider impact, or a long cooldown.
  Structural counts are allowed when they define behavior, such as one return pass.

Do not add capability reports, implementation constraints, rules recaps, proposal disclaimers, or review verdicts to the output.

## Upgrade notation

Always write upgrade and build references in top-middle-bottom order.

- `0-0-0` is the base unit.
- Top-path upgrade headings are `1-x-x` through `5-x-x`.
- Middle-path upgrade headings are `x-1-x` through `x-5-x`.
- Bottom-path upgrade headings are `x-x-1` through `x-x-5`.
- Concrete builds use digits, such as `1-2-0` or `4-0-2`.
- `3+-x-x`, `x-3+-x`, and `x-x-3+` are crosspath group headings covering main-path upgrades three through five.

An `x` leaves that path unspecified; it does not grant an upgrade or permit an illegal build.
Use concrete digits when describing a particular combination.
Always include both hyphens: `1-0-2`, not `1-02`.

Use these references instead of "T3," "tier 4," or "Path 1 tier 5" in the generated unit.
Each upgrade description states what that purchase changes.
Crosspath entries resolve its behavior with other purchased upgrades.

## Start with the character

Use the character's actual attacks, weapons, transformations, sensory powers, and physical interactions.
Use supplied or verified references for source facts and honor source-period limits.
Do not present invented tower-defense interactions as additional canon facts.

Do not turn a name into an unrelated power.
Hardening a fist does not inherently amplify every ally's damage.
Sensing enemies does not inherently slow them.

Do not require every known move to appear.
Do not make every path climb the same sequence of forms.
The strongest canonical form does not have to produce a tower that does everything best.

Privately consider different combinations of path concepts.
Reject combinations whose finished paths differ only in which statistic is largest.
Choose different attack behaviors, targeting interactions, terrain interactions, or support functions.
Do not print this exploration.

## Progression without a mandatory formula

The first two purchases on each path should be readable upgrades and useful crosspaths.
Plain stat increases are welcome.
Small, bounded utility mechanics are also acceptable.

The third purchase must establish a reason to commit.
It can introduce a behavior, substantially develop an existing one, or establish a different role.
It does not have to introduce a transformation, resource, mark, combo counter, or minigame.

The fourth and fifth purchases can develop the existing mechanic, add a separate attack, change delivery, introduce support, or add another system.
They are not restricted to making the third purchase larger or stronger.
Earlier attacks can remain alongside new ones.

Possible progressions include piercing projectiles becoming bouncing and splitting attacks; sacrifices determining additional attacks and support; a placed support effect leading into ally-targeted boosts; placement conditions developing alongside a weapon change; or freezing developing into fragments and damage vulnerability.

These are examples, not a checklist.
Use different progression patterns across the unit.
Do not make every path follow "new attack, more targets, bigger version."

Do not force a new mechanic into every purchase either.
A stronger version can be the right upgrade when the existing behavior is worth developing.
More rules do not automatically create a better design.

Keep judgments about role, placement value, purchase timing, and path distinctness in the private review.
The descriptions report the resulting mechanics, not those judgments.

## Describe changes, not unchanged defaults

Each upgrade must include all behavior it adds, changes, or removes.
It should not repeat the unit's entire existing behavior.

Use enough detail to identify the action, trigger, affected targets, delivery, secondary effects, and relevant limits.
Do not print these as a form with repeated labels.
A simple upgrade may need one sentence.
A complicated upgrade needs enough sentences to preserve its mechanics.

State actual conditions and timing. "The transformation has a long cooldown" is useful. "The long cooldown makes players choose the right moment" is usage advice and must be removed.

Do not explain basic terms twice. "Increases fist width and pierce" does not need "so the fists can hit more enemies."

Do not list absent mechanics.
Omit "does not splash yet," "starts without stealth detection," and "nearby enemies are not slowed" when nothing in the description grants those effects.

Do not restate unchanged attacks to contrast them with a new one.
When a new projectile deals return-pass damage, describe that projectile's return pass.
Do not add that every other projectile still lacks return damage.

A negative clarification belongs only when it records an actual removal, a relevant exception to a stated effect, or resolves an ambiguity created by the described mechanic.
Prefer a precise positive scope when possible: "The pierce bonus applies to the impact and its fragments."

Do not describe how to place the unit, which waves to buy it against, why an effect is useful, or what decision it forces.
Retain the factual part of such sentences and remove the advice.

## Additions, replacements, and simultaneous attacks

Privately track which attacks and effects currently exist and what each upgrade does to them.

Use "adds" for an additional attack, "changes" for a modified attack, and "replaces" when an old attack actually stops existing.
Do not call an attack an addition while silently making it disable an earlier system.

An added attack coexists with existing attacks by default.
Do not invent a shared cooldown, shared pierce, animation lock, or recovery period that shuts an earlier attack down.
A shared trigger does not automatically mean shared capacity or mutually exclusive attacks.

For example, adding a heavy-enemy grab alongside a crowd-launching patch must not quietly turn the patch into a single-target attack or suspend its launches.
A limitation on how many enemies the grab holds belongs to the grab, not automatically to the patch.

When an upgrade intentionally replaces or links systems, describe the change where it is introduced.
Check that the replacement retains the useful purchased behavior or deliberately improves on it.
Explicitly describing a regression does not make it a good upgrade.

Do not add damage reductions, slower attacks, or interruptions merely to compensate for a new mechanic.
Do not force every upgrade to have a drawback.

A transformation description must establish which attacks it changes and what happens when it ends.
If `x-5-x` changes both normal attacks and the ability, describe both.
Mention carried-over effects only when their transfer would otherwise be ambiguous.

## Targeting and prediction

Preserve the player's targeting selection.
Give an added attack its own targeting description when that targeting differs.

Assume enemy-targeted attacks account for movement during windup and travel.
A windup does not freeze aim at the enemy's previous position.
Do not invent misses simply because the enemy keeps moving.

Do not make slow attacks deliberately inaccurate as a balancing drawback.
Do not add claims about fast enemies outrunning prediction unless that behavior is specifically part of the requested design.

Distinguish an enemy-targeted attack from an explicitly ground-targeted effect.
A placed trap can occupy a chosen location, but that does not make other attacks stop predicting movement.
Do not introduce fixed-point aiming as an unrequested limitation.

Keep the prediction assumption internal unless an upgrade changes it.
Describe unusual targeting behavior, not the normal targeting system or instructions for aiming around an invented defect.

## Pierce

Treat damage, projectile count, width, area size, and pierce as separate properties.
Pierce is an attack's hit capacity.
A wide attack can have low pierce, and a narrow attack can have high pierce.

Establish the base attack's pierce qualitatively.
Mention later pierce changes without repeatedly explaining what pierce means.

For new chains, bounces, fragments, explosions, returning attacks, or enemy collisions, decide how their hit capacities work.
Explain shared or separate pierce where the new behavior would otherwise be ambiguous.
Do not silently turn an area attack into unlimited targets.

A bounce or retarget does not automatically replenish pierce.
If an upgrade creates a refresh, say so.
Distinguish finding another enemy from hitting the same enemy again.

When a return pass or secondary impact can hit an enemy again, explain that repeat-hit behavior.
Do not recap unrelated attacks that remain unchanged.

Give control effects an application capacity where relevant.
Keep the capacities of independent effects separate.
Holding one enemy does not imply that a separate crowd attack now has one pierce.
Range alone does not bound the number of enemies affected by an aura.

## Crosspath coverage

Describe the actual combined behavior, including partial crosspaths.
Do not collapse a first purchase and a completed two-upgrade crosspath into one vague "speed" or "utility" option.

Every early combination below needs its own identifiable entry:

`1-1-0`, `1-2-0`, `1-0-1`, `1-0-2`
`2-1-0`, `2-2-0`, `2-0-1`, `2-0-2`
`0-1-1`, `0-1-2`, `0-2-1`, `0-2-2`

Describe each once.
A concrete build is the same build regardless of which path the player purchased first.
Do not invent purchase-order effects.

Cover every advanced combination under these groups:

### `3+-x-x`

`3-1-0`, `3-2-0`, `3-0-1`, `3-0-2`
`4-1-0`, `4-2-0`, `4-0-1`, `4-0-2`
`5-1-0`, `5-2-0`, `5-0-1`, `5-0-2`

### `x-3+-x`

`1-3-0`, `2-3-0`, `0-3-1`, `0-3-2`
`1-4-0`, `2-4-0`, `0-4-1`, `0-4-2`
`1-5-0`, `2-5-0`, `0-5-1`, `0-5-2`

### `x-x-3+`

`1-0-3`, `2-0-3`, `0-1-3`, `0-2-3`
`1-0-4`, `2-0-4`, `0-1-4`, `0-2-4`
`1-0-5`, `2-0-5`, `0-1-5`, `0-2-5`

Advanced builds can share a description only when the crosspath applies the same way.
List every covered build explicitly, such as `3-1-0 / 4-1-0 / 5-1-0`.
Split out later differences, including new attacks, transformations, secondary effects, and ability changes.

Describe the effect of the crosspath on the existing main-path description.
Do not repeat the entire main path.
A completed two-upgrade crosspath includes its first purchase as well as its second.

Name the receiving attack or effect when inheritance matters.
State where damage, speed, pierce, range, detection, and status effects apply.
A simple, direct combination is valid; do not invent a special synergy for every build.

Crosspaths can create specific interactions, such as faster production of a placed effect or extending a purchased status effect to a secondary attack.
Describe that interaction where it occurs rather than granting unrelated bonuses.

Do not automatically make attack speed reduce ability cooldowns or make range enlarge an independent explosion.
Define positive scope where needed rather than printing a list of unrelated systems that stay unchanged.

Review whether the two crosspath options meaningfully affect the main purchase.
Keep that comparison private.
The final descriptions must say what changes, not why the player should choose one option.

## Private design review

Review power and progression without assigning numbers or adding a commentary section.

Reject broad all-source damage amplification in `1-x-x`, `2-x-x`, `x-1-x`, `x-2-x`, `x-x-1`, or `x-x-2`.
Early utility should affect bounded targets, require an application condition, or address a specific problem.

Check what happens when the player places many copies.
Look for overlapping support, permanent coverage, repeated displacement, and control loops.
Low personal damage is not a sufficient weakness for support that benefits an entire defense.

Use limits that belong to the mechanic.
Target eligibility, application capacity, consumed effects, downtime, or protection against repeated control can bound a powerful effect.
Do not attach every possible restriction to every upgrade or disable independent attacks to manufacture a weakness.

The final purchase may overcome earlier restrictions without replacing every other unit's function.
Do not use price as the explanation at the concept stage.

Check for missing effects, illegal builds, extra activation buttons, lost purchased benefits, unsupported character claims, ambiguous pierce, and crosspaths that stop working after transformations.

Check that all 12 early combinations and all 36 advanced combinations are covered.
Grouped advanced entries must still identify each covered build and account for later changes.

Check that additions actually coexist, replacements are explicit, prediction is not silently removed, and later upgrades do not accidentally reduce an existing crowd attack to one target.

Reject repetitive paths whose only changes are larger numbers, and reject complicated mechanics that add rules without worthwhile behavior.
Revise failures privately.
Do not print scores, verdicts, balance claims, or forced weakness paragraphs.

## Prose unslop

Apply this pass to every output. It is part of the skill, not a separate optional instruction.

1. Scan for the problems below.
2. Rewrite while preserving mechanics, meaning, and requested organization.
3. Ask what still sounds autogenerated. Fix it.

Write what the unit does.
Replace "expresses its identity," "rewards strategic placement," "provides meaningful utility," and "delivers a powerful payoff" with the actual action or condition.
Do not narrate compliance with the design brief.

Delete usage advice and design reasoning from descriptions.
Keep the underlying mechanic: "has a long cooldown," not "has a long cooldown, so players must time it carefully."

Delete unchanged defaults, irrelevant absences, and explanations that merely restate a technical term.
Do not shorten the text by removing actual attacks, triggers, effects, restrictions, or crosspath interactions.

Remove filler and shallow phrases such as "highlighting," "ensuring," "showcasing," and "fostering." Remove vague attribution; name a source when a factual claim needs one.

Use plain words and direct verbs.
Prefer "is," "has," "uses," and "helps" to "serves as," "boasts," "utilizes," "leverages," and "facilitates."

Avoid stock vocabulary such as additionally, crucial, delve, enduring, enhance, interplay, intricate, pivotal, tapestry, testament, underscore, and vibrant.

Use complete sentences, usually with one main idea each.
Name the actor.
Remove weak adverbs and unnecessary passive voice.
Do not compress descriptions into fragments or symbol-heavy shorthand.
Build codes are identifiers, not substitutes for sentences.

Use one term consistently. Do not alternate among debuff, mark, curse, and status for the same mechanic.

Do not force ideas into groups of three.
Avoid fake "from X to Y" ranges, "not just X, but Y," rhetorical fragments, dramatic claims, and stock metaphors.

Avoid abstract technical metaphors such as substrate, wedge, vector, nexus, primitive, scaffolding, paradigm, north star, and flywheel.

Use straight quotes.
Do not use em dashes or en dashes.
Use sentence case for headings.
Limit bold text, colons, and lists to places where they help readers find information.
Do not repeat a label in the sentence immediately following it.

Remove decorative emojis, greetings, praise, agreement filler, generic conclusions, and offers to continue.
Use one clear qualifier where uncertainty matters instead of layers of hedging.

The final check must remove vague claims, repetition, filler, forced structure, unnecessary metaphors, dense sentences, decorative formatting, chatbot phrasing, and redundant statements of absence.
Preserve the design information.

## Output

Write the character name, followed by `0-0-0` and a short description of its starting attack.

Give each path a name.
Label every upgrade with its wildcard build reference and upgrade name, such as `3-x-x Jet Gatling`.
Describe its full changes in ordinary paragraphs.
Put relevant conditions and limits with the effect they belong to.

Follow with crosspaths.
Include the 12 early-build entries and the `3+-x-x`, `x-3+-x`, and `x-x-3+` groups.
Use concrete build references inside those groups.
Group identical advanced interactions only under the coverage rules above.

Do not force mechanics into narrow table cells.
Do not add a rules section, strategy guide, weakness summary, proposal disclaimer, capability report, or review certificate.
Do not explain why a purchase is good.

When patch notes are requested, name the affected build or upgrade, state the changed behavior and relevant crosspath changes, and distinguish changes from fixes.
Include a reason only when requested.
Do not invent a revision history.
