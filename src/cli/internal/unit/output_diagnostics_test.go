package unit_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/research"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// testdataValue decodes a JSON file of testdata.
func testdataValue(t *testing.T, name string) any {
	t.Helper()
	data, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}
	value, err := s.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	return value
}

// escanorV34 is Escanor's v34 run (OPUS-NET-61-13): its prepared request
// under the Default, its retained plan and its design attempt 4, whose only
// schema issue is x-5-x's activeFollowUp radius of 0.
func escanorV34(t *testing.T) (unit.Prepared, any, unit.DesignPlan, any) {
	t.Helper()
	sources, err := research.ParseSources(testdataValue(t, "escanor.sources.json"))
	if err != nil {
		t.Fatal(err)
	}
	prepared, err := sources.Prepare(unit.DefaultProfile())
	if err != nil {
		t.Fatal(err)
	}
	planOutput := testdataValue(t, "escanor-v34.plan.json")
	var plan unit.DesignPlan
	if err := s.ParseInto(unit.DesignPlanSchemaV2, planOutput, &plan); err != nil {
		t.Fatal(err)
	}
	return prepared, planOutput, plan, testdataValue(t, "escanor-v34-design.output.json")
}

const escanorRadius = "paths.path2.tiers.tier5.activeFollowUp.radius: Too small: expected number to be >0"

// escanorSplash are the nine top-path builds whose splash has pierce 1 in
// the same design output.
var escanorSplash = []string{"3-0-0", "3-1-0", "3-2-0", "4-0-0", "4-1-0", "4-2-0", "5-0-0", "5-1-0", "5-2-0"}

// A schema error no longer hides the draft's semantic issues: Escanor's v34
// design attempt reported only its radius, and the nine latent splash and
// pierce issues surfaced after the repair, with no repair left
// (OPUS-NET-61-14). A number outside its bound is read in bounds for the
// other checks, and the attempt still fails on it.
func TestSchemaErrorKeepsSemanticIssues(t *testing.T) {
	prepared, _, plan, output := escanorV34(t)
	issues, valid := unit.MechanicsIssues(output, &prepared.Request, plan)
	if valid {
		t.Fatal("an output with a radius of 0 passed its schema")
	}
	if len(issues) != 1+len(escanorSplash) || issues[0] != escanorRadius+". The other checks ran on every purchase and build that does not include this value." {
		t.Fatalf("issues: %q", issues)
	}
	for i, build := range escanorSplash {
		if !strings.HasPrefix(issues[i+1], "builds."+build+".baseAttack.stats.splashRadius: Splash requires pierce of at least 2") {
			t.Errorf("issue %d: %s", i+1, issues[i+1])
		}
	}

	// Any other schema issue is reported alone, as before: the typed
	// checks need a complete blueprint.
	missing := s.Clone(output).(*s.Object)
	at(missing, "paths", "path2", "tiers", "tier5", "activeFollowUp").(*s.Object).Delete("inheritStatuses")
	issues, valid = unit.MechanicsIssues(missing, &prepared.Request, plan)
	if valid || len(issues) != 2 || issues[0] != escanorRadius || !strings.HasPrefix(issues[1], "paths.path2.tiers.tier5.activeFollowUp.inheritStatuses: ") {
		t.Errorf("a missing field: %q", issues)
	}
}

