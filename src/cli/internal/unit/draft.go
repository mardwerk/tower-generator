package unit

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

var (
	countIssue    = regexp.MustCompile(`^(builds\.[^:]+: Resolved (projectiles|pierce) is \S+; must be a positive integer\.)`)
	countModifier = regexp.MustCompile(`Purchased modifiers: (.*?)\. Starting base `)
)

// failureSummary keeps surfaced failures readable; repairs keep the full issues.
func failureSummary(issues []string) string {
	type count struct {
		first, modifier string
		occurrences     int
	}
	var order []string
	counts := map[string]*count{}
	var other []string
	for _, issue := range issues {
		match := countIssue.FindStringSubmatch(issue)
		if match == nil {
			other = append(other, issue)
			continue
		}
		stat := match[2]
		if c, ok := counts[stat]; ok {
			c.occurrences++
			continue
		}
		c := &count{first: match[1], occurrences: 1}
		if mod := countModifier.FindStringSubmatch(issue); mod != nil && mod[1] != "" {
			c.modifier = strings.Split(mod[1], "; ")[0]
		}
		counts[stat] = c
		order = append(order, stat)
	}
	var parts []string
	for _, stat := range order {
		c := counts[stat]
		lines := []string{c.first, fmt.Sprintf("Correct the authored %s upgrades so every legal build has a positive whole-number count. Counts are never rounded.", stat)}
		if c.modifier != "" {
			lines = append(lines, "First purchased modifier: "+c.modifier+".")
		}
		if c.occurrences > 1 {
			lines = append(lines, fmt.Sprintf("%d additional %s count checks failed.", c.occurrences-1, stat))
		}
		parts = append(parts, strings.Join(lines, " "))
	}
	for i, issue := range other {
		if i == 4 {
			break
		}
		parts = append(parts, s.SliceUTF16(issue, 0, 220))
	}
	return strings.Join(parts, " ")
}

func issueStrings(issues []s.Issue) []string {
	out := make([]string, len(issues))
	for i, issue := range issues {
		out[i] = issue.PathString() + ": " + issue.Message
	}
	return out
}

func cancelled(ctx context.Context, attempts []Attempt) error {
	if ctx.Err() == nil {
		return nil
	}
	return StageFailure(ctx.Err(), "draft", totalUsage(attempts), false)
}

func withUsage(a Attempt, usage *Usage) Attempt {
	if usage != nil {
		a.Usage = usage
	}
	return a
}

