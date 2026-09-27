package unit_test

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// The review reads every legal build as code resolved it, so it need not
// recompute a hit count or guess which builds exist (reported on #27: a
// review said pierce 2 cannot reach a second enemy and judged 3-3-0).
func TestReviewReadsResolvedLegalBuilds(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	facts := unit.LegalBuildFacts(stages.Checked.Draft.Candidate.Blueprint, unit.DefaultAuthoringDefinition())
	if len(facts) != 64 {
		t.Fatalf("%d legal builds", len(facts))
	}
	var base *s.Object
	for _, fact := range facts {
		if code, _ := fact.(*s.Object).Get("code"); code == "0-0-0" {
			base = fact.(*s.Object)
		}
	}
	if got := s.Stringify(base); base == nil || !strings.Contains(got, `"cost":200`) || !strings.Contains(got, `"pierce":2`) || !strings.Contains(got, `"maxEnemiesHitPerAttack":2`) {
		t.Errorf("0-0-0 facts %s", got)
	}
	prompt := unit.BlueprintReviewRequest(stages.Checked).Prompt
	if !strings.Contains(prompt, `"legalBuilds":[`) || !strings.Contains(prompt, `"code":"3-2-0"`) || strings.Contains(prompt, `"code":"3-3-0"`) || !strings.Contains(prompt, "such as 3-3-0, is illegal") {
		t.Error("the review prompt lacks the legal build facts")
	}
	// A copy bound is context for a capstone, not a verdict, the draft's
	// proposals are claims to check, not facts, and the period decision is
	// context, not text the Unit must print (reported on #27).
	for _, want := range []string{"A same-budget copy bound is context, not a verdict", "They are the drafting model's claims and can be wrong", "A name taken from the source", "not text the Unit must print",
		// The reviewer checks the planned adaptation the unit shows, as a claim.
		`"adaptation":"Replaces the dart with a heavier spiked ball that deals more damage and pierces far more enemies, at a slower throw."`,
		"It is a claim, not proof",
		// A path theme is judged across the path (reported on #27: a review
		// failed "Higher damage per hit" at x-4-x, which develops the boost).
		"judge them across its purchases together, not as a promise every purchase repeats",
		// Live Luffy runs on #27: compression credited with range, and a
		// Snakeman purchase whose text only said the punch hits harder.
		"a mechanism credited with an effect the source ties to another",
		// A TD adaptation of a cited effect needs no canon wording (a Luffy
		// review on #27 failed Gear 3's bounded splash for lacking it).
		"is a supported connection even when no source sentence names the mechanic", "not by canon wording",
		"is a fail when the cited evidence explains the technique", "unresolved only when the evidence lacks the detail to judge the connection",
		// Each purchased tier carries what the plan says about it and is
		// judged on its own build code (reported on #27: compression credited
		// with range at x-x-1 was flagged only in crosspath proposals).
		`"technique":"Dart Throw","sourceIds":["source1:2","source1:3"],"promises":{"improves":["pierce"],"unlock":"none"},"plannedChange":"1-x-x raises dart pierce by one."}`,
		`"technique":"Spiked Ball","sourceIds":["source1:6"],"promises":{"improves":["damage","pierce"],"unlock":"none","lowers":["attack-rate"]},"promisesWithoutEffect":[],"adaptation":"Replaces the dart`,
		"Judge every purchased tier on its own build code, the first and second purchases included",
		"against the passages its sourceIds cite", "which designPlan does not repeat",
		// Each technique's effects are judged against its passages.
		`"effects":[{"effect":"A heavier spiked ball replaces the dart and deals more damage.","adaptedAs":["damage"]`,
		"fail a described effect the list leaves out, an omission whose reason does not hold",
		// The review sees which entries code checked (Kyle on #27: an entry
		// no purchase names was skipped while the review was told code had
		// checked every adapted promise).
		`"adaptedBy":["3-x-x","4-x-x","5-x-x"]}`, `"adaptedBy":["x-4-x","x-5-x"]}`,
		"Code checks no entry whose adaptedBy is empty: fail a nonempty adaptedAs there",
		"except for an entry named like the base attack: judge that entry's adaptations against 0-0-0"} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
}

