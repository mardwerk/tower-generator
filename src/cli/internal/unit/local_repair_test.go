package unit_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	"github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/research"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// v39Run is run a of one character under Default v39 (OPUS-NET-61-24):
// its prepared request from the same Sources under the Default, Luffy's
// with the required concepts Gear 4 and Gear 5, its retained plan, its
// design and repair outputs and the issues each attempt recorded
// (testdata/<character>-v39.*.json, from <character>.result.failure.json).
type v39Run struct {
	prepared               unit.Prepared
	planOutput             any
	plan                   unit.DesignPlan
	design, repair         any
	designIssues, repaired []string
}

func loadV39Run(t *testing.T, character string, required ...string) v39Run {
	t.Helper()
	sources, err := research.ParseSources(testdataValue(t, character+".sources.json"))
	if err != nil {
		t.Fatal(err)
	}
	prepared, err := sources.Prepare(unit.DefaultProfile())
	if err != nil {
		t.Fatal(err)
	}
	if len(required) > 0 {
		if prepared, err = unit.Prepare(s.FromGoValue(requiring(prepared.Request, required...))); err != nil {
			t.Fatal(err)
		}
	}
	run := v39Run{prepared: prepared, planOutput: testdataValue(t, character+"-v39.plan.json")}
	if run.plan, err = unit.DecodeDesignPlan(run.planOutput, &prepared.Request); err != nil {
		t.Fatalf("the retained plan: %v", err)
	}
	run.design = testdataValue(t, character+"-v39-design.output.json")
	run.repair = testdataValue(t, character+"-v39-repair.output.json")
	issues := testdataValue(t, character+"-v39.issues.json")
	run.designIssues, run.repaired = stringValues(at(issues, "design")), stringValues(at(issues, "repair"))
	return run
}

func luffyV39(t *testing.T) v39Run { return loadV39Run(t, "luffy", "Gear 4", "Gear 5") }

func escanorV39(t *testing.T) v39Run { return loadV39Run(t, "escanor") }

func stringValues(value any) []string {
	var out []string
	for _, item := range value.([]any) {
		out = append(out, item.(string))
	}
	return out
}

// tierPatch is a targeted repair's answer: one tier of an output under its
// path and tier.
func tierPatch(tier any, path, key string) *s.Object {
	return s.NewObject().Set("paths", s.NewObject().Set(path, s.NewObject().Set("tiers", s.NewObject().Set(key, s.Clone(tier)))))
}

// luffyFollowUp is Luffy a's x-5-x as the repair left it, with its
// activeFollowUp moved to followUp: the follow-up its plan promises.
func luffyFollowUp(t *testing.T, run v39Run) *s.Object {
	t.Helper()
	tier := s.Clone(at(run.repair, "paths", "path2", "tiers", "tier5")).(*s.Object)
	follow := at(tier, "activeFollowUp").(*s.Object)
	follow.Set("name", "Python Follow-Up Punch")
	return tier.Set("followUp", follow).Set("activeFollowUp", nil)
}

// blueprintOf decodes a mechanics output bound to its plan.
func blueprintOf(t *testing.T, run v39Run, output any) *s.Object {
	t.Helper()
	blueprint, err := unit.DecodeBlueprintOutput(unit.BindDesignPlan(s.Clone(output), run.plan), &run.prepared.Request)
	if err != nil {
		t.Fatal(err)
	}
	return s.FromGoValue(blueprint).(*s.Object)
}

// sameExcept fails unless two blueprints are byte-identical but for one
// purchase.
func sameExcept(t *testing.T, got, want *s.Object, path, tier string) {
	t.Helper()
	for _, key := range want.Keys() {
		if key == "paths" {
			continue
		}
		g, _ := got.Get(key)
		w, _ := want.Get(key)
		if s.Stringify(g) != s.Stringify(w) {
			t.Errorf("%s changed: %s", key, s.Stringify(g))
		}
	}
	for _, p := range mechanics.PathKeys {
		for _, k := range mechanics.TierKeys {
			g, w := s.Stringify(at(got, "paths", p, "tiers", k)), s.Stringify(at(want, "paths", p, "tiers", k))
			if p == path && k == tier {
				if g == w {
					t.Errorf("%s.%s did not change", p, k)
				}
			} else if g != w {
				t.Errorf("%s.%s changed: %s", p, k, g)
			}
		}
	}
	for _, p := range mechanics.PathKeys {
		for _, key := range at(want, "paths", p).(*s.Object).Keys() {
			if key != "tiers" && s.Stringify(at(got, "paths", p, key)) != s.Stringify(at(want, "paths", p, key)) {
				t.Errorf("%s.%s changed", p, key)
			}
		}
	}
}

