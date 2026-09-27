package unit_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// fixtureSourceTechniques are source techniques as prepare would list them
// for the Dart Monkey brief. Spike-o-pult and Super Monkey Fan Club are
// listed by the passages their repertoire entries cite, Triple Shot and
// Crossbow by name, and Juggernaut by nothing in the fixture plan.
var fixtureSourceTechniques = []unit.SourceTechnique{
	{Name: "Spike-o-pult", PassageIDs: []string{"source1:6"}},
	{Name: "Triple Shot", PassageIDs: []string{"source1:11"}, Signature: true},
	{Name: "Super Monkey Fan Club", PassageIDs: []string{"source1:12", "source1:13"}},
	{Name: "Crossbow", PassageIDs: []string{"source1:16", "source1:17"}},
	{Name: "Juggernaut", PassageIDs: []string{"source1:7", "source1:8"}},
}

// coreRequest is the fixture request with source techniques.
func coreRequest(t *testing.T, techniques []unit.SourceTechnique) unit.Request {
	t.Helper()
	prepared, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	request := prepared.Request
	request.SourceTechniques = &techniques
	return request
}

func omission(name, importance, reason string) *s.Object {
	return s.NewObject().Set("name", name).Set("importance", importance).Set("reason", reason)
}

func omit(plan *s.Object, entries ...*s.Object) {
	omitted, _ := plan.Get("omittedTechniques")
	list := omitted.([]any)
	for _, entry := range entries {
		list = append(list, entry)
	}
	plan.Set("omittedTechniques", list)
}

func repertoireEntry(plan *s.Object, name string) *s.Object {
	repertoire, _ := plan.Get("repertoire")
	for _, entry := range repertoire.([]any) {
		if at(entry, "name") == name {
			return entry.(*s.Object)
		}
	}
	return nil
}

var juggernautOmitted = omission("Juggernaut", "minor", "Juggernaut's bonus damage to Ceramic and Fortified needs properties this brief does not use; Spiked Ball already carries its heavier ball.")

// Under requireCoreConcepts a plan ranks its repertoire and omitted
// techniques, lists every source technique, and adapts each core entry on a
// purchase (#61). Code checks only presence and naming; the review judges
// the spirit.
func TestCoreConceptPlanCheck(t *testing.T) {
	request := coreRequest(t, fixtureSourceTechniques)
	decode := func(edit func(plan *s.Object)) error {
		plan := recordedOutput(t, "plan")
		edit(plan)
		_, err := unit.DecodeDesignPlan(plan, &request)
		return err
	}
	cases := []struct {
		name string
		edit func(plan *s.Object)
		want string
	}{
		{"a well-formed plan", func(plan *s.Object) { omit(plan, juggernautOmitted) }, ""},
		{"a core entry no purchase adapts", func(plan *s.Object) {
			repertoireEntry(plan, "Crossbow").Set("importance", "major")
			repertoire, _ := plan.Get("repertoire")
			plan.Set("repertoire", append(repertoire.([]any), s.NewObject().
				Set("name", "Juggernaut").Set("importance", "core").Set("sourceIds", []any{"source1:7"}).
				Set("limitation", "Its bonus damage needs properties this brief does not use.").
				Set("effects", []any{s.NewObject().Set("effect", "A larger ball hits many more bloons.").Set("adaptedAs", []any{"pierce"}).Set("reason", "More bloons hit by one ball is pierce.")})))
		}, `repertoire.4: No purchase names the core entry "Juggernaut" as its technique. ` + m.CoreConceptsRule},
		{"a core technique omitted whole", func(plan *s.Object) {
			omit(plan, omission("Juggernaut", "core", "The Definition cannot express Juggernaut."))
		}, `omittedTechniques.2: omittedTechniques omits the core technique "Juggernaut" whole. ` + m.CoreConceptsRule},
		{"a core entry omitted whole", func(plan *s.Object) {
			omit(plan, juggernautOmitted, omission("crossbow", "major", "Counted critical shots have no operator."))
		}, `omittedTechniques.3: omittedTechniques omits the core entry "Crossbow" whole. ` + m.CoreConceptsRule},
		{"an omitted aspect of a core entry", func(plan *s.Object) {
			omit(plan, juggernautOmitted, omission("Crossbow critical shots", "minor", "Every tenth shot dealing extra damage needs a shot counter operator."))
		}, ""},
		{"a source technique listed nowhere", func(*s.Object) {},
			`repertoire: The plan leaves out the source technique "Juggernaut", which is neither in the repertoire nor in omittedTechniques. ` + m.CoreConceptsRule},
		{"a name variant lists it", func(plan *s.Object) {
			omit(plan, omission("Juggernaut ball", "minor", "Juggernaut's bonus damage needs properties this brief does not use."))
		}, ""},
		{"an unranked entry", func(plan *s.Object) {
			omit(plan, juggernautOmitted)
			repertoireEntry(plan, "Fan Club").Delete("importance")
		}, `repertoire: Repertoire "Fan Club" has no importance. ` + m.CoreConceptsRule},
		{"no core entry", func(plan *s.Object) {
			omit(plan, juggernautOmitted)
			for _, name := range []string{"Spiked Ball", "Triple Throw", "Crossbow"} {
				repertoireEntry(plan, name).Set("importance", "major")
			}
		}, "repertoire: No repertoire entry is core. " + m.CoreConceptsRule},
		{"four core entries", func(plan *s.Object) {
			omit(plan, juggernautOmitted)
			repertoireEntry(plan, "Fan Club").Set("importance", "core")
		}, `repertoire: 4 repertoire entries are core: "Spiked Ball", "Triple Throw", "Fan Club" and "Crossbow". ` + m.CoreConceptsRule},
	}
	for _, c := range cases {
		err := decode(c.edit)
		switch {
		case c.want == "" && err != nil:
			t.Errorf("%s: rejected: %v", c.name, err)
		case c.want != "" && (err == nil || !strings.Contains(err.Error(), c.want)):
			t.Errorf("%s: want %q in %v", c.name, c.want, err)
		}
	}
}

