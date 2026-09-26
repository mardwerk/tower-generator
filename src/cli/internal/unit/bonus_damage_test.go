package unit_test

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// The Default Profile adds Hardened, adapted from Ceramic, and lists only
// it for bonus damage: a Blimp is not Hardened (SOL-31-25 on #31, #42).
func TestDefaultProfileListsOnlyHardenedForBonusDamage(t *testing.T) {
	profile := unit.DefaultProfile()
	vocabulary := profile.MechanicsDefinition.Vocabulary
	var ids []string
	for _, term := range vocabulary.EnemyProperties {
		ids = append(ids, term.ID)
	}
	if strings.Join(ids, ",") != "lead,frozen,purple,black,zebra,hardened,blimp,boss" {
		t.Errorf("enemy properties %v", ids)
	}
	if strings.Join(vocabulary.BonusDamageProperties, ",") != "hardened" || vocabulary.PropertyName("hardened") != "Hardened" {
		t.Errorf("bonus damage properties %v", vocabulary.BonusDamageProperties)
	}
	for _, damageType := range vocabulary.DamageTypes {
		if slices.Contains(damageType.IneffectiveAgainst, "hardened") {
			t.Errorf("%s cannot hurt Hardened", damageType.ID)
		}
	}
	if profile.Rules.ID != "default-td-profile-v22" || profile.MechanicsDefinition.Revision != "2026-09-27-atlas-56.3-v22" {
		t.Errorf("rules %s, Definition %s", profile.Rules.ID, profile.MechanicsDefinition.Revision)
	}
	for _, want := range []string{
		// Deadly Precision returns as a qualifying third purchase with its bonus.
		"Deadly Precision nearly triples damage and adds +50 damage against Ceramic",
		"This Profile adapts a bonus against Ceramic as bonus damage against Hardened; bonuses against MOAB-class, Fortified, Lead or other classes have no counterpart here.",
	} {
		if !strings.Contains(profile.Rules.Text, want) {
			t.Errorf("the default rules lack %q", want)
		}
	}
	if issues := m.VocabularyIssues(*vocabulary); len(issues) > 0 {
		t.Errorf("vocabulary issues %v", issues)
	}
}

// bonusOutputs are the fixture's recorded outputs with Juggernaut, 4-x-x,
// promising and dealing +3 damage against Hardened, as its atlas file has
// +3 to Ceramic.
func bonusOutputs(t *testing.T) (plan, mechanics, review *s.Object) {
	plan, mechanics, review = recordedOutput(t, "plan"), recordedOutput(t, "mechanics"), recordedOutput(t, "review")
	path := func(root *s.Object, keys ...string) *s.Object {
		node := root
		for _, key := range keys {
			value, _ := node.Get(key)
			node = value.(*s.Object)
		}
		return node
	}
	milestone := path(plan, "paths", "path1", "milestones", "tier4")
	improves, _ := milestone.Get("improves")
	milestone.Set("improves", append(improves.([]any), unit.BonusDamagePromise))
	tier := path(mechanics, "paths", "path1", "tiers", "tier4")
	tier.Set("bonusDamage", []any{s.NewObject().Set("property", "hardened").Set("operation", "add").Set("value", 3.0)})
	return plan, mechanics, review
}

