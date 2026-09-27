package render_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/render"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// A revision's patch notes name a proposed mechanic it adds to a purchase,
// apart from the unchanged typed mechanics.
func TestRevisionNotesNameNewProposedMechanics(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	request := stages.Prepared.Request
	feedback := "Keep the critical bolt on x-x-4 as a proposed mechanic."
	request.Previous = &unit.Previous{ResultID: stages.Result.ID, Draft: stages.Result.Candidate, Findings: stages.Result.Findings}
	request.Feedback = &feedback
	prepared, err := unit.Prepare(s.FromGoValue(request))
	if err != nil {
		t.Fatal(err)
	}
	mechanics, _ := fixture.JSON("mechanics")
	paths, _ := mechanics.(*s.Object).Get("paths")
	path3, _ := paths.(*s.Object).Get("path3")
	tiers, _ := path3.(*s.Object).Get("tiers")
	tier4, _ := tiers.(*s.Object).Get("tier4")
	tier4.(*s.Object).Set("proposedMechanics", []any{s.NewObject().
		Set("name", "Critical bolt").Set("effect", "Every tenth bolt deals five times its damage.").Set("sourceIds", []any{"source1:17"})})
	plan, _ := fixture.JSON("plan")
	draft, err := unit.DraftUnit(context.Background(), prepared, &fixture.Model{Outputs: []any{plan, mechanics}}, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	view, err := render.ReadView(s.FromGoValue(draft))
	if err != nil {
		t.Fatal(err)
	}
	notes := render.Revision(view)
	if notes == nil || strings.Join(notes.Mechanics, "\n") != "x-x-4 new proposed mechanic: Critical bolt." || len(notes.ChangedBuilds) != 0 {
		t.Errorf("notes %+v", notes)
	}
}
