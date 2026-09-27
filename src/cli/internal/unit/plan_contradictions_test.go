package unit_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// luffyV36Plan is the retained plan of the v36 Luffy Result 706999ee
// (OPUS-NET-61-16), as plan output.
func luffyV36Plan(t *testing.T) *s.Object {
	t.Helper()
	data, err := os.ReadFile("testdata/luffy-v36.plan.json")
	if err != nil {
		t.Fatal(err)
	}
	output, err := s.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	return output.(*s.Object)
}

// planIssues are the issues DecodeDesignPlan reports for a plan output,
// each as its code, path and message.
func planIssues(t *testing.T, output any, request *unit.Request) []string {
	t.Helper()
	_, err := unit.DecodeDesignPlan(output, request)
	if err == nil {
		return nil
	}
	var validation *s.Error
	if !errors.As(err, &validation) {
		t.Fatal(err)
	}
	var out []string
	for _, issue := range validation.Issues {
		out = append(out, issue.Code+" "+issue.PathString()+": "+issue.Message)
	}
	return out
}

// parsedPlan reads plan output as a retained plan, without the plan checks.
func parsedPlan(t *testing.T, output any) unit.DesignPlan {
	t.Helper()
	var plan unit.DesignPlan
	if err := s.ParseInto(unit.DesignPlanSchemaV2, output, &plan); err != nil {
		t.Fatal(err)
	}
	return plan
}

// addOmission appends a whole-technique omission to plan output.
func addOmission(plan *s.Object, name, importance string) {
	omitted, _ := plan.Get("omittedTechniques")
	plan.Set("omittedTechniques", append(omitted.([]any), s.NewObject().Set("name", name).Set("importance", importance).Set("reason", "The Definition cannot express its central behavior in any form.")))
}

// The v36 Luffy plan lists "Gear 4 Python" twice (OPUS-NET-61-16). An exact
// duplicate fails, alone, as an item of the targeted correction. Its Gear 4
// and Python omissions are no exact names of anything it selects, so the
// broader contradiction stays the review's (SOL-61-13).
func TestLuffyV36PlanListsGear4PythonTwice(t *testing.T) {
	request := savedSources(t, "luffy")
	output := luffyV36Plan(t)
	issues := planIssues(t, output, &request)
	want := `plan-correctable repertoire.7: The repertoire lists "Gear 4 Python" twice, as repertoire.3 and repertoire.7. ` + unit.PlanNamesRule + ` Keep one entry named "Gear 4 Python", with the sourceIds and effects of both, and remove the other.`
	if len(issues) != 1 || issues[0] != want {
		t.Errorf("the v36 plan's issues:\n%s", strings.Join(issues, "\n"))
	}
	items := s.Stringify(s.FromGoValue(unit.CorrectionItems(parsedPlan(t, output), &request)))
	for _, want := range []string{
		`"path":"repertoire.7"`,
		`Keep one entry named \"Gear 4 Python\", with the sourceIds and effects of both, and remove the other; each milestone whose technique is \"Gear 4 Python\" keeps it.`,
	} {
		if !strings.Contains(items, want) {
			t.Errorf("the correction items lack %q: %s", want, items)
		}
	}
	if strings.Contains(items, `"path":"omittedTechniques.`) {
		t.Errorf("an inexact name is a correction item: %s", items)
	}

	// The next call is the targeted correction, not a full re-plan, and
	// its deduplicated plan is accepted: the third call is the mechanics
	// call, which this test does not script.
	prepared := savedPrepared(t, "luffy")
	fixed := luffyV36Plan(t)
	repertoire, _ := fixed.Get("repertoire")
	fixed.Set("repertoire", repertoire.([]any)[:7])
	model := &fixture.Model{Outputs: []any{luffyV36Plan(t), fixed}}
	_, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if len(model.Requests) != 3 || err == nil || !strings.Contains(model.Requests[2].Prompt, "Implement the supplied designPlan") {
		t.Fatalf("%d calls: %v", len(model.Requests), err)
	}
	if prompt := model.Requests[1].Prompt; !strings.Contains(prompt, "The previous plan failed only on the items below.") || strings.Contains(prompt, "Correct this invalid design plan") {
		t.Error("the duplicate is not corrected by the targeted correction")
	}
}

