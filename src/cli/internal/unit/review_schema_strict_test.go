package unit_test

import (
	"encoding/json"
	"slices"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// strictGaps lists every object in a JSON schema whose properties are not
// all required. Strict structured output rejects such a schema: OpenRouter
// answered Escanor's v38 review, whose Request had no required concepts,
// with a 400 while requiredConceptVerdicts stayed an optional property.
func strictGaps(node any, path string) []string {
	var gaps []string
	switch v := node.(type) {
	case map[string]any:
		if properties, ok := v["properties"].(map[string]any); ok {
			required := map[string]bool{}
			for _, name := range v["required"].([]any) {
				required[name.(string)] = true
			}
			for name := range properties {
				if !required[name] {
					gaps = append(gaps, path+"."+name)
				}
			}
		}
		for key, child := range v {
			gaps = append(gaps, strictGaps(child, path+"."+key)...)
		}
	case []any:
		for _, child := range v {
			gaps = append(gaps, strictGaps(child, path)...)
		}
	}
	return gaps
}

func TestReviewSchemaRequiresEveryProperty(t *testing.T) {
	stages, err := fixture.Build()
	if err != nil {
		t.Fatal(err)
	}
	check := func(name string, checked unit.Checked) {
		var schema any
		if err := json.Unmarshal([]byte(s.Stringify(unit.BlueprintReviewRequest(checked).Schema)), &schema); err != nil {
			t.Fatal(err)
		}
		gaps := strictGaps(schema, "")
		slices.Sort(gaps)
		if len(gaps) > 0 {
			t.Errorf("%s: optional properties in the review schema: %v", name, gaps)
		}
	}
	check("without required concepts", stages.Checked)
	required := stages.Checked
	required.Draft.Prepared.Request.RequiredConcepts = []unit.RequiredConcept{{Name: "Fan Club"}}
	check("with required concepts", required)
}