// planDesign makes the planning call and at most repairs full-plan retries.
// A plan attempt that fails only on items a targeted correction can fix is
// followed by one such correction (planCorrections), which does not use
// the retry budget and is not sent when the budget is 0. A full-plan retry
// that fails only on fifth purchases with no new behavior, with at most
// such items besides, is followed by one capstone correction
// (capstoneCorrectionFor), whose output is merged into the retry's so that
// no other milestone changes.
func planDesign(ctx context.Context, prepared Prepared, model Model, repairs int, attempts *[]Attempt) (DesignPlan, error) {
	correction := ""
	purpose, full, corrected, capstoned := "plan", 0, false, false
	var capstone capstoneCorrection
	var previous any
	for purpose == PlanCorrectionPurpose || purpose == CapstoneCorrectionPurpose || full <= repairs {
		if purpose == "plan" {
			full++
		}
		if err := cancelled(ctx, *attempts); err != nil {
			return DesignPlan{}, err
		}
		request, err := DesignPlanRequest(prepared)
		var response ModelResponse
		if err == nil {
			request.Prompt += correction
			response, err = model.Generate(ctx, request)
			if err == nil && ctx.Err() != nil {
				err = ctx.Err()
			}
		}
		var plan DesignPlan
		output := response.Output
		if err == nil {
			if purpose == CapstoneCorrectionPurpose {
				output = capstone.merge(previous, output)
			}
			plan, err = DecodeDesignPlan(output, &prepared.Request)
			if err == nil {
				*attempts = append(*attempts, withUsage(Attempt{Number: len(*attempts) + 1, Purpose: purpose, Issues: []string{}}, response.Usage))
				return plan, nil
			}
		}
		var validation *s.Error
		issues := []string{}
		if errors.As(err, &validation) {
			compact := false
			if obj, ok := output.(*s.Object); ok && obj.Has("contract") {
				compact = true
			}
			for _, issue := range validation.Issues {
				path := issue.Path
				if compact && len(path) >= 3 && path[0] == "upgradeIntents" {
					path = append([]any{"paths", path[1], "milestones", path[2]}, path[3:]...)
				}
				message := issue.Message
				// Quoting the rejected text lets the correction find it
				// without counting array positions.
				if value, ok := valueAt(output, issue.Path).(string); ok {
					message += ". Current value: " + s.Stringify(s.SliceUTF16(value, 0, 160))
				}
				issues = append(issues, s.Issue{Path: path}.PathString()+": "+message)
			}
		}
		billed := response.Usage
		var modelErr *ModelError
		if errors.As(err, &modelErr) && modelErr.Usage != nil {
			billed = modelErr.Usage
		}
		rejected := withUsage(Attempt{Number: len(*attempts) + 1, Purpose: purpose, Issues: issues}, billed)
		if validation != nil {
			rejected.Output = output
		}
		*attempts = append(*attempts, rejected)
		if validation == nil {
			failure := StageFailure(err, "draft", totalUsage(*attempts), false)
			return DesignPlan{}, &ModelError{Message: failure.Message, Usage: totalUsage(*attempts), Failure: failure.Failure, Cause: failure, Evidence: failureEvidence(&prepared.Request, nil, *attempts)}
		}
		if repairs > 0 && !corrected && onlyCorrectable(validation.Issues) {
			if items := planCorrections(output, &prepared.Request); len(items) > 0 {
				purpose, corrected = PlanCorrectionPurpose, true
				correction = targetedPlanCorrection(items, output)
				continue
			}
		}
		// Only after a full-plan retry: the first plan gets the retry.
		if purpose == "plan" && full > 1 && !capstoned && onlyCapstone(validation.Issues) {
			if c, ok := capstoneCorrectionFor(output, &prepared.Request, validation.Issues); ok {
				purpose, capstoned, capstone, previous = CapstoneCorrectionPurpose, true, c, output
				correction = c.prompt
				continue
			}
		}
		purpose = "plan"
		// A Luffy retry on #27 returned only the one repertoire entry an
		// issue named, and the shortened plan failed.
		correction = "\n\nCorrect this invalid design plan while retaining supported character identity. Return the complete plan with every field and entry, changing only what the issues require: " +
			s.Stringify(s.NewObject().Set("issues", stringList(issues)).Set("previous", output))
	}
	last := (*attempts)[len(*attempts)-1]
	message := "The character design plan could not be validated. " + failureSummary(last.Issues)
	return DesignPlan{}, &ModelError{Message: message, Usage: totalUsage(*attempts), Failure: &Failure{Code: CodeOutputInvalid, Stage: "draft", Message: message}, Evidence: failureEvidence(&prepared.Request, nil, *attempts)}
}

// failureEvidence copies the attempts of a failed draft with the accepted
// plan, if any, and the source passages the model was given.
func failureEvidence(request *Request, plan *DesignPlan, attempts []Attempt) *FailureEvidence {
	evidence := &FailureEvidence{Attempts: append([]Attempt{}, attempts...), SourcePassages: AuthorEvidence(request)}
	if plan != nil {
		retained := *plan
		evidence.Plan = &retained
	}
	return evidence
}

// valueAt follows a validation issue path through decoded model output.
func valueAt(value any, path []any) any {
	for _, key := range path {
		switch k := key.(type) {
		case string:
			object, ok := value.(*s.Object)
			if !ok {
				return nil
			}
			value, _ = object.Get(k)
		case int:
			list, ok := value.([]any)
			if !ok || k < 0 || k >= len(list) {
				return nil
			}
			value = list[k]
		default:
			return nil
		}
	}
	return value
}

func stringList(values []string) []any {
	out := make([]any, len(values))
	for i, v := range values {
		out[i] = v
	}
	return out
}

// DraftUnit plans a unit, authors its mechanics, repairs within the budget and
// publishes it only when every check passes.
func DraftUnit(ctx context.Context, prepared Prepared, model Model, options Options) (Draft, error) {
	if _, err := ParsePrepared(s.FromGoValue(prepared)); err != nil {
		return Draft{}, err
	}
	if err := VerifyPrepared(prepared); err != nil {
		return Draft{}, err
	}
	if err := ctx.Err(); err != nil {
		return Draft{}, err
	}
	if prepared.Request.MechanicsDefinition == nil {
		return Draft{}, errors.New("This request has no mechanics Definition. Prepare it again under a Profile.")
	}
	return draftBlueprint(ctx, prepared, model, options)
}

