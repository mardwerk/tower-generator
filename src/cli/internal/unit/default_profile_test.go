package unit_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// The default Profile cites the pinned btd6-atlas capture and none of the
// retired dataset or deleted snapshots.
func TestDefaultProfileCitesThePinnedAtlasCapture(t *testing.T) {
	profile := unit.DefaultProfile()
	rules := profile.Rules.Text
	for _, want := range []string{
		"btd6-atlas capture 56.3, Steam build 24829026, repository revision a380413",
		"DartMonkey-100 to -500", "BoomerangMonkey.json and -010 to -050", "SniperMonkey.json and -100 to -500",
		"IceMonkey.json and -100 to -500", "TackShooter.json and -010 to -050", "MonkeyVillage.json",
		"DartMonkey-400 and SuperMonkey-001, KnockbackModel",
		"0-0-0", "Private design checks, never printed in the unit",
		"a substantial improvement of one dimension may qualify", "no universal capstone multiplier",
	} {
		if !strings.Contains(rules, want) {
			t.Errorf("the default rules lack %q", want)
		}
	}
	// The Profile keeps design taste and cited scale. Build codes, legality,
	// crosspaths, Engine limits and output format belong to the Definition,
	// the Engine's stage instructions and the renderer, so a custom Profile
	// receives them too.
	for _, generic := range []string{"Build codes.", "Crosspaths.", "Engine boundary.", "Unit output.", "Activation.", "Actions.", "Source fidelity.", "version 13"} {
		if strings.Contains(rules, generic) {
			t.Errorf("the default rules still hold the generic section %q", generic)
		}
	}
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan, err := unit.DesignPlanRequest(prepared)
	if err != nil {
		t.Fatal(err)
	}
	everything := rules + profile.Task + plan.Prompt + plan.System
	for _, retired := range []string{"btd6_towers.json", "a2a5e2bb4591", "source-snapshots", "bloons.fandom.com", "cyberquincy", "workedExamples", "universal threefold"} {
		if strings.Contains(everything, retired) {
			t.Errorf("the default Profile or plan prompt still cites %q", retired)
		}
	}
	for _, derived := range []string{
		"Only the middle path may have a player-activated ability, first at x-4-x.",
		"A Unit buys at most 2 paths, and at most 1 of them beyond its second purchase; 3-3-0 and 1-1-1 are illegal. Code resolves the 12 early and 36 advanced crosspath builds",
	} {
		if !strings.Contains(plan.Prompt, derived) {
			t.Errorf("the plan prompt lacks the Definition's %q", derived)
		}
	}
	definition := profile.MechanicsDefinition
	if definition.Revision != "2026-09-26-atlas-56.3-v17" || !strings.Contains(definition.Label, "btd6-atlas 56.3") || definition.Profile.MaxChangesPerTier != 5 {
		t.Errorf("Definition %s %q", definition.Revision, definition.Label)
	}
	if scale := definition.Profile.ReferenceScale; scale.BaseCost != 200 || scale.BaseDamage != 1 || scale.BaseIntervalSeconds != 0.95 || scale.BaseRange != 32 || scale.BasePierce != 2 ||
		scale.IncrementalUpgradeCosts != [5]float64{140, 200, 320, 1800, 15000} {
		t.Errorf("reference scale %+v does not match DartMonkey.json and its top-path upgrades", scale)
	}
	if policy := definition.Profile.DesignPolicy; policy.MinTier5SpecialtyMultiplier != nil {
		t.Error("the default Profile demands a universal capstone multiplier")
	}
}

// Every stage's prompt states which instructions are private checks and
// which text reaches the unit, and none demands one recipe for every path.
func TestPromptsSeparatePrivateChecksFromOutput(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	model := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), recordedOutput(t, "mechanics"), recordedOutput(t, "review")}}
	draft, err := unit.DraftUnit(context.Background(), stages.Prepared, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	checked, _ := unit.CheckDraft(draft)
	if _, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options()); err != nil {
		t.Fatal(err)
	}
	plan, mechanics, review := model.Requests[0].Prompt, model.Requests[1].Prompt, model.Requests[2].Prompt
	for name, test := range map[string]struct {
		prompt string
		want   []string
	}{
		"plan": {plan, []string{
			"buyFor, weakness and capstoneValue are private design checks and never appear in the unit description",
			"Name every purchase by build code", "a substantial improvement of one dimension may qualify",
			"A second Active at the fifth purchase", "12 early and 36 advanced crosspath builds",
			"more shots per attack to projectiles, heavier hits to damage",
		}},
		"mechanics": {mechanics, []string{
			"Refer to purchases by build code", "The boost is the only activated ability this Definition expresses",
			"List in unsupportedMechanics", "a substantial improvement of one dimension may qualify",
			"No universal capstone multiplier applies", "Scale references from btd6-atlas capture 56.3",
			"Code checks each milestone's improves and unlock in every legal build", "more projectiles do not count",
			"Tier1 to Tier2 allow 1 to 3 primitive changes total; Tier3 to Tier5 allow up to 5.",
		}},
		"review": {review, []string{
			"Check privately and report only concrete problems", "Findings are review data kept apart from the unit description",
			"Do not demand a universal capstone multiplier, a new subsystem at every purchase or one recipe for every path",
		}},
	} {
		for _, want := range test.want {
			if !strings.Contains(test.prompt, want) {
				t.Errorf("%s prompt lacks %q", name, want)
			}
		}
		for _, banned := range []string{"—", "–"} {
			if strings.Contains(strings.SplitN(test.prompt, "{", 2)[0], banned) {
				t.Errorf("%s instructions use a dash %q", name, banned)
			}
		}
	}
}

