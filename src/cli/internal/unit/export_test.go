package unit

import m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"

// CapstoneOrdering is the review's purchaseComparisonOrdering for one
// blueprint, for tests that build their own unit.
func CapstoneOrdering(blueprint *m.Blueprint) []any {
	return capstoneOrdering(m.CompareCapstonePurchasesWith(blueprint, nil), blueprint)
}

// LegacySourceTechniques is a request's sourceTechniques in the form
// requests were prepared with under Default v30 and earlier.
func LegacySourceTechniques(request *Request) []SourceTechnique {
	return sourceTechniques(request, true)
}

// CorrectionItems are the items a targeted plan correction names for a
// parsed plan.
func CorrectionItems(plan DesignPlan, request *Request) []any { return correctionItems(plan, request) }
