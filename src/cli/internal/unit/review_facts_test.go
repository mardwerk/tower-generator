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

// The review reads every legal build as code resolved it, so it need not
// recompute a hit count or guess which builds exist (reported on #27: a
// review said pierce 2 cannot reach a second enemy and judged 3-3-0).
func TestReviewReadsResolvedLegalBuilds(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	facts := unit.LegalBuildFacts(stages.Checked.Draft.Candidate.Blueprint, unit.DefaultAuthoringDefinition())
	if len(facts) != 64 {
		t.Fatalf("%d legal builds", len(facts))
	}
	var base *s.Object
	for _, fact := range facts {
		if code, _ := fact.(*s.Object).Get("code"); code == "0-0-0" {
			base = fact.(*s.Object)
		}
	}
	if got := s.Stringify(base); base == nil || !strings.Contains(got, `"cost":200`) || !strings.Contains(got, `"pierce":2`) || !strings.Contains(got, `"maxEnemiesHitPerAttack":2`) {
		t.Errorf("0-0-0 facts %s", got)
	}
	prompt := unit.BlueprintReviewRequest(stages.Checked).Prompt
	if !strings.Contains(prompt, `"legalBuilds":[`) || !strings.Contains(prompt, `"code":"3-2-0"`) || strings.Contains(prompt, `"code":"3-3-0"`) || !strings.Contains(prompt, "such as 3-3-0, is illegal") {
		t.Error("the review prompt lacks the legal build facts")
	}
	// A copy bound is context for a capstone, not a verdict, the draft's
	// proposals are claims to check, not facts, and the period decision is
	// context, not text the Unit must print (reported on #27).
	for _, want := range []string{"A same-budget copy bound is context, not a verdict", "They are the drafting model's claims and can be wrong", "A name taken from the source", "not text the Unit must print"} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
}

// A review that cites a build the Definition does not allow is corrected
// once and then rejected.
func TestReviewsCitingIllegalBuildsAreCorrectedOnce(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	valid := recordedOutput(t, "review")
	illegal := s.Clone(valid).(*s.Object)
	findings, _ := illegal.Get("findings")
	findings.([]any)[0].(*s.Object).Set("message", "The 3-3-0 crosspath loses the frenzy.")

	model := &fixture.Model{Outputs: []any{illegal, valid}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil {
		t.Fatalf("a corrected review failed: %v", err)
	}
	if len(model.Requests) != 2 || !strings.Contains(model.Requests[1].Prompt, "It cites builds that are not legal under this Definition: 3-3-0.") {
		t.Fatalf("%d calls; correction missing", len(model.Requests))
	}

	model = &fixture.Model{Outputs: []any{illegal, illegal}}
	_, err = unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	var failure *unit.ModelError
	if !errors.As(err, &failure) || !strings.Contains(failure.Message, "cited builds that are not legal (3-3-0) after one correction") || len(model.Requests) != 2 {
		t.Errorf("a review citing 3-3-0 twice: %v after %d calls", err, len(model.Requests))
	}
}

// A finding cites the resolved values it relies on and code checks them, so
// a review cannot report that 2-x-x lacks a change it has (reported on #27).
func TestReviewFindingsCiteCheckedFacts(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	citing := func(build, field, value string) *s.Object {
		review := recordedOutput(t, "review")
		findings, _ := review.Get("findings")
		findings.([]any)[0].(*s.Object).Set("facts", []any{s.NewObject().Set("build", build).Set("field", field).Set("value", value)})
		return review
	}
	model := &fixture.Model{Outputs: []any{citing("2-0-0", "attack.pierce", "5")}}
	result, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	if err != nil || len(model.Requests) != 1 {
		t.Fatalf("a correct citation: %v after %d calls", err, len(model.Requests))
	}
	last := result.Findings[len(result.Findings)-1]
	if len(last.Facts) != 1 || last.Facts[0].Value != "5" {
		t.Errorf("the result lost the cited facts: %+v", last)
	}

	model = &fixture.Model{Outputs: []any{citing("2-0-0", "attack.pierce", "3"), citing("2-0-0", "attack.pierce", "5")}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil {
		t.Fatalf("a corrected citation: %v", err)
	}
	if len(model.Requests) != 2 || !strings.Contains(model.Requests[1].Prompt, "cites 2-0-0 attack.pierce as 3, but it is 5.") {
		t.Errorf("%d calls; the correction does not name the resolved value", len(model.Requests))
	}

	model = &fixture.Model{Outputs: []any{citing("2-0-0", "attack.width", "3"), citing("2-0-0", "attack.width", "3")}}
	_, err = unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	var failure *unit.ModelError
	if !errors.As(err, &failure) || !strings.Contains(failure.Message, "a field legalBuilds does not list for that build") {
		t.Errorf("an unknown field: %v", err)
	}
}
