package unit

import (
	"fmt"
	"slices"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// Core concepts (#61): under requireCoreConcepts the plan ranks every
// repertoire entry and omitted technique core, major or minor, lists every
// source technique in one of the two, and adapts each core entry on a
// purchase. Code checks only the ranking, the listing and that a purchase
// names each core entry as its technique. Whether that purchase fulfils the
// concept in spirit, and whether an omission's reason holds, is the review's
// judgment: prose heuristics for either produced false positives
// (SOL-57-01).

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

// sameConcept reports two names whose words are equal, or where one name's
// words appear in the other's in order, such as "Gear 4" in "Gear 4 forms".
func sameConcept(a, b string) bool {
	x, y := conceptKey(a), conceptKey(b)
	if len(x) == 0 || len(y) == 0 {
		return false
	}
	if len(x) > len(y) {
		x, y = y, x
	}
	for start := 0; start+len(x) <= len(y); start++ {
		if slices.Equal(x, y[start:start+len(x)]) {
			return true
		}
	}
	return false
}

// specificPassages are the passages of a source technique that no other
// source technique lists, or all of its passages when every one is shared.
// A plan entry that cites one of them covers the technique; a passage that
// several techniques share, such as a sentence listing all three kinds of
// Haki, does not cover each of them.
func specificPassages(techniques []SourceTechnique, index int) []string {
	var specific []string
	for _, id := range techniques[index].PassageIDs {
		shared := false
		for other, technique := range techniques {
			if other != index && slices.Contains(technique.PassageIDs, id) {
				shared = true
				break
			}
		}
		if !shared {
			specific = append(specific, id)
		}
	}
	if len(specific) == 0 {
		return techniques[index].PassageIDs
	}
	return specific
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
	if request.MechanicsDefinition == nil || !coreConceptsOn(*request.MechanicsDefinition) || request.SourceTechniques == nil {
		return nil
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
		// Only the whole entry counts: an omitted aspect such as "Gear 4
		// flight" beside a core "Gear 4" is allowed.
		for _, name := range core {
			if conceptEqual(omission.Name, name) {
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
			named := false
			for pathIndex := range m.PathKeys {
				for tier := 1; tier <= len(m.TierKeys); tier++ {
					named = named || sameTechnique(plan.UpgradeIntents.At(pathIndex).At(tier).Technique, entry.Name)
				}
			}
			if !named {
				issues = append(issues, coreIssue(fmt.Sprintf("repertoire.%d", i), fmt.Sprintf("No purchase names the core entry %q as its technique.", entry.Name), "Make it the technique of a purchase whose typed changes or proposed mechanics do what it does in play, ideally one that defines a path."))
			}
		}
	}
	techniques := *request.SourceTechniques
	var missing []string
	for index, technique := range techniques {
		if !sourceTechniqueListed(plan, techniques, index) {
			missing = append(missing, technique.Name)
		}
	}
	if len(missing) > 0 {
		issues = append(issues, coreIssue("repertoire", fmt.Sprintf("The plan leaves out the source %s, which %s neither in the repertoire nor in omittedTechniques.", pluralWord(len(missing), "technique", "techniques")+" "+joinWith(quoted(missing), "and", ""), pluralWord(len(missing), "is", "are")), "Add each to the repertoire, or to an entry that cites its passages, or omit it in omittedTechniques with the exact effect the Definition cannot express."))
	}
	return issues
}

// sourceTechniqueListed reports a source technique that the plan's base
// attack, a repertoire entry or an omitted technique names, or that the base
// attack or a repertoire entry cites by one of its specific passages.
func sourceTechniqueListed(plan DesignPlan, techniques []SourceTechnique, index int) bool {
	technique := techniques[index]
	specific := specificPassages(techniques, index)
	cites := func(ids []string) bool {
		for _, id := range ids {
			if slices.Contains(specific, id) {
				return true
			}
		}
		return false
	}
	if sameConcept(plan.Base.Name, technique.Name) || cites(plan.Base.SourceIDs) {
		return true
	}
	for _, entry := range plan.Repertoire {
		if sameConcept(entry.Name, technique.Name) || cites(entry.SourceIDs) {
			return true
		}
	}
	for _, omission := range plan.OmittedTechniques {
		if sameConcept(omission.Name, technique.Name) {
			return true
		}
	}
	return false
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
