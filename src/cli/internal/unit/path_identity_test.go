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

var rebound = m.ProposedMechanic{Name: "Rebound", Effect: "A ball that hits an obstacle rebounds off it and flies on, hitting enemies it has not hit yet with its remaining pierce.", SourceIDs: []string{"source1:6"}}

func identityPlanIssues(plan unit.DesignPlan, definition m.Definition) []string {
	var found []string
	for _, issue := range unit.PlanFeasibilityIssues(plan, definition) {
		if strings.Contains(issue.Message, m.PathIdentityRule) {
			found = append(found, issue.Path+": "+issue.Message)
		}
	}
	return found
}

func withIdentityPolicy(d m.Definition, on *bool) m.Definition {
	policy := *d.Profile.DesignPolicy
	policy.RequireTier3PathIdentity = on
	d.Profile.DesignPolicy = &policy
	return d
}

// Under requireTier3PathIdentity each third purchase promises an
// improvement, unlock or proposed mechanic that no purchase of the other two
// paths promises (#61). Dart Monkey's x-3-x is the only purchase that adds
// projectiles; its 3-x-x has the rebound it proposes, and its x-x-3 the range
// no other path raises.
func TestPathIdentityPlanCheck(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan := *stages.Draft.Run.DesignPlan
	base := *plan.UpgradeIntents
	def := unit.DefaultAuthoringDefinition()
	with := func(i unit.UpgradeIntent, proposed ...m.ProposedMechanic) unit.UpgradeIntent {
		i.ProposedMechanics = proposed
		return i
	}
	burn := intent("none", "burn")
	arc := m.ProposedMechanic{Name: "Arc", Effect: "Each throw spreads its darts over a 30 degree arc, so enemies beside the target can be hit.", SourceIDs: []string{"source1:11"}}
	cases := []struct {
		name       string
		edit       func(*unit.UpgradeIntents)
		definition m.Definition
		want       string
	}{
		{"fixture", func(*unit.UpgradeIntents) {}, def, ""},
		{"only another path's dimension", func(i *unit.UpgradeIntents) { i.Path2.Tier3 = intent("none", "attack-rate") }, def,
			"upgradeIntents.path2.tier3: x-3-x promises attack-rate, and a purchase of another path promises it too: attack-rate at 4-x-x. " + m.PathIdentityRule + " Give x-3-x an improvement, unlock or proposed mechanic from this path's own technique that no purchase of the other two paths promises, such as projectiles, splash or a proposed mechanic the sources describe."},
		{"stats without the proposed mechanic", func(i *unit.UpgradeIntents) { i.Path1.Tier3 = with(i.Path1.Tier3) }, def,
			"upgradeIntents.path1.tier3: 3-x-x promises damage and pierce, and a purchase of another path promises each too: damage at x-x-3 and pierce at x-x-3. " + m.PathIdentityRule},
		{"a proposed mechanic is enough", func(i *unit.UpgradeIntents) { i.Path2.Tier3 = with(intent("none", "attack-rate"), arc) }, def, ""},
		{"the same proposal on another path", func(i *unit.UpgradeIntents) {
			renamed := rebound
			renamed.Name = " REBOUND "
			i.Path3.Tier4 = with(i.Path3.Tier4, renamed)
		}, def, "upgradeIntents.path1.tier3: 3-x-x promises damage, pierce and proposed mechanic Rebound, and a purchase of another path promises each too: damage at x-x-3, pierce at x-x-3 and proposed mechanic Rebound at x-x-4."},
		{"an unlock is the dimension it opens", func(i *unit.UpgradeIntents) {
			i.Path2.Tier3 = intent("burn", "attack-rate")
			i.Path3.Tier5 = burn
		}, def, "upgradeIntents.path2.tier3: x-3-x promises attack-rate and burn, and a purchase of another path promises each too: attack-rate at 4-x-x and burn at x-x-5."},
		{"a targeting change does not count", func(i *unit.UpgradeIntents) { i.Path2.Tier3 = intent("targeting-change", "attack-rate") }, def,
			"upgradeIntents.path2.tier3: x-3-x promises attack-rate, and a purchase of another path promises it too"},
		{"a later purchase of its own path does not take it", func(i *unit.UpgradeIntents) { i.Path2.Tier5 = intent("none", "projectiles") }, def, ""},
		{"policy off", func(i *unit.UpgradeIntents) { i.Path2.Tier3 = intent("none", "attack-rate") }, withIdentityPolicy(def, nil), ""},
	}
	for _, c := range cases {
		intents := base
		c.edit(&intents)
		p := plan
		p.UpgradeIntents = &intents
		found := identityPlanIssues(p, c.definition)
		switch {
		case c.want == "" && len(found) != 0:
			t.Errorf("%s: rejected: %v", c.name, found)
		case c.want != "" && (len(found) != 1 || !strings.HasPrefix(found[0], c.want)):
			t.Errorf("%s: issues %v, want %q", c.name, found, c.want)
		}
	}
}

