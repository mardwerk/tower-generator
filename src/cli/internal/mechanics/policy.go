package mechanics

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

func atLeast(value, minimum float64) bool {
	return value >= minimum || math.Abs(value-minimum) <= 1e-12*math.Max(1, minimum)
}

// Sustained is an effect's combined magnitude when an attack reapplies it
// every interval. Resetting stacks build up to the limit once the duration
// covers the interval; independent stacks are limited by duration over
// interval; extending adds duration, not magnitude. The stacked cap, when
// set, bounds the total. With one stack this is magnitude × uptime.
func Sustained(effect StatusEffect, status StatusApplication, interval float64) float64 {
	ratio := status.Seconds / interval
	stacks := float64(effect.Stacking.MaxStacks)
	var total float64
	switch {
	case stacks <= 1 || effect.Stacking.Refresh == RefreshExtend:
		total = status.Strength() * math.Min(1, ratio)
	case effect.Stacking.Refresh == RefreshReset && ratio >= 1:
		total = status.Strength() * stacks
	default:
		total = status.Strength() * math.Min(stacks, ratio)
	}
	if limit := effect.Stacking.MaxMagnitude; limit != nil {
		total = math.Min(total, *limit)
	}
	return total
}

// damageOverTime is the sustained damage per second of a version 2
// attack's damage-over-time effects.
func damageOverTime(attack Attack, vocabulary *Vocabulary) float64 {
	total := 0.0
	if vocabulary == nil {
		return total
	}
	for _, status := range attack.AppliedStatuses() {
		if effect, ok := vocabulary.Effect(status.Effect); ok && effect.Kind == KindDamageOverTime {
			total += Sustained(effect, status, attack.Stats.IntervalSeconds)
		}
	}
	return total
}

func directDamage(attack Attack, vocabulary *Vocabulary) float64 {
	st := attack.Stats
	count := st.Projectiles
	if attack.Distribution == "distinct-targets" {
		count = 1
	}
	if attack.IsV2() {
		return (st.Damage*count)/st.IntervalSeconds + damageOverTime(attack, vocabulary)
	}
	return (st.Damage*count)/st.IntervalSeconds + st.BurnDamagePerSecond*math.Min(1, st.BurnSeconds/st.IntervalSeconds)
}

func groupDamage(attack Attack, vocabulary *Vocabulary) float64 {
	st := attack.Stats
	spread := 1.0
	if attack.Distribution == "distinct-targets" {
		spread = st.Projectiles
	}
	direct := directDamage(attack, vocabulary) * st.Pierce * spread
	secondary := 0.0
	if f := attack.FollowUp; f != nil {
		inherited := 0.0
		if f.InheritStatuses {
			if attack.IsV2() {
				inherited = damageOverTime(attack, vocabulary)
			} else {
				inherited = st.BurnDamagePerSecond * math.Min(1, st.BurnSeconds/st.IntervalSeconds)
			}
		}
		secondary = f.Count * ((st.Damage*f.DamageMultiplier)/st.IntervalSeconds + inherited)
	}
	return direct + secondary
}

// SpecialtyMetricsWith measures a build with its Definition's vocabulary.
// Control counts every movement, disable and damage-taken effect of the
// vocabulary, zero when the attack does not apply it.
func SpecialtyMetricsWith(build Build, specialization string, vocabulary *Vocabulary) *s.Object {
	st := build.BaseAttack.Stats
	out := s.NewObject()
	switch specialization {
	case "direct-damage":
		out.Set("direct damage rate", directDamage(build.BaseAttack, vocabulary))
	case "group-damage":
		out.Set("group damage rate upper bound", groupDamage(build.BaseAttack, vocabulary))
	case "attack-speed":
		out.Set("attacks per second", 1/st.IntervalSeconds)
	case "range":
		out.Set("range", st.Range)
	case "control":
		if !build.BaseAttack.IsV2() {
			out.Set("slow coverage upper bound", st.SlowPercent*math.Min(1, st.SlowSeconds/st.IntervalSeconds)*st.Pierce)
			out.Set("stun coverage upper bound", math.Min(1, st.StunSeconds/st.IntervalSeconds)*st.Pierce)
			break
		}
		if vocabulary == nil {
			break
		}
		for _, effect := range vocabulary.StatusEffects {
			if effect.Kind != KindMoveSpeed && effect.Kind != KindDisable && effect.Kind != KindDamageTaken && effect.Kind != KindKnockback {
				continue
			}
			coverage := 0.0
			if status, ok := build.BaseAttack.Status(effect.ID); ok {
				if effect.Kind == KindDisable {
					coverage = math.Min(1, status.Seconds/st.IntervalSeconds)
				} else {
					coverage = Sustained(effect, status, st.IntervalSeconds)
				}
			}
			out.Set(effect.ID+" coverage upper bound", coverage*st.Pierce)
		}
	case "ability-burst":
		if len(build.Abilities) == 0 {
			return out
		}
		a := build.Abilities[0]
		out.Set("active peak direct damage rate", directDamage(a.BoostedAttack, vocabulary))
		out.Set("active peak group damage rate upper bound", groupDamage(a.BoostedAttack, vocabulary))
		out.Set("active duty fraction", math.Min(1, a.DurationSeconds/a.CooldownSeconds))
	}
	return out
}

