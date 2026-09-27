package render

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// Purchase sentences describe resolved mechanics for readers: each change a
// purchase makes, with the exact numbers of its pure path. They add no rule
// of their own; the blueprint and its Definition decide every value.

// pathPositions name the paths in build-code order.
var pathPositions = []string{"Top", "Middle", "Bottom"}

// Purchase is one bought upgrade as the unit sheet shows it.
type Purchase struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
	// Technique is the technique the retained plan maps the purchase to, and
	// BaseTechnique whether it is the base attack; see unit.PurchaseTechnique.
	// Both are plan intent, shown apart from the resolved Effects.
	Technique     string `json:"technique,omitempty"`
	BaseTechnique bool   `json:"baseTechnique,omitempty"`
	// Adaptation is the plan's description of how the purchase adapts its
	// source technique, a claim the review checks; see
	// unit.PurchaseAdaptation.
	Adaptation string `json:"adaptation,omitempty"`
	// NameMentions are other repertoire entries the purchase's name points
	// to, an advisory flag for review; see unit.PurchaseNameMentions.
	NameMentions []string `json:"nameMentions,omitempty"`
	Effects      []string `json:"effects"`
	// ProposedMechanics are what the purchase needs that its Definition
	// cannot express yet. No build grants them; they follow the resolved
	// effects, labeled as proposed.
	ProposedMechanics []m.ProposedMechanic `json:"proposedMechanics,omitempty"`
	// Text is Passage, derived when read, so every view words the purchase
	// the same way.
	Text string `json:"text"`
}

// Passage is the purchase as one passage. With a planned technique it labels
// the plan's intent apart from the resolved effects: "Plan: adapts Gear 3.
// {adaptation} Resolved: {effects}". Without one, as for Results made before
// plans typed techniques, it is the effects alone. Each proposed mechanic
// follows the effects as "Proposed (not yet supported): {name}: {effect}".
func (p Purchase) Passage() string {
	effects := strings.Join(p.Effects, " ") + p.proposedPassage()
	if p.Technique == "" {
		return strings.TrimSpace(p.Adaptation + " " + effects)
	}
	plan := "Plan: adapts " + p.Technique
	if p.BaseTechnique {
		plan += " (base attack)"
	}
	plan += "."
	if p.Adaptation != "" {
		plan += " " + p.Adaptation
	}
	if len(p.NameMentions) > 0 {
		plan += " Review: the name suggests " + joinAnd(p.NameMentions) + ", which the plan does not map to this purchase."
	}
	return plan + " Resolved: " + effects
}

// proposedPassage labels each proposed mechanic as not yet supported:
// " Proposed (not yet supported): Boundman bounce: each hit bounces to a
// second enemy."
func (p Purchase) proposedPassage() string {
	var text string
	for _, proposed := range p.ProposedMechanics {
		effect := strings.TrimSpace(proposed.Effect)
		if effect != "" && !strings.HasSuffix(effect, ".") {
			effect += "."
		}
		text += " Proposed (not yet supported): " + strings.TrimSpace(proposed.Name) + ": " + effect
	}
	return text
}

// PathPurchases are a path's five purchases in order.
type PathPurchases struct {
	Position  string     `json:"position"`
	Name      string     `json:"name"`
	Purchases []Purchase `json:"purchases"`
}

// sheet resolves one blueprint under its Definition for the sentences.
type sheet struct {
	blueprint  *m.Blueprint
	definition m.Definition
	vocabulary m.Vocabulary
	currency   string
	// plan, when set, names the technique each purchase adapts.
	plan *unit.DesignPlan
}

func newSheet(blueprint *m.Blueprint, definition *m.Definition) *sheet {
	if blueprint == nil {
		return nil
	}
	d := m.DefaultDefinition()
	if definition != nil {
		d = *definition
	}
	// A reading view checks structure only; the authoring checks judge how a
	// unit is made, and AuthoringIssues reports them beside the sheet.
	if len(m.ValidateStructure(blueprint, d)) > 0 {
		return nil
	}
	return &sheet{blueprint: blueprint, definition: d, vocabulary: d.Terms(), currency: d.Profile.Currency}
}

