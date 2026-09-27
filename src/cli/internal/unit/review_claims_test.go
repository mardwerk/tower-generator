package unit

import (
	"strings"
	"testing"
)

// Trimmed from Sol's v22 Luffy diagnostic on #57 (Result e27c4546): the
// plan's base attack, repertoire names and bottom-path techniques, the
// names and planned text of x-x-1 and x-x-2, and the againstCapstone rows
// of the middle path's x-5-x whose side gain rounds to 4.66, with the rows
// of 2-x-x at 1-4-0 and x-x-2 at 0-5-1 beside them.

const (
	solTimingMessage = "At x-x-1, “Driving Pistol” and its plan text claim only a harder Gum-Gum Pistol, but this purchase begins the Gear 4 Tankman path whose cited source describes a larger form that sacrifices speed and uses recoil to blast opponents away. The only Tankman mechanic, knockback, arrives at x-x-3. Revise x-x-1's technique and adaptation to explain the transition, or start the Tankman specialization later if the design permits."
	solTimingAction  = "Clarify how x-x-1 introduces Tankman, or defer the named path transition to the purchase that delivers recoil control."
	solPriceMessage  = "The x-5-x purchase costs 18,000 Gold and adds 6.21 time-averaged direct damage rate over x-4-x. Its gain is smaller than the path1 x-2-x side purchase's 4.66 gain, which costs 180 Gold, providing evidence that the capstone adds too little for its price. Reconsider the price or increase the burst window payoff while keeping the single-target Gear 2 role."
	solPriceAction   = "Reprice x-5-x or strengthen its resolved active benefit, then reassess against the x-2-x side purchase."
)

func solPlan() DesignPlan {
	pistol, tankman := "Gum-Gum Pistol", "Gear 4 Tankman"
	intent := func(technique string, improves ...string) UpgradeIntent {
		return UpgradeIntent{Improves: improves, Unlock: "none", Technique: technique}
	}
	path := func(later string) PathIntents {
		return PathIntents{Tier1: intent(pistol, "damage"), Tier2: intent(pistol, "damage", "pierce"),
			Tier3: intent(later, "damage"), Tier4: intent(later, "damage"), Tier5: intent(later, "damage")}
	}
	plan := DesignPlan{Base: PlanBase{Name: pistol}, UpgradeIntents: &UpgradeIntents{Path1: path("Gear 3"), Path2: path("Gear 2"), Path3: path(tankman)}}
	for _, name := range []string{pistol, "Gear 2", "Gear 3", "Gear 4 Boundman", tankman, "Gear 4 Snakeman", "Gum-Gum Gatling", "Conqueror's Haki"} {
		plan.Repertoire = append(plan.Repertoire, PlanRepertoire{Name: name})
	}
	plan.Paths.Path3.Milestones.Tier1 = "x-x-1: Luffy drives the Gum-Gum Pistol with more force into the player's selected target, increasing damage while retaining its single-target delivery."
	plan.Paths.Path3.Milestones.Tier2 = "x-x-2: Luffy drives the Gum-Gum Pistol through the selected target and into one additional distinct eligible enemy in its path, increasing damage and hit capacity while preserving the stronger single-target punch when no secondary enemy is present."
	return plan
}

