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
// for the Dart Monkey brief. The fixture plan names each but Juggernaut in
// its repertoire.
var fixtureSourceTechniques = []unit.SourceTechnique{
	{Name: "Spiked Ball", PassageIDs: []string{"source1:6"}},
	{Name: "Triple Throw", PassageIDs: []string{"source1:11"}, Signature: true},
	{Name: "Fan Club", PassageIDs: []string{"source1:12", "source1:13"}},
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

// unadapt leaves every effect of a repertoire entry unadapted.
func unadapt(entry *s.Object) {
	effects, _ := entry.Get("effects")
	for _, effect := range effects.([]any) {
		effect.(*s.Object).Set("adaptedAs", []any{})
	}
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
		{"a core entry omitted whole with a qualifier word", func(plan *s.Object) {
			omit(plan, juggernautOmitted, omission("Crossbow attacks", "major", "Counted critical shots have no operator."))
		}, `omittedTechniques.3: omittedTechniques omits the core entry "Crossbow" whole. ` + m.CoreConceptsRule},
		{"an omitted aspect of a core entry", func(plan *s.Object) {
			omit(plan, juggernautOmitted, omission("Crossbow critical shots", "minor", "Every tenth shot dealing extra damage needs a shot counter operator."))
		}, ""},
		{"a source technique listed nowhere", func(*s.Object) {},
			`repertoire: The plan leaves out the source technique "Juggernaut", which is neither in the repertoire nor in omittedTechniques. ` + m.CoreConceptsRule},
		{"a name with a qualifier word lists it", func(plan *s.Object) {
			omit(plan, omission("Juggernaut attacks", "minor", "Juggernaut's bonus damage needs properties this brief does not use."))
		}, ""},
		{"a more specific name does not list it", func(plan *s.Object) {
			omit(plan, omission("Juggernaut ball", "minor", "Juggernaut's bonus damage needs properties this brief does not use."))
		}, `The plan leaves out the source technique "Juggernaut"`},
		{"a core entry adapted only by a proposed mechanic", func(plan *s.Object) {
			omit(plan, juggernautOmitted)
			repertoireEntry(plan, "Crossbow").Set("importance", "major")
			repertoireEntry(plan, "Fan Club").Set("importance", "core")
			unadapt(repertoireEntry(plan, "Fan Club"))
		}, ""},
		{"a core entry adapted with neither", func(plan *s.Object) {
			omit(plan, juggernautOmitted)
			unadapt(repertoireEntry(plan, "Crossbow"))
		}, `repertoire.3: No purchase whose technique is the core entry "Crossbow" promises what one of its effects is adapted as or names a proposed mechanic. ` + m.CoreConceptsRule},
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

// luffySections is a sectioned Luffy request: a character-wiki page whose
// passages carry their section, and the source techniques prepare lists
// from them.
func luffySections(t *testing.T) (unit.Request, map[string]string) {
	t.Helper()
	request := sectionedRequest(t,
		sourceDocument("character-wiki:onepiece.fandom.com:Monkey_D._Luffy/Abilities_and_Powers", strings.Join([]string{
			"Devil Fruit\nGum-Gum Pistol: Luffy stretches his arm back and snaps it forward into a punch.",
			"Devil Fruit\nGear 4\nGear 4 inflates his muscles with Armament Haki and makes him bounce.",
			"Devil Fruit\nGear 4\nKong Gun: Luffy compresses his fist into his arm and fires it as a giant punch.",
			"Devil Fruit\nGear 4\nPython: Luffy bends the path of a punch around a guard.",
			"Devil Fruit\nGear 4\nIn Gear 4 his Gum-Gum Pistol becomes the Kong Gun.",
			"Haki\nHaki lets Luffy use three kinds of power.",
			"Haki\nArmament Haki\nArmament Haki lets him hit Logia users.",
			"Haki\nObservation Haki\nObservation Haki lets him foresee attacks.",
			"Haki\nSupreme King Haki\nSupreme King Haki knocks out weaker foes.",
		}, "\n\n")),
	)
	ids := map[string]string{}
	for _, span := range unit.AuthorEvidence(&request) {
		ids[span.Text] = span.ID
	}
	return request, ids
}

// leftOut is the listing check's message naming the source techniques a
// plan leaves out, or empty when it lists them all.
func leftOut(t *testing.T, plan unit.DesignPlan, request unit.Request) string {
	t.Helper()
	for _, issue := range unit.CoreConceptIssues(plan, &request) {
		if strings.Contains(issue.Message, "leaves out") {
			return issue.Message
		}
	}
	return ""
}

func corePlan(base unit.PlanBase, entries ...unit.PlanRepertoire) unit.DesignPlan {
	return unit.DesignPlan{Base: base, Repertoire: entries, OmittedTechniques: []unit.PlanOmission{}}
}

// A generic entry lists no more specific source technique: a repertoire
// entry named "Haki" does not list Armament, Observation or Supreme King
// Haki, even when it cites their passages, and "Armament Haki" lists only
// itself. A more specific entry lists a less specific name only through a
// qualifier word, as "Gear 4 forms" lists Gear 4 (SOL-67-01).
func TestGenericEntryListsNoSpecificTechnique(t *testing.T) {
	request, ids := luffySections(t)
	names := map[string]bool{}
	for _, technique := range *request.SourceTechniques {
		names[technique.Name] = true
	}
	for _, want := range []string{"Gear 4", "Kong Gun", "Python", "Armament Haki", "Observation Haki", "Supreme King Haki"} {
		if !names[want] {
			t.Fatalf("prepare does not list %q: %v", want, *request.SourceTechniques)
		}
	}
	base := unit.PlanBase{Name: "Gum-Gum Pistol", SourceIDs: []string{ids["Gum-Gum Pistol: Luffy stretches his arm back and snaps it forward into a punch."]}}
	haki := unit.PlanRepertoire{Name: "Haki", Importance: "core", SourceIDs: []string{
		ids["Haki lets Luffy use three kinds of power."],
		ids["Armament Haki lets him hit Logia users."],
		ids["Observation Haki lets him foresee attacks."],
		ids["Supreme King Haki knocks out weaker foes."],
	}}
	gear4 := unit.PlanRepertoire{Name: "Gear 4 forms", Importance: "core", SourceIDs: []string{ids["Gear 4 inflates his muscles with Armament Haki and makes him bounce."]}}
	got := leftOut(t, corePlan(base, haki, gear4), request)
	for _, kind := range []string{"Armament Haki", "Observation Haki", "Supreme King Haki"} {
		if !strings.Contains(got, `"`+kind+`"`) {
			t.Errorf("a generic Haki entry lists %s:\n%s", kind, got)
		}
	}
	if strings.Contains(got, `"Gear 4"`) {
		t.Errorf("a Gear 4 forms entry does not list Gear 4:\n%s", got)
	}
	armament := unit.PlanRepertoire{Name: "Armament Haki", Importance: "major", SourceIDs: []string{ids["Armament Haki lets him hit Logia users."]}}
	got = leftOut(t, corePlan(base, armament, gear4), request)
	if strings.Contains(got, `"Armament Haki"`) || !strings.Contains(got, `"Observation Haki"`) || !strings.Contains(got, `"Supreme King Haki"`) {
		t.Errorf("an Armament Haki entry lists another kind of Haki, or not itself:\n%s", got)
	}
}

// Citing a passage lists a source technique only when the technique is a
// part of a form inside the form's own section and the citing entry is that
// form: a Gear 4 entry citing "Kong Gun: …" and "Python: …" under Gear 4
// lists both, while a Pistol base attack citing a Gear 4 passage lists
// neither Gear 4 nor Kong Gun, and a Gear 4 entry citing a Gear 4 passage
// that mentions the Pistol does not list the Pistol (SOL-67-01).
func TestCitationListsOnlyAFormsOwnSubtechniques(t *testing.T) {
	request, ids := luffySections(t)
	gear4Passage := ids["Gear 4 inflates his muscles with Armament Haki and makes him bounce."]
	kongGun := ids["Kong Gun: Luffy compresses his fist into his arm and fires it as a giant punch."]
	python := ids["Python: Luffy bends the path of a punch around a guard."]
	mention := ids["In Gear 4 his Gum-Gum Pistol becomes the Kong Gun."]
	haki := []unit.PlanRepertoire{
		{Name: "Armament Haki", Importance: "major", SourceIDs: []string{ids["Armament Haki lets him hit Logia users."]}},
		{Name: "Observation Haki", Importance: "minor", SourceIDs: []string{ids["Observation Haki lets him foresee attacks."]}},
		{Name: "Supreme King Haki", Importance: "minor", SourceIDs: []string{ids["Supreme King Haki knocks out weaker foes."]}},
	}

	pistol := unit.PlanBase{Name: "Pistol", SourceIDs: []string{gear4Passage, kongGun, mention}}
	got := leftOut(t, corePlan(pistol, haki...), request)
	for _, name := range []string{"Gear 4", "Kong Gun", "Python", "Gum-Gum Pistol"} {
		if !strings.Contains(got, `"`+name+`"`) {
			t.Errorf("a Pistol entry citing Gear 4 passages lists %s:\n%s", name, got)
		}
	}
	unrelated := unit.PlanRepertoire{Name: "Pistol", Importance: "core", SourceIDs: []string{gear4Passage}}
	if got := leftOut(t, corePlan(unit.PlanBase{Name: "Gum-Gum Pistol"}, append(haki, unrelated)...), request); !strings.Contains(got, `"Gear 4"`) {
		t.Errorf("a Pistol repertoire entry citing a Gear 4 passage lists Gear 4:\n%s", got)
	}

	gear4 := unit.PlanRepertoire{Name: "Gear 4", Importance: "core", SourceIDs: []string{gear4Passage, kongGun, python, mention}}
	got = leftOut(t, corePlan(unit.PlanBase{Name: "Stretch punch"}, append(haki, gear4)...), request)
	if strings.Contains(got, `"Kong Gun"`) || strings.Contains(got, `"Python"`) || strings.Contains(got, `"Gear 4"`) {
		t.Errorf("a Gear 4 entry does not list its own subtechniques:\n%s", got)
	}
	if !strings.Contains(got, `"Gum-Gum Pistol"`) {
		t.Errorf("a Gear 4 entry citing a passage that mentions the Pistol lists it:\n%s", got)
	}
	gear4.SourceIDs = []string{gear4Passage}
	if got := leftOut(t, corePlan(unit.PlanBase{Name: "Gum-Gum Pistol"}, append(haki, gear4)...), request); !strings.Contains(got, `the source techniques "Kong Gun" and "Python"`) {
		t.Errorf("a Gear 4 entry lists subtechniques whose passages it does not cite:\n%s", got)
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
	if !strings.Contains(request.Prompt, `"sourceTechniques":[{"name":"Spiked Ball","passageIds":["source1:6"],"signature":false}`) {
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

// coreFindings checks the fixture draft with its plan edited and returns
// its core concept Findings.
func coreFindings(t *testing.T, edit func(plan *unit.DesignPlan)) []unit.Finding {
	t.Helper()
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	draft := stages.Draft
	plan := *draft.Run.DesignPlan
	plan.Repertoire = append([]unit.PlanRepertoire(nil), plan.Repertoire...)
	for i := range plan.Repertoire {
		plan.Repertoire[i].Effects = append([]unit.PlanEffect(nil), plan.Repertoire[i].Effects...)
	}
	edit(&plan)
	draft.Run.DesignPlan = &plan
	checked, err := unit.CheckDraft(draft)
	if err != nil {
		t.Fatal(err)
	}
	var out []unit.Finding
	for _, f := range checked.Findings {
		if f.Rule == unit.CoreConceptFindingRule {
			out = append(out, f)
		}
	}
	return out
}

// omitEffects leaves every effect of a repertoire entry unadapted.
func omitEffects(plan *unit.DesignPlan, name string) {
	for i := range plan.Repertoire {
		if plan.Repertoire[i].Name != name {
			continue
		}
		for j := range plan.Repertoire[i].Effects {
			plan.Repertoire[i].Effects[j].AdaptedAs = []string{}
		}
	}
}

func rank(plan *unit.DesignPlan, name, importance string) {
	for i := range plan.Repertoire {
		if plan.Repertoire[i].Name == name {
			plan.Repertoire[i].Importance = importance
		}
	}
}

// A core concept that a purchase adapts with a typed change is embodied and
// passes clean. One that its purchases adapt only with a proposed mechanic
// is an unresolved design gap: a proposal grants no behavior, so the Unit
// does not embody it yet. One adapted with neither fails (SOL-67-01).
func TestCoreConceptImplementedOrOnlyProposed(t *testing.T) {
	if found := coreFindings(t, func(*unit.DesignPlan) {}); len(found) > 0 {
		t.Errorf("typed core concepts found %v", found)
	}

	// Fan Club's effects are left unadapted, so only x-5-x's proposed
	// Plasma transformation adapts it.
	found := coreFindings(t, func(plan *unit.DesignPlan) {
		rank(plan, "Crossbow", "major")
		rank(plan, "Fan Club", "core")
		omitEffects(plan, "Fan Club")
	})
	if len(found) != 1 || found[0].Outcome != "unresolved" || found[0].Subject != "paths.path2.tiers.tier5" ||
		!strings.HasPrefix(found[0].Message, `The core concept "Fan Club" is only proposed: x-5-x `) ||
		!strings.Contains(found[0].Message, "until the Definition supports it, no build grants it and the Unit does not embody this core concept. "+m.CoreConceptsRule) {
		t.Errorf("a proposed-only core concept: %+v", found)
	}

	// Crossbow's effects are left unadapted and its purchases propose
	// nothing.
	found = coreFindings(t, func(plan *unit.DesignPlan) { omitEffects(plan, "Crossbow") })
	if len(found) != 1 || found[0].Outcome != "fail" || !strings.Contains(found[0].Message, `No purchase adapts the core concept "Crossbow": x-x-3 `) {
		t.Errorf("a core concept adapted with neither: %+v", found)
	}
}
