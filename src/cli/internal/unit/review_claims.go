package unit

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Prose cues can identify claims worth checking, but cannot prove their
// meaning. Negation and qualifications such as "per Gold" can reverse the
// apparent assertion. These checks only add human-review advisories; hard
// fact correction is restricted to structured citations in review_facts.go.

// reviewClaims holds what the two checks read: the retained plan, the
// purchase names and the side-purchase comparisons against each capstone.
type reviewClaims struct {
	plan  *DesignPlan
	names [3][5]string
	rows  []comparisonRow
}

// comparisonRow is one metric of one againstCapstone comparison: the side
// purchase that takes a build from one legal build to another, its gain and
// price, the capstone's gain and price, and sideGainAtLeastCapstone.
type comparisonRow struct {
	from, to                 string
	side, capstone           string
	sidePrice, capstonePrice any
	metric                   string
	sideGain, capstoneGain   float64
	atLeast                  bool
}

func newReviewClaims(checked Checked) reviewClaims {
	claims := reviewClaims{plan: checked.Draft.Run.DesignPlan}
	if blueprint := checked.Draft.Candidate.Blueprint; blueprint != nil {
		for path := range m.PathKeys {
			for tier := 1; tier <= len(m.TierKeys); tier++ {
				claims.names[path][tier-1] = blueprint.Paths.At(path).Tiers.At(tier).Name
			}
		}
	}
	currency := ""
	if d := checked.Draft.Prepared.Request.MechanicsDefinition; d != nil {
		currency = d.Profile.Currency
	}
	evidence := ReviewPurchaseEvidence(checked.Draft.Run.DesignEvaluation, currency)
	if evidence == nil {
		return claims
	}
	paths, _ := evidence.Get("paths")
	for _, value := range asList(paths) {
		path, ok := value.(*s.Object)
		if !ok {
			continue
		}
		crosspaths, _ := path.Get("crosspaths")
		for _, value := range asList(crosspaths) {
			crosspath, ok := value.(*s.Object)
			if !ok {
				continue
			}
			against, ok := objectAt(crosspath, AgainstCapstone)
			if !ok {
				continue
			}
			side, _ := objectAt(against, "sidePurchase")
			capstone, _ := objectAt(against, "capstone")
			if side == nil || capstone == nil {
				continue
			}
			from, _ := crosspath.Get("from")
			to, _ := crosspath.Get("to")
			sideCode, _ := side.Get("code")
			capstoneCode, _ := capstone.Get("code")
			sidePrice, _ := side.Get("price")
			capstonePrice, _ := capstone.Get("price")
			for _, metric := range timeAveragedMetrics {
				gains, _ := objectAt(against, metric.name)
				if gains == nil {
					continue
				}
				sideValue, _ := gains.Get("sidePurchase")
				capstoneValue, _ := gains.Get("capstone")
				atLeast, _ := gains.Get(SideGainAtLeastCapstone)
				x, okSide := evidenceNumber(sideValue)
				y, okCapstone := evidenceNumber(capstoneValue)
				checkedValue, okAtLeast := atLeast.(bool)
				code, _ := sideCode.(string)
				capstoneName, _ := capstoneCode.(string)
				if !okSide || !okCapstone || !okAtLeast || code == "" || capstoneName == "" {
					continue
				}
				claims.rows = append(claims.rows, comparisonRow{
					from: selectionText(from), to: selectionText(to), side: code, capstone: capstoneName,
					sidePrice: sidePrice, capstonePrice: capstonePrice, metric: metric.name,
					sideGain: x, capstoneGain: y, atLeast: checkedValue,
				})
			}
		}
	}
	return claims
}

// selectionText writes a retained build such as [0, 5, 0] as 0-5-0.
func selectionText(value any) string {
	var parts []string
	for _, tier := range asList(value) {
		number, _ := tier.(float64)
		parts = append(parts, s.FormatNumber(number))
	}
	return strings.Join(parts, "-")
}

