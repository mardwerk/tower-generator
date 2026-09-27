package unit_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/render"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// The fixture's Fan Club transforms nearby Dart Monkeys and its Crossbow
// scores a critical hit every tenth shot; the Definition expresses neither.
var (
	allyFrenzy = mechanics.ProposedMechanic{
		Name:      "Fan Club transformation",
		Effect:    "Up to 10 nearby Dart Monkeys join the frenzy and attack as fast as this Unit while it lasts.",
		SourceIDs: []string{"source1:12"},
	}
	criticalBolt = mechanics.ProposedMechanic{
		Name:      "Critical bolt",
		Effect:    "Every tenth bolt deals five times its damage",
		SourceIDs: []string{"source1:17"},
	}
)

func proposedValue(p mechanics.ProposedMechanic) *s.Object {
	return s.FromGoValue(p).(*s.Object)
}

// proposalOutputs are the fixture's plan and mechanics outputs with the
// Fan Club transformation named by the x-4-x milestone and the critical
// bolt added by the mechanics stage to x-x-4.
func proposalOutputs(t *testing.T) (plan, design *s.Object) {
	t.Helper()
	plan, design = recordedOutput(t, "plan"), recordedOutput(t, "mechanics")
	at(plan, "paths", "path2", "milestones", "tier4").(*s.Object).Set("proposedMechanics", []any{proposedValue(allyFrenzy)})
	at(design, "paths", "path3", "tiers", "tier4").(*s.Object).Set("proposedMechanics", []any{proposedValue(criticalBolt)})
	return plan, design
}

func draftWith(t *testing.T, outputs ...any) (unit.Draft, *fixture.Model) {
	t.Helper()
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	model := &fixture.Model{Outputs: outputs}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v\n%s", err, retryReason(model.Requests[len(model.Requests)-1].Prompt))
	}
	return draft, model
}

