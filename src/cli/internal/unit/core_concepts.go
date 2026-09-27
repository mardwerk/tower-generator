package unit

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// Core concepts (#61): under requireCoreConcepts the plan ranks every
// repertoire entry and omitted technique core, major or minor, lists every
// source technique in one of the two, and adapts each core entry on a
// purchase. Code checks only the ranking, the listing, that a purchase names
// each core entry as its technique and adapts it, and, on the checked Unit,
// whether a typed change embodies it or only a proposed mechanic does.
// Whether that purchase fulfils the concept in spirit, and whether an
// omission's reason holds, is the review's judgment: prose heuristics for
// either produced false positives (SOL-57-01).

// coreConceptsOn reports whether a Definition selects requireCoreConcepts.
func coreConceptsOn(definition m.Definition) bool {
	return definition.Profile.DesignPolicy.RequiresCoreConcepts()
}

// maxCoreConcepts bounds the core entries of a plan.
const maxCoreConcepts = 3

// conceptKey is a technique name without case, punctuation, parentheses or a
// trailing "Techniques", as words.
func conceptKey(name string) []string {
	return strings.Fields(techniqueKey(name))
}

// qualifierWord is a word that names a kind of entry, not a technique, so a
// plan entry may add it to a source technique's name and still name that
// technique: "Gear 4 forms" names "Gear 4".
var qualifierWord = regexp.MustCompile(`^(?:forms?|modes?|techniques?|transformations?|attacks?|moves?)$`)

// listsByName reports a plan entry named for a source technique: both names
// have the same words, or the entry's name is the technique's with only
// qualifier words added, as "Gear 4 forms" names "Gear 4". A name with any
// other word is a different concept, so a generic entry never lists a more
// specific technique ("Haki" lists neither "Armament Haki" nor "Observation
// Haki"), and a specific entry never lists a generic one ("Armament Haki"
// does not list "Haki"; "Juggernaut ball" does not list "Juggernaut").
func listsByName(entry, technique string) bool {
	x, y := conceptKey(entry), conceptKey(technique)
	if len(y) == 0 || len(x) < len(y) {
		return false
	}
	for start := 0; start+len(y) <= len(x); start++ {
		if !slices.Equal(y, x[start:start+len(y)]) {
			continue
		}
		extra := append(slices.Clone(x[:start]), x[start+len(y):]...)
		if !slices.ContainsFunc(extra, func(word string) bool { return !qualifierWord.MatchString(word) }) {
			return true
		}
	}
	return false
}

// generalizes reports an entry whose name is a proper part of a technique's
// name, such as "Haki" of "Armament Haki": the technique is a kind of the
// entry, not a part of it, and the entry lists it neither by name nor by
// citation.
func generalizes(entry, technique string) bool {
	x, y := conceptKey(entry), conceptKey(technique)
	if len(x) == 0 || len(x) >= len(y) {
		return false
	}
	for start := 0; start+len(x) <= len(y); start++ {
		if slices.Equal(x, y[start:start+len(x)]) {
			return true
		}
	}
	return false
}

// passageLeadingTerm is the name a passage starts with as an entry, as in
// "Python (…): …", after a technique page's title; empty when it starts
// with none.
func passageLeadingTerm(span EvidenceSpan) string {
	text := span.Text
	if strings.HasPrefix(span.DocumentID, "character-technique:") {
		text = strings.TrimPrefix(text, techniqueTitle(span.DocumentID)+": ")
	}
	if match := leadingTerm.FindStringSubmatch(text); match != nil {
		return sourceTechniqueName(match[1])
	}
	return ""
}

// namesSubtechnique reports a passage that names a technique as a part of a
// form inside the form's own section: a heading of the passage's section
// path is the form's name, and below it a heading or the passage's leading
// term is the technique's name. "Kong Gun: …" under "Devil Fruit > Gear 4"
// names Kong Gun as a part of Gear 4; a Gear 4 passage that only mentions
// Pistol does not.
func namesSubtechnique(span EvidenceSpan, form, technique string) bool {
	path := sectionPath(span.Section)
	for depth, heading := range path {
		if !conceptEqual(heading, form) {
			continue
		}
		for _, below := range path[depth+1:] {
			if conceptEqual(below, technique) {
				return true
			}
		}
		return conceptEqual(passageLeadingTerm(span), technique)
	}
	return false
}

// coreIssue is one core concept failure: its facts, the rule, then the fix.
func coreIssue(path, facts, fix string) m.Issue {
	return m.Issue{Path: path, Message: facts + " " + m.CoreConceptsRule + " " + fix}
}

