package unit_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Every whole-technique omission, major included, reaches the review with
// its importance and reason, and one named for a source technique carries
// that technique's salience and passages. The review judges each against
// the vocabulary and purchase-level proposed mechanics, fails a circular
// reason and challenges rankings against the cited passages (SOL-61-05).
func TestOmissionsReachTheReview(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	checked := stages.Checked
	techniques := append([]unit.SourceTechnique(nil), fixtureSourceTechniques...)
	techniques[4].Salience = unit.SalienceStrong // Juggernaut
	checked.Draft.Prepared.Request.SourceTechniques = &techniques
	plan := *checked.Draft.Run.DesignPlan
	plan.OmittedTechniques = append(append([]unit.PlanOmission(nil), plan.OmittedTechniques...),
		unit.PlanOmission{Name: "Juggernaut", Importance: "major", Reason: "This plan does not implement it."})
	checked.Draft.Run.DesignPlan = &plan
	prompt := unit.BlueprintReviewRequest(checked).Prompt
	for _, want := range []string{
		`{"name":"Juggernaut","importance":"major","reason":"This plan does not implement it.","sourceTechnique":{"name":"Juggernaut","salience":"strong","passageIds":["source1:7","source1:8"]}}`,
		`{"name":"Critical shots","importance":"minor","reason":"Every tenth or fifth shot dealing extra damage needs a shot counter operator."}`,
		"judge every whole-technique omission in designPlan.omittedTechniques, major and minor alike, against definition.vocabulary and the proposed mechanics a purchase could carry",
		unit.OmissionRule,
		`a circular reason never holds, as "this plan does not implement it"`,
		"Challenge every ranking against the passages its source technique cites",
		"an implausible downgrade included",
		unit.SalienceRule,
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	planRequest, err := unit.DesignPlanRequest(checked.Draft.Prepared)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(planRequest.Prompt, unit.OmissionRule) {
		t.Error("the plan prompt does not state the omission rule")
	}
}

// The review compares each third purchase with its own path's first two
// purchases and the other paths' purchases, keeps the Dart reference and
// judges a proposed-only capstone as a design gap (#66).
func TestReviewComparesThirdPurchases(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prompt := unit.BlueprintReviewRequest(stages.Checked).Prompt
	for _, want := range []string{
		"comparing each third purchase explicitly with its own path's first and second purchases and with every purchase of the other two paths",
		"Fail a third purchase that only adds numbers its path already had",
		"a punch that homes around obstructions where every delivery needs a clear path",
		"Dart Monkey's x-3-x Triple Shot only adds projectiles and no other Dart Monkey path adds any",
		"Code reports a fifth purchase whose only new capability is a proposed mechanic as an unresolved design gap",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
}

// The plan, mechanics and review prompts state what an Active Ability can
// scope, and the review reads each purchase's typed changes by scope, so a
// text that limits a base change to the Active can fail (OPUS-NET-61-7).
func TestActiveAbilityScope(t *testing.T) {
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
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
	for name, prompt := range map[string]string{"plan": model.Requests[0].Prompt, "mechanics": model.Requests[1].Prompt, "review": review} {
		if !strings.Contains(prompt, unit.ActiveScopeRule) {
			t.Errorf("the %s prompt does not state the Active Ability scope", name)
		}
	}
	for _, want := range []string{
		// x-1-x attacks faster, a base change; x-4-x unlocks the Active and
		// x-5-x changes it, boost changes.
		`"plannedChange":"x-1-x shortens the attack interval.","changeScope":{"base":["intervalSeconds multiply 0.85"],"boost":[]}}`,
		`"changeScope":{"base":["intervalSeconds multiply 0.5"],"boost":["unlocks the Active Ability Fan Club Frenzy"]}}`,
		`"changeScope":{"base":[],"boost":["boost damageMultiplier`,
		"Fail, on that build code, a name, adaptation or plannedChange whose scope contradicts changeScope",
	} {
		if !strings.Contains(review, want) {
			t.Errorf("the review context lacks %q", want)
		}
	}
}

// escanorV38Checked is the v38 Escanor Result 427dab9f as a checked draft:
// its retained plan and blueprint, prepared from the same Sources under the
// Default, without required concepts, as its Request had none
// (OPUS-NET-61-22).
func escanorV38Checked(t *testing.T) unit.Checked {
	t.Helper()
	return savedChecked(t, "escanor", "v38", unit.Run{ID: "381d1f48-5c18-4d03-b22f-d837ffe452bf", ModelID: "openrouter:openai/gpt-6-luna", StartedAt: "2026-09-27T14:33:18.365Z", CompletedAt: "2026-09-27T14:34:38.836Z"})
}

// The v38 Escanor plan omitted Sunshine whole because the Definition has no
// time-of-day cycle, ranked a basic Super Slash core, and its review passed
// the omission on the literal schedule, though the passages Sunshine cites
// describe a rise to a peak that x-4-x's Active Ability The One already
// carries in part (SOL-61-17). The review gets the omission's reason and
// passages, the core ranking and x-4-x's typed changes, and is asked to
// judge the central effect, say what is covered and what is absent, and
// fail an unjustified whole omission. A model's verdict cannot be tested
// here; this pins what the review is given and asked.
func TestOmissionReviewJudgesTheCentralEffectNotALiteralDetail(t *testing.T) {
	checked := escanorV38Checked(t)
	prompt, context := reviewContext(t, checked)
	sunshine := at(context, "designPlan", "omittedTechniques").([]any)[0]
	if at(sunshine, "name") != "Sunshine" || at(sunshine, "importance") != "major" ||
		!strings.Contains(at(sunshine, "reason").(string), "The Definition cannot express a time-of-day cycle") ||
		s.Stringify(at(sunshine, "sourceTechnique")) != `{"name":"Sunshine","salience":"strong","passageIds":["source2:12","source2:14","source2:16","source2:18","source2:21","source2:27","source2:32","source2:33","source2:36","source2:38","source2:40","source4:19"]}` {
		t.Errorf("Sunshine omission %s", s.Stringify(sunshine))
	}
	for _, passage := range at(context, "sourcePassages").([]any) {
		if at(passage, "id") == "source2:27" && !strings.Contains(at(passage, "text").(string), "At high noon, Escanor reaches his absolute peak and becomes The One") {
			t.Errorf("source2:27 reads %q", at(passage, "text"))
		}
	}
	var ranks []string
	for _, entry := range at(context, "designPlan", "repertoire").([]any) {
		ranks = append(ranks, at(entry, "name").(string)+"="+at(entry, "importance").(string))
	}
	if got := strings.Join(ranks, ","); got != "Super Slash=core,Cruel Sun=major,The One=core,Crazy Prominence=major" {
		t.Errorf("ranking %s", got)
	}
	// What the Unit already covers of the rise to a peak: x-4-x's Active
	// Ability.
	peak := at(context, "unit", "paths").([]any)[1]
	if tier := at(peak, "tiers").([]any)[3]; s.Stringify(at(tier, "changeScope")) != `{"base":["damage add 2","intervalSeconds multiply 0.8"],"boost":["unlocks the Active Ability The One"]}` {
		t.Errorf("x-4-x %s", s.Stringify(tier))
	}
	for _, want := range []string{
		unit.OmissionRule + " " + unit.OmittedNameRule,
		"a literal detail the Definition cannot express, such as a time-of-day clock, justifies omitting only that aspect when a typed change or a proposed mechanic on a purchase could carry the central effect",
		"carries the technique's central effect, the source-backed effect that identifies it against the character's other forms and techniques",
		"its reason names the technique's central effect from the passages its sourceTechnique cites",
		"not whether a literal detail or another aspect of it is unsupported",
		"says exactly what the Unit already covers of that effect, naming each purchase by build code with its change as changeScope lists it, and what is still absent",
		"fail a whole omission whose central effect a typed change or a proposed mechanic could carry, with the action to adapt that effect and omit only the unsupported aspect",
		"the core ranking beside it, against the passages its sourceTechnique cites and the character's identity, and fails an omission those passages show to be more central than an entry the plan ranks core",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	plan, err := unit.DesignPlanRequest(checked.Draft.Prepared)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(plan.Prompt, unit.OmissionRule+" "+unit.OmittedNameRule) {
		t.Error("the plan prompt does not state the omission and omitted name rules")
	}
}

// The v38 Escanor plan omitted Rhitta whole while its base attack and the
// first purchase of every path say "Rhitta slash", and named its omitted
// Sunshine as its signature and in 1-x-x (SOL-61-17). Code lists each text
// that names an omitted technique, by its name or an alias of its source
// technique, beside the omission's verdict subject; it gates nothing.
func TestOmittedTechniqueNamedInPlanTextsIsCued(t *testing.T) {
	checked := escanorV38Checked(t)
	prompt, context := reviewContext(t, checked)
	namedBy := func(context *s.Object) map[string]string {
		out := map[string]string{}
		for _, omission := range at(context, "requiredVerdicts", "omissions").([]any) {
			if omission.(*s.Object).Has("namedBy") {
				out[at(omission, "technique").(string)] = s.Stringify(at(omission, "namedBy"))
			}
		}
		return out
	}
	got := namedBy(context)
	for technique, want := range map[string]string{
		"Rhitta":   `["designPlan.base.behavior","1-x-x adaptation","x-1-x adaptation","x-x-1 adaptation"]`,
		"Sunshine": `["designPlan.signature.name","1-x-x adaptation"]`,
		"Daytime":  `["designPlan.base.behavior"]`,
	} {
		if got[technique] != want {
			t.Errorf("%s namedBy %s, want %s", technique, got[technique], want)
		}
	}
	// Final Prominence shares a word with Crazy Prominence, never its name.
	if len(got) != 3 {
		t.Errorf("namedBy %v", got)
	}
	for index, want := range []string{"1-x-x makes each selected-target Rhitta slash hit harder", "x-1-x makes Escanor's ordinary Rhitta slash arrive more frequently", "x-x-1 increases the number of enemies a single Rhitta slash can hit"} {
		if milestone := checked.Draft.Run.DesignPlan.Paths.At(index).Milestones.At(1); !strings.HasPrefix(milestone, want) {
			t.Errorf("milestone %q", milestone)
		}
	}
	for _, want := range []string{
		"requiredVerdicts.omissions gives an omission namedBy when code found its name, or an alias of its source technique, as whole words in designPlan.signature, designPlan.base or a purchase's name, adaptation or plannedChange",
		"fail an omission that such a text claims as part of what the Unit does, as the omitted name rule says",
		unit.OmittedNameRule,
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}

	// An omission by an alias of its source technique is found by the
	// technique's name.
	plan := *checked.Draft.Run.DesignPlan
	plan.OmittedTechniques = append([]unit.PlanOmission(nil), plan.OmittedTechniques...)
	plan.OmittedTechniques[3].Name = "The Divine Axe Rhitta"
	checked.Draft.Run.DesignPlan = &plan
	_, context = reviewContext(t, checked)
	if got := namedBy(context)["The Divine Axe Rhitta"]; got != `["designPlan.base.behavior","1-x-x adaptation","x-1-x adaptation","x-x-1 adaptation"]` {
		t.Errorf("The Divine Axe Rhitta namedBy %s", got)
	}
}

// Fail verdicts on the Sunshine and Rhitta omissions are recorded as open
// omission-verdict Findings with their reasons and actions. The verdicts are
// scripted as the clarified rules ask for them on the v38 Escanor Unit
// (SOL-61-17); they show what the Result keeps, not what a model returns.
func TestUnjustifiedOmissionFailVerdictsAreRecorded(t *testing.T) {
	checked := escanorV38Checked(t)
	sunshine := "character-technique:nanatsu-no-taizai.fandom.com:Sunshine"
	rhitta := "character-technique:nanatsu-no-taizai.fandom.com:Rhitta"
	verdict := func(technique, reason, action, evidence string) *s.Object {
		return s.NewObject().Set("technique", technique).Set("outcome", "fail").Set("reason", reason).Set("action", action).Set("evidence", []any{evidence})
	}
	output := scriptedReview(t, checked, map[string]*s.Object{
		"Sunshine": verdict("Sunshine", "source2:27 identifies Sunshine by power that rises to a peak at noon. x-4-x covers the peak with its Active Ability The One (unlocks the Active Ability The One); the rise before the peak is still absent. No day and night clock justifies omitting only that aspect.", "Adapt Sunshine's rise on a purchase, as a proposed ramp, keep The One as its peak and omit only the time-of-day schedule.", sunshine),
		"Rhitta":   verdict("Rhitta", "designPlan.base.behavior, 1-x-x, x-1-x and x-x-1 adaptation all name a Rhitta slash while Rhitta is omitted whole; heat storage and command release are only an aspect.", "Select Rhitta as the base attack's weapon and omit only its heat storage, or stop naming Rhitta in those texts.", rhitta),
	}, nil)
	result, err := unit.ReviewDraft(context.Background(), checked, &fixture.Model{Outputs: []any{output}}, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]unit.Finding{}
	for _, f := range verdictFindings(result) {
		if f.Rule == unit.OmissionVerdictRule && f.Outcome != "pass" {
			got[f.Subject] = f
		}
	}
	if len(got) != 2 {
		t.Fatalf("open omission verdicts %+v", got)
	}
	for subject, want := range map[string]struct{ id, message, action string }{
		"omittedTechniques, Sunshine (major)": {"verdict.omission.1", "the rise before the peak is still absent", "omit only the time-of-day schedule"},
		"omittedTechniques, Rhitta (major)":   {"verdict.omission.4", "name a Rhitta slash while Rhitta is omitted whole", "stop naming Rhitta"},
	} {
		f := got[subject]
		if f.ID != want.id || f.Outcome != "fail" || f.Severity != "error" || f.Method != "model" ||
			!strings.Contains(f.Message, want.message) || f.Action == nil || !strings.Contains(*f.Action, want.action) {
			t.Errorf("%s: %+v", subject, f)
		}
	}
}
