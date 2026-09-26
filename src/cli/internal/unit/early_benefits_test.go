package unit_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

func intent(unlock string, improves ...string) unit.UpgradeIntent {
	return unit.UpgradeIntent{Improves: improves, Unlock: unlock, Technique: "Dart Throw"}
}

// withEarlyPolicy sets distinctEarlyBenefits and distinctFirstUpgrades.
func withEarlyPolicy(d m.Definition, benefits *bool, first bool) m.Definition {
	policy := *d.Profile.DesignPolicy
	policy.DistinctEarlyBenefits, policy.DistinctFirstUpgrades = benefits, first
	d.Profile.DesignPolicy = &policy
	return d
}

func earlyPlanIssues(plan unit.DesignPlan, definition m.Definition) []string {
	var found []string
	for _, issue := range unit.PlanFeasibilityIssues(plan, definition) {
		if strings.Contains(issue.Message, unit.EarlyBenefitsRule) {
			found = append(found, issue.Path+": "+issue.Message)
		}
	}
	return found
}

// Under distinctEarlyBenefits two paths' first two purchases may not
// promise the same multiset of improvement dimensions and unlocks, in any
// order; repeats count and lowers do not (#32, SOL-32-01, SOL-32-02).
func TestEarlyBenefitsPlanCheck(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan := *stages.Draft.Run.DesignPlan
	base := *plan.UpgradeIntents
	on := true
	def := unit.DefaultAuthoringDefinition()
	slowed := intent("none", "damage")
	slowed.Lowers = []string{"attack-rate"}
	renamed := func(i unit.UpgradeIntent) unit.UpgradeIntent {
		i.Technique = "Other Technique"
		return i
	}
	cases := []struct {
		name           string
		a1, a2, b1, b2 unit.UpgradeIntent
		definition     m.Definition
		want           string
	}{
		{"fixture", base.Path1.Tier1, base.Path1.Tier2, base.Path3.Tier1, base.Path3.Tier2, def, ""},
		{"same", intent("none", "range"), intent("none", "damage"), renamed(intent("none", "range")), renamed(intent("none", "damage")), def,
			"upgradeIntents.path3.tier2: x-x-1 and x-x-2 together promise damage and range, the same as 1-x-x and 2-x-x. " + unit.EarlyBenefitsRule + " Redesign x-x-2 or x-x-1"},
		{"reversed", intent("none", "range"), intent("none", "damage"), intent("none", "damage"), intent("none", "range"), def,
			"upgradeIntents.path3.tier2: x-x-1 and x-x-2 together promise damage and range, the same as 1-x-x and 2-x-x."},
		{"split across tiers", intent("none", "range", "damage"), intent("none", "damage"), intent("none", "damage"), intent("none", "range", "damage"), def,
			"upgradeIntents.path3.tier2: x-x-1 and x-x-2 together promise damage twice and range, the same as 1-x-x and 2-x-x."},
		{"interval-only twins (#27)", base.Path1.Tier1, base.Path1.Tier2, intent("none", "attack-rate"), intent("none", "attack-rate"), def,
			"upgradeIntents.path3.tier2: x-x-1 and x-x-2 together promise attack-rate twice, the same as x-1-x and x-2-x."},
		{"repeat kept", intent("none", "damage"), intent("none", "damage"), intent("none", "damage"), intent("none", "range", "damage"), def, ""},
		{"added improvement", intent("none", "range"), intent("none", "damage"), intent("none", "range"), intent("none", "damage", "pierce"), def, ""},
		{"distinct unlock", intent("none", "range"), intent("none", "damage"), intent("none", "range"), intent("camo", "damage"), def, ""},
		{"same unlock", intent("camo", "range"), intent("none", "damage"), intent("none", "damage"), intent("camo", "range"), def,
			"upgradeIntents.path3.tier2: x-x-1 and x-x-2 together promise damage, range and unlock camo, the same as 1-x-x and 2-x-x."},
		{"lowers only", intent("none", "range"), intent("none", "damage"), intent("none", "range"), slowed, def,
			"upgradeIntents.path3.tier2: x-x-1 and x-x-2 together promise damage and range, the same as 1-x-x and 2-x-x."},
		{"policy off", intent("none", "range"), intent("none", "damage"), intent("none", "damage"), intent("none", "range"), withEarlyPolicy(def, nil, true), ""},
		{"explicit false", intent("none", "range"), intent("none", "damage"), intent("none", "range"), intent("none", "damage"), withEarlyPolicy(def, new(bool), true), ""},
		{"only the new field", intent("none", "damage"), intent("none", "damage"), intent("none", "damage"), intent("none", "damage"), withEarlyPolicy(def, &on, false),
			"upgradeIntents.path3.tier2: x-x-1 and x-x-2 together promise damage twice, the same as 1-x-x and 2-x-x."},
	}
	for _, c := range cases {
		intents := base
		intents.Path1.Tier1, intents.Path1.Tier2 = c.a1, c.a2
		intents.Path3.Tier1, intents.Path3.Tier2 = c.b1, c.b2
		p := plan
		p.UpgradeIntents = &intents
		found := earlyPlanIssues(p, c.definition)
		switch {
		case c.want == "" && len(found) != 0:
			t.Errorf("%s: rejected: %v", c.name, found)
		case c.want != "" && (len(found) != 1 || !strings.HasPrefix(found[0], c.want)):
			t.Errorf("%s: issues %v, want %q", c.name, found, c.want)
		}
	}
}

