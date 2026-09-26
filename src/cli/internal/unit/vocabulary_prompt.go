package unit

// Prompt lines for version 2 Definitions. They replace the version 1 lines
// that name slow, burn, stun and Camo; the vocabulary supplies the rest.
const (
	draftStatusBudgetV2 = "Each tier has statChanges and statuses plus nullable detect/delivery/damageType/targeting fields. Null or an empty list means unchanged; do not repeat current values. %s Count each statChanges or boostChanges entry and every nonnull enum/unlockBoost field as one. Each statuses entry counts as TWO changes when it has a magnitude and one otherwise, but one new capability. Tier1 and Tier2 may add only one new capability."
	draftStatusFormV2   = "Use statuses:[] when no status effect changes. A statuses entry {effect, magnitude, seconds} introduces or replaces one of the Definition's status effects: a positive magnitude in the effect's unit and within its bounds, or null for an effect without magnitude, and a positive duration within its limit. Never put status fields in statChanges. detect names a hidden trait the attack starts detecting. Zero disables splash; multiplying zero cannot enable an effect."
	draftExampleV2      = "Example Tier1: {\"name\":\"Focused Strike\",\"cost\":100,\"statChanges\":[{\"stat\":\"damage\",\"operation\":\"add\",\"value\":1}],\"statuses\":[],\"detect\":null,\"delivery\":null,\"damageType\":null,\"targeting\":null,\"unlockBoost\":null,\"boostChanges\":[]}. A Tier4 boost could use damageMultiplier:2, intervalMultiplier:0.8, durationSeconds:8, cooldownSeconds:30, rangeBonus:0."
	repairBudgetV2      = "Correct the violations below with the smallest coherent tier changes. Preserve the established role, path themes and valid benefits. Count every statChanges item, boostChanges item and nonnull detect/delivery/damageType/targeting/unlockBoost field toward the tier effect budget. Each statuses entry also counts as two primitive changes when it has a magnitude, otherwise one, and must keep positive values. Removing an effect must preserve at least one meaningful benefit. Choose subsets that preserve the tactical purpose of each path and valid inherited mechanics. A subset menu guarantees the effect count only; avoid ineffective or harmful combinations. Output only the requested repair fields."
	reviewStatusesV2    = "Inspect meaningful early upgrades and crosspaths. Repeated-primary volleys can hit the same target; explicitly distinct-target volleys assign one initial shot per eligible target. Bounded follow-ups happen once after primary hits and exclude those targets, with declared status inheritance and no recursion; each status effect's kind, bounds, duration limit, stacking and immunities are defined in the Definition's vocabulary. Boosts modify the already purchased attack. Distinguish a numerical power concern from a proven contradiction. Numerical values are on an experimental starter scale; do not claim playtested balance."
)

// isV2 reports a request under a version 2 Definition.
func isV2(request *Request) bool {
	return request.MechanicsDefinition != nil && request.MechanicsDefinition.IsV2()
}

// VocabularyGuidance says how to use a version 2 Definition's vocabulary.
// The vocabulary itself, with IDs, names, aliases, bounds, stacking,
// immunities and descriptions, is in the prompt context's
// definition.vocabulary, so it is not restated here. It is empty for version
// 1 Definitions, whose vocabulary is fixed in the prompts.
func VocabularyGuidance(request *Request) []string {
	if !isV2(request) {
		return nil
	}
	text := "definition.vocabulary lists every status effect, damage type, targeting mode and detection trait by ID; no other exists. "
	if len(request.MechanicsDefinition.Vocabulary.StatusEffects) == 0 {
		text += "It defines no status effects; express none. "
	} else {
		text += "Map source wording to a status effect through its name and aliases and keep within its magnitude, maxSeconds and stacking; describe anything else in unsupportedMechanics. "
	}
	return []string{text + "Choose a damage type from its description; ineffectiveAgainst lists the enemy properties it cannot damage. A hidden enemy can be targeted only by an attack with the matching detection trait."}
}