// Code does not judge payoff: a fourth or fifth purchase that promises one
// dimension passes the plan check, and the review judges it from the
// resolved purchase evidence (decided on #27). Impossible promises still fail.
func TestPlansLeavePayoffToTheReview(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan := recordedOutput(t, "plan")
	for _, path := range []string{"path1", "path3"} {
		for _, tier := range []string{"tier4", "tier5"} {
			at(plan, "paths", path, "milestones", tier).(*s.Object).Set("improves", []any{"damage"}).Set("unlock", "none")
		}
	}
	if _, err := unit.DecodeDesignPlan(plan, &prepared.Request); err != nil {
		t.Fatalf("single-dimension fourth and fifth purchases were rejected: %v", err)
	}
	at(plan, "paths", "path3", "milestones", "tier5").(*s.Object).Set("improves", []any{"damage", "active-damage"})
	if _, err := unit.DecodeDesignPlan(plan, &prepared.Request); err == nil || !strings.Contains(err.Error(), "same-path manual-boost") {
		t.Errorf("an active promise without a boost was accepted: %v", err)
	}
	if request, _ := unit.DesignPlanRequest(prepared); strings.Contains(request.Prompt, "code rejects") {
		t.Error("the plan prompt still says code rejects a token step")
	}
}

// Every blueprint proposal becomes an explicit unsupported-mechanic finding.
func TestUnsupportedMechanicsAreFindings(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, finding := range stages.Checked.Findings {
		if finding.Rule == unit.UnsupportedMechanicRule {
			if finding.Category != "unsupported" || finding.Outcome != "unresolved" {
				t.Errorf("finding %+v", finding)
			}
			names = append(names, finding.Message)
		}
	}
	if len(names) != 3 || !strings.Contains(names[1], "Unsupported mechanic Critical shot counter") {
		t.Errorf("unsupported findings %v", names)
	}
}

// The change budget belongs to the Definition profile: the Default Profile
// allows the five changes of a real Crossbow Master, a Definition that sets
// four rejects them, and no version 2 Definition may exceed the limit.
func TestTheDefinitionProfileSetsTheChangeBudget(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	blueprint := *stages.Result.Candidate.Blueprint
	tier := &blueprint.Paths.Path3.Tiers.Tier5
	tier.Changes = append(append([]m.Change(nil), tier.Changes...), m.Change{Kind: "stat", Target: "base", Stat: "range", Operation: "add", Number: 20})
	if len(tier.Changes) != 5 {
		t.Fatalf("x-x-5 has %d changes", len(tier.Changes))
	}
	definition := unit.DefaultAuthoringDefinition()
	if issues := m.ValidateBlueprint(s.FromGoValue(blueprint), s.FromGoValue(definition)); len(issues) > 0 {
		t.Errorf("five changes rejected under the Default Profile: %v", issues)
	}
	definition.Profile.MaxChangesPerTier = 4
	issues := m.ValidateBlueprint(s.FromGoValue(blueprint), s.FromGoValue(definition))
	if len(issues) == 0 || !strings.Contains(fmt.Sprint(issues), "Exceeds the Definition change budget.") {
		t.Errorf("five changes accepted under a four-change Definition: %v", issues)
	}
	definition.Profile.MaxChangesPerTier = m.MaxChangesLimit + 1
	if _, issues := s.Parse(m.DefinitionV2Schema, s.FromGoValue(definition)); len(issues) == 0 {
		t.Error("a Definition exceeded the change limit")
	}
}