func draftBlueprint(ctx context.Context, prepared Prepared, model Model, options Options) (Draft, error) {
	repairs := 1
	if options.MaxRepairAttempts != nil {
		repairs = *options.MaxRepairAttempts
	}
	if repairs < 0 || repairs > 2 {
		return Draft{}, errors.New("maxRepairAttempts must be 0, 1 or 2.")
	}
	request := &prepared.Request
	definition := *request.MechanicsDefinition
	startedAt := options.now()
	attempts := []Attempt{}
	var previous any
	issues := []string{}
	plan, err := planDesign(ctx, prepared, model, repairs, &attempts)
	if err != nil {
		return Draft{}, err
	}
	fail := func(err error, usage *Usage, _ bool, purpose string) error {
		var validation *s.Error
		invalid := errors.As(err, &validation)
		failedUsage := usage
		var modelErr *ModelError
		if errors.As(err, &modelErr) && modelErr.Usage != nil {
			failedUsage = modelErr.Usage
		}
		attempts = append(attempts, withUsage(Attempt{Number: len(attempts) + 1, Purpose: purpose, Issues: []string{}}, failedUsage))
		failure := StageFailure(err, "draft", totalUsage(attempts), invalid)
		return &ModelError{Message: failure.Message, Usage: totalUsage(attempts), Failure: failure.Failure, Cause: failure, Evidence: failureEvidence(request, &plan, attempts)}
	}
	for attempt := 0; attempt <= repairs; attempt++ {
		if err := cancelled(ctx, attempts); err != nil {
			return Draft{}, err
		}
		purpose := "design"
		if attempt > 0 {
			purpose = "repair"
		}
		var repair *TierRepair
		if attempt > 0 {
			if repair, err = TargetedTierRepair(request, previous, issues); err != nil {
				return Draft{}, fail(err, nil, false, purpose)
			}
		}
		var call ModelRequest
		if repair != nil {
			call = repair.Request
			call.Prompt += "\n\nPreserve the retained character plan while correcting these tiers. Do not trade its branch purpose for easier arithmetic: " + s.Stringify(MechanicsPlan(plan))
		} else if call, err = blueprintRequest(prepared, previous, issues, plan); err != nil {
			return Draft{}, fail(err, nil, false, purpose)
		}
		response, err := model.Generate(ctx, call)
		if err == nil && ctx.Err() != nil {
			err = ctx.Err()
		}
		if err != nil {
			return Draft{}, fail(err, response.Usage, false, purpose)
		}
		usage := response.Usage
		authored := response.Output
		var decodeErr error
		if repair != nil {
			var merged *s.Object
			merged, decodeErr = repair.Apply(response.Output)
			authored = merged
		}
		var blueprint m.Blueprint
		valid := false
		var validation *s.Error
		if decodeErr == nil {
			// A schema failure still gets the semantic checks when it
			// is only numbers outside their bounds (mechanicsIssues).
			previous = BindDesignPlan(authored, plan)
			blueprint, issues, valid, decodeErr = mechanicsIssues(previous, request, plan)
		} else if errors.As(decodeErr, &validation) {
			issues = issueStrings(validation.Issues)
		}
		if decodeErr != nil && validation == nil {
			return Draft{}, fail(decodeErr, usage, false, purpose)
		}
		var draft *Draft
		if valid && len(issues) == 0 {
			candidate, err := CompileBlueprint(blueprint, *request)
			if err != nil {
				return Draft{}, fail(err, usage, false, purpose)
			}
			evaluation, err := EvaluateUnitDesign(blueprint, &plan, definition)
			if err != nil {
				return Draft{}, fail(err, usage, false, purpose)
			}
			planCopy := plan
			d := Draft{
				SchemaVersion: prepared.SchemaVersion, Kind: "draft", Prepared: prepared, Candidate: candidate,
				Run: Run{ID: options.id(), ModelID: model.ID(), StartedAt: startedAt, CompletedAt: options.now(), DesignPlan: &planCopy, DesignEvaluation: evaluation},
			}
			value := s.FromGoValue(d)
			if err := s.ParseInto(Versioned(value, DraftSchema, DraftSchemaV2), value, &d); err != nil {
				return Draft{}, fail(err, usage, true, purpose)
			}
			checked, err := CheckDraft(d)
			if err != nil {
				return Draft{}, fail(err, usage, false, purpose)
			}
			issues = []string{}
			for _, f := range checked.Findings {
				if f.Outcome == "fail" {
					issues = append(issues, f.Subject+": "+f.Message)
				}
			}
			draft = &d
		}
		attempt := withUsage(Attempt{Number: len(attempts) + 1, Purpose: purpose, Issues: issues}, usage)
		if len(issues) > 0 {
			attempt.Output = authored
		}
		attempts = append(attempts, attempt)
		if draft != nil && len(issues) == 0 {
			all := append([]Attempt(nil), attempts...)
			for i := range all {
				all[i].Output = nil
			}
			draft.Run.Attempts = &all
			draft.Run.Usage = totalUsage(attempts)
			return *draft, nil
		}
	}
	designAttempts := 0
	for _, a := range attempts {
		if a.Purpose != "plan" && a.Purpose != PlanCorrectionPurpose && a.Purpose != CapstoneCorrectionPurpose {
			designAttempts++
		}
	}
	message := fmt.Sprintf("The draft still failed mechanics checks after %d attempts. %s No invalid Unit was published.", designAttempts, failureSummary(issues))
	return Draft{}, &ModelError{Message: message, Usage: totalUsage(attempts), Failure: &Failure{Code: CodeOutputInvalid, Stage: "draft", Message: message}, Evidence: failureEvidence(request, &plan, attempts)}
}

