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

// changeDimension is what one typed change changes (#61, SOL-61-20): a stat
// with the direction the change moves it, as "damage raised" or
// "intervalSeconds lowered", or "set" when its operation sets it; a status
// effect; bonus damage against a property; detection; a follow-up; the
// Active Ability; or another change's kind. Amounts are left out, so two
// purchases that raise damage by different amounts change one dimension,
// and a purchase that lengthens the interval does not share a shortened one.
func changeDimension(c m.Change) string {
	switch c.Kind {
	case "stat", "modifyBoost":
		name := c.Stat
		if c.Kind == "modifyBoost" {
			name = "boost " + c.Stat
		}
		switch {
		case c.Operation == "set":
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
