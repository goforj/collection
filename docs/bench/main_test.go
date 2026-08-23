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
	if got := formatRatio(200, 100); got != "**2.0x faster**" {
		t.Fatalf("formatRatio(faster) = %q", got)
	}
	if got := formatRatio(100, 200); got != "2.0x slower" {
		t.Fatalf("formatRatio(slower) = %q", got)
	}
	if got := formatRatio(0, 100); got != "∞x slower" {
		t.Fatalf("formatRatio(zero baseline) = %q", got)
	}
	if got := formatRatio(100, 0); got != "**∞x faster**" {
		t.Fatalf("formatRatio(zero collection) = %q", got)
	}
	if got := formatSpeed(100, 0, false, false); got != "∞x faster" {
		t.Fatalf("formatSpeed(zero collection) = %q", got)
	}
	if got := formatSpeed(100, 0, true, false); got != "**∞x faster**" {
		t.Fatalf("formatSpeed(zero collection, bold) = %q", got)
	}
}

// TestFormatRatioDoesNotHideSubFloorDifferences avoids claiming equivalence below the reporting floor.
func TestFormatRatioDoesNotHideSubFloorDifferences(t *testing.T) {
	if got := formatRatio(49, 1); got != "below floor" {
		t.Fatalf("formatRatio(sub-floor difference) = %q", got)
	}
}

// TestScalarSummaryUsesDocumentedTolerance verifies the condensed scalar table's wider noise band.
func TestScalarSummaryUsesDocumentedTolerance(t *testing.T) {
	raw := formatBenchmarkRatio("All", benchBorrow, 112, 100, false)
	if raw != "**1.1x faster**" {
		t.Fatalf("raw timing classification = %q", raw)
	}
	summary := formatBenchmarkSpeed("All", benchBorrow, 112, 100, false, false, true)
	if summary != "≈" {
		t.Fatalf("scalar summary classification = %q", summary)
	}
}

// TestFormatRatioBytesUsesCorrectDirection reports collection memory relative to lo.
func TestFormatRatioBytesUsesCorrectDirection(t *testing.T) {
	if got := formatRatioBytes(100, 200); got != "2.00x more" {
		t.Fatalf("formatRatioBytes(more) = %q", got)
	}
	if got := formatRatioBytes(200, 100); got != "**2.00x less**" {
		t.Fatalf("formatRatioBytes(less) = %q", got)
	}
	if got := formatRatioBytes(100, 105); got != "1.05x more" {
		t.Fatalf("formatRatioBytes(nearby values) = %q", got)
	}
}
