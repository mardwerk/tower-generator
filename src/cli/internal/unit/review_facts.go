package unit

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// The model review reads the unit through facts that code resolved: every
// legal build with its price and resolved attack. A review that cites a build
// the Definition does not allow, or a resolved value that is wrong, is
// corrected once and then rejected, so it cannot report findings about builds
// such as 3-3-0. The correction must return the findings whose citations
// held unchanged in every field, so its summary describes them. A flagged
// finding whose only fault is build notation, such as 5-x-2 for x-x-5, must
// return under its ID with only the flagged codes replaced; one that names
// an impossible build or a wrong fact may be corrected or withdrawn
// (keepCheckedFindings). No impossible build is ever published.

// buildCodePattern matches what a review writes as a build code: a
// concrete build such as 3-2-0, a purchase such as x-4-x, or a malformed mix
// such as 1-x-5.
var buildCodePattern = regexp.MustCompile(`\b[0-9x]-[0-9x]-[0-9x]\b`)

// purchaseCode reports a purchase written in build-code notation: one path
// at a tier from 1 to 5, the others x, as in 3-x-x or x-x-5.
func purchaseCode(code string) bool {
	digits := 0
	for _, part := range strings.Split(code, "-") {
		switch {
		case part == "x":
		case part >= "1" && part <= string(rune('0'+len(m.TierKeys))):
			digits++
		default:
			return false
		}
	}
	return digits == 1
}

// attackFacts is one resolved attack as the review reads it.
func attackFacts(attack m.Attack) *s.Object {
	st := attack.Stats
	facts := s.NewObject().
		Set("delivery", attack.Delivery).
		Set("damageType", attack.DamageType).
		Set("targeting", attack.Targeting).
		Set("damage", st.Damage).
		Set("intervalSeconds", st.IntervalSeconds).
		Set("range", st.Range).
		Set("projectiles", st.Projectiles).
		Set("pierce", st.Pierce).
		Set("maxEnemiesHitPerAttack", st.Projectiles*st.Pierce)
	if st.SplashRadius > 0 {
		facts.Set("splashRadius", st.SplashRadius)
	}
	if st.Projectiles > 1 && attack.Distribution != "" {
		facts.Set("distribution", attack.Distribution)
	}
	if f := attack.FollowUp; f != nil {
		facts.Set("followUp", s.NewObject().Set("count", f.Count).Set("damage", f.DamageMultiplier*st.Damage).Set("radius", f.Radius))
	}
	if statuses := attack.AppliedStatuses(); len(statuses) > 0 {
		facts.Set("statuses", s.FromGoValue(statuses))
	}
	// Bonus damage per hit by enemy property, so a finding can cite
	// attack.bonusDamage.hardened or attack.bonusDamage.blimp.
	if len(attack.BonusDamage) > 0 {
		bonuses := s.NewObject()
		for _, bonus := range attack.BonusDamage {
			bonuses.Set(bonus.Property, bonus.Damage)
		}
		facts.Set("bonusDamage", bonuses)
	}
	var detects []string
	if attack.Detects != nil {
		detects = *attack.Detects
	} else if attack.Camo {
		detects = []string{"camo"}
	}
	if len(detects) > 0 {
		facts.Set("detects", s.FromGoValue(detects))
	}
	return facts
}

// LegalBuildFacts lists every build the Definition allows with its total
// price, resolved attack and Active Ability.
func LegalBuildFacts(blueprint *m.Blueprint, definition m.Definition) []any {
	var out []any
	for _, selection := range m.AllLegalBuilds(definition) {
		build := m.ResolveUnchecked(blueprint, selection)
		entry := s.NewObject().
			Set("code", fmt.Sprintf("%d-%d-%d", selection[0], selection[1], selection[2])).
			Set("cost", build.CumulativeCost).
			Set("attack", attackFacts(build.BaseAttack))
		for _, ability := range build.Abilities {
			boosted := ability.BoostedAttack.Stats
			entry.Set("activeAbility", s.NewObject().
				Set("name", ability.Name).
				Set("durationSeconds", ability.DurationSeconds).
				Set("cooldownSeconds", ability.CooldownSeconds).
				Set("boostedDamage", boosted.Damage).
				Set("boostedIntervalSeconds", boosted.IntervalSeconds).
				Set("boostedRange", boosted.Range))
		}
		out = append(out, entry)
	}
	return out
}

