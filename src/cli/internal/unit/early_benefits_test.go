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

// withEarlyPolicy sets distinctEarlyBenefits and distinctFirstUpgrades and
// clears exclusiveEarlyBenefits, which the Default sets and which subsumes
// distinctEarlyBenefits (#61).
func withEarlyPolicy(d m.Definition, benefits *bool, first bool) m.Definition {
	policy := *d.Profile.DesignPolicy
	policy.DistinctEarlyBenefits, policy.DistinctFirstUpgrades, policy.ExclusiveEarlyBenefits = benefits, first, nil
	d.Profile.DesignPolicy = &policy
	return d
}

// multisetPolicy selects distinctEarlyBenefits alone, the early benefits
// rule of the Default before v29.
func multisetPolicy(p *m.DesignPolicy) {
	on := true
	p.DistinctEarlyBenefits, p.ExclusiveEarlyBenefits = &on, nil
}

// multisetDraft is the fixture drafted under multisetPolicy.
func multisetDraft(t *testing.T) unit.Draft {
	t.Helper()
	model := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), preparedUnder(t, multisetPolicy), model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	return draft
}

func earlyPlanIssues(plan unit.DesignPlan, definition m.Definition) []string {
	var found []string
	for _, issue := range unit.PlanFeasibilityIssues(plan, definition) {
		if strings.Contains(issue.Message, m.EarlyBenefitsRule) {
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
	def := withEarlyPolicy(unit.DefaultAuthoringDefinition(), &on, true)
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
			"upgradeIntents.path3.tier2: x-x-1 and x-x-2 together promise damage and range, the same as 1-x-x and 2-x-x. " + m.EarlyBenefitsRule + " Redesign x-x-2 or x-x-1"},
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
	if err != nil || strings.Contains(request.Prompt, m.EarlyBenefitsRule) {
		t.Errorf("the plan prompt states the early benefits rule without its field: %v", err)
	}
	guidance := strings.Join(unit.DesignGuidance(&prepared.Request), "\n")
	if strings.Contains(guidance, m.EarlyBenefitsRule) || !strings.Contains(guidance, "No two first purchases (1-x-x, x-1-x, x-x-1) may produce the same resolved attack.") {
		t.Errorf("draft guidance under distinctFirstUpgrades alone: %s", guidance)
	}
	blueprint := *stages.Draft.Candidate.Blueprint
	blueprint.Paths.Path3.Tiers.Tier1.Changes = blueprint.Paths.Path1.Tiers.Tier1.Changes
	if got := unit.EarlyBenefitsIssues(blueprint, &intents, definition); len(got) != 0 {
		t.Errorf("resolved early benefits checked without the field: %v", got)
	}
	duplicate := false
	for _, issue := range unit.ValidateBlueprintRequest(blueprint, prepared.Request) {
		if issue.Path == "paths.path3.tiers.tier1" && strings.HasPrefix(issue.Message, "Resolved x-x-1 behaves exactly like 1-x-x. "+m.DistinctFirstPurchasesRule) {
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
	on := true
	definition := withEarlyPolicy(unit.DefaultAuthoringDefinition(), &on, true)
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
			!strings.Contains(got[0].Message, m.EarlyBenefitsRule+" Change what x-x-2 or x-x-1 improves or unlocks so the two paths' early benefits differ.")):
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
	prepared := preparedUnder(t, multisetPolicy)
	zero := 0
	scripted := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), earlyMatchMechanics(t)}}
	if _, err := unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{MaxRepairAttempts: &zero}); err == nil ||
		!strings.Contains(err.Error(), "together improve attack-rate twice and pierce twice") || !strings.Contains(err.Error(), "No invalid Unit was published.") {
		t.Errorf("published a resolved match: %v", err)
	}
	scripted = &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), earlyMatchMechanics(t)}}
	_, _ = unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{})
	want := "paths.path2.tiers.tier2: " + resolvedEarlyFacts + " " + m.EarlyBenefitsRule + " " + resolvedEarlyCause
	if len(scripted.Requests) != 3 {
		t.Fatalf("%d requests", len(scripted.Requests))
	}
	repair := scripted.Requests[2].Prompt
	if !strings.Contains(repair, want) || !strings.Contains(repair, "Preserve the retained character plan while correcting these tiers.") || strings.Contains(repair, "Implement only the planned promises") {
		t.Errorf("repair: %s", retryReason(repair))
	}

	// check runs the same rule on a draft whose mechanics match, beside the
	// promise check and outside ValidateTyped.
	failed := earlyMatchCheckFailures(t, true)
	if len(failed) != 1 || failed[0].rule != "distinct-early-benefits "+want ||
		!strings.Contains(failed[0].action, "while keeping the retained plan true") {
		t.Errorf("check findings: %v", failed)
	}
}

