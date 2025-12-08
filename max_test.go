package collection

import "testing"

func TestMax_Ints(t *testing.T) {
	c := NewNumeric([]int{3, 1, 2})
	val, ok := c.Max()

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}

	if val != 3 {
		t.Fatalf("expected 3, got %v", val)
	}
}

func TestMax_Floats(t *testing.T) {
	c := NewNumeric([]float64{3.5, 1.1, 2.2})
	val, ok := c.Max()

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}

	if val != 3.5 {
		t.Fatalf("expected 3.5, got %v", val)
	}
}

func TestMax_Empty(t *testing.T) {
	c := NewNumeric([]uint{})
	val, ok := c.Max()

	if ok {
		t.Fatalf("expected ok=false for empty collection")
	}

	if val != 0 {
		t.Fatalf("expected zero value for empty, got %v", val)
	}
}

func TestMax_SingleValue(t *testing.T) {
	c := NewNumeric([]int{42})
	val, ok := c.Max()

	if !ok {
		t.Fatalf("expected ok=true")
	}

	if val != 42 {
		t.Fatalf("expected 42, got %v", val)
	}
}

func TestMax_BranchUpdate(t *testing.T) {
	// Here the first item is NOT the max.
	// This forces the branch (v > max) to execute.
	c := NewNumeric([]int{1, 9, 3})
	max, ok := c.Max()

	if !ok {
		t.Fatalf("expected ok=true")
	}

	if max != 9 {
		t.Fatalf("expected max=9, got %v", max)
	}
}
