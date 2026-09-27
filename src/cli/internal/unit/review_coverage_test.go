package unit_test

import (
	"os"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/research"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// luffyV36Checked is the v36 Luffy Result 706999ee as a checked draft: its
// retained plan and blueprint, prepared from the same Sources under the
// Default, compiled and evaluated again (OPUS-NET-61-16).
func luffyV36Checked(t *testing.T) unit.Checked {
	t.Helper()
	return luffyChecked(t, "v36", unit.Run{ID: "706999ee-b344-4be1-85be-4e171b78007f", ModelID: "openai/gpt-6-luna", StartedAt: "2026-09-27T13:35:00Z", CompletedAt: "2026-09-27T13:38:00Z"})
}

// luffyChecked is a saved Luffy Result of one Default revision as a checked
// draft: its retained plan and blueprint (testdata/luffy-<revision>.*.json),
// prepared from the same Sources under the Default with the given required
// concepts, compiled and evaluated again.
func luffyChecked(t *testing.T, revision string, run unit.Run, required ...string) unit.Checked {
	t.Helper()
	decode := func(name string) any {
		data, err := os.ReadFile("testdata/" + name)
		if err != nil {
			t.Fatal(err)
		}
		value, err := s.Decode(data)
		if err != nil {
			t.Fatal(err)
		}
		return value
	}
	sources, err := research.ParseSources(decode("luffy.sources.json"))
	if err != nil {
		t.Fatal(err)
	}
	prepared, err := sources.Prepare(unit.DefaultProfile())
	if err != nil {
		t.Fatal(err)
	}
	if len(required) > 0 {
		if prepared, err = unit.Prepare(s.FromGoValue(requiring(prepared.Request, required...))); err != nil {
			t.Fatal(err)
		}
	}
	definition := *prepared.Request.MechanicsDefinition
	var plan unit.DesignPlan
	if err := s.ParseInto(unit.DesignPlanSchemaV2, decode("luffy-"+revision+".plan.json"), &plan); err != nil {
		t.Fatal(err)
	}
	var blueprint mechanics.Blueprint
	if err := s.ParseInto(mechanics.BlueprintSchemaFor(definition), decode("luffy-"+revision+".blueprint.json"), &blueprint); err != nil {
		t.Fatal(err)
	}
	candidate, err := unit.CompileBlueprint(blueprint, prepared.Request)
	if err != nil {
		t.Fatal(err)
	}
	evaluation, err := unit.EvaluateUnitDesign(blueprint, &plan, definition)
	if err != nil {
		t.Fatal(err)
	}
	run.DesignPlan, run.DesignEvaluation = &plan, evaluation
	draft := unit.Draft{SchemaVersion: prepared.SchemaVersion, Kind: "draft", Prepared: prepared, Candidate: candidate, Run: run}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	return checked
}

// requiredVerdicts is the requiredVerdicts object of a review request.
func requiredVerdicts(t *testing.T, checked unit.Checked) *s.Object {
	t.Helper()
	prompt := unit.BlueprintReviewRequest(checked).Prompt
	value, err := s.Decode([]byte(prompt[strings.LastIndex(prompt, "\n\n")+2:]))
	if err != nil {
		t.Fatal(err)
	}
	return at(value, "requiredVerdicts").(*s.Object)
}

// The v36 Luffy x-x-5 costs 21500 Gold and adds 0.387 time-averaged direct
// and 2.709 group damage per second over 0-0-4; adding 1-x-x to 0-0-4 costs
// 140 Gold and adds 1.548 and 3.096. Its fifth purchase verdict gets that
// comparison beside its subject. x-5-x Jet Bazooka, 35000 Gold for 12.66
// and 63.86, has no such cheaper side purchase (SOL-61-13).
func TestFifthPurchaseVerdictGetsPayoffAndPrice(t *testing.T) {
	checked := luffyV36Checked(t)
	fifth := at(requiredVerdicts(t, checked), "fifthPurchases").([]any)
	payoff := func(index int) *s.Object { return at(fifth[index], "payoff").(*s.Object) }
	number := func(value any) string { return s.FormatNumber(float64(int(value.(float64)*1000+0.5)) / 1000) }

	bottom := payoff(2)
	if at(fifth[2], "build") != "x-x-5" || at(bottom, "fourthPurchase") != "0-0-4" || at(bottom, "price") != 21500.0 ||
		number(at(bottom, "gains", unit.TimeAveragedDirect)) != "0.387" || number(at(bottom, "gains", unit.TimeAveragedGroup)) != "2.709" {
		t.Errorf("x-x-5 payoff %s", s.Stringify(bottom))
	}
	var codes []string
	for _, side := range at(bottom, "cheaperSidePurchases").([]any) {
		codes = append(codes, at(side, "code").(string)+"="+s.Stringify(at(side, "atLeastCapstoneGains")))
	}
	if got := strings.Join(codes, ","); got != "1-x-x=true,2-x-x=false,x-1-x=false,x-2-x=false" {
		t.Errorf("x-x-5 side purchases %s", got)
	}
	facts := s.Stringify(at(bottom, "facts"))
	if want := `["Adding 1-x-x to 0-0-4 costs 140 Gold and adds 1.548 time-averaged direct and 3.096 group damage per second, at least the 0.387 and 2.709 that x-x-5 adds for 21500 Gold."]`; facts != want {
		t.Errorf("x-x-5 facts %s", facts)
	}

	middle := payoff(1)
	if at(fifth[1], "build") != "x-5-x" || at(middle, "price") != 35000.0 ||
		number(at(middle, "gains", unit.TimeAveragedDirect)) != "12.658" || number(at(middle, "gains", unit.TimeAveragedGroup)) != "63.857" {
		t.Errorf("x-5-x payoff %s", s.Stringify(middle))
	}
	if facts := s.Stringify(at(middle, "facts")); facts != "[]" {
		t.Errorf("x-5-x has a cheaper side purchase with at least its gains: %s", facts)
	}

	prompt := unit.BlueprintReviewRequest(checked).Prompt
	if !strings.Contains(prompt, unit.CapstonePayoffRule) || !strings.Contains(prompt, "a follow-up, volley or splash that adds damage is judged by the damage it adds") {
		t.Error("the review is not asked to judge payoff and price in the fifth purchase verdict")
	}
	if strings.Contains(prompt, "tier5ToFourthPriceRatio") || strings.Contains(s.Stringify(bottom), "atio") {
		t.Error("the payoff carries a ratio")
	}
}

// Every purchase-level proposed mechanic of the v36 Luffy Unit gets a
// required verdict, unresolved or fail, outside the eight findings: 3-x-x's
// "Massive-scale impact" restates its typed splash and x-x-4's "Continuous
// momentum" has no player-visible effect, which the review should fail
// (SOL-61-13).
func TestEveryProposedMechanicGetsAVerdict(t *testing.T) {
	checked := luffyV36Checked(t)
	var listed []string
	for _, p := range at(requiredVerdicts(t, checked), "proposals").([]any) {
		listed = append(listed, at(p, "build").(string)+" "+at(p, "proposal").(string))
	}
	want := "3-x-x Massive-scale impact, x-3-x Jet Pistol launch, x-5-x Jet Bazooka charge and stance, x-x-3 Python redirection, x-x-4 Continuous momentum, x-x-5 Bending trajectory"
	if got := strings.Join(listed, ", "); got != want {
		t.Errorf("required proposal verdicts %s", got)
	}
	request := unit.BlueprintReviewRequest(checked)
	for _, want := range []string{
		`{"build":"x-x-4","path":"path3","name":"Python Continuous Stretch","technique":"Gear 4 Python","proposal":"Continuous momentum","effect":"The stretching attack continues without losing momentum while in Boundman form; the Definition does not represent trajectory momentum.","sourceIds":["source4:18"]}`,
		unit.ProposalOutcomeRule,
		"it fails when it restates what a supported change of the same purchase already does, when it defines no effect a player could see, such as a momentum or state with no consequence in play",
		unit.ProposalDistinctRule,
	} {
		if !strings.Contains(request.Prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	schema := s.Stringify(request.Schema)
	for _, want := range []string{
		`"proposalVerdicts":{"minItems":6,"maxItems":6,"type":"array"`,
		`"build":{"type":"string","enum":["3-x-x","x-3-x","x-5-x","x-x-3","x-x-4","x-x-5"]}`,
		`"proposal":{"type":"string","enum":["Massive-scale impact","Jet Pistol launch","Jet Bazooka charge and stance","Python redirection","Continuous momentum","Bending trajectory"]}`,
		`"outcome":{"type":"string","enum":["unresolved","fail"]}`,
	} {
		if !strings.Contains(schema, want) {
			t.Errorf("the review schema lacks %s", want)
		}
	}
}
