package unit_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// billedModel is the scripted model with a cost on every call, so a test
// can see each attempt's usage.
type billedModel struct{ *fixture.Model }

func (b billedModel) Generate(ctx context.Context, request unit.ModelRequest) (unit.ModelResponse, error) {
	response, err := b.Model.Generate(ctx, request)
	cost := float64(len(b.Requests)) / 1000
	response.Usage = &unit.Usage{CostUSD: &cost}
	return response, err
}

// fourCorePlan is the fixture plan with four core entries, an issue no
// correction fixes, beside x-5-x with no new behavior under the Default.
func fourCorePlan(t *testing.T) *s.Object {
	t.Helper()
	plan := recordedOutput(t, "plan")
	repertoireEntry(plan, "Fan Club").Set("importance", "core")
	return plan
}

// correctedCapstone is the fixture plan whose x-5-x adapts Fan Club's
// allied frenzy as a proposed mechanic and says so in its change.
func correctedCapstone(t *testing.T) *s.Object {
	t.Helper()
	plan := recordedOutput(t, "plan")
	milestone := at(plan, "paths", "path2", "milestones", "tier5").(*s.Object)
	milestone.Set("change", "x-5-x adapts Plasma Monkey Fan Club's allied transformation: the frenzy doubles dart damage, lasts longer and turns nearby Dart Monkeys into plasma throwers.")
	milestone.Set("proposedMechanics", []any{plasmaTransformation()})
	return plan
}

func planJSON(t *testing.T, plan *unit.DesignPlan) string {
	t.Helper()
	if plan == nil {
		t.Fatal("no plan")
	}
	return s.Stringify(s.FromGoValue(*plan))
}

// When the full-plan retry fails only on a fifth purchase with no new
// behavior, one capstone correction names that purchase, its milestone and
// its technique's effects, and asks for the source effect and the playable
// behavior it adds. Its attempt and cost are recorded, and the corrected
// plan passes the same validation (SOL-61-07).
func TestCapstoneCorrectionAfterFullRetry(t *testing.T) {
	prepared := underTheDefault(t)
	model := billedModel{&fixture.Model{Outputs: []any{fourCorePlan(t), recordedOutput(t, "plan"), correctedCapstone(t), recordedOutput(t, "mechanics")}}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if got := strings.Join(purposes(draft), ","); got != "plan,plan,capstone-correction,design" {
		t.Fatalf("attempts %s", got)
	}
	attempts := *draft.Run.Attempts
	if len(attempts[1].Issues) != 1 || !strings.HasPrefix(attempts[1].Issues[0], "paths.path2.milestones.tier5: x-5-x promises no new behavior or access") {
		t.Errorf("the retry's issues: %v", attempts[1].Issues)
	}
	if u := attempts[2].Usage; u == nil || u.CostUSD == nil || *u.CostUSD != 0.003 || len(attempts[2].Issues) != 0 {
		t.Errorf("the capstone correction attempt: %+v", attempts[2])
	}
	if u := draft.Run.Usage; u == nil || u.CostUSD == nil || *u.CostUSD < 0.00999 || *u.CostUSD > 0.01001 {
		t.Errorf("the draft's usage does not include every attempt: %+v", u)
	}
	prompt := model.Requests[2].Prompt
	for _, want := range []string{
		"The previous plan failed only on the items below. Each capstone item is a fifth purchase that promises no new behavior or access",
		"A capstone with no new behavior is a design choice as well as a missing field, so do not add an unlock only to pass this check.",
		"choose the source effect the capstone adapts",
		"Change only the fifth milestone each capstone item names",
		"Code keeps every other milestone and field exactly as it was and validates the complete plan again.",
		`"path":"paths.path2.milestones.tier5","purchase":"x-5-x","pathName":"Fan Club Line"`,
		`"milestone":{"change":"x-5-x doubles dart damage during the frenzy and lengthens the frenzy window."`,
		`"technique":{"name":"Fan Club"`,
		`"effects":[`,
		`"pathUnlocks":["manual-boost"]`,
		"not an unlock in pathUnlocks",
		"name it in the proposedMechanics of x-5-x",
		`"previous":{`,
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the capstone correction lacks %q", want)
		}
	}
	for _, unwanted := range []string{"Correct this invalid design plan", `"purchase":"5-x-x"`, `"purchase":"x-x-5"`} {
		if strings.Contains(prompt, unwanted) {
			t.Errorf("the capstone correction has %q", unwanted)
		}
	}
	want, err := unit.DecodeDesignPlan(correctedCapstone(t), &prepared.Request)
	if err != nil {
		t.Fatal(err)
	}
	if planJSON(t, draft.Run.DesignPlan) != planJSON(t, &want) {
		t.Error("the draft's plan is not the retry's plan with the corrected x-5-x")
	}
}