// A proposed mechanic stays on its purchase through plan, design, check,
// review and render; it grants nothing and changes no resolved stat.
func TestProposedMechanicsStayOnTheirPurchase(t *testing.T) {
	plan, design := proposalOutputs(t)
	review := recordedOutput(t, "review")
	draft, model := draftWith(t, plan, design, review)
	baseline, _ := draftWith(t, recordedOutput(t, "plan"), recordedOutput(t, "mechanics"))

	intent := draft.Run.DesignPlan.UpgradeIntents.Path2.Tier4
	if len(intent.ProposedMechanics) != 1 || intent.ProposedMechanics[0].Name != allyFrenzy.Name {
		t.Fatalf("the plan lost its proposed mechanic: %+v", intent.ProposedMechanics)
	}
	if !strings.Contains(model.Requests[1].Prompt, `"proposedMechanics":[{"name":"Fan Club transformation"`) {
		t.Error("the mechanics stage does not see the milestone's proposed mechanic")
	}
	blueprint := draft.Candidate.Blueprint
	// Code binds the plan's proposal although the mechanics output left it out.
	if got := blueprint.Paths.Path2.Tiers.Tier4.ProposedMechanics; len(got) != 1 || got[0].Name != allyFrenzy.Name || got[0].Effect != allyFrenzy.Effect {
		t.Errorf("x-4-x proposed mechanics %+v", got)
	}
	if got := blueprint.Paths.Path3.Tiers.Tier4.ProposedMechanics; len(got) != 1 || got[0].Name != criticalBolt.Name {
		t.Errorf("x-x-4 proposed mechanics %+v", got)
	}

	// The draft decodes and validates as saved.
	parsed, err := unit.ParseDraft(s.FromGoValue(draft))
	if err != nil {
		t.Fatal(err)
	}
	definition := *parsed.Prepared.Request.MechanicsDefinition
	if issues := unit.ValidateBlueprintRequest(*parsed.Candidate.Blueprint, parsed.Prepared.Request); len(issues) > 0 {
		t.Errorf("validation issues %+v", issues)
	}

	// Build resolution ignores proposed mechanics.
	for _, selection := range mechanics.AllLegalBuilds(definition) {
		with, err := mechanics.ResolveBuild(blueprint, selection, definition)
		if err != nil {
			t.Fatal(err)
		}
		without, err := mechanics.ResolveBuild(baseline.Candidate.Blueprint, selection, definition)
		if err != nil {
			t.Fatal(err)
		}
		if s.Stringify(s.FromGoValue(with)) != s.Stringify(s.FromGoValue(without)) {
			t.Errorf("%s resolves differently with proposed mechanics", unit.SelectionCode(selection))
		}
	}
	if s.Stringify(s.FromGoValue(draft.Candidate.Paths)) != s.Stringify(s.FromGoValue(baseline.Candidate.Paths)) ||
		s.Canonical(draft.Run.DesignEvaluation) != s.Canonical(baseline.Run.DesignEvaluation) {
		t.Error("proposed mechanics changed the compiled purchases or the purchase evidence")
	}

	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	var proposed []unit.Finding
	total := 0
	for _, finding := range checked.Findings {
		if finding.Outcome == "fail" {
			t.Errorf("%s %s: %s", finding.Rule, finding.Subject, finding.Message)
		}
		// The fixture's own proposals, Rebound at 3-x-x and the Plasma
		// transformation at x-5-x, are reported too.
		if finding.Rule == unit.ProposedMechanicRule {
			total++
			if strings.HasSuffix(finding.Subject, "tier4.proposedMechanics.0") {
				proposed = append(proposed, finding)
			}
		}
	}
	if len(proposed) != 2 || total != 4 {
		t.Fatalf("proposed-mechanic findings %+v of %d", proposed, total)
	}
	first := proposed[0]
	if first.Outcome != "unresolved" || first.Method != "deterministic" || first.Subject != "paths.path2.tiers.tier4.proposedMechanics.0" ||
		!strings.HasPrefix(first.Message, "x-4-x Super Monkey Fan Club proposes Fan Club transformation: Up to 10 nearby") ||
		!strings.Contains(first.Message, "The Definition must be expanded before this purchase can grant it") ||
		strings.Join(first.Evidence, ",") != "dart-monkey-atlas-56-3" {
		t.Errorf("x-4-x finding %+v", first)
	}
	if !strings.Contains(proposed[1].Message, "x-x-4 ") || !strings.Contains(proposed[1].Message, "five times its damage. The Definition") {
		t.Errorf("x-x-4 finding %+v", proposed[1])
	}

	result, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	reviewPrompt := model.Requests[2].Prompt
	if !strings.Contains(reviewPrompt, `"proposedMechanics":[{"name":"Critical bolt","effect":"Every tenth bolt deals five times its damage","sourceIds":["source1:17"]}]`) {
		t.Error("the review context lacks x-x-4's proposed mechanic")
	}
	if !strings.Contains(reviewPrompt, "fail one that makes no sense for its purchase's technique and the passages its sourceIds cite, or that is lore rather than a mechanic") {
		t.Error("the review is not asked to judge proposed mechanics")
	}

	markdown, err := render.Markdown(s.FromGoValue(result), false)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"Proposed (not yet supported): Fan Club transformation: Up to 10 nearby Dart Monkeys join the frenzy and attack as fast as this Unit while it lasts.",
		"Proposed (not yet supported): Critical bolt: Every tenth bolt deals five times its damage.",
	} {
		line := ""
		for _, l := range strings.Split(markdown, "\n") {
			if strings.Contains(l, want) {
				line = l
			}
		}
		if line == "" {
			t.Errorf("the sheet lacks %q", want)
		} else if !strings.HasPrefix(line, "**x-4-x ") && !strings.HasPrefix(line, "**x-x-4 ") {
			t.Errorf("%q is not on its purchase: %s", want, line)
		}
	}
	if strings.Count(markdown, "Proposed (not yet supported)") != 4 {
		t.Error("a proposed mechanic appears outside its purchase, such as in a crosspath row")
	}

	purchases := render.Purchases(result.Candidate, result.Prepared.Request.MechanicsDefinition, result.Run.Draft.DesignPlan)
	fanClub := purchases[1].Purchases[3]
	if fanClub.Code != "x-4-x" || len(fanClub.ProposedMechanics) != 1 || !strings.HasSuffix(fanClub.Text, "Proposed (not yet supported): Fan Club transformation: "+allyFrenzy.Effect) {
		t.Errorf("x-4-x view %+v", fanClub)
	}
	if other := purchases[0].Purchases[3]; len(other.ProposedMechanics) != 0 || strings.Contains(other.Text, "Proposed") {
		t.Errorf("4-x-x shows a proposed mechanic: %+v", other)
	}

	// A saved draft whose purchase drops a planned proposal fails check.
	edited := *blueprint
	edited.Paths.Path2.Tiers.Tier4.ProposedMechanics = nil
	draft.Candidate.Blueprint = &edited
	if draft.Candidate, err = unit.CompileBlueprint(edited, draft.Prepared.Request); err != nil {
		t.Fatal(err)
	}
	rechecked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	dropped := false
	for _, finding := range rechecked.Findings {
		if finding.Outcome == "fail" && finding.Rule == "planned-upgrade-intent" && finding.Subject == "paths.path2.tiers.tier4.proposedMechanics" &&
			strings.HasPrefix(finding.Message, "The retained plan proposes Fan Club transformation at x-4-x, but the purchase does not carry it.") {
			dropped = true
		}
	}
	if !dropped {
		t.Error("check does not report the dropped planned proposal")
	}
}

