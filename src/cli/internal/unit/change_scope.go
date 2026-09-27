package unit

import (
	"fmt"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Change scope (#61, OPUS-NET-61-7). Luffy's v30 x-5-x said its burn and
// extra damage applied only during Gear 2 Burst, while its burn and +2
// damage were changes of the automatic attack, owned in every build, and
// the review missed it. An Active Ability scopes only its unlock, its
// modifyBoost stats and a follow-up that targets the boost; every other
// change applies to the automatic attack and the Active inherits it. The
// review reads each purchase's typed changes by scope (changeScope) beside
// its planned text, so it can compare the two without code reading prose.

// ChangeScope lists a purchase's typed changes by scope: base, what applies
// to the automatic attack in every build that owns the purchase (the
// Active's boosted attack inherits it), and boost, what applies only to the
// Active Ability. Each change is a short label, such as "damage add 2",
// "burn damagePerSecond set 2" or "boost damageMultiplier multiply 1.5".
func ChangeScope(tier m.Tier) *s.Object {
	base, boost := []any{}, []any{}
	for _, change := range tier.Changes {
		label := changeLabel(change)
		if change.Kind == "unlockBoost" || change.Kind == "modifyBoost" || (change.Kind == "followUp" && change.Target == "boost") {
			boost = append(boost, label)
		} else {
			base = append(base, label)
		}
	}
	return s.NewObject().Set("base", base).Set("boost", boost)
}

// changeDimension is what one typed change of the purchase at a path's
// tier changes (#61, SOL-61-20): a stat with the direction the change moves
// it, as "damage raised" or "intervalSeconds lowered"; a status effect;
// bonus damage against a property; detection; a follow-up; the Active
// Ability; or another change's kind. Amounts are left out, so two purchases
// that raise damage by different amounts change one dimension, and a
// purchase that lengthens the interval does not share a shortened one. A
// set takes the direction its purchase resolves on its pure path
// (setDirection), so a set to a higher damage and an added damage change
// one dimension (SOL-82-01); a set that moves nothing there, or a boost stat
// its purchase unlocks, keeps "set".
func changeDimension(blueprint *m.Blueprint, pathIndex, tier int, c m.Change) string {
	switch c.Kind {
	case "stat", "modifyBoost":
		name := c.Stat
		if c.Kind == "modifyBoost" {
			name = "boost " + c.Stat
		}
		switch {
		case c.Operation == "set":
			if direction := setDirection(blueprint, pathIndex, tier, c); direction != "" {
				return name + " " + direction
			}
			return name + " set"
		case c.Operation == "add" && c.Number > 0, c.Operation == "multiply" && c.Number > 1:
			return name + " raised"
		case c.Operation == "add" && c.Number < 0, c.Operation == "multiply" && c.Number < 1:
			return name + " lowered"
		}
		return name
	case "status":
		return c.Effect
	case "bonusDamage":
		return "bonus damage against " + c.Property
	case "camo":
		return "camo detection"
	case "detection":
		return "detects " + c.Trait
	case "followUp":
		if c.Target == "boost" {
			return "Active Ability follow-up"
		}
		return "follow-up"
	case "unlockBoost":
		return "Active Ability"
	default:
		return c.Kind
	}
}

// setDirection is the direction in which the purchase at a path's tier
// moves the stat a set change of it sets, on its pure path: "raised" when
// the stat's resolved value with that path's purchases through the tier
// alone is higher than through the tier before, "lowered" when lower, and
// "" when it is equal or, for a boost stat, the path owns no Active Ability
// before the tier.
func setDirection(blueprint *m.Blueprint, pathIndex, tier int, c m.Change) string {
	if blueprint == nil {
		return ""
	}
	value := func(bought int) (float64, bool) {
		var selection m.Selection
		selection[pathIndex] = bought
		build := m.ResolveUnchecked(blueprint, selection)
		if c.Kind == "stat" {
			return build.BaseAttack.Stats.Get(c.Stat), true
		}
		for _, ability := range build.Abilities {
			if ability.Path == m.PathKeys[pathIndex] {
				boost := m.Boost{DurationSeconds: ability.DurationSeconds, CooldownSeconds: ability.CooldownSeconds, DamageMultiplier: ability.DamageMultiplier, IntervalMultiplier: ability.IntervalMultiplier, RangeBonus: ability.RangeBonus}
				return boost.Get(c.Stat), true
			}
		}
		return 0, false
	}
	before, hadBefore := value(tier - 1)
	after, hasAfter := value(tier)
	switch {
	case !hadBefore || !hasAfter || after == before:
		return ""
	case after > before:
		return "raised"
	default:
		return "lowered"
	}
}

// changeLabel is a short label of one typed change.
func changeLabel(c m.Change) string {
	number := s.FormatNumber(c.Number)
	switch c.Kind {
	case "stat":
		return fmt.Sprintf("%s %s %s", c.Stat, c.Operation, number)
	case "modifyBoost":
		return fmt.Sprintf("boost %s %s %s", c.Stat, c.Operation, number)
	case "status":
		return fmt.Sprintf("%s %s %s %s", c.Effect, c.Field, c.Operation, number)
	case "bonusDamage":
		return fmt.Sprintf("bonus damage against %s add %s", c.Property, number)
	case "camo":
		return fmt.Sprintf("camo detection %t", c.Bool)
	case "detection":
		return fmt.Sprintf("detects %s %t", c.Trait, c.Bool)
	case "followUp":
		if c.FollowUp == nil {
			return "follow-up"
		}
		return fmt.Sprintf("follow-up %s: %s hits of %s damage within %s", c.FollowUp.Name, s.FormatNumber(c.FollowUp.Count), s.FormatNumber(c.FollowUp.DamageMultiplier), s.FormatNumber(c.FollowUp.Radius))
	case "unlockBoost":
		if c.Boost == nil {
			return "unlocks the Active Ability"
		}
		return fmt.Sprintf("unlocks the Active Ability %s", c.Boost.Name)
	default:
		return fmt.Sprintf("%s %s", c.Kind, c.Text)
	}
}