// TimeAveraged is a damage rate averaged over one Active Ability cycle when
// the Active is used whenever it is ready: the boosted rate for the duty
// fraction of the cooldown and the ordinary rate for the rest. duty is the
// "active duty fraction", min(1, duration / cooldown), so this is
// (duration × peak + (cooldown - duration) × ordinary) / cooldown. With no
// Active, duty and peak are zero and the result is the ordinary rate.
func TimeAveraged(ordinary, peak, duty float64) float64 {
	return peak*duty + ordinary*(1-duty)
}

func attackBehavior(attack Attack) []any {
	distribution := attack.Distribution
	if distribution == "" {
		distribution = "same-primary"
	}
	var follow any
	if f := attack.FollowUp; f != nil {
		follow = []any{f.Count, f.DamageMultiplier, f.Radius, f.InheritStatuses}
	}
	if attack.IsV2() {
		out := []any{attack.Delivery, attack.DamageType, attack.Targeting, s.FromGoValue(attack.DetectionTraits()), distribution, follow}
		for _, stat := range CoreStatKeys {
			out = append(out, attack.Stats.Get(stat))
		}
		out = append(out, s.FromGoValue(attack.AppliedStatuses()))
		// Only an attack with bonus damage lists it, so earlier attacks keep
		// their behavior value.
		if len(attack.BonusDamage) > 0 {
			out = append(out, s.FromGoValue(attack.BonusDamage))
		}
		return out
	}
	out := []any{attack.Delivery, attack.DamageType, attack.Targeting, attack.Camo, distribution, follow}
	for _, stat := range StatKeys {
		out = append(out, attack.Stats.Get(stat))
	}
	return out
}

// volleyDistribution is an attack's distribution, same-primary when unset.
func volleyDistribution(attack Attack) string {
	if attack.Distribution == "" {
		return "same-primary"
	}
	return attack.Distribution
}

// HasBehaviorTransition reports a new attack shape or capability.
func HasBehaviorTransition(before, after Build) bool {
	a, b := before.BaseAttack, after.BaseAttack
	// A targeting priority is the player's choice, not a behavior, and
	// distinct targets change nothing while the attack fires one projectile.
	if a.Delivery != b.Delivery || (volleyDistribution(a) != volleyDistribution(b) && b.Stats.Projectiles > 1) {
		return true
	}
	if a.FollowUp == nil && b.FollowUp != nil {
		return true
	}
	// New access is a behavior too: a damage type that hurts other enemies,
	// or a newly detected trait.
	if a.DamageType != b.DamageType {
		return true
	}
	for _, trait := range b.DetectionTraits() {
		if !a.DetectsTrait(trait) {
			return true
		}
	}
	if a.Stats.Projectiles == 1 && b.Stats.Projectiles > 1 {
		return true
	}
	for _, key := range []string{"splashRadius", "slowPercent", "burnDamagePerSecond", "stunSeconds"} {
		if a.Stats.Get(key) == 0 && b.Stats.Get(key) > 0 {
			return true
		}
	}
	if b.IsV2() {
		for _, status := range b.AppliedStatuses() {
			if _, ok := a.Status(status.Effect); !ok {
				return true
			}
		}
		// Bonus damage against a new enemy property is new access, like a
		// new status.
		for _, bonus := range b.BonusDamage {
			if a.Bonus(bonus.Property) == 0 {
				return true
			}
		}
	}
	for _, ability := range after.Abilities {
		var prior *ResolvedAbility
		for i := range before.Abilities {
			if before.Abilities[i].Path == ability.Path {
				prior = &before.Abilities[i]
				break
			}
		}
		if (prior == nil || prior.BoostedAttack.FollowUp == nil) && ability.BoostedAttack.FollowUp != nil {
			return true
		}
	}
	return false
}

