package collection

import (
	"math"
	"reflect"
	"testing"
)

func TestMin_Ints(t *testing.T) {
	c := New([]int{5, 3, 9, 1, 4})

	val, ok := Min(c)

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if val != 1 {
		t.Fatalf("expected 1, got %v", val)
	}
}

func TestMin_IntsWithNegatives(t *testing.T) {
	c := New([]int{-1, -50, 10, 0, -3})

	val, ok := Min(c)

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if val != -50 {
		t.Fatalf("expected -50, got %v", val)
	}
}

func TestMin_SingleElement(t *testing.T) {
	c := New([]int{42})

	val, ok := Min(c)

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if val != 42 {
		t.Fatalf("expected 42, got %v", val)
	}
}

func TestMin_Empty(t *testing.T) {
	c := New([]int{})

	val, ok := Min(c)

	if ok {
		t.Fatalf("expected ok=false, got true with value=%v", val)
	}

	if val != 0 {
		t.Fatalf("expected zero value, got %v", val)
	}
}

func TestMin_Floats(t *testing.T) {
	c := New([]float64{3.3, 1.1, 2.2})

	val, ok := Min(c)

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if val != 1.1 {
		t.Fatalf("expected 1.1, got %v", val)
	}
}

func TestMin_LargeNumbers(t *testing.T) {
	c := New([]int64{math.MaxInt64, math.MaxInt64 - 10, 0, -500})

	val, ok := Min(c)

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if val != -500 {
		t.Fatalf("expected -500, got %v", val)
	}
}

func TestMin_NoMutation(t *testing.T) {
	c := New([]int{3, 1, 2})
	orig := append([]int{}, c.items...) // copy original

	_, _ = Min(c)

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("original collection was mutated: %v", c.items)
	}
}
