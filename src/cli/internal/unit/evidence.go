package unit

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/clipperhouse/uax29/v2/sentences"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// EvidenceSpan is a quotable source passage with a stable ID. Section is
// its heading path in the source ("Devil Fruit > Gear 4"), present only for
// a request prepared with sectioned passages; the introduction has none.
type EvidenceSpan struct {
	ID         string `json:"id"`
	DocumentID string `json:"documentId"`
	Section    string `json:"section,omitempty"`
	Text       string `json:"text"`
}

// TechniqueContext is the period context of a retrieved technique article.
type TechniqueContext struct {
	DocumentID string `json:"documentId"`
	Quote      string `json:"quote"`
	Technique  string `json:"technique"`
}

var formerWord = regexp.MustCompile(`(?i)\bformer\b`)

// jsLines splits text at JavaScript line terminators, like a multiline regex.
func jsLines(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return r == '\n' || r == '\r' || r == 0x2028 || r == 0x2029
	})
}

// firstLine returns the first line starting with prefix and having content after it.
func firstLine(text, prefix string) (string, bool) {
	for _, line := range jsLines(text) {
		if strings.HasPrefix(line, prefix) && len(line) > len(prefix) {
			return line, true
		}
	}
	return "", false
}

// HistoricalTechniqueContext keeps the explicit period of an observed technique link.
func HistoricalTechniqueContext(request *Request, documentID string) *TechniqueContext {
	var document *Document
	for i := range request.Documents {
		if request.Documents[i].ID == documentID {
			document = &request.Documents[i]
			break
		}
	}
	if document == nil || !strings.HasPrefix(document.ID, "character-technique:") {
		return nil
	}
	section, ok := firstLine(document.Text, "Section: ")
	if !ok || !formerWord.MatchString(section) {
		return nil
	}
	technique := "Selected technique"
	if line, ok := firstLine(document.Text, "Link text: "); ok {
		technique = strings.TrimPrefix(line, "Link text: ")
	}
	return &TechniqueContext{DocumentID: documentID, Quote: section, Technique: technique}
}

// u16 is a JavaScript string: UTF-16 code units.
type u16 []uint16

func toU16(text string) u16  { return utf16.Encode([]rune(text)) }
func (t u16) String() string { return string(utf16.Decode(t)) }
func (t u16) trim() u16      { return toU16(s.Trim(t.String())) }
func (t u16) lastSpace(from int) int {
	if from >= len(t) {
		from = len(t) - 1
	}
	for i := from; i >= 0; i-- {
		if t[i] == ' ' {
			return i
		}
	}
	return -1
}

// splitPassages splits text into sentences, joining a sentence shorter
// than 15 characters to the next, and splits passages longer than 450
// characters at a space.
func splitPassages(text string) []string {
	var passages []string
	segments := sentences.FromString(text)
	for segments.Next() {
		segment := segments.Value()
		if n := len(passages); n > 0 && s.UTF16Len(s.Trim(passages[n-1])) < 15 {
			passages[n-1] += segment
		} else {
			passages = append(passages, segment)
		}
	}
	if n := len(passages); n > 1 && s.UTF16Len(s.Trim(passages[n-1])) < 15 {
		tail := passages[n-1]
		passages = passages[:n-1]
		passages[n-2] += tail
	}
	var out []string
	for _, passage := range passages {
		remaining := toU16(passage).trim()
		for len(remaining) > 0 {
			end := len(remaining)
			if end > 450 {
				boundary := remaining.lastSpace(450)
				if boundary >= 15 {
					end = boundary
				} else {
					end = 450
				}
				if s.UTF16Len(s.Trim(remaining[end:].String())) < 15 {
					end = max(15, end-15)
				}
			}
			if text := remaining[:end].trim(); len(text) > 0 {
				out = append(out, text.String())
			}
			remaining = remaining[end:].trim()
		}
	}
	return out
}

