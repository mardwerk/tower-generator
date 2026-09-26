package unit

import (
	"context"
	"errors"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// ParseChecked validates a checked artifact value.
func ParseChecked(value any) (Checked, error) {
	var c Checked
	return c, s.ParseInto(Versioned(value, CheckedSchema, CheckedSchemaV2), value, &c)
}

// ParseResult validates a Result value.
func ParseResult(value any) (Result, error) {
	var r Result
	return r, s.ParseInto(Versioned(value, ResultSchema, ResultSchemaV2), value, &r)
}

// ErrNoBlueprint means an older prose draft reached a stage that needs typed mechanics.
var ErrNoBlueprint = errors.New("This draft has no typed blueprint. It can be rendered, but reviewing needs a unit generated under a Profile.")

// BlueprintReviewRequest asks a separate model call for a semantic review.
func BlueprintReviewRequest(checked Checked) ModelRequest {
	request := &checked.Draft.Prepared.Request
	candidate := checked.Draft.Candidate
	blueprint := candidate.Blueprint
	var documentIDs []string
	for _, d := range request.Documents {
		documentIDs = append(documentIDs, d.ID)
	}
	finding := ReviewFindingSchema().Extend(s.F("evidence", s.Array(s.Enum(documentIDs...))))
	schema := SemanticReviewSchema.Extend(s.F("findings", s.Array(finding).Max(8)))
	mechanicsID := "mechanics:undefined"
	if request.MechanicsDefinition != nil {
		mechanicsID = "mechanics:" + request.MechanicsDefinition.ID
	}
	documents := []any{}
	for _, d := range request.Documents {
		if d.ID == mechanicsID {
			continue
		}
		entry := s.NewObject().Set("id", d.ID).Set("kind", d.Kind)
		if d.Kind != "source" {
			entry.Set("text", d.Text)
		}
		documents = append(documents, entry.Set("origin", s.FromGoValue(d.Origin)))
	}
	context := s.NewObject()
	if checked.Draft.Run.DesignPlan != nil {
		context.Set("designPlan", s.FromGoValue(checked.Draft.Run.DesignPlan))
	}
	if checked.Draft.Run.DesignEvaluation != nil {
		// The retained evidence plus the time-averaged Active rates and side
		// purchase gains derived from it; the saved draft is unchanged.
		context.Set("purchaseEvidence", ReviewPurchaseEvidence(checked.Draft.Run.DesignEvaluation))
	}
	comparisons := []any{}
	if blueprint != nil {
		var vocabulary *m.Vocabulary
		if d := request.MechanicsDefinition; d != nil {
			terms := d.Terms()
			vocabulary = &terms
		}
		comparisons = m.CompareCapstonePurchasesWith(blueprint, vocabulary)
	}
	// The raw capstone comparisons are in purchaseEvidence per path; the
	// review also gets their checked ordering.
	context.Set("purchaseComparisonOrdering", capstoneOrdering(comparisons)).
		Set("character", s.FromGoValue(request.Character)).
		Set("task", request.Task).
		Set("constraints", s.FromGoValue(request.Constraints)).
		Set("previous", previousValue(request)).
		Set("previousFindings", previousFindings(request)).
		Set("feedback", nullableString(request.Feedback))
	if request.MechanicsDefinition != nil {
		context.Set("definition", s.FromGoValue(request.MechanicsDefinition))
		if blueprint != nil {
			context.Set("legalBuilds", LegalBuildFacts(blueprint, *request.MechanicsDefinition))
		}
	}
	evidence := AuthorEvidence(request)
	context.Set("documents", documents).
		Set("sourcePassages", s.FromGoValue(evidence)).
		Set("sourceScope", s.NewObject().
			Set("selected", float64(len(evidence))).
			Set("available", float64(len(EvidenceSpans(request)))).
			Set("note", "Bounded source selection; do not claim exhaustive canon coverage.")).
		Set("mechanicsEvidenceId", mechanicsID)
	pathEvidence := []any{}
	if blueprint != nil {
		context.Set("sourceFacts", s.FromGoValue(blueprint.SourceFacts))
		for index, path := range m.PathKeys {
			design := blueprint.Paths.At(index)
			entry := s.NewObject().Set("path", path)
			if design.Specialization != "" {
				entry.Set("specialization", design.Specialization)
			}
			quotes := []any{}
			for _, i := range design.SourceFactIndices {
				if i < len(blueprint.SourceFacts) {
					quotes = append(quotes, s.FromGoValue(blueprint.SourceFacts[i]))
				} else {
					quotes = append(quotes, nil)
				}
			}
			pathEvidence = append(pathEvidence, entry.Set("quotes", quotes))
		}
	}
	context.Set("pathEvidence", pathEvidence)
	var paths []any
	for _, p := range candidate.Paths {
		var tiers []any
		for _, t := range p.Tiers {
			tiers = append(tiers, s.NewObject().Set("tier", float64(t.Tier)).Set("name", t.Name).Set("benefit", t.Benefit))
		}
		if tiers == nil {
			tiers = []any{}
		}
		paths = append(paths, s.NewObject().Set("id", p.ID).Set("name", p.Name).Set("theme", p.Theme).Set("tiers", tiers))
	}
	if paths == nil {
		paths = []any{}
	}
	context.Set("unit", s.NewObject().
		Set("role", candidate.Role).
		Set("basicAttack", s.FromGoValue(candidate.BasicAttack)).
		Set("paths", paths).
		Set("abilities", s.FromGoValue(candidate.Abilities)).
		Set("mechanics", s.FromGoValue(candidate.Mechanics)).
		Set("unresolvedQuestions", s.FromGoValue(candidate.UnresolvedQuestions)))
	failed := []Finding{}
	for _, f := range checked.Findings {
		if f.Outcome == "fail" {
			failed = append(failed, f)
		}
	}
	context.Set("deterministicFindings", s.FromGoValue(failed))
	statuses := reviewStatusesV1
	if isV2(request) {
		statuses = reviewStatusesV2
	}
	prompt := []string{reviewStyle, reviewScope, reviewGrounding, reviewPlan, reviewPrivate, reviewAdaptation, reviewPeriod, reviewReading, statuses, reviewPolicy, reviewFindings, s.Stringify(context)}
	return ModelRequest{System: reviewSystem, Prompt: strings.Join(prompt, "\n\n"), Schema: s.JSONSchema(schema)}
}

// capstoneOrdering gives review a checked numerical comparison. The raw
// metrics remain in purchaseEvidence; this is prompt context, not saved
// design evidence, so older drafts retain their exact evidence.
func capstoneOrdering(comparisons []any) []any {
	out := []any{}
	metric := func(object *s.Object, key string) (float64, bool) {
		value, ok := object.Get(key)
		number, numeric := value.(float64)
		return number, ok && numeric
	}
	for _, value := range comparisons {
		comparison := value.(*s.Object)
		path, _ := comparison.Get("path")
		tier5, _ := comparison.Get("tier5")
		copies, _ := comparison.Get("sameBudgetTier4Copies")
		tier5Metrics, _ := tier5.(*s.Object).Get("metrics")
		bounds, _ := copies.(*s.Object).Get("additiveThroughputUpperBounds")
		row := s.NewObject().Set("path", path)
		for _, key := range []string{"direct damage rate", "group damage rate upper bound"} {
			tier5Value, validTier5 := metric(tier5Metrics.(*s.Object), key)
			copiesValue, validCopies := metric(bounds.(*s.Object), key)
			if !validTier5 || !validCopies {
				continue
			}
			ordering := "equal"
			if tier5Value > copiesValue {
				ordering = "higher"
			} else if tier5Value < copiesValue {
				ordering = "lower"
			}
			var ratio any
			if copiesValue > 0 {
				ratio = tier5Value / copiesValue
			}
			row.Set(key, s.NewObject().Set("ordering", ordering).Set("tier5ToCopiesRatio", ratio))
		}
		out = append(out, row)
	}
	return out
}

func invalidReview() *ModelError {
	message := "The model returned a review with invalid finding or evidence references. The draft is retained. Retry the review or choose another model."
	return &ModelError{Message: message, Failure: &Failure{Code: CodeOutputInvalid, Message: message, Stage: "review"}}
}

// ReviewDraft adds a model's semantic review to a checked draft.
func ReviewDraft(ctx context.Context, input Checked, model Model, options Options) (Result, error) {
	checked, err := ParseChecked(s.FromGoValue(input))
	if err != nil {
		return Result{}, err
	}
	verified, err := CheckDraft(checked.Draft)
	if err != nil {
		return Result{}, err
	}
	if s.Stringify(s.FromGoValue(verified.Findings)) != s.Stringify(s.FromGoValue(checked.Findings)) {
		return Result{}, errors.New("Checked findings do not match deterministic checks of the retained draft")
	}
	if err := ctx.Err(); err != nil {
		return Result{}, err
	}
	if checked.Draft.Candidate.Blueprint == nil {
		return Result{}, ErrNoBlueprint
	}
	startedAt := options.now()
	request := BlueprintReviewRequest(checked)
	definition := checked.Draft.Prepared.Request.MechanicsDefinition
	var calls []Attempt
	review, err := reviewOnce(ctx, checked, model, request, &calls)
	if err == nil && definition != nil {
		// A review that cites builds the Definition does not allow, or
		// resolved values that are wrong, is corrected once, then rejected.
		// Findings whose citations hold must come back unchanged.
		citations := newReviewCitations(checked.Draft.Candidate.Blueprint, *definition)
		if problems := citations.problems(review); len(problems) > 0 {
			request.Prompt += correctionPrompt(review, problems)
			var corrected SemanticReview
			if corrected, err = reviewOnce(ctx, checked, model, request, &calls); err == nil {
				if review, err = keepCheckedFindings(review, corrected, problems); err == nil {
					err = rejectedCitations(citations.problems(review))
				}
			}
		}
	}
	usage := totalUsage(calls)
	if err != nil {
		var validation *s.Error
		return Result{}, StageFailure(err, "review", usage, errors.As(err, &validation))
	}
	result := Result{
		SchemaVersion: checked.SchemaVersion, Kind: "result", ID: options.id(),
		Prepared: checked.Draft.Prepared, Candidate: checked.Draft.Candidate,
		Findings:      append(append([]Finding{}, checked.Findings...), review.Findings...),
		ReviewSummary: review.Summary,
		Run: ResultRuns{
			Draft:  checked.Draft.Run,
			Review: Run{ID: options.id(), ModelID: model.ID(), StartedAt: startedAt, CompletedAt: options.now(), Usage: usage},
		},
	}
	value := s.FromGoValue(result)
	if err := s.ParseInto(Versioned(value, ResultSchema, ResultSchemaV2), value, &result); err != nil {
		return Result{}, err
	}
	return result, nil
}

// reviewOnce makes one review call and checks its finding IDs and evidence
// references. calls records the billed call.
func reviewOnce(ctx context.Context, checked Checked, model Model, request ModelRequest, calls *[]Attempt) (SemanticReview, error) {
	response, err := model.Generate(ctx, request)
	*calls = append(*calls, withUsage(Attempt{Number: len(*calls) + 1, Purpose: "review", Issues: []string{}}, response.Usage))
	var review SemanticReview
	if err == nil {
		err = s.ParseInto(SemanticReviewSchema, response.Output, &review)
	}
	if err == nil {
		documents := map[string]bool{}
		for _, d := range checked.Draft.Prepared.Request.Documents {
			documents[d.ID] = true
		}
		ids := map[string]bool{}
		for _, f := range review.Findings {
			if !strings.HasPrefix(f.ID, "model.") || ids[f.ID] {
				err = invalidReview()
				break
			}
			ids[f.ID] = true
			for _, id := range f.Evidence {
				if !documents[id] {
					err = invalidReview()
				}
			}
			if err != nil {
				break
			}
		}
	}
	if err == nil && ctx.Err() != nil {
		err = ctx.Err()
	}
	return review, err
}

// Author prepares, drafts, checks and reviews one request.
func Author(ctx context.Context, request any, model Model, options Options) (Result, error) {
	prepared, err := Prepare(request)
	if err != nil {
		return Result{}, err
	}
	draft, err := DraftUnit(ctx, prepared, model, options)
	if err != nil {
		return Result{}, err
	}
	checked, err := CheckDraft(draft)
	if err != nil {
		return Result{}, err
	}
	return ReviewDraft(ctx, checked, model, options)
}