// purchasePosition reads a purchase code such as x-x-1 as its path index
// and tier.
func purchasePosition(code string) (path, tier int, ok bool) {
	if !purchaseCode(code) {
		return 0, 0, false
	}
	for index, part := range strings.Split(code, "-") {
		if part != "x" {
			return index, int(part[0] - '0'), true
		}
	}
	return 0, 0, false
}

// sentenceEnd ends a sentence: a full stop, question or exclamation mark
// before a space or the end, so 4.66 stays one number, or a line break.
var sentenceEnd = regexp.MustCompile(`[.!?]+(?:\s+|$)|\n+`)

// textSentence is one sentence of a finding's text and where it starts.
type textSentence struct {
	text  string
	start int
}

func findingSentences(text string) []textSentence {
	var out []textSentence
	last := 0
	for _, at := range sentenceEnd.FindAllStringIndex(text, -1) {
		if strings.TrimSpace(text[last:at[1]]) != "" {
			out = append(out, textSentence{text: text[last:at[1]], start: last})
		}
		last = at[1]
	}
	if strings.TrimSpace(text[last:]) != "" {
		out = append(out, textSentence{text: text[last:], start: last})
	}
	return out
}

// findingText is what a finding claims: its message and action.
func findingText(f Finding) string {
	text := f.Message
	if f.Action != nil {
		text += "\n" + *f.Action
	}
	return text
}

// faulting reports an outcome that counts against the purchase.
func faulting(f Finding) bool {
	return f.Outcome == "fail" || f.Outcome == "unresolved"
}

// timingClaim is a finding that faults an earlier purchase, which adapts the
// base attack, for a technique its path adapts from a later purchase.
type timingClaim struct {
	purchase, technique, later, starts string
}

// timingCue is wording that faults a purchase for a technique: it begins,
// opens, introduces or enters the technique, or lacks or omits it.
var timingCue = regexp.MustCompile(`(?i)\b(begin|begins|began|beginning|start|starts|started|starting|open|opens|introduce|introduces|introducing|enter|enters|entry|lack|lacks|lacking|missing|omit|omits|absent)\b`)

// laterWording acknowledges that a technique comes later on the path.
var laterWording = regexp.MustCompile(`(?i)\b(later|afterwards?|until|not yet|from then)\b`)

// pronounSubject is a sentence opening with "It" or "Its", which refers to
// the finding's subject.
var pronounSubject = regexp.MustCompile(`(?i)^\s*(it|its)\b`)

// thisPurchase names the subject's purchase without its code.
var thisPurchase = regexp.MustCompile(`(?i)\bthis (purchase|tier|upgrade)\b`)