// reviewCitations checks what a review cites against the Definition's legal
// builds: the build codes a finding or the summary names, and the resolved
// facts a finding cites.
type reviewCitations struct {
	builds map[string]*s.Object
	// claims checks what a finding says about technique timing and the
	// side-purchase comparisons against a capstone (review_claims.go).
	claims   reviewClaims
	currency string
}

func newReviewCitations(blueprint *m.Blueprint, definition m.Definition) reviewCitations {
	c := reviewCitations{builds: map[string]*s.Object{}}
	for _, entry := range LegalBuildFacts(blueprint, definition) {
		object := entry.(*s.Object)
		code, _ := object.Get("code")
		c.builds[code.(string)] = object
	}
	return c
}

// citationProblem is what code rejected in one finding, or in the summary
// when finding is empty: impossible builds, malformed codes, wrong facts,
// or purchases the text names in place of the one its subject names.
type citationProblem struct {
	finding    string
	illegal    []string
	malformed  []string
	wrong      []string
	mismatched []purchaseMismatch
}

// purchaseMismatch is a purchase a finding's subject names and the purchase
// on another path at the same tier that its text names instead.
type purchaseMismatch struct {
	finding, subject, text string
}

func (m purchaseMismatch) String() string {
	return fmt.Sprintf("%s's subject names %s, but its text names %s instead.", m.finding, m.subject, m.text)
}

// mismatchSentences are the mismatches as the prompt and errors state them.
func mismatchSentences(mismatches []purchaseMismatch) []string {
	var out []string
	for _, m := range mismatches {
		out = append(out, m.String())
	}
	return out
}

// problems lists the review's rejected citations, the summary first, then
// the findings in order.
func (c reviewCitations) problems(review SemanticReview) []citationProblem {
	var out []citationProblem
	if illegal, malformed := c.illegal(review.Summary); len(illegal)+len(malformed) > 0 {
		out = append(out, citationProblem{illegal: illegal, malformed: malformed})
	}
	for _, f := range review.Findings {
		texts := []string{f.Subject, f.Rule, f.Message}
		if f.Action != nil {
			texts = append(texts, *f.Action)
		}
		illegal, malformed := c.illegal(texts...)
		problem := citationProblem{finding: f.ID, illegal: illegal, malformed: malformed, wrong: c.wrongFacts(f), mismatched: mismatchedPurchases(f)}
		if len(problem.illegal) > 0 || len(problem.malformed) > 0 || len(problem.wrong) > 0 || len(problem.mismatched) > 0 {
			out = append(out, problem)
		}
	}
	return out
}

// purchaseCodes lists the purchases a text names in build-code notation.
func purchaseCodes(text string) []string {
	var out []string
	for _, code := range buildCodePattern.FindAllString(text, -1) {
		if purchaseCode(code) && !slices.Contains(out, code) {
			out = append(out, code)
		}
	}
	return out
}

// purchaseTier is the tier a purchase code names: 4 for x-4-x.
func purchaseTier(code string) string {
	return strings.Trim(strings.ReplaceAll(code, "-", ""), "x")
}

// mismatchedPurchases reports a finding whose message and action never name
// the purchase its subject names but name the same tier on another path
// instead, such as a subject x-5-x with a message about 5-x-x (a live Luffy
// review on #27). A text that names the subject's purchase may compare it
// with any other.
func mismatchedPurchases(f Finding) []purchaseMismatch {
	subject := purchaseCodes(f.Subject)
	text := f.Message
	if f.Action != nil {
		text += "\n" + *f.Action
	}
	named := purchaseCodes(text)
	var out []purchaseMismatch
	for _, code := range subject {
		if slices.Contains(named, code) {
			continue
		}
		for _, other := range named {
			if purchaseTier(other) == purchaseTier(code) && !slices.Contains(subject, other) {
				out = append(out, purchaseMismatch{finding: f.ID, subject: code, text: other})
				break
			}
		}
	}
	return out
}

