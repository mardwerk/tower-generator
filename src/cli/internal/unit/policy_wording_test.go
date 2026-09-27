package unit_test

import (
	"context"
	"encoding/json"
	"slices"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Each design policy field's rule is stated once in mechanics/policy_text.go
// and reaches the plan prompt, the draft and repair guidance and the field's
// failures verbatim (#33). These tests pin the sentences, so a wording change
// is deliberate, and check that every stage carries the same one.

// withPolicy returns the fixture's prepared request under an edited copy of
// the Default design policy.
func withPolicy(t *testing.T, edit func(*m.DesignPolicy)) unit.Prepared {
	t.Helper()
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	definition := *prepared.Request.MechanicsDefinition
	policy := *definition.Profile.DesignPolicy
	edit(&policy)
	definition.Profile.DesignPolicy = &policy
	prepared.Request.MechanicsDefinition = &definition
	return prepared
}

// policyTexts are the plan prompt and the draft guidance of a request.
func policyTexts(t *testing.T, prepared unit.Prepared) (string, string) {
	t.Helper()
	plan, err := unit.DesignPlanRequest(prepared)
	if err != nil {
		t.Fatal(err)
	}
	return plan.Prompt, strings.Join(unit.DesignGuidance(&prepared.Request), "\n")
}

func messages(issues []m.Issue) string {
	var out []string
	for _, issue := range issues {
		out = append(out, issue.Path+": "+issue.Message)
	}
	return strings.Join(out, "\n")
}

func TestPolicyRuleSentencesArePinned(t *testing.T) {
	v2 := unit.DefaultAuthoringDefinition()
	v1 := m.DefaultDefinition()
	active := func(path m.NullableString, maximum int) *m.Definition {
		d := unit.DefaultAuthoringDefinition()
		policy := *d.Profile.DesignPolicy
		policy.ManualAbilityPath, policy.MaxManualAbilityPaths = path, maximum
		d.Profile.DesignPolicy = &policy
		return &d
	}
	noBonus := unit.DefaultAuthoringDefinition()
	vocabulary := *noBonus.Vocabulary
	vocabulary.BonusDamageProperties = nil
	noBonus.Vocabulary = &vocabulary
	absent, none := m.NullableString{}, m.NullableString{Present: true, Null: true}
	middle, bottom := m.NullableString{Present: true, Value: "path2"}, m.NullableString{Present: true, Value: "path3"}
	for _, pin := range []struct{ name, got, want string }{
		{"specializations", m.DistinctSpecializationsRule, "The three paths declare three different specializations."},
		{"first purchases", m.DistinctFirstPurchasesRule, "No two first purchases (1-x-x, x-1-x, x-x-1) may produce the same resolved attack. Names, prices and different arithmetic expressions do not make identical effects different."},
		{"capstones", m.DistinctCapstonesRule, "Pure 5-0-0, 0-5-0 and 0-0-5 builds must differ mechanically even after ignoring names and costs."},
		{"early benefits", m.EarlyBenefitsRule, "Two paths' first two purchases together must differ in what they improve or unlock, whatever their order, names, prices or amounts; lowers do not count."},
		{"early identity", m.EarlyIdentityRule(&v2), "The first and second purchase of each path keep the base attack's form: they add no new status, bonus damage, splash, follow-up or distinct-target volley, change no delivery, targeting or damage type, and keep a single-projectile attack single. Improving existing stats and effects and adding personal detection remain allowed."},
		{"early identity, no bonus damage", m.EarlyIdentityRule(&noBonus), "The first and second purchase of each path keep the base attack's form: they add no new status, splash, follow-up or distinct-target volley, change no delivery, targeting or damage type, and keep a single-projectile attack single. Improving existing stats and effects and adding personal detection remain allowed."},
		{"early identity, version 1", m.EarlyIdentityRule(&v1), "The first and second purchase of each path keep the base attack's form: they add no new status, splash, follow-up or distinct-target volley, change no delivery, targeting or damage type, and keep a single-projectile attack single. Improving existing stats and effects and adding personal Camo detection remain allowed."},
		{"third purchase", m.BehaviorChangeRule(3), "The third purchase of every path must add a supported behavior or access: a new delivery, a distinct-target volley of more than one projectile, more than one projectile, splash, a status effect, a bounded follow-up, a new damage type or a newly detected trait; larger existing numbers, a targeting change, a new name or a change with no effect do not count."},
		{"fifth purchase", m.BehaviorChangeRule(5), "The fifth purchase of every path must add a supported behavior or access: a new delivery, a distinct-target volley of more than one projectile, more than one projectile, splash, a status effect, a bounded follow-up, a new damage type or a newly detected trait; larger existing numbers, a targeting change, a new name or a change with no effect do not count."},
		{"capstone multiplier", m.CapstoneMultiplierRule(3), "The fifth purchase of every path must improve an established specialty metric of the path's pure fourth purchase by at least 3 times; a gain in active duty fraction alone must also keep the active peak."},
		{"Active Ability, no Definition", m.ActiveAbilityRule(nil), "An Active Ability exists only where the Definition allows one."},
		{"Active Ability, no policy", m.ActiveAbilityRule(&v1), "An Active Ability may unlock at a path's fourth purchase."},
		{"Active Ability, the Default", m.ActiveAbilityRule(&v2), "Only the middle path may have an Active Ability, first at x-4-x; the other paths stay automatic."},
		{"Active Ability, bottom path", m.ActiveAbilityRule(active(bottom, 1)), "Only the bottom path may have an Active Ability, first at x-x-4; the other paths stay automatic."},
		{"Active Ability, any 1 path", m.ActiveAbilityRule(active(absent, 1)), "At most 1 path may have an Active Ability, first at a path's fourth purchase; automatic paths are complete designs."},
		{"Active Ability, any 2 paths", m.ActiveAbilityRule(active(absent, 2)), "At most 2 paths may have an Active Ability, first at a path's fourth purchase; automatic paths are complete designs."},
		{"Active Ability, null", m.ActiveAbilityRule(active(none, 1)), "No path may have an Active Ability; every path stays automatic."},
		{"Active Ability, absent with maximum 0", m.ActiveAbilityRule(active(absent, 0)), "No path may have an Active Ability; every path stays automatic."},
		{"Active Ability, middle with maximum 0", m.ActiveAbilityRule(active(middle, 0)), "No path may have an Active Ability; every path stays automatic."},
	} {
		if pin.got != pin.want {
			t.Errorf("%s:\n got %q\nwant %q", pin.name, pin.got, pin.want)
		}
	}
}

// unlockEnum reads the unlocks the plan output schema offers a milestone.
func unlockEnum(t *testing.T, request *unit.Request, path, tier string) []string {
	t.Helper()
	var schema map[string]any
	if err := json.Unmarshal([]byte(s.Stringify(s.JSONSchema(unit.PurchasePlanOutputSchema(request)))), &schema); err != nil {
		t.Fatal(err)
	}
	value := any(schema)
	for _, key := range []string{"properties", "paths", "properties", path, "properties", "milestones", "properties", tier, "properties", "unlock", "enum"} {
		value = value.(map[string]any)[key]
	}
	var out []string
	for _, unlock := range value.([]any) {
		out = append(out, unlock.(string))
	}
	return out
}

// With maxManualAbilityPaths 0 the plan prompt used to permit an Active
// Ability that the plan schema and both checks reject. Every state now gives
// the prompt, the guidance, the schema and the failures one rule.
func TestActiveAbilityRuleAgrees(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	for _, state := range []struct {
		name    string
		path    m.NullableString
		maximum int
	}{
		{"any path, at most 1", m.NullableString{}, 1},
		{"any path, at most 0", m.NullableString{}, 0},
		{"null", m.NullableString{Present: true, Null: true}, 1},
		{"middle path", m.NullableString{Present: true, Value: "path2"}, 1},
		{"middle path, at most 0", m.NullableString{Present: true, Value: "path2"}, 0},
	} {
		prepared := withPolicy(t, func(p *m.DesignPolicy) { p.ManualAbilityPath, p.MaxManualAbilityPaths = state.path, state.maximum })
		definition := prepared.Request.MechanicsDefinition
		policy := definition.Profile.DesignPolicy
		rule := m.ActiveAbilityRule(definition)
		forbidden := state.maximum == 0 || state.path.Null
		plan, guidance := policyTexts(t, prepared)
		if !strings.Contains(plan, rule) || !strings.Contains(guidance, rule) {
			t.Errorf("%s: the plan prompt or guidance lacks %q", state.name, rule)
		}
		if forbidden == strings.Contains(plan, "unlocked at that purchase") || forbidden != strings.Contains(plan, "Adapt an activated technique from the sources as permanent purchases") {
			t.Errorf("%s: the plan prompt's Active paragraph does not follow the rule", state.name)
		}
		for _, path := range m.PathKeys {
			allowed := slices.Contains(unlockEnum(t, &prepared.Request, path, "tier4"), "manual-boost")
			if allowed != m.ActiveAbilityAllowed(policy, path) || (forbidden && allowed) {
				t.Errorf("%s: the plan schema offers manual-boost on %s: %v", state.name, path, allowed)
			}
		}

		// The fixture's Active is on the middle path; a top-path Active is
		// forbidden or one too many in every state.
		withTop := recordedOutput(t, "plan")
		at(withTop, "paths", "path1", "milestones", "tier4").(*s.Object).Set("unlock", "manual-boost")
		_, planErr := unit.DecodeDesignPlan(withTop, &prepared.Request)
		if planErr == nil || !strings.Contains(planErr.Error(), rule) {
			t.Errorf("%s: plan issue lacks the rule: %v", state.name, planErr)
		}
		blueprint := *stages.Draft.Candidate.Blueprint
		blueprint.Paths.Path1.Tiers.Tier4.Changes = blueprint.Paths.Path2.Tiers.Tier4.Changes
		resolved := messages(unit.ValidateBlueprintRequest(blueprint, prepared.Request))
		if !strings.Contains(resolved, "paths.path1.tiers.tier4: Resolved 4-x-x unlocks an Active Ability. "+rule) &&
			!strings.Contains(resolved, "paths.path2.tiers.tier4: Resolved 4-x-x and x-4-x each unlock an Active Ability. "+rule) {
			t.Errorf("%s: resolved issue lacks the rule:\n%s", state.name, resolved)
		}
		everything := plan + guidance + planErr.Error() + resolved
		for _, retired := range []string{"player-activated", "manual abilit", "active boosts", "At most 0", "1 paths", "manual boosts"} {
			if strings.Contains(everything, retired) {
				t.Errorf("%s: a policy text still says %q", state.name, retired)
			}
		}
	}
}

// The early-identity failures used to list different forbidden changes,
// and the resolved one omitted targeting, damage type and a second
// projectile although the check rejects them.
func TestEarlyIdentityRuleAgrees(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prepared := stages.Draft.Prepared
	rule := m.EarlyIdentityRule(prepared.Request.MechanicsDefinition)
	plan, guidance := policyTexts(t, prepared)
	if !strings.Contains(plan, rule+" At a first or second purchase, promise no unlock except a detection trait, and no projectiles.") || !strings.Contains(guidance, rule) {
		t.Error("the plan prompt or guidance lacks the early-identity rule")
	}
	off := withPolicy(t, func(p *m.DesignPolicy) { p.PreserveEarlyAttackIdentity = nil })
	if plan, guidance := policyTexts(t, off); strings.Contains(plan, rule) || strings.Contains(guidance, rule) {
		t.Error("the early-identity rule is stated while the field is off")
	}

	followUp := &m.FollowUp{Name: "Shockwave", Count: 2, DamageMultiplier: 0.5, Radius: 20}
	for _, kind := range []struct {
		fact   string
		change m.Change
	}{
		{"adds burn", m.Change{Kind: "status", Target: "base", Effect: "burn", Field: "magnitude", Operation: "set", Number: 1}},
		{"adds splash", statChange("splashRadius", "set", 10)},
		{"adds a follow-up", m.Change{Kind: "followUp", Target: "base", FollowUp: followUp}},
		{"adds a distinct-target volley", m.Change{Kind: "distribution", Target: "base", Text: "distinct-targets"}},
		{"changes delivery", m.Change{Kind: "delivery", Target: "base", Text: "beam"}},
		{"changes targeting", m.Change{Kind: "targeting", Target: "base", Text: "strong"}},
		{"changes damage type", m.Change{Kind: "damageType", Target: "base", Text: "energy"}},
		{"fires more than one projectile", statChange("projectiles", "set", 2)},
	} {
		blueprint := *stages.Draft.Candidate.Blueprint
		tier2 := &blueprint.Paths.Path1.Tiers.Tier2
		tier2.Changes = append(append([]m.Change{}, tier2.Changes...), kind.change)
		want := "paths.path1.tiers.tier2.changes: Resolved 2-x-x " + kind.fact + ". " + rule + " Move these to the third purchase or later."
		if got := messages(unit.ValidateBlueprintRequest(blueprint, prepared.Request)); !strings.Contains(got, want) {
			t.Errorf("%s: want %q in\n%s", kind.fact, want, got)
		}
	}

	burn := recordedOutput(t, "plan")
	at(burn, "paths", "path1", "milestones", "tier2").(*s.Object).Set("unlock", "burn")
	if _, err := unit.DecodeDesignPlan(burn, &prepared.Request); err == nil || !strings.Contains(err.Error(), "upgradeIntents.path1.tier2: 2-x-x promises to unlock burn. "+rule+" Unlock burn at the third purchase or later.") {
		t.Errorf("burn plan issue: %v", err)
	}
	projectiles := recordedOutput(t, "plan")
	at(projectiles, "paths", "path1", "milestones", "tier1").(*s.Object).Set("improves", []any{"pierce", "projectiles"})
	if _, err := unit.DecodeDesignPlan(projectiles, &prepared.Request); err == nil || !strings.Contains(err.Error(), "upgradeIntents.path1.tier1: 1-x-x promises projectiles. "+rule) {
		t.Errorf("projectiles plan issue: %v", err)
	}
	intents := *stages.Draft.Run.DesignPlan.UpgradeIntents
	intents.Path1.Tier1.Improves = append(append([]string{}, intents.Path1.Tier1.Improves...), "splash")
	promise := messages(unit.PlanIntentIssues(*stages.Draft.Candidate.Blueprint, &intents, *prepared.Request.MechanicsDefinition))
	if !strings.Contains(promise, "paths.path1.tiers.tier1.planIntent: The retained plan promises improved splash at 1-x-x, but the base attack has none. "+rule+" Give the base attack a splashRadius") {
		t.Errorf("promise issue:\n%s", promise)
	}

	// The plan issue reaches the full-plan correction with the rule.
	scripted := &fixture.Model{Outputs: []any{burn, recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	if _, err := unit.DraftUnit(context.Background(), prepared, scripted, unit.Options{}); err != nil {
		t.Fatal(err)
	}
	if len(scripted.Requests) != 3 || !strings.Contains(scripted.Requests[1].Prompt, "paths.path1.milestones.tier2: 2-x-x promises to unlock burn. "+rule) {
		t.Errorf("plan correction: %s", retryReason(scripted.Requests[1].Prompt))
	}
}

// requireTier3BehaviorChange and requireTier5BehaviorChange each state their
// rule only under their field, and the resolved failure carries it. The
// fifth-purchase rule had a check but no prompt sentence.
func TestBehaviorChangeRulesAgree(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	yes := true
	for _, tier := range []int{3, 5} {
		rule, other := m.BehaviorChangeRule(tier), m.BehaviorChangeRule(8-tier)
		defaults := stages.Draft.Prepared
		if plan, guidance := policyTexts(t, defaults); strings.Contains(plan, rule) || strings.Contains(guidance, rule) {
			t.Errorf("tier %d: the Default Profile states %q", tier, rule)
		}
		prepared := withPolicy(t, func(p *m.DesignPolicy) {
			if tier == 3 {
				p.RequireTier3BehaviorChange = &yes
			} else {
				p.RequireTier5BehaviorChange = &yes
			}
		})
		plan, guidance := policyTexts(t, prepared)
		if !strings.Contains(plan, rule) || !strings.Contains(guidance, rule) || strings.Contains(plan+guidance, other) {
			t.Errorf("tier %d: the plan prompt or guidance does not state exactly its rule", tier)
		}
		blueprint := *stages.Draft.Candidate.Blueprint
		blueprint.Paths.Path1.Tiers.At(tier).Changes = []m.Change{statChange("damage", "add", 1)}
		code, before := unit.BuildCode(0, tier), unit.BuildCode(0, tier-1)
		want := "paths.path1.tiers.tier" + string(rune('0'+tier)) + ": Resolved " + code + " adds no new behavior or access over " + before + ". " + rule
		if got := messages(unit.ValidateBlueprintRequest(blueprint, prepared.Request)); !strings.Contains(got, want) {
			t.Errorf("tier %d: want %q in\n%s", tier, want, got)
		}
	}
}

// The remaining fields state their rule in the guidance and in each failure.
func TestDistinctPurchaseRulesAgree(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	multiplier := 3.0
	prepared := withPolicy(t, func(p *m.DesignPolicy) {
		p.DistinctPathSpecializations, p.MinTier5SpecialtyMultiplier = true, &multiplier
	})
	_, guidance := policyTexts(t, prepared)
	for _, rule := range []string{m.DistinctSpecializationsRule, m.DistinctFirstPurchasesRule, m.DistinctCapstonesRule, m.CapstoneMultiplierRule(3)} {
		if !strings.Contains(guidance, rule) {
			t.Errorf("the guidance lacks %q", rule)
		}
	}
	blueprint := *stages.Draft.Candidate.Blueprint
	blueprint.Paths.Path3 = blueprint.Paths.Path1
	got := messages(unit.ValidateBlueprintRequest(blueprint, prepared.Request))
	specialization := blueprint.Paths.Path1.Specialization
	for _, want := range []string{
		"paths.path3.specialization: The top path already declares " + specialization + ". " + m.DistinctSpecializationsRule + " Choose another specialization for the bottom path.",
		"paths.path3.tiers.tier1: Resolved x-x-1 behaves exactly like 1-x-x. " + m.DistinctFirstPurchasesRule + " Change what x-x-1 improves, or by how much, so the resolved attack differs.",
		"paths.path3.tiers.tier5: Resolved x-x-5 behaves exactly like 5-x-x. " + m.DistinctCapstonesRule + " Change what x-x-5 improves, or by how much, so the resolved attack differs.",
		m.CapstoneMultiplierRule(3) + " These are capacity heuristics",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("want %q in\n%s", want, got)
		}
	}
}
