package research

import (
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// techniqueLink is a combat technique linked from a character's own page.
type techniqueLink struct {
	url       *url.URL
	parent    string
	character string
	title     string
	linkText  string
	context   string
	headings  []string
	family    string
	// lead marks the page a power section names as its own: the bold term
	// an entry starts with, or a "Further information" or "Main article"
	// link.
	lead bool
}

const techniqueNoise = ".article-tabs,script,style,iframe,noscript,nav,aside,table,figure,.thumb,.gallery,.portable-infobox,.navbox,.navibox,.navigation,.toc,#toc,.mw-editsection,.mw-references-wrap,.references,.reference,.printfooter,.catlinks"

var (
	attackFamilies = []struct {
		name    string
		pattern *regexp.Regexp
	}{
		{"cutting", regexp.MustCompile(`(?i)\b(blade|slash|sword|kunai|shuriken|arrow)\b`)},
		{"fire", regexp.MustCompile(`(?i)\b(fire|flame|fireball|burning)\b`)},
		{"impact", regexp.MustCompile(`(?i)\b(punch|kick|strike|fist)\b`)},
		{"energy", regexp.MustCompile(`(?i)\b(beam|lightning|thunder|bolt)\b`)},
		{"freezing", regexp.MustCompile(`(?i)\b(ice|frost|freezing)\b`)},
	}
	combatSection   = regexp.MustCompile(`(?i)\b(abilities|powers|skills|techniques|combat|arts|magic)\b`)
	excludedSection = regexp.MustCompile(`(?i)\b(subordinates?|analy[sz]ed|references|navigation|trivia|gallery)\b`)
	// equipmentSection excludes an equipment section, unless its heading
	// also names abilities, as in "Abilities and Equipment".
	equipmentSection = regexp.MustCompile(`(?i)\bequipment\b`)
	mainArticle      = regexp.MustCompile(`(?i)^(?:further information|main articles?|main page)\s*:`)
	fandomHost       = regexp.MustCompile(`^[a-z0-9-]+\.fandom\.com$`)
	descriptionPart  = regexp.MustCompile(`(?i)\b(abilities|powers|usage|effects|description|techniques?)\b`)
	unrelatedSection = regexp.MustCompile(`(?i)\b(users|trivia|references|navigation|related|gallery)\b`)
)

func family(title string) string {
	for _, f := range attackFamilies {
		if f.pattern.MatchString(title) {
			return f.name
		}
	}
	return ""
}

// excluded reports a section whose links are never followed.
func excluded(headings []section) bool {
	for _, h := range headings {
		if excludedSection.MatchString(h.text) || (equipmentSection.MatchString(h.text) && !combatSection.MatchString(h.text)) {
			return true
		}
	}
	return false
}

// linkedTechniques lists technique links observed in the combat sections of
// an identity-checked character page: links whose title names an attack
// family, and the lead links of power and ability sections, whatever their
// title.
func linkedTechniques(content string, page *url.URL, character string) []techniqueLink {
	document := parseHTML(content)
	document.Find(techniqueNoise).Remove()
	root := articleRoot(document)
	var headings []section
	var links []techniqueLink
	ability := false
	if title, ok := fandomTitle(page); ok {
		if _, suffix, found := strings.Cut(title, "/"); found {
			ability = abilityPage.MatchString(suffix)
		}
	}
	root.Find("h2,h3,h4,h5,h6,a[href]").Each(func(_ int, node *goquery.Selection) {
		if level := headingLevel(node); level > 0 {
			headings = pushHeading(headings, level, clean(node.Text()))
			return
		}
		path := headingTexts(headings)
		power := unit.SectionRank(path) != unit.OtherSection
		if !(power || anyHeadingMatches(headings, combatSection) || (ability && len(headings) > 0)) || excluded(headings) {
			return
		}
		context := clean(node.Closest("li,p,dd").Text())
		linkText := clean(node.Text())
		if context == "" || s.UTF16Len(context) > 800 || linkText == "" || s.UTF16Len(linkText) > 80 {
			return
		}
		hatnote := mainArticle.MatchString(context)
		bold := node.ParentsFiltered("b,strong").Length() > 0 && strings.HasPrefix(context, linkText)
		lead := power && (hatnote || bold)
		href, _ := node.Attr("href")
		target, err := page.Parse(href)
		if err != nil || target.Scheme != "https" || target.Hostname() != page.Hostname() || !fandomHost.MatchString(target.Hostname()) ||
			target.User != nil || target.Port() != "" || target.RawQuery != "" || !strings.HasPrefix(target.EscapedPath(), "/wiki/") {
			return
		}
		title, err := url.PathUnescape(strings.TrimPrefix(target.EscapedPath(), "/wiki/"))
		if err != nil {
			return
		}
		title = strings.ReplaceAll(title, "_", " ")
		// Only a main-article link may name a subpage, and never one of the
		// character's own pages.
		if title == "" || strings.Contains(title, ":") || (strings.Contains(title, "/") && !hatnote) ||
			strings.EqualFold(title, character) || sameName(strings.Split(title, "/")[0], character) {
			return
		}
		behavior := family(title)
		if behavior == "" && !lead {
			return
		}
		target.Fragment, target.RawFragment = "", ""
		for i, existing := range links {
			if existing.url.String() == target.String() {
				links[i].lead = links[i].lead || lead
				return
			}
		}
		links = append(links, techniqueLink{target, page.String(), character, title, linkText, context, path, behavior, lead})
	})
	return links
}

// maxTechniquePages bounds the technique pages followed for one character.
const maxTechniquePages = 4

// followedTechniques chooses the technique pages to read: lead links first,
// those of the most specific sections before their parents' and otherwise
// in page order, then the attack-family choice of selectTechniques, at most
// maxTechniquePages in all.
func followedTechniques(links []techniqueLink) []techniqueLink {
	var leads []techniqueLink
	for _, link := range links {
		if link.lead {
			leads = append(leads, link)
		}
	}
	sort.SliceStable(leads, func(i, j int) bool { return len(leads[i].headings) > len(leads[j].headings) })
	var out []techniqueLink
	for _, link := range append(leads, selectTechniques(links)...) {
		duplicate := false
		for _, chosen := range out {
			duplicate = duplicate || chosen.url.String() == link.url.String()
		}
		if !duplicate && len(out) < maxTechniquePages {
			out = append(out, link)
		}
	}
	return out
}

// selectTechniques prefers two distinct attack families without following
// further links.
func selectTechniques(links []techniqueLink) []techniqueLink {
	priority := map[string]int{"cutting": 0, "fire": 1, "impact": 2, "energy": 3, "freezing": 4}
	complexity := func(title string) int {
		n := 0
		for _, f := range attackFamilies {
			if f.pattern.MatchString(title) {
				n++
			}
		}
		return n
	}
	var sorted []techniqueLink
	for _, link := range links {
		if link.family != "" {
			sorted = append(sorted, link)
		}
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		if priority[sorted[i].family] != priority[sorted[j].family] {
			return priority[sorted[i].family] < priority[sorted[j].family]
		}
		return complexity(sorted[i].title) < complexity(sorted[j].title)
	})
	var selected []techniqueLink
	for _, link := range sorted {
		duplicate := false
		for _, entry := range selected {
			if entry.url.String() == link.url.String() || entry.family == link.family {
				duplicate = true
			}
		}
		if duplicate {
			continue
		}
		selected = append(selected, link)
		if len(selected) == 2 {
			break
		}
	}
	return selected
}

