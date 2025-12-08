package collection

import (
	"math"
	"testing"
)

func TestMedian_OddCount(t *testing.T) {
	c := NewNumeric([]int{3, 1, 2})
	median, ok := c.Median()

	if !ok {
		t.Fatalf("expected ok=true")
	}

	if median != 2 {
		t.Fatalf("expected 2, got %v", median)
	}
}

func TestMedian_EvenCount(t *testing.T) {
	c := NewNumeric([]int{4, 1, 2, 3})
	median, ok := c.Median()

	if !ok {
		t.Fatalf("expected ok=true")
	}

	expected := (2 + 3) / 2.0
	if median != expected {
		t.Fatalf("expected %v, got %v", expected, median)
	}
}

func TestMedian_Floats(t *testing.T) {
	c := NewNumeric([]float64{1.5, 3.5, 2.5})
	median, ok := c.Median()

	if !ok {
		t.Fatalf("expected ok=true")
	}

	if math.Abs(median-2.5) > 1e-9 {
		t.Fatalf("expected 2.5, got %v", median)
	}
}

func TestMedian_Empty(t *testing.T) {
	c := NewNumeric([]int{})
	median, ok := c.Median()

	if ok {
		t.Fatalf("expected ok=false for empty collection")
	}

	if median != 0 {
		t.Fatalf("expected median=0 for empty collection, got %v", median)
	}
}

func TestMedian_DoesNotMutate(t *testing.T) {
	orig := []int{10, 1, 5}
	c := NewNumeric(orig)

	_, _ = c.Median()

	// Ensure original slice unchanged
	if orig[0] != 10 || orig[1] != 1 || orig[2] != 5 {
		t.Fatalf("Median() mutated original slice: %v", orig)
	}
}