// Without a retained plan the check reports the plan-free cause and an
// Action that names no plan (SOL-PR43-01).
func TestEarlyBenefitsCheckWithoutPlan(t *testing.T) {
	failed := earlyMatchCheckFailures(t, false)
	want := "distinct-early-benefits paths.path2.tiers.tier2: " + resolvedEarlyFacts + " " + m.EarlyBenefitsRule +
		" Change what x-2-x or x-1-x improves or unlocks so the two paths' early benefits differ."
	if len(failed) != 1 || failed[0].rule != want ||
		failed[0].action != "Make the two paths' resolved early benefits distinct and compile again." {
		t.Errorf("check findings: %v", failed)
	}
}

type checkFailure struct{ rule, action string }

// earlyMatchCheckFailures checks the recorded draft after giving path1 and
// path2 the same two early benefits, with or without its retained plan.
func earlyMatchCheckFailures(t *testing.T, withPlan bool) []checkFailure {
	t.Helper()
	draft := multisetDraft(t)
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
	if !withPlan {
		draft.Run.DesignPlan = nil
	}
	evaluation, err := unit.EvaluateUnitDesign(blueprint, draft.Run.DesignPlan, *draft.Prepared.Request.MechanicsDefinition)
	if err != nil {
		t.Fatal(err)
	}
	draft.Run.DesignEvaluation = evaluation
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	var failed []checkFailure
	for _, f := range checked.Findings {
		if f.Outcome == "fail" {
			action := ""
			if f.Action != nil {
				action = *f.Action
			}
			failed = append(failed, checkFailure{f.Rule + " " + f.Subject + ": " + f.Message, action})
		}
	}
	return failed
}

