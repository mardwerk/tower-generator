package unit_test

import (
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Under the early-identity policy a first purchase cannot add splash, so a
// plan that improves splash at 1-x-x holds only if the base attack splashes.
// The issue names that fix, once, so a repair does not move splash to a
// later purchase without pierce (a failed Luffy run on #27).
func TestEarlyEffectPromisesNeedTheBaseAttack(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	blueprint := *stages.Result.Candidate.Blueprint
	intents := *stages.Draft.Run.DesignPlan.UpgradeIntents
	intents.Path1.Tier1.Improves = append(append([]string{}, intents.Path1.Tier1.Improves...), "splash")
	definition := unit.DefaultAuthoringDefinition()
	var splash []string
	for _, issue := range unit.PlanIntentIssues(blueprint, &intents, definition) {
		if strings.Contains(issue.Message, "improved splash") {
			splash = append(splash, issue.Path+": "+issue.Message)
		}
	}
	want := "paths.path1.tiers.tier1.planIntent: The retained plan promises improved splash at 1-x-x, but the base attack has none"
	if len(splash) != 1 || !strings.HasPrefix(splash[0], want) || !strings.Contains(splash[0], "Give the base attack a splashRadius, with pierce of at least 2") {
		t.Errorf("splash issues %v", splash)
	}
	off := false
	definition.Profile.DesignPolicy.PreserveEarlyAttackIdentity = &off
	for _, issue := range unit.PlanIntentIssues(blueprint, &intents, definition) {
		if strings.Contains(issue.Message, "the base attack has none") {
			t.Errorf("the early-effect issue applied without the early-identity policy: %s", issue.Message)
		}
	}
}

// Under distinctFirstUpgrades two paths whose first two purchases only
// improve the same dimension give no distinct early crosspath value, even
// with different names, prices and amounts: a Luffy plan on #27 made
// x-1-x, x-2-x, x-x-1 and x-x-2 interval-only. The plan is sent back before
// drafting; a third purchase may still lead with that stat.
func TestMirroredEarlyPurchasesAreRejected(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan := *stages.Draft.Run.DesignPlan
	intents := *plan.UpgradeIntents
	plan.UpgradeIntents = &intents
	definition := unit.DefaultAuthoringDefinition()
	mirrored := func() []string {
		var found []string
		for _, issue := range unit.PlanFeasibilityIssues(plan, definition) {
			if strings.Contains(issue.Message, "distinct early crosspath value") {
				found = append(found, issue.Path+": "+issue.Message)
			}
		}
		return found
	}
	// The fixture's early purchases improve pierce, attack-rate and range.
	if found := mirrored(); len(found) != 0 {
		t.Fatalf("distinct early purchases were rejected: %v", found)
	}
	// A stat-led third purchase is not an early purchase.
	intents.Path3.Tier3 = unit.UpgradeIntent{Improves: []string{"attack-rate"}, Unlock: "none", Technique: "Crossbow"}
	if found := mirrored(); len(found) != 0 {
		t.Fatalf("a stat-led third purchase was rejected: %v", found)
	}
	intents.Path3.Tier1 = unit.UpgradeIntent{Improves: []string{"attack-rate"}, Unlock: "none", Technique: "Dart Throw"}
	intents.Path3.Tier2 = unit.UpgradeIntent{Improves: []string{"attack-rate"}, Unlock: "none", Technique: "Dart Throw"}
	want := "upgradeIntents.path3.tier2: x-x-1 and x-x-2 only improve attack-rate, as x-1-x and x-2-x do."
	if found := mirrored(); len(found) != 1 || !strings.HasPrefix(found[0], want) || !strings.Contains(found[0], "such as damage, pierce or range, or personal detection") {
		t.Errorf("mirrored issues %v", found)
	}
	// Another early improvement on either purchase makes the paths distinct.
	intents.Path3.Tier2.Improves = []string{"attack-rate", "range"}
	if found := mirrored(); len(found) != 0 {
		t.Errorf("an added early improvement was rejected: %v", found)
	}
	intents.Path3.Tier2.Improves = []string{"attack-rate"}
	off := definition
	policy := *definition.Profile.DesignPolicy
	policy.DistinctFirstUpgrades = false
	off.Profile.DesignPolicy = &policy
	definition = off
	if found := mirrored(); len(found) != 0 {
		t.Errorf("the check applied without distinctFirstUpgrades: %v", found)
	}
}
