package unit_test

import (
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// A capstone repair reads the specialty metrics the Tier 5 gate measures,
// with the Definition's vocabulary, so a version 2 control path lists its
// status coverage instead of no metrics.
func TestCapstoneRepairMeasuresTheDefinitionsStatuses(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	blueprint := *stages.Checked.Draft.Candidate.Blueprint
	half := 50.0
	statuses := []m.StatusApplication{{Effect: "slow", Magnitude: &half, Seconds: 2}}
	blueprint.BaseAttack.Statuses = &statuses
	blueprint.Paths.At(0).Specialization = "control"
	definition := unit.DefaultAuthoringDefinition()
	multiplier := 2.0
	definition.Profile.DesignPolicy.MinTier5SpecialtyMultiplier = &multiplier
	got := s.Stringify(unit.CapstoneRepairContext(blueprint, &unit.Request{MechanicsDefinition: &definition}))
	if !strings.Contains(got, `"metric":"slow coverage upper bound"`) {
		t.Errorf("the repair context lacks the slow coverage: %s", got)
	}
}
