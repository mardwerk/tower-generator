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

// verdictFindings are a Result's review verdicts, in recorded order.
func verdictFindings(result unit.Result) []unit.Finding {
	var out []unit.Finding
	for _, f := range result.Findings {
		if unit.IsReviewVerdict(f) {
			out = append(out, f)
		}
	}
	return out
}

// The review context lists every whole-technique omission and every third
// purchase to judge, the prompt asks for one verdict on each outside the
// eight findings, and the schema allows exactly those subjects (SOL-61-08:
// the v32 Luffy review spent its eight findings on purchases and judged no
// omission and no third purchase).
func TestReviewAsksForAVerdictPerOmissionAndThirdPurchase(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	request := unit.BlueprintReviewRequest(stages.Checked)
	for _, want := range []string{
		`"requiredVerdicts":{"omissions":[{"technique":"Allied Fan Club","importance":"minor"},{"technique":"Critical shots","importance":"minor"}],"thirdPurchases":[{"build":"3-x-x","path":"path1","name":"Spike-o-pult","technique":"Spiked Ball"},{"build":"x-3-x","path":"path2","name":"Triple Shot","technique":"Triple Throw"},{"build":"x-x-3","path":"path3","name":"Crossbow","technique":"Crossbow"}]}`,
		"Give at most eight useful findings; the verdicts on omissions and third purchases are apart from them.",
		"give one verdict on each subject requiredVerdicts lists, and no other",
		"could adapt the technique's central effect, as the omission rule says, not whether some aspect of it is unsupported",
		"whether its importance is plausible against the passages its sourceTechnique cites",
		"comparing it with its own path's first and second purchases and with the other paths' purchases",
		unit.OmissionRule,
	} {
		if !strings.Contains(request.Prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	schema := s.Stringify(request.Schema)
	for _, want := range []string{
		`"omissionVerdicts":{"minItems":2,"maxItems":2,"type":"array","items":{"type":"object","properties":{"technique":{"type":"string","enum":["Allied Fan Club","Critical shots"]}`,
		`"thirdPurchaseVerdicts":{"minItems":3,"maxItems":3,"type":"array","items":{"type":"object","properties":{"build":{"type":"string","enum":["3-x-x","x-3-x","x-x-3"]}`,
		`"outcome":{"type":"string","enum":["pass","fail","unresolved"]}`,
		`"required":["summary","findings","omissionVerdicts","thirdPurchaseVerdicts"]`,
	} {
		if !strings.Contains(schema, want) {
			t.Errorf("the review schema lacks %s", want)
		}
	}
	// Findings keep their limit of eight.
	if !strings.Contains(schema, `"findings":{"maxItems":8`) {
		t.Error("the findings lost their limit")
	}
}

// A review with one verdict per subject is recorded whole: every verdict is
// a model Finding, a pass included, after the deterministic findings and
// before the review's own.
func TestReviewRecordsEveryVerdict(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	verdicts := verdictFindings(stages.Result)
	var got []string
	for _, f := range verdicts {
		got = append(got, f.ID+" "+f.Rule+" "+f.Outcome+" "+f.Severity+" "+f.Subject)
	}
	want := []string{
		"verdict.omission.1 omission-verdict pass info omittedTechniques, Allied Fan Club (minor)",
		"verdict.omission.2 omission-verdict pass info omittedTechniques, Critical shots (minor)",
		"verdict.third-purchase.path1 path-identity-verdict pass info path1, 3-x-x Spike-o-pult",
		"verdict.third-purchase.path2 path-identity-verdict pass info path2, x-3-x Triple Shot",
		"verdict.third-purchase.path3 path-identity-verdict pass info path3, x-x-3 Crossbow",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("verdicts:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
	if verdicts[2].Message != "3-x-x Spike-o-pult turns the top path's pierce into a heavy spiked ball with 13 more pierce and more damage; no other path's purchase reaches that pierce." || verdicts[2].Action != nil || strings.Join(verdicts[2].Evidence, ",") != "dart-monkey-atlas-56-3" {
		t.Errorf("3-x-x verdict %+v", verdicts[2])
	}
	checked := len(stages.Checked.Findings)
	if first := stages.Result.Findings[checked]; first.ID != "verdict.omission.1" {
		t.Errorf("the first finding after the deterministic ones is %s", first.ID)
	}
	if last := stages.Result.Findings[len(stages.Result.Findings)-1]; last.ID != "model.fan-club-allies" {
		t.Errorf("the last finding is %s", last.ID)
	}
	// The Result reads back with its verdicts.
	reread, err := unit.ParseResult(s.FromGoValue(stages.Result))
	if err != nil || len(verdictFindings(reread)) != 5 {
		t.Errorf("the Result does not read back with its verdicts: %v", err)
	}
}

// A fail or unresolved verdict is an open Finding with its action.
func TestNonPassVerdictBecomesAFinding(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	output := recordedOutput(t, "review")
	omission := at(output, "omissionVerdicts").([]any)[1].(*s.Object)
	omission.Set("outcome", "fail").
		Set("reason", "A shot counter is unsupported, but the critical hit's central effect, larger damage on some shots, is damage a purchase could raise, or a proposed mechanic on x-x-4 could carry.").
		Set("action", "Adapt Critical shots on x-x-4 as damage with a proposed shot counter, or give a reason that holds.")
	purchase := at(output, "thirdPurchaseVerdicts").([]any)[2].(*s.Object)
	purchase.Set("outcome", "unresolved").Set("reason", "x-x-3 Crossbow adds damage and pierce; the sources do not say whether the crossbow's reach defines the path.").Set("action", "Decide what x-x-3 adds that x-x-1 and x-x-2 lack.")
	model := &fixture.Model{Outputs: []any{output}}
	result, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	byID := map[string]unit.Finding{}
	for _, f := range verdictFindings(result) {
		byID[f.ID] = f
	}
	failed := byID["verdict.omission.2"]
	if failed.Outcome != "fail" || failed.Severity != "error" || failed.Category != "coverage" || failed.Method != "model" ||
		failed.Action == nil || *failed.Action != "Adapt Critical shots on x-x-4 as damage with a proposed shot counter, or give a reason that holds." ||
		!strings.HasPrefix(failed.Message, "A shot counter is unsupported, but the critical hit's central effect") {
		t.Errorf("the failed omission verdict %+v", failed)
	}
	open := byID["verdict.third-purchase.path3"]
	if open.Outcome != "unresolved" || open.Severity != "warning" || open.Category != "conflict" || open.Action == nil {
		t.Errorf("the unresolved third purchase verdict %+v", open)
	}
	if byID["verdict.omission.1"].Outcome != "pass" {
		t.Error("the passing verdict was not kept")
	}
}

// A review that leaves out a verdict, repeats one or judges a subject that
// is not asked for is malformed output: it fails like a review with invalid
// finding references, and Author keeps the checked draft.
func TestIncompleteVerdictsAreRejected(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	cases := map[string]struct {
		edit func(*s.Object)
		want string
	}{
		"a missing third purchase": {func(o *s.Object) {
			o.Set("thirdPurchaseVerdicts", at(o, "thirdPurchaseVerdicts").([]any)[:2])
		}, "no verdict on the third purchase x-x-3"},
		"a missing omission": {func(o *s.Object) {
			o.Set("omissionVerdicts", at(o, "omissionVerdicts").([]any)[1:])
		}, `no verdict on the omission of "Allied Fan Club"`},
		"no verdicts": {func(o *s.Object) {
			o.Delete("omissionVerdicts")
			o.Delete("thirdPurchaseVerdicts")
		}, `no verdict on the omission of "Allied Fan Club"; no verdict on the omission of "Critical shots"; no verdict on the third purchase 3-x-x`},
		"a repeated verdict": {func(o *s.Object) {
			verdicts := at(o, "thirdPurchaseVerdicts").([]any)
			o.Set("thirdPurchaseVerdicts", append(verdicts, s.Clone(verdicts[0])))
		}, "2 verdicts on the third purchase 3-x-x"},
		"an unknown subject": {func(o *s.Object) {
			verdicts := at(o, "omissionVerdicts").([]any)
			extra := s.Clone(verdicts[0]).(*s.Object).Set("technique", "Monkey Knowledge")
			o.Set("omissionVerdicts", append(verdicts, extra))
		}, `a verdict on "Monkey Knowledge", which designPlan.omittedTechniques does not list`},
	}
	for name, c := range cases {
		output := recordedOutput(t, "review")
		c.edit(output)
		model := &fixture.Model{Outputs: []any{output}}
		_, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
		var failure *unit.ModelError
		if !errors.As(err, &failure) || failure.Failure == nil || failure.Failure.Code != unit.CodeOutputInvalid || failure.Failure.Stage != "review" ||
			!strings.Contains(failure.Message, "did not return exactly one verdict per whole-technique omission and per third purchase") ||
			!strings.Contains(failure.Message, c.want) || len(model.Requests) != 1 {
			t.Errorf("%s: %v after %d calls", name, err, len(model.Requests))
		}
	}
	// A verdict citing an unknown document is invalid like such a finding.
	output := recordedOutput(t, "review")
	at(output, "thirdPurchaseVerdicts").([]any)[0].(*s.Object).Set("evidence", []any{"no-such-document"})
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, &fixture.Model{Outputs: []any{output}}, fixture.Options()); err == nil || !strings.Contains(err.Error(), "invalid finding or evidence references") {
		t.Errorf("a verdict citing an unknown document: %v", err)
	}
}

// The citation correction concerns findings only: the first review's
// verdicts are published even when the correction returns others.
func TestCorrectionKeepsTheFirstVerdicts(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	valid := recordedOutput(t, "review")
	illegal := s.Clone(valid).(*s.Object)
	at(illegal, "findings").([]any)[0].(*s.Object).Set("message", "The 3-3-0 crosspath loses the frenzy.")
	changed := s.Clone(valid).(*s.Object)
	at(changed, "thirdPurchaseVerdicts").([]any)[0].(*s.Object).Set("outcome", "fail").Set("action", "Rename it.")
	model := &fixture.Model{Outputs: []any{illegal, changed}}
	result, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(model.Requests[1].Prompt, "with omissionVerdicts and thirdPurchaseVerdicts as in your previous review") {
		t.Error("the correction does not ask for the verdicts")
	}
	for _, f := range verdictFindings(result) {
		if f.Outcome != "pass" {
			t.Errorf("the correction changed the verdict %s", f.ID)
		}
	}
}

// A plan without omissions asks for none; its third purchases
// still need verdicts.
func TestVerdictsWithoutOmissions(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	checked := stages.Checked
	plan := *checked.Draft.Run.DesignPlan
	plan.OmittedTechniques = nil
	checked.Draft.Run.DesignPlan = &plan
	request := unit.BlueprintReviewRequest(checked)
	if !strings.Contains(request.Prompt, `"requiredVerdicts":{"omissions":[],"thirdPurchases":[{"build":"3-x-x"`) ||
		!strings.Contains(s.Stringify(request.Schema), `"omissionVerdicts":{"minItems":0,"maxItems":0`) {
		t.Error("a plan without omissions still asks for omission verdicts")
	}
}

// A coherent, source-backed proposed mechanic that is a third purchase's only
// new capability is an unresolved design gap, not a fail; an invalid or
// numbers-only proposal still fails (SOL-72-01).
func TestProposedOnlyThirdPurchaseVerdictIsUnresolved(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prompt := unit.BlueprintReviewRequest(stages.Checked).Prompt
	for _, want := range []string{
		"When a third purchase's only new capability is a proposed mechanic, give its verdict unresolved, a design gap until the Definition supports it, if the proposal is coherent, fits the technique's cited source and adds a real capability",
		"give it fail if the proposal is invalid, only renames larger numbers, conflicts with the Definition's rules",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	if strings.Contains(prompt, "fail one whose only distinction is a proposed mechanic that is not playable") {
		t.Error("the review prompt still fails every proposed-only third purchase")
	}
}
