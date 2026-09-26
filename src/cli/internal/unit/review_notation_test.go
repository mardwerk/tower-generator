package unit_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Regressions for #34. A correction treats a flagged finding by its class:
//   (a) malformed notation such as 5-x-2, or another path's purchase, with
//       no fact disproven: it must return under the same ID with only its
//       flagged build codes replaced, or the review is rejected; when only
//       notation was wrong, the summary may change only at its flagged codes
//       and no finding may be added (SOL-34-01, SOL-34-04, SOL-34-05);
//   (b) an impossible numeric build such as 3-3-0: fixed or withdrawn, and
//       never published (SOL-34-03);
//   (c) a wrong or unverifiable fact: may be withdrawn (SOL-34-02).

// notationReview is the 13c shape on the scripted fixture: the recorded
// finding whose citations hold, then extra findings in order.
func notationReview(t *testing.T, summary string, extra ...*s.Object) (*s.Object, *s.Object) {
	t.Helper()
	review := recordedOutput(t, "review")
	findings, _ := review.Get("findings")
	kept := findings.([]any)[0].(*s.Object)
	list := []any{kept}
	for _, f := range extra {
		list = append(list, f)
	}
	return review.Set("summary", summary).Set("findings", list), kept
}

func factList(build, field, value string) []any {
	return []any{s.NewObject().Set("build", build).Set("field", field).Set("value", value)}
}

// snakeFinding is model.snake-crosspath from run 13c (evidence 003-output),
// class (a): its action writes 5-x-2, neither a purchase nor a build.
func snakeFinding(kept *s.Object, action string) *s.Object {
	return s.Clone(kept).(*s.Object).
		Set("id", "model.snake-crosspath").
		Set("outcome", "fail").
		Set("subject", "path3 x-2-x").
		Set("message", "At 0-5-2, the x-2-x side purchase adds greater time-averaged direct and group damage than x-x-5's own capstone gain for only 240 Gold, while x-x-5 costs 14000 Gold.").
		Set("action", action).
		Set("facts", []any{})
}

const malformedAction = "Reassess path3 5-x-2's payoff or cost relative to x-2-x."

// fixedAction replaces the code and nothing else. x-x-5 is what the message
// compares; code cannot tell it from another legal purchase such as x-5-x.
const fixedAction = "Reassess path3 x-x-5's payoff or cost relative to x-2-x."

// crosspathFinding is class (b): its text names 3-3-0, which this
// Definition does not allow, while its structured fact is correct.
func crosspathFinding(kept *s.Object, build string) *s.Object {
	return s.Clone(kept).(*s.Object).
		Set("id", "model.crosspath").
		Set("outcome", "fail").
		Set("subject", "path1 tier 3").
		Set("message", "At "+build+", pierce stays at 5 while the price climbs.").
		Set("facts", factList("2-0-0", "attack.pierce", "5"))
}

// misreadFinding is class (c): a structured fact code contradicts.
func misreadFinding(kept *s.Object) *s.Object {
	return s.Clone(kept).(*s.Object).Set("id", "model.interval").Set("facts", factList("2-0-0", "attack.pierce", "3"))
}

func review(t *testing.T, checked unit.Checked, outputs ...any) (unit.Result, error, *fixture.Model) {
	t.Helper()
	model := &fixture.Model{Outputs: outputs}
	result, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options())
	return result, err, model
}

// rejected asserts a review-stage rejection whose message holds every want
// and none of the unwanted IDs, after the review and one correction, with
// no Result.
func rejected(t *testing.T, name string, result unit.Result, err error, calls int, want []string, unwanted ...string) {
	t.Helper()
	var failure *unit.ModelError
	if !errors.As(err, &failure) || failure.Failure == nil || failure.Failure.Code != unit.CodeOutputInvalid || failure.Failure.Stage != "review" || !strings.Contains(failure.Message, "The draft is retained") {
		t.Errorf("%s: want a rejected review, got %v", name, err)
	}
	for _, w := range want {
		if failure == nil || !strings.Contains(failure.Message, w) {
			t.Errorf("%s: the error lacks %q: %v", name, w, err)
		}
	}
	for _, other := range unwanted {
		if failure != nil && strings.Contains(failure.Message, other) {
			t.Errorf("%s: the error also names %s, which may be withdrawn", name, other)
		}
	}
	if calls != 2 {
		t.Errorf("%s: %d calls, want the review and one correction", name, calls)
	}
	if result.ID != "" || result.ReviewSummary != "" || len(result.Findings) != 0 {
		t.Errorf("%s: a rejected correction published a Result: %q", name, result.ReviewSummary)
	}
}

