package unit_test

import (
	"math"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Regressions for #57, on the scripted Dart Monkey: a promise without a
// source effect is a review cue, an Active-only capstone is compared by its
// time-averaged rate, and two claims code can check against the plan and
// purchaseEvidence are corrected like a wrong fact, or, when code cannot read
// them, kept and marked for review.

// reviewTier is one purchase in the review context's unit.paths.
func reviewTier(context *s.Object, path, tier int) *s.Object {
	paths := at(context, "unit", "paths").([]any)
	return at(paths[path], "tiers").([]any)[tier-1].(*s.Object)
}

// Fix 1: each purchase whose technique is a repertoire entry lists the
// promises no effect of that technique adapts. 3-x-x promises range from
// Spiked Ball, whose effects adapt damage, pierce and a damage type change.
func TestReviewTiersListPromisesWithoutEffect(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prompt, context := reviewContext(t, stages.Checked)
	for _, test := range []struct {
		path, tier int
		want       any
	}{
		{0, 3, []any{"range"}},
		{2, 3, []any{}},
		{1, 5, []any{"active-damage", "active-duration"}},
		// A purchase that adapts the base attack has no cue.
		{0, 1, nil},
	} {
		got, ok := reviewTier(context, test.path, test.tier).Get("promisesWithoutEffect")
		if s.Stringify(s.FromGoValue(got)) != s.Stringify(s.FromGoValue(test.want)) || ok != (test.want != nil) {
			t.Errorf("%s: promisesWithoutEffect %v", unit.BuildCode(test.path, test.tier), got)
		}
	}
	if !strings.Contains(prompt, "A tier's promisesWithoutEffect lists what it promises that no effect of its own technique adapts: fail an adaptation or name that credits the technique with such a promise, unless a passage its sourceIds cite describes it.") {
		t.Error("the review is not told how to read promisesWithoutEffect")
	}
	// It is a cue, not a gate: the plan and check are unchanged.
	if issues := unit.PlanEffectIssues(*stages.Checked.Draft.Run.DesignPlan); len(issues) > 0 {
		t.Errorf("the plan check rejects the fixture: %v", issues)
	}
}

// Fix 4: an Active-only capstone scores 1/count on ordinary rates, like one
// that adds nothing; its time-averaged ordering shows what it adds.
func TestCapstoneOrderingComparesActiveOnlyCapstonesOverTime(t *testing.T) {
	blueprint := burstUnit()
	rows := unit.CapstoneOrdering(&blueprint)
	top, middle := rows[0].(*s.Object), rows[1].(*s.Object)
	// x-5-x "Long Burst" only lengthens the Active: 4 ordinary, 8
	// time-averaged, against 7 copies of x-4-x at 4 and 6.
	if at(middle, "activeOnly") != true || at(top, "activeOnly") != false {
		t.Errorf("activeOnly: 5-x-x %v, x-5-x %v", at(top, "activeOnly"), at(middle, "activeOnly"))
	}
	for key, want := range map[string]float64{"direct damage rate": 4.0 / 28, unit.TimeAveragedDirect: 8.0 / 42, unit.TimeAveragedGroup: 8.0 / 42} {
		ratio, _ := at(middle, key, "tier5ToCopiesRatio").(float64)
		if at(middle, key, "ordering") != "lower" || math.Abs(ratio-want) > 1e-12 {
			t.Errorf("x-5-x %s: %s", key, s.Stringify(at(middle, key)))
		}
	}
	// In the scripted review, x-5-x Plasma Monkey Fan Club changes only the
	// Active and 5-x-x and x-x-5 change the attack.
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prompt, context := reviewContext(t, stages.Checked)
	ordering := at(context, "purchaseComparisonOrdering").([]any)
	for index, want := range []bool{false, true, false} {
		if at(ordering[index], "activeOnly") != want || at(ordering[index], unit.TimeAveragedDirect, "ordering") == nil {
			t.Errorf("path %d: %s", index+1, s.Stringify(ordering[index]))
		}
	}
	if !strings.Contains(prompt, "Its activeOnly marks a fifth purchase that changes only the Active Ability, whose ordinary rates equal the fourth purchase's while every copy owns that Active: for such a capstone, state its time-averaged ordering against the copies and whether its typed changes deliver the plan's capstoneValue, and report the capstone as unresolved when they do not; this is a checked comparison, not a balance verdict.") {
		t.Error("the review is not told how to judge an Active-only capstone")
	}
}

// timingFinding faults 1-x-x, which adapts Dart Throw, for Spiked Ball,
// which its path adapts from 3-x-x.
func timingFinding(kept *s.Object, message string) *s.Object {
	return s.Clone(kept).(*s.Object).
		Set("id", "model.sharp-timing").
		Set("outcome", "fail").
		Set("subject", "1-x-x Sharp Shots").
		Set("message", message).
		Set("action", "Revise 1-x-x.").
		Set("facts", []any{})
}

const faultedTiming = "1-x-x Sharp Shots claims only a sharper Dart Throw, but this purchase begins the Spiked Ball path."

// Fix 2: a finding that faults 1-x-x for Spiked Ball goes through the
// correction like a wrong fact: withdrawn, it is published without it;
// returned unchanged, the review is rejected.
func TestReviewCorrectsAPurchaseFaultedForALaterTechnique(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	first, _ := notationReview(t, "One timing concern.", timingFinding(kept, faultedTiming))
	withdrawn, _ := notationReview(t, "No timing concern.")
	result, err, model := review(t, stages.Checked, first, withdrawn)
	accepted(t, "withdrawn", result, err, len(model.Requests), "model.fan-club-allies", "No timing concern.")
	contains(t, "timing", model.Requests[len(model.Requests)-1].Prompt, []string{
		"Some findings fault a purchase for a technique its path adapts only at a later purchase: model.sharp-timing faults 1-x-x for Spiked Ball, but 1-x-x adapts Dart Throw; Spiked Ball starts at 3-x-x.",
		"Return model.sharp-timing corrected under the same ID, or omit it",
	}, nil)
	result, err, model = review(t, stages.Checked, first, first)
	rejected(t, "returned unchanged", result, err, len(model.Requests), []string{"model.sharp-timing faults 1-x-x for Spiked Ball"})
	// A contextual mention is published as written, in one call.
	context, _ := notationReview(t, "One pierce concern.", timingFinding(kept, "1-x-x Sharp Shots adds pierce the Dart Throw passage does not describe. Spiked Ball replaces the dart at 3-x-x."))
	result, err, model = review(t, stages.Checked, context)
	if err != nil || len(model.Requests) != 1 || modelIDs(result) != "model.fan-club-allies,model.sharp-timing" {
		t.Errorf("contextual mention: %v after %d calls, %s", err, len(model.Requests), modelIDs(result))
	}
}

// Fix 3: a finding that says x-x-5's gain is smaller than the x-2-x side
// purchase's, where purchaseEvidence says the side gain is smaller, goes
// through the same correction.
func TestReviewCorrectsAReversedCapstoneComparison(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	reversed := s.Clone(kept).(*s.Object).
		Set("id", "model.crossbow-price").
		Set("outcome", "fail").
		Set("subject", "x-x-5 Crossbow Master").
		Set("message", "x-x-5 costs 21,500 Gold and adds 21.05 time-averaged direct damage rate over x-x-4. Its gain is smaller than the x-2-x side purchase's 10.53 gain, which costs 190 Gold.").
		Set("facts", []any{})
	first, _ := notationReview(t, "One price concern.", reversed)
	withdrawn, _ := notationReview(t, "No price concern.")
	result, err, model := review(t, stages.Checked, first, withdrawn)
	accepted(t, "withdrawn", result, err, len(model.Requests), "model.fan-club-allies", "No price concern.")
	contains(t, "comparison", model.Requests[len(model.Requests)-1].Prompt, []string{
		"Some findings state the opposite of a checked againstCapstone comparison in purchaseEvidence: model.crossbow-price says the x-2-x side purchase's 10.53 gain is at least x-x-5's, but sideGainAtLeastCapstone is false for the time-averaged direct damage rate: from 0-1-5 to 0-2-5, x-2-x adds 10.53 for 190 Gold and x-x-5 adds 21.05 for 21500 Gold.",
		"A side purchase that adds less absolute gain than the capstone is no evidence against it",
	}, nil)
	// The same gain quoted in the right direction is published as written.
	right := s.Clone(reversed).(*s.Object).Set("message", "x-x-5 costs 21,500 Gold and adds 21.05 time-averaged direct damage rate over x-x-4. The x-2-x side purchase's 10.53 gain, for 190 Gold, is smaller than the capstone's.")
	correct, _ := notationReview(t, "One price concern.", right)
	result, err, model = review(t, stages.Checked, correct)
	if err != nil || len(model.Requests) != 1 || modelIDs(result) != "model.fan-club-allies,model.crossbow-price" {
		t.Errorf("correct direction: %v after %d calls, %s", err, len(model.Requests), modelIDs(result))
	}
}

// A claim code finds but cannot read is not corrected: the Result keeps the
// finding as written and adds an unresolved Finding that asks for it to be
// checked.
func TestReviewRetainsAnUnreadClaimForReview(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	unread := timingFinding(kept, "1-x-x Sharp Shots feels unrelated to Spiked Ball.")
	first, _ := notationReview(t, "One timing concern.", unread)
	result, err, model := review(t, stages.Checked, first)
	if err != nil || len(model.Requests) != 1 {
		t.Fatalf("%v after %d calls", err, len(model.Requests))
	}
	if modelIDs(result) != "model.fan-club-allies,model.sharp-timing" || result.Findings[len(result.Findings)-2].Message != "1-x-x Sharp Shots feels unrelated to Spiked Ball." {
		t.Errorf("the finding was not kept as written: %s", modelIDs(result))
	}
	marked := result.Findings[len(result.Findings)-1]
	if marked.ID != "deterministic.review.sharp-timing" || marked.Method != "deterministic" || marked.Outcome != "unresolved" || marked.Rule != unit.HumanReviewRule ||
		!strings.Contains(marked.Message, "Code could not tell whether model.sharp-timing faults 1-x-x for Spiked Ball: 1-x-x adapts Dart Throw, and Spiked Ball starts at 3-x-x.") {
		t.Errorf("review Finding %+v", marked)
	}
	// The recorded review needs no mark.
	for _, f := range stages.Result.Findings {
		if f.Rule == unit.HumanReviewRule {
			t.Errorf("the scripted review is marked: %+v", f)
		}
	}
}