// illegal lists the build codes the texts cite that the Definition does not
// allow, in order of first mention: concrete builds legalBuilds does not
// list, and codes that are neither a build nor a purchase, such as 1-x-5
// for 1-5-0 (reported on #27).
func (c reviewCitations) illegal(texts ...string) (impossible, malformed []string) {
	seen := map[string]bool{}
	for _, code := range buildCodePattern.FindAllString(strings.Join(texts, "\n"), -1) {
		if _, legal := c.builds[code]; !legal && !purchaseCode(code) && !seen[code] {
			seen[code] = true
			if strings.Contains(code, "x") {
				malformed = append(malformed, code)
			} else {
				impossible = append(impossible, code)
			}
		}
	}
	return impossible, malformed
}

// flagClass is what a correction may do with a flagged finding.
type flagClass int

const (
	flagNotation   flagClass = iota // (a) same ID, notation fixed, or reject
	flagImpossible                  // (b) fix or withdraw; never published
	flagFact                        // (c) correct or withdraw
)

// class determines the permitted correction for a citation problem.
func (p citationProblem) class() flagClass {
	switch {
	case len(p.wrong) > 0:
		return flagFact
	case len(p.illegal) > 0:
		return flagImpossible
	}
	return flagNotation
}

// notationOnly reports a review whose only rejected citations are notation:
// malformed codes or another path's purchase, with no wrong fact and no
// impossible build in any finding or the summary. Its correction may change
// build codes and nothing else.
func notationOnly(problems []citationProblem) bool {
	for _, p := range problems {
		if p.class() != flagNotation {
			return false
		}
	}
	return true
}

// splitCodes splits a text at its build codes: the words around them, one
// more than there are codes, and the codes in order.
func splitCodes(text string) (words, codes []string) {
	last := 0
	for _, at := range buildCodePattern.FindAllStringIndex(text, -1) {
		words = append(words, text[last:at[0]])
		codes = append(codes, text[at[0]:at[1]])
		last = at[1]
	}
	return append(words, text[last:]), codes
}

// onlyFlaggedCodesChanged reports whether after is before with a build code
// replaced only where before wrote one of the flagged codes. Every word and
// every other code stays exact, so a correction cannot change a comparison
// code that was valid while it fixes a malformed one (SOL-34-05). Whether a
// replacement is itself valid is checked again after the correction.
func onlyFlaggedCodesChanged(before, after string, flagged []string) bool {
	beforeWords, beforeCodes := splitCodes(before)
	afterWords, afterCodes := splitCodes(after)
	if !slices.Equal(beforeWords, afterWords) || len(beforeCodes) != len(afterCodes) {
		return false
	}
	for i, code := range beforeCodes {
		if afterCodes[i] != code && !slices.Contains(flagged, code) {
			return false
		}
	}
	return true
}

// notationFixed reports whether a class (a) finding came back with only its
// flagged build codes replaced: its malformed codes anywhere in its text,
// and for a mismatch the subject's purchase in the subject and the other
// path's purchase in the message and action. Every other word and code of
// subject, rule, message and action, and every other field, must stay as it
// was (SOL-34-04: a correction kept the ID, outcome and facts but replaced
// the claim with "No payoff concern").
func notationFixed(before, after Finding, p citationProblem) bool {
	var subject, text []string
	for _, m := range p.mismatched {
		subject, text = append(subject, m.subject), append(text, m.text)
	}
	action := func(f Finding) string {
		if f.Action == nil {
			return ""
		}
		return *f.Action
	}
	if (before.Action == nil) != (after.Action == nil) ||
		!onlyFlaggedCodesChanged(before.Subject, after.Subject, slices.Concat(p.malformed, subject)) ||
		!onlyFlaggedCodesChanged(before.Rule, after.Rule, p.malformed) ||
		!onlyFlaggedCodesChanged(before.Message, after.Message, slices.Concat(p.malformed, text)) ||
		!onlyFlaggedCodesChanged(action(before), action(after), slices.Concat(p.malformed, text)) {
		return false
	}
	before.Subject, before.Rule, before.Message, before.Action = "", "", "", nil
	after.Subject, after.Rule, after.Message, after.Action = "", "", "", nil
	return s.Stringify(s.FromGoValue(before)) == s.Stringify(s.FromGoValue(after))
}

