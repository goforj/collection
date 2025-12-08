package collection

import (
	"math"
	"reflect"
	"testing"
)

func TestMax_Ints(t *testing.T) {
	c := New([]int{1, 5, 3, 9, 2})

	val, ok := Max(c)

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if val != 9 {
		t.Fatalf("expected 9, got %v", val)
	}
}

func TestMax_IntsWithNegatives(t *testing.T) {
	c := New([]int{-10, -3, -50, -2, -99})

	val, ok := Max(c)

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if val != -2 {
		t.Fatalf("expected -2, got %v", val)
	}
}

func TestMax_SingleElement(t *testing.T) {
	c := New([]int{42})

	val, ok := Max(c)

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if val != 42 {
		t.Fatalf("expected 42, got %v", val)
	}
}

func TestMax_Empty(t *testing.T) {
	c := New([]int{})

	val, ok := Max(c)

	if ok {
		t.Fatalf("expected ok=false for empty collection, got true with %v", val)
	}

	if val != 0 {
		t.Fatalf("expected zero value, got %v", val)
	}
}

func TestMax_Floats(t *testing.T) {
	c := New([]float64{1.1, 5.5, 2.2})

	val, ok := Max(c)

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if val != 5.5 {
		t.Fatalf("expected 5.5, got %v", val)
	}
}

func TestMax_LargeNumbers(t *testing.T) {
	c := New([]int64{10, math.MaxInt64 - 5, math.MaxInt64})

	val, ok := Max(c)

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if val != math.MaxInt64 {
		t.Fatalf("expected %v, got %v", math.MaxInt64, val)
	}
}

func TestMax_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})
	orig := append([]int{}, c.items...) // copy original

	_, _ = Max(c)

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("original collection was mutated: %v", c.items)
	}
}