func accepted(t *testing.T, name string, result unit.Result, err error, calls int, ids, summary string) {
	t.Helper()
	if err != nil || calls != 2 {
		t.Errorf("%s: want a published Result after 2 calls, got %v after %d calls", name, err, calls)
		return
	}
	if got := modelIDs(result); got != ids || result.ReviewSummary != summary {
		t.Errorf("%s: model findings %s, summary %q", name, got, result.ReviewSummary)
	}
}

func contains(t *testing.T, name, prompt string, want, unwanted []string) {
	t.Helper()
	for _, w := range want {
		if !strings.Contains(prompt, w) {
			t.Errorf("%s: the correction lacks %q", name, w)
		}
	}
	for _, u := range unwanted {
		if strings.Contains(prompt, u) {
			t.Errorf("%s: the correction still says %q", name, u)
		}
	}
}

func modelIDs(result unit.Result) string {
	var ids []string
	for _, f := range result.Findings {
		if strings.HasPrefix(f.ID, "model.") {
			ids = append(ids, f.ID)
		}
	}
	return strings.Join(ids, ",")
}

var notationRejected = []string{"model.snake-crosspath", "notation"}

// 1. Run 13c, class (a): the correction omits the finding and says no
// finding is made about an illegal crosspath. Rejected, not published.
func TestReviewCorrectionCannotWithdrawMalformedNotation(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	first, _ := notationReview(t, "Found a capstone payoff concern.", snakeFinding(kept, malformedAction))
	omitted, _ := notationReview(t, "No finding is made about an illegal Snakeman crosspath.")
	result, err, model := review(t, stages.Checked, first, omitted)
	rejected(t, "13c omission", result, err, len(model.Requests), notationRejected)
	contains(t, "13c", model.Requests[1].Prompt,
		[]string{"5-x-2 is neither a purchase nor a build", "Return model.snake-crosspath under the same ID with its build codes corrected"},
		[]string{"not legal under this Definition: 5-x-2", "or omit it when the resolved facts contradict its claim"})

	// A path mismatch is notation too: subject x-5-x, text 5-x-x.
	swapped := s.Clone(kept).(*s.Object).Set("id", "model.swapped").Set("subject", "x-5-x").Set("message", "5-x-x adds little over its price.")
	first, _ = notationReview(t, "One payoff concern.", swapped)
	result, err, model = review(t, stages.Checked, first, omitted)
	rejected(t, "mismatch omission", result, err, len(model.Requests), []string{"model.swapped", "notation"})
}

// 2. Class (a), corrected under the same ID: published with the
// correction's summary. Returning it with its outcome or facts changed is a
// withdrawal in disguise and is rejected.
func TestReviewCorrectionPublishesCorrectedNotation(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	first, _ := notationReview(t, "Found a capstone payoff concern.", snakeFinding(kept, malformedAction).Set("facts", factList("2-0-0", "attack.pierce", "5")))
	summary := "Found a capstone payoff concern."
	fixed := snakeFinding(kept, fixedAction).Set("facts", factList("2-0-0", "attack.pierce", "5"))
	corrected, _ := notationReview(t, summary, fixed)
	result, err, model := review(t, stages.Checked, first, corrected)
	accepted(t, "corrected notation", result, err, len(model.Requests), "model.fan-club-allies,model.snake-crosspath", summary)
	last := result.Findings[len(result.Findings)-1]
	if last.Action == nil || *last.Action != fixedAction || last.Outcome != "fail" || len(last.Facts) != 1 || last.Facts[0].Value != "5" {
		t.Errorf("the published finding %+v", last)
	}
	for name, change := range map[string]func(*s.Object){
		"outcome pass":  func(f *s.Object) { f.Set("outcome", "pass") },
		"facts dropped": func(f *s.Object) { f.Set("facts", []any{}) },
	} {
		returned := s.Clone(fixed).(*s.Object)
		change(returned)
		correction, _ := notationReview(t, "No concrete issue.", returned)
		result, err, model := review(t, stages.Checked, first, correction)
		rejected(t, name, result, err, len(model.Requests), notationRejected)
	}
}

