package mechanics

import (
	"fmt"
	"slices"
	"strings"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Rule sentences of the design policy. Each designPolicy field's rule is
// stated once here. The plan prompt, the draft and repair guidance and every
// plan, promise and resolved failure of that field use the same sentence: a
// failure gives its path-specific facts, then the sentence, then the fix. A
// prompt therefore cannot permit what a check rejects.

// DistinctSpecializationsRule states distinctPathSpecializations.
const DistinctSpecializationsRule = "The three paths declare three different specializations."

// DistinctFirstPurchasesRule states distinctFirstUpgrades.
const DistinctFirstPurchasesRule = "No two first purchases (1-x-x, x-1-x, x-x-1) may produce the same resolved attack. Names, prices and different arithmetic expressions do not make identical effects different."

// DistinctCapstonesRule states distinctCapstones.
const DistinctCapstonesRule = "Pure 5-0-0, 0-5-0 and 0-0-5 builds must differ mechanically even after ignoring names and costs."

// EarlyBenefitsRule states distinctEarlyBenefits.
const EarlyBenefitsRule = "Two paths' first two purchases together must differ in what they improve or unlock, whatever their order, names, prices or amounts; lowers do not count."

// ExclusiveEarlyBenefitsRule states exclusiveEarlyBenefits, which subsumes
// distinctEarlyBenefits: two paths whose early benefits share nothing cannot
// have the same early benefits.
const ExclusiveEarlyBenefitsRule = "No dimension or capability that one path's first two purchases improve or unlock may be improved or unlocked by another path's first two purchases, whatever their order, names, prices or amounts; a path may repeat its own, any path's third to fifth purchases may improve it, and lowers do not count."

// NoDefinitionActiveAbilityRule stands for ActiveAbilityRule when a request
// carries no Definition.
const NoDefinitionActiveAbilityRule = "An Active Ability exists only where the Definition allows one."

var ordinals = []string{"", "first", "second", "third", "fourth", "fifth"}

// BuildCode names a purchase in top-middle-bottom notation: the second path
// at tier 4 is x-4-x.
func BuildCode(pathIndex, tier int) string {
	parts := []string{"x", "x", "x"}
	parts[pathIndex] = string(rune('0' + tier))
	return strings.Join(parts, "-")
}

// PathPosition names a path key's position in build codes: path2 is middle.
func PathPosition(path string) string {
	switch path {
	case "path1":
		return "top"
	case "path2":
		return "middle"
	case "path3":
		return "bottom"
	}
	return path
}

// ActiveAbilityAllowed reports whether a design policy lets a path unlock an
// Active Ability. Without a policy every path may.
func ActiveAbilityAllowed(policy *DesignPolicy, path string) bool {
	if policy == nil {
		return true
	}
	if policy.MaxManualAbilityPaths == 0 || policy.ManualAbilityPath.Null {
		return false
	}
	return !policy.ManualAbilityPath.Present || policy.ManualAbilityPath.Value == path
}

// ForbidsActiveAbility reports a policy under which no path may unlock an
// Active Ability: a null manualAbilityPath or maxManualAbilityPaths 0.
func (p *DesignPolicy) ForbidsActiveAbility() bool {
	return p != nil && (p.ManualAbilityPath.Null || p.MaxManualAbilityPaths == 0)
}

// PreservesEarlyIdentity reports whether preserveEarlyAttackIdentity is on.
func (p *DesignPolicy) PreservesEarlyIdentity() bool {
	return p != nil && p.PreserveEarlyAttackIdentity != nil && *p.PreserveEarlyAttackIdentity
}

// ExcludesSharedEarlyBenefits reports whether exclusiveEarlyBenefits is on.
func (p *DesignPolicy) ExcludesSharedEarlyBenefits() bool {
	return p != nil && p.ExclusiveEarlyBenefits != nil && *p.ExclusiveEarlyBenefits
}

// RequiresBehaviorChange reports whether the policy requires the third or
// fifth purchase to add a behavior or access.
func (p *DesignPolicy) RequiresBehaviorChange(tier int) bool {
	if p == nil {
		return false
	}
	required := p.RequireTier3BehaviorChange
	if tier == 5 {
		required = p.RequireTier5BehaviorChange
	}
	return required != nil && *required
}

// ActiveAbilityRule states manualAbilityPath and maxManualAbilityPaths as
// the Engine applies them: the plan schema, the plan check and the resolved
// check permit an Active Ability on exactly the paths it names.
func ActiveAbilityRule(d *Definition) string {
	if d == nil {
		return NoDefinitionActiveAbilityRule
	}
	tier := d.Rules.ManualBoostUnlockTier
	policy := d.Profile.DesignPolicy
	switch {
	case policy == nil:
		return fmt.Sprintf("An Active Ability may unlock at a path's %s purchase.", ordinals[tier])
	case policy.ForbidsActiveAbility():
		return "No path may have an Active Ability; every path stays automatic."
	case policy.ManualAbilityPath.Present:
		path := policy.ManualAbilityPath.Value
		return fmt.Sprintf("Only the %s path may have an Active Ability, first at %s; the other paths stay automatic.", PathPosition(path), BuildCode(slices.Index(PathKeys, path), tier))
	}
	return fmt.Sprintf("At most %s may have an Active Ability, first at a path's %s purchase; automatic paths are complete designs.", pathCount(policy.MaxManualAbilityPaths), ordinals[tier])
}

// EarlyIdentityRule states preserveEarlyAttackIdentity as the plan schema,
// the plan checks and the resolved check enforce it. A Definition with bonus
// damage names it too: an early purchase may raise an existing bonus but not
// add one.
func EarlyIdentityRule(d *Definition) string {
	detection, bonus := "personal detection", ""
	if d == nil || !d.IsV2() {
		detection = "personal Camo detection"
	} else if len(d.Vocabulary.BonusDamageProperties) > 0 {
		bonus = "bonus damage, "
	}
	return "The first and second purchase of each path keep the base attack's form: they add no new status, " + bonus + "splash, follow-up or distinct-target volley, change no delivery, targeting or damage type, and keep a single-projectile attack single. Improving existing stats and effects and adding " + detection + " remain allowed."
}

// BehaviorChangeRule states requireTier3BehaviorChange (tier 3) and
// requireTier5BehaviorChange (tier 5) as HasBehaviorTransition judges them.
// Larger numbers of the Active Ability do not count. A proposed mechanic
// grants no behavior, so it does not meet the rule; a purchase whose only
// new capability is proposed is an unresolved design gap, not a failure
// (ProposedCapabilityGaps).
func BehaviorChangeRule(tier int) string {
	return "The " + ordinals[tier] + " purchase of every path must add a supported behavior or access: a new delivery, a distinct-target volley of more than one projectile, more than one projectile, splash, a status effect, a bounded follow-up, a new damage type, a newly detected trait or new bonus damage against an eligible enemy property; larger existing numbers, the Active Ability's included, a targeting change, a new name or a change with no effect do not count. A proposed mechanic grants no behavior in any build: a purchase whose only new capability is proposed is an unresolved design gap, not playable until the Definition supports it."
}

// CapstoneMultiplierRule states minTier5SpecialtyMultiplier.
func CapstoneMultiplierRule(minimum float64) string {
	return "The fifth purchase of every path must improve an established specialty metric of the path's pure fourth purchase by at least " + s.FormatNumber(minimum) + " times; a gain in active duty fraction alone must also keep the active peak."
}

func pathCount(n int) string {
	if n == 1 {
		return "1 path"
	}
	return fmt.Sprintf("%d paths", n)
}

func joinAnd(items []string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return items[0]
	}
	return strings.Join(items[:len(items)-1], ", ") + " and " + items[len(items)-1]
}