// The resolved form compares what the pure builds actually gain and the
// proposed mechanics each purchase carries. With a retained plan it targets
// the purchase that departs from it.
func TestPathIdentityResolvedCheck(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	definition := *stages.Draft.Prepared.Request.MechanicsDefinition
	intents := stages.Draft.Run.DesignPlan.UpgradeIntents
	fixtureBlueprint := *stages.Draft.Candidate.Blueprint
	if got := unit.PathIdentityIssues(fixtureBlueprint, intents, definition); len(got) != 0 {
		t.Fatalf("the fixture fails: %v", got)
	}
	if got := unit.PathIdentityIssues(fixtureBlueprint, nil, definition); len(got) != 0 {
		t.Fatalf("the fixture fails without its plan: %v", got)
	}
	projectiles := statChange("projectiles", "add", 1)
	cases := []struct {
		name    string
		edit    func(*m.Blueprint)
		intents *unit.UpgradeIntents
		d       m.Definition
		want    string
	}{
		{"another path gives it without the plan", func(b *m.Blueprint) {
			b.Paths.Path3.Tiers.Tier4.Changes = append(append([]m.Change{}, b.Paths.Path3.Tiers.Tier4.Changes...), projectiles)
		}, intents, definition,
			"paths.path3.tiers.tier4: Resolved x-3-x gives attack-rate and projectiles, and a purchase of another path gives each too: attack-rate at 4-x-x and projectiles at x-x-4. " + m.PathIdentityRule +
				" The retained plan promises projectiles at x-3-x and at no purchase of the other paths, but resolved x-x-4 also gives projectiles, which the plan does not promise there. Remove that projectiles from x-x-4 and keep x-x-4's promises, so x-3-x keeps the path's identity."},
		{"without a plan", func(b *m.Blueprint) {
			b.Paths.Path3.Tiers.Tier4.Changes = append(append([]m.Change{}, b.Paths.Path3.Tiers.Tier4.Changes...), projectiles)
		}, nil, definition,
			"paths.path2.tiers.tier3: Resolved x-3-x gives attack-rate and projectiles, and a purchase of another path gives each too: attack-rate at 4-x-x and projectiles at x-x-4. " + m.PathIdentityRule +
				" Give x-3-x an improvement, unlock or proposed mechanic from this path's own technique that no purchase of the other two paths gives."},
		{"the third purchase drops its planned identity", func(b *m.Blueprint) {
			b.Paths.Path2.Tiers.Tier3.Changes = []m.Change{statChange("intervalSeconds", "multiply", 0.75)}
		}, intents, definition,
			"paths.path2.tiers.tier3: Resolved x-3-x gives attack-rate, and a purchase of another path gives it too: attack-rate at 4-x-x. " + m.PathIdentityRule +
				" The retained plan promises projectiles at x-3-x and at no purchase of the other paths; the resolved purchase does not give it. Deliver x-3-x's promised projectiles so the path keeps its identity."},
		{"only larger numbers", func(b *m.Blueprint) { b.Paths.Path1.Tiers.Tier3.ProposedMechanics = nil }, nil, definition,
			"paths.path1.tiers.tier3: Resolved 3-x-x gives damage and pierce, and a purchase of another path gives each too: damage at x-x-3 and pierce at x-x-3. " + m.PathIdentityRule},
		{"policy off", func(b *m.Blueprint) { b.Paths.Path1.Tiers.Tier3.ProposedMechanics = nil }, nil, withIdentityPolicy(definition, nil), ""},
	}
	for _, c := range cases {
		b := fixtureBlueprint
		c.edit(&b)
		got := unit.PathIdentityIssues(b, c.intents, c.d)
		switch {
		case c.want == "" && len(got) != 0:
			t.Errorf("%s: rejected: %v", c.name, got)
		case c.want != "" && (len(got) != 1 || !strings.HasPrefix(got[0].Path+": "+got[0].Message, c.want)):
			t.Errorf("%s: issues %v, want %q", c.name, got, c.want)
		}
	}
}

