package unit_test

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/render"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// The Default Profile adds Hardened, adapted from Ceramic, and lists it and
// Blimp, adapted from the MOAB-class Moabs tag, for bonus damage as two
// separate properties: a Blimp is not Hardened (SOL-42-02 on #42).
func TestDefaultProfileListsHardenedAndBlimpForBonusDamage(t *testing.T) {
	profile := unit.DefaultProfile()
	vocabulary := profile.MechanicsDefinition.Vocabulary
	var ids []string
	for _, term := range vocabulary.EnemyProperties {
		ids = append(ids, term.ID)
	}
	if strings.Join(ids, ",") != "lead,frozen,purple,black,zebra,hardened,blimp,boss" {
		t.Errorf("enemy properties %v", ids)
	}
	if strings.Join(vocabulary.BonusDamageProperties, ",") != "hardened,blimp" || vocabulary.PropertyName("hardened") != "Hardened" || vocabulary.PropertyName("blimp") != "Blimp" {
		t.Errorf("bonus damage properties %v", vocabulary.BonusDamageProperties)
	}
	for _, term := range vocabulary.EnemyProperties {
		if (term.ID == "hardened" && !strings.Contains(term.Description, "it is not a Blimp")) || (term.ID == "blimp" && !strings.Contains(term.Description, "it is not Hardened")) {
			t.Errorf("%s reads %q", term.ID, term.Description)
		}
	}
	for _, damageType := range vocabulary.DamageTypes {
		for _, property := range vocabulary.BonusDamageProperties {
			if slices.Contains(damageType.IneffectiveAgainst, property) {
				t.Errorf("%s cannot hurt %s", damageType.ID, property)
			}
		}
	}
	if profile.Rules.ID != "default-td-profile-v37" || profile.MechanicsDefinition.Revision != "2026-09-27-atlas-56.3-v37" {
		t.Errorf("rules %s, Definition %s", profile.Rules.ID, profile.MechanicsDefinition.Revision)
	}
	for _, want := range []string{
		// Deadly Precision returns as a qualifying third purchase with its bonus.
		"Deadly Precision nearly triples damage and adds +50 damage against Ceramic",
		"0-3-0 Bionic Boomerang 1250: every 0.15 s, 8 times the base frequency, and +1 damage to MOAB-class enemies.",
		"This Profile adapts a bonus against Ceramic as bonus damage against Hardened and a bonus against MOAB-class enemies, such as Bionic Boomerang's +1, as bonus damage against Blimp. The two are separate properties: a Blimp is not Hardened, and a bonus against one never applies to the other. Bonuses against Fortified, Lead or other classes have no counterpart here.",
	} {
		if !strings.Contains(profile.Rules.Text, want) {
			t.Errorf("the default rules lack %q", want)
		}
	}
	if issues := m.VocabularyIssues(*vocabulary); len(issues) > 0 {
		t.Errorf("vocabulary issues %v", issues)
	}
}

// bonusAddition is a bonus a test adds to one recorded purchase, with the
// plan's promise to match.
type bonusAddition struct {
	path, tier, property string
	value                float64
}

// juggernautHardened is Juggernaut, 4-x-x, dealing +3 damage against
// Hardened, as its atlas file (DartMonkey-400) has +3 to Ceramic.
var juggernautHardened = bonusAddition{"path1", "tier4", "hardened", 3}

// pressBlimp is the fourth bottom purchase, x-x-4, dealing +4 damage against
// Blimp, as MOAB Press (BoomerangMonkey-004) has +4 to Moabs.
var pressBlimp = bonusAddition{"path3", "tier4", "blimp", 4}

// bonusOutputs are the fixture's recorded outputs with each addition's
// purchase promising and dealing its bonus damage.
func bonusOutputs(t *testing.T, additions ...bonusAddition) (plan, mechanics, review *s.Object) {
	plan, mechanics, review = recordedOutput(t, "plan"), recordedOutput(t, "mechanics"), recordedOutput(t, "review")
	path := func(root *s.Object, keys ...string) *s.Object {
		node := root
		for _, key := range keys {
			value, _ := node.Get(key)
			node = value.(*s.Object)
		}
		return node
	}
	for _, addition := range additions {
		milestone := path(plan, "paths", addition.path, "milestones", addition.tier)
		improves, _ := milestone.Get("improves")
		milestone.Set("improves", append(improves.([]any), unit.BonusDamagePromise))
		tier := path(mechanics, "paths", addition.path, "tiers", addition.tier)
		tier.Set("bonusDamage", []any{s.NewObject().Set("property", addition.property).Set("operation", "add").Set("value", addition.value)})
	}
	return plan, mechanics, review
}

// bonusRun drafts, checks and reviews the fixture with the additions. It
// fails the test on any failing Finding.
func bonusRun(t *testing.T, additions ...bonusAddition) (unit.Draft, m.Definition, *fixture.Model) {
	t.Helper()
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan, mechanics, review := bonusOutputs(t, additions...)
	model := &fixture.Model{Outputs: []any{plan, mechanics, review}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
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
	if _, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options()); err != nil {
		t.Fatal(err)
	}
	return draft, *prepared.Request.MechanicsDefinition, model
}

