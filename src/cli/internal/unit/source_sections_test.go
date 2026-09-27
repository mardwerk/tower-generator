package unit_test

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

func sourceDocument(id, text string) unit.Document {
	return unit.Document{ID: id, Kind: "source", Text: text, Origin: unit.Origin{Location: "https://example.test/" + id, Access: "retrieved"}}
}

func sectionedRequest(t *testing.T, documents ...unit.Document) unit.Request {
	t.Helper()
	request := unit.Request{
		SchemaVersion: "1",
		Task:          "Adapt the character.",
		Character:     unit.Character{Name: "Monkey D. Luffy", Work: "One Piece", Scope: "Adapt the supplied sources."},
		Documents:     documents,
		Constraints:   []unit.Constraint{},
	}
	prepared, err := unit.Prepare(s.FromGoValue(unit.ApplyProfile(request, unit.DefaultProfile())))
	if err != nil {
		t.Fatal(err)
	}
	return prepared.Request
}

func TestSectionedPassagesCarryTheirSectionAndNoHeadingPassages(t *testing.T) {
	request := sectionedRequest(t,
		sourceDocument("character-reference", strings.Join([]string{
			"Luffy is a pirate whose body is rubber.",
			"",
			"== Description ==",
			"",
			"=== Abilities ===",
			"",
			"==== Devil Fruit ====",
			"His signature attack, the Gum-Gum Pistol, stretches his arm into a punch.",
			"",
			"== Reception ==",
			"Critics praised the character.",
		}, "\n")),
		sourceDocument("character-wiki:onepiece.fandom.com:Monkey_D._Luffy/Abilities_and_Powers", strings.Join([]string{
			"Devil Fruit\nGear 4\nGear 4 inflates his muscles with Armament Haki.",
			"Devil Fruit\nGear 4\nBoundman bounces around the battlefield.",
			"Haki\nArmament Haki\nArmament Haki lets him hit Logia users.",
		}, "\n\n")),
	)
	if request.SourceTechniques == nil {
		t.Fatal("a prepared request carries its source techniques")
	}
	spans := unit.EvidenceSpans(&request)
	sections := map[string]string{}
	for _, span := range spans {
		if regexp.MustCompile(`==|^(Devil Fruit|Gear 4|Haki|Armament Haki)$`).MatchString(span.Text) || strings.Contains(span.Text, "Devil Fruit\n") {
			t.Errorf("heading in passage %s %q", span.ID, span.Text)
		}
		sections[span.Text] = span.Section
	}
	for text, want := range map[string]string{
		"Luffy is a pirate whose body is rubber.":                                   "",
		"His signature attack, the Gum-Gum Pistol, stretches his arm into a punch.": "Description > Abilities > Devil Fruit",
		"Critics praised the character.":                                            "Reception",
		"Gear 4 inflates his muscles with Armament Haki.":                           "Devil Fruit > Gear 4",
		"Armament Haki lets him hit Logia users.":                                   "Haki > Armament Haki",
	} {
		if got, ok := sections[text]; !ok || got != want {
			t.Errorf("%q: section %q (found %v), want %q", text, got, ok, want)
		}
	}
	if spans[0].ID != "source1:0" || spans[len(spans)-1].ID != "source2:2" {
		t.Errorf("ids %s .. %s", spans[0].ID, spans[len(spans)-1].ID)
	}
}

