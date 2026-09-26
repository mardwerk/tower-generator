package unit

import (
	"math"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Review evidence derived from retained purchase evidence. Every value comes
// from numbers the evidence already holds, which code computed from the
// resolved builds. Deriving it when read keeps saved artifacts exact: older
// drafts gain it without a recompute, and check still compares only the
// retained fields. It is evidence for review and reading, never a gate, a
// finding or a pass/fail outcome.

// Time-averaged metrics: a damage rate averaged over one Active Ability
// cycle with the Active used whenever it is ready. Without an owned Active
// they equal the ordinary rate.
const (
	TimeAveragedDirect = "time-averaged direct damage rate"
	TimeAveragedGroup  = "time-averaged group damage rate upper bound"
)

// Per100GoldAgainstCapstone is the key a side purchase's comparison with its
// main path's fifth purchase is stored under.
const Per100GoldAgainstCapstone = "per100GoldAgainstCapstone"

const activeDuty = "active duty fraction"

// timeAveragedMetrics names each time-averaged metric with the ordinary rate
// and Active peak it is averaged from.
var timeAveragedMetrics = []struct{ name, ordinary, peak string }{
	{TimeAveragedDirect, "direct damage rate", "active peak direct damage rate"},
	{TimeAveragedGroup, "group damage rate upper bound", "active peak group damage rate upper bound"},
}

const derivedLimitation = "Time-averaged rates use the Active Ability whenever it is ready and equal the ordinary rate without one. per100GoldAgainstCapstone compares a side purchase's time-averaged gain per 100 Gold with the gain of its main path's own fifth purchase. Both are derived evidence, not thresholds."

// ReviewPurchaseEvidence returns a copy of retained purchase evidence with
// two derived readings. Wherever the evidence shows an Active Ability's
// peak and duty fraction, the time-averaged direct and group rates follow
// them, with their change for a purchase. Each crosspath purchase also
// gets per100GoldAgainstCapstone: its time-averaged direct and group gain
// per 100 Gold beside the same gain of its main path's fifth purchase
// (4-0-0 to 5-0-0 for the top path). A value that cannot be computed is null.
func ReviewPurchaseEvidence(evaluation *s.Object) *s.Object {
	if evaluation == nil {
		return nil
	}
	out := s.Clone(evaluation).(*s.Object)
	paths, _ := out.Get("paths")
	for _, value := range asList(paths) {
		path, ok := value.(*s.Object)
		if !ok {
			continue
		}
		milestones, _ := path.Get("milestones")
		for _, milestone := range asList(milestones) {
			averageDeltas(milestone)
		}
		if comparison, ok := objectAt(path, "capstoneComparison"); ok {
			for _, key := range [][]string{{"tier4", "metrics"}, {"tier5", "metrics"}, {"sameBudgetTier4Copies", "perCopyMetrics"}} {
				if holder, ok := objectAt(comparison, key[0]); ok {
					if metrics, ok := objectAt(holder, key[1]); ok {
						holder.Set(key[1], averageMetrics(metrics))
					}
				}
			}
		}
		capstone, code := capstoneStep(path, asList(milestones))
		crosspaths, _ := path.Get("crosspaths")
		for _, value := range asList(crosspaths) {
			averageDeltas(value)
			crosspath, ok := value.(*s.Object)
			if !ok || capstone == nil {
				continue
			}
			side := per100Gold(crosspath)
			step := per100Gold(capstone)
			against := s.NewObject().Set("capstone", code)
			for _, metric := range timeAveragedMetrics {
				against.Set(metric.name, s.NewObject().Set("sidePurchase", side[metric.name]).Set("capstone", step[metric.name]))
			}
			crosspath.Set(Per100GoldAgainstCapstone, against)
		}
	}
	if limitations, ok := out.Get("limitations"); ok {
		if items, ok := limitations.([]any); ok {
			out.Set("limitations", append(items, derivedLimitation))
		}
	}
	return out
}

func asList(value any) []any {
	items, _ := value.([]any)
	return items
}

func objectAt(object *s.Object, key string) (*s.Object, bool) {
	value, _ := object.Get(key)
	child, ok := value.(*s.Object)
	return child, ok
}

func evidenceNumber(value any) (float64, bool) {
	f, ok := value.(float64)
	return f, ok && !math.IsNaN(f) && !math.IsInf(f, 0)
}

// timeAveraged reads the time-averaged rate from one side of a metric set.
// read returns a metric's value and whether the set names it.
func timeAveraged(read func(string) (any, bool), ordinary, peak string) any {
	base, ok := read(ordinary)
	rate, valid := evidenceNumber(base)
	if !ok || !valid {
		return nil
	}
	dutyValue, owned := read(activeDuty)
	if !owned {
		return rate
	}
	peakValue, _ := read(peak)
	duty, validDuty := evidenceNumber(dutyValue)
	top, validPeak := evidenceNumber(peakValue)
	if !validDuty || !validPeak {
		return nil
	}
	if average := m.TimeAveraged(rate, top, duty); !math.IsNaN(average) && !math.IsInf(average, 0) {
		return average
	}
	return nil
}

// averageMetrics copies a metric set, placing the time-averaged rates right
// after the duty fraction when the build owns an Active.
func averageMetrics(metrics *s.Object) *s.Object {
	if !metrics.Has(activeDuty) {
		return metrics
	}
	out := s.NewObject()
	for _, key := range metrics.Keys() {
		value, _ := metrics.Get(key)
		out.Set(key, value)
		if key == activeDuty {
			for _, metric := range timeAveragedMetrics {
				out.Set(metric.name, timeAveraged(metrics.Get, metric.ordinary, metric.peak))
			}
		}
	}
	return out
}

// deltaSide reads the before or after value of each metric delta.
func deltaSide(deltas *s.Object, side string) func(string) (any, bool) {
	return func(key string) (any, bool) {
		delta, ok := objectAt(deltas, key)
		if !ok {
			return nil, false
		}
		return delta.Get(side)
	}
}

// averagedDelta is the before, after and change of one time-averaged rate.
func averagedDelta(deltas *s.Object, ordinary, peak string) *s.Object {
	before := timeAveraged(deltaSide(deltas, "before"), ordinary, peak)
	after := timeAveraged(deltaSide(deltas, "after"), ordinary, peak)
	var change any
	x, okBefore := evidenceNumber(before)
	y, okAfter := evidenceNumber(after)
	if okBefore && okAfter {
		if delta := y - x; !math.IsNaN(delta) && !math.IsInf(delta, 0) {
			change = delta
		}
	}
	return s.NewObject().Set("before", before).Set("after", after).Set("change", change)
}

// averageDeltas places the time-averaged deltas right after the duty
// fraction when either build of a purchase owns an Active.
func averageDeltas(value any) {
	purchase, ok := value.(*s.Object)
	if !ok {
		return
	}
	deltas, ok := objectAt(purchase, "metricDeltas")
	if !ok || !deltas.Has(activeDuty) {
		return
	}
	out := s.NewObject()
	for _, key := range deltas.Keys() {
		delta, _ := deltas.Get(key)
		out.Set(key, delta)
		if key == activeDuty {
			for _, metric := range timeAveragedMetrics {
				out.Set(metric.name, averagedDelta(deltas, metric.ordinary, metric.peak))
			}
		}
	}
	purchase.Set("metricDeltas", out)
}

// per100Gold is a purchase's time-averaged gain per 100 Gold, by metric.
// It is null without a positive price or a finite gain.
func per100Gold(purchase *s.Object) map[string]any {
	out := map[string]any{}
	deltas, _ := objectAt(purchase, "metricDeltas")
	price, _ := purchase.Get("incrementalGold")
	gold, validGold := evidenceNumber(price)
	for _, metric := range timeAveragedMetrics {
		out[metric.name] = nil
		if deltas == nil || !validGold || gold <= 0 {
			continue
		}
		delta, _ := averagedDelta(deltas, metric.ordinary, metric.peak).Get("change")
		if change, valid := evidenceNumber(delta); valid {
			out[metric.name] = change / gold * 100
		}
	}
	return out
}

// capstoneStep finds a path's own fifth purchase among its milestones and
// its build code, such as 5-x-x.
func capstoneStep(path *s.Object, milestones []any) (*s.Object, string) {
	key, _ := path.Get("path")
	for index, name := range m.PathKeys {
		if name != key {
			continue
		}
		for _, value := range milestones {
			milestone, ok := value.(*s.Object)
			if !ok {
				continue
			}
			to, _ := milestone.Get("to")
			selection := asList(to)
			if len(selection) == len(m.PathKeys) && selection[index] == 5.0 {
				return milestone, BuildCode(index, 5)
			}
		}
	}
	return nil, ""
}