// Drafting sends a resolved identity issue to the targeted repair of the
// departing purchase, a draft that keeps it publishes nothing, and check
// reports it as a path-identity Finding.
func TestPathIdentityIsRepairedAndChecked(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	taken := func() *s.Object {
		mech := recordedOutput(t, "mechanics")
		tier := at(mech, "paths", "path3", "tiers", "tier4").(*s.Object)
		changes, _ := tier.Get("statChanges")
		tier.Set("statChanges", append(changes.([]any), s.NewObject().Set("stat", "projectiles").Set("operation", "add").Set("value", 1.0)))
		return mech
	}
	zero := 0
	scripted := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), taken()}}
	if _, err := unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{MaxRepairAttempts: &zero}); err == nil ||
		!strings.Contains(err.Error(), "paths.path3.tiers.tier4: Resolved x-3-x gives attack-rate and projectiles") {
		t.Errorf("published a unit whose third purchase lost its identity: %v", err)
	}
	scripted = &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), taken()}}
	_, _ = unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{})
	if len(scripted.Requests) != 3 || !strings.Contains(scripted.Requests[2].Prompt, `"violations":["paths.path3.tiers.tier4: Resolved x-3-x gives attack-rate and projectiles`) ||
		!strings.Contains(scripted.Requests[2].Prompt, m.PathIdentityRule) {
		t.Errorf("targeted repair: %s", retryReason(scripted.Requests[len(scripted.Requests)-1].Prompt))
	}

	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	draft := stages.Draft
	blueprint := *draft.Candidate.Blueprint
	blueprint.Paths.Path3.Tiers.Tier4.Changes = append(append([]m.Change{}, blueprint.Paths.Path3.Tiers.Tier4.Changes...), statChange("projectiles", "add", 1))
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
	if len(failed) != 1 || !strings.HasPrefix(failed[0], unit.PathIdentityFindingRule+" paths.path3.tiers.tier4: Resolved x-3-x gives attack-rate and projectiles") {
		t.Errorf("check findings: %v", failed)
	}
}