// timing checks a faulting finding whose subject names only purchases that
// adapt the base attack. For each such purchase and each technique its path
// adapts only later, which neither the purchase's name nor its planned text
// names, it reads the sentences of the message and action that name the
// technique:
//   - a sentence that also names a later purchase of that path, or says
//     later, acknowledges the timing and is context;
//   - one that refers to the purchase, by its code, its name, "this
//     purchase" or an opening "It", and faults it with a cue such as
//     "begins" or "lacks" is the claim (wrong);
//   - one that refers to the purchase without a cue is a claim code cannot
//     read (uncertain);
//   - any other sentence is context.
//
// A later technique is named by its full name or by a word of it that no
// other repertoire entry or the base attack uses, such as Tankman in Gear 4
// Tankman. The purchase names it by its full name or any word of it that the
// base attack does not use, so x-x-1 "Boundman Force" names Gear 4, Boundman
// and Snakeman traits.
func (c reviewClaims) timing(f Finding) (wrong, uncertain []timingClaim) {
	if c.plan == nil || c.plan.UpgradeIntents == nil || !faulting(f) {
		return nil, nil
	}
	plan := *c.plan
	subject := purchaseCodes(f.Subject)
	if len(subject) == 0 {
		return nil, nil
	}
	for _, code := range subject {
		path, tier, _ := purchasePosition(code)
		if !sameTechnique(plan.UpgradeIntents.At(path).At(tier).Technique, plan.Base.Name) {
			return nil, nil
		}
	}
	text := findingText(f)
	for _, code := range subject {
		path, tier, _ := purchasePosition(code)
		name := c.names[path][tier-1]
		own := name + "\n" + plan.Paths.At(path).Milestones.At(tier)
		seen := map[string]bool{}
		for later := tier + 1; later <= len(m.TierKeys); later++ {
			technique := strings.TrimSpace(plan.UpgradeIntents.At(path).At(later).Technique)
			key := strings.ToLower(strings.Join(strings.Fields(technique), " "))
			if technique == "" || seen[key] || sameTechnique(technique, plan.Base.Name) {
				continue
			}
			seen[key] = true
			if c.purchaseNames(own, technique) {
				continue
			}
			var laterCodes []string
			for t := tier + 1; t <= len(m.TierKeys); t++ {
				laterCodes = append(laterCodes, BuildCode(path, t))
			}
			fault, unclear := false, false
			for _, sentence := range findingSentences(text) {
				if !c.namesLater(sentence.text, technique) {
					continue
				}
				named := purchaseCodes(sentence.text)
				switch {
				case laterWording.MatchString(sentence.text) || slices.ContainsFunc(named, func(n string) bool { return slices.Contains(laterCodes, n) }):
				case !refersTo(sentence.text, code, name):
				case timingCue.MatchString(sentence.text):
					fault = true
				default:
					unclear = true
				}
			}
			claim := timingClaim{purchase: code, technique: strings.TrimSpace(plan.UpgradeIntents.At(path).At(tier).Technique), later: technique, starts: BuildCode(path, later)}
			if fault {
				wrong = append(wrong, claim)
			} else if unclear {
				uncertain = append(uncertain, claim)
			}
		}
	}
	return wrong, uncertain
}

// refersTo reports a sentence that names a purchase: its code, its name,
// "this purchase" or an opening "It".
func refersTo(sentence, code, name string) bool {
	return slices.Contains(purchaseCodes(sentence), code) ||
		(strings.TrimSpace(name) != "" && containsWords(sentence, name)) ||
		thisPurchase.MatchString(sentence) || pronounSubject.MatchString(sentence)
}

// techniqueStopWords are words of a technique name that do not identify it.
var techniqueStopWords = map[string]bool{"and": true, "the": true, "of": true, "with": true, "for": true, "from": true, "form": true, "forms": true, "mode": true, "trait": true, "traits": true, "technique": true, "style": true}

// techniqueWords are the words of a name that may identify a technique:
// lower case, at least three characters, not a stop word and not a word of
// the base attack's name.
func (c reviewClaims) techniqueWords(name string) []string {
	base := map[string]bool{}
	for _, word := range words(c.plan.Base.Name) {
		base[word] = true
	}
	var out []string
	for _, word := range words(name) {
		if len(word) >= 3 && !techniqueStopWords[word] && !base[word] && !slices.Contains(out, word) {
			out = append(out, word)
		}
	}
	return out
}

func words(text string) []string {
	return strings.FieldsFunc(strings.ToLower(text), func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) })
}

// containsWords reports text holding phrase as whole words, ignoring case
// and spacing.
func containsWords(text, phrase string) bool {
	fields := strings.Fields(phrase)
	if len(fields) == 0 {
		return false
	}
	quoted := make([]string, len(fields))
	for i, field := range fields {
		quoted[i] = regexp.QuoteMeta(field)
	}
	return regexp.MustCompile(`(?i)(^|[^\pL\pN])` + strings.Join(quoted, `\s+`) + `($|[^\pL\pN])`).MatchString(text)
}

