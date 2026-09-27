package render_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/render"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// hardenedJuggernaut is the fixture unit with Juggernaut, 4-x-x, adding +3
// damage against Hardened, as its atlas file adds +3 to Ceramic.
func hardenedJuggernaut(t *testing.T) (unit.Candidate, *mechanics.Definition) {
	t.Helper()
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	candidate := stages.Result.Candidate
	blueprint := *candidate.Blueprint
	tier := &blueprint.Paths.Path1.Tiers.Tier4
	tier.Changes = append(append([]mechanics.Change(nil), tier.Changes...), mechanics.Change{Kind: "bonusDamage", Target: "base", Property: "hardened", Operation: "add", Number: 3})
	candidate.Blueprint = &blueprint
	return candidate, stages.Result.Prepared.Request.MechanicsDefinition
}

func TestBonusDamageRendersOnTheSheet(t *testing.T) {
	candidate, definition := hardenedJuggernaut(t)
	purchases := render.Purchases(candidate, definition, nil)
	if purchases == nil {
		t.Fatal("the sheet does not render")
	}
	juggernaut := strings.Join(purchases[0].Purchases[3].Effects, " ")
	if !strings.Contains(juggernaut, "Adds +3 damage against Hardened enemies (2 to 5 per hit).") {
		t.Errorf("4-x-x reads %q", juggernaut)
	}
	stats := render.Stats(candidate, definition)
	found := false
	for _, change := range stats.Tiers[render.TierKey("path-1", 4)].Changes {
		if change.Key == "bonusDamage.hardened" {
			found = change.Label == "Damage against Hardened" && change.Kind == render.BonusDamageKind && change.After == 3.0 && change.Before == nil
		}
	}
	if !found {
		t.Errorf("the kit stats lack the Hardened bonus: %+v", stats.Tiers[render.TierKey("path-1", 4)].Changes)
	}
	crosspaths := render.ResolveCrosspaths(candidate, definition)
	for _, row := range crosspaths.Advanced {
		if row.Code == "4-1-0" && !strings.Contains(row.Attack, "+3 damage against Hardened enemies") {
			t.Errorf("4-1-0 attack %q", row.Attack)
		}
	}

	// A bonus on the base attack is described with its damage per hit, and
	// a later raise as a number change.
	blueprint := *candidate.Blueprint
	blueprint.BaseAttack.BonusDamage = []mechanics.DamageBonus{{Property: "hardened", Damage: 1}}
	candidate.Blueprint = &blueprint
	base := render.Base(candidate, definition)
	if base == nil || !strings.Contains(base.Text, "Each hit deals +1 damage against Hardened enemies (2 per hit).") {
		t.Errorf("0-0-0 reads %+v", base)
	}
	juggernaut = strings.Join(render.Purchases(candidate, definition, nil)[0].Purchases[3].Effects, " ")
	if !strings.Contains(juggernaut, "Raises bonus damage against Hardened enemies from 1 to 4 (+3).") {
		t.Errorf("4-x-x reads %q", juggernaut)
	}

	// The bonus is added to every hit unscaled: a follow-up hit takes its
	// multiple of ordinary damage plus the bonus, and the Active Ability
	// multiplies ordinary damage only.
	sheet := render.Purchases(candidate, definition, nil)
	split := strings.Join(sheet[0].Purchases[4].Effects, " ")
	if !strings.Contains(split, "take 0.4 times the hit damage (2), plus +4 damage against Hardened enemies, once each") {
		t.Errorf("5-x-x reads %q", split)
	}
	active := strings.Join(sheet[1].Purchases[3].Effects, " ")
	if !strings.Contains(active, "so the attack deals 1 damage, plus +1 damage against Hardened enemies, every 0.015 s") {
		t.Errorf("x-4-x reads %q", active)
	}
}

// A bonus never overrides an immunity; the sheet says when the damage type
// cannot deal it.
func TestBonusDamageSheetNamesAnImmunity(t *testing.T) {
	candidate, definition := hardenedJuggernaut(t)
	immune := *definition
	vocabulary := *definition.Vocabulary
	vocabulary.DamageTypes = append([]mechanics.DamageType(nil), vocabulary.DamageTypes...)
	for i := range vocabulary.DamageTypes {
		if vocabulary.DamageTypes[i].ID == "normal" {
			vocabulary.DamageTypes[i].IneffectiveAgainst = []string{"hardened"}
		}
	}
	immune.Vocabulary = &vocabulary
	juggernaut := strings.Join(render.Purchases(candidate, &immune, nil)[0].Purchases[3].Effects, " ")
	if !strings.Contains(juggernaut, "Adds +3 damage against Hardened enemies, which its Normal damage cannot deal.") {
		t.Errorf("4-x-x reads %q", juggernaut)
	}
}

