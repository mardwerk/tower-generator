package unit

import (
	"slices"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// benefitSet is what one or more purchases improve, unlock or propose, as
// the exclusive early benefits and path identity rules compare it. Each
// benefit has a key, which the comparison uses, and the text an issue shows.
// The plan and resolved stages key a benefit the same way:
//
//   - an improved dimension is its promise, such as damage or burn;
//   - an unlocked capability that a purchase could also improve, such as
//     splash, a follow-up, a status effect or bonus damage, is that same
//     dimension, so unlocking splash on one path and raising it on another
//     share splash;
//   - another unlock reads "unlock camo" or "unlock delivery-change";
//   - a proposed mechanic reads "proposed mechanic Rebound" and is keyed by
//     its name, ignoring case and spacing.
//
// Lowers, names, prices and amounts are left out, and so is a targeting
// change: a targeting priority is the player's choice, not a benefit.
type benefitSet map[string]string

func (b benefitSet) add(key, text string) {
	if _, ok := b[key]; !ok {
		b[key] = text
	}
}

func (b benefitSet) addAll(other benefitSet) {
	for key, text := range other {
		b.add(key, text)
	}
}

// keys lists the benefit keys in sorted order.
func (b benefitSet) keys() []string {
	out := make([]string, 0, len(b))
	for key := range b {
		out = append(out, key)
	}
	slices.Sort(out)
	return out
}

// text reads the benefits with the given keys, such as "damage and unlock
// camo", or nothing.
func (b benefitSet) text(keys []string) string {
	parts := make([]string, len(keys))
	for i, key := range keys {
		parts[i] = b[key]
	}
	return joinWith(parts, "and", "nothing")
}

// unlockBenefit keys an unlock: a capability the Definition's promises can
// also improve is that dimension, any other unlock reads "unlock X".
func unlockBenefit(unlock string, definition m.Definition) string {
	if slices.Contains(ImprovementsFor(&definition), unlock) {
		return unlock
	}
	return "unlock " + unlock
}

// countsAsUnlock reports an unlock that is a benefit: none and a targeting
// change are not.
func countsAsUnlock(unlock string) bool {
	return unlock != "" && unlock != "none" && unlock != "targeting-change"
}

func (b benefitSet) addUnlock(unlock string, definition m.Definition) {
	if countsAsUnlock(unlock) {
		key := unlockBenefit(unlock, definition)
		b.add(key, key)
	}
}

func (b benefitSet) addProposed(proposed []m.ProposedMechanic) {
	for _, p := range proposed {
		key := "proposed mechanic " + strings.ToLower(strings.Join(strings.Fields(p.Name), " "))
		b.add(key, "proposed mechanic "+strings.Join(strings.Fields(p.Name), " "))
	}
}

// stepBenefits are an early step's improvements and unlocks.
func stepBenefits(step earlyStep, definition m.Definition) benefitSet {
	out := benefitSet{}
	for _, d := range step.improves {
		out.add(d, d)
	}
	for _, u := range step.unlocks {
		out.addUnlock(u, definition)
	}
	return out
}

// plannedBenefits are what a milestone promises: its improvements, its
// unlock and its proposed mechanics.
func plannedBenefits(intent UpgradeIntent, definition m.Definition) benefitSet {
	out := benefitSet{}
	for _, d := range intent.Improves {
		out.add(d, d)
	}
	out.addUnlock(intent.Unlock, definition)
	out.addProposed(intent.ProposedMechanics)
	return out
}

// resolvedBenefits are what a purchase gives in its pure build, measured
// against the pure build before it with the promise check's measures, plus
// the proposed mechanics it carries.
func resolvedBenefits(blueprint *m.Blueprint, definition m.Definition, pathIndex, tier int) benefitSet {
	var before, after m.Selection
	before[pathIndex], after[pathIndex] = tier-1, tier
	b, a := m.ResolveUnchecked(blueprint, before), m.ResolveUnchecked(blueprint, after)
	path := m.PathKeys[pathIndex]
	out := benefitSet{}
	for _, d := range ImprovementsFor(&definition) {
		prior := measures(b, path, d, definition)
		for i, v := range measures(a, path, d, definition) {
			if i < len(prior) && v > prior[i] {
				out.add(d, d)
				break
			}
		}
	}
	for _, u := range UnlocksFor(&definition) {
		if countsAsUnlock(u) && unlockedIntent(b, a, u, pathIndex, tier, blueprint, definition) {
			out.addUnlock(u, definition)
		}
	}
	out.addProposed(blueprint.Paths.At(pathIndex).Tiers.At(tier).ProposedMechanics)
	return out
}

// benefitOwners maps each benefit to the first purchase, in build-code
// order, of the paths other than pathIndex that has it.
func benefitOwners(all [3][5]benefitSet, pathIndex int) map[string]string {
	owners := map[string]string{}
	for other := range m.PathKeys {
		if other == pathIndex {
			continue
		}
		for tier := 1; tier <= 5; tier++ {
			for _, key := range all[other][tier-1].keys() {
				if _, ok := owners[key]; !ok {
					owners[key] = BuildCode(other, tier)
				}
			}
		}
	}
	return owners
}

// sharedText reads benefits with the purchase that also has each, such as
// "damage at 1-x-x and attack-rate at x-x-4".
func sharedText(b benefitSet, keys []string, owners map[string]string) string {
	parts := make([]string, len(keys))
	for i, key := range keys {
		parts[i] = b[key] + " at " + owners[key]
	}
	return joinWith(parts, "and", "nothing")
}
