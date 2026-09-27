package unit

import (
	"fmt"
	"slices"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Proposed mechanics are what a purchase needs that its Definition cannot
// express yet (#61). Each stays on its purchase beside the typed changes
// that approximate it, is flagged for Definition expansion by an unresolved
// Finding, and is judged by the review. No build grants one: build
// resolution reads only typed changes. A milestone may name a proposed
// mechanic, and BindDesignPlan keeps it on the purchase.

// ProposedMechanicRule marks the Finding for each proposed mechanic of a
// purchase.
const ProposedMechanicRule = "proposed-mechanic"

// bindProposedMechanics keeps each proposed mechanic a milestone names on
// its purchase: the plan's come first, then those the mechanics stage adds
// under another name. Names match without regard to case or spacing.
func bindProposedMechanics(path *s.Object, intents *PathIntents) {
	tiers, ok := field(path, "tiers").(*s.Object)
	if !ok {
		return
	}
	for number, key := range m.TierKeys {
		planned := intents.At(number + 1).ProposedMechanics
		tier, ok := field(tiers, key).(*s.Object)
		if !ok || len(planned) == 0 {
			continue
		}
		merged := []any{}
		names := map[string]bool{}
		for _, proposed := range planned {
			names[proposalKey(proposed.Name)] = true
			merged = append(merged, s.FromGoValue(proposed))
		}
		authored, _ := field(tier, "proposedMechanics").([]any)
		for _, raw := range authored {
			if object, ok := raw.(*s.Object); ok {
				if name, _ := field(object, "name").(string); names[proposalKey(name)] {
					continue
				}
			}
			merged = append(merged, raw)
		}
		tier.Set("proposedMechanics", merged)
	}
}

func proposalKey(name string) string { return strings.ToLower(strings.Join(strings.Fields(name), " ")) }

// proposedMechanicIssues checks that each proposed mechanic of a purchase
// cites passages of the request's evidence and that no purchase names one
// twice.
func proposedMechanicIssues(blueprint m.Blueprint, request Request) []m.Issue {
	known := map[string]bool{}
	for _, span := range AuthorEvidence(&request) {
		known[span.ID] = true
	}
	var issues []m.Issue
	for pathIndex, path := range m.PathKeys {
		for number, tier := range m.TierKeys {
			names := map[string]bool{}
			for i, proposed := range blueprint.Paths.At(pathIndex).Tiers.At(number + 1).ProposedMechanics {
				at := fmt.Sprintf("paths.%s.tiers.%s.proposedMechanics.%d", path, tier, i)
				if key := proposalKey(proposed.Name); names[key] {
					issues = append(issues, m.Issue{Path: at, Message: "Name each proposed mechanic of a purchase once."})
				} else {
					names[key] = true
				}
				for _, id := range proposed.SourceIDs {
					if !known[id] {
						issues = append(issues, m.Issue{Path: at + ".sourceIds", Message: "Cite a supplied source passage ID; " + id + " is not one."})
					}
				}
			}
		}
	}
	return issues
}

// PlannedProposalIssues reports a proposed mechanic a milestone names that
// its purchase no longer carries. BindDesignPlan keeps each one when a unit
// is drafted, so only an edited blueprint can lose one.
func PlannedProposalIssues(blueprint m.Blueprint, intents *UpgradeIntents) []m.Issue {
	if intents == nil {
		return nil
	}
	var issues []m.Issue
	for pathIndex, path := range m.PathKeys {
		for number, tier := range m.TierKeys {
			carried := map[string]bool{}
			for _, proposed := range blueprint.Paths.At(pathIndex).Tiers.At(number + 1).ProposedMechanics {
				carried[proposalKey(proposed.Name)] = true
			}
			for _, planned := range intents.At(pathIndex).At(number + 1).ProposedMechanics {
				if !carried[proposalKey(planned.Name)] {
					issues = append(issues, m.Issue{
						Path:    fmt.Sprintf("paths.%s.tiers.%s.proposedMechanics", path, tier),
						Message: fmt.Sprintf("The retained plan proposes %s at %s, but the purchase does not carry it. Keep it on the purchase as a proposed mechanic.", planned.Name, BuildCode(pathIndex, number+1)),
					})
				}
			}
		}
	}
	return issues
}

// checkProposedMechanics reports each proposed mechanic of a purchase as an
// unresolved Finding: the Definition must be expanded before any build can
// grant it.
func checkProposedMechanics(blueprint *m.Blueprint, request Request, report reporter) {
	if blueprint == nil {
		return
	}
	documents := map[string]string{}
	for _, span := range AuthorEvidence(&request) {
		documents[span.ID] = span.DocumentID
	}
	for pathIndex, path := range m.PathKeys {
		for number, tier := range m.TierKeys {
			upgrade := blueprint.Paths.At(pathIndex).Tiers.At(number + 1)
			for i, proposed := range upgrade.ProposedMechanics {
				evidence := []string{}
				for _, id := range proposed.SourceIDs {
					if document, ok := documents[id]; ok && !slices.Contains(evidence, document) {
						evidence = append(evidence, document)
					}
				}
				code := BuildCode(pathIndex, number+1)
				report(checkFinding{
					Category: "unsupported", Outcome: "unresolved",
					Subject: fmt.Sprintf("paths.%s.tiers.%s.proposedMechanics.%d", path, tier, i), Rule: ProposedMechanicRule,
					Message: fmt.Sprintf("%s %s proposes %s: %s The Definition must be expanded before this purchase can grant it; until then no build grants it.",
						code, upgrade.Name, proposed.Name, sentence(proposed.Effect)),
					Action:   act("Expand the Definition with this mechanic, or keep it proposed. The review judges whether it fits the purchase and its cited passages."),
					Evidence: evidence,
				})
			}
		}
	}
}

// sentence ends text with a full stop.
func sentence(text string) string {
	text = strings.TrimSpace(text)
	if text == "" || strings.HasSuffix(text, ".") || strings.HasSuffix(text, "!") || strings.HasSuffix(text, "?") {
		return text
	}
	return text + "."
}
