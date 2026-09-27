package unit

import (
	"fmt"
	"regexp"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Review verdicts are the review's judgments that code requires one of per
// subject, outside the eight free-form findings: one per whole-technique
// omission in designPlan.omittedTechniques, one per path's third purchase and
// one per path's fifth purchase. The v32 Luffy review spent all eight
// findings on purchases and judged no omission and no third purchase
// (SOL-61-08). The v34 Luffy x-5-x and x-x-5 both bought a follow-up, which
// the review did not judge together (OPUS-NET-61-13): a fifth purchase
// verdict judges each capstone against the other two capstones and its
// path's identity, with no typed ban on two capstones sharing a capability
// kind (SOL-61-10). Code checks that the review returns exactly one verdict
// per subject and records each one in the Result as a model Finding: a pass
// as a pass, so the owner can read every verdict, and a fail or unresolved
// verdict as an open Finding.

const (
	// OmissionVerdictRule is the rule of the Finding that records the
	// review's verdict on one whole-technique omission.
	OmissionVerdictRule = "omission-verdict"
	// PathIdentityVerdictRule is the rule of the Finding that records the
	// review's verdict on one path's third purchase, its Path identity.
	PathIdentityVerdictRule = "path-identity-verdict"
	// CapstoneVerdictRule is the rule of the Finding that records the
	// review's verdict on one path's fifth purchase, judged with the other
	// paths' fifth purchases and its path's identity.
	CapstoneVerdictRule = "capstone-verdict"
)

// OmissionVerdict is the review's verdict on one whole-technique omission.
type OmissionVerdict struct {
	Technique string   `json:"technique"`
	Outcome   string   `json:"outcome"`
	Reason    string   `json:"reason"`
	Action    *string  `json:"action"`
	Evidence  []string `json:"evidence"`
}

// PurchaseVerdict is the review's verdict on one purchase, a path's third or
// fifth, by its build code.
type PurchaseVerdict struct {
	Build    string   `json:"build"`
	Outcome  string   `json:"outcome"`
	Reason   string   `json:"reason"`
	Action   *string  `json:"action"`
	Evidence []string `json:"evidence"`
}

// ThirdPurchaseVerdict is the review's verdict on one path's third purchase.
type ThirdPurchaseVerdict = PurchaseVerdict

// FifthPurchaseVerdict is the review's verdict on one path's fifth purchase.
type FifthPurchaseVerdict = PurchaseVerdict

var verdictOutcome = s.Enum("pass", "fail", "unresolved")

// omissionVerdictSchema and purchaseVerdictSchema read any well-formed
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

func purchaseVerdictSchema(build, evidence s.Schema) *s.ObjectSchema {
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
// path's third and fifth purchase by build code, with its name and
// technique.
type verdictSubjects struct {
	omissions      []PlanOmission
	thirdPurchases []purchaseSubject
	fifthPurchases []purchaseSubject
}

type purchaseSubject struct {
	path             int
	build, name, key string
	technique        string
}

// purchaseKind is one kind of purchase verdict: the tier it judges, the
// field it is returned in, the rule and ID of its Finding, and the words its
// problems use.
type purchaseKind struct {
	tier                int
	field, rule, prefix string
	category, noun      string
}

var (
	thirdPurchaseKind = purchaseKind{3, "thirdPurchaseVerdicts", PathIdentityVerdictRule, "verdict.third-purchase.", "conflict", "third purchase"}
	fifthPurchaseKind = purchaseKind{5, "fifthPurchaseVerdicts", CapstoneVerdictRule, "verdict.fifth-purchase.", "conflict", "fifth purchase"}
)

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
	purchase := func(index, number int) purchaseSubject {
		purchase := purchaseSubject{path: index, build: BuildCode(index, number), key: m.PathKeys[index]}
		for _, tier := range candidate.Paths[index].Tiers {
			if tier.Tier == number {
				purchase.name = tier.Name
			}
		}
		if plan := checked.Draft.Run.DesignPlan; plan != nil && plan.UpgradeIntents != nil {
			purchase.technique = strings.TrimSpace(plan.UpgradeIntents.At(index).At(number).Technique)
		}
		return purchase
	}
	for index := range m.PathKeys {
		subjects.thirdPurchases = append(subjects.thirdPurchases, purchase(index, 3))
		subjects.fifthPurchases = append(subjects.fifthPurchases, purchase(index, 5))
	}
	return subjects
}

// purchases are the subjects of one kind of purchase verdict.
func (v verdictSubjects) purchases(kind purchaseKind) []purchaseSubject {
	if kind.tier == 5 {
		return v.fifthPurchases
	}
	return v.thirdPurchases
}

// purchaseVerdicts are a review's verdicts of one kind.
func purchaseVerdicts(review SemanticReview, kind purchaseKind) []PurchaseVerdict {
	if kind.tier == 5 {
		return review.FifthPurchaseVerdicts
	}
	return review.ThirdPurchaseVerdicts
}

// context lists the verdicts the review must give, for requiredVerdicts in
// the review context.
func (v verdictSubjects) context() *s.Object {
	omissions := []any{}
	for _, omission := range v.omissions {
		entry := s.NewObject().Set("technique", omission.Name)
		if omission.Importance != "" {
			entry.Set("importance", omission.Importance)
		}
		omissions = append(omissions, entry)
	}
	purchases := func(subjects []purchaseSubject) []any {
		out := []any{}
		for _, purchase := range subjects {
			entry := s.NewObject().Set("build", purchase.build).Set("path", purchase.key).Set("name", purchase.name)
			if purchase.technique != "" {
				entry.Set("technique", purchase.technique)
			}
			out = append(out, entry)
		}
		return out
	}
	return s.NewObject().Set("omissions", omissions).Set("thirdPurchases", purchases(v.thirdPurchases)).Set("fifthPurchases", purchases(v.fifthPurchases))
}

// schemaFields narrow the review's verdict fields to exactly its subjects.
func (v verdictSubjects) schemaFields(evidence s.Schema) []s.Field {
	var names []string
	for _, omission := range v.omissions {
		names = append(names, omission.Name)
	}
	omission := omissionVerdictSchema(text(), evidence)
	if len(names) > 0 {
		omission = omissionVerdictSchema(s.Enum(names...), evidence)
	}
	fields := []s.Field{s.F("omissionVerdicts", s.Array(omission).Length(len(names)))}
	for _, kind := range []purchaseKind{thirdPurchaseKind, fifthPurchaseKind} {
		var builds []string
		for _, purchase := range v.purchases(kind) {
			builds = append(builds, purchase.build)
		}
		purchase := purchaseVerdictSchema(text(), evidence)
		if len(builds) > 0 {
			purchase = purchaseVerdictSchema(s.Enum(builds...), evidence)
		}
		fields = append(fields, s.F(kind.field, s.Array(purchase).Length(len(builds))))
	}
	return fields
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
	for _, kind := range []purchaseKind{thirdPurchaseKind, fifthPurchaseKind} {
		subjects := v.purchases(kind)
		counts = make([]int, len(subjects))
		for _, verdict := range purchaseVerdicts(review, kind) {
			known := false
			for i, purchase := range subjects {
				if strings.TrimSpace(verdict.Build) == purchase.build {
					counts[i]++
					known = true
				}
			}
			if !known {
				problems = append(problems, fmt.Sprintf("a %s verdict on %q, which is no path's %s", kind.noun, verdict.Build, kind.noun))
			}
		}
		for i, purchase := range subjects {
			switch {
			case counts[i] == 0:
				problems = append(problems, "no verdict on the "+kind.noun+" "+purchase.build)
			case counts[i] > 1:
				problems = append(problems, fmt.Sprintf("%d verdicts on the %s %s", counts[i], kind.noun, purchase.build))
			}
		}
	}
	return problems
}

// incompleteVerdicts is the error for a review whose verdicts are not
// exactly one per subject; like other malformed review output it fails the
// review and keeps the checked draft.
func incompleteVerdicts(problems []string) *ModelError {
	message := "The model review did not return exactly one verdict per whole-technique omission and per third and fifth purchase: it gave " + strings.Join(problems, "; ") + ". The draft is retained. Retry the review or choose another model."
	return &ModelError{Message: message, Failure: &Failure{Code: CodeOutputInvalid, Message: message, Stage: "review"}}
}

// findings records every verdict as a model Finding, in subject order: an
// omission's verdict under OmissionVerdictRule, a third purchase's under
// PathIdentityVerdictRule and a fifth purchase's under CapstoneVerdictRule.
// A pass stays a pass; a fail or unresolved verdict is an open Finding like
// any other.
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
	for _, kind := range []purchaseKind{thirdPurchaseKind, fifthPurchaseKind} {
		for _, purchase := range v.purchases(kind) {
			for _, verdict := range purchaseVerdicts(review, kind) {
				if strings.TrimSpace(verdict.Build) != purchase.build {
					continue
				}
				subject := m.PathKeys[purchase.path] + ", " + purchase.build
				if purchase.name != "" {
					subject += " " + purchase.name
				}
				out = append(out, verdictFinding(kind.prefix+purchase.key, kind.category, verdict.Outcome, subject, kind.rule, verdict.Reason, verdict.Action, verdict.Evidence))
				break
			}
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
// omission or a third or fifth purchase.
func IsReviewVerdict(finding Finding) bool {
	return finding.Method == "model" && (finding.Rule == OmissionVerdictRule || finding.Rule == PathIdentityVerdictRule || finding.Rule == CapstoneVerdictRule)
}

// checkVerdicts rejects a review whose verdicts cite an unknown document or
// are not exactly one per subject.
func checkVerdicts(checked Checked, review SemanticReview, documents map[string]bool) error {
	var evidence [][]string
	for _, verdict := range review.OmissionVerdicts {
		evidence = append(evidence, verdict.Evidence)
	}
	for _, verdict := range append(append([]PurchaseVerdict{}, review.ThirdPurchaseVerdicts...), review.FifthPurchaseVerdicts...) {
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

// Verdict repeats (SOL-61-10). The v34 Luffy review repeated four omission
// fails and the x-x-3 fail as free-form findings despite the instruction not
// to, spending five of its eight findings on verdicts already recorded
// (OPUS-NET-61-13). Code drops a free-form model finding only when it
// repeats a verdict: it names exactly that verdict's subject, as the
// omitted technique or the purchase's build code with at most its path and
// name, it has the verdict's outcome, and its message shares at least half
// of the content words of the shorter of it and the verdict's reason
// (repeatedReasonShare). In those five v34 pairs the share is 0.56 to 0.95.
// A finding on the same subject with another outcome, or with its own
// reason, such as a capstone's price beside its capstone verdict, stays.

// repeatedReasonShare is the least share of shared content words at which a
// finding on a verdict's subject with its outcome repeats its reason.
const repeatedReasonShare = 0.5

// withoutRepeats drops each finding that repeats one of the review's
// verdicts.
func (v verdictSubjects) withoutRepeats(review SemanticReview, findings []Finding) []Finding {
	kept := []Finding{}
	for _, finding := range findings {
		if !v.repeatsVerdict(review, finding) {
			kept = append(kept, finding)
		}
	}
	return kept
}

// repeatsVerdict reports a free-form finding that repeats one verdict's
// subject, outcome and reason.
func (v verdictSubjects) repeatsVerdict(review SemanticReview, finding Finding) bool {
	if finding.Method != "model" || IsReviewVerdict(finding) {
		return false
	}
	repeats := func(outcome, reason string) bool {
		return finding.Outcome == outcome && sharedReason(finding.Message, reason) >= repeatedReasonShare
	}
	for _, omission := range v.omissions {
		for _, verdict := range review.OmissionVerdicts {
			if sameTechnique(verdict.Technique, omission.Name) && namesOmission(finding.Subject, omission.Name) && repeats(verdict.Outcome, verdict.Reason) {
				return true
			}
		}
	}
	for _, kind := range []purchaseKind{thirdPurchaseKind, fifthPurchaseKind} {
		for _, purchase := range v.purchases(kind) {
			for _, verdict := range purchaseVerdicts(review, kind) {
				if strings.TrimSpace(verdict.Build) == purchase.build && namesPurchase(finding.Subject, purchase) && repeats(verdict.Outcome, verdict.Reason) {
					return true
				}
			}
		}
	}
	return false
}

// omissionLabel is a leading word a finding's subject may give an omission,
// as "Omission: Gear 5" or "omittedTechniques, Gear 5 (major)".
var (
	omissionLabel = regexp.MustCompile(`(?i)^\s*(?:omittedTechniques|omitted techniques?|omission)\s*[:,]?\s*`)
	trailingRank  = regexp.MustCompile(`(?i)\s*\((?:core|major|minor)\)\s*$`)
)

// namesOmission reports a subject that names exactly one omitted
// technique, with at most a leading omission label and a trailing rank, and
// no build code.
func namesOmission(subject, technique string) bool {
	if buildCodePattern.MatchString(subject) {
		return false
	}
	name := trailingRank.ReplaceAllString(omissionLabel.ReplaceAllString(subject, ""), "")
	return sameTechnique(name, technique)
}

// namesPurchase reports a subject that names exactly one purchase: its build
// code and no other, with at most its path key and its name besides, as
// "x-x-3, Gear 4 Fast Punch" or "path3, x-x-3 Gear 4 Fast Punch".
func namesPurchase(subject string, purchase purchaseSubject) bool {
	codes := buildCodePattern.FindAllString(subject, -1)
	if len(codes) != 1 || codes[0] != purchase.build {
		return false
	}
	rest := strings.Replace(subject, purchase.build, " ", 1)
	for _, part := range []string{purchase.name, purchase.key} {
		if part == "" {
			continue
		}
		if at := strings.Index(strings.ToLower(rest), strings.ToLower(part)); at >= 0 {
			rest = rest[:at] + " " + rest[at+len(part):]
		}
	}
	return strings.Trim(rest, " ,:;.-") == ""
}

// reasonStopWords are frequent words that say nothing of a reason.
var reasonStopWords = map[string]bool{
	"the": true, "and": true, "that": true, "this": true, "with": true, "from": true, "into": true, "have": true, "does": true,
	"their": true, "there": true, "which": true, "when": true, "than": true, "then": true, "also": true, "only": true, "each": true,
	"other": true, "more": true, "most": true, "will": true, "would": true, "been": true, "were": true, "what": true, "them": true,
	"they": true, "not": true, "but": true, "for": true, "are": true, "was": true, "can": true, "one": true, "two": true, "any": true,
}

var reasonWord = regexp.MustCompile(`[\pL\pN]+(?:-[\pL\pN]+)*`)

// contentWords are the distinct words of a text with four or more
// characters, a build code included, other than frequent words.
func contentWords(text string) map[string]bool {
	words := map[string]bool{}
	for _, word := range reasonWord.FindAllString(strings.ToLower(text), -1) {
		if len([]rune(word)) >= 4 && !reasonStopWords[word] {
			words[word] = true
		}
	}
	return words
}

// sharedReason is the share of the content words of the shorter of two
// texts that the other has too.
func sharedReason(a, b string) float64 {
	x, y := contentWords(a), contentWords(b)
	if len(x) == 0 || len(y) == 0 {
		return 0
	}
	if len(y) < len(x) {
		x, y = y, x
	}
	shared := 0
	for word := range x {
		if y[word] {
			shared++
		}
	}
	return float64(shared) / float64(len(x))
}
