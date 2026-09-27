package unit

import (
	"fmt"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// Every planned purchase names the technique it adapts: a repertoire entry or
// the base attack. The names are typed plan data, so the checks below do not
// depend on how a purchase or Active Ability is later named.
//
// A path needs a technique of its own. Two paths may adapt the same technique
// as crosspath synergy, but only where a legal build owns both purchases;
// otherwise the technique is split between purchases that never meet, such as
// a middle-path Active Ability boosting a form the top path reaches at its
// third purchase.

// adaptation is one purchase's technique.
type adaptation struct {
	path, tier int
}

func sameTechnique(a, b string) bool {
	return strings.EqualFold(strings.Join(strings.Fields(a), " "), strings.Join(strings.Fields(b), " "))
}

// ownedTogether reports a legal build that owns both purchases.
func ownedTogether(a, b adaptation, builds []m.Selection) bool {
	for _, build := range builds {
		if build[a.path] >= a.tier && build[b.path] >= b.tier {
			return true
		}
	}
	return false
}

// PlanTechniqueIssues checks the typed technique of every planned purchase.
// Plans without techniques predate the field and are not checked.
func PlanTechniqueIssues(plan DesignPlan, definition m.Definition) []m.Issue {
	if plan.UpgradeIntents == nil {
		return nil
	}
	var issues []m.Issue
	known := func(name string) (canonical string, base bool, ok bool) {
		if sameTechnique(name, plan.Base.Name) {
			return plan.Base.Name, true, true
		}
		for _, entry := range plan.Repertoire {
			if sameTechnique(name, entry.Name) {
				return entry.Name, false, true
			}
		}
		return "", false, false
	}
	var order []string
	uses := map[string][]adaptation{}
	for pathIndex, path := range m.PathKeys {
		for tier := 1; tier <= 5; tier++ {
			technique := plan.UpgradeIntents.At(pathIndex).At(tier).Technique
			if technique == "" {
				return nil
			}
			name, base, ok := known(technique)
			if !ok {
				issues = append(issues, m.Issue{
					Path:    "upgradeIntents." + path + "." + m.TierKeys[tier-1] + ".technique",
					Message: fmt.Sprintf("%q is neither the base attack nor a repertoire entry. Name the technique this purchase adapts exactly as the base or repertoire names it.", technique),
				})
				continue
			}
			if base {
				continue
			}
			if _, seen := uses[name]; !seen {
				order = append(order, name)
			}
			uses[name] = append(uses[name], adaptation{pathIndex, tier})
		}
	}
	builds := m.AllLegalBuilds(definition)
	own := make([]bool, len(m.PathKeys))
	for _, name := range order {
		paths := map[int]bool{}
		for _, use := range uses[name] {
			paths[use.path] = true
		}
		if len(paths) == 1 {
			own[uses[name][0].path] = true
			continue
		}
		reported := false
		for i, a := range uses[name] {
			for _, b := range uses[name][i+1:] {
				if reported || a.path == b.path || ownedTogether(a, b, builds) {
					continue
				}
				reported = true
				issues = append(issues, m.Issue{
					Path:    "upgradeIntents." + m.PathKeys[b.path] + "." + m.TierKeys[b.tier-1] + ".technique",
					Message: fmt.Sprintf("%s and %s both adapt %q, and no legal build owns both, so neither can build on the other. Two paths share a technique only as crosspath synergy: give it to one path, or keep the other path's use to a purchase a crosspath can own with it (a first or second purchase).", BuildCode(a.path, a.tier), BuildCode(b.path, b.tier), name),
				})
			}
		}
	}
	for pathIndex, path := range m.PathKeys {
		if !own[pathIndex] {
			issues = append(issues, m.Issue{
				Path:    "upgradeIntents." + path,
				Message: fmt.Sprintf("The %s path adapts no technique of its own: every purchase adapts the base attack or a technique another path also adapts. Give the path a repertoire technique that is its identity; it may still build on other paths' techniques in crosspaths.", m.PathPosition(path)),
			})
		}
	}
	return issues
}
