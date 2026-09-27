package unit_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Every whole-technique omission, major included, reaches the review with
// its importance and reason, and one named for a source technique carries
// that technique's salience and passages. The review judges each against
// the vocabulary and purchase-level proposed mechanics, fails a circular
// reason and challenges rankings against the cited passages (SOL-61-05).
func TestOmissionsReachTheReview(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	checked := stages.Checked
	techniques := append([]unit.SourceTechnique(nil), fixtureSourceTechniques...)
	techniques[4].Salience = unit.SalienceStrong // Juggernaut
	checked.Draft.Prepared.Request.SourceTechniques = &techniques
	plan := *checked.Draft.Run.DesignPlan
	plan.OmittedTechniques = append(append([]unit.PlanOmission(nil), plan.OmittedTechniques...),
		unit.PlanOmission{Name: "Juggernaut", Importance: "major", Reason: "This plan does not implement it."})
	checked.Draft.Run.DesignPlan = &plan
	prompt := unit.BlueprintReviewRequest(checked).Prompt
	for _, want := range []string{
		`{"name":"Juggernaut","importance":"major","reason":"This plan does not implement it.","sourceTechnique":{"name":"Juggernaut","salience":"strong","passageIds":["source1:7","source1:8"]}}`,
		`{"name":"Critical shots","importance":"minor","reason":"Every tenth or fifth shot dealing extra damage needs a shot counter operator."}`,
		"judge every whole-technique omission in designPlan.omittedTechniques, major and minor alike, against definition.vocabulary and the proposed mechanics a purchase could carry",
		unit.OmissionRule,
		`a circular reason never holds, as "this plan does not implement it"`,
		"Challenge every ranking against the passages its source technique cites",
		"an implausible downgrade included",
		unit.SalienceRule,
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	planRequest, err := unit.DesignPlanRequest(checked.Draft.Prepared)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(planRequest.Prompt, unit.OmissionRule) {
		t.Error("the plan prompt does not state the omission rule")
	}
}

// The review compares each third purchase with its own path's first two
// purchases and the other paths' purchases, keeps the Dart reference and
// judges a proposed-only capstone as a design gap (#66).
func TestReviewComparesThirdPurchases(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prompt := unit.BlueprintReviewRequest(stages.Checked).Prompt
	for _, want := range []string{
		"comparing each third purchase explicitly with its own path's first and second purchases and with every purchase of the other two paths",
		"Fail a third purchase that only adds numbers its path already had",
		"a punch that homes around obstructions where every delivery needs a clear path",
		"Dart Monkey's x-3-x Triple Shot only adds projectiles and no other Dart Monkey path adds any",
		"Code reports a fifth purchase whose only new capability is a proposed mechanic as an unresolved design gap",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
}

// The plan, mechanics and review prompts state what an Active Ability can
// scope, and the review reads each purchase's typed changes by scope, so a
// text that limits a base change to the Active can fail (OPUS-NET-61-7).
func TestActiveAbilityScope(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	model := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	review := unit.BlueprintReviewRequest(checked).Prompt
	for name, prompt := range map[string]string{"plan": model.Requests[0].Prompt, "mechanics": model.Requests[1].Prompt, "review": review} {
		if !strings.Contains(prompt, unit.ActiveScopeRule) {
			t.Errorf("the %s prompt does not state the Active Ability scope", name)
		}
	}
	for _, want := range []string{
		// x-1-x attacks faster, a base change; x-4-x unlocks the Active and
		// x-5-x changes it, boost changes.
		`"plannedChange":"x-1-x shortens the attack interval.","changeScope":{"base":["intervalSeconds multiply 0.85"],"boost":[]}}`,
		`"changeScope":{"base":["intervalSeconds multiply 0.5"],"boost":["unlocks the Active Ability Fan Club Frenzy"]}}`,
		`"changeScope":{"base":[],"boost":["boost damageMultiplier`,
		"Fail, on that build code, a name, adaptation or plannedChange whose scope contradicts changeScope",
	} {
		if !strings.Contains(review, want) {
			t.Errorf("the review context lacks %q", want)
		}
	}
}