func (sh *sheet) resolve(selection m.Selection) m.Build {
	return m.ResolveUnchecked(sh.blueprint, selection)
}

// decimal writes a stat with at most four decimals.
func decimal(value float64) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return s.FormatNumber(value)
	}
	rounded, _ := strconv.ParseFloat(strconv.FormatFloat(value, 'f', 4, 64), 64)
	return s.FormatNumber(rounded)
}

// price writes a price with thousands separators: 15000 is 15,000.
func price(value float64) string {
	text := decimal(value)
	whole, fraction, _ := strings.Cut(text, ".")
	negative := strings.HasPrefix(whole, "-")
	whole = strings.TrimPrefix(whole, "-")
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + "," + whole[i:]
	}
	if negative {
		whole = "-" + whole
	}
	if fraction != "" {
		return whole + "." + fraction
	}
	return whole
}

func (sh *sheet) money(value float64) string { return price(value) + " " + sh.currency }

func (sh *sheet) damageTypeName(id string) string {
	if damageType, ok := sh.vocabulary.DamageType(id); ok && damageType.Name != "" {
		return damageType.Name
	}
	return capitalized(id)
}

func (sh *sheet) effect(id string) m.StatusEffect {
	if effect, ok := sh.vocabulary.Effect(id); ok {
		return effect
	}
	return m.StatusEffect{ID: id, Name: capitalized(id), Stacking: m.Stacking{MaxStacks: 1}}
}

func capitalized(text string) string {
	if text == "" {
		return text
	}
	return strings.ToUpper(text[:1]) + text[1:]
}

// magnitude writes an effect's magnitude with its unit: 30% or 5 damage/s.
func (sh *sheet) magnitude(value float64, effect m.StatusEffect) string {
	return decimal(value) + magnitudeSuffix(effect)
}

// status describes an applied status: "Slow 30% for 2 s".
func (sh *sheet) status(status m.StatusApplication) string {
	effect := sh.effect(status.Effect)
	text := effect.Name + " for " + decimal(status.Seconds) + " s"
	if status.Magnitude != nil {
		text = effect.Name + " " + sh.magnitude(*status.Magnitude, effect) + " for " + decimal(status.Seconds) + " s"
	}
	if effect.Stacking.MaxStacks > 1 {
		text += fmt.Sprintf(", stacking up to %d times", effect.Stacking.MaxStacks)
	}
	if limit := effect.Stacking.MaxMagnitude; limit != nil {
		text += ", at most " + sh.magnitude(*limit, effect) + " combined"
	}
	return text
}

// bonusAgainst names bonus damage against a property: "+50 damage against
// Hardened enemies".
func (sh *sheet) bonusAgainst(damage float64, property string) string {
	return "+" + decimal(damage) + " damage against " + sh.vocabulary.PropertyName(property) + " enemies"
}

// bonusSentence says what an attack's bonus damage adds to each hit, with
// the resulting damage per hit: "Each hit deals +50 damage against Hardened
// enemies (70 per hit)." A bonus the damage type cannot deal says so.
func (sh *sheet) bonusSentence(attack m.Attack) string {
	var parts []string
	for _, bonus := range attack.BonusDamage {
		text := sh.bonusAgainst(bonus.Damage, bonus.Property) + " (" + decimal(attack.Stats.Damage+bonus.Damage) + " per hit)"
		if !sh.vocabulary.CanDamage(attack.DamageType, bonus.Property) {
			text = sh.bonusAgainst(bonus.Damage, bonus.Property) + ", which its " + sh.damageTypeName(attack.DamageType) + " damage cannot deal"
		}
		parts = append(parts, text)
	}
	if len(parts) == 0 {
		return ""
	}
	return "Each hit deals " + joinAnd(parts) + "."
}

func (sh *sheet) detectionName(trait string) string {
	return termName(trait, sh.vocabulary.Detection)
}

