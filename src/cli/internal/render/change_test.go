package render_test

import (
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/render"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Every change reads as "value → value (delta)", and multipliers as
// percentages (#61): the delta says what the purchase does, the values what
// the Unit has when only that purchase is bought.

// fixtureUnit is the scripted reference unit with an edited blueprint.
func fixtureUnit(t *testing.T, edit func(*mechanics.Blueprint)) (unit.Candidate, *mechanics.Definition) {
	t.Helper()
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	candidate := stages.Result.Candidate
	blueprint := *candidate.Blueprint
	if edit != nil {
		edit(&blueprint)
	}
	candidate.Blueprint = &blueprint
	return candidate, stages.Result.Prepared.Request.MechanicsDefinition
}

func effects(t *testing.T, candidate unit.Candidate, definition *mechanics.Definition, path, tier int) string {
	t.Helper()
	purchases := render.Purchases(candidate, definition, nil)
	if purchases == nil {
		t.Fatal("the blueprint is invalid")
	}
	return strings.Join(purchases[path].Purchases[tier-1].Effects, " ")
}

func tierChanges(t *testing.T, candidate unit.Candidate, definition *mechanics.Definition, key string) string {
	t.Helper()
	stats := render.Stats(candidate, definition)
	if stats == nil {
		t.Fatal("no kit stats")
	}
	return s.Stringify(s.FromGoValue(stats.Tiers[key].Changes))
}

// An interval multiplier reads as the attack speed it gives: ×0.85 attacks
// 1/0.85 - 1, about 18%, faster. An added interval keeps its seconds too.
func TestIntervalMultiplierReadsAsAttackSpeed(t *testing.T) {
	candidate, definition := fixtureUnit(t, nil)
	if got := effects(t, candidate, definition, 1, 1); got != "Shortens the attack interval from 0.95 s to 0.8075 s (attacks 18% faster)." {
		t.Errorf("x-1-x reads %q", got)
	}
	if got := effects(t, candidate, definition, 0, 3); !strings.Contains(got, "Lengthens the attack interval from 0.95 s to 1.15 s (+0.2 s, attacks 17% slower).") {
		t.Errorf("3-x-x reads %q", got)
	}
	want := `[{"key":"intervalSeconds","before":0.95,"after":0.8075,"delta":"attacks 18% faster","improvement":true}]`
	if got := tierChanges(t, candidate, definition, "path-2:1"); got != want {
		t.Errorf("x-1-x kit stats\n got %s\nwant %s", got, want)
	}
}

// An Active Ability's multipliers read as percentages: a damage multiplier
// of 2 is +100% damage and an interval multiplier of 0.5 attacks 100%
// faster.
func TestActiveMultipliersReadAsPercentages(t *testing.T) {
	candidate, definition := fixtureUnit(t, func(blueprint *mechanics.Blueprint) {
		tier := &blueprint.Paths.Path2.Tiers.Tier4
		tier.Changes = append([]mechanics.Change(nil), tier.Changes...)
		for i, change := range tier.Changes {
			if change.Kind == "unlockBoost" {
				boost := *change.Boost
				boost.DamageMultiplier, boost.IntervalMultiplier, boost.RangeBonus = 2, 0.5, 0
				tier.Changes[i].Boost = &boost
			}
		}
	})
	got := effects(t, candidate, definition, 1, 4)
	if !strings.Contains(got, "Adds Fan Club Frenzy, this Unit's Active Ability: for 15 s it gives the purchased attack +100% damage and 100% faster attacks, so the attack deals 2 damage every 0.1196 s at range 32.") || strings.Contains(got, "multipl") {
		t.Errorf("x-4-x reads %q", got)
	}
	changes := tierChanges(t, candidate, definition, "path-2:4")
	for _, want := range []string{`{"key":"damageMultiplier","after":"+100%"}`, `{"key":"intervalMultiplier","after":"100% faster"}`} {
		if !strings.Contains(changes, want) {
			t.Errorf("x-4-x kit stats lack %s: %s", want, changes)
		}
	}
}

// A purchase that changes an owned Active Ability's multipliers reads in
// percentages too, with the delta in percentage points.
func TestModifyingTheActiveShowsPercentagesAndTheDelta(t *testing.T) {
	candidate, definition := fixtureUnit(t, func(blueprint *mechanics.Blueprint) {
		tier := &blueprint.Paths.Path2.Tiers.Tier5
		tier.Changes = append(append([]mechanics.Change(nil), tier.Changes...),
			mechanics.Change{Kind: "modifyBoost", Target: "base", Stat: "intervalMultiplier", Operation: "set", Number: 0.05})
	})
	got := effects(t, candidate, definition, 1, 5)
	for _, want := range []string{
		"Raises Fan Club Frenzy's damage bonus from +0% to +100% (+100 percentage points).",
		"Raises Fan Club Frenzy duration from 15 s to 20 s (+5 s).",
		"Raises Fan Club Frenzy's attack speed from 1500% faster to 1900% faster (+400 percentage points).",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("x-5-x lacks %q: %s", want, got)
		}
	}
	changes := tierChanges(t, candidate, definition, "path-2:5")
	for _, want := range []string{
		`{"key":"damageMultiplier","before":"+0%","after":"+100%","delta":"+100 percentage points","improvement":true}`,
		`{"key":"intervalMultiplier","before":"1500% faster","after":"1900% faster","delta":"+400 percentage points","improvement":true}`,
	} {
		if !strings.Contains(changes, want) {
			t.Errorf("x-5-x kit stats lack %s: %s", want, changes)
		}
	}
}

