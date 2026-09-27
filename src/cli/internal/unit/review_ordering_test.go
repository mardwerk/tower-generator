package unit

import (
	"math"
	"testing"

	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
)

func TestCapstoneOrderingDoesNotInvertTheCopyComparison(t *testing.T) {
	comparison := s.NewObject().
		Set("path", "path1").
		Set("tier5", s.NewObject().Set("metrics", s.NewObject().
			Set("group damage rate upper bound", 511.58))).
		Set("sameBudgetTier4Copies", s.NewObject().
			Set("additiveThroughputUpperBounds", s.NewObject().
				Set("group damage rate upper bound", 454.74)))
	rows := capstoneOrdering([]any{comparison}, nil)
	row := rows[0].(*s.Object)
	group, _ := row.Get("group damage rate upper bound")
	ordering, _ := group.(*s.Object).Get("ordering")
	ratio, _ := group.(*s.Object).Get("tier5ToCopiesRatio")
	if ordering != "higher" || math.Abs(ratio.(float64)-511.58/454.74) > 1e-9 {
		t.Fatalf("comparison ordering %s", s.Stringify(row))
	}
}
