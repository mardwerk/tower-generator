package render_test

import (
	"context"
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

func TestMarkdownRendersEveryStage(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		name     string
		artifact any
		status   string
	}{
		{"draft", s.FromGoValue(stages.Draft), "Draft with checked upgrade mechanics"},
		{"checked", s.FromGoValue(stages.Checked), "Structural checks complete. Model review has not run."},
		{"result", s.FromGoValue(stages.Result), "Structural checks and model review complete."},
	} {
		compact, err := render.Markdown(test.artifact, false)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.HasPrefix(compact, "# Dart Monkey\n") || strings.Contains(compact, test.status) {
			t.Errorf("%s compact render is not only the unit:\n%s", test.name, compact)
		}
		for _, name := range []string{"Juggernaut Line", "Fan Club Line", "Crossbow Line", "Ultra-Juggernaut"} {
			if !strings.Contains(compact, name) {
				t.Errorf("%s render lacks %s", test.name, name)
			}
		}
		detailed, err := render.Markdown(test.artifact, true)
		if err != nil {
			t.Fatal(err)
		}
		// Gold totals are exact: 24,615, not 24,620 (reported on #27).
		for _, section := range []string{"## Evidence", "## Purchase evidence", "## Generation usage", "dart-monkey-atlas-56-3", "Definition BTD6-inspired Gold and Health starter", "T5 total: 24,615 Gold", "| 4-0-0 → 5-0-0 | 15,000 |"} {
			if !strings.Contains(detailed, section) {
				t.Errorf("%s details lack %q", test.name, section)
			}
		}
	}
	if _, err := render.Markdown(s.FromGoValue(stages.Prepared), false); err == nil {
		t.Error("a prepared request rendered as a unit")
	}
}

// Purchase evidence shows the Active's time-averaged rates next to its peak
// and duty fraction, and states each crosspath purchase's absolute gain and
// price, in the Definition's currency (Gold in the fixture), beside its main
// path's fifth purchase. Fan Club Frenzy lasts 15 of 50 s at
// x-4-x and 20 of 50 s at x-5-x: 0.6 × 12.54 + 0.4 × 401.3 = 168.
func TestDetailsShowTimeAveragedAndSidePurchaseEvidence(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	detailed, err := render.Markdown(s.FromGoValue(stages.Result), true)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		// A milestone is the path's own purchase: no comparison follows it.
		"| 0-4-0 → 0-5-0 | 45,000 | active peak direct damage rate: 200.7 → 401.3; active peak group damage rate upper bound: 401.3 → 802.6; active duty fraction: 0.3 → 0.4; time-averaged direct damage rate: 68.97 → 168; time-averaged group damage rate upper bound: 137.9 → 336.1 |",
		"| 0-5-0 → 1-5-0 | 140 | group damage rate upper bound: 25.08 → 37.62; active peak group damage rate upper bound: 802.6 → 1204; time-averaged group damage rate upper bound: 336.1 → 504.1. Side 1-x-x adds direct +0, group +168 for 140 Gold; the path's x-5-x adds direct +99.07, group +198.1 for 45,000 Gold. |",
		"| 0-5-0 → 0-5-1 | 90 | range: 32 → 40. Side x-x-1 adds direct +0, group +0 for 90 Gold; the path's x-5-x adds direct +99.07, group +198.1 for 45,000 Gold. |",
	} {
		if !strings.Contains(detailed, want) {
			t.Errorf("details lack %q", want)
		}
	}
}

// A 140 Gold side purchase adding direct +1.05 on top of a 45,000 Gold
// capstone adding +12.64, the shape of a live run on #27, reads as the
// smaller gain it is: absolute gains and prices, no gain per currency.
func TestDetailsShowACheapSidePurchaseAsTheSmallerGain(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	evaluation, err := s.Decode([]byte(`{"paths":[{"path":"path3","name":"Bottom","milestones":[
{"from":[0,0,4],"to":[0,0,5],"incrementalGold":45000,"metricDeltas":{
"direct damage rate":{"before":10,"after":22.64,"change":12.64},
"group damage rate upper bound":{"before":20,"after":45.28,"change":25.28}}}],
"crosspaths":[{"from":[0,0,5],"to":[1,0,5],"incrementalGold":140,"metricDeltas":{
"direct damage rate":{"before":22.64,"after":23.69,"change":1.05},
"group damage rate upper bound":{"before":45.28,"after":47.38,"change":2.1}}}]}],
"limitations":[]}`))
	if err != nil {
		t.Fatal(err)
	}
	view, err := render.ReadView(s.FromGoValue(stages.Result))
	if err != nil {
		t.Fatal(err)
	}
	view.DesignEvaluation = evaluation.(*s.Object)
	detailed := render.Detailed(view)
	want := "| 0-0-5 → 1-0-5 | 140 | direct damage rate: 22.64 → 23.69; group damage rate upper bound: 45.28 → 47.38. Side 1-x-x adds direct +1.05, group +2.1 for 140 Gold; the path's x-x-5 adds direct +12.64, group +25.28 for 45,000 Gold. |"
	if !strings.Contains(detailed, want) {
		t.Errorf("details lack %q", want)
	}
	for _, ratio := range []string{"Per 100", "per 100", "stronger"} {
		if strings.Contains(detailed, ratio) {
			t.Errorf("details contain %q", ratio)
		}
	}
}