// The review reads each purchase's planned text once, on its tier. Before,
// the designPlan repeated every milestone and a compact plan's crosspath
// contributions repeated each first and second milestone as "Proposed early
// purchases", and a Luffy review on #27 flagged compression credited with
// range there instead of on the purchased x-x-1.
func TestReviewReadsEachPlannedPurchaseOnItsTier(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prompt := unit.BlueprintReviewRequest(stages.Checked).Prompt
	plan := stages.Checked.Draft.Run.DesignPlan
	for path := 0; path < 3; path++ {
		for tier := 1; tier <= 5; tier++ {
			change := plan.Paths.At(path).Milestones.At(tier)
			shown := unit.PurchaseAdaptation(plan, path, tier)
			if shown != "" && (strings.Count(prompt, shown) != 1 || strings.Contains(prompt, change) && change != shown) {
				t.Errorf("%s: the shown adaptation appears %d times, the raw milestone %v", unit.BuildCode(path, tier), strings.Count(prompt, shown), strings.Contains(prompt, change))
			}
			if shown == "" && strings.Count(prompt, change) != 1 {
				t.Errorf("%s: the planned change appears %d times", unit.BuildCode(path, tier), strings.Count(prompt, change))
			}
		}
	}
	paragraphs := strings.Split(prompt, "\n\n")
	var context map[string]any
	if err := json.Unmarshal([]byte(paragraphs[len(paragraphs)-1]), &context); err != nil {
		t.Fatal(err)
	}
	designPlan, _ := json.Marshal(context["designPlan"])
	for _, banned := range []string{"Proposed early purchases", `"upgradeIntents"`, `"milestones"`, `"crosspaths"`, `"referenceExample"`} {
		if strings.Contains(string(designPlan), banned) {
			t.Errorf("the review's designPlan still holds %s", banned)
		}
	}
	if strings.Contains(prompt, "Proposed early purchases") {
		t.Error("the review context repeats early purchases as crosspath proposals")
	}
	// The plan's decisions across purchases stay.
	for _, want := range []string{`"designPlan":{"contract":"purchase-plan-v1","concept":`, `"buyFor":"Dense lanes where one throw can pass through many enemies."`, `"omittedTechniques":`} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review context lacks %s", want)
		}
	}
	// An entry no purchase names shows an empty adaptedBy, so the review can
	// fail an adaptation it claims; code does not check it.
	checked := stages.Checked
	unused := *checked.Draft.Run.DesignPlan
	unused.Repertoire = append(append([]unit.PlanRepertoire(nil), unused.Repertoire...), unit.PlanRepertoire{
		Name: "Dart Monkey Legend", SourceIDs: []string{"source1:2"}, Limitation: "Not purchased.",
		Effects: []unit.PlanEffect{{Effect: "The dart flies far.", AdaptedAs: []string{"range"}, Reason: "Reach is range."}},
	})
	if issues := unit.PlanEffectIssues(unused); len(issues) > 0 {
		t.Errorf("code checked an entry no purchase names: %v", issues)
	}
	checked.Draft.Run.DesignPlan = &unused
	if got := unit.BlueprintReviewRequest(checked).Prompt; !strings.Contains(got, `"name":"Dart Monkey Legend","sourceIds":["source1:2"],"limitation":"Not purchased.","effects":[{"effect":"The dart flies far.","adaptedAs":["range"],"reason":"Reach is range."}],"adaptedBy":[]}`) {
		t.Error("the review does not see that no purchase adapts the entry")
	}
	// The saved draft keeps the whole plan.
	if saved := s.Stringify(s.FromGoValue(stages.Checked.Draft.Run.DesignPlan)); !strings.Contains(saved, "Proposed early purchases") || !strings.Contains(saved, `"upgradeIntents"`) {
		t.Error("the saved plan lost its crosspaths or promises")
	}
}

