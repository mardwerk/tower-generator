package unit

import (
	"fmt"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Review verdicts are the review's judgments that code requires one of per
// subject, outside the eight free-form findings: one per whole-technique
// omission in designPlan.omittedTechniques and one per path's third purchase.
// The v32 Luffy review spent all eight findings on purchases and judged no
// omission and no third purchase (SOL-61-08). Code checks that the review
// returns exactly one verdict per subject and records each one in the Result
// as a model Finding: a pass as a pass, so the owner can read every verdict,
// and a fail or unresolved verdict as an open Finding.

const (
	// OmissionVerdictRule is the rule of the Finding that records the
	// review's verdict on one whole-technique omission.
	OmissionVerdictRule = "omission-verdict"
	// PathIdentityVerdictRule is the rule of the Finding that records the
	// review's verdict on one path's third purchase, its Path identity.
	PathIdentityVerdictRule = "path-identity-verdict"
)

// OmissionVerdict is the review's verdict on one whole-technique omission.
type OmissionVerdict struct {
	Technique string   `json:"technique"`
	Outcome   string   `json:"outcome"`
	Reason    string   `json:"reason"`
	Action    *string  `json:"action"`
	Evidence  []string `json:"evidence"`
}

// ThirdPurchaseVerdict is the review's verdict on one path's third purchase.
type ThirdPurchaseVerdict struct {
	Build    string   `json:"build"`
	Outcome  string   `json:"outcome"`
	Reason   string   `json:"reason"`
	Action   *string  `json:"action"`
	Evidence []string `json:"evidence"`
}

var verdictOutcome = s.Enum("pass", "fail", "unresolved")

// omissionVerdictSchema and thirdPurchaseVerdictSchema read any well-formed
// verdict; the review request narrows the technique and build to its own
// subjects, and code checks that each subject has exactly one.
func omissionVerdictSchema(technique, evidence s.Schema) *s.ObjectSchema {
	return s.StrictObject(
		s.F("technique", technique),
		s.F("outcome", verdictOutcome),
		s.F("reason", text()),
		s.F("action", s.Nullable(text())),
		s.F("evidence", s.Array(evidence)),
	)
}

func thirdPurchaseVerdictSchema(build, evidence s.Schema) *s.ObjectSchema {
	return s.StrictObject(
		s.F("build", build),
		s.F("outcome", verdictOutcome),
		s.F("reason", text()),
		s.F("action", s.Nullable(text())),
		s.F("evidence", s.Array(evidence)),
	)
}

// verdictSubjects are what a review must give a verdict on: the distinct
// names of the plan's whole-technique omissions with their rank, and each
// path's third purchase by build code, with its name and technique.
type verdictSubjects struct {
	omissions      []PlanOmission
	thirdPurchases []thirdPurchase
}

type thirdPurchase struct {
	path             int
	build, name, key string
	technique        string
}

func reviewVerdictSubjects(checked Checked) verdictSubjects {
	var subjects verdictSubjects
	if plan := checked.Draft.Run.DesignPlan; plan != nil {
		for _, omission := range plan.OmittedTechniques {
			duplicate := false
			for _, listed := range subjects.omissions {
				duplicate = duplicate || sameTechnique(listed.Name, omission.Name)
			}
			if !duplicate && strings.TrimSpace(omission.Name) != "" {
				subjects.omissions = append(subjects.omissions, omission)
			}
		}
	}
	candidate := checked.Draft.Candidate
	if candidate.Blueprint == nil || len(candidate.Paths) != len(m.PathKeys) {
		return subjects
	}
	for index, key := range m.PathKeys {
		purchase := thirdPurchase{path: index, build: BuildCode(index, 3), key: key}
		for _, tier := range candidate.Paths[index].Tiers {
			if tier.Tier == 3 {
				purchase.name = tier.Name
			}
		}
		if plan := checked.Draft.Run.DesignPlan; plan != nil && plan.UpgradeIntents != nil {
			purchase.technique = strings.TrimSpace(plan.UpgradeIntents.At(index).At(3).Technique)
		}
		subjects.thirdPurchases = append(subjects.thirdPurchases, purchase)
	}
	return subjects
}

// context lists the verdicts the review must give, for requiredVerdicts in
// the review context.
func (v verdictSubjects) context() *s.Object {
	omissions, purchases := []any{}, []any{}
	for _, omission := range v.omissions {
		entry := s.NewObject().Set("technique", omission.Name)
		if omission.Importance != "" {
			entry.Set("importance", omission.Importance)
		}
		omissions = append(omissions, entry)
	}
	for _, purchase := range v.thirdPurchases {
		entry := s.NewObject().Set("build", purchase.build).Set("path", purchase.key).Set("name", purchase.name)
		if purchase.technique != "" {
			entry.Set("technique", purchase.technique)
		}
		purchases = append(purchases, entry)
	}
	return s.NewObject().Set("omissions", omissions).Set("thirdPurchases", purchases)
}

// schemaFields narrow the review's verdict fields to exactly its subjects.
func (v verdictSubjects) schemaFields(evidence s.Schema) []s.Field {
	var names, builds []string
	for _, omission := range v.omissions {
		names = append(names, omission.Name)
	}
	for _, purchase := range v.thirdPurchases {
		builds = append(builds, purchase.build)
	}
	omission := omissionVerdictSchema(text(), evidence)
	if len(names) > 0 {
		omission = omissionVerdictSchema(s.Enum(names...), evidence)
	}
	purchase := thirdPurchaseVerdictSchema(text(), evidence)
	if len(builds) > 0 {
		purchase = thirdPurchaseVerdictSchema(s.Enum(builds...), evidence)
	}
	return []s.Field{
		s.F("omissionVerdicts", s.Array(omission).Length(len(names))),
		s.F("thirdPurchaseVerdicts", s.Array(purchase).Length(len(builds))),
	}
}

// problems lists what keeps a review's verdicts from being exactly one per
// subject: a missing, repeated or unknown verdict.
func (v verdictSubjects) problems(review SemanticReview) []string {
	var problems []string
	counts := make([]int, len(v.omissions))
	for _, verdict := range review.OmissionVerdicts {
		known := false
		for i, omission := range v.omissions {
			if sameTechnique(verdict.Technique, omission.Name) {
				counts[i]++
				known = true
			}
		}
		if !known {
			problems = append(problems, fmt.Sprintf("a verdict on %q, which designPlan.omittedTechniques does not list", verdict.Technique))
		}
	}
	for i, omission := range v.omissions {
		switch {
		case counts[i] == 0:
			problems = append(problems, fmt.Sprintf("no verdict on the omission of %q", omission.Name))
		case counts[i] > 1:
			problems = append(problems, fmt.Sprintf("%d verdicts on the omission of %q", counts[i], omission.Name))
		}
	}
	counts = make([]int, len(v.thirdPurchases))
	for _, verdict := range review.ThirdPurchaseVerdicts {
		known := false
		for i, purchase := range v.thirdPurchases {
			if strings.TrimSpace(verdict.Build) == purchase.build {
				counts[i]++
				known = true
			}
		}
		if !known {
			problems = append(problems, fmt.Sprintf("a verdict on %q, which is no path's third purchase", verdict.Build))
		}
	}
	for i, purchase := range v.thirdPurchases {
		switch {
		case counts[i] == 0:
			problems = append(problems, "no verdict on the third purchase "+purchase.build)
		case counts[i] > 1:
			problems = append(problems, fmt.Sprintf("%d verdicts on the third purchase %s", counts[i], purchase.build))
		}
	}
	return problems
}

// incompleteVerdicts is the error for a review whose verdicts are not
// exactly one per subject; like other malformed review output it fails the
// review and keeps the checked draft.
func incompleteVerdicts(problems []string) *ModelError {
	message := "The model review did not return exactly one verdict per whole-technique omission and per third purchase: it gave " + strings.Join(problems, "; ") + ". The draft is retained. Retry the review or choose another model."
	return &ModelError{Message: message, Failure: &Failure{Code: CodeOutputInvalid, Message: message, Stage: "review"}}
}

// findings records every verdict as a model Finding, in subject order: an
// omission's verdict under OmissionVerdictRule and a third purchase's under
// PathIdentityVerdictRule. A pass stays a pass; a fail or unresolved verdict
// is an open Finding like any other.
func (v verdictSubjects) findings(review SemanticReview) []Finding {
	out := []Finding{}
	for index, omission := range v.omissions {
		for _, verdict := range review.OmissionVerdicts {
			if !sameTechnique(verdict.Technique, omission.Name) {
				continue
			}
			subject := "omittedTechniques, " + omission.Name
			if omission.Importance != "" {
				subject += " (" + omission.Importance + ")"
			}
			out = append(out, verdictFinding(fmt.Sprintf("verdict.omission.%d", index+1), "coverage", verdict.Outcome, subject, OmissionVerdictRule, verdict.Reason, verdict.Action, verdict.Evidence))
			break
		}
	}
	for _, purchase := range v.thirdPurchases {
		for _, verdict := range review.ThirdPurchaseVerdicts {
			if strings.TrimSpace(verdict.Build) != purchase.build {
				continue
			}
			subject := m.PathKeys[purchase.path] + ", " + purchase.build
			if purchase.name != "" {
				subject += " " + purchase.name
			}
			out = append(out, verdictFinding("verdict.third-purchase."+purchase.key, "conflict", verdict.Outcome, subject, PathIdentityVerdictRule, verdict.Reason, verdict.Action, verdict.Evidence))
			break
		}
	}
	return out
}

func verdictFinding(id, category, outcome, subject, rule, reason string, action *string, evidence []string) Finding {
	severity := map[string]string{"pass": "info", "unresolved": "warning", "fail": "error"}[outcome]
	if evidence == nil {
		evidence = []string{}
	}
	if action != nil && strings.TrimSpace(*action) == "" {
		action = nil
	}
	return Finding{ID: id, Method: "model", Category: category, Severity: severity, Outcome: outcome, Subject: subject, Rule: rule, Message: reason, Evidence: evidence, Action: action}
}

// IsReviewVerdict reports a Finding that records a review verdict on an
// omission or a third purchase.
func IsReviewVerdict(finding Finding) bool {
	return finding.Method == "model" && (finding.Rule == OmissionVerdictRule || finding.Rule == PathIdentityVerdictRule)
}

// checkVerdicts rejects a review whose verdicts cite an unknown document or
// are not exactly one per subject.
func checkVerdicts(checked Checked, review SemanticReview, documents map[string]bool) error {
	var evidence [][]string
	for _, verdict := range review.OmissionVerdicts {
		evidence = append(evidence, verdict.Evidence)
	}
	for _, verdict := range review.ThirdPurchaseVerdicts {
		evidence = append(evidence, verdict.Evidence)
	}
	for _, ids := range evidence {
		for _, id := range ids {
			if !documents[id] {
				return invalidReview()
			}
		}
	}
	if problems := reviewVerdictSubjects(checked).problems(review); len(problems) > 0 {
		return incompleteVerdicts(problems)
	}
	return nil
}