// EvidenceSpans splits source documents into deterministic, quotable
// passages. A request prepared with sectioned passages splits each section
// on its own and labels its passages with the section; headings are never
// passages. Other requests split each document as a whole, as when they
// were prepared.
func EvidenceSpans(request *Request) []EvidenceSpan {
	var spans []EvidenceSpan
	sectioned := sectionedPassages(request)
	for docIndex, document := range request.Documents {
		if document.Kind != "source" {
			continue
		}
		sections := []sourceSection{{text: document.Text}}
		if sectioned {
			sections = sourceSections(document)
		}
		index := 0
		for _, section := range sections {
			label := ""
			if sectioned {
				label = strings.Join(section.path, " > ")
			}
			for _, text := range splitPassages(section.text) {
				spans = append(spans, EvidenceSpan{
					ID:         "source" + strconv.Itoa(docIndex+1) + ":" + strconv.Itoa(index),
					DocumentID: document.ID,
					Section:    label,
					Text:       text,
				})
				index++
			}
		}
	}
	return spans
}

var (
	combatWords = regexp.MustCompile(`(?i)\b(?:abilit\w*|power\w*|attack\w*|combat|strength|speed|technique\w*|transform\w*|form\w*|damage|control|weapon\w*|punch\w*|beam\w*|stretc\w*|absor\w*|mimic\w*|summon\w*|limit\w*|weak\w*|cannot|unable|immune|immunity)\b`)
	// effectWords mark what an attack does to its targets, which a Tower
	// Defense adaptation maps to area, status, rate and capacity.
	effectWords   = regexp.MustCompile(`(?i)\b(?:flames?|fire|ignit\w*|burn\w*|explo\w*|freez\w*|ice|poison\w*|electri\w*|lightning|shock\w*|paraly\w*|stun\w*|slow\w*|knock\w*|blast\w*|area|gigantic|inflat\w*|multiple|barrage|rapid\w*|flurry|pierc\w*|bounc\w*)\b`)
	signatureWord = regexp.MustCompile(`(?i)\b(?:signature|primary|characteristic)\b`)
	actionWords   = regexp.MustCompile(`(?i)\b(?:fires?|firing|shoots?|shooting|launch\w*|emits?|emitting|strikes?|striking|throws?|throwing|projectiles?|slashes?|slashing)\b`)
	limitWords    = regexp.MustCompile(`(?i)\b(?:cannot|unable|requires?|only|former|limitations?|weakness\w*)\b`)
)

// AuthorEvidence is a bounded lexical selection of spans for the model.
func AuthorEvidence(request *Request) []EvidenceSpan {
	all := EvidenceSpans(request)
	const limit, maxSpans = 32000, 200
	total := 0
	for _, span := range all {
		total += s.UTF16Len(span.Text)
	}
	if len(all) <= maxSpans && total <= limit {
		return all
	}
	if sectionedPassages(request) {
		return sectionedEvidence(request, all, limit, maxSpans)
	}
	type ranked struct {
		span  EvidenceSpan
		index int
		score int
	}
	scores := evidenceScores(all)
	list := make([]ranked, len(all))
	for i, span := range all {
		list[i] = ranked{span, i, scores[i]}
	}
	sort.SliceStable(list, func(a, b int) bool {
		if list[a].score != list[b].score {
			return list[a].score > list[b].score
		}
		return list[a].index < list[b].index
	})
	used, selected := 0, 0
	reserved := map[string]bool{}
	for _, document := range request.Documents {
		if !strings.HasPrefix(document.ID, "character-technique:") {
			continue
		}
		var packet []EvidenceSpan
		size := 0
		for _, span := range all {
			if span.DocumentID == document.ID {
				packet = append(packet, span)
				size += s.UTF16Len(span.Text)
			}
		}
		if len(packet) > 32 || size > 6000 || selected+len(packet) > maxSpans || used+size > limit {
			continue
		}
		for _, span := range packet {
			reserved[span.ID] = true
		}
		used += size
		selected += len(packet)
	}
	var kept []ranked
	for _, item := range list {
		if reserved[item.span.ID] {
			kept = append(kept, item)
			continue
		}
		length := s.UTF16Len(item.span.Text)
		if selected >= maxSpans || used+length > limit {
			continue
		}
		used += length
		selected++
		kept = append(kept, item)
	}
	sort.SliceStable(kept, func(a, b int) bool { return kept[a].index < kept[b].index })
	out := make([]EvidenceSpan, len(kept))
	for i, item := range kept {
		out[i] = item.span
	}
	return out
}

