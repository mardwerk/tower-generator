package unit

import (
	"fmt"
	"regexp"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// A source technique or form belongs to one path. Two paths that develop the
// same technique read as one path split in two, and an Active Ability that
// boosts another path's form makes one path depend on the other's identity.
// The base attack is shared by every path and belongs to none.

// ownedTechnique is a repertoire technique and the paths whose own plan text
// names it.
type ownedTechnique struct {
	name     string
	patterns []*regexp.Regexp
	paths    []int
}

var parenthetical = regexp.MustCompile(`\s*\([^)]*\)`)

// techniquePatterns match a technique name as whole words, ignoring case and
// parenthetical notes; "A / B" names two spellings of one technique.
func techniquePatterns(name string) []*regexp.Regexp {
	var out []*regexp.Regexp
	for _, spelling := range strings.Split(parenthetical.ReplaceAllString(name, ""), "/") {
		words := strings.Fields(spelling)
		if len(strings.Join(words, " ")) < 3 {
			continue
		}
		for i, word := range words {
			words[i] = regexp.QuoteMeta(word)
		}
		out = append(out, regexp.MustCompile(`(?i)(?:^|[^\p{L}\p{N}])`+strings.Join(words, `\s+`)+`(?:$|[^\p{L}\p{N}])`))
	}
	return out
}

func mentions(patterns []*regexp.Regexp, text string) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(text) {
			return true
		}
	}
	return false
}

// pathText is a planned path's own description: its name, reason to buy,
// purchases and capstone. Its weakness may compare it with another path and
// is left out.
func pathText(branch PlanBranch) string {
	texts := []string{branch.Name, branch.BuyFor, branch.CapstoneValue}
	for tier := 1; tier <= 5; tier++ {
		texts = append(texts, branch.Milestones.At(tier))
	}
	return strings.Join(texts, "\n")
}

// techniqueOwnership lists the repertoire techniques other than the base
// attack with the paths that name them.
func techniqueOwnership(plan DesignPlan) []ownedTechnique {
	var out []ownedTechnique
	for _, entry := range plan.Repertoire {
		patterns := techniquePatterns(entry.Name)
		if len(patterns) == 0 || mentions(patterns, plan.Base.Name) {
			continue
		}
		owned := ownedTechnique{name: entry.Name, patterns: patterns}
		for index := range m.PathKeys {
			if mentions(patterns, pathText(*plan.Paths.At(index))) {
				owned.paths = append(owned.paths, index)
			}
		}
		out = append(out, owned)
	}
	return out
}

func pathList(paths []int) string {
	names := make([]string, len(paths))
	for i, index := range paths {
		names[i] = pathPosition(m.PathKeys[index])
	}
	if len(names) == 2 {
		return names[0] + " and " + names[1]
	}
	return strings.Join(names[:len(names)-1], ", ") + " and " + names[len(names)-1]
}

// PlanOwnershipIssues rejects a plan whose paths share a repertoire technique.
func PlanOwnershipIssues(plan DesignPlan) []m.Issue {
	var issues []m.Issue
	for _, technique := range techniqueOwnership(plan) {
		if len(technique.paths) < 2 {
			continue
		}
		issues = append(issues, m.Issue{
			Path:    "paths." + m.PathKeys[technique.paths[1]],
			Message: fmt.Sprintf("The %s paths both name %q. A source technique or form belongs to one path: give it to one path and build the others from their own techniques. The middle path's Active Ability comes from its own technique, not from a form another path develops. List distinct techniques separately in the repertoire when the source distinguishes them.", pathList(technique.paths), technique.name),
		})
	}
	return issues
}

// NamedTechniqueIssues rejects a purchase or Active Ability named after a
// technique that the plan gives to another path.
func NamedTechniqueIssues(blueprint m.Blueprint, plan DesignPlan) []m.Issue {
	var issues []m.Issue
	for _, technique := range techniqueOwnership(plan) {
		if len(technique.paths) != 1 {
			continue
		}
		owner := technique.paths[0]
		for index, path := range m.PathKeys {
			if index == owner {
				continue
			}
			report := func(where, name string) {
				issues = append(issues, m.Issue{
					Path:    "paths." + path + ".tiers." + where,
					Message: fmt.Sprintf("%q names %q, which the plan gives to the %s path. Name this path's purchases and Active Ability after its own techniques.", name, technique.name, pathPosition(m.PathKeys[owner])),
				})
			}
			for tier := 1; tier <= 5; tier++ {
				purchase := blueprint.Paths.At(index).Tiers.At(tier)
				key := m.TierKeys[tier-1]
				if mentions(technique.patterns, purchase.Name) {
					report(key+".name", purchase.Name)
				}
				for number, change := range purchase.Changes {
					if change.Boost != nil && mentions(technique.patterns, change.Boost.Name) {
						report(fmt.Sprintf("%s.changes.%d.boost.name", key, number), change.Boost.Name)
					}
				}
			}
		}
	}
	return issues
}