// 2b. SOL-34-04, class (a): the ID, outcome and empty facts are kept, but the
// payoff claim becomes "No payoff concern" and the summary "No concrete
// issue". Only build codes may change, so the review is rejected, nothing is
// published and Author keeps the checked draft. So are a code fix with a
// substantive rewording, a dropped word beside the code, a changed summary
// and an added finding. Replacing only 5-x-2, in the finding and in the
// summary, is published.
func TestReviewCorrectionCannotRewriteANotationOnlyFinding(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	snake := snakeFinding(kept, malformedAction)
	first, _ := notationReview(t, "Found a capstone payoff concern.", snake)
	noClaim := snakeFinding(kept, fixedAction).Set("message", "No payoff concern.")
	sol, _ := notationReview(t, "No concrete issue.", noClaim)
	result, err, model := review(t, stages.Checked, first, sol)
	rejected(t, "SOL-34-04", result, err, len(model.Requests), []string{"model.snake-crosspath", "notation", "changed its summary"})
	contains(t, "SOL-34-04", model.Requests[1].Prompt,
		[]string{"Replace only the codes named above as notation errors or as another path's purchase, each with one code, and keep every other build code and word exactly", "return the summary as before with only the codes named above corrected, and add no finding"}, nil)
	retained(t, stages, first, sol, "model.snake-crosspath")

	reworded := snakeFinding(kept, fixedAction).Set("message", strings.Replace(must(snake.Get("message")).(string), "adds greater", "adds similar", 1))
	extra := s.Clone(kept).(*s.Object).Set("id", "model.extra").Set("message", "Another concern.")
	for name, correction := range map[string][]any{
		"claim only":       {"Found a capstone payoff concern.", noClaim},
		"summary only":     {"No concrete issue.", snakeFinding(kept, fixedAction)},
		"code and wording": {"Found a capstone payoff concern.", reworded},
		"word beside code": {"Found a capstone payoff concern.", snakeFinding(kept, "Reassess x-x-5's payoff or cost relative to x-2-x.")},
		"finding added":    {"Found a capstone payoff concern.", snakeFinding(kept, fixedAction), extra},
	} {
		var findings []*s.Object
		for _, f := range correction[1:] {
			findings = append(findings, f.(*s.Object))
		}
		returned, _ := notationReview(t, correction[0].(string), findings...)
		result, err, model := review(t, stages.Checked, first, returned)
		rejected(t, name, result, err, len(model.Requests), []string{"build codes"})
	}

	for name, code := range map[string]string{"x-x-5": "x-x-5", "x-5-x, also a legal purchase": "x-5-x"} {
		action := strings.Replace(malformedAction, "5-x-2", code, 1)
		corrected, _ := notationReview(t, "Found a capstone payoff concern.", snakeFinding(kept, action))
		result, err, model := review(t, stages.Checked, first, corrected)
		accepted(t, name, result, err, len(model.Requests), "model.fan-club-allies,model.snake-crosspath", "Found a capstone payoff concern.")
		if err == nil {
			last := result.Findings[len(result.Findings)-1]
			if *last.Action != action || last.Message != must(snake.Get("message")).(string) || last.Outcome != "fail" || len(last.Facts) != 0 {
				t.Errorf("%s: the published finding %+v", name, last)
			}
		}
	}
	inSummary, _ := notationReview(t, "Found a capstone payoff concern at 5-x-2.", snake)
	fixedSummary, _ := notationReview(t, "Found a capstone payoff concern at x-x-5.", snakeFinding(kept, fixedAction))
	result, err, model = review(t, stages.Checked, inSummary, fixedSummary)
	accepted(t, "code in the summary", result, err, len(model.Requests), "model.fan-club-allies,model.snake-crosspath", "Found a capstone payoff concern at x-x-5.")
}

