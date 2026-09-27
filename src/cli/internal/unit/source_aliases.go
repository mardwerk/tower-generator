package unit

import (
	"regexp"
	"slices"
	"sort"
	"strings"
)

// Source aliases (#61, SOL-61-05): two source technique names are one
// technique only when one passage proves it. Containment alone proves
// nothing: "Gear 2 Buso" is not "Gear 2", and "Armament Haki" is not
// "Haki". A passage proves an alias in one of two ways:
//
//   - It opens the page or section of one name and defines its subject by
//     the other: the first passage of the technique page "Rhitta" reads
//     "The Divine Axe Rhitta「…」 is a Sacred Treasure", so "The Divine Axe
//     Rhitta" is that page's subject, Rhitta. The same holds for the first
//     passage under a heading that names a technique.
//   - It says so in one sentence: "X, also called Y", "X (also known as
//     Y)", "X, otherwise known as Y".
//
// Both names must already be source techniques; an alias adds no name.

var (
	// techniqueNamePattern is a capitalized name of up to six words, as a
	// leading term is.
	techniqueNamePattern = `[A-Z0-9][\w'’.-]*(?:\s+(?:[A-Z0-9][\w'’.-]*|of|the|no|and|de|a)){0,5}`
	// subjectDefinition is a passage that opens by defining its subject:
	// "The Divine Axe Rhitta「…」 is a Sacred Treasure".
	subjectDefinition = regexp.MustCompile(`^(` + techniqueNamePattern + `)\s*(?:「[^」]*」|\([^)]{0,120}\))?\s*(?:is|was|are|were)\s+(?:a|an|the|one)\s`)
	// alsoCalled is one sentence naming a technique by a second name.
	alsoCalled = regexp.MustCompile(`(` + techniqueNamePattern + `)(?:\s*「[^」]*」)?\s*[,(]\s*(?:also|otherwise|commonly)\s+(?:called|named|known\s+as|referred\s+to\s+as)\s+(?:the\s+)?["“'‘]?(` + techniqueNamePattern + `)`)
)

// sourceCandidate is a source technique while aliases are merged: its
// passages as span indexes and whether a technique page is titled for it.
type sourceCandidate struct {
	technique SourceTechnique
	key       string
	keys      []string
	passages  []int
	page      bool
}

// aliasPair is two candidates, by position, that one passage names as one
// technique, and that passage's position (proof).
type aliasPair struct{ a, b, proof int }

// sourceAliases finds the alias pairs the selected passages prove. body is
// each passage's text after a technique page's title, empty for a page's
// header, and lookup the position of the candidate a name names, or -1.
func sourceAliases(request *Request, spans []EvidenceSpan, body []string, lookup func(string) int) []aliasPair {
	first := openingPassages(request)
	var pairs []aliasPair
	for i, span := range spans {
		if body[i] == "" {
			continue
		}
		if first[span.ID] {
			owner := -1
			if strings.HasPrefix(span.DocumentID, "character-technique:") {
				owner = lookup(techniqueTitle(span.DocumentID))
			} else if path := sectionPath(span.Section); len(path) > 0 {
				owner = lookup(path[len(path)-1])
			}
			if match := subjectDefinition.FindStringSubmatch(body[i]); owner >= 0 && match != nil {
				if subject := lookupWords(match[1], lookup, true); subject >= 0 && subject != owner {
					pairs = append(pairs, aliasPair{owner, subject, i})
				}
			}
		}
		for _, match := range alsoCalled.FindAllStringSubmatch(body[i], -1) {
			a, b := lookupWords(match[1], lookup, false), lookupWords(match[2], lookup, true)
			if a >= 0 && b >= 0 && a != b {
				pairs = append(pairs, aliasPair{a, b, i})
			}
		}
	}
	return pairs
}

// lookupWords finds the candidate a captured name names. A capture can
// take in capitalized words around the name, as "In Rhitta" or "Rhitta
// Charge": with prefix it tries the longest leading words first, otherwise
// the longest trailing words.
func lookupWords(name string, lookup func(string) int, prefix bool) int {
	fields := strings.Fields(name)
	for n := len(fields); n > 0; n-- {
		part := fields[len(fields)-n:]
		if prefix {
			part = fields[:n]
		}
		if found := lookup(strings.Join(part, " ")); found >= 0 {
			return found
		}
	}
	return -1
}

// openingPassages marks the first passage of each technique page's text and
// of each section, among all passages of the request.
func openingPassages(request *Request) map[string]bool {
	out := map[string]bool{}
	seen := map[string]bool{}
	for _, span := range EvidenceSpans(request) {
		key := sectionKey(span)
		if strings.HasPrefix(span.DocumentID, "character-technique:") {
			if !strings.HasPrefix(span.Text, techniqueTitle(span.DocumentID)+": ") {
				continue
			}
			key = span.DocumentID
		}
		if !seen[key] {
			seen[key] = true
			out[span.ID] = true
		}
	}
	return out
}

// mergeSourceAliases joins the candidates each pair names. A group keeps
// the name its passages name first and lists the others as aliases, keeps
// every passage each had, and is a signature one or titled by a page when
// any of them is.
func mergeSourceAliases(candidates []*sourceCandidate, pairs []aliasPair) []*sourceCandidate {
	parent := make([]int, len(candidates))
	for i := range parent {
		parent[i] = i
	}
	var root func(int) int
	root = func(i int) int {
		for parent[i] != i {
			i = parent[i]
		}
		return i
	}
	for _, pair := range pairs {
		a, b := root(pair.a), root(pair.b)
		if a == b {
			continue
		}
		// Candidates are in the order the passages first name them, so
		// the lower position is named first.
		parent[max(a, b)] = min(a, b)
	}
	// Each merged group keeps the passage that proves each of its aliases,
	// even beyond the passages it keeps per name, so the plan and the
	// review can see why the names were joined.
	proofs := map[int][]int{}
	for _, pair := range pairs {
		if r := root(pair.a); r == root(pair.b) && !slices.Contains(proofs[r], pair.proof) {
			proofs[r] = append(proofs[r], pair.proof)
		}
	}
	var out []*sourceCandidate
	for i, c := range candidates {
		c.keys = []string{c.key}
		r := root(i)
		if r == i {
			out = append(out, c)
			continue
		}
		group := candidates[r]
		group.technique.Aliases = append(group.technique.Aliases, c.technique.Name)
		group.technique.Aliases = append(group.technique.Aliases, c.technique.Aliases...)
		group.keys = append(group.keys, c.key)
		group.technique.Signature = group.technique.Signature || c.technique.Signature
		group.page = group.page || c.page
		for _, p := range c.passages {
			if !slices.Contains(group.passages, p) {
				group.passages = append(group.passages, p)
			}
		}
		sort.Ints(group.passages)
	}
	for r, proof := range proofs {
		group := candidates[r]
		for _, p := range proof {
			if !slices.Contains(group.passages, p) {
				group.passages = append(group.passages, p)
			}
		}
		sort.Ints(group.passages)
	}
	return out
}

// Names lists a source technique's name and then its aliases.
func (t SourceTechnique) Names() []string {
	return append([]string{t.Name}, t.Aliases...)
}