// evidenceScores rates each passage by its combat, action, effect,
// signature and limit words. The first passage of each document scores 100
// more, and a passage repeating an earlier one in its document scores -1.
func evidenceScores(all []EvidenceSpan) []int {
	first := map[string]bool{}
	seen := map[string]bool{}
	scores := make([]int, len(all))
	for i, span := range all {
		identity := !first[span.DocumentID]
		first[span.DocumentID] = true
		key := span.DocumentID + "\x00" + span.Text
		repeated := seen[key]
		seen[key] = true
		combat := len(combatWords.FindAllStringIndex(span.Text, -1))
		relevance := 0
		if signatureWord.MatchString(span.Text) {
			relevance += 3
		}
		relevance += min(len(actionWords.FindAllStringIndex(span.Text, -1)), 3)
		relevance += min(len(effectWords.FindAllStringIndex(span.Text, -1)), 3)
		if limitWords.MatchString(span.Text) {
			relevance += 4
		}
		score := -1
		if !repeated {
			score = min(combat, 12) + relevance
			if identity {
				score += 100
			}
		}
		scores[i] = score
	}
	return scores
}

// Budgets of a sectioned selection, within its passage and text limits.
const (
	// techniqueEvidenceBudget is the text reserved for technique pages, and
	// techniquePageText the most one page reserves.
	techniqueEvidenceBudget = 8_000
	techniquePageText       = 6_000
	// sectionQuotaSpans and sectionQuotaText bound what each power or
	// ability section reserves; sectionQuotaBudget bounds all of them.
	sectionQuotaSpans  = 4
	sectionQuotaText   = 1_500
	sectionQuotaBudget = 16_000
)

