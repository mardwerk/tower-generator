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

// clauseBuildCode matches "In build CODE," or "In CODE," later in a planned
// change. The unit shows the code already, so code drops it where it opens a
// sentence or clause, and inside a clause only with "build", as in "the arm
// inflates, and in build 3-x-x, the punch lands harder"; "the range it gained
// in 1-x-x, and" keeps its code.
var clauseBuildCode = regexp.MustCompile(`(?i)\bin\s+(?:(?:the\s+)?(build)\s+)?[0-9x]-[0-9x]-[0-9x]\b(?:\s+purchase)?\s*,\s*`)

// dropClauseBuildCodes removes the "In build CODE," clauses clauseBuildCode
// matches and capitalizes a sentence the removal opens.
func dropClauseBuildCodes(text string) string {
	var out strings.Builder
	last := 0
	for _, match := range clauseBuildCode.FindAllStringSubmatchIndex(text, -1) {
		before := strings.TrimRightFunc(text[:match[0]], unicode.IsSpace)
		end := ""
		if before != "" {
			end = before[len(before)-1:]
		}
		if !strings.ContainsAny(end, ".;:!?") && match[2] < 0 {
			continue
		}
		out.WriteString(text[last:match[0]])
		last = match[1]
		if strings.ContainsAny(end, ".!?") && last < len(text) {
			first, size := utf8.DecodeRuneInString(text[last:])
			out.WriteRune(unicode.ToUpper(first))
			last += size
		}
	}
	out.WriteString(text[last:])
	return out.String()
}

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
	text := strings.TrimSpace(dropClauseBuildCodes(leadingBuildCode.ReplaceAllString(plan.Paths.At(pathIndex).Milestones.At(tier), "")))
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
