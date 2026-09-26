package unit_test

import (
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// A purchase that adapts a named technique says how its typed changes
// represent it, naming only what they do. A Luffy review on #27 failed 3-x-x
// and x-x-3, whose sheets listed interval, damage and splash, or a
// follow-up, without the Gear 3 and Gear 4 they adapt.
func TestPurchasesSayWhatTheyAdapt(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	blueprint := *stages.Result.Candidate.Blueprint
	plan := *stages.Draft.Run.DesignPlan
	intents := *plan.UpgradeIntents
	plan.UpgradeIntents = &intents
	definition := unit.DefaultAuthoringDefinition()

	blueprint.Paths.Path1.Tiers.Tier3.Changes = []m.Change{
		{Kind: "stat", Target: "base", Stat: "intervalSeconds", Operation: "multiply", Number: 1.5},
		{Kind: "stat", Target: "base", Stat: "damage", Operation: "add", Number: 2},
		{Kind: "stat", Target: "base", Stat: "splashRadius", Operation: "set", Number: 8},
	}
	intents.Path1.Tier3.Technique = "Gear 3"
	blueprint.Paths.Path3.Tiers.Tier3.Changes = []m.Change{
		{Kind: "followUp", Target: "base", FollowUp: &m.FollowUp{Name: "Rebound", Count: 2, DamageMultiplier: 0.5, Radius: 10}},
	}
	intents.Path3.Tier3.Technique = "Gear 4"
	for _, test := range []struct {
		path, tier int
		want       string
	}{
		{0, 3, "Adapts Gear 3: the attack becomes slower and heavier and now hits enemies around the target."},
		{2, 3, "Adapts Gear 4: the attack is followed by a strike on up to 2 nearby enemies."},
		{1, 4, "Adapts Fan Club: the attack becomes faster and the Unit gains the Fan Club Frenzy Active Ability."},
		{1, 5, "Adapts Fan Club: Fan Club Frenzy becomes stronger."},
		// A purchase of the base attack alone: its effects say it all.
		{0, 1, ""},
	} {
		if got := unit.PurchaseAdaptation(&blueprint, &plan, definition, test.path, test.tier); got != test.want {
			t.Errorf("path %d tier %d: %q, want %q", test.path+1, test.tier, got, test.want)
		}
	}
	// Plans made before techniques were typed get no line.
	intents.Path1.Tier3.Technique = ""
	if got := unit.PurchaseAdaptation(&blueprint, &plan, definition, 0, 3); got != "" {
		t.Errorf("a purchase without a technique: %q", got)
	}
}