func solRows() []comparisonRow {
	row := func(from, to, side string, price float64, metric string, gain float64, atLeast bool) comparisonRow {
		return comparisonRow{from: from, to: to, side: side, capstone: "x-5-x", sidePrice: price, capstonePrice: 18000.0,
			metric: metric, sideGain: gain, capstoneGain: 6.2077, atLeast: atLeast}
	}
	return []comparisonRow{
		row("1-4-0", "2-4-0", "2-x-x", 180, TimeAveragedDirect, 2.5905, false),
		row("1-4-0", "2-4-0", "2-x-x", 180, TimeAveragedGroup, 2.5905, false),
		row("1-5-0", "2-5-0", "2-x-x", 180, TimeAveragedDirect, 4.6598, false),
		row("1-5-0", "2-5-0", "2-x-x", 180, TimeAveragedGroup, 4.6598, false),
		row("0-5-0", "0-5-1", "x-x-1", 120, TimeAveragedDirect, 4.6598, false),
		row("0-5-0", "0-5-1", "x-x-1", 120, TimeAveragedGroup, 4.6598, false),
		row("0-5-1", "0-5-2", "x-x-2", 240, TimeAveragedDirect, 4.6598, false),
		row("0-5-1", "0-5-2", "x-x-2", 240, TimeAveragedGroup, 27.9586, true),
	}
}

func solClaims() reviewClaims {
	plan := solPlan()
	claims := reviewClaims{plan: &plan, rows: solRows()}
	claims.names[2] = [5]string{"Driving Pistol", "Piercing Pistol", "Tankman Recoil Punch", "Heavier Recoil Punch", "Long-Range Tankman Punch"}
	return claims
}

func modelFinding(id, outcome, subject, message, action string) Finding {
	return Finding{ID: id, Method: "model", Category: "evidence", Severity: severity(outcome), Outcome: outcome,
		Subject: subject, Rule: "A rule.", Message: message, Evidence: []string{}, Action: act(action)}
}

// Miss 2 on #57: Sol's review failed x-x-1 for not being Gear 4 Tankman,
// which its path adapts from x-x-3. The guard states the timing.
func TestTimingGuardCorrectsAPurchaseFaultedForALaterTechnique(t *testing.T) {
	claims := solClaims()
	wrong, uncertain := claims.timing(modelFinding("model.path3-t1-source-fit", "fail", "x-x-1 Tankman entry purchase", solTimingMessage, solTimingAction))
	if len(wrong) != 1 || len(uncertain) != 0 {
		t.Fatalf("wrong %v, uncertain %v", wrong, uncertain)
	}
	if got, want := wrong[0].String(), "model.path3-t1-source-fit faults x-x-1 for Gear 4 Tankman, but x-x-1 adapts Gum-Gum Pistol; Gear 4 Tankman starts at x-x-3."; got != want {
		t.Errorf("correction %q, want %q", got, want)
	}
}

