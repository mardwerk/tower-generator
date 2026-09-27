package unit_test

import (
	"context"
	"errors"
	"slices"
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
		`"requiredConcepts":[{"name":"Fan Club","reason":"The owner's signature support form.","baseAttack":false,"entries":["Fan Club"],"sourceIds":["source1:12"],"adaptedBy":["x-4-x","x-5-x"],"purchases":[` +
			`{"build":"x-4-x","name":"Super Monkey Fan Club","changeScope":{"base":["intervalSeconds multiply 0.5"],"boost":["unlocks the Active Ability Fan Club Frenzy"]},` +
			`"dimensions":[{"dimension":"intervalSeconds lowered","alsoChangedBy":["Spiked Ball","Dart Throw","Triple Throw","Crossbow"]},{"dimension":"Active Ability","alsoChangedBy":[]}]},` +
			`{"build":"x-5-x","name":"Plasma Monkey Fan Club","changeScope":{"base":[],"boost":["boost damageMultiplier set 2","boost durationSeconds add 5"]},` +
			`"dimensions":[{"dimension":"boost damageMultiplier set","alsoChangedBy":[]},{"dimension":"boost durationSeconds raised","alsoChangedBy":[]}]}]}]`,
		"requiredConcepts lists the concepts the owner requires this character's Unit to adapt",
		unit.RequiredConceptRule,
		`"requiredConcepts":[{"concept":"Fan Club"}]`,
		"fail when the Unit adapts it only in name, only for a peripheral effect or only with stats its cited passages do not make its identifying effect",
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

// proposedFanClub drafts and checks the fixture with Fan Club required. Its
// effects are left unadapted, so only x-5-x's proposed Plasma
// transformation adapts it; core ranks it core, and Crossbow major,
// instead of the recorded major.
func proposedFanClub(t *testing.T, core bool) unit.Checked {
	t.Helper()
	plain, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	request := plain.Request
	request.RequiredConcepts = []unit.RequiredConcept{{Name: "Fan Club", Reason: "The owner's signature support form."}}
	prepared, err := unit.Prepare(s.FromGoValue(request))
	if err != nil {
		t.Fatal(err)
	}
	output := recordedOutput(t, "plan")
	if core {
		repertoireEntry(output, "Crossbow").Set("importance", "major")
		repertoireEntry(output, "Fan Club").Set("importance", "core")
	}
	unadapt(repertoireEntry(output, "Fan Club"))
	at(output, "paths", "path2", "milestones", "tier5").(*s.Object).Set("proposedMechanics", []any{plasmaTransformation()})
	model := &fixture.Model{Outputs: []any{output, recordedOutput(t, "mechanics")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if issues := unit.RequiredConceptIssues(*draft.Run.DesignPlan, &draft.Prepared.Request); len(issues) > 0 {
		t.Fatalf("the plan fails the required concept check: %v", issues)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	return checked
}

// gapFindings are a checked artifact's core concept and required concept
// Findings.
func gapFindings(checked unit.Checked) []unit.Finding {
	var out []unit.Finding
	for _, f := range checked.Findings {
		if f.Rule == unit.CoreConceptFindingRule || f.Rule == unit.RequiredConceptFindingRule {
			out = append(out, f)
		}
	}
	return out
}

// withPlasmaVerdict gives a scripted review its proposal verdict on x-5-x's
// Plasma transformation, the only proposed mechanic.
func withPlasmaVerdict(review *s.Object) *s.Object {
	return withProposalVerdicts(review, proposalVerdict{"x-5-x", "Plasma transformation", "unresolved", "Transforming nearby Dart Monkeys is the Fan Club's allied effect, a coherent capability the Definition cannot express."})
}

// fanClubVerdict is the review's verdict on the required Fan Club.
func fanClubVerdict(concept, outcome string) *s.Object {
	return s.NewObject().Set("concept", concept).Set("outcome", outcome).
		Set("reason", "Only x-5-x's proposed Plasma transformation carries Fan Club's central effect, transforming nearby Dart Monkeys; x-4-x and x-5-x otherwise only attack faster.").
		Set("action", "Expand the Definition with an allied transformation, or adapt the frenzy on a typed change.").
		Set("evidence", []any{"dart-monkey-atlas-56-3"})
}

// A required major concept carried only by a proposed mechanic passes the
// plan check, which accepts a proposal, but the checked Unit reports it
// once as an unresolved design gap under its own rule, since checkCoreConcepts
// judges only core entries. The review gives it a required verdict,
// recorded as a model Finding after the omission verdicts; a review that
// leaves it out, or judges a concept the Request does not require, is
// rejected (SOL-76-01).
func TestRequiredMajorConceptOnlyProposed(t *testing.T) {
	checked := proposedFanClub(t, false)
	gaps := gapFindings(checked)
	if len(gaps) != 1 || gaps[0].Rule != unit.RequiredConceptFindingRule || gaps[0].Outcome != "unresolved" || gaps[0].Severity != "warning" ||
		gaps[0].Category != "missing_specification" || gaps[0].Subject != "paths.path2.tiers.tier5" ||
		!strings.HasPrefix(gaps[0].Message, `The required concept "Fan Club" is only proposed: x-5-x Plasma Monkey Fan Club carries it as the proposed mechanic Plasma transformation, and no purchase whose technique it is adapts one of its effects with a typed change. `) ||
		!strings.HasSuffix(gaps[0].Message, "until the Definition supports it, no build grants it and the Unit does not embody this required concept. "+unit.RequiredConceptRule) ||
		strings.Join(gaps[0].Evidence, ",") != "dart-monkey-atlas-56-3" {
		t.Fatalf("gap findings: %+v", gaps)
	}

	request := unit.BlueprintReviewRequest(checked)
	for _, want := range []string{
		`"requiredVerdicts":{"omissions":[{"technique":"Allied Fan Club","importance":"minor"},{"technique":"Critical shots","importance":"minor"}],"requiredConcepts":[{"concept":"Fan Club"}],"thirdPurchases":`,
		"in requiredConceptVerdicts, one verdict per concept by the name requiredVerdicts.requiredConcepts gives, apart from the eight findings",
		"fail when the Unit adapts it only in name, only for a peripheral effect or only with stats its cited passages do not make its identifying effect",
		"unresolved, never pass, when only a coherent, source-fitting proposed mechanic carries its central effect, a design gap until the Definition supports it",
	} {
		if !strings.Contains(request.Prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	if schema := s.Stringify(request.Schema); !strings.Contains(schema, `"requiredConceptVerdicts":{"minItems":1,"maxItems":1,"type":"array","items":{"type":"object","properties":{"concept":{"type":"string","enum":["Fan Club"]}`) ||
		!strings.Contains(schema, `"required":["summary","findings","omissionVerdicts","requiredConceptVerdicts","thirdPurchaseVerdicts","fifthPurchaseVerdicts","proposalVerdicts"]`) {
		t.Errorf("the review schema does not require the verdict: %s", schema)
	}

	// A free-form finding that repeats the verdict's subject, outcome and
	// reason is dropped; one on the concept with its own reason is kept.
	output := withPlasmaVerdict(recordedOutput(t, "review"))
	verdictValue := fanClubVerdict("Fan Club", "unresolved")
	output.Set("requiredConceptVerdicts", []any{verdictValue})
	findings := at(output, "findings").([]any)
	finding := func(id, subject, message string) *s.Object {
		return s.Clone(findings[0]).(*s.Object).Set("id", id).Set("category", "coverage").Set("severity", "warning").Set("outcome", "unresolved").
			Set("subject", subject).Set("message", message).Set("rule", "source-fit")
	}
	output.Set("findings", append(findings,
		finding("model.fan-club-repeat", "Required concept: Fan Club", at(verdictValue, "reason").(string)),
		finding("model.fan-club-name", "Fan Club", "x-4-x keeps the Fan Club name while only raising attack rate; name it for its frenzy."),
	))
	result, err := unit.ReviewDraft(context.Background(), checked, &fixture.Model{Outputs: []any{output}}, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	verdicts := verdictFindings(result)
	var ids []string
	for _, f := range verdicts {
		ids = append(ids, f.ID)
	}
	if got := strings.Join(ids[:4], ","); got != "verdict.omission.1,verdict.omission.2,verdict.required-concept.1,verdict.third-purchase.path1" {
		t.Errorf("verdict order %s", got)
	}
	verdict := verdicts[2]
	if verdict.Rule != unit.RequiredConceptVerdictRule || verdict.Method != "model" || verdict.Outcome != "unresolved" || verdict.Severity != "warning" ||
		verdict.Category != "coverage" || verdict.Subject != "requiredConcepts, Fan Club" || verdict.Action == nil ||
		!strings.HasPrefix(verdict.Message, "Only x-5-x's proposed Plasma transformation carries") {
		t.Errorf("the required concept verdict %+v", verdict)
	}
	gapCount := 0
	for _, f := range result.Findings {
		if f.Rule == unit.RequiredConceptFindingRule {
			gapCount++
		}
	}
	if gapCount != 1 {
		t.Errorf("the Result has %d required concept gaps", gapCount)
	}
	var kept []string
	for _, f := range result.Findings {
		if f.Method == "model" && !unit.IsReviewVerdict(f) {
			kept = append(kept, f.ID)
		}
	}
	if got := strings.Join(kept, ","); got != "model.fan-club-allies,model.fan-club-name" {
		t.Errorf("model findings %s", got)
	}
	if reread, err := unit.ParseResult(s.FromGoValue(result)); err != nil || len(verdictFindings(reread)) != len(verdicts) {
		t.Errorf("the Result does not read back with its verdicts: %v", err)
	}

	cases := map[string]struct {
		verdicts []any
		want     string
	}{
		"no verdict":         {nil, `no verdict on the required concept "Fan Club"`},
		"an unknown concept": {[]any{fanClubVerdict("Fan Club", "unresolved"), fanClubVerdict("Gear 5", "fail")}, `a verdict on "Gear 5", which requiredConcepts does not list`},
		"a repeated verdict": {[]any{fanClubVerdict("Fan Club", "unresolved"), fanClubVerdict("fan club", "fail")}, `2 verdicts on the required concept "Fan Club"`},
	}
	for name, c := range cases {
		output := withPlasmaVerdict(recordedOutput(t, "review"))
		if c.verdicts != nil {
			output.Set("requiredConceptVerdicts", c.verdicts)
		}
		model := &fixture.Model{Outputs: []any{output}}
		_, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options())
		var failure *unit.ModelError
		if !errors.As(err, &failure) || failure.Failure == nil || failure.Failure.Code != unit.CodeOutputInvalid || failure.Failure.Stage != "review" ||
			!strings.Contains(failure.Message, "per proposed mechanic: it gave "+c.want) || len(model.Requests) != 1 {
			t.Errorf("%s: %v after %d calls", name, err, len(model.Requests))
		}
	}
}

// A required concept that is also a core entry, carried only by a
// proposed mechanic, has exactly one gap Finding: the core concept's.
func TestRequiredCoreConceptOnlyProposedHasOneGap(t *testing.T) {
	gaps := gapFindings(proposedFanClub(t, true))
	if len(gaps) != 1 || gaps[0].Rule != unit.CoreConceptFindingRule || gaps[0].Outcome != "unresolved" ||
		!strings.HasPrefix(gaps[0].Message, `The core concept "Fan Club" is only proposed: `) {
		t.Errorf("gap findings: %+v", gaps)
	}
}

// luffyV38Checked is the v38 Luffy Result 55b028fa as a checked draft, with
// Gear 4 and Gear 5 required as its Request required them (OPUS-NET-61-21).
func luffyV38Checked(t *testing.T) unit.Checked {
	t.Helper()
	return luffyChecked(t, "v38", unit.Run{ID: "7b135155-1602-4110-99cf-223cbb7ed4c6", ModelID: "openrouter:openai/gpt-6-luna", StartedAt: "2026-09-27T14:31:37.867Z", CompletedAt: "2026-09-27T14:32:47.283Z"}, "Gear 4", "Gear 5")
}

// The v38 Luffy review passed Gear 4 and Gear 5 as "increased power and
// speed", though x-x-1 to x-x-4 buy only pierce, damage and attack rate and
// x-x-5 damage, attack rate and a nearby follow-up, stats the Gear 3 and
// Gear 2 paths buy too, while the cited passages identify Gear 4 by
// redirected punches and Gear 5 by rubber properties it gives its
// surroundings (SOL-61-16). The review's context for each required concept
// now gives the passages its entry cites, its source technique's passages
// and each purchase that carries it with its typed changes, and the prompt
// asks for the effect that identifies the concept and a pass only when a
// named typed change carries it. A model's verdict cannot be tested here;
// this pins what the review is given and asked.
func TestRequiredConceptReviewGetsTheIdentifyingEffectEvidence(t *testing.T) {
	prompt, context := reviewContext(t, luffyV38Checked(t))
	concepts := at(context, "requiredConcepts").([]any)
	if len(concepts) != 2 {
		t.Fatalf("required concepts %s", s.Stringify(concepts))
	}
	for i, want := range []string{
		`{"name":"Gear 4","reason":"The owner holds it essential to the character.","baseAttack":false,"entries":["Gear 4"],` +
			`"sourceIds":["source4:9","source4:10","source4:13","source4:17","source4:18","source7:42","source7:47","source7:48","source7:49","source7:51","source7:52"],` +
			`"sourceTechnique":{"name":"Gear 4","passageIds":["source4:9","source4:13","source4:17","source5:18","source7:42","source7:46","source7:47","source7:48","source7:49","source7:51","source7:52","source7:89"]},` +
			`"adaptedBy":["x-x-1","x-x-2","x-x-3","x-x-4"],"purchases":[` +
			`{"build":"x-x-1","name":"Gear 4 Reach","changeScope":{"base":["pierce add 1"],"boost":[]},"dimensions":[{"dimension":"pierce raised","alsoChangedBy":["Gear 3"]}]},` +
			`{"build":"x-x-2","name":"Snakeman Reach","changeScope":{"base":["pierce add 1"],"boost":[]},"dimensions":[{"dimension":"pierce raised","alsoChangedBy":["Gear 3"]}]},` +
			`{"build":"x-x-3","name":"Snakeman Pistol","changeScope":{"base":["damage add 1","intervalSeconds multiply 0.85"],"boost":[]},` +
			`"dimensions":[{"dimension":"damage raised","alsoChangedBy":["Gear 3","Gear 2","Gear 5"]},{"dimension":"intervalSeconds lowered","alsoChangedBy":["Gear 2","Gear 5"]}]},` +
			`{"build":"x-x-4","name":"Gear 4 Power","changeScope":{"base":["damage add 3"],"boost":[]},"dimensions":[{"dimension":"damage raised","alsoChangedBy":["Gear 3","Gear 2","Gear 5"]}]}]}`,
		`{"name":"Gear 5","reason":"The owner holds it essential to the character.","baseAttack":false,"entries":["Gear 5"],` +
			`"sourceIds":["source5:9","source5:12","source5:18","source7:53","source7:54","source7:55"],` +
			`"sourceTechnique":{"name":"Gear 5","passageIds":["source5:9","source5:12","source5:18","source5:28","source7:53","source7:54","source7:55","source7:56"]},` +
			`"adaptedBy":["x-x-5"],"purchases":[` +
			`{"build":"x-x-5","name":"Gear 5 Follow-up","changeScope":{"base":["damage add 4","intervalSeconds multiply 0.8","follow-up Gear 5 follow-up strike: 1 hits of 0.5 damage within 6"],"boost":[]},` +
			`"dimensions":[{"dimension":"damage raised","alsoChangedBy":["Gear 3","Gear 2","Gear 4"]},{"dimension":"intervalSeconds lowered","alsoChangedBy":["Gear 2","Gear 4"]},{"dimension":"follow-up","alsoChangedBy":[]}]}]}`,
	} {
		if got := s.Stringify(concepts[i]); got != want {
			t.Errorf("required concept %d:\n got %s\nwant %s", i, got, want)
		}
	}

	// The passages that identify each form are among those its entry
	// cites, and the review reads their text.
	passages := map[string]string{}
	for _, passage := range at(context, "sourcePassages").([]any) {
		passages[at(passage, "id").(string)] = at(passage, "text").(string)
	}
	for id, want := range map[string]string{
		"source4:17": "redirect his punches during an attack",
		"source7:54": "grant the environment around him the same properties as rubber",
	} {
		if !strings.Contains(passages[id], want) {
			t.Errorf("the review does not read %s: %q", id, passages[id])
		}
	}

	// Every stat Gear 4's purchases change, and x-x-5's damage and attack
	// rate, some other path's purchases change too: shared stats, which
	// alone cannot prove either form. x-x-5's follow-up is its only change
	// of its own, and the review judges it against Gear 5's passages.
	stat := func(label any) string { return strings.Fields(label.(string))[0] }
	shared := map[string]bool{}
	for _, path := range at(context, "unit", "paths").([]any)[:2] {
		for _, tier := range at(path, "tiers").([]any) {
			for _, label := range at(tier, "changeScope", "base").([]any) {
				shared[stat(label)] = true
			}
		}
	}
	var own []string
	for _, concept := range concepts {
		for _, purchase := range at(concept, "purchases").([]any) {
			for _, label := range at(purchase, "changeScope", "base").([]any) {
				if !shared[stat(label)] {
					own = append(own, at(purchase, "build").(string)+" "+label.(string))
				}
			}
		}
	}
	if got := strings.Join(own, "; "); got != "x-x-5 follow-up Gear 5 follow-up strike: 1 hits of 0.5 damage within 6" {
		t.Errorf("changes no other path makes: %s", got)
	}

	for _, want := range []string{
		unit.CoreSpiritRule,
		"the passages those entries cite (sourceIds), the source technique it names with its passageIds, when sourceTechniques has one, and the purchases whose technique those entries are (adaptedBy), each with its name, changeScope and dimensions",
		"first name its central effect from its sourceIds and passageIds, against the character's other forms and techniques, then give pass only when a typed change of one of its purchases, or the base attack, carries that effect, naming in the reason that purchase by build code, that change as its changeScope lists it and the passage that makes it the concept's identifying effect",
		"fail when the Unit adapts it only in name, only for a peripheral effect or only with stats its cited passages do not make its identifying effect, naming that effect and the purchase that should carry it by build code",
		"unresolved, never pass, when only a coherent, source-fitting proposed mechanic carries its central effect",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
}

// The v38 Luffy x-2-x Soru only shortens the attack interval, and the Gear
// 2 passage it cites describes Soru as movement at disappearing speeds
// (OPUS-NET-61-21, W3). The review reads that passage beside the purchase's
// typed changes and applies the naming rule, which fails a source name its
// cited passages do not support for what the purchase does.
func TestReviewGetsTheNamingEvidenceForSoru(t *testing.T) {
	prompt, context := reviewContext(t, luffyV38Checked(t))
	soru := at(context, "unit", "paths").([]any)[1].(*s.Object)
	tier := at(soru, "tiers").([]any)[1]
	if at(tier, "name") != "Soru" || s.Stringify(at(tier, "changeScope")) != `{"base":["intervalSeconds multiply 0.8"],"boost":[]}` ||
		!strings.Contains(s.Stringify(at(tier, "sourceIds")), `"source2:13"`) {
		t.Fatalf("x-2-x %s", s.Stringify(tier))
	}
	for _, passage := range at(context, "sourcePassages").([]any) {
		if at(passage, "id") == "source2:13" && !strings.Contains(at(passage, "text").(string), "Soru, allowing him to move at disappearing speeds") {
			t.Errorf("source2:13 reads %q", at(passage, "text"))
		}
	}
	if !strings.Contains(prompt, unit.NamingRule) {
		t.Error("the review prompt lacks the naming rule")
	}
}

// luffyV39Checked is the v39 Luffy Result ce07c903, run b, as a checked
// draft, with Gear 4 and Gear 5 required as its Request required them: its
// plan and blueprint (testdata/luffy-v39b.*) are those of the checked draft
// rebuilt from it (ug-acc/data/runs/v39/luffy-b.checked.json, SHA-256
// d363ecfb961e45bf), which the low and medium reviews read (OPUS-NET-61-25).
func luffyV39Checked(t *testing.T) unit.Checked {
	t.Helper()
	return luffyChecked(t, "v39b", unit.Run{ID: "c5711ce1-2f62-4e60-947d-687793356dee", ModelID: "openrouter:openai/gpt-6-luna", StartedAt: "2026-09-27T15:20:04.168Z", CompletedAt: "2026-09-27T15:22:00.725Z"}, "Gear 4", "Gear 5")
}

// Both v39 Luffy reviews, at low and medium reasoning, passed Gear 4 on
// x-3-x's damage and attack rate and Gear 5 on damage, which Gum-Gum
// Pistol, Gear 3 and the other form's purchases raise too, and the medium
// review passed 5-x-x's 15,000-Gold knockback as control, though Blimps and
// Bosses ignore knockback and x-2-x adds as much for 200 Gold
// (OPUS-NET-61-25, SOL-61-20). Each purchase of a required concept now
// lists the dimensions its typed changes change with the other techniques
// whose purchases change them too, and each fifth purchase its status
// effects with the enemy properties immune to them. The cue is descriptive:
// it pins what the review is given and asked, not a verdict.
func TestReviewGetsSharedDimensionsAndStatusImmunities(t *testing.T) {
	prompt, context := reviewContext(t, luffyV39Checked(t))
	dimensions := map[string]string{}
	for _, concept := range at(context, "requiredConcepts").([]any) {
		for _, purchase := range at(concept, "purchases").([]any) {
			dimensions[at(concept, "name").(string)+" "+at(purchase, "build").(string)] = s.Stringify(at(purchase, "dimensions"))
		}
	}
	for key, want := range map[string]string{
		"Gear 4 x-3-x": `[{"dimension":"damage raised","alsoChangedBy":["Gum-Gum Pistol","Gear 3","Gear 5"]},{"dimension":"intervalSeconds lowered","alsoChangedBy":["Gum-Gum Pistol","Python"]}]`,
		"Gear 4 x-4-x": `[{"dimension":"Active Ability","alsoChangedBy":[]}]`,
		"Gear 5 x-x-3": `[{"dimension":"damage raised","alsoChangedBy":["Gum-Gum Pistol","Gear 3","Gear 4"]}]`,
		"Gear 5 x-x-4": `[{"dimension":"damage raised","alsoChangedBy":["Gum-Gum Pistol","Gear 3","Gear 4"]},{"dimension":"range raised","alsoChangedBy":["Gum-Gum Pistol"]}]`,
		"Gear 5 x-x-5": `[{"dimension":"damage raised","alsoChangedBy":["Gum-Gum Pistol","Gear 3","Gear 4"]},{"dimension":"splashRadius raised","alsoChangedBy":["Gear 3"]}]`,
	} {
		if dimensions[key] != want {
			t.Errorf("%s dimensions:\n got %s\nwant %s", key, dimensions[key], want)
		}
	}
	if len(dimensions) != 5 {
		t.Errorf("required concept purchases %v", dimensions)
	}

	fifth := at(context, "requiredVerdicts", "fifthPurchases").([]any)
	if at(fifth[0], "build") != "5-x-x" || s.Stringify(at(fifth[0], "statusEffects")) != `[{"effect":"Knockback","immune":["Blimp","Boss"]}]` {
		t.Errorf("5-x-x %s", s.Stringify(fifth[0]))
	}
	if facts := s.Stringify(at(fifth[0], "payoff", "facts")); !strings.Contains(facts, "Adding x-2-x to") || !strings.Contains(facts, "that 5-x-x adds for 15000 Gold") {
		t.Errorf("5-x-x facts %s", facts)
	}
	for _, capstone := range fifth[1:] {
		if at(capstone, "statusEffects") != nil {
			t.Errorf("%s changes no status: %s", at(capstone, "build"), s.Stringify(capstone))
		}
	}

	for _, want := range []string{unit.SharedDimensionRule, unit.CapstonePayoffRule, "each with its name, changeScope and dimensions, each dimension its typed changes change, such as damage raised or knockback, with alsoChangedBy, the other techniques whose purchases change it too (purchases)"} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
}

// scriptedLuffyV38Review is a scripted review of the v38 Luffy checked
// draft: the given required concept verdicts and a pass on every other
// subject requiredVerdicts lists. It is not a model's judgment.
func scriptedLuffyV38Review(t *testing.T, checked unit.Checked, required ...any) *s.Object {
	t.Helper()
	return scriptedReview(t, checked, nil, required)
}

// scriptedReview is a scripted review of a checked draft: the given
// omission verdicts, by technique, and required concept verdicts, a pass on
// every other omission and on every third and fifth purchase
// requiredVerdicts lists, and an unresolved verdict on every proposal. A
// Request without required concepts gets no requiredConceptVerdicts. It is
// not a model's judgment.
func scriptedReview(t *testing.T, checked unit.Checked, omissions map[string]*s.Object, required []any) *s.Object {
	t.Helper()
	_, context := reviewContext(t, checked)
	verdicts := at(context, "requiredVerdicts").(*s.Object)
	verdict := func(key, field, outcome string, subject any) *s.Object {
		return s.NewObject().Set(key, at(subject, field)).Set("outcome", outcome).Set("reason", "Scripted.").Set("action", nil).Set("evidence", []any{})
	}
	list := func(name, key, field string) []any {
		out := []any{}
		for _, subject := range at(verdicts, name).([]any) {
			out = append(out, verdict(key, field, "pass", subject))
		}
		return out
	}
	omitted := []any{}
	for _, subject := range at(verdicts, "omissions").([]any) {
		if given, ok := omissions[at(subject, "technique").(string)]; ok {
			omitted = append(omitted, given)
		} else {
			omitted = append(omitted, verdict("technique", "technique", "pass", subject))
		}
	}
	proposals := []any{}
	for _, subject := range at(verdicts, "proposals").([]any) {
		proposals = append(proposals, verdict("build", "build", "unresolved", subject).Set("proposal", at(subject, "proposal")))
	}
	review := s.NewObject().Set("summary", "Scripted review.").Set("findings", []any{}).
		Set("omissionVerdicts", omitted)
	if verdicts.Has("requiredConcepts") {
		review.Set("requiredConceptVerdicts", required)
	}
	return review.Set("thirdPurchaseVerdicts", list("thirdPurchases", "build", "build")).
		Set("fifthPurchaseVerdicts", list("fifthPurchases", "build", "build")).
		Set("proposalVerdicts", proposals)
}

// A fail verdict on a required form that its purchases carry only with
// shared stats is recorded as an open Finding under the verdict rule, its
// action naming the purchase that should carry the identifying effect. The
// verdicts are scripted as the clarified rule asks for them on the v38
// Luffy Unit (SOL-61-16); they show what the Result keeps, not what a model
// returns.
func TestStatOnlyRequiredFormFailVerdictIsRecorded(t *testing.T) {
	checked := luffyV38Checked(t)
	gear4 := "character-technique:onepiece.fandom.com:Gomu%20Gomu%20no%20Mi%2FGear%204%20Techniques"
	gear5 := "character-technique:onepiece.fandom.com:Gomu%20Gomu%20no%20Mi%2FGear%205%20Techniques"
	verdict := func(concept, reason, action, evidence string) *s.Object {
		return s.NewObject().Set("concept", concept).Set("outcome", "fail").Set("reason", reason).Set("action", action).Set("evidence", []any{evidence})
	}
	output := scriptedLuffyV38Review(t, checked,
		verdict("Gear 4", "source4:17 identifies Gear 4 by punches redirected during an attack; x-x-1 to x-x-4 buy only pierce, damage and attack rate, which the Gear 3 and Gear 2 paths buy too, so no typed change carries it.", "Carry the redirected punch on x-x-3 with a typed change, such as a bounded follow-up, or propose it there.", gear4),
		verdict("Gear 5", "source7:54 identifies Gear 5 by the rubber properties it gives its surroundings; x-x-5 buys damage, attack rate and a nearby follow-up that no cited passage ties to Gear 5.", "Carry that effect on x-x-5 with a typed change, such as knockback, or propose it there.", gear5),
	)
	result, err := unit.ReviewDraft(context.Background(), checked, &fixture.Model{Outputs: []any{output}}, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	var got []unit.Finding
	for _, f := range verdictFindings(result) {
		if f.Rule == unit.RequiredConceptVerdictRule {
			got = append(got, f)
		}
	}
	if len(got) != 2 {
		t.Fatalf("required concept verdicts %+v", got)
	}
	for i, want := range []struct{ id, subject, build string }{
		{"verdict.required-concept.1", "requiredConcepts, Gear 4", "x-x-3"},
		{"verdict.required-concept.2", "requiredConcepts, Gear 5", "x-x-5"},
	} {
		f := got[i]
		if f.ID != want.id || f.Subject != want.subject || f.Outcome != "fail" || f.Severity != "error" || f.Method != "model" ||
			f.Action == nil || !strings.Contains(*f.Action, want.build) || !strings.Contains(f.Message, "identifies") {
			t.Errorf("verdict %d: %+v", i, f)
		}
	}
}

// A required concept the base attack adapts cites the base attack's
// passages, so a pass verdict can name one even without a repertoire entry.
// The review may cite the identifying passage from either the source
// technique's passages or the concept's sourceIds, and gives unresolved only
// when neither establishes the effect (SOL-80-01, SOL-80-03).
func TestBaseAttackRequiredConceptCarriesItsPassages(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	checked := stages.Checked
	base := checked.Draft.Run.DesignPlan.Base
	checked.Draft.Prepared.Request.RequiredConcepts = []unit.RequiredConcept{{Name: base.Name}}
	prompt, context := reviewContext(t, checked)
	concept := at(context, "requiredConcepts").([]any)[0].(*s.Object)
	if value, _ := concept.Get("baseAttack"); value != true {
		t.Fatalf("%s is not the base attack: %s", base.Name, s.Stringify(concept))
	}
	ids, _ := concept.Get("sourceIds")
	for _, id := range base.SourceIDs {
		if !slices.Contains(ids.([]any), any(id)) {
			t.Errorf("the base concept's sourceIds %v lack the base attack's %s", ids, id)
		}
	}
	for _, want := range []string{
		"citing it from sourceTechnique.passageIds or from its sourceIds, whichever holds the identifying passage, when both exist too; its sourceIds include the base attack's passages when the base attack adapts it",
		"and unresolved when neither sourceTechnique nor sourceIds supplies a passage that identifies the concept, saying that the supplied sources do not establish its central effect.",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
}