func TestSourceTechniquesNameHeadingsEntriesAndSignatures(t *testing.T) {
	request := sectionedRequest(t,
		sourceDocument("character-reference", "Luffy is a pirate. In his signature attack, the Gum-Gum Pistol, he slingshots punches."),
		sourceDocument("character-wiki:onepiece.fandom.com:Monkey_D._Luffy/Abilities_and_Powers", strings.Join([]string{
			"Devil Fruit\nThe fruit makes his body rubber.",
			"Devil Fruit\nGear 4\nGear 4 inflates his muscles with Armament Haki.",
			"Devil Fruit\nGear 4\nGear 4 lets him fly by bouncing.",
			"Haki\nTechniques\nRed Hawk「火拳銃, Reddo Hōku」: Luffy ignites his fist through friction.",
			"Miscellaneous Abilities\nLuck\nLuffy is very lucky in many situations.",
		}, "\n\n")),
		sourceDocument("character-technique:onepiece.fandom.com:"+"Gomu%20Gomu%20no%20Mi%2FGear%205%20Techniques", strings.Join([]string{
			"Observed link on Monkey D. Luffy's article: https://onepiece.fandom.com/wiki/Monkey_D._Luffy\nSection: Devil Fruit > Gear 5\nLink text: Gear 5\nParent passage: Further information: Gomu Gomu no Mi/Gear 5 Techniques",
			"Ownership and period limit: the parent lists this technique in the section above.",
			"Gomu Gomu no Mi/Gear 5 Techniques: Gear 5 is the awakened form of the fruit.",
		}, "\n\n")),
	)
	techniques := *request.SourceTechniques
	got := map[string]unit.SourceTechnique{}
	var names []string
	for _, technique := range techniques {
		got[technique.Name] = technique
		names = append(names, technique.Name)
	}
	for _, want := range []string{"Gum-Gum Pistol", "Gear 4", "Red Hawk", "Gear 5"} {
		if _, ok := got[want]; !ok {
			t.Errorf("missing %s in %v", want, names)
		}
	}
	for _, unwanted := range []string{"Devil Fruit", "Haki", "Techniques", "Luck", "Miscellaneous Abilities", "Section", "Observed link on Monkey D"} {
		if _, ok := got[unwanted]; ok {
			t.Errorf("%s listed in %v", unwanted, names)
		}
	}
	if !got["Gum-Gum Pistol"].Signature || got["Gear 4"].Signature {
		t.Errorf("signatures %+v", techniques)
	}
	ids := map[string]unit.EvidenceSpan{}
	for _, span := range unit.AuthorEvidence(&request) {
		ids[span.ID] = span
	}
	for _, technique := range techniques {
		if len(technique.PassageIDs) == 0 {
			t.Errorf("%s has no passages", technique.Name)
		}
		for _, id := range technique.PassageIDs {
			if _, ok := ids[id]; !ok {
				t.Errorf("%s cites %s, which the selection does not keep", technique.Name, id)
			}
		}
	}
	if gear4 := got["Gear 4"].PassageIDs; len(gear4) != 2 || ids[gear4[0]].Section != "Devil Fruit > Gear 4" {
		t.Errorf("Gear 4 passages %v", gear4)
	}
	// The derived list is checked against the passages when a prepared
	// request is verified.
	edited := request
	changed := append([]unit.SourceTechnique{}, techniques...)
	changed[0].Name = "Invented"
	edited.SourceTechniques = &changed
	hash, _ := unit.HashRequest(edited)
	if err := unit.VerifyPrepared(unit.Prepared{SchemaVersion: edited.SchemaVersion, Kind: "prepared", InputHash: hash, Request: edited}); err == nil || !strings.Contains(err.Error(), "sourceTechniques") {
		t.Errorf("edited source techniques: %v", err)
	}
}

func TestSectionRanksPowersBeforeMiscellany(t *testing.T) {
	for _, c := range []struct {
		path []string
		rank int
	}{
		{[]string{"Devil Fruit", "Gear 4"}, unit.PowerSection},
		{[]string{"Haki", "Armament Haki", "Techniques"}, unit.PowerSection},
		{[]string{"Abilities and Equipment", "Weapons"}, unit.PowerSection},
		{[]string{"Abilities and Equipment", "Abilities"}, unit.AbilitySection},
		{[]string{"Physical Abilities", "Strength"}, unit.AbilitySection},
		{[]string{"Miscellaneous Abilities", "Luck"}, unit.OtherSection},
		{[]string{"Miscellaneous Abilities", "Artistic Skill"}, unit.OtherSection},
		{[]string{"Overview"}, unit.OtherSection},
		{[]string{"Creation and conception", "Development"}, unit.OtherSection},
		{nil, unit.OtherSection},
	} {
		if got := unit.SectionRank(c.path); got != c.rank {
			t.Errorf("%v: rank %d, want %d", c.path, got, c.rank)
		}
	}
}

// filler is a paragraph of n sentences without combat words, which scores
// zero, like the plot and reception passages that used to fill the
// selection in page order.
func filler(topic string, n int) string {
	var sentences []string
	for i := range n {
		sentences = append(sentences, fmt.Sprintf("The %s chapter %d was published in a weekly magazine that year.", topic, i))
	}
	return strings.Join(sentences, " ")
}