// distinctFirstUpgrades keeps its meaning from before 5a1883b: the resolved
// 1-0-0, 0-1-0 and 0-0-1 builds differ. Alone it neither states nor checks
// the early benefits rule (#32).
func TestDistinctFirstUpgradesAloneKeepsItsMeaning(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prepared := stages.Draft.Prepared
	definition := withEarlyPolicy(*prepared.Request.MechanicsDefinition, nil, true)
	prepared.Request.MechanicsDefinition = &definition
	plan := *stages.Draft.Run.DesignPlan
	intents := *plan.UpgradeIntents
	intents.Path1.Tier1, intents.Path1.Tier2 = intent("none", "damage"), intent("none", "damage")
	intents.Path3.Tier1, intents.Path3.Tier2 = intent("none", "damage"), intent("none", "damage")
	plan.UpgradeIntents = &intents
	if found := earlyPlanIssues(plan, definition); len(found) != 0 {
		t.Errorf("distinctFirstUpgrades rejected matching early benefits: %v", found)
	}
	request, err := unit.DesignPlanRequest(prepared)
	if err != nil || strings.Contains(request.Prompt, unit.EarlyBenefitsRule) {
		t.Errorf("the plan prompt states the early benefits rule without its field: %v", err)
	}
	guidance := strings.Join(unit.DesignGuidance(&prepared.Request), "\n")
	if strings.Contains(guidance, unit.EarlyBenefitsRule) || !strings.Contains(guidance, "No two first purchases (1-x-x, x-1-x, x-x-1) may produce the same resolved attack.") {
		t.Errorf("draft guidance under distinctFirstUpgrades alone: %s", guidance)
	}
	blueprint := *stages.Draft.Candidate.Blueprint
	blueprint.Paths.Path3.Tiers.Tier1.Changes = blueprint.Paths.Path1.Tiers.Tier1.Changes
	if got := unit.EarlyBenefitsIssues(blueprint, &intents, definition); len(got) != 0 {
		t.Errorf("resolved early benefits checked without the field: %v", got)
	}
	duplicate := false
	for _, issue := range unit.ValidateBlueprintRequest(blueprint, prepared.Request) {
		if issue.Path == "paths.path3.tiers.tier1" && strings.HasPrefix(issue.Message, "Resolved tier 1 behavior duplicates path1") {
			duplicate = true
		}
	}
	if !duplicate {
		t.Error("identical resolved first purchases were not rejected under distinctFirstUpgrades")
	}
}

// An early or late milestone that promises no improvement and no unlock is
// rejected by the common plan authoring floor, planAuthoringFloor, under
// every policy (SOL-32-03).
func TestEmptyMilestoneIsRejectedWhateverThePolicy(t *testing.T) {
	on := true
	for _, tier := range []string{"tier1", "tier2", "tier4"} {
		for _, benefits := range []*bool{nil, new(bool), &on} {
			prepared, err := fixture.Prepare()
			if err != nil {
				t.Fatal(err)
			}
			d := withEarlyPolicy(*prepared.Request.MechanicsDefinition, benefits, false)
			prepared.Request.MechanicsDefinition = &d
			output := recordedOutput(t, "plan")
			at(output, "paths", "path2", "milestones", tier).(*s.Object).Set("improves", []any{}).Set("unlock", "none")
			_, err = unit.DecodeDesignPlan(output, &prepared.Request)
			want := "upgradeIntents.path2." + tier + ": Declare at least one supported improvement or unlock for this milestone."
			if err == nil || !strings.Contains(err.Error(), want) {
				t.Errorf("%s, distinctEarlyBenefits %v: %v", tier, benefits, err)
			}
		}
	}
}