func policyBehavior(build Build) string {
	abilities := []any{}
	for _, a := range build.Abilities {
		abilities = append(abilities, []any{a.DurationSeconds, a.CooldownSeconds, attackBehavior(a.BoostedAttack)})
	}
	return s.Stringify([]any{attackBehavior(build.BaseAttack), abilities})
}

func pureBuild(blueprint *Blueprint, pathIndex, tier int) Build {
	selection := Selection{}
	selection[pathIndex] = tier
	return ResolveUnchecked(blueprint, selection)
}

// toPrecision4 is Number(x.toPrecision(4)) rendered as a JavaScript number.
func toPrecision4(x float64) string {
	rounded, _ := strconv.ParseFloat(strconv.FormatFloat(x, 'g', 4, 64), 64)
	return s.FormatNumber(rounded)
}

// missesBehavior reports a third or fifth purchase that a behavior rule
// covers and whose pure build adds no supported behavior or access over the
// purchase before it.
func missesBehavior(blueprint *Blueprint, policy *DesignPolicy, pathIndex, tier int) bool {
	return policy.RequiresBehaviorChange(tier) && !HasBehaviorTransition(pureBuild(blueprint, pathIndex, tier-1), pureBuild(blueprint, pathIndex, tier))
}

// ProposedCapabilityGaps lists each third or fifth purchase that a behavior
// rule covers, that adds no supported behavior or access and whose only new
// capability is a proposed mechanic. No build grants a proposed mechanic, so
// the purchase adds only larger numbers in play until the Definition
// supports it: a design gap the caller reports as unresolved, not a
// failure. It assumes a blueprint whose legal builds resolve.
func ProposedCapabilityGaps(blueprint *Blueprint, definition Definition) []Issue {
	policy := definition.Profile.DesignPolicy
	if policy == nil {
		return nil
	}
	var gaps []Issue
	for index, path := range PathKeys {
		for _, tier := range []int{3, 5} {
			proposed := blueprint.Paths.At(index).Tiers.At(tier).ProposedMechanics
			if len(proposed) == 0 || !missesBehavior(blueprint, policy, index, tier) {
				continue
			}
			names := make([]string, len(proposed))
			for i, p := range proposed {
				names[i] = p.Name
			}
			code := BuildCode(index, tier)
			gaps = append(gaps, Issue{fmt.Sprintf("paths.%s.tiers.tier%d", path, tier), fmt.Sprintf("Resolved %s adds no supported behavior or access over %s. Its new capability, %s, is proposed and not yet playable: until the Definition supports it, no build grants it and %s adds only larger numbers in play. %s", code, BuildCode(index, tier-1), joinAnd(names), code, BehaviorChangeRule(tier))})
		}
	}
	return gaps
}