// The saved outputs reproduce each attempt's issues verbatim: the design
// attempt's, then what the repair left, one undelivered unlock on one fifth
// purchase (OPUS-NET-61-24).
func TestV39FailuresReproduce(t *testing.T) {
	for name, run := range map[string]v39Run{"luffy": luffyV39(t), "escanor": escanorV39(t)} {
		design, _ := unit.MechanicsIssues(run.design, &run.prepared.Request, run.plan)
		repair, valid := unit.MechanicsIssues(run.repair, &run.prepared.Request, run.plan)
		if strings.Join(design, "\n") != strings.Join(run.designIssues, "\n") {
			t.Errorf("%s design issues %q", name, design)
		}
		if !valid || strings.Join(repair, "\n") != strings.Join(run.repaired, "\n") {
			t.Errorf("%s repair issues %q", name, repair)
		}
	}
}

// Luffy a's repair left x-5-x's promised follow-up undelivered: it bought
// an activeFollowUp. One local repair names that promise, the purchase's
// current changes and the followUp the Definition offers; its answer is
// merged into x-5-x alone, the whole blueprint passes every check and the
// draft succeeds. Every other tier is byte-identical to the repair's
// (SOL-61-19).
func TestLocalRepairDeliversLuffyFollowUp(t *testing.T) {
	run := luffyV39(t)
	model := billedModel{&fixture.Model{Outputs: []any{run.planOutput, run.design, run.repair, tierPatch(luffyFollowUp(t, run), "path2", "tier5")}}}
	draft, err := unit.DraftUnit(context.Background(), run.prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	if got := strings.Join(purposes(draft), ","); got != "plan,design,repair,local-repair" {
		t.Fatalf("attempts %s", got)
	}
	attempts := *draft.Run.Attempts
	if strings.Join(attempts[1].Issues, "\n") != strings.Join(run.designIssues, "\n") || strings.Join(attempts[2].Issues, "\n") != strings.Join(run.repaired, "\n") {
		t.Errorf("the design and repair issues: %q, %q", attempts[1].Issues, attempts[2].Issues)
	}
	if u := attempts[3].Usage; u == nil || u.CostUSD == nil || *u.CostUSD != 0.004 || len(attempts[3].Issues) != 0 {
		t.Errorf("the local repair attempt: %+v", attempts[3])
	}
	sameExcept(t, s.FromGoValue(*draft.Candidate.Blueprint).(*s.Object), blueprintOf(t, run, run.repair), "path2", "tier5")
	if at(s.FromGoValue(*draft.Candidate.Blueprint), "paths", "path2", "tiers", "tier5") == nil {
		t.Fatal("no x-5-x")
	}

	request := model.Requests[3]
	schema := s.Stringify(request.Schema)
	if !strings.Contains(schema, `"required":["path2"]`) || !strings.Contains(schema, `"required":["tier5"]`) || strings.Contains(schema, `"tier4"`) || strings.Contains(schema, `"path1"`) {
		t.Errorf("the local repair's schema is not scoped to x-5-x: %s", schema[:400])
	}
	prompt := request.Prompt
	for _, want := range []string{
		"Repair exactly one purchase",
		`"purchase":"x-5-x","path":"path2","tier":"tier5"`,
		`"currentPurchase":{"name":"Gear 4 Follow-Up","cost":35000,`,
		`"activeFollowUp":{"name":"Boosted Follow-Up Punch"`,
		`"promise":"unlock follow-up","build":"0-5-0","deliverWith":"Set followUp, the automatic attack's bounded follow-up`,
		"An activeFollowUp belongs to the Active Ability and does not deliver this promise.",
		`"violations":["` + run.repaired[0],
		"Preserve the retained character plan",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("the local repair lacks %q", want)
		}
	}
}

// Code merges only the scoped purchase: a local repair that also rewrites
// another tier or path has that change discarded, not rejected.
func TestLocalRepairDiscardsOtherTiers(t *testing.T) {
	run := luffyV39(t)
	patch := tierPatch(luffyFollowUp(t, run), "path2", "tier5")
	tiers := at(patch, "paths", "path2", "tiers").(*s.Object)
	tier4 := s.Clone(at(run.repair, "paths", "path2", "tiers", "tier4")).(*s.Object)
	tiers.Set("tier4", tier4.Set("cost", 1.0).Set("name", "Cheap Burst"))
	top := s.Clone(at(run.repair, "paths", "path1", "tiers", "tier1")).(*s.Object)
	at(patch, "paths").(*s.Object).Set("path1", s.NewObject().Set("tiers", s.NewObject().Set("tier1", top.Set("cost", 1.0))))
	patch.Set("baseAttack", s.NewObject().Set("name", "Gum-Gum Bazooka"))
	model := &fixture.Model{Outputs: []any{run.planOutput, run.design, run.repair, patch}}
	draft, err := unit.DraftUnit(context.Background(), run.prepared, model, fixture.Options())
	if err != nil {
		t.Fatalf("draft: %v", err)
	}
	sameExcept(t, s.FromGoValue(*draft.Candidate.Blueprint).(*s.Object), blueprintOf(t, run, run.repair), "path2", "tier5")
}

