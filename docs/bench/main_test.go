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