// usedRepertoire is a plan output whose repertoire keeps only the entries
// the base attack or a purchase adapts: an entry omittedTechniques omits too
// leaves the repertoire, and any other moves to omittedTechniques.
func usedRepertoire(t *testing.T, output any, request *unit.Request) any {
	t.Helper()
	fixed := s.Clone(output).(*s.Object)
	var plan unit.DesignPlan
	if err := s.ParseInto(unit.DesignPlanSchemaV2, output, &plan); err != nil {
		t.Fatal(err)
	}
	unused := map[int]bool{}
	for _, issue := range unit.RepertoireUseIssues(plan, request) {
		var index int
		if _, err := fmt.Sscanf(issue.Path, "repertoire.%d", &index); err != nil {
			t.Fatal(err)
		}
		unused[index] = true
	}
	repertoire, omitted := []any{}, at(fixed, "omittedTechniques").([]any)
	for index, entry := range at(fixed, "repertoire").([]any) {
		switch {
		case !unused[index]:
			repertoire = append(repertoire, entry)
		case !strings.Contains(issueFor(t, plan, request, index), "omits it too"):
			omitted = append(omitted, s.NewObject().Set("name", at(entry, "name")).Set("importance", at(entry, "importance")).
				Set("reason", "Its daylight growth needs a time-of-day rule the Definition lacks."))
		}
	}
	return fixed.Set("repertoire", repertoire).Set("omittedTechniques", omitted)
}

func issueFor(t *testing.T, plan unit.DesignPlan, request *unit.Request, index int) string {
	t.Helper()
	for _, issue := range unit.RepertoireUseIssues(plan, request) {
		if issue.Path == fmt.Sprintf("repertoire.%d", index) {
			return issue.Message
		}
	}
	return ""
}

// Escanor's v34 plan passed with seven repertoire entries no purchase
// adapts; six are in omittedTechniques too. Each is now an issue.
func TestEscanorV34UnusedEntries(t *testing.T) {
	prepared, output, plan, _ := escanorV34(t)
	var got []string
	for _, issue := range unit.RepertoireUseIssues(plan, &prepared.Request) {
		got = append(got, issue.Path+" "+fmt.Sprint(strings.Contains(issue.Message, "omits it too")))
	}
	if want := "repertoire.6 false,repertoire.8 true,repertoire.10 true,repertoire.11 true,repertoire.12 true,repertoire.13 true,repertoire.14 true"; strings.Join(got, ",") != want {
		t.Errorf("unused entries %v", got)
	}
	if _, err := unit.DecodeDesignPlan(usedRepertoire(t, output, &prepared.Request), &prepared.Request); err != nil {
		t.Errorf("the plan with its repertoire used: %v", err)
	}
}

// Drafting sends every issue to the first repair together, and records them
// on the design attempt.
func TestFirstRepairSeesEveryIssue(t *testing.T) {
	prepared, planOutput, _, output := escanorV34(t)
	model := &fixture.Model{Outputs: []any{usedRepertoire(t, planOutput, &prepared.Request), output, output}}
	_, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	var modelErr *unit.ModelError
	if !errors.As(err, &modelErr) || modelErr.Evidence == nil || len(model.Requests) != 3 {
		t.Fatalf("draft: %v after %d calls", err, len(model.Requests))
	}
	attempts := modelErr.Evidence.Attempts
	if len(attempts) != 3 || attempts[0].Purpose != "plan" || len(attempts[0].Issues) != 0 || attempts[1].Purpose != "design" {
		t.Fatalf("attempts %+v", attempts)
	}
	design := strings.Join(attempts[1].Issues, "\n")
	repair := model.Requests[2].Prompt
	for _, want := range append([]string{escanorRadius}, "builds.3-0-0.baseAttack.stats.splashRadius", "builds.5-2-0.baseAttack.stats.splashRadius") {
		if !strings.Contains(design, want) {
			t.Errorf("the design attempt's issues lack %q", want)
		}
		if !strings.Contains(repair, want) {
			t.Errorf("the first repair lacks %q", want)
		}
	}
	// The repair sees the output as the model wrote it, radius 0.
	if !strings.Contains(repair, `"radius":0,`) {
		t.Error("the repair saw the diagnostic copy")
	}
}

