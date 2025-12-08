package collection

import (
	"math"
	"testing"
)

func TestAvg_Ints(t *testing.T) {
	c := NewNumeric([]int{2, 4, 6})
	out := c.Avg()

	if out != 4 {
		t.Fatalf("expected 4, got %v", out)
	}
}

func TestAvg_Floats(t *testing.T) {
	c := NewNumeric([]float64{1.5, 2.5, 3.0})
	out := c.Avg()

	expected := (1.5 + 2.5 + 3.0) / 3
	if math.Abs(out-expected) > 1e-9 {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestAvg_Empty(t *testing.T) {
	c := NewNumeric([]int{})
	out := c.Avg()

	if out != 0 {
		t.Fatalf("expected 0 for empty collection, got %v", out)
	}
}

func TestAvg_FluentUsage(t *testing.T) {
	// Ensures Avg behaves correctly when chained after other methods
	c := NewNumeric([]int{1, 2, 3})
	avg := c.Avg() // numeric terminal op

	if avg != 2 {
		t.Fatalf("expected 2, got %v", avg)
	}
}

func TestAvg_LargeNumbers(t *testing.T) {
	c := NewNumeric([]int64{1_000_000_000, 2_000_000_000})
	out := c.Avg()

	if out != 1_500_000_000 {
		t.Fatalf("expected 1500000000, got %v", out)
	}
}
