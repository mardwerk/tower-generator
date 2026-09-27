package unit_test

import (
	"slices"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// A purchase shows the technique the plan maps it to, and a name that points
// to another repertoire entry is flagged for review, never remapped (#35). A
// Luffy Result named base-attack purchases "Boundman Force" and "Longer
// Boundman Reach". Only a word that identifies exactly one entry counts, so
// shared words such as Gear or Haki flag nothing.
func TestPurchaseNamesThatPointToAnotherTechniqueAreFlagged(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan := *stages.Draft.Run.DesignPlan
	intents := *plan.UpgradeIntents
	plan.UpgradeIntents = &intents
	plan.Base.Name = "Gum-Gum Pistol"
	plan.Repertoire = []unit.PlanRepertoire{
		{Name: "Gear 2"}, {Name: "Gear 3"}, {Name: "Gear 4 Boundman"}, {Name: "Gear 4 Snakeman"},
		{Name: "Armament Haki"}, {Name: "Observation Haki"},
	}
	intents.Path3.Tier1.Technique = "Gum-Gum Pistol"
	intents.Path3.Tier4.Technique = "Gear 4 Snakeman"
	intents.Path2.Tier3.Technique = "Gear 2"
	for _, test := range []struct {
		path, tier int
		name       string
		want       []string
	}{
		{2, 1, "Boundman Force", []string{"Gear 4 Boundman"}},
		// Shared words identify no single entry.
		{2, 1, "Haki Punch", nil},
		{2, 1, "Gear Rush", nil},
		// The purchase's own technique and the base attack never count.
		{2, 4, "Snakeman Barrage", nil},
		{1, 3, "Jet Pistol", nil},
		// A distinctive word of another entry does, whatever the technique.
		{1, 3, "Armament Jet", []string{"Armament Haki"}},
		{1, 3, "Boundman Armament Jet", []string{"Gear 4 Boundman", "Armament Haki"}},
	} {
		if got := unit.PurchaseNameMentions(&plan, test.path, test.tier, test.name); !slices.Equal(got, test.want) {
			t.Errorf("%s at path %d tier %d: %q, want %q", test.name, test.path+1, test.tier, got, test.want)
		}
	}
	if technique, base := unit.PurchaseTechnique(&plan, 2, 1); technique != "Gum-Gum Pistol" || !base {
		t.Errorf("x-x-1 maps to %q, base %v", technique, base)
	}
	if technique, base := unit.PurchaseTechnique(&plan, 1, 3); technique != "Gear 2" || base {
		t.Errorf("x-3-x maps to %q, base %v", technique, base)
	}
	// Plans made before techniques were typed show no mapping and no flag.
	intents.Path3.Tier1.Technique = ""
	if technique, _ := unit.PurchaseTechnique(&plan, 2, 1); technique != "" {
		t.Errorf("a purchase without a technique maps to %q", technique)
	}
	if got := unit.PurchaseNameMentions(&plan, 2, 1, "Boundman Force"); got != nil {
		t.Errorf("a purchase without a technique is flagged: %q", got)
	}
	if technique, _ := unit.PurchaseTechnique(nil, 0, 1); technique != "" {
		t.Errorf("no plan maps to %q", technique)
	}
}
