package mechanics

import (
	"math"
	"strconv"
	"strings"
	"testing"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// bonusDefinition is the version 2 starter with a Hardened property that
// accepts bonus damage, as the Default Profile has it, plus Blimp so a
// custom Profile can list two properties.
func bonusDefinition(properties ...string) Definition {
	d := UpgradeDefinition(extended())
	d.Vocabulary.EnemyProperties = append(d.Vocabulary.EnemyProperties, Term{ID: "hardened", Name: "Hardened", Description: "Tough."})
	for i := range d.Vocabulary.EnemyProperties {
		if d.Vocabulary.EnemyProperties[i].ID == "blimp" {
			d.Vocabulary.EnemyProperties[i].Name = "Blimp"
		}
	}
	d.Vocabulary.BonusDamageProperties = properties
	return d
}

func bonus(property, operation string, value float64) Change {
	return Change{Kind: "bonusDamage", Target: "base", Property: property, Operation: operation, Number: value}
}

// bonusStarter is the starter in version 2 form. Its third top purchase is
// Deadly Precision-like: damage set to 20 and +50 against Hardened; the
// fifth adds another +50. The first two middle purchases add a Blimp bonus
// and a Hardened one, so crosspaths sum the additions of two paths.
func bonusStarter(t *testing.T) *Blueprint {
	t.Helper()
	var blueprint Blueprint
	if err := s.ToGo(upgradeBlueprint(s.FromGoValue(starter())), &blueprint); err != nil {
		t.Fatal(err)
	}
	top := &blueprint.Paths.Path1.Tiers
	top.Tier3.Changes = append(top.Tier3.Changes, bonus("hardened", "add", 50))
	top.Tier5.Changes = append(top.Tier5.Changes, bonus("hardened", "add", 50))
	middle := &blueprint.Paths.Path2.Tiers
	middle.Tier1.Changes = append(middle.Tier1.Changes, bonus("blimp", "add", 4))
	middle.Tier2.Changes = append(middle.Tier2.Changes, bonus("hardened", "add", 10))
	return &blueprint
}

func TestBonusDamageResolvesAsASum(t *testing.T) {
	blueprint, definition := bonusStarter(t), bonusDefinition("hardened", "blimp")
	if issues := ValidateTyped(blueprint, definition); len(issues) > 0 {
		t.Fatalf("the bonus starter is invalid: %v", issues)
	}
	for selection, want := range map[Selection]string{
		{0, 0, 0}: `[]`,
		{3, 0, 0}: `[{"property":"hardened","damage":50}]`,
		{5, 0, 0}: `[{"property":"hardened","damage":100}]`,
		{0, 1, 0}: `[{"property":"blimp","damage":4}]`,
		// Two paths' additions sum, and properties are sorted by ID.
		{3, 2, 0}: `[{"property":"blimp","damage":4},{"property":"hardened","damage":60}]`,
		{5, 2, 0}: `[{"property":"blimp","damage":4},{"property":"hardened","damage":110}]`,
		{0, 2, 0}: `[{"property":"blimp","damage":4},{"property":"hardened","damage":10}]`,
	} {
		attack := ResolveUnchecked(blueprint, selection).BaseAttack
		if got := s.Stringify(s.FromGoValue(attack.BonusDamage)); got != want {
			t.Errorf("%s resolves bonus damage %s, want %s", label(selection), got, want)
		}
	}
	// The Active Ability multiplies ordinary damage only: 22 × 3, +60 added
	// unscaled.
	build := ResolveUnchecked(blueprint, Selection{3, 4, 0})
	if len(build.Abilities) != 1 || build.Abilities[0].BoostedAttack.Bonus("hardened") != 60 || build.Abilities[0].BoostedAttack.Stats.Damage != 66 {
		t.Errorf("3-4-0 boosted attack %+v", build.Abilities)
	}
	// Only an attack with bonus damage writes the field.
	plain := s.Stringify(ResolveUnchecked(blueprint, Selection{}).BaseAttack.JSONValue())
	strong := s.Stringify(ResolveUnchecked(blueprint, Selection{3, 0, 0}).BaseAttack.JSONValue())
	if strings.Contains(plain, "bonusDamage") || !strings.Contains(strong, `"bonusDamage":[{"property":"hardened","damage":50}]`) {
		t.Errorf("attack JSON: %s / %s", plain, strong)
	}
	var decoded Attack
	if err := decoded.UnmarshalJSON([]byte(strong)); err != nil || decoded.Bonus("hardened") != 50 {
		t.Errorf("the bonus does not survive a round trip: %v %+v", err, decoded)
	}
}

// Bonus damage never overrides a damage-type immunity: an enemy the damage
// type cannot hurt takes no damage, bonus included, so a bonus against it
// changes nothing and is reported like any change with no effect.
func TestBonusDamageNeverOverridesImmunity(t *testing.T) {
	definition := bonusDefinition("hardened")
	for i := range definition.Vocabulary.DamageTypes {
		if definition.Vocabulary.DamageTypes[i].ID == "sharp" {
			definition.Vocabulary.DamageTypes[i].IneffectiveAgainst = append(definition.Vocabulary.DamageTypes[i].IneffectiveAgainst, "hardened")
		}
	}
	blueprint := bonusStarter(t)
	blueprint.Paths.Path2.Tiers.Tier1.Changes = blueprint.Paths.Path2.Tiers.Tier1.Changes[:1]
	blueprint.Paths.Path2.Tiers.Tier2.Changes = blueprint.Paths.Path2.Tiers.Tier2.Changes[:1]
	// The starter's attack is sharp, so +50 against Hardened deals nothing.
	issues := ValidateTyped(blueprint, definition)
	found := false
	for _, issue := range issues {
		if issue.Path == "paths.path1.tiers.tier3.changes.2" && strings.Contains(issue.Message, "no effect") {
			found = true
		}
	}
	if !found {
		t.Errorf("a bonus the damage type cannot deal is not reported: %v", issues)
	}
	vocabulary := definition.Terms()
	attack := ResolveUnchecked(blueprint, Selection{3, 0, 0}).BaseAttack
	if got := directDamageAgainst(attack, &vocabulary, "hardened"); got != 0 {
		t.Errorf("an immune Hardened enemy takes %v damage per second", got)
	}
	attack.DamageType = "normal"
	if got := directDamageAgainst(attack, &vocabulary, "hardened"); got != 72 {
		t.Errorf("normal damage deals %v per second to Hardened, want (22 + 50) / 1", got)
	}
}

func TestBonusDamageValidation(t *testing.T) {
	definition := bonusDefinition("hardened")
	issuesOf := func(edit func(*Blueprint)) string {
		blueprint := bonusStarter(t)
		blueprint.Paths.Path2.Tiers.Tier1.Changes = blueprint.Paths.Path2.Tiers.Tier1.Changes[:1]
		edit(blueprint)
		var out []string
		for _, issue := range ValidateTyped(blueprint, definition) {
			out = append(out, issue.Path+": "+issue.Message)
		}
		return strings.Join(out, "\n")
	}
	if got := issuesOf(func(*Blueprint) {}); got != "" {
		t.Fatalf("the Hardened-only starter is invalid:\n%s", got)
	}
	for name, test := range map[string]struct {
		edit func(*Blueprint)
		want string
	}{
		"an unlisted property": {func(b *Blueprint) {
			b.Paths.Path1.Tiers.Tier3.Changes[2].Property = "blimp"
		}, "paths.path1.tiers.tier3.changes.2.property"},
		"a negative value": {func(b *Blueprint) {
			b.Paths.Path1.Tiers.Tier3.Changes[2].Number = -5
		}, "paths.path1.tiers.tier3.changes.2.value"},
		"a value above the ceiling": {func(b *Blueprint) {
			b.Paths.Path1.Tiers.Tier3.Changes[2].Number = 2000000
		}, "bonusDamage.0.damage: Exceeds the Definition stat ceiling."},
		"an attack without damage": {func(b *Blueprint) {
			b.BaseAttack.Stats.Damage = 0
			b.Paths.Path1.Tiers.Tier1.Changes[0] = stat("pierce", "add", 1)
			b.Paths.Path1.Tiers.Tier3.Changes[0] = stat("pierce", "add", 1)
			b.BaseAttack.Statuses = &[]StatusApplication{{Effect: "stun", Seconds: 1}}
		}, "Bonus damage adds to the attack's hits, so the attack must deal damage."},
		"a base bonus listed twice": {func(b *Blueprint) {
			b.BaseAttack.BonusDamage = []DamageBonus{{"hardened", 1}, {"hardened", 2}}
		}, "Each enemy property may have one bonus per attack."},
	} {
		if got := issuesOf(test.edit); !strings.Contains(got, test.want) {
			t.Errorf("%s: want %q in\n%s", name, test.want, got)
		}
	}

	// Under the early-identity policy a new bonus waits for the third
	// purchase; raising a bonus the base attack already has is allowed.
	preserve := true
	definition.Profile.DesignPolicy = &DesignPolicy{Version: "1", PreserveEarlyAttackIdentity: &preserve, MaxManualAbilityPaths: 1, Tier5Uniqueness: "one-per-player-unit-type-and-path"}
	early := issuesOf(func(b *Blueprint) {
		b.Paths.Path1.Tiers.Tier1.Changes = append(b.Paths.Path1.Tiers.Tier1.Changes, bonus("hardened", "add", 5))
	})
	if !strings.Contains(early, "paths.path1.tiers.tier1.changes: T1 and T2 improve the existing basic attack. They cannot introduce a new attack pattern, status, bonus damage or delivery") {
		t.Errorf("a first-purchase bonus is not rejected:\n%s", early)
	}
	raised := issuesOf(func(b *Blueprint) {
		b.BaseAttack.BonusDamage = []DamageBonus{{"hardened", 5}}
		b.Paths.Path1.Tiers.Tier1.Changes = append(b.Paths.Path1.Tiers.Tier1.Changes, bonus("hardened", "add", 5))
	})
	if strings.Contains(raised, "tier1.changes: T1 and T2") {
		t.Errorf("raising the base attack's bonus at the first purchase is rejected:\n%s", raised)
	}
}

// Bonus damage is an additive +N per hit (SOL-42-01 on #42): a bonusDamage
// change that multiplies or sets the bonus fails validation at its
// operation, with a message that says only add applies.
func TestBonusDamageOnlyAdds(t *testing.T) {
	definition := bonusDefinition("hardened", "blimp")
	for _, operation := range []string{"multiply", "set"} {
		blueprint := bonusStarter(t)
		blueprint.Paths.Path1.Tiers.Tier5.Changes[len(blueprint.Paths.Path1.Tiers.Tier5.Changes)-1].Operation = operation
		var found []string
		for _, issue := range ValidateTyped(blueprint, definition) {
			found = append(found, issue.Path+": "+issue.Message)
		}
		want := `paths.path1.tiers.tier5.changes.` + strconv.Itoa(len(blueprint.Paths.Path1.Tiers.Tier5.Changes)-1) + `.operation: Bonus damage only adds: use operation "add" with a positive value; multiply and set do not apply to bonus damage.`
		if !strings.Contains(strings.Join(found, "\n"), want) {
			t.Errorf("%s: want %q in\n%s", operation, want, strings.Join(found, "\n"))
		}
		if _, issues := s.Parse(ChangeSchemaV2(nil), bonus("hardened", operation, 2).JSONValue()); len(issues) != 1 || issues[0].PathString() != "operation" {
			t.Errorf("reading accepts a %s bonus: %v", operation, issues)
		}
	}
	if _, issues := s.Parse(ChangeSchemaV2(nil), bonus("hardened", "add", 2).JSONValue()); len(issues) > 0 {
		t.Errorf("an added bonus is rejected: %v", issues)
	}
	if schema := s.Stringify(s.JSONSchema(ChangeSchemaV2(bonusDefinition("hardened").Vocabulary))); !strings.Contains(schema, `"property":{"type":"string","enum":["hardened"]},"operation":{"type":"string","enum":["add"]}`) {
		t.Errorf("the change schema offers another bonus operation: %s", schema)
	}
}

// Definitions without bonus damage properties, version 1 included, keep
// their schemas: no bonusDamage change or attack field exists for them.
func TestDefinitionsWithoutBonusPropertiesAreUnchanged(t *testing.T) {
	blueprint := bonusStarter(t)
	for name, definition := range map[string]Definition{"version 2 without the list": bonusDefinition()} {
		issues := ValidateTyped(blueprint, definition)
		if len(issues) == 0 || !strings.Contains(issues[0].Path, "changes") {
			t.Errorf("%s accepts bonus damage: %v", name, issues)
		}
	}
	if _, issues := s.Parse(ChangeSchema, bonus("hardened", "add", 1).JSONValue()); len(issues) == 0 {
		t.Error("the version 1 change schema accepts bonus damage")
	}
	offers := func(v *Vocabulary) bool {
		return strings.Contains(s.Stringify(s.JSONSchema(AttackSchemaV2(v))), `"bonusDamage"`)
	}
	if offers(bonusDefinition().Vocabulary) || !offers(nil) {
		t.Error("the attack schema offers bonusDamage without a listed property, or refuses it when reading")
	}
	d := bonusDefinition()
	if strings.Contains(s.Stringify(d.JSONValue()), "bonusDamageProperties") {
		t.Error("a Definition without bonus damage properties writes the field")
	}
	vocabulary := d.Terms()
	if PurchaseMetricsWith(ResolveUnchecked(blueprint, Selection{3, 0, 0}), &vocabulary).Has("direct damage rate against Hardened") {
		t.Error("a Definition without bonus damage properties measures bonus damage")
	}
}

func TestBonusDamageVocabularyIssues(t *testing.T) {
	d := bonusDefinition("hardened", "ceramic", "hardened")
	var got []string
	for _, issue := range VocabularyIssues(*d.Vocabulary) {
		got = append(got, issue.Path+": "+issue.Message)
	}
	want := "vocabulary.bonusDamageProperties.1: ceramic is not a declared enemy property.\nvocabulary.bonusDamageProperties.2: hardened is already listed."
	if strings.Join(got, "\n") != want {
		t.Errorf("issues:\n%s\nwant:\n%s", strings.Join(got, "\n"), want)
	}
}

// Purchase metrics add the direct damage rate against each listed property;
// the other rates stay bonus-free.
func TestBonusDamageMetrics(t *testing.T) {
	blueprint, definition := bonusStarter(t), bonusDefinition("hardened")
	vocabulary := definition.Terms()
	metrics := PurchaseMetricsWith(ResolveUnchecked(blueprint, Selection{3, 0, 0}), &vocabulary)
	against, _ := metrics.Get("direct damage rate against Hardened")
	direct, _ := metrics.Get("direct damage rate")
	if against != 72.0 || direct != 22.0 {
		t.Errorf("3-0-0 direct %v, against Hardened %v; want 22 and 72", direct, against)
	}
	// With Hardened and Blimp listed, each rate counts only its own
	// property's bonus: 3-2-0 has +60 against Hardened and +4 against Blimp,
	// and 3-0-0's Blimp rate is its ordinary direct rate.
	both := bonusDefinition("hardened", "blimp").Terms()
	metrics = PurchaseMetricsWith(ResolveUnchecked(blueprint, Selection{3, 0, 0}), &both)
	if blimp, _ := metrics.Get("direct damage rate against Blimp"); blimp != 22.0 {
		t.Errorf("3-0-0 against Blimp %v; want 22", blimp)
	}
	build := ResolveUnchecked(blueprint, Selection{3, 2, 0})
	metrics = PurchaseMetricsWith(build, &both)
	direct, _ = metrics.Get("direct damage rate")
	hardened, _ := metrics.Get("direct damage rate against Hardened")
	blimp, _ := metrics.Get("direct damage rate against Blimp")
	perHit := build.BaseAttack.Stats.Projectiles / build.BaseAttack.Stats.IntervalSeconds
	if math.Abs(hardened.(float64)-direct.(float64)-60*perHit) > 1e-9 || math.Abs(blimp.(float64)-direct.(float64)-4*perHit) > 1e-9 {
		t.Errorf("3-2-0 direct %v, against Hardened %v, against Blimp %v", direct, hardened, blimp)
	}
	before := ResolveUnchecked(blueprint, Selection{2, 0, 0})
	after := ResolveUnchecked(blueprint, Selection{3, 0, 0})
	if !HasBehaviorTransition(before, after) {
		t.Error("a new bonus is not a behavior transition")
	}
}