// Code keeps every other milestone: a correction that also rewrites
// another milestone, or the repertoire, has those changes discarded and
// only its fifth milestone and capstoneValue merged, and the merged plan
// validates.
func TestCapstoneCorrectionKeepsOtherMilestones(t *testing.T) {
	prepared := underTheDefault(t)
	corrected := correctedCapstone(t)
	at(corrected, "paths", "path2").(*s.Object).Set("capstoneValue", "The allied frenzy is the reason to buy x-5-x.")
	at(corrected, "paths", "path1", "milestones", "tier1").(*s.Object).Set("change", "1-x-x throws a spiked ball.").Set("unlock", "splash")
	at(corrected, "paths", "path3", "milestones", "tier4").(*s.Object).Set("improves", []any{"range"})
	repertoireEntry(corrected, "Fan Club").Set("importance", "core")
	model := &fixture.Model{Outputs: []any{fourCorePlan(t), recordedOutput(t, "plan"), corrected, recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if got := strings.Join(purposes(draft), ","); got != "plan,plan,capstone-correction,design" {
		t.Fatalf("attempts %s", got)
	}
	expected := correctedCapstone(t)
	at(expected, "paths", "path2").(*s.Object).Set("capstoneValue", "The allied frenzy is the reason to buy x-5-x.")
	want, err := unit.DecodeDesignPlan(expected, &prepared.Request)
	if err != nil {
		t.Fatal(err)
	}
	if planJSON(t, draft.Run.DesignPlan) != planJSON(t, &want) {
		t.Error("the merged plan kept a change outside x-5-x and its capstoneValue")
	}
}

// No capstone correction is sent when any other kind of issue remains,
// after the first plan, which gets the full retry, with a retry budget of 0
// or a second time in one draft.
func TestNoCapstoneCorrection(t *testing.T) {
	prepared := underTheDefault(t)

	// Another kind of issue remains after the retry.
	model := &fixture.Model{Outputs: []any{fourCorePlan(t), fourCorePlan(t)}}
	if _, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options()); err == nil || len(model.Requests) != 2 {
		t.Errorf("mixed issues after the retry: %v after %d calls", err, len(model.Requests))
	}

	// The first plan fails only on x-5-x: the full retry follows.
	model = &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), correctedCapstone(t), recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if got := strings.Join(purposes(draft), ","); got != "plan,plan,design" || !strings.Contains(model.Requests[1].Prompt, "Correct this invalid design plan") {
		t.Errorf("the first plan's retry: %s", got)
	}

	// No retry budget.
	none := 0
	options := fixture.Options()
	options.MaxRepairAttempts = &none
	model = &fixture.Model{Outputs: []any{recordedOutput(t, "plan")}}
	if _, err := unit.DraftUnit(context.Background(), prepared, model, options); err == nil || len(model.Requests) != 1 {
		t.Errorf("budget 0: %v after %d calls", err, len(model.Requests))
	}

	// At most one per draft: a correction that still fails is followed by
	// the remaining full retry, and final validation stays strict.
	two := 2
	options = fixture.Options()
	options.MaxRepairAttempts = &two
	model = &fixture.Model{Outputs: []any{fourCorePlan(t), recordedOutput(t, "plan"), recordedOutput(t, "plan"), recordedOutput(t, "plan")}}
	_, err = unit.DraftUnit(context.Background(), prepared, model, options)
	var failure *unit.ModelError
	if err == nil || !strings.Contains(err.Error(), "could not be validated") || len(model.Requests) != 4 {
		t.Fatalf("an uncorrected capstone: %v after %d calls", err, len(model.Requests))
	}
	if !errors.As(err, &failure) || failure.Evidence == nil {
		t.Fatal("no failure evidence")
	}
	var got []string
	for _, attempt := range failure.Evidence.Attempts {
		got = append(got, attempt.Purpose)
	}
	if strings.Join(got, ",") != "plan,plan,capstone-correction,plan" {
		t.Errorf("attempts %v", got)
	}
}

// A capstone correction also names the items a targeted correction fixes,
// and an unpromised adaptedAs may gain its promise only on a fifth purchase
// the correction changes; otherwise its adaptedAs loses it. The corrected
// repertoire is then merged too.
func TestCapstoneCorrectionWithNarrowItems(t *testing.T) {
	prepared := underTheDefault(t)
	retry := unpromisedPlan(t)
	prompt, merge, ok := unit.CapstoneCorrection(retry, &prepared.Request)
	if !ok {
		t.Fatal("no capstone correction")
	}
	for _, want := range []string{
		`"purchase":"x-5-x"`,
		`"path":"repertoire.3.effects.0.adaptedAs"`,
		`"allowed":["Remove burn from this adaptedAs`,
		"correct each other item with one of its allowed corrections, in the repertoire or omittedTechniques or on a fifth milestone this correction changes.",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the correction lacks %q", want)
		}
	}
	if strings.Contains(prompt, `Promise burn on a milestone`) {
		t.Error("the correction allows a promise on a milestone it does not change")
	}
	if _, err := unit.DecodeDesignPlan(merge(correctedCapstone(t)), &prepared.Request); err != nil {
		t.Errorf("the merged correction: %v", err)
	}

	// Another kind of issue: no capstone correction.
	if _, _, ok := unit.CapstoneCorrection(fourCorePlan(t), &prepared.Request); ok {
		t.Error("a capstone correction beside four core entries")
	}
}

// Escanor's v32 plan retry (OPUS-NET-61-10) fails only on 5-x-x, and its
// capstone correction names Pride Flare and its effects.
func TestEscanorCapstoneCorrection(t *testing.T) {
	request := savedSources(t, "escanor")
	data, err := os.ReadFile("testdata/escanor-v32-retry.plan.json")
	if err != nil {
		t.Fatal(err)
	}
	output, err := s.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	_, err = unit.DecodeDesignPlan(output, &request)
	if err == nil || strings.Count(err.Error(), "upgradeIntents.") != 1 || !strings.Contains(err.Error(), "upgradeIntents.path1.tier5: 5-x-x promises no new behavior") {
		t.Fatalf("the retry's issues: %v", err)
	}
	prompt, _, ok := unit.CapstoneCorrection(output, &request)
	if !ok {
		t.Fatal("no capstone correction")
	}
	for _, want := range []string{`"purchase":"5-x-x","pathName":"Sunfire Impact"`, `"technique":{"name":"Pride Flare"`, `"pathUnlocks":["splash","burn"]`} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the correction lacks %q", want)
		}
	}
}