// Mentioning a later technique is not by itself a false claim (SOL-57-01).
func TestTimingGuardSparesValidMentions(t *testing.T) {
	for _, test := range []struct {
		name, outcome, subject, message string
		rename                          string
		plan                            func(*DesignPlan)
	}{
		// A contextual mention that states the timing.
		{name: "context", outcome: "fail", subject: "x-x-1 Driving Pistol",
			message: "x-x-1 Driving Pistol credits the punch with more force, but its cited Pistol passages describe only a stretched punch. This purchase keeps the Gum-Gum Pistol, and the heavier Gear 4 Tankman blow starts at x-x-3."},
		// A mention in a sentence about the path, not the purchase.
		{name: "path context", outcome: "fail", subject: "x-x-1",
			message: "x-x-1 cites only the Pistol passage for its extra force. The Gear 4 Tankman path is well sourced."},
		// The control on #57: a purchase whose name claims the technique.
		{name: "name claims it", outcome: "fail", subject: "x-x-1 Tankman entry purchase", message: solTimingMessage, rename: "Tankman Warm-up"},
		// luffy10a: x-x-1 "Boundman Force" names Gear 4, Boundman and
		// Snakeman traits, so a finding that it lacks that source is valid.
		{name: "luffy10a", outcome: "fail", subject: "x-x-1 Boundman Force", rename: "Boundman Force",
			message: "x-x-1 calls this a Boundman Force and says Luffy hardens the punch, but its cited Pistol passages support a stretched punch, not Boundman or hardening. Cite Gear 4 evidence for this tier and describe the supported compression or Haki trait, or rename it as a Pistol damage upgrade.",
			plan: func(plan *DesignPlan) {
				traits := "Gear 4, Boundman and Snakeman traits"
				plan.Repertoire = []PlanRepertoire{{Name: "Gum-Gum Pistol"}, {Name: "Gear 2"}, {Name: "Gear 3"}, {Name: traits}}
				for _, intent := range []*UpgradeIntent{&plan.UpgradeIntents.Path3.Tier3, &plan.UpgradeIntents.Path3.Tier4, &plan.UpgradeIntents.Path3.Tier5} {
					intent.Technique = traits
				}
				plan.Paths.Path3.Milestones.Tier1 = "x-x-1: Luffy hardens and drives his existing punch more forcefully, increasing its normal damage."
			}},
		// A pass is not a fault.
		{name: "pass", outcome: "pass", subject: "x-x-1 Tankman entry purchase", message: solTimingMessage},
		// A subject that adapts the later technique is judged on it.
		{name: "own technique", outcome: "fail", subject: "x-x-3", message: "x-x-3 begins the Gear 4 Tankman path without its defense."},
	} {
		claims := solClaims()
		if test.rename != "" {
			claims.names[2][0] = test.rename
		}
		if test.plan != nil {
			test.plan(claims.plan)
		}
		f := modelFinding("model.timing", test.outcome, test.subject, test.message, "Revise x-x-1.")
		if wrong, uncertain := claims.timing(f); len(wrong)+len(uncertain) > 0 {
			t.Errorf("%s: wrong %v, uncertain %v", test.name, wrong, uncertain)
		}
		if extra := claims.humanReviewFindings([]Finding{f}, "Gold"); len(extra) > 0 {
			t.Errorf("%s: marked for review: %s", test.name, extra[0].Message)
		}
	}
}

// A finding that relates x-x-1 to a later technique without a readable
// fault is kept as written, and an unresolved Finding asks for it to be
// checked instead of a speculative correction.
func TestTimingGuardRetainsAnUnreadClaimForReview(t *testing.T) {
	claims := solClaims()
	f := modelFinding("model.path3-t1-fit", "fail", "x-x-1 Driving Pistol", "x-x-1 Driving Pistol feels disconnected from Gear 4 Tankman.", "Revise x-x-1.")
	wrong, uncertain := claims.timing(f)
	if len(wrong) != 0 || len(uncertain) != 1 {
		t.Fatalf("wrong %v, uncertain %v", wrong, uncertain)
	}
	extra := claims.humanReviewFindings([]Finding{f}, "Gold")
	if len(extra) != 1 {
		t.Fatalf("%d review Findings", len(extra))
	}
	got := extra[0]
	if got.ID != "deterministic.review.path3-t1-fit" || got.Method != "deterministic" || got.Outcome != "unresolved" || got.Severity != "warning" || got.Rule != HumanReviewRule || got.Subject != f.Subject ||
		!strings.Contains(got.Message, "Code could not tell whether model.path3-t1-fit faults x-x-1 for Gear 4 Tankman: x-x-1 adapts Gum-Gum Pistol, and Gear 4 Tankman starts at x-x-3.") ||
		got.Action == nil || !strings.Contains(*got.Action, "Check model.path3-t1-fit") {
		t.Errorf("review Finding %+v", got)
	}
}

// Miss 3 on #57: Sol's review said x-5-x's 6.21 gain is smaller than a side
// purchase's 4.66. The quoted price 180 Gold and "direct" leave one row,
// 2-x-x at 1-5-0, whose sideGainAtLeastCapstone is false.
func TestComparisonGuardCorrectsAReversedComparison(t *testing.T) {
	claims := solClaims()
	wrong, uncertain := claims.comparison(modelFinding("model.path2-capstone-price", "fail", "x-5-x Extended Gear 2 Burst", solPriceMessage, solPriceAction), "Gold")
	if len(wrong) != 1 || len(uncertain) != 0 {
		t.Fatalf("wrong %v, uncertain %v", wrong, uncertain)
	}
	want := "model.path2-capstone-price says the 2-x-x side purchase's 4.66 gain is at least x-5-x's, but sideGainAtLeastCapstone is false for the time-averaged direct damage rate: from 1-5-0 to 2-5-0, 2-x-x adds 4.66 for 180 Gold and x-5-x adds 6.21 for 18000 Gold."
	if got := wrong[0].String(); got != want {
		t.Errorf("correction %q, want %q", got, want)
	}
}

