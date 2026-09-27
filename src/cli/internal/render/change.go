package render

import (
	"math"
	"strconv"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// A change reads as "value → value (delta)": the values show the result of
// buying only that change, and the delta what the change itself does.
// Multipliers read as percentages. The unit sheet, its crosspath rows and
// UnitLab's kit stats all use these helpers, so every view words a change
// the same way.
//
// Percentages round to whole numbers from 10% up, to one decimal from 1% and
// to two below that, so a small change never reads as 0%.

// percentNumber writes a percentage without its sign: 17.647 is "18".
func percentNumber(value float64) string {
	value = math.Abs(value)
	digits := 2
	switch {
	case value >= 10:
		digits = 0
	case value >= 1:
		digits = 1
	}
	rounded, _ := strconv.ParseFloat(strconv.FormatFloat(value, 'f', digits, 64), 64)
	return s.FormatNumber(rounded)
}

// signedPercent writes a fraction as a signed percentage: 1 is "+100%" and
// -0.2 is "-20%".
func signedPercent(fraction float64) string {
	if fraction < 0 {
		return "-" + percentNumber(fraction*100) + "%"
	}
	return "+" + percentNumber(fraction*100) + "%"
}

// signedNumber writes a difference with its sign: "+1" or "-0.5".
func signedNumber(value float64) string {
	if value < 0 {
		return "-" + decimal(-value)
	}
	return "+" + decimal(value)
}

// delta is the difference between two shown values with their unit: "+1",
// "-0.5 s" or, for values in percent, "+20 percentage points".
func delta(before, after float64, unitText string) string {
	if unitText == "%" {
		return pointsDelta(before, after)
	}
	return signedNumber(after-before) + unitText
}

// pointsDelta is the difference between two percentages: 100 and 150 are
// "+50 percentage points".
func pointsDelta(before, after float64) string {
	difference := after - before
	sign := "+"
	if difference < 0 {
		sign = "-"
	}
	return sign + percentNumber(difference) + " percentage points"
}

// rateChange is how much an attack's rate changes when its interval goes
// from before to after: the rate is 1/interval, so the change is
// before/after - 1. An interval multiplied by 0.85 attacks 1/0.85 - 1, about
// 18%, faster.
func rateChange(before, after float64) float64 {
	if after <= 0 || before <= 0 {
		return 0
	}
	return before/after - 1
}

// speed words a rate change: "18% faster" or "17% slower".
func speed(change float64) string {
	if change < 0 {
		return percentNumber(change*100) + "% slower"
	}
	return percentNumber(change*100) + "% faster"
}

// attackSpeedChange is the delta of an attack interval: "attacks 18%
// faster". It is empty when the interval is unchanged.
func attackSpeedChange(before, after float64) string {
	if before == after {
		return ""
	}
	return "attacks " + speed(rateChange(before, after))
}

// activeSpeed words an Active Ability's interval multiplier as the attack
// speed it gives: 0.5 is "100% faster".
func activeSpeed(multiplier float64) string {
	return speed(rateChange(1, multiplier))
}

// activeDamage words an Active Ability's damage multiplier as the damage it
// adds: 2 is "+100%".
func activeDamage(multiplier float64) string {
	return signedPercent(multiplier - 1)
}

// percentOfDamage words a follow-up's damage multiplier as a share of the
// hit's damage: 0.5 is "50%".
func percentOfDamage(multiplier float64) string {
	return percentNumber(multiplier*100) + "%"
}
