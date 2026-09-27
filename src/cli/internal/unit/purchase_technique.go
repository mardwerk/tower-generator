package unit

import (
	"slices"
	"strings"
	"unicode"
)

// PurchaseTechnique is the technique the retained plan maps a purchase to,
// and whether that technique is the base attack. It is the plan's typed
// mapping, which the technique checks hold to the plan's ownership rules; the
// unit shows it beside the purchase's resolved effects, never in place of
// them (#35). It is empty for a plan without typed techniques.
func PurchaseTechnique(plan *DesignPlan, pathIndex, tier int) (technique string, base bool) {
	if plan == nil || plan.UpgradeIntents == nil {
		return "", false
	}
	technique = strings.TrimSpace(plan.UpgradeIntents.At(pathIndex).At(tier).Technique)
	return technique, technique != "" && strings.EqualFold(technique, strings.TrimSpace(plan.Base.Name))
}

// nameWords are a name's lowercase words of at least four letters or
// digits; shorter words such as "of" or "2" identify nothing.
func nameWords(name string) []string {
	var words []string
	for _, word := range strings.FieldsFunc(strings.ToLower(name), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}) {
		if len([]rune(word)) >= 4 && !slices.Contains(words, word) {
			words = append(words, word)
		}
	}
	return words
}

// PurchaseNameMentions lists the repertoire entries a purchase's name points
// to although the plan maps the purchase to another technique, such as
// "Boundman Force" on a purchase of the base attack when the repertoire has
// "Gear 4 Boundman" (a Luffy Result on #31). Only a distinctive word counts:
// one that appears in exactly one name among the repertoire entries and the
// base attack. Shared words such as "Gear" or "Haki" implicate no entry.
// Words of the purchase's own technique never count, nor words of the base
// attack, which every purchase changes. It is advisory, shown for review: it
// never remaps a purchase or fails a check. It is empty for a plan without
// typed techniques.
func PurchaseNameMentions(plan *DesignPlan, pathIndex, tier int, name string) []string {
	technique, _ := PurchaseTechnique(plan, pathIndex, tier)
	if technique == "" {
		return nil
	}
	owners := map[string][]string{}
	for _, entry := range append([]string{plan.Base.Name}, repertoireNames(plan)...) {
		entry = strings.TrimSpace(entry)
		for _, word := range nameWords(entry) {
			if !slices.ContainsFunc(owners[word], func(owner string) bool { return strings.EqualFold(owner, entry) }) {
				owners[word] = append(owners[word], entry)
			}
		}
	}
	own := append(nameWords(technique), nameWords(plan.Base.Name)...)
	var mentions []string
	for _, word := range nameWords(name) {
		if slices.Contains(own, word) || len(owners[word]) != 1 {
			continue
		}
		if entry := owners[word][0]; !strings.EqualFold(entry, technique) && !slices.Contains(mentions, entry) {
			mentions = append(mentions, entry)
		}
	}
	return mentions
}

func repertoireNames(plan *DesignPlan) []string {
	names := make([]string, len(plan.Repertoire))
	for i, entry := range plan.Repertoire {
		names[i] = entry.Name
	}
	return names
}