// A third purchase needs a significant reason to commit, which the model
// review judges; a substantial stat change qualifies, so the Default Profile
// accepts one. A Profile can opt in to requiring a behavior or access, and
// then a stat-only or targeting-only third purchase is rejected at the plan.
func TestThirdPurchaseBehaviorIsAnOptInGate(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	statOnly := recordedOutput(t, "plan")
	at(statOnly, "paths", "path3", "milestones", "tier3").(*s.Object).Set("improves", []any{"damage", "range"}).Set("unlock", "none")
	if _, err := unit.DecodeDesignPlan(statOnly, &prepared.Request); err != nil {
		t.Fatalf("the Default Profile rejected a stat-only third purchase: %v", err)
	}
	if request, _ := unit.DesignPlanRequest(prepared); strings.Contains(request.Prompt, "This Profile requires the third purchase") {
		t.Error("the Default Profile's plan prompt requires a third-purchase behavior")
	}
	optIn := prepared.Request
	definition := *optIn.MechanicsDefinition
	policy := *definition.Profile.DesignPolicy
	required := true
	policy.RequireTier3BehaviorChange = &required
	definition.Profile.DesignPolicy = &policy
	optIn.MechanicsDefinition = &definition
	for unlock, rejected := range map[string]bool{"none": true, "targeting-change": true, "splash": false} {
		plan := recordedOutput(t, "plan")
		at(plan, "paths", "path3", "milestones", "tier3").(*s.Object).Set("improves", []any{"damage", "range"}).Set("unlock", unlock)
		_, err := unit.DecodeDesignPlan(plan, &optIn)
		if got := err != nil && strings.Contains(err.Error(), "x-x-3 must add a supported behavior or access"); got != rejected {
			t.Errorf("opt-in, unlock %s: %v", unlock, err)
		}
	}
}

// Under the early-identity policy the first and second purchase keep a
// single-projectile attack single, so a plan cannot promise projectiles
// there; the mechanics could never keep that promise (reported on #27).
func TestEarlyPurchasesCannotPromiseProjectiles(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan := recordedOutput(t, "plan")
	at(plan, "paths", "path1", "milestones", "tier2").(*s.Object).Set("improves", []any{"pierce", "projectiles"})
	_, err = unit.DecodeDesignPlan(plan, &prepared.Request)
	if err == nil || !strings.Contains(err.Error(), "2-x-x cannot promise projectiles") {
		t.Errorf("an early projectiles promise was accepted: %v", err)
	}
	schema := s.Stringify(s.JSONSchema(unit.PurchasePlanOutputSchema(&prepared.Request)))
	if !strings.Contains(schema, `"projectiles"`) {
		t.Error("no purchase may promise projectiles at all")
	}
}

// A coherent stat-led path is allowed: code does not require a behavior or
// access on every path, and the review judges its third purchase and
// capstone (decided on #27).
func TestAStatLedPathIsAllowed(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan := recordedOutput(t, "plan")
	for _, tier := range []string{"tier2", "tier5"} {
		at(plan, "paths", "path3", "milestones", tier).(*s.Object).Set("unlock", "none")
	}
	if _, err := unit.DecodeDesignPlan(plan, &prepared.Request); err != nil {
		t.Errorf("a stat-led bottom path was rejected: %v", err)
	}
}

// Every purchase names the technique it adapts. Two paths may share a
// technique as crosspath synergy only where a legal build owns both
// purchases, and every path needs a technique of its own. The checks read the
// plan's typed techniques, so a shared name alone is no issue and a renamed
// purchase cannot hide a split (decided on #27).
func TestPathsShareTechniquesOnlyAsCrosspathSynergy(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	if request, _ := unit.DesignPlanRequest(prepared); !strings.Contains(request.Prompt, "technique") {
		t.Error("the plan prompt does not ask for techniques")
	}
	decode := func(edit func(plan *s.Object)) error {
		plan := recordedOutput(t, "plan")
		edit(plan)
		_, err := unit.DecodeDesignPlan(plan, &prepared.Request)
		return err
	}
	milestone := func(plan *s.Object, path, tier string) *s.Object {
		return at(plan, "paths", path, "milestones", tier).(*s.Object)
	}
	// x-4-x builds on the top path's 2-x-x: 2-4-0 owns both.
	if err := decode(func(plan *s.Object) {
		milestone(plan, "path1", "tier2").Set("technique", "Fan Club")
	}); err != nil {
		t.Errorf("crosspath synergy was rejected: %v", err)
	}
	// x-4-x and 3-x-x never meet in a legal build.
	err = decode(func(plan *s.Object) {
		milestone(plan, "path1", "tier3").Set("technique", "fan club")
	})
	if err == nil || !strings.Contains(err.Error(), `3-x-x and x-4-x both adapt "Fan Club", and no legal build owns both`) {
		t.Errorf("a technique split between exclusive purchases: %v", err)
	}
	// The middle path keeps only the base attack and the top path's technique.
	err = decode(func(plan *s.Object) {
		for _, tier := range []string{"tier3", "tier4", "tier5"} {
			milestone(plan, "path2", tier).Set("technique", "Dart Throw")
		}
		milestone(plan, "path2", "tier1").Set("technique", "Spiked Ball")
	})
	if err == nil || !strings.Contains(err.Error(), "The middle path adapts no technique of its own") {
		t.Errorf("a path without identity: %v", err)
	}
	err = decode(func(plan *s.Object) {
		milestone(plan, "path3", "tier3").Set("technique", "Crossbow Mastery")
	})
	if err == nil || !strings.Contains(err.Error(), `"Crossbow Mastery" is neither the base attack nor a repertoire entry`) {
		t.Errorf("an unknown technique: %v", err)
	}
	err = decode(func(plan *s.Object) {
		milestone(plan, "path3", "tier3").Delete("technique")
	})
	if err == nil {
		t.Error("a plan milestone without a technique was accepted")
	}
}