// A review that cites a build the Definition does not allow is corrected
// once and then rejected.
func TestReviewsCitingIllegalBuildsAreCorrectedOnce(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	valid := recordedOutput(t, "review")
	illegal := s.Clone(valid).(*s.Object)
	findings, _ := illegal.Get("findings")
	findings.([]any)[0].(*s.Object).Set("message", "The 3-3-0 crosspath loses the frenzy.")

	model := &fixture.Model{Outputs: []any{illegal, valid}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil {
		t.Fatalf("a corrected review failed: %v", err)
	}
	if len(model.Requests) != 2 || !strings.Contains(model.Requests[1].Prompt, "It cites builds that are not legal under this Definition: 3-3-0.") {
		t.Fatalf("%d calls; correction missing", len(model.Requests))
	}

	model = &fixture.Model{Outputs: []any{illegal, illegal}}
	_, err = unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	var failure *unit.ModelError
	if !errors.As(err, &failure) || !strings.Contains(failure.Message, "cited builds that are not legal (3-3-0) after one correction") || len(model.Requests) != 2 {
		t.Errorf("a review citing 3-3-0 twice: %v after %d calls", err, len(model.Requests))
	}
}

// A finding cites the resolved values it relies on and code checks them, so
// a review cannot report that 2-x-x lacks a change it has (reported on #27).
func TestReviewFindingsCiteCheckedFacts(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	citing := func(build, field, value string) *s.Object {
		review := recordedOutput(t, "review")
		findings, _ := review.Get("findings")
		findings.([]any)[0].(*s.Object).Set("facts", []any{s.NewObject().Set("build", build).Set("field", field).Set("value", value)})
		return review
	}
	model := &fixture.Model{Outputs: []any{citing("2-0-0", "attack.pierce", "5")}}
	result, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	if err != nil || len(model.Requests) != 1 {
		t.Fatalf("a correct citation: %v after %d calls", err, len(model.Requests))
	}
	last := result.Findings[len(result.Findings)-1]
	if len(last.Facts) != 1 || last.Facts[0].Value != "5" {
		t.Errorf("the result lost the cited facts: %+v", last)
	}

	model = &fixture.Model{Outputs: []any{citing("2-0-0", "attack.pierce", "3"), citing("2-0-0", "attack.pierce", "5")}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil {
		t.Fatalf("a corrected citation: %v", err)
	}
	if len(model.Requests) != 2 || !strings.Contains(model.Requests[1].Prompt, "cites 2-0-0 attack.pierce as 3, but it is 5.") {
		t.Errorf("%d calls; the correction does not name the resolved value", len(model.Requests))
	}

	model = &fixture.Model{Outputs: []any{citing("2-0-0", "attack.width", "3"), citing("2-0-0", "attack.width", "3")}}
	_, err = unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	var failure *unit.ModelError
	if !errors.As(err, &failure) || !strings.Contains(failure.Message, "a field legalBuilds does not list for that build") {
		t.Errorf("an unknown field: %v", err)
	}
}

// A correction cannot drop or change a finding whose citations hold, and the
// published summary describes the published findings. A live Luffy review on #27
// flagged a real Gatling capacity mismatch with correct facts and misread an
// interval in another finding; the correction returned no findings and "No
// concrete issue", and the Result reported a clean review.
func TestReviewCorrectionKeepsCheckedFindings(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	fact := func(build, field, value string) []any {
		return []any{s.NewObject().Set("build", build).Set("field", field).Set("value", value)}
	}
	first := recordedOutput(t, "review")
	findings, _ := first.Get("findings")
	valid := findings.([]any)[0].(*s.Object)
	valid.Set("facts", fact("2-0-0", "attack.pierce", "5"))
	misread := s.Clone(valid).(*s.Object).Set("id", "model.interval").Set("facts", fact("2-0-0", "attack.pierce", "3"))
	first.Set("findings", []any{valid, misread})

	var failure *unit.ModelError
	clean := s.Clone(first).(*s.Object).Set("summary", "No concrete issue.").Set("findings", []any{})
	changed := s.Clone(first).(*s.Object)
	findings, _ = changed.Get("findings")
	findings.([]any)[0].(*s.Object).Set("outcome", "pass")
	changed.Set("findings", findings.([]any)[:1])
	// A kept finding must return unchanged in every field: a correction
	// that rewrites its message may summarize the rewritten text, which the
	// Result would not hold (reported on #27).
	rewritten := s.Clone(first).(*s.Object).Set("summary", "No concrete issue remains.")
	findings, _ = rewritten.Get("findings")
	findings.([]any)[0].(*s.Object).Set("message", "Rewritten.")
	findings.([]any)[1].(*s.Object).Set("facts", fact("2-0-0", "attack.pierce", "5"))
	for name, correction := range map[string]*s.Object{"dropped": clean, "changed outcome": changed, "rewritten message": rewritten} {
		model := &fixture.Model{Outputs: []any{first, correction}}
		result, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
		if !errors.As(err, &failure) || !strings.Contains(failure.Message, "dropped or changed findings whose citations held (model.fan-club-allies)") || len(model.Requests) != 2 {
			t.Errorf("%s: %v after %d calls", name, err, len(model.Requests))
		}
		if result.ID != "" || result.ReviewSummary != "" || len(result.Findings) != 0 {
			t.Errorf("%s: a rejected correction published a Result: %+v", name, result)
		}
	}

	// A correction that returns the kept finding verbatim publishes its
	// summary with it; the flagged finding is replaced, and a new finding is
	// checked like the others.
	corrected := s.Clone(first).(*s.Object).Set("summary", "One unresolved scope limit.")
	findings, _ = corrected.Get("findings")
	findings.([]any)[1].(*s.Object).Set("facts", fact("2-0-0", "attack.pierce", "5"))
	extra := s.Clone(valid).(*s.Object).Set("id", "model.extra")
	corrected.Set("findings", append(findings.([]any), extra))
	model := &fixture.Model{Outputs: []any{first, corrected}}
	result, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	prompt := model.Requests[1].Prompt
	for _, want := range []string{"Your previous review was: {", "model.interval cites 2-0-0 attack.pierce as 3, but it is 5.", "a summary that describes exactly the findings you return", "Return model.fan-club-allies verbatim, every field exactly as in your previous review", "Return model.interval corrected under the same ID"} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the correction lacks %q", want)
		}
	}
	var ids []string
	for _, f := range result.Findings {
		if strings.HasPrefix(f.ID, "model.") {
			ids = append(ids, f.ID)
			if f.ID == "model.fan-club-allies" && s.Stringify(s.FromGoValue(f)) != s.Stringify(valid) {
				t.Errorf("the kept finding changed: %+v", f)
			}
			if f.ID == "model.interval" && f.Facts[0].Value != "5" {
				t.Errorf("the flagged finding was not replaced: %+v", f.Facts)
			}
		}
	}
	if strings.Join(ids, ",") != "model.fan-club-allies,model.interval,model.extra" || result.ReviewSummary != "One unresolved scope limit." {
		t.Errorf("model findings %v, summary %q", ids, result.ReviewSummary)
	}
}

