package unit

import (
	"fmt"
	"regexp"
	"slices"
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
// kind (SOL-61-10). A fifth purchase verdict also judges payoff and price
// from the numbers beside its subject (capstone_payoff.go, SOL-61-13). A
// proposal verdict judges each purchase-level proposed mechanic: the v36
// Luffy review caught 3-x-x's "Massive-scale impact", which restates its
// typed splash, but not x-x-4's "Continuous momentum", which has no
// player-visible effect (OPUS-NET-61-16, SOL-61-13). Code checks that the
// review returns exactly one verdict per subject and records each one in
// the Result as a model Finding: a pass as a pass, so the owner can read
// every verdict, and a fail or unresolved verdict as an open Finding.

const (
	// OmissionVerdictRule is the rule of the Finding that records the
	// review's verdict on one whole-technique omission.
	OmissionVerdictRule = "omission-verdict"
	// RequiredConceptVerdictRule is the rule of the Finding that records
	// the review's verdict on one Required concept of the Request.
	RequiredConceptVerdictRule = "required-concept-verdict"
	// PathIdentityVerdictRule is the rule of the Finding that records the
	// review's verdict on one path's third purchase, its Path identity.
	PathIdentityVerdictRule = "path-identity-verdict"
	// CapstoneVerdictRule is the rule of the Finding that records the
	// review's verdict on one path's fifth purchase, judged with the other
	// paths' fifth purchases and its path's identity.
	CapstoneVerdictRule = "capstone-verdict"
	// ProposalVerdictRule is the rule of the Finding that records the
	// review's verdict on one purchase-level proposed mechanic.
	ProposalVerdictRule = "proposal-verdict"
)

// OmissionVerdict is the review's verdict on one whole-technique omission.
type OmissionVerdict struct {
	Technique string   `json:"technique"`
	Outcome   string   `json:"outcome"`
	Reason    string   `json:"reason"`
	Action    *string  `json:"action"`
	Evidence  []string `json:"evidence"`
}

// RequiredConceptVerdict is the review's verdict on one required concept
// of the Request, by its name: pass when a typed change carries its central
// effect, fail when the Unit adapts it only in name or only for a
// peripheral effect, and unresolved when only a coherent, source-fitting
// proposed mechanic carries it, a design gap (SOL-76-01). The plan check
// establishes only that a purchase names it and promises or proposes
// something for it, so without this verdict a review could leave a
// required concept unjudged.
type RequiredConceptVerdict struct {
	Concept  string   `json:"concept"`
	Outcome  string   `json:"outcome"`
	Reason   string   `json:"reason"`
	Action   *string  `json:"action"`
	Evidence []string `json:"evidence"`
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

// ProposalVerdict is the review's verdict on one proposed mechanic of a
// purchase, by the purchase's build code and the proposal's name. A
// proposal grants nothing, so it never passes: a coherent one is
// unresolved, a Design gap, and any other fails.
type ProposalVerdict struct {
	Build    string   `json:"build"`
	Proposal string   `json:"proposal"`
	Outcome  string   `json:"outcome"`
	Reason   string   `json:"reason"`
	Action   *string  `json:"action"`
	Evidence []string `json:"evidence"`
}

var (
	verdictOutcome  = s.Enum("pass", "fail", "unresolved")
	proposalOutcome = s.Enum("unresolved", "fail")
)

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

func requiredConceptVerdictSchema(concept, evidence s.Schema) *s.ObjectSchema {
	return s.StrictObject(
		s.F("concept", concept),
		s.F("outcome", verdictOutcome),
		s.F("reason", text()),
		s.F("action", s.Nullable(text())),
		s.F("evidence", s.Array(evidence)),
	)
}

func proposalVerdictSchema(build, proposal, evidence s.Schema) *s.ObjectSchema {
	return s.StrictObject(
		s.F("build", build),
		s.F("proposal", proposal),
		s.F("outcome", proposalOutcome),
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
// names of the plan's whole-technique omissions with their rank, the
// Request's required concepts (required), each path's third and fifth
// purchase by build code, with its name and technique, and each proposed
// mechanic of a purchase.
type verdictSubjects struct {
	omissions      []PlanOmission
	required       []RequiredConcept
	thirdPurchases []purchaseSubject
	fifthPurchases []purchaseSubject
	proposals      []proposalSubject
}

// proposalSubject is one proposed mechanic of a purchase: the purchase, its
// tier and the proposal, and its index among the purchase's proposals.
type proposalSubject struct {
	purchaseSubject
	tier, index int
	proposal    m.ProposedMechanic
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
	subjects.required = checked.Draft.Prepared.Request.RequiredConcepts
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
	for index := range m.PathKeys {
		for tier := 1; tier <= len(m.TierKeys); tier++ {
			var seen []string
			for i, proposed := range candidate.Blueprint.Paths.At(index).Tiers.At(tier).ProposedMechanics {
				// Two proposals of one purchase with the same name are
				// one subject.
				if slices.ContainsFunc(seen, func(name string) bool { return sameTechnique(name, proposed.Name) }) || strings.TrimSpace(proposed.Name) == "" {
					continue
				}
				seen = append(seen, proposed.Name)
				subjects.proposals = append(subjects.proposals, proposalSubject{purchaseSubject: purchase(index, tier), tier: tier, index: i, proposal: proposed})
			}
		}
	}
	return subjects
}

// judgedBy reports a proposal verdict on this proposal subject.
func (p proposalSubject) judgedBy(verdict ProposalVerdict) bool {
	return strings.TrimSpace(verdict.Build) == p.build && sameTechnique(verdict.Proposal, p.proposal.Name)
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
// the review context. Each fifth purchase carries its payoff numbers
// (CapstonePayoffs) when the review has purchase evidence, and each
// proposal what it says it does.
func (v verdictSubjects) context(payoffs map[string]*s.Object) *s.Object {
	omissions := []any{}
	for _, omission := range v.omissions {
		entry := s.NewObject().Set("technique", omission.Name)
		if omission.Importance != "" {
			entry.Set("importance", omission.Importance)
		}
		omissions = append(omissions, entry)
	}
	required := []any{}
	for _, concept := range v.required {
		required = append(required, s.NewObject().Set("concept", concept.Name))
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
	fifth := purchases(v.fifthPurchases)
	for i, purchase := range v.fifthPurchases {
		if payoff, ok := payoffs[purchase.key]; ok {
			fifth[i].(*s.Object).Set("payoff", payoff)
		}
	}
	proposals := []any{}
	for _, p := range v.proposals {
		entry := s.NewObject().Set("build", p.build).Set("path", p.key).Set("name", p.name)
		if p.technique != "" {
			entry.Set("technique", p.technique)
		}
		proposals = append(proposals, entry.Set("proposal", p.proposal.Name).Set("effect", p.proposal.Effect).Set("sourceIds", anyStrings(p.proposal.SourceIDs)))
	}
	out := s.NewObject().Set("omissions", omissions)
	if len(required) > 0 {
		out.Set("requiredConcepts", required)
	}
	return out.Set("thirdPurchases", purchases(v.thirdPurchases)).Set("fifthPurchases", fifth).Set("proposals", proposals)
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
	// A Request without required concepts asks for no verdict on them, so
	// its review schema is unchanged.
	if len(v.required) > 0 {
		var concepts []string
		for _, concept := range v.required {
			concepts = append(concepts, concept.Name)
		}
		fields = append(fields, s.F("requiredConceptVerdicts", s.Array(requiredConceptVerdictSchema(s.Enum(concepts...), evidence)).Length(len(concepts))))
	}
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
	var builds, proposals []string
	for _, p := range v.proposals {
		if !slices.Contains(builds, p.build) {
			builds = append(builds, p.build)
		}
		if !slices.Contains(proposals, p.proposal.Name) {
			proposals = append(proposals, p.proposal.Name)
		}
	}
	proposal := proposalVerdictSchema(text(), text(), evidence)
	if len(builds) > 0 {
		proposal = proposalVerdictSchema(s.Enum(builds...), s.Enum(proposals...), evidence)
	}
	return append(fields, s.F("proposalVerdicts", s.Array(proposal).Length(len(v.proposals))))
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
	counts = make([]int, len(v.required))
	for _, verdict := range review.RequiredConceptVerdicts {
		known := false
		for i, concept := range v.required {
			if sameTechnique(verdict.Concept, concept.Name) {
				counts[i]++
				known = true
			}
		}
		if !known {
			problems = append(problems, fmt.Sprintf("a verdict on %q, which requiredConcepts does not list", verdict.Concept))
		}
	}
	for i, concept := range v.required {
		switch {
		case counts[i] == 0:
			problems = append(problems, fmt.Sprintf("no verdict on the required concept %q", concept.Name))
		case counts[i] > 1:
			problems = append(problems, fmt.Sprintf("%d verdicts on the required concept %q", counts[i], concept.Name))
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
	counts = make([]int, len(v.proposals))
	for _, verdict := range review.ProposalVerdicts {
		known := false
		for i, p := range v.proposals {
			if p.judgedBy(verdict) {
				counts[i]++
				known = true
			}
		}
		if !known {
			problems = append(problems, fmt.Sprintf("a proposal verdict on %q at %q, which that purchase does not propose", verdict.Proposal, verdict.Build))
		}
	}
	for i, p := range v.proposals {
		switch {
		case counts[i] == 0:
			problems = append(problems, fmt.Sprintf("no verdict on the proposed mechanic %q of %s", p.proposal.Name, p.build))
		case counts[i] > 1:
			problems = append(problems, fmt.Sprintf("%d verdicts on the proposed mechanic %q of %s", counts[i], p.proposal.Name, p.build))
		}
	}
	return problems
}

// incompleteVerdicts is the error for a review whose verdicts are not
// exactly one per subject; like other malformed review output it fails the
// review and keeps the checked draft.
func incompleteVerdicts(problems []string) *ModelError {
	message := "The model review did not return exactly one verdict per whole-technique omission, per required concept of the Request, per third and fifth purchase and per proposed mechanic: it gave " + strings.Join(problems, "; ") + ". The draft is retained. Retry the review or choose another model."
	return &ModelError{Message: message, Failure: &Failure{Code: CodeOutputInvalid, Message: message, Stage: "review"}}
}

// findings records every verdict as a model Finding, in subject order: an
// omission's verdict under OmissionVerdictRule, a required concept's under
// RequiredConceptVerdictRule, a third purchase's under
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
	for index, concept := range v.required {
		for _, verdict := range review.RequiredConceptVerdicts {
			if sameTechnique(verdict.Concept, concept.Name) {
				out = append(out, verdictFinding(fmt.Sprintf("verdict.required-concept.%d", index+1), "coverage", verdict.Outcome, "requiredConcepts, "+concept.Name, RequiredConceptVerdictRule, verdict.Reason, verdict.Action, verdict.Evidence))
				break
			}
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
	for _, p := range v.proposals {
		for _, verdict := range review.ProposalVerdicts {
			if !p.judgedBy(verdict) {
				continue
			}
			id := fmt.Sprintf("verdict.proposal.%s.%s.%d", p.key, m.TierKeys[p.tier-1], p.index+1)
			out = append(out, verdictFinding(id, "unsupported", verdict.Outcome, p.label(), ProposalVerdictRule, verdict.Reason, verdict.Action, verdict.Evidence))
			break
		}
	}
	return out
}

// label is a proposal verdict's subject, as "path3, x-x-4 Python Continuous
// Stretch: Continuous momentum".
func (p proposalSubject) label() string {
	subject := p.key + ", " + p.build
	if p.name != "" {
		subject += " " + p.name
	}
	return subject + ": " + p.proposal.Name
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
// omission, a required concept, a third or fifth purchase or a proposed
// mechanic.
func IsReviewVerdict(finding Finding) bool {
	return finding.Method == "model" && (finding.Rule == OmissionVerdictRule || finding.Rule == RequiredConceptVerdictRule || finding.Rule == PathIdentityVerdictRule || finding.Rule == CapstoneVerdictRule || finding.Rule == ProposalVerdictRule)
}

// checkVerdicts rejects a review whose verdicts cite an unknown document or
// are not exactly one per subject.
func checkVerdicts(checked Checked, review SemanticReview, documents map[string]bool) error {
	var evidence [][]string
	for _, verdict := range review.OmissionVerdicts {
		evidence = append(evidence, verdict.Evidence)
	}
	for _, verdict := range review.RequiredConceptVerdicts {
		evidence = append(evidence, verdict.Evidence)
	}
	for _, verdict := range append(append([]PurchaseVerdict{}, review.ThirdPurchaseVerdicts...), review.FifthPurchaseVerdicts...) {
		evidence = append(evidence, verdict.Evidence)
	}
	for _, verdict := range review.ProposalVerdicts {
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
// (OPUS-NET-61-13). Word overlap cannot establish that a finding repeats a
// verdict: "does not redirect" and "redirects" share every content word
// (SOL-74-01). Code therefore drops a free-form model finding only when it
// names exactly a verdict's subject, as the omitted technique, the required
// concept or the purchase's build code with at most its path and name, has
// the verdict's outcome, and its message is the verdict's reason word for
// word once case, spacing and final punctuation are normalized
// (sameReason). Every other finding stays, a paraphrase included; the
// prompt asks the review to spend its findings on issues the verdicts do
// not cover.

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
		return finding.Outcome == outcome && sameReason(finding.Message, reason)
	}
	for _, omission := range v.omissions {
		for _, verdict := range review.OmissionVerdicts {
			if sameTechnique(verdict.Technique, omission.Name) && namesOmission(finding.Subject, omission.Name) && repeats(verdict.Outcome, verdict.Reason) {
				return true
			}
		}
	}
	for _, concept := range v.required {
		for _, verdict := range review.RequiredConceptVerdicts {
			if sameTechnique(verdict.Concept, concept.Name) && namesRequiredConcept(finding.Subject, concept.Name) && repeats(verdict.Outcome, verdict.Reason) {
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
	for _, p := range v.proposals {
		for _, verdict := range review.ProposalVerdicts {
			if p.judgedBy(verdict) && namesProposal(finding.Subject, p) && repeats(verdict.Outcome, verdict.Reason) {
				return true
			}
		}
	}
	return false
}

// namesProposal reports a subject that names exactly one proposed
// mechanic: its purchase as namesPurchase reads it, and the proposal's
// name.
func namesProposal(subject string, p proposalSubject) bool {
	at := strings.Index(strings.ToLower(subject), strings.ToLower(p.proposal.Name))
	if at < 0 {
		return false
	}
	return namesPurchase(subject[:at]+" "+subject[at+len(p.proposal.Name):], p.purchaseSubject)
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

// requiredLabel is a leading word a finding's subject may give a required
// concept, as "requiredConcepts, Gear 4" or "Required concept: Gear 4".
var requiredLabel = regexp.MustCompile(`(?i)^\s*(?:requiredConcepts|required concepts?)\s*[:,]?\s*`)

// namesRequiredConcept reports a subject that names exactly one required
// concept, with at most a leading required concept label, and no build
// code.
func namesRequiredConcept(subject, concept string) bool {
	return !buildCodePattern.MatchString(subject) && sameTechnique(requiredLabel.ReplaceAllString(subject, ""), concept)
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

// sameReason reports two reasons that are the same text once case, runs of
// spaces and final punctuation are normalized.
func sameReason(a, b string) bool {
	normalize := func(text string) string {
		return strings.TrimRight(strings.Join(strings.Fields(strings.ToLower(text)), " "), ".!;: ")
	}
	return normalize(a) != "" && normalize(a) == normalize(b)
}
