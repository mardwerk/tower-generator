package unit

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// leadingBuildCode matches the build code a planned change opens with, as
// "3-x-x", "x-x-3:" or "In build 1-x-x,"; the unit already shows the code.
var leadingBuildCode = regexp.MustCompile(`(?i)^\s*(?:(?:in|at|for)\s+)?(?:(?:the\s+)?build\s+)?[0-9x]-[0-9x]-[0-9x]\b(?:\s+purchase)?[\s:,.;]*`)

// PurchaseAdaptation is the retained plan's description of how a purchase
// adapts its source technique, such as "Replaces the dart with a heavier
// spiked ball that deals more damage, reaches farther and pierces far more
// enemies, at a slower throw." The plan wrote it from the cited sources
// before the numbers were chosen, so it is a claim: the unit shows it beside
// the resolved effects, and the review checks it against them and the
// evidence (reported on #27: Luffy's 3-x-x and x-x-3 listed stat changes
// without the inflated or compressed limb they adapt). It is empty when the
// plan names no technique for the purchase, or only the base attack, which
// the purchase's effects already describe.
func PurchaseAdaptation(plan *DesignPlan, pathIndex, tier int) string {
	if plan == nil || plan.UpgradeIntents == nil {
		return ""
	}
	technique := strings.TrimSpace(plan.UpgradeIntents.At(pathIndex).At(tier).Technique)
	if technique == "" || strings.EqualFold(technique, strings.TrimSpace(plan.Base.Name)) {
		return ""
	}
	text := strings.TrimSpace(leadingBuildCode.ReplaceAllString(plan.Paths.At(pathIndex).Milestones.At(tier), ""))
	if text == "" {
		return ""
	}
	first, size := utf8.DecodeRuneInString(text)
	text = string(unicode.ToUpper(first)) + text[size:]
	if !strings.HasSuffix(text, ".") {
		text += "."
	}
	return text
}
