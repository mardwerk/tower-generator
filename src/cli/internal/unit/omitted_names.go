package unit

import (
	"slices"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// Omitted names (#61, SOL-61-17): the v38 Escanor base attack and the
// first purchase of every path said "Rhitta slash" while the plan omitted
// Rhitta whole. A plan that names a technique it omits whole contradicts
// the omission (OmittedNameRule), but a text may also name a technique to
// set it apart, so code only lists where the plan names it, as a cue beside
// the omission's Review verdict, and gates nothing.

// omissionNamedBy lists where the plan names an omitted technique, by its
// name or an alias of the source technique it is named for: the
// signature's name and adaptation, the base attack's name and behavior, and
// each purchase's name and planned text, as the review sees them
// (designPlan.signature.name, 0-0-0 is designPlan.base, "1-x-x name",
// "1-x-x adaptation" or "1-x-x plannedChange"). A name matches as whole
// words, case and punctuation aside, as conceptKey reads it.
func omissionNamedBy(plan DesignPlan, candidate Candidate, techniques []SourceTechnique, omission PlanOmission) []string {
	names := []string{omission.Name}
	for _, technique := range techniques {
		if listsTechnique(omission.Name, technique) {
			names = append(names, technique.Names()...)
			break
		}
	}
	var keys [][]string
	for _, name := range names {
		if key := conceptKey(name); len(key) > 0 && !slices.ContainsFunc(keys, func(k []string) bool { return slices.Equal(k, key) }) {
			keys = append(keys, key)
		}
	}
	namesIt := func(text string) bool {
		found := words(text)
		return slices.ContainsFunc(keys, func(key []string) bool { return containsRun(found, key) })
	}
	var out []string
	add := func(where, text string) {
		if namesIt(text) {
			out = append(out, where)
		}
	}
	add("designPlan.signature.name", plan.Signature.Name)
	add("designPlan.signature.adaptation", plan.Signature.Adaptation)
	add("designPlan.base.name", plan.Base.Name)
	add("designPlan.base.behavior", plan.Base.Behavior)
	for pathIndex := range m.PathKeys {
		for tier := 1; tier <= len(m.TierKeys); tier++ {
			code := BuildCode(pathIndex, tier)
			if len(candidate.Paths) == len(m.PathKeys) {
				for _, t := range candidate.Paths[pathIndex].Tiers {
					if t.Tier == tier {
						add(code+" name", t.Name)
					}
				}
			}
			field := "plannedChange"
			if PurchaseAdaptation(&plan, pathIndex, tier) != "" {
				field = "adaptation"
			}
			add(code+" "+field, plan.Paths.At(pathIndex).Milestones.At(tier))
		}
	}
	return out
}

// containsRun reports words holding run as consecutive words.
func containsRun(words, run []string) bool {
	for start := 0; start+len(run) <= len(words); start++ {
		if slices.Equal(words[start:start+len(run)], run) {
			return true
		}
	}
	return false
}
