package unit_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// reviewContext is the JSON context at the end of a review prompt.
func reviewContext(t *testing.T, checked unit.Checked) (instructions string, context *s.Object) {
	t.Helper()
	prompt := unit.BlueprintReviewRequest(checked).Prompt
	split := strings.LastIndex(prompt, "\n\n")
	value, err := s.Decode([]byte(prompt[split+2:]))
	if err != nil {
		t.Fatal(err)
	}
	return prompt[:split], value.(*s.Object)
}

// The review reads a path whose five prices copy the Definition's reference
// sequence as a fact it can name (reported on #27: a Luffy run priced its
// top and bottom paths 140, 200, 320, 1,800 and 15,000 Gold, exactly the Dart
// Monkey reference). The fact is context for a judgment, not a finding.
func TestReviewContextStatesReferencePriceCopies(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	// The fixture's top path uses the reference prices; the others do not.
	_, context := reviewContext(t, stages.Checked)
	prices := at(context, "referencePrices").(*s.Object)
	if got := s.Stringify(at(prices, "referenceSequence")); got != "[140,200,320,1800,15000]" {
		t.Errorf("reference sequence %s", got)
	}
	want := []bool{true, false, false}
	for index, value := range at(prices, "paths").([]any) {
		if got := at(value, "equalsReferenceSequence"); got != want[index] {
			t.Errorf("path %d equalsReferenceSequence %v, want %v", index+1, got, want[index])
		}
	}
	if got := s.Stringify(at(prices, "facts")); got != `["Top path prices (1-x-x to 5-x-x) equal the reference sequence 140, 200, 320, 1800 and 15000 Gold exactly."]` {
		t.Errorf("facts %s", got)
	}
	for _, f := range stages.Checked.Findings {
		if strings.Contains(f.Message, "reference sequence") {
			t.Errorf("code reported a price finding: %s", f.Message)
		}
	}

	// One price that differs is no copy.
	checked := stages.Checked
	blueprint := *checked.Draft.Candidate.Blueprint
	blueprint.Paths.Path1.Tiers.Tier5.Cost = 16000
	checked.Draft.Candidate.Blueprint = &blueprint
	_, context = reviewContext(t, checked)
	if got := s.Stringify(at(context, "referencePrices", "facts")); got != "[]" {
		t.Errorf("facts for Unit-specific prices %s", got)
	}
	if at(at(context, "referencePrices", "paths").([]any)[0], "equalsReferenceSequence") != false {
		t.Error("a repriced top path still equals the reference sequence")
	}

	// A Definition without a reference scale gives no price facts.
	definition := *checked.Draft.Prepared.Request.MechanicsDefinition
	definition.Profile.ReferenceScale = nil
	checked.Draft.Prepared.Request.MechanicsDefinition = &definition
	if _, context = reviewContext(t, checked); context.Has("referencePrices") {
		t.Error("a Definition without a reference scale has price facts")
	}
}

// The review judges prices and the capstone's reason to buy from the
// purchase evidence and the reference-price fact, as a judgment without a
// threshold; the mechanics prompt asks for prices derived by value, not a
// copied reference sequence.
func TestPriceGuidanceAsksForUnitSpecificPrices(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	instructions, _ := reviewContext(t, stages.Checked)
	for _, want := range []string{
		"whether each fourth and fifth purchase's price fits what it adds, and whether the capstone keeps a reason to buy at its price",
		"Compare absolute gains and prices: a cheap side purchase that adds as much absolute time-averaged gain as the capstone, or more, is evidence that the capstone is overpriced or adds too little, and only then does a side purchase count against the capstone.",
		"A side purchase that adds less absolute gain than the capstone is no evidence against it.",
		"Never use a gain per unit of currency as capstone evidence by itself",
		"referencePrices.facts states each path whose five prices equal the Definition's reference sequence exactly",
		"not thresholds or a required multiplier",
	} {
		if !strings.Contains(instructions, want) {
			t.Errorf("the review instructions lack %q", want)
		}
	}
	if strings.Count(instructions, "referencePrices") != 1 {
		t.Error("more than one review instruction owns the reference-price rule")
	}
	// One instruction owns the side-purchase rule, and no instruction offers
	// a normalized ratio or calls a cheap side purchase stronger.
	if strings.Count(instructions, "againstCapstone") != 1 {
		t.Error("more than one review instruction owns the side-purchase rule")
	}
	for _, ratio := range []string{"per 100", "per100", "stronger"} {
		if strings.Contains(instructions, ratio) {
			t.Errorf("the review instructions contain %q", ratio)
		}
	}

	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan, _ := fixture.JSON("plan")
	mechanics, _ := fixture.JSON("mechanics")
	model := &fixture.Model{Outputs: []any{plan, mechanics}}
	if _, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options()); err != nil {
		t.Fatal(err)
	}
	prompt := model.Requests[1].Prompt
	for _, want := range []string{
		"Price each incremental purchase for this Unit's own payoff.",
		"describe reference units, not this one: use them as scale by value, pricing a purchase above or below a reference purchase as its gain compares, and do not copy a reference sequence.",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the mechanics prompt lacks %q", want)
		}
	}
	if strings.Contains(prompt, "not a curve to copy") {
		t.Error("the mechanics prompt keeps the older price wording")
	}
}