// impossibleBuilds lists the numeric builds this Definition does not allow
// anywhere in what a Result would publish; evidence holds document IDs.
func (c reviewCitations) impossibleBuilds(summary string, findings []Finding) []string {
	texts := []string{summary}
	for _, f := range findings {
		texts = append(texts, f.ID, f.Subject, f.Rule, f.Message)
		if f.Action != nil {
			texts = append(texts, *f.Action)
		}
		for _, fact := range f.Facts {
			texts = append(texts, fact.Build, fact.Field, fact.Value)
		}
	}
	impossible, _ := c.illegal(texts...)
	return impossible
}

// wrongFacts checks the resolved values a finding cites against legalBuilds,
// so a finding cannot rest on a misread number.
func (c reviewCitations) wrongFacts(finding Finding) []string {
	var out []string
	for _, fact := range finding.Facts {
		build, field := strings.TrimSpace(fact.Build), strings.TrimSpace(fact.Field)
		entry, ok := c.builds[build]
		if !ok {
			out = append(out, fmt.Sprintf("%s cites build %s, which legalBuilds does not list.", finding.ID, build))
			continue
		}
		var value any = entry
		for _, key := range strings.Split(field, ".") {
			object, ok := value.(*s.Object)
			if !ok {
				value = nil
				break
			}
			value, _ = object.Get(key)
		}
		if value == nil {
			out = append(out, fmt.Sprintf("%s cites %s %s, a field legalBuilds does not list for that build.", finding.ID, build, field))
		} else if !sameFact(value, fact.Value) {
			out = append(out, fmt.Sprintf("%s cites %s %s as %s, but it is %s.", finding.ID, build, field, strings.TrimSpace(fact.Value), s.Stringify(value)))
		}
	}
	return out
}

