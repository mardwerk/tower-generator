package render

import (
	"fmt"
	"strings"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Detailed renders the full reading view: purchases, usage, abilities,
// mechanics, review counts, open findings, the authoring checks that fail
// today (AuthoringIssues) and evidence.
func Detailed(view View) string {
	var lines []string
	lines = append(lines, describeUnit(view)...)
	lines = append(lines, describePurchases(view)...)
	lines = append(lines, describeUsage(view)...)
	lines = append(lines, describeAbilities(view)...)
	lines = append(lines, describeMechanics(view)...)
	lines = append(lines, describeReview(view)...)
	lines = append(lines, currentChecks(view)...)
	lines = append(lines, describeEvidence(view)...)
	return strings.Join(lines, "\n") + "\n"
}

func describeUnit(view View) []string {
	candidate, prepared := view.Candidate, view.Prepared
	result := "Intermediate artifact."
	if view.ResultID != nil && *view.ResultID != "" {
		result = "Result: `" + *view.ResultID + "`."
	}
	lines := []string{
		"# " + Escape(candidate.Character.Name),
		"",
		"Authoring candidate for " + Escape(candidate.Character.Work) + ". Scope: " + Escape(candidate.Character.Scope),
		"",
		"Input: `" + prepared.InputHash + "`. " + result,
		"",
	}
	if definition := prepared.Request.MechanicsDefinition; definition != nil {
		rules := []string{}
		for _, document := range prepared.Request.Documents {
			if document.Kind == "rules" && !strings.HasPrefix(document.ID, "mechanics:") {
				rules = append(rules, document.ID)
			}
		}
		lines = append(lines, Escape(fmt.Sprintf("Definition %s, %s, revision %s. Rules: %s.", definition.Label, definition.ID, definition.Revision, strings.Join(rules, ", "))), "")
	}
	lines = append(lines, Escape(candidate.Role), "")
	if candidate.Blueprint != nil && candidate.Blueprint.ReferencePattern != nil {
		pattern := candidate.Blueprint.ReferencePattern
		lines = append(lines,
			"Recorded reference pattern: "+Escape(pattern.ID)+", version "+Escape(pattern.Version)+". Initial numerical mechanics and prices came from this fixed proposed pattern. This is authoring history, not proof that external edits preserved those mechanics or that the design is balanced.",
			"")
	}
	attack := candidate.BasicAttack
	lines = append(lines,
		"## Basic attack", "",
		Escape(attack.Name)+" ("+attack.Status+"). "+Escape(attack.Behavior), "",
		"Delivery: "+Escape(attack.Delivery), "",
		"Targeting: "+Escape(attack.Targeting), "",
		"Limits: "+Escape(attack.Limitations), "",
		"## Upgrade paths", "",
	)
	resolved := resolvedEffects(view)
	for index, path := range candidate.Paths {
		lines = append(lines,
			"### "+Escape(path.Name), "",
			Escape(path.Theme), "",
			"| Build | Upgrade | Effect | Status |",
			"| --- | --- | --- | --- |",
		)
		for _, tier := range path.Tiers {
			effect := tier.Benefit
			if purchase, ok := resolved[TierKey(path.ID, tier.Tier)]; ok {
				effect = purchase.price + ". " + purchase.effects
			}
			lines = append(lines, fmt.Sprintf("| %s | %s | %s | %s |", purchaseLabel(len(candidate.Paths), index, tier.Tier), Escape(tier.Name), Escape(effect), tier.Status))
		}
		lines = append(lines, "")
	}
	return lines
}

// resolvedPurchase is one purchase of a typed unit as the unit sheet words
// it: its build code, price and effects.
type resolvedPurchase struct{ code, price, effects string }

// resolvedEffects words each purchase of a typed unit as the unit sheet
// does, keyed by path ID and tier ("path-2:4") as KitStats names paths, so
// the report shows its changes with their deltas and multipliers as
// percentages. The tier benefits and ability descriptions stored in the
// candidate keep their compiled wording, which the review reads; the report
// shows them only for a unit without valid typed mechanics.
func resolvedEffects(view View) map[string]resolvedPurchase {
	sh := newSheet(view.Candidate.Blueprint, view.Prepared.Request.MechanicsDefinition)
	if sh == nil {
		return nil
	}
	out := map[string]resolvedPurchase{}
	for index, path := range sh.purchases() {
		for tier, purchase := range path.Purchases {
			out[TierKey("path-"+itoa(index+1), tier+1)] = resolvedPurchase{purchase.Code, sh.money(purchase.Cost), strings.Join(purchase.Effects, " ")}
		}
	}
	return out
}

// field reads a key path from a JSON value; nil when absent.
func field(value any, path ...string) any {
	for _, key := range path {
		object, ok := value.(*s.Object)
		if !ok {
			return nil
		}
		value, _ = object.Get(key)
	}
	return value
}

func list(value any) []any {
	items, _ := value.([]any)
	return items
}

func text(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case float64:
		return s.FormatNumber(v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	case nil:
		return "null"
	}
	return fmt.Sprint(value)
}

func joined(value any, separator string) string {
	items := list(value)
	parts := make([]string, len(items))
	for i, item := range items {
		if item != nil {
			parts[i] = text(item)
		}
	}
	return strings.Join(parts, separator)
}

func measured(value any) string {
	number, ok := value.(float64)
	if !ok {
		return "unavailable"
	}
	return Escape(s.FormatNumber(toPrecision(number, 4)))
}

// gold writes a price exactly, as the unit sheet does: rounding 24615 to
// four significant digits would misstate a total.
func gold(value any) string {
	number, ok := value.(float64)
	if !ok {
		return "unavailable"
	}
	return Escape(price(number))
}

func describePurchases(view View) []string {
	if view.DesignEvaluation == nil {
		return nil
	}
	currency := ""
	if definition := view.Prepared.Request.MechanicsDefinition; definition != nil {
		currency = definition.Profile.Currency
	}
	evaluation := any(unit.ReviewPurchaseEvidence(view.DesignEvaluation, currency))
	if currency == "" {
		currency = "currency"
	}
	currency = Escape(currency)
	lines := []string{
		"## Purchase evidence",
		"",
		"Calculated from the resolved builds. Throughput assumes eligible targets continuously in reach; group values are capacity upper bounds. Time-averaged rates use the Active Ability whenever it is ready and equal the ordinary rate without one. Each crosspath purchase states its absolute time-averaged gain and price beside those of its main path's own fifth purchase. These comparisons do not prove balance, source fidelity or player preference.",
		"",
	}
	for _, path := range list(field(evaluation, "paths")) {
		// The plan's purchase reasons, weaknesses and capstone notes are
		// private design checks; the JSON artifact keeps them for review.
		lines = append(lines, "### "+Escape(text(field(path, "name"))), "")
		lines = append(lines, "| Purchase | Added "+currency+" | Changed capacities |", "| --- | --- | --- |")
		purchases := append(append([]any{}, list(field(path, "milestones"))...), list(field(path, "crosspaths"))...)
		for _, purchase := range purchases {
			var deltas []string
			if metrics, ok := field(purchase, "metricDeltas").(*s.Object); ok {
				for _, metric := range metrics.Keys() {
					delta, _ := metrics.Get(metric)
					if change, ok := field(delta, "change").(float64); ok && change == 0 {
						continue
					}
					deltas = append(deltas, Escape(metric)+": "+measured(field(delta, "before"))+" → "+measured(field(delta, "after")))
				}
			}
			changed := strings.Join(deltas, "; ")
			if len(deltas) == 0 {
				changed = joined(field(purchase, "capabilityChanges"), "; ")
				if changed == "" {
					changed = "No change in measured capacities."
				}
				changed = Escape(changed)
			}
			if against := againstCapstoneText(field(purchase, unit.AgainstCapstone), currency); against != "" {
				if !strings.HasSuffix(changed, ".") {
					changed += "."
				}
				changed += " " + against
			}
			lines = append(lines, "| "+joined(field(purchase, "from"), "-")+" → "+joined(field(purchase, "to"), "-")+" | "+gold(field(purchase, "incrementalGold"))+" | "+changed+" |")
		}
		comparison := field(path, "capstoneComparison")
		copies := "an undefined number of"
		if value := field(comparison, "tier4CopiesAtTier5Budget"); value != nil {
			copies = text(value)
		}
		lines = append(lines,
			"",
			"T5 total: "+gold(field(comparison, "tier5", "totalGold"))+" "+currency+". The same budget buys "+copies+" pure T4 copies. Extra copies need extra placement space and target access. Range and active uptime do not add across copies.",
			"",
		)
	}
	return lines
}

// againstCapstoneText states a side purchase's absolute time-averaged gain
// and price beside those of its main path's fifth purchase, such as "Side
// 1-x-x adds direct +1.05, group +2.1 for 140 Gold; the path's x-5-x adds
// direct +12.64, group +25.28 for 45,000 Gold."; empty without a comparison.
func againstCapstoneText(against any, currency string) string {
	if against == nil {
		return ""
	}
	adds := func(side string) string {
		return Escape(text(field(against, side, "code"))) + " adds direct " + signed(field(against, unit.TimeAveragedDirect, side)) + ", group " + signed(field(against, unit.TimeAveragedGroup, side)) + " for " + gold(field(against, side, "price")) + " " + currency
	}
	return "Side " + adds("sidePurchase") + "; the path's " + adds("capstone") + "."
}

// signed writes a measured change with its sign.
func signed(value any) string {
	if number, ok := value.(float64); ok && number >= 0 {
		return "+" + measured(number)
	}
	return measured(value)
}

func describeUsage(view View) []string {
	lines := []string{
		"",
		"## Generation usage",
		"",
		UsageSummaryText(view.Usage),
		"",
		"Costs are reported USD, not estimates. Partial totals include only reported stages of this revision. Failed or cancelled attempts are excluded; unavailable reports are not zero. Reasoning and cached input tokens are breakdowns, not additional totals.",
		"",
	}
	for _, stage := range view.Usage.Stages {
		lines = append(lines, "### "+stage.Stage, "", "| Metric | Reported value |", "| --- | --- |")
		for _, row := range StageUsageRows(stage) {
			lines = append(lines, "| "+row[0]+" | "+Escape(row[1])+" |")
		}
		lines = append(lines, "")
	}
	return lines
}

func describeAbilities(view View) []string {
	if len(view.Candidate.Abilities) == 0 {
		return nil
	}
	lines := []string{"## Abilities", ""}
	resolved := resolvedEffects(view)
	for _, ability := range view.Candidate.Abilities {
		description := ability.Description
		if ability.PathID != nil && ability.Tier != nil {
			if purchase, ok := resolved[TierKey(*ability.PathID, *ability.Tier)]; ok {
				description = "At " + purchase.code + ": " + purchase.effects
			}
		}
		assignment := "No single path assignment"
		if ability.PathID != nil {
			tier := "unspecified"
			if ability.Tier != nil {
				tier = itoa(*ability.Tier)
			}
			assignment = *ability.PathID + ", tier " + tier
		}
		prerequisites := strings.Join(ability.PrerequisiteAbilityIDs, ", ")
		if prerequisites == "" {
			prerequisites = "None declared"
		}
		lines = append(lines,
			"### "+Escape(ability.Name), "",
			ability.Status+"; "+ability.Placement+". "+Escape(description), "",
			"Assignment: "+Escape(assignment)+". Prerequisites: "+Escape(prerequisites)+".", "",
			"Availability: "+Escape(ability.Availability), "",
			"Delivery: "+Escape(ability.Delivery), "",
			"Targeting: "+Escape(ability.Targeting), "",
			"Limits: "+Escape(ability.Limitations), "",
		)
	}
	return lines
}

func describeMechanics(view View) []string {
	lines := []string{
		"## Mechanics",
		"",
		"| Mechanic | Status | Behavior | Required decision |",
		"| --- | --- | --- | --- |",
	}
	for _, mechanic := range view.Candidate.Mechanics {
		decision := ""
		if mechanic.RequiredDecision != nil {
			decision = *mechanic.RequiredDecision
		}
		lines = append(lines, "| "+Escape(mechanic.Name)+" | "+mechanic.Status+" | "+Escape(mechanic.Behavior)+" | "+Escape(decision)+" |")
	}
	lines = append(lines, "", "## Representative builds", "", "| Build | Selection | Reason |", "| --- | --- | --- |")
	for _, build := range view.Candidate.RepresentativeBuilds {
		selections := make([]string, len(build.Selections))
		for i, selection := range build.Selections {
			selections[i] = selection.PathID + ": " + itoa(selection.Tier)
		}
		lines = append(lines, "| "+Escape(build.Name)+" | "+Escape(strings.Join(selections, ", "))+" | "+Escape(build.Rationale)+" |")
	}
	return lines
}

func describeReview(view View) []string {
	lines := []string{"", "## Review", "", reviewStatus(view), ""}
	if view.ReviewSummary != nil && *view.ReviewSummary != "" {
		lines = append(lines, Escape(*view.ReviewSummary), "")
	}
	for _, method := range []string{"deterministic", "model"} {
		counts := map[string]int{}
		for _, finding := range view.Findings {
			if finding.Method == method {
				counts[finding.Outcome]++
			}
		}
		name := "Deterministic checks"
		if method == "model" {
			name = "Model review"
		}
		lines = append(lines, fmt.Sprintf("%s: %d passed, %d failed, %d unresolved, %d not checked.", name, counts["pass"], counts["fail"], counts["unresolved"], counts["not_checked"]), "")
	}
	lines = append(lines,
		"Passing authoring checks does not certify runtime behavior or balance. The JSON artifact retains every finding, including successful checks and their evidence.",
		"",
	)
	var attention []string
	for _, finding := range view.Findings {
		// Verdicts have their own sections below, passes included.
		if finding.Outcome == "pass" || unit.IsReviewVerdict(finding) {
			continue
		}
		action := ""
		if finding.Action != nil {
			action = *finding.Action
		}
		method := finding.Method
		if finding.Rule == unit.HumanReviewRule {
			method = "human review needed"
		}
		attention = append(attention, "| "+finding.Outcome+" | "+method+" | "+Escape(finding.Subject)+" | "+Escape(finding.Message+" "+action)+" |")
	}
	if len(attention) > 0 {
		lines = append(lines, "| Outcome | Method | Subject | Finding and next action |", "| --- | --- | --- | --- |")
		lines = append(append(lines, attention...), "")
	}
	lines = append(lines, verdictSection(view.Findings, unit.OmissionVerdictRule, "### Omission verdicts",
		"The model review's verdict on each whole-technique omission: whether a supported typed change or a proposed mechanic on a purchase could adapt the technique's central effect, and whether its rank fits the passages it cites.")...)
	lines = append(lines, verdictSection(view.Findings, unit.PathIdentityVerdictRule, "### Third purchase verdicts",
		"The model review's verdict on each path's third purchase: whether it defines or distinguishes its path against its own first and second purchases and the other paths' purchases.")...)
	lines = append(lines, verdictSection(view.Findings, unit.CapstoneVerdictRule, "### Fifth purchase verdicts",
		"The model review's verdict on each path's fifth purchase: whether it is its path's pinnacle, judged with the other two fifth purchases and its path's identity, and a distinct play reason when two fifth purchases buy the same capability kind.")...)
	return lines
}

// verdictSection lists a Result's review verdicts under one rule, a pass
// included; nil when it has none, as a Result reviewed before verdicts.
func verdictSection(findings []unit.Finding, rule, heading, intro string) []string {
	var rows []string
	for _, finding := range findings {
		if finding.Method != "model" || finding.Rule != rule {
			continue
		}
		verdict := finding.Message
		if finding.Action != nil && *finding.Action != "" {
			verdict += " " + *finding.Action
		}
		rows = append(rows, "| "+finding.Outcome+" | "+Escape(finding.Subject)+" | "+Escape(verdict)+" |")
	}
	if len(rows) == 0 {
		return nil
	}
	lines := []string{heading, "", intro, "", "| Outcome | Subject | Verdict and next action |", "| --- | --- | --- |"}
	return append(append(lines, rows...), "")
}

func describeEvidence(view View) []string {
	var lines []string
	if questions := view.Candidate.UnresolvedQuestions; len(questions) > 0 {
		lines = append(lines, "## Open questions", "")
		for _, question := range questions {
			lines = append(lines, "- "+Escape(question.Question)+" Affects: "+Escape(question.Affected)+".")
		}
		lines = append(lines, "")
	}
	lines = append(lines, "## Evidence", "", "| ID | Kind | Origin | Access |", "| --- | --- | --- | --- |")
	for _, document := range view.Prepared.Request.Documents {
		note := ""
		if document.Origin.Note != nil {
			note = *document.Origin.Note
		}
		lines = append(lines, "| "+Escape(document.ID)+" | "+document.Kind+" | "+Escape(document.Origin.Location)+" | "+Escape(document.Origin.Access+". "+note)+" |")
	}
	lines = append(lines, "")
	for _, source := range view.Candidate.Sources {
		lines = append(lines, "- "+Escape(source.DocumentID)+": "+Escape(strings.Join(source.Claims, "; "))+" Limits: "+Escape(source.Limitations))
	}
	return lines
}

// purchaseLabel names a purchase by build code when the unit has the three
// paths build codes describe, and by tier number otherwise.
func purchaseLabel(paths, index, tier int) string {
	if paths != 3 || tier < 1 || tier > 9 {
		return itoa(tier)
	}
	return unit.BuildCode(index, tier)
}