// The whole blueprint is validated again, and every promise that held
// before still has to: a local repair that delivers the follow-up but drops
// the boost changes that kept x-5-x's active-damage and active-attack-rate
// promises fails the draft, with every issue reported. No further call is
// made.
func TestLocalRepairThatBreaksAPromiseFails(t *testing.T) {
	run := luffyV39(t)
	fix := luffyFollowUp(t, run).Set("boostChanges", []any{})
	model := &fixture.Model{Outputs: []any{run.planOutput, run.design, run.repair, tierPatch(fix, "path2", "tier5")}}
	_, err := unit.DraftUnit(context.Background(), run.prepared, model, fixture.Options())
	var modelErr *unit.ModelError
	if !errors.As(err, &modelErr) || modelErr.Evidence == nil || len(model.Requests) != 4 {
		t.Fatalf("draft: %v after %d calls", err, len(model.Requests))
	}
	attempts := modelErr.Evidence.Attempts
	last := attempts[len(attempts)-1]
	if last.Purpose != unit.LocalRepairPurpose || last.Output == nil {
		t.Fatalf("the last attempt %+v", last)
	}
	issues := strings.Join(last.Issues, "\n")
	for _, want := range []string{"promises improved active-damage", "promises improved active-attack-rate"} {
		if !strings.Contains(issues, want) {
			t.Errorf("the local repair's issues lack %q: %s", want, issues)
		}
	}
	if strings.Contains(issues, "unlock follow-up") {
		t.Errorf("the follow-up is still reported: %s", issues)
	}
	if !strings.Contains(err.Error(), "after 3 attempts") || !strings.Contains(err.Error(), "active-damage") {
		t.Errorf("the failure %v", err)
	}
}

// A local repair is sent at most once per draft: when it leaves the
// promise undelivered, the draft fails with its issues, with a repair
// budget of 1 or 2 alike.
func TestLocalRepairIsSentOnce(t *testing.T) {
	run := luffyV39(t)
	unchanged := tierPatch(at(run.repair, "paths", "path2", "tiers", "tier5"), "path2", "tier5")
	for budget, outputs := range map[int][]any{
		1: {run.planOutput, run.design, run.repair, unchanged},
		2: {run.planOutput, run.design, run.repair, unchanged, unchanged},
	} {
		model := &fixture.Model{Outputs: outputs}
		options := fixture.Options()
		options.MaxRepairAttempts = &budget
		_, err := unit.DraftUnit(context.Background(), run.prepared, model, options)
		var modelErr *unit.ModelError
		if !errors.As(err, &modelErr) || modelErr.Evidence == nil || len(model.Requests) != len(outputs) {
			t.Fatalf("budget %d: %v after %d calls", budget, err, len(model.Requests))
		}
		var got []string
		for _, attempt := range modelErr.Evidence.Attempts {
			got = append(got, attempt.Purpose)
		}
		want := "plan,design,repair,local-repair"
		if budget == 2 {
			want = "plan,design,repair,repair,local-repair"
		}
		if strings.Join(got, ",") != want {
			t.Errorf("budget %d: attempts %v", budget, got)
		}
		last := modelErr.Evidence.Attempts[len(got)-1]
		if strings.Join(last.Issues, "\n") != strings.Join(run.repaired, "\n") {
			t.Errorf("budget %d: the local repair's issues %q", budget, last.Issues)
		}
	}
}

// With a repair budget of 0 no local repair is sent, even when the design
// leaves only one undelivered promise.
func TestNoLocalRepairWithoutABudget(t *testing.T) {
	run := luffyV39(t)
	budget := 0
	options := fixture.Options()
	options.MaxRepairAttempts = &budget
	model := &fixture.Model{Outputs: []any{run.planOutput, run.repair}}
	_, err := unit.DraftUnit(context.Background(), run.prepared, model, options)
	var modelErr *unit.ModelError
	if !errors.As(err, &modelErr) || len(model.Requests) != 2 || !strings.Contains(err.Error(), "unlock follow-up") {
		t.Fatalf("draft: %v after %d calls", err, len(model.Requests))
	}
}