// reviewBonuses maps each legal build code to its review facts'
// attack.bonusDamage, for builds that have one.
func reviewBonuses(blueprint *m.Blueprint, definition m.Definition) map[string]string {
	out := map[string]string{}
	for _, entry := range unit.LegalBuildFacts(blueprint, definition) {
		code, _ := entry.(*s.Object).Get("code")
		attack, _ := entry.(*s.Object).Get("attack")
		if bonus, ok := attack.(*s.Object).Get("bonusDamage"); ok {
			out[code.(string)] = s.Stringify(bonus)
		}
	}
	return out
}

func TestHardenedBonusRunsThroughThePipeline(t *testing.T) {
	draft, definition, model := bonusRun(t, juggernautHardened)
	blueprint := draft.Candidate.Blueprint
	changes := blueprint.Paths.Path1.Tiers.Tier4.Changes
	at := slices.IndexFunc(changes, func(c m.Change) bool { return c.Kind == "bonusDamage" })
	if at < 0 || changes[at].Property != "hardened" || changes[at].Operation != "add" || changes[at].Number != 3 {
		t.Fatalf("Juggernaut decodes %+v", changes)
	}
	if got := m.ResolveUnchecked(blueprint, m.Selection{4, 0, 0}).BaseAttack.Bonus("hardened"); got != 3 {
		t.Errorf("4-0-0 deals +%v against Hardened", got)
	}
	// Purchase evidence measures the bonus; the other rates stay bonus-free.
	evidence := s.Stringify(draft.Run.DesignEvaluation)
	if !strings.Contains(evidence, `"direct damage rate against Hardened"`) || !strings.Contains(evidence, "bonusDamage hardened: 0 → 3") ||
		!strings.Contains(evidence, "Every other rate, the capstone copy bounds and the design policy ignore bonus damage.") {
		t.Error("purchase evidence lacks the Hardened rate, its change or its limitation")
	}
	planPrompt, mechanicsPrompt, reviewPrompt := model.Requests[0].Prompt, model.Requests[1].Prompt, model.Requests[2].Prompt
	for name, prompt := range map[string]string{"plan": planPrompt, "mechanics": mechanicsPrompt} {
		if !strings.Contains(prompt, "Bonus damage is typed: an attack may deal +N damage per hit to enemies with a property definition.vocabulary.bonusDamageProperties lists") ||
			!strings.Contains(prompt, "Each listed property has its own bonus: a bonus against one property never applies to another") {
			t.Errorf("the %s prompt does not explain bonus damage", name)
		}
	}
	// The review reads the resolved bonus and is told to hold a claimed
	// class bonus to it.
	if !strings.Contains(reviewPrompt, "keeps that claim only when its builds' attack.bonusDamage shows a bonus against that property") ||
		!strings.Contains(reviewPrompt, "Each property's bonus is separate: a bonus against one property never counts for another.") {
		t.Error("the review prompt cannot check a claimed class bonus")
	}
	bonuses := reviewBonuses(blueprint, definition)
	for _, entry := range unit.LegalBuildFacts(blueprint, definition) {
		code, _ := entry.(*s.Object).Get("code")
		bonus, ok := bonuses[code.(string)]
		want := strings.HasPrefix(code.(string), "4-") || strings.HasPrefix(code.(string), "5-")
		if want != ok || (ok && bonus != `{"hardened":3}`) {
			t.Errorf("%s review facts bonusDamage %v", code, bonus)
		}
	}
}

// A Blimp bonus, MOAB Press's +4 to Moabs adapted, runs through the same
// pipeline as a Hardened one, with its own rate in the purchase evidence.
func TestBlimpBonusRunsThroughThePipeline(t *testing.T) {
	draft, definition, _ := bonusRun(t, pressBlimp)
	blueprint := draft.Candidate.Blueprint
	if got := m.ResolveUnchecked(blueprint, m.Selection{0, 0, 4}).BaseAttack.Bonus("blimp"); got != 4 {
		t.Errorf("0-0-4 deals +%v against Blimp", got)
	}
	evidence := s.Stringify(draft.Run.DesignEvaluation)
	if !strings.Contains(evidence, `"direct damage rate against Blimp"`) || !strings.Contains(evidence, "bonusDamage blimp: 0 → 4") ||
		strings.Contains(evidence, "bonusDamage hardened") {
		t.Error("purchase evidence lacks the Blimp rate or its change, or shows a Hardened bonus")
	}
	bonuses := reviewBonuses(blueprint, definition)
	if bonuses["0-0-4"] != `{"blimp":4}` || bonuses["0-0-5"] != `{"blimp":4}` {
		t.Errorf("review facts bonusDamage %v", bonuses)
	}
	for code, bonus := range bonuses {
		if !strings.HasSuffix(code, "-4") && !strings.HasSuffix(code, "-5") || bonus != `{"blimp":4}` {
			t.Errorf("%s review facts bonusDamage %v", code, bonus)
		}
	}
}