func TestViewCarriesStatsForEveryTier(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	view, err := render.ReadView(s.FromGoValue(stages.Result))
	if err != nil || view.Kind != "result" {
		t.Fatalf("view %s: %v", view.Kind, err)
	}
	stats := render.Stats(view.Candidate, view.Prepared.Request.MechanicsDefinition)
	if stats == nil {
		t.Fatal("no stats")
	}
	for _, path := range []string{"path-1", "path-2", "path-3"} {
		for tier := 1; tier <= 5; tier++ {
			if _, ok := stats.Tiers[render.TierKey(path, tier)]; !ok {
				t.Errorf("no stats for %s tier %d", path, tier)
			}
		}
	}
	if text := render.UsageSummaryText(view.Usage); text == "" {
		t.Error("no usage summary")
	}
}

// The unit sheet starts with the name and 0-0-0, names purchases by build
// code, lists every early and advanced crosspath build and holds nothing
// but the unit: no checks, costs, unsupported or reserved lists, defaults
// or the plan's private purchase notes.
func TestUnitSheetUsesBuildCodesAndEveryCrosspath(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	compact, err := render.Markdown(s.FromGoValue(stages.Result), false)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(compact, "# Dart Monkey\n\n## 0-0-0: Dart Throw\n\nPlacement costs 200 Gold.") {
		t.Fatalf("the unit does not start with its name and 0-0-0:\n%s", compact[:200])
	}
	unitPart := compact
	for _, diagnostic := range []string{
		"Structural checks", "## About this artifact", "## Unsupported mechanics", "## Reserved techniques", "Critical shot counter",
		"Gold and Health starter", "cannot target", "clear path", "First targeting", "damage by 1 ", "USD", "render --details",
	} {
		if strings.Contains(compact, diagnostic) {
			t.Errorf("the unit sheet shows %q", diagnostic)
		}
	}
	for _, want := range []string{
		// A purchase that adapts a named technique shows the plan's
		// description of that adaptation before its effects; one of the base
		// attack alone shows only its effects.
		"## Top path: Juggernaut Line", "**3-x-x Spike-o-pult** (320 Gold). Replaces the dart with a heavier spiked ball that deals more damage, reaches farther and pierces far more enemies, at a slower throw. Raises damage from 1 to 2 (+1).",
		"**1-x-x Sharp Shots** (140 Gold). Raises pierce from 2 to 3 (+1).",
		"**x-4-x Super Monkey Fan Club** (7,200 Gold).", "Adds Fan Club Frenzy, this Unit's Active Ability: for 15 s it multiplies its interval by 0.0625 and adds 8 range",
		"**x-x-5 Crossbow Master** (21,500 Gold).", "Switches damage from Sharp to Normal.",
		"### Early builds (12)", "### Advanced builds (36)",
	} {
		if !strings.Contains(unitPart, want) {
			t.Errorf("the unit lacks %q", want)
		}
	}
	for _, code := range []string{
		"1-1-0", "1-2-0", "1-0-1", "1-0-2", "2-1-0", "2-2-0", "2-0-1", "2-0-2", "0-1-1", "0-1-2", "0-2-1", "0-2-2",
		"3-1-0", "3-2-0", "3-0-1", "3-0-2", "4-1-0", "4-2-0", "4-0-1", "4-0-2", "5-1-0", "5-2-0", "5-0-1", "5-0-2",
		"1-3-0", "2-3-0", "0-3-1", "0-3-2", "1-4-0", "2-4-0", "0-4-1", "0-4-2", "1-5-0", "2-5-0", "0-5-1", "0-5-2",
		"1-0-3", "2-0-3", "0-1-3", "0-2-3", "1-0-4", "2-0-4", "0-1-4", "0-2-4", "1-0-5", "2-0-5", "0-1-5", "0-2-5",
	} {
		if !strings.Contains(unitPart, "| "+code+" |") {
			t.Errorf("crosspath %s is missing", code)
		}
	}
	// The side purchase's effect reaches the x-4-x boost window too.
	if !strings.Contains(unitPart, "| 1-4-0 | 8,280 Gold | 1-x-x: pierce 2 → 3, during Fan Club Frenzy: pierce 2 → 3 |") {
		t.Error("the 1-4-0 row does not show what the boost window receives")
	}
	plan := stages.Result.Run.Draft.DesignPlan
	for _, private := range []string{plan.Paths.Path1.BuyFor, plan.Paths.Path1.Weakness, plan.Paths.Path1.CapstoneValue, "Tier 1", "T3", "Path 1"} {
		if strings.Contains(compact, private) {
			t.Errorf("the render shows %q", private)
		}
	}
	detailed, _ := render.Markdown(s.FromGoValue(stages.Result), true)
	if strings.Contains(detailed, plan.Paths.Path1.BuyFor) || strings.Contains(detailed, "Capstone intention") {
		t.Error("the detailed render prints private purchase notes")
	}
	for _, diagnostic := range []string{"Structural checks and model review complete", "Draft claim: the Definition cannot express Critical shot counter."} {
		if !strings.Contains(detailed, diagnostic) {
			t.Errorf("the diagnostics lack %q", diagnostic)
		}
	}
	crosspaths := render.ResolveCrosspaths(stages.Result.Candidate, stages.Result.Prepared.Request.MechanicsDefinition)
	if len(crosspaths.Early) != 12 || len(crosspaths.Advanced) != 36 || crosspaths.Advanced[0].Code != "3-1-0" || crosspaths.Early[0].Code != "1-1-0" {
		t.Errorf("crosspaths %d early, %d advanced", len(crosspaths.Early), len(crosspaths.Advanced))
	}
}

