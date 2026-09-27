package unit

import (
	"errors"
	"slices"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Validation completeness (#61, OPUS-NET-61-14, SOL-61-10). Escanor's v34
// design attempt failed only on the mechanics output schema: x-5-x's
// activeFollowUp radius was 0 where the schema requires more than 0. The
// semantic checks never ran on that draft, so the repair fixed the radius,
// and only then did the nine top-path builds whose splash had pierce 1 in
// the same draft fail, with no repair left.
//
// When every schema issue of a mechanics output is a finite number outside
// one of its bounds (s.Issue.NumberBound), the rest of the output parsed, so
// code still runs the semantic checks: on a copy in which each such number
// is moved into its bound, to the limit of an inclusive bound and one past
// the limit of an exclusive one, so a radius of 0 reads as 1. The attempt
// still fails on its schema issues, each of which says what value the other
// checks read, and the repair receives every issue together. The copy is
// diagnostic only: code never publishes, records or repairs it, the repair
// sees the output as the model wrote it, and final validation is as strict
// as before. Partial decoding of other schema failures, such as a missing
// field or a wrong type, is not attempted: the typed checks assume a
// complete blueprint.

// clampedNumber is one number moved into its bound for the semantic checks.
type clampedNumber struct {
	path  []any
	value float64
}

// numbersInBounds returns a copy of value with the number at each issue's
// path moved into the bound it broke. ok is false unless every issue is a
// number outside a bound.
func numbersInBounds(value any, issues []s.Issue) (any, []clampedNumber, bool) {
	if len(issues) == 0 {
		return nil, nil, false
	}
	copied := s.Clone(value)
	var clamped []clampedNumber
	for _, issue := range issues {
		bound, ok := issue.NumberBound()
		if !ok || len(issue.Path) == 0 {
			return nil, nil, false
		}
		current, ok := valueAt(copied, issue.Path).(float64)
		if !ok {
			return nil, nil, false
		}
		next := current
		switch bound.Op {
		case "gte", "lte":
			next = bound.Limit
		case "gt":
			next = bound.Limit + 1
		case "lt":
			next = bound.Limit - 1
		}
		switch parent := valueAt(copied, issue.Path[:len(issue.Path)-1]).(type) {
		case *s.Object:
			key, ok := issue.Path[len(issue.Path)-1].(string)
			if !ok {
				return nil, nil, false
			}
			parent.Set(key, next)
		case []any:
			index, ok := issue.Path[len(issue.Path)-1].(int)
			if !ok {
				return nil, nil, false
			}
			parent[index] = next
		default:
			return nil, nil, false
		}
		clamped = append(clamped, clampedNumber{slices.Clone(issue.Path), next})
	}
	return copied, clamped, true
}

// mechanicsIssues decodes one mechanics output, bound to its plan, and
// lists every issue it has. valid is false when the output fails its schema:
// then its schema issues come first and, when each is a number outside a
// bound, the semantic issues of the copy with those numbers in bounds
// follow. err is set only when the output cannot be read as a validation
// failure.
func mechanicsIssues(output any, request *Request, plan DesignPlan) (blueprint m.Blueprint, issues []string, valid bool, err error) {
	blueprint, budget, err := DecodeForDiagnostics(output, request)
	var validation *s.Error
	if err != nil && !errors.As(err, &validation) {
		return m.Blueprint{}, nil, false, err
	}
	if validation == nil {
		return blueprint, semanticIssues(blueprint, budget, request, plan), true, nil
	}
	return m.Blueprint{}, schemaFailureIssues(output, request, plan, validation), false, nil
}

// schemaFailureIssues lists the issues of an output that fails its schema,
// with the semantic issues of its copy in bounds when every schema issue is
// a number outside a bound.
func schemaFailureIssues(output any, request *Request, plan DesignPlan, validation *s.Error) []string {
	issues := issueStrings(validation.Issues)
	schema, err := ModelOutputSchema(request)
	if err != nil {
		return issues
	}
	_, outputIssues := s.Parse(schema, output)
	copied, clamped, ok := numbersInBounds(output, outputIssues)
	if !ok {
		return issues
	}
	if _, again := s.Parse(schema, copied); len(again) > 0 {
		return issues
	}
	blueprint, budget, err := DecodeForDiagnostics(copied, request)
	if err != nil {
		return issues
	}
	issues = []string{}
	for i, issue := range outputIssues {
		issues = append(issues, issue.PathString()+": "+issue.Message+". The other checks read it as "+s.FormatNumber(clamped[i].value)+".")
	}
	return append(issues, semanticIssues(blueprint, budget, request, plan)...)
}

// semanticIssues are the checks of a blueprint that passed the output
// schema: its change budgets, the typed validation of every legal build and
// the design policy, the plan's promises and the early benefits.
func semanticIssues(blueprint m.Blueprint, budget []s.Issue, request *Request, plan DesignPlan) []string {
	definition := *request.MechanicsDefinition
	budgetPaths := map[string]bool{}
	for _, issue := range budget {
		budgetPaths[s.Issue{Path: issue.Path[:len(issue.Path)-1]}.PathString()+".changes"] = true
	}
	issues := issueStrings(budget)
	for _, issue := range ValidateBlueprintRequest(blueprint, *request) {
		if budgetPaths[issue.Path] && (issue.Message == "Exceeds the Definition change budget." || issue.Message == "Tier must contain at least one effect.") {
			continue
		}
		issues = append(issues, issue.Path+": "+issue.Message)
	}
	for _, issue := range PlanIntentIssues(blueprint, plan.UpgradeIntents, definition) {
		issues = append(issues, issue.Path+": "+issue.Message)
	}
	for _, issue := range EarlyBenefitsIssues(blueprint, plan.UpgradeIntents, definition) {
		issues = append(issues, issue.Path+": "+issue.Message)
	}
	return issues
}