// namesLater reports a sentence that names a technique: its full name or a
// word of it no other repertoire entry uses.
func (c reviewClaims) namesLater(sentence, technique string) bool {
	if containsWords(sentence, technique) {
		return true
	}
	for _, word := range c.techniqueWords(technique) {
		unique := true
		for _, entry := range c.plan.Repertoire {
			if !sameTechnique(entry.Name, technique) && slices.Contains(words(entry.Name), word) {
				unique = false
			}
		}
		if unique && containsWords(sentence, word) {
			return true
		}
	}
	return false
}

// purchaseNames reports a purchase name or planned text that names a
// technique: its full name or any word of it.
func (c reviewClaims) purchaseNames(text, technique string) bool {
	if containsWords(text, technique) {
		return true
	}
	return slices.ContainsFunc(c.techniqueWords(technique), func(word string) bool { return containsWords(text, word) })
}

// comparisonClaim retains the evidence for a suspected prose comparison.
type comparisonClaim struct {
	row comparisonRow
}

// gainText writes a gain to two decimals, as the review quotes them.
func gainText(value float64) string {
	return s.FormatNumber(math.Round(value*100) / 100)
}

// namePart matches text that ends in a word a number completes as a name or
// position: a capitalized word inside a sentence, as in "the Gear 2 role",
// or tier, path, level or purchase.
var namePart = regexp.MustCompile(`(?:[^.!?\s]\s+\p{Lu}\pL*|(?i:\b(?:tier|tiers|path|paths|level|purchase|purchases)))\s+$`)

// quotedNumber matches a number a finding writes, such as 4.66 or 18,000.
var quotedNumber = regexp.MustCompile(`\d+(?:,\d{3})*(?:\.\d+)?`)

// readNumber reads a quoted number and the rounding its written digits
// allow: 4.66 stands for any value from 4.655 to 4.665.
func readNumber(written string) (value, tolerance float64, ok bool) {
	value, err := strconv.ParseFloat(strings.ReplaceAll(written, ",", ""), 64)
	if err != nil {
		return 0, 0, false
	}
	decimals := 0
	if dot := strings.IndexByte(written, '.'); dot >= 0 {
		decimals = len(written) - dot - 1
	}
	return value, 0.5*math.Pow(10, -float64(decimals)) + 1e-9, true
}

// comparators is wording that orders two gains: more, less or equal.
var comparators = []struct {
	polarity string
	pattern  *regexp.Regexp
}{
	{"more", regexp.MustCompile(`(?i)\b(more than|greater than|larger than|higher than|bigger than|exceeds?|exceeding|outgains?|surpass(?:es)?|outpaces?|outperforms?|beats?)\b`)},
	{"less", regexp.MustCompile(`(?i)\b(less than|smaller than|lower than|fewer than|below|short of|trails?)\b`)},
	{"equal", regexp.MustCompile(`(?i)\b(at least|as much as|as large as|as high as|as big as|matches|equals?|equal to|on par with|rivals?)\b`)},
}

// negation before a comparator makes its direction unreadable.
var negation = regexp.MustCompile(`(?i)\b(not|no|never|nor)\s+(\w+\s+)?$|n't\s+(\w+\s+)?$`)

var sideWords = regexp.MustCompile(`(?i)\bside[- ](purchase|path|upgrade)s?\b`)
var capstoneWords = regexp.MustCompile(`(?i)\b(capstone|fifth purchase)\b`)
var capstonePronoun = regexp.MustCompile(`(?i)\bthis (purchase|capstone)\b`)

