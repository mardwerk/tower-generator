package unit

import (
	"fmt"
	"slices"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Unlock order (#61, SOL-61-21). Escanor a's v39 plan promised improves
// splash at x-x-3 and unlock splash at x-x-5 (OPUS-NET-61-26): once x-x-3
// keeps its promise, every build that owns x-x-5 already has splash, so no
// x-x-5 can keep the unlock, and the draft failed at the design stage. The
// check reads the plan's own promises only: a milestone's unlock fails when
// an earlier milestone of the same path promises the same name, as an
// improvement or as a capability unlock. It infers nothing from stats, so
// the unlock of another capability, such as follow-up after an improved
// splash, is unaffected. bonus-damage is the one promise whose unlock
// adds a bonus against an enemy property the attack had none against
// (bonusDamageGuidance), so it may follow an improved bonus-damage; a
// second unlock of it still fails, as a capability unlocked twice always
// has. The issue is on the later milestone. When the Definition lists the
// name as an improvement, the targeted plan correction can fix it
// (planCorrections): move the unlock to the earlier milestone, or make the
// later promise an improvement.

// UnlockOrderRule is the rule the plan prompt and the plan check state.
const UnlockOrderRule = "A path unlocks a capability at its first promise of it: no milestone promises unlock X after an earlier milestone of its path promises improves X or unlock X, since the path has X from that earlier purchase on; a later purchase that develops X promises improves X. Only the plan's own promise names are compared, so the unlock of another capability, such as follow-up after an improved splash, is unaffected, and so is unlock bonus-damage after improves bonus-damage, which adds a bonus against another enemy property."

// lateUnlock is a milestone whose unlock an earlier milestone of its path
// already promises.
type lateUnlock struct {
	path, tier int
	// earlier is the tier of the first earlier milestone that promises it,
	// and improved whether that promise is an improvement.
	earlier  int
	improved bool
}

func (u lateUnlock) unlock(plan DesignPlan) string {
	return plan.UpgradeIntents.At(u.path).At(u.tier).Unlock
}

// earlierPromise says what the earlier milestone promises.
func (u lateUnlock) earlierPromise(plan DesignPlan) string {
	if u.improved {
		return "improves " + u.unlock(plan)
	}
	return "unlock " + u.unlock(plan)
}

// correctable reports a late unlock the targeted correction can fix: the
// Definition lists its name as an improvement, so the later milestone can
// promise that instead.
func (u lateUnlock) correctable(plan DesignPlan, definition m.Definition) bool {
	return slices.Contains(ImprovementsFor(&definition), u.unlock(plan))
}

// lateUnlocks lists each milestone whose capability unlock an earlier
// milestone of the same path promises as an improvement or an unlock,
// compared by exact promise name.
func lateUnlocks(plan DesignPlan, definition m.Definition) []lateUnlock {
	if plan.UpgradeIntents == nil {
		return nil
	}
	var out []lateUnlock
	for pathIndex := range m.PathKeys {
		intents := plan.UpgradeIntents.At(pathIndex)
		for tier := 2; tier <= len(m.TierKeys); tier++ {
			unlock := intents.At(tier).Unlock
			if !isCapabilityUnlock(definition, unlock) {
				continue
			}
			for earlier := 1; earlier < tier; earlier++ {
				intent := intents.At(earlier)
				if intent.Unlock == unlock {
					out = append(out, lateUnlock{pathIndex, tier, earlier, false})
					break
				}
				if unlock != BonusDamagePromise && slices.Contains(intent.Improves, unlock) {
					out = append(out, lateUnlock{pathIndex, tier, earlier, true})
					break
				}
			}
		}
	}
	return out
}

// UnlockOrderIssues reports each milestone that promises to unlock what an
// earlier milestone of its path already promises, on the later milestone.
// correctable are those a targeted plan correction can fix; the others,
// a capability the Definition cannot improve unlocked twice, need a full
// plan.
func UnlockOrderIssues(plan DesignPlan, definition m.Definition) (correctable, other []m.Issue) {
	for _, u := range lateUnlocks(plan, definition) {
		unlock := u.unlock(plan)
		code, earlier := BuildCode(u.path, u.tier), BuildCode(u.path, u.earlier)
		fix := fmt.Sprintf("Promise unlock %s only at %s, and give %s another unlock or none.", unlock, earlier, code)
		if u.correctable(plan, definition) {
			fix = fmt.Sprintf("Change %s's promise from unlock %s to improves %s.", code, unlock, unlock)
			if u.improved && canMoveUnlock(plan, definition, u) {
				fix = fmt.Sprintf("Move the unlock of %s to %s, or change %s's promise from unlock %s to improves %s.", unlock, earlier, code, unlock, unlock)
			}
		}
		issue := m.Issue{
			Path:    fmt.Sprintf("upgradeIntents.%s.%s", m.PathKeys[u.path], m.TierKeys[u.tier-1]),
			Message: fmt.Sprintf("%s promises unlock %s, but %s already promises %s. %s %s", code, unlock, earlier, u.earlierPromise(plan), UnlockOrderRule, fix),
		}
		if u.correctable(plan, definition) {
			correctable = append(correctable, issue)
		} else {
			other = append(other, issue)
		}
	}
	return correctable, other
}

// canMoveUnlock reports whether the earlier milestone can take the unlock:
// it unlocks nothing now, and the early-identity policy lets its tier
// unlock it.
func canMoveUnlock(plan DesignPlan, definition m.Definition, u lateUnlock) bool {
	if plan.UpgradeIntents.At(u.path).At(u.earlier).Unlock != "none" {
		return false
	}
	return u.earlier > 2 || !definition.Profile.DesignPolicy.PreservesEarlyIdentity() || !earlyIdentityBreaking(&definition, u.unlock(plan))
}

// unlockOrderItems are the targeted-correction items of the correctable
// late unlocks. With open paths, a capstone correction's, the only
// milestones the correction may change are those paths' fifth purchases,
// so an item is named only when its later milestone is one of them, and
// only the improvement is allowed.
func unlockOrderItems(plan DesignPlan, request *Request, open []int) []any {
	items := []any{}
	definition := request.MechanicsDefinition
	if definition == nil {
		return items
	}
	for _, u := range lateUnlocks(plan, *definition) {
		if !u.correctable(plan, *definition) || (open != nil && (u.tier != 5 || !slices.Contains(open, u.path))) {
			continue
		}
		unlock := u.unlock(plan)
		code, earlier := BuildCode(u.path, u.tier), BuildCode(u.path, u.earlier)
		behavior := ""
		if policy := definition.Profile.DesignPolicy; policy.RequiresBehaviorChange(u.tier) && open == nil {
			intent := *plan.UpgradeIntents.At(u.path).At(u.tier)
			intent.Unlock = "none"
			if !addsBehavior(intent) && len(intent.ProposedMechanics) == 0 {
				behavior = fmt.Sprintf(" %s must still add new behavior or access: %s Without unlock %s it has none, so give it an unlock its path does not promise before it, or, only when the Definition cannot express the capability the sources describe, a proposed mechanic.", code, m.BehaviorChangeRule(u.tier), unlock)
			}
		}
		allowed := []any{}
		if u.improved && open == nil && canMoveUnlock(plan, *definition, u) {
			allowed = append(allowed, fmt.Sprintf("Move the unlock to %s: set its unlock to %s in place of none and keep its other promises, and take unlock %s from %s, which keeps improves %s when it still develops it. Say in each milestone's change what it now adds.%s", earlier, unlock, unlock, code, unlock, behavior))
		}
		allowed = append(allowed, fmt.Sprintf("Change %s's promise from unlock %s to improves %s, and say in its change that it develops the %s %s adds.%s", code, unlock, unlock, unlock, earlier, behavior))
		items = append(items, s.NewObject().
			Set("path", fmt.Sprintf("paths.%s.milestones.%s", m.PathKeys[u.path], m.TierKeys[u.tier-1])).
			Set("problem", fmt.Sprintf("%s promises unlock %s, but %s already promises %s. %s", code, unlock, earlier, u.earlierPromise(plan), UnlockOrderRule)).
			Set("allowed", allowed))
	}
	return items
}
