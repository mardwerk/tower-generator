package unit_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// unpromisedPlan is the fixture plan with Crossbow's first effect adapted as
// burn, which no bottom path milestone promises: the only failure a targeted
// correction can fix.
func unpromisedPlan(t *testing.T) *s.Object {
	t.Helper()
	plan := recordedOutput(t, "plan")
	effects, _ := repertoireEntry(plan, "Crossbow").Get("effects")
	effects.([]any)[0].(*s.Object).Set("adaptedAs", []any{"damage", "burn"})
	return plan
}

func purposes(draft unit.Draft) []string {
	var out []string
	for _, attempt := range *draft.Run.Attempts {
		out = append(out, attempt.Purpose)
	}
	return out
}

// A plan that fails only on an unpromised adaptedAs gets one targeted
// correction naming that item and its allowed corrections, outside the
// full-plan retry budget. Its result passes the same plan validation, and
// the attempt is recorded with its purpose and issues (SOL-61-05).
func TestTargetedPlanCorrection(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	model := &fixture.Model{Outputs: []any{unpromisedPlan(t), recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if got := strings.Join(purposes(draft), ","); got != "plan,plan-correction,design" {
		t.Errorf("attempts %s", got)
	}
	first := (*draft.Run.Attempts)[0]
	if len(first.Issues) != 1 || !strings.HasPrefix(first.Issues[0], "repertoire.3.effects.0.adaptedAs: ") || !strings.Contains(first.Issues[0], "is adapted as burn") {
		t.Errorf("the failed plan attempt's issues: %v", first.Issues)
	}
	prompt := model.Requests[1].Prompt
	for _, want := range []string{
		"The previous plan failed only on the items below.",
		`"path":"repertoire.3.effects.0.adaptedAs"`,
		`Promise burn on a milestone whose technique is \"Crossbow\"`,
		`Remove burn from this adaptedAs`,
		`"previous":{`,
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the targeted correction lacks %q", want)
		}
	}
	if strings.Contains(prompt, "Correct this invalid design plan") {
		t.Error("the targeted correction also asks for a full re-plan")
	}

	// The correction's result is validated like any plan: when it still
	// fails, the full-plan retry follows, and final validation stays
	// strict.
	model = &fixture.Model{Outputs: []any{unpromisedPlan(t), unpromisedPlan(t), unpromisedPlan(t)}}
	_, err = unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err == nil || !strings.Contains(err.Error(), "could not be validated") || len(model.Requests) != 3 {
		t.Fatalf("an uncorrected plan: %v after %d calls", err, len(model.Requests))
	}
	if !strings.Contains(model.Requests[2].Prompt, "Correct this invalid design plan") {
		t.Error("the full retry after a failed correction is not a full re-plan")
	}

	// No correction is sent without a retry budget.
	none := 0
	options := fixture.Options()
	options.MaxRepairAttempts = &none
	model = &fixture.Model{Outputs: []any{unpromisedPlan(t)}}
	if _, err := unit.DraftUnit(context.Background(), prepared, model, options); err == nil || len(model.Requests) != 1 {
		t.Errorf("budget 0: %v after %d calls", err, len(model.Requests))
	}
}

// A plan that also fails on anything a targeted correction cannot fix gets
// the ordinary full-plan retry.
func TestMixedPlanFailureGetsFullRetry(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	mixed := unpromisedPlan(t)
	repertoireEntry(mixed, "Fan Club").Set("importance", "core") // four core entries
	model := &fixture.Model{Outputs: []any{mixed, recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if got := strings.Join(purposes(draft), ","); got != "plan,plan,design" {
		t.Errorf("attempts %s", got)
	}
	if !strings.Contains(model.Requests[1].Prompt, "Correct this invalid design plan") {
		t.Error("the retry is not a full re-plan")
	}
}

// For a source technique listed nowhere, the correction offers listing it
// with a rank, by its name or an alias, major or core when its salience is
// strong.
func TestCorrectionListsUnlistedTechnique(t *testing.T) {
	request := savedSources(t, "escanor")
	plan := corePlan(unit.PlanBase{Name: "Sunshine"})
	items := s.Stringify(s.FromGoValue(unit.CorrectionItems(plan, &request)))
	for _, want := range []string{
		`The source technique \"Rhitta\" is listed neither in the repertoire nor in omittedTechniques.`,
		`by the name \"Rhitta\" or one of its aliases, \"The Divine Axe Rhitta\", ranked major or core, since its salience is strong`,
		`The source technique \"Cruel Sun\" is listed`,
		`by the name \"Cruel Sun\", ranked core, major or minor`,
	} {
		if !strings.Contains(items, want) {
			t.Errorf("the correction items lack %q", want)
		}
	}
	if strings.Contains(items, `The source technique \"Sunshine\"`) {
		t.Error("the base attack's technique is named as unlisted")
	}
}

// A source technique of strong salience is ranked major or core wherever
// the plan lists it: ranked minor, in the repertoire or in
// omittedTechniques, by its name or an alias, the plan fails. A normal one
// may be minor, and strong salience never requires core (SOL-61-05).
func TestStrongSalienceIsAMajorFloor(t *testing.T) {
	request := savedSources(t, "escanor")
	floor := func(plan unit.DesignPlan) string {
		var out []string
		for _, issue := range unit.CoreConceptIssues(plan, &request) {
			if strings.Contains(issue.Message, "whose salience is strong") {
				out = append(out, issue.Path+": "+issue.Message)
			}
		}
		return strings.Join(out, "\n")
	}
	base := unit.PlanBase{Name: "Sunshine"}
	for _, c := range []struct {
		name string
		plan unit.DesignPlan
		want string
	}{
		{"omitted minor", unit.DesignPlan{Base: base, OmittedTechniques: []unit.PlanOmission{{Name: "Rhitta", Importance: "minor", Reason: "An axe."}}},
			`omittedTechniques.0.importance: "Rhitta" is ranked minor, but it lists the source technique "Rhitta", whose salience is strong. ` + unit.SalienceRule + " Rank it major or core."},
		{"an alias in the repertoire ranked minor", corePlan(base, unit.PlanRepertoire{Name: "The Divine Axe Rhitta", Importance: "minor"}),
			`repertoire.0.importance: "The Divine Axe Rhitta" is ranked minor, but it lists the source technique "Rhitta"`},
		{"The One ranked minor", corePlan(base, unit.PlanRepertoire{Name: "The One", Importance: "minor"}), `lists the source technique "The One"`},
		{"ranked major", unit.DesignPlan{Base: base, OmittedTechniques: []unit.PlanOmission{{Name: "Rhitta", Importance: "major", Reason: "An axe."}}}, ""},
		{"a normal technique ranked minor", corePlan(base, unit.PlanRepertoire{Name: "Cruel Sun", Importance: "minor"}), ""},
	} {
		got := floor(c.plan)
		if (c.want == "") != (got == "") || !strings.Contains(got, c.want) {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}
