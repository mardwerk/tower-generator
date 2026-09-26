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
	if strings.Contains(s.Stringify(retained), "time-averaged") || strings.Contains(s.Stringify(retained), unit.AgainstCapstone) {
		t.Error("deriving the review evidence changed the retained evidence")
	}
}

func TestSidePurchasesAgainstCapstone(t *testing.T) {
	_, derived := burstEvidence(t)
	for _, test := range []struct {
		path                        int
		from, to                    string
		sideCode, capstoneCode      string
		sidePrice, capstonePrice    float64
		sideDirect, capstoneDirect  float64
		sideGroup, capstoneGroup    float64
		directAtLeast, groupAtLeast bool
	}{
		// 1-x-x adds 1 damage while the Active doubles it half the time:
		// +2 for 100 Gold, as much as x-5-x adds for 20,000 Gold.
		{1, "0-5-0", "1-5-0", "1-x-x", "x-5-x", 100, 20000, 2, 2, 2, 2, true, true},
		// x-x-2 only adds range.
		{1, "0-4-1", "0-4-2", "x-x-2", "x-5-x", 100, 20000, 0, 2, 0, 2, false, false},
		// 4-0-0 → 5-0-0 doubles 4 damage at pierce 2 for 10,000 Gold: +4
		// direct and +8 group. The middle path's x-1-x only adds range.
		{0, "5-0-0", "5-1-0", "x-1-x", "5-x-x", 100, 10000, 0, 4, 0, 8, false, false},
	} {
		against := at(purchase(t, derived, test.path, "crosspaths", test.from, test.to), unit.AgainstCapstone)
		if at(against, "sidePurchase", "code") != test.sideCode || at(against, "capstone", "code") != test.capstoneCode ||
			!near(at(against, "sidePurchase", "price"), test.sidePrice) || !near(at(against, "capstone", "price"), test.capstonePrice) ||
			!near(at(against, unit.TimeAveragedDirect, "sidePurchase"), test.sideDirect) || !near(at(against, unit.TimeAveragedDirect, "capstone"), test.capstoneDirect) ||
			!near(at(against, unit.TimeAveragedGroup, "sidePurchase"), test.sideGroup) || !near(at(against, unit.TimeAveragedGroup, "capstone"), test.capstoneGroup) ||
			at(against, unit.TimeAveragedDirect, unit.SideGainAtLeastCapstone) != test.directAtLeast ||
			at(against, unit.TimeAveragedGroup, unit.SideGainAtLeastCapstone) != test.groupAtLeast {
			t.Errorf("%s → %s: %s", test.from, test.to, s.Stringify(against))
		}
	}
	// Milestones are the path's own purchases and carry no comparison.
	if at(purchase(t, derived, 1, "milestones", "0-4-0", "0-5-0"), unit.AgainstCapstone) != nil {
		t.Error("a milestone carries a side-purchase comparison")
	}
}

// cheapSideEvidence has the shape of a live run on #27: the bottom path's
// x-x-5 adds time-averaged direct +12.64 for 45,000 Gold, and a 140 Gold
// 1-x-x on top of it adds only +1.05. Per 100 Gold the side purchase would
// look far stronger (0.75 against 0.028); in absolute terms the capstone
// adds twelve times as much.
const cheapSideEvidence = `{"paths":[{"path":"path3","name":"Bottom","milestones":[
{"from":[0,0,4],"to":[0,0,5],"incrementalGold":45000,"metricDeltas":{
"direct damage rate":{"before":10,"after":22.64,"change":12.64},
"group damage rate upper bound":{"before":20,"after":45.28,"change":25.28}}}],
"crosspaths":[{"from":[0,0,5],"to":[1,0,5],"incrementalGold":140,"metricDeltas":{
"direct damage rate":{"before":22.64,"after":23.69,"change":1.05},
"group damage rate upper bound":{"before":45.28,"after":47.38,"change":2.1}}}]}],
"limitations":[]}`

// A cheap side purchase that adds less than the capstone is shown as less:
// the evidence states absolute gains and prices, no gain per currency, and
// never marks the side purchase as at least the capstone.
func TestCheapSidePurchaseAddingLessThanTheCapstone(t *testing.T) {
	retained, err := s.Decode([]byte(cheapSideEvidence))
	if err != nil {
		t.Fatal(err)
	}
	derived := unit.ReviewPurchaseEvidence(retained.(*s.Object), "Gold")
	against := at(purchase(t, derived, 0, "crosspaths", "0-0-5", "1-0-5"), unit.AgainstCapstone)
	side, capstone := at(against, unit.TimeAveragedDirect, "sidePurchase"), at(against, unit.TimeAveragedDirect, "capstone")
	if !near(side, 1.05) || !near(capstone, 12.64) || side.(float64) >= capstone.(float64) {
		t.Errorf("direct gains %v and %v, want side +1.05 below capstone +12.64", side, capstone)
	}
	if !near(at(against, unit.TimeAveragedGroup, "sidePurchase"), 2.1) || !near(at(against, unit.TimeAveragedGroup, "capstone"), 25.28) {
		t.Errorf("group gains: %s", s.Stringify(against))
	}
	if at(against, "sidePurchase", "code") != "1-x-x" || at(against, "capstone", "code") != "x-x-5" ||
		!near(at(against, "sidePurchase", "price"), 140) || !near(at(against, "capstone", "price"), 45000) {
		t.Errorf("codes and prices: %s", s.Stringify(against))
	}
	for _, metric := range []string{unit.TimeAveragedDirect, unit.TimeAveragedGroup} {
		if at(against, metric, unit.SideGainAtLeastCapstone) != false {
			t.Errorf("%s: the side purchase is marked as adding at least the capstone's gain", metric)
		}
	}
	text := strings.ToLower(s.Stringify(derived))
	for _, ratio := range []string{"per 100", "per100", "per unit", "0.75", "stronger"} {
		if strings.Contains(text, ratio) {
			t.Errorf("the derived evidence contains %q", ratio)
		}
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
	for _, want := range []string{`"` + unit.TimeAveragedDirect + `"`, `"` + unit.TimeAveragedGroup + `"`, `"` + unit.AgainstCapstone + `"`, `"capstone":{"code":"x-5-x","price":45000}`} {
		if !strings.Contains(text, want) {
			t.Errorf("the review's purchase evidence lacks %s", want)
		}
	}
	if !strings.Contains(prompt, "againstCapstone gives that side purchase's absolute time-averaged direct and group gain and its price in Gold") {
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
	if !strings.Contains(instructions, "againstCapstone gives that side purchase's absolute time-averaged direct and group gain and its price in Berries") {
		t.Error("the review instructions do not name the Definition's currency")
	}
	context, err := s.Decode([]byte(request.Prompt[split+2:]))
	if err != nil {
		t.Fatal(err)
	}
	evidence := s.Stringify(at(context, "purchaseEvidence"))
	if !strings.Contains(evidence, "its price in Berries") {
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
	if !strings.Contains(evidence, "its price in the Definition's currency") {
		t.Error("the derived limitation is not currency-neutral without a currency")
	}
}
