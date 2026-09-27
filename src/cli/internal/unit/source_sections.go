package unit

import (
	"net/url"
	"regexp"
	"slices"
	"sort"
	"strings"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Source sections: the heading path of each passage, how directly a section
// describes the character's fighting, and the source techniques that the
// headings and passages name. A request that carries sourceTechniques was
// prepared with sectioned passages; a request without it keeps the earlier
// passages, so the citations of earlier drafts stay valid.

// SourceTechnique is a technique or form that the selected passages name,
// found by fixed rules: a technique or form section heading, the leading
// term of a list entry in a power section, a linked technique page, or a
// sentence that calls it the character's signature. PassageIDs are selected
// passages that name it in one of these ways or sit in its section, then
// passages that mention it, at most 12 in page order for each of its names.
type SourceTechnique struct {
	Name string `json:"name"`
	// Aliases are other names that one passage proves name the same
	// technique (sourceAliases); absent when it has none.
	Aliases    []string `json:"aliases,omitempty"`
	PassageIDs []string `json:"passageIds"`
	Signature  bool     `json:"signature"`
	// Salience is "strong" or "normal" (techniqueSalience); absent in a
	// request prepared before it existed.
	Salience string `json:"salience,omitempty"`
}

// SourceTechniqueSchema is one entry of a request's sourceTechniques.
var SourceTechniqueSchema = s.StrictObject(
	s.F("name", text()),
	s.F("aliases", s.Optional(s.Array(text()).Min(1))),
	s.F("passageIds", s.Array(text()).Min(1)),
	s.F("signature", s.Bool()),
	s.F("salience", s.Optional(s.Enum(SalienceStrong, SalienceNormal))),
)

var (
	// powerHeading marks a section about a specific power, form, technique
	// or weapon.
	powerHeading = regexp.MustCompile(`(?i)\b(?:devil fruits?|haki|gears?|forms?|transformations?|techniques?|fighting styles?|weapons?|moves?|attacks?|magic|spells?|jutsu|quirks?|sacred treasures?|awakening|modes?|graces?)\b`)
	// abilityHeading marks a general ability section.
	abilityHeading = regexp.MustCompile(`(?i)\b(?:abilit(?:y|ies)|powers?|skills?|combat|arts)\b`)
	// miscHeading marks overview and miscellaneous sections, which share
	// what the power and ability sections leave.
	miscHeading     = regexp.MustCompile(`(?i)\b(?:overview|miscellaneous|other|trivia|personality|appearances?|relationships?|history|gallery|luck|artistic|gluttony|cooking|reception|development|portrayals?|popularity|merchandise|voice|references|navigation|bounty|charisma)\b`)
	wikiHeadingLine = regexp.MustCompile(`^(={2,6})\s*(.*?)\s*(={2,6})$`)
)

// Section ranks.
const (
	PowerSection   = 0
	AbilitySection = 1
	OtherSection   = 2
)

// SectionRank orders a source section by how directly it describes the
// character's fighting: PowerSection for powers, forms, techniques,
// Devil Fruit, Haki and weapons; AbilitySection for other ability sections;
// OtherSection for the introduction, overview and miscellaneous sections.
// A miscellaneous heading anywhere in the path ranks the section as
// OtherSection unless its own heading names a power.
func SectionRank(path []string) int {
	if len(path) == 0 {
		return OtherSection
	}
	last := path[len(path)-1]
	for _, heading := range path {
		if miscHeading.MatchString(heading) && !powerHeading.MatchString(last) {
			return OtherSection
		}
	}
	for _, heading := range path {
		if powerHeading.MatchString(heading) {
			return PowerSection
		}
	}
	for _, heading := range path {
		if abilityHeading.MatchString(heading) {
			return AbilitySection
		}
	}
	return OtherSection
}

// sectionedPassages reports whether a request's passages carry sections.
func sectionedPassages(request *Request) bool { return request.SourceTechniques != nil }

// sourceSection is consecutive text under one heading path.
type sourceSection struct {
	path []string
	text string
}

// sourceSections splits a source document into its sections. Character-wiki
// documents write each passage as its heading path, one heading per line,
// followed by the passage; technique documents belong to their parent
// section and technique; other documents mark headings as wiki headings
// ("== Heading =="). Headings are never passage text.
func sourceSections(document Document) []sourceSection {
	var sections []sourceSection
	add := func(path []string, text string) {
		text = s.Trim(text)
		if text == "" {
			return
		}
		if n := len(sections); n > 0 && strings.Join(sections[n-1].path, "\n") == strings.Join(path, "\n") {
			sections[n-1].text += "\n" + text
			return
		}
		sections = append(sections, sourceSection{append([]string{}, path...), text})
	}
	switch {
	case strings.HasPrefix(document.ID, "character-wiki:"):
		for _, block := range strings.Split(document.Text, "\n\n") {
			lines := jsLines(block)
			if len(lines) == 0 {
				continue
			}
			var path []string
			for _, line := range lines[:len(lines)-1] {
				if heading := s.Trim(line); heading != "" {
					path = append(path, heading)
				}
			}
			add(path, lines[len(lines)-1])
		}
	case strings.HasPrefix(document.ID, "character-technique:"):
		var path []string
		if line, ok := firstLine(document.Text, "Section: "); ok {
			for _, heading := range strings.Split(strings.TrimPrefix(line, "Section: "), " > ") {
				if heading = s.Trim(heading); heading != "" {
					path = append(path, heading)
				}
			}
		}
		if title := techniqueTitle(document.ID); title != "" {
			path = append(path, title)
		}
		add(path, document.Text)
	default:
		type heading struct {
			level int
			text  string
		}
		var headings []heading
		path := func() []string {
			out := make([]string, len(headings))
			for i, h := range headings {
				out[i] = h.text
			}
			return out
		}
		var current []string
		flush := func() {
			add(path(), strings.Join(current, "\n"))
			current = nil
		}
		for _, line := range jsLines(document.Text) {
			if match := wikiHeadingLine.FindStringSubmatch(s.Trim(line)); match != nil && match[1] == match[3] && match[2] != "" {
				flush()
				for len(headings) > 0 && headings[len(headings)-1].level >= len(match[1]) {
					headings = headings[:len(headings)-1]
				}
				headings = append(headings, heading{len(match[1]), match[2]})
				continue
			}
			current = append(current, line)
		}
		flush()
	}
	return sections
}

// techniqueTitle is the page title of a technique document ID
// ("character-technique:HOST:TITLE", the title URI-encoded).
func techniqueTitle(id string) string {
	rest := strings.TrimPrefix(id, "character-technique:")
	_, encoded, ok := strings.Cut(rest, ":")
	if !ok {
		return ""
	}
	title, err := url.PathUnescape(encoded)
	if err != nil {
		return ""
	}
	return s.Trim(title)
}

// sectionPath splits a passage's section back into its headings.
func sectionPath(section string) []string {
	if section == "" {
		return nil
	}
	return strings.Split(section, " > ")
}

var (
	// leadingTerm is the name a list entry or definition starts with, as
	// in "Cruel Sun「…」: Escanor creates" (the passage may end inside the
	// brackets), "Gear 2 (…): …" or "Rhitta: A giant axe".
	leadingTerm = regexp.MustCompile(`^([A-Z0-9][\w'’.-]*(?:\s+(?:[A-Z0-9][\w'’.-]*|of|the|no|and|de|a)){0,5})\s*(?:「|\([^)]{0,120}\)\s*:\s|:\s)`)
	// signatureCue marks a sentence that names what the character is best
	// known for.
	signatureCue = regexp.MustCompile(`(?i)\b(?:signature|most iconic|iconic|trademark|best[- ]known|most famous)\b`)
	// signatureName is the capitalized name that follows a signature cue
	// within a few words, as in "his signature attack, the Gum-Gum Pistol".
	signatureName = regexp.MustCompile(`^(?:\s+[a-z][\w-]*){0,3},?\s+(?:(?:the|called|named|known as)\s+)?["“'‘]?((?:[A-Z][\w'’-]*)(?:\s+(?:[A-Z0-9][\w'’-]*|no|of|the))*)`)
	parentheses   = regexp.MustCompile(`\s*(?:\([^)]*\)|「[^」]*」)`)
	techniqueWord = regexp.MustCompile(`(?i)\s+techniques?$`)
)

const (
	maxSourceTechniques   = 40
	maxTechniquePassages  = 12
	maxTechniqueNameUTF16 = 60
)

// categoryWord is a heading word that names no technique on its own.
var categoryWord = regexp.MustCompile(`(?i)\b(?:and|of|the|physical|situational|former|other|main|known|list|equipment|usage|description|effects?)\b`)

// genericHeading reports whether a heading only names a category, such as
// "Devil Fruit", "Techniques" or "Physical Abilities".
func genericHeading(heading string) bool {
	rest := powerHeading.ReplaceAllString(heading, " ")
	rest = abilityHeading.ReplaceAllString(rest, " ")
	rest = categoryWord.ReplaceAllString(rest, " ")
	return len(words(parentheses.ReplaceAllString(rest, " "))) == 0
}

// namesTechnique reports text naming a technique by its name without
// parentheses, as whole words with the same capitals.
func namesTechnique(text, name string) bool {
	fields := strings.Fields(parentheses.ReplaceAllString(name, ""))
	if len(fields) == 0 || s.UTF16Len(strings.Join(fields, " ")) < 4 {
		return false
	}
	quoted := make([]string, len(fields))
	for i, field := range fields {
		quoted[i] = regexp.QuoteMeta(field)
	}
	return regexp.MustCompile(`(^|[^\pL\pN])` + strings.Join(quoted, `\s+`) + `($|[^\pL\pN])`).MatchString(text)
}

// techniqueKey identifies names that differ only in case, parentheses or a
// trailing "Techniques".
func techniqueKey(name string) string {
	name = parentheses.ReplaceAllString(name, "")
	if i := strings.LastIndex(name, "/"); i >= 0 {
		name = name[i+1:]
	}
	return strings.Join(words(techniqueWord.ReplaceAllString(s.Trim(name), "")), " ")
}

// sourceTechniqueName is the display name of a heading or technique page title.
func sourceTechniqueName(name string) string {
	if i := strings.LastIndex(name, "/"); i >= 0 {
		name = name[i+1:]
	}
	if short := s.Trim(techniqueWord.ReplaceAllString(name, "")); short != "" && !genericHeading(short) {
		name = short
	}
	return s.Trim(name)
}

// SourceTechniques lists the techniques and forms the selected passages of a
// sectioned request name, in the order the passages first name them. A
// technique is named by a power section heading that is not only a
// category, the leading term of an entry in a power or ability section
// ("Cruel Sun「…」: …"), a technique page's title, or the capitalized name
// that follows a signature cue ("his signature attack, the Gum-Gum
// Pistol"). Its passages are those that name it this way or sit in its
// section, then those that mention its name, at most 12. A technique is a
// signature one when a sentence with a signature cue names it.
//
// Two names are one technique when one passage proves it (sourceAliases):
// the merged technique keeps the first name, lists the other in Aliases
// and keeps every passage ID of both. Each technique then carries its
// Salience (techniqueSalience). The list is empty for a request whose
// passages carry no sections.
func SourceTechniques(request *Request) []SourceTechnique {
	return sourceTechniques(request, false)
}

// sourceTechniques derives the list, or with legacy the form requests were
// prepared with before aliases and salience (Default v30 and earlier):
// every name its own technique, without Aliases or Salience. Verifying such
// a saved request compares it with the legacy form, so its hash and the
// citations of its draft stay valid.
func sourceTechniques(request *Request, legacy bool) []SourceTechnique {
	out := []SourceTechnique{}
	if !sectionedPassages(request) {
		return out
	}
	spans := AuthorEvidence(request)
	character := techniqueKey(request.Character.Name)
	type entry struct {
		technique SourceTechnique
		first     int
		defining  map[int]bool
		mentions  map[int]bool
		page      bool
	}
	var entries []*entry
	index := map[string]*entry{}
	add := func(name string, span int, signature bool) *entry {
		name = sourceTechniqueName(name)
		key := techniqueKey(name)
		if key == "" || key == character || s.UTF16Len(name) > maxTechniqueNameUTF16 {
			return nil
		}
		e := index[key]
		if e == nil {
			e = &entry{technique: SourceTechnique{Name: name}, first: span, defining: map[int]bool{}, mentions: map[int]bool{}}
			index[key] = e
			entries = append(entries, e)
		}
		e.defining[span] = true
		e.technique.Signature = e.technique.Signature || signature
		return e
	}
	body := make([]string, len(spans))
	for i, span := range spans {
		path := sectionPath(span.Section)
		rank := SectionRank(path)
		text := span.Text
		technique := strings.HasPrefix(span.DocumentID, "character-technique:")
		if technique {
			// A technique page's passages start with its title; its header
			// lines name the parent article, not a technique.
			title := techniqueTitle(span.DocumentID)
			if title == "" || !strings.HasPrefix(text, title+": ") {
				continue
			}
			text = strings.TrimPrefix(text, title+": ")
			if e := add(title, i, false); e != nil {
				e.page = true
			}
		}
		body[i] = text
		if rank == PowerSection {
			for depth, heading := range path {
				if technique && depth == len(path)-1 {
					break
				}
				if SectionRank(path[:depth+1]) == PowerSection && !genericHeading(heading) {
					add(heading, i, false)
				}
			}
		}
		if rank != OtherSection || technique {
			if match := leadingTerm.FindStringSubmatch(text); match != nil && !genericHeading(match[1]) {
				add(match[1], i, false)
			}
		}
		if cue := signatureCue.FindStringIndex(text); cue != nil {
			if match := signatureName.FindStringSubmatch(text[cue[1]:]); match != nil {
				add(match[1], i, true)
			}
		}
	}
	for i := range spans {
		if body[i] == "" {
			continue
		}
		signature := signatureCue.MatchString(body[i])
		for _, e := range entries {
			if namesTechnique(body[i], e.technique.Name) {
				e.mentions[i] = true
				e.technique.Signature = e.technique.Signature || signature
			}
		}
	}
	sort.SliceStable(entries, func(a, b int) bool { return entries[a].first < entries[b].first })
	passages := func(e *entry) []int {
		var chosen []int
		for _, group := range []map[int]bool{e.defining, e.mentions} {
			var members []int
			for i := range group {
				if !slices.Contains(chosen, i) {
					members = append(members, i)
				}
			}
			sort.Ints(members)
			for _, i := range members {
				if len(chosen) < maxTechniquePassages {
					chosen = append(chosen, i)
				}
			}
		}
		sort.Ints(chosen)
		return chosen
	}
	ids := func(chosen []int) []string {
		out := make([]string, len(chosen))
		for n, i := range chosen {
			out[n] = spans[i].ID
		}
		return out
	}
	if legacy {
		for _, e := range entries {
			if len(out) == maxSourceTechniques {
				break
			}
			e.technique.PassageIDs = ids(passages(e))
			out = append(out, e.technique)
		}
		return out
	}
	// Merge proven aliases into the technique named first; each keeps the
	// passages it had on its own.
	candidates := make([]*sourceCandidate, len(entries))
	position := map[*entry]int{}
	for n, e := range entries {
		position[e] = n
		candidates[n] = &sourceCandidate{technique: e.technique, key: techniqueKey(e.technique.Name), passages: passages(e), page: e.page}
	}
	lookup := func(name string) int {
		if e := index[techniqueKey(sourceTechniqueName(name))]; e != nil {
			return position[e]
		}
		return -1
	}
	merged := mergeSourceAliases(candidates, sourceAliases(request, spans, body, lookup))
	shares := powerSectionShares(request)
	for _, c := range merged {
		if len(out) == maxSourceTechniques {
			break
		}
		c.technique.PassageIDs = ids(c.passages)
		c.technique.Salience = techniqueSalience(c, shares)
		out = append(out, c.technique)
	}
	return out
}

// WithSourceTechniques marks a request's passages as sectioned and records
// the source techniques they name.
func WithSourceTechniques(request Request) Request {
	request.SourceTechniques = &[]SourceTechnique{}
	techniques := SourceTechniques(&request)
	request.SourceTechniques = &techniques
	return request
}

// verifiedSourceTechniques is the list a prepared request's
// sourceTechniques must equal: derived again from its documents, in the
// legacy form when every stored technique lacks a salience, as in requests
// prepared under Default v30 and earlier.
func verifiedSourceTechniques(request Request) []SourceTechnique {
	stored := *request.SourceTechniques
	legacy := len(stored) > 0 && !slices.ContainsFunc(stored, func(t SourceTechnique) bool { return t.Salience != "" })
	request.SourceTechniques = &[]SourceTechnique{}
	return sourceTechniques(&request, legacy)
}

// sectionKey identifies a passage's section within its document.
func sectionKey(span EvidenceSpan) string { return span.DocumentID + "\x00" + span.Section }
