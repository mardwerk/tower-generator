package unit

import (
	"errors"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

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
func CorrectionItems(plan DesignPlan, request *Request) []any {
	return correctionItems(plan, request, nil)
}

// CapstoneCorrection is the capstone correction of a plan output that
// DecodeDesignPlan rejects, its prompt addition and its merge of a
// corrected output; ok is false when none applies.
func CapstoneCorrection(output any, request *Request) (prompt string, merge func(any) any, ok bool) {
	_, err := DecodeDesignPlan(output, request)
	var validation *s.Error
	if !errors.As(err, &validation) || !onlyCapstone(validation.Issues) {
		return "", nil, false
	}
	correction, ok := capstoneCorrectionFor(output, request, validation.Issues)
	return correction.prompt, func(corrected any) any { return correction.merge(output, corrected) }, ok
}

// MechanicsIssues lists every issue of one mechanics output bound to its
// plan, as a design or repair attempt reports them, and whether the output
// passed its schema.
func MechanicsIssues(output any, request *Request, plan DesignPlan) ([]string, bool) {
	_, issues, valid, err := mechanicsIssues(BindDesignPlan(output, plan), request, plan)
	if err != nil {
		return []string{err.Error()}, false
	}
	return issues, valid
}

// LocalRepair is the local repair of a bound mechanics output with these
// issues, or nil when none applies.
func LocalRepair(request *Request, previous any, issues []string) (*TierRepair, error) {
	return localRepairFor(request, previous, issues)
}

// CapstoneStatusEffects is the review's statusEffects of each path's fifth
// purchase, by path key, for tests that build their own unit.
func CapstoneStatusEffects(blueprint *m.Blueprint, vocabulary *m.Vocabulary) map[string][]any {
	return capstoneStatusEffects(blueprint, vocabulary)
}