func blueprintRequest(prepared Prepared, previous any, issues []string, plan DesignPlan) (ModelRequest, error) {
	request := &prepared.Request
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
		documents = append(documents, entry.Set("access", s.FromGoValue(d.Origin)))
	}
	evidence := AuthorEvidence(request)
	context := s.NewObject().
		Set("designPlan", MechanicsPlan(plan)).
		Set("character", s.FromGoValue(request.Character)).
		Set("task", request.Task).
		Set("constraints", s.FromGoValue(request.Constraints))
	if request.MechanicsDefinition != nil {
		context.Set("definition", s.FromGoValue(request.MechanicsDefinition))
	}
	context.
		Set("documents", documents).
		Set("evidenceSpans", s.FromGoValue(evidence)).
		Set("sourceScope", s.NewObject().
			Set("selectedPassages", float64(len(evidence))).
			Set("availablePassages", float64(len(EvidenceSpans(request)))).
			Set("note", draftSourceScope)).
		Set("previous", previousValue(request)).
		Set("previousFindings", previousFindings(request)).
		Set("feedback", nullableString(request.Feedback))
	budget, form, example := draftBudget, draftStatusForm, draftExample
	if isV2(request) {
		budget, form, example = draftStatusBudgetV2, draftStatusFormV2, draftExampleV2
	}
	budget = fmt.Sprintf(budget, budgetSentence(request))
	prompt := []string{draftPlan, draftPromises, draftStyle, draftNames, draftOwnership, draftShape, draftTruth, budget, form, draftArithmetic, draftExtensions, CountArithmeticGuidance}
	if d := request.MechanicsDefinition; d != nil {
		prompt = append(prompt, BuildShape(*d))
	}
	prompt = append(prompt, VocabularyGuidance(request)...)
	prompt = append(prompt, DesignGuidance(request)...)
	prompt = append(prompt, draftBoost, example, draftUnsupported, s.Stringify(context))
	if previous != nil {
		shown := issues
		if len(shown) > 20 {
			shown = shown[:20]
		}
		prompt = append(prompt, draftRetry, s.Stringify(s.NewObject().Set("issues", stringList(shown)).Set("previous", previous)))
	}
	schema, err := ModelOutputJSONSchema(request)
	if err != nil {
		return ModelRequest{}, err
	}
	return ModelRequest{System: draftSystem, Prompt: strings.Join(prompt, "\n\n"), Schema: schema}, nil
}

// ParsePrepared validates a prepared request value.
func ParsePrepared(value any) (Prepared, error) {
	var p Prepared
	return p, s.ParseInto(Versioned(value, PreparedSchema, PreparedSchemaV2), value, &p)
}
