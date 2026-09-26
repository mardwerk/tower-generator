package unit

import (
	"fmt"
	"slices"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// BuildCode names a purchase in top-middle-bottom notation: the second path
// at tier 4 is x-4-x.
func BuildCode(pathIndex, tier int) string {
	parts := []string{"x", "x", "x"}
	parts[pathIndex] = string(rune('0' + tier))
	return strings.Join(parts, "-")
}

// SelectionCode names a concrete build: {3, 1, 0} is 3-1-0.
func SelectionCode(selection m.Selection) string {
	parts := make([]string, len(selection))
	for i, tier := range selection {
		parts[i] = string(rune('0' + tier))
	}
	return strings.Join(parts, "-")
}

var ordinals = []string{"", "first", "second", "third", "fourth", "fifth"}

// BuildShape states the Definition's legal builds for a model: how many paths
// a Unit may buy, how many may pass the crosspath tier, and how many
// crosspath builds code resolves. The Default Definition reads "at most 2
// paths ... the 12 early and 36 advanced crosspath builds".
func BuildShape(d m.Definition) string {
	p := d.Progression
	early, advanced := 0, 0
	for _, selection := range m.AllLegalBuilds(d) {
		bought, beyond := 0, 0
		for _, tier := range selection {
			if tier > 0 {
				bought++
			}
			if tier > p.CrosspathTier {
				beyond++
			}
		}
		switch {
		case bought < 2:
		case beyond == 0:
			early++
		default:
			advanced++
		}
	}
	var illegal []string
	if p.MaxAdvancedPaths < 2 {
		illegal = append(illegal, SelectionCode(m.Selection{p.CrosspathTier + 1, p.CrosspathTier + 1, 0}))
	}
	if p.MaxPurchasedPaths < 3 {
		illegal = append(illegal, SelectionCode(m.Selection{1, 1, 1}))
	}
	text := fmt.Sprintf("A Unit buys at most %d paths, and at most %d of them beyond its %s purchase", p.MaxPurchasedPaths, p.MaxAdvancedPaths, ordinals[p.CrosspathTier])
	if len(illegal) > 0 {
		text += "; " + strings.Join(illegal, " and ") + " are illegal"
	}
	return text + fmt.Sprintf(". Code resolves the %d early and %d advanced crosspath builds; do not output a crosspath tree. A side purchase changes the one attack and, when owned, the boost of that attack.", early, advanced)
}

// ActivationShape states which path may have the Active Ability and from
// which purchase, as the Definition's design policy and rules set it.
func ActivationShape(d m.Definition) string {
	tier := d.Rules.ManualBoostUnlockTier
	policy := d.Profile.DesignPolicy
	switch {
	case policy != nil && policy.ManualAbilityPath.Present && policy.ManualAbilityPath.Null:
		return "No path may have a player-activated ability."
	case policy != nil && policy.ManualAbilityPath.Present:
		index := slices.Index(m.PathKeys, policy.ManualAbilityPath.Value)
		return fmt.Sprintf("Only the %s path may have a player-activated ability, first at %s.", pathPosition(policy.ManualAbilityPath.Value), BuildCode(index, tier))
	}
	return fmt.Sprintf("A player-activated ability may unlock at a path's %s purchase.", ordinals[tier])
}