// A passage several source techniques share lists none of them: an entry
// citing a sentence that names all three kinds of Haki does not list
// Armament Haki. A technique whose passages are all shared is listed by any
// of them, as an attack of a form is by the form's entry.
func TestSharedPassagesListOnlyTheirOwnTechnique(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan := *stages.Draft.Run.DesignPlan
	request := coreRequest(t, []unit.SourceTechnique{
		{Name: "Heavy Ball", PassageIDs: []string{"source1:6", "source1:7"}},
		{Name: "Ultra Ball", PassageIDs: []string{"source1:7", "source1:8"}},
	})
	plan.Repertoire[0].SourceIDs = []string{"source1:7"}
	got := messages(unit.CoreConceptIssues(plan, &request))
	if !strings.Contains(got, `the source techniques "Heavy Ball" and "Ultra Ball"`) {
		t.Errorf("a shared passage listed a technique:\n%s", got)
	}
	plan.Repertoire[0].SourceIDs = []string{"source1:6", "source1:7"}
	request.SourceTechniques = &[]unit.SourceTechnique{
		{Name: "Heavy Ball", PassageIDs: []string{"source1:6", "source1:7"}},
		{Name: "Heavy Ball Throw", PassageIDs: []string{"source1:7"}},
	}
	if got := messages(unit.CoreConceptIssues(plan, &request)); strings.Contains(got, "leaves out") {
		t.Errorf("an entry citing a technique's shared passage did not list it:\n%s", got)
	}
}

// The checks apply only under requireCoreConcepts and to a request that
// lists its source techniques. A request prepared before source techniques
// existed, and a saved plan without importance, are not checked, and
// checking a saved draft whose plan has no ranking finds what it found
// before.
func TestCoreConceptChecksKeepOlderArtifacts(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	plan := *stages.Draft.Run.DesignPlan
	for i := range plan.Repertoire {
		plan.Repertoire[i].Importance = ""
	}
	older := coreRequest(t, nil)
	older.SourceTechniques = nil
	if issues := unit.CoreConceptIssues(plan, &older); len(issues) > 0 {
		t.Errorf("a request without source techniques was checked:\n%s", messages(issues))
	}
	off := coreRequest(t, fixtureSourceTechniques)
	definition := *off.MechanicsDefinition
	policy := *definition.Profile.DesignPolicy
	policy.RequireCoreConcepts = nil
	definition.Profile.DesignPolicy = &policy
	off.MechanicsDefinition = &definition
	if issues := unit.CoreConceptIssues(plan, &off); len(issues) > 0 {
		t.Errorf("the policy off still checked:\n%s", messages(issues))
	}

	// A saved draft whose plan has no importance loads and checks as before.
	raw, err := json.Marshal(s.FromGoValue(stages.Draft))
	if err != nil {
		t.Fatal(err)
	}
	unranked := strings.NewReplacer(`"importance":"core",`, "", `"importance":"major",`, "", `"importance":"minor",`, "").Replace(string(raw))
	if unranked == string(raw) {
		t.Fatal("the fixture plan carries no importance")
	}
	value, err := s.Decode([]byte(unranked))
	if err != nil {
		t.Fatal(err)
	}
	draft, err := unit.ParseDraft(value)
	if err != nil {
		t.Fatalf("a draft without importance does not load: %v", err)
	}
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	if len(checked.Findings) != len(stages.Checked.Findings) {
		t.Errorf("checking an unranked draft found %d findings, the ranked one %d", len(checked.Findings), len(stages.Checked.Findings))
	}
}