// sectionedEvidence selects the passages of a sectioned request within the
// same limits as before. It keeps, in order:
//
//   - the first passage of each document;
//   - for each technique page, its header, its first passage and its best
//     passages, the pages sharing techniqueEvidenceBudget equally;
//   - each document's introduction, in page order, and then the best
//     passages of each power section and each ability section, each within
//     the section quota;
//   - the remaining passages by score, sections taking turns among passages
//     of equal score, power sections first.
func sectionedEvidence(request *Request, all []EvidenceSpan, limit, maxSpans int) []EvidenceSpan {
	scores := evidenceScores(all)
	rank := make([]int, len(all))
	for i, span := range all {
		rank[i] = SectionRank(sectionPath(span.Section))
	}
	kept := make([]bool, len(all))
	used, selected := 0, 0
	take := func(i int) bool {
		size := s.UTF16Len(all[i].Text)
		if kept[i] || selected >= maxSpans || used+size > limit {
			return false
		}
		kept[i] = true
		used += size
		selected++
		return true
	}
	for i := range all {
		if scores[i] >= 100 {
			take(i)
		}
	}
	// Technique pages share techniqueEvidenceBudget equally. Each keeps its
	// header, which states whose page it is and what the character may
	// claim from it, and its first passage, then its best passages by
	// score.
	var techniques []string
	for _, document := range request.Documents {
		if strings.HasPrefix(document.ID, "character-technique:") && document.Kind == "source" {
			techniques = append(techniques, document.ID)
		}
	}
	for _, id := range techniques {
		share, size := min(techniqueEvidenceBudget/len(techniques), techniquePageText), 0
		prefix := techniqueTitle(id) + ": "
		var header, rest []int
		body := false
		for i, span := range all {
			if span.DocumentID != id {
				continue
			}
			if !body {
				header = append(header, i)
				body = strings.HasPrefix(span.Text, prefix)
				continue
			}
			rest = append(rest, i)
		}
		sort.SliceStable(rest, func(a, b int) bool { return scores[rest[a]] > scores[rest[b]] })
		for _, i := range append(header, rest...) {
			length := s.UTF16Len(all[i].Text)
			if kept[i] {
				size += length
				continue
			}
			if scores[i] < 0 || size+length > share {
				continue
			}
			if take(i) {
				size += length
			}
		}
	}
	// byScore orders passages by score, then page order.
	byScore := func(indices []int) {
		sort.SliceStable(indices, func(a, b int) bool {
			if scores[indices[a]] != scores[indices[b]] {
				return scores[indices[a]] > scores[indices[b]]
			}
			return indices[a] < indices[b]
		})
	}
	// Sections in page order, then power sections first.
	sections := map[string][]int{}
	var keys []string
	for i, span := range all {
		key := sectionKey(span)
		if _, ok := sections[key]; !ok {
			keys = append(keys, key)
		}
		sections[key] = append(sections[key], i)
	}
	sort.SliceStable(keys, func(a, b int) bool { return rank[sections[keys[a]][0]] < rank[sections[keys[b]][0]] })
	// Each document's introduction, which summarizes the character, takes
	// its quota before the power sections and the ability sections.
	quotaKeys := append([]string{}, keys...)
	sort.SliceStable(quotaKeys, func(a, b int) bool {
		return all[sections[quotaKeys[a]][0]].Section == "" && all[sections[quotaKeys[b]][0]].Section != ""
	})
	quota := 0
	for _, key := range quotaKeys {
		members := append([]int{}, sections[key]...)
		first := members[0]
		introduction := all[first].Section == ""
		if (rank[first] == OtherSection && !introduction) || strings.HasPrefix(all[first].DocumentID, "character-technique:") {
			continue
		}
		// An introduction keeps its opening passages; a section its best.
		most := sectionQuotaSpans
		if introduction {
			most = len(members)
		} else {
			byScore(members)
		}
		count, size := 0, 0
		for _, i := range members {
			length := s.UTF16Len(all[i].Text)
			if kept[i] {
				count++
				size += length
				continue
			}
			if introduction && (scores[i] < 0 || size+length > sectionQuotaText) {
				break
			}
			if scores[i] < 0 || count >= most || size+length > sectionQuotaText || quota+length > sectionQuotaBudget {
				continue
			}
			if take(i) {
				count++
				size += length
				quota += length
			}
		}
	}
	// The remainder goes by score. Among passages of equal score, each
	// section offers its next passage in turn, power sections first, so no
	// single long section fills the remainder in page order.
	turn := make([]int, len(all))
	position := map[string]int{}
	for p, key := range keys {
		position[key] = p
		members := append([]int{}, sections[key]...)
		byScore(members)
		n := 0
		for _, i := range members {
			if !kept[i] {
				turn[i] = n
				n++
			}
		}
	}
	var rest []int
	for i := range all {
		if !kept[i] {
			rest = append(rest, i)
		}
	}
	sort.SliceStable(rest, func(a, b int) bool {
		i, j := rest[a], rest[b]
		switch {
		case scores[i] != scores[j]:
			return scores[i] > scores[j]
		case turn[i] != turn[j]:
			return turn[i] < turn[j]
		default:
			return position[sectionKey(all[i])] < position[sectionKey(all[j])]
		}
	})
	for _, i := range rest {
		take(i)
	}
	var out []EvidenceSpan
	for i, span := range all {
		if kept[i] {
			out = append(out, span)
		}
	}
	return out
}