// 2c. SOL-34-05, class (a): replacement is allowed only where the first
// review wrote a flagged code. A correction that fixes 5-x-2 and also turns
// the valid comparison x-2-x into another legal purchase, with every word and
// the ID unchanged, alters the claim and is rejected; so is a summary that
// changes a code nobody flagged. Fixing a mismatch in the subject or in the
// text is published.
func TestReviewCorrectionReplacesOnlyFlaggedCodes(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	snake := snakeFinding(kept, malformedAction)
	message := must(snake.Get("message")).(string)
	first, _ := notationReview(t, "Found a capstone payoff concern.", snake)
	for name, returned := range map[string]*s.Object{
		"comparison in the action":  snakeFinding(kept, "Reassess path3 x-x-5's payoff or cost relative to x-3-x."),
		"comparison in the message": snakeFinding(kept, fixedAction).Set("message", strings.Replace(message, "the x-2-x side purchase", "the x-3-x side purchase", 1)),
		"build in the message":      snakeFinding(kept, fixedAction).Set("message", strings.Replace(message, "At 0-5-2", "At 0-5-1", 1)),
		"code in the subject":       snakeFinding(kept, fixedAction).Set("subject", "path3 x-3-x"),
	} {
		corrected, _ := notationReview(t, "Found a capstone payoff concern.", returned)
		result, err, model := review(t, stages.Checked, first, corrected)
		rejected(t, name, result, err, len(model.Requests), []string{"model.snake-crosspath", "notation", "only the build codes code flagged replaced"})
	}

	for name, summaries := range map[string][2]string{
		"unflagged code beside a flagged one": {"Found a capstone payoff concern at 5-x-2 against x-2-x.", "Found a capstone payoff concern at x-x-5 against x-3-x."},
		"no flagged code in the summary":      {"Found a capstone payoff concern at x-2-x.", "Found a capstone payoff concern at x-3-x."},
	} {
		first, _ := notationReview(t, summaries[0], snake)
		corrected, _ := notationReview(t, summaries[1], snakeFinding(kept, fixedAction))
		result, err, model := review(t, stages.Checked, first, corrected)
		rejected(t, name, result, err, len(model.Requests), []string{"changed its summary", "only the build codes code flagged in it replaced"}, "model.snake-crosspath")
	}

	// A mismatch flags the subject's purchase and the text's: either may be
	// the one corrected.
	swapped := s.Clone(kept).(*s.Object).Set("id", "model.swapped").Set("subject", "x-5-x").Set("message", "5-x-x adds little over x-2-x.")
	first, _ = notationReview(t, "One payoff concern.", swapped)
	for name, fixed := range map[string]*s.Object{
		"subject corrected": s.Clone(swapped).(*s.Object).Set("subject", "5-x-x"),
		"text corrected":    s.Clone(swapped).(*s.Object).Set("message", "x-5-x adds little over x-2-x."),
	} {
		corrected, _ := notationReview(t, "One payoff concern.", fixed)
		result, err, model := review(t, stages.Checked, first, corrected)
		accepted(t, name, result, err, len(model.Requests), "model.fan-club-allies,model.swapped", "One payoff concern.")
	}
	moved, _ := notationReview(t, "One payoff concern.", s.Clone(swapped).(*s.Object).Set("message", "x-5-x adds little over x-3-x."))
	result, err, model := review(t, stages.Checked, first, moved)
	rejected(t, "mismatch fixed with a comparison changed", result, err, len(model.Requests), []string{"model.swapped", "notation"})
}

// retained asserts that Author, given this review and correction, fails at
// review with the checked draft kept and names the finding.
func retained(t *testing.T, stages fixture.Stages, first, correction *s.Object, id string) {
	t.Helper()
	request, err := fixture.Request()
	if err != nil {
		t.Fatal(err)
	}
	model := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), recordedOutput(t, "mechanics"), first, correction}}
	result, err := unit.Author(context.Background(), s.FromGoValue(unit.ApplyProfile(request, unit.DefaultProfile())), model, fixture.Options())
	var failure *unit.ModelError
	if !errors.As(err, &failure) || failure.Failure == nil || failure.Failure.Stage != "review" || !strings.Contains(failure.Message, id) || len(model.Requests) != 4 || result.ID != "" {
		t.Errorf("author: %v after %d calls", err, len(model.Requests))
		return
	}
	if failure.Checked == nil {
		t.Error("the review failure does not carry the checked draft")
		return
	}
	checked := *failure.Checked
	if s.Stringify(s.FromGoValue(checked)) != s.Stringify(s.FromGoValue(stages.Checked)) {
		t.Error("the retained draft is not the checked draft")
	}
	if _, err := unit.ReviewDraft(context.Background(), checked, &fixture.Model{Outputs: []any{recordedOutput(t, "review")}}, fixture.Options()); err != nil {
		t.Errorf("the retained draft cannot be reviewed again: %v", err)
	}
}