// A broken promise names the change that keeps it. The boost's multipliers
// never keep a permanent promise, also at the purchase that unlocks the boost
// (seen in a failed Luffy run on #27).
func TestBrokenPromisesNameTheirFix(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	blueprint, plan := *stages.Result.Candidate.Blueprint, *stages.Draft.Run.DesignPlan
	tier := &blueprint.Paths.Path2.Tiers.Tier4
	var kept []m.Change
	for _, change := range tier.Changes {
		if change.Kind != "stat" || change.Stat != "intervalSeconds" {
			kept = append(kept, change)
		}
	}
	tier.Changes = kept
	issues := unit.PlanIntentIssues(blueprint, plan.UpgradeIntents, unit.DefaultAuthoringDefinition())
	if len(issues) != 1 || issues[0].Path != "paths.path2.tiers.tier4.planIntent" ||
		!strings.Contains(issues[0].Message, "lowers intervalSeconds") || !strings.Contains(issues[0].Message, "intervalMultiplier is active-attack-rate and does not count") {
		t.Errorf("issues %v", issues)
	}
}

// A tradeoff the plan states is typed in lowers and checked like a promise:
// Spike-o-pult's slower throw must slow the attack in every legal build that
// owns it (a Gear 4 run on #27 promised a slower, heavier attack in words and
// never changed the interval).
func TestTradeoffsAreCheckedLikePromises(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	blueprint, plan := *stages.Result.Candidate.Blueprint, *stages.Draft.Run.DesignPlan
	if lowers := plan.UpgradeIntents.Path1.Tier3.Lowers; len(lowers) != 1 || lowers[0] != "attack-rate" {
		t.Fatalf("the fixture's 3-x-x tradeoff was not retained: %v", lowers)
	}
	tier := &blueprint.Paths.Path1.Tiers.Tier3
	var kept []m.Change
	for _, change := range tier.Changes {
		if change.Stat != "intervalSeconds" {
			kept = append(kept, change)
		}
	}
	tier.Changes = kept
	issues := unit.PlanIntentIssues(blueprint, plan.UpgradeIntents, unit.DefaultAuthoringDefinition())
	if len(issues) != 1 || issues[0].Path != "paths.path1.tiers.tier3.planIntent" || !strings.Contains(issues[0].Message, "the tradeoff lowered attack-rate") || !strings.Contains(issues[0].Message, "raises intervalSeconds") {
		t.Errorf("issues %v", issues)
	}
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	contradictory := recordedOutput(t, "plan")
	at(contradictory, "paths", "path2", "milestones", "tier1").(*s.Object).Set("lowers", []any{"attack-rate"})
	if _, err := unit.DecodeDesignPlan(contradictory, &prepared.Request); err == nil || !strings.Contains(err.Error(), "cannot both raise and lower attack-rate") {
		t.Errorf("a purchase raising and lowering attack-rate: %v", err)
	}
}

// Build legality, crosspath counts and the Active Ability slot come from the
// Definition: a Profile that moves the Active to the bottom path gets that
// sentence, not the Default's middle path.
func TestPlanPromptDerivesTheActiveSlotFromTheDefinition(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	definition := *prepared.Request.MechanicsDefinition
	policy := *definition.Profile.DesignPolicy
	policy.ManualAbilityPath = m.NullableString{Present: true, Value: "path3"}
	definition.Profile.DesignPolicy = &policy
	prepared.Request.MechanicsDefinition = &definition
	request, err := unit.DesignPlanRequest(prepared)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(request.Prompt, "Only the bottom path may have a player-activated ability, first at x-x-4.") || strings.Contains(request.Prompt, "middle path may have a player-activated") {
		t.Error("the plan prompt does not follow the Definition's manualAbilityPath")
	}
}
