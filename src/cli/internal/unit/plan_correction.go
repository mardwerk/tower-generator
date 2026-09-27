package unit

import (
	"fmt"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Targeted plan correction (#61, SOL-61-05). Five plan failures have a
// known, local fix: a source technique the plan lists nowhere, an effect
// adapted as a promise that no purchase of its technique makes, a
// repertoire entry that neither the base attack nor any purchase adapts
// (RepertoireUseRule, SOL-61-10), a repertoire name listed twice and an
// omitted name the plan also selects (PlanNamesRule, SOL-61-13).
// Escanor's v30 plans failed on these after the full-plan retry (escanor-m
// in OPUS-NET-61-6). When a plan attempt fails on these alone, the next
// call is a correction that names each failed item and the corrections it
// allows, instead of another full re-plan. It is at most one call per
// draft, it does not use the full-plan retry budget, and none is sent when
// that budget is 0. Its output passes DecodeDesignPlan like any plan; the
// attempt is recorded with purpose "plan-correction".

// correctableIssue is the issue code of a plan failure that a targeted
// correction can fix.
const correctableIssue = "plan-correctable"

// PlanCorrectionPurpose is the attempt purpose of a targeted plan
// correction.
const PlanCorrectionPurpose = "plan-correction"

// onlyCorrectable reports issues that a targeted correction can fix, all of
// them.
func onlyCorrectable(issues []s.Issue) bool {
	for _, issue := range issues {
		if issue.Code != correctableIssue {
			return false
		}
	}
	return len(issues) > 0
}

// planCorrections lists each failed item of a plan output that a targeted
// correction can fix, with what it may do about it: a source technique
// listed nowhere is listed with a rank, an unpromised adaptedAs gains the
// promise on a purchase of that technique or loses that adaptedAs, an
// unused repertoire entry moves to omittedTechniques with its reason or
// becomes a purchase's technique, a repertoire name listed twice keeps one
// entry, and an omitted name the plan selects leaves omittedTechniques.
func planCorrections(output any, request *Request) []any {
	plan, err := parseDesignPlan(output, request)
	if err != nil {
		return nil
	}
	return correctionItems(plan, request, nil)
}

// correctionItems are planCorrections of a parsed plan. With open paths, a
// capstone correction's, the only milestones the correction may change are
// those paths' fifth purchases, so an unpromised adaptedAs may gain its
// promise, and an unused entry become the technique, only on one of them.
func correctionItems(plan DesignPlan, request *Request, open []int) []any {
	items := []any{}
	for _, u := range unpromisedAdaptations(plan) {
		entry := plan.Repertoire[u.entry]
		remove := fmt.Sprintf("Remove %s from this adaptedAs and say in the effect's reason why it is not adapted.", u.promise)
		allowed := []any{
			fmt.Sprintf("Promise %s on a milestone whose technique is %q, in its improves, unlock or lowers as the promise requires, and keep that milestone's other promises.", u.promise, entry.Name),
			remove,
		}
		if open != nil {
			allowed = []any{}
			for _, index := range open {
				if sameTechnique(plan.UpgradeIntents.At(index).At(5).Technique, entry.Name) {
					allowed = append(allowed, fmt.Sprintf("Promise %s on %s, the fifth purchase this correction changes, whose technique is %q, as the promise requires.", u.promise, BuildCode(index, 5), entry.Name))
				}
			}
			allowed = append(allowed, remove)
		}
		items = append(items, s.NewObject().
			Set("path", u.path()).
			Set("problem", fmt.Sprintf("The effect %q of %q is adapted as %s, but no milestone whose technique is %q promises %s.", entry.Effects[u.effect].Effect, entry.Name, u.promise, entry.Name, u.promise)).
			Set("allowed", allowed))
	}
	if d := request.MechanicsDefinition; d != nil && d.Profile.DesignPolicy != nil {
		for _, index := range unusedRepertoireEntries(plan, request) {
			entry := plan.Repertoire[index]
			rank := ""
			if entry.Importance != "" {
				rank = ", ranked " + entry.Importance + " as it is now"
			}
			problem := fmt.Sprintf("The repertoire entry %q is neither the base attack nor the technique of any milestone.", entry.Name)
			omit := fmt.Sprintf("Move it from the repertoire to omittedTechniques by the name %q%s, with a reason that names the specific behavior the Definition cannot express. %s", entry.Name, rank, OmissionRule)
			drop := ""
			if alsoOmitted(plan, entry) {
				problem = fmt.Sprintf("The repertoire entry %q is neither the base attack nor the technique of any milestone, and omittedTechniques omits it too.", entry.Name)
				omit = fmt.Sprintf("Remove it from the repertoire and keep its entry in omittedTechniques, whose reason names the specific behavior the Definition cannot express. %s", OmissionRule)
				drop = ", and remove it from omittedTechniques"
			}
			allowed := []any{omit}
			if open == nil {
				allowed = append(allowed, fmt.Sprintf("Name %q as the technique of a milestone that adapts what it does, promising there what its effects are adapted as, and keep each path's own technique%s.", entry.Name, drop))
			}
			for _, pathIndex := range open {
				allowed = append(allowed, fmt.Sprintf("Name %q as the technique of %s, the fifth purchase this correction changes, when that purchase adapts what it does, promising there what its effects are adapted as%s.", entry.Name, BuildCode(pathIndex, 5), drop))
			}
			items = append(items, s.NewObject().
				Set("path", fmt.Sprintf("repertoire.%d", index)).
				Set("problem", problem+" "+RepertoireUseRule).
				Set("allowed", allowed))
		}
	}
	items = append(items, contradictionItems(plan, request)...)
	if request.SourceTechniques != nil && request.MechanicsDefinition != nil && coreConceptsOn(*request.MechanicsDefinition) {
		for _, technique := range unlistedSourceTechniques(plan, request) {
			rank := "core, major or minor"
			if technique.Salience == SalienceStrong {
				rank = "major or core, since its salience is strong"
			}
			names := fmt.Sprintf("%q", technique.Name)
			if len(technique.Aliases) > 0 {
				names += " or one of its aliases, " + joinWith(quoted(technique.Aliases), "or", "")
			}
			items = append(items, s.NewObject().
				Set("path", "repertoire").
				Set("problem", fmt.Sprintf("The source technique %q is listed neither in the repertoire nor in omittedTechniques.", technique.Name)).
				Set("sourceTechnique", s.FromGoValue(technique)).
				Set("allowed", []any{
					fmt.Sprintf("Add it to omittedTechniques by the name %s, ranked %s, with a reason that names the specific behavior the Definition cannot express. %s", names, rank, OmissionRule),
					fmt.Sprintf("Add it to the repertoire by the name %s, ranked %s, citing its passageIds, with its effects; an effect no milestone whose technique it is promises keeps adaptedAs empty.", names, rank),
				}))
		}
	}
	return items
}

// targetedPlanCorrection is the prompt addition of a targeted correction.
func targetedPlanCorrection(items []any, previous any) string {
	return "\n\nThe previous plan failed only on the items below. Correct each item with one of its allowed corrections and change nothing else: return the complete plan with every other field, entry and milestone exactly as it was. " +
		s.Stringify(s.NewObject().Set("items", items).Set("previous", previous))
}
