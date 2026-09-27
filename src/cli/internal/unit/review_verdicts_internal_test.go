package unit

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
)

// The v34 Luffy review (OPUS-NET-61-13) spent five of its eight free-form
// findings on four omission fails and the x-x-3 fail that its verdicts
// already recorded, in their own words. Word overlap cannot prove a
// paraphrase repeats a verdict (SOL-74-01), so code keeps all eight; the
// prompt steers the findings away from the verdicts instead.
func TestLuffyV34ParaphrasedRepeatsAreKept(t *testing.T) {
	data, err := os.ReadFile("testdata/luffy-v34.review.json")
	if err != nil {
		t.Fatal(err)
	}
	var saved struct {
		OmissionVerdicts []struct {
			OmissionVerdict
			Importance string `json:"importance"`
		} `json:"omissionVerdicts"`
		ThirdPurchaseVerdicts []struct {
			PurchaseVerdict
			Path string `json:"path"`
			Name string `json:"name"`
		} `json:"thirdPurchaseVerdicts"`
		Findings []Finding `json:"findings"`
	}
	if err := json.Unmarshal(data, &saved); err != nil {
		t.Fatal(err)
	}
	var subjects verdictSubjects
	var review SemanticReview
	for _, verdict := range saved.OmissionVerdicts {
		subjects.omissions = append(subjects.omissions, omissionSubject{PlanOmission: PlanOmission{Name: verdict.Technique, Importance: verdict.Importance}})
		review.OmissionVerdicts = append(review.OmissionVerdicts, verdict.OmissionVerdict)
	}
	for index, verdict := range saved.ThirdPurchaseVerdicts {
		subjects.thirdPurchases = append(subjects.thirdPurchases, purchaseSubject{path: index, build: verdict.Build, name: verdict.Name, key: verdict.Path})
		review.ThirdPurchaseVerdicts = append(review.ThirdPurchaseVerdicts, verdict.PurchaseVerdict)
	}
	var kept []string
	for _, finding := range subjects.withoutRepeats(review, saved.Findings) {
		kept = append(kept, finding.ID)
	}
	if got, want := strings.Join(kept, ","), "model.omit-jetpistol,model.omit-gigantpistol,model.omit-atamabuso,model.omit-gearbuso,model.path3-identity,model.path3-tier4-proposal,model.pierce-noeffect,model.knockback-noeffect"; got != want {
		t.Errorf("kept %s, want %s", got, want)
	}
}

// Code drops a finding only when it names exactly a verdict's subject, has
// its outcome and gives its reason word for word. A negated reason, a
// paraphrase, a finding on the same subject with its own reason, such as a
// capstone's price, another outcome, or a subject that names a second
// purchase stays (SOL-74-01).
func TestVerdictRepeatNeedsSubjectOutcomeAndSameReason(t *testing.T) {
	subjects := verdictSubjects{
		omissions:      []omissionSubject{{PlanOmission: PlanOmission{Name: "Gear 5", Importance: "major"}}},
		fifthPurchases: []purchaseSubject{{path: 1, build: "x-5-x", name: "Jet Pistol Follow-Through", key: "path2"}},
	}
	capstone := "The fifth purchase follow-up redirects the attack to another target."
	omission := "Gear 5's rubbery, cartoonish strikes that send enemies flying are expressible as knockback and splash, so the omission does not hold."
	review := SemanticReview{
		OmissionVerdicts:      []OmissionVerdict{{Technique: "Gear 5", Outcome: "fail", Reason: omission}},
		FifthPurchaseVerdicts: []FifthPurchaseVerdict{{Build: "x-5-x", Outcome: "fail", Reason: capstone}},
	}
	finding := func(id, outcome, subject, message string) Finding {
		return Finding{ID: id, Method: "model", Outcome: outcome, Subject: subject, Message: message}
	}
	findings := []Finding{
		finding("model.repeat-capstone", "fail", "x-5-x, Jet Pistol Follow-Through", "  the fifth purchase follow-up  REDIRECTS the attack to another target"),
		finding("model.repeat-omission", "fail", "Omission: Gear 5", omission),
		finding("model.negated", "fail", "x-5-x", "The fifth purchase follow-up does not redirect the attack to another target."),
		finding("model.paraphrase", "fail", "Omission: Gear 5", "Gear 5 sends enemies flying with rubbery strikes, which knockback and splash express, so omitting it does not hold."),
		finding("model.capstone-price", "fail", "x-5-x", "x-5-x costs 12,000 Gold for one extra hit at half damage, while x-4-x costs 4,200; price x-5-x for its gain."),
		finding("model.capstone-pass", "pass", "x-5-x, Jet Pistol Follow-Through", capstone),
		finding("model.two-capstones", "fail", "x-5-x and x-x-5", capstone),
	}
	var kept []string
	for _, f := range subjects.withoutRepeats(review, findings) {
		kept = append(kept, f.ID)
	}
	if got, want := strings.Join(kept, ","), "model.negated,model.paraphrase,model.capstone-price,model.capstone-pass,model.two-capstones"; got != want {
		t.Errorf("kept %s, want %s", got, want)
	}
}

// A free-form finding repeats a proposal verdict only when it names the
// purchase and the proposal, has the verdict's outcome and gives its reason
// word for word.
func TestProposalVerdictRepeat(t *testing.T) {
	subject := proposalSubject{
		purchaseSubject: purchaseSubject{path: 2, build: "x-x-4", name: "Python Continuous Stretch", key: "path3"},
		tier:            4,
		proposal:        m.ProposedMechanic{Name: "Continuous momentum"},
	}
	subjects := verdictSubjects{proposals: []proposalSubject{subject}}
	reason := "Continuous momentum names a state with no effect a player could see."
	review := SemanticReview{ProposalVerdicts: []ProposalVerdict{{Build: "x-x-4", Proposal: "Continuous momentum", Outcome: "fail", Reason: reason}}}
	finding := func(id, subject, outcome, message string) Finding {
		return Finding{ID: id, Method: "model", Subject: subject, Outcome: outcome, Message: message}
	}
	var kept []string
	for _, f := range subjects.withoutRepeats(review, []Finding{
		finding("model.repeat", "x-x-4 Python Continuous Stretch: Continuous momentum", "fail", reason+"."),
		finding("model.purchase", "x-x-4 Python Continuous Stretch", "fail", reason),
		finding("model.other", "x-x-4: Continuous momentum", "fail", "Its price does not fit."),
		finding("model.outcome", "path3, x-x-4: Continuous momentum", "unresolved", reason),
	}) {
		kept = append(kept, f.ID)
	}
	if got := strings.Join(kept, ","); got != "model.purchase,model.other,model.outcome" {
		t.Errorf("kept %s", got)
	}
	if f := subjects.findings(review); len(f) != 1 || f[0].ID != "verdict.proposal.path3.tier4.1" || f[0].Subject != "path3, x-x-4 Python Continuous Stretch: Continuous momentum" || f[0].Rule != ProposalVerdictRule || f[0].Severity != "error" {
		t.Errorf("recorded %+v", f)
	}
}
