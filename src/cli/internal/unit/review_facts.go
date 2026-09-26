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
// held unchanged in every field, so its summary describes them.

var buildCodePattern = regexp.MustCompile(`\b[0-9]-[0-9]-[0-9]\b`)

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
// when finding is empty: illegal build codes and wrong facts.
type citationProblem struct {
	finding string
	illegal []string
	wrong   []string
}

// problems lists the review's rejected citations, the summary first, then
// the findings in order.
func (c reviewCitations) problems(review SemanticReview) []citationProblem {
	var out []citationProblem
	if illegal := c.illegal(review.Summary); len(illegal) > 0 {
		out = append(out, citationProblem{illegal: illegal})
	}
	for _, f := range review.Findings {
		texts := []string{f.Subject, f.Message}
		if f.Action != nil {
			texts = append(texts, *f.Action)
		}
		problem := citationProblem{finding: f.ID, illegal: c.illegal(texts...), wrong: c.wrongFacts(f)}
		if len(problem.illegal) > 0 || len(problem.wrong) > 0 {
			out = append(out, problem)
		}
	}
	return out
}

// illegal lists the build codes the texts cite that the Definition does not
// allow, in order of first mention.
func (c reviewCitations) illegal(texts ...string) []string {
	var out []string
	seen := map[string]bool{}
	for _, code := range buildCodePattern.FindAllString(strings.Join(texts, "\n"), -1) {
		if _, legal := c.builds[code]; !legal && !seen[code] {
			seen[code] = true
			out = append(out, code)
		}
	}
	return out
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
// publish a summary about text the Result does not hold). A flagged finding
// comes back corrected under its ID or not at all. Because the published
// findings are exactly the ones the correction returned, its summary is
// published with them.
func keepCheckedFindings(previous, corrected SemanticReview, problems []citationProblem) (SemanticReview, error) {
	flagged := map[string]bool{}
	for _, p := range problems {
		if p.finding != "" {
			flagged[p.finding] = true
		}
	}
	returned := map[string]string{}
	for _, f := range corrected.Findings {
		returned[f.ID] = s.Stringify(s.FromGoValue(f))
	}
	var lost []string
	for _, f := range previous.Findings {
		if got, ok := returned[f.ID]; !flagged[f.ID] && (!ok || got != s.Stringify(s.FromGoValue(f))) {
			lost = append(lost, f.ID)
		}
	}
	if len(lost) > 0 {
		message := "The model review's correction dropped or changed findings whose citations held (" + strings.Join(lost, ", ") + "); they must return unchanged in every field, or its summary may not describe them. The draft is retained. Retry the review or choose another model."
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
	var illegal, wrong []string
	for _, p := range problems {
		for _, code := range p.illegal {
			if !slices.Contains(illegal, code) {
				illegal = append(illegal, code)
			}
		}
		wrong = append(wrong, p.wrong...)
	}
	message := "The model review cited resolved facts that are wrong after one correction: " + strings.Join(wrong, " ") + " The draft is retained. Retry the review or choose another model."
	if len(illegal) > 0 {
		message = "The model review cited builds that are not legal (" + strings.Join(illegal, ", ") + ") after one correction. The draft is retained. Retry the review or choose another model."
	}
	return &ModelError{Message: message, Failure: &Failure{Code: CodeOutputInvalid, Message: message, Stage: "review"}}
}

// correctionPrompt asks for the flagged findings again, with the previous
// review and what code rejected in it.
func correctionPrompt(previous SemanticReview, problems []citationProblem) string {
	var illegal, wrong, flagged []string
	seen := map[string]bool{}
	for _, p := range problems {
		for _, code := range p.illegal {
			if !seen[code] {
				seen[code] = true
				illegal = append(illegal, code)
			}
		}
		wrong = append(wrong, p.wrong...)
		if p.finding != "" {
			flagged = append(flagged, p.finding)
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
	if len(wrong) > 0 {
		text += "\nSome cited facts are wrong: " + strings.Join(wrong, " ") + " Read each value in legalBuilds."
	}
	text += "\nReturn the whole review again, with a summary that describes exactly the findings you return."
	if len(kept) > 0 {
		text += " Return " + strings.Join(kept, ", ") + " verbatim, every field exactly as in your previous review, including message, subject, action, evidence and facts: their citations hold, and a review that drops or changes any field of one is rejected."
	}
	if len(flagged) > 0 {
		text += " Return " + strings.Join(flagged, ", ") + " corrected under the same ID, or omit it when the resolved facts contradict its claim."
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