// A revision's notes keep changed mechanics apart from renamed purchases.
func TestRevisionNotesSeparateMechanicsFromWording(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	request := stages.Prepared.Request
	feedback := "Rename the first top purchase and make Spike-o-pult cheaper."
	request.Previous = &unit.Previous{ResultID: stages.Result.ID, Draft: stages.Result.Candidate, Findings: stages.Result.Findings}
	request.Feedback = &feedback
	prepared, err := unit.Prepare(s.FromGoValue(request))
	if err != nil {
		t.Fatal(err)
	}
	mechanics, _ := fixture.JSON("mechanics")
	tiers := func(o any) *s.Object {
		paths, _ := o.(*s.Object).Get("paths")
		path, _ := paths.(*s.Object).Get("path1")
		value, _ := path.(*s.Object).Get("tiers")
		return value.(*s.Object)
	}
	tier1, _ := tiers(mechanics).Get("tier1")
	tier1.(*s.Object).Set("name", "Sharper Shots")
	tier3, _ := tiers(mechanics).Get("tier3")
	tier3.(*s.Object).Set("cost", 300.0)
	plan, _ := fixture.JSON("plan")
	model := &fixture.Model{Outputs: []any{plan, mechanics}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	view, err := render.ReadView(s.FromGoValue(draft))
	if err != nil {
		t.Fatal(err)
	}
	notes := render.Revision(view)
	if notes == nil {
		t.Fatal("no revision notes")
	}
	if strings.Join(notes.Mechanics, "\n") != "3-x-x price 320 Gold to 300 Gold." || strings.Join(notes.Wording, "\n") != "1-x-x renamed from Sharp Shots to Sharper Shots." || len(notes.ChangedBuilds) != 0 {
		t.Errorf("notes %+v", notes)
	}
	compact, _ := render.Markdown(s.FromGoValue(draft), false)
	if !strings.Contains(compact, "## Patch notes\n\n### Mechanics\n\n- 3-x-x price 320 Gold to 300 Gold.") || !strings.Contains(compact, "### Wording\n\n- 1-x-x renamed from Sharp Shots to Sharper Shots.") {
		t.Error("the render lacks separate revision notes")
	}
}

// The reference captures render as committed, so their sheets, and the
// README excerpt quoted from one, stay current with the renderer. After a
// deliberate renderer change, render each capture again.
func TestReferenceCapturesRenderAsCommitted(t *testing.T) {
	files, _ := filepath.Glob("../../../../data/reference/captures/*.json")
	if len(files) == 0 {
		t.Fatal("no reference captures")
	}
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		value, err := s.Decode(data)
		if err != nil {
			t.Fatalf("%s: %v", file, err)
		}
		rendered, err := render.Markdown(value, false)
		if err != nil {
			t.Fatalf("%s: %v", file, err)
		}
		committed, err := os.ReadFile(strings.TrimSuffix(file, ".json") + ".md")
		if err != nil || rendered != string(committed) {
			t.Errorf("%s no longer renders as its committed sheet; run render on it again", filepath.Base(file))
		}
	}
}

