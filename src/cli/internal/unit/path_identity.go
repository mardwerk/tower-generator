package unit

import (
	"fmt"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// guidePathIdentity is the draft and repair guidance under
// requireTier3PathIdentity.
const guidePathIdentity = m.PathIdentityRule + " Code compares what each path's resolved purchases improve or unlock and the proposed mechanics they carry, so a change the plan does not promise on another path can take a third purchase's identity."

// PathIdentityFindingRule names the check Finding of requireTier3PathIdentity.
const PathIdentityFindingRule = "path-identity"

// pathIdentityOn reports whether a Definition selects requireTier3PathIdentity.
func pathIdentityOn(definition m.Definition) bool {
	return definition.Profile.DesignPolicy.RequiresPathIdentity()
}

// exclusiveKeys lists the benefits of a third purchase that no purchase of
// the other paths has.
func exclusiveKeys(all [3][5]benefitSet, pathIndex int) []string {
	owners := benefitOwners(all, pathIndex)
	var out []string
	for _, key := range all[pathIndex][2].keys() {
		if _, shared := owners[key]; !shared {
			out = append(out, key)
		}
	}
	return out
}

// identityFacts reads what a third purchase has and where the other paths
// have it too.
func identityFacts(all [3][5]benefitSet, pathIndex int, verb string) string {
	third := all[pathIndex][2]
	code := BuildCode(pathIndex, 3)
	if len(third) == 0 {
		return fmt.Sprintf("%s %s nothing that counts.", code, verb)
	}
	keys := third.keys()
	each := "it"
	if len(keys) > 1 {
		each = "each"
	}
	return fmt.Sprintf("%s %s %s, and a purchase of another path %s %s too: %s.", code, verb, third.text(keys), verb, each, sharedText(third, keys, benefitOwners(all, pathIndex)))
}

// identitySuggestions names improvements no purchase of the other paths
// promises, then a proposed mechanic.
func identitySuggestions(all [3][5]benefitSet, pathIndex int) string {
	owners := benefitOwners(all, pathIndex)
	var out []string
	for _, c := range []string{"projectiles", "splash", "pierce", "range", "attack-rate", "damage"} {
		if _, taken := owners[c]; !taken && len(out) < 2 {
			out = append(out, c)
		}
	}
	return joinWith(append(out, "a proposed mechanic the sources describe"), "or", "")
}

// pathIdentityIssues is the plan form of requireTier3PathIdentity: a third
// milestone must promise an improvement, unlock or proposed mechanic that no
// milestone of the other two paths promises.
func pathIdentityIssues(intents *UpgradeIntents, definition m.Definition) []m.Issue {
	var all [3][5]benefitSet
	for index := range m.PathKeys {
		for tier := 1; tier <= 5; tier++ {
			all[index][tier-1] = plannedBenefits(*intents.At(index).At(tier), definition)
		}
	}
	var issues []m.Issue
	for index, path := range m.PathKeys {
		if len(exclusiveKeys(all, index)) > 0 {
			continue
		}
		code := BuildCode(index, 3)
		issues = append(issues, m.Issue{
			Path: "upgradeIntents." + path + ".tier3",
			Message: fmt.Sprintf("%s %s Give %s an improvement, unlock or proposed mechanic from this path's own technique that no purchase of the other two paths promises, such as %s.",
				identityFacts(all, index, "promises"), m.PathIdentityRule, code, identitySuggestions(all, index)),
		})
	}
	return issues
}

// PathIdentityIssues is the resolved form of requireTier3PathIdentity: each
// path's resolved third purchase must give an improvement, unlock or
// proposed mechanic that no resolved purchase of the other paths gives. With
// a retained plan it targets the purchase that departs from it, so
// TargetedTierRepair rebuilds that tier: the third purchase when it drops its
// planned identity, or another path's purchase that gives it without the
// plan promising it there. Like EarlyBenefitsIssues it runs beside
// PlanIntentIssues in drafting and check, never in ValidateTyped. intents
// may be nil for a draft without a plan.
func PathIdentityIssues(blueprint m.Blueprint, intents *UpgradeIntents, definition m.Definition) []m.Issue {
	if !pathIdentityOn(definition) {
		return nil
	}
	var resolved, planned [3][5]benefitSet
	for index := range m.PathKeys {
		for tier := 1; tier <= 5; tier++ {
			resolved[index][tier-1] = resolvedBenefits(&blueprint, definition, index, tier)
			if intents != nil {
				planned[index][tier-1] = plannedBenefits(*intents.At(index).At(tier), definition)
			}
		}
	}
	var issues []m.Issue
	for index := range m.PathKeys {
		if len(exclusiveKeys(resolved, index)) > 0 {
			continue
		}
		code := BuildCode(index, 3)
		target, tier := index, 3
		fix := fmt.Sprintf("Give %s an improvement, unlock or proposed mechanic from this path's own technique that no purchase of the other two paths gives.", code)
		if intents != nil {
		search:
			for _, key := range exclusiveKeys(planned, index) {
				text := planned[index][2][key]
				if _, kept := resolved[index][2][key]; !kept {
					fix = fmt.Sprintf("The retained plan promises %s at %s and at no purchase of the other paths; the resolved purchase does not give it. Deliver %s's promised %s so the path keeps its identity.", text, code, code, text)
					break
				}
				for other := range m.PathKeys {
					if other == index {
						continue
					}
					for t := 1; t <= 5; t++ {
						_, gives := resolved[other][t-1][key]
						_, promised := planned[other][t-1][key]
						if gives && !promised {
							target, tier = other, t
							other := BuildCode(other, t)
							fix = fmt.Sprintf("The retained plan promises %s at %s and at no purchase of the other paths, but resolved %s also gives %s, which the plan does not promise there. Remove that %s from %s and keep %s's promises, so %s keeps the path's identity.", text, code, other, text, text, other, other, code)
							break search
						}
					}
				}
			}
		}
		issues = append(issues, m.Issue{
			Path:    fmt.Sprintf("paths.%s.tiers.%s", m.PathKeys[target], m.TierKeys[tier-1]),
			Message: "Resolved " + identityFacts(resolved, index, "gives") + " " + m.PathIdentityRule + " " + fix,
		})
	}
	return issues
}
