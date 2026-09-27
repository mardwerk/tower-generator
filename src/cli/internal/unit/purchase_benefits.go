package unit

import (
	"slices"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// benefitSet is what one or more purchases improve or unlock, as the
// exclusive early benefits rule compares it. Each benefit has a key, which
// the comparison uses, and the text an issue shows. The plan and resolved
// stages key a benefit the same way:
//
//   - an improved dimension is its promise, such as damage or burn;
//   - an unlocked capability that a purchase could also improve, such as
//     splash, a follow-up, a status effect or bonus damage, is that same
//     dimension, so unlocking splash on one path and raising it on another
//     share splash;
//   - another unlock reads "unlock camo" or "unlock delivery-change".
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
