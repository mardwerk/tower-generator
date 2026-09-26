package unit_test

import (
	"math"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// burstUnit is a small hand-built Unit. Its base attack deals 2 damage every
// second with pierce 1. The middle path reaches 4 damage at x-3-x, unlocks
// an Active at x-4-x (triple damage for 10 s every 40 s: peak 12, duty 0.25)
// and extends it to 20 s at x-5-x for 20,000 Gold (duty 0.5). The top
// path's x-1-x adds 1 damage for 100 Gold; the bottom path's first
// purchase halves the interval for 100 Gold.
func burstUnit() m.Blueprint {
	stat := func(name, operation string, value float64) m.Change {
		return m.Change{Kind: "stat", Target: "base", Stat: name, Operation: operation, Number: value}
	}
	tier := func(name string, cost float64, changes ...m.Change) m.Tier {
		return m.Tier{Name: name, Cost: cost, Changes: changes}
	}
	path := func(name string, tiers ...m.Tier) m.Path {
		return m.Path{Name: name, Theme: name + " theme.", Rationale: name + " rationale.", SourceFactIndices: []int{0},
			Tiers: m.Tiers{Tier1: tiers[0], Tier2: tiers[1], Tier3: tiers[2], Tier4: tiers[3], Tier5: tiers[4]}}
	}
	return m.Blueprint{
		Name: "Burst", Role: "Test unit.", Weakness: "Slow base.",
		SourceFacts:        []m.SourceFact{{DocumentID: "source", Quote: "A quoted source passage."}},
		ConstraintCoverage: []m.Coverage{},
		BaseAttack: m.Attack{Name: "Punch", Cost: 200, Delivery: "projectile", DamageType: "normal", Targeting: "first",
			Stats: m.AttackStats{Damage: 2, IntervalSeconds: 1, Range: 30, Pierce: 1, Projectiles: 1}},
		Paths: m.Paths{
			Path1: path("Power",
				tier("Heavier", 100, stat("damage", "add", 1)),
				tier("Longer", 100, stat("range", "add", 4)),
				tier("Hammer", 300, stat("damage", "add", 1)),
				tier("Wide", 1000, stat("pierce", "add", 1)),
				tier("Titan", 10000, stat("damage", "multiply", 2)),
			),
			Path2: path("Burst",
				tier("Reach", 100, stat("range", "add", 2)),
				tier("Further", 100, stat("range", "add", 2)),
				tier("Strong", 500, stat("damage", "add", 2)),
				tier("Burst", 2000, m.Change{Kind: "unlockBoost", Target: "base", Boost: &m.Boost{Name: "Burst", DurationSeconds: 10, CooldownSeconds: 40, DamageMultiplier: 3, IntervalMultiplier: 1}}),
				tier("Long Burst", 20000, m.Change{Kind: "modifyBoost", Target: "base", Stat: "durationSeconds", Operation: "add", Number: 10}),
			),
			Path3: path("Speed",
				tier("Quick", 100, stat("intervalSeconds", "multiply", 0.5)),
				tier("Sight", 100, stat("range", "add", 1)),
				tier("Heavy", 300, stat("damage", "add", 1)),
				tier("Heavier", 1000, stat("damage", "add", 1)),
				tier("Heaviest", 10000, stat("damage", "add", 1)),
			),
		},
		Proposals:          []m.Proposal{},
		ReservedTechniques: []m.Proposal{},
	}
}

func burstEvidence(t *testing.T) (retained, derived *s.Object) {
	t.Helper()
	blueprint := burstUnit()
	definition := m.DefaultDefinition()
	if issues := m.ValidateTyped(&blueprint, definition); len(issues) > 0 {
		t.Fatalf("the hand-built unit is invalid: %v", issues)
	}
	retained, err := unit.EvaluateUnitDesign(blueprint, nil, definition)
	if err != nil {
		t.Fatal(err)
	}
	return retained, unit.ReviewPurchaseEvidence(retained, definition.Profile.Currency)
}

// purchase finds the milestone or crosspath comparison from one build to another.
func purchase(t *testing.T, evidence *s.Object, path int, list, from, to string) *s.Object {
	t.Helper()
	items, _ := at(at(evidence, "paths").([]any)[path], list).([]any)
	for _, value := range items {
		code := func(key string) string {
			parts := []string{}
			for _, tier := range at(value, key).([]any) {
				parts = append(parts, s.FormatNumber(tier.(float64)))
			}
			return strings.Join(parts, "-")
		}
		if code("from") == from && code("to") == to {
			return value.(*s.Object)
		}
	}
	t.Fatalf("no %s comparison %s → %s", list, from, to)
	return nil
}

func near(value any, want float64) bool {
	got, ok := value.(float64)
	return ok && math.Abs(got-want) <= 1e-12
}

func TestTimeAveragedActiveOutput(t *testing.T) {
	if got := m.TimeAveraged(4, 12, 0.25); got != 6 {
		t.Errorf("time-averaged rate %v, want 6", got)
	}
	retained, derived := burstEvidence(t)
	for _, test := range []struct {
		list, from, to        string
		before, after, change float64
	}{
		// Unlocking: 4 ordinary, then 12 for a quarter of the cycle.
		{"milestones", "0-3-0", "0-4-0", 4, 6, 2},
		// The capstone doubles the duty fraction: 12 × 0.5 + 4 × 0.5.
		{"milestones", "0-4-0", "0-5-0", 6, 8, 2},
		// A side purchase raises both rates: 15 × 0.5 + 5 × 0.5.
		{"crosspaths", "0-5-0", "1-5-0", 8, 10, 2},
		// Halving the interval doubles both rates.
		{"crosspaths", "0-4-0", "0-4-1", 6, 12, 6},
	} {
		for _, key := range []string{unit.TimeAveragedDirect, unit.TimeAveragedGroup} {
			delta := at(purchase(t, derived, 1, test.list, test.from, test.to), "metricDeltas", key)
			if !near(at(delta, "before"), test.before) || !near(at(delta, "after"), test.after) || !near(at(delta, "change"), test.change) {
				t.Errorf("%s → %s %s: %s", test.from, test.to, key, s.Stringify(delta))
			}
		}
	}
	deltas := at(purchase(t, derived, 1, "milestones", "0-4-0", "0-5-0"), "metricDeltas").(*s.Object).Keys()
	if strings.Join(deltas[len(deltas)-3:], "|") != "active duty fraction|"+unit.TimeAveragedDirect+"|"+unit.TimeAveragedGroup {
		t.Errorf("the time-averaged rates do not follow the duty fraction: %v", deltas)
	}
	if !near(at(at(derived, "paths").([]any)[1], "capstoneComparison", "tier5", "metrics", unit.TimeAveragedDirect), 8) {
		t.Error("the capstone comparison lacks the time-averaged rate")
	}
	// A path without an Active gains no time-averaged rows of its own.
	if at(purchase(t, derived, 0, "milestones", "3-0-0", "4-0-0"), "metricDeltas", unit.TimeAveragedDirect) != nil {
		t.Error("a purchase without an Active shows a time-averaged rate")
	}
	// The retained evidence is not changed, so check still compares it exactly.
	if strings.Contains(s.Stringify(retained), "time-averaged") || strings.Contains(s.Stringify(retained), unit.Per100CurrencyAgainstCapstone) {
		t.Error("deriving the review evidence changed the retained evidence")
	}
}

func TestSidePurchasesAgainstCapstone(t *testing.T) {
	_, derived := burstEvidence(t)
	for _, test := range []struct {
		path             int
		from, to, code   string
		side, capstone   float64
		sideGroupGain    float64
		capstoneGroupPer float64
	}{
		// x-1-x: +2 time-averaged for 100 Gold; x-5-x: +2 for 20,000 Gold.
		{1, "0-5-0", "1-5-0", "x-5-x", 2, 0.01, 2, 0.01},
		// x-x-2 only adds range.
		{1, "0-4-1", "0-4-2", "x-5-x", 0, 0.01, 0, 0.01},
		// 4-0-0 → 5-0-0 doubles 4 damage at pierce 2 for 10,000 Gold: +4
		// direct and +8 group. The middle path's x-1-x only adds range.
		{0, "5-0-0", "5-1-0", "5-x-x", 0, 0.04, 0, 0.08},
	} {
		against := at(purchase(t, derived, test.path, "crosspaths", test.from, test.to), unit.Per100CurrencyAgainstCapstone)
		if at(against, "capstone") != test.code ||
			!near(at(against, unit.TimeAveragedDirect, "sidePurchase"), test.side) || !near(at(against, unit.TimeAveragedDirect, "capstone"), test.capstone) ||
			!near(at(against, unit.TimeAveragedGroup, "sidePurchase"), test.sideGroupGain) || !near(at(against, unit.TimeAveragedGroup, "capstone"), test.capstoneGroupPer) {
			t.Errorf("%s → %s: %s", test.from, test.to, s.Stringify(against))
		}
	}
	// Milestones are the path's own purchases and carry no comparison.
	if at(purchase(t, derived, 1, "milestones", "0-4-0", "0-5-0"), unit.Per100CurrencyAgainstCapstone) != nil {
		t.Error("a milestone carries a side-purchase comparison")
	}
}

// The review reads the derived purchase evidence; the draft keeps its own.
func TestReviewContextCarriesDerivedPurchaseEvidence(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prompt := unit.BlueprintReviewRequest(stages.Checked).Prompt
	context, err := s.Decode([]byte(prompt[strings.LastIndex(prompt, "\n\n")+2:]))
	if err != nil {
		t.Fatal(err)
	}
	evidence := at(context, "purchaseEvidence")
	if s.Canonical(evidence) != s.Canonical(unit.ReviewPurchaseEvidence(stages.Checked.Draft.Run.DesignEvaluation, "Gold")) {
		t.Error("the review context does not carry the derived purchase evidence")
	}
	text := s.Stringify(evidence)
	for _, want := range []string{`"` + unit.TimeAveragedDirect + `"`, `"` + unit.TimeAveragedGroup + `"`, `"` + unit.Per100CurrencyAgainstCapstone + `"`, `"capstone":"x-5-x"`} {
		if !strings.Contains(text, want) {
			t.Errorf("the review's purchase evidence lacks %s", want)
		}
	}
	if !strings.Contains(prompt, "per100CurrencyAgainstCapstone gives that side purchase's time-averaged gain per 100 Gold") {
		t.Error("the review prompt does not explain the derived evidence")
	}
}

// A Definition with another currency gets review guidance and derived
// evidence in that currency. Only the retained price fields, whose saved
// names are incrementalGold and totalGold, keep the word.
func TestReviewContextUsesTheDefinitionCurrency(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	checked := stages.Checked
	definition := *checked.Draft.Prepared.Request.MechanicsDefinition
	definition.Profile.Currency = "Berries"
	checked.Draft.Prepared.Request.MechanicsDefinition = &definition
	request := unit.BlueprintReviewRequest(checked)
	split := strings.LastIndex(request.Prompt, "\n\n")
	instructions := request.System + "\n\n" + request.Prompt[:split]
	if strings.Contains(instructions, "Gold") {
		t.Error("the review instructions name Gold for a Definition in Berries")
	}
	if !strings.Contains(instructions, "per100CurrencyAgainstCapstone gives that side purchase's time-averaged gain per 100 Berries") {
		t.Error("the review instructions do not name the Definition's currency")
	}
	context, err := s.Decode([]byte(request.Prompt[split+2:]))
	if err != nil {
		t.Fatal(err)
	}
	evidence := s.Stringify(at(context, "purchaseEvidence"))
	if !strings.Contains(evidence, "gain per 100 Berries") {
		t.Error("the derived limitation does not name the Definition's currency")
	}
	if strings.Contains(strings.NewReplacer(`"incrementalGold"`, "", `"totalGold"`, "").Replace(evidence), "Gold") {
		t.Error("the review's purchase evidence names Gold for a Definition in Berries")
	}
}

// Without a named currency the derived limitation stays neutral.
func TestDerivedEvidenceWithoutACurrencyIsNeutral(t *testing.T) {
	retained, _ := burstEvidence(t)
	evidence := s.Stringify(unit.ReviewPurchaseEvidence(retained, ""))
	if !strings.Contains(evidence, "gain per 100 of the Definition's currency") {
		t.Error("the derived limitation is not currency-neutral without a currency")
	}
}
