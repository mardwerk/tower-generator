package unit

import "testing"

// A semantic issue about a purchase whose number was out of bounds, or a
// build that owns it, is dropped, so nothing is judged from a value the
// model did not write; issues about other purchases and builds stay
// (SOL-61-11).
func TestIssuesOnAnOutOfBoundsPurchaseAreDropped(t *testing.T) {
	ref, ok := tierOf([]any{"paths", "path2", "tiers", "tier5", "activeFollowUp", "radius"})
	if !ok || ref != (tierRef{1, 5}) {
		t.Fatalf("tier of x-5-x: %+v %v", ref, ok)
	}
	if _, ok := tierOf([]any{"baseAttack", "stats", "range"}); ok {
		t.Error("the base attack has no purchase")
	}
	invalid := []tierRef{ref}
	for issue, want := range map[string]bool{
		"paths.path2.tiers.tier5.activeFollowUp: needs the tier 4 boost":           true,
		"builds.1-5-0.baseAttack.stats.range: Exceeds the Definition stat ceiling": true,
		"builds.0-5-2.abilities.0: invalid":                                        true,
		"builds.3-2-0.baseAttack.stats.splashRadius: Splash requires pierce":       false,
		"builds.0-4-1.baseAttack.stats.range: fine":                                false,
		"paths.path1.tiers.tier3.changes: Tier must contain at least one effect.":  false,
	} {
		if got := touchesTier(issue, invalid); got != want {
			t.Errorf("%s: dropped %v, want %v", issue, got, want)
		}
	}
}
