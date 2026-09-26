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
// technique promises. Plans without effects or techniques predate the fields
// and are not checked.
func PlanEffectIssues(plan DesignPlan) []m.Issue {
	if plan.UpgradeIntents == nil {
		return nil
	}
	var issues []m.Issue
	for index, entry := range plan.Repertoire {
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
		for effectIndex, effect := range entry.Effects {
			for _, id := range effect.AdaptedAs {
				if slices.ContainsFunc(promised, func(p string) bool { return promiseMatches(id, p) }) {
					continue
				}
				message := fmt.Sprintf("%q is adapted as %s, but no purchase that adapts %q promises %s. Promise %s on a purchase whose technique is %q, or leave adaptedAs empty and give the reason the effect is omitted.", effect.Effect, id, entry.Name, id, id, entry.Name)
				if !adapted {
					message = fmt.Sprintf("%q is adapted as %s, but no purchase adapts %q. Name it as the technique of the purchases that adapt it, or leave adaptedAs empty and give the reason the effect is omitted.", effect.Effect, id, entry.Name)
				}
				issues = append(issues, m.Issue{Path: fmt.Sprintf("repertoire.%d.effects.%d.adaptedAs", index, effectIndex), Message: message})
			}
		}
	}
	return issues
}
