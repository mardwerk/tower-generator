package unit_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// requiring is a request with the given required concepts.
func requiring(request unit.Request, names ...string) unit.Request {
	request.RequiredConcepts = nil
	for _, name := range names {
		request.RequiredConcepts = append(request.RequiredConcepts, unit.RequiredConcept{Name: name, Reason: "The owner holds it essential to the character."})
	}
	return request
}

// gear4Plan is the v36 Luffy plan with its bottom path's technique renamed
// from "Gear 4 Python" to "Gear 4", listed once and no longer omitted. Its
// purchases keep what they did: range, then attack rate, then a follow-up.
func gear4Plan(t *testing.T) *s.Object {
	t.Helper()
	plan := luffyV36Plan(t)
	repertoire, _ := plan.Get("repertoire")
	entries := repertoire.([]any)[:7]
	entries[3].(*s.Object).Set("name", "Gear 4")
	plan.Set("repertoire", entries)
	omitted, _ := plan.Get("omittedTechniques")
	var kept []any
	for _, omission := range omitted.([]any) {
		if at(omission, "name") != "Gear 4" {
			kept = append(kept, omission)
		}
	}
	plan.Set("omittedTechniques", kept)
	intents, _ := plan.Get("upgradeIntents")
	bottom, _ := intents.(*s.Object).Get("path3")
	for _, tier := range []string{"tier1", "tier2", "tier3", "tier4", "tier5"} {
		intent, _ := bottom.(*s.Object).Get(tier)
		intent.(*s.Object).Set("technique", "Gear 4")
	}
	return plan
}

// requiredIssues are a plan's issues that state the required concept rule.
func requiredIssues(t *testing.T, output any, request unit.Request) []string {
	t.Helper()
	var out []string
	for _, issue := range planIssues(t, output, &request) {
		if strings.Contains(issue, unit.RequiredConceptRule) {
			out = append(out, issue)
		}
	}
	return out
}

// With Gear 4 and Gear 5 required, the v36 Luffy plan fails: it omits both
// whole and lists neither in its repertoire, since "Gear 4 Python" is not
// named for Gear 4 (SOL-61-13).
func TestPlanOmittingARequiredConceptFails(t *testing.T) {
	request := requiring(savedSources(t, "luffy"), "Gear 4", "Gear 5")
	got := requiredIssues(t, luffyV36Plan(t), request)
	want := []string{
		`custom omittedTechniques.8: omittedTechniques omits "Gear 4", the required concept "Gear 4", whole. ` + unit.RequiredConceptRule + ` Remove it from omittedTechniques, list it in the repertoire and adapt its central effect on a purchase whose technique it is, as a typed change or a proposed mechanic; name each aspect the Definition cannot express as an effect of that entry with adaptedAs empty.`,
		`custom repertoire: No repertoire entry is named for the required concept "Gear 4". ` + unit.RequiredConceptRule + ` Add it to the repertoire by the name "Gear 4", or by the name or an alias sourceTechniques gives it, citing its passages, and make it the technique of a purchase that adapts its central effect.`,
		`custom omittedTechniques.14: omittedTechniques omits "Gear 5", the required concept "Gear 5", whole.`,
		`custom repertoire: No repertoire entry is named for the required concept "Gear 5".`,
	}
	if len(got) != len(want) {
		t.Fatalf("issues:\n%s", strings.Join(got, "\n"))
	}
	for i := range want {
		if !strings.HasPrefix(got[i], want[i]) {
			t.Errorf("issue %d:\n got %s\nwant %s", i, got[i], want[i])
		}
	}
	// Without required concepts the plan has no such issue.
	if got := requiredIssues(t, luffyV36Plan(t), savedSources(t, "luffy")); len(got) > 0 {
		t.Errorf("unrequired: %v", got)
	}
}

// A plan whose purchases adapt a required concept by name, with typed
// promises its effects are adapted as, passes the plan check, whether or
// not those promises carry its central effect: that is the review's
// judgment. The bottom path renamed Gear 4 buys range and attack rate,
// which Gear 4's bounce-powered blows do not describe, and code cannot tell.
func TestRequiredConceptAdaptedByNameIsLeftToTheReview(t *testing.T) {
	request := requiring(savedSources(t, "luffy"), "Gear 4")
	if got := planIssues(t, gear4Plan(t), &request); len(got) > 0 {
		t.Errorf("issues:\n%s", strings.Join(got, "\n"))
	}
}

