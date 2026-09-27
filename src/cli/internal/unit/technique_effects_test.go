package unit_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// repertoireEffect is effect index of repertoire entry entry in a plan output.
func repertoireEffect(plan *s.Object, entry, index int) *s.Object {
	repertoire, _ := plan.Get("repertoire")
	effects, _ := repertoire.([]any)[entry].(*s.Object).Get("effects")
	return effects.([]any)[index].(*s.Object)
}

// Each repertoire entry lists the effects its cited passages describe, each
// adapted as promises or omitted with a reason, and code finds every adapted
// promise on a purchase of that technique. Luffy's Gear 2 source says it
// raises strength and speed; 7 of 10 plans on #27 promised only speed on
// every Gear 2 purchase and recorded no omission.
func TestTechniqueEffectsAreAdaptedOrOmitted(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	decode := func(edit func(plan *s.Object)) (unit.DesignPlan, error) {
		plan := recordedOutput(t, "plan")
		edit(plan)
		return unit.DecodeDesignPlan(plan, &prepared.Request)
	}
	// Adapted and omitted effects are retained with the plan.
	plan, err := decode(func(*s.Object) {})
	if err != nil {
		t.Fatal(err)
	}
	if got := plan.Repertoire[0].Effects; len(got) != 4 || got[0].AdaptedAs[0] != "damage" || len(got[3].AdaptedAs) != 0 || !strings.Contains(got[3].Reason, "no wall geometry") {
		t.Errorf("Spiked Ball effects %+v", got)
	}
	// A form that raises strength and speed, adapted as damage and
	// attack-rate, on a path whose purchases promise only speed.
	strength := func(plan *s.Object) {
		repertoireEffect(plan, 1, 1).Set("effect", "The darts are thrown harder and faster.").Set("adaptedAs", []any{"damage", "attack-rate"})
	}
	_, err = decode(strength)
	if err == nil || !strings.Contains(err.Error(), `repertoire.1.effects.1.adaptedAs: "The darts are thrown harder and faster." is adapted as damage, but no purchase that adapts "Triple Throw" promises damage`) {
		t.Errorf("a dropped effect was accepted: %v", err)
	}
	if strings.Contains(err.Error(), "adapted as attack-rate") {
		t.Errorf("a kept effect was reported: %v", err)
	}
	// The same effect kept on a purchase of that technique passes.
	if _, err := decode(func(plan *s.Object) {
		strength(plan)
		at(plan, "paths", "path2", "milestones", "tier3").(*s.Object).Set("improves", []any{"projectiles", "attack-rate", "damage"})
	}); err != nil {
		t.Errorf("a kept effect was rejected: %v", err)
	}
	// Omitted with a reason passes.
	if _, err := decode(func(plan *s.Object) {
		repertoireEffect(plan, 1, 1).Set("effect", "The darts are thrown harder.").Set("adaptedAs", []any{}).Set("reason", "This path keeps each dart's damage so that Crossbow owns heavier hits.")
	}); err != nil {
		t.Errorf("an omission with a reason was rejected: %v", err)
	}
	// The Active's promise keeps its plain form: x-5-x promises active-damage.
	if _, err := decode(func(plan *s.Object) {
		repertoireEffect(plan, 2, 1).Set("adaptedAs", []any{"damage"})
	}); err != nil {
		t.Errorf("active-damage did not keep damage: %v", err)
	}
	// Code checks only entries that purchases adapt: an entry named like the
	// base attack is shown by 0-0-0, and an entry no purchase names adds
	// nothing to the unit, so the review judges its list (4 of 5 Luffy plan
	// retries on #27 came from these two cases).
	unused := func(name string) func(*s.Object) {
		return func(plan *s.Object) {
			repertoire, _ := plan.Get("repertoire")
			plan.Set("repertoire", append(repertoire.([]any), s.NewObject().
				Set("name", name).Set("importance", "minor").Set("sourceIds", []any{"source1:2"}).
				Set("limitation", "The ordinary throw is described, not a purchase.").
				Set("effects", []any{s.NewObject().Set("effect", "The dart flies to a distant target.").Set("adaptedAs", []any{"range"}).Set("reason", "Reach is range.")})))
		}
	}
	for _, name := range []string{"Dart Monkey Legend", "dart throw"} {
		if _, err := decode(unused(name)); err != nil {
			t.Errorf("an entry no purchase adapts, %q, was checked: %v", name, err)
		}
	}
	// Another path's promise does not keep it: Crossbow promises range.
	if _, err := decode(func(plan *s.Object) {
		repertoireEffect(plan, 1, 1).Set("adaptedAs", []any{"range"})
	}); err == nil || !strings.Contains(err.Error(), `no purchase that adapts "Triple Throw" promises range`) {
		t.Errorf("another technique's promise kept an effect: %v", err)
	}
	// An authored entry lists effects, and each gives a reason.
	if _, err := decode(func(plan *s.Object) {
		repertoire, _ := plan.Get("repertoire")
		repertoire.([]any)[1].(*s.Object).Delete("effects")
	}); err == nil || !strings.Contains(err.Error(), "repertoire.1.effects") {
		t.Errorf("an entry without effects was accepted: %v", err)
	}
	if _, err := decode(func(plan *s.Object) {
		repertoireEffect(plan, 0, 3).Set("reason", "")
	}); err == nil || !strings.Contains(err.Error(), "repertoire.0.effects.3.reason") {
		t.Errorf("an omission without a reason was accepted: %v", err)
	}
	// A capstone may promise five dimensions within its change budget
	// (Luffy 7b on #27 was rejected at four).
	if _, err := decode(func(plan *s.Object) {
		at(plan, "paths", "path2", "milestones", "tier5").(*s.Object).Set("improves", []any{"damage", "attack-rate", "active-damage", "active-attack-rate", "active-duration"})
	}); err != nil {
		t.Errorf("five promises within the budget were rejected: %v", err)
	}
	// The provider schema offers the Definition's promises, never none.
	request, err := unit.DesignPlanRequest(prepared)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		// A combined example beside the speed-only one (#27).
		"a sudden speed increase or barrage can attack faster or fire more shots; a form that raises both strength and speed can hit harder and attack faster",
		"In each repertoire entry's effects, list every effect its cited passages describe",
		"code finds each on a purchase whose technique is that entry, in its plain or active form",
		"including every effect of an entry that no purchase names as its technique; an entry named like the base attack is shown by 0-0-0",
		"reason says how the effect is adapted or why it is omitted",
	} {
		if strings.Count(request.Prompt, want) != 1 {
			t.Errorf("the plan prompt holds %q %d times", want, strings.Count(request.Prompt, want))
		}
	}
	schema := s.Stringify(request.Schema)
	if !strings.Contains(schema, `"adaptedAs":{"type":"array","items":{"type":"string","enum":["damage","attack-rate"`) || strings.Contains(schema, `"enum":["none","damage"`) {
		t.Errorf("adaptedAs schema: %s", schema[strings.Index(schema, `"adaptedAs"`):][:300])
	}
}

// Plans saved before effects existed still load and keep their exact JSON.
func TestPlansWithoutEffectsStillLoad(t *testing.T) {
	data, err := os.ReadFile("../../../../data/reference/captures/dart-monkey.result.json")
	if err != nil {
		t.Fatal(err)
	}
	var value any
	if err := json.Unmarshal(data, &value); err != nil {
		t.Fatal(err)
	}
	result, err := unit.ParseResult(s.FromGoValue(value))
	if err != nil {
		t.Fatal(err)
	}
	plan := result.Run.Draft.DesignPlan
	if plan == nil || len(plan.Repertoire) == 0 || plan.Repertoire[0].Effects != nil {
		t.Fatalf("saved plan %+v", plan)
	}
	if strings.Contains(s.Stringify(s.FromGoValue(result)), `"effects"`) {
		t.Error("a saved plan gained an effects field")
	}
	if issues := unit.PlanEffectIssues(*plan); len(issues) > 0 {
		t.Errorf("a saved plan without effects has issues: %v", issues)
	}
}
