package unit

import m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"

// DesignGuidance is the Definition policy guidance shared by drafting and repair.
func DesignGuidance(request *Request) []string {
	if request.MechanicsDefinition == nil || request.MechanicsDefinition.Profile.DesignPolicy == nil {
		return nil
	}
	definition := request.MechanicsDefinition
	policy := definition.Profile.DesignPolicy
	specializations := "Shared specializations are allowed."
	if policy.DistinctPathSpecializations {
		specializations = m.DistinctSpecializationsRule
	}
	automatic := guideAutomaticPaths
	if policy.ForbidsActiveAbility() {
		automatic = guideAutomaticOnly
	}
	out := []string{
		"The selected designPolicy is a hard authoring constraint with heuristic metrics, not a balance law. Declare a specialization for each path: direct-damage, group-damage, attack-speed, control, range or ability-burst. " + specializations,
		m.ActiveAbilityRule(definition) + " " + automatic,
	}
	if policy.MinTier5SpecialtyMultiplier == nil {
		out = append(out, guideNoCapstoneMultiplier)
	} else {
		out = append(out, m.CapstoneMultiplierRule(*policy.MinTier5SpecialtyMultiplier)+" "+guideCapstoneMultiplier)
	}
	first := "Repeated early behavior is allowed."
	if policy.DistinctFirstUpgrades {
		first = m.DistinctFirstPurchasesRule
	}
	capstones := ""
	if policy.DistinctCapstones {
		capstones = m.DistinctCapstonesRule
	}
	out = append(out, "Distinct purchases: "+first+" "+capstones+" Strengthen the purchased branch rather than granting the other branches. A high-tier generalist can still be classified basic_dps; classification is not a power or quality grade.")
	if exclusiveEarlyOn(*definition) {
		out = append(out, guideExclusiveEarly)
	}
	if earlyBenefitsOn(*definition) {
		out = append(out, guideEarlyBenefits)
	}
	if policy.PreservesEarlyIdentity() {
		out = append(out, m.EarlyIdentityRule(definition))
	}
	if policy.RequiresBehaviorChange(3) {
		out = append(out, m.BehaviorChangeRule(3)+" "+guideTier3Promise)
	}
	if policy.RequiresBehaviorChange(5) {
		out = append(out, m.BehaviorChangeRule(5)+" "+guideTier5Promise)
	}
	if coreConceptsOn(*definition) {
		out = append(out, guideCoreConcepts)
	}
	return append(out, guideDistinctPaths, guideNames, guidePrices)
}
