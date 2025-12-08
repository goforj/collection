package collection

import (
	"testing"
)

func TestLastWhere_WithPredicate(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	v, ok := c.LastWhere(func(v int, i int) bool {
		return v < 3
	})

	if !ok {
		t.Fatalf("expected ok == true, got false")
	}
	if v != 2 {
		t.Fatalf("expected value 2, got %v", v)
	}
}

func TestLastWhere_NoPredicate(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	v, ok := c.LastWhere(nil)

	if !ok {
		t.Fatalf("expected ok == true, got false")
	}
	if v != 4 {
		t.Fatalf("expected last element 4, got %v", v)
	}
}

func TestLastWhere_PredicateNoMatch(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	v, ok := c.LastWhere(func(v int, i int) bool {
		return v > 10
	})

	if ok {
		t.Fatalf("expected ok == false when no match, got true with value %v", v)
	}
}

func TestLastWhere_EmptyCollection(t *testing.T) {
	c := New([]int{})

	v, ok := c.LastWhere(nil)

	if ok {
		t.Fatalf("expected ok == false for empty collection, got true with value %v", v)
	}
	if v != 0 { // zero-value for int
		t.Fatalf("expected zero-value, got %v", v)
	}
}
