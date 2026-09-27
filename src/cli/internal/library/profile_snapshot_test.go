package library

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// A Result keeps the full Profile it was made under. Editing the saved
// Profile, here to turn on distinctEarlyBenefits and requireTier3BehaviorChange,
// or deleting it, changes neither the Result's design policy nor what
// CheckDraft finds for its draft (OPUS-NET-33-2 step 6).
func TestResultsKeepTheirProfileSnapshot(t *testing.T) {
	profiles, _ := OpenProfiles(filepath.Join(t.TempDir(), "profiles"))
	profile := unit.DefaultProfile()
	profile.ID, profile.Name, profile.Rules.ID = "snapshot-copy", "Snapshot copy", "profile:snapshot-copy"
	policy := *profile.MechanicsDefinition.Profile.DesignPolicy
	policy.DistinctEarlyBenefits = nil
	profile.MechanicsDefinition.Profile.DesignPolicy = &policy
	if _, err := profiles.Save(s.FromGoValue(profile)); err != nil {
		t.Fatal(err)
	}
	saved, err := profiles.Get("snapshot-copy")
	if err != nil {
		t.Fatal(err)
	}
	request, err := fixture.Request()
	if err != nil {
		t.Fatal(err)
	}
	prepared, err := unit.Prepare(s.FromGoValue(unit.ApplyProfile(request, saved)))
	if err != nil {
		t.Fatal(err)
	}
	var outputs []any
	for _, name := range []string{"plan", "mechanics", "review"} {
		output, err := fixture.JSON(name)
		if err != nil {
			t.Fatal(err)
		}
		outputs = append(outputs, output)
	}
	model := &fixture.Model{Outputs: outputs}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	result, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	library, _ := open(t)
	entry, err := library.Save(s.FromGoValue(result))
	if err != nil {
		t.Fatal(err)
	}
	snapshot := s.Stringify(s.FromGoValue(policy))
	findings := s.Stringify(s.FromGoValue(checked.Findings))

	keeps := func(when string) {
		t.Helper()
		loaded, err := library.Load(entry.ID)
		if err != nil {
			t.Fatalf("%s: %v", when, err)
		}
		retained := s.Stringify(s.FromGoValue(result.Prepared.Request.MechanicsDefinition.Profile.DesignPolicy))
		stored := loaded
		for _, key := range []string{"prepared", "request", "mechanicsDefinition", "profile", "designPolicy"} {
			stored, _ = stored.(*s.Object).Get(key)
		}
		if retained != snapshot || s.Stringify(stored) != snapshot {
			t.Errorf("%s: the Result's design policy changed:\n%s\n%s", when, retained, s.Stringify(stored))
		}
		again, err := unit.CheckDraft(draft)
		if err != nil || s.Stringify(s.FromGoValue(again.Findings)) != findings {
			t.Errorf("%s: CheckDraft found something else: %v", when, err)
		}
	}
	keeps("before the edit")

	edited := saved
	yes := true
	editedPolicy := *saved.MechanicsDefinition.Profile.DesignPolicy
	editedPolicy.DistinctEarlyBenefits, editedPolicy.RequireTier3BehaviorChange = &yes, &yes
	edited.MechanicsDefinition.Profile.DesignPolicy = &editedPolicy
	if _, err := profiles.Save(s.FromGoValue(edited)); err != nil {
		t.Fatal(err)
	}
	current, err := profiles.Get("snapshot-copy")
	if err != nil || s.Stringify(s.FromGoValue(current.MechanicsDefinition.Profile.DesignPolicy)) == snapshot {
		t.Fatalf("the Profile edit was not saved: %v", err)
	}
	if again, err := unit.Prepare(s.FromGoValue(unit.ApplyProfile(request, current))); err != nil || again.InputHash == prepared.InputHash {
		t.Errorf("the edited Profile prepares the same Request: %v", err)
	}
	keeps("after the edit")

	if _, err := profiles.Delete("snapshot-copy"); err != nil {
		t.Fatal(err)
	}
	keeps("after the delete")
}
