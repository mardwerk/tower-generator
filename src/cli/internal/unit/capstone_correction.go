package unit

import (
	"fmt"
	"slices"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Capstone correction (#61, SOL-61-07). Under requireTier5BehaviorChange a
// fifth purchase must add a supported behavior or access, or name a
// proposed mechanic. Escanor's v32 plan attempt 1 failed on six mixed
// issues, so no targeted correction applied; the full-plan retry left only
// 5-x-x with no new behavior, and the run failed (OPUS-NET-61-10). When a
// full-plan retry fails only on fifth purchases with no new behavior,
// together at most with items a targeted correction can fix, the next call
// is one capstone correction. It names each failing fifth purchase, its
// milestone and its technique's effects, and asks for the source effect the
// capstone adapts and the playable behavior it adds: a supported unlock
// that adds behavior or access, or, only when the Definition cannot express
// it, a proposed mechanic of that purchase. A capstone with no new behavior
// is a design choice as well as a missing field, so it does not ask for any
// unlock that passes the check.
//
// Code, not the model, keeps every other milestone: the corrected output is
// merged into the rejected one, taking only the named fifth milestones and
// their paths' capstoneValue, and the repertoire and omittedTechniques when
// the correction also names targeted-correction items. Every other change
// the model makes is discarded, not rejected, so a correction that also
// rewrites another milestone still gets its fifth purchase judged. The
// merged plan passes DecodeDesignPlan like any plan. It is at most one call
// per draft, sent only after a full-plan retry, so never with a retry budget
// of 0, and the attempt is recorded with purpose "capstone-correction", its
// usage and its issues.

// capstoneIssue is the issue code of a fifth purchase that promises no new
// behavior or access and names no proposed mechanic.
const capstoneIssue = "plan-capstone"

// CapstoneCorrectionPurpose is the attempt purpose of a capstone
// correction.
const CapstoneCorrectionPurpose = "capstone-correction"

// onlyCapstone reports issues a capstone correction can fix, all of them:
// at least one fifth purchase with no new behavior, and otherwise only items
// a targeted correction can fix.
func onlyCapstone(issues []s.Issue) bool {
	capstone := false
	for _, issue := range issues {
		switch issue.Code {
		case capstoneIssue:
			capstone = true
		case correctableIssue:
		default:
			return false
		}
	}
	return capstone
}

// capstoneCorrection is the capstone correction of one rejected plan
// output.
type capstoneCorrection struct {
	// paths are the indexes of the paths whose fifth milestone it may
	// change.
	paths []int
	// narrow is whether it also names targeted-correction items, so the
	// repertoire and omittedTechniques may change too.
	narrow bool
	// prompt is the addition to the planning prompt.
	prompt string
}

// capstoneCorrectionFor builds the capstone correction of a compact plan
// output that failed on issues, all of which onlyCapstone accepts. ok is
// false when the output is not a compact plan or names no fifth purchase.
func capstoneCorrectionFor(output any, request *Request, issues []s.Issue) (capstoneCorrection, bool) {
	previous, ok := output.(*s.Object)
	if !ok || !previous.Has("contract") {
		return capstoneCorrection{}, false
	}
	plan, err := parseDesignPlan(output, request)
	if err != nil || plan.UpgradeIntents == nil {
		return capstoneCorrection{}, false
	}
	var correction capstoneCorrection
	for _, issue := range issues {
		if issue.Code == correctableIssue {
			correction.narrow = true
			continue
		}
		if len(issue.Path) != 3 {
			continue
		}
		if index := slices.Index(m.PathKeys, fmt.Sprint(issue.Path[1])); index >= 0 && !slices.Contains(correction.paths, index) {
			correction.paths = append(correction.paths, index)
		}
	}
	if len(correction.paths) == 0 {
		return capstoneCorrection{}, false
	}
	slices.Sort(correction.paths)
	items := []any{}
	for _, index := range correction.paths {
		items = append(items, capstoneItem(plan, previous, index))
	}
	scope := "Change only the fifth milestone each capstone item names, and that path's capstoneValue to match it."
	if correction.narrow {
		items = append(items, correctionItems(plan, request, correction.paths)...)
		scope = "Change only the fifth milestone each capstone item names, and that path's capstoneValue to match it; correct each other item with one of its allowed corrections, in the repertoire or omittedTechniques or on a fifth milestone this correction changes."
	}
	correction.prompt = "\n\nThe previous plan failed only on the items below. Each capstone item is a fifth purchase that promises no new behavior or access and names no proposed mechanic. " + m.BehaviorChangeRule(5) +
		" A capstone with no new behavior is a design choice as well as a missing field, so do not add an unlock only to pass this check. For each capstone item, choose the source effect the capstone adapts, an effect of its technique listed with the item or described by the passages that technique cites, and the playable behavior that effect adds at this purchase. Write both in the milestone's change, and promise the behavior with one of the item's allowed corrections. " +
		scope + " Code keeps every other milestone and field exactly as it was and validates the complete plan again. Return the complete plan. " +
		s.Stringify(s.NewObject().Set("items", items).Set("previous", output))
	return correction, true
}

// capstoneItem names one failing fifth purchase: its milestone as the
// rejected output wrote it, the technique it adapts with that technique's
// effects, the unlocks its path already has and the corrections allowed.
func capstoneItem(plan DesignPlan, previous *s.Object, index int) *s.Object {
	key := m.PathKeys[index]
	code := BuildCode(index, 5)
	intents := plan.UpgradeIntents.At(index)
	intent := intents.At(5)
	item := s.NewObject().
		Set("path", "paths."+key+".milestones.tier5").
		Set("purchase", code).
		Set("pathName", plan.Paths.At(index).Name).
		Set("problem", fmt.Sprintf("%s promises no new behavior or access and names no proposed mechanic.", code)).
		Set("milestone", valueAt(previous, []any{"paths", key, "milestones", "tier5"})).
		Set("capstoneValue", plan.Paths.At(index).CapstoneValue)
	if sameTechnique(intent.Technique, plan.Base.Name) {
		item.Set("technique", s.NewObject().Set("name", plan.Base.Name).Set("baseAttack", true).Set("sourceIds", s.FromGoValue(plan.Base.SourceIDs)).Set("behavior", plan.Base.Behavior))
	} else {
		for _, entry := range plan.Repertoire {
			if sameTechnique(entry.Name, intent.Technique) {
				item.Set("technique", s.NewObject().Set("name", entry.Name).Set("sourceIds", s.FromGoValue(entry.SourceIDs)).Set("effects", s.FromGoValue(entry.Effects)))
				break
			}
		}
	}
	earlier := []any{}
	for tier := 1; tier <= 4; tier++ {
		if unlock := intents.At(tier).Unlock; unlock != "none" && !slices.Contains(earlier, any(unlock)) {
			earlier = append(earlier, unlock)
		}
	}
	return item.
		Set("pathUnlocks", earlier).
		Set("allowed", []any{
			fmt.Sprintf("Promise the behavior the chosen source effect adds as a supported unlock of %s that adds behavior or access: not none, not targeting-change and not an unlock in pathUnlocks, which the path already has; or promise projectiles while the path fires one projectile. Keep the milestone's other promises that still hold.", code),
			fmt.Sprintf("Only when the Definition cannot express the behavior the chosen source effect describes, name it in the proposedMechanics of %s: what it does in play, as a playable mechanic, and the sourceIds of the passages that describe it. No build grants it, so %s stays an unresolved Design gap until the Definition supports it.", code, code),
		})
}

// merge keeps every milestone of the rejected output except the fifth
// milestones the correction may change: it takes those, with their paths'
// capstoneValue, from the corrected output, and its repertoire and
// omittedTechniques when the correction also names targeted-correction
// items. Anything else the corrected output changes is discarded.
func (c capstoneCorrection) merge(previous, corrected any) any {
	merged, ok := s.Clone(previous).(*s.Object)
	out, isObject := corrected.(*s.Object)
	if !ok || !isObject {
		return corrected
	}
	for _, index := range c.paths {
		key := m.PathKeys[index]
		milestones, ok := valueAt(merged, []any{"paths", key, "milestones"}).(*s.Object)
		if !ok {
			continue
		}
		if milestone, ok := valueAt(out, []any{"paths", key, "milestones", "tier5"}).(*s.Object); ok {
			milestones.Set("tier5", s.Clone(milestone))
		}
		if value, ok := valueAt(out, []any{"paths", key, "capstoneValue"}).(string); ok {
			valueAt(merged, []any{"paths", key}).(*s.Object).Set("capstoneValue", value)
		}
	}
	if c.narrow {
		for _, key := range []string{"repertoire", "omittedTechniques"} {
			if value, ok := out.Get(key); ok {
				merged.Set(key, s.Clone(value))
			}
		}
	}
	return merged
}