// savedPrepared prepares a Sources file of testdata under the Default.
func savedPrepared(t *testing.T, name string) unit.Prepared {
	t.Helper()
	prepared, err := unit.Prepare(s.FromGoValue(savedSources(t, name)))
	if err != nil {
		t.Fatal(err)
	}
	return prepared
}

// An omitted name the plan selects exactly, case and spacing aside, fails:
// a repertoire entry a purchase adapts, the base attack, or an entry no
// purchase adapts. A substring is no match.
func TestSelectedAndOmittedNameFails(t *testing.T) {
	request := savedSources(t, "luffy")
	for _, c := range []struct {
		name, omission, importance, want, item string
	}{
		{"a purchase technique", "gomu gomu no  jet pistol", "minor",
			`plan-correctable omittedTechniques.17: omittedTechniques omits "gomu gomu no  jet pistol", which the plan also selects as repertoire.4, the technique of x-3-x. ` + unit.PlanNamesRule + ` Remove it from omittedTechniques and name each aspect the Definition cannot express as an effect of its entry with adaptedAs empty and its reason.`,
			`Remove \"gomu gomu no  jet pistol\" from omittedTechniques`},
		{"a substring", "Jet", "minor", "", ""},
	} {
		output := luffyV36Plan(t)
		addOmission(output, c.omission, c.importance)
		var found []string
		for _, issue := range planIssues(t, output, &request) {
			if strings.Contains(issue, "omittedTechniques omits") {
				found = append(found, issue)
			}
		}
		if c.want == "" {
			if len(found) > 0 {
				t.Errorf("%s: %v", c.name, found)
			}
			continue
		}
		if len(found) != 1 || found[0] != c.want {
			t.Errorf("%s: got\n%s", c.name, strings.Join(found, "\n"))
		}
		if items := s.Stringify(s.FromGoValue(unit.CorrectionItems(parsedPlan(t, output), &request))); !strings.Contains(items, c.item) {
			t.Errorf("%s: the correction items lack %q", c.name, c.item)
		}
	}

	// The base attack, when no repertoire entry has its name.
	plan := parsedPlan(t, luffyV36Plan(t))
	plan.Base.Name = "Gum-Gum Rifle"
	for i := range plan.Repertoire {
		if plan.Repertoire[i].Name == "Gum-Gum Pistol" {
			plan.Repertoire[i].Name = "Gum-Gum Rifle Pistol"
		}
	}
	plan.OmittedTechniques = append(plan.OmittedTechniques, unit.PlanOmission{Name: "Gum-Gum Rifle", Importance: "minor", Reason: "Not a mechanic."})
	issues := messages(unit.PlanContradictionIssues(plan, &request))
	if !strings.Contains(issues, `omittedTechniques.17: omittedTechniques omits "Gum-Gum Rifle", which the plan also selects as the base attack.`) || !strings.Contains(issues, "since 0-0-0 adapts it as the base attack") {
		t.Errorf("the base attack's omission: %s", issues)
	}

	// A core entry omitted whole is the core concept check's, reported
	// once.
	output := luffyV36Plan(t)
	addOmission(output, "Gear 2", "major")
	all := strings.Join(planIssues(t, output, &request), "\n")
	if strings.Contains(all, `omits "Gear 2", which the plan also selects`) || !strings.Contains(all, `omittedTechniques omits the core entry "Gear 2" whole.`) {
		t.Errorf("a core entry omitted whole: %s", all)
	}
}

// The plan prompt states the rule once, with what code rejects, and the
// review states it beside the broader contradiction it judges.
func TestPlanNamesRuleReachesPlanAndReview(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan, err := unit.DesignPlanRequest(stages.Prepared)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(plan.Prompt, unit.PlanNamesRule) != 1 || !strings.Contains(plan.Prompt, unit.PlanNamesRule+" Code rejects a plan whose repertoire lists a name twice") {
		t.Error("the plan prompt does not state the rule once")
	}
	review := unit.BlueprintReviewRequest(stages.Checked).Prompt
	if !strings.Contains(review, unit.PlanNamesRule+" Code has rejected a name repeated exactly; judge the broader contradiction") || !strings.Contains(review, `as "Gear 4 Python" beside omitted "Gear 4" and "Python"`) {
		t.Error("the review does not judge the broader contradiction")
	}
}