// keepCheckedFindings checks a correction against the review it corrects and
// returns the correction. The correction returns the whole review, so its
// summary describes the findings it returns. Every finding whose citations
// held must come back under its ID identical in every field; a correction
// that drops or changes one is rejected, because its summary was written
// about something other than the checked finding (reported on #27: a
// correction returned no findings and "No concrete issue", losing a finding
// with valid facts, and one that rewrote a kept finding's message would
// publish a summary about text the Result does not hold). What a flagged
// finding may do depends on its class (#34):
//   - (a) notation, a malformed code or another path's purchase with no
//     wrong fact: it must come back under its ID with only its flagged build
//     codes replaced (notationFixed), or the review is rejected;
//   - (b) impossible, a numeric build legalBuilds does not list: it may be
//     fixed under any ID or withdrawn, and is never published;
//   - (c) fact, a wrong or unverifiable structured fact: it may be corrected under
//     its ID or withdrawn.
//
// When every problem is notation, the correction adds no finding and its
// summary may change only at the codes the first review flagged there.
// Because the published findings are exactly the ones the correction
// returned, its summary is published with them.
func keepCheckedFindings(previous, corrected SemanticReview, problems []citationProblem) (SemanticReview, error) {
	flagged := map[string]citationProblem{}
	var summary citationProblem
	for _, p := range problems {
		if p.finding == "" {
			summary = p
		} else {
			flagged[p.finding] = p
		}
	}
	returned := map[string]Finding{}
	for _, f := range corrected.Findings {
		returned[f.ID] = f
	}
	// Only notation was wrong: the correction may change flagged build codes
	// and nothing else, so it adds no finding and keeps its summary but for
	// the codes flagged in it.
	onlyNotation := notationOnly(problems)
	var lost, notation, added []string
	summaryChanged := onlyNotation && !onlyFlaggedCodesChanged(previous.Summary, corrected.Summary, summary.malformed)
	if onlyNotation {
		before := map[string]bool{}
		for _, f := range previous.Findings {
			before[f.ID] = true
		}
		for _, f := range corrected.Findings {
			if !before[f.ID] {
				added = append(added, f.ID)
			}
		}
	}
	for _, f := range previous.Findings {
		got, ok := returned[f.ID]
		problem, isFlagged := flagged[f.ID]
		switch {
		case !isFlagged && (!ok || s.Stringify(s.FromGoValue(got)) != s.Stringify(s.FromGoValue(f))):
			lost = append(lost, f.ID)
		case isFlagged && problem.class() == flagNotation && (!ok || !notationFixed(f, got, problem)):
			notation = append(notation, f.ID)
		}
	}
	var parts []string
	if len(lost) > 0 {
		parts = append(parts, "The model review's correction dropped or changed findings whose citations held ("+strings.Join(lost, ", ")+"); they must return unchanged in every field, or its summary may not describe them.")
	}
	if len(notation) > 0 {
		parts = append(parts, "The model review's correction withdrew or changed findings whose facts held but whose build notation was wrong ("+strings.Join(notation, ", ")+"); they must return under the same ID with only the build codes code flagged replaced: every other build code, word and field must stay as it was.")
	}
	if summaryChanged || len(added) > 0 {
		part := "The model review's correction changed its summary"
		if len(added) > 0 {
			part = "The model review's correction added findings (" + strings.Join(added, ", ") + ")"
			if summaryChanged {
				part += " and changed its summary"
			}
		}
		parts = append(parts, part+" when only build notation was wrong; the summary must return with only the build codes code flagged in it replaced, and no finding may be added.")
	}
	if len(parts) > 0 {
		message := strings.Join(parts, " ") + " The draft is retained. Retry the review or choose another model."
		return SemanticReview{}, &ModelError{Message: message, Failure: &Failure{Code: CodeOutputInvalid, Message: message, Stage: "review"}}
	}
	return SemanticReview{Summary: corrected.Summary, Findings: append([]Finding{}, corrected.Findings...)}, nil
}

// rejectedCitations is the error for citations that are still wrong after
// one correction, or nil when none are.
func rejectedCitations(problems []citationProblem) error {
	if len(problems) == 0 {
		return nil
	}
	var illegal, malformed, wrong, mismatched []string
	for _, p := range problems {
		for _, code := range p.illegal {
			if !slices.Contains(illegal, code) {
				illegal = append(illegal, code)
			}
		}
		for _, code := range p.malformed {
			if !slices.Contains(malformed, code) {
				malformed = append(malformed, code)
			}
		}
		wrong = append(wrong, p.wrong...)
		mismatched = append(mismatched, mismatchSentences(p.mismatched)...)
	}
	message := "The model review cited resolved facts that are wrong after one correction: " + strings.Join(wrong, " ") + " The draft is retained. Retry the review or choose another model."
	if len(wrong) == 0 {
		message = "The model review named another path's purchase than its subject after one correction: " + strings.Join(mismatched, " ") + " The draft is retained. Retry the review or choose another model."
	}
	if len(malformed) > 0 {
		message = "The model review wrote build codes that are neither a purchase nor a build (" + strings.Join(malformed, ", ") + ") after one correction. The draft is retained. Retry the review or choose another model."
	}
	if len(illegal) > 0 {
		message = "The model review cited builds that are not legal (" + strings.Join(illegal, ", ") + ") after one correction. The draft is retained. Retry the review or choose another model."
	}
	return &ModelError{Message: message, Failure: &Failure{Code: CodeOutputInvalid, Message: message, Stage: "review"}}
}