// A repair patch whose only schema issues are numbers outside their bounds
// is merged as written, so the merged output reports them with the semantic
// issues; a patch with any other schema issue is rejected as before.
func TestRepairPatchWithNumberOutOfBounds(t *testing.T) {
	prepared, _, plan, output := escanorV34(t)
	valid := unit.BindDesignPlan(s.Clone(output), plan)
	follow := at(valid, "paths", "path2", "tiers", "tier5", "activeFollowUp").(*s.Object)
	follow.Set("radius", 1.0)
	repair, err := unit.TargetedTierRepair(&prepared.Request, valid, []string{"paths.path2.tiers.tier5: Rebuild this capstone."})
	if err != nil || repair == nil {
		t.Fatalf("repair %v, %v", repair, err)
	}
	tier := s.Clone(at(output, "paths", "path2", "tiers", "tier5")).(*s.Object)
	patch := s.NewObject().Set("paths", s.NewObject().Set("path2", s.NewObject().Set("tiers", s.NewObject().Set("tier5", tier))))
	merged, err := repair.Apply(patch)
	if err != nil {
		t.Fatalf("a patch with a radius of 0: %v", err)
	}
	issues, ok := unit.MechanicsIssues(merged, &prepared.Request, plan)
	if ok || len(issues) != 1+len(escanorSplash) || issues[0] != escanorRadius+". The other checks ran on every purchase and build that does not include this value." {
		t.Errorf("the merged repair's issues: %q", issues)
	}
	at(tier, "activeFollowUp").(*s.Object).Delete("inheritStatuses")
	if _, err := repair.Apply(patch); err == nil {
		t.Error("a patch with a missing field was merged")
	}
}

// Only per-build checks run on the copy in bounds. A check that compares
// purchases or paths, such as exclusive early benefits, can consume a
// substituted value and report it on another purchase, so it waits for an
// output that passes its schema, while every build that does not own the
// invalid purchase is still judged (SOL-73-01).
func TestCrossPurchaseIssuesWaitForAValidOutput(t *testing.T) {
	prepared, _, plan, output := escanorV34(t)
	tier := func(value any, path, name string) *s.Object {
		paths, _ := value.(*s.Object).Get("paths")
		p, _ := paths.(*s.Object).Get(path)
		tiers, _ := p.(*s.Object).Get("tiers")
		found, _ := tiers.(*s.Object).Get(name)
		return found.(*s.Object)
	}
	// path3's first two purchases buy exactly what path1's do.
	shared := s.Clone(output)
	for _, name := range []string{"tier1", "tier2"} {
		changes, _ := tier(shared, "path1", name).Get("statChanges")
		tier(shared, "path3", name).Set("statChanges", s.Clone(changes))
	}
	crossPurchase := func(issues []string) []string {
		var found []string
		for _, issue := range issues {
			if !strings.HasPrefix(issue, "builds.") && !strings.HasPrefix(issue, "paths.path2.tiers.tier5.activeFollowUp.radius") {
				found = append(found, issue)
			}
		}
		return found
	}
	valid := s.Clone(shared)
	followUp, _ := tier(valid, "path2", "tier5").Get("activeFollowUp")
	followUp.(*s.Object).Set("radius", 1.0)
	issues, _ := unit.MechanicsIssues(valid, &prepared.Request, plan)
	if len(crossPurchase(issues)) == 0 {
		t.Fatalf("the valid output reports no cross-purchase issue: %q", issues)
	}
	issues, ok := unit.MechanicsIssues(shared, &prepared.Request, plan)
	if ok || issues[0] != escanorRadius+". The other checks ran on every purchase and build that does not include this value." {
		t.Fatalf("the schema issue: %q", issues)
	}
	if found := crossPurchase(issues); len(found) != 0 {
		t.Errorf("a cross-purchase issue survived the invalid number: %q", found)
	}
	for _, build := range escanorSplash {
		if !slices.ContainsFunc(issues, func(issue string) bool { return strings.HasPrefix(issue, "builds."+build+".") }) {
			t.Errorf("build %s lost its splash issue", build)
		}
	}
}
