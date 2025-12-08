package collection

import "testing"

func TestSum_Ints(t *testing.T) {
	c := NewNumeric([]int{1, 2, 3, 4})
	out := c.Sum()

	if out != 10 {
		t.Fatalf("expected 10, got %v", out)
	}
}

func TestSum_Floats(t *testing.T) {
	c := NewNumeric([]float64{1.5, 2.5, 3.0})
	out := c.Sum()

	if out != 7.0 {
		t.Fatalf("expected 7.0, got %v", out)
	}
}

func TestSum_Empty(t *testing.T) {
	c := NewNumeric([]int{})
	out := c.Sum()

	if out != 0 {
		t.Fatalf("expected 0 for empty collection, got %v", out)
	}
}

func TestSum_DoesNotMutate(t *testing.T) {
	c := NewNumeric([]int{5, 5})

	_ = c.Sum()

	if len(c.items) != 2 {
		t.Fatalf("Sum() mutated the collection length, got len=%d", len(c.items))
	}
}