func TestHardenedBonusRunsThroughThePipeline(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan, mechanics, review := bonusOutputs(t)
	model := &fixture.Model{Outputs: []any{plan, mechanics, review}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	blueprint := draft.Candidate.Blueprint
	changes := blueprint.Paths.Path1.Tiers.Tier4.Changes
	at := slices.IndexFunc(changes, func(c m.Change) bool { return c.Kind == "bonusDamage" })
	if at < 0 || changes[at].Property != "hardened" || changes[at].Operation != "add" || changes[at].Number != 3 {
		t.Fatalf("Juggernaut decodes %+v", changes)
	}
	definition := *prepared.Request.MechanicsDefinition
	if got := m.ResolveUnchecked(blueprint, m.Selection{4, 0, 0}).BaseAttack.Bonus("hardened"); got != 3 {
		t.Errorf("4-0-0 deals +%v against Hardened", got)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	for _, finding := range checked.Findings {
		if finding.Outcome == "fail" {
			t.Errorf("finding %s: %s", finding.Subject, finding.Message)
		}
	}
	// Purchase evidence measures the bonus; the other rates stay bonus-free.
	evidence := s.Stringify(draft.Run.DesignEvaluation)
	if !strings.Contains(evidence, `"direct damage rate against Hardened"`) || !strings.Contains(evidence, "bonusDamage hardened: 0 → 3") ||
		!strings.Contains(evidence, "Every other rate, the capstone copy bounds and the design policy ignore bonus damage.") {
		t.Error("purchase evidence lacks the Hardened rate, its change or its limitation")
	}
	if _, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options()); err != nil {
		t.Fatal(err)
	}
	planPrompt, mechanicsPrompt, reviewPrompt := model.Requests[0].Prompt, model.Requests[1].Prompt, model.Requests[2].Prompt
	for name, prompt := range map[string]string{"plan": planPrompt, "mechanics": mechanicsPrompt} {
		if !strings.Contains(prompt, "Bonus damage is typed: an attack may deal +N damage per hit to enemies with a property definition.vocabulary.bonusDamageProperties lists") {
			t.Errorf("the %s prompt does not explain bonus damage", name)
		}
	}
	// The review reads the resolved bonus and is told to hold a claimed
	// class bonus to it.
	if !strings.Contains(reviewPrompt, "keeps that claim only when its builds' attack.bonusDamage shows a bonus against that property") {
		t.Error("the review prompt cannot check a claimed class bonus")
	}
	for _, entry := range unit.LegalBuildFacts(blueprint, definition) {
		code, _ := entry.(*s.Object).Get("code")
		attack, _ := entry.(*s.Object).Get("attack")
		bonus, ok := attack.(*s.Object).Get("bonusDamage")
		want := code == "4-0-0" || code == "5-0-0" || strings.HasPrefix(code.(string), "4-") || strings.HasPrefix(code.(string), "5-")
		if want != ok || (ok && s.Stringify(bonus) != `{"hardened":3}`) {
			t.Errorf("%s review facts bonusDamage %v", code, bonus)
		}
	}
}

// A plan that promises bonus damage holds the purchase to it, and under
// the early-identity policy a first or second purchase cannot unlock it.
func TestBonusDamagePromise(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	definition := *prepared.Request.MechanicsDefinition
	if !slices.Contains(unit.ImprovementsFor(&definition), unit.BonusDamagePromise) || !slices.Contains(unit.UnlocksFor(&definition), unit.BonusDamagePromise) {
		t.Error("the Default Profile's plans cannot promise bonus-damage")
	}
	without := definition
	vocabulary := *definition.Vocabulary
	vocabulary.BonusDamageProperties = nil
	without.Vocabulary = &vocabulary
	if slices.Contains(unit.ImprovementsFor(&without), unit.BonusDamagePromise) || slices.Contains(unit.UnlocksFor(&without), unit.BonusDamagePromise) {
		t.Error("a Definition without bonus damage properties offers the bonus-damage promise")
	}

	plan, _, _ := bonusOutputs(t)
	decoded, err := unit.DecodeDesignPlan(plan, &prepared.Request)
	if err != nil {
		t.Fatal(err)
	}
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	// The recorded Juggernaut lacks the bonus the plan now promises.
	issues := unit.PlanIntentIssues(*stages.Result.Candidate.Blueprint, decoded.UpgradeIntents, definition)
	if len(issues) != 1 || issues[0].Path != "paths.path1.tiers.tier4.planIntent" || !strings.Contains(issues[0].Message, "improved bonus-damage") ||
		!strings.Contains(issues[0].Message, "Add a bonusDamage entry that raises the bonus damage") {
		t.Errorf("promise issues %+v", issues)
	}

	early := recordedOutput(t, "plan")
	value, _ := early.Get("paths")
	value, _ = value.(*s.Object).Get("path1")
	value, _ = value.(*s.Object).Get("milestones")
	value, _ = value.(*s.Object).Get("tier1")
	value.(*s.Object).Set("unlock", unit.BonusDamagePromise)
	_, err = unit.DecodeDesignPlan(early, &prepared.Request)
	if err == nil || !strings.Contains(err.Error(), "1-x-x must preserve the existing attack identity: the first and second purchase of a path add no new status, bonus damage, attack pattern") {
		t.Errorf("a first-purchase bonus-damage unlock is not rejected: %v", err)
	}
}

// The mechanics wire offers bonusDamage only under a vocabulary that lists
// bonus damage properties, and holds it to them.
func TestBonusDamageWire(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	request := prepared.Request
	schema, err := unit.ModelOutputJSONSchema(&request)
	if err != nil || !strings.Contains(s.Stringify(schema), `"bonusDamage"`) {
		t.Fatalf("the Default wire lacks bonusDamage: %v", err)
	}
	_, mechanics, _ := bonusOutputs(t)
	value, _ := mechanics.Get("paths")
	value, _ = value.(*s.Object).Get("path1")
	value, _ = value.(*s.Object).Get("tiers")
	value, _ = value.(*s.Object).Get("tier4")
	value.(*s.Object).Set("bonusDamage", []any{s.NewObject().Set("property", "blimp").Set("operation", "add").Set("value", 3.0)})
	if _, err := unit.DecodeBlueprintOutput(mechanics, &request); err == nil {
		t.Error("a bonus against Blimp decodes under the Default Profile")
	}
	vocabulary := *request.MechanicsDefinition.Vocabulary
	vocabulary.BonusDamageProperties = nil
	definition := *request.MechanicsDefinition
	definition.Vocabulary = &vocabulary
	request.MechanicsDefinition = &definition
	if schema, err := unit.ModelOutputJSONSchema(&request); err != nil || strings.Contains(s.Stringify(schema), `"bonusDamage"`) {
		t.Errorf("a Definition without bonus damage properties offers bonusDamage: %v", err)
	}
	if strings.Contains(strings.Join(unit.VocabularyGuidance(&request), " "), "Bonus damage is typed") {
		t.Error("a Definition without bonus damage properties is told about bonus damage")
	}
}