// The plan prompt states the ranking rule with the source techniques, the
// draft and repair guidance state the spirit rule, and the plan schema asks
// for importance, only under requireCoreConcepts. The omission rule is
// stated whatever the policy.
func TestCoreConceptPrompts(t *testing.T) {
	on, err := fixture.Prepare()
	if err != nil {
		t.Fatal(err)
	}
	techniques := fixtureSourceTechniques
	on.Request.SourceTechniques = &techniques
	off := withPolicy(t, func(p *m.DesignPolicy) { p.RequireCoreConcepts = nil })
	for _, c := range []struct {
		name     string
		prepared unit.Prepared
		want     bool
	}{{"on", on, true}, {"off", off, false}} {
		request, err := unit.DesignPlanRequest(c.prepared)
		if err != nil {
			t.Fatal(err)
		}
		guidance := strings.Join(unit.DesignGuidance(&c.prepared.Request), "\n")
		schema := s.Stringify(request.Schema)
		if strings.Contains(request.Prompt, m.CoreConceptsRule) != c.want || strings.Contains(request.Prompt, unit.CoreSpiritRule) != c.want {
			t.Errorf("%s: the plan prompt states the core concept rules: %v", c.name, !c.want)
		}
		if strings.Contains(guidance, unit.CoreSpiritRule) != c.want {
			t.Errorf("%s: the guidance states the spirit rule: %v", c.name, !c.want)
		}
		if strings.Contains(schema, `"importance"`) != c.want {
			t.Errorf("%s: the plan schema asks for importance: %v", c.name, !c.want)
		}
		if !strings.Contains(request.Prompt, unit.OmissionRule) {
			t.Errorf("%s: the plan prompt does not state the omission rule", c.name)
		}
	}
	request, err := unit.DesignPlanRequest(on)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(request.Prompt, `"sourceTechniques":[{"name":"Spike-o-pult","passageIds":["source1:6"],"signature":false}`) {
		t.Error("the plan context does not list the source techniques")
	}
}

// The review reads each entry's importance, the omitted techniques with
// their reasons and the source techniques, and states the core concept,
// spirit and omission rules. A plan made before the ranking gets no core
// concept rule, and nothing exempts a core entry from being adapted.
func TestReviewReadsCoreConceptsAndOmissions(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	prompt := unit.BlueprintReviewRequest(stages.Checked).Prompt
	for _, want := range []string{
		`{"name":"Spiked Ball","importance":"core","sourceIds":["source1:6"]`,
		`{"name":"Fan Club","importance":"major"`,
		`{"name":"Critical shots","importance":"minor","reason":"Every tenth or fifth shot dealing extra damage needs a shot counter operator."}`,
		`"sourceTechniques":[]`,
		m.CoreConceptsRule,
		unit.CoreSpiritRule,
		unit.OmissionRule,
		"except a core entry of designPlan.repertoire, which the unit must adapt",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the review prompt lacks %q", want)
		}
	}
	checked := stages.Checked
	plan := *checked.Draft.Run.DesignPlan
	plan.Repertoire = append([]unit.PlanRepertoire(nil), plan.Repertoire...)
	plan.OmittedTechniques = append([]unit.PlanOmission(nil), plan.OmittedTechniques...)
	for i := range plan.Repertoire {
		plan.Repertoire[i].Importance = ""
	}
	for i := range plan.OmittedTechniques {
		plan.OmittedTechniques[i].Importance = ""
	}
	checked.Draft.Run.DesignPlan = &plan
	if prompt := unit.BlueprintReviewRequest(checked).Prompt; strings.Contains(prompt, m.CoreConceptsRule) || !strings.Contains(prompt, unit.OmissionRule) {
		t.Error("an unranked plan's review states the core concept rule, or omits the omission rule")
	}
}