// Crosspath rows show each number's delta and an interval's attack speed,
// in the early and the advanced tables and inside the Active window.
func TestCrosspathRowsShowDeltas(t *testing.T) {
	candidate, definition := fixtureUnit(t, nil)
	crosspaths := render.ResolveCrosspaths(candidate, definition)
	if crosspaths == nil {
		t.Fatal("no crosspaths")
	}
	rows := map[string]string{}
	for _, row := range append(append([]render.BuildRow{}, crosspaths.Early...), crosspaths.Advanced...) {
		rows[row.Code] = s.Stringify(s.FromGoValue(row.Contributions))
	}
	for code, want := range map[string]string{
		"1-1-0": `[{"from":"1-x-x","changes":["pierce 2 → 3 (+1)"]},{"from":"x-1-x","changes":["interval 0.95 s → 0.8075 s (attacks 18% faster)"]}]`,
		"0-2-1": `[{"from":"x-1-x and x-2-x","changes":["interval 0.95 s → 0.6379 s (attacks 49% faster)"]},{"from":"x-x-1","changes":["range 32 → 40 (+8)"]}]`,
		"3-1-0": `[{"from":"x-1-x","changes":["interval 1.15 s → 0.9775 s (attacks 18% faster)"]}]`,
		"1-4-0": `[{"from":"1-x-x","changes":["pierce 2 → 3 (+1)","during Fan Club Frenzy: pierce 2 → 3 (+1)"]}]`,
	} {
		if rows[code] != want {
			t.Errorf("%s\n got %s\nwant %s", code, rows[code], want)
		}
	}
}

// The details report words a typed unit's purchases and Active Ability as
// the sheet does, not in the compiled "multiply by" wording the candidate
// keeps for review.
func TestDetailsWordChangesAsTheSheet(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	details, err := render.Markdown(s.FromGoValue(stages.Result), true)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"| x-1-x | Quick Shots | 100 Gold. Shortens the attack interval from 0.95 s to 0.8075 s (attacks 18% faster). |",
		"At x-4-x: Shortens the attack interval from 0.4784 s to 0.2392 s (attacks 100% faster). Adds Fan Club Frenzy, this Unit's Active Ability: for 15 s it gives the purchased attack 1500% faster attacks and +8 range,",
	} {
		if !strings.Contains(details, want) {
			t.Errorf("the details lack %q", want)
		}
	}
	if strings.Contains(details, "multiply by") || strings.Contains(details, "multiply the purchased") {
		t.Error("the details keep the compiled multiplier wording")
	}
}