// Quoting a side gain is not by itself a false claim, and a number that
// matches several rows or metrics is never corrected (SOL-57-01).
func TestComparisonGuardSparesCorrectAndAmbiguousQuotes(t *testing.T) {
	for _, test := range []struct {
		name, message string
		review        bool
	}{
		// The side gain quoted in the right direction.
		{name: "correct direction", message: "x-5-x adds 6.21 time-averaged direct damage rate for 18,000 Gold. The 2-x-x side purchase's 4.66 direct gain, for 180 Gold, is smaller than the capstone's."},
		// A true row quoted as at least the capstone's gain.
		{name: "true row", message: "x-5-x adds 6.21 for 18,000 Gold. The x-x-2 side purchase adds 27.96 group gain for 240 Gold, which exceeds the capstone's."},
		// 4.66 is the side gain of three comparisons and two metrics.
		{name: "several rows", message: "x-5-x adds 6.21 over x-4-x. Its gain is smaller than a side purchase's 4.66 gain.", review: true},
		// The price leaves one comparison, but both its metrics are 4.66.
		{name: "several metrics", message: "x-5-x adds 6.21 over x-4-x. Its gain is smaller than the side purchase's 4.66 gain, which costs 180 Gold.", review: true},
		// No comparator: the gain is only quoted.
		{name: "no comparison", message: "x-5-x adds 6.21 for 18,000 Gold, and it keeps the burst role."},
	} {
		claims := solClaims()
		f := modelFinding("model.capstone", "fail", "x-5-x", test.message, "Reprice x-5-x.")
		wrong, _ := claims.comparison(f, "Gold")
		if len(wrong) > 0 {
			t.Errorf("%s: corrected: %v", test.name, wrong)
		}
		if extra := claims.humanReviewFindings([]Finding{f}, "Gold"); (len(extra) > 0) != test.review {
			t.Errorf("%s: %d review Findings, want review %v", test.name, len(extra), test.review)
		}
	}
}

// A capstone finding resting on a smaller side gain without a readable
// comparison is kept as written and marked for review.
func TestComparisonGuardRetainsAnUnreadClaimForReview(t *testing.T) {
	claims := solClaims()
	f := modelFinding("model.path2-capstone-price", "fail", "x-5-x", "x-5-x adds 6.21 time-averaged direct damage rate for 18,000 Gold; 2-x-x adds 4.66 for 180 Gold.", "Reprice x-5-x.")
	wrong, uncertain := claims.comparison(f, "Gold")
	if len(wrong) != 0 || len(uncertain) != 1 {
		t.Fatalf("wrong %v, uncertain %v", wrong, uncertain)
	}
	extra := claims.humanReviewFindings([]Finding{f}, "Gold")
	if len(extra) != 1 || extra[0].Outcome != "unresolved" || !strings.Contains(extra[0].Message, "model.path2-capstone-price quotes 4.66, the 2-x-x side purchase's time-averaged direct damage rate gain from 1-5-0 to 2-5-0, which is below x-5-x's 6.21") {
		t.Errorf("review Findings %+v", extra)
	}
	// A pass is not marked.
	f.Outcome = "pass"
	if extra := claims.humanReviewFindings([]Finding{f}, "Gold"); len(extra) > 0 {
		t.Errorf("a pass is marked for review: %+v", extra)
	}
}
