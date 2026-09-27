package unit

import (
	"fmt"
	"slices"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// A repertoire entry lists the effects its cited passages describe, each
// adapted as promises or omitted with a reason. Nothing checked that before:
// Luffy's Gear 2 cites a passage saying it raises strength and speed, and 7
// of 10 plans on #27 promised only speed on every Gear 2 purchase without
// recording the dropped strength. Code cannot tell whether the list covers
// the passages; the review judges that. Code checks that the list is kept:
// every promise an effect is adapted as appears on a purchase that adapts
// the technique, and the mechanics promise checks then hold that purchase to
// a typed change.

// promiseMatches reports a promise that keeps an adaptedAs ID: the same ID,
// or its plain or active- form, so strength adapted as damage is kept by
// active-damage on the Active's purchase.
func promiseMatches(adapted, promised string) bool {
	plain := func(id string) string { return strings.TrimPrefix(id, "active-") }
	return adapted == promised || plain(adapted) == plain(promised)
}

// PlanEffectIssues reports an adapted effect that no purchase of its
// technique promises. It checks only entries that purchases adapt: an entry
// named like the base attack is shown by 0-0-0, and an entry no purchase
// names adds nothing to the unit, so its list is the review's to judge (in
// 5 Luffy plans on #27 these two cases sent 4 plans back to the plan retry,
// and one retry failed). Plans without effects or techniques predate the
// fields and are not checked.
func PlanEffectIssues(plan DesignPlan) []m.Issue {
	if plan.UpgradeIntents == nil {
		return nil
	}
	var issues []m.Issue
	for index, entry := range plan.Repertoire {
		if sameTechnique(entry.Name, plan.Base.Name) {
			continue
		}
		var promised []string
		adapted := false
		for pathIndex := range m.PathKeys {
			for tier := 1; tier <= 5; tier++ {
				intent := plan.UpgradeIntents.At(pathIndex).At(tier)
				if intent.Technique == "" {
					return nil
				}
				if !sameTechnique(intent.Technique, entry.Name) {
					continue
				}
				adapted = true
				promised = append(append(promised, intent.Improves...), intent.Lowers...)
				if intent.Unlock != "none" {
					promised = append(promised, intent.Unlock)
				}
			}
		}
		if !adapted {
			continue
		}
		for effectIndex, effect := range entry.Effects {
			for _, id := range effect.AdaptedAs {
				if slices.ContainsFunc(promised, func(p string) bool { return promiseMatches(id, p) }) {
					continue
				}
				issues = append(issues, m.Issue{
					Path:    fmt.Sprintf("repertoire.%d.effects.%d.adaptedAs", index, effectIndex),
					Message: fmt.Sprintf("%q is adapted as %s, but no purchase that adapts %q promises %s. Promise %s on a purchase whose technique is %q, or leave adaptedAs empty and give the reason the effect is omitted.", effect.Effect, id, entry.Name, id, id, entry.Name),
				})
			}
		}
	}
	return issues
}

// PromisesWithoutEffect lists what one planned purchase improves or unlocks
// that no effect of its own technique adapts, matched as PlanEffectIssues
// matches them (promiseMatches). PlanEffectIssues checks the other
// direction, and a plan gate on this one would reject nearly every saved
// Luffy plan and the reference fixture (#57: 81 of 436 promises on technique
// purchases), so it is review context, not a finding: Sol's x-x-5 promised
// range from a Gear 4 Tankman entry whose one effect is adapted as damage,
// attack rate and knockback. A tradeoff in lowers is not credited to the
// technique and is left out. ok is false for a purchase that adapts the base
// attack, whose early stat steps are ordinary adaptations the review judges
// against 0-0-0, and when code cannot tell: the plan predates techniques or
// effects, or the purchase's technique names no repertoire entry.
func PromisesWithoutEffect(plan DesignPlan, pathIndex, tier int) (missing []string, ok bool) {
	if plan.UpgradeIntents == nil {
		return nil, false
	}
	intent := plan.UpgradeIntents.At(pathIndex).At(tier)
	withEffects := false
	var entry *PlanRepertoire
	for index := range plan.Repertoire {
		withEffects = withEffects || len(plan.Repertoire[index].Effects) > 0
		if entry == nil && sameTechnique(plan.Repertoire[index].Name, intent.Technique) {
			entry = &plan.Repertoire[index]
		}
	}
	if strings.TrimSpace(intent.Technique) == "" || entry == nil || !withEffects || sameTechnique(intent.Technique, plan.Base.Name) {
		return nil, false
	}
	promised := append([]string{}, intent.Improves...)
	if intent.Unlock != "" && intent.Unlock != "none" {
		promised = append(promised, intent.Unlock)
	}
	missing = []string{}
	for _, promise := range promised {
		adapted := false
		for _, effect := range entry.Effects {
			adapted = adapted || slices.ContainsFunc(effect.AdaptedAs, func(id string) bool { return promiseMatches(id, promise) })
		}
		if !adapted && !slices.Contains(missing, promise) {
			missing = append(missing, promise)
		}
	}
	return missing, true
}
