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
