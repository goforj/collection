package main

import "testing"

// TestRatioIsUncertain verifies that only consistent differences escape the equivalence label.
func TestRatioIsUncertain(t *testing.T) {
	tests := []struct {
		name   string
		ratios []float64
		want   bool
	}{
		{name: "equivalent sample", ratios: []float64{1.20, 1.02, 1.18}, want: true},
		{name: "conflicting directions", ratios: []float64{1.20, 0.80, 1.18}, want: true},
		{name: "consistently faster", ratios: []float64{1.20, 1.15, 1.18}},
		{name: "consistently slower", ratios: []float64{0.80, 0.85, 0.82}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := ratioIsUncertain(test.ratios); got != test.want {
				t.Fatalf("ratioIsUncertain(%v) = %v, want %v", test.ratios, got, test.want)
			}
		})
	}
}

// TestFormatUncertainTiming distinguishes equivalent medians from unstable differences.
func TestFormatUncertainTiming(t *testing.T) {
	if got := formatUncertainTiming(102, 100, equivalentEpsilon); got != "≈" {
		t.Fatalf("formatUncertainTiming(equivalent) = %q", got)
	}
	if got := formatUncertainTiming(118, 100, equivalentEpsilon); got != "inconclusive" {
		t.Fatalf("formatUncertainTiming(inconsistent) = %q", got)
	}
}

// TestFormatRatioUsesConservativeCategories avoids publishing build-sensitive multipliers.
func TestFormatRatioUsesConservativeCategories(t *testing.T) {
	if got := formatRatio(200, 100); got != "**faster**" {
		t.Fatalf("formatRatio(faster) = %q", got)
	}
	if got := formatRatio(100, 200); got != "slower" {
		t.Fatalf("formatRatio(slower) = %q", got)
	}
}

// TestScalarSummaryUsesDocumentedTolerance verifies the condensed scalar table's wider noise band.
func TestScalarSummaryUsesDocumentedTolerance(t *testing.T) {
	raw := formatBenchmarkRatio("All", benchBorrow, 112, 100, false)
	if raw != "**faster**" {
		t.Fatalf("raw timing classification = %q", raw)
	}
	summary := formatBenchmarkSpeed("All", benchBorrow, 112, 100, false, false, true)
	if summary != "≈" {
		t.Fatalf("scalar summary classification = %q", summary)
	}
}