// requireTier5BehaviorChange, on in the Default, makes the fifth purchase
// the pinnacle of its path. A proposed mechanic counts beside a supported
// behavior, so a capstone is never only larger numbers; an Active-only
// capstone that only raises the Active's numbers fails (#61).
func TestFifthPurchaseIsThePinnacle(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	definition := *stages.Draft.Prepared.Request.MechanicsDefinition
	rule := m.BehaviorChangeRule(5)

	// Plan stage.
	plan := *stages.Draft.Run.DesignPlan
	fifth := func(edit func(*unit.UpgradeIntents)) []string {
		intents := *plan.UpgradeIntents
		edit(&intents)
		p := plan
		p.UpgradeIntents = &intents
		var found []string
		for _, issue := range unit.PlanFeasibilityIssues(p, definition) {
			if strings.Contains(issue.Message, rule) {
				found = append(found, issue.Path+": "+issue.Message)
			}
		}
		return found
	}
	if got := fifth(func(*unit.UpgradeIntents) {}); len(got) != 0 {
		t.Errorf("the fixture's capstones fail: %v", got)
	}
	if got := fifth(func(i *unit.UpgradeIntents) { i.Path2.Tier5.ProposedMechanics = nil }); len(got) != 1 ||
		!strings.HasPrefix(got[0], "upgradeIntents.path2.tier5: x-5-x promises no new behavior or access and names no proposed mechanic. "+rule) {
		t.Errorf("an Active-only capstone without a proposal: %v", got)
	}
	if got := fifth(func(i *unit.UpgradeIntents) {
		i.Path1.Tier5 = intent("none", "damage", "pierce")
		i.Path1.Tier5.ProposedMechanics = []m.ProposedMechanic{rebound}
	}); len(got) != 0 {
		t.Errorf("a larger-number capstone with a proposed mechanic: %v", got)
	}
	if got := fifth(func(i *unit.UpgradeIntents) { i.Path1.Tier5 = intent("none", "damage", "pierce") }); len(got) != 1 ||
		!strings.HasPrefix(got[0], "upgradeIntents.path1.tier5: 5-x-x promises no new behavior or access and names no proposed mechanic.") {
		t.Errorf("a larger-number capstone: %v", got)
	}

	// Resolved stage.
	resolved := func(edit func(*m.Blueprint)) []string {
		b := *stages.Draft.Candidate.Blueprint
		edit(&b)
		var found []string
		for _, issue := range m.DesignPolicyIssues(&b, definition) {
			if strings.Contains(issue.Message, rule) {
				found = append(found, issue.Path+": "+issue.Message)
			}
		}
		return found
	}
	if got := resolved(func(*m.Blueprint) {}); len(got) != 0 {
		t.Errorf("the fixture's capstones fail: %v", got)
	}
	if got := resolved(func(b *m.Blueprint) { b.Paths.Path2.Tiers.Tier5.ProposedMechanics = nil }); len(got) != 1 ||
		got[0] != "paths.path2.tiers.tier5: Resolved x-5-x adds no new behavior or access over x-4-x and carries no proposed mechanic. "+rule {
		t.Errorf("an Active-only capstone that only raises the Active's numbers: %v", got)
	}
	larger := []m.Change{statChange("damage", "add", 3), statChange("pierce", "add", 150)}
	if got := resolved(func(b *m.Blueprint) { b.Paths.Path1.Tiers.Tier5.Changes = larger }); len(got) != 1 ||
		!strings.HasPrefix(got[0], "paths.path1.tiers.tier5: Resolved 5-x-x adds no new behavior or access over 4-x-x and carries no proposed mechanic.") {
		t.Errorf("a larger-number capstone: %v", got)
	}
	if got := resolved(func(b *m.Blueprint) {
		b.Paths.Path1.Tiers.Tier5.Changes = larger
		b.Paths.Path1.Tiers.Tier5.ProposedMechanics = []m.ProposedMechanic{rebound}
	}); len(got) != 0 {
		t.Errorf("a larger-number capstone with a proposed mechanic: %v", got)
	}
}

// The naming rule is prompt and review text (#61): the mechanics prompt
// states it and asks a revision to rename from a failed finding, and the
// review fails a borrowed source concept with the replacement in its action.
// Code judges no name, so such a name raises no deterministic finding.
func TestNamingRuleIsStatedAndReviewed(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	model := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	if _, err := unit.DraftUnit(context.Background(), stages.Prepared, model, fixture.Options()); err != nil {
		t.Fatal(err)
	}
	mechanicsPrompt := model.Requests[1].Prompt
	for _, want := range []string{unit.NamingRule, "In a revision, rename each purchase whose name a failed previous finding says breaks this rule, as that finding's action says."} {
		if !strings.Contains(mechanicsPrompt, want) {
			t.Errorf("the mechanics prompt lacks %q", want)
		}
	}
	review := unit.BlueprintReviewRequest(stages.Checked).Prompt
	for _, want := range []string{
		"or when the name breaks the naming rule. " + unit.NamingRule,
		"A name taken from the source still fails when the source uses it for something other than what the purchase does.",
		"give the replacement in the action: the source's name for what the purchase does, or a descriptive word that matches its typed changes or proposed mechanics",
		"its third purchase defines the path or at least distinguishes it, its fourth develops that identity, and its fifth is the path's pinnacle",
	} {
		if !strings.Contains(review, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	draft := stages.Draft
	blueprint := *draft.Candidate.Blueprint
	blueprint.Paths.Path3.Tiers.Tier5.Name = "Farseeing Reach"
	if draft.Candidate, err = unit.CompileBlueprint(blueprint, draft.Prepared.Request); err != nil {
		t.Fatal(err)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range checked.Findings {
		if f.Outcome == "fail" {
			t.Errorf("code judged a name: %s %s: %s", f.Rule, f.Subject, f.Message)
		}
	}
}
