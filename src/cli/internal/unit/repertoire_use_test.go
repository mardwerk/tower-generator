package unit_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// withEntry is the fixture plan with one more repertoire entry, which no
// milestone names as its technique.
func withEntry(t *testing.T, name, importance string) *s.Object {
	t.Helper()
	plan := recordedOutput(t, "plan")
	repertoire, _ := plan.Get("repertoire")
	plan.Set("repertoire", append(repertoire.([]any), s.NewObject().
		Set("name", name).Set("importance", importance).Set("sourceIds", []any{"source1:2"}).
		Set("limitation", "The ordinary throw is described, not a purchase.").
		Set("effects", []any{s.NewObject().Set("effect", "The dart flies to a distant target.").Set("adaptedAs", []any{}).Set("reason", "No purchase adapts it.")})))
	return plan
}

// Every repertoire entry, minor included, is the base attack or the
// technique of a purchase; Luffy's v34 plan ranked Gear 5 and three kinds of
// Haki major and adapted none, so no omission verdict judged them
// (OPUS-NET-61-13, SOL-61-10).
func TestRepertoireEntryIsUsed(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	for _, importance := range []string{"major", "minor"} {
		_, err := unit.DecodeDesignPlan(withEntry(t, "Dart Monkey Legend", importance), &prepared.Request)
		want := `repertoire.4: The repertoire entry "Dart Monkey Legend", ranked ` + importance + `, is neither the base attack nor the technique of any purchase. ` + unit.RepertoireUseRule
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("an unused %s entry: %v", importance, err)
		}
	}
	// An entry named like the base attack is adapted by 0-0-0.
	if _, err := unit.DecodeDesignPlan(withEntry(t, "dart throw", "major"), &prepared.Request); err != nil {
		t.Errorf("an entry the base attack adapts: %v", err)
	}
	// The plan prompt and the review state the rule.
	request, err := unit.DesignPlanRequest(prepared)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(request.Prompt, unit.RepertoireUseRule) != 1 {
		t.Errorf("the plan prompt holds the rule %d times", strings.Count(request.Prompt, unit.RepertoireUseRule))
	}
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	review := unit.BlueprintReviewRequest(stages.Checked).Prompt
	for _, want := range []string{unit.RepertoireUseRule, "An entry whose adaptedBy is empty and that is not named like the base attack adapts nothing"} {
		if !strings.Contains(review, want) {
			t.Errorf("the review lacks %q", want)
		}
	}
}

// An unused entry is a narrow item: the targeted correction moves it to
// omittedTechniques with its reason or names it as a purchase's technique,
// and the moved entry then gets an omission verdict.
func TestTargetedCorrectionMovesUnusedEntry(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	corrected := recordedOutput(t, "plan")
	omitted, _ := corrected.Get("omittedTechniques")
	corrected.Set("omittedTechniques", append(omitted.([]any), s.NewObject().
		Set("name", "Dart Monkey Legend").Set("importance", "major").
		Set("reason", "Its far-flying dart needs a map-wide range the Definition caps.")))
	model := &fixture.Model{Outputs: []any{withEntry(t, "Dart Monkey Legend", "major"), corrected, recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if got := strings.Join(purposes(draft), ","); got != "plan,plan-correction,design" {
		t.Errorf("attempts %s", got)
	}
	if first := (*draft.Run.Attempts)[0]; len(first.Issues) != 1 || !strings.HasPrefix(first.Issues[0], "repertoire.4: ") {
		t.Errorf("the failed plan attempt's issues: %v", first.Issues)
	}
	prompt := model.Requests[1].Prompt
	for _, want := range []string{
		"The previous plan failed only on the items below.",
		`"path":"repertoire.4"`,
		`Move it from the repertoire to omittedTechniques by the name \"Dart Monkey Legend\", ranked major as it is now`,
		`Name \"Dart Monkey Legend\" as the technique of a milestone that adapts what it does`,
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the targeted correction lacks %q", want)
		}
	}
	if strings.Contains(prompt, "Correct this invalid design plan") {
		t.Error("the targeted correction also asks for a full re-plan")
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	if review := unit.BlueprintReviewRequest(checked).Prompt; !strings.Contains(review, `{"technique":"Dart Monkey Legend","importance":"major"}`) {
		t.Error("the moved entry gets no omission verdict")
	}
}

// Luffy's v34 plan (Result a6d93b5c, OPUS-NET-61-13) fails only on its four
// unused major entries, each an item of the targeted correction.
func TestLuffyV34PlanUnusedEntries(t *testing.T) {
	request := savedSources(t, "luffy")
	data, err := os.ReadFile("testdata/luffy-v34.plan.json")
	if err != nil {
		t.Fatal(err)
	}
	output, err := s.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	_, err = unit.DecodeDesignPlan(output, &request)
	if err == nil {
		t.Fatal("the v34 plan passes")
	}
	lines := strings.Split(err.Error(), "\n")
	unused := []string{"Gear 5", "Observation Haki", "Armament Haki", "Supreme King Haki"}
	if len(lines) != len(unused) {
		t.Errorf("the v34 plan's issues: %v", err)
	}
	for i, name := range unused {
		if want := `The repertoire entry "` + name + `", ranked major, is neither`; i >= len(lines) || !strings.Contains(lines[i], want) {
			t.Errorf("issue %d lacks %q", i, want)
		}
	}
	var plan unit.DesignPlan
	if err := s.ParseInto(unit.DesignPlanSchemaV2, output, &plan); err != nil {
		t.Fatal(err)
	}
	items := s.Stringify(s.FromGoValue(unit.CorrectionItems(plan, &request)))
	for i, name := range unused {
		if !strings.Contains(items, `"path":"repertoire.`+string(rune('4'+i))+`"`) || !strings.Contains(items, `by the name \"`+name+`\", ranked major as it is now`) {
			t.Errorf("the correction does not name %s", name)
		}
	}
}