// The review of a revision judges the current plan and unit. It gets the
// earlier unit and the requested change, but not the earlier review's
// findings: a review of an edited Luffy plan on #27 repeated an earlier
// verdict that Gear 2's speed was only temporary, which the revised plan
// no longer said.
func TestRevisionReviewsJudgeTheCurrentPlan(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	rule := "a verdict about the plan must hold for the current designPlan text"
	first := unit.BlueprintReviewRequest(stages.Checked).Prompt
	if strings.Contains(first, `"revision":{`) || strings.Contains(first, rule) || strings.Contains(first, "previousFindings") {
		t.Error("a first generation's review carries revision context")
	}
	checked := stages.Checked
	request := checked.Draft.Prepared.Request
	feedback := "Make Gear 2 speed permanent."
	request.Feedback = &feedback
	request.Previous = &unit.Previous{ResultID: "earlier", Draft: stages.Result.Candidate, Findings: []unit.Finding{
		{ID: "model.gear2", Method: "model", Outcome: "fail", Subject: "x-3-x", Message: "The plan limits Gear 2 speed to a temporary boost."},
	}}
	checked.Draft.Prepared.Request = request
	prompt := unit.BlueprintReviewRequest(checked).Prompt
	for _, want := range []string{`"revision":{"earlierUnit":{`, `"feedback":"Make Gear 2 speed permanent."`, rule} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the revision review lacks %q", want)
		}
	}
	if strings.Contains(prompt, "limits Gear 2 speed to a temporary boost") || strings.Contains(prompt, "previousFindings") {
		t.Error("the revision review carries the earlier review's findings")
	}
}