// The prompt, the scoped issues and the retry requests share one rule
// sentence (SOL-33-03): the plan prompt and draft guidance state it only
// under the field, the plan issue carries it into the full-plan correction,
// and the resolved issue carries it into the targeted tier repair.
func TestEarlyBenefitsPromptIssueAndRetryAgree(t *testing.T) {
	const rule = "Two paths' first two purchases together must differ in what they improve or unlock, whatever their order, names, prices or amounts; lowers do not count."
	if m.EarlyBenefitsRule != rule {
		t.Fatalf("rule changed: %q", m.EarlyBenefitsRule)
	}
	prepared := preparedUnder(t, multisetPolicy)
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

// A follow-up promise means the follow-up's own fields improve: count,
// damage multiplier, radius or inheritance. A base damage raise does not
// keep it, so the plan and resolved checks agree and the resolved repair
// targets the purchase that skipped its follow-up (OPUS-CODE-43-1,
// SOL-PR43-02).
func TestDamageRaiseDoesNotImproveTheFollowUp(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	on := true
	definition := withEarlyPolicy(unit.DefaultAuthoringDefinition(), &on, true)
	b := *stages.Draft.Candidate.Blueprint
	b.BaseAttack.FollowUp = &m.FollowUp{Name: "Shockwave", Count: 2, DamageMultiplier: 0.5, Radius: 20}
	early := [2][]m.Change{{statChange("range", "add", 2)}, {statChange("damage", "add", 1)}}
	b.Paths.Path1.Tiers.Tier1.Changes, b.Paths.Path1.Tiers.Tier2.Changes = early[0], early[1]
	b.Paths.Path3.Tiers.Tier1.Changes, b.Paths.Path3.Tiers.Tier2.Changes = early[0], early[1]
	intents := *stages.Draft.Run.DesignPlan.UpgradeIntents
	intents.Path1.Tier1, intents.Path1.Tier2 = intent("none", "range"), intent("none", "damage")
	intents.Path3.Tier1, intents.Path3.Tier2 = intent("none", "range"), intent("none", "damage", "follow-up")

	var early3 []string
	for _, issue := range unit.PlanIntentIssues(b, &intents, definition) {
		if strings.HasPrefix(issue.Path, "paths.path1.tiers.tier1") || strings.HasPrefix(issue.Path, "paths.path1.tiers.tier2") ||
			strings.HasPrefix(issue.Path, "paths.path3.tiers.tier1") || strings.HasPrefix(issue.Path, "paths.path3.tiers.tier2") {
			early3 = append(early3, issue.Path+": "+issue.Message)
		}
	}
	if len(early3) != 1 || !strings.HasPrefix(early3[0], "paths.path3.tiers.tier2.planIntent: The retained plan promises improved follow-up,") {
		t.Errorf("plan intent issues: %v", early3)
	}

	got := unit.EarlyBenefitsIssues(b, &intents, definition)
	if len(got) != 1 || got[0].Path != "paths.path3.tiers.tier2" ||
		!strings.HasSuffix(got[0].Message, "The retained plan promises damage and follow-up at x-x-2; the resolved purchase gives damage. Deliver x-x-2's promised follow-up so the two paths' early benefits differ as the plan does.") {
		t.Errorf("resolved issues: %v", got)
	}

	// Raising the follow-up's own multiplier keeps the promise and makes
	// the two paths distinct.
	stronger := *b.BaseAttack.FollowUp
	stronger.DamageMultiplier = 0.75
	b.Paths.Path3.Tiers.Tier2.Changes = []m.Change{statChange("damage", "add", 1), {Kind: "followUp", Target: "base", FollowUp: &stronger}}
	if got := unit.EarlyBenefitsIssues(b, &intents, definition); len(got) != 0 {
		t.Errorf("a kept follow-up promise: %v", got)
	}
}

// ---- exclusiveEarlyBenefits ----

func exclusivePlanIssues(plan unit.DesignPlan, definition m.Definition) []string {
	var found []string
	for _, issue := range unit.PlanFeasibilityIssues(plan, definition) {
		if strings.Contains(issue.Message, m.ExclusiveEarlyBenefitsRule) {
			found = append(found, issue.Path+": "+issue.Message)
		}
	}
	return found
}

// Under exclusiveEarlyBenefits, on in the Default, nothing one path's first
// two purchases improve or unlock may be improved or unlocked by another
// path's first two purchases (#61). A path may repeat its own dimension, as
// the fixture's middle path attacks faster twice, and later purchases may
// add another path's dimension, as x-x-3 raises damage and pierce.
func TestExclusiveEarlyBenefitsPlanCheck(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan := *stages.Draft.Run.DesignPlan
	base := *plan.UpgradeIntents
	def := unit.DefaultAuthoringDefinition()
	on := true
	lowersPierce := intent("none", "attack-rate")
	lowersPierce.Lowers = []string{"pierce"}
	cases := []struct {
		name       string
		edit       func(*unit.UpgradeIntents)
		definition m.Definition
		want       []string
	}{
		{"fixture: each path double-dips its own dimension", func(*unit.UpgradeIntents) {}, def, nil},
		{"two first purchases add damage", func(i *unit.UpgradeIntents) {
			i.Path1.Tier1, i.Path2.Tier1 = intent("none", "damage"), intent("none", "damage")
		}, def, []string{"upgradeIntents.path2.tier1: x-1-x and x-2-x promise damage, which 1-x-x and 2-x-x also promise. " + m.ExclusiveEarlyBenefitsRule +
			" Redesign x-1-x or x-2-x: move damage to x-3-x or later, or replace it with an improvement or unlock from this path's own technique that no other path's first two purchases promise."}},
		{"shared across tiers", func(i *unit.UpgradeIntents) { i.Path1.Tier2 = intent("none", "range") }, def,
			[]string{"upgradeIntents.path3.tier2: x-x-1 and x-x-2 promise range, which 1-x-x and 2-x-x also promise."}},
		{"the same unlock", func(i *unit.UpgradeIntents) { i.Path1.Tier2 = intent("camo", "pierce") }, def,
			[]string{"upgradeIntents.path3.tier2: x-x-1 and x-x-2 promise unlock camo, which 1-x-x and 2-x-x also promise."}},
		{"a later purchase may add it", func(i *unit.UpgradeIntents) { i.Path1.Tier1 = intent("none", "damage") }, def, nil},
		{"lowers do not count", func(i *unit.UpgradeIntents) { i.Path2.Tier1 = lowersPierce }, def, nil},
		{"distinct multisets still share", func(i *unit.UpgradeIntents) {
			i.Path1.Tier1, i.Path1.Tier2 = intent("none", "range"), intent("none", "damage")
			i.Path3.Tier1, i.Path3.Tier2 = intent("none", "range"), intent("none", "damage", "pierce")
		}, def, []string{"upgradeIntents.path3.tier2: x-x-1 and x-x-2 promise damage and range, which 1-x-x and 2-x-x also promise."}},
		{"both fields: the exclusive rule only", func(i *unit.UpgradeIntents) { i.Path1.Tier1 = intent("none", "range") }, func() m.Definition {
			d := def
			policy := *d.Profile.DesignPolicy
			policy.DistinctEarlyBenefits = &on
			d.Profile.DesignPolicy = &policy
			return d
		}(), []string{"upgradeIntents.path3.tier2: x-x-1 and x-x-2 promise range, which 1-x-x and 2-x-x also promise."}},
		{"policy off", func(i *unit.UpgradeIntents) { i.Path1.Tier1 = intent("none", "range") }, withEarlyPolicy(def, nil, true), nil},
	}
	for _, c := range cases {
		intents := base
		c.edit(&intents)
		p := plan
		p.UpgradeIntents = &intents
		found := exclusivePlanIssues(p, c.definition)
		if len(found) != len(c.want) {
			t.Errorf("%s: issues %v, want %v", c.name, found, c.want)
			continue
		}
		for i, want := range c.want {
			if !strings.HasPrefix(found[i], want) {
				t.Errorf("%s: issue %q, want %q", c.name, found[i], want)
			}
		}
		if multiset := earlyPlanIssues(p, c.definition); len(multiset) != 0 {
			t.Errorf("%s: the subsumed multiset rule ran: %v", c.name, multiset)
		}
	}
}

// The resolved form compares what the pure builds up to each second
// purchase gain, whatever the amounts.
func TestExclusiveEarlyBenefitsResolvedCheck(t *testing.T) {
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
	pierceTwice := [2][]m.Change{{statChange("pierce", "add", 1)}, {statChange("pierce", "add", 2)}}
	camo := m.Change{Kind: "detection", Target: "base", Trait: "camo", Bool: true}
	cases := []struct {
		name string
		p3   [2][]m.Change
		d    m.Definition
		want string
	}{
		{"each path double-dips its own dimension", [2][]m.Change{{statChange("range", "add", 8)}, {statChange("range", "add", 8)}}, definition, ""},
		{"a distinct unlock beside its own dimension", [2][]m.Change{{statChange("range", "add", 8)}, {statChange("range", "add", 8), camo}}, definition, ""},
		{"a shared early dimension", [2][]m.Change{{statChange("range", "add", 8)}, {statChange("pierce", "add", 1)}}, definition,
			"paths.path3.tiers.tier2: Resolved x-x-1 and x-x-2 give pierce, which resolved 1-x-x and 2-x-x also give: 1-x-x gives pierce, 2-x-x gives pierce, x-x-1 gives range and x-x-2 gives pierce. " +
				m.ExclusiveEarlyBenefitsRule + " Change what x-x-2 or x-x-1 improves or unlocks so it shares nothing with 1-x-x and 2-x-x."},
		{"a slower interval does not share attack-rate", [2][]m.Change{{statChange("range", "add", 8)}, {statChange("range", "add", 8), statChange("intervalSeconds", "multiply", 1.2)}}, definition, ""},
		{"policy off", [2][]m.Change{{statChange("pierce", "add", 1)}, {statChange("pierce", "add", 1)}}, withEarlyPolicy(definition, nil, true), ""},
	}
	for _, c := range cases {
		got := resolved(pierceTwice, c.p3, c.d)
		switch {
		case c.want == "" && len(got) != 0:
			t.Errorf("%s: rejected: %v", c.name, got)
		case c.want != "" && (len(got) != 1 || got[0].Path+": "+got[0].Message != c.want):
			t.Errorf("%s: issues %v, want %q", c.name, got, c.want)
		}
	}
}

// earlyShareMechanics keeps the fixture plan and adds a pierce raise, which
// that plan does not promise, to x-2-x: the middle path's early purchases
// then share pierce with the top path's.
func earlyShareMechanics(t *testing.T) *s.Object {
	mech := recordedOutput(t, "mechanics")
	tier := at(mech, "paths", "path2", "tiers", "tier2").(*s.Object)
	changes, _ := tier.Get("statChanges")
	tier.Set("statChanges", append(changes.([]any), s.NewObject().Set("stat", "pierce").Set("operation", "add").Set("value", 1.0)))
	return mech
}

const exclusiveEarlyFacts = "Resolved x-1-x and x-2-x give pierce, which resolved 1-x-x and 2-x-x also give: 1-x-x gives pierce, 2-x-x gives pierce, x-1-x gives attack-rate and x-2-x gives attack-rate and pierce."

const exclusiveEarlyCause = "The retained plan does not promise pierce at x-2-x; the resolved purchase gives it, and the other path's first two purchases give it too. Remove that pierce from x-2-x and keep x-2-x's promises."

// The plan prompt, the draft guidance, the plan issue with its full-plan
// correction and the resolved issue with its targeted repair share one rule
// sentence, and check reports the resolved issue.
func TestExclusiveEarlyBenefitsPromptIssueAndRetryAgree(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	rule := m.ExclusiveEarlyBenefitsRule
	planPrompt, err := unit.DesignPlanRequest(prepared)
	if err != nil || !strings.Contains(planPrompt.Prompt, rule) || strings.Contains(planPrompt.Prompt, m.EarlyBenefitsRule) {
		t.Errorf("the plan prompt does not state exactly the exclusive rule: %v", err)
	}
	if guidance := strings.Join(unit.DesignGuidance(&prepared.Request), "\n"); !strings.Contains(guidance, rule) {
		t.Errorf("draft guidance lacks the rule: %s", guidance)
	}

	// Plan stage.
	sharing := recordedOutput(t, "plan")
	at(sharing, "paths", "path2", "milestones", "tier1").(*s.Object).Set("improves", []any{"pierce"})
	scripted := &fixture.Model{Outputs: []any{sharing, recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	if _, err := unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{}); err != nil {
		t.Fatal(err)
	}
	if len(scripted.Requests) != 3 || !strings.Contains(scripted.Requests[1].Prompt, "paths.path2.milestones.tier1: x-1-x and x-2-x promise pierce, which 1-x-x and 2-x-x also promise. "+rule) {
		t.Errorf("plan correction: %s", retryReason(scripted.Requests[1].Prompt))
	}

	// Resolved stage.
	zero := 0
	scripted = &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), earlyShareMechanics(t)}}
	if _, err := unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{MaxRepairAttempts: &zero}); err == nil || !strings.Contains(err.Error(), "No invalid Unit was published.") {
		t.Errorf("published a shared early benefit: %v", err)
	}
	want := "paths.path2.tiers.tier2: " + exclusiveEarlyFacts + " " + rule + " " + exclusiveEarlyCause
	scripted = &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), earlyShareMechanics(t)}}
	_, _ = unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{})
	if len(scripted.Requests) != 3 || !strings.Contains(scripted.Requests[2].Prompt, `"violations":["`+want+`"]`) {
		t.Errorf("targeted repair: %s", retryReason(scripted.Requests[len(scripted.Requests)-1].Prompt))
	}

	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	draft := stages.Draft
	blueprint := *draft.Candidate.Blueprint
	tier := &blueprint.Paths.Path2.Tiers.Tier2
	tier.Changes = append(append([]m.Change{}, tier.Changes...), statChange("pierce", "add", 1))
	if issues := m.ValidateTyped(&blueprint, *draft.Prepared.Request.MechanicsDefinition); len(issues) != 0 {
		t.Fatalf("the rule reached ValidateTyped, which rendering calls: %v", issues)
	}
	if draft.Candidate, err = unit.CompileBlueprint(blueprint, draft.Prepared.Request); err != nil {
		t.Fatal(err)
	}
	if draft.Run.DesignEvaluation, err = unit.EvaluateUnitDesign(blueprint, draft.Run.DesignPlan, *draft.Prepared.Request.MechanicsDefinition); err != nil {
		t.Fatal(err)
	}
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
	if len(failed) != 1 || failed[0] != "exclusive-early-benefits "+want {
		t.Errorf("check findings: %v", failed)
	}
}
