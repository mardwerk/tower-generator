package unit_test

import (
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/research"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// savedSources prepares a Sources file of testdata under the Default: the
// Luffy and Escanor Sources researched for the v30 runs of #61
// (OPUS-NET-61-6).
func savedSources(t *testing.T, name string) unit.Request {
	t.Helper()
	data, err := os.ReadFile("testdata/" + name + ".sources.json")
	if err != nil {
		t.Fatal(err)
	}
	value, err := s.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	sources, err := research.ParseSources(value)
	if err != nil {
		t.Fatal(err)
	}
	prepared, err := sources.Prepare(unit.DefaultProfile())
	if err != nil {
		t.Fatal(err)
	}
	return prepared.Request
}

func techniquesByName(request unit.Request) map[string]unit.SourceTechnique {
	out := map[string]unit.SourceTechnique{}
	for _, technique := range *request.SourceTechniques {
		out[technique.Name] = technique
	}
	return out
}

// One passage proves that Rhitta and The Divine Axe Rhitta are one axe: the
// first passage of the Rhitta page defines its subject by the longer name.
// The candidate keeps the page's name, lists the other as an alias and keeps
// the passages of both (SOL-61-05).
func TestEscanorRhittaAliasesMerge(t *testing.T) {
	request := savedSources(t, "escanor")
	techniques := techniquesByName(request)
	if _, ok := techniques["The Divine Axe Rhitta"]; ok {
		t.Fatalf("The Divine Axe Rhitta is still its own candidate: %v", *request.SourceTechniques)
	}
	rhitta, ok := techniques["Rhitta"]
	if !ok || !slices.Equal(rhitta.Aliases, []string{"The Divine Axe Rhitta"}) {
		t.Fatalf("Rhitta %+v", rhitta)
	}
	legacy := unit.LegacySourceTechniques(&request)
	for _, technique := range legacy {
		if technique.Name != "Rhitta" && technique.Name != "The Divine Axe Rhitta" {
			continue
		}
		for _, id := range technique.PassageIDs {
			if !slices.Contains(rhitta.PassageIDs, id) {
				t.Errorf("the merged Rhitta lost %s of %s", id, technique.Name)
			}
		}
	}
	// Every other candidate stays its own: containment such as "Divine
	// Sword Escanor" and "Divine Spear Escanor" proves nothing.
	if len(legacy) != len(techniques)+1 {
		t.Errorf("legacy %d candidates, merged %d", len(legacy), len(techniques))
	}
}

// Names that share words are not aliases without a passage that says so:
// Gear 2 and Gear 2 Buso, and the three kinds of Haki, stay apart.
func TestSharedWordsDoNotMerge(t *testing.T) {
	luffy := techniquesByName(savedSources(t, "luffy"))
	for _, name := range []string{"Gear 2", "Gear 2 Buso", "Armament Haki", "Observation Haki", "Supreme King Haki", "Gear 4", "Kin'niku Fusen"} {
		technique, ok := luffy[name]
		if !ok {
			t.Errorf("Luffy's Sources lack %s", name)
			continue
		}
		if len(technique.Aliases) > 0 {
			t.Errorf("%s has aliases %v", name, technique.Aliases)
		}
	}
	request := sectionedRequest(t,
		sourceDocument("character-wiki:onepiece.fandom.com:Monkey_D._Luffy/Abilities_and_Powers", strings.Join([]string{
			"Haki\nHaki is a power every living being has.",
			"Haki\nArmament Haki\nArmament Haki is a form of Haki that hardens the body.",
			"Haki\nObservation Haki\nObservation Haki is a form of Haki that senses others.",
			"Devil Fruit\nGear 2\nGear 2 is a technique that pumps blood faster.",
			"Devil Fruit\nGear 2\nGear 2 Buso: Luffy coats his limbs in Armament Haki while in Gear 2.",
		}, "\n\n")),
	)
	for _, technique := range *request.SourceTechniques {
		if len(technique.Aliases) > 0 {
			t.Errorf("%s has aliases %v", technique.Name, technique.Aliases)
		}
	}
}

// A sentence that calls one technique by another's name merges the two; the
// first named keeps its name.
func TestAlsoCalledMerges(t *testing.T) {
	request := sectionedRequest(t,
		sourceDocument("character-wiki:onepiece.fandom.com:Monkey_D._Luffy/Abilities_and_Powers", strings.Join([]string{
			"Haki\nBusoshoku Haki\nBusoshoku Haki is his strongest power.",
			"Haki\nArmament Haki\nArmament Haki lets him hit Logia users.",
			"Haki\nNotes\nBusoshoku Haki, also known as Armament Haki, hardens his fists.",
		}, "\n\n")),
	)
	techniques := techniquesByName(request)
	if _, ok := techniques["Armament Haki"]; ok {
		t.Fatalf("Armament Haki stays its own: %v", *request.SourceTechniques)
	}
	if merged := techniques["Busoshoku Haki"]; !slices.Equal(merged.Aliases, []string{"Armament Haki"}) {
		t.Errorf("Busoshoku Haki %+v", merged)
	}
}

// Salience is strong for a followed page, a signature technique or a
// technique whose own sections hold a tenth of the character article's
// power and ability sections; otherwise normal.
func TestSourceTechniqueSalience(t *testing.T) {
	for character, want := range map[string]map[string]string{
		"luffy": {
			"Gum-Gum Pistol": unit.SalienceStrong, // signature
			"Gear 2":         unit.SalienceStrong, // followed page
			"Gear 5":         unit.SalienceStrong,
			"Armament Haki":  unit.SalienceStrong, // 14% of the sections
			"Python":         unit.SalienceNormal,
			"Gear 2 Buso":    unit.SalienceNormal,
		},
		"escanor": {
			"Sunshine":    unit.SalienceStrong,
			"Rhitta":      unit.SalienceStrong,
			"The One":     unit.SalienceStrong, // 14% of the sections
			"Cruel Sun":   unit.SalienceNormal,
			"Large Spear": unit.SalienceNormal,
		},
	} {
		techniques := techniquesByName(savedSources(t, character))
		for name, salience := range want {
			if got := techniques[name].Salience; got != salience {
				t.Errorf("%s %s: salience %q, want %q", character, name, got, salience)
			}
		}
		for _, technique := range techniques {
			if technique.Salience != unit.SalienceStrong && technique.Salience != unit.SalienceNormal {
				t.Errorf("%s %s: salience %q", character, technique.Name, technique.Salience)
			}
		}
	}
}

// The plan lists a source technique by any of its names.
func TestAliasListsSourceTechnique(t *testing.T) {
	request := savedSources(t, "escanor")
	techniques := *request.SourceTechniques
	var rhitta unit.SourceTechnique
	for _, technique := range techniques {
		if technique.Name == "Rhitta" {
			rhitta = technique
		}
	}
	for _, name := range []string{"Rhitta", "The Divine Axe Rhitta"} {
		plan := corePlan(unit.PlanBase{Name: "Sunshine Punch"}, unit.PlanRepertoire{Name: name, Importance: "major", SourceIDs: rhitta.PassageIDs[:1]})
		if message := leftOut(t, plan, request); strings.Contains(message, `"Rhitta"`) {
			t.Errorf("%s does not list Rhitta: %s", name, message)
		}
	}
	plan := corePlan(unit.PlanBase{Name: "Sunshine Punch"}, unit.PlanRepertoire{Name: "Axe", Importance: "major", SourceIDs: rhitta.PassageIDs[:1]})
	if message := leftOut(t, plan, request); !strings.Contains(message, `"Rhitta"`) {
		t.Errorf("a generic Axe entry lists Rhitta: %s", message)
	}
}

// A request prepared before aliases and salience keeps verifying: its
// sourceTechniques are compared with the list in the form it was made in.
// A current list edited to that form, or any other edit, does not verify.
func TestLegacySourceTechniquesVerify(t *testing.T) {
	request := savedSources(t, "escanor")
	verify := func(techniques []unit.SourceTechnique) error {
		edited := request
		edited.SourceTechniques = &techniques
		hash, err := unit.HashRequest(edited)
		if err != nil {
			t.Fatal(err)
		}
		return unit.VerifyPrepared(unit.Prepared{SchemaVersion: edited.SchemaVersion, Kind: "prepared", InputHash: hash, Request: edited})
	}
	if err := verify(*request.SourceTechniques); err != nil {
		t.Errorf("current list: %v", err)
	}
	legacy := unit.LegacySourceTechniques(&request)
	if err := verify(legacy); err != nil {
		t.Errorf("legacy list: %v", err)
	}
	stripped := slices.Clone(*request.SourceTechniques)
	for i := range stripped {
		stripped[i].Salience, stripped[i].Aliases = "", nil
	}
	if err := verify(stripped); err == nil {
		t.Error("the merged list without salience verifies")
	}
	flipped := slices.Clone(*request.SourceTechniques)
	flipped[0].Salience = unit.SalienceNormal
	if flipped[0].Salience == (*request.SourceTechniques)[0].Salience {
		flipped[0].Salience = unit.SalienceStrong
	}
	if err := verify(flipped); err == nil {
		t.Error("an edited salience verifies")
	}
}