// Results saved before bonus damage existed have no bonusDamage field and
// no bonus damage properties. They still load, resolve and render, and
// nothing on their sheets mentions bonus damage; that they render exactly
// as committed is TestReferenceCapturesRenderAsCommitted.
func TestSavedResultsWithoutBonusDamageStillRender(t *testing.T) {
	files, _ := filepath.Glob("../../../../data/reference/captures/*.json")
	if len(files) == 0 {
		t.Fatal("no reference captures")
	}
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(data), "bonusDamage") {
			t.Fatalf("%s already has bonus damage", filepath.Base(file))
		}
		value, err := s.Decode(data)
		if err != nil {
			t.Fatal(err)
		}
		result, err := unit.ParseResult(value)
		if err != nil {
			t.Fatalf("%s: %v", filepath.Base(file), err)
		}
		candidate, definition := result.Candidate, result.Prepared.Request.MechanicsDefinition
		if render.Stats(candidate, definition) == nil || render.Purchases(candidate, definition, nil) == nil || render.ResolveCrosspaths(candidate, definition) == nil {
			t.Errorf("%s lost its sheet", filepath.Base(file))
		}
		compact, err := render.Markdown(value, false)
		if err != nil || strings.Contains(compact, "Hardened") || strings.Contains(compact, "bonus damage") {
			t.Errorf("%s renders bonus damage (err %v)", filepath.Base(file), err)
		}
		if definition != nil && strings.Contains(s.Stringify(definition.JSONValue()), "bonusDamageProperties") {
			t.Errorf("%s gained bonus damage properties on reading", filepath.Base(file))
		}
	}
}

// Hardened and Blimp bonuses stay separate on the sheet, in the kit stats
// and on crosspath rows: each names its own property, and one attack may
// show both (SOL-42-02).
func TestHardenedAndBlimpBonusesRenderApart(t *testing.T) {
	candidate, definition := hardenedJuggernaut(t)
	blueprint := *candidate.Blueprint
	tier := &blueprint.Paths.Path3.Tiers.Tier4
	tier.Changes = append(append([]mechanics.Change(nil), tier.Changes...), mechanics.Change{Kind: "bonusDamage", Target: "base", Property: "blimp", Operation: "add", Number: 4})
	candidate.Blueprint = &blueprint
	sheet := render.Purchases(candidate, definition, nil)
	juggernaut, press := strings.Join(sheet[0].Purchases[3].Effects, " "), strings.Join(sheet[2].Purchases[3].Effects, " ")
	if !strings.Contains(juggernaut, "Adds +3 damage against Hardened enemies (2 to 5 per hit).") || strings.Contains(juggernaut, "Blimp") {
		t.Errorf("4-x-x reads %q", juggernaut)
	}
	if !strings.Contains(press, "Adds +4 damage against Blimp enemies (6 to 10 per hit).") || strings.Contains(press, "Hardened") {
		t.Errorf("x-x-4 reads %q", press)
	}
	stats := render.Stats(candidate, definition)
	for _, tc := range []struct{ tier, key, other, label string }{
		{render.TierKey("path-1", 4), "bonusDamage.hardened", "bonusDamage.blimp", "Damage against Hardened"},
		{render.TierKey("path-3", 4), "bonusDamage.blimp", "bonusDamage.hardened", "Damage against Blimp"},
	} {
		found := false
		for _, change := range stats.Tiers[tc.tier].Changes {
			if change.Key == tc.key {
				found = change.Label == tc.label && change.Kind == render.BonusDamageKind
			}
			if change.Key == tc.other {
				t.Errorf("%s shows %s", tc.tier, tc.other)
			}
		}
		if !found {
			t.Errorf("%s lacks %s: %+v", tc.tier, tc.label, stats.Tiers[tc.tier].Changes)
		}
	}
	for _, row := range render.ResolveCrosspaths(candidate, definition).Advanced {
		if (row.Code == "4-0-2" && (!strings.Contains(row.Attack, "+3 damage against Hardened enemies") || strings.Contains(row.Attack, "Blimp"))) ||
			(row.Code == "2-0-4" && (!strings.Contains(row.Attack, "+4 damage against Blimp enemies") || strings.Contains(row.Attack, "Hardened"))) {
			t.Errorf("%s attack %q", row.Code, row.Attack)
		}
	}
	// One attack with both bonuses names each with its own damage per hit.
	blueprint.BaseAttack.BonusDamage = []mechanics.DamageBonus{{Property: "blimp", Damage: 2}, {Property: "hardened", Damage: 1}}
	if base := render.Base(candidate, definition); base == nil || !strings.Contains(base.Text, "Each hit deals +2 damage against Blimp enemies (3 per hit) and +1 damage against Hardened enemies (2 per hit).") {
		t.Errorf("0-0-0 reads %+v", base)
	}
}