// comparison checks the side gains a finding quotes against the
// againstCapstone comparisons of the capstones it names, by purchase code
// (x-5-x) or pure build (0-5-0). A row is one metric of one comparison. A
// quoted number matches a row's side gain when it equals the gain rounded to
// the digits quoted (4.66 for 4.6598). The finding's own words narrow the
// rows: the metric it names, when it names only direct or only group, the
// side purchase's price it quotes (180 Gold) and the side purchase's code.
// The sentence holding the number is read for one comparator (more, less or
// equal) between a side operand (the number, the side purchase's code or
// "side purchase") and a capstone operand (the capstone's code or gain,
// "capstone", or, when the subject names the capstone, "this purchase" or an
// opening "It"). The finding states the opposite of the comparison (wrong)
// only when one row is left, the number matches no capstone gain, and the
// sentence says the side gain is at least the capstone's while
// sideGainAtLeastCapstone is false, or less while it is true. A number that
// matches several rows or metrics is ambiguous and is never corrected. A
// faulting finding whose claim code cannot read is uncertain: a stated
// direction that one of several matched rows contradicts, or one matched row
// whose side gain is below the capstone's with no readable direction.
func (c reviewClaims) comparison(f Finding, currency string) (wrong []comparisonClaim, uncertain []string) {
	text := findingText(f)
	named := buildCodePattern.FindAllString(f.Subject+"\n"+text, -1)
	var rows []int
	for index, row := range c.rows {
		if namesCapstone(named, row.capstone) {
			rows = append(rows, index)
		}
	}
	if len(rows) == 0 {
		return nil, nil
	}
	gains, prices := quotedNumbers(text, currency)
	// The metric the finding names, when it names only one.
	direct, group := strings.Contains(strings.ToLower(text), "direct"), strings.Contains(strings.ToLower(text), "group")
	metricNamed := func(metric string) bool {
		switch {
		case direct && !group:
			return metric == TimeAveragedDirect
		case group && !direct:
			return metric == TimeAveragedGroup
		}
		return true
	}
	for _, number := range gains {
		var sides []int
		capstoneGain := false
		for _, index := range rows {
			row := c.rows[index]
			if !metricNamed(row.metric) {
				continue
			}
			if math.Abs(row.sideGain-number.value) <= number.tolerance {
				sides = append(sides, index)
			}
			if math.Abs(row.capstoneGain-number.value) <= number.tolerance {
				capstoneGain = true
			}
		}
		// The side purchase's price or code the finding quotes narrows the
		// comparisons its gain may come from.
		sides = narrowed(sides, func(index int) bool {
			return slices.ContainsFunc(prices, func(p quotedValue) bool {
				spent, ok := evidenceNumber(c.rows[index].sidePrice)
				return ok && math.Abs(spent-p.value) <= p.tolerance
			})
		})
		sides = narrowed(sides, func(index int) bool { return slices.Contains(named, c.rows[index].side) })
		if len(sides) == 0 {
			continue
		}
		unique := len(sides) == 1 && !capstoneGain
		var sentence textSentence
		for _, candidate := range findingSentences(text) {
			if number.start >= candidate.start && number.start < candidate.start+len(candidate.text) {
				sentence = candidate
			}
		}
		saysAtLeast, readable := c.direction(sentence, number.start, number.end, sides, buildCodePattern.FindAllString(f.Subject, -1))
		switch {
		case readable && unique && saysAtLeast != c.rows[sides[0]].atLeast:
			wrong = append(wrong, comparisonClaim{row: c.rows[sides[0]]})
		case !faulting(f) || unique && (readable || c.rows[sides[0]].atLeast):
		case readable && slices.ContainsFunc(sides, func(index int) bool { return c.rows[index].atLeast != saysAtLeast }):
			uncertain = append(uncertain, fmt.Sprintf("%s quotes %s, which matches more than one gain in the %s comparisons of purchaseEvidence, so code cannot tell which comparison it states the opposite of.", f.ID, number.text, AgainstCapstone))
		case unique && !readable:
			row := c.rows[sides[0]]
			uncertain = append(uncertain, fmt.Sprintf("%s quotes %s, the %s side purchase's %s gain from %s to %s, which is below %s's %s (%s is false), and code cannot tell what it says about it.", f.ID, number.text, row.side, row.metric, row.from, row.to, row.capstone, gainText(row.capstoneGain), SideGainAtLeastCapstone))
		}
	}
	if len(wrong) > 0 {
		return wrong, nil
	}
	return nil, uncertain
}

