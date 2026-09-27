package unit

import (
	"fmt"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// Every repertoire entry is used (#61, SOL-61-10). The v34 Luffy plan ranked
// Gear 5 and all three kinds of Haki major in its Repertoire, but no purchase
// named any of them as its technique, so the plan had in effect omitted
// them, and no omission verdict judged them, because the verdicts cover only
// omittedTechniques (OPUS-NET-61-13). A Repertoire entry, whatever its rank,
// is therefore the base attack, by its name, or the technique of at least one
// purchase; one that neither adapts is a whole-technique omission and
// belongs in omittedTechniques with its reason, where the review records a
// verdict on it. Code checks only that a purchase names the entry: whether
// that purchase adapts what the entry does in play is the review's judgment.

// RepertoireUseRule is the rule the plan prompt, the plan check and the
// review state.
const RepertoireUseRule = "Every repertoire entry, core, major or minor, is used: it is named like the base attack, which 0-0-0 adapts, or it is the technique of at least one purchase; a technique that neither the base attack nor any purchase adapts is omitted whole, so it belongs in omittedTechniques with its reason, where the review judges the omission, and not in the repertoire."

// unusedRepertoireEntries lists the index of each repertoire entry that is
// not the base attack and that no purchase names as its technique. A core
// entry is left out under the core concept check, which reports it
// (coreConceptIssues). Plans without typed techniques predate the field and
// are not checked.
func unusedRepertoireEntries(plan DesignPlan, request *Request) []int {
	if plan.UpgradeIntents == nil {
		return nil
	}
	coreChecked := request.MechanicsDefinition != nil && coreConceptsOn(*request.MechanicsDefinition) && request.SourceTechniques != nil
	var unused []int
	for index, entry := range plan.Repertoire {
		if sameTechnique(entry.Name, plan.Base.Name) || (coreChecked && entry.Importance == "core") {
			continue
		}
		named := false
		for pathIndex := range m.PathKeys {
			for tier := 1; tier <= len(m.TierKeys); tier++ {
				technique := plan.UpgradeIntents.At(pathIndex).At(tier).Technique
				if technique == "" {
					return nil
				}
				named = named || sameTechnique(technique, entry.Name)
			}
		}
		if !named {
			unused = append(unused, index)
		}
	}
	return unused
}

// alsoOmitted reports a repertoire entry that omittedTechniques omits by
// the same name too, as six of the seven unused entries of Escanor's v34
// plan were.
func alsoOmitted(plan DesignPlan, entry PlanRepertoire) bool {
	for _, omission := range plan.OmittedTechniques {
		if sameTechnique(omission.Name, entry.Name) {
			return true
		}
	}
	return false
}

// RepertoireUseIssues reports each repertoire entry that neither the base
// attack nor any purchase adapts, one issue per entry. A targeted plan
// correction can fix it (planCorrections).
func RepertoireUseIssues(plan DesignPlan, request *Request) []m.Issue {
	var issues []m.Issue
	for _, index := range unusedRepertoireEntries(plan, request) {
		entry := plan.Repertoire[index]
		rank := ""
		if entry.Importance != "" {
			rank = ", ranked " + entry.Importance + ","
		}
		facts := fmt.Sprintf("The repertoire entry %q%s is neither the base attack nor the technique of any purchase.", entry.Name, rank)
		fix := "Move it to omittedTechniques with its rank and the reason no purchase adapts it, or name it as the technique of a purchase that adapts it."
		if alsoOmitted(plan, entry) {
			facts = fmt.Sprintf("The repertoire entry %q%s is neither the base attack nor the technique of any purchase, and omittedTechniques omits it too.", entry.Name, rank)
			fix = "Remove it from the repertoire and keep its omission, or name it as the technique of a purchase that adapts it and remove it from omittedTechniques."
		}
		issues = append(issues, m.Issue{Path: fmt.Sprintf("repertoire.%d", index), Message: facts + " " + RepertoireUseRule + " " + fix})
	}
	return issues
}