func (sh *sheet) targetingName(id string) string {
	return termName(id, sh.vocabulary.Targeting)
}

// immunity says which enemy properties a damage type cannot hurt.
func (sh *sheet) immunity(damageType string) string {
	kind, ok := sh.vocabulary.DamageType(damageType)
	name := sh.damageTypeName(damageType)
	if !ok || len(kind.IneffectiveAgainst) == 0 {
		return name + " damage can hurt every enemy property."
	}
	var names []string
	for _, property := range kind.IneffectiveAgainst {
		names = append(names, termName(property, sh.vocabulary.EnemyProperties))
	}
	return name + " damage cannot hurt " + joinAnd(names) + " enemies."
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

func plural(count float64, one, many string) string {
	if count == 1 {
		return one
	}
	return many
}

// shot names what one attack emits.
func shot(attack m.Attack) string {
	if attack.Delivery == "projectile" {
		return "projectile"
	}
	return "pulse"
}

// followUp describes a bounded follow-up of an attack. Its multiplier,
// shown as a percentage of the hit damage, scales the hit's ordinary damage; the attack's bonus damage is added to
// each follow-up hit unscaled.
func (sh *sheet) followUp(f m.FollowUp, attack m.Attack) string {
	inherit := "it applies no statuses"
	if f.InheritStatuses {
		inherit = "it applies the attack's purchased statuses"
	}
	return fmt.Sprintf("%s: after each volley hits, up to %s other detected %s within %s of the primary impact take %s of the hit damage (%s)%s once each; %s, never recurses and inherits no pierce, splash or volley count",
		f.Name, decimal(f.Count), plural(f.Count, "enemy", "enemies"), decimal(f.Radius), percentOfDamage(f.DamageMultiplier), decimal(attack.Stats.Damage*f.DamageMultiplier), sh.plusBonus(attack), inherit)
}

// plusBonus names the bonus damage an attack's damage type can deal, added
// to a hit whose ordinary damage a multiplier scales: ", plus +50 damage
// against Hardened enemies". It is empty without such a bonus.
func (sh *sheet) plusBonus(attack m.Attack) string {
	var parts []string
	for _, bonus := range attack.BonusDamage {
		if sh.vocabulary.CanDamage(attack.DamageType, bonus.Property) {
			parts = append(parts, sh.bonusAgainst(bonus.Damage, bonus.Property))
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return ", plus " + joinAnd(parts) + ","
}

// attackSentences describe a whole resolved attack, as the base Unit has it.
// attackSentences describes what an attack does. Defaults the reader need
// not act on stay internal: the Definition's first targeting priority, the
// clear-path delivery rule and detection the attack lacks.
func (sh *sheet) attackSentences(attack m.Attack) []string {
	st := attack.Stats
	aim := "its target"
	if attack.Distribution == "distinct-targets" {
		aim = "different detected enemies in range, its target first; unused shots are lost"
	}
	targeting := ""
	if len(sh.vocabulary.Targeting) > 0 && attack.Targeting != sh.vocabulary.Targeting[0].ID {
		targeting = " with " + sh.targetingName(attack.Targeting) + " targeting"
	}
	out := []string{fmt.Sprintf("%s is an automatic %s attack%s. Every %s s it fires %s %s at %s; each deals %s %s damage with pierce %s, hitting up to %s %s, at range %s.",
		attack.Name, attack.Delivery, targeting, decimal(st.IntervalSeconds), decimal(st.Projectiles), plural(st.Projectiles, shot(attack), shot(attack)+"s"), aim,
		decimal(st.Damage), sh.damageTypeName(attack.DamageType), decimal(st.Pierce), decimal(st.Pierce), plural(st.Pierce, "enemy", "enemies"), decimal(st.Range))}
	if st.SplashRadius > 0 {
		out = append(out, "Its splash radius is "+decimal(st.SplashRadius)+", shared within the same target cap.")
	}
	var statuses []string
	for _, status := range attack.AppliedStatuses() {
		statuses = append(statuses, sh.status(status))
	}
	if len(statuses) > 0 {
		out = append(out, "Each hit applies "+joinAnd(statuses)+".")
	}
	if bonus := sh.bonusSentence(attack); bonus != "" {
		out = append(out, bonus)
	}
	if attack.FollowUp != nil {
		out = append(out, capitalized(sh.followUp(*attack.FollowUp, attack))+".")
	}
	var detected []string
	for _, trait := range sh.vocabulary.Detection {
		if attack.DetectsTrait(trait.ID) {
			detected = append(detected, sh.detectionName(trait.ID))
		}
	}
	if len(detected) > 0 {
		out = append(out, "It detects "+joinAnd(detected)+" enemies.")
	}
	return append(out, sh.immunity(attack.DamageType))
}

var statNames = map[string]string{
	"damage":              "damage",
	"intervalSeconds":     "the attack interval",
	"range":               "range",
	"pierce":              "pierce",
	"projectiles":         "projectiles per attack",
	"splashRadius":        "splash radius",
	"slowPercent":         "Slow",
	"slowSeconds":         "Slow duration",
	"burnDamagePerSecond": "Burn damage per second",
	"burnSeconds":         "Burn duration",
	"stunSeconds":         "Stun duration",
}

func statUnit(stat string) string {
	switch stat {
	case "intervalSeconds", "slowSeconds", "burnSeconds", "stunSeconds":
		return " s"
	case "slowPercent":
		return "%"
	case "burnDamagePerSecond":
		return " damage/s"
	}
	return ""
}

// statHow is the delta of a purchase's change to one stat, written in the
// parentheses after its values. Additions read as the difference ("+1"),
// multipliers as a percentage ("+50%") and an attack interval as the attack
// speed it gives ("attacks 18% faster"). When the delta alone does not show
// what the purchase's own changes do, it follows them: additions apply
// before multipliers, so +2 under an earlier purchase's multiplier of 3
// reads "+6: its +2 scaled by an earlier purchase's +200%".
func statHow(stat string, changes []m.Change, before, after float64) string {
	unitText := statUnit(stat)
	interval := stat == "intervalSeconds"
	var sets, adds, multiplies []m.Change
	sum := 0.0
	for _, change := range changes {
		switch change.Operation {
		case "add":
			adds = append(adds, change)
			sum += change.Number
		case "multiply":
			multiplies = append(multiplies, change)
		default:
			sets = append(sets, change)
		}
	}
	onlyMultiplies := len(sets) == 0 && len(adds) == 0
	var primary string
	switch {
	case interval && onlyMultiplies:
		primary = attackSpeedChange(before, after)
	case interval:
		primary = delta(before, after, unitText)
		if speed := attackSpeedChange(before, after); speed != "" {
			primary += ", " + speed
		}
	case onlyMultiplies && before != 0:
		primary = signedPercent(after/before - 1)
	default:
		primary = delta(before, after, unitText)
	}
	if onlyMultiplies {
		return primary
	}
	if len(sets) == 0 && len(multiplies) == 0 {
		factor := (after - before) / sum
		if sum == 0 || factor <= 0 || math.Abs(factor-1) < 1e-9 {
			return primary
		}
		return primary + ": its " + signedNumber(sum) + unitText + " scaled by an earlier purchase's " + signedPercent(factor-1)
	}
	var parts []string
	for _, change := range sets {
		parts = append(parts, "sets the base value to "+decimal(change.Number)+unitText)
	}
	for _, change := range adds {
		parts = append(parts, signedNumber(change.Number)+unitText)
	}
	for _, change := range multiplies {
		if interval {
			parts = append(parts, "attacks "+activeSpeed(change.Number))
		} else {
			parts = append(parts, signedPercent(change.Number-1))
		}
	}
	how := primary + ": " + strings.Join(parts, ", then ")
	if len(sets) > 0 {
		how += "; purchased additions and multipliers still apply"
	}
	return how
}

// numberChange writes a resolved number's change as a sentence.
func numberChange(label, unitText string, before, after float64, lowerIsBetter bool, how string) string {
	verb := "Raises"
	switch {
	case before == after:
		verb = "Keeps"
	case lowerIsBetter && after < before:
		verb = "Shortens"
	case lowerIsBetter:
		verb = "Lengthens"
	case after < before:
		verb = "Lowers"
	}
	text := fmt.Sprintf("%s %s from %s%s to %s%s", verb, label, decimal(before), unitText, decimal(after), unitText)
	if verb == "Keeps" {
		text = fmt.Sprintf("Keeps %s at %s%s", label, decimal(after), unitText)
	}
	if how != "" {
		text += " (" + how + ")"
	}
	return text + "."
}

// purchaseEffects describes one purchase with the resolved values of the
// builds before and after it.
func (sh *sheet) purchaseEffects(changes []m.Change, before, after m.Build) []string {
	var order []string
	groups := map[string][]m.Change{}
	for _, change := range changes {
		key := change.Kind + "." + change.Target
		switch change.Kind {
		case "stat", "modifyBoost":
			key = change.Kind + "." + change.Stat
		case "status":
			key = change.Kind + "." + change.Effect
		case "detection":
			key = change.Kind + "." + change.Trait
		case "bonusDamage":
			key = change.Kind + "." + change.Property
		}
		if _, ok := groups[key]; !ok {
			order = append(order, key)
		}
		groups[key] = append(groups[key], change)
	}
	prior, next := before.BaseAttack, after.BaseAttack
	var out []string
	for _, key := range order {
		group := groups[key]
		change := group[len(group)-1]
		switch change.Kind {
		case "stat":
			label := statNames[change.Stat]
			if change.Stat == "projectiles" && next.Delivery != "projectile" {
				label = "pulses per attack"
			}
			was, now := prior.Stats.Get(change.Stat), next.Stats.Get(change.Stat)
			out = append(out, numberChange(label, statUnit(change.Stat), was, now, change.Stat == "intervalSeconds", statHow(change.Stat, group, was, now)))
		case "status":
			effect := sh.effect(change.Effect)
			was, had := prior.Status(change.Effect)
			now, has := next.Status(change.Effect)
			switch {
			case !had && has:
				out = append(out, "Adds "+sh.status(now)+" to each hit.")
			case had && !has:
				out = append(out, "Removes "+effect.Name+".")
			default:
				if was.Strength() != now.Strength() {
					suffix := magnitudeSuffix(effect)
					out = append(out, numberChange(effect.Name, suffix, was.Strength(), now.Strength(), false, delta(was.Strength(), now.Strength(), suffix)))
				}
				if was.Seconds != now.Seconds {
					out = append(out, numberChange(effect.Name+" duration", " s", was.Seconds, now.Seconds, false, delta(was.Seconds, now.Seconds, " s")))
				}
			}
		case "bonusDamage":
			// Bonus damage only adds, so a purchase either adds a bonus
			// against a property or raises the one the attack has.
			was, now := prior.Bonus(change.Property), next.Bonus(change.Property)
			if was == 0 {
				text := fmt.Sprintf("Adds %s (%s to %s per hit).", sh.bonusAgainst(now, change.Property), decimal(next.Stats.Damage), decimal(next.Stats.Damage+now))
				if !sh.vocabulary.CanDamage(next.DamageType, change.Property) {
					text = fmt.Sprintf("Adds %s, which its %s damage cannot deal.", sh.bonusAgainst(now, change.Property), sh.damageTypeName(next.DamageType))
				}
				out = append(out, text)
			} else {
				out = append(out, numberChange("bonus damage against "+sh.vocabulary.PropertyName(change.Property)+" enemies", "", was, now, false, delta(was, now, "")))
			}
		case "detection", "camo":
			trait := change.Trait
			if change.Kind == "camo" {
				trait = "camo"
			}
			if change.Bool {
				out = append(out, "Adds "+sh.detectionName(trait)+" detection to this Unit's attack.")
			} else {
				out = append(out, "Removes "+sh.detectionName(trait)+" detection.")
			}
		case "damageType":
			out = append(out, fmt.Sprintf("Switches damage from %s to %s. %s", sh.damageTypeName(prior.DamageType), sh.damageTypeName(next.DamageType), sh.immunity(next.DamageType)))
		case "delivery":
			out = append(out, fmt.Sprintf("Replaces %s delivery with %s delivery.", prior.Delivery, next.Delivery))
		case "targeting":
			out = append(out, fmt.Sprintf("Changes targeting from %s to %s.", sh.targetingName(prior.Targeting), sh.targetingName(next.Targeting)))
		case "distribution":
			if next.Distribution == "distinct-targets" {
				out = append(out, "Aims each projectile at a different detected enemy in range, the selected target first; unused shots are lost.")
			} else {
				out = append(out, "Aims every projectile at the selected target again.")
			}
		case "followUp":
			if change.Target == "boost" {
				name := "the Active Ability"
				if ability := newAbilityOrLast(after); ability != nil {
					name = ability.Name
				}
				out = append(out, "While "+name+" is active, adds "+sh.followUp(*change.FollowUp, next)+".")
			} else if prior.FollowUp != nil {
				out = append(out, "Replaces "+prior.FollowUp.Name+" with "+sh.followUp(*change.FollowUp, next)+".")
			} else {
				out = append(out, "Adds "+sh.followUp(*change.FollowUp, next)+".")
			}
		case "unlockBoost":
			if ability := newAbility(before, after); ability != nil {
				out = append(out, sh.abilitySentence(*ability))
			}
		case "modifyBoost":
			if ability := newAbilityOrLast(after); ability != nil {
				var old *m.ResolvedAbility
				for i := range before.Abilities {
					if before.Abilities[i].Path == ability.Path {
						old = &before.Abilities[i]
					}
				}
				if old != nil {
					out = append(out, boostChange(ability.Name, change.Stat, boostStat(old, change.Stat), boostStat(ability, change.Stat)))
				}
			}
		}
	}
	return out
}

// magnitudeSuffix follows a magnitude: "%" or " damage/s".
func magnitudeSuffix(effect m.StatusEffect) string {
	if effect.Magnitude == nil || effect.Magnitude.Unit == "percent" {
		return "%"
	}
	return " " + effect.Magnitude.Unit
}

// boostChange describes a purchase's change to an owned Active Ability. Its
// multipliers read as percentages: a damage multiplier of 2 is a +100%
// damage bonus and an interval multiplier of 0.5 attacks 100% faster, so the
// delta is in percentage points.
func boostChange(name, stat string, before, after float64) string {
	switch stat {
	case "durationSeconds":
		return numberChange(name+" duration", " s", before, after, false, delta(before, after, " s"))
	case "cooldownSeconds":
		return numberChange(name+" recharge", " s", before, after, true, delta(before, after, " s"))
	case "damageMultiplier":
		return percentChange(name+"'s damage bonus", (before-1)*100, (after-1)*100, activeDamage(before), activeDamage(after))
	case "intervalMultiplier":
		return percentChange(name+"'s attack speed", rateChange(1, before)*100, rateChange(1, after)*100, activeSpeed(before), activeSpeed(after))
	}
	return numberChange(name+"'s range bonus", "", before, after, false, delta(before, after, ""))
}

// percentChange writes a change between two percentages, shown as before
// and after text, with its delta in percentage points.
func percentChange(label string, before, after float64, beforeText, afterText string) string {
	if before == after {
		return "Keeps " + label + " at " + afterText + "."
	}
	verb := "Raises"
	if after < before {
		verb = "Lowers"
	}
	return verb + " " + label + " from " + beforeText + " to " + afterText + " (" + pointsDelta(before, after) + ")."
}

// abilitySentence describes the Active Ability a purchase unlocks: a boost
// of the purchased attack, with its multipliers as percentages.
func (sh *sheet) abilitySentence(ability m.ResolvedAbility) string {
	boosted := ability.BoostedAttack.Stats
	var effects []string
	if ability.DamageMultiplier != 1 {
		effects = append(effects, activeDamage(ability.DamageMultiplier)+" damage")
	}
	if ability.IntervalMultiplier != 1 {
		effects = append(effects, activeSpeed(ability.IntervalMultiplier)+" attacks")
	}
	if ability.RangeBonus != 0 {
		effects = append(effects, signedNumber(ability.RangeBonus)+" range")
	}
	boost := "gives the purchased attack " + joinAnd(effects)
	if len(effects) == 0 {
		boost = "leaves the purchased attack unchanged"
	}
	return fmt.Sprintf("Adds %s, this Unit's Active Ability: for %s s it %s, so the attack deals %s damage%s every %s s at range %s. It is ready on purchase, recharges %s s after activation and cannot reactivate while active; it grants no separate attack.",
		ability.Name, decimal(ability.DurationSeconds), boost,
		decimal(boosted.Damage), sh.plusBonus(ability.BoostedAttack), decimal(boosted.IntervalSeconds), decimal(boosted.Range), decimal(ability.CooldownSeconds))
}

func newAbility(before, after m.Build) *m.ResolvedAbility {
	for i := range after.Abilities {
		found := false
		for _, old := range before.Abilities {
			if old.Path == after.Abilities[i].Path {
				found = true
			}
		}
		if !found {
			return &after.Abilities[i]
		}
	}
	return nil
}

func newAbilityOrLast(after m.Build) *m.ResolvedAbility {
	if len(after.Abilities) == 0 {
		return nil
	}
	return &after.Abilities[len(after.Abilities)-1]
}

// BaseUnit is 0-0-0 as the unit sheet describes it.
type BaseUnit struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Text string `json:"text"`
}

// Base describes 0-0-0 of a valid blueprint: its price and attack. It
// returns nil for a candidate without valid typed mechanics.
func Base(candidate unit.Candidate, definition *m.Definition) *BaseUnit {
	sh := newSheet(candidate.Blueprint, definition)
	if sh == nil {
		return nil
	}
	base := sh.resolve(m.Selection{}).BaseAttack
	return &BaseUnit{Code: "0-0-0", Name: base.Name, Text: sh.baseText(base)}
}

func (sh *sheet) baseText(base m.Attack) string {
	return "Placement costs " + sh.money(base.Cost) + ". " + strings.Join(sh.attackSentences(base), " ")
}

// Purchases describes every purchase of a valid blueprint along its pure
// path, with the technique each adapts when the plan names one. It returns
// nil for a candidate without valid typed mechanics.
func Purchases(candidate unit.Candidate, definition *m.Definition, plan *unit.DesignPlan) []PathPurchases {
	sh := newSheet(candidate.Blueprint, definition)
	if sh == nil {
		return nil
	}
	sh.plan = plan
	return sh.purchases()
}

func (sh *sheet) purchases() []PathPurchases {
	var out []PathPurchases
	for index := range m.PathKeys {
		path := sh.blueprint.Paths.At(index)
		entry := PathPurchases{Position: pathPositions[index], Name: path.Name}
		for tier := 1; tier <= len(m.TierKeys); tier++ {
			var before, after m.Selection
			before[index], after[index] = tier-1, tier
			upgrade := path.Tiers.At(tier)
			technique, base := unit.PurchaseTechnique(sh.plan, index, tier)
			entry.Purchases = append(entry.Purchases, Purchase{
				Code: unit.BuildCode(index, tier), Name: upgrade.Name, Cost: upgrade.Cost,
				Technique: technique, BaseTechnique: base,
				Adaptation:   unit.PurchaseAdaptation(sh.plan, index, tier),
				NameMentions: unit.PurchaseNameMentions(sh.plan, index, tier, upgrade.Name),
				Effects:      sh.purchaseEffects(upgrade.Changes, sh.resolve(before), sh.resolve(after)),
			})
			last := &entry.Purchases[len(entry.Purchases)-1]
			last.ProposedMechanics = upgrade.ProposedMechanics
			last.Text = last.Passage()
		}
		out = append(out, entry)
	}
	return out
}