// A saved unit without proposed mechanics loads, checks and renders as
// before: no proposed-mechanic finding and no proposed line. The fixture
// proposes a rebound and the Plasma transformation, so this unit is drafted
// without them, under a policy that does not need them.
func TestUnitWithoutProposedMechanicsIsUnchanged(t *testing.T) {
	prepared := preparedUnder(t, func(p *mechanics.DesignPolicy) { p.RequireTier3PathIdentity, p.RequireTier5BehaviorChange = nil, nil })
	model := &fixture.Model{Outputs: []any{planWithoutProposals(t), recordedOutput(t, "mechanics"), recordedOutput(t, "review")}}
	draft, err := unit.DraftUnit(context.Background(), prepared, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	for _, finding := range checked.Findings {
		if finding.Rule == unit.ProposedMechanicRule {
			t.Errorf("unexpected finding %+v", finding)
		}
	}
	result, err := unit.ReviewDraft(context.Background(), checked, model, fixture.Options())
	if err != nil {
		t.Fatal(err)
	}
	value := s.FromGoValue(result)
	if strings.Contains(s.Stringify(value), "proposedMechanics") {
		t.Error("a Result without proposed mechanics writes the field")
	}
	markdown, err := render.Markdown(value, false)
	if err != nil || strings.Contains(markdown, "Proposed (not yet supported)") {
		t.Errorf("render: %v", err)
	}
}

// A proposed mechanic must cite supplied passages, and a purchase names
// each one once.
func TestProposedMechanicValidation(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	request := stages.Draft.Prepared.Request
	blueprint := *stages.Draft.Candidate.Blueprint
	uncited := criticalBolt
	uncited.SourceIDs = []string{"source9:99"}
	blueprint.Paths.Path3.Tiers.Tier4.ProposedMechanics = []mechanics.ProposedMechanic{criticalBolt, uncited}
	var messages []string
	for _, issue := range unit.ValidateBlueprintRequest(blueprint, request) {
		messages = append(messages, issue.Path+": "+issue.Message)
	}
	got := strings.Join(messages, "\n")
	for _, want := range []string{
		"paths.path3.tiers.tier4.proposedMechanics.1: Name each proposed mechanic of a purchase once.",
		"paths.path3.tiers.tier4.proposedMechanics.1.sourceIds: Cite a supplied source passage ID; source9:99 is not one.",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in\n%s", want, got)
		}
	}
	// A proposed mechanic is no typed change: a tier with only a proposal
	// still fails the change minimum.
	schema := mechanics.BlueprintSchemaFor(*request.MechanicsDefinition)
	value := s.FromGoValue(blueprint).(*s.Object)
	tier := at(value, "paths", "path3", "tiers", "tier4").(*s.Object)
	tier.Set("changes", []any{})
	if _, issues := s.Parse(schema, value); len(issues) == 0 {
		t.Error("a tier with a proposed mechanic and no typed change parses")
	}
}

// The prompts keep an unsupported mechanic on the purchase that needs it
// instead of telling the model to drop or park it.
func TestPromptsKeepProposedMechanicsOnPurchases(t *testing.T) {
	plan, design := proposalOutputs(t)
	_, model := draftWith(t, plan, design)
	planPrompt, designPrompt := model.Requests[0].Prompt, model.Requests[1].Prompt
	for _, banned := range []string{
		"omit unsupported techniques",
		"list that concrete gap in unsupportedMechanics",
		"list those as unsupported mechanics",
		"List in unsupportedMechanics, with a concrete reason, every planned or sourced behavior",
		"describe anything else in unsupportedMechanics",
		"Unsupported wishes remain explicit omissions",
	} {
		if strings.Contains(planPrompt, banned) || strings.Contains(designPrompt, banned) {
			t.Errorf("a prompt still says %q", banned)
		}
	}
	for name, test := range map[string]struct {
		prompt string
		want   []string
	}{
		"plan": {planPrompt, []string{
			"name it in that milestone's proposedMechanics as {name, effect, sourceIds}",
			"a bodily strain or a blood-flow mechanism is not a mechanic",
		}},
		"mechanics": {designPrompt, []string{
			"keep it on the purchase that needs it as a proposed mechanic",
			"that tier's proposedMechanics lists {name, effect, sourceIds}",
			"List in unsupportedMechanics, with a concrete reason, only such behavior that belongs to no purchase.",
		}},
	} {
		for _, want := range test.want {
			if !strings.Contains(test.prompt, want) {
				t.Errorf("%s prompt lacks %q", name, want)
			}
		}
	}
	// Both model-facing schemas accept proposed mechanics.
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	planSchema := s.Stringify(unit.ProviderJSONSchema(unit.PurchasePlanOutputSchema(&prepared.Request)))
	designSchema, err := unit.ModelOutputJSONSchema(&prepared.Request)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(planSchema, `"proposedMechanics"`) || !strings.Contains(s.Stringify(designSchema), `"proposedMechanics"`) {
		t.Error("a model-facing schema lacks proposedMechanics")
	}
}
