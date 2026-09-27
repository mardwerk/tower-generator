package unit

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// BuildCode names a purchase in top-middle-bottom notation: the second path
// at tier 4 is x-4-x.
func BuildCode(pathIndex, tier int) string { return m.BuildCode(pathIndex, tier) }

// SelectionCode names a concrete build: {3, 1, 0} is 3-1-0.
func SelectionCode(selection m.Selection) string {
	parts := make([]string, len(selection))
	for i, tier := range selection {
		parts[i] = string(rune('0' + tier))
	}
	return strings.Join(parts, "-")
}

// NameWithoutBuildCode is a base attack or purchase name without the build
// code it opens with, as "0-0-0 Gum-Gum Pistol", "x-4-x: Jet Boost" or "In
// build 3-x-x, Gigant Pistol" (leadingBuildCode), or with a dash after it.
// The Unit sheet names the base attack and each purchase by its build code
// already, so the v32 Luffy plan's base "0-0-0 Gum-Gum Pistol" printed as
// "0-0-0: 0-0-0 Gum-Gum Pistol" (#61, OPUS-NET-61-10). A name that is only
// a build code is kept.
func NameWithoutBuildCode(name string) string {
	stripped := leadingBuildCode.ReplaceAllString(name, "")
	if stripped == name {
		return name
	}
	stripped = strings.TrimSpace(strings.TrimLeft(stripped, "-\u2013\u2014 "))
	if stripped == "" {
		return name
	}
	first, size := utf8.DecodeRuneInString(stripped)
	return string(unicode.ToUpper(first)) + stripped[size:]
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