// Escanor a's repair left x-x-5's promised splash unlock undelivered, but
// x-x-3 already gives 0-0-4 splash, so no change of x-x-5 alone can unlock
// it: the local repair is not sent and the draft fails with the repair's
// issue, as before.
func TestNoLocalRepairWhenThePurchaseCannotDeliver(t *testing.T) {
	run := escanorV39(t)
	patch := tierPatch(at(run.repair, "paths", "path3", "tiers", "tier5"), "path3", "tier5")
	model := &fixture.Model{Outputs: []any{run.planOutput, run.design, patch}}
	_, err := unit.DraftUnit(context.Background(), run.prepared, model, fixture.Options())
	var modelErr *unit.ModelError
	if !errors.As(err, &modelErr) || modelErr.Evidence == nil || len(model.Requests) != 3 {
		t.Fatalf("draft: %v after %d calls", err, len(model.Requests))
	}
	attempts := modelErr.Evidence.Attempts
	if strings.Join(attempts[1].Issues, "\n") != strings.Join(run.designIssues, "\n") || strings.Join(attempts[2].Issues, "\n") != strings.Join(run.repaired, "\n") {
		t.Errorf("the design and repair issues: %q, %q", attempts[1].Issues, attempts[2].Issues)
	}
	previous := unit.BindDesignPlan(s.Clone(run.repair), run.plan)
	if repair, err := unit.LocalRepair(&run.prepared.Request, previous, run.repaired); err != nil || repair != nil {
		t.Errorf("a local repair for Escanor: %v, %v", repair, err)
	}
	// Were x-x-3 to raise damage instead of splash, x-x-5 alone could
	// unlock it.
	blocked := s.Clone(run.repair).(*s.Object)
	damage := s.NewObject().Set("stat", "damage").Set("operation", "add").Set("value", 1.0)
	at(blocked, "paths", "path3", "tiers", "tier3").(*s.Object).Set("statChanges", []any{damage})
	if repair, err := unit.LocalRepair(&run.prepared.Request, unit.BindDesignPlan(blocked, run.plan), run.repaired); err != nil || repair == nil {
		t.Errorf("no local repair once 0-0-4 has no splash: %v", err)
	} else if !strings.Contains(repair.Request.Prompt, "Splash requires pierce of at least 2 in every legal build that owns this purchase") {
		t.Error("the splash offer does not name its pierce")
	}
}

// Only a lone purchase's undelivered promises start a local repair: not
// promises on two purchases, and not a promise beside another kind of
// issue, a promised tradeoff included.
func TestLocalRepairTrigger(t *testing.T) {
	run := luffyV39(t)
	request := &run.prepared.Request
	previous := unit.BindDesignPlan(s.Clone(run.repair), run.plan)
	promise := run.repaired[0]
	if repair, err := unit.LocalRepair(request, previous, run.repaired); err != nil || repair == nil {
		t.Fatalf("the saved issue: %v, %v", repair, err)
	}
	for name, issues := range map[string][]string{
		"two purchases":         {promise, strings.Replace(strings.Replace(promise, "path2", "path3", 1), "0-5-0", "0-0-5", 1)},
		"a build issue":         {promise, "builds.0-5-0.baseAttack.stats.splashRadius: Splash requires pierce of at least 2."},
		"a policy issue":        {promise, "paths.path2.tiers.tier5: Resolved x-5-x adds no new behavior or access over x-4-x and carries no proposed mechanic."},
		"a tradeoff":            {promise, "paths.path2.tiers.tier5.planIntent: The retained plan promises the tradeoff lowered range, but this purchase does not implement it in legal build 0-5-0. Add a statChanges entry that lowers range. Keep the purchase's other promised changes; an unrelated benefit does not satisfy it."},
		"a base attack promise": {"paths.path2.tiers.tier1.planIntent: The retained plan promises improved splash at x-1-x, but the base attack has none."},
		"no issue":              {},
	} {
		if repair, err := unit.LocalRepair(request, previous, issues); err != nil || repair != nil {
			t.Errorf("%s: %v, %v", name, repair, err)
		}
	}
	// Two promises of the same purchase are repaired together.
	improved := "paths.path2.tiers.tier5.planIntent: The retained plan promises improved active-damage, but this purchase does not implement it in legal build 0-5-1. Raise the boost's damageMultiplier in boostChanges, or damage while the boost is owned. Keep the purchase's other promised changes; an unrelated benefit does not satisfy it."
	repair, err := unit.LocalRepair(request, previous, []string{promise, improved})
	if err != nil || repair == nil || !strings.Contains(repair.Request.Prompt, `"promise":"improved active-damage","build":"0-5-1","deliverWith":"Raise the boost's damageMultiplier`) {
		t.Errorf("two promises of x-5-x: %v, %v", repair, err)
	}
}
