package unit

import (
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// PurchaseAdaptation says which source technique a purchase adapts and what
// its typed changes make of the attack: "Adapts Gear 3: the attack becomes
// slower and heavier and now hits enemies around the target." It is derived
// when read, from the retained plan's technique and the resolved purchase on
// its own path, so it names only behavior the Unit has and implies no
// delivery the typed changes lack (reported on #27: a Luffy review failed
// purchases that listed stat changes without the Gear 3 and Gear 4 they
// adapt). It is empty when the plan names no technique for the purchase, or
// only the base attack, which the purchase's effects already describe, or
// when the purchase changes nothing it can describe.
func PurchaseAdaptation(blueprint *m.Blueprint, plan *DesignPlan, definition m.Definition, pathIndex, tier int) string {
	if blueprint == nil || plan == nil || plan.UpgradeIntents == nil {
		return ""
	}
	technique := strings.TrimSpace(plan.UpgradeIntents.At(pathIndex).At(tier).Technique)
	if technique == "" || strings.EqualFold(technique, strings.TrimSpace(plan.Base.Name)) {
		return ""
	}
	var before, after m.Selection
	before[pathIndex], after[pathIndex] = tier-1, tier
	vocabulary := definition.Terms()
	parts := adaptationClauses(m.ResolveUnchecked(blueprint, before), m.ResolveUnchecked(blueprint, after), &vocabulary)
	if len(parts) == 0 {
		return ""
	}
	return "Adapts " + technique + ": " + joinAnd(parts) + "."
}

// deliveryNouns name a delivery as the attack it becomes.
var deliveryNouns = map[string]string{"projectile": "a projectile", "instant": "an instant strike", "area": "an area attack", "beam": "a beam"}

// adaptationClauses describes, in words, what one purchase changes: one
// part about the attack, then one per Active Ability it adds or improves.
func adaptationClauses(before, after m.Build, v *m.Vocabulary) []string {
	a, b := before.BaseAttack, after.BaseAttack
	sa, sb := a.Stats, b.Stats
	var becomes []string
	switch {
	case sb.IntervalSeconds > sa.IntervalSeconds:
		becomes = append(becomes, "slower")
	case sb.IntervalSeconds < sa.IntervalSeconds:
		becomes = append(becomes, "faster")
	}
	switch {
	case sb.Damage > sa.Damage:
		becomes = append(becomes, "heavier")
	case sb.Damage < sa.Damage:
		becomes = append(becomes, "lighter")
	}
	switch {
	case sb.Range > sa.Range:
		becomes = append(becomes, "longer-reaching")
	case sb.Range < sa.Range:
		becomes = append(becomes, "shorter-reaching")
	}
	var clauses []string
	if len(becomes) > 0 {
		clauses = append(clauses, "becomes "+joinAnd(becomes))
	}
	if a.Delivery != b.Delivery {
		noun := deliveryNouns[b.Delivery]
		if noun == "" {
			noun = b.Delivery
		}
		clauses = append(clauses, "is delivered as "+noun)
	}
	switch {
	case sa.SplashRadius <= 0 && sb.SplashRadius > 0:
		clauses = append(clauses, "now hits enemies around the target")
	case sb.SplashRadius > sa.SplashRadius:
		clauses = append(clauses, "hits a wider area around the target")
	}
	if sb.Projectiles > sa.Projectiles {
		clauses = append(clauses, "fires "+num(sb.Projectiles)+" shots per attack")
	}
	if b.Distribution == "distinct-targets" && a.Distribution != b.Distribution && sb.Projectiles > 1 {
		clauses = append(clauses, "spreads its shots across separate enemies")
	}
	if sb.Pierce > sa.Pierce {
		clauses = append(clauses, "can hit up to "+num(sb.Pierce)+" enemies per shot")
	}
	switch f, g := a.FollowUp, b.FollowUp; {
	case f == nil && g != nil:
		clauses = append(clauses, "is followed by a strike on up to "+num(g.Count)+" nearby enemies")
	case f != nil && g != nil && (g.Count > f.Count || g.DamageMultiplier > f.DamageMultiplier || g.Radius > f.Radius || g.InheritStatuses && !f.InheritStatuses):
		clauses = append(clauses, "gets a stronger follow-up strike")
	}
	for _, status := range b.AppliedStatuses() {
		name := effectOf(status.Effect, v).Name
		prior, had := a.Status(status.Effect)
		switch {
		case !had:
			clauses = append(clauses, "applies "+name)
		case status.Strength() > prior.Strength() || status.Seconds > prior.Seconds:
			clauses = append(clauses, "applies stronger "+name)
		}
	}
	for _, trait := range b.DetectionTraits() {
		if !a.DetectsTrait(trait) {
			clauses = append(clauses, "can target "+termName(trait, v.Detection)+" enemies")
		}
	}
	if a.DamageType != b.DamageType {
		name := b.DamageType
		if t, ok := v.DamageType(b.DamageType); ok && t.Name != "" {
			name = t.Name
		}
		clauses = append(clauses, "deals "+name+" damage")
	}
	if a.Targeting != b.Targeting {
		clauses = append(clauses, "targets the "+termName(b.Targeting, v.Targeting)+" enemy")
	}
	var parts []string
	if len(clauses) > 0 {
		parts = append(parts, "the attack "+joinAnd(clauses))
	}
	owned := map[string]m.ResolvedAbility{}
	for _, ability := range before.Abilities {
		owned[ability.Name] = ability
	}
	for _, ability := range after.Abilities {
		prior, had := owned[ability.Name]
		switch {
		case !had:
			parts = append(parts, "the Unit gains the "+ability.Name+" Active Ability")
		case ability.DurationSeconds > prior.DurationSeconds || ability.CooldownSeconds < prior.CooldownSeconds ||
			ability.DamageMultiplier > prior.DamageMultiplier || ability.IntervalMultiplier < prior.IntervalMultiplier || ability.RangeBonus > prior.RangeBonus:
			parts = append(parts, ability.Name+" becomes stronger")
		}
	}
	return parts
}

// joinAnd joins phrases as "a", "a and b" or "a, b and c".
func joinAnd(parts []string) string {
	if len(parts) < 2 {
		return strings.Join(parts, "")
	}
	return strings.Join(parts[:len(parts)-1], ", ") + " and " + parts[len(parts)-1]
}