// CoreConceptIssues checks a plan under requireCoreConcepts for a request
// whose passages name its source techniques (sourceTechniques). A request
// prepared before source techniques existed is not checked, and a saved
// plan is never checked again, so older Results keep their findings.
func CoreConceptIssues(plan DesignPlan, request *Request) []m.Issue {
	ranking, listing := coreConceptIssues(plan, request)
	return append(ranking, listing...)
}

// coreConceptIssues are CoreConceptIssues apart: the ranking and adaptation
// issues, and the listing issue that a targeted plan correction can fix.
func coreConceptIssues(plan DesignPlan, request *Request) (ranking, listing []m.Issue) {
	if request.MechanicsDefinition == nil || !coreConceptsOn(*request.MechanicsDefinition) || request.SourceTechniques == nil {
		return nil, nil
	}
	var issues []m.Issue
	var unranked, core []string
	for _, entry := range plan.Repertoire {
		switch entry.Importance {
		case "":
			unranked = append(unranked, entry.Name)
		case "core":
			core = append(core, entry.Name)
		}
	}
	if len(unranked) > 0 {
		issues = append(issues, coreIssue("repertoire", fmt.Sprintf("Repertoire %s no importance.", rankedList(unranked)), "Rank each entry core, major or minor from the sources."))
	}
	var unrankedOmissions []string
	for i, omission := range plan.OmittedTechniques {
		switch {
		case omission.Importance == "":
			unrankedOmissions = append(unrankedOmissions, omission.Name)
		case omission.Importance == "core":
			issues = append(issues, coreIssue(fmt.Sprintf("omittedTechniques.%d", i), fmt.Sprintf("omittedTechniques omits the core technique %q whole.", omission.Name), "Adapt it on a purchase, as a typed change or a proposed mechanic, and omit only the aspects the Definition cannot express."))
		}
		// Only the whole entry counts, by its name or with a qualifier word
		// ("Gear 4 forms"): an omitted aspect such as "Gear 4 flight" beside
		// a core "Gear 4" is allowed.
		for _, name := range core {
			if listsByName(omission.Name, name) {
				issues = append(issues, coreIssue(fmt.Sprintf("omittedTechniques.%d", i), fmt.Sprintf("omittedTechniques omits the core entry %q whole.", name), "Remove it from omittedTechniques and name each aspect the Definition cannot express as an effect of that entry with adaptedAs empty."))
			}
		}
	}
	if len(unrankedOmissions) > 0 {
		issues = append(issues, coreIssue("omittedTechniques", fmt.Sprintf("omittedTechniques %s no importance.", rankedList(unrankedOmissions)), "Rank each omitted technique major or minor; a core technique is adapted, not omitted."))
	}
	switch {
	case len(core) == 0:
		issues = append(issues, coreIssue("repertoire", "No repertoire entry is core.", "Rank as core the one to three concepts the sources show the character is known for, and adapt each on a purchase."))
	case len(core) > maxCoreConcepts:
		issues = append(issues, coreIssue("repertoire", fmt.Sprintf("%d repertoire entries are core: %s.", len(core), joinWith(quoted(core), "and", "")), "Keep the one to three concepts vital to the character as core and rank the others major or minor."))
	}
	if plan.UpgradeIntents != nil {
		for i, entry := range plan.Repertoire {
			if entry.Importance != "core" || sameTechnique(entry.Name, plan.Base.Name) {
				continue
			}
			named, adapted := false, false
			for pathIndex := range m.PathKeys {
				for tier := 1; tier <= len(m.TierKeys); tier++ {
					intent := *plan.UpgradeIntents.At(pathIndex).At(tier)
					if sameTechnique(intent.Technique, entry.Name) {
						named = true
						adapted = adapted || typedAdaptation(entry, intent) || len(intent.ProposedMechanics) > 0
					}
				}
			}
			switch {
			case !named:
				issues = append(issues, coreIssue(fmt.Sprintf("repertoire.%d", i), fmt.Sprintf("No purchase names the core entry %q as its technique.", entry.Name), "Make it the technique of a purchase whose typed changes do what it does in play, ideally one that defines a path, or name the mechanic it needs in that milestone's proposedMechanics."))
			case !adapted:
				issues = append(issues, coreIssue(fmt.Sprintf("repertoire.%d", i), fmt.Sprintf("No purchase whose technique is the core entry %q promises what one of its effects is adapted as or names a proposed mechanic.", entry.Name), "Adapt one of its effects as a promise of a purchase whose technique it is, or name the mechanic it needs in that milestone's proposedMechanics."))
			}
		}
	}
	issues = append(issues, majorFloorIssues(plan, *request.SourceTechniques)...)
	return issues, sourceListingIssues(plan, request)
}