// 3. Class (b), fixed: 3-3-0 replaced by the legal 3-2-0 and the finding
// revised, under its ID or under a new one. Published.
func TestReviewCorrectionMayFixAnImpossibleBuild(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	first, _ := notationReview(t, "A crosspath concern at 3-3-0.", crosspathFinding(kept, "3-3-0"))
	fixed := crosspathFinding(kept, "3-2-0").Set("severity", "info")
	for name, returned := range map[string]*s.Object{"same ID": fixed, "new ID": s.Clone(fixed).(*s.Object).Set("id", "model.crosspath-legal")} {
		id := must(returned.Get("id")).(string)
		corrected, _ := notationReview(t, "A crosspath concern at 3-2-0.", returned)
		result, err, model := review(t, stages.Checked, first, corrected)
		accepted(t, name, result, err, len(model.Requests), "model.fan-club-allies,"+id, "A crosspath concern at 3-2-0.")
		contains(t, name, model.Requests[1].Prompt,
			[]string{"It cites builds that are not legal under this Definition: 3-3-0.", "Replace each build legalBuilds does not list in model.crosspath with a legal build and revise the finding to match that build's resolved facts, or omit the finding."},
			[]string{"Return model.crosspath under the same ID with its build codes corrected", "Write a purchase as 3-x-x"})
	}
}

// 4. Class (b), withdrawn, even though its structured fact is correct, and
// even when it also has malformed notation: 3-3-0 outranks 5-x-2.
func TestReviewCorrectionMayWithdrawAnImpossibleBuild(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	both := crosspathFinding(kept, "3-3-0").Set("action", malformedAction)
	for name, flagged := range map[string]*s.Object{"3-3-0": crosspathFinding(kept, "3-3-0"), "3-3-0 and 5-x-2": both} {
		first, _ := notationReview(t, "A crosspath concern.", flagged)
		withdrawn, _ := notationReview(t, "One scope concern.")
		result, err, model := review(t, stages.Checked, first, withdrawn)
		accepted(t, name, result, err, len(model.Requests), "model.fan-club-allies", "One scope concern.")
	}
}

// 5. Class (b), left in: never published. In the text after a correction
// (rejected as today); in the rule, which main never reads (corrected, then
// rejected); in a new finding ID (rejected by the publish guard).
func TestReviewNeverPublishesAnImpossibleBuild(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	first, _ := notationReview(t, "A crosspath concern.", crosspathFinding(kept, "3-3-0"))
	left, _ := notationReview(t, "A crosspath concern.", crosspathFinding(kept, "3-3-0"))
	result, err, model := review(t, stages.Checked, first, left)
	rejected(t, "left in the text", result, err, len(model.Requests), []string{"cited builds that are not legal (3-3-0) after one correction"})

	ruled := crosspathFinding(kept, "3-2-0").Set("rule", "crosspath 3-3-0")
	first, _ = notationReview(t, "A crosspath concern.", ruled)
	result, err, model = review(t, stages.Checked, first, first)
	rejected(t, "left in the rule", result, err, len(model.Requests), []string{"3-3-0"})
	fixedRule, _ := notationReview(t, "A crosspath concern.", s.Clone(ruled).(*s.Object).Set("rule", "crosspath"))
	result, err, model = review(t, stages.Checked, first, fixedRule)
	accepted(t, "rule corrected", result, err, len(model.Requests), "model.fan-club-allies,model.crosspath", "A crosspath concern.")

	first, _ = notationReview(t, "A crosspath concern.", crosspathFinding(kept, "3-3-0"))
	renamed, _ := notationReview(t, "A crosspath concern.", crosspathFinding(kept, "3-2-0").Set("id", "model.3-3-0-crosspath"))
	result, err, model = review(t, stages.Checked, first, renamed)
	rejected(t, "moved into an ID", result, err, len(model.Requests), []string{"3-3-0", "findings it would publish"})
}

