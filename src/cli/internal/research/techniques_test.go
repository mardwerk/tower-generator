package research

import (
	"regexp"
	"strings"
	"testing"
)

var techniquePage = mustURL("https://tensura.fandom.com/wiki/Rimuru_Tempest/Abilities_%26_gear")

const techniqueHTML = `<div class="mw-parser-output"><nav><a href="/wiki/Nav_Blade">Nav</a></nav>
<h2>Skills</h2><h3>Former</h3><h4>Extra skills</h4>
<ul><li><a href="/wiki/Black_Flame_Thunder">Black Flame Thunder</a> is a compound variant.</li><li><a href="/wiki/Black_Flame">Black Flame</a> (Absorbed into Black Flame-Thunder)</li>
<li><a href="/wiki/Other_Flame">Other Flame</a></li><li><a href="/wiki/Black_Lightning">Black Lightning</a></li></ul>
<h4>Common skills</h4><ul><li><a href="/wiki/Water_Blade">Water Blade</a> (Evolved into Water Manipulation)</li>
<li><a href="https://evil.test/wiki/Evil_Blade">Evil Blade</a></li><li><a href="/wiki/File:Blade">File</a></li>
<li><a href="/wiki/Bad%2FBlade">Subpage</a></li><li><a href="/wiki/Bad_Blade?x=1">Query</a></li></ul>
<h2>Analyzed and Subordinates Skills</h2><p><a href="/wiki/Other_Punch">Other Punch</a></p>
<h2>Equipment</h2><p><a href="/wiki/Sword">Sword</a></p></div>`

func titles(links []techniqueLink) string {
	var out []string
	for _, link := range links {
		out = append(out, link.title)
	}
	return strings.Join(out, ", ")
}

func TestObservedCombatLinksKeepOwnershipContext(t *testing.T) {
	links := linkedTechniques(techniqueHTML, techniquePage, "Rimuru Tempest")
	if got := titles(links); got != "Black Flame Thunder, Black Flame, Other Flame, Black Lightning, Water Blade" {
		t.Errorf("links %s", got)
	}
	selected := selectTechniques(append(links, links...))
	if got := titles(selected); got != "Water Blade, Black Flame" {
		t.Fatalf("selected %s", got)
	}
	water := selected[0]
	if strings.Join(water.headings, " > ") != "Skills > Former > Common skills" || water.linkText != "Water Blade" || water.context != "Water Blade (Evolved into Water Manipulation)" || water.parent != techniquePage.String() {
		t.Errorf("water %+v", water)
	}
}

func TestTechniqueDescriptionsKeepBehaviorAndOmitOtherUsers(t *testing.T) {
	link := selectTechniques(linkedTechniques(techniqueHTML, techniquePage, "Rimuru Tempest"))[0]
	quote := "The user adds rotation to magicule-infused water and sprays it at high-velocity. The thin water blade has a severing effect and deals only physical damage."
	source := techniqueSource(`<div class="mw-parser-output"><h2>Abilities</h2><p>`+quote+`</p><h2>Known Users</h2><p>Someone else also has an unrelated supreme power.</p><h2>Trivia</h2><p>Unrelated trivia should not enter the extraction.</p><nav><p>Navigation fake attack description should never appear.</p></nav></div>`, link)
	if source == nil || !strings.Contains(source.Text, quote) || !strings.Contains(source.Text, "Skills > Former > Common skills") ||
		!strings.Contains(source.Text, "Former, evolved or absorbed entries do not establish current availability") ||
		!strings.Contains(source.Text, "Parent passage: Water Blade (Evolved into Water Manipulation)") ||
		regexp.MustCompile(`supreme power|Unrelated trivia|Navigation fake`).MatchString(source.Text) ||
		source.Origin.Location != link.url.String() || !strings.Contains(*source.Origin.Note, "not independently verified canon") {
		t.Fatalf("source %+v", source)
	}
	if techniqueSource(`<nav><p>Only navigation is available on this missing page.</p></nav>`, link) != nil {
		t.Error("navigation became a source")
	}
	bounded := techniqueSource(`<h2>Abilities</h2><p>Creates super-heated flames to attack the target.</p><p>`+strings.Repeat("x", 5000)+`</p><p>See <a href="/wiki/Extra_Blade">another technique</a> for an unrelated variant.</p>`, linkedTechniques(techniqueHTML, techniquePage, "Rimuru Tempest")[0])
	if bounded == nil || len(bounded.Text) >= 5500 || strings.Contains(bounded.Text, strings.Repeat("x", 100)) {
		t.Errorf("bounded %v", bounded != nil)
	}
}

