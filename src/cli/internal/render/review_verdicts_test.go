package render_test

import (
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/render"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// The details report lists every review verdict, a pass included, in an
// omission section, a third purchase section and a fifth purchase section,
// and a failed verdict only there, not again among the other findings. The
// unit sheet shows none (SOL-61-08, SOL-61-10).
func TestDetailsListReviewVerdicts(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	result := stages.Result
	for i, f := range result.Findings {
		if f.ID == "verdict.omission.2" {
			action := "Adapt Critical shots as damage with a proposed shot counter."
			result.Findings[i].Outcome, result.Findings[i].Severity, result.Findings[i].Action = "fail", "error", &action
		}
	}
	text, err := render.Markdown(s.FromGoValue(result), true)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"Model review: 6 passed, 1 failed, 2 unresolved, 0 not checked.",
		"### Omission verdicts\n\nThe model review's verdict on each whole-technique omission",
		"| Outcome | Subject | Verdict and next action |\n| --- | --- | --- |\n| pass | omittedTechniques, Allied Fan Club (minor) | Transforming nearby allied Units is the technique's central effect",
		"| fail | omittedTechniques, Critical shots (minor) | A hit counter is the central effect, and damage or a proposed mechanic on x-4-x would carry only its size, not its every-tenth-shot timing; minor fits a brief that names it only on the bottom path. Adapt Critical shots as damage with a proposed shot counter. |",
		"### Third purchase verdicts\n\nThe model review's verdict on each path's third purchase",
		"| pass | path2, x-3-x Triple Shot | x-3-x Triple Shot is the only purchase that fires more than one projectile; x-1-x and x-2-x only attack faster. |",
		"### Fifth purchase verdicts\n\nThe model review's verdict on each path's fifth purchase",
		"| unresolved | path2, x-5-x Plasma Monkey Fan Club | x-5-x Plasma Monkey Fan Club only doubles the frenzy's damage and lengthens it;",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("the details report lacks %q", want)
		}
	}
	if strings.Contains(text, "| fail | model | omittedTechniques") || strings.Contains(text, "| unresolved | model | path2, x-5-x") {
		t.Error("the failed verdict is listed again among the findings")
	}
	sheet, err := render.Markdown(s.FromGoValue(result), false)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(sheet, "verdict") || strings.Contains(sheet, "Allied Fan Club") {
		t.Error("the unit sheet shows review verdicts")
	}
}

// A Result saved before review verdicts reads and renders as before, without
// verdict sections.
func TestResultWithoutVerdictsRenders(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	result := stages.Result
	var kept []unit.Finding
	for _, f := range result.Findings {
		if !unit.IsReviewVerdict(f) {
			kept = append(kept, f)
		}
	}
	result.Findings = kept
	value := s.FromGoValue(result)
	if _, err := unit.ParseResult(value); err != nil {
		t.Fatalf("a Result without verdicts does not read: %v", err)
	}
	text, err := render.Markdown(value, true)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(text, "verdicts") || !strings.Contains(text, "Model review: 0 passed, 0 failed, 1 unresolved, 0 not checked.") || !strings.Contains(text, "| unresolved | model | path-2 tier 4 |") {
		t.Errorf("a Result without verdicts renders differently:\n%s", text)
	}
	if _, err := render.Markdown(value, false); err != nil {
		t.Fatal(err)
	}
}
