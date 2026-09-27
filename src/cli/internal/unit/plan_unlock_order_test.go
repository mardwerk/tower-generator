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

// escanorPath3 sets one bottom path milestone of Escanor a's saved plan
// output, which keeps its upgradeIntents.
func escanorPath3(plan any, tier string, improves []any, unlock string) {
	at(plan, "upgradeIntents", "path3", tier).(*s.Object).Set("improves", improves).Set("unlock", unlock)
}

// Escanor a's v39 plan promises improves splash at x-x-3 and unlock splash
// at x-x-5 (OPUS-NET-61-26). The plan check rejects that, alone, on x-x-5,
// as an item of the targeted correction (SOL-61-21).
func TestEscanorV39PlanUnlocksWhatX3Improves(t *testing.T) {
	run := escanorV39(t)
	issues := planIssues(t, run.planOutput, &run.prepared.Request)
	message := "x-x-5 promises unlock splash, but x-x-3 already promises improves splash. " + unit.UnlockOrderRule + " Move the unlock of splash to x-x-3, or change x-x-5's promise from unlock splash to improves splash."
	if len(issues) != 1 || issues[0] != "plan-correctable upgradeIntents.path3.tier5: "+message {
		t.Fatalf("the v39 plan's issues:\n%s", strings.Join(issues, "\n"))
	}

	// The next call is the targeted correction, not a full re-plan. Its
	// plan unlocks splash at x-x-3 and gives x-x-5 a stun beside the
	// splash and burn it develops; it is accepted, and the third call is
	// the mechanics call, which this test does not script.
	corrected := s.Clone(run.planOutput)
	escanorPath3(corrected, "tier3", []any{"splash"}, "splash")
	escanorPath3(corrected, "tier5", []any{"splash", "burn"}, "stun")
	model := &fixture.Model{Outputs: []any{run.planOutput, corrected}}
	_, err := unit.DraftUnit(context.Background(), run.prepared, model, fixture.Options())
	var modelErr *unit.ModelError
	if !errors.As(err, &modelErr) || modelErr.Evidence == nil || len(model.Requests) != 3 || !strings.Contains(model.Requests[2].Prompt, "Implement the supplied designPlan") {
		t.Fatalf("%d calls: %v", len(model.Requests), err)
	}
	attempts := modelErr.Evidence.Attempts
	var got []string
	for _, attempt := range attempts {
		got = append(got, attempt.Purpose)
	}
	if strings.Join(got, ",") != "plan,plan-correction,design" {
		t.Errorf("attempts %v", got)
	}
	if len(attempts[0].Issues) != 1 || attempts[0].Issues[0] != "paths.path3.milestones.tier5: "+message {
		t.Errorf("the plan attempt's issues: %q", attempts[0].Issues)
	}
	if len(attempts[1].Issues) != 0 {
		t.Errorf("the correction's issues: %q", attempts[1].Issues)
	}
	prompt := model.Requests[1].Prompt
	for _, want := range []string{
		"The previous plan failed only on the items below.",
		`"path":"paths.path3.milestones.tier5"`,
		`"problem":"x-x-5 promises unlock splash, but x-x-3 already promises improves splash. ` + unit.UnlockOrderRule + `"`,
		"Move the unlock to x-x-3: set its unlock to splash in place of none and keep its other promises, and take unlock splash from x-x-5, which keeps improves splash when it still develops it.",
		"Change x-x-5's promise from unlock splash to improves splash, and say in its change that it develops the splash x-x-3 adds.",
		"x-x-5 must still add new behavior or access: The fifth purchase of every path must add a supported behavior or access",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the correction lacks %q", want)
		}
	}
	if strings.Contains(prompt, "Correct this invalid design plan") {
		t.Error("the late unlock is corrected by a full re-plan")
	}

	// The other allowed correction, x-x-5 promising improves splash with
	// another unlock for its new behavior, is accepted too.
	improved := s.Clone(run.planOutput)
	escanorPath3(improved, "tier5", []any{"splash", "burn"}, "stun")
	if issues := planIssues(t, improved, &run.prepared.Request); len(issues) != 0 {
		t.Errorf("x-x-5 improving the splash x-x-3 improves: %v", issues)
	}
}

