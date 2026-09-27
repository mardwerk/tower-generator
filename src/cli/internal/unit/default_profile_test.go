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
		"a substantial improvement of one dimension other than damage may qualify",
		// Whole-number damage is the Profile's taste, not a mechanic rule.
		"Keep damage per hit a whole number in every build", "Code accepts fractions",
	} {
		if !strings.Contains(rules, want) {
			t.Errorf("the default rules lack %q", want)
		}
	}
	// The Profile keeps design taste and cited scale. Build codes, legality,
	// crosspaths, Engine limits and output format belong to the Definition,
	// the Engine's stage instructions and the renderer, so a custom Profile
	// receives them too.
	// The review checklist, distinct paths, technique ownership and the
	// capstone multiplier each have one owner in the Engine's stage
	// instructions or the design policy.
	for _, generic := range []string{"Build codes.", "Crosspaths.", "Engine boundary.", "Unit output.", "Activation.", "Actions.", "Source fidelity.", "version 13",
		"a review for missing effects", "Keep the three paths behaviorally distinct", "universal capstone multiplier", "mechanicsDefinition owns"} {
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
		"Only the middle path may have an Active Ability, first at x-4-x; the other paths stay automatic.",
		"A Unit buys at most 2 paths, and at most 1 of them beyond its second purchase; 3-3-0 and 1-1-1 are illegal. Code resolves the 12 early and 36 advanced crosspath builds",
	} {
		if !strings.Contains(plan.Prompt, derived) {
			t.Errorf("the plan prompt lacks the Definition's %q", derived)
		}
	}
	definition := profile.MechanicsDefinition
	if definition.Revision != "2026-09-27-atlas-56.3-v39" || !strings.Contains(definition.Label, "btd6-atlas 56.3") || definition.Profile.MaxChangesPerTier != 5 {
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
	if strings.Contains(review, "energy rules") {
		t.Error("review guidance still assumes a Default Profile damage type")
	}
	for name, test := range map[string]struct {
		prompt string
		want   []string
	}{
		"plan": {plan, []string{
			"buyFor, weakness and capstoneValue are private design checks and never appear in the unit description",
			"Name every purchase by build code", "a substantial improvement of one dimension other than damage may qualify",
			"A second Active at the fifth purchase", "12 early and 36 advanced crosspath builds",
			"more shots per attack to projectiles, heavier hits to damage",
			"The unit shows the change of every purchase whose technique is not the base attack",
		}},
		"mechanics": {mechanics, []string{
			"Refer to purchases by build code", "The boost is the only activated ability this Definition expresses",
			"on the purchase that needs it: that tier's proposedMechanics lists", "a substantial improvement of one dimension other than damage may qualify",
			"No universal capstone multiplier applies", "Scale references from btd6-atlas capture 56.3",
			`"id":"sharp","name":"Sharp","description":"Darts, blades, spikes and arrows.","ineffectiveAgainst":["lead","frozen"]`,
			"ineffectiveAgainst lists the enemy properties it cannot damage",
			"Code checks each milestone's improves and unlock in every legal build", "more projectiles do not count",
			"Tier1 to Tier2 allow 1 to 3 primitive changes total; Tier3 to Tier5 allow up to 5.",
			// A Gatling plan promised 3 distinct enemies and got 3 shots with pierce 3 (#27).
			"each shot keeps its pierce, so a volley can hit up to projectiles × pierce enemies",
			"The unit shows each technique purchase's planned change beside its effects",
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
		// The vocabulary is described once, in the Definition JSON.
		if n := strings.Count(test.prompt, "Darts, blades, spikes and arrows."); n != 1 {
			t.Errorf("%s prompt describes Sharp %d times", name, n)
		}
	}
}

// Code does not judge payoff: a fourth purchase that promises one dimension
// passes the plan check, and the review judges it from the resolved purchase
// evidence (decided on #27). Under the Default's requireTier5BehaviorChange
// a fifth purchase that only raises damage is rejected, and one whose new
// capability is a proposed mechanic passes the plan check, to be reported as
// a design gap once resolved (#61); its size is still the review's
// judgment. Impossible promises still fail.
func TestPlansLeavePayoffToTheReview(t *testing.T) {
	yes := true
	prepared := withPolicy(t, func(p *m.DesignPolicy) { p.RequireTier5BehaviorChange = &yes })
	plan := recordedOutput(t, "plan")
	for _, path := range []string{"path1", "path3"} {
		for _, tier := range []string{"tier4", "tier5"} {
			at(plan, "paths", path, "milestones", tier).(*s.Object).Set("improves", []any{"damage"}).Set("unlock", "none")
		}
	}
	// 4-x-x no longer changes the damage type, so Spiked Ball omits Frozen access.
	repertoireEffect(plan, 0, 2).Set("adaptedAs", []any{})
	_, err := unit.DecodeDesignPlan(plan, &prepared.Request)
	for _, code := range []string{"5-x-x", "x-x-5"} {
		if err == nil || !strings.Contains(err.Error(), code+" promises no new behavior or access and names no proposed mechanic. "+m.BehaviorChangeRule(5)) {
			t.Errorf("a %s that only raises damage was accepted: %v", code, err)
		}
	}
	split := s.NewObject().Set("name", "Split").Set("effect", "When the ball's pierce runs out it splits into twelve smaller balls that each hit one more enemy.").Set("sourceIds", []any{"source1:8"})
	critical := s.NewObject().Set("name", "Critical bolt").Set("effect", "Every fifth bolt deals ten times its damage.").Set("sourceIds", []any{"source1:18"})
	at(plan, "paths", "path1", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{split})
	at(plan, "paths", "path3", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{critical})
	// The scripted x-5-x only raises its Active Ability's numbers.
	at(plan, "paths", "path2", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{plasmaTransformation()})
	if _, err := unit.DecodeDesignPlan(plan, &prepared.Request); err != nil {
		t.Fatalf("single-dimension fourth purchases and fifth purchases with a proposed capability were rejected: %v", err)
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
	if len(names) != 3 || !strings.Contains(names[1], "Draft claim: the Definition cannot express Critical shot counter.") {
		t.Errorf("unsupported findings %v", names)
	}
}

// The change budget belongs to the Definition profile: the Default Profile
// allows the five changes of a real Crossbow Master, a Definition that sets
// four rejects them, and no version 2 Definition may exceed the limit. The
// fixture's Profile is the Default without requireTier5BehaviorChange, whose
// rule the scripted x-5-x fails.
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
	definition := fixture.Profile().MechanicsDefinition
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
	if request, _ := unit.DesignPlanRequest(prepared); strings.Contains(request.Prompt, m.BehaviorChangeRule(3)) {
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
		if got := err != nil && strings.Contains(err.Error(), "x-x-3 promises no new behavior or access and names no proposed mechanic. "+m.BehaviorChangeRule(3)); got != rejected {
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
	if err == nil || !strings.Contains(err.Error(), "2-x-x promises projectiles. "+m.EarlyIdentityRule(prepared.Request.MechanicsDefinition)) {
		t.Errorf("an early projectiles promise was accepted: %v", err)
	}
	schema := s.Stringify(s.JSONSchema(unit.PurchasePlanOutputSchema(&prepared.Request)))
	if !strings.Contains(schema, `"projectiles"`) {
		t.Error("no purchase may promise projectiles at all")
	}
}

// A coherent stat-led path is allowed: code does not require a behavior or
// access at every purchase, and the review judges its third purchase and
// capstone (decided on #27). Its third purchase needs a benefit of its own
// and its capstone a distinct capability, which may be a proposed mechanic
// (#61).
func TestAStatLedPathIsAllowed(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan := recordedOutput(t, "plan")
	for _, tier := range []string{"tier2", "tier5"} {
		at(plan, "paths", "path3", "milestones", tier).(*s.Object).Set("unlock", "none")
	}
	critical := s.NewObject().Set("name", "Critical bolt").Set("effect", "Every fifth bolt deals ten times its damage.").Set("sourceIds", []any{"source1:18"})
	at(plan, "paths", "path3", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{critical})
	// The scripted x-5-x only raises its Active Ability's numbers.
	at(plan, "paths", "path2", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{plasmaTransformation()})
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
	if !strings.Contains(request.Prompt, "Only the bottom path may have an Active Ability, first at x-x-4; the other paths stay automatic.") || strings.Contains(request.Prompt, "middle path may have an Active Ability") || strings.Contains(request.Prompt, "middle path's Active Ability") {
		t.Error("the plan prompt does not follow the Definition's manualAbilityPath")
	}
}

// The Dart Monkey Fan Club prices pay for transforming allied towers, which
// the Definition cannot express. On #27 a Luffy plan copied them for a
// self-only Gear 2 Active; the rules say once, in the reference that owns
// them, that they must not price a self Active, and point to the self
// Active references instead.
func TestAlliedActivePricesDoNotPriceASelfActive(t *testing.T) {
	rules := unit.DefaultProfile().Rules.Text
	rule := "The 7200 and 45000 prices of x-4-x and x-5-x pay mostly for transforming allied towers, which this Definition cannot express, so they must not price or shape a self-only Active Ability; the Boomerang Monkey and Tack Shooter middle paths below are the self Active references."
	for _, want := range []string{rule, "0-4-0 Super Monkey Fan Club 7200", "0-5-0 Plasma Monkey Fan Club 45000",
		"Attack speed with a self active, Boomerang Monkey middle", "0-4-0 Turbo Charge 4200", "Burst active, Tack Shooter middle"} {
		if strings.Count(rules, want) != 1 {
			t.Errorf("the default rules hold %q %d times, want once", want, strings.Count(rules, want))
		}
	}
	if strings.Index(rules, rule) > strings.Index(rules, "Attack speed with a self active") {
		t.Error("the allied-price rule does not precede the self Active references it points to")
	}
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan, err := unit.DesignPlanRequest(prepared)
	if err != nil {
		t.Fatal(err)
	}
	// The plan receives the rule through the rules document only.
	if got := strings.Count(plan.Prompt, "must not price or shape a self-only Active Ability"); got != 1 {
		t.Errorf("the plan prompt states the allied-price rule %d times", got)
	}
}

// A Luffy plan on #27 said Gear 2's speed appeared only in the Active while
// its own x-3-x made that speed permanent. The plan prompt asks that a
// path's permanent purchases and its Active develop one form.
func TestPlanPromptKeepsAFormConsistentWithItsActive(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan, err := unit.DesignPlanRequest(prepared)
	if err != nil {
		t.Fatal(err)
	}
	want := "Plan a path's permanent purchases and its Active as one form: both develop it, and the Active amplifies what the path already owns, so never describe an effect a permanent purchase grants as appearing only in the Active."
	if strings.Count(plan.Prompt, want) != 1 {
		t.Errorf("the plan prompt lacks, or repeats, %q", want)
	}
	// It sits in the Active paragraph, after the Definition's Active slot.
	paragraph := ""
	for _, p := range strings.Split(plan.Prompt, "\n\n") {
		if strings.Contains(p, want) {
			paragraph = p
		}
	}
	if !strings.HasPrefix(paragraph, "Only the middle path may have an Active Ability, first at x-4-x; the other paths stay automatic. A targeting choice is not an activation.") {
		t.Errorf("the form sentence is outside the Active paragraph: %q", paragraph)
	}
}

func TestPlanPromptAppliesEarlyIdentityOnlyWhenSelected(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	line := m.EarlyIdentityRule(prepared.Request.MechanicsDefinition)
	withPolicy, err := unit.DesignPlanRequest(prepared)
	if err != nil || !strings.Contains(withPolicy.Prompt, line) {
		t.Fatalf("early-identity guidance missing: %v", err)
	}
	definition := *prepared.Request.MechanicsDefinition
	policy := *definition.Profile.DesignPolicy
	disabled := false
	policy.PreserveEarlyAttackIdentity = &disabled
	definition.Profile.DesignPolicy = &policy
	prepared.Request.MechanicsDefinition = &definition
	withoutPolicy, err := unit.DesignPlanRequest(prepared)
	if err != nil || strings.Contains(withoutPolicy.Prompt, line) {
		t.Fatalf("early-identity guidance remained after disabling the policy: %v", err)
	}
}

// The Default states exclusiveEarlyBenefits, which subsumes
// distinctEarlyBenefits (#61); each rule is stated only under its field, and
// with both set only the exclusive one.
func TestPlanPromptStatesDistinctEarlyPurchasesOnlyWhenSelected(t *testing.T) {
	on := true
	for _, state := range []struct {
		name                string
		distinct, exclusive *bool
		want                string
	}{
		{"the Default", nil, &on, m.ExclusiveEarlyBenefitsRule},
		{"both set", &on, &on, m.ExclusiveEarlyBenefitsRule},
		{"distinct only", &on, nil, m.EarlyBenefitsRule},
		{"neither", nil, nil, ""},
	} {
		prepared, err := fixture.Prepare()
		if err != nil {
			t.Fatal(err)
		}
		definition := *prepared.Request.MechanicsDefinition
		policy := *definition.Profile.DesignPolicy
		policy.DistinctEarlyBenefits, policy.ExclusiveEarlyBenefits = state.distinct, state.exclusive
		definition.Profile.DesignPolicy = &policy
		prepared.Request.MechanicsDefinition = &definition
		request, err := unit.DesignPlanRequest(prepared)
		if err != nil {
			t.Fatal(err)
		}
		guidance := strings.Join(unit.DesignGuidance(&prepared.Request), "\n")
		for _, rule := range []string{m.EarlyBenefitsRule, m.ExclusiveEarlyBenefitsRule} {
			if want := rule == state.want; strings.Contains(request.Prompt, rule) != want || strings.Contains(guidance, rule) != want {
				t.Errorf("%s: the plan prompt or guidance states %q: %v", state.name, rule, !want)
			}
		}
	}
}

// Every BTD6 tier 3 to 5 purchase that raises damage also changes something
// else, and range is never its only companion (R1, #31). The rules document
// owns that rule and reaches the plan and the review; the review flags a
// purchase whose typed changes only raise damage, of any size, and no longer
// exempts a capstone whose damage step is not "token".
func TestAdvancedDamagePurchasesChangeSomethingElse(t *testing.T) {
	rule := "From the third purchase on, a purchase that raises damage also changes how the attack reaches or affects enemies: pierce, attack rate, projectiles, splash, a follow-up, a status, bonus damage against an enemy property, damage type, delivery or the Active Ability, as the scale references show; more range alone does not count."
	clause := "When the purchase roles ask a damage increase to come with another change, flag a third, fourth or fifth purchase whose typed changes only raise damage, of any size."
	contradiction := "flag its payoff only when it barely develops its own path over the fourth purchase, such as a token damage step"
	rules := unit.DefaultProfile().Rules.Text
	if strings.Count(rules, rule) != 1 {
		t.Errorf("the default rules hold the damage rule %d times, want once", strings.Count(rules, rule))
	}
	if strings.Contains(rules, "a small ordinary-damage step is not a capstone") {
		t.Error("the default rules keep the capstone damage-step rule the new sentence replaces")
	}
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
	plan, review := model.Requests[0].Prompt, model.Requests[2].Prompt
	for name, prompt := range map[string]string{"plan": plan, "review": review} {
		if got := strings.Count(prompt, rule); got != 1 {
			t.Errorf("the %s prompt states the damage rule %d times, want once", name, got)
		}
	}
	if !strings.Contains(review, clause) {
		t.Error("the review prompt lacks the damage-only clause")
	}
	if !strings.Contains(review, "Do not fail a capstone only because same-budget copies out-throughput it.") {
		t.Error("the review prompt lost the copy-bound exemption")
	}
	for name, prompt := range map[string]string{"plan": plan, "mechanics": model.Requests[1].Prompt, "review": review} {
		if strings.Contains(prompt, contradiction) || strings.Contains(prompt, "token damage step") {
			t.Errorf("the %s prompt still exempts a capstone with a small damage step", name)
		}
	}
}
