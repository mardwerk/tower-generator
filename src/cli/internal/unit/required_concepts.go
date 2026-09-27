package unit

import (
	"fmt"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Required concepts (#61, SOL-61-13). The v36 Luffy plan omitted Gear 4 and
// Gear 5, which the owner holds essential to Luffy; a failing omission
// verdict is a Finding and does not stop publication. Strong salience stays
// only a major floor (SOL-61-05), and the Default adds no global rule:
// the owner names the concepts a character needs in its Request instead,
// with an optional reason. Code checks the plan the way it checks a core
// concept: each required concept is the base attack or a repertoire entry
// named for it, matched by name or by an alias of the source technique it
// names (listsTechnique), is never omitted whole, and is the technique of a
// purchase that promises what one of its effects is adapted as or carries a
// proposed mechanic. Whether that purchase carries the concept's central
// effect is the review's judgment (CoreSpiritRule).

// MaxRequiredConcepts bounds a Request's required concepts.
const MaxRequiredConcepts = 8

// RequiredConceptSchema is one entry of a Request's requiredConcepts.
var RequiredConceptSchema = s.StrictObject(
	s.F("name", s.String().Trim().Min(1).Max(80)),
	s.F("reason", s.Optional(s.String().Trim().Min(1).Max(300))),
)

// RequiredConceptRule is the rule the plan prompt, the plan check and the
// review state for a Request with required concepts.
const RequiredConceptRule = "Each required concept of the Request is the base attack, or a repertoire entry named for it that is the technique of at least one purchase adapting its central effect, as a typed change or a proposed mechanic of that purchase; it is never omitted whole, only in named aspects, and one adapted only by a proposed mechanic is an unresolved design gap, because a proposal grants no behavior."

// requiredConceptNames checks a Request's required concepts: each name once.
func requiredConceptNames(concepts []RequiredConcept) error {
	for i, concept := range concepts {
		for _, earlier := range concepts[:i] {
			if sameTechnique(earlier.Name, concept.Name) {
				return fmt.Errorf("Required concepts must name each concept once; %q is listed twice.", concept.Name)
			}
		}
	}
	return nil
}

// conceptTechnique is the source technique a required concept names, by
// its name or an alias; nil when it names none.
func conceptTechnique(concept RequiredConcept, request *Request) *SourceTechnique {
	if request.SourceTechniques == nil {
		return nil
	}
	for i, technique := range *request.SourceTechniques {
		if listsTechnique(concept.Name, technique) {
			return &(*request.SourceTechniques)[i]
		}
	}
	return nil
}

// namesConcept reports a plan name for a required concept: named for it,
// or for the source technique it names or one of that technique's aliases.
func namesConcept(name string, concept RequiredConcept, technique *SourceTechnique) bool {
	return listsByName(name, concept.Name) || (technique != nil && listsTechnique(name, *technique))
}

// conceptAdaptation is what a plan does with one required concept.
type conceptAdaptation struct {
	concept   RequiredConcept
	base      bool
	entries   []int
	omissions []int
	// adaptedBy are the build codes of the purchases whose technique is
	// one of entries; adapting are those that promise what one of that
	// entry's effects is adapted as or carry a proposed mechanic.
	adaptedBy, adapting []string
}

// requiredConceptAdaptations lists, for each required concept, the plan's
// entries and omissions named for it and the purchases that adapt it.
func requiredConceptAdaptations(plan DesignPlan, request *Request) []conceptAdaptation {
	var out []conceptAdaptation
	for _, concept := range request.RequiredConcepts {
		technique := conceptTechnique(concept, request)
		a := conceptAdaptation{concept: concept, base: namesConcept(plan.Base.Name, concept, technique)}
		for index, entry := range plan.Repertoire {
			if namesConcept(entry.Name, concept, technique) {
				a.entries = append(a.entries, index)
			}
		}
		for index, omission := range plan.OmittedTechniques {
			if namesConcept(omission.Name, concept, technique) {
				a.omissions = append(a.omissions, index)
			}
		}
		if plan.UpgradeIntents != nil {
			for pathIndex := range m.PathKeys {
				for tier := 1; tier <= len(m.TierKeys); tier++ {
					intent := *plan.UpgradeIntents.At(pathIndex).At(tier)
					for _, index := range a.entries {
						if !sameTechnique(intent.Technique, plan.Repertoire[index].Name) {
							continue
						}
						code := BuildCode(pathIndex, tier)
						a.adaptedBy = append(a.adaptedBy, code)
						if typedAdaptation(plan.Repertoire[index], intent) || len(intent.ProposedMechanics) > 0 {
							a.adapting = append(a.adapting, code)
						}
						break
					}
				}
			}
		}
		out = append(out, a)
	}
	return out
}

// requiredIssue is one required concept failure: its facts, the rule, then
// the fix.
func requiredIssue(path, facts, fix string) m.Issue {
	return m.Issue{Path: path, Message: facts + " " + RequiredConceptRule + " " + fix}
}

// RequiredConceptIssues checks a plan against its Request's required
// concepts; which purchases adapt them only when the plan names its
// purchases' techniques. A saved plan is never checked again.
func RequiredConceptIssues(plan DesignPlan, request *Request) []m.Issue {
	if len(request.RequiredConcepts) == 0 {
		return nil
	}
	var issues []m.Issue
	for _, a := range requiredConceptAdaptations(plan, request) {
		name := a.concept.Name
		for _, index := range a.omissions {
			issues = append(issues, requiredIssue(fmt.Sprintf("omittedTechniques.%d", index),
				fmt.Sprintf("omittedTechniques omits %q, the required concept %q, whole.", plan.OmittedTechniques[index].Name, name),
				"Remove it from omittedTechniques, list it in the repertoire and adapt its central effect on a purchase whose technique it is, as a typed change or a proposed mechanic; name each aspect the Definition cannot express as an effect of that entry with adaptedAs empty."))
		}
		switch {
		case a.base || plan.UpgradeIntents == nil:
		case len(a.entries) == 0:
			issues = append(issues, requiredIssue("repertoire",
				fmt.Sprintf("No repertoire entry is named for the required concept %q.", name),
				fmt.Sprintf("Add it to the repertoire by the name %q, or by the name or an alias sourceTechniques gives it, citing its passages, and make it the technique of a purchase that adapts its central effect.", name)))
		case len(a.adaptedBy) == 0:
			issues = append(issues, requiredIssue(fmt.Sprintf("repertoire.%d", a.entries[0]),
				fmt.Sprintf("No purchase names %q, the entry of the required concept %q, as its technique.", plan.Repertoire[a.entries[0]].Name, name),
				"Make it the technique of a purchase whose typed changes do what it does in play, ideally one that defines a path, or name the mechanic it needs in that milestone's proposedMechanics."))
		case len(a.adapting) == 0:
			issues = append(issues, requiredIssue(fmt.Sprintf("repertoire.%d", a.entries[0]),
				fmt.Sprintf("%s %s %q, the entry of the required concept %q, but %s nothing one of its effects is adapted as and %s no proposed mechanic.", joinWith(a.adaptedBy, "and", ""), pluralWord(len(a.adaptedBy), "adapts", "adapt"), plan.Repertoire[a.entries[0]].Name, name, pluralWord(len(a.adaptedBy), "promises", "promise"), pluralWord(len(a.adaptedBy), "carries", "carry")),
				"Adapt one of its effects as a promise of a purchase whose technique it is, or name the mechanic it needs in that milestone's proposedMechanics."))
		}
	}
	return issues
}

// requiredConceptsContext is the Request's required concepts as a prompt
// gives them.
func requiredConceptsContext(request *Request) []any {
	out := []any{}
	for _, concept := range request.RequiredConcepts {
		entry := s.NewObject().Set("name", concept.Name)
		if concept.Reason != "" {
			entry.Set("reason", concept.Reason)
		}
		out = append(out, entry)
	}
	return out
}

// reviewRequiredConcepts is the review's view of each required concept:
// the plan's entries named for it and the purchases whose technique those
// entries are (adaptedBy), as code found them, and whether it is the base
// attack.
func reviewRequiredConcepts(plan *DesignPlan, request *Request) []any {
	out := requiredConceptsContext(request)
	if plan == nil {
		return out
	}
	for i, a := range requiredConceptAdaptations(*plan, request) {
		entry := out[i].(*s.Object)
		var names []string
		for _, index := range a.entries {
			names = append(names, plan.Repertoire[index].Name)
		}
		entry.Set("baseAttack", a.base).Set("entries", anyStrings(names)).Set("adaptedBy", anyStrings(a.adaptedBy))
	}
	return out
}
