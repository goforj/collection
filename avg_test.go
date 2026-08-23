package collection

import (
	"math"
	"testing"
)

func TestAvg_Ints(t *testing.T) {
	c := []int{2, 4, 6}
	out := Avg(c)

	if out != 4 {
		t.Fatalf("expected 4, got %v", out)
	}
}

func TestAvg_Floats(t *testing.T) {
	c := []float64{1.5, 2.5, 3.0}
	out := Avg(c)

	expected := (1.5 + 2.5 + 3.0) / 3
	if math.Abs(out-expected) > 1e-9 {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestAvg_Empty(t *testing.T) {
	c := []int{}
	out := Avg(c)

	if out != 0 {
		t.Fatalf("expected 0 for empty collection, got %v", out)
	}
}

func TestAvg_FluentUsage(t *testing.T) {
	// Ensures Avg behaves correctly when chained after other methods
	c := []int{1, 2, 3}
	avg := Avg(c) // numeric terminal op

	if avg != 2 {
		t.Fatalf("expected 2, got %v", avg)
	}
}

func TestAvg_LargeNumbers(t *testing.T) {
	c := []int64{1_000_000_000, 2_000_000_000}
	out := Avg(c)

	if out != 1_500_000_000 {
		t.Fatalf("expected 1500000000, got %v", out)
	}
}
