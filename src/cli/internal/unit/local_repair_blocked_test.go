package unit

import (
	"testing"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// The scoped repair is skipped only when the build before the purchase
// already has what the unlock adds, with unlockedIntent's own presence
// semantics (SOL-81-01).
func TestUnlockBlockedMatchesUnlockedIntent(t *testing.T) {
	definition := DefaultAuthoringDefinition()
	magnitude := 2.0
	v2 := func(statuses ...m.StatusApplication) m.Attack {
		detects := []string{}
		return m.Attack{Statuses: &statuses, Detects: &detects, Stats: m.AttackStats{Projectiles: 1, Pierce: 2}}
	}
	burning := v2(m.StatusApplication{Effect: "burn", Magnitude: &magnitude, Seconds: 3})
	if !burning.IsV2() {
		t.Fatal("the test attack is not version 2")
	}
	cases := []struct {
		name    string
		before  m.Attack
		after   m.Attack
		intent  string
		blocked bool
	}{
		// An attack that already burns cannot unlock burn again.
		{"an owned v2 status", burning, burning, "burn", true},
		{"a v2 status not yet owned", v2(), burning, "burn", false},
		// Distinct targets with one projectile are not yet a volley; the
		// next purchase can add shots and unlock it.
		{"distinct targets with one projectile", withVolley(v2(), 1), withVolley(v2(), 3), "distinct-volley", false},
		{"a distinct volley already owned", withVolley(v2(), 3), withVolley(v2(), 3), "distinct-volley", true},
	}
	for _, c := range cases {
		before := m.Build{BaseAttack: c.before}
		after := m.Build{BaseAttack: c.after}
		if got := unlockBlocked(before, c.intent, 0, 3, nil, definition); got != c.blocked {
			t.Errorf("%s: blocked %v, want %v", c.name, got, c.blocked)
		}
		// Whenever the unlock is not blocked and the next purchase adds
		// the capability, unlockedIntent credits it.
		if !c.blocked && !unlockedIntent(before, after, c.intent, 0, 3, nil, definition) {
			t.Errorf("%s: not blocked, but unlockedIntent does not credit the unlock", c.name)
		}
	}
	// A manual boost unlocks only at the Definition's unlock tier, and not
	// when the path already has its Active Ability.
	unlockTier := definition.Rules.ManualBoostUnlockTier
	if !unlockBlocked(m.Build{}, "manual-boost", 1, unlockTier+1, nil, definition) {
		t.Error("a manual boost is not blocked away from its unlock tier")
	}
	if unlockBlocked(m.Build{}, "manual-boost", 1, unlockTier, nil, definition) {
		t.Error("a manual boost is blocked at its unlock tier on a path without one")
	}
	owned := m.Build{Abilities: []m.ResolvedAbility{{Path: "path2"}}}
	if !unlockBlocked(owned, "manual-boost", 1, unlockTier, nil, definition) {
		t.Error("a manual boost is not blocked when the path already has its Active Ability")
	}
}

func withVolley(a m.Attack, projectiles float64) m.Attack {
	a.Distribution = "distinct-targets"
	a.Stats.Projectiles = projectiles
	return a
}
