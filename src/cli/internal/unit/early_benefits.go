package unit

import (
	"fmt"
	"slices"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// guideEarlyBenefits is the draft and repair guidance under the field.
const guideEarlyBenefits = m.EarlyBenefitsRule + " Code compares what each path's resolved first and second purchases improve or unlock, so a change the plan does not promise there can make two paths match."

// guideExclusiveEarly is the draft and repair guidance under
// exclusiveEarlyBenefits.
const guideExclusiveEarly = m.ExclusiveEarlyBenefitsRule + " Code compares what each path's resolved first and second purchases improve or unlock, so a change the plan does not promise there can make two paths share one."

// earlyBenefitsOn reports whether a Definition selects distinctEarlyBenefits
// and not exclusiveEarlyBenefits, which subsumes it: early benefits that
// share nothing differ. With both set only the exclusive rule applies.
func earlyBenefitsOn(definition m.Definition) bool {
	policy := definition.Profile.DesignPolicy
	return policy != nil && policy.DistinctEarlyBenefits != nil && *policy.DistinctEarlyBenefits && !policy.ExcludesSharedEarlyBenefits()
}

// exclusiveEarlyOn reports whether a Definition selects exclusiveEarlyBenefits.
func exclusiveEarlyOn(definition m.Definition) bool {
	return definition.Profile.DesignPolicy.ExcludesSharedEarlyBenefits()
}

// EarlyBenefitsFindingRule names the check Finding of the early benefits
// rule a Definition selects, or "" when it selects none.
func EarlyBenefitsFindingRule(definition m.Definition) string {
	switch {
	case exclusiveEarlyOn(definition):
		return "exclusive-early-benefits"
	case earlyBenefitsOn(definition):
		return "distinct-early-benefits"
	}
	return ""
}

// earlyStep is one early purchase: the dimensions it improves and the
// capabilities it unlocks, each at most once. lowers never count.
type earlyStep struct {
	improves []string
	unlocks  []string
}

func (e earlyStep) key() string {
	return strings.Join(e.improves, ",") + "|" + strings.Join(e.unlocks, ",")
}

// text reads range and unlock camo, or nothing.
func (e earlyStep) text() string {
	parts := append([]string{}, e.improves...)
	for _, u := range e.unlocks {
		parts = append(parts, "unlock "+u)
	}
	return joinWith(parts, "and", "nothing")
}

func joinWith(parts []string, conjunction, empty string) string {
	switch len(parts) {
	case 0:
		return empty
	case 1:
		return parts[0]
	}
	return strings.Join(parts[:len(parts)-1], ", ") + " " + conjunction + " " + parts[len(parts)-1]
}

func sortedSet(values []string) []string {
	out := slices.Clone(values)
	slices.Sort(out)
	return slices.Compact(out)
}

func plannedStep(intent UpgradeIntent) earlyStep {
	step := earlyStep{improves: sortedSet(intent.Improves)}
	if intent.Unlock != "" && intent.Unlock != "none" {
		step.unlocks = []string{intent.Unlock}
	}
	return step
}

// earlyBenefits is a path's first two purchases taken together, in no
// order, keeping a dimension both purchases improve twice.
func earlyBenefits(steps [2]earlyStep) []string {
	var items []string
	for _, step := range steps {
		items = append(items, step.improves...)
		for _, u := range step.unlocks {
			items = append(items, "unlock "+u)
		}
	}
	slices.Sort(items)
	return items
}

func benefitsKey(steps [2]earlyStep) string { return strings.Join(earlyBenefits(steps), ",") }

// benefitsText reads damage twice and range for [damage damage range].
func benefitsText(steps [2]earlyStep) string {
	items := earlyBenefits(steps)
	var parts []string
	for i := 0; i < len(items); {
		j := i
		for j < len(items) && items[j] == items[i] {
			j++
		}
		if j-i == 1 {
			parts = append(parts, items[i])
		} else {
			parts = append(parts, items[i]+" twice")
		}
		i = j
	}
	return joinWith(parts, "and", "nothing")
}

// earlySuggestions names improvements the matching paths do not already share.
func earlySuggestions(steps [2]earlyStep) string {
	used := map[string]bool{}
	for _, s := range steps {
		for _, d := range s.improves {
			used[d] = true
		}
	}
	var out []string
	for _, c := range []string{"damage", "pierce", "range", "attack-rate"} {
		if !used[c] && len(out) < 2 {
			out = append(out, c)
		}
	}
	return joinWith(append(out, "personal detection"), "or", "")
}

// earlyBenefitsIssues is the plan form of distinctEarlyBenefits: two paths
// whose first two purchases promise the same multiset of improvement
// dimensions and unlocks. The later path in build-code order gets the issue.
func earlyBenefitsIssues(intents *UpgradeIntents) []m.Issue {
	var issues []m.Issue
	owner := map[string]int{}
	for index, path := range m.PathKeys {
		p := intents.At(index)
		steps := [2]earlyStep{plannedStep(*p.At(1)), plannedStep(*p.At(2))}
		key := benefitsKey(steps)
		other, taken := owner[key]
		if !taken {
			owner[key] = index
			continue
		}
		issues = append(issues, m.Issue{
			Path: "upgradeIntents." + path + ".tier2",
			Message: fmt.Sprintf("%s and %s together promise %s, the same as %s and %s. %s Redesign %s or %s: add or swap an improvement or unlock from this path's own technique, such as %s.",
				BuildCode(index, 1), BuildCode(index, 2), benefitsText(steps), BuildCode(other, 1), BuildCode(other, 2), m.EarlyBenefitsRule, BuildCode(index, 2), BuildCode(index, 1), earlySuggestions(steps)),
		})
	}
	return issues
}

// resolvedStep derives an early purchase's improvements and unlocks from
// the pure builds before and after it, with the promise check's measures.
func resolvedStep(blueprint *m.Blueprint, definition m.Definition, pathIndex, tier int) earlyStep {
	var before, after m.Selection
	before[pathIndex], after[pathIndex] = tier-1, tier
	b, a := m.ResolveUnchecked(blueprint, before), m.ResolveUnchecked(blueprint, after)
	path := m.PathKeys[pathIndex]
	var step earlyStep
	for _, d := range ImprovementsFor(&definition) {
		// The Active unlocks at the fourth purchase at the earliest.
		if strings.HasPrefix(d, "active-") {
			continue
		}
		prior := measures(b, path, d, definition)
		for i, v := range measures(a, path, d, definition) {
			if i < len(prior) && v > prior[i] {
				step.improves = append(step.improves, d)
				break
			}
		}
	}
	for _, u := range UnlocksFor(&definition) {
		if u == "none" || u == "manual-boost" || u == "active-follow-up" {
			continue
		}
		if unlockedIntent(b, a, u, pathIndex, tier, blueprint, definition) {
			step.unlocks = append(step.unlocks, u)
		}
	}
	step.improves, step.unlocks = sortedSet(step.improves), sortedSet(step.unlocks)
	return step
}

// stepDifference lists what a has and b lacks, as dimensions and unlocks.
func stepDifference(a, b earlyStep) []string {
	var out []string
	for _, d := range a.improves {
		if !slices.Contains(b.improves, d) {
			out = append(out, d)
		}
	}
	for _, u := range a.unlocks {
		if !slices.Contains(b.unlocks, u) {
			out = append(out, "unlock "+u)
		}
	}
	return out
}

// EarlyBenefitsIssues is the resolved form of distinctEarlyBenefits: the
// finished mechanics may not give two paths' first two purchases the same
// improvements and unlocks, for example through a change a distinct plan
// does not promise. It targets the purchase that departs from the retained
// plan, so TargetedTierRepair rebuilds that tier. It runs beside
// PlanIntentIssues in drafting and check, never in ValidateTyped, which
// rendering also calls. intents may be nil for a draft without a plan.
func EarlyBenefitsIssues(blueprint m.Blueprint, intents *UpgradeIntents, definition m.Definition) []m.Issue {
	if exclusiveEarlyOn(definition) {
		return exclusiveEarlyResolvedIssues(blueprint, intents, definition)
	}
	if !earlyBenefitsOn(definition) {
		return nil
	}
	var issues []m.Issue
	owner := map[string]int{}
	var resolved [3][2]earlyStep
	for index := range m.PathKeys {
		resolved[index] = [2]earlyStep{resolvedStep(&blueprint, definition, index, 1), resolvedStep(&blueprint, definition, index, 2)}
		key := benefitsKey(resolved[index])
		other, taken := owner[key]
		if !taken {
			owner[key] = index
			continue
		}
		facts := fmt.Sprintf("Resolved %s and %s together improve %s, the same as resolved %s and %s: %s gives %s, %s gives %s, %s gives %s and %s gives %s.",
			BuildCode(index, 1), BuildCode(index, 2), benefitsText(resolved[index]), BuildCode(other, 1), BuildCode(other, 2),
			BuildCode(other, 1), resolved[other][0].text(), BuildCode(other, 2), resolved[other][1].text(),
			BuildCode(index, 1), resolved[index][0].text(), BuildCode(index, 2), resolved[index][1].text())
		// Target the purchase that departs from its plan: this path's
		// second, then first, then the other path's. Else this path's second.
		target, tier := index, 2
		fix := fmt.Sprintf("Change what %s or %s improves or unlocks so the two paths' early benefits differ.", BuildCode(index, 2), BuildCode(index, 1))
		if intents != nil {
			fix = fmt.Sprintf("Change what %s or %s improves or unlocks so the two paths' early benefits differ, and keep every promise the retained plan makes there.", BuildCode(index, 2), BuildCode(index, 1))
		search:
			for _, candidate := range []int{index, other} {
				for t := 2; t >= 1; t-- {
					planned, actual := plannedStep(*intents.At(candidate).At(t)), resolved[candidate][t-1]
					if planned.key() == actual.key() {
						continue
					}
					target, tier = candidate, t
					code := BuildCode(candidate, t)
					fix = fmt.Sprintf("The retained plan promises %s at %s; the resolved purchase gives %s.", planned.text(), code, actual.text())
					if extra := stepDifference(actual, planned); len(extra) > 0 {
						fix += fmt.Sprintf(" Its %s, which the plan does not promise there, makes the two paths' early benefits match. Keep %s's promises and make its early benefits distinct: remove that %s from %s, or add another improvement or unlock from this path's own technique.",
							joinWith(extra, "and", ""), code, joinWith(extra, "and", ""), code)
					} else {
						fix += fmt.Sprintf(" Deliver %s's promised %s so the two paths' early benefits differ as the plan does.", code, joinWith(stepDifference(planned, actual), "and", ""))
					}
					break search
				}
			}
		}
		issues = append(issues, m.Issue{
			Path:    fmt.Sprintf("paths.%s.tiers.%s", m.PathKeys[target], m.TierKeys[tier-1]),
			Message: facts + " " + m.EarlyBenefitsRule + " " + fix,
		})
	}
	return issues
}

// ---- exclusiveEarlyBenefits ----

// earlySet is a path's first two steps taken together as benefits.
func earlySet(steps [2]earlyStep, definition m.Definition) benefitSet {
	out := stepBenefits(steps[0], definition)
	out.addAll(stepBenefits(steps[1], definition))
	return out
}

// sharedKeys lists the benefits two sets have in common.
func sharedKeys(a, b benefitSet) []string {
	var out []string
	for _, key := range a.keys() {
		if _, ok := b[key]; ok {
			out = append(out, key)
		}
	}
	return out
}

// sharingTier is the later of a path's first two steps that has one of the
// shared benefits.
func sharingTier(steps [2]earlyStep, keys []string, definition m.Definition) int {
	second := stepBenefits(steps[1], definition)
	for _, key := range keys {
		if _, ok := second[key]; ok {
			return 2
		}
	}
	return 1
}

// exclusiveSuggestions names improvements no path's first two purchases
// have, then personal detection when no path unlocks detection early, as a
// ", such as" clause; empty when every one is taken.
func exclusiveSuggestions(sets [3]benefitSet, definition m.Definition) string {
	used := benefitSet{}
	for _, set := range sets {
		used.addAll(set)
	}
	var out []string
	for _, c := range []string{"damage", "pierce", "range", "attack-rate"} {
		if _, taken := used[c]; !taken && len(out) < 2 {
			out = append(out, c)
		}
	}
	detects := false
	for _, key := range used.keys() {
		trait := strings.TrimPrefix(key, "unlock ")
		if trait != key && (trait == "camo" || (definition.IsV2() && isDetection(&definition, trait))) {
			detects = true
		}
	}
	if !detects {
		out = append(out, "personal detection")
	}
	if len(out) == 0 {
		return ""
	}
	return ", such as " + joinWith(out, "or", "")
}

// exclusiveEarlyIssues is the plan form of exclusiveEarlyBenefits: a
// benefit that one path's first two milestones promise may not be promised
// by another path's first two milestones. The later path in build-code
// order gets the issue, at the milestone that promises the shared benefit.
func exclusiveEarlyIssues(intents *UpgradeIntents, definition m.Definition) []m.Issue {
	var steps [3][2]earlyStep
	var sets [3]benefitSet
	for index := range m.PathKeys {
		p := intents.At(index)
		steps[index] = [2]earlyStep{plannedStep(*p.At(1)), plannedStep(*p.At(2))}
		sets[index] = earlySet(steps[index], definition)
	}
	var issues []m.Issue
	for index, path := range m.PathKeys {
		for other := 0; other < index; other++ {
			keys := sharedKeys(sets[index], sets[other])
			if len(keys) == 0 {
				continue
			}
			shared := sets[index].text(keys)
			tier := sharingTier(steps[index], keys, definition)
			pronoun := "it"
			if len(keys) > 1 {
				pronoun = "them"
			}
			issues = append(issues, m.Issue{
				Path: fmt.Sprintf("upgradeIntents.%s.tier%d", path, tier),
				Message: fmt.Sprintf("%s and %s promise %s, which %s and %s also promise. %s Redesign %s or %s: move %s to %s or later, or replace %s with an improvement or unlock from this path's own technique that no other path's first two purchases promise%s.",
					BuildCode(index, 1), BuildCode(index, 2), shared, BuildCode(other, 1), BuildCode(other, 2), m.ExclusiveEarlyBenefitsRule,
					BuildCode(index, tier), BuildCode(index, 3-tier), shared, BuildCode(index, 3), pronoun, exclusiveSuggestions(sets, definition)),
			})
		}
	}
	return issues
}

// exclusiveEarlyResolvedIssues is the resolved form of exclusiveEarlyBenefits.
// It compares the pure builds up to each path's second purchase with the
// promise check's measures and targets the purchase that departs from the
// retained plan by giving a shared benefit the plan does not promise there,
// so TargetedTierRepair rebuilds that tier.
func exclusiveEarlyResolvedIssues(blueprint m.Blueprint, intents *UpgradeIntents, definition m.Definition) []m.Issue {
	var resolved [3][2]earlyStep
	var sets [3]benefitSet
	for index := range m.PathKeys {
		resolved[index] = [2]earlyStep{resolvedStep(&blueprint, definition, index, 1), resolvedStep(&blueprint, definition, index, 2)}
		sets[index] = earlySet(resolved[index], definition)
	}
	var issues []m.Issue
	for index := range m.PathKeys {
		for other := 0; other < index; other++ {
			keys := sharedKeys(sets[index], sets[other])
			if len(keys) == 0 {
				continue
			}
			shared := sets[index].text(keys)
			facts := fmt.Sprintf("Resolved %s and %s give %s, which resolved %s and %s also give: %s gives %s, %s gives %s, %s gives %s and %s gives %s.",
				BuildCode(index, 1), BuildCode(index, 2), shared, BuildCode(other, 1), BuildCode(other, 2),
				BuildCode(other, 1), resolved[other][0].text(), BuildCode(other, 2), resolved[other][1].text(),
				BuildCode(index, 1), resolved[index][0].text(), BuildCode(index, 2), resolved[index][1].text())
			target, tier := index, sharingTier(resolved[index], keys, definition)
			fix := fmt.Sprintf("Change what %s or %s improves or unlocks so it shares nothing with %s and %s.", BuildCode(index, tier), BuildCode(index, 3-tier), BuildCode(other, 1), BuildCode(other, 2))
			if intents != nil {
			search:
				for _, candidate := range []int{index, other} {
					for t := 2; t >= 1; t-- {
						actual := stepBenefits(resolved[candidate][t-1], definition)
						planned := stepBenefits(plannedStep(*intents.At(candidate).At(t)), definition)
						for _, key := range keys {
							_, gives := actual[key]
							_, promised := planned[key]
							if !gives || promised {
								continue
							}
							target, tier = candidate, t
							code := BuildCode(candidate, t)
							fix = fmt.Sprintf("The retained plan does not promise %s at %s; the resolved purchase gives it, and the other path's first two purchases give it too. Remove that %s from %s and keep %s's promises.", actual[key], code, actual[key], code, code)
							break search
						}
					}
				}
			}
			issues = append(issues, m.Issue{
				Path:    fmt.Sprintf("paths.%s.tiers.%s", m.PathKeys[target], m.TierKeys[tier-1]),
				Message: facts + " " + m.ExclusiveEarlyBenefitsRule + " " + fix,
			})
		}
	}
	return issues
}
