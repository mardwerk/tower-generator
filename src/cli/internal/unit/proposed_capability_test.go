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

// plasmaTransformation is a proposed mechanic for the scripted x-5-x: the
// allied transformation Plasma Monkey Fan Club pays for, which the
// Definition cannot express.
func plasmaTransformation() *s.Object {
	return s.NewObject().Set("name", "Plasma transformation").
		Set("effect", "While the frenzy runs, up to 20 nearby Dart Monkeys with no path above tier 2 attack about 16 times as often, and their darts deal 1 more damage and hit 3 more enemies.").
		Set("sourceIds", []any{"source1:12", "source1:13"})
}

// underTheDefault prepares the fixture request under the Default Profile,
// with requireTier5BehaviorChange on.
func underTheDefault(t *testing.T) unit.Prepared {
	t.Helper()
	request, err := fixture.Request()
	if err != nil {
		t.Fatal(err)
	}
	prepared, err := unit.Prepare(s.FromGoValue(unit.ApplyProfile(request, unit.DefaultProfile())))
	if err != nil {
		t.Fatal(err)
	}
	if !prepared.Request.MechanicsDefinition.Profile.DesignPolicy.RequiresBehaviorChange(5) {
		t.Fatal("the Default does not require a fifth-purchase behavior")
	}
	return prepared
}

// The Default judges the unchanged Dart Monkey fixture (#61, SOL-66-01).
// 5-x-x adds a follow-up and x-x-5 a new damage type, supported
// capabilities that pass. x-5-x, Plasma Monkey Fan Club, only doubles the
// frenzy's damage and lengthens it: its allied transformation is a
// blueprint-level unsupported mechanic, not a proposed mechanic of the
// purchase, so it fails the fifth-purchase rule, in the plan and in the
// resolved builds. Its exclusive early benefits hold, and no third purchase
// is judged by code.
func TestDartFixtureUnderTheDefault(t *testing.T) {
	prepared := underTheDefault(t)
	rule := m.BehaviorChangeRule(5)

	_, err := unit.DecodeDesignPlan(recordedOutput(t, "plan"), &prepared.Request)
	want := "upgradeIntents.path2.tier5: x-5-x promises no new behavior or access and names no proposed mechanic. " + rule
	if err == nil || !strings.Contains(err.Error(), want) || strings.Count(err.Error(), "upgradeIntents.") != 1 {
		t.Errorf("plan check: %v", err)
	}

	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	blueprint := *stages.Draft.Candidate.Blueprint
	want = "paths.path2.tiers.tier5: Resolved x-5-x adds no new behavior or access over x-4-x and carries no proposed mechanic. " + rule
	if got := messages(unit.ValidateBlueprintRequest(blueprint, prepared.Request)); got != want {
		t.Errorf("resolved check:\n%s\nwant\n%s", got, want)
	}
	intents := stages.Draft.Run.DesignPlan.UpgradeIntents
	if got := messages(unit.EarlyBenefitsIssues(blueprint, intents, *prepared.Request.MechanicsDefinition)); got != "" {
		t.Errorf("early benefits: %s", got)
	}
	if gaps := m.ProposedCapabilityGaps(&blueprint, *prepared.Request.MechanicsDefinition); len(gaps) != 0 {
		t.Errorf("design gaps %v", gaps)
	}
}

// A fifth purchase whose only new capability is a proposed mechanic passes
// generation under the Default: the plan check and the resolved check accept
// it, and check reports it as an unresolved design gap, because no build
// grants the mechanic until the Definition supports it. A proposed mechanic
// beside a supported capability, as on 5-x-x, is no gap.
func TestProposedOnlyCapstoneIsADesignGap(t *testing.T) {
	prepared := underTheDefault(t)
	plan := recordedOutput(t, "plan")
	at(plan, "paths", "path2", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{plasmaTransformation()})
	split := s.NewObject().Set("name", "Split").Set("effect", "When the ball's pierce runs out it splits into twelve smaller balls that each hit one more enemy.").Set("sourceIds", []any{"source1:8"})
	at(plan, "paths", "path1", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{split})
	review := withProposalVerdicts(recordedOutput(t, "review"),
		proposalVerdict{"5-x-x", "Split", "unresolved", "Splitting into smaller balls fits the Juggernaut ball and is a capability no typed change of 5-x-x has."},
		proposalVerdict{"x-5-x", "Plasma transformation", "unresolved", "Transforming allied Units is a coherent, source-backed capability the Definition cannot express."})
	model := &fixture.Model{Outputs: []any{plan, recordedOutput(t, "mechanics"), review}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	var gaps []unit.Finding
	for _, finding := range checked.Findings {
		if finding.Outcome == "fail" {
			t.Errorf("%s %s: %s", finding.Rule, finding.Subject, finding.Message)
		}
		if finding.Rule == unit.ProposedCapabilityRule {
			gaps = append(gaps, finding)
		}
	}
	want := "Resolved x-5-x adds no supported behavior or access over x-4-x. Its new capability, Plasma transformation, is proposed and not yet playable: until the Definition supports it, no build grants it and x-5-x adds only larger numbers in play. " + m.BehaviorChangeRule(5)
	if len(gaps) != 1 || gaps[0].Subject != "paths.path2.tiers.tier5" || gaps[0].Outcome != "unresolved" || gaps[0].Category != "missing_specification" ||
		gaps[0].Method != "deterministic" || gaps[0].Message != want {
		t.Fatalf("design gaps %+v", gaps)
	}
	if _, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options()); err != nil {
		t.Fatalf("review: %v", err)
	}
	if len(model.Requests) != 3 {
		t.Errorf("%d model calls, want plan, mechanics and review without a retry or repair", len(model.Requests))
	}
}

// A targeted repair of a third or fourth purchase rebuilds the fifth
// purchase under requireTier5BehaviorChange, also when the fifth purchase
// carries a proposed mechanic: the mechanic grants no behavior, so an
// earlier tier that took the fifth purchase's supported one would leave a
// design gap. Under early identity a first or second purchase leaves it out.
func TestRepairRebuildsAProposingCapstone(t *testing.T) {
	prepared := underTheDefault(t)
	output := recordedOutput(t, "plan")
	at(output, "paths", "path2", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{plasmaTransformation()})
	split := s.NewObject().Set("name", "Split").Set("effect", "When the ball's pierce runs out it splits into twelve smaller balls that each hit one more enemy.").Set("sourceIds", []any{"source1:8"})
	at(output, "paths", "path1", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{split})
	plan, err := unit.DecodeDesignPlan(output, &prepared.Request)
	if err != nil {
		t.Fatal(err)
	}
	previous := unit.BindDesignPlan(recordedOutput(t, "mechanics"), plan)
	for tier, dependent := range map[string]bool{"tier1": false, "tier2": false, "tier3": true, "tier4": true} {
		repair, err := unit.TargetedTierRepair(&prepared.Request, previous, []string{"paths.path1.tiers." + tier + ": Upgrade changes no behavior in legal build 1-0-0."})
		if err != nil || repair == nil {
			t.Fatalf("%s: %v", tier, err)
		}
		if got := strings.Contains(repair.Request.Prompt, `"dependentCapstones":["path1.tier5"]`); got != dependent {
			t.Errorf("repairing %s rebuilds 5-x-x: %v, want %v", tier, got, dependent)
		}
	}
}