// 6. Class (c): a wrong value, or a build or field legalBuilds does not
// list, is not a verified fact. The finding may be withdrawn, even with a
// notation error.
func TestReviewCorrectionMayWithdrawAnUnverifiableFact(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	for name, flagged := range map[string]*s.Object{
		"wrong value":           misreadFinding(kept),
		"wrong value and 5-x-2": snakeFinding(kept, malformedAction).Set("facts", factList("2-0-0", "attack.pierce", "3")),
		"unlisted build":        snakeFinding(kept, malformedAction).Set("facts", factList("3-3-0", "attack.pierce", "5")),
		"unlisted field":        snakeFinding(kept, malformedAction).Set("facts", factList("2-0-0", "attack.width", "3")),
	} {
		first, _ := notationReview(t, "Two concerns.", flagged)
		withdrawn, _ := notationReview(t, "One scope concern.")
		result, err, model := review(t, stages.Checked, first, withdrawn)
		accepted(t, name, result, err, len(model.Requests), "model.fan-club-allies", "One scope concern.")
		id := must(flagged.Get("id")).(string)
		contains(t, name, model.Requests[1].Prompt,
			[]string{"Some cited facts are wrong or cannot be checked:", "Return " + id + " corrected under the same ID, or omit it when the resolved facts contradict its claim or do not list what it cites."},
			[]string{"Return " + id + " under the same ID with its build codes corrected"})
	}
}

// 7. Mixed: one finding of each class. Fixing (a) and (b) and withdrawing
// (c) is published; withdrawing (b) and (c) but not (a) is published;
// withdrawing (a) is rejected naming only it; leaving 3-3-0 is rejected.
func TestReviewCorrectionMixedFlaggedFindings(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	first, _ := notationReview(t, "Four concerns.", snakeFinding(kept, malformedAction), crosspathFinding(kept, "3-3-0"), misreadFinding(kept))

	fixed, _ := notationReview(t, "Three concerns.", snakeFinding(kept, fixedAction), crosspathFinding(kept, "3-2-0"))
	result, err, model := review(t, stages.Checked, first, fixed)
	accepted(t, "a and b fixed, c withdrawn", result, err, len(model.Requests), "model.fan-club-allies,model.snake-crosspath,model.crosspath", "Three concerns.")
	contains(t, "mixed", model.Requests[1].Prompt, []string{
		"Return model.snake-crosspath under the same ID with its build codes corrected",
		"Replace each build legalBuilds does not list in model.crosspath with a legal build",
		"Return model.interval corrected under the same ID, or omit it when the resolved facts contradict its claim",
	}, []string{"Return model.snake-crosspath, "})

	onlyA, _ := notationReview(t, "Two concerns.", snakeFinding(kept, fixedAction))
	result, err, model = review(t, stages.Checked, first, onlyA)
	accepted(t, "b and c withdrawn", result, err, len(model.Requests), "model.fan-club-allies,model.snake-crosspath", "Two concerns.")

	noA, _ := notationReview(t, "Two concerns.", crosspathFinding(kept, "3-2-0"), misreadFinding(kept).Set("facts", factList("2-0-0", "attack.pierce", "5")))
	result, err, model = review(t, stages.Checked, first, noA)
	rejected(t, "a withdrawn", result, err, len(model.Requests), notationRejected, "model.crosspath", "model.interval")

	leftB, _ := notationReview(t, "Two concerns.", snakeFinding(kept, fixedAction), crosspathFinding(kept, "3-3-0"))
	result, err, model = review(t, stages.Checked, first, leftB)
	rejected(t, "3-3-0 left", result, err, len(model.Requests), []string{"3-3-0"})
}

// 8. The draft is retained: Author returns the checked draft with a
// review-stage failure, so the CLI can write it for another review.
func TestAuthorKeepsTheCheckedDraftOfARejectedReview(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	_, kept := notationReview(t, "")
	first, _ := notationReview(t, "Found a capstone payoff concern.", snakeFinding(kept, malformedAction))
	omitted, _ := notationReview(t, "No finding is made about an illegal Snakeman crosspath.")
	retained(t, stages, first, omitted, "model.snake-crosspath")
}

func must(value any, _ bool) any { return value }
