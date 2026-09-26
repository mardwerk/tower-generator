package unit

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
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

// factIssues checks the resolved values that review findings cite against
// legalBuilds, so a finding cannot rest on a misread number.
func factIssues(review SemanticReview, blueprint *m.Blueprint, definition m.Definition) []string {
	builds := map[string]*s.Object{}
	for _, entry := range LegalBuildFacts(blueprint, definition) {
		object := entry.(*s.Object)
		code, _ := object.Get("code")
		builds[code.(string)] = object
	}
	var out []string
	for _, finding := range review.Findings {
		for _, fact := range finding.Facts {
			build, field := strings.TrimSpace(fact.Build), strings.TrimSpace(fact.Field)
			entry, ok := builds[build]
			if !ok {
				out = append(out, fmt.Sprintf("%s cites build %s, which legalBuilds does not list.", finding.ID, build))
				continue
			}
			var value any = entry
			for _, key := range strings.Split(field, ".") {
				object, ok := value.(*s.Object)
				if !ok {
					value = nil
					break
				}
				value, _ = object.Get(key)
			}
			if value == nil {
				out = append(out, fmt.Sprintf("%s cites %s %s, a field legalBuilds does not list for that build.", finding.ID, build, field))
			} else if !sameFact(value, fact.Value) {
				out = append(out, fmt.Sprintf("%s cites %s %s as %s, but it is %s.", finding.ID, build, field, strings.TrimSpace(fact.Value), s.Stringify(value)))
			}
		}
	}
	return out
}

// sameFact compares a cited value with the resolved one; numbers may differ
// by rounding in the last listed digit.
func sameFact(actual any, cited string) bool {
	cited = strings.TrimSpace(cited)
	switch value := actual.(type) {
	case float64:
		number, err := strconv.ParseFloat(cited, 64)
		return err == nil && math.Abs(number-value) <= math.Max(1e-9, 1e-3*math.Abs(value))
	case string:
		return strings.EqualFold(cited, value)
	case bool:
		return strings.EqualFold(cited, strconv.FormatBool(value))
	}
	return cited == s.Stringify(actual)
}