func TestPowerSectionsLeadToTheirTechniquePages(t *testing.T) {
	links := linkedTechniques(escanorHTML, escanorPage, "Escanor")
	var leads []string
	for _, link := range links {
		if link.lead {
			leads = append(leads, link.title)
		}
	}
	// Sunshine has no attack-family word; the Power Level link to it is not
	// a lead, and Merlin's bold link sits in Relationships.
	if strings.Join(leads, ", ") != "Sunshine, Rhitta" || strings.Contains(titles(links), "Merlin") {
		t.Fatalf("leads %v links %s", leads, titles(links))
	}
	followed := followedTechniques(links)
	if titles(followed) != "Sunshine, Rhitta" || strings.Join(followed[0].headings, " > ") != "Abilities and Equipment > Abilities" {
		t.Errorf("followed %s", titles(followed))
	}
	// A "Further information" link may name a subpage; the deepest sections
	// come first and at most four pages are followed.
	luffy := `<div class="mw-parser-output"><h2>Devil Fruit</h2><dl><dd><i>Further information: <a href="/wiki/Gomu_Gomu_no_Mi">Gomu Gomu no Mi</a></i></dd></dl>
<h3>Gear 2</h3><dl><dd><i>Further information: <a href="/wiki/Gomu_Gomu_no_Mi/Gear_2_Techniques">Gomu Gomu no Mi/Gear 2 Techniques</a></i></dd></dl>
<h3>Gear 3</h3><dl><dd><i>Further information: <a href="/wiki/Gomu_Gomu_no_Mi/Gear_3_Techniques">Gomu Gomu no Mi/Gear 3 Techniques</a></i></dd></dl>
<h3>Gear 4</h3><dl><dd><i>Further information: <a href="/wiki/Gomu_Gomu_no_Mi/Gear_4_Techniques">Gomu Gomu no Mi/Gear 4 Techniques</a></i></dd></dl>
<h3>Gear 5</h3><dl><dd><i>Further information: <a href="/wiki/Gomu_Gomu_no_Mi/Gear_5_Techniques">Gomu Gomu no Mi/Gear 5 Techniques</a></i></dd></dl>
<p>See <a href="/wiki/Monkey_D._Luffy/History">his history</a> and <a href="/wiki/Other/Subpage">a subpage</a>.</p>
<h2>Haki</h2><dl><dd><i>Further information: <a href="/wiki/Haki">Haki</a></i></dd></dl></div>`
	chosen := followedTechniques(linkedTechniques(luffy, mustURL(luffyPage.String()+"/Abilities_and_Powers"), "Monkey D. Luffy"))
	if got := titles(chosen); got != "Gomu Gomu no Mi/Gear 2 Techniques, Gomu Gomu no Mi/Gear 3 Techniques, Gomu Gomu no Mi/Gear 4 Techniques, Gomu Gomu no Mi/Gear 5 Techniques" {
		t.Errorf("chosen %s", got)
	}
}

func TestTechniquePagesKeepTheCharactersOwnEntries(t *testing.T) {
	link := followedTechniques(linkedTechniques(escanorHTML, escanorPage, "Escanor"))[0]
	source := techniqueSource(`<div class="mw-parser-output"><p>Sunshine is one of the four Graces created by the Supreme Deity.</p>
<h2>Description</h2><h3>Escanor</h3><p>Starting from sunrise, Escanor's power increases and he grows larger.</p>
<h3>Mael</h3><p>MAEL EXCLUDED: Mael does not undergo any physical change.</p>
<h2>Techniques</h2><h3>Escanor</h3>
<ul><li><b>Cruel Sun</b>: Escanor creates a miniature Sun that melts nearby armor.
<ul><li><b>Pride Flare</b>: Escanor causes his Cruel Sun to flare with intense heat.</li></ul></li></ul>
<h3>Mael</h3><ul><li><b>Greatest Sun</b>: MAEL EXCLUDED: Mael creates a huge miniature sun.</li></ul>
<h2>Trivia</h2><p>TRIVIA EXCLUDED: an unrelated trivia item about the anime.</p></div>`, link)
	if source == nil {
		t.Fatal("no source")
	}
	for _, want := range []string{"Sunshine: Sunshine is one of the four Graces", "Sunshine: Starting from sunrise", "Sunshine: Cruel Sun: Escanor creates a miniature Sun that melts nearby armor.", "Sunshine: Pride Flare: Escanor causes"} {
		if !strings.Contains(source.Text, want) {
			t.Errorf("missing %q in %q", want, source.Text)
		}
	}
	if strings.Contains(source.Text, "EXCLUDED") || !strings.Contains(source.Text, "Section: Abilities and Equipment > Abilities") {
		t.Errorf("text %q", source.Text)
	}
}
