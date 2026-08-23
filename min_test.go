package collection

import "testing"

func TestMin_Ints(t *testing.T) {
	c := []int{3, 1, 2}
	min, ok := Min(c)

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}

	if min != 1 {
		t.Fatalf("expected 1, got %v", min)
	}
}

func TestMin_Floats(t *testing.T) {
	c := []float64{3.5, 1.1, 2.2}
	min, ok := Min(c)

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}

	if min != 1.1 {
		t.Fatalf("expected 1.1, got %v", min)
	}
}

func TestMin_Empty(t *testing.T) {
	c := []uint{}
	min, ok := Min(c)

	if ok {
		t.Fatalf("expected ok=false on empty collection")
	}

	if min != 0 {
		t.Fatalf("expected zero value for empty collection, got %v", min)
	}
}

func TestMin_SingleValue(t *testing.T) {
	c := []int{42}
	min, ok := Min(c)

	if !ok {
		t.Fatalf("expected ok=true")
	}

	if min != 42 {
		t.Fatalf("expected 42, got %v", min)
	}
}
