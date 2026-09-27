package unit

import (
	"regexp"
	"strconv"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// Local repair (#61, SOL-61-19). Three of four v39 failed runs ended with
// one undelivered planned unlock on one fifth purchase after the single
// design repair: Luffy a's x-5-x promised unlock follow-up and bought an
// activeFollowUp, and Escanor a's x-x-5 promised unlock splash
// (OPUS-NET-61-24). When the last repair's output passes its schema and
// every issue left is a promise to unlock or improve that the same single
// purchase does not implement (PlanIntentIssues), one more repair call is
// sent, scoped to that purchase. It names each undelivered promise, the
// purchase's current changes, the resolved attack with and without it, and
// the change the Definition offers to deliver it.
//
// Code merges back only that purchase's tier and discards everything else
// the call returns, then validates the whole blueprint again as every
// attempt is validated: typed validation, every legal build, the design
// policy and every plan promise, including the ones that held before. So a
// fix that breaks a promise that held, or leaves this one undelivered,
// fails the draft with every issue reported, and the promise is never
// loosened.
//
// It is at most one call per draft, never sent with a repair budget of 0,
// and uses no retry budget. It is not sent when the purchase alone cannot
// deliver an unlock: when a legal build that owns it already has the
// capability without it, from an earlier purchase of its path or from
// another path. Escanor a's x-x-3 already gives 0-0-4 splash, so no change
// of x-x-5 unlocks splash. The attempt is recorded with purpose
// "local-repair", its usage and its issues.

// LocalRepairPurpose is the attempt purpose of a local repair.
const LocalRepairPurpose = "local-repair"

// promiseIssue is an undelivered promise to unlock or improve, as
// PlanIntentIssues reports it on a purchase.
var promiseIssue = regexp.MustCompile(`^paths\.(path[1-3])\.tiers\.(tier[1-5])\.planIntent: The retained plan promises (unlock|improved) ([^,]+), but this purchase does not implement it in legal build ([0-5])-([0-5])-([0-5])\. `)

// undeliveredPromise is one promise a local repair names.
type undeliveredPromise struct {
	kind, dimension string
	build           m.Selection
}

// localRepairFor is the local repair of a bound mechanics output whose
// schema passed and whose issues are all undelivered promises of one
// purchase, or nil when none applies.
func localRepairFor(request *Request, previous any, issues []string) (*TierRepair, error) {
	if len(issues) == 0 || request.MechanicsDefinition == nil {
		return nil, nil
	}
	pathIndex, tier := -1, 0
	var promises []undeliveredPromise
	for _, issue := range issues {
		match := promiseIssue.FindStringSubmatch(issue)
		if match == nil {
			return nil, nil
		}
		p, t := int(match[1][4]-'1'), int(match[2][4]-'0')
		if pathIndex >= 0 && (p != pathIndex || t != tier) {
			return nil, nil
		}
		pathIndex, tier = p, t
		var build m.Selection
		for i := range build {
			build[i], _ = strconv.Atoi(match[5+i])
		}
		if build[pathIndex] != tier {
			return nil, nil
		}
		promises = append(promises, undeliveredPromise{match[3], match[4], build})
	}
	full, err := ModelOutputSchema(request)
	if err != nil {
		return nil, err
	}
	parsedValue, parseIssues := s.Parse(full, previous)
	if len(parseIssues) > 0 {
		return nil, nil
	}
	parsed := parsedValue.(*s.Object)
	blueprint, err := DecodeBlueprintOutput(parsed, request)
	if err != nil {
		return nil, nil
	}
	definition := *request.MechanicsDefinition
	for _, promise := range promises {
		if promise.kind != "unlock" {
			continue
		}
		for _, selection := range m.AllLegalBuilds(definition) {
			if selection[pathIndex] != tier {
				continue
			}
			selection[pathIndex] = tier - 1
			if unlockBlocked(m.ResolveUnchecked(&blueprint, selection).BaseAttack, promise.dimension) {
				return nil, nil
			}
		}
	}

	path, key := m.PathKeys[pathIndex], m.TierKeys[tier-1]
	tierSchema := full.Shape("paths").(*s.ObjectSchema).Shape(path).(*s.ObjectSchema).Shape("tiers").(*s.ObjectSchema).Shape(key)
	schema := s.StrictObject(s.F("paths", s.StrictObject(s.F(path, s.StrictObject(s.F("tiers", s.StrictObject(s.F(key, tierSchema))))))))
	current := field(field(field(field(parsed, "paths").(*s.Object), path).(*s.Object), "tiers").(*s.Object), key)
	listed := []any{}
	for _, promise := range promises {
		without := promise.build
		without[pathIndex] = tier - 1
		listed = append(listed, s.NewObject().
			Set("promise", promise.kind+" "+promise.dimension).
			Set("build", buildLabel(promise.build)).
			Set("deliverWith", promiseOffer(promise)).
			Set("attackWithoutPurchase", s.FromGoValue(m.ResolveUnchecked(&blueprint, without).BaseAttack)).
			Set("attackWithPurchase", s.FromGoValue(m.ResolveUnchecked(&blueprint, promise.build).BaseAttack)))
	}
	scope := s.NewObject().
		Set("purchase", BuildCode(pathIndex, tier)).
		Set("path", path).
		Set("tier", key).
		Set("currentPurchase", s.Clone(current)).
		Set("undeliveredPromises", listed).
		Set("violations", stringList(issues))
	prompt := append([]string{localRepairScope, s.Stringify(repairContext(request, parsed))}, repairGuidance(request)...)
	prompt = append(prompt, s.Stringify(scope))
	return &TierRepair{
		Request: ModelRequest{System: repairSystem, Prompt: strings.Join(prompt, "\n\n"), Schema: ProviderJSONSchema(schema)},
		// Apply takes only this purchase's tier from the output and
		// discards every other path, tier or key. A tier whose only
		// schema issues are numbers outside their bounds is merged as
		// written, as a targeted repair's is; any other schema issue
		// rejects it.
		Apply: func(patch any) (*s.Object, error) {
			value := valueAt(patch, []any{"paths", path, "tiers", key})
			replacement, issues := s.Parse(tierSchema, value)
			if _, _, ok := numbersInBounds(value, issues); len(issues) > 0 && !ok {
				for i := range issues {
					issues[i].Path = append([]any{"paths", path, "tiers", key}, issues[i].Path...)
				}
				return nil, &s.Error{Issues: issues}
			}
			merged := s.Clone(parsed).(*s.Object)
			field(field(field(merged, "paths").(*s.Object), path).(*s.Object), "tiers").(*s.Object).Set(key, s.Clone(replacement))
			return merged, nil
		},
	}, nil
}

// unlockBlocked reports that an attack already has what an unlock adds, so
// no change of the next purchase can unlock it (unlockedIntent).
func unlockBlocked(attack m.Attack, intent string) bool {
	if attack.IsV2() && !coreUnlocks[intent] && intent != BonusDamagePromise {
		_, had := attack.Status(intent)
		return had && attack.DetectsTrait(intent)
	}
	switch intent {
	case "follow-up":
		return attack.FollowUp != nil
	case "distinct-volley":
		return attack.Distribution == "distinct-targets"
	case "splash":
		return attack.Stats.SplashRadius > 0
	case "camo":
		return attack.Camo
	case "slow":
		return attack.Stats.SlowPercent > 0
	case "burn":
		return attack.Stats.BurnDamagePerSecond > 0
	case "stun":
		return attack.Stats.StunSeconds > 0
	}
	return false
}

// promiseOffer names what the Definition offers to deliver one undelivered
// promise at its purchase.
func promiseOffer(promise undeliveredPromise) string {
	if promise.kind == "improved" {
		return promiseFix(promise.dimension)
	}
	if offer, ok := unlockOffers[promise.dimension]; ok {
		return offer
	}
	return unlockOfferOther
}

func buildLabel(selection m.Selection) string {
	return strconv.Itoa(selection[0]) + "-" + strconv.Itoa(selection[1]) + "-" + strconv.Itoa(selection[2])
}