// correctionPrompt asks for the flagged findings again, with the previous
// review and what code rejected in it.
func correctionPrompt(previous SemanticReview, problems []citationProblem) string {
	var illegal, malformed, wrong, mismatched, flagged []string
	byClass := map[flagClass][]string{}
	seen := map[string]bool{}
	for _, p := range problems {
		for _, code := range append(append([]string{}, p.illegal...), p.malformed...) {
			if !seen[code] {
				seen[code] = true
				if strings.Contains(code, "x") {
					malformed = append(malformed, code)
				} else {
					illegal = append(illegal, code)
				}
			}
		}
		wrong = append(wrong, p.wrong...)
		mismatched = append(mismatched, mismatchSentences(p.mismatched)...)
		if p.finding != "" {
			flagged = append(flagged, p.finding)
			byClass[p.class()] = append(byClass[p.class()], p.finding)
		}
	}
	var kept []string
	for _, f := range previous.Findings {
		if !slices.Contains(flagged, f.ID) {
			kept = append(kept, f.ID)
		}
	}
	text := "\n\nCorrect this review. Your previous review was: " + s.Stringify(s.FromGoValue(previous))
	if len(illegal) > 0 {
		text += "\nIt cites builds that are not legal under this Definition: " + strings.Join(illegal, ", ") + ". legalBuilds lists every legal build; judge only those, and read counts from their resolved facts."
	}
	if len(malformed) > 0 {
		var each []string
		for _, code := range malformed {
			each = append(each, code+" is neither a purchase nor a build")
		}
		text += "\nSome build codes are notation errors, not builds: " + strings.Join(each, "; ") + ". Fix the notation, not the claim: write a purchase as 3-x-x, x-4-x or x-x-5 and a build as 1-5-0, and use a build legalBuilds lists."
	}
	if len(wrong) > 0 {
		text += "\nSome cited facts are wrong or cannot be checked: " + strings.Join(wrong, " ") + " Read each value in legalBuilds, and cite only builds and fields it lists."
	}
	if len(mismatched) > 0 {
		text += "\nSome findings name another path's purchase than their subject: " + strings.Join(mismatched, " ") + " Name the purchase the finding is about with the build code its subject uses."
	}
	text += "\nReturn the whole review again, with a summary that describes exactly the findings you return, and with omissionVerdicts, thirdPurchaseVerdicts and fifthPurchaseVerdicts as in your previous review: this correction concerns findings only, and code keeps your previous verdicts."
	if len(kept) > 0 {
		text += " Return " + strings.Join(kept, ", ") + " verbatim, every field exactly as in your previous review, including message, subject, action, evidence and facts: their citations hold, and a review that drops or changes any field of one is rejected."
	}
	if len(previous.RequiredConceptVerdicts) > 0 {
		text += " Return requiredConceptVerdicts as in your previous review too."
	}
	if ids := byClass[flagNotation]; len(ids) > 0 {
		text += " Return " + strings.Join(ids, ", ") + " under the same ID with its build codes corrected and every other field unchanged, including outcome, severity, evidence and facts: code found no wrong fact in it, so a review that omits it or changes those fields is rejected. Replace only the codes named above as notation errors or as another path's purchase, each with one code, and keep every other build code and word exactly, in subject, rule, message and action."
	}
	if notationOnly(problems) {
		text += " Only build notation was wrong: return the summary as before with only the codes named above corrected, and add no finding."
	}
	if ids := byClass[flagImpossible]; len(ids) > 0 {
		text += " Replace each build legalBuilds does not list in " + strings.Join(ids, ", ") + " with a legal build and revise the finding to match that build's resolved facts, or omit the finding. No build legalBuilds does not list may appear anywhere in the review, including IDs and rules."
	}
	if ids := byClass[flagFact]; len(ids) > 0 {
		text += " Return " + strings.Join(ids, ", ") + " corrected under the same ID, or omit it when the resolved facts contradict its claim or do not list what it cites."
	}
	return text
}

// sameFact compares a cited value with the resolved one; numbers may differ
// by rounding in the last listed digit.
func sameFact(actual any, cited string) bool {
	cited = strings.TrimSpace(cited)
	switch value := actual.(type) {
	case float64:
		number, err := strconv.ParseFloat(cited, 64)
		return err == nil && math.Abs(number-value) <= math.Max(1e-9, 1e-3*math.Abs(value))
	case string:
		return strings.EqualFold(cited, value)
	case bool:
		return strings.EqualFold(cited, strconv.FormatBool(value))
	}
	return cited == s.Stringify(actual)
}
