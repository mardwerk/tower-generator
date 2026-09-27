package unit

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// The v34 Luffy review (OPUS-NET-61-13) spent five of its eight free-form
// findings on four omission fails and the x-x-3 fail that its verdicts
// already recorded. Those five are dropped; the other three, a proposal on
// x-x-4 and two misread promisesWithoutEffect on 4-x-x and 5-x-x, are kept.
func TestLuffyV34VerdictRepeatsAreDropped(t *testing.T) {
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
		subjects.omissions = append(subjects.omissions, PlanOmission{Name: verdict.Technique, Importance: verdict.Importance})
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
	if got, want := strings.Join(kept, ","), "model.path3-tier4-proposal,model.pierce-noeffect,model.knockback-noeffect"; got != want {
		t.Errorf("kept %s, want %s", got, want)
	}
}

// Code drops a finding only when it names exactly a verdict's subject, has
// its outcome and shares its reason. A finding on the same subject with its
// own reason, such as a capstone's price, another outcome, or a subject that
// names a second purchase stays.
func TestVerdictRepeatNeedsSubjectOutcomeAndReason(t *testing.T) {
	subjects := verdictSubjects{
		omissions:      []PlanOmission{{Name: "Gear 5", Importance: "major"}},
		fifthPurchases: []purchaseSubject{{path: 1, build: "x-5-x", name: "Jet Pistol Follow-Through", key: "path2"}},
	}
	review := SemanticReview{
		OmissionVerdicts:      []OmissionVerdict{{Technique: "Gear 5", Outcome: "fail", Reason: "Gear 5's rubbery, cartoonish strikes that send enemies flying are expressible as knockback and splash, so the omission does not hold."}},
		FifthPurchaseVerdicts: []FifthPurchaseVerdict{{Build: "x-5-x", Outcome: "fail", Reason: "x-5-x adds a bounded follow-up, and x-x-5 also buys a follow-up; neither gives a distinct play reason, so the two capstones buy the same capability."}},
	}
	finding := func(id, outcome, subject, message string) Finding {
		return Finding{ID: id, Method: "model", Outcome: outcome, Subject: subject, Message: message}
	}
	findings := []Finding{
		finding("model.repeat-capstone", "fail", "x-5-x, Jet Pistol Follow-Through", "x-5-x adds a bounded follow-up that x-x-5 also buys, so the two capstones buy the same capability with no distinct play reason."),
		finding("model.repeat-omission", "fail", "Omission: Gear 5", "Gear 5 sends enemies flying with rubbery strikes, which knockback and splash express, so omitting it does not hold."),
		finding("model.capstone-price", "fail", "x-5-x", "x-5-x costs 12,000 Gold for one extra hit at half damage, while x-4-x costs 4,200; price x-5-x for its gain."),
		finding("model.capstone-pass", "pass", "x-5-x, Jet Pistol Follow-Through", "x-5-x adds a bounded follow-up that x-x-5 also buys, so the two capstones buy the same capability with no distinct play reason."),
		finding("model.two-capstones", "fail", "x-5-x and x-x-5", "x-5-x adds a bounded follow-up that x-x-5 also buys, so the two capstones buy the same capability with no distinct play reason."),
	}
	var kept []string
	for _, f := range subjects.withoutRepeats(review, findings) {
		kept = append(kept, f.ID)
	}
	if got, want := strings.Join(kept, ","), "model.capstone-price,model.capstone-pass,model.two-capstones"; got != want {
		t.Errorf("kept %s, want %s", got, want)
	}
}