// namesCapstone reports build codes that name a capstone: its purchase code,
// such as x-5-x, or the build that buys only it, such as 0-5-0.
func namesCapstone(codes []string, capstone string) bool {
	return slices.Contains(codes, capstone) || slices.Contains(codes, pureBuild(capstone))
}

// pureBuild is the build that buys only one purchase's path: 0-5-0 for x-5-x.
func pureBuild(code string) string {
	return strings.ReplaceAll(code, "x", "0")
}

// narrowed keeps the rows that pass keep, or all of them when none does.
func narrowed(rows []int, keep func(int) bool) []int {
	var out []int
	for _, index := range rows {
		if keep(index) {
			out = append(out, index)
		}
	}
	if len(out) == 0 {
		return rows
	}
	return out
}

// quotedValue is one number a finding writes, where it stands and the
// rounding its digits allow.
type quotedValue struct {
	text             string
	start, end       int
	value, tolerance float64
}

// quotedNumbers lists the numbers a finding's text writes, split into gains
// and prices: a number followed by the currency, such as 180 Gold, is a
// price. Digits inside a build code such as x-5-x or a word such as path1
// are not numbers, nor is a number that ends a name or a position, such as
// Gear 2 or tier 5.
func quotedNumbers(text, currency string) (gains, prices []quotedValue) {
	masked := buildCodePattern.ReplaceAllStringFunc(text, func(code string) string { return strings.Repeat("#", len(code)) })
	var price *regexp.Regexp
	if currency != "" {
		price = regexp.MustCompile(`(?i)^\s*` + regexp.QuoteMeta(currency) + `\b`)
	}
	for _, at := range quotedNumber.FindAllStringIndex(masked, -1) {
		before, after := rune(0), rune(0)
		if at[0] > 0 {
			before = rune(masked[at[0]-1])
		}
		if at[1] < len(masked) {
			after = rune(masked[at[1]])
		}
		if unicode.IsLetter(before) || unicode.IsDigit(before) || before == '.' || before == ',' || before == '#' || unicode.IsLetter(after) || unicode.IsDigit(after) || after == '#' {
			continue
		}
		if namePart.MatchString(masked[:at[0]]) {
			continue
		}
		written := masked[at[0]:at[1]]
		value, tolerance, ok := readNumber(written)
		if !ok {
			continue
		}
		entry := quotedValue{text: written, start: at[0], end: at[1], value: value, tolerance: tolerance}
		if price != nil && price.MatchString(masked[at[1]:]) {
			prices = append(prices, entry)
		} else {
			gains = append(gains, entry)
		}
	}
	return gains, prices
}

