package unit

import (
	"errors"
	"slices"
	"strconv"
	"strings"

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
// the limit of an exclusive one, only so the rest decodes. Semantic issues
// about a purchase with such a number, or a build that owns it, are
// dropped, so nothing is judged from a value the model did not write
// (SOL-61-11). The attempt still fails on its schema issues, and the repair
// receives every issue together. The copy is
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
	// The copy in bounds only lets the rest decode. No semantic issue may
	// judge a purchase whose value was out of bounds, or a build that owns
	// it, so no issue is invented from a value the model did not write
	// (SOL-61-11); the base attack is in every build, so a number out of
	// bounds there leaves only the schema issues.
	var invalid []tierRef
	for _, c := range clamped {
		ref, ok := tierOf(c.path)
		if !ok {
			return issues
		}
		invalid = append(invalid, ref)
	}
	issues = []string{}
	for _, issue := range outputIssues {
		issues = append(issues, issue.PathString()+": "+issue.Message+". The other checks ran on every purchase and build that does not include this value.")
	}
	for _, issue := range semanticIssues(blueprint, budget, request, plan) {
		if !touchesTier(issue, invalid) {
			issues = append(issues, issue)
		}
	}
	return issues
}

// tierRef is one purchase, by path index (0 to 2) and tier (1 to 5).
type tierRef struct{ path, tier int }

// tierOf finds the purchase an output path lies in, as in
// paths.path2.tiers.tier5.activeFollowUp.radius. A path outside a purchase,
// such as the base attack, has none.
func tierOf(path []any) (tierRef, bool) {
	if len(path) < 4 || path[0] != "paths" || path[2] != "tiers" {
		return tierRef{}, false
	}
	pathKey, _ := path[1].(string)
	tierKey, _ := path[3].(string)
	p := slices.Index(m.PathKeys[:], pathKey)
	t := slices.Index(m.TierKeys[:], tierKey)
	if p < 0 || t < 0 {
		return tierRef{}, false
	}
	return tierRef{p, t + 1}, true
}

// touchesTier reports an issue about one of the purchases, or about a legal
// build that owns one of them, as builds.1-5-0 owns x-5-x.
func touchesTier(issue string, tiers []tierRef) bool {
	for _, ref := range tiers {
		if strings.HasPrefix(issue, "paths."+m.PathKeys[ref.path]+".tiers."+m.TierKeys[ref.tier-1]) {
			return true
		}
		if rest, ok := strings.CutPrefix(issue, "builds."); ok {
			code, _, _ := strings.Cut(rest, ".")
			if parts := strings.Split(code, "-"); len(parts) == 3 {
				if n, err := strconv.Atoi(parts[ref.path]); err == nil && n >= ref.tier {
					return true
				}
			}
		}
	}
	return false
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
