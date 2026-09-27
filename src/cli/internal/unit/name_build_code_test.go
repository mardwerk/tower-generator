package unit_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/render"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// A name drops the build code it opens with, which the Unit sheet already
// shows; other names, and a name that is only a code, are kept.
func TestNameWithoutBuildCode(t *testing.T) {
	for name, want := range map[string]string{
		"0-0-0 Gum-Gum Pistol":                     "Gum-Gum Pistol",
		"1-x-x Heavy Pistol":                       "Heavy Pistol",
		"x-4-x: Jet Boost":                         "Jet Boost",
		"X-X-5 - Python Redirected Pistol":         "Python Redirected Pistol",
		"In build 3-x-x, Gigant Pistol":            "Gigant Pistol",
		"in build x-x-4, python":                   "Python",
		"Gum-Gum Pistol":                           "Gum-Gum Pistol",
		"Gear 2":                                   "Gear 2",
		"X-Ray Eyes":                               "X-Ray Eyes",
		"Jet Pistol 1-x-x":                         "Jet Pistol 1-x-x",
		"3-x-x":                                    "3-x-x",
		"Gomu Gomu no Jet Pistol":                  "Gomu Gomu no Jet Pistol",
		"Plasma Monkey Fan Club":                   "Plasma Monkey Fan Club",
		"0-0-0 Gum-Gum Pistol (Gomu Gomu no Pisu)": "Gum-Gum Pistol (Gomu Gomu no Pisu)",
	} {
		if got := unit.NameWithoutBuildCode(name); got != want {
			t.Errorf("%q: got %q, want %q", name, got, want)
		}
	}
}

// The v32 Luffy plan named its base attack "0-0-0 Gum-Gum Pistol", and the
// sheet printed "0-0-0: 0-0-0 Gum-Gum Pistol" (OPUS-NET-61-10). Drafting
// drops a build code that opens the base attack's name or a purchase name,
// without a retry: the plan, the milestones that adapt the base attack by
// that name, the blueprint and the sheet carry the name alone.
func TestDraftDropsBuildCodesFromNames(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	plan := recordedOutput(t, "plan")
	at(plan, "base").(*s.Object).Set("name", "0-0-0 Gum-Gum Pistol")
	for _, path := range []string{"path1", "path2", "path3"} {
		for _, tier := range []string{"tier1", "tier2", "tier3", "tier4", "tier5"} {
			milestone := at(plan, "paths", path, "milestones", tier).(*s.Object)
			if at(milestone, "technique") == "Dart Throw" {
				milestone.Set("technique", "0-0-0 Gum-Gum Pistol")
			}
		}
	}
	mechanics := recordedOutput(t, "mechanics")
	at(mechanics, "paths", "path1", "tiers", "tier1").(*s.Object).Set("name", "1-x-x Sharp Shots")
	at(mechanics, "paths", "path2", "tiers", "tier4").(*s.Object).Set("name", "In build x-4-x, Super Monkey Fan Club")
	at(mechanics, "paths", "path3", "tiers", "tier2").(*s.Object).Set("name", "x-x-2: Enhanced Eyesight")
	model := &fixture.Model{Outputs: []any{plan, mechanics}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if len(model.Requests) != 2 {
		t.Errorf("%d model calls, want the plan and the mechanics without a retry", len(model.Requests))
	}
	designPlan := draft.Run.DesignPlan
	if designPlan.Base.Name != "Gum-Gum Pistol" || designPlan.UpgradeIntents.At(0).At(1).Technique != "Gum-Gum Pistol" {
		t.Errorf("the plan's base %q and 1-x-x technique %q", designPlan.Base.Name, designPlan.UpgradeIntents.At(0).At(1).Technique)
	}
	blueprint := draft.Candidate.Blueprint
	if blueprint.BaseAttack.Name != "Gum-Gum Pistol" || draft.Candidate.BasicAttack.Name != "Gum-Gum Pistol" {
		t.Errorf("the base attack %q, %q", blueprint.BaseAttack.Name, draft.Candidate.BasicAttack.Name)
	}
	for _, c := range []struct {
		code       string
		path, tier int
		want       string
	}{{"1-x-x", 0, 1, "Sharp Shots"}, {"x-4-x", 1, 4, "Super Monkey Fan Club"}, {"x-x-2", 2, 2, "Enhanced Eyesight"}} {
		if got := blueprint.Paths.At(c.path).Tiers.At(c.tier).Name; got != c.want {
			t.Errorf("%s is named %q, want %q", c.code, got, c.want)
		}
	}
	sheet, err := render.Markdown(s.FromGoValue(draft), false)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"## 0-0-0: Gum-Gum Pistol\n", "**1-x-x Sharp Shots**", "**x-4-x Super Monkey Fan Club**", "**x-x-2 Enhanced Eyesight**"} {
		if !strings.Contains(sheet, want) {
			t.Errorf("the sheet lacks %q", want)
		}
	}
	if strings.Contains(sheet, "0-0-0 Gum-Gum") || strings.Contains(sheet, "In build") {
		t.Error("the sheet keeps a build code in a name")
	}
}

// A saved Result whose names carry a build code, as Luffy's b37d6e09 does,
// still loads and renders as it was saved: reading applies no authoring
// normalization.
func TestSavedResultKeepsItsNames(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	value := s.FromGoValue(stages.Result).(*s.Object)
	at(value, "candidate", "blueprint", "baseAttack").(*s.Object).Set("name", "0-0-0 Gum-Gum Pistol")
	at(value, "candidate", "basicAttack").(*s.Object).Set("name", "0-0-0 Gum-Gum Pistol")
	at(value, "run", "draft", "designPlan", "base").(*s.Object).Set("name", "0-0-0 Gum-Gum Pistol")
	if _, err := unit.ParseResult(value); err != nil {
		t.Fatalf("the saved Result does not load: %v", err)
	}
	sheet, err := render.Markdown(value, false)
	if err != nil || !strings.Contains(sheet, "## 0-0-0: 0-0-0 Gum-Gum Pistol\n") {
		t.Errorf("the saved Result does not render as saved (err %v)", err)
	}
}