// DesignPolicyIssues applies the Definition's optional authoring gates.
func DesignPolicyIssues(blueprint *Blueprint, definition Definition) []Issue {
	policy := definition.Profile.DesignPolicy
	if policy == nil {
		return nil
	}
	var issues []Issue
	specializations := map[string]string{}
	firstUpgrades := map[string]int{}
	capstones := map[string]int{}
	var activePaths []string
	activeRule := ActiveAbilityRule(&definition)
	for index, path := range PathKeys {
		branch := blueprint.Paths.At(index)
		prefix := "paths." + path
		specialization := branch.Specialization
		if specialization == "" {
			issues = append(issues, Issue{prefix + ".specialization", "The design policy requires an explicit path specialization."})
		} else if previous, ok := specializations[specialization]; ok && policy.DistinctPathSpecializations {
			issues = append(issues, Issue{prefix + ".specialization", fmt.Sprintf("The %s path already declares %s. %s Choose another specialization for the %s path.", PathPosition(previous), specialization, DistinctSpecializationsRule, PathPosition(path))})
		} else if !ok {
			specializations[specialization] = path
		}
		tier4 := pureBuild(blueprint, index, 4)
		tier5 := pureBuild(blueprint, index, 5)
		for _, tier := range []int{3, 5} {
			// A purchase whose only new capability is a proposed mechanic
			// is a design gap that ProposedCapabilityGaps reports, not a
			// failure; one with neither fails.
			if missesBehavior(blueprint, policy, index, tier) && len(branch.Tiers.At(tier).ProposedMechanics) == 0 {
				issues = append(issues, Issue{fmt.Sprintf("%s.tiers.tier%d", prefix, tier), fmt.Sprintf("Resolved %s adds no new behavior or access over %s and carries no proposed mechanic. %s", BuildCode(index, tier), BuildCode(index, tier-1), BehaviorChangeRule(tier))})
			}
		}
		for _, check := range []struct {
			enabled bool
			tier    int
			rule    string
			build   Build
			seen    map[string]int
		}{{policy.DistinctFirstUpgrades, 1, DistinctFirstPurchasesRule, pureBuild(blueprint, index, 1), firstUpgrades}, {policy.DistinctCapstones, 5, DistinctCapstonesRule, tier5, capstones}} {
			if !check.enabled {
				continue
			}
			signature := policyBehavior(check.build)
			if previous, ok := check.seen[signature]; ok {
				code := BuildCode(index, check.tier)
				issues = append(issues, Issue{fmt.Sprintf("%s.tiers.tier%d", prefix, check.tier), fmt.Sprintf("Resolved %s behaves exactly like %s. %s Change what %s improves, or by how much, so the resolved attack differs.", code, BuildCode(previous, check.tier), check.rule, code)})
			} else {
				check.seen[signature] = index
			}
		}
		if len(tier4.Abilities) > 0 {
			code := BuildCode(index, 4)
			manual := policy.ManualAbilityPath
			if manual.Present && (manual.Null || manual.Value != path) {
				issues = append(issues, Issue{prefix + ".tiers.tier4", fmt.Sprintf("Resolved %s unlocks an Active Ability. %s Set this path's unlockBoost to null and its boostChanges to empty.", code, activeRule)})
			}
			activePaths = append(activePaths, code)
			if len(activePaths) > policy.MaxManualAbilityPaths {
				facts := fmt.Sprintf("Resolved %s unlocks an Active Ability.", code)
				fix := "Set every unlockBoost to null and every boostChanges to empty."
				if len(activePaths) > 1 {
					facts = fmt.Sprintf("Resolved %s each unlock an Active Ability.", joinAnd(activePaths))
				}
				if policy.MaxManualAbilityPaths > 0 {
					fix = fmt.Sprintf("Keep it on at most %s; on the others set unlockBoost to null and boostChanges to empty.", pathCount(policy.MaxManualAbilityPaths))
				}
				issues = append(issues, Issue{prefix + ".tiers.tier4", facts + " " + activeRule + " " + fix})
			}
		}
		if specialization == "" || policy.MinTier5SpecialtyMultiplier == nil {
			continue
		}
		minimum := *policy.MinTier5SpecialtyMultiplier
		vocabulary := definition.Terms()
		before := SpecialtyMetricsWith(tier4, specialization, &vocabulary)
		after := SpecialtyMetricsWith(tier5, specialization, &vocabulary)
		type ratio struct {
			metric string
			value  float64
		}
		var ratios, peaks []ratio
		for _, metric := range before.Keys() {
			v, _ := before.Get(metric)
			value := v.(float64)
			n, ok := after.Get(metric)
			if !ok {
				continue
			}
			next := n.(float64)
			r := next / value
			if finite(value) && value > 0 && finite(next) && finite(r) {
				ratios = append(ratios, ratio{metric, r})
				if strings.HasPrefix(metric, "active peak") {
					peaks = append(peaks, ratio{metric, r})
				}
			}
		}
		improved := false
		for _, r := range ratios {
			peaksHold := len(peaks) > 0
			for _, p := range peaks {
				if !atLeast(p.value, 1) {
					peaksHold = false
				}
			}
			if atLeast(r.value, minimum) && (r.metric != "active duty fraction" || peaksHold) {
				improved = true
			}
		}
		if !improved {
			facts := fmt.Sprintf("Resolved %s establishes no finite positive %s specialty metric.", BuildCode(index, 4), specialization)
			if len(ratios) > 0 {
				parts := make([]string, len(ratios))
				for i, r := range ratios {
					parts[i] = fmt.Sprintf("%s %sx", r.metric, toPrecision4(r.value))
				}
				facts = fmt.Sprintf("Resolved %s multiplies the %s specialty metrics of %s by these ratios: %s.", BuildCode(index, 5), specialization, BuildCode(index, 4), joinAnd(parts))
			}
			issues = append(issues, Issue{prefix + ".tiers.tier5", facts + " " + CapstoneMultiplierRule(minimum) + " These are capacity heuristics, including group/control upper bounds, not simulated combat power or a universal BTD6 balance rule."})
		}
	}
	return issues
}

func finite(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }

