package unit

import "testing"

// A follow-up promise costs its own change in the budget floor, since a
// damage raise does not also improve the follow-up (SOL-PR43-02).
func TestFollowUpPromiseCostsItsOwnChange(t *testing.T) {
	definition := DefaultAuthoringDefinition()
	for _, c := range []struct {
		improves []string
		want     int
	}{
		{[]string{"damage"}, 1},
		{[]string{"follow-up"}, 1},
		{[]string{"damage", "follow-up"}, 2},
	} {
		if got := minimumEffects(UpgradeIntent{Improves: c.improves, Unlock: "none"}, definition); got != c.want {
			t.Errorf("%v: floor %d, want %d", c.improves, got, c.want)
		}
	}
}