func statChange(stat, operation string, value float64) m.Change {
	return m.Change{Kind: "stat", Target: "base", Stat: stat, Operation: operation, Number: value}
}

// The resolved form compares what the pure builds actually gain, whatever
// the amounts.
func TestResolvedEarlyBenefits(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	definition := unit.DefaultAuthoringDefinition()
	resolved := func(p1, p3 [2][]m.Change, d m.Definition) []m.Issue {
		b := *stages.Draft.Candidate.Blueprint
		b.Paths.Path1.Tiers.Tier1.Changes, b.Paths.Path1.Tiers.Tier2.Changes = p1[0], p1[1]
		b.Paths.Path3.Tiers.Tier1.Changes, b.Paths.Path3.Tiers.Tier2.Changes = p3[0], p3[1]
		return unit.EarlyBenefitsIssues(b, nil, d)
	}
	rangeThenDamage := [2][]m.Change{{statChange("range", "add", 2)}, {statChange("damage", "add", 1)}}
	camo := m.Change{Kind: "detection", Target: "base", Trait: "camo", Bool: true}
	cases := []struct {
		name   string
		p3     [2][]m.Change
		d      m.Definition
		reject bool
	}{
		{"same benefits, different amounts", [2][]m.Change{{statChange("range", "add", 8)}, {statChange("damage", "add", 3)}}, definition, true},
		{"reversed", [2][]m.Change{{statChange("damage", "add", 2)}, {statChange("range", "add", 5)}}, definition, true},
		{"a slower interval beside the same gains", [2][]m.Change{{statChange("range", "add", 2)}, {statChange("damage", "add", 1), statChange("intervalSeconds", "multiply", 1.2)}}, definition, true},
		{"an added improvement", [2][]m.Change{{statChange("range", "add", 2)}, {statChange("damage", "add", 1), statChange("pierce", "add", 1)}}, definition, false},
		{"a distinct unlock", [2][]m.Change{{statChange("range", "add", 2)}, {statChange("damage", "add", 1), camo}}, definition, false},
		{"policy off", rangeThenDamage, withEarlyPolicy(definition, nil, true), false},
	}
	for _, c := range cases {
		got := resolved(rangeThenDamage, c.p3, c.d)
		switch {
		case !c.reject && len(got) != 0:
			t.Errorf("%s: rejected: %v", c.name, got)
		case c.reject && (len(got) != 1 || got[0].Path != "paths.path3.tiers.tier2" ||
			!strings.HasPrefix(got[0].Message, "Resolved x-x-1 and x-x-2 together improve damage and range, the same as resolved 1-x-x and 2-x-x: 1-x-x gives range, 2-x-x gives damage, ") ||
			!strings.Contains(got[0].Message, unit.EarlyBenefitsRule+" Change what x-x-2 or x-x-1 improves or unlocks so the two paths' early benefits differ.")):
			t.Errorf("%s: issues %v", c.name, got)
		}
	}
}

// earlyMatchMechanics keeps the fixture plan, whose early benefits are
// distinct, and adds changes it does not promise: a faster interval to
// 1-x-x and 2-x-x, and pierce to x-1-x and x-2-x. Both paths then resolve
// to attack-rate twice and pierce twice.
func earlyMatchMechanics(t *testing.T) *s.Object {
	mech := recordedOutput(t, "mechanics")
	for _, tier := range []string{"tier1", "tier2"} {
		p1 := at(mech, "paths", "path1", "tiers", tier).(*s.Object)
		p2 := at(mech, "paths", "path2", "tiers", tier).(*s.Object)
		c1, _ := p1.Get("statChanges")
		c2, _ := p2.Get("statChanges")
		p1.Set("statChanges", append(c1.([]any), s.NewObject().Set("stat", "intervalSeconds").Set("operation", "multiply").Set("value", 0.9)))
		p2.Set("statChanges", append(c2.([]any), s.NewObject().Set("stat", "pierce").Set("operation", "add").Set("value", 1.0)))
	}
	return mech
}

