package unit

import (
	"fmt"
	"strconv"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Capstone payoff (#61, SOL-61-13). The v36 Luffy x-x-5 cost 21500 Gold and
// its 4 to 5 step added 0.387 time-averaged direct and 2.709 group damage per
// second, while adding 1-x-x to 0-0-4 for 140 Gold added 1.548 and 3.096.
// Its fifth purchase verdict judged its distinct follow-up kind and passed
// it (OPUS-NET-61-16). Each fifth purchase verdict now also judges payoff
// and price, and gets the numbers beside its subject: the capstone's price
// and absolute time-averaged gains, and every side purchase added to its
// fourth purchase that costs less, with its price and gains and whether it
// adds at least as much on both. The numbers come from the reviewed
// purchase evidence (ReviewPurchaseEvidence), so they are the ones
// purchaseEvidence shows. No ratio is computed and nothing here is a gate:
// the verdict is the review's judgment.

// CapstonePayoffs are the payoff numbers of each path's fifth purchase by
// path key, read from reviewed purchase evidence; a path without its fifth
// purchase step is left out.
func CapstonePayoffs(evidence *s.Object, currency string) map[string]*s.Object {
	out := map[string]*s.Object{}
	if evidence == nil {
		return out
	}
	paths, _ := evidence.Get("paths")
	for _, value := range asList(paths) {
		path, ok := value.(*s.Object)
		if !ok {
			continue
		}
		key, _ := path.Get("path")
		index := -1
		for i, name := range m.PathKeys {
			if name == key {
				index = i
			}
		}
		milestones, _ := path.Get("milestones")
		capstone, code := capstoneStep(path, asList(milestones))
		if index < 0 || capstone == nil {
			continue
		}
		price := incrementalPrice(capstone)
		gain := gains(capstone)
		var fourth m.Selection
		fourth[index] = 4
		payoff := s.NewObject().
			Set("fourthPurchase", selectionCode(fourth)).
			Set("price", price).
			Set("gains", metricObject(gain))
		sides, facts := []any{}, []any{}
		crosspaths, _ := path.Get("crosspaths")
		for _, value := range asList(crosspaths) {
			crosspath, ok := value.(*s.Object)
			if !ok || !atMainTier(crosspath, index, 4) {
				continue
			}
			sidePrice, okSide := evidenceNumber(incrementalPrice(crosspath))
			capstonePrice, okCapstone := evidenceNumber(price)
			if !okSide || !okCapstone || sidePrice >= capstonePrice {
				continue
			}
			sideGain := gains(crosspath)
			atLeast := true
			for _, metric := range timeAveragedMetrics {
				x, okX := evidenceNumber(sideGain[metric.name])
				y, okY := evidenceNumber(gain[metric.name])
				atLeast = atLeast && okX && okY && x >= y
			}
			to, _ := crosspath.Get("to")
			side := sideCode(crosspath)
			sides = append(sides, s.NewObject().
				Set("code", side).
				Set("build", listCode(asList(to))).
				Set("price", sidePrice).
				Set("gains", metricObject(sideGain)).
				Set("atLeastCapstoneGains", atLeast))
			if atLeast {
				facts = append(facts, fmt.Sprintf("Adding %s to %s costs %s %s and adds %s time-averaged direct and %s group damage per second, at least the %s and %s that %s adds for %s %s.",
					side, listCode(asList(valueOf(crosspath, "from"))), s.FormatNumber(sidePrice), inCurrency(currency),
					payoffGain(sideGain[TimeAveragedDirect]), payoffGain(sideGain[TimeAveragedGroup]),
					payoffGain(gain[TimeAveragedDirect]), payoffGain(gain[TimeAveragedGroup]), code, s.FormatNumber(capstonePrice), inCurrency(currency)))
			}
		}
		out[m.PathKeys[index]] = payoff.Set("cheaperSidePurchases", sides).Set("facts", facts)
	}
	return out
}

// atMainTier reports a crosspath comparison whose main path stays at tier.
func atMainTier(crosspath *s.Object, index, tier int) bool {
	from, to := asList(valueOf(crosspath, "from")), asList(valueOf(crosspath, "to"))
	return len(from) == len(m.PathKeys) && len(to) == len(m.PathKeys) && from[index] == float64(tier) && to[index] == float64(tier)
}

func valueOf(object *s.Object, key string) any {
	value, _ := object.Get(key)
	return value
}

// metricObject lists gains by time-averaged metric, in their order.
func metricObject(gain map[string]any) *s.Object {
	out := s.NewObject()
	for _, metric := range timeAveragedMetrics {
		out.Set(metric.name, gain[metric.name])
	}
	return out
}

// selectionCode is a concrete build code, such as 0-0-4.
func selectionCode(selection m.Selection) string {
	return fmt.Sprintf("%d-%d-%d", selection[0], selection[1], selection[2])
}

// listCode is the build code of a selection as evidence stores it, such as
// [1, 0, 4]; empty when it is not three tiers.
func listCode(selection []any) string {
	var tiers m.Selection
	if len(selection) != len(tiers) {
		return ""
	}
	for i, value := range selection {
		tier, ok := value.(float64)
		if !ok {
			return ""
		}
		tiers[i] = int(tier)
	}
	return selectionCode(tiers)
}

// payoffGain reads a gain to three decimals, or "an unknown amount".
func payoffGain(value any) string {
	if gain, ok := evidenceNumber(value); ok {
		return strconv.FormatFloat(gain, 'f', 3, 64)
	}
	return "an unknown amount"
}