// Kyle's diagnostic edit of a Luffy Result on #27: the earlier version said
// Gear 2's speed appeared only in the Active, the revised plan made it
// permanent from x-3-x, and the review still judged the earlier claim. The
// review of that edit presents only the revised plan as current; the
// earlier claim is left only inside revision.earlierUnit, and the earlier
// review's finding is gone.
func TestRevisionReviewTreatsOnlyTheNewPlanAsCurrent(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	earlier := "Gear 2 speed appears only in the Active."
	current := "Gear 2 speed is permanent from x-3-x, and the Active amplifies it."
	previous := stages.Result.Candidate
	blueprint := *previous.Blueprint
	blueprint.Paths.Path2.Rationale = earlier
	previous.Blueprint = &blueprint
	checked := stages.Checked
	plan := *checked.Draft.Run.DesignPlan
	plan.Repertoire = append([]unit.PlanRepertoire(nil), plan.Repertoire...)
	plan.Repertoire[0].Limitation = current
	checked.Draft.Run.DesignPlan = &plan
	request := checked.Draft.Prepared.Request
	feedback := "Gear 2 is a form that raises strength and speed; do not keep its speed only in the Active."
	request.Feedback = &feedback
	request.Previous = &unit.Previous{ResultID: "earlier", Draft: previous, Findings: []unit.Finding{
		{ID: "model.gear2", Method: "model", Outcome: "fail", Subject: "x-4-x", Message: "The plan limits Gear 2 speed to a temporary boost: " + earlier},
	}}
	checked.Draft.Prepared.Request = request
	prompt := unit.BlueprintReviewRequest(checked).Prompt
	paragraphs := strings.Split(prompt, "\n\n")
	var context map[string]any
	if err := json.Unmarshal([]byte(paragraphs[len(paragraphs)-1]), &context); err != nil {
		t.Fatal(err)
	}
	text := func(value any) string {
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		return string(encoded)
	}
	if got := text(context["designPlan"]); !strings.Contains(got, current) || strings.Contains(got, earlier) {
		t.Errorf("the current designPlan is not the revised plan: %s", got)
	}
	for _, key := range []string{"previous", "previousFindings", "feedback"} {
		if _, ok := context[key]; ok {
			t.Errorf("the review context holds %s as a current fact", key)
		}
	}
	revision, _ := context["revision"].(map[string]any)
	if !strings.Contains(text(revision["earlierUnit"]), earlier) || revision["feedback"] != feedback {
		t.Errorf("revision does not hold the earlier unit and the requested change: %s", text(revision))
	}
	if got := strings.Count(prompt, earlier); got != 1 {
		t.Errorf("the earlier claim appears %d times, want once inside revision.earlierUnit", got)
	}
	if strings.Contains(prompt, "limits Gear 2 speed to a temporary boost") {
		t.Error("the review carries the earlier review's finding")
	}
	if !strings.Contains(prompt, "a verdict about the plan must hold for the current designPlan text, and nothing an earlier version said is evidence about this one") {
		t.Error("the review lacks the rule to judge the current plan")
	}
}

