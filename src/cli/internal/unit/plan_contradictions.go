package unit

import (
	"fmt"
	"slices"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Plan contradictions (#61, SOL-61-13). The v36 Luffy plan listed "Gear 4
// Python" twice in its repertoire (OPUS-NET-61-16). A technique has one
// entry, and a name the plan selects, as the base attack or a repertoire
// entry, which is every name a purchase may give as its technique, is never
// also a whole-technique omission. Code compares names exactly, as
// sameTechnique does, case and spacing aside: a substring is no match, so
// "Gear 4 Python" beside an omitted "Gear 4" or "Python" is left to the
// review, which judges whether the entry claims a form the plan omits.
// Both failures have a local fix, so each is an item of the targeted plan
// correction (planCorrections).

// PlanNamesRule is the rule the plan prompt, the plan check and the review
// state.
const PlanNamesRule = "Each technique has one entry: a name appears once in the repertoire, and a name the plan selects, as the base attack or a repertoire entry and so as a purchase's technique, is never also in omittedTechniques; an aspect of a selected technique that the Definition cannot express is an effect of its entry with adaptedAs empty and its reason, not an omission of the same name."

// duplicateEntry is a repertoire entry whose name an earlier entry has.
type duplicateEntry struct{ first, index int }

// duplicateRepertoireEntries lists each repertoire entry named exactly like
// an earlier one, with that earlier entry.
func duplicateRepertoireEntries(plan DesignPlan) []duplicateEntry {
	var out []duplicateEntry
	for index, entry := range plan.Repertoire {
		for first := 0; first < index; first++ {
			if sameTechnique(plan.Repertoire[first].Name, entry.Name) {
				out = append(out, duplicateEntry{first, index})
				break
			}
		}
	}
	return out
}

// selectedOmission is an omitted technique named exactly like a repertoire
// entry, or like the base attack when no entry has its name (entry -1).
type selectedOmission struct {
	omission, entry int
	// used is whether the base attack or a purchase adapts it.
	used bool
}

// selectedOmissions lists each omitted technique that the plan also
// selects. An unused repertoire entry that omittedTechniques omits too is
// left to RepertoireUseIssues, and a core entry omitted whole to the core
// concept check, when those checks apply, since each already reports it.
func selectedOmissions(plan DesignPlan, request *Request) []selectedOmission {
	unused := map[int]bool{}
	coreChecked := false
	if d := request.MechanicsDefinition; d != nil && d.Profile.DesignPolicy != nil {
		for _, index := range unusedRepertoireEntries(plan, request) {
			unused[index] = true
		}
		coreChecked = coreConceptsOn(*d) && request.SourceTechniques != nil
	}
	named := func(name string) bool {
		if plan.UpgradeIntents == nil {
			return false
		}
		for pathIndex := range m.PathKeys {
			for tier := 1; tier <= len(m.TierKeys); tier++ {
				if sameTechnique(plan.UpgradeIntents.At(pathIndex).At(tier).Technique, name) {
					return true
				}
			}
		}
		return false
	}
	var out []selectedOmission
	for index, omission := range plan.OmittedTechniques {
		entryIndex := slices.IndexFunc(plan.Repertoire, func(entry PlanRepertoire) bool { return sameTechnique(entry.Name, omission.Name) })
		switch {
		case entryIndex >= 0:
			entry := plan.Repertoire[entryIndex]
			if !unused[entryIndex] && !(coreChecked && entry.Importance == "core") {
				out = append(out, selectedOmission{index, entryIndex, sameTechnique(entry.Name, plan.Base.Name) || named(entry.Name)})
			}
		case sameTechnique(omission.Name, plan.Base.Name):
			out = append(out, selectedOmission{index, -1, true})
		}
	}
	return out
}

// PlanContradictionIssues reports each repertoire entry named exactly like
// an earlier one and each omitted technique named exactly like the base
// attack or a repertoire entry. A targeted plan correction can fix both
// (planCorrections).
func PlanContradictionIssues(plan DesignPlan, request *Request) []m.Issue {
	var issues []m.Issue
	for _, d := range duplicateRepertoireEntries(plan) {
		name := plan.Repertoire[d.index].Name
		issues = append(issues, m.Issue{
			Path:    fmt.Sprintf("repertoire.%d", d.index),
			Message: fmt.Sprintf("The repertoire lists %q twice, as repertoire.%d and repertoire.%d. %s Keep one entry named %q, with the sourceIds and effects of both, and remove the other.", name, d.first, d.index, PlanNamesRule, name),
		})
	}
	for _, o := range selectedOmissions(plan, request) {
		name := plan.OmittedTechniques[o.omission].Name
		issues = append(issues, m.Issue{
			Path:    fmt.Sprintf("omittedTechniques.%d", o.omission),
			Message: fmt.Sprintf("omittedTechniques omits %q, which the plan also selects as %s. %s %s", name, selectedAs(plan, o), PlanNamesRule, omissionFix(o)),
		})
	}
	return issues
}

// selectedAs says how the plan selects an omitted name.
func selectedAs(plan DesignPlan, o selectedOmission) string {
	if o.entry < 0 {
		return "the base attack"
	}
	var codes []string
	if plan.UpgradeIntents != nil {
		for pathIndex := range m.PathKeys {
			for tier := 1; tier <= len(m.TierKeys); tier++ {
				if sameTechnique(plan.UpgradeIntents.At(pathIndex).At(tier).Technique, plan.Repertoire[o.entry].Name) {
					codes = append(codes, BuildCode(pathIndex, tier))
				}
			}
		}
	}
	if len(codes) == 0 {
		return fmt.Sprintf("repertoire.%d", o.entry)
	}
	return fmt.Sprintf("repertoire.%d, the technique of %s", o.entry, joinWith(codes, "and", ""))
}

// baseOmissionFix is the fix of an omitted name the base attack has.
const baseOmissionFix = "Remove it from omittedTechniques, since 0-0-0 adapts it as the base attack, and state each aspect the Definition cannot express in scopeLimits."

// omissionFix is the fix of an omitted name the plan selects: a used one
// leaves omittedTechniques, and an unused one may instead leave the
// repertoire.
func omissionFix(o selectedOmission) string {
	if o.entry < 0 {
		return baseOmissionFix
	}
	keep := "Remove it from omittedTechniques and name each aspect the Definition cannot express as an effect of its entry with adaptedAs empty and its reason."
	if o.used {
		return keep
	}
	return keep + " Or, since no purchase adapts it, remove its repertoire entry and keep the omission."
}

// contradictionItems are the targeted-correction items of the plan
// contradictions, one per issue of PlanContradictionIssues.
func contradictionItems(plan DesignPlan, request *Request) []any {
	items := []any{}
	for _, d := range duplicateRepertoireEntries(plan) {
		name := plan.Repertoire[d.index].Name
		items = append(items, s.NewObject().
			Set("path", fmt.Sprintf("repertoire.%d", d.index)).
			Set("problem", fmt.Sprintf("The repertoire lists %q twice, as repertoire.%d and repertoire.%d. %s", name, d.first, d.index, PlanNamesRule)).
			Set("allowed", []any{fmt.Sprintf("Keep one entry named %q, with the sourceIds and effects of both, and remove the other; each milestone whose technique is %q keeps it.", name, name)}))
	}
	for _, o := range selectedOmissions(plan, request) {
		name := plan.OmittedTechniques[o.omission].Name
		allowed := []any{fmt.Sprintf("Remove %q from omittedTechniques and name each aspect the Definition cannot express as an effect of its entry with adaptedAs empty and its reason; every other omission stays.", name)}
		if o.entry < 0 {
			allowed = []any{baseOmissionFix}
		} else if !o.used {
			allowed = append(allowed, fmt.Sprintf("Remove the repertoire entry %q, which no purchase adapts, and keep its omission.", name))
		}
		items = append(items, s.NewObject().
			Set("path", fmt.Sprintf("omittedTechniques.%d", o.omission)).
			Set("problem", fmt.Sprintf("omittedTechniques omits %q, which the plan also selects as %s. %s", name, selectedAs(plan, o), PlanNamesRule)).
			Set("allowed", allowed))
	}
	return items
}
