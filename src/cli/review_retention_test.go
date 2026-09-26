package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/mardwerk/unit-generator/src/cli/internal/fixture"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

// sequencedCodex writes a stand-in codex that answers call n with outputs[n].
func sequencedCodex(t *testing.T, dir string, outputs ...any) string {
	t.Helper()
	for i, output := range outputs {
		if err := os.WriteFile(filepath.Join(dir, "answer-"+strconv.Itoa(i+1)+".json"), []byte(s.Stringify(output)), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	counter := filepath.Join(dir, "calls")
	codex := filepath.Join(dir, "codex")
	script := "#!/bin/sh\nout=\"\"\nwhile [ $# -gt 0 ]; do\n  if [ \"$1\" = --output-last-message ]; then out=\"$2\"; shift; fi\n  shift\ndone\ncat >/dev/null\n" +
		"n=$(( $(cat " + counter + " 2>/dev/null || echo 0) + 1 ))\necho $n > " + counter + "\ncp " + dir + "/answer-$n.json \"$out\"\n"
	if err := os.WriteFile(codex, []byte(script), 0o700); err != nil {
		t.Fatal(err)
	}
	return codex
}

// Run 13c through author: the correction withdraws a finding whose only
// fault is 5-x-2. The review is rejected, no Result is written, and the
// checked draft is kept beside the output so review can run again.
func TestFailedReviewsKeepTheCheckedDraft(t *testing.T) {
	dir := scratch(t)
	request := filepath.Join(dir, "request.json")
	raw, _ := fixture.JSON("request")
	_ = os.WriteFile(request, []byte(s.Stringify(raw)), 0o600)
	plan, _ := fixture.JSON("plan")
	mechanics, _ := fixture.JSON("mechanics")
	review, _ := fixture.JSON("review")
	findings, _ := review.(*s.Object).Get("findings")
	kept := findings.([]any)[0].(*s.Object)
	snake := s.Clone(kept).(*s.Object).Set("id", "model.snake-crosspath").Set("outcome", "fail").Set("subject", "path3 x-2-x").
		Set("message", "At 0-5-2, the x-2-x side purchase adds more than x-x-5 for only 240 Gold.").
		Set("action", "Reassess path3 5-x-2's payoff or cost relative to x-2-x.").Set("facts", []any{})
	first := s.Clone(review).(*s.Object).Set("findings", []any{kept, snake})
	omitted := s.Clone(review).(*s.Object).Set("summary", "No finding is made about an illegal Snakeman crosspath.")
	codex := sequencedCodex(t, dir, plan, mechanics, first, omitted, review)

	output := filepath.Join(dir, "dart.json")
	flags := []string{"--provider", "codex", "--codex", codex, "--timeout", "20"}
	_, stderr, err := cli(t, append([]string{"author", request, "--profile", "default", "-o", output}, flags...)...)
	if err == nil || !strings.Contains(err.Error(), "model.snake-crosspath") {
		t.Fatalf("author: %v\n%s", err, stderr)
	}
	if _, err := os.Stat(output); err == nil {
		t.Error("a rejected review wrote its Result")
	}
	checked := filepath.Join(dir, "dart.checked.json")
	want := "Wrote the checked draft to " + checked + ". Retry the review with the same model options: mardwerk-unit review " + checked + " -o " + output
	if _, err := os.Stat(checked); err != nil || !strings.Contains(stderr, want) {
		t.Fatalf("no checked draft: %v\n%s", err, stderr)
	}
	// Call 5 answers the rerun with the recorded review.
	if _, stderr, err := cli(t, append([]string{"review", checked, "-o", output}, flags...)...); err != nil {
		t.Fatalf("review of the kept draft: %v\n%s", err, stderr)
	}
	if _, err := os.Stat(output); err != nil {
		t.Error("the rerun wrote no Result")
	}
}