// direction reads what one sentence says about a quoted side gain: whether
// it is at least the capstone's gain. readable is false unless the sentence
// holds one comparator, without a negation before it, between a side
// operand and a capstone operand.
func (c reviewClaims) direction(sentence textSentence, start, end int, sides []int, subjectCodes []string) (saysAtLeast, readable bool) {
	type span struct {
		start, end int
		kind       string
	}
	text := sentence.text
	var found []span
	for _, comparator := range comparators {
		for _, at := range comparator.pattern.FindAllStringIndex(text, -1) {
			found = append(found, span{at[0], at[1], comparator.polarity})
		}
	}
	if len(found) == 0 {
		return false, false
	}
	slices.SortFunc(found, func(a, b span) int { return a.start - b.start })
	// Adjacent comparators such as "at least as much as" are one.
	group := found[0]
	for _, next := range found[1:] {
		if next.start-group.end > 4 {
			return false, false
		}
		switch {
		case group.kind == "equal":
			group.kind = next.kind
		case next.kind != "equal" && next.kind != group.kind:
			return false, false
		}
		group.end = next.end
	}
	if negation.MatchString(text[:group.start]) {
		return false, false
	}
	var operands []span
	add := func(pattern *regexp.Regexp, kind string) {
		for _, at := range pattern.FindAllStringIndex(text, -1) {
			operands = append(operands, span{at[0], at[1], kind})
		}
	}
	add(sideWords, "side")
	add(capstoneWords, "capstone")
	operands = append(operands, span{start - sentence.start, end - sentence.start, "side"})
	for _, at := range buildCodePattern.FindAllStringIndex(text, -1) {
		code := text[at[0]:at[1]]
		for _, index := range sides {
			switch code {
			case c.rows[index].side:
				operands = append(operands, span{at[0], at[1], "side"})
			case c.rows[index].capstone, pureBuild(c.rows[index].capstone):
				operands = append(operands, span{at[0], at[1], "capstone"})
			}
		}
	}
	for _, index := range sides {
		for _, at := range quotedNumber.FindAllStringIndex(text, -1) {
			value, tolerance, ok := readNumber(text[at[0]:at[1]])
			if ok && at[0] != start-sentence.start && math.Abs(c.rows[index].capstoneGain-value) <= tolerance {
				operands = append(operands, span{at[0], at[1], "capstone"})
			}
		}
	}
	subjectIsCapstone := slices.ContainsFunc(sides, func(index int) bool {
		return namesCapstone(subjectCodes, c.rows[index].capstone) && !slices.Contains(subjectCodes, c.rows[index].side)
	})
	if subjectIsCapstone {
		add(capstonePronoun, "capstone")
		add(pronounSubject, "capstone")
	}
	var before, after *span
	for i := range operands {
		operand := &operands[i]
		if operand.end <= group.start && (before == nil || operand.end > before.end) {
			before = operand
		}
		if operand.start >= group.end && (after == nil || operand.start < after.start) {
			after = operand
		}
	}
	if before == nil || after == nil || before.kind == after.kind {
		return false, false
	}
	switch group.kind {
	case "equal":
		return true, true
	case "more":
		return before.kind == "side", true
	}
	return before.kind == "capstone", true
}

// humanReviewFindings are the unresolved Findings code adds to a Result for
// model findings whose timing or comparison claim it could not read. Each
// keeps the model finding as the review wrote it and asks for it to be
// checked before it is acted on.
func (c reviewClaims) humanReviewFindings(findings []Finding, currency string) []Finding {
	var out []Finding
	for _, f := range findings {
		var reasons []string
		if suspected, uncertain := c.timing(f); len(suspected)+len(uncertain) > 0 {
			for _, claim := range append(suspected, uncertain...) {
				reasons = append(reasons, fmt.Sprintf("Code could not tell whether %s faults %s for %s: %s adapts %s, and %s starts at %s.", f.ID, claim.purchase, claim.later, claim.purchase, claim.technique, claim.later, claim.starts))
			}
		}
		if suspected, uncertain := c.comparison(f, currency); faulting(f) {
			reasons = append(reasons, uncertain...)
			for _, claim := range suspected {
				row := claim.row
				reasons = append(reasons, fmt.Sprintf("Code could not verify the meaning of %s's comparison: %s adds %s and %s adds %s %s. Check whether the prose compares these absolute gains or a different measure.", f.ID, row.side, gainText(row.sideGain), row.capstone, gainText(row.capstoneGain), row.metric))
			}
		}
		if len(reasons) == 0 {
			continue
		}
		out = append(out, Finding{
			ID: "deterministic.review." + strings.TrimPrefix(f.ID, "model."), Method: "deterministic", Category: "evidence",
			Severity: severity("unresolved"), Outcome: "unresolved", Subject: f.Subject, Rule: HumanReviewRule,
			Message:  strings.Join(reasons, " ") + " The finding is published as the review wrote it.",
			Evidence: []string{},
			Action:   act("Check " + f.ID + " against the plan and purchaseEvidence before acting on it."),
		})
	}
	return out
}

// HumanReviewRule is the rule of a Finding code adds when it cannot read a
// model finding's timing or comparison claim.
const HumanReviewRule = "review-claim-unread"
