package render

import (
	m "github.com/mardwerk/unit-generator/src/cli/internal/mechanics"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// AuthoringIssues are the authoring checks that a candidate's typed
// mechanics fail today under the Definition it carries, such as a change
// with no effect. Reading views check structure only, so such a candidate
// still has its unit sheet; these issues say that today's checks would not
// accept it as made. They are computed when the artifact is read and are not
// Findings: a Result's stored Findings stay as recorded. Nil when the
// candidate has no typed mechanics, when its structure is invalid (the sheet
// then falls back to prose) or when every authoring check passes.
func AuthoringIssues(candidate unit.Candidate, definition *m.Definition) []m.Issue {
	if candidate.Blueprint == nil {
		return nil
	}
	d := m.DefaultDefinition()
	if definition != nil {
		d = *definition
	}
	issues := m.ValidateTyped(candidate.Blueprint, d)
	if len(issues) == 0 || len(m.ValidateStructure(candidate.Blueprint, d)) > 0 {
		return nil
	}
	return issues
}

// currentChecks lists the authoring issues for the detailed render, apart
// from the stored Findings; nil when there are none.
func currentChecks(view View) []string {
	issues := AuthoringIssues(view.Candidate, view.Prepared.Request.MechanicsDefinition)
	if len(issues) == 0 {
		return nil
	}
	lines := []string{
		"## Current checks",
		"",
		"Today's authoring checks flag this unit's typed mechanics. These issues are computed when the artifact is read and are not stored Findings; the Findings above stay as recorded. Every legal build still resolves to valid values, so `render` without `--details` still draws the unit sheet.",
		"",
		"| Subject | Issue |",
		"| --- | --- |",
	}
	for _, issue := range issues {
		lines = append(lines, "| "+Escape(issue.Path)+" | "+Escape(issue.Message)+" |")
	}
	return append(lines, "")
}