func TestSectionedSelectionKeepsEveryPowerSection(t *testing.T) {
	wiki := []string{"Luffy is a pirate whose body is rubber.", "== Appearances =="}
	for i := range 40 {
		wiki = append(wiki, filler(fmt.Sprintf("arc %d", i), 8))
	}
	var fandom []string
	for _, form := range []string{"Gear 2", "Gear 3", "Gear 4", "Gear 5"} {
		// Neutral text scores zero, like the Gear 5 passages that never
		// say "Gear 5".
		for i := range 4 {
			fandom = append(fandom, fmt.Sprintf("Devil Fruit\n%s\n%s passage %d was drawn in color for the anniversary.", form, form, i))
		}
	}
	for i := range 30 {
		fandom = append(fandom, "Miscellaneous Abilities\nLuck\n"+filler(fmt.Sprintf("luck %d", i), 1))
	}
	request := sectionedRequest(t,
		sourceDocument("character-reference", strings.Join(wiki, "\n")),
		sourceDocument("character-wiki:onepiece.fandom.com:Monkey_D._Luffy/Abilities_and_Powers", strings.Join(fandom, "\n\n")),
	)
	all := unit.EvidenceSpans(&request)
	selected := unit.AuthorEvidence(&request)
	if len(all) <= 200 || len(selected) > 200 {
		t.Fatalf("available %d, selected %d", len(all), len(selected))
	}
	kept := map[string]int{}
	for _, span := range selected {
		kept[span.Section]++
	}
	for _, form := range []string{"Gear 2", "Gear 3", "Gear 4", "Gear 5"} {
		if kept["Devil Fruit > "+form] != 4 {
			t.Errorf("%s kept %d of 4: %v", form, kept["Devil Fruit > "+form], kept)
		}
	}
	// Passages of equal score come from each section in turn, not from
	// the longest section in page order.
	if kept["Miscellaneous Abilities > Luck"] == 0 || kept["Appearances"] == 0 {
		t.Errorf("sections did not take turns: %v", kept)
	}
	// The legacy selection of the same documents fills its slots in page
	// order and loses the forms.
	legacy := request
	legacy.SourceTechniques = nil
	for _, span := range unit.AuthorEvidence(&legacy) {
		if strings.Contains(span.Text, "Gear 5 passage") {
			t.Errorf("the legacy selection kept %q", span.Text)
		}
	}
}

// Saved artifacts keep their passages: a request without sourceTechniques
// splits and selects as before, so the citations of its draft stay valid
// and its hash is unchanged.
func TestSavedRequestsKeepTheirPassages(t *testing.T) {
	data, err := os.ReadFile("../../../../data/reference/captures/luffy-3b.result.json")
	if err != nil {
		t.Fatal(err)
	}
	value, err := s.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	result, err := unit.ParseResult(value)
	if err != nil {
		t.Fatal(err)
	}
	request := result.Prepared.Request
	if request.SourceTechniques != nil {
		t.Fatal("the capture predates source techniques")
	}
	if err := unit.VerifyPrepared(result.Prepared); err != nil {
		t.Fatalf("saved hash: %v", err)
	}
	selected := map[string]bool{}
	for _, span := range unit.AuthorEvidence(&request) {
		if span.Section != "" {
			t.Fatalf("legacy passage %s carries a section", span.ID)
		}
		selected[span.ID] = true
	}
	cited := regexp.MustCompile(`"(source[0-9]+:[0-9]+)"`).FindAllStringSubmatch(string(data), -1)
	if len(cited) == 0 {
		t.Fatal("the capture cites no passages")
	}
	for _, match := range cited {
		if !selected[match[1]] {
			t.Errorf("citation %s is no longer a selected passage", match[1])
		}
	}
	if slices.ContainsFunc(unit.EvidenceSpans(&request), func(span unit.EvidenceSpan) bool { return span.Section != "" }) {
		t.Error("legacy passages carry sections")
	}
	if strings.Contains(s.Stringify(s.FromGoValue(request)), "sourceTechniques") {
		t.Error("reading added sourceTechniques")
	}
}