// unlistedSourceTechniques are the source techniques a plan lists nowhere
// (sourceTechniqueListed).
func unlistedSourceTechniques(plan DesignPlan, request *Request) []SourceTechnique {
	techniques := *request.SourceTechniques
	spans := map[string]EvidenceSpan{}
	for _, span := range AuthorEvidence(request) {
		spans[span.ID] = span
	}
	var missing []SourceTechnique
	for index, technique := range techniques {
		if !sourceTechniqueListed(plan, techniques, index, spans) {
			missing = append(missing, technique)
		}
	}
	return missing
}

// sourceListingIssues names the source techniques a plan lists nowhere, in
// one issue. A targeted plan correction can fix it (planCorrections).
func sourceListingIssues(plan DesignPlan, request *Request) []m.Issue {
	var missing []string
	for _, technique := range unlistedSourceTechniques(plan, request) {
		missing = append(missing, technique.Name)
	}
	if len(missing) == 0 {
		return nil
	}
	return []m.Issue{coreIssue("repertoire", fmt.Sprintf("The plan leaves out the source %s, which %s neither in the repertoire nor in omittedTechniques.", pluralWord(len(missing), "technique", "techniques")+" "+joinWith(quoted(missing), "and", ""), pluralWord(len(missing), "is", "are")), "Name each in the repertoire, or omit it in omittedTechniques with the exact effect the Definition cannot express; a technique that is a part of a form in the form's own section is also listed by that form's entry citing the passage that names it, unless its salience is strong: a strong technique needs an entry of its own.")}
}

// majorFloorIssues reports each plan entry ranked minor that is named for a
// source technique of strong salience, or for one of its aliases: strong
// salience is a major floor, never a core requirement (SOL-61-05). The base
// attack carries no rank and is not checked.
func majorFloorIssues(plan DesignPlan, techniques []SourceTechnique) []m.Issue {
	var issues []m.Issue
	check := func(path, name, importance string) {
		if importance != "minor" {
			return
		}
		for _, technique := range techniques {
			if technique.Salience == SalienceStrong && listsTechnique(name, technique) {
				issues = append(issues, m.Issue{Path: path, Message: fmt.Sprintf("%q is ranked minor, but it lists the source technique %q, whose salience is strong. %s Rank it major or core.", name, technique.Name, SalienceRule)})
				return
			}
		}
	}
	for i, entry := range plan.Repertoire {
		check(fmt.Sprintf("repertoire.%d.importance", i), entry.Name, entry.Importance)
	}
	for i, omission := range plan.OmittedTechniques {
		check(fmt.Sprintf("omittedTechniques.%d.importance", i), omission.Name, omission.Importance)
	}
	return issues
}