// otherUsers marks, in page order, the headings of a technique page that
// another user's subsection owns: where a section splits by user and one
// subsection is named for the character, its siblings named for others are
// skipped.
func otherUsers(root *goquery.Selection, character string) []bool {
	type node struct {
		level, parent int
		named         bool
	}
	var nodes []node
	var stack []int
	root.Find("h2,h3,h4,h5,h6").Each(func(_ int, heading *goquery.Selection) {
		level := headingLevel(heading)
		for len(stack) > 0 && nodes[stack[len(stack)-1]].level >= level {
			stack = stack[:len(stack)-1]
		}
		parent := -1
		if len(stack) > 0 {
			parent = stack[len(stack)-1]
		}
		nodes = append(nodes, node{level, parent, namesCharacter(clean(heading.Text()), character)})
		stack = append(stack, len(nodes)-1)
	})
	skip := make([]bool, len(nodes))
	for i, n := range nodes {
		if n.named {
			continue
		}
		for j, sibling := range nodes {
			if j != i && sibling.named && sibling.parent == n.parent && sibling.level == n.level {
				skip[i] = true
				break
			}
		}
	}
	return skip
}

// techniqueSource extracts a technique page's introduction, description and
// technique entries, keeping the parent page's ownership and period
// context. Where the page splits a section by user, only the character's
// own subsection is read.
func techniqueSource(content string, link techniqueLink) *unit.Document {
	document := parseHTML(content)
	document.Find(techniqueNoise).Remove()
	root := articleRoot(document)
	skip := otherUsers(root, link.character)
	var headings []section
	var skipped []bool
	var passages []string
	length, count := 0, 0
	root.Find("h2,h3,h4,h5,h6,p,li").Each(func(_ int, node *goquery.Selection) {
		if level := headingLevel(node); level > 0 {
			other := count < len(skip) && skip[count]
			count++
			for len(headings) > 0 && headings[len(headings)-1].level >= level {
				headings, skipped = headings[:len(headings)-1], skipped[:len(skipped)-1]
			}
			headings = append(headings, section{level, clean(node.Text())})
			skipped = append(skipped, other || (len(skipped) > 0 && skipped[len(skipped)-1]))
			return
		}
		text := ownText(node)
		if len(headings) > 0 && (skipped[len(skipped)-1] || !anyHeadingMatches(headings, descriptionPart) || anyHeadingMatches(headings, unrelatedSection)) {
			return
		}
		size := s.UTF16Len(text)
		if size < 30 || contains(passages, text) || length+size > 4000 {
			return
		}
		passages = append(passages, text)
		length += size
	})
	if len(passages) == 0 {
		return nil
	}
	section := strings.Join(link.headings, " > ")
	quoted := make([]string, len(passages))
	for i, text := range passages {
		quoted[i] = link.title + ": " + text
	}
	note := "Retrieved through the public MediaWiki parse API from an observed combat-section link on " + link.parent + ". Parent section: " + section + ". Introduction, description and technique entries only, capped at 4000 characters; where the page splits a section by user, only this character's subsection; navigation and user lists omitted. Fan-maintained secondary source, not independently verified canon. Ownership and story-period limits remain those of the parent article."
	return &unit.Document{
		ID:   "character-technique:" + link.url.Hostname() + ":" + encodeURIComponent(link.title),
		Kind: "source",
		Text: "Observed link on " + link.character + "'s article: " + link.parent + "\nSection: " + section + "\nLink text: " + link.linkText + "\nParent passage: " + link.context +
			"\n\nOwnership and period limit: the parent lists this technique in the section above. Former, evolved or absorbed entries do not establish current availability. The shared technique description below does not transfer other users' powers to this character.\n\n" +
			strings.Join(quoted, "\n\n"),
		Origin: unit.Origin{Location: link.url.String(), Access: "retrieved", Note: &note},
	}
}