// A path that unlocks splash first and improves it later passes, and so
// does the unlock of another capability after an improved splash: the
// check compares promise names and infers nothing from stats.
func TestUnlockOrderComparesPromiseNames(t *testing.T) {
	run := escanorV39(t)
	request := &run.prepared.Request
	first := s.Clone(run.planOutput)
	escanorPath3(first, "tier3", []any{}, "splash")
	escanorPath3(first, "tier5", []any{"splash", "burn"}, "stun")
	if issues := planIssues(t, first, request); len(issues) != 0 {
		t.Errorf("unlock splash at x-x-3, improves splash at x-x-5: %v", issues)
	}
	followUp := s.Clone(run.planOutput)
	escanorPath3(followUp, "tier5", []any{"splash"}, "follow-up")
	if issues := planIssues(t, followUp, request); len(issues) != 0 {
		t.Errorf("unlock follow-up at x-x-5 after improves splash at x-x-3: %v", issues)
	}

	definition := *request.MechanicsDefinition
	plan := parsedPlan(t, run.planOutput)
	path3 := plan.UpgradeIntents.At(2)
	// bonus-damage unlocks a bonus against another enemy property, so it
	// may follow an improved bonus-damage.
	path3.At(3).Improves, path3.At(5).Unlock = []string{"bonus-damage"}, "bonus-damage"
	if correctable, other := unit.UnlockOrderIssues(plan, definition); len(correctable)+len(other) != 0 {
		t.Errorf("unlock bonus-damage after improves bonus-damage: %s %s", messages(correctable), messages(other))
	}
	// A capability unlocked twice fails as before, on the later purchase,
	// and the correction may only make that promise an improvement.
	path3.At(3).Improves, path3.At(5).Unlock = []string{"splash"}, "burn"
	correctable, other := unit.UnlockOrderIssues(plan, definition)
	if len(other) != 0 || messages(correctable) != "upgradeIntents.path3.tier5: x-x-5 promises unlock burn, but x-x-4 already promises unlock burn. "+unit.UnlockOrderRule+" Change x-x-5's promise from unlock burn to improves burn." {
		t.Errorf("burn unlocked twice: %s %s", messages(correctable), messages(other))
	}
	items := s.Stringify(s.FromGoValue(unit.CorrectionItems(plan, request)))
	if !strings.Contains(items, "Change x-x-5's promise from unlock burn to improves burn") || strings.Contains(items, "Move the unlock") {
		t.Errorf("the correction items: %s", items)
	}
	// A capability the Definition cannot improve, unlocked twice, needs a
	// full plan.
	path3.At(4).Unlock, path3.At(5).Unlock = "camo", "camo"
	correctable, other = unit.UnlockOrderIssues(plan, definition)
	if len(correctable) != 0 || messages(other) != "upgradeIntents.path3.tier5: x-x-5 promises unlock camo, but x-x-4 already promises unlock camo. "+unit.UnlockOrderRule+" Promise unlock camo only at x-x-4, and give x-x-5 another unlock or none." {
		t.Errorf("camo unlocked twice: %s %s", messages(correctable), messages(other))
	}
}

// The plan prompt states the rule once, with what code rejects.
func TestUnlockOrderRuleReachesThePlan(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan, err := unit.DesignPlanRequest(stages.Prepared)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(plan.Prompt, unit.UnlockOrderRule) != 1 || !strings.Contains(plan.Prompt, unit.UnlockOrderRule+" For example, when x-x-3 promises improves splash") {
		t.Error("the plan prompt does not state the rule once")
	}
}