// sourceTechniqueListed reports a source technique that the plan lists: the
// base attack, a repertoire entry or an omitted technique is named for it or
// for one of its aliases (listsByName), or the base attack or a repertoire
// entry is named for a form and cites a passage that names the technique as
// a part of that form inside the form's own section (namesSubtechnique), as
// a Gear 4 entry citing "Python (…): …" under "Gear 4" lists Python. Citing
// a passage alone lists nothing: a Pistol entry citing a Gear 4 passage does
// not list Gear 4. A strong technique is listed only by name, so it always has
// an entry of its own whose rank the major floor checks (SOL-69-01).
func sourceTechniqueListed(plan DesignPlan, techniques []SourceTechnique, index int, spans map[string]EvidenceSpan) bool {
	technique := techniques[index]
	type entry struct {
		name string
		ids  []string
	}
	citing := []entry{{plan.Base.Name, plan.Base.SourceIDs}}
	for _, repertoire := range plan.Repertoire {
		citing = append(citing, entry{repertoire.Name, repertoire.SourceIDs})
	}
	for _, e := range citing {
		if listsTechnique(e.name, technique) {
			return true
		}
	}
	for _, omission := range plan.OmittedTechniques {
		if listsTechnique(omission.Name, technique) {
			return true
		}
	}
	if technique.Salience == SalienceStrong {
		return false
	}
	for _, e := range citing {
		if slices.ContainsFunc(technique.Names(), func(name string) bool { return generalizes(e.name, name) }) {
			continue
		}
		for formIndex, form := range techniques {
			if formIndex == index || !listsTechnique(e.name, form) {
				continue
			}
			for _, id := range e.ids {
				span, ok := spans[id]
				if !ok || !slices.Contains(technique.PassageIDs, id) {
					continue
				}
				for _, formName := range form.Names() {
					for _, name := range technique.Names() {
						if namesSubtechnique(span, formName, name) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

// listsTechnique reports a plan entry named for a source technique or for
// one of its aliases (listsByName): an entry named "The Divine Axe Rhitta"
// lists Rhitta when one passage proves the two names are one axe.
func listsTechnique(entry string, technique SourceTechnique) bool {
	return slices.ContainsFunc(technique.Names(), func(name string) bool { return listsByName(entry, name) })
}

// conceptEqual reports two names with the same words.
func conceptEqual(a, b string) bool { return slices.Equal(conceptKey(a), conceptKey(b)) }

// rankedList names entries without a rank: `"A" has` or `"A" and "B" have`.
func rankedList(names []string) string {
	return joinWith(quoted(names), "and", "") + " " + pluralWord(len(names), "has", "have")
}

func quoted(names []string) []string {
	out := make([]string, len(names))
	for i, name := range names {
		out[i] = fmt.Sprintf("%q", name)
	}
	return out
}

func pluralWord(n int, one, many string) string {
	if n == 1 {
		return one
	}
	return many
}

// CoreConceptFindingRule names the check Finding of a core concept that the
// Unit does not embody.
const CoreConceptFindingRule = "core-concept"

// typedAdaptation reports a planned purchase that adapts an entry with a
// typed change: it promises an improvement or unlock that one of the
// entry's effects is adapted as (promiseMatches), or any improvement or
// unlock when the entry lists no effects. The promise checks then hold the
// purchase to that typed change.
func typedAdaptation(entry PlanRepertoire, intent UpgradeIntent) bool {
	promised := slices.Clone(intent.Improves)
	if intent.Unlock != "" && intent.Unlock != "none" {
		promised = append(promised, intent.Unlock)
	}
	if len(entry.Effects) == 0 {
		return len(promised) > 0
	}
	for _, effect := range entry.Effects {
		for _, id := range effect.AdaptedAs {
			if slices.ContainsFunc(promised, func(p string) bool { return promiseMatches(id, p) }) {
				return true
			}
		}
	}
	return false
}

// checkCoreConcepts reports each core entry of a retained plan that the
// Unit does not embody. A core entry is embodied when it is the base attack
// or when a purchase whose technique it is adapts one of its effects with a
// typed change (typedAdaptation). One that those purchases adapt only with
// proposed mechanics is an unresolved design gap (onlyProposedFinding): no
// build grants a proposal, so the Unit embodies the concept only once the
// Definition supports the mechanic. One they adapt with neither fails.
// Whether the typed change carries the concept's central effect is the
// review's judgment (CoreSpiritRule). Like CoreConceptIssues it runs only
// under requireCoreConcepts for a request that lists its source techniques.
func checkCoreConcepts(blueprint *m.Blueprint, plan *DesignPlan, request Request, report reporter) {
	if !coreConceptsChecked(blueprint, plan, request) {
		return
	}
	documents := evidenceDocuments(&request)
	for index, entry := range plan.Repertoire {
		if entry.Importance != "core" || sameTechnique(entry.Name, plan.Base.Name) {
			continue
		}
		e := embodiment(blueprint, plan, []int{index}, documents)
		subject := fmt.Sprintf("designPlan.repertoire.%d", index)
		switch {
		case e.typed:
			continue
		case e.onlyProposed():
			report(e.onlyProposedFinding("core concept", entry.Name, CoreConceptFindingRule, m.CoreConceptsRule))
		case len(e.named) == 0:
			report(checkFinding{
				Category: "conflict", Outcome: "fail", Subject: subject, Rule: CoreConceptFindingRule,
				Message: fmt.Sprintf("No purchase names the core concept %q as its technique. %s", entry.Name, m.CoreConceptsRule),
				Action:  act("Make it the technique of a purchase whose typed changes do what it does in play."),
			})
		default:
			report(checkFinding{
				Category: "conflict", Outcome: "fail", Subject: subject, Rule: CoreConceptFindingRule,
				Message: fmt.Sprintf("No purchase adapts the core concept %q: %s %s none of its effects with a typed change and %s no proposed mechanic. %s", entry.Name, joinWith(e.named, "and", ""), pluralWord(len(e.named), "adapts", "adapt"), pluralWord(len(e.named), "carries", "carry"), m.CoreConceptsRule),
				Action:  act("Promise, on a purchase whose technique it is, a typed change that one of its effects is adapted as, or carry what the Definition cannot express as a proposed mechanic."),
			})
		}
	}
}

// coreConceptsChecked reports whether checkCoreConcepts judges a checked
// Unit's core entries: a typed Unit with a retained plan that names its
// purchases' techniques, under requireCoreConcepts, for a request that
// lists its source techniques.
func coreConceptsChecked(blueprint *m.Blueprint, plan *DesignPlan, request Request) bool {
	definition := request.MechanicsDefinition
	return blueprint != nil && plan != nil && plan.UpgradeIntents != nil && definition != nil && coreConceptsOn(*definition) && request.SourceTechniques != nil
}

// coreGap reports a core entry that checkCoreConcepts reports as only
// proposed, so a required concept named for it needs no gap of its own.
func coreGap(blueprint *m.Blueprint, plan *DesignPlan, request Request, index int, documents map[string]string) bool {
	entry := plan.Repertoire[index]
	return coreConceptsChecked(blueprint, plan, request) && entry.Importance == "core" && !sameTechnique(entry.Name, plan.Base.Name) &&
		embodiment(blueprint, plan, []int{index}, documents).onlyProposed()
}

// evidenceDocuments maps each passage ID of a request to its document.
func evidenceDocuments(request *Request) map[string]string {
	documents := map[string]string{}
	for _, span := range AuthorEvidence(request) {
		documents[span.ID] = span.DocumentID
	}
	return documents
}

// conceptEmbodiment is how the purchases whose technique is one of a
// plan's repertoire entries adapt it on the checked Unit: whether one
// adapts one of its effects with a typed change, the purchases named for
// it, those that carry proposed mechanics with the proposals' names and the
// documents they cite, and the blueprint path of the first proposing
// purchase.
type conceptEmbodiment struct {
	typed                                 bool
	named, proposing, proposals, evidence []string
	subject                               string
}

// embodiment is how a checked Unit's purchases adapt the given repertoire
// entries, which name one concept.
func embodiment(blueprint *m.Blueprint, plan *DesignPlan, entries []int, documents map[string]string) conceptEmbodiment {
	e := conceptEmbodiment{evidence: []string{}}
	for pathIndex, path := range m.PathKeys {
		for tier := 1; tier <= len(m.TierKeys); tier++ {
			intent := *plan.UpgradeIntents.At(pathIndex).At(tier)
			index := slices.IndexFunc(entries, func(i int) bool { return sameTechnique(intent.Technique, plan.Repertoire[i].Name) })
			if index < 0 {
				continue
			}
			upgrade := blueprint.Paths.At(pathIndex).Tiers.At(tier)
			code := BuildCode(pathIndex, tier) + " " + upgrade.Name
			e.named = append(e.named, code)
			e.typed = e.typed || typedAdaptation(plan.Repertoire[entries[index]], intent)
			if len(upgrade.ProposedMechanics) == 0 {
				continue
			}
			if len(e.proposing) == 0 {
				e.subject = fmt.Sprintf("paths.%s.tiers.%s", path, m.TierKeys[tier-1])
			}
			e.proposing = append(e.proposing, code)
			for _, proposed := range upgrade.ProposedMechanics {
				e.proposals = append(e.proposals, proposed.Name)
				for _, id := range proposed.SourceIDs {
					if document, ok := documents[id]; ok && !slices.Contains(e.evidence, document) {
						e.evidence = append(e.evidence, document)
					}
				}
			}
		}
	}
	return e
}

// onlyProposed reports a concept that its purchases adapt only with
// proposed mechanics: a design gap.
func (e conceptEmbodiment) onlyProposed() bool { return !e.typed && len(e.proposing) > 0 }

// onlyProposedFinding is the unresolved design gap of a concept, a core
// concept or a required concept, that its purchases adapt only with
// proposed mechanics, under its rule and stating its rule sentence.
func (e conceptEmbodiment) onlyProposedFinding(kind, name, rule, statement string) checkFinding {
	return checkFinding{
		Category: "missing_specification", Outcome: "unresolved", Subject: e.subject, Rule: rule,
		Message: fmt.Sprintf("The %s %q is only proposed: %s %s it as the proposed %s %s, and no purchase whose technique it is adapts one of its effects with a typed change. A proposed mechanic is not yet playable: until the Definition supports it, no build grants it and the Unit does not embody this %s. %s",
			kind, name, joinWith(e.proposing, "and", ""), pluralWord(len(e.proposing), "carries", "carry"), pluralWord(len(e.proposals), "mechanic", "mechanics"), joinWith(e.proposals, "and", ""), kind, statement),
		Action:   act("Adapt one of its effects with a typed change on a purchase whose technique it is, or expand the Definition with its proposed mechanic. Until then the Unit does not embody the concept in play."),
		Evidence: e.evidence,
	}
}