// BonusMetric names the direct damage rate against an enemy property that
// accepts bonus damage, such as "direct damage rate against Hardened".
func BonusMetric(vocabulary *Vocabulary, property string) string {
	return "direct damage rate against " + vocabulary.PropertyName(property)
}

// directDamageAgainst is the direct damage rate against one enemy with a
// property: every primary hit adds the attack's bonus against it. An enemy
// the damage type cannot hurt takes nothing, and a damage-over-time effect
// it is immune to adds nothing.
func directDamageAgainst(attack Attack, vocabulary *Vocabulary, property string) float64 {
	if !vocabulary.CanDamage(attack.DamageType, property) {
		return 0
	}
	st := attack.Stats
	count := st.Projectiles
	if attack.Distribution == "distinct-targets" {
		count = 1
	}
	total := ((st.Damage + attack.Bonus(property)) * count) / st.IntervalSeconds
	for _, status := range attack.AppliedStatuses() {
		if effect, ok := vocabulary.Effect(status.Effect); ok && effect.Kind == KindDamageOverTime && !slices.Contains(effect.Immune, property) {
			total += Sustained(effect, status, st.IntervalSeconds)
		}
	}
	return total
}

// PurchaseMetricsWith measures a build with its Definition's vocabulary.
// After the specialty metrics it lists the direct damage rate against each
// property that accepts bonus damage; the other metrics ignore bonus damage.
func PurchaseMetricsWith(build Build, vocabulary *Vocabulary) *s.Object {
	out := s.NewObject()
	parts := []string{"direct-damage", "group-damage", "control", "range", "attack-speed"}
	if len(build.Abilities) > 0 {
		parts = append(parts, "ability-burst")
	}
	for _, part := range parts {
		m := SpecialtyMetricsWith(build, part, vocabulary)
		for _, key := range m.Keys() {
			v, _ := m.Get(key)
			if f := v.(float64); finite(f) {
				out.Set(key, f)
			} else {
				out.Set(key, nil)
			}
		}
	}
	if vocabulary != nil && build.BaseAttack.IsV2() {
		for _, property := range vocabulary.BonusDamageProperties {
			if f := directDamageAgainst(build.BaseAttack, vocabulary, property); finite(f) {
				out.Set(BonusMetric(vocabulary, property), f)
			} else {
				out.Set(BonusMetric(vocabulary, property), nil)
			}
		}
	}
	return out
}

// CompareCapstonePurchasesWith compares capstones with a vocabulary.
func CompareCapstonePurchasesWith(blueprint *Blueprint, vocabulary *Vocabulary) []any {
	var out []any
	for index, path := range PathKeys {
		selection := Selection{}
		selection[index] = 4
		before := ResolveUnchecked(blueprint, selection)
		selection[index] = 5
		after := ResolveUnchecked(blueprint, selection)
		ratio := math.NaN()
		if before.CumulativeCost > 0 {
			ratio = after.CumulativeCost / before.CumulativeCost
		}
		var count any
		if finite(ratio) {
			count = math.Floor(ratio)
		}
		metrics := PurchaseMetricsWith(before, vocabulary)
		bounds := s.NewObject()
		if count != nil {
			for _, key := range []string{"direct damage rate", "group damage rate upper bound"} {
				value, _ := metrics.Get(key)
				product := math.NaN()
				if f, ok := value.(float64); ok {
					product = f * count.(float64)
				}
				if finite(product) {
					bounds.Set(key, product)
				} else {
					bounds.Set(key, nil)
				}
			}
		}
		out = append(out, s.NewObject().
			Set("path", path).
			Set("tier4", s.NewObject().Set("totalGold", before.CumulativeCost).Set("metrics", metrics)).
			Set("tier5", s.NewObject().Set("totalGold", after.CumulativeCost).Set("metrics", PurchaseMetricsWith(after, vocabulary))).
			Set("tier4CopiesAtTier5Budget", count).
			Set("sameBudgetTier4Copies", s.NewObject().Set("count", count).Set("additiveThroughputUpperBounds", bounds).Set("perCopyMetrics", metrics)).
			Set("assumption", "Ideal sustained access to eligible targets. Group/control metrics are capacity upper bounds; active peaks are not sustained output. Copy throughput assumes independent target access and extra placement space; range, attack frequency, control coverage and active duty are per-copy values, not summed. A zero-cost tier-four build has no finite budget-limited copy count. Null metrics or counts are unavailable because the calculation has no finite numeric result. No waves, buffs, geometry, actual crowd density or balance are simulated."))
	}
	return out
}
