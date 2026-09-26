package render_test

import (
	"os"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/render"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

const sheetHeading = "## 0-0-0: "

// Saved Result 4d938ae2 (luffy3b, v14) was made before d7f8950 added
// ineffectiveChanges. Its tier 3 sets distinct-targets on a one-projectile
// attack, which today's authoring check flags. Every legal build still
// resolves to valid values, so reading paths draw its sheet while check
// still reports the change (#37).
func TestSavedResultKeepsSheetUnderLaterAuthoringCheck(t *testing.T) {
	data, err := os.ReadFile("../../../../data/reference/captures/luffy-3b.result.json")
	if err != nil {
		t.Fatal(err)
	}
	value, err := s.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	result, err := unit.ParseResult(value)
	if err != nil {
		t.Fatal(err)
	}
	candidate, definition := result.Candidate, result.Prepared.Request.MechanicsDefinition
	const flagged = "paths.path1.tiers.tier3.changes.3"

	compact, err := render.Markdown(value, false)
	if err != nil || !strings.Contains(compact, sheetHeading) || !strings.Contains(compact, "## Crosspaths") {
		t.Errorf("the Markdown render lost the unit sheet (err %v)", err)
	}
	if render.Stats(candidate, definition) == nil || render.Base(candidate, definition) == nil ||
		render.Purchases(candidate, definition, nil) == nil || render.ResolveCrosspaths(candidate, definition) == nil {
		t.Error("the view lost stats, base, purchases or crosspaths")
	}
	if _, err := mechanics.ResolveBuild(candidate.Blueprint, mechanics.Selection{3, 0, 0}, *definition); err != nil {
		t.Errorf("3-0-0 does not resolve: %v", err)
	}

	issues := render.AuthoringIssues(candidate, definition)
	if len(issues) != 1 || issues[0].Path != flagged {
		t.Errorf("authoring issues %+v", issues)
	}
	// Detailed lists the current checks apart from the stored findings; the
	// compact render stays sheet-only.
	detailed, err := render.Markdown(value, true)
	if err != nil || !strings.Contains(detailed, "## Current checks") || !strings.Contains(detailed, "| "+render.Escape(flagged)+" | This change has no effect") {
		t.Errorf("the detailed render does not list the current checks (err %v)", err)
	}
	if strings.Contains(compact, "Current checks") || strings.Contains(compact, "This change has no effect") {
		t.Error("the compact render lists the current checks")
	}

	draft := unit.Draft{SchemaVersion: result.SchemaVersion, Kind: "draft", Prepared: result.Prepared, Candidate: candidate, Run: result.Run.Draft}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	reported := false
	for _, f := range checked.Findings {
		if f.Rule == "typed-mechanics" && f.Outcome == "fail" && f.Subject == flagged {
			reported = true
		}
	}
	if !reported {
		t.Errorf("check no longer reports the ineffective change at %s", flagged)
	}
}

// A unit that passes today's authoring checks has no current checks section.
func TestPassingUnitHasNoCurrentChecks(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	result := stages.Result
	if issues := render.AuthoringIssues(result.Candidate, result.Prepared.Request.MechanicsDefinition); issues != nil {
		t.Errorf("authoring issues %+v", issues)
	}
	if detailed, _ := render.Markdown(s.FromGoValue(result), true); strings.Contains(detailed, "## Current checks") {
		t.Error("the detailed render lists current checks for a passing unit")
	}
}

// Reading paths still reject a blueprint whose structure is invalid or
// whose legal builds do not resolve to valid values.
func TestReadingPathsStillRejectInvalidBlueprints(t *testing.T) {
	for _, test := range []struct {
		name string
		edit func(*mechanics.Blueprint)
	}{
		{"missing source fact", func(b *mechanics.Blueprint) {
			b.Paths.Path1.SourceFactIndices = []int{len(b.SourceFacts)}
		}},
		{"unresolvable pierce", func(b *mechanics.Blueprint) {
			b.Paths.Path1.Tiers.Tier1.Changes = []mechanics.Change{{Kind: "stat", Target: "base", Stat: "pierce", Operation: "add", Number: 0.5}}
		}},
	} {
		stages, err := fixture.Build()
		if err != nil {
			t.Fatal(err)
		}
		result := stages.Result
		blueprint := *result.Candidate.Blueprint
		blueprint.Paths.Path1.Tiers.Tier1.Changes = append([]mechanics.Change(nil), blueprint.Paths.Path1.Tiers.Tier1.Changes...)
		test.edit(&blueprint)
		result.Candidate.Blueprint = &blueprint
		definition := result.Prepared.Request.MechanicsDefinition

		if len(mechanics.ValidateStructure(&blueprint, *definition)) == 0 {
			t.Errorf("%s: the structure check passes", test.name)
		}
		if markdown, err := render.Markdown(s.FromGoValue(result), false); err == nil && strings.Contains(markdown, sheetHeading) {
			t.Errorf("%s: the Markdown render draws a sheet", test.name)
		}
		if render.Stats(result.Candidate, definition) != nil || render.Base(result.Candidate, definition) != nil {
			t.Errorf("%s: the view resolves stats", test.name)
		}
		if _, err := mechanics.ResolveBuild(&blueprint, mechanics.Selection{1, 0, 0}, *definition); err == nil {
			t.Errorf("%s: 1-0-0 resolves", test.name)
		}
		if issues := render.AuthoringIssues(result.Candidate, definition); issues != nil {
			t.Errorf("%s: authoring issues reported for an unreadable blueprint: %+v", test.name, issues)
		}
	}
}
