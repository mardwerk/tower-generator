package unit

import (
	"slices"
	"strings"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Technique salience (#61, SOL-61-05): a deterministic signal of how much
// source evidence a source technique has. It is evidence for a major floor
// only: page availability and cross-link structure are no rule for what is
// canonical, so a strong technique is never core for that alone.

// Salience values of a source technique.
const (
	SalienceStrong = "strong"
	SalienceNormal = "normal"
)

// StrongSectionShare is the share of the power sections' text at or above
// which a technique's own sections make it strong. Luffy's Gear and Haki
// sections each hold 5 to 14 percent of his power sections, and his single
// named attacks below 2 percent.
const StrongSectionShare = 0.10

// techniqueSalience is strong when research followed a technique or form
// page titled for the technique, when it is a signature technique, or when
// its own sections hold at least StrongSectionShare of the power sections'
// text (powerSectionShares); otherwise normal.
func techniqueSalience(c *sourceCandidate, shares sectionShares) string {
	if c.page || c.technique.Signature || shares.share(c.keys) >= StrongSectionShare {
		return SalienceStrong
	}
	return SalienceNormal
}

// sectionShares holds, for each technique key, the passages of its own
// sections, and the text of every power section passage.
type sectionShares struct {
	length  []int
	owned   map[string][]int
	allText int
}

// share is the part of the power sections' text that the sections of any
// of keys hold, each passage counted once.
func (shares sectionShares) share(keys []string) float64 {
	if shares.allText == 0 {
		return 0
	}
	var passages []int
	for _, key := range keys {
		for _, p := range shares.owned[key] {
			if !slices.Contains(passages, p) {
				passages = append(passages, p)
			}
		}
	}
	total := 0
	for _, p := range passages {
		total += shares.length[p]
	}
	return float64(total) / float64(shares.allText)
}

// powerSectionShares measures the power and ability sections of a
// request's character articles, the documents that are not technique pages
// (a technique page is its own signal), in full, selected or not. A
// technique's own sections are the text under a heading that names it,
// subsections included, and the list entries that start with its name, each
// entry with all its text. Lengths are in UTF-16 units.
func powerSectionShares(request *Request) sectionShares {
	shares := sectionShares{owned: map[string][]int{}}
	for _, document := range request.Documents {
		if document.Kind != "source" || strings.HasPrefix(document.ID, "character-technique:") {
			continue
		}
		for _, part := range sectionEntries(document) {
			if SectionRank(part.path) == OtherSection {
				continue
			}
			i := len(shares.length)
			shares.length = append(shares.length, s.UTF16Len(part.text))
			shares.allText += shares.length[i]
			var keys []string
			for depth, heading := range part.path {
				if SectionRank(part.path[:depth+1]) == PowerSection && !genericHeading(heading) {
					keys = append(keys, techniqueKey(sourceTechniqueName(heading)))
				}
			}
			if match := leadingTerm.FindStringSubmatch(part.text); match != nil && !genericHeading(match[1]) {
				keys = append(keys, techniqueKey(sourceTechniqueName(match[1])))
			}
			for _, key := range keys {
				if key != "" && !slices.Contains(shares.owned[key], i) {
					shares.owned[key] = append(shares.owned[key], i)
				}
			}
		}
	}
	return shares
}

// sectionEntries splits a character article into the text of each entry
// with its heading path: a character-wiki document's blocks, one entry
// each, or another document's sections.
func sectionEntries(document Document) []sourceSection {
	if !strings.HasPrefix(document.ID, "character-wiki:") {
		return sourceSections(document)
	}
	var out []sourceSection
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
		if text := s.Trim(lines[len(lines)-1]); text != "" {
			out = append(out, sourceSection{path, text})
		}
	}
	return out
}
