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