// One Unit may buy both: Juggernaut's Hardened bonus on the top path and
// MOAB Press's Blimp bonus on the bottom path. They stay separate: each
// build's review facts, sheet and rates show only the bonus it bought, and
// each property's rate counts only its own bonus (SOL-42-02).
func TestHardenedAndBlimpBonusesStaySeparate(t *testing.T) {
	draft, definition, _ := bonusRun(t, juggernautHardened, pressBlimp)
	blueprint := draft.Candidate.Blueprint
	bonuses := reviewBonuses(blueprint, definition)
	if bonuses["4-0-0"] != `{"hardened":3}` || bonuses["4-0-2"] != `{"hardened":3}` || bonuses["0-0-4"] != `{"blimp":4}` || bonuses["2-0-4"] != `{"blimp":4}` || bonuses["0-4-0"] != "" {
		t.Errorf("review facts bonusDamage %v", bonuses)
	}
	evidence := s.Stringify(draft.Run.DesignEvaluation)
	if !strings.Contains(evidence, "bonusDamage hardened: 0 → 3") || !strings.Contains(evidence, "bonusDamage blimp: 0 → 4") ||
		!strings.Contains(evidence, "Each of direct damage rate against Hardened and direct damage rate against Blimp counts only its own property's bonus damage") {
		t.Error("purchase evidence lacks a bonus change or the per-property limitation")
	}
	vocabulary := definition.Terms()
	for _, tc := range []struct {
		selection       m.Selection
		hardened, blimp float64
	}{{m.Selection{4, 0, 0}, 3, 0}, {m.Selection{0, 0, 4}, 0, 4}} {
		build := m.ResolveUnchecked(blueprint, tc.selection)
		attack := build.BaseAttack
		metrics := m.PurchaseMetricsWith(build, &vocabulary)
		direct, _ := metrics.Get("direct damage rate")
		hardened, _ := metrics.Get("direct damage rate against Hardened")
		blimp, _ := metrics.Get("direct damage rate against Blimp")
		perBonus := attack.Stats.Projectiles / attack.Stats.IntervalSeconds
		if attack.Distribution == "distinct-targets" {
			perBonus = 1 / attack.Stats.IntervalSeconds
		}
		if !near(hardened.(float64)-direct.(float64), tc.hardened*perBonus) || !near(blimp.(float64)-direct.(float64), tc.blimp*perBonus) {
			t.Errorf("%v rates: direct %v, Hardened %v, Blimp %v", tc.selection, direct, hardened, blimp)
		}
	}
	// The sheet names each bonus with its own property.
	sheet := render.Purchases(draft.Candidate, &definition, nil)
	top, bottom := strings.Join(sheet[0].Purchases[3].Effects, " "), strings.Join(sheet[2].Purchases[3].Effects, " ")
	if !strings.Contains(top, "+3 damage against Hardened enemies") || strings.Contains(top, "Blimp") ||
		!strings.Contains(bottom, "+4 damage against Blimp enemies") || strings.Contains(bottom, "Hardened") {
		t.Errorf("4-x-x reads %q; x-x-4 reads %q", top, bottom)
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

	plan, _, _ := bonusOutputs(t, juggernautHardened)
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
	if err == nil || !strings.Contains(err.Error(), "1-x-x promises to unlock bonus-damage. The first and second purchase of each path keep the base attack's form: they add no new status, bonus damage, splash") {
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
	_, mechanics, _ := bonusOutputs(t, juggernautHardened)
	value, _ := mechanics.Get("paths")
	value, _ = value.(*s.Object).Get("path1")
	value, _ = value.(*s.Object).Get("tiers")
	value, _ = value.(*s.Object).Get("tier4")
	value.(*s.Object).Set("bonusDamage", []any{s.NewObject().Set("property", "boss").Set("operation", "add").Set("value", 3.0)})
	if _, err := unit.DecodeBlueprintOutput(mechanics, &request); err == nil {
		t.Error("a bonus against Boss decodes under the Default Profile")
	} else if !strings.Contains(err.Error(), `paths.path1.tiers.tier4.bonusDamage.0.property: Invalid option: expected one of "hardened"|"blimp"`) {
		t.Errorf("an unlisted property fails unclearly: %v", err)
	}
	// The wire offers only add, and multiply or set fail at the operation.
	if !strings.Contains(s.Stringify(schema), `"operation":{"type":"string","enum":["add"]}`) {
		t.Error("the wire offers another bonus damage operation")
	}
	for _, operation := range []string{"multiply", "set"} {
		value.(*s.Object).Set("bonusDamage", []any{s.NewObject().Set("property", "hardened").Set("operation", operation).Set("value", 2.0)})
		_, err := unit.DecodeBlueprintOutput(mechanics, &request)
		if err == nil || !strings.Contains(err.Error(), "paths.path1.tiers.tier4.bonusDamage.0.operation") || !strings.Contains(err.Error(), `Bonus damage only adds: use operation "add"`) {
			t.Errorf("a %s bonus decodes or fails unclearly: %v", operation, err)
		}
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
