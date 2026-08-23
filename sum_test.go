package collection

import "testing"

func TestSum_Ints(t *testing.T) {
	c := []int{1, 2, 3, 4}
	out := Sum(c)

	if out != 10 {
		t.Fatalf("expected 10, got %v", out)
	}
}

func TestSum_Floats(t *testing.T) {
	c := []float64{1.5, 2.5, 3.0}
	out := Sum(c)

	if out != 7.0 {
		t.Fatalf("expected 7.0, got %v", out)
	}
}

func TestSum_Empty(t *testing.T) {
	c := []int{}
	out := Sum(c)

	if out != 0 {
		t.Fatalf("expected 0 for empty collection, got %v", out)
	}
}

func TestSum_DoesNotMutate(t *testing.T) {
	c := []int{5, 5}

	_ = Sum(c)

	if len(c) != 2 {
		t.Fatalf("Sum() mutated the collection length, got len=%d", len(c))
	}
}
