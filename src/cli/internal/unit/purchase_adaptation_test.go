package unit_test

import (
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// A purchase that adapts a named technique shows the plan's source-backed
// description of that adaptation, which the review then checks against the
// resolved mechanics. A Luffy review on #27 failed 3-x-x and x-x-3, whose
// sheets listed stat changes without the inflated or compressed limb they
// adapt; a code-written restatement of the stats did not show the source
// connection either.
func TestPurchasesShowThePlannedAdaptation(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan := *stages.Draft.Run.DesignPlan
	intents := *plan.UpgradeIntents
	plan.UpgradeIntents = &intents
	intents.Path1.Tier3.Technique = "Gear 3"
	plan.Paths.Path1.Milestones.Tier3 = "3-x-x Gear 3 inflates Luffy's arm into a giant limb: the punch lands slower and harder, and its impact hits enemies around the target."
	intents.Path3.Tier3.Technique = "Gear 4"
	plan.Paths.Path3.Milestones.Tier3 = "x-x-3: Gear 4 compresses the limb and releases it, so each hit rebounds onto two nearby enemies"
	for _, test := range []struct {
		path, tier int
		want       string
	}{
		{0, 3, "Gear 3 inflates Luffy's arm into a giant limb: the punch lands slower and harder, and its impact hits enemies around the target."},
		{2, 3, "Gear 4 compresses the limb and releases it, so each hit rebounds onto two nearby enemies."},
		// The build code the plan starts with is dropped; the sentence keeps its capital.
		{0, 4, "Throws the ball faster, multiplies its pierce and switches it to normal damage so it can hit Lead."},
		// A purchase of the base attack alone: its effects say it all.
		{0, 1, ""},
	} {
		if got := unit.PurchaseAdaptation(&plan, test.path, test.tier); got != test.want {
			t.Errorf("path %d tier %d: %q, want %q", test.path+1, test.tier, got, test.want)
		}
	}
	// Plans made before techniques were typed get no line.
	intents.Path1.Tier3.Technique = ""
	if got := unit.PurchaseAdaptation(&plan, 0, 3); got != "" {
		t.Errorf("a purchase without a technique: %q", got)
	}
}