// A required concept is adapted by a typed promise one of its effects is
// adapted as, or by a proposed mechanic of a purchase whose technique it
// is; a purchase that carries only its name, with neither, fails.
func TestRequiredConceptNeedsAnAdaptation(t *testing.T) {
	request := requiring(savedSources(t, "luffy"), "Gear 4")
	unadapted := gear4Plan(t)
	unadapt(repertoireEntry(unadapted, "Gear 4"))
	got := requiredIssues(t, unadapted, request)
	want := `custom repertoire.3: x-x-1, x-x-2, x-x-3, x-x-4 and x-x-5 adapt "Gear 4", the entry of the required concept "Gear 4", but promise nothing one of its effects is adapted as and carry no proposed mechanic. ` + unit.RequiredConceptRule + ` Adapt one of its effects as a promise of a purchase whose technique it is, or name the mechanic it needs in that milestone's proposedMechanics.`
	if len(got) != 1 || got[0] != want {
		t.Errorf("issues:\n%s", strings.Join(got, "\n"))
	}

	proposed := gear4Plan(t)
	unadapt(repertoireEntry(proposed, "Gear 4"))
	intents, _ := proposed.Get("upgradeIntents")
	bottom, _ := intents.(*s.Object).Get("path3")
	capstone, _ := bottom.(*s.Object).Get("tier5")
	capstone.(*s.Object).Set("proposedMechanics", []any{s.NewObject().
		Set("name", "Boundman bounce").
		Set("effect", "The inflated fist rebounds off the ground and strikes the next enemy in range once for full damage.").
		Set("sourceIds", []any{"source4:17"})})
	if got := requiredIssues(t, proposed, request); len(got) > 0 {
		t.Errorf("a proposed adaptation: %v", got)
	}
}

// A required concept named for a source technique's alias, or with a
// qualifier word, is that technique.
func TestRequiredConceptMatchesAliases(t *testing.T) {
	request := requiring(savedSources(t, "escanor"), "The Divine Axe Rhitta")
	plan := corePlan(unit.PlanBase{Name: "Sunshine"})
	plan.OmittedTechniques = []unit.PlanOmission{{Name: "Rhitta", Importance: "major", Reason: "An axe."}}
	issues := messages(unit.RequiredConceptIssues(plan, &request))
	if !strings.Contains(issues, `omittedTechniques.0: omittedTechniques omits "Rhitta", the required concept "The Divine Axe Rhitta", whole.`) {
		t.Errorf("the alias: %s", issues)
	}
}

// Required concepts reach the plan prompt and the review, with the
// entries and purchases code found for each; a Request without them keeps
// its hash and prompts.
func TestRequiredConceptsReachThePromptsAndKeepTheHash(t *testing.T) {
	plain, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	again, err := unit.Prepare(s.FromGoValue(plain.Request))
	if err != nil || again.InputHash != plain.InputHash || strings.Contains(s.Stringify(s.FromGoValue(plain)), "requiredConcepts") {
		t.Fatalf("a Request without required concepts: %v, %s and %s", err, again.InputHash, plain.InputHash)
	}
	planRequest, err := unit.DesignPlanRequest(plain)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(planRequest.Prompt, "requiredConcepts") || strings.Contains(planRequest.Prompt, unit.RequiredConceptRule) {
		t.Error("the plan prompt names required concepts the Request lacks")
	}

	request := plain.Request
	request.RequiredConcepts = []unit.RequiredConcept{{Name: "Fan Club", Reason: "The owner's signature support form."}}
	prepared, err := unit.Prepare(s.FromGoValue(request))
	if err != nil {
		t.Fatal(err)
	}
	if prepared.InputHash == plain.InputHash {
		t.Error("required concepts leave the hash unchanged")
	}
	planRequest, err = unit.DesignPlanRequest(prepared)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`"requiredConcepts":[{"name":"Fan Club","reason":"The owner's signature support form."}]`,
		unit.RequiredConceptRule + " requiredConcepts lists the concepts the owner requires",
		unit.CoreSpiritRule,
	} {
		if !strings.Contains(planRequest.Prompt, want) {
			t.Errorf("the plan prompt lacks %q", want)
		}
	}
	model := &fixture.Model{Outputs: []any{recordedOutput(t, "plan"), recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	review := unit.BlueprintReviewRequest(checked).Prompt
	for _, want := range []string{
		`"requiredConcepts":[{"name":"Fan Club","reason":"The owner's signature support form.","baseAttack":false,"entries":["Fan Club"],"adaptedBy":["x-4-x","x-5-x"]}]`,
		"requiredConcepts lists the concepts the owner requires this character's Unit to adapt",
		unit.RequiredConceptRule,
		"fail one whose purchases carry its name but not its central effect",
	} {
		if !strings.Contains(review, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(unit.BlueprintReviewRequest(stages.Checked).Prompt, unit.RequiredConceptRule) {
		t.Error("the review states the rule for a Request without required concepts")
	}

	// A concept named twice is rejected.
	request.RequiredConcepts = append(request.RequiredConcepts, unit.RequiredConcept{Name: "fan  club"})
	if _, err := unit.Prepare(s.FromGoValue(request)); err == nil || !strings.Contains(err.Error(), `"fan  club" is listed twice`) {
		t.Errorf("a repeated concept: %v", err)
	}
}