// A code that is neither a legal build nor a purchase, such as 1-x-5 for
// 1-5-0 (a live Luffy review on #27), is corrected like an illegal build.
// Purchases such as x-4-x and legal builds such as 0-5-1 are not.
func TestReviewsCitingMalformedCodesAreCorrected(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	citing := func(message string) *s.Object {
		review := recordedOutput(t, "review")
		findings, _ := review.Get("findings")
		findings.([]any)[0].(*s.Object).Set("message", message)
		return review
	}
	valid := citing("x-4-x and 3-x-x are fine, and 0-5-1 is a legal build.")
	model := &fixture.Model{Outputs: []any{valid}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil || len(model.Requests) != 1 {
		t.Fatalf("purchase codes and legal builds: %v after %d calls", err, len(model.Requests))
	}
	// A notation fix changes the codes and nothing else (SOL-34-04).
	model = &fixture.Model{Outputs: []any{citing("The side purchase 1-x-5 adds more than x-5-x."), citing("The side purchase 1-5-0 adds more than x-5-x.")}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil {
		t.Fatalf("a corrected code: %v", err)
	}
	if len(model.Requests) != 2 || !strings.Contains(model.Requests[1].Prompt, "1-x-5 is neither a purchase nor a build.") || !strings.Contains(model.Requests[1].Prompt, "write a purchase as 3-x-x, x-4-x or x-x-5 and a build as 1-5-0") {
		t.Errorf("%d calls; the correction does not name 1-x-5", len(model.Requests))
	}
	// The subject and action are read too: a 7a review on #27 wrote "3-x-3"
	// and "4-x-4", which are neither purchases nor builds.
	both := recordedOutput(t, "review")
	findings, _ := both.Get("findings")
	findings.([]any)[0].(*s.Object).Set("subject", "3-x-3").Set("action", "Compare 4-x-4 as well.")
	fixed := recordedOutput(t, "review")
	findings, _ = fixed.Get("findings")
	findings.([]any)[0].(*s.Object).Set("subject", "3-x-x").Set("action", "Compare x-4-x as well.")
	model = &fixture.Model{Outputs: []any{both, fixed}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil {
		t.Fatalf("a corrected subject and action: %v", err)
	}
	if len(model.Requests) != 2 || !strings.Contains(model.Requests[1].Prompt, "3-x-3 is neither a purchase nor a build; 4-x-4 is neither a purchase nor a build.") {
		t.Errorf("%d calls; the correction does not name 3-x-3 and 4-x-4", len(model.Requests))
	}
}

// A finding names the purchases its subject names. A live Luffy review on
// #27 had subject x-4-x and x-5-x but wrote 4-x-x and 5-x-x in its
// message, legal codes on the wrong path; it is corrected once. A message
// that names the subject's purchase may compare it with another path.
func TestReviewsNameTheSubjectsPurchase(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	finding := func(subject, message string) *s.Object {
		review := recordedOutput(t, "review")
		findings, _ := review.Get("findings")
		findings.([]any)[0].(*s.Object).Set("subject", subject).Set("message", message)
		return review
	}
	comparing := finding("x-5-x", "x-5-x adds less than the 140 Gold 1-x-x side purchase; 5-x-x is not comparable.")
	model := &fixture.Model{Outputs: []any{comparing}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil || len(model.Requests) != 1 {
		t.Fatalf("a comparison that names its subject: %v after %d calls", err, len(model.Requests))
	}
	swapped := finding("x-4-x and x-5-x", "4-x-x costs 7,200 Gold and 5-x-x adds little over it.")
	model = &fixture.Model{Outputs: []any{swapped, finding("x-4-x and x-5-x", "x-4-x costs 7,200 Gold and x-5-x adds little over it.")}}
	if _, err := unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options()); err != nil {
		t.Fatalf("a corrected finding: %v", err)
	}
	if len(model.Requests) != 2 || !strings.Contains(model.Requests[1].Prompt, "model.fan-club-allies's subject names x-4-x, but its text names 4-x-x instead. model.fan-club-allies's subject names x-5-x, but its text names 5-x-x instead.") {
		t.Errorf("%d calls; the correction does not name the swapped purchases", len(model.Requests))
	}
	model = &fixture.Model{Outputs: []any{swapped, swapped}}
	_, err = unit.ReviewDraft(context.Background(), stages.Checked, model, fixture.Options())
	var failure *unit.ModelError
	if !errors.As(err, &failure) || !strings.Contains(failure.Message, "named another path's purchase than its subject after one correction") {
		t.Errorf("a finding swapped twice: %v", err)
	}
}