const resolvedEarlyFacts = "Resolved x-1-x and x-2-x together improve attack-rate twice and pierce twice, the same as resolved 1-x-x and 2-x-x: 1-x-x gives attack-rate and pierce, 2-x-x gives attack-rate and pierce, x-1-x gives attack-rate and pierce and x-2-x gives attack-rate and pierce."

const resolvedEarlyCause = "The retained plan promises attack-rate at x-2-x; the resolved purchase gives attack-rate and pierce. Its pierce, which the plan does not promise there, makes the two paths' early benefits match. Keep x-2-x's promises and make its early benefits distinct: remove that pierce from x-2-x, or add another improvement or unlock from this path's own technique."

// The resolved-stage catch. The issue names both paths and their resolved
// early benefits, states the rule and says which unplanned benefit caused
// the match, without the global "Implement only the planned promises"
// (SOL-33-04). It targets the departing purchase for TargetedTierRepair,
// and a draft that keeps the match publishes nothing. check reports it too.
func TestResolvedEarlyBenefitsAreCaughtAndRepaired(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	zero := 0
	scripted := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), earlyMatchMechanics(t)}}
	if _, err := unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{MaxRepairAttempts: &zero}); err == nil ||
		!strings.Contains(err.Error(), "together improve attack-rate twice and pierce twice") || !strings.Contains(err.Error(), "No invalid Unit was published.") {
		t.Errorf("published a resolved match: %v", err)
	}
	scripted = &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), earlyMatchMechanics(t)}}
	_, _ = unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{})
	want := "paths.path2.tiers.tier2: " + resolvedEarlyFacts + " " + unit.EarlyBenefitsRule + " " + resolvedEarlyCause
	if len(scripted.Requests) != 3 {
		t.Fatalf("%d requests", len(scripted.Requests))
	}
	repair := scripted.Requests[2].Prompt
	if !strings.Contains(repair, want) || !strings.Contains(repair, "Preserve the retained character plan while correcting these tiers.") || strings.Contains(repair, "Implement only the planned promises") {
		t.Errorf("repair: %s", retryReason(repair))
	}

	// check runs the same rule on a draft whose mechanics match, beside the
	// promise check and outside ValidateTyped.
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	draft := stages.Draft
	blueprint := *draft.Candidate.Blueprint
	for tier := 1; tier <= 2; tier++ {
		p1, p2 := blueprint.Paths.Path1.Tiers.At(tier), blueprint.Paths.Path2.Tiers.At(tier)
		p1.Changes = append(append([]m.Change{}, p1.Changes...), statChange("intervalSeconds", "multiply", 0.9))
		p2.Changes = append(append([]m.Change{}, p2.Changes...), statChange("pierce", "add", 1))
	}
	if issues := m.ValidateTyped(&blueprint, *draft.Prepared.Request.MechanicsDefinition); len(issues) != 0 {
		t.Fatalf("the rule reached ValidateTyped, which rendering calls: %v", issues)
	}
	candidate, err := unit.CompileBlueprint(blueprint, draft.Prepared.Request)
	if err != nil {
		t.Fatal(err)
	}
	draft.Candidate = candidate
	evaluation, err := unit.EvaluateUnitDesign(blueprint, draft.Run.DesignPlan, *draft.Prepared.Request.MechanicsDefinition)
	if err != nil {
		t.Fatal(err)
	}
	draft.Run.DesignEvaluation = evaluation
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	var failed []string
	for _, f := range checked.Findings {
		if f.Outcome == "fail" {
			failed = append(failed, f.Rule+" "+f.Subject+": "+f.Message)
		}
	}
	if len(failed) != 1 || failed[0] != "distinct-early-benefits "+want {
		t.Errorf("check findings: %v", failed)
	}
}