// An early crosspath row reports what each path does on its own, so a stat
// both paths raise is not counted twice; the resulting attack combines them.
func TestEarlyCrosspathsDoNotRepeatSharedChanges(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	candidate := stages.Result.Candidate
	blueprint := *candidate.Blueprint
	tier := &blueprint.Paths.Path3.Tiers.Tier1
	tier.Changes = append(append([]mechanics.Change(nil), tier.Changes...), mechanics.Change{Kind: "stat", Target: "base", Stat: "pierce", Operation: "add", Number: 1})
	candidate.Blueprint = &blueprint
	crosspaths := render.ResolveCrosspaths(candidate, stages.Result.Prepared.Request.MechanicsDefinition)
	if crosspaths == nil {
		t.Fatal("the edited blueprint is invalid")
	}
	for _, row := range crosspaths.Early {
		if row.Code != "1-0-1" {
			continue
		}
		got := s.Stringify(s.FromGoValue(row.Contributions))
		if !strings.Contains(got, `{"from":"1-x-x","changes":["pierce 2 → 3"]}`) || !strings.Contains(got, `"pierce 2 → 3"`) || !strings.Contains(row.Attack, "pierce 4") {
			t.Errorf("1-0-1 reports %s, %s", got, row.Attack)
		}
		return
	}
	t.Fatal("no 1-0-1 row")
}

// Additions apply before multipliers, so an addition under an earlier
// multiplier raises the value by more than it says; the sentence names the
// multiplier (seen on #27: "Raises damage from 6 to 18 (+4)").
func TestScaledAdditionsNameTheMultiplier(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	candidate := stages.Result.Candidate
	blueprint := *candidate.Blueprint
	tier := &blueprint.Paths.Path3.Tiers.Tier4
	tier.Changes = append([]mechanics.Change(nil), tier.Changes...)
	for i, change := range tier.Changes {
		if change.Kind == "stat" && change.Stat == "damage" {
			tier.Changes[i].Operation, tier.Changes[i].Number = "multiply", 2
		}
	}
	candidate.Blueprint = &blueprint
	purchases := render.Purchases(candidate, stages.Result.Prepared.Request.MechanicsDefinition, nil)
	if purchases == nil {
		t.Fatal("the edited blueprint is invalid")
	}
	effects := strings.Join(purchases[2].Purchases[4].Effects, " ")
	if !strings.Contains(effects, "(+2, multiplied by the earlier ×2).") || strings.Contains(strings.Join(purchases[2].Purchases[3].Effects, " "), "multiplied by the earlier") {
		t.Errorf("x-x-4 %v, x-x-5 %s", purchases[2].Purchases[3].Effects, effects)
	}
}

// The Default Profile's Knockback is a control status with a speed-multiple
// magnitude: it validates in every legal build, reads as a sentence, counts
// as control and does not reach Blimp or Boss enemies (requested on #27 for
// Luffy's Bazooka and Kong Gun, and Juggernaut's knockback).
func TestKnockbackResolvesAndReads(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	definition := stages.Result.Prepared.Request.MechanicsDefinition
	terms := definition.Terms()
	effect, ok := terms.Effect("knockback")
	if !ok || effect.Kind != mechanics.KindKnockback || strings.Join(effect.Immune, ",") != "blimp,boss" {
		t.Fatalf("the Default Profile's Knockback: %+v", effect)
	}
	candidate := stages.Result.Candidate
	blueprint := *candidate.Blueprint
	tier := &blueprint.Paths.Path3.Tiers.Tier3
	tier.Changes = append(append([]mechanics.Change(nil), tier.Changes...),
		mechanics.Change{Kind: "status", Target: "base", Effect: "knockback", Field: "magnitude", Operation: "set", Number: 1.25},
		mechanics.Change{Kind: "status", Target: "base", Effect: "knockback", Field: "seconds", Operation: "set", Number: 0.5},
	)
	if issues := mechanics.ValidateBlueprint(s.FromGoValue(blueprint), s.FromGoValue(*definition)); len(issues) > 0 {
		t.Fatalf("a knockback purchase: %v", issues)
	}
	candidate.Blueprint = &blueprint
	effects := strings.Join(render.Purchases(candidate, definition, nil)[2].Purchases[2].Effects, " ")
	if !strings.Contains(effects, "Knockback 1.25 times enemy speed for 0.5 s") {
		t.Errorf("x-x-3 reads %s", effects)
	}
	build := mechanics.ResolveUnchecked(&blueprint, mechanics.Selection{0, 0, 3})
	metrics := s.Stringify(mechanics.SpecialtyMetricsWith(build, "control", &terms))
	if !strings.Contains(metrics, `"knockback coverage upper bound"`) || strings.Contains(metrics, `"knockback coverage upper bound":0`) {
		t.Errorf("control metrics %s", metrics)
	}
}
