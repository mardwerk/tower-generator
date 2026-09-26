package unit

import (
	"fmt"
	"regexp"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// The model review reads the unit through facts that code resolved: every
// legal build with its price and resolved attack. A review that cites a build
// the Definition does not allow is corrected once and then rejected, so it
// cannot report findings about builds such as 3-3-0.

var buildCodePattern = regexp.MustCompile(`\b[0-9]-[0-9]-[0-9]\b`)

// attackFacts is one resolved attack as the review reads it.
func attackFacts(attack m.Attack) *s.Object {
	st := attack.Stats
	facts := s.NewObject().
		Set("delivery", attack.Delivery).
		Set("damageType", attack.DamageType).
		Set("targeting", attack.Targeting).
		Set("damage", st.Damage).
		Set("intervalSeconds", st.IntervalSeconds).
		Set("range", st.Range).
		Set("projectiles", st.Projectiles).
		Set("pierce", st.Pierce).
		Set("maxEnemiesHitPerAttack", st.Projectiles*st.Pierce)
	if st.SplashRadius > 0 {
		facts.Set("splashRadius", st.SplashRadius)
	}
	if st.Projectiles > 1 && attack.Distribution != "" {
		facts.Set("distribution", attack.Distribution)
	}
	if f := attack.FollowUp; f != nil {
		facts.Set("followUp", s.NewObject().Set("count", f.Count).Set("damage", f.DamageMultiplier*st.Damage).Set("radius", f.Radius))
	}
	if statuses := attack.AppliedStatuses(); len(statuses) > 0 {
		facts.Set("statuses", s.FromGoValue(statuses))
	}
	var detects []string
	if attack.Detects != nil {
		detects = *attack.Detects
	} else if attack.Camo {
		detects = []string{"camo"}
	}
	if len(detects) > 0 {
		facts.Set("detects", s.FromGoValue(detects))
	}
	return facts
}

// LegalBuildFacts lists every build the Definition allows with its total
// price, resolved attack and Active Ability.
func LegalBuildFacts(blueprint *m.Blueprint, definition m.Definition) []any {
	var out []any
	for _, selection := range m.AllLegalBuilds(definition) {
		build := m.ResolveUnchecked(blueprint, selection)
		entry := s.NewObject().
			Set("code", fmt.Sprintf("%d-%d-%d", selection[0], selection[1], selection[2])).
			Set("cost", build.CumulativeCost).
			Set("attack", attackFacts(build.BaseAttack))
		for _, ability := range build.Abilities {
			boosted := ability.BoostedAttack.Stats
			entry.Set("activeAbility", s.NewObject().
				Set("name", ability.Name).
				Set("durationSeconds", ability.DurationSeconds).
				Set("cooldownSeconds", ability.CooldownSeconds).
				Set("boostedDamage", boosted.Damage).
				Set("boostedIntervalSeconds", boosted.IntervalSeconds).
				Set("boostedRange", boosted.Range))
		}
		out = append(out, entry)
	}
	return out
}

// illegalBuildCodes lists the build codes a review cites that the
// Definition does not allow, in order of first mention.
func illegalBuildCodes(review SemanticReview, definition m.Definition) []string {
	legal := map[string]bool{}
	for _, selection := range m.AllLegalBuilds(definition) {
		legal[fmt.Sprintf("%d-%d-%d", selection[0], selection[1], selection[2])] = true
	}
	texts := []string{review.Summary}
	for _, f := range review.Findings {
		texts = append(texts, f.Subject, f.Message)
		if f.Action != nil {
			texts = append(texts, *f.Action)
		}
	}
	var out []string
	seen := map[string]bool{}
	for _, code := range buildCodePattern.FindAllString(strings.Join(texts, "\n"), -1) {
		if !legal[code] && !seen[code] {
			seen[code] = true
			out = append(out, code)
		}
	}
	return out
}