// The prompt, the scoped issues and the retry requests share one rule
// sentence (SOL-33-03): the plan prompt and draft guidance state it only
// under the field, the plan issue carries it into the full-plan correction,
// and the resolved issue carries it into the targeted tier repair.
func TestEarlyBenefitsPromptIssueAndRetryAgree(t *testing.T) {
	const rule = "Two paths' first two purchases together must differ in what they improve or unlock, whatever their order, names, prices or amounts; lowers do not count."
	if unit.EarlyBenefitsRule != rule {
		t.Fatalf("rule changed: %q", unit.EarlyBenefitsRule)
	}
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	planPrompt, err := unit.DesignPlanRequest(prepared)
	if err != nil || !strings.Contains(planPrompt.Prompt, rule) {
		t.Errorf("plan prompt lacks the rule: %v", err)
	}
	if guidance := strings.Join(unit.DesignGuidance(&prepared.Request), "\n"); !strings.Contains(guidance, rule) {
		t.Errorf("draft guidance lacks the rule: %s", guidance)
	}
	off := prepared
	offDefinition := withEarlyPolicy(*prepared.Request.MechanicsDefinition, nil, true)
	off.Request.MechanicsDefinition = &offDefinition
	if request, err := unit.DesignPlanRequest(off); err != nil || strings.Contains(request.Prompt, rule) {
		t.Errorf("plan prompt states the rule while the field is off: %v", err)
	}
	if guidance := strings.Join(unit.DesignGuidance(&off.Request), "\n"); strings.Contains(guidance, rule) {
		t.Errorf("draft guidance states the rule while the field is off")
	}

	// Plan stage: the scoped issue, then the full-plan correction.
	matching := recordedOutput(t, "plan")
	set := func(path, tier string, improves ...any) {
		at(matching, "paths", path, "milestones", tier).(*s.Object).Set("improves", improves).Set("unlock", "none")
	}
	set("path1", "tier1", "pierce")
	set("path1", "tier2", "range")
	set("path3", "tier1", "range")
	set("path3", "tier2", "pierce")
	facts := "x-x-1 and x-x-2 together promise pierce and range, the same as 1-x-x and 2-x-x. " + rule +
		" Redesign x-x-2 or x-x-1: add or swap an improvement or unlock from this path's own technique, such as damage, attack-rate or personal detection."
	if _, err := unit.DecodeDesignPlan(matching, &prepared.Request); err == nil || !strings.Contains(err.Error(), "upgradeIntents.path3.tier2: "+facts) {
		t.Errorf("plan issue: %v", err)
	}
	scripted := &fixture.Model{Outputs: []any{matching, recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	if _, err := unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{}); err != nil {
		t.Fatal(err)
	}
	if len(scripted.Requests) != 3 || !strings.Contains(scripted.Requests[1].Prompt, "Correct this invalid design plan") ||
		!strings.Contains(scripted.Requests[1].Prompt, "paths.path3.milestones.tier2: "+facts) {
		t.Errorf("plan correction: %s", retryReason(scripted.Requests[1].Prompt))
	}
	scripted = &fixture.Model{Outputs: []any{matching, matching}}
	if _, err := unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{}); err == nil || len(scripted.Requests) != 2 ||
		!strings.Contains(err.Error(), "The character design plan could not be validated. paths.path3.milestones.tier2: x-x-1 and x-x-2 together promise pierce and range") {
		t.Errorf("want failure after 2 plan calls, got %v after %d", err, len(scripted.Requests))
	}

	// Resolved stage: the scoped issue, then the targeted tier repair.
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	blueprint := *stages.Draft.Candidate.Blueprint
	for tier := 1; tier <= 2; tier++ {
		p1, p2 := blueprint.Paths.Path1.Tiers.At(tier), blueprint.Paths.Path2.Tiers.At(tier)
		p1.Changes = append(append([]m.Change{}, p1.Changes...), statChange("intervalSeconds", "multiply", 0.9))
		p2.Changes = append(append([]m.Change{}, p2.Changes...), statChange("pierce", "add", 1))
	}
	issues := unit.EarlyBenefitsIssues(blueprint, stages.Draft.Run.DesignPlan.UpgradeIntents, *prepared.Request.MechanicsDefinition)
	want := "paths.path2.tiers.tier2: " + resolvedEarlyFacts + " " + rule + " " + resolvedEarlyCause
	if len(issues) != 1 || issues[0].Path+": "+issues[0].Message != want {
		t.Fatalf("resolved issue: %v", issues)
	}
	scripted = &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), earlyMatchMechanics(t)}}
	_, _ = unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{})
	if len(scripted.Requests) != 3 || !strings.Contains(scripted.Requests[2].Prompt, `"violations":["`+want+`"]`) || !strings.Contains(scripted.Requests[2].Prompt, rule) {
		t.Errorf("targeted repair: %s", retryReason(scripted.Requests[len(scripted.Requests)-1].Prompt))
	}
}
